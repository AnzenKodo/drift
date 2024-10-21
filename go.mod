module drift-cli

go 1.22.0

require github.com/coder/websocket v1.8.12 // direct

require (
	github.com/gen2brain/raylib-go/raygui v0.0.0-00010101000000-000000000000
	github.com/gen2brain/raylib-go/raylib v0.0.0-20231118125650-a1c890e8cbfc
	github.com/nbd-wtf/go-nostr v0.13.2
)

require (
	github.com/AllenDang/cimgui-go v1.0.2 // indirect
	github.com/AllenDang/giu v0.9.0 // indirect
	github.com/AllenDang/go-findfont v0.0.0-20200702051237-9f180485aeb8 // indirect
	github.com/SaveTheRbtz/generic-sync-map-go v0.0.0-20230201052002-6c5833b989be // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.3.2 // indirect
	github.com/btcsuite/btcd/chaincfg/chainhash v1.0.2 // indirect
	github.com/decred/dcrd/crypto/blake256 v1.0.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.1.0 // indirect
	github.com/ebitengine/gomobile v0.0.0-20240911145611-4856209ac325 // indirect
	github.com/ebitengine/hideconsole v1.0.0 // indirect
	github.com/ebitengine/purego v0.8.0 // indirect
	github.com/faiface/mainthread v0.0.0-20171120011319-8b78f0a41ae3 // indirect
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20231223183121-56fa3ac82ce7 // indirect
	github.com/go-text/typesetting v0.2.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hajimehoshi/ebiten v1.12.12 // indirect
	github.com/hajimehoshi/ebiten/v2 v2.8.1 // indirect
	github.com/jezek/xgb v1.1.1 // indirect
	github.com/mazznoer/csscolorparser v0.1.5 // indirect
	github.com/napsy/go-css v0.0.0-20221107082635-4ed403047a64 // indirect
	github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8 // indirect
	github.com/sahilm/fuzzy v0.1.1 // indirect
	github.com/valyala/fastjson v1.6.4 // indirect
	github.com/zeozeozeo/ebitengine-microui-go v1.0.0 // indirect
	github.com/zeozeozeo/microui-go v0.0.0-20240828161410-f386f91fa9b0 // indirect
	golang.design/x/hotkey v0.4.1 // indirect
	golang.design/x/mainthread v0.3.0 // indirect
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842 // indirect
	golang.org/x/exp/shiny v0.0.0-20241009180824-f66d83c29e7c // indirect
	golang.org/x/image v0.20.0 // indirect
	golang.org/x/mobile v0.0.0-20231127183840-76ac6878050a // indirect
	golang.org/x/net v0.0.0-20211015210444-4f30a5c0130f // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/text v0.18.0 // indirect
	gopkg.in/eapache/queue.v1 v1.1.0 // indirect
)

replace (
	github.com/gen2brain/raylib-go/raygui => ./libs/raygui
	github.com/gen2brain/raylib-go/raylib => ./libs/raylib
)
