// Code generated by counterfeiter. DO NOT EDIT.
package gitprovidersfakes

import (
	"sync"

	"github.com/fluxcd/go-git-providers/gitprovider"
	"github.com/weaveworks/weave-gitops/pkg/gitproviders"
)

type FakeGitProvider struct {
	CreatePullRequestToOrgRepoStub        func(gitprovider.OrgRepositoryRef, string, string, []gitprovider.CommitFile, string, string, string) (gitprovider.PullRequest, error)
	createPullRequestToOrgRepoMutex       sync.RWMutex
	createPullRequestToOrgRepoArgsForCall []struct {
		arg1 gitprovider.OrgRepositoryRef
		arg2 string
		arg3 string
		arg4 []gitprovider.CommitFile
		arg5 string
		arg6 string
		arg7 string
	}
	createPullRequestToOrgRepoReturns struct {
		result1 gitprovider.PullRequest
		result2 error
	}
	createPullRequestToOrgRepoReturnsOnCall map[int]struct {
		result1 gitprovider.PullRequest
		result2 error
	}
	CreatePullRequestToUserRepoStub        func(gitprovider.UserRepositoryRef, string, string, []gitprovider.CommitFile, string, string, string) (gitprovider.PullRequest, error)
	createPullRequestToUserRepoMutex       sync.RWMutex
	createPullRequestToUserRepoArgsForCall []struct {
		arg1 gitprovider.UserRepositoryRef
		arg2 string
		arg3 string
		arg4 []gitprovider.CommitFile
		arg5 string
		arg6 string
		arg7 string
	}
	createPullRequestToUserRepoReturns struct {
		result1 gitprovider.PullRequest
		result2 error
	}
	createPullRequestToUserRepoReturnsOnCall map[int]struct {
		result1 gitprovider.PullRequest
		result2 error
	}
	CreateRepositoryStub        func(string, string, bool) error
	createRepositoryMutex       sync.RWMutex
	createRepositoryArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 bool
	}
	createRepositoryReturns struct {
		result1 error
	}
	createRepositoryReturnsOnCall map[int]struct {
		result1 error
	}
	DeployKeyExistsStub        func(string, string) (bool, error)
	deployKeyExistsMutex       sync.RWMutex
	deployKeyExistsArgsForCall []struct {
		arg1 string
		arg2 string
	}
	deployKeyExistsReturns struct {
		result1 bool
		result2 error
	}
	deployKeyExistsReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	GetAccountTypeStub        func(string) (gitproviders.ProviderAccountType, error)
	getAccountTypeMutex       sync.RWMutex
	getAccountTypeArgsForCall []struct {
		arg1 string
	}
	getAccountTypeReturns struct {
		result1 gitproviders.ProviderAccountType
		result2 error
	}
	getAccountTypeReturnsOnCall map[int]struct {
		result1 gitproviders.ProviderAccountType
		result2 error
	}
	GetCommitsFromOrgRepoStub        func(gitprovider.OrgRepositoryRef, string, int, int) ([]gitprovider.Commit, error)
	getCommitsFromOrgRepoMutex       sync.RWMutex
	getCommitsFromOrgRepoArgsForCall []struct {
		arg1 gitprovider.OrgRepositoryRef
		arg2 string
		arg3 int
		arg4 int
	}
	getCommitsFromOrgRepoReturns struct {
		result1 []gitprovider.Commit
		result2 error
	}
	getCommitsFromOrgRepoReturnsOnCall map[int]struct {
		result1 []gitprovider.Commit
		result2 error
	}
	GetCommitsFromUserRepoStub        func(gitprovider.UserRepositoryRef, string, int, int) ([]gitprovider.Commit, error)
	getCommitsFromUserRepoMutex       sync.RWMutex
	getCommitsFromUserRepoArgsForCall []struct {
		arg1 gitprovider.UserRepositoryRef
		arg2 string
		arg3 int
		arg4 int
	}
	getCommitsFromUserRepoReturns struct {
		result1 []gitprovider.Commit
		result2 error
	}
	getCommitsFromUserRepoReturnsOnCall map[int]struct {
		result1 []gitprovider.Commit
		result2 error
	}
	GetDefaultBranchStub        func(string) (string, error)
	getDefaultBranchMutex       sync.RWMutex
	getDefaultBranchArgsForCall []struct {
		arg1 string
	}
	getDefaultBranchReturns struct {
		result1 string
		result2 error
	}
	getDefaultBranchReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetRepoInfoStub        func(gitproviders.ProviderAccountType, string, string) (*gitprovider.RepositoryInfo, error)
	getRepoInfoMutex       sync.RWMutex
	getRepoInfoArgsForCall []struct {
		arg1 gitproviders.ProviderAccountType
		arg2 string
		arg3 string
	}
	getRepoInfoReturns struct {
		result1 *gitprovider.RepositoryInfo
		result2 error
	}
	getRepoInfoReturnsOnCall map[int]struct {
		result1 *gitprovider.RepositoryInfo
		result2 error
	}
	GetRepoInfoFromUrlStub        func(string) (*gitprovider.RepositoryInfo, error)
	getRepoInfoFromUrlMutex       sync.RWMutex
	getRepoInfoFromUrlArgsForCall []struct {
		arg1 string
	}
	getRepoInfoFromUrlReturns struct {
		result1 *gitprovider.RepositoryInfo
		result2 error
	}
	getRepoInfoFromUrlReturnsOnCall map[int]struct {
		result1 *gitprovider.RepositoryInfo
		result2 error
	}
	GetRepoVisibilityStub        func(string) (*gitprovider.RepositoryVisibility, error)
	getRepoVisibilityMutex       sync.RWMutex
	getRepoVisibilityArgsForCall []struct {
		arg1 string
	}
	getRepoVisibilityReturns struct {
		result1 *gitprovider.RepositoryVisibility
		result2 error
	}
	getRepoVisibilityReturnsOnCall map[int]struct {
		result1 *gitprovider.RepositoryVisibility
		result2 error
	}
	RepositoryExistsStub        func(string, string) (bool, error)
	repositoryExistsMutex       sync.RWMutex
	repositoryExistsArgsForCall []struct {
		arg1 string
		arg2 string
	}
	repositoryExistsReturns struct {
		result1 bool
		result2 error
	}
	repositoryExistsReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	UploadDeployKeyStub        func(string, string, []byte) error
	uploadDeployKeyMutex       sync.RWMutex
	uploadDeployKeyArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 []byte
	}
	uploadDeployKeyReturns struct {
		result1 error
	}
	uploadDeployKeyReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGitProvider) CreatePullRequestToOrgRepo(arg1 gitprovider.OrgRepositoryRef, arg2 string, arg3 string, arg4 []gitprovider.CommitFile, arg5 string, arg6 string, arg7 string) (gitprovider.PullRequest, error) {
	var arg4Copy []gitprovider.CommitFile
	if arg4 != nil {
		arg4Copy = make([]gitprovider.CommitFile, len(arg4))
		copy(arg4Copy, arg4)
	}
	fake.createPullRequestToOrgRepoMutex.Lock()
	ret, specificReturn := fake.createPullRequestToOrgRepoReturnsOnCall[len(fake.createPullRequestToOrgRepoArgsForCall)]
	fake.createPullRequestToOrgRepoArgsForCall = append(fake.createPullRequestToOrgRepoArgsForCall, struct {
		arg1 gitprovider.OrgRepositoryRef
		arg2 string
		arg3 string
		arg4 []gitprovider.CommitFile
		arg5 string
		arg6 string
		arg7 string
	}{arg1, arg2, arg3, arg4Copy, arg5, arg6, arg7})
	stub := fake.CreatePullRequestToOrgRepoStub
	fakeReturns := fake.createPullRequestToOrgRepoReturns
	fake.recordInvocation("CreatePullRequestToOrgRepo", []interface{}{arg1, arg2, arg3, arg4Copy, arg5, arg6, arg7})
	fake.createPullRequestToOrgRepoMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGitProvider) CreatePullRequestToOrgRepoCallCount() int {
	fake.createPullRequestToOrgRepoMutex.RLock()
	defer fake.createPullRequestToOrgRepoMutex.RUnlock()
	return len(fake.createPullRequestToOrgRepoArgsForCall)
}

