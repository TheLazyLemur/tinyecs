package ecs

import (
	"sync"
)

type FilterType int

const (
	All FilterType = iota
	One
	EntityTag
)

type World struct {
	Entities map[int64]Entity
	Systems  map[int64]System

	entityLock      *sync.Mutex
	entitesToAdd    []Entity
	entitesToRemove []int64

	systemLock      *sync.Mutex
	systemsToRemove []int64
	systemsToAdd    map[int64]System
}

func NewWorld() World {
	return World{
		Entities:        make(map[int64]Entity),
		Systems:         make(map[int64]System),
		entityLock:      &sync.Mutex{},
		systemLock:      &sync.Mutex{},
		systemsToAdd:    make(map[int64]System),
		systemsToRemove: make([]int64, 0),
		entitesToRemove: make([]int64, 0),
		entitesToAdd:    make([]Entity, 0),
	}
}

var lastSystemID int64

func (w *World) AddSystem(s System) {
	w.systemLock.Lock()
	defer w.systemLock.Unlock()

	lastSystemID++
	w.systemsToAdd[lastSystemID] = s
}

func (w *World) AddEntity(e Entity) {
	w.entityLock.Lock()
	defer w.entityLock.Unlock()

	w.entitesToAdd = append(w.entitesToAdd, e)
}

func (w *World) RemoveEntity(id int64) {
	w.entityLock.Lock()
	defer w.entityLock.Unlock()

	w.entitesToRemove = append(w.entitesToRemove, id)
}

func (w *World) Update(dt float32) {
	w.cleanEntities()
	w.addNewEnties()

	w.cleanSystems()
	w.addNewSystems()

	for _, s := range w.Systems {
		fType, filters := s.GetFilters()

		ents := make([]Entity, 0)

		for _, e := range w.Entities {
			switch fType {
			case One:
				shouldAdd := false
				for _, f := range filters {
					if e.Components[f] != nil {
						shouldAdd = true
					}
				}
				if shouldAdd {
					ents = append(ents, e)
				}
			case All:
				shouldAdd := true

				for _, f := range filters {
					if e.Components[f] == nil {
						shouldAdd = false
					}
				}

				if shouldAdd {
					ents = append(ents, e)
				}
			case EntityTag:
				shouldAdd := false
				for _, f := range filters {
					if e.Tag == f {
						shouldAdd = true
					}
				}

				if shouldAdd {
					ents = append(ents, e)
				}
			}
		}

		s.Update(w, dt, ents)
	}
}

func (w *World) cleanEntities() {
	w.entityLock.Lock()
	defer w.entityLock.Unlock()

	for _, id := range w.entitesToRemove {
		ent := w.Entities[id]
		pool.Release(ent)
		delete(w.Entities, id)
	}
	w.entitesToRemove = make([]int64, 0)
}

func (w *World) addNewEnties() {
	w.entityLock.Lock()
	defer w.entityLock.Unlock()

	for _, e := range w.entitesToAdd {
		w.Entities[e.ID] = e
	}

	w.entitesToAdd = make([]Entity, 0)
}

func (w *World) cleanSystems() {
	w.systemLock.Lock()
	defer w.systemLock.Unlock()
}

func (w *World) addNewSystems() {
	w.systemLock.Lock()
	defer w.systemLock.Unlock()

	for id, s := range w.systemsToAdd {
		w.Systems[id] = s
	}

	w.systemsToAdd = make(map[int64]System)
}
