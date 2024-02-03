package ecs

import (
	"fmt"
	"os"
	"unsafe"
)

type FilterType int

const (
	All FilterType = iota
	One
)

type Component interface{}

type Entity struct {
	Components map[string]Component
}

type BaseSystem struct {
	World *World
}

type System interface {
	Update(e []*Entity, dt float32)
	GetFilter() (FilterType, []string)
	SetBaseSystem(bs *BaseSystem)
}

type World struct {
	entities         map[*Entity]struct{}
	entitiesToAdd    []*Entity
	entitiesToRemove []*Entity
	systems          map[System]struct{}
	systemsToAdd     []System
	systemsToRemove  []System

	SceneSwapCallback func(string)
}

func (w *World) CallSceneSwap(scene string) {
	if w.SceneSwapCallback != nil {
		w.SceneSwapCallback(scene)
	}
}

func (w *World) RemoveAllSystems() {
	for s := range w.systems {
		w.systemsToRemove = append(w.systemsToRemove, s)
	}
}

func (w *World) RemoveSystem(s System) {
	w.systemsToRemove = append(w.systemsToRemove, s)
}

func (w *World) RemoveAllSystemsExcept(s ...System) {
	for systems := range w.systems {
		for _, ss := range s {
			if systems != ss {
				w.systemsToRemove = append(w.systemsToRemove, systems)
			}
		}
	}
}

func (w *World) AddSystem(s ...System) {
	if w.systems == nil {
		w.systems = make(map[System]struct{})
	}

	bs := &BaseSystem{
		World: w,
	}

	for _, ss := range s {
		ss.SetBaseSystem(bs)
	}

	w.systemsToAdd = append(w.systemsToAdd, s...)
}

func (w *World) AddEntity(e ...*Entity) {
	if w.entities == nil {
		w.entities = make(map[*Entity]struct{})
	}

	w.entitiesToAdd = append(w.entitiesToAdd, e...)
}

func (w *World) RemoveAllEntities() {
	for e := range w.entities {
		w.entitiesToRemove = append(w.entitiesToRemove, e)
	}
}

func (w *World) RemoveEntity(e *Entity) {
	w.entitiesToRemove = append(w.entitiesToRemove, e)
}

func (w *World) Update(dt float32) {
	for _, s := range w.systemsToRemove {
		delete(w.systems, s)
	}
	w.systemsToRemove = []System{}

	for _, s := range w.systemsToAdd {
		w.systems[s] = struct{}{}
	}
	w.systemsToAdd = []System{}

	for _, e := range w.entitiesToRemove {
		delete(w.entities, e)
	}

	w.entitiesToRemove = []*Entity{}

	for _, e := range w.entitiesToAdd {
		w.entities[e] = struct{}{}
	}
	w.entitiesToAdd = []*Entity{}

	str := fmt.Sprintf("Entities: %v, Systems: %v\n", len(w.entities), len(w.systems))
	fmt.Println(str)

	objectSize := unsafe.Sizeof(w.entities)
	objectSizeInMB := float64(objectSize) / (1024 * 1024)
	fmt.Fprintf(os.Stdout, "Size of World: %v bytes or %.2f MB\n", []any{objectSize, objectSizeInMB}...)

	for s := range w.systems {
		typ, filters := s.GetFilter()
		ents := []*Entity{}

		switch typ {
		case All:
			for e := range w.entities {
				if w.hasAll(e, filters) {
					ents = append(ents, e)
				}
			}
		case One:
			for e := range w.entities {
				if w.hasOne(e, filters) {
					ents = append(ents, e)
				}
			}
		default:
			panic("unknown filter type")
		}

		s.Update(ents, dt)
	}
}

func (w *World) hasAll(e *Entity, filters []string) bool {
	shouldUpdate := true
	for _, f := range filters {
		_, ok := e.Components[f]
		if !ok {
			shouldUpdate = false
		}
	}

	return shouldUpdate
}

func (w *World) hasOne(e *Entity, filters []string) bool {
	shouldUpdate := false
	for _, f := range filters {
		_, ok := e.Components[f]
		if ok {
			shouldUpdate = true
		}
	}

	return shouldUpdate
}
