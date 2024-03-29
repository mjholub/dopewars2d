package core

import (
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/lafriks/go-tiled"
	"github.com/lafriks/go-tiled/render"
)

func (g Game) Close() {
	os.Exit(0)
}

func (g Game) Update() error {
	if !g.init {
		g.init = true
	}

	g.gameUI.UpdateWithSize(ebiten.WindowSize())

	return nil
}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 960, 540
}

func (g Game) Draw(screen *ebiten.Image) {
	level0, err := loadAssets()
	if err != nil {
		fmt.Printf("fatal runtime error: %v", err)
		os.Exit(2)
	}

	screen.DrawImage(level0, nil)
}

// FIXME: unsupported orientation error here
func loadAssets() (*ebiten.Image, error) {
	gameMap, err := tiled.LoadFile("assets/tiles/Building Tiles/map.tmx")
	if err != nil {
		return nil, fmt.Errorf("error opening map: %v", err)
	}
	renderer, err := render.NewRenderer(gameMap)
	if err != nil {
		return nil, fmt.Errorf("error starting renderer: %w", err)
	}

	if err = renderer.RenderVisibleLayers(); err != nil {
		return nil, fmt.Errorf("error rendering layers: %w", err)
	}

	img := renderer.Result

	renderer.Clear()

	return ebiten.NewImageFromImage(img), nil
}

func (h *GameHandler) DrawMenu(screen *ebiten.Image) {
	if bg == nil {
		if err := initBackground(); err != nil {
			h.log.Fatal().Err(err).Msg("Failed to load background image")
		}
	}

	screen.DrawImage(bg, nil)
	h.g.UI.DrawMenu(*ebiten.NewImageFromImage(bg))

	buttonNames := []string{"newgame", "loadsave", "donate", "issues", "quit"}
	buttonPositions := []struct{ x, y, w, h int }{
		{340, 150, 200, 50},
		{340, 200, 200, 50},
		{340, 250, 200, 50},
		{340, 300, 200, 50},
		{340, 350, 280, 50},
	}

	if err := initButtonImages(); err != nil {
		h.log.Fatal().Err(err).Msg("Failed to initialize button images")
	}

	buttons := make([]button, len(buttonNames))
	for i := range buttonNames {
		buttons[i] = button{
			img:      buttonImages[i*2],
			imgHover: buttonImages[i*2+1],
			x:        &buttonPositions[i].x,
			y:        &buttonPositions[i].y,
			w:        &buttonPositions[i].w,
			h:        &buttonPositions[i].h,
		}
	}
	for _, btn := range buttons {
		pos := &ebiten.DrawImageOptions{}
		pos.GeoM.Translate(float64(*btn.x), float64(*btn.y))
		screen.DrawImage(btn.img, pos)
		if mouseOverButton(*btn.x, *btn.y, *btn.w, *btn.h) {
			screen.DrawImage(btn.imgHover, pos)
		}
	}
	if err := h.setupUI(); err != nil {
		h.log.Fatal().Err(err).Msg("failed to set up menu")
	}

	ebiten.RunGame(h.g)
}

// create a function that checks if the mouse is over a button
func mouseOverButton(x, y, width, height int) bool {
	// get the mouse position
	mouseX, mouseY := ebiten.CursorPosition()
	// check if the mouse is within the button's x and y bounds
	if mouseX >= x && mouseX <= x+width {
		if mouseY >= y && mouseY <= y+height {
			return true
		}
	}
	return false
}
