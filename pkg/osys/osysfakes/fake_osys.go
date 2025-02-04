// Code generated by counterfeiter. DO NOT EDIT.
package osysfakes

import (
	"os"
	"sync"

	"github.com/weaveworks/weave-gitops/pkg/osys"
)

type FakeOsys struct {
	ExitStub        func(int)
	exitMutex       sync.RWMutex
	exitArgsForCall []struct {
		arg1 int
	}
	GetGitProviderTokenStub        func() (string, error)
	getGitProviderTokenMutex       sync.RWMutex
	getGitProviderTokenArgsForCall []struct {
	}
	getGitProviderTokenReturns struct {
		result1 string
		result2 error
	}
	getGitProviderTokenReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetenvStub        func(string) string
	getenvMutex       sync.RWMutex
	getenvArgsForCall []struct {
		arg1 string
	}
	getenvReturns struct {
		result1 string
	}
	getenvReturnsOnCall map[int]struct {
		result1 string
	}
	LookupEnvStub        func(string) (string, bool)
	lookupEnvMutex       sync.RWMutex
	lookupEnvArgsForCall []struct {
		arg1 string
	}
	lookupEnvReturns struct {
		result1 string
		result2 bool
	}
	lookupEnvReturnsOnCall map[int]struct {
		result1 string
		result2 bool
	}
	SetenvStub        func(string, string) error
	setenvMutex       sync.RWMutex
	setenvArgsForCall []struct {
		arg1 string
		arg2 string
	}
	setenvReturns struct {
		result1 error
	}
	setenvReturnsOnCall map[int]struct {
		result1 error
	}
	StderrStub        func() *os.File
	stderrMutex       sync.RWMutex
	stderrArgsForCall []struct {
	}
	stderrReturns struct {
		result1 *os.File
	}
	stderrReturnsOnCall map[int]struct {
		result1 *os.File
	}
	StdinStub        func() *os.File
	stdinMutex       sync.RWMutex
	stdinArgsForCall []struct {
	}
	stdinReturns struct {
		result1 *os.File
	}
	stdinReturnsOnCall map[int]struct {
		result1 *os.File
	}
	StdoutStub        func() *os.File
	stdoutMutex       sync.RWMutex
	stdoutArgsForCall []struct {
	}
	stdoutReturns struct {
		result1 *os.File
	}
	stdoutReturnsOnCall map[int]struct {
		result1 *os.File
	}
	UserHomeDirStub        func() (string, error)
	userHomeDirMutex       sync.RWMutex
	userHomeDirArgsForCall []struct {
	}
	userHomeDirReturns struct {
		result1 string
		result2 error
	}
	userHomeDirReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOsys) Exit(arg1 int) {
	fake.exitMutex.Lock()
	fake.exitArgsForCall = append(fake.exitArgsForCall, struct {
		arg1 int
	}{arg1})
	stub := fake.ExitStub
	fake.recordInvocation("Exit", []interface{}{arg1})
	fake.exitMutex.Unlock()
	if stub != nil {
		fake.ExitStub(arg1)
	}
}

func (fake *FakeOsys) ExitCallCount() int {
	fake.exitMutex.RLock()
	defer fake.exitMutex.RUnlock()
	return len(fake.exitArgsForCall)
}

func (fake *FakeOsys) ExitCalls(stub func(int)) {
	fake.exitMutex.Lock()
	defer fake.exitMutex.Unlock()
	fake.ExitStub = stub
}

func (fake *FakeOsys) ExitArgsForCall(i int) int {
	fake.exitMutex.RLock()
	defer fake.exitMutex.RUnlock()
	argsForCall := fake.exitArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeOsys) GetGitProviderToken() (string, error) {
	fake.getGitProviderTokenMutex.Lock()
	ret, specificReturn := fake.getGitProviderTokenReturnsOnCall[len(fake.getGitProviderTokenArgsForCall)]
	fake.getGitProviderTokenArgsForCall = append(fake.getGitProviderTokenArgsForCall, struct {
	}{})
	stub := fake.GetGitProviderTokenStub
	fakeReturns := fake.getGitProviderTokenReturns
	fake.recordInvocation("GetGitProviderToken", []interface{}{})
	fake.getGitProviderTokenMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeOsys) GetGitProviderTokenCallCount() int {
	fake.getGitProviderTokenMutex.RLock()
	defer fake.getGitProviderTokenMutex.RUnlock()
	return len(fake.getGitProviderTokenArgsForCall)
}

