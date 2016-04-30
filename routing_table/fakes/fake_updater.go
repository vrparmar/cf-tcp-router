// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/cf-tcp-router/routing_table"
	"github.com/cloudfoundry-incubator/routing-api"
)

type FakeUpdater struct {
	HandleEventStub        func(event routing_api.TcpEvent) error
	handleEventMutex       sync.RWMutex
	handleEventArgsForCall []struct {
		event routing_api.TcpEvent
	}
	handleEventReturns struct {
		result1 error
	}
	SyncStub           func()
	syncMutex          sync.RWMutex
	syncArgsForCall    []struct{}
	SyncingStub        func() bool
	syncingMutex       sync.RWMutex
	syncingArgsForCall []struct{}
	syncingReturns     struct {
		result1 bool
	}
	PruneStaleRoutesStub        func()
	pruneStaleRoutesMutex       sync.RWMutex
	pruneStaleRoutesArgsForCall []struct{}
}

func (fake *FakeUpdater) HandleEvent(event routing_api.TcpEvent) error {
	fake.handleEventMutex.Lock()
	fake.handleEventArgsForCall = append(fake.handleEventArgsForCall, struct {
		event routing_api.TcpEvent
	}{event})
	fake.handleEventMutex.Unlock()
	if fake.HandleEventStub != nil {
		return fake.HandleEventStub(event)
	} else {
		return fake.handleEventReturns.result1
	}
}

func (fake *FakeUpdater) HandleEventCallCount() int {
	fake.handleEventMutex.RLock()
	defer fake.handleEventMutex.RUnlock()
	return len(fake.handleEventArgsForCall)
}

func (fake *FakeUpdater) HandleEventArgsForCall(i int) routing_api.TcpEvent {
	fake.handleEventMutex.RLock()
	defer fake.handleEventMutex.RUnlock()
	return fake.handleEventArgsForCall[i].event
}

func (fake *FakeUpdater) HandleEventReturns(result1 error) {
	fake.HandleEventStub = nil
	fake.handleEventReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUpdater) Sync() {
	fake.syncMutex.Lock()
	fake.syncArgsForCall = append(fake.syncArgsForCall, struct{}{})
	fake.syncMutex.Unlock()
	if fake.SyncStub != nil {
		fake.SyncStub()
	}
}

func (fake *FakeUpdater) SyncCallCount() int {
	fake.syncMutex.RLock()
	defer fake.syncMutex.RUnlock()
	return len(fake.syncArgsForCall)
}

func (fake *FakeUpdater) Syncing() bool {
	fake.syncingMutex.Lock()
	fake.syncingArgsForCall = append(fake.syncingArgsForCall, struct{}{})
	fake.syncingMutex.Unlock()
	if fake.SyncingStub != nil {
		return fake.SyncingStub()
	} else {
		return fake.syncingReturns.result1
	}
}

func (fake *FakeUpdater) SyncingCallCount() int {
	fake.syncingMutex.RLock()
	defer fake.syncingMutex.RUnlock()
	return len(fake.syncingArgsForCall)
}

func (fake *FakeUpdater) SyncingReturns(result1 bool) {
	fake.SyncingStub = nil
	fake.syncingReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeUpdater) PruneStaleRoutes() {
	fake.pruneStaleRoutesMutex.Lock()
	fake.pruneStaleRoutesArgsForCall = append(fake.pruneStaleRoutesArgsForCall, struct{}{})
	fake.pruneStaleRoutesMutex.Unlock()
	if fake.PruneStaleRoutesStub != nil {
		fake.PruneStaleRoutesStub()
	}
}

func (fake *FakeUpdater) PruneStaleRoutesCallCount() int {
	fake.pruneStaleRoutesMutex.RLock()
	defer fake.pruneStaleRoutesMutex.RUnlock()
	return len(fake.pruneStaleRoutesArgsForCall)
}

var _ routing_table.Updater = new(FakeUpdater)
