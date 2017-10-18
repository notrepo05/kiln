// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/pivotal-cf/kiln/commands"
)

type TileWriter struct {
	WriteStub        func(generatedMetadataContents []byte, config commands.BakeConfig) error
	writeMutex       sync.RWMutex
	writeArgsForCall []struct {
		generatedMetadataContents []byte
		config                    commands.BakeConfig
	}
	writeReturns struct {
		result1 error
	}
	writeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *TileWriter) Write(generatedMetadataContents []byte, config commands.BakeConfig) error {
	var generatedMetadataContentsCopy []byte
	if generatedMetadataContents != nil {
		generatedMetadataContentsCopy = make([]byte, len(generatedMetadataContents))
		copy(generatedMetadataContentsCopy, generatedMetadataContents)
	}
	fake.writeMutex.Lock()
	ret, specificReturn := fake.writeReturnsOnCall[len(fake.writeArgsForCall)]
	fake.writeArgsForCall = append(fake.writeArgsForCall, struct {
		generatedMetadataContents []byte
		config                    commands.BakeConfig
	}{generatedMetadataContentsCopy, config})
	fake.recordInvocation("Write", []interface{}{generatedMetadataContentsCopy, config})
	fake.writeMutex.Unlock()
	if fake.WriteStub != nil {
		return fake.WriteStub(generatedMetadataContents, config)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.writeReturns.result1
}

func (fake *TileWriter) WriteCallCount() int {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return len(fake.writeArgsForCall)
}

func (fake *TileWriter) WriteArgsForCall(i int) ([]byte, commands.BakeConfig) {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return fake.writeArgsForCall[i].generatedMetadataContents, fake.writeArgsForCall[i].config
}

func (fake *TileWriter) WriteReturns(result1 error) {
	fake.WriteStub = nil
	fake.writeReturns = struct {
		result1 error
	}{result1}
}

func (fake *TileWriter) WriteReturnsOnCall(i int, result1 error) {
	fake.WriteStub = nil
	if fake.writeReturnsOnCall == nil {
		fake.writeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *TileWriter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return fake.invocations
}

func (fake *TileWriter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
