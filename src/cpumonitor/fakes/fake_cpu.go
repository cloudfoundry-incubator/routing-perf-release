// This file was generated by counterfeiter
package fakes

import (
	"sync"
	"time"
)

type FakeCpu struct {
	PercentStub        func(time.Duration, bool) ([]float64, error)
	percentMutex       sync.RWMutex
	percentArgsForCall []struct {
		arg1 time.Duration
		arg2 bool
	}
	percentReturns struct {
		result1 []float64
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCpu) Percent(arg1 time.Duration, arg2 bool) ([]float64, error) {
	fake.percentMutex.Lock()
	fake.percentArgsForCall = append(fake.percentArgsForCall, struct {
		arg1 time.Duration
		arg2 bool
	}{arg1, arg2})
	fake.recordInvocation("Percent", []interface{}{arg1, arg2})
	fake.percentMutex.Unlock()
	if fake.PercentStub != nil {
		return fake.PercentStub(arg1, arg2)
	} else {
		return fake.percentReturns.result1, fake.percentReturns.result2
	}
}

func (fake *FakeCpu) PercentCallCount() int {
	fake.percentMutex.RLock()
	defer fake.percentMutex.RUnlock()
	return len(fake.percentArgsForCall)
}

func (fake *FakeCpu) PercentArgsForCall(i int) (time.Duration, bool) {
	fake.percentMutex.RLock()
	defer fake.percentMutex.RUnlock()
	return fake.percentArgsForCall[i].arg1, fake.percentArgsForCall[i].arg2
}

func (fake *FakeCpu) PercentReturns(result1 []float64, result2 error) {
	fake.PercentStub = nil
	fake.percentReturns = struct {
		result1 []float64
		result2 error
	}{result1, result2}
}

func (fake *FakeCpu) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.percentMutex.RLock()
	defer fake.percentMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeCpu) recordInvocation(key string, args []interface{}) {
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
