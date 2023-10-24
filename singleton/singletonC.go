package singleton

import "sync"

type SingletonC interface {
	AddOne() int
}

type singletonC struct {
	count int
}

func (s *singleton) AddOneConcurrency() int {
	s.count++
	return s.count
}

var instanceCon *singleton
var once sync.Once

func GetInstanceConcurrency() Singleton {
	once.Do(func() {
		instanceCon = new(singleton)
	})
	return instance
}
