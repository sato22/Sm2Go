// This is a test file for testing state transitions

package stopwatch

import (
"testing"
"time"
"sync"
)

func TestStateTrans(t *testing.T) {
var wg sync.WaitGroup
wg.Add(1)
go func() {
for {
time.Sleep(1 * time.Millisecond)
task()
}
wg.Done()
}()
wg.Wait()
}

