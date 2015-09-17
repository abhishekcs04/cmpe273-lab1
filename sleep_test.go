package main

import "testing"
import "time"

type sleepTest struct {
        n        int
        expected time.Duration
}

var sleepTests = []sleepTest {
        {1, 1*time.Second}, {2, 2*time.Second}, {3, 3*time.Second}, 
}


func TestSleep(t *testing.T) {

	for _, tt := range sleepTests {
                startingTime := time.Now().UTC()
				Sleep(tt.n)
				endingTime := time.Now().UTC()
				var duration time.Duration = endingTime.Sub(startingTime)
				if duration < tt.expected{
				t.Errorf("Fail ")
                }
    
	}
}