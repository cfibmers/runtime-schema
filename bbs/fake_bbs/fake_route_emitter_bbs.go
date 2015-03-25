// This file was generated by counterfeiter
package fake_bbs

import (
	"sync"
	"time"

	"github.com/cloudfoundry-incubator/runtime-schema/bbs"
	"github.com/tedsuo/ifrit"
)

type FakeRouteEmitterBBS struct {
	NewRouteEmitterLockStub        func(emitterID string, ttl, retryInterval time.Duration) ifrit.Runner
	newRouteEmitterLockMutex       sync.RWMutex
	newRouteEmitterLockArgsForCall []struct {
		emitterID     string
		ttl           time.Duration
		retryInterval time.Duration
	}
	newRouteEmitterLockReturns struct {
		result1 ifrit.Runner
	}
}

func (fake *FakeRouteEmitterBBS) NewRouteEmitterLock(emitterID string, ttl time.Duration, retryInterval time.Duration) ifrit.Runner {
	fake.newRouteEmitterLockMutex.Lock()
	fake.newRouteEmitterLockArgsForCall = append(fake.newRouteEmitterLockArgsForCall, struct {
		emitterID     string
		ttl           time.Duration
		retryInterval time.Duration
	}{emitterID, ttl, retryInterval})
	fake.newRouteEmitterLockMutex.Unlock()
	if fake.NewRouteEmitterLockStub != nil {
		return fake.NewRouteEmitterLockStub(emitterID, ttl, retryInterval)
	} else {
		return fake.newRouteEmitterLockReturns.result1
	}
}

func (fake *FakeRouteEmitterBBS) NewRouteEmitterLockCallCount() int {
	fake.newRouteEmitterLockMutex.RLock()
	defer fake.newRouteEmitterLockMutex.RUnlock()
	return len(fake.newRouteEmitterLockArgsForCall)
}

func (fake *FakeRouteEmitterBBS) NewRouteEmitterLockArgsForCall(i int) (string, time.Duration, time.Duration) {
	fake.newRouteEmitterLockMutex.RLock()
	defer fake.newRouteEmitterLockMutex.RUnlock()
	return fake.newRouteEmitterLockArgsForCall[i].emitterID, fake.newRouteEmitterLockArgsForCall[i].ttl, fake.newRouteEmitterLockArgsForCall[i].retryInterval
}

func (fake *FakeRouteEmitterBBS) NewRouteEmitterLockReturns(result1 ifrit.Runner) {
	fake.NewRouteEmitterLockStub = nil
	fake.newRouteEmitterLockReturns = struct {
		result1 ifrit.Runner
	}{result1}
}

var _ bbs.RouteEmitterBBS = new(FakeRouteEmitterBBS)
