package game

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/saleh/game/assets"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600

	meteorSpawnTimer  = 1 * time.Second
	meteorSpeedUpTime = 5 * time.Second

	meteorSpeedUpAmount = 0.1
	baseMeterVelocity   = 0.25
)

type Game struct {
	player              *Player
	meteorSpawnTimer    Timer
	meteorSpeedUpAmount Timer
	meteors             []*Meteor
	bullets             []*Bullet
	score               int

	baseVelocity  float64
	velocityTimer *Timer
	offscreen     *ebiten.Image
}

func NewGame() *Game {
	game := &Game{
		meteorSpawnTimer:    *NewTimer(meteorSpawnTimer),
		meteorSpeedUpAmount: *NewTimer(meteorSpeedUpTime),
		meteors:             make([]*Meteor, 0),
		bullets:             make([]*Bullet, 0),
		baseVelocity:        baseMeterVelocity,
		velocityTimer:       NewTimer(meteorSpeedUpTime),
	}

	game.player = NewPlayer(game)

	return game
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	g.velocityTimer.Update()
	if g.velocityTimer.IsReady() {
		g.velocityTimer.Reset()
		g.baseVelocity += meteorSpeedUpAmount
	}

	g.player.Update()

	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()

		m := NewMeteor(g.baseVelocity)
		g.meteors = append(g.meteors, m)
	}

	for _, m := range g.meteors {
		m.Update()
	}

	for _, b := range g.bullets {
		b.Update()
	}

	for i, m := range g.meteors {
		for j, b := range g.bullets {
			if m.Collider().Intersets(b.Collider()) {
				g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
				g.bullets = append(g.bullets[:j], g.bullets[j+1:]...)
				g.score++
			}
		}
	}

	for _, m := range g.meteors {
		if m.Collider().Intersets(g.player.Collider()) {
			g.Reset()
			break
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()

	g.player.Draw(screen)

	for _, m := range g.meteors {
		m.Draw(screen)
	}

	for _, b := range g.bullets {
		b.Draw(screen)
	}

	text.Draw(screen, fmt.Sprintf("%06d", g.score), assets.ScoreFont, ScreenWidth/2-100, 50, color.White)
}

func (g *Game) Reset() {
	g.player = NewPlayer(g)
	g.meteors = nil
	g.bullets = nil
	g.meteorSpawnTimer.Reset()
	g.score = 0
	g.velocityTimer.Reset()
	g.baseVelocity = baseMeterVelocity

}
func (g *Game) AddBullet(b *Bullet) {
	g.bullets = append(g.bullets, b)
}
