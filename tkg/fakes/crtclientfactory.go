// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/vmware-tanzu/tanzu-framework/tkg/clusterclient"
)

type CrtClientFactory struct {
	NewClientStub        func(*rest.Config, client.Options) (client.Client, error)
	newClientMutex       sync.RWMutex
	newClientArgsForCall []struct {
		arg1 *rest.Config
		arg2 client.Options
	}
	newClientReturns struct {
		result1 client.Client
		result2 error
	}
	newClientReturnsOnCall map[int]struct {
		result1 client.Client
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CrtClientFactory) NewClient(arg1 *rest.Config, arg2 client.Options) (client.Client, error) {
	fake.newClientMutex.Lock()
	ret, specificReturn := fake.newClientReturnsOnCall[len(fake.newClientArgsForCall)]
	fake.newClientArgsForCall = append(fake.newClientArgsForCall, struct {
		arg1 *rest.Config
		arg2 client.Options
	}{arg1, arg2})
	stub := fake.NewClientStub
	fakeReturns := fake.newClientReturns
	fake.recordInvocation("NewClient", []interface{}{arg1, arg2})
	fake.newClientMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CrtClientFactory) NewClientCallCount() int {
	fake.newClientMutex.RLock()
	defer fake.newClientMutex.RUnlock()
	return len(fake.newClientArgsForCall)
}

func (fake *CrtClientFactory) NewClientCalls(stub func(*rest.Config, client.Options) (client.Client, error)) {
	fake.newClientMutex.Lock()
	defer fake.newClientMutex.Unlock()
	fake.NewClientStub = stub
}

func (fake *CrtClientFactory) NewClientArgsForCall(i int) (*rest.Config, client.Options) {
	fake.newClientMutex.RLock()
	defer fake.newClientMutex.RUnlock()
	argsForCall := fake.newClientArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *CrtClientFactory) NewClientReturns(result1 client.Client, result2 error) {
	fake.newClientMutex.Lock()
	defer fake.newClientMutex.Unlock()
	fake.NewClientStub = nil
	fake.newClientReturns = struct {
		result1 client.Client
		result2 error
	}{result1, result2}
}

func (fake *CrtClientFactory) NewClientReturnsOnCall(i int, result1 client.Client, result2 error) {
	fake.newClientMutex.Lock()
	defer fake.newClientMutex.Unlock()
	fake.NewClientStub = nil
	if fake.newClientReturnsOnCall == nil {
		fake.newClientReturnsOnCall = make(map[int]struct {
			result1 client.Client
			result2 error
		})
	}
	fake.newClientReturnsOnCall[i] = struct {
		result1 client.Client
		result2 error
	}{result1, result2}
}

func (fake *CrtClientFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newClientMutex.RLock()
	defer fake.newClientMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CrtClientFactory) recordInvocation(key string, args []interface{}) {
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

var _ clusterclient.CrtClientFactory = new(CrtClientFactory)
