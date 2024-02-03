package system

import (
	"github.com/TheLazyLemur/tinyecs/ecs"
	"github.com/TheLazyLemur/tinyecs/example/entities"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type SpawnSystem struct {
	BaseSystem *ecs.BaseSystem
	SpawnTime  float32
}

func (s *SpawnSystem) Update(entites []*ecs.Entity, dt float32) {
	s.SpawnTime -= dt
	if s.SpawnTime < 0 {
		s.SpawnTime = 1
		x := float32(800 - 100)
		y := float32(rl.GetRandomValue(0, 450))
		s.BaseSystem.World.AddEntity(entities.Enemy(x, y))
	}
}

func (s *SpawnSystem) SetBaseSystem(bs *ecs.BaseSystem) {
	s.BaseSystem = bs
}

func (s *SpawnSystem) GetFilter() (ecs.FilterType, []string) {
	return ecs.One, []string{}
}
