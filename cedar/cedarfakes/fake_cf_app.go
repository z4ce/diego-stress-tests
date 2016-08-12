// This file was generated by counterfeiter
package cedarfakes

import (
	"sync"
	"time"

	"code.cloudfoundry.org/diego-stress-tests/cedar"
	"golang.org/x/net/context"
)

type FakeCfApp struct {
	AppNameStub        func() string
	appNameMutex       sync.RWMutex
	appNameArgsForCall []struct{}
	appNameReturns     struct {
		result1 string
	}
	PushStub        func(ctx context.Context, payload string, timeout time.Duration) error
	pushMutex       sync.RWMutex
	pushArgsForCall []struct {
		ctx     context.Context
		payload string
		timeout time.Duration
	}
	pushReturns struct {
		result1 error
	}
	StartStub        func(ctx context.Context, timeout time.Duration) error
	startMutex       sync.RWMutex
	startArgsForCall []struct {
		ctx     context.Context
		timeout time.Duration
	}
	startReturns struct {
		result1 error
	}
	GuidStub        func(ctx context.Context, timeout time.Duration) (string, error)
	guidMutex       sync.RWMutex
	guidArgsForCall []struct {
		ctx     context.Context
		timeout time.Duration
	}
	guidReturns struct {
		result1 string
		result2 error
	}
}

func (fake *FakeCfApp) AppName() string {
	fake.appNameMutex.Lock()
	fake.appNameArgsForCall = append(fake.appNameArgsForCall, struct{}{})
	fake.appNameMutex.Unlock()
	if fake.AppNameStub != nil {
		return fake.AppNameStub()
	} else {
		return fake.appNameReturns.result1
	}
}

func (fake *FakeCfApp) AppNameCallCount() int {
	fake.appNameMutex.RLock()
	defer fake.appNameMutex.RUnlock()
	return len(fake.appNameArgsForCall)
}

func (fake *FakeCfApp) AppNameReturns(result1 string) {
	fake.AppNameStub = nil
	fake.appNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCfApp) Push(ctx context.Context, payload string, timeout time.Duration) error {
	fake.pushMutex.Lock()
	fake.pushArgsForCall = append(fake.pushArgsForCall, struct {
		ctx     context.Context
		payload string
		timeout time.Duration
	}{ctx, payload, timeout})
	fake.pushMutex.Unlock()
	if fake.PushStub != nil {
		return fake.PushStub(ctx, payload, timeout)
	} else {
		return fake.pushReturns.result1
	}
}

func (fake *FakeCfApp) PushCallCount() int {
	fake.pushMutex.RLock()
	defer fake.pushMutex.RUnlock()
	return len(fake.pushArgsForCall)
}

func (fake *FakeCfApp) PushArgsForCall(i int) (context.Context, string, time.Duration) {
	fake.pushMutex.RLock()
	defer fake.pushMutex.RUnlock()
	return fake.pushArgsForCall[i].ctx, fake.pushArgsForCall[i].payload, fake.pushArgsForCall[i].timeout
}

func (fake *FakeCfApp) PushReturns(result1 error) {
	fake.PushStub = nil
	fake.pushReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCfApp) Start(ctx context.Context, timeout time.Duration) error {
	fake.startMutex.Lock()
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
		ctx     context.Context
		timeout time.Duration
	}{ctx, timeout})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		return fake.StartStub(ctx, timeout)
	} else {
		return fake.startReturns.result1
	}
}

func (fake *FakeCfApp) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeCfApp) StartArgsForCall(i int) (context.Context, time.Duration) {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return fake.startArgsForCall[i].ctx, fake.startArgsForCall[i].timeout
}

func (fake *FakeCfApp) StartReturns(result1 error) {
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCfApp) Guid(ctx context.Context, timeout time.Duration) (string, error) {
	fake.guidMutex.Lock()
	fake.guidArgsForCall = append(fake.guidArgsForCall, struct {
		ctx     context.Context
		timeout time.Duration
	}{ctx, timeout})
	fake.guidMutex.Unlock()
	if fake.GuidStub != nil {
		return fake.GuidStub(ctx, timeout)
	} else {
		return fake.guidReturns.result1, fake.guidReturns.result2
	}
}

func (fake *FakeCfApp) GuidCallCount() int {
	fake.guidMutex.RLock()
	defer fake.guidMutex.RUnlock()
	return len(fake.guidArgsForCall)
}

func (fake *FakeCfApp) GuidArgsForCall(i int) (context.Context, time.Duration) {
	fake.guidMutex.RLock()
	defer fake.guidMutex.RUnlock()
	return fake.guidArgsForCall[i].ctx, fake.guidArgsForCall[i].timeout
}

func (fake *FakeCfApp) GuidReturns(result1 string, result2 error) {
	fake.GuidStub = nil
	fake.guidReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

var _ cedar.CfApp = new(FakeCfApp)
