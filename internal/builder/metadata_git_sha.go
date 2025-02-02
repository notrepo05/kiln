package builder

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const dirtyStateSHAValue = "DEVELOPMENT"

func GitMetadataSHA(repositoryDirectory string, isDev bool) (string, error) {
	if err := ensureGitExecutableIsFound(); err != nil {
		return "", err
	}
	gitStatus := exec.Command("git", "status", "--porcelain")
	gitStatus.Dir = repositoryDirectory
	err := gitStatus.Run()
	if err != nil {
		if gitStatus.ProcessState.ExitCode() == 1 && isDev {
			_, _ = fmt.Fprintf(os.Stderr, "WARNING: git working directory has un-commited changes: the variable %q has has development only value %q", MetadataGitSHAVariable, dirtyStateSHAValue)
			return dirtyStateSHAValue, nil
		}
		return "", fmt.Errorf("failed to run `%s %s`: %w", gitStatus.Path, strings.Join(gitStatus.Args, " "), err)
	}
	return gitHeadRevision(repositoryDirectory)
}

func gitHeadRevision(repositoryDirectory string) (string, error) {
	var out bytes.Buffer
	gitRevParseHead := exec.Command("git", "rev-parse", "HEAD")
	gitRevParseHead.Dir = repositoryDirectory
	gitRevParseHead.Stdout = &out
	err := gitRevParseHead.Run()
	if err != nil {
		return "", fmt.Errorf("failed to get HEAD revision hash: %w", err)
	}
	return strings.TrimSpace(out.String()), nil
}

func ensureGitExecutableIsFound() error {
	if _, err := exec.LookPath("git"); err != nil {
		return fmt.Errorf("could not calculate %q: %w", MetadataGitSHAVariable, err)
	}
	return nil
}
