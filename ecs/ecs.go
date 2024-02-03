package ecs

type Component interface{}

type Entity struct {
	ID         int64
	Components map[string]Component
}

var (
	pool   = newEntityPool(1024)
	lastID int64
)

func NewEntity(components map[string]Component) Entity {
	ent := pool.acquire()
	lastID++
	ent.ID = lastID
	ent.Components = components

	return ent
}

type System interface {
	Update(world *World, entity []Entity)
	GetFilters() (FilterType, []string)
}
