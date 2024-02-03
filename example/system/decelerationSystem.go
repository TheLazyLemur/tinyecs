package system

import (
	comp "github.com/TheLazyLemur/tinyecs/example/component"
	"github.com/TheLazyLemur/tinyecs/ecs"
)

type DecelerationSystem struct{}

func (s *DecelerationSystem) Update(entites []*ecs.Entity, dt float32) {
	for _, ent := range entites {
		s.updateOne(ent, dt)
	}
}

func (s *DecelerationSystem) updateOne(ent *ecs.Entity, dt float32) {
	vel := ent.Components[comp.VelocityComponentName].(*comp.VelocityComponent)

	vel.X *= 0.9
	vel.Y *= 0.9
}

func (s *DecelerationSystem) GetFilter() (ecs.FilterType, []string) {
	return ecs.One, []string{
		comp.VelocityComponentName,
	}
}

func(s *DecelerationSystem) SetBaseSystem(_ *ecs.BaseSystem) () {}
