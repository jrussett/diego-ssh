// Code generated by counterfeiter. DO NOT EDIT.
package fake_handlers

import (
	"sync"

	"code.cloudfoundry.org/diego-ssh/handlers"
	"code.cloudfoundry.org/diego-ssh/helpers"
	"code.cloudfoundry.org/lager"
	"golang.org/x/crypto/ssh"
)

type FakeGlobalRequestHandler struct {
	HandleRequestStub        func(logger lager.Logger, request *ssh.Request, conn ssh.Conn, lnStore *helpers.TCPIPListenerStore)
	handleRequestMutex       sync.RWMutex
	handleRequestArgsForCall []struct {
		logger  lager.Logger
		request *ssh.Request
		conn    ssh.Conn
		lnStore *helpers.TCPIPListenerStore
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGlobalRequestHandler) HandleRequest(logger lager.Logger, request *ssh.Request, conn ssh.Conn, lnStore *helpers.TCPIPListenerStore) {
	fake.handleRequestMutex.Lock()
	fake.handleRequestArgsForCall = append(fake.handleRequestArgsForCall, struct {
		logger  lager.Logger
		request *ssh.Request
		conn    ssh.Conn
		lnStore *helpers.TCPIPListenerStore
	}{logger, request, conn, lnStore})
	fake.recordInvocation("HandleRequest", []interface{}{logger, request, conn, lnStore})
	fake.handleRequestMutex.Unlock()
	if fake.HandleRequestStub != nil {
		fake.HandleRequestStub(logger, request, conn, lnStore)
	}
}

func (fake *FakeGlobalRequestHandler) HandleRequestCallCount() int {
	fake.handleRequestMutex.RLock()
	defer fake.handleRequestMutex.RUnlock()
	return len(fake.handleRequestArgsForCall)
}

func (fake *FakeGlobalRequestHandler) HandleRequestArgsForCall(i int) (lager.Logger, *ssh.Request, ssh.Conn, *helpers.TCPIPListenerStore) {
	fake.handleRequestMutex.RLock()
	defer fake.handleRequestMutex.RUnlock()
	return fake.handleRequestArgsForCall[i].logger, fake.handleRequestArgsForCall[i].request, fake.handleRequestArgsForCall[i].conn, fake.handleRequestArgsForCall[i].lnStore
}

func (fake *FakeGlobalRequestHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.handleRequestMutex.RLock()
	defer fake.handleRequestMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGlobalRequestHandler) recordInvocation(key string, args []interface{}) {
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

var _ handlers.GlobalRequestHandler = new(FakeGlobalRequestHandler)
