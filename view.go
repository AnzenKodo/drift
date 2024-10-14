package main

import (
	"fmt"
	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var font rl.Font

func view_style_init() {
	rl.InitWindow(800, 450, viewInfo.name)
	rl.SetTargetFPS(60)
    rg.LoadStyleDefault()

    font = rl.LoadFontEx("./assets/Urbanist.ttf", 100, nil, 0)
    rl.GenTextureMipmaps(&font.Texture);
    rl.SetTextureFilter(font.Texture, rl.FilterBilinear);
    rg.SetFont(font)

    // rg.LoadStyle("./assets/cyber.rgs")
    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 20)
}

func view_start() {
    rl.SetTraceLogLevel(rl.LogNone)
    view_style_init()

	var button bool
	fontPosition := rl.NewVector2(100, 100)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

        rl.ClearBackground(viewInfo.bg_color)
		button = rg.Button(rl.NewRectangle(50, 150, 100, 40), "Click")
		rl.DrawTextEx(font, "Hello", fontPosition, float32(50), 0, rl.Color{0, 0, 0, 255})

		if button {
			fmt.Println("Clicked on button")
		}

		rl.EndDrawing()
	}

    rl.UnloadFont(font);
	rl.CloseWindow()
}
