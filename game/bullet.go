package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/saleh/game/assets"
)

const (
	bulletSpeedPerSecond = 450.0
)

type Bullet struct {
	position Vector
	sprite   *ebiten.Image
	rotation float64
}

func NewBullet(position Vector, rotation float64) *Bullet {
	sprite := assets.BulletSprite

	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx() / 2)
	halfH := float64(bounds.Dy() / 2)

	position.X -= halfW
	position.Y -= halfH

	bullet := &Bullet{
		position: position,
		sprite:   sprite,
		rotation: rotation,
	}

	return bullet
}

func (b *Bullet) Update() {
	speed := bulletSpeedPerSecond / float64(ebiten.TPS())

	b.position.X += math.Sin(b.rotation) * speed
	b.position.Y += math.Cos(b.rotation) * -speed
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	bounds := b.sprite.Bounds()

	halfW := float64(bounds.Dx() / 2)
	halfH := float64(bounds.Dy() / 2)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(b.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(b.position.X, b.position.Y)

	screen.DrawImage(b.sprite, op)
}

func (b *Bullet) Collider() Rect {
	bounds := b.sprite.Bounds()

	return NewRect(
		b.position.X,
		b.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
