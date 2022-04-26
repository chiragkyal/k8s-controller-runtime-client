// Code generated by counterfeiter. DO NOT EDIT.
package clientfakes

import (
	"context"
	"sync"

	"github.com/chiragkyal/k8s-controller-runtime-client/client"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	clienta "sigs.k8s.io/controller-runtime/pkg/client"
)

type FakeK8sRestClient struct {
	CreateStub        func(context.Context, clienta.Object, ...clienta.CreateOption) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 context.Context
		arg2 clienta.Object
		arg3 []clienta.CreateOption
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func(context.Context, clienta.Object, ...clienta.DeleteOption) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 context.Context
		arg2 clienta.Object
		arg3 []clienta.DeleteOption
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteAllOfStub        func(context.Context, clienta.Object, ...clienta.DeleteAllOfOption) error
	deleteAllOfMutex       sync.RWMutex
	deleteAllOfArgsForCall []struct {
		arg1 context.Context
		arg2 clienta.Object
		arg3 []clienta.DeleteAllOfOption
	}
	deleteAllOfReturns struct {
		result1 error
	}
	deleteAllOfReturnsOnCall map[int]struct {
		result1 error
	}
	GetStub        func(context.Context, types.NamespacedName, clienta.Object) error
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 context.Context
		arg2 types.NamespacedName
		arg3 clienta.Object
	}
	getReturns struct {
		result1 error
	}
	getReturnsOnCall map[int]struct {
		result1 error
	}
	ListStub        func(context.Context, clienta.ObjectList, ...clienta.ListOption) error
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 context.Context
		arg2 clienta.ObjectList
		arg3 []clienta.ListOption
	}
	listReturns struct {
		result1 error
	}
	listReturnsOnCall map[int]struct {
		result1 error
	}
	PatchStub        func(context.Context, clienta.Object, clienta.Patch, ...clienta.PatchOption) error
	patchMutex       sync.RWMutex
	patchArgsForCall []struct {
		arg1 context.Context
		arg2 clienta.Object
		arg3 clienta.Patch
		arg4 []clienta.PatchOption
	}
	patchReturns struct {
		result1 error
	}
	patchReturnsOnCall map[int]struct {
		result1 error
	}
	RESTMapperStub        func() meta.RESTMapper
	rESTMapperMutex       sync.RWMutex
	rESTMapperArgsForCall []struct {
	}
	rESTMapperReturns struct {
		result1 meta.RESTMapper
	}
	rESTMapperReturnsOnCall map[int]struct {
		result1 meta.RESTMapper
	}
	SchemeStub        func() *runtime.Scheme
	schemeMutex       sync.RWMutex
	schemeArgsForCall []struct {
	}
	schemeReturns struct {
		result1 *runtime.Scheme
	}
	schemeReturnsOnCall map[int]struct {
		result1 *runtime.Scheme
	}
	StatusStub        func() clienta.StatusWriter
	statusMutex       sync.RWMutex
	statusArgsForCall []struct {
	}
	statusReturns struct {
		result1 clienta.StatusWriter
	}
	statusReturnsOnCall map[int]struct {
		result1 clienta.StatusWriter
	}
	UpdateStub        func(context.Context, clienta.Object, ...clienta.UpdateOption) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 context.Context
		arg2 clienta.Object
		arg3 []clienta.UpdateOption
	}
	updateReturns struct {
		result1 error
	}
	updateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeK8sRestClient) Create(arg1 context.Context, arg2 clienta.Object, arg3 ...clienta.CreateOption) error {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 context.Context
		arg2 clienta.Object
		arg3 []clienta.CreateOption
	}{arg1, arg2, arg3})
	stub := fake.CreateStub
	fakeReturns := fake.createReturns
	fake.recordInvocation("Create", []interface{}{arg1, arg2, arg3})
	fake.createMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeK8sRestClient) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeK8sRestClient) CreateCalls(stub func(context.Context, clienta.Object, ...clienta.CreateOption) error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeK8sRestClient) CreateArgsForCall(i int) (context.Context, clienta.Object, []clienta.CreateOption) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeK8sRestClient) CreateReturns(result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeK8sRestClient) CreateReturnsOnCall(i int, result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeK8sRestClient) Delete(arg1 context.Context, arg2 clienta.Object, arg3 ...clienta.DeleteOption) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 context.Context
		arg2 clienta.Object
		arg3 []clienta.DeleteOption
	}{arg1, arg2, arg3})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1, arg2, arg3})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeK8sRestClient) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeK8sRestClient) DeleteCalls(stub func(context.Context, clienta.Object, ...clienta.DeleteOption) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeK8sRestClient) DeleteArgsForCall(i int) (context.Context, clienta.Object, []clienta.DeleteOption) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeK8sRestClient) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeK8sRestClient) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeK8sRestClient) DeleteAllOf(arg1 context.Context, arg2 clienta.Object, arg3 ...clienta.DeleteAllOfOption) error {
	fake.deleteAllOfMutex.Lock()
	ret, specificReturn := fake.deleteAllOfReturnsOnCall[len(fake.deleteAllOfArgsForCall)]
	fake.deleteAllOfArgsForCall = append(fake.deleteAllOfArgsForCall, struct {
		arg1 context.Context
		arg2 clienta.Object
		arg3 []clienta.DeleteAllOfOption
	}{arg1, arg2, arg3})
	stub := fake.DeleteAllOfStub
	fakeReturns := fake.deleteAllOfReturns
	fake.recordInvocation("DeleteAllOf", []interface{}{arg1, arg2, arg3})
	fake.deleteAllOfMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeK8sRestClient) DeleteAllOfCallCount() int {
	fake.deleteAllOfMutex.RLock()
	defer fake.deleteAllOfMutex.RUnlock()
	return len(fake.deleteAllOfArgsForCall)
}

