package main

import (
    "fmt"
    rl "github.com/gen2brain/raylib-go/raylib"
    rg "github.com/gen2brain/raylib-go/raygui"
)

type MouseBtn int
const (
	MouseBtnLeft    = MouseBtn(rl.MouseButtonLeft)
	MouseBtnRight   = MouseBtn(rl.MouseButtonRight)
	MouseBtnMiddle  = MouseBtn(rl.MouseButtonMiddle)
	MouseBtnSide    = MouseBtn(rl.MouseButtonSide)
	MouseBtnExtra   = MouseBtn(rl.MouseButtonExtra)
	MouseBtnForward = MouseBtn(rl.MouseButtonForward)
	MouseBtnBack    = MouseBtn(rl.MouseButtonBack)
)

type DefaultStyle struct {
    theme_color rl.Color
    bg_color    rl.Color
    fg_color    rl.Color
    font        rl.Font
    font_size   int
}
var default_style = DefaultStyle{
    theme_color: rl.Color{235, 118, 48, 225},
    bg_color:    rl.Color{32, 29, 42, 225},
    fg_color:    rl.Color{127, 169, 162, 225},
    font_size:   20,
}

func get_win_height() int {
    return rl.GetScreenHeight()
}
func get_win_width() int {
    return rl.GetScreenWidth()
}

func get_mouse_x() int32 {
    return rl.GetMouseX()
}
func get_mouse_y() int32 {
    return rl.GetMouseY()
}

func is_mouse_btn_down(btn MouseBtn) bool {
    return rl.IsMouseButtonDown(rl.MouseButton(btn))
}

func draw_text_main(msg string, x float32, y float32, size float32) {
    rl.DrawTextEx(default_style.font, msg, rl.Vector2{ x, y }, size, 0, default_style.theme_color)
}
func draw_text(msg string, x float32, y float32, size float32) {
    rl.DrawTextEx(default_style.font, msg, rl.Vector2{ x, y }, size, 0, default_style.theme_color)
}

func draw_img(img *rl.Image, width int32, height int32, x int32, y int32) {
    mouse_x := get_mouse_x()
    mouse_y := get_mouse_y()
    img_area :=  mouse_x > x && mouse_x < width && mouse_y > y && mouse_y < height

    rl.ImageResize(img, width, height)
    img_tex := rl.LoadTextureFromImage(img)

    if is_mouse_btn_down(MouseBtnLeft) && img_area {
        rg.Panel(rl.Rectangle{ float32(x)-1, float32(y)-1, float32(width)+2, float32(height)+2 }, "")
    }

    if img_area {
        rl.DrawTexture(img_tex, x, y, rl.Gray)
    } else {
        rl.DrawTexture(img_tex, x, y, rl.White)
    }
}

func view_style_init() {
    rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(800, 450, viewInfo.name)
	rl.SetTargetFPS(200)
    rg.LoadStyleDefault()

    rg.LoadStyle("./assets/cyber.rgs")
    default_style.font = rl.LoadFontEx("./assets/Urbanist.ttf", 100, nil, 0)
    rl.GenTextureMipmaps(&default_style.font.Texture)
    rl.SetTextureFilter(default_style.font.Texture, rl.FilterBilinear)
    rg.SetFont(default_style.font)

    rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 20)
}

func view_logo(img *rl.Image) {
    draw_img(img, 350, 80, 10, 10)
}


func view_stats(x float32, y float32, show_stats *bool) {
    show_stats_btn := rg.Button(rl.Rectangle{x, y, 100, 30}, "Show Stats")
    if show_stats_btn {
        *show_stats = !*show_stats
    }
    if *show_stats {
        stats_win_pos_x := x
        stats_win_pos_y := y-220

        *show_stats = !rg.WindowBox(rl.Rectangle{ stats_win_pos_x, stats_win_pos_y, 200, 250 }, "Stats")

        stats_win_pos_x += 10
        stats_win_pos_y += 30

        stats_msg := fmt.Sprint("Window Width: ", rl.GetScreenWidth(), "\n",
            "Window Height: ", rl.GetScreenHeight(), "\n",
            "Window Position X: ", rl.GetWindowPosition().X, "\n",
            "Window Position Y: ", rl.GetWindowPosition().Y, "\n",
            "Mouse Position X: ", rl.GetMouseX(), "\n",
            "Mouse Position Y: ", rl.GetMouseY(), "\n",
            "Mouse Wheel X: ", rl.GetMouseWheelMoveV().X, "\n",
            "Mouse Wheel Y: ", rl.GetMouseWheelMoveV().Y, "\n",
            "FPS: ", rl.GetFPS(), "\n",
            "Key Pressed Code: ", rl.GetKeyPressed(), "\n")
        draw_text(stats_msg, stats_win_pos_x, stats_win_pos_y, 20)
    }

}

func view_post_box(msg *string, state *bool) {
    rg.SetStyle(rg.TEXTBOX, rg.TEXT_PADDING, 5)
    rg.Panel(rl.Rectangle{ 360, 10, 400, 100 }, "")
    if (rg.TextBox(rl.Rectangle{ 360, 10, 400, 100 }, msg, 100, *state)) {
        *state = !*state
    }
}

func view_start() {
    rl.SetTraceLogLevel(rl.LogNone)
    view_style_init()

	show_stats := false
    // textbox_edit := false

    // post_box_placeholder := "Write what's on your mind!"
    logo_img := rl.LoadImage("./assets/logo.png")
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
            mouse_cell := rl.Vector2{ float32(get_mouse_x()), float32(get_mouse_y()) }
            win_height := get_win_height()

            rl.ClearBackground(viewInfo.bg_color)
            rg.Grid(rl.Rectangle{ 0, 0, float32(get_win_width()), float32(get_win_height()) }, "hello", 50, 10, &mouse_cell)
            view_logo(logo_img)
            view_stats(10, float32(win_height)-40, &show_stats)

            // view_post_box(&post_box_placeholder, &textbox_edit)
            // rl.ImageClearBackground(logo_img, rl.Red)
		rl.EndDrawing()
	}

    rl.UnloadImage(logo_img)
    rl.UnloadFont(default_style.font)
	rl.CloseWindow()
}
