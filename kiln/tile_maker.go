package kiln

import (
	"io/ioutil"
	"path/filepath"

	"github.com/pivotal-cf/kiln/builder"
	"github.com/pivotal-cf/kiln/commands"
	yaml "gopkg.in/yaml.v2"
)

//go:generate counterfeiter -o ./fakes/tile_writer.go --fake-name TileWriter . tileWriter
type tileWriter interface {
	Write(generatedMetadataContents []byte, config commands.BakeConfig) error
}

//go:generate counterfeiter -o ./fakes/metadata_builder.go --fake-name MetadataBuilder . metadataBuilder
type metadataBuilder interface {
	Build(releaseTarballs []string, pathToStemcell, pathToMetadata, name, version, pathToTile string) (builder.GeneratedMetadata, error)
}

type logger interface {
	Println(v ...interface{})
}

type TileMaker struct {
	metadataBuilder metadataBuilder
	tileWriter      tileWriter
	logger          logger
}

func NewTileMaker(metadataBuilder metadataBuilder, tileWriter tileWriter, logger logger) TileMaker {
	return TileMaker{
		metadataBuilder: metadataBuilder,
		tileWriter:      tileWriter,
		logger:          logger,
	}
}

func (t TileMaker) Make(config commands.BakeConfig) error {
	var releaseTarballs []string
	for _, releasesDirectory := range config.ReleaseDirectories {
		files, err := ioutil.ReadDir(releasesDirectory)
		if err != nil {
			return err
		}

		for _, file := range files {
			releaseTarballs = append(releaseTarballs, filepath.Join(releasesDirectory, file.Name()))
		}
	}

	generatedMetadata, err := t.metadataBuilder.Build(
		releaseTarballs,
		config.StemcellTarball,
		config.Metadata,
		config.ProductName,
		config.Version,
		config.OutputFile,
	)
	if err != nil {
		return err
	}

	t.logger.Println("Marshaling metadata file...")
	generatedMetadataYAML, err := yaml.Marshal(generatedMetadata)
	if err != nil {
		return err
	}

	err = t.tileWriter.Write(generatedMetadataYAML, config)
	if err != nil {
		return err
	}

	return nil

}
