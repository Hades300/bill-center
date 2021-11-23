package flow

import (
	"errors"
	"time"
)

var (
	ErrExceedRate = errors.New("exceed rate")
	ErrTimeout    = errors.New("timeout")
)

// flow control based on leaky bucket
type Leaky struct {
	// flow rate each time gap, minimum 1
	rate int64
	// time gap, minimum 1
	gap int64
	// last time
	last int64
	// current flow in bucket
	flow int64
}

// NewLeaky returns a leaky bucket flow control
func NewLeaky(rate int64, gap int64) *Leaky {
	return &Leaky{rate: rate, gap: gap}
}

func (l *Leaky) Get(n int64) bool {
	if n > l.rate {
		// exceed rate
		return false
	}
	// if last is not set, set it to now
	if l.last == 0 {
		l.last = time.Now().Unix()
	}

	// calc grow and update current flow
	now := time.Now().Unix()
	grow := int64(float64(now-l.last) * float64(l.rate) / float64(l.gap))
	l.flow = min(l.flow+grow, l.rate)

	if l.flow >= n {
		l.flow -= n
		l.last = now
		return true
	}
	return false
}

func (l *Leaky) Wait(n int64) error {
	if n > l.rate {
		// exceed rate
		return ErrExceedRate
	}
	try := l.Get(n)
	if try {
		return nil
	}
	// calc next time
	waitTime := int(float64(n-l.flow) / float64(l.rate) * float64(l.gap) * 1000) // millisecond

	time.Sleep(time.Duration(waitTime) * time.Millisecond)
	now := time.Now().Unix()
	l.last = now
	return nil
}

// willing to cost time, but if we can't afford, then fail fast
func (l *Leaky) TryWait(n int64, millisecond time.Duration) bool {
	try := l.Get(n)
	if try {
		return true
	}
	// calc next time
	waitTime := int(float64(n-l.flow) / float64(l.rate) * float64(l.gap) * 1000) // millisecond
	if waitTime > int(millisecond/time.Millisecond) {
		return false
	}
	time.Sleep(time.Duration(waitTime) * time.Millisecond)
	now := time.Now().Unix()
	l.last = now
	return true
}

func min(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
