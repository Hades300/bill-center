package flow

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestLeaky_Wait(t *testing.T) {
	var err error
	var f = NewLeaky(2, 3)
	begin := time.Now()
	cnt := 20
	for i := 0; i < cnt; i++ {
		err = f.Wait(1)
		if err != nil {
			t.Error(err)
		}
		t.Logf("at %f second for no.%d request", time.Since(begin).Seconds(), i+1)
	}
}

func TestLeaky_WaitParallel(t *testing.T) {
	var f = NewLeaky(2, 3)
	begin := time.Now()
	wg := sync.WaitGroup{}
	waiter := func(id int, ti time.Time) {
		var current = time.Now()
		err := f.Wait(1)
		if err != nil {
			assert.NoError(t, err)
		}
		t.Logf("at %f second for no.%d request", current.Sub(begin).Seconds(), id)
		wg.Done()
	}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go waiter(i+1, begin)
	}
	doneC := make(chan bool)
	go func() {
		wg.Wait()
		doneC <- true
	}()
	select {
	case <-time.After(3 * time.Second):
		t.Log("all 100 waiter get their req more than 3 seconds,success")
	case <-doneC:
		t.Fatal("all 100 waiter get their req in 3 seconds,failed")
		return
	}
}
