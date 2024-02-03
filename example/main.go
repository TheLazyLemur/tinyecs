package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/TheLazyLemur/tinyecs/ecs"
	sys "github.com/TheLazyLemur/tinyecs/example/system"

	rl "github.com/gen2brain/raylib-go/raylib"
)


var lock = &sync.Mutex{}

func printMemoryUsage() {
	lock.Lock()
	defer lock.Unlock()

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
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			printMemoryUsage()
		}
	}()

	w := ecs.World{}

	w.AddSystem(&sys.SceneSystem{})

	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	rl.SetConfigFlags(rl.FlagVsyncHint)

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		lock.Lock()

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		w.Update(rl.GetFrameTime())

		rl.DrawFPS(10, 10)
		rl.EndDrawing()

		lock.Unlock()
	}
}
