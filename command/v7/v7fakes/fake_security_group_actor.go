// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeSecurityGroupActor struct {
	GetSecurityGroupStub        func(string) (v7action.SecurityGroupSummary, v7action.Warnings, error)
	getSecurityGroupMutex       sync.RWMutex
	getSecurityGroupArgsForCall []struct {
		arg1 string
	}
	getSecurityGroupReturns struct {
		result1 v7action.SecurityGroupSummary
		result2 v7action.Warnings
		result3 error
	}
	getSecurityGroupReturnsOnCall map[int]struct {
		result1 v7action.SecurityGroupSummary
		result2 v7action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSecurityGroupActor) GetSecurityGroup(arg1 string) (v7action.SecurityGroupSummary, v7action.Warnings, error) {
	fake.getSecurityGroupMutex.Lock()
	ret, specificReturn := fake.getSecurityGroupReturnsOnCall[len(fake.getSecurityGroupArgsForCall)]
	fake.getSecurityGroupArgsForCall = append(fake.getSecurityGroupArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetSecurityGroup", []interface{}{arg1})
	fake.getSecurityGroupMutex.Unlock()
	if fake.GetSecurityGroupStub != nil {
		return fake.GetSecurityGroupStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getSecurityGroupReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeSecurityGroupActor) GetSecurityGroupCallCount() int {
	fake.getSecurityGroupMutex.RLock()
	defer fake.getSecurityGroupMutex.RUnlock()
	return len(fake.getSecurityGroupArgsForCall)
}

func (fake *FakeSecurityGroupActor) GetSecurityGroupCalls(stub func(string) (v7action.SecurityGroupSummary, v7action.Warnings, error)) {
	fake.getSecurityGroupMutex.Lock()
	defer fake.getSecurityGroupMutex.Unlock()
	fake.GetSecurityGroupStub = stub
}

func (fake *FakeSecurityGroupActor) GetSecurityGroupArgsForCall(i int) string {
	fake.getSecurityGroupMutex.RLock()
	defer fake.getSecurityGroupMutex.RUnlock()
	argsForCall := fake.getSecurityGroupArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSecurityGroupActor) GetSecurityGroupReturns(result1 v7action.SecurityGroupSummary, result2 v7action.Warnings, result3 error) {
	fake.getSecurityGroupMutex.Lock()
	defer fake.getSecurityGroupMutex.Unlock()
	fake.GetSecurityGroupStub = nil
	fake.getSecurityGroupReturns = struct {
		result1 v7action.SecurityGroupSummary
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSecurityGroupActor) GetSecurityGroupReturnsOnCall(i int, result1 v7action.SecurityGroupSummary, result2 v7action.Warnings, result3 error) {
	fake.getSecurityGroupMutex.Lock()
	defer fake.getSecurityGroupMutex.Unlock()
	fake.GetSecurityGroupStub = nil
	if fake.getSecurityGroupReturnsOnCall == nil {
		fake.getSecurityGroupReturnsOnCall = make(map[int]struct {
			result1 v7action.SecurityGroupSummary
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getSecurityGroupReturnsOnCall[i] = struct {
		result1 v7action.SecurityGroupSummary
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSecurityGroupActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getSecurityGroupMutex.RLock()
	defer fake.getSecurityGroupMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSecurityGroupActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.SecurityGroupActor = new(FakeSecurityGroupActor)