func (fake *FakeGitProvider) CreatePullRequestToOrgRepoCalls(stub func(gitprovider.OrgRepositoryRef, string, string, []gitprovider.CommitFile, string, string, string) (gitprovider.PullRequest, error)) {
	fake.createPullRequestToOrgRepoMutex.Lock()
	defer fake.createPullRequestToOrgRepoMutex.Unlock()
	fake.CreatePullRequestToOrgRepoStub = stub
}

func (fake *FakeGitProvider) CreatePullRequestToOrgRepoArgsForCall(i int) (gitprovider.OrgRepositoryRef, string, string, []gitprovider.CommitFile, string, string, string) {
	fake.createPullRequestToOrgRepoMutex.RLock()
	defer fake.createPullRequestToOrgRepoMutex.RUnlock()
	argsForCall := fake.createPullRequestToOrgRepoArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7
}

func (fake *FakeGitProvider) CreatePullRequestToOrgRepoReturns(result1 gitprovider.PullRequest, result2 error) {
	fake.createPullRequestToOrgRepoMutex.Lock()
	defer fake.createPullRequestToOrgRepoMutex.Unlock()
	fake.CreatePullRequestToOrgRepoStub = nil
	fake.createPullRequestToOrgRepoReturns = struct {
		result1 gitprovider.PullRequest
		result2 error
	}{result1, result2}
}

func (fake *FakeGitProvider) CreatePullRequestToOrgRepoReturnsOnCall(i int, result1 gitprovider.PullRequest, result2 error) {
	fake.createPullRequestToOrgRepoMutex.Lock()
	defer fake.createPullRequestToOrgRepoMutex.Unlock()
	fake.CreatePullRequestToOrgRepoStub = nil
	if fake.createPullRequestToOrgRepoReturnsOnCall == nil {
		fake.createPullRequestToOrgRepoReturnsOnCall = make(map[int]struct {
			result1 gitprovider.PullRequest
			result2 error
		})
	}
	fake.createPullRequestToOrgRepoReturnsOnCall[i] = struct {
		result1 gitprovider.PullRequest
		result2 error
	}{result1, result2}
}

