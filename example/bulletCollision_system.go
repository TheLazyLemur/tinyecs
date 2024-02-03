package main

import (
	"github.com/TheLazyLemur/tinyecs/ecs"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type BulletCollisionSystem struct{}

func (b BulletCollisionSystem) Update(world *ecs.World, entity []ecs.Entity) {
	for _, ent := range entity {

		bulletRec := rl.Rectangle{}
		if ent.Components[BulletComponentID] != nil {
			bulletSize := ent.Components[SizeComponentID].(*SizeComponent)
			bulletPos := ent.Components[PositionComponentID].(*PositionComponent)

			bulletRec = rl.NewRectangle(float32(bulletPos.x), float32(bulletPos.y), float32(bulletSize.w), float32(bulletSize.h))
		}

		for _, e := range entity {
			if e.Components[EnemyComponentID] != nil {
				enemySize := e.Components[SizeComponentID].(*SizeComponent)
				enemyPos := e.Components[PositionComponentID].(*PositionComponent)

				enemyRec := rl.NewRectangle(float32(enemyPos.x), float32(enemyPos.y), float32(enemySize.w), float32(enemySize.h))
				if rl.CheckCollisionRecs(bulletRec, enemyRec) {
					world.RemoveEntity(ent.ID)
					world.RemoveEntity(e.ID)
				}
			}
		}
	}
}

func (b BulletCollisionSystem) GetFilters() (ecs.FilterType, []string) {
	return ecs.One, []string{EnemyComponentID, BulletComponentID}
}
