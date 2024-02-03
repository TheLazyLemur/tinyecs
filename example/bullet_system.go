package main

import (
	"github.com/TheLazyLemur/tinyecs/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)


type BulletSystem struct{}

func (b BulletSystem) Update(world *ecs.World, dt float32, entity []ecs.Entity) {
	for _, e := range entity {
		pc := e.Components[PositionComponentID].(*PositionComponent)
		sz := e.Components[SizeComponentID].(*SizeComponent)
		tc := e.Components[TimerComponentID].(*TimerComponent)
		sc := e.Components[SpeedComponentID].(*SpeedComponent)

		tc.timer -= rl.GetFrameTime()
		if tc.timer <= 0 {
			world.RemoveEntity(e.ID)
		}

		pc.x += sc.speed * rl.GetFrameTime()

		rl.DrawRectangle(int32(pc.x), int32(pc.y), int32(sz.w), int32(sz.h), rl.Red)
	}
}

func (b BulletSystem) GetFilters() (ecs.FilterType, []string) {
	return ecs.All, []string{PositionComponentID, SizeComponentID, TimerComponentID, SpeedComponentID}
}
