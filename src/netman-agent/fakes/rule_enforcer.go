// This file was generated by counterfeiter
package fakes

import (
	"netman-agent/rules"
	"sync"
)

type RuleEnforcer struct {
	EnforceStub        func(chain string, r []rules.Rule) error
	enforceMutex       sync.RWMutex
	enforceArgsForCall []struct {
		chain string
		r     []rules.Rule
	}
	enforceReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *RuleEnforcer) Enforce(chain string, r []rules.Rule) error {
	var rCopy []rules.Rule
	if r != nil {
		rCopy = make([]rules.Rule, len(r))
		copy(rCopy, r)
	}
	fake.enforceMutex.Lock()
	fake.enforceArgsForCall = append(fake.enforceArgsForCall, struct {
		chain string
		r     []rules.Rule
	}{chain, rCopy})
	fake.recordInvocation("Enforce", []interface{}{chain, rCopy})
	fake.enforceMutex.Unlock()
	if fake.EnforceStub != nil {
		return fake.EnforceStub(chain, r)
	} else {
		return fake.enforceReturns.result1
	}
}

func (fake *RuleEnforcer) EnforceCallCount() int {
	fake.enforceMutex.RLock()
	defer fake.enforceMutex.RUnlock()
	return len(fake.enforceArgsForCall)
}

func (fake *RuleEnforcer) EnforceArgsForCall(i int) (string, []rules.Rule) {
	fake.enforceMutex.RLock()
	defer fake.enforceMutex.RUnlock()
	return fake.enforceArgsForCall[i].chain, fake.enforceArgsForCall[i].r
}

func (fake *RuleEnforcer) EnforceReturns(result1 error) {
	fake.EnforceStub = nil
	fake.enforceReturns = struct {
		result1 error
	}{result1}
}

func (fake *RuleEnforcer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.enforceMutex.RLock()
	defer fake.enforceMutex.RUnlock()
	return fake.invocations
}

func (fake *RuleEnforcer) recordInvocation(key string, args []interface{}) {
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

var _ rules.RuleEnforcer = new(RuleEnforcer)
