// Code generated by counterfeiter. DO NOT EDIT.
package resourcefakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/concourse/atc/creds"
	"github.com/concourse/atc/resource"
	"github.com/concourse/atc/worker"
)

type FakeFetchSourceProviderFactory struct {
	NewFetchSourceProviderStub        func(logger lager.Logger, session resource.Session, metadata resource.Metadata, tags atc.Tags, teamID int, resourceTypes creds.VersionedResourceTypes, resourceInstance resource.ResourceInstance, imageFetchingDelegate worker.ImageFetchingDelegate) resource.FetchSourceProvider
	newFetchSourceProviderMutex       sync.RWMutex
	newFetchSourceProviderArgsForCall []struct {
		logger                lager.Logger
		session               resource.Session
		metadata              resource.Metadata
		tags                  atc.Tags
		teamID                int
		resourceTypes         creds.VersionedResourceTypes
		resourceInstance      resource.ResourceInstance
		imageFetchingDelegate worker.ImageFetchingDelegate
	}
	newFetchSourceProviderReturns struct {
		result1 resource.FetchSourceProvider
	}
	newFetchSourceProviderReturnsOnCall map[int]struct {
		result1 resource.FetchSourceProvider
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFetchSourceProviderFactory) NewFetchSourceProvider(logger lager.Logger, session resource.Session, metadata resource.Metadata, tags atc.Tags, teamID int, resourceTypes creds.VersionedResourceTypes, resourceInstance resource.ResourceInstance, imageFetchingDelegate worker.ImageFetchingDelegate) resource.FetchSourceProvider {
	fake.newFetchSourceProviderMutex.Lock()
	ret, specificReturn := fake.newFetchSourceProviderReturnsOnCall[len(fake.newFetchSourceProviderArgsForCall)]
	fake.newFetchSourceProviderArgsForCall = append(fake.newFetchSourceProviderArgsForCall, struct {
		logger                lager.Logger
		session               resource.Session
		metadata              resource.Metadata
		tags                  atc.Tags
		teamID                int
		resourceTypes         creds.VersionedResourceTypes
		resourceInstance      resource.ResourceInstance
		imageFetchingDelegate worker.ImageFetchingDelegate
	}{logger, session, metadata, tags, teamID, resourceTypes, resourceInstance, imageFetchingDelegate})
	fake.recordInvocation("NewFetchSourceProvider", []interface{}{logger, session, metadata, tags, teamID, resourceTypes, resourceInstance, imageFetchingDelegate})
	fake.newFetchSourceProviderMutex.Unlock()
	if fake.NewFetchSourceProviderStub != nil {
		return fake.NewFetchSourceProviderStub(logger, session, metadata, tags, teamID, resourceTypes, resourceInstance, imageFetchingDelegate)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.newFetchSourceProviderReturns.result1
}

func (fake *FakeFetchSourceProviderFactory) NewFetchSourceProviderCallCount() int {
	fake.newFetchSourceProviderMutex.RLock()
	defer fake.newFetchSourceProviderMutex.RUnlock()
	return len(fake.newFetchSourceProviderArgsForCall)
}

func (fake *FakeFetchSourceProviderFactory) NewFetchSourceProviderArgsForCall(i int) (lager.Logger, resource.Session, resource.Metadata, atc.Tags, int, creds.VersionedResourceTypes, resource.ResourceInstance, worker.ImageFetchingDelegate) {
	fake.newFetchSourceProviderMutex.RLock()
	defer fake.newFetchSourceProviderMutex.RUnlock()
	return fake.newFetchSourceProviderArgsForCall[i].logger, fake.newFetchSourceProviderArgsForCall[i].session, fake.newFetchSourceProviderArgsForCall[i].metadata, fake.newFetchSourceProviderArgsForCall[i].tags, fake.newFetchSourceProviderArgsForCall[i].teamID, fake.newFetchSourceProviderArgsForCall[i].resourceTypes, fake.newFetchSourceProviderArgsForCall[i].resourceInstance, fake.newFetchSourceProviderArgsForCall[i].imageFetchingDelegate
}

func (fake *FakeFetchSourceProviderFactory) NewFetchSourceProviderReturns(result1 resource.FetchSourceProvider) {
	fake.NewFetchSourceProviderStub = nil
	fake.newFetchSourceProviderReturns = struct {
		result1 resource.FetchSourceProvider
	}{result1}
}

func (fake *FakeFetchSourceProviderFactory) NewFetchSourceProviderReturnsOnCall(i int, result1 resource.FetchSourceProvider) {
	fake.NewFetchSourceProviderStub = nil
	if fake.newFetchSourceProviderReturnsOnCall == nil {
		fake.newFetchSourceProviderReturnsOnCall = make(map[int]struct {
			result1 resource.FetchSourceProvider
		})
	}
	fake.newFetchSourceProviderReturnsOnCall[i] = struct {
		result1 resource.FetchSourceProvider
	}{result1}
}

func (fake *FakeFetchSourceProviderFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newFetchSourceProviderMutex.RLock()
	defer fake.newFetchSourceProviderMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFetchSourceProviderFactory) recordInvocation(key string, args []interface{}) {
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

var _ resource.FetchSourceProviderFactory = new(FakeFetchSourceProviderFactory)
