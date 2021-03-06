// This file was generated by counterfeiter
package fake_listener

import (
	"sync"

	"github.com/cloudfoundry-incubator/garden-linux/container_daemon"
	"github.com/cloudfoundry-incubator/garden-linux/container_daemon/unix_socket"
)

type FakeListener struct {
	InitStub        func() error
	initMutex       sync.RWMutex
	initArgsForCall []struct{}
	initReturns     struct {
		result1 error
	}
	ListenStub        func(ch unix_socket.ConnectionHandler) error
	listenMutex       sync.RWMutex
	listenArgsForCall []struct {
		ch unix_socket.ConnectionHandler
	}
	listenReturns struct {
		result1 error
	}
	StopStub        func() error
	stopMutex       sync.RWMutex
	stopArgsForCall []struct{}
	stopReturns     struct {
		result1 error
	}
}

func (fake *FakeListener) Init() error {
	fake.initMutex.Lock()
	fake.initArgsForCall = append(fake.initArgsForCall, struct{}{})
	fake.initMutex.Unlock()
	if fake.InitStub != nil {
		return fake.InitStub()
	} else {
		return fake.initReturns.result1
	}
}

func (fake *FakeListener) InitCallCount() int {
	fake.initMutex.RLock()
	defer fake.initMutex.RUnlock()
	return len(fake.initArgsForCall)
}

func (fake *FakeListener) InitReturns(result1 error) {
	fake.InitStub = nil
	fake.initReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeListener) Listen(ch unix_socket.ConnectionHandler) error {
	fake.listenMutex.Lock()
	fake.listenArgsForCall = append(fake.listenArgsForCall, struct {
		ch unix_socket.ConnectionHandler
	}{ch})
	fake.listenMutex.Unlock()
	if fake.ListenStub != nil {
		return fake.ListenStub(ch)
	} else {
		return fake.listenReturns.result1
	}
}

func (fake *FakeListener) ListenCallCount() int {
	fake.listenMutex.RLock()
	defer fake.listenMutex.RUnlock()
	return len(fake.listenArgsForCall)
}

func (fake *FakeListener) ListenArgsForCall(i int) unix_socket.ConnectionHandler {
	fake.listenMutex.RLock()
	defer fake.listenMutex.RUnlock()
	return fake.listenArgsForCall[i].ch
}

func (fake *FakeListener) ListenReturns(result1 error) {
	fake.ListenStub = nil
	fake.listenReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeListener) Stop() error {
	fake.stopMutex.Lock()
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct{}{})
	fake.stopMutex.Unlock()
	if fake.StopStub != nil {
		return fake.StopStub()
	} else {
		return fake.stopReturns.result1
	}
}

func (fake *FakeListener) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeListener) StopReturns(result1 error) {
	fake.StopStub = nil
	fake.stopReturns = struct {
		result1 error
	}{result1}
}

var _ container_daemon.Listener = new(FakeListener)
