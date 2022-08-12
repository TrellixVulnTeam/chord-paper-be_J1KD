// Code generated by counterfeiter. DO NOT EDIT.
package publishfakes

import (
	"github.com/veedubyou/chord-paper-be/worker/src/internal/application/publish"
	"sync"

	"github.com/rabbitmq/amqp091-go"
)

type FakePublisher struct {
	PublishStub        func(amqp091.Publishing) error
	publishMutex       sync.RWMutex
	publishArgsForCall []struct {
		arg1 amqp091.Publishing
	}
	publishReturns struct {
		result1 error
	}
	publishReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePublisher) Publish(arg1 amqp091.Publishing) error {
	fake.publishMutex.Lock()
	ret, specificReturn := fake.publishReturnsOnCall[len(fake.publishArgsForCall)]
	fake.publishArgsForCall = append(fake.publishArgsForCall, struct {
		arg1 amqp091.Publishing
	}{arg1})
	stub := fake.PublishStub
	fakeReturns := fake.publishReturns
	fake.recordInvocation("Publish", []interface{}{arg1})
	fake.publishMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePublisher) PublishCallCount() int {
	fake.publishMutex.RLock()
	defer fake.publishMutex.RUnlock()
	return len(fake.publishArgsForCall)
}

func (fake *FakePublisher) PublishCalls(stub func(amqp091.Publishing) error) {
	fake.publishMutex.Lock()
	defer fake.publishMutex.Unlock()
	fake.PublishStub = stub
}

func (fake *FakePublisher) PublishArgsForCall(i int) amqp091.Publishing {
	fake.publishMutex.RLock()
	defer fake.publishMutex.RUnlock()
	argsForCall := fake.publishArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePublisher) PublishReturns(result1 error) {
	fake.publishMutex.Lock()
	defer fake.publishMutex.Unlock()
	fake.PublishStub = nil
	fake.publishReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePublisher) PublishReturnsOnCall(i int, result1 error) {
	fake.publishMutex.Lock()
	defer fake.publishMutex.Unlock()
	fake.PublishStub = nil
	if fake.publishReturnsOnCall == nil {
		fake.publishReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePublisher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.publishMutex.RLock()
	defer fake.publishMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePublisher) recordInvocation(key string, args []interface{}) {
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

var _ publish.Publisher = new(FakePublisher)