func (fake *FakeGitProvider) CreatePullRequestToUserRepo(arg1 gitprovider.UserRepositoryRef, arg2 string, arg3 string, arg4 []gitprovider.CommitFile, arg5 string, arg6 string, arg7 string) (gitprovider.PullRequest, error) {
	var arg4Copy []gitprovider.CommitFile
	if arg4 != nil {
		arg4Copy = make([]gitprovider.CommitFile, len(arg4))
		copy(arg4Copy, arg4)
	}
	fake.createPullRequestToUserRepoMutex.Lock()
	ret, specificReturn := fake.createPullRequestToUserRepoReturnsOnCall[len(fake.createPullRequestToUserRepoArgsForCall)]
	fake.createPullRequestToUserRepoArgsForCall = append(fake.createPullRequestToUserRepoArgsForCall, struct {
		arg1 gitprovider.UserRepositoryRef
		arg2 string
		arg3 string
		arg4 []gitprovider.CommitFile
		arg5 string
		arg6 string
		arg7 string
	}{arg1, arg2, arg3, arg4Copy, arg5, arg6, arg7})
	stub := fake.CreatePullRequestToUserRepoStub
	fakeReturns := fake.createPullRequestToUserRepoReturns
	fake.recordInvocation("CreatePullRequestToUserRepo", []interface{}{arg1, arg2, arg3, arg4Copy, arg5, arg6, arg7})
	fake.createPullRequestToUserRepoMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGitProvider) CreatePullRequestToUserRepoCallCount() int {
	fake.createPullRequestToUserRepoMutex.RLock()
	defer fake.createPullRequestToUserRepoMutex.RUnlock()
	return len(fake.createPullRequestToUserRepoArgsForCall)
}

func (fake *FakeGitProvider) CreatePullRequestToUserRepoCalls(stub func(gitprovider.UserRepositoryRef, string, string, []gitprovider.CommitFile, string, string, string) (gitprovider.PullRequest, error)) {
	fake.createPullRequestToUserRepoMutex.Lock()
	defer fake.createPullRequestToUserRepoMutex.Unlock()
	fake.CreatePullRequestToUserRepoStub = stub
}

func (fake *FakeGitProvider) CreatePullRequestToUserRepoArgsForCall(i int) (gitprovider.UserRepositoryRef, string, string, []gitprovider.CommitFile, string, string, string) {
	fake.createPullRequestToUserRepoMutex.RLock()
	defer fake.createPullRequestToUserRepoMutex.RUnlock()
	argsForCall := fake.createPullRequestToUserRepoArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7
}

func (fake *FakeGitProvider) CreatePullRequestToUserRepoReturns(result1 gitprovider.PullRequest, result2 error) {
	fake.createPullRequestToUserRepoMutex.Lock()
	defer fake.createPullRequestToUserRepoMutex.Unlock()
	fake.CreatePullRequestToUserRepoStub = nil
	fake.createPullRequestToUserRepoReturns = struct {
		result1 gitprovider.PullRequest
		result2 error
	}{result1, result2}
}

func (fake *FakeGitProvider) CreatePullRequestToUserRepoReturnsOnCall(i int, result1 gitprovider.PullRequest, result2 error) {
	fake.createPullRequestToUserRepoMutex.Lock()
	defer fake.createPullRequestToUserRepoMutex.Unlock()
	fake.CreatePullRequestToUserRepoStub = nil
	if fake.createPullRequestToUserRepoReturnsOnCall == nil {
		fake.createPullRequestToUserRepoReturnsOnCall = make(map[int]struct {
			result1 gitprovider.PullRequest
			result2 error
		})
	}
	fake.createPullRequestToUserRepoReturnsOnCall[i] = struct {
		result1 gitprovider.PullRequest
		result2 error
	}{result1, result2}
}

func (fake *FakeGitProvider) CreateRepository(arg1 string, arg2 string, arg3 bool) error {
	fake.createRepositoryMutex.Lock()
	ret, specificReturn := fake.createRepositoryReturnsOnCall[len(fake.createRepositoryArgsForCall)]
	fake.createRepositoryArgsForCall = append(fake.createRepositoryArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 bool
	}{arg1, arg2, arg3})
	stub := fake.CreateRepositoryStub
	fakeReturns := fake.createRepositoryReturns
	fake.recordInvocation("CreateRepository", []interface{}{arg1, arg2, arg3})
	fake.createRepositoryMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeGitProvider) CreateRepositoryCallCount() int {
	fake.createRepositoryMutex.RLock()
	defer fake.createRepositoryMutex.RUnlock()
	return len(fake.createRepositoryArgsForCall)
}

func (fake *FakeGitProvider) CreateRepositoryCalls(stub func(string, string, bool) error) {
	fake.createRepositoryMutex.Lock()
	defer fake.createRepositoryMutex.Unlock()
	fake.CreateRepositoryStub = stub
}

