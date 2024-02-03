package ecs

import (
	"fmt"
	"sync"
)

type EntityPool struct {
	entities chan Entity
	lock     sync.Mutex
}

func newEntityPool(size int) *EntityPool {
	return &EntityPool{
		entities: make(chan Entity, size),
	}
}

func (pool *EntityPool) acquire() Entity {
	select {
	case obj := <-pool.entities:
		fmt.Println("Acquired object from pool")
		return obj
	default:
		fmt.Println("Created new object")
		return Entity{}
	}
}

func (pool *EntityPool) Release(obj Entity) {
	pool.lock.Lock()
	defer pool.lock.Unlock()

	select {
	case pool.entities <- obj:
		fmt.Println("Returned object to pool")
	default:
		fmt.Println("Discarded object")
	}
}
