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

func TestBasicTab(t *testing.T) {
	now, _ := time.Parse("15:04", "16:30")

	tab := NewTab()

	// Runs at 12:30.
	event := NewEventSimpleFormatValues(-1, -1, "")
	tab.PushEvent(event)

	// Ensure we have no events yet.
	q := tab.buildQueue(&now)
	if len(q) != 0 {
		t.Fatal("Shouldn't be running event yet.")
	}
}
