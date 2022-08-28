// Code generated by counterfeiter. DO NOT EDIT.
package transferfakes

import (
	"github.com/veedubyou/chord-paper-be/src/worker/internal/application/jobs/transfer"
	"sync"
)

type FakeTransferJobHandler struct {
	HandleTransferJobStub        func([]byte) (transfer.JobParams, string, error)
	handleTransferJobMutex       sync.RWMutex
	handleTransferJobArgsForCall []struct {
		arg1 []byte
	}
	handleTransferJobReturns struct {
		result1 transfer.JobParams
		result2 string
		result3 error
	}
	handleTransferJobReturnsOnCall map[int]struct {
		result1 transfer.JobParams
		result2 string
		result3 error
	}
	invocations      map[string][][]any
	invocationsMutex sync.RWMutex
}

func (fake *FakeTransferJobHandler) HandleTransferJob(arg1 []byte) (transfer.JobParams, string, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.handleTransferJobMutex.Lock()
	ret, specificReturn := fake.handleTransferJobReturnsOnCall[len(fake.handleTransferJobArgsForCall)]
	fake.handleTransferJobArgsForCall = append(fake.handleTransferJobArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	stub := fake.HandleTransferJobStub
	fakeReturns := fake.handleTransferJobReturns
	fake.recordInvocation("HandleTransferJob", []any{arg1Copy})
	fake.handleTransferJobMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeTransferJobHandler) HandleTransferJobCallCount() int {
	fake.handleTransferJobMutex.RLock()
	defer fake.handleTransferJobMutex.RUnlock()
	return len(fake.handleTransferJobArgsForCall)
}

func (fake *FakeTransferJobHandler) HandleTransferJobCalls(stub func([]byte) (transfer.JobParams, string, error)) {
	fake.handleTransferJobMutex.Lock()
	defer fake.handleTransferJobMutex.Unlock()
	fake.HandleTransferJobStub = stub
}

func (fake *FakeTransferJobHandler) HandleTransferJobArgsForCall(i int) []byte {
	fake.handleTransferJobMutex.RLock()
	defer fake.handleTransferJobMutex.RUnlock()
	argsForCall := fake.handleTransferJobArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeTransferJobHandler) HandleTransferJobReturns(result1 transfer.JobParams, result2 string, result3 error) {
	fake.handleTransferJobMutex.Lock()
	defer fake.handleTransferJobMutex.Unlock()
	fake.HandleTransferJobStub = nil
	fake.handleTransferJobReturns = struct {
		result1 transfer.JobParams
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTransferJobHandler) HandleTransferJobReturnsOnCall(i int, result1 transfer.JobParams, result2 string, result3 error) {
	fake.handleTransferJobMutex.Lock()
	defer fake.handleTransferJobMutex.Unlock()
	fake.HandleTransferJobStub = nil
	if fake.handleTransferJobReturnsOnCall == nil {
		fake.handleTransferJobReturnsOnCall = make(map[int]struct {
			result1 transfer.JobParams
			result2 string
			result3 error
		})
	}
	fake.handleTransferJobReturnsOnCall[i] = struct {
		result1 transfer.JobParams
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTransferJobHandler) Invocations() map[string][][]any {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.handleTransferJobMutex.RLock()
	defer fake.handleTransferJobMutex.RUnlock()
	copiedInvocations := map[string][][]any{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTransferJobHandler) recordInvocation(key string, args []any) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]any{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]any{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ transfer.TransferJobHandler = new(FakeTransferJobHandler)
