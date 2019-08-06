//
// Copyright © 2019 Ikey Doherty <contactikeydoherty@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package cron

import (
	"fmt"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

// A Tab is used internally to store cron events, which are then
// available for query and scheduling.
//
// As a design decision, the Tab is not responsible for the parsing
// of crontab files, as it would force the scheduler component to become
// I/O bound. Instead it becomes the responsibility of the calling program
// to feed parsed events.
type Tab struct {
	// Map our entries by their TID to allow expirations.
	entries map[int64]Event

	mut *sync.RWMutex
	tid int64
}

// NewTab will construct a new Tab type and prepare it for
// use within the calling process.
func NewTab() *Tab {
	return &Tab{

		entries: make(map[int64]Event),
		mut:     &sync.RWMutex{},
	}
}

// PushEvent will attempt to insert the event into the crontab for
// future scheduling
func (t *Tab) PushEvent(e Event) {
	tid := t.nextTID()
	e.setTID(tid)

	defer t.mut.Unlock()
	t.mut.Lock()

	// Insert the event
	t.entries[tid] = e
}

// buildQueue will construct a time-linear queue of events to process
// This ensures we clear out unnecessary entries
func (t *Tab) buildQueue(whence *time.Time) EventQueue {
	defer t.mut.RUnlock()
	t.mut.RLock()

	var ret EventQueue
	for _, val := range t.entries {
		if whence != nil && !val.Timing().ShouldRun(*whence) {
			continue
		}
		ret = append(ret, val)
	}
	sort.Sort(ret)
	return ret
}

// expireEvent will remove the event from the list of known events
func (t *Tab) expireEvent(e Event) {
	// Set the next run-time if appropriate
	now := time.Now()
	e.Timing().NextTimestamp(now, true)

	// If we repeat (basically, always) then don't expire.
	// This just opens the scope to one-shot events being added
	// to the scheduler.
	if e.Timing().Repeats() {
		return
	}

	defer t.mut.Unlock()
	t.mut.Lock()

	if _, ok := t.entries[e.TID()]; !ok {
		return
	}

	delete(t.entries, e.TID())
}

func (t *Tab) nextTID() int64 {
	return atomic.AddInt64(&t.tid, 1)
}

// Run is a dummy function that pretends to run all the events.
func (t *Tab) Run(tm time.Time) {
	var queue EventQueue

	queue = t.buildQueue(&tm)

	if len(queue) > 0 {
		fmt.Printf("Got %v runnables\n", len(queue))
	}

	// Run all runnable events
	for _, event := range queue {
		// Execute it
		event.Execute(time.Now())

		// Expire it if required
		t.expireEvent(event)
	}
}
