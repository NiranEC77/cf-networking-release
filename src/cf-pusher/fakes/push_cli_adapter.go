// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
)

type PushCLIAdapter struct {
	CheckAppStub        func(guid string) ([]byte, error)
	checkAppMutex       sync.RWMutex
	checkAppArgsForCall []struct {
		guid string
	}
	checkAppReturns struct {
		result1 []byte
		result2 error
	}
	checkAppReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	AppGuidStub        func(name string) (string, error)
	appGuidMutex       sync.RWMutex
	appGuidArgsForCall []struct {
		name string
	}
	appGuidReturns struct {
		result1 string
		result2 error
	}
	appGuidReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	PushStub        func(name, directory, manifestFile string) error
	pushMutex       sync.RWMutex
	pushArgsForCall []struct {
		name         string
		directory    string
		manifestFile string
	}
	pushReturns struct {
		result1 error
	}
	pushReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PushCLIAdapter) CheckApp(guid string) ([]byte, error) {
	fake.checkAppMutex.Lock()
	ret, specificReturn := fake.checkAppReturnsOnCall[len(fake.checkAppArgsForCall)]
	fake.checkAppArgsForCall = append(fake.checkAppArgsForCall, struct {
		guid string
	}{guid})
	fake.recordInvocation("CheckApp", []interface{}{guid})
	fake.checkAppMutex.Unlock()
	if fake.CheckAppStub != nil {
		return fake.CheckAppStub(guid)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.checkAppReturns.result1, fake.checkAppReturns.result2
}

func (fake *PushCLIAdapter) CheckAppCallCount() int {
	fake.checkAppMutex.RLock()
	defer fake.checkAppMutex.RUnlock()
	return len(fake.checkAppArgsForCall)
}

func (fake *PushCLIAdapter) CheckAppArgsForCall(i int) string {
	fake.checkAppMutex.RLock()
	defer fake.checkAppMutex.RUnlock()
	return fake.checkAppArgsForCall[i].guid
}

func (fake *PushCLIAdapter) CheckAppReturns(result1 []byte, result2 error) {
	fake.CheckAppStub = nil
	fake.checkAppReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *PushCLIAdapter) CheckAppReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.CheckAppStub = nil
	if fake.checkAppReturnsOnCall == nil {
		fake.checkAppReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.checkAppReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *PushCLIAdapter) AppGuid(name string) (string, error) {
	fake.appGuidMutex.Lock()
	ret, specificReturn := fake.appGuidReturnsOnCall[len(fake.appGuidArgsForCall)]
	fake.appGuidArgsForCall = append(fake.appGuidArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("AppGuid", []interface{}{name})
	fake.appGuidMutex.Unlock()
	if fake.AppGuidStub != nil {
		return fake.AppGuidStub(name)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.appGuidReturns.result1, fake.appGuidReturns.result2
}

func (fake *PushCLIAdapter) AppGuidCallCount() int {
	fake.appGuidMutex.RLock()
	defer fake.appGuidMutex.RUnlock()
	return len(fake.appGuidArgsForCall)
}

func (fake *PushCLIAdapter) AppGuidArgsForCall(i int) string {
	fake.appGuidMutex.RLock()
	defer fake.appGuidMutex.RUnlock()
	return fake.appGuidArgsForCall[i].name
}

func (fake *PushCLIAdapter) AppGuidReturns(result1 string, result2 error) {
	fake.AppGuidStub = nil
	fake.appGuidReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *PushCLIAdapter) AppGuidReturnsOnCall(i int, result1 string, result2 error) {
	fake.AppGuidStub = nil
	if fake.appGuidReturnsOnCall == nil {
		fake.appGuidReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.appGuidReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *PushCLIAdapter) Push(name string, directory string, manifestFile string) error {
	fake.pushMutex.Lock()
	ret, specificReturn := fake.pushReturnsOnCall[len(fake.pushArgsForCall)]
	fake.pushArgsForCall = append(fake.pushArgsForCall, struct {
		name         string
		directory    string
		manifestFile string
	}{name, directory, manifestFile})
	fake.recordInvocation("Push", []interface{}{name, directory, manifestFile})
	fake.pushMutex.Unlock()
	if fake.PushStub != nil {
		return fake.PushStub(name, directory, manifestFile)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.pushReturns.result1
}

func (fake *PushCLIAdapter) PushCallCount() int {
	fake.pushMutex.RLock()
	defer fake.pushMutex.RUnlock()
	return len(fake.pushArgsForCall)
}

func (fake *PushCLIAdapter) PushArgsForCall(i int) (string, string, string) {
	fake.pushMutex.RLock()
	defer fake.pushMutex.RUnlock()
	return fake.pushArgsForCall[i].name, fake.pushArgsForCall[i].directory, fake.pushArgsForCall[i].manifestFile
}

func (fake *PushCLIAdapter) PushReturns(result1 error) {
	fake.PushStub = nil
	fake.pushReturns = struct {
		result1 error
	}{result1}
}

func (fake *PushCLIAdapter) PushReturnsOnCall(i int, result1 error) {
	fake.PushStub = nil
	if fake.pushReturnsOnCall == nil {
		fake.pushReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pushReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *PushCLIAdapter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkAppMutex.RLock()
	defer fake.checkAppMutex.RUnlock()
	fake.appGuidMutex.RLock()
	defer fake.appGuidMutex.RUnlock()
	fake.pushMutex.RLock()
	defer fake.pushMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *PushCLIAdapter) recordInvocation(key string, args []interface{}) {
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
