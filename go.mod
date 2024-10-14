module drift-cli

go 1.22.0

require github.com/coder/websocket v1.8.12 // direct

require (
	github.com/gen2brain/raylib-go/raygui v0.0.0-00010101000000-000000000000
	github.com/gen2brain/raylib-go/raylib v0.0.0-20231118125650-a1c890e8cbfc
	github.com/nbd-wtf/go-nostr v0.13.2
)

require (
	github.com/SaveTheRbtz/generic-sync-map-go v0.0.0-20230201052002-6c5833b989be // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.3.2 // indirect
	github.com/btcsuite/btcd/chaincfg/chainhash v1.0.2 // indirect
	github.com/decred/dcrd/crypto/blake256 v1.0.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.1.0 // indirect
	github.com/ebitengine/purego v0.7.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/valyala/fastjson v1.6.4 // indirect
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842 // indirect
	golang.org/x/net v0.0.0-20211015210444-4f30a5c0130f // indirect
	golang.org/x/sys v0.20.0 // indirect
)

replace (
	github.com/gen2brain/raylib-go/raygui => ./libs/raygui
	github.com/gen2brain/raylib-go/raylib => ./libs/raylib
)
