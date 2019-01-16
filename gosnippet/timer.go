package gosnippet

import "time"

var (
	TimerMockTime = time.Second * 10
)

// the right way to use Timer
func RightWayToUseTimer() {
	tm := time.NewTimer(TimerMockTime)
	for {
		select {
		case <-tm.C:
			// the first way
			if !tm.Stop() && len(tm.C) > 0 {
				<-tm.C
			}
			tm.Reset(TimerMockTime)
			// the second way
			if !tm.Stop() {
				select {
				case <-tm.C:
				default:
				}
			}
			tm.Reset(TimerMockTime)
			// else...
		}
	}
}