func (fake *FakeK8sRestClient) DeleteAllOfCalls(stub func(context.Context, clienta.Object, ...clienta.DeleteAllOfOption) error) {
	fake.deleteAllOfMutex.Lock()
	defer fake.deleteAllOfMutex.Unlock()
	fake.DeleteAllOfStub = stub
}

func (fake *FakeK8sRestClient) DeleteAllOfArgsForCall(i int) (context.Context, clienta.Object, []clienta.DeleteAllOfOption) {
	fake.deleteAllOfMutex.RLock()
	defer fake.deleteAllOfMutex.RUnlock()
	argsForCall := fake.deleteAllOfArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeK8sRestClient) DeleteAllOfReturns(result1 error) {
	fake.deleteAllOfMutex.Lock()
	defer fake.deleteAllOfMutex.Unlock()
	fake.DeleteAllOfStub = nil
	fake.deleteAllOfReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeK8sRestClient) DeleteAllOfReturnsOnCall(i int, result1 error) {
	fake.deleteAllOfMutex.Lock()
	defer fake.deleteAllOfMutex.Unlock()
	fake.DeleteAllOfStub = nil
	if fake.deleteAllOfReturnsOnCall == nil {
		fake.deleteAllOfReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteAllOfReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeK8sRestClient) Get(arg1 context.Context, arg2 types.NamespacedName, arg3 clienta.Object) error {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 context.Context
		arg2 types.NamespacedName
		arg3 clienta.Object
	}{arg1, arg2, arg3})
	stub := fake.GetStub
	fakeReturns := fake.getReturns
	fake.recordInvocation("Get", []interface{}{arg1, arg2, arg3})
	fake.getMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeK8sRestClient) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeK8sRestClient) GetCalls(stub func(context.Context, types.NamespacedName, clienta.Object) error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeK8sRestClient) GetArgsForCall(i int) (context.Context, types.NamespacedName, clienta.Object) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeK8sRestClient) GetReturns(result1 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeK8sRestClient) GetReturnsOnCall(i int, result1 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeK8sRestClient) List(arg1 context.Context, arg2 clienta.ObjectList, arg3 ...clienta.ListOption) error {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 context.Context
		arg2 clienta.ObjectList
		arg3 []clienta.ListOption
	}{arg1, arg2, arg3})
	stub := fake.ListStub
	fakeReturns := fake.listReturns
	fake.recordInvocation("List", []interface{}{arg1, arg2, arg3})
	fake.listMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeK8sRestClient) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeK8sRestClient) ListCalls(stub func(context.Context, clienta.ObjectList, ...clienta.ListOption) error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakeK8sRestClient) ListArgsForCall(i int) (context.Context, clienta.ObjectList, []clienta.ListOption) {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	argsForCall := fake.listArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeK8sRestClient) ListReturns(result1 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeK8sRestClient) ListReturnsOnCall(i int, result1 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeK8sRestClient) Patch(arg1 context.Context, arg2 clienta.Object, arg3 clienta.Patch, arg4 ...clienta.PatchOption) error {
	fake.patchMutex.Lock()
	ret, specificReturn := fake.patchReturnsOnCall[len(fake.patchArgsForCall)]
	fake.patchArgsForCall = append(fake.patchArgsForCall, struct {
		arg1 context.Context
		arg2 clienta.Object
		arg3 clienta.Patch
		arg4 []clienta.PatchOption
	}{arg1, arg2, arg3, arg4})
	stub := fake.PatchStub
	fakeReturns := fake.patchReturns
	fake.recordInvocation("Patch", []interface{}{arg1, arg2, arg3, arg4})
	fake.patchMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeK8sRestClient) PatchCallCount() int {
	fake.patchMutex.RLock()
	defer fake.patchMutex.RUnlock()
	return len(fake.patchArgsForCall)
}