func (fake *FakeOsys) GetGitProviderTokenCalls(stub func() (string, error)) {
	fake.getGitProviderTokenMutex.Lock()
	defer fake.getGitProviderTokenMutex.Unlock()
	fake.GetGitProviderTokenStub = stub
}

func (fake *FakeOsys) GetGitProviderTokenReturns(result1 string, result2 error) {
	fake.getGitProviderTokenMutex.Lock()
	defer fake.getGitProviderTokenMutex.Unlock()
	fake.GetGitProviderTokenStub = nil
	fake.getGitProviderTokenReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeOsys) GetGitProviderTokenReturnsOnCall(i int, result1 string, result2 error) {
	fake.getGitProviderTokenMutex.Lock()
	defer fake.getGitProviderTokenMutex.Unlock()
	fake.GetGitProviderTokenStub = nil
	if fake.getGitProviderTokenReturnsOnCall == nil {
		fake.getGitProviderTokenReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getGitProviderTokenReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeOsys) Getenv(arg1 string) string {
	fake.getenvMutex.Lock()
	ret, specificReturn := fake.getenvReturnsOnCall[len(fake.getenvArgsForCall)]
	fake.getenvArgsForCall = append(fake.getenvArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetenvStub
	fakeReturns := fake.getenvReturns
	fake.recordInvocation("Getenv", []interface{}{arg1})
	fake.getenvMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeOsys) GetenvCallCount() int {
	fake.getenvMutex.RLock()
	defer fake.getenvMutex.RUnlock()
	return len(fake.getenvArgsForCall)
}

func (fake *FakeOsys) GetenvCalls(stub func(string) string) {
	fake.getenvMutex.Lock()
	defer fake.getenvMutex.Unlock()
	fake.GetenvStub = stub
}

func (fake *FakeOsys) GetenvArgsForCall(i int) string {
	fake.getenvMutex.RLock()
	defer fake.getenvMutex.RUnlock()
	argsForCall := fake.getenvArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeOsys) GetenvReturns(result1 string) {
	fake.getenvMutex.Lock()
	defer fake.getenvMutex.Unlock()
	fake.GetenvStub = nil
	fake.getenvReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeOsys) GetenvReturnsOnCall(i int, result1 string) {
	fake.getenvMutex.Lock()
	defer fake.getenvMutex.Unlock()
	fake.GetenvStub = nil
	if fake.getenvReturnsOnCall == nil {
		fake.getenvReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getenvReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeOsys) LookupEnv(arg1 string) (string, bool) {
	fake.lookupEnvMutex.Lock()
	ret, specificReturn := fake.lookupEnvReturnsOnCall[len(fake.lookupEnvArgsForCall)]
	fake.lookupEnvArgsForCall = append(fake.lookupEnvArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.LookupEnvStub
	fakeReturns := fake.lookupEnvReturns
	fake.recordInvocation("LookupEnv", []interface{}{arg1})
	fake.lookupEnvMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeOsys) LookupEnvCallCount() int {
	fake.lookupEnvMutex.RLock()
	defer fake.lookupEnvMutex.RUnlock()
	return len(fake.lookupEnvArgsForCall)
}

func (fake *FakeOsys) LookupEnvCalls(stub func(string) (string, bool)) {
	fake.lookupEnvMutex.Lock()
	defer fake.lookupEnvMutex.Unlock()
	fake.LookupEnvStub = stub
}

func (fake *FakeOsys) LookupEnvArgsForCall(i int) string {
	fake.lookupEnvMutex.RLock()
	defer fake.lookupEnvMutex.RUnlock()
	argsForCall := fake.lookupEnvArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeOsys) LookupEnvReturns(result1 string, result2 bool) {
	fake.lookupEnvMutex.Lock()
	defer fake.lookupEnvMutex.Unlock()
	fake.LookupEnvStub = nil
	fake.lookupEnvReturns = struct {
		result1 string
		result2 bool
	}{result1, result2}
}

func (fake *FakeOsys) LookupEnvReturnsOnCall(i int, result1 string, result2 bool) {
	fake.lookupEnvMutex.Lock()
	defer fake.lookupEnvMutex.Unlock()
	fake.LookupEnvStub = nil
	if fake.lookupEnvReturnsOnCall == nil {
		fake.lookupEnvReturnsOnCall = make(map[int]struct {
			result1 string
			result2 bool
		})
	}
	fake.lookupEnvReturnsOnCall[i] = struct {
		result1 string
		result2 bool
	}{result1, result2}
}

func (fake *FakeOsys) Setenv(arg1 string, arg2 string) error {
	fake.setenvMutex.Lock()
	ret, specificReturn := fake.setenvReturnsOnCall[len(fake.setenvArgsForCall)]
	fake.setenvArgsForCall = append(fake.setenvArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.SetenvStub
	fakeReturns := fake.setenvReturns
	fake.recordInvocation("Setenv", []interface{}{arg1, arg2})
	fake.setenvMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeOsys) SetenvCallCount() int {
	fake.setenvMutex.RLock()
	defer fake.setenvMutex.RUnlock()
	return len(fake.setenvArgsForCall)
}

func (fake *FakeOsys) SetenvCalls(stub func(string, string) error) {
	fake.setenvMutex.Lock()
	defer fake.setenvMutex.Unlock()
	fake.SetenvStub = stub
}

func (fake *FakeOsys) SetenvArgsForCall(i int) (string, string) {
	fake.setenvMutex.RLock()
	defer fake.setenvMutex.RUnlock()
	argsForCall := fake.setenvArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeOsys) SetenvReturns(result1 error) {
	fake.setenvMutex.Lock()
	defer fake.setenvMutex.Unlock()
	fake.SetenvStub = nil
	fake.setenvReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOsys) SetenvReturnsOnCall(i int, result1 error) {
	fake.setenvMutex.Lock()
	defer fake.setenvMutex.Unlock()
	fake.SetenvStub = nil
	if fake.setenvReturnsOnCall == nil {
		fake.setenvReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setenvReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOsys) Stderr() *os.File {
	fake.stderrMutex.Lock()
	ret, specificReturn := fake.stderrReturnsOnCall[len(fake.stderrArgsForCall)]
	fake.stderrArgsForCall = append(fake.stderrArgsForCall, struct {
	}{})
	stub := fake.StderrStub
	fakeReturns := fake.stderrReturns
	fake.recordInvocation("Stderr", []interface{}{})
	fake.stderrMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeOsys) StderrCallCount() int {
	fake.stderrMutex.RLock()
	defer fake.stderrMutex.RUnlock()
	return len(fake.stderrArgsForCall)
}

func (fake *FakeOsys) StderrCalls(stub func() *os.File) {
	fake.stderrMutex.Lock()
	defer fake.stderrMutex.Unlock()
	fake.StderrStub = stub
}

func (fake *FakeOsys) StderrReturns(result1 *os.File) {
	fake.stderrMutex.Lock()
	defer fake.stderrMutex.Unlock()
	fake.StderrStub = nil
	fake.stderrReturns = struct {
		result1 *os.File
	}{result1}
}

func (fake *FakeOsys) StderrReturnsOnCall(i int, result1 *os.File) {
	fake.stderrMutex.Lock()
	defer fake.stderrMutex.Unlock()
	fake.StderrStub = nil
	if fake.stderrReturnsOnCall == nil {
		fake.stderrReturnsOnCall = make(map[int]struct {
			result1 *os.File
		})
	}
	fake.stderrReturnsOnCall[i] = struct {
		result1 *os.File
	}{result1}
}

func (fake *FakeOsys) Stdin() *os.File {
	fake.stdinMutex.Lock()
	ret, specificReturn := fake.stdinReturnsOnCall[len(fake.stdinArgsForCall)]
	fake.stdinArgsForCall = append(fake.stdinArgsForCall, struct {
	}{})
	stub := fake.StdinStub
	fakeReturns := fake.stdinReturns
	fake.recordInvocation("Stdin", []interface{}{})
	fake.stdinMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeOsys) StdinCallCount() int {
	fake.stdinMutex.RLock()
	defer fake.stdinMutex.RUnlock()
	return len(fake.stdinArgsForCall)
}

func (fake *FakeOsys) StdinCalls(stub func() *os.File) {
	fake.stdinMutex.Lock()
	defer fake.stdinMutex.Unlock()
	fake.StdinStub = stub
}

func (fake *FakeOsys) StdinReturns(result1 *os.File) {
	fake.stdinMutex.Lock()
	defer fake.stdinMutex.Unlock()
	fake.StdinStub = nil
	fake.stdinReturns = struct {
		result1 *os.File
	}{result1}
}

func (fake *FakeOsys) StdinReturnsOnCall(i int, result1 *os.File) {
	fake.stdinMutex.Lock()
	defer fake.stdinMutex.Unlock()
	fake.StdinStub = nil
	if fake.stdinReturnsOnCall == nil {
		fake.stdinReturnsOnCall = make(map[int]struct {
			result1 *os.File
		})
	}
	fake.stdinReturnsOnCall[i] = struct {
		result1 *os.File
	}{result1}
}

func (fake *FakeOsys) Stdout() *os.File {
	fake.stdoutMutex.Lock()
	ret, specificReturn := fake.stdoutReturnsOnCall[len(fake.stdoutArgsForCall)]
	fake.stdoutArgsForCall = append(fake.stdoutArgsForCall, struct {
	}{})
	stub := fake.StdoutStub
	fakeReturns := fake.stdoutReturns
	fake.recordInvocation("Stdout", []interface{}{})
	fake.stdoutMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeOsys) StdoutCallCount() int {
	fake.stdoutMutex.RLock()
	defer fake.stdoutMutex.RUnlock()
	return len(fake.stdoutArgsForCall)
}

func (fake *FakeOsys) StdoutCalls(stub func() *os.File) {
	fake.stdoutMutex.Lock()
	defer fake.stdoutMutex.Unlock()
	fake.StdoutStub = stub
}

func (fake *FakeOsys) StdoutReturns(result1 *os.File) {
	fake.stdoutMutex.Lock()
	defer fake.stdoutMutex.Unlock()
	fake.StdoutStub = nil
	fake.stdoutReturns = struct {
		result1 *os.File
	}{result1}
}

func (fake *FakeOsys) StdoutReturnsOnCall(i int, result1 *os.File) {
	fake.stdoutMutex.Lock()
	defer fake.stdoutMutex.Unlock()
	fake.StdoutStub = nil
	if fake.stdoutReturnsOnCall == nil {
		fake.stdoutReturnsOnCall = make(map[int]struct {
			result1 *os.File
		})
	}
	fake.stdoutReturnsOnCall[i] = struct {
		result1 *os.File
	}{result1}
}

func (fake *FakeOsys) UserHomeDir() (string, error) {
	fake.userHomeDirMutex.Lock()
	ret, specificReturn := fake.userHomeDirReturnsOnCall[len(fake.userHomeDirArgsForCall)]
	fake.userHomeDirArgsForCall = append(fake.userHomeDirArgsForCall, struct {
	}{})
	stub := fake.UserHomeDirStub
	fakeReturns := fake.userHomeDirReturns
	fake.recordInvocation("UserHomeDir", []interface{}{})
	fake.userHomeDirMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeOsys) UserHomeDirCallCount() int {
	fake.userHomeDirMutex.RLock()
	defer fake.userHomeDirMutex.RUnlock()
	return len(fake.userHomeDirArgsForCall)
}

func (fake *FakeOsys) UserHomeDirCalls(stub func() (string, error)) {
	fake.userHomeDirMutex.Lock()
	defer fake.userHomeDirMutex.Unlock()
	fake.UserHomeDirStub = stub
}

func (fake *FakeOsys) UserHomeDirReturns(result1 string, result2 error) {
	fake.userHomeDirMutex.Lock()
	defer fake.userHomeDirMutex.Unlock()
	fake.UserHomeDirStub = nil
	fake.userHomeDirReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeOsys) UserHomeDirReturnsOnCall(i int, result1 string, result2 error) {
	fake.userHomeDirMutex.Lock()
	defer fake.userHomeDirMutex.Unlock()
	fake.UserHomeDirStub = nil
	if fake.userHomeDirReturnsOnCall == nil {
		fake.userHomeDirReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.userHomeDirReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeOsys) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.exitMutex.RLock()
	defer fake.exitMutex.RUnlock()
	fake.getGitProviderTokenMutex.RLock()
	defer fake.getGitProviderTokenMutex.RUnlock()
	fake.getenvMutex.RLock()
	defer fake.getenvMutex.RUnlock()
	fake.lookupEnvMutex.RLock()
	defer fake.lookupEnvMutex.RUnlock()
	fake.setenvMutex.RLock()
	defer fake.setenvMutex.RUnlock()
	fake.stderrMutex.RLock()
	defer fake.stderrMutex.RUnlock()
	fake.stdinMutex.RLock()
	defer fake.stdinMutex.RUnlock()
	fake.stdoutMutex.RLock()
	defer fake.stdoutMutex.RUnlock()
	fake.userHomeDirMutex.RLock()
	defer fake.userHomeDirMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeOsys) recordInvocation(key string, args []interface{}) {
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

var _ osys.Osys = new(FakeOsys)