func (fake *FakeGitProvider) CreateRepositoryArgsForCall(i int) (string, string, bool) {
	fake.createRepositoryMutex.RLock()
	defer fake.createRepositoryMutex.RUnlock()
	argsForCall := fake.createRepositoryArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeGitProvider) CreateRepositoryReturns(result1 error) {
	fake.createRepositoryMutex.Lock()
	defer fake.createRepositoryMutex.Unlock()
	fake.CreateRepositoryStub = nil
	fake.createRepositoryReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGitProvider) CreateRepositoryReturnsOnCall(i int, result1 error) {
	fake.createRepositoryMutex.Lock()
	defer fake.createRepositoryMutex.Unlock()
	fake.CreateRepositoryStub = nil
	if fake.createRepositoryReturnsOnCall == nil {
		fake.createRepositoryReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createRepositoryReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeGitProvider) DeployKeyExists(arg1 string, arg2 string) (bool, error) {
	fake.deployKeyExistsMutex.Lock()
	ret, specificReturn := fake.deployKeyExistsReturnsOnCall[len(fake.deployKeyExistsArgsForCall)]
	fake.deployKeyExistsArgsForCall = append(fake.deployKeyExistsArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.DeployKeyExistsStub
	fakeReturns := fake.deployKeyExistsReturns
	fake.recordInvocation("DeployKeyExists", []interface{}{arg1, arg2})
	fake.deployKeyExistsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGitProvider) DeployKeyExistsCallCount() int {
	fake.deployKeyExistsMutex.RLock()
	defer fake.deployKeyExistsMutex.RUnlock()
	return len(fake.deployKeyExistsArgsForCall)
}

func (fake *FakeGitProvider) DeployKeyExistsCalls(stub func(string, string) (bool, error)) {
	fake.deployKeyExistsMutex.Lock()
	defer fake.deployKeyExistsMutex.Unlock()
	fake.DeployKeyExistsStub = stub
}

func (fake *FakeGitProvider) DeployKeyExistsArgsForCall(i int) (string, string) {
	fake.deployKeyExistsMutex.RLock()
	defer fake.deployKeyExistsMutex.RUnlock()
	argsForCall := fake.deployKeyExistsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeGitProvider) DeployKeyExistsReturns(result1 bool, result2 error) {
	fake.deployKeyExistsMutex.Lock()
	defer fake.deployKeyExistsMutex.Unlock()
	fake.DeployKeyExistsStub = nil
	fake.deployKeyExistsReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeGitProvider) DeployKeyExistsReturnsOnCall(i int, result1 bool, result2 error) {
	fake.deployKeyExistsMutex.Lock()
	defer fake.deployKeyExistsMutex.Unlock()
	fake.DeployKeyExistsStub = nil
	if fake.deployKeyExistsReturnsOnCall == nil {
		fake.deployKeyExistsReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.deployKeyExistsReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeGitProvider) GetAccountType(arg1 string) (gitproviders.ProviderAccountType, error) {
	fake.getAccountTypeMutex.Lock()
	ret, specificReturn := fake.getAccountTypeReturnsOnCall[len(fake.getAccountTypeArgsForCall)]
	fake.getAccountTypeArgsForCall = append(fake.getAccountTypeArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetAccountTypeStub
	fakeReturns := fake.getAccountTypeReturns
	fake.recordInvocation("GetAccountType", []interface{}{arg1})
	fake.getAccountTypeMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGitProvider) GetAccountTypeCallCount() int {
	fake.getAccountTypeMutex.RLock()
	defer fake.getAccountTypeMutex.RUnlock()
	return len(fake.getAccountTypeArgsForCall)
}

func (fake *FakeGitProvider) GetAccountTypeCalls(stub func(string) (gitproviders.ProviderAccountType, error)) {
	fake.getAccountTypeMutex.Lock()
	defer fake.getAccountTypeMutex.Unlock()
	fake.GetAccountTypeStub = stub
}

func (fake *FakeGitProvider) GetAccountTypeArgsForCall(i int) string {
	fake.getAccountTypeMutex.RLock()
	defer fake.getAccountTypeMutex.RUnlock()
	argsForCall := fake.getAccountTypeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeGitProvider) GetAccountTypeReturns(result1 gitproviders.ProviderAccountType, result2 error) {
	fake.getAccountTypeMutex.Lock()
	defer fake.getAccountTypeMutex.Unlock()
	fake.GetAccountTypeStub = nil
	fake.getAccountTypeReturns = struct {
		result1 gitproviders.ProviderAccountType
		result2 error
	}{result1, result2}
}

func (fake *FakeGitProvider) GetAccountTypeReturnsOnCall(i int, result1 gitproviders.ProviderAccountType, result2 error) {
	fake.getAccountTypeMutex.Lock()
	defer fake.getAccountTypeMutex.Unlock()
	fake.GetAccountTypeStub = nil
	if fake.getAccountTypeReturnsOnCall == nil {
		fake.getAccountTypeReturnsOnCall = make(map[int]struct {
			result1 gitproviders.ProviderAccountType
			result2 error
		})
	}
	fake.getAccountTypeReturnsOnCall[i] = struct {
		result1 gitproviders.ProviderAccountType
		result2 error
	}{result1, result2}
}

func (fake *FakeGitProvider) GetCommitsFromOrgRepo(arg1 gitprovider.OrgRepositoryRef, arg2 string, arg3 int, arg4 int) ([]gitprovider.Commit, error) {
	fake.getCommitsFromOrgRepoMutex.Lock()
	ret, specificReturn := fake.getCommitsFromOrgRepoReturnsOnCall[len(fake.getCommitsFromOrgRepoArgsForCall)]
	fake.getCommitsFromOrgRepoArgsForCall = append(fake.getCommitsFromOrgRepoArgsForCall, struct {
		arg1 gitprovider.OrgRepositoryRef
		arg2 string
		arg3 int
		arg4 int
	}{arg1, arg2, arg3, arg4})
	stub := fake.GetCommitsFromOrgRepoStub
	fakeReturns := fake.getCommitsFromOrgRepoReturns
	fake.recordInvocation("GetCommitsFromOrgRepo", []interface{}{arg1, arg2, arg3, arg4})
	fake.getCommitsFromOrgRepoMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGitProvider) GetCommitsFromOrgRepoCallCount() int {
	fake.getCommitsFromOrgRepoMutex.RLock()
	defer fake.getCommitsFromOrgRepoMutex.RUnlock()
	return len(fake.getCommitsFromOrgRepoArgsForCall)
}

func (fake *FakeGitProvider) GetCommitsFromOrgRepoCalls(stub func(gitprovider.OrgRepositoryRef, string, int, int) ([]gitprovider.Commit, error)) {
	fake.getCommitsFromOrgRepoMutex.Lock()
	defer fake.getCommitsFromOrgRepoMutex.Unlock()
	fake.GetCommitsFromOrgRepoStub = stub
}

func (fake *FakeGitProvider) GetCommitsFromOrgRepoArgsForCall(i int) (gitprovider.OrgRepositoryRef, string, int, int) {
	fake.getCommitsFromOrgRepoMutex.RLock()
	defer fake.getCommitsFromOrgRepoMutex.RUnlock()
	argsForCall := fake.getCommitsFromOrgRepoArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeGitProvider) GetCommitsFromOrgRepoReturns(result1 []gitprovider.Commit, result2 error) {
	fake.getCommitsFromOrgRepoMutex.Lock()
	defer fake.getCommitsFromOrgRepoMutex.Unlock()
	fake.GetCommitsFromOrgRepoStub = nil
	fake.getCommitsFromOrgRepoReturns = struct {
		result1 []gitprovider.Commit
		result2 error
	}{result1, result2}
}

func (fake *FakeGitProvider) GetCommitsFromOrgRepoReturnsOnCall(i int, result1 []gitprovider.Commit, result2 error) {
	fake.getCommitsFromOrgRepoMutex.Lock()
	defer fake.getCommitsFromOrgRepoMutex.Unlock()
	fake.GetCommitsFromOrgRepoStub = nil
	if fake.getCommitsFromOrgRepoReturnsOnCall == nil {
		fake.getCommitsFromOrgRepoReturnsOnCall = make(map[int]struct {
			result1 []gitprovider.Commit
			result2 error
		})
	}
	fake.getCommitsFromOrgRepoReturnsOnCall[i] = struct {
		result1 []gitprovider.Commit
		result2 error
	}{result1, result2}
}

func (fake *FakeGitProvider) GetCommitsFromUserRepo(arg1 gitprovider.UserRepositoryRef, arg2 string, arg3 int, arg4 int) ([]gitprovider.Commit, error) {
	fake.getCommitsFromUserRepoMutex.Lock()
	ret, specificReturn := fake.getCommitsFromUserRepoReturnsOnCall[len(fake.getCommitsFromUserRepoArgsForCall)]
	fake.getCommitsFromUserRepoArgsForCall = append(fake.getCommitsFromUserRepoArgsForCall, struct {
		arg1 gitprovider.UserRepositoryRef
		arg2 string
		arg3 int
		arg4 int
	}{arg1, arg2, arg3, arg4})
	stub := fake.GetCommitsFromUserRepoStub
	fakeReturns := fake.getCommitsFromUserRepoReturns
	fake.recordInvocation("GetCommitsFromUserRepo", []interface{}{arg1, arg2, arg3, arg4})
	fake.getCommitsFromUserRepoMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGitProvider) GetCommitsFromUserRepoCallCount() int {
	fake.getCommitsFromUserRepoMutex.RLock()
	defer fake.getCommitsFromUserRepoMutex.RUnlock()
	return len(fake.getCommitsFromUserRepoArgsForCall)
}

func (fake *FakeGitProvider) GetCommitsFromUserRepoCalls(stub func(gitprovider.UserRepositoryRef, string, int, int) ([]gitprovider.Commit, error)) {
	fake.getCommitsFromUserRepoMutex.Lock()
	defer fake.getCommitsFromUserRepoMutex.Unlock()
	fake.GetCommitsFromUserRepoStub = stub
}

func (fake *FakeGitProvider) GetCommitsFromUserRepoArgsForCall(i int) (gitprovider.UserRepositoryRef, string, int, int) {
	fake.getCommitsFromUserRepoMutex.RLock()
	defer fake.getCommitsFromUserRepoMutex.RUnlock()
	argsForCall := fake.getCommitsFromUserRepoArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeGitProvider) GetCommitsFromUserRepoReturns(result1 []gitprovider.Commit, result2 error) {
	fake.getCommitsFromUserRepoMutex.Lock()
	defer fake.getCommitsFromUserRepoMutex.Unlock()
	fake.GetCommitsFromUserRepoStub = nil
	fake.getCommitsFromUserRepoReturns = struct {
		result1 []gitprovider.Commit
		result2 error
	}{result1, result2}
}

func (fake *FakeGitProvider) GetCommitsFromUserRepoReturnsOnCall(i int, result1 []gitprovider.Commit, result2 error) {
	fake.getCommitsFromUserRepoMutex.Lock()
	defer fake.getCommitsFromUserRepoMutex.Unlock()
	fake.GetCommitsFromUserRepoStub = nil
	if fake.getCommitsFromUserRepoReturnsOnCall == nil {
		fake.getCommitsFromUserRepoReturnsOnCall = make(map[int]struct {
			result1 []gitprovider.Commit
			result2 error
		})
	}
	fake.getCommitsFromUserRepoReturnsOnCall[i] = struct {
		result1 []gitprovider.Commit
		result2 error
	}{result1, result2}
}

func (fake *FakeGitProvider) GetDefaultBranch(arg1 string) (string, error) {
	fake.getDefaultBranchMutex.Lock()
	ret, specificReturn := fake.getDefaultBranchReturnsOnCall[len(fake.getDefaultBranchArgsForCall)]
	fake.getDefaultBranchArgsForCall = append(fake.getDefaultBranchArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetDefaultBranchStub
	fakeReturns := fake.getDefaultBranchReturns
	fake.recordInvocation("GetDefaultBranch", []interface{}{arg1})
	fake.getDefaultBranchMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGitProvider) GetDefaultBranchCallCount() int {
	fake.getDefaultBranchMutex.RLock()
	defer fake.getDefaultBranchMutex.RUnlock()
	return len(fake.getDefaultBranchArgsForCall)
}

func (fake *FakeGitProvider) GetDefaultBranchCalls(stub func(string) (string, error)) {
	fake.getDefaultBranchMutex.Lock()
	defer fake.getDefaultBranchMutex.Unlock()
	fake.GetDefaultBranchStub = stub
}

func (fake *FakeGitProvider) GetDefaultBranchArgsForCall(i int) string {
	fake.getDefaultBranchMutex.RLock()
	defer fake.getDefaultBranchMutex.RUnlock()
	argsForCall := fake.getDefaultBranchArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeGitProvider) GetDefaultBranchReturns(result1 string, result2 error) {
	fake.getDefaultBranchMutex.Lock()
	defer fake.getDefaultBranchMutex.Unlock()
	fake.GetDefaultBranchStub = nil
	fake.getDefaultBranchReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeGitProvider) GetDefaultBranchReturnsOnCall(i int, result1 string, result2 error) {
	fake.getDefaultBranchMutex.Lock()
	defer fake.getDefaultBranchMutex.Unlock()
	fake.GetDefaultBranchStub = nil
	if fake.getDefaultBranchReturnsOnCall == nil {
		fake.getDefaultBranchReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getDefaultBranchReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeGitProvider) GetRepoInfo(arg1 gitproviders.ProviderAccountType, arg2 string, arg3 string) (*gitprovider.RepositoryInfo, error) {
	fake.getRepoInfoMutex.Lock()
	ret, specificReturn := fake.getRepoInfoReturnsOnCall[len(fake.getRepoInfoArgsForCall)]
	fake.getRepoInfoArgsForCall = append(fake.getRepoInfoArgsForCall, struct {
		arg1 gitproviders.ProviderAccountType
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.GetRepoInfoStub
	fakeReturns := fake.getRepoInfoReturns
	fake.recordInvocation("GetRepoInfo", []interface{}{arg1, arg2, arg3})
	fake.getRepoInfoMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGitProvider) GetRepoInfoCallCount() int {
	fake.getRepoInfoMutex.RLock()
	defer fake.getRepoInfoMutex.RUnlock()
	return len(fake.getRepoInfoArgsForCall)
}

func (fake *FakeGitProvider) GetRepoInfoCalls(stub func(gitproviders.ProviderAccountType, string, string) (*gitprovider.RepositoryInfo, error)) {
	fake.getRepoInfoMutex.Lock()
	defer fake.getRepoInfoMutex.Unlock()
	fake.GetRepoInfoStub = stub
}

func (fake *FakeGitProvider) GetRepoInfoArgsForCall(i int) (gitproviders.ProviderAccountType, string, string) {
	fake.getRepoInfoMutex.RLock()
	defer fake.getRepoInfoMutex.RUnlock()
	argsForCall := fake.getRepoInfoArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeGitProvider) GetRepoInfoReturns(result1 *gitprovider.RepositoryInfo, result2 error) {
	fake.getRepoInfoMutex.Lock()
	defer fake.getRepoInfoMutex.Unlock()
	fake.GetRepoInfoStub = nil
	fake.getRepoInfoReturns = struct {
		result1 *gitprovider.RepositoryInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeGitProvider) GetRepoInfoReturnsOnCall(i int, result1 *gitprovider.RepositoryInfo, result2 error) {
	fake.getRepoInfoMutex.Lock()
	defer fake.getRepoInfoMutex.Unlock()
	fake.GetRepoInfoStub = nil
	if fake.getRepoInfoReturnsOnCall == nil {
		fake.getRepoInfoReturnsOnCall = make(map[int]struct {
			result1 *gitprovider.RepositoryInfo
			result2 error
		})
	}
	fake.getRepoInfoReturnsOnCall[i] = struct {
		result1 *gitprovider.RepositoryInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeGitProvider) GetRepoInfoFromUrl(arg1 string) (*gitprovider.RepositoryInfo, error) {
	fake.getRepoInfoFromUrlMutex.Lock()
	ret, specificReturn := fake.getRepoInfoFromUrlReturnsOnCall[len(fake.getRepoInfoFromUrlArgsForCall)]
	fake.getRepoInfoFromUrlArgsForCall = append(fake.getRepoInfoFromUrlArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetRepoInfoFromUrlStub
	fakeReturns := fake.getRepoInfoFromUrlReturns
	fake.recordInvocation("GetRepoInfoFromUrl", []interface{}{arg1})
	fake.getRepoInfoFromUrlMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGitProvider) GetRepoInfoFromUrlCallCount() int {
	fake.getRepoInfoFromUrlMutex.RLock()
	defer fake.getRepoInfoFromUrlMutex.RUnlock()
	return len(fake.getRepoInfoFromUrlArgsForCall)
}

func (fake *FakeGitProvider) GetRepoInfoFromUrlCalls(stub func(string) (*gitprovider.RepositoryInfo, error)) {
	fake.getRepoInfoFromUrlMutex.Lock()
	defer fake.getRepoInfoFromUrlMutex.Unlock()
	fake.GetRepoInfoFromUrlStub = stub
}

func (fake *FakeGitProvider) GetRepoInfoFromUrlArgsForCall(i int) string {
	fake.getRepoInfoFromUrlMutex.RLock()
	defer fake.getRepoInfoFromUrlMutex.RUnlock()
	argsForCall := fake.getRepoInfoFromUrlArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeGitProvider) GetRepoInfoFromUrlReturns(result1 *gitprovider.RepositoryInfo, result2 error) {
	fake.getRepoInfoFromUrlMutex.Lock()
	defer fake.getRepoInfoFromUrlMutex.Unlock()
	fake.GetRepoInfoFromUrlStub = nil
	fake.getRepoInfoFromUrlReturns = struct {
		result1 *gitprovider.RepositoryInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeGitProvider) GetRepoInfoFromUrlReturnsOnCall(i int, result1 *gitprovider.RepositoryInfo, result2 error) {
	fake.getRepoInfoFromUrlMutex.Lock()
	defer fake.getRepoInfoFromUrlMutex.Unlock()
	fake.GetRepoInfoFromUrlStub = nil
	if fake.getRepoInfoFromUrlReturnsOnCall == nil {
		fake.getRepoInfoFromUrlReturnsOnCall = make(map[int]struct {
			result1 *gitprovider.RepositoryInfo
			result2 error
		})
	}
	fake.getRepoInfoFromUrlReturnsOnCall[i] = struct {
		result1 *gitprovider.RepositoryInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeGitProvider) GetRepoVisibility(arg1 string) (*gitprovider.RepositoryVisibility, error) {
	fake.getRepoVisibilityMutex.Lock()
	ret, specificReturn := fake.getRepoVisibilityReturnsOnCall[len(fake.getRepoVisibilityArgsForCall)]
	fake.getRepoVisibilityArgsForCall = append(fake.getRepoVisibilityArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetRepoVisibilityStub
	fakeReturns := fake.getRepoVisibilityReturns
	fake.recordInvocation("GetRepoVisibility", []interface{}{arg1})
	fake.getRepoVisibilityMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGitProvider) GetRepoVisibilityCallCount() int {
	fake.getRepoVisibilityMutex.RLock()
	defer fake.getRepoVisibilityMutex.RUnlock()
	return len(fake.getRepoVisibilityArgsForCall)
}

func (fake *FakeGitProvider) GetRepoVisibilityCalls(stub func(string) (*gitprovider.RepositoryVisibility, error)) {
	fake.getRepoVisibilityMutex.Lock()
	defer fake.getRepoVisibilityMutex.Unlock()
	fake.GetRepoVisibilityStub = stub
}

func (fake *FakeGitProvider) GetRepoVisibilityArgsForCall(i int) string {
	fake.getRepoVisibilityMutex.RLock()
	defer fake.getRepoVisibilityMutex.RUnlock()
	argsForCall := fake.getRepoVisibilityArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeGitProvider) GetRepoVisibilityReturns(result1 *gitprovider.RepositoryVisibility, result2 error) {
	fake.getRepoVisibilityMutex.Lock()
	defer fake.getRepoVisibilityMutex.Unlock()
	fake.GetRepoVisibilityStub = nil
	fake.getRepoVisibilityReturns = struct {
		result1 *gitprovider.RepositoryVisibility
		result2 error
	}{result1, result2}
}

func (fake *FakeGitProvider) GetRepoVisibilityReturnsOnCall(i int, result1 *gitprovider.RepositoryVisibility, result2 error) {
	fake.getRepoVisibilityMutex.Lock()
	defer fake.getRepoVisibilityMutex.Unlock()
	fake.GetRepoVisibilityStub = nil
	if fake.getRepoVisibilityReturnsOnCall == nil {
		fake.getRepoVisibilityReturnsOnCall = make(map[int]struct {
			result1 *gitprovider.RepositoryVisibility
			result2 error
		})
	}
	fake.getRepoVisibilityReturnsOnCall[i] = struct {
		result1 *gitprovider.RepositoryVisibility
		result2 error
	}{result1, result2}
}

func (fake *FakeGitProvider) RepositoryExists(arg1 string, arg2 string) (bool, error) {
	fake.repositoryExistsMutex.Lock()
	ret, specificReturn := fake.repositoryExistsReturnsOnCall[len(fake.repositoryExistsArgsForCall)]
	fake.repositoryExistsArgsForCall = append(fake.repositoryExistsArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.RepositoryExistsStub
	fakeReturns := fake.repositoryExistsReturns
	fake.recordInvocation("RepositoryExists", []interface{}{arg1, arg2})
	fake.repositoryExistsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGitProvider) RepositoryExistsCallCount() int {
	fake.repositoryExistsMutex.RLock()
	defer fake.repositoryExistsMutex.RUnlock()
	return len(fake.repositoryExistsArgsForCall)
}

func (fake *FakeGitProvider) RepositoryExistsCalls(stub func(string, string) (bool, error)) {
	fake.repositoryExistsMutex.Lock()
	defer fake.repositoryExistsMutex.Unlock()
	fake.RepositoryExistsStub = stub
}

func (fake *FakeGitProvider) RepositoryExistsArgsForCall(i int) (string, string) {
	fake.repositoryExistsMutex.RLock()
	defer fake.repositoryExistsMutex.RUnlock()
	argsForCall := fake.repositoryExistsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeGitProvider) RepositoryExistsReturns(result1 bool, result2 error) {
	fake.repositoryExistsMutex.Lock()
	defer fake.repositoryExistsMutex.Unlock()
	fake.RepositoryExistsStub = nil
	fake.repositoryExistsReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeGitProvider) RepositoryExistsReturnsOnCall(i int, result1 bool, result2 error) {
	fake.repositoryExistsMutex.Lock()
	defer fake.repositoryExistsMutex.Unlock()
	fake.RepositoryExistsStub = nil
	if fake.repositoryExistsReturnsOnCall == nil {
		fake.repositoryExistsReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.repositoryExistsReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeGitProvider) UploadDeployKey(arg1 string, arg2 string, arg3 []byte) error {
	var arg3Copy []byte
	if arg3 != nil {
		arg3Copy = make([]byte, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.uploadDeployKeyMutex.Lock()
	ret, specificReturn := fake.uploadDeployKeyReturnsOnCall[len(fake.uploadDeployKeyArgsForCall)]
	fake.uploadDeployKeyArgsForCall = append(fake.uploadDeployKeyArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 []byte
	}{arg1, arg2, arg3Copy})
	stub := fake.UploadDeployKeyStub
	fakeReturns := fake.uploadDeployKeyReturns
	fake.recordInvocation("UploadDeployKey", []interface{}{arg1, arg2, arg3Copy})
	fake.uploadDeployKeyMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeGitProvider) UploadDeployKeyCallCount() int {
	fake.uploadDeployKeyMutex.RLock()
	defer fake.uploadDeployKeyMutex.RUnlock()
	return len(fake.uploadDeployKeyArgsForCall)
}

func (fake *FakeGitProvider) UploadDeployKeyCalls(stub func(string, string, []byte) error) {
	fake.uploadDeployKeyMutex.Lock()
	defer fake.uploadDeployKeyMutex.Unlock()
	fake.UploadDeployKeyStub = stub
}

func (fake *FakeGitProvider) UploadDeployKeyArgsForCall(i int) (string, string, []byte) {
	fake.uploadDeployKeyMutex.RLock()
	defer fake.uploadDeployKeyMutex.RUnlock()
	argsForCall := fake.uploadDeployKeyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeGitProvider) UploadDeployKeyReturns(result1 error) {
	fake.uploadDeployKeyMutex.Lock()
	defer fake.uploadDeployKeyMutex.Unlock()
	fake.UploadDeployKeyStub = nil
	fake.uploadDeployKeyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeGitProvider) UploadDeployKeyReturnsOnCall(i int, result1 error) {
	fake.uploadDeployKeyMutex.Lock()
	defer fake.uploadDeployKeyMutex.Unlock()
	fake.UploadDeployKeyStub = nil
	if fake.uploadDeployKeyReturnsOnCall == nil {
		fake.uploadDeployKeyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.uploadDeployKeyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeGitProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createPullRequestToOrgRepoMutex.RLock()
	defer fake.createPullRequestToOrgRepoMutex.RUnlock()
	fake.createPullRequestToUserRepoMutex.RLock()
	defer fake.createPullRequestToUserRepoMutex.RUnlock()
	fake.createRepositoryMutex.RLock()
	defer fake.createRepositoryMutex.RUnlock()
	fake.deployKeyExistsMutex.RLock()
	defer fake.deployKeyExistsMutex.RUnlock()
	fake.getAccountTypeMutex.RLock()
	defer fake.getAccountTypeMutex.RUnlock()
	fake.getCommitsFromOrgRepoMutex.RLock()
	defer fake.getCommitsFromOrgRepoMutex.RUnlock()
	fake.getCommitsFromUserRepoMutex.RLock()
	defer fake.getCommitsFromUserRepoMutex.RUnlock()
	fake.getDefaultBranchMutex.RLock()
	defer fake.getDefaultBranchMutex.RUnlock()
	fake.getRepoInfoMutex.RLock()
	defer fake.getRepoInfoMutex.RUnlock()
	fake.getRepoInfoFromUrlMutex.RLock()
	defer fake.getRepoInfoFromUrlMutex.RUnlock()
	fake.getRepoVisibilityMutex.RLock()
	defer fake.getRepoVisibilityMutex.RUnlock()
	fake.repositoryExistsMutex.RLock()
	defer fake.repositoryExistsMutex.RUnlock()
	fake.uploadDeployKeyMutex.RLock()
	defer fake.uploadDeployKeyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGitProvider) recordInvocation(key string, args []interface{}) {
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

var _ gitproviders.GitProvider = new(FakeGitProvider)
