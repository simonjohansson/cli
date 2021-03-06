// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeCreateSpaceQuotaActor struct {
	CreateSpaceQuotaStub        func(string, string, v7action.QuotaLimits) (v7action.Warnings, error)
	createSpaceQuotaMutex       sync.RWMutex
	createSpaceQuotaArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 v7action.QuotaLimits
	}
	createSpaceQuotaReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	createSpaceQuotaReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCreateSpaceQuotaActor) CreateSpaceQuota(arg1 string, arg2 string, arg3 v7action.QuotaLimits) (v7action.Warnings, error) {
	fake.createSpaceQuotaMutex.Lock()
	ret, specificReturn := fake.createSpaceQuotaReturnsOnCall[len(fake.createSpaceQuotaArgsForCall)]
	fake.createSpaceQuotaArgsForCall = append(fake.createSpaceQuotaArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 v7action.QuotaLimits
	}{arg1, arg2, arg3})
	fake.recordInvocation("CreateSpaceQuota", []interface{}{arg1, arg2, arg3})
	fake.createSpaceQuotaMutex.Unlock()
	if fake.CreateSpaceQuotaStub != nil {
		return fake.CreateSpaceQuotaStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createSpaceQuotaReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCreateSpaceQuotaActor) CreateSpaceQuotaCallCount() int {
	fake.createSpaceQuotaMutex.RLock()
	defer fake.createSpaceQuotaMutex.RUnlock()
	return len(fake.createSpaceQuotaArgsForCall)
}

func (fake *FakeCreateSpaceQuotaActor) CreateSpaceQuotaCalls(stub func(string, string, v7action.QuotaLimits) (v7action.Warnings, error)) {
	fake.createSpaceQuotaMutex.Lock()
	defer fake.createSpaceQuotaMutex.Unlock()
	fake.CreateSpaceQuotaStub = stub
}

func (fake *FakeCreateSpaceQuotaActor) CreateSpaceQuotaArgsForCall(i int) (string, string, v7action.QuotaLimits) {
	fake.createSpaceQuotaMutex.RLock()
	defer fake.createSpaceQuotaMutex.RUnlock()
	argsForCall := fake.createSpaceQuotaArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCreateSpaceQuotaActor) CreateSpaceQuotaReturns(result1 v7action.Warnings, result2 error) {
	fake.createSpaceQuotaMutex.Lock()
	defer fake.createSpaceQuotaMutex.Unlock()
	fake.CreateSpaceQuotaStub = nil
	fake.createSpaceQuotaReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeCreateSpaceQuotaActor) CreateSpaceQuotaReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.createSpaceQuotaMutex.Lock()
	defer fake.createSpaceQuotaMutex.Unlock()
	fake.CreateSpaceQuotaStub = nil
	if fake.createSpaceQuotaReturnsOnCall == nil {
		fake.createSpaceQuotaReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.createSpaceQuotaReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeCreateSpaceQuotaActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createSpaceQuotaMutex.RLock()
	defer fake.createSpaceQuotaMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCreateSpaceQuotaActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.CreateSpaceQuotaActor = new(FakeCreateSpaceQuotaActor)
