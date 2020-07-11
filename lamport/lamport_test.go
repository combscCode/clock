// Test the lamport.go clock struct
package lamport

import "testing"

func TestIncrement(t *testing.T) {
	var clock Lamport
	for i := 0; i < 100; i++ {
		if clock.time != i {
			t.Errorf("Lamport failed, expected internal time of %d, was actually %d", i, clock.time)
		}
		clock.Increment()
	}
}
