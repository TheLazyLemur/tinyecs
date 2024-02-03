package system

import (
	"fmt"
	"github.com/TheLazyLemur/tinyecs/ecs"
	comp "github.com/TheLazyLemur/tinyecs/example/component"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type EnemySystem struct {
	baseSystem *ecs.BaseSystem
}

func (s *EnemySystem) Update(entites []*ecs.Entity, dt float32) {
	for _, ent := range entites {
		s.updateOne(ent, entites, dt)
	}
}

func (s *EnemySystem) updateOne(ent *ecs.Entity, entities []*ecs.Entity, dt float32) {
	c, isEnemy := ent.Components[comp.PlayerFinderComponentName].(*comp.PlayerFinderComponent)
	if !isEnemy {
		return
	}

	var player *ecs.Entity
	for _, e := range entities {
		_, isEnemy := e.Components[comp.PlayerFinderComponentName]
		if !isEnemy {
			playerPos := e.Components[comp.PositionComponentName].(*comp.PositionComponent)
			if c.PlayerPosition == nil {
				c.PlayerPosition = playerPos
			}

			player = e

			break
		}
	}

	s.moveEnemy(ent, dt)

	if s.checkForCollisionWithPlayer(ent, c) {
		fmt.Println("Collision")
		s.baseSystem.World.RemoveEntity(player)
		s.baseSystem.World.RemoveEntity(ent)
		c.PlayerPosition = nil
	}

	pos := ent.Components[comp.PositionComponentName].(*comp.PositionComponent)
	rl.DrawRectangleV(rl.Vector2{X: float32(pos.X), Y: float32(pos.Y)}, rl.Vector2{X: 100, Y: 100}, rl.Red)

	if c.PlayerPosition == nil {
		return
	}

}

func (s *EnemySystem) moveEnemy(ent *ecs.Entity, dt float32) {
	pos := ent.Components[comp.PositionComponentName].(*comp.PositionComponent)
	pos.X -= 300 * dt
	if pos.X < -100 {
		s.baseSystem.World.RemoveEntity(ent)
	}
}

func (s *EnemySystem) checkForCollisionWithPlayer(ent *ecs.Entity, c *comp.PlayerFinderComponent) bool {
	pos, ok := ent.Components[comp.PositionComponentName].(*comp.PositionComponent)
	if !ok {
		return false
	}

	playerRect := rl.Rectangle{
		X:      float32(c.PlayerPosition.X),
		Y:      float32(c.PlayerPosition.Y),
		Width:  10,
		Height: 10,
	}

	enemyRect := rl.Rectangle{
		X:      float32(pos.X),
		Y:      float32(pos.Y),
		Width:  100,
		Height: 100,
	}

	return rl.CheckCollisionRecs(playerRect, enemyRect)
}

func (s *EnemySystem) SetBaseSystem(bs *ecs.BaseSystem) {
	s.baseSystem = bs
}

func (s *EnemySystem) GetFilter() (ecs.FilterType, []string) {
	return ecs.One, []string{
		comp.PositionComponentName,
		comp.PlayerFinderComponentName,
	}
}
