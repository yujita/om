// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	sync "sync"

	stow "github.com/graymeta/stow"
	commands "github.com/pivotal-cf/om/commands"
)

type Stower struct {
	DialStub        func(string, commands.Config) (stow.Location, error)
	dialMutex       sync.RWMutex
	dialArgsForCall []struct {
		arg1 string
		arg2 commands.Config
	}
	dialReturns struct {
		result1 stow.Location
		result2 error
	}
	dialReturnsOnCall map[int]struct {
		result1 stow.Location
		result2 error
	}
	WalkStub        func(stow.Container, string, int, stow.WalkFunc) error
	walkMutex       sync.RWMutex
	walkArgsForCall []struct {
		arg1 stow.Container
		arg2 string
		arg3 int
		arg4 stow.WalkFunc
	}
	walkReturns struct {
		result1 error
	}
	walkReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Stower) Dial(arg1 string, arg2 commands.Config) (stow.Location, error) {
	fake.dialMutex.Lock()
	ret, specificReturn := fake.dialReturnsOnCall[len(fake.dialArgsForCall)]
	fake.dialArgsForCall = append(fake.dialArgsForCall, struct {
		arg1 string
		arg2 commands.Config
	}{arg1, arg2})
	fake.recordInvocation("Dial", []interface{}{arg1, arg2})
	fake.dialMutex.Unlock()
	if fake.DialStub != nil {
		return fake.DialStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.dialReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Stower) DialCallCount() int {
	fake.dialMutex.RLock()
	defer fake.dialMutex.RUnlock()
	return len(fake.dialArgsForCall)
}

func (fake *Stower) DialCalls(stub func(string, commands.Config) (stow.Location, error)) {
	fake.dialMutex.Lock()
	defer fake.dialMutex.Unlock()
	fake.DialStub = stub
}

func (fake *Stower) DialArgsForCall(i int) (string, commands.Config) {
	fake.dialMutex.RLock()
	defer fake.dialMutex.RUnlock()
	argsForCall := fake.dialArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *Stower) DialReturns(result1 stow.Location, result2 error) {
	fake.dialMutex.Lock()
	defer fake.dialMutex.Unlock()
	fake.DialStub = nil
	fake.dialReturns = struct {
		result1 stow.Location
		result2 error
	}{result1, result2}
}

func (fake *Stower) DialReturnsOnCall(i int, result1 stow.Location, result2 error) {
	fake.dialMutex.Lock()
	defer fake.dialMutex.Unlock()
	fake.DialStub = nil
	if fake.dialReturnsOnCall == nil {
		fake.dialReturnsOnCall = make(map[int]struct {
			result1 stow.Location
			result2 error
		})
	}
	fake.dialReturnsOnCall[i] = struct {
		result1 stow.Location
		result2 error
	}{result1, result2}
}

func (fake *Stower) Walk(arg1 stow.Container, arg2 string, arg3 int, arg4 stow.WalkFunc) error {
	fake.walkMutex.Lock()
	ret, specificReturn := fake.walkReturnsOnCall[len(fake.walkArgsForCall)]
	fake.walkArgsForCall = append(fake.walkArgsForCall, struct {
		arg1 stow.Container
		arg2 string
		arg3 int
		arg4 stow.WalkFunc
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Walk", []interface{}{arg1, arg2, arg3, arg4})
	fake.walkMutex.Unlock()
	if fake.WalkStub != nil {
		return fake.WalkStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.walkReturns
	return fakeReturns.result1
}

func (fake *Stower) WalkCallCount() int {
	fake.walkMutex.RLock()
	defer fake.walkMutex.RUnlock()
	return len(fake.walkArgsForCall)
}

func (fake *Stower) WalkCalls(stub func(stow.Container, string, int, stow.WalkFunc) error) {
	fake.walkMutex.Lock()
	defer fake.walkMutex.Unlock()
	fake.WalkStub = stub
}

func (fake *Stower) WalkArgsForCall(i int) (stow.Container, string, int, stow.WalkFunc) {
	fake.walkMutex.RLock()
	defer fake.walkMutex.RUnlock()
	argsForCall := fake.walkArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *Stower) WalkReturns(result1 error) {
	fake.walkMutex.Lock()
	defer fake.walkMutex.Unlock()
	fake.WalkStub = nil
	fake.walkReturns = struct {
		result1 error
	}{result1}
}

func (fake *Stower) WalkReturnsOnCall(i int, result1 error) {
	fake.walkMutex.Lock()
	defer fake.walkMutex.Unlock()
	fake.WalkStub = nil
	if fake.walkReturnsOnCall == nil {
		fake.walkReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.walkReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Stower) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.dialMutex.RLock()
	defer fake.dialMutex.RUnlock()
	fake.walkMutex.RLock()
	defer fake.walkMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Stower) recordInvocation(key string, args []interface{}) {
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

var _ commands.Stower = new(Stower)
