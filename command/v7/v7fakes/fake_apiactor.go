// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeAPIActor struct {
	ClearTargetStub        func()
	clearTargetMutex       sync.RWMutex
	clearTargetArgsForCall []struct {
	}
	SetTargetStub        func(v7action.TargetSettings) (v7action.Warnings, error)
	setTargetMutex       sync.RWMutex
	setTargetArgsForCall []struct {
		arg1 v7action.TargetSettings
	}
	setTargetReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	setTargetReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAPIActor) ClearTarget() {
	fake.clearTargetMutex.Lock()
	fake.clearTargetArgsForCall = append(fake.clearTargetArgsForCall, struct {
	}{})
	fake.recordInvocation("ClearTarget", []interface{}{})
	fake.clearTargetMutex.Unlock()
	if fake.ClearTargetStub != nil {
		fake.ClearTargetStub()
	}
}

func (fake *FakeAPIActor) ClearTargetCallCount() int {
	fake.clearTargetMutex.RLock()
	defer fake.clearTargetMutex.RUnlock()
	return len(fake.clearTargetArgsForCall)
}

func (fake *FakeAPIActor) ClearTargetCalls(stub func()) {
	fake.clearTargetMutex.Lock()
	defer fake.clearTargetMutex.Unlock()
	fake.ClearTargetStub = stub
}

func (fake *FakeAPIActor) SetTarget(arg1 v7action.TargetSettings) (v7action.Warnings, error) {
	fake.setTargetMutex.Lock()
	ret, specificReturn := fake.setTargetReturnsOnCall[len(fake.setTargetArgsForCall)]
	fake.setTargetArgsForCall = append(fake.setTargetArgsForCall, struct {
		arg1 v7action.TargetSettings
	}{arg1})
	fake.recordInvocation("SetTarget", []interface{}{arg1})
	fake.setTargetMutex.Unlock()
	if fake.SetTargetStub != nil {
		return fake.SetTargetStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.setTargetReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAPIActor) SetTargetCallCount() int {
	fake.setTargetMutex.RLock()
	defer fake.setTargetMutex.RUnlock()
	return len(fake.setTargetArgsForCall)
}

func (fake *FakeAPIActor) SetTargetCalls(stub func(v7action.TargetSettings) (v7action.Warnings, error)) {
	fake.setTargetMutex.Lock()
	defer fake.setTargetMutex.Unlock()
	fake.SetTargetStub = stub
}

func (fake *FakeAPIActor) SetTargetArgsForCall(i int) v7action.TargetSettings {
	fake.setTargetMutex.RLock()
	defer fake.setTargetMutex.RUnlock()
	argsForCall := fake.setTargetArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAPIActor) SetTargetReturns(result1 v7action.Warnings, result2 error) {
	fake.setTargetMutex.Lock()
	defer fake.setTargetMutex.Unlock()
	fake.SetTargetStub = nil
	fake.setTargetReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeAPIActor) SetTargetReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.setTargetMutex.Lock()
	defer fake.setTargetMutex.Unlock()
	fake.SetTargetStub = nil
	if fake.setTargetReturnsOnCall == nil {
		fake.setTargetReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.setTargetReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeAPIActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.clearTargetMutex.RLock()
	defer fake.clearTargetMutex.RUnlock()
	fake.setTargetMutex.RLock()
	defer fake.setTargetMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAPIActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.APIActor = new(FakeAPIActor)
