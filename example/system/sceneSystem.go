package system

import (
	"github.com/TheLazyLemur/tinyecs/ecs"
	"github.com/TheLazyLemur/tinyecs/example/entities"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type SceneSystem struct {
	BaseSystem *ecs.BaseSystem
	hasStarted bool
}

func (s *SceneSystem) Update(entites []*ecs.Entity, dt float32) {
	if !s.hasStarted {
		s.BaseSystem.World.SceneSwapCallback = func(str string) {
			switch str {
			case "game":
				s.BaseSystem.World.RemoveAllSystemsExcept(s)
				s.BaseSystem.World.RemoveAllEntities()

				s.BaseSystem.World.AddSystem(&PlayerSystem{})
				s.BaseSystem.World.AddSystem(&DecelerationSystem{})
				s.BaseSystem.World.AddSystem(&EnemySystem{})
				s.BaseSystem.World.AddSystem(&ShootingSystem{})
				s.BaseSystem.World.AddSystem(&BulletSystem{})
				s.BaseSystem.World.AddSystem(&SpawnSystem{})
				s.BaseSystem.World.AddSystem(&DropSystem{})
				s.BaseSystem.World.AddSystem(&PickupSystem{})

				s.BaseSystem.World.AddEntity(entities.Player())
			}
		}
		s.hasStarted = true
	}

	if rl.IsKeyPressed(rl.KeyR) {
		s.BaseSystem.World.RemoveAllSystemsExcept(s)
		s.BaseSystem.World.RemoveAllEntities()

		s.BaseSystem.World.AddSystem(&UiSystem{})
	}

}

func (s *SceneSystem) SetBaseSystem(bs *ecs.BaseSystem) {
	s.BaseSystem = bs
}

func (s *SceneSystem) GetFilter() (ecs.FilterType, []string) {
	return ecs.All, []string{}
}
