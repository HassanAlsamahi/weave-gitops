// Code generated by counterfeiter. DO NOT EDIT.
package kubefakes

import (
	"context"
	"sync"

	"github.com/weaveworks/weave-gitops/api/v1alpha1"
	"github.com/weaveworks/weave-gitops/pkg/kube"
	"k8s.io/apimachinery/pkg/types"
)

type FakeKube struct {
	ApplyStub        func([]byte, string) ([]byte, error)
	applyMutex       sync.RWMutex
	applyArgsForCall []struct {
		arg1 []byte
		arg2 string
	}
	applyReturns struct {
		result1 []byte
		result2 error
	}
	applyReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	DeleteStub        func([]byte, string) ([]byte, error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 []byte
		arg2 string
	}
	deleteReturns struct {
		result1 []byte
		result2 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	DeleteByNameStub        func(string, string, string) ([]byte, error)
	deleteByNameMutex       sync.RWMutex
	deleteByNameArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	deleteByNameReturns struct {
		result1 []byte
		result2 error
	}
	deleteByNameReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	FluxPresentStub        func(context.Context) (bool, error)
	fluxPresentMutex       sync.RWMutex
	fluxPresentArgsForCall []struct {
		arg1 context.Context
	}
	fluxPresentReturns struct {
		result1 bool
		result2 error
	}
	fluxPresentReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	GetApplicationStub        func(context.Context, types.NamespacedName) (*v1alpha1.Application, error)
	getApplicationMutex       sync.RWMutex
	getApplicationArgsForCall []struct {
		arg1 context.Context
		arg2 types.NamespacedName
	}
	getApplicationReturns struct {
		result1 *v1alpha1.Application
		result2 error
	}
	getApplicationReturnsOnCall map[int]struct {
		result1 *v1alpha1.Application
		result2 error
	}
	GetApplicationsStub        func(context.Context, string) ([]v1alpha1.Application, error)
	getApplicationsMutex       sync.RWMutex
	getApplicationsArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getApplicationsReturns struct {
		result1 []v1alpha1.Application
		result2 error
	}
	getApplicationsReturnsOnCall map[int]struct {
		result1 []v1alpha1.Application
		result2 error
	}
	GetClusterNameStub        func(context.Context) (string, error)
	getClusterNameMutex       sync.RWMutex
	getClusterNameArgsForCall []struct {
		arg1 context.Context
	}
	getClusterNameReturns struct {
		result1 string
		result2 error
	}
	getClusterNameReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetClusterStatusStub        func(context.Context) kube.ClusterStatus
	getClusterStatusMutex       sync.RWMutex
	getClusterStatusArgsForCall []struct {
		arg1 context.Context
	}
	getClusterStatusReturns struct {
		result1 kube.ClusterStatus
	}
	getClusterStatusReturnsOnCall map[int]struct {
		result1 kube.ClusterStatus
	}
	GetResourceStub        func(context.Context, types.NamespacedName, kube.Resource) error
	getResourceMutex       sync.RWMutex
	getResourceArgsForCall []struct {
		arg1 context.Context
		arg2 types.NamespacedName
		arg3 kube.Resource
	}
	getResourceReturns struct {
		result1 error
	}
	getResourceReturnsOnCall map[int]struct {
		result1 error
	}
	LabelExistsInClusterStub        func(context.Context, string) (bool, error)
	labelExistsInClusterMutex       sync.RWMutex
	labelExistsInClusterArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	labelExistsInClusterReturns struct {
		result1 bool
		result2 error
	}
	labelExistsInClusterReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	SecretPresentStub        func(context.Context, string, string) (bool, error)
	secretPresentMutex       sync.RWMutex
	secretPresentArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}
	secretPresentReturns struct {
		result1 bool
		result2 error
	}
	secretPresentReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeKube) Apply(arg1 []byte, arg2 string) ([]byte, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.applyMutex.Lock()
	ret, specificReturn := fake.applyReturnsOnCall[len(fake.applyArgsForCall)]
	fake.applyArgsForCall = append(fake.applyArgsForCall, struct {
		arg1 []byte
		arg2 string
	}{arg1Copy, arg2})
	stub := fake.ApplyStub
	fakeReturns := fake.applyReturns
	fake.recordInvocation("Apply", []interface{}{arg1Copy, arg2})
	fake.applyMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeKube) ApplyCallCount() int {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	return len(fake.applyArgsForCall)
}

