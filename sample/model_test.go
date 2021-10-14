// This is a test file for testing state transitions

package test

import (
"fmt"
"testing"
"time"
)

func TestStateTrans(t *testing.T) {
go func() {
for {
time.Sleep(1 * time.Millisecond)
task()
}
}()
for {
fmt.Scan(&input)
if input == "q" {
fmt.Println("quit")
break
}
}
}

