package contextimpl

import (
	"errors"
	"reflect"
	"sync"
	"time"
)

// Context ...
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}

type emptyCtx int

func (ctx emptyCtx) Deadline() (deadline time.Time, ok bool) {
	return
}
func (ctx emptyCtx) Done() <-chan struct{} {
	return nil
}
func (ctx emptyCtx) Err() error {
	return nil
}
func (ctx emptyCtx) Value(key interface{}) interface{} {
	return nil
}

var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)

// Background ...
func Background() Context {
	return background
}

// TODO ...
func TODO() Context {
	return todo
}

type cancelCtx struct {
	Context
	done chan struct{}
	err  error
	mu   sync.Mutex
}

func (ctx *cancelCtx) Done() <-chan struct{} {
	return ctx.done
}
func (ctx *cancelCtx) Err() error {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	return ctx.err
}

// ErrCanceled ...
var (
	ErrCanceled        = errors.New("context canceled")
	ErrDeadlineExceded = errors.New("deadline exceded")
)

// CancelFunc ...
type CancelFunc func()

// WithCancel ...
func WithCancel(parent Context) (Context, CancelFunc) {
	ctx := &cancelCtx{
		Context: parent,
		done:    make(chan struct{}),
	}

	cancel := func() {
		ctx.cancel(ErrCanceled)
	}

	go func() {
		select {
		case <-parent.Done():
			ctx.cancel(parent.Err())
		case <-ctx.Done():

		}
	}()

	return ctx, cancel
}

func (ctx *cancelCtx) cancel(err error) {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	if ctx.err != nil {
		return
	}

	ctx.err = err
	close(ctx.done)
}

type deadlineCtx struct {
	*cancelCtx
	deadline time.Time
}

func (ctx *deadlineCtx) Deadline() (deadline time.Time, ok bool) {
	return ctx.deadline, true
}
func (ctx *deadlineCtx) Done() <-chan struct{} {
	return ctx.Done()
}
func (ctx *deadlineCtx) Err() error {
	return ctx.Err()
}
func (ctx *deadlineCtx) Value(key interface{}) interface{} {
	return ctx.Value(key)
}

// WithDeadline ...
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc) {
	cctx, cancel := WithCancel(parent)

	ctx := &deadlineCtx{
		cancelCtx: cctx.(*cancelCtx),
		deadline:  deadline,
	}

	t := time.AfterFunc(time.Until(deadline), func() {
		ctx.cancel(ErrDeadlineExceded)
	})

	stop := func() {
		t.Stop()
		cancel()
	}

	return ctx, stop
}

// WithTimeout ...
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
	return WithDeadline(parent, time.Now().Add(timeout))
}

type valueCtx struct {
	Context
	value, key interface{}
}

func (ctx *valueCtx) Value(key interface{}) interface{} {
	if key == ctx.key {
		return ctx.value
	}

	return ctx.Context.Value(key)
}

//WithValue ...
func WithValue(parent Context, key, value interface{}) Context {

	if key == nil {
		panic("key is nil")
	}

	if reflect.TypeOf(key).Comparable() {
		panic("key is not comparable")
	}

	return &valueCtx{
		Context: parent,
		key:     key,
		value:   value,
	}
}
