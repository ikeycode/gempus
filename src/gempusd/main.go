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

package main

import (
	"cron"
	"fmt"
	"time"
)

// pollJobs is a dumb function to pull jobs in order.
func pollJobs(c *cron.Tab) {
	fmt.Println("Checking jobs.")
	c.Run()
}

func main() {
	tab := cron.NewTab()

	// We don't actually use it.
	q := make(chan bool)

	tab.PushEvent(cron.NewEventSimpleFormatValues(1, 30, "/bin/run_me_daily"))
	tab.PushEvent(cron.NewEventSimpleFormatValues(-1, 45, "/bin/run_me_hourly"))
	tab.PushEvent(cron.NewEventSimpleFormatValues(-1, -1, "/bin/run_me_every_minute"))
	tab.PushEvent(cron.NewEventSimpleFormatValues(19, -1, "/bin/run_me_sixty_times"))

	for {
		select {
		case <-time.After(1 * time.Second):
			pollJobs(tab)
		case <-q:
			break
		}
	}
}
