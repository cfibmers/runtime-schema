// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/runtime-schema/cb"
	"github.com/cloudfoundry-incubator/runtime-schema/models"
)

type FakeTaskClient struct {
	CompleteTaskStub        func(receptorURL string, task *models.Task) error
	completeTaskMutex       sync.RWMutex
	completeTaskArgsForCall []struct {
		receptorURL string
		task        *models.Task
	}
	completeTaskReturns struct {
		result1 error
	}
}

func (fake *FakeTaskClient) CompleteTask(receptorURL string, task *models.Task) error {
	fake.completeTaskMutex.Lock()
	fake.completeTaskArgsForCall = append(fake.completeTaskArgsForCall, struct {
		receptorURL string
		task        *models.Task
	}{receptorURL, task})
	fake.completeTaskMutex.Unlock()
	if fake.CompleteTaskStub != nil {
		return fake.CompleteTaskStub(receptorURL, task)
	} else {
		return fake.completeTaskReturns.result1
	}
}

func (fake *FakeTaskClient) CompleteTaskCallCount() int {
	fake.completeTaskMutex.RLock()
	defer fake.completeTaskMutex.RUnlock()
	return len(fake.completeTaskArgsForCall)
}

func (fake *FakeTaskClient) CompleteTaskArgsForCall(i int) (string, *models.Task) {
	fake.completeTaskMutex.RLock()
	defer fake.completeTaskMutex.RUnlock()
	return fake.completeTaskArgsForCall[i].receptorURL, fake.completeTaskArgsForCall[i].task
}

func (fake *FakeTaskClient) CompleteTaskReturns(result1 error) {
	fake.CompleteTaskStub = nil
	fake.completeTaskReturns = struct {
		result1 error
	}{result1}
}

var _ cb.TaskClient = new(FakeTaskClient)