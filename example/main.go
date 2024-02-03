package main

import (
	"github.com/TheLazyLemur/tinyecs/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	w := ecs.NewWorld()

	w.AddEntity(NewPlayer())
	w.AddEntity(NewEnemy())

	w.AddSystem(PlayerSystem{})
	w.AddSystem(BulletSystem{})
	w.AddSystem(EnemySystem{})
	w.AddSystem(BulletCollisionSystem{})

	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
		w.Update()

		rl.EndDrawing()
	}
}