func (fake *FakeKube) ApplyCalls(stub func([]byte, string) ([]byte, error)) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = stub
}

func (fake *FakeKube) ApplyArgsForCall(i int) ([]byte, string) {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	argsForCall := fake.applyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeKube) ApplyReturns(result1 []byte, result2 error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = nil
	fake.applyReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) ApplyReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = nil
	if fake.applyReturnsOnCall == nil {
		fake.applyReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.applyReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) Delete(arg1 []byte, arg2 string) ([]byte, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 []byte
		arg2 string
	}{arg1Copy, arg2})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1Copy, arg2})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeKube) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeKube) DeleteCalls(stub func([]byte, string) ([]byte, error)) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeKube) DeleteArgsForCall(i int) ([]byte, string) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeKube) DeleteReturns(result1 []byte, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) DeleteReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) DeleteByName(arg1 string, arg2 string, arg3 string) ([]byte, error) {
	fake.deleteByNameMutex.Lock()
	ret, specificReturn := fake.deleteByNameReturnsOnCall[len(fake.deleteByNameArgsForCall)]
	fake.deleteByNameArgsForCall = append(fake.deleteByNameArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.DeleteByNameStub
	fakeReturns := fake.deleteByNameReturns
	fake.recordInvocation("DeleteByName", []interface{}{arg1, arg2, arg3})
	fake.deleteByNameMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeKube) DeleteByNameCallCount() int {
	fake.deleteByNameMutex.RLock()
	defer fake.deleteByNameMutex.RUnlock()
	return len(fake.deleteByNameArgsForCall)
}

func (fake *FakeKube) DeleteByNameCalls(stub func(string, string, string) ([]byte, error)) {
	fake.deleteByNameMutex.Lock()
	defer fake.deleteByNameMutex.Unlock()
	fake.DeleteByNameStub = stub
}

func (fake *FakeKube) DeleteByNameArgsForCall(i int) (string, string, string) {
	fake.deleteByNameMutex.RLock()
	defer fake.deleteByNameMutex.RUnlock()
	argsForCall := fake.deleteByNameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeKube) DeleteByNameReturns(result1 []byte, result2 error) {
	fake.deleteByNameMutex.Lock()
	defer fake.deleteByNameMutex.Unlock()
	fake.DeleteByNameStub = nil
	fake.deleteByNameReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) DeleteByNameReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.deleteByNameMutex.Lock()
	defer fake.deleteByNameMutex.Unlock()
	fake.DeleteByNameStub = nil
	if fake.deleteByNameReturnsOnCall == nil {
		fake.deleteByNameReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.deleteByNameReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) FluxPresent(arg1 context.Context) (bool, error) {
	fake.fluxPresentMutex.Lock()
	ret, specificReturn := fake.fluxPresentReturnsOnCall[len(fake.fluxPresentArgsForCall)]
	fake.fluxPresentArgsForCall = append(fake.fluxPresentArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.FluxPresentStub
	fakeReturns := fake.fluxPresentReturns
	fake.recordInvocation("FluxPresent", []interface{}{arg1})
	fake.fluxPresentMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeKube) FluxPresentCallCount() int {
	fake.fluxPresentMutex.RLock()
	defer fake.fluxPresentMutex.RUnlock()
	return len(fake.fluxPresentArgsForCall)
}

func (fake *FakeKube) FluxPresentCalls(stub func(context.Context) (bool, error)) {
	fake.fluxPresentMutex.Lock()
	defer fake.fluxPresentMutex.Unlock()
	fake.FluxPresentStub = stub
}

func (fake *FakeKube) FluxPresentArgsForCall(i int) context.Context {
	fake.fluxPresentMutex.RLock()
	defer fake.fluxPresentMutex.RUnlock()
	argsForCall := fake.fluxPresentArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeKube) FluxPresentReturns(result1 bool, result2 error) {
	fake.fluxPresentMutex.Lock()
	defer fake.fluxPresentMutex.Unlock()
	fake.FluxPresentStub = nil
	fake.fluxPresentReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) FluxPresentReturnsOnCall(i int, result1 bool, result2 error) {
	fake.fluxPresentMutex.Lock()
	defer fake.fluxPresentMutex.Unlock()
	fake.FluxPresentStub = nil
	if fake.fluxPresentReturnsOnCall == nil {
		fake.fluxPresentReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.fluxPresentReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) GetApplication(arg1 context.Context, arg2 types.NamespacedName) (*v1alpha1.Application, error) {
	fake.getApplicationMutex.Lock()
	ret, specificReturn := fake.getApplicationReturnsOnCall[len(fake.getApplicationArgsForCall)]
	fake.getApplicationArgsForCall = append(fake.getApplicationArgsForCall, struct {
		arg1 context.Context
		arg2 types.NamespacedName
	}{arg1, arg2})
	stub := fake.GetApplicationStub
	fakeReturns := fake.getApplicationReturns
	fake.recordInvocation("GetApplication", []interface{}{arg1, arg2})
	fake.getApplicationMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeKube) GetApplicationCallCount() int {
	fake.getApplicationMutex.RLock()
	defer fake.getApplicationMutex.RUnlock()
	return len(fake.getApplicationArgsForCall)
}

func (fake *FakeKube) GetApplicationCalls(stub func(context.Context, types.NamespacedName) (*v1alpha1.Application, error)) {
	fake.getApplicationMutex.Lock()
	defer fake.getApplicationMutex.Unlock()
	fake.GetApplicationStub = stub
}

func (fake *FakeKube) GetApplicationArgsForCall(i int) (context.Context, types.NamespacedName) {
	fake.getApplicationMutex.RLock()
	defer fake.getApplicationMutex.RUnlock()
	argsForCall := fake.getApplicationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeKube) GetApplicationReturns(result1 *v1alpha1.Application, result2 error) {
	fake.getApplicationMutex.Lock()
	defer fake.getApplicationMutex.Unlock()
	fake.GetApplicationStub = nil
	fake.getApplicationReturns = struct {
		result1 *v1alpha1.Application
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) GetApplicationReturnsOnCall(i int, result1 *v1alpha1.Application, result2 error) {
	fake.getApplicationMutex.Lock()
	defer fake.getApplicationMutex.Unlock()
	fake.GetApplicationStub = nil
	if fake.getApplicationReturnsOnCall == nil {
		fake.getApplicationReturnsOnCall = make(map[int]struct {
			result1 *v1alpha1.Application
			result2 error
		})
	}
	fake.getApplicationReturnsOnCall[i] = struct {
		result1 *v1alpha1.Application
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) GetApplications(arg1 context.Context, arg2 string) ([]v1alpha1.Application, error) {
	fake.getApplicationsMutex.Lock()
	ret, specificReturn := fake.getApplicationsReturnsOnCall[len(fake.getApplicationsArgsForCall)]
	fake.getApplicationsArgsForCall = append(fake.getApplicationsArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.GetApplicationsStub
	fakeReturns := fake.getApplicationsReturns
	fake.recordInvocation("GetApplications", []interface{}{arg1, arg2})
	fake.getApplicationsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeKube) GetApplicationsCallCount() int {
	fake.getApplicationsMutex.RLock()
	defer fake.getApplicationsMutex.RUnlock()
	return len(fake.getApplicationsArgsForCall)
}

func (fake *FakeKube) GetApplicationsCalls(stub func(context.Context, string) ([]v1alpha1.Application, error)) {
	fake.getApplicationsMutex.Lock()
	defer fake.getApplicationsMutex.Unlock()
	fake.GetApplicationsStub = stub
}

func (fake *FakeKube) GetApplicationsArgsForCall(i int) (context.Context, string) {
	fake.getApplicationsMutex.RLock()
	defer fake.getApplicationsMutex.RUnlock()
	argsForCall := fake.getApplicationsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeKube) GetApplicationsReturns(result1 []v1alpha1.Application, result2 error) {
	fake.getApplicationsMutex.Lock()
	defer fake.getApplicationsMutex.Unlock()
	fake.GetApplicationsStub = nil
	fake.getApplicationsReturns = struct {
		result1 []v1alpha1.Application
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) GetApplicationsReturnsOnCall(i int, result1 []v1alpha1.Application, result2 error) {
	fake.getApplicationsMutex.Lock()
	defer fake.getApplicationsMutex.Unlock()
	fake.GetApplicationsStub = nil
	if fake.getApplicationsReturnsOnCall == nil {
		fake.getApplicationsReturnsOnCall = make(map[int]struct {
			result1 []v1alpha1.Application
			result2 error
		})
	}
	fake.getApplicationsReturnsOnCall[i] = struct {
		result1 []v1alpha1.Application
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) GetClusterName(arg1 context.Context) (string, error) {
	fake.getClusterNameMutex.Lock()
	ret, specificReturn := fake.getClusterNameReturnsOnCall[len(fake.getClusterNameArgsForCall)]
	fake.getClusterNameArgsForCall = append(fake.getClusterNameArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.GetClusterNameStub
	fakeReturns := fake.getClusterNameReturns
	fake.recordInvocation("GetClusterName", []interface{}{arg1})
	fake.getClusterNameMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeKube) GetClusterNameCallCount() int {
	fake.getClusterNameMutex.RLock()
	defer fake.getClusterNameMutex.RUnlock()
	return len(fake.getClusterNameArgsForCall)
}

func (fake *FakeKube) GetClusterNameCalls(stub func(context.Context) (string, error)) {
	fake.getClusterNameMutex.Lock()
	defer fake.getClusterNameMutex.Unlock()
	fake.GetClusterNameStub = stub
}

func (fake *FakeKube) GetClusterNameArgsForCall(i int) context.Context {
	fake.getClusterNameMutex.RLock()
	defer fake.getClusterNameMutex.RUnlock()
	argsForCall := fake.getClusterNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeKube) GetClusterNameReturns(result1 string, result2 error) {
	fake.getClusterNameMutex.Lock()
	defer fake.getClusterNameMutex.Unlock()
	fake.GetClusterNameStub = nil
	fake.getClusterNameReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) GetClusterNameReturnsOnCall(i int, result1 string, result2 error) {
	fake.getClusterNameMutex.Lock()
	defer fake.getClusterNameMutex.Unlock()
	fake.GetClusterNameStub = nil
	if fake.getClusterNameReturnsOnCall == nil {
		fake.getClusterNameReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getClusterNameReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) GetClusterStatus(arg1 context.Context) kube.ClusterStatus {
	fake.getClusterStatusMutex.Lock()
	ret, specificReturn := fake.getClusterStatusReturnsOnCall[len(fake.getClusterStatusArgsForCall)]
	fake.getClusterStatusArgsForCall = append(fake.getClusterStatusArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.GetClusterStatusStub
	fakeReturns := fake.getClusterStatusReturns
	fake.recordInvocation("GetClusterStatus", []interface{}{arg1})
	fake.getClusterStatusMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeKube) GetClusterStatusCallCount() int {
	fake.getClusterStatusMutex.RLock()
	defer fake.getClusterStatusMutex.RUnlock()
	return len(fake.getClusterStatusArgsForCall)
}

func (fake *FakeKube) GetClusterStatusCalls(stub func(context.Context) kube.ClusterStatus) {
	fake.getClusterStatusMutex.Lock()
	defer fake.getClusterStatusMutex.Unlock()
	fake.GetClusterStatusStub = stub
}

func (fake *FakeKube) GetClusterStatusArgsForCall(i int) context.Context {
	fake.getClusterStatusMutex.RLock()
	defer fake.getClusterStatusMutex.RUnlock()
	argsForCall := fake.getClusterStatusArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeKube) GetClusterStatusReturns(result1 kube.ClusterStatus) {
	fake.getClusterStatusMutex.Lock()
	defer fake.getClusterStatusMutex.Unlock()
	fake.GetClusterStatusStub = nil
	fake.getClusterStatusReturns = struct {
		result1 kube.ClusterStatus
	}{result1}
}

func (fake *FakeKube) GetClusterStatusReturnsOnCall(i int, result1 kube.ClusterStatus) {
	fake.getClusterStatusMutex.Lock()
	defer fake.getClusterStatusMutex.Unlock()
	fake.GetClusterStatusStub = nil
	if fake.getClusterStatusReturnsOnCall == nil {
		fake.getClusterStatusReturnsOnCall = make(map[int]struct {
			result1 kube.ClusterStatus
		})
	}
	fake.getClusterStatusReturnsOnCall[i] = struct {
		result1 kube.ClusterStatus
	}{result1}
}

func (fake *FakeKube) GetResource(arg1 context.Context, arg2 types.NamespacedName, arg3 kube.Resource) error {
	fake.getResourceMutex.Lock()
	ret, specificReturn := fake.getResourceReturnsOnCall[len(fake.getResourceArgsForCall)]
	fake.getResourceArgsForCall = append(fake.getResourceArgsForCall, struct {
		arg1 context.Context
		arg2 types.NamespacedName
		arg3 kube.Resource
	}{arg1, arg2, arg3})
	stub := fake.GetResourceStub
	fakeReturns := fake.getResourceReturns
	fake.recordInvocation("GetResource", []interface{}{arg1, arg2, arg3})
	fake.getResourceMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeKube) GetResourceCallCount() int {
	fake.getResourceMutex.RLock()
	defer fake.getResourceMutex.RUnlock()
	return len(fake.getResourceArgsForCall)
}

func (fake *FakeKube) GetResourceCalls(stub func(context.Context, types.NamespacedName, kube.Resource) error) {
	fake.getResourceMutex.Lock()
	defer fake.getResourceMutex.Unlock()
	fake.GetResourceStub = stub
}

func (fake *FakeKube) GetResourceArgsForCall(i int) (context.Context, types.NamespacedName, kube.Resource) {
	fake.getResourceMutex.RLock()
	defer fake.getResourceMutex.RUnlock()
	argsForCall := fake.getResourceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeKube) GetResourceReturns(result1 error) {
	fake.getResourceMutex.Lock()
	defer fake.getResourceMutex.Unlock()
	fake.GetResourceStub = nil
	fake.getResourceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeKube) GetResourceReturnsOnCall(i int, result1 error) {
	fake.getResourceMutex.Lock()
	defer fake.getResourceMutex.Unlock()
	fake.GetResourceStub = nil
	if fake.getResourceReturnsOnCall == nil {
		fake.getResourceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.getResourceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeKube) LabelExistsInCluster(arg1 context.Context, arg2 string) (bool, error) {
	fake.labelExistsInClusterMutex.Lock()
	ret, specificReturn := fake.labelExistsInClusterReturnsOnCall[len(fake.labelExistsInClusterArgsForCall)]
	fake.labelExistsInClusterArgsForCall = append(fake.labelExistsInClusterArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.LabelExistsInClusterStub
	fakeReturns := fake.labelExistsInClusterReturns
	fake.recordInvocation("LabelExistsInCluster", []interface{}{arg1, arg2})
	fake.labelExistsInClusterMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeKube) LabelExistsInClusterCallCount() int {
	fake.labelExistsInClusterMutex.RLock()
	defer fake.labelExistsInClusterMutex.RUnlock()
	return len(fake.labelExistsInClusterArgsForCall)
}

func (fake *FakeKube) LabelExistsInClusterCalls(stub func(context.Context, string) (bool, error)) {
	fake.labelExistsInClusterMutex.Lock()
	defer fake.labelExistsInClusterMutex.Unlock()
	fake.LabelExistsInClusterStub = stub
}

func (fake *FakeKube) LabelExistsInClusterArgsForCall(i int) (context.Context, string) {
	fake.labelExistsInClusterMutex.RLock()
	defer fake.labelExistsInClusterMutex.RUnlock()
	argsForCall := fake.labelExistsInClusterArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeKube) LabelExistsInClusterReturns(result1 bool, result2 error) {
	fake.labelExistsInClusterMutex.Lock()
	defer fake.labelExistsInClusterMutex.Unlock()
	fake.LabelExistsInClusterStub = nil
	fake.labelExistsInClusterReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) LabelExistsInClusterReturnsOnCall(i int, result1 bool, result2 error) {
	fake.labelExistsInClusterMutex.Lock()
	defer fake.labelExistsInClusterMutex.Unlock()
	fake.LabelExistsInClusterStub = nil
	if fake.labelExistsInClusterReturnsOnCall == nil {
		fake.labelExistsInClusterReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.labelExistsInClusterReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) SecretPresent(arg1 context.Context, arg2 string, arg3 string) (bool, error) {
	fake.secretPresentMutex.Lock()
	ret, specificReturn := fake.secretPresentReturnsOnCall[len(fake.secretPresentArgsForCall)]
	fake.secretPresentArgsForCall = append(fake.secretPresentArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.SecretPresentStub
	fakeReturns := fake.secretPresentReturns
	fake.recordInvocation("SecretPresent", []interface{}{arg1, arg2, arg3})
	fake.secretPresentMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeKube) SecretPresentCallCount() int {
	fake.secretPresentMutex.RLock()
	defer fake.secretPresentMutex.RUnlock()
	return len(fake.secretPresentArgsForCall)
}

func (fake *FakeKube) SecretPresentCalls(stub func(context.Context, string, string) (bool, error)) {
	fake.secretPresentMutex.Lock()
	defer fake.secretPresentMutex.Unlock()
	fake.SecretPresentStub = stub
}

func (fake *FakeKube) SecretPresentArgsForCall(i int) (context.Context, string, string) {
	fake.secretPresentMutex.RLock()
	defer fake.secretPresentMutex.RUnlock()
	argsForCall := fake.secretPresentArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeKube) SecretPresentReturns(result1 bool, result2 error) {
	fake.secretPresentMutex.Lock()
	defer fake.secretPresentMutex.Unlock()
	fake.SecretPresentStub = nil
	fake.secretPresentReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) SecretPresentReturnsOnCall(i int, result1 bool, result2 error) {
	fake.secretPresentMutex.Lock()
	defer fake.secretPresentMutex.Unlock()
	fake.SecretPresentStub = nil
	if fake.secretPresentReturnsOnCall == nil {
		fake.secretPresentReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.secretPresentReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.deleteByNameMutex.RLock()
	defer fake.deleteByNameMutex.RUnlock()
	fake.fluxPresentMutex.RLock()
	defer fake.fluxPresentMutex.RUnlock()
	fake.getApplicationMutex.RLock()
	defer fake.getApplicationMutex.RUnlock()
	fake.getApplicationsMutex.RLock()
	defer fake.getApplicationsMutex.RUnlock()
	fake.getClusterNameMutex.RLock()
	defer fake.getClusterNameMutex.RUnlock()
	fake.getClusterStatusMutex.RLock()
	defer fake.getClusterStatusMutex.RUnlock()
	fake.getResourceMutex.RLock()
	defer fake.getResourceMutex.RUnlock()
	fake.labelExistsInClusterMutex.RLock()
	defer fake.labelExistsInClusterMutex.RUnlock()
	fake.secretPresentMutex.RLock()
	defer fake.secretPresentMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeKube) recordInvocation(key string, args []interface{}) {
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

var _ kube.Kube = new(FakeKube)
