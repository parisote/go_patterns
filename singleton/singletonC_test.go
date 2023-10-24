package singleton

import (
	"sync"
	"testing"
)

func TestGetInstanceConcurrency(t *testing.T) {
	instances := make(chan Singleton, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		instances <- GetInstanceConcurrency()
		wg.Done()
	}()

	go func() {
		instances <- GetInstanceConcurrency()
		wg.Done()
	}()

	wg.Wait()
	close(instances)

	instance1 := <-instances
	instance2 := <-instances

	if instance1 != instance2 {
		t.Errorf("Expected same instance in concurrent GetInstance call, but got different instances.")
	}
}
