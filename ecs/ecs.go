package ecs

type Component interface{}

type Entity struct {
	ID         int64
	Tag        string
	Components map[string]Component
}

var (
	pool   = newEntityPool(1024)
	lastID int64
)

func NewEntity(tag string, components map[string]Component) Entity {
	ent := pool.acquire()
	lastID++
	ent.ID = lastID
	ent.Components = components
	ent.Tag = tag

	return ent
}

type System interface {
	Update(world *World, dt float32, entity []Entity)
	GetFilters() (FilterType, []string)
}

func GetComponent[T Component](entity Entity, id string) T {
	return entity.Components[id].(T)
}
