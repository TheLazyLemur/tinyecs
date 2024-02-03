package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/TheLazyLemur/tinyecs/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func printMemoryUsage() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	fmt.Printf("Memory Usage:\n")
	fmt.Printf("Alloc: %v MB\n", bToMb(memStats.Alloc))
	fmt.Printf("TotalAlloc: %v MB\n", bToMb(memStats.TotalAlloc))
	fmt.Printf("Sys: %v MB\n", bToMb(memStats.Sys))
	fmt.Printf("NumGC: %v\n", memStats.NumGC)
	fmt.Printf("LastGC: %v seconds ago\n", time.Since(time.Unix(0, int64(memStats.LastGC))))
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func main() {
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			printMemoryUsage()
		}
	}()

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
		w.Update(rl.GetFrameTime())

		rl.EndDrawing()
	}
}
