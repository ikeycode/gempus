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
	"testing"
	"time"
)

// Add basic test for simple format
func TestBasicEvent(t *testing.T) {
	event := NewEventSimpleFormatValues(-1, -1, "")
	if !event.Timing().HasFlag(EventPerHour) {
		t.Fatal("Missing hour flag")
	}
	if !event.Timing().HasFlag(EventPerMinute) {
		t.Fatal("Missing minute flag")
	}
	if !event.Timing().HasFlag(EventRepeats) {
		t.Fatal("Missing repeat flag")
	}

	event = NewEventSimpleFormatValues(12, -1, "")
	if !event.Timing().HasFlag(EventPerMinute) {
		t.Fatal("Missing minute flag")
	}
	if event.Timing().HasFlag(EventPerHour) {
		t.Fatal("Extra hour flag")
	}
}

// TestTextEvent will attempt parsing of events from text input
func TestTextEvent(t *testing.T) {
	// Good event
	event, err := NewEventSimpleFormat("* * /bin/bash")
	if err != nil {
		t.Fatal("Couldn't parse good string")
	}
	if !event.Timing().HasFlag(EventPerHour | EventPerMinute) {
		t.Fatal("Missing flags on event")
	}

	event, err = NewEventSimpleFormat("30 * /bin/bash")
	if err != nil {
		t.Fatal("Couldn't parse good string")
	}
	if !event.Timing().HasFlag(EventPerHour | EventRepeats) {
		t.Fatal("Missing flags on event")
	}
	if event.Timing().HasFlag(EventPerMinute) {
		t.Fatal("Extra flags on event")
	}
	if event.Timing().Minute != 30 {
		t.Fatal("Invalid minute on event")
	}

	event, err = NewEventSimpleFormat("s* * /bin/fail")
	if err == nil {
		t.Fatal("Shouldn't parse event")
	}
}

// TestTimings will ensure timings are correct.
func TestTimings(t *testing.T) {
	// Set the current time to 16:10
	now, _ := time.Parse("15:04", "16:10")

	// Run command at 13:10, which must be YESTERDAY
	event := NewEventSimpleFormatValues(13, 10, "/bin/bash")
	event.Timing().NextTimestamp(now, false)

	if event.Timing().tm.Day() != now.Day()+1 {
		t.Fatal("incorrect day for event")
	}

	// Runs at 11 minutes every hour. Must run THIS hour too
	event = NewEventSimpleFormatValues(-1, 11, "/bin/bash")
	event.Timing().NextTimestamp(now, false)

	if event.Timing().tm.Hour() != now.Hour() {
		t.Fatal("incorrect hour for event")
	}

	// Runs at 9 minutes past the hour. Must be NEXT hour
	event = NewEventSimpleFormatValues(-1, 9, "/bin/bash")
	event.Timing().NextTimestamp(now, false)

	if event.Timing().tm.Hour() != now.Hour()+1 {
		t.Fatal("incorrect hour for event")
	}
}