func (fake *FakeK8sRestClient) PatchCalls(stub func(context.Context, clienta.Object, clienta.Patch, ...clienta.PatchOption) error) {
	fake.patchMutex.Lock()
	defer fake.patchMutex.Unlock()
	fake.PatchStub = stub
}

func (fake *FakeK8sRestClient) PatchArgsForCall(i int) (context.Context, clienta.Object, clienta.Patch, []clienta.PatchOption) {
	fake.patchMutex.RLock()
	defer fake.patchMutex.RUnlock()
	argsForCall := fake.patchArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeK8sRestClient) PatchReturns(result1 error) {
	fake.patchMutex.Lock()
	defer fake.patchMutex.Unlock()
	fake.PatchStub = nil
	fake.patchReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeK8sRestClient) PatchReturnsOnCall(i int, result1 error) {
	fake.patchMutex.Lock()
	defer fake.patchMutex.Unlock()
	fake.PatchStub = nil
	if fake.patchReturnsOnCall == nil {
		fake.patchReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.patchReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeK8sRestClient) RESTMapper() meta.RESTMapper {
	fake.rESTMapperMutex.Lock()
	ret, specificReturn := fake.rESTMapperReturnsOnCall[len(fake.rESTMapperArgsForCall)]
	fake.rESTMapperArgsForCall = append(fake.rESTMapperArgsForCall, struct {
	}{})
	stub := fake.RESTMapperStub
	fakeReturns := fake.rESTMapperReturns
	fake.recordInvocation("RESTMapper", []interface{}{})
	fake.rESTMapperMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeK8sRestClient) RESTMapperCallCount() int {
	fake.rESTMapperMutex.RLock()
	defer fake.rESTMapperMutex.RUnlock()
	return len(fake.rESTMapperArgsForCall)
}

func (fake *FakeK8sRestClient) RESTMapperCalls(stub func() meta.RESTMapper) {
	fake.rESTMapperMutex.Lock()
	defer fake.rESTMapperMutex.Unlock()
	fake.RESTMapperStub = stub
}

func (fake *FakeK8sRestClient) RESTMapperReturns(result1 meta.RESTMapper) {
	fake.rESTMapperMutex.Lock()
	defer fake.rESTMapperMutex.Unlock()
	fake.RESTMapperStub = nil
	fake.rESTMapperReturns = struct {
		result1 meta.RESTMapper
	}{result1}
}

func (fake *FakeK8sRestClient) RESTMapperReturnsOnCall(i int, result1 meta.RESTMapper) {
	fake.rESTMapperMutex.Lock()
	defer fake.rESTMapperMutex.Unlock()
	fake.RESTMapperStub = nil
	if fake.rESTMapperReturnsOnCall == nil {
		fake.rESTMapperReturnsOnCall = make(map[int]struct {
			result1 meta.RESTMapper
		})
	}
	fake.rESTMapperReturnsOnCall[i] = struct {
		result1 meta.RESTMapper
	}{result1}
}

func (fake *FakeK8sRestClient) Scheme() *runtime.Scheme {
	fake.schemeMutex.Lock()
	ret, specificReturn := fake.schemeReturnsOnCall[len(fake.schemeArgsForCall)]
	fake.schemeArgsForCall = append(fake.schemeArgsForCall, struct {
	}{})
	stub := fake.SchemeStub
	fakeReturns := fake.schemeReturns
	fake.recordInvocation("Scheme", []interface{}{})
	fake.schemeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeK8sRestClient) SchemeCallCount() int {
	fake.schemeMutex.RLock()
	defer fake.schemeMutex.RUnlock()
	return len(fake.schemeArgsForCall)
}

func (fake *FakeK8sRestClient) SchemeCalls(stub func() *runtime.Scheme) {
	fake.schemeMutex.Lock()
	defer fake.schemeMutex.Unlock()
	fake.SchemeStub = stub
}

func (fake *FakeK8sRestClient) SchemeReturns(result1 *runtime.Scheme) {
	fake.schemeMutex.Lock()
	defer fake.schemeMutex.Unlock()
	fake.SchemeStub = nil
	fake.schemeReturns = struct {
		result1 *runtime.Scheme
	}{result1}
}

func (fake *FakeK8sRestClient) SchemeReturnsOnCall(i int, result1 *runtime.Scheme) {
	fake.schemeMutex.Lock()
	defer fake.schemeMutex.Unlock()
	fake.SchemeStub = nil
	if fake.schemeReturnsOnCall == nil {
		fake.schemeReturnsOnCall = make(map[int]struct {
			result1 *runtime.Scheme
		})
	}
	fake.schemeReturnsOnCall[i] = struct {
		result1 *runtime.Scheme
	}{result1}
}

func (fake *FakeK8sRestClient) Status() clienta.StatusWriter {
	fake.statusMutex.Lock()
	ret, specificReturn := fake.statusReturnsOnCall[len(fake.statusArgsForCall)]
	fake.statusArgsForCall = append(fake.statusArgsForCall, struct {
	}{})
	stub := fake.StatusStub
	fakeReturns := fake.statusReturns
	fake.recordInvocation("Status", []interface{}{})
	fake.statusMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeK8sRestClient) StatusCallCount() int {
	fake.statusMutex.RLock()
	defer fake.statusMutex.RUnlock()
	return len(fake.statusArgsForCall)
}

func (fake *FakeK8sRestClient) StatusCalls(stub func() clienta.StatusWriter) {
	fake.statusMutex.Lock()
	defer fake.statusMutex.Unlock()
	fake.StatusStub = stub
}

func (fake *FakeK8sRestClient) StatusReturns(result1 clienta.StatusWriter) {
	fake.statusMutex.Lock()
	defer fake.statusMutex.Unlock()
	fake.StatusStub = nil
	fake.statusReturns = struct {
		result1 clienta.StatusWriter
	}{result1}
}

func (fake *FakeK8sRestClient) StatusReturnsOnCall(i int, result1 clienta.StatusWriter) {
	fake.statusMutex.Lock()
	defer fake.statusMutex.Unlock()
	fake.StatusStub = nil
	if fake.statusReturnsOnCall == nil {
		fake.statusReturnsOnCall = make(map[int]struct {
			result1 clienta.StatusWriter
		})
	}
	fake.statusReturnsOnCall[i] = struct {
		result1 clienta.StatusWriter
	}{result1}
}

func (fake *FakeK8sRestClient) Update(arg1 context.Context, arg2 clienta.Object, arg3 ...clienta.UpdateOption) error {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 context.Context
		arg2 clienta.Object
		arg3 []clienta.UpdateOption
	}{arg1, arg2, arg3})
	stub := fake.UpdateStub
	fakeReturns := fake.updateReturns
	fake.recordInvocation("Update", []interface{}{arg1, arg2, arg3})
	fake.updateMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeK8sRestClient) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeK8sRestClient) UpdateCalls(stub func(context.Context, clienta.Object, ...clienta.UpdateOption) error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakeK8sRestClient) UpdateArgsForCall(i int) (context.Context, clienta.Object, []clienta.UpdateOption) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeK8sRestClient) UpdateReturns(result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeK8sRestClient) UpdateReturnsOnCall(i int, result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeK8sRestClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.deleteAllOfMutex.RLock()
	defer fake.deleteAllOfMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.patchMutex.RLock()
	defer fake.patchMutex.RUnlock()
	fake.rESTMapperMutex.RLock()
	defer fake.rESTMapperMutex.RUnlock()
	fake.schemeMutex.RLock()
	defer fake.schemeMutex.RUnlock()
	fake.statusMutex.RLock()
	defer fake.statusMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeK8sRestClient) recordInvocation(key string, args []interface{}) {
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

var _ client.K8sRestClient = new(FakeK8sRestClient)