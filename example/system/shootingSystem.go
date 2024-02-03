package system

import (
	"github.com/TheLazyLemur/tinyecs/ecs"
	"github.com/TheLazyLemur/tinyecs/example/entities"

	comp "github.com/TheLazyLemur/tinyecs/example/component"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ShootingSystem struct {
	BaseSystem *ecs.BaseSystem
}

func (s *ShootingSystem) Update(entites []*ecs.Entity, dt float32) {
	for _, ent := range entites {
		s.updateOne(ent, dt)
	}
}

func (s *ShootingSystem) updateOne(ent *ecs.Entity, dt float32) {
	shooting := ent.Components[comp.ShootingComponentName].(*comp.ShootingComponent)
	pos := ent.Components[comp.PositionComponentName].(*comp.PositionComponent)

	shooting.Cooldown -= dt
	if shooting.Cooldown < 0 && rl.IsKeyDown(rl.KeySpace) {
		s.BaseSystem.World.AddEntity(entities.Bullet(pos.X, pos.Y))
		shooting.Cooldown = 0.3
	}
}

func (s *ShootingSystem) SetBaseSystem(bs *ecs.BaseSystem) {
	s.BaseSystem = bs
}

func (s *ShootingSystem) GetFilter() (ecs.FilterType, []string) {
	return ecs.All, []string{
		comp.ShootingComponentName,
		comp.PositionComponentName,
	}
}
