package generating_random_avatar

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"
)

const avatarSize = 800

const brushSize = 80

func GenerateAvatar(count int) {
	for i := 0; i < count; i++ {
		rand.Seed(time.Now().UnixNano())

		avatar := image.NewRGBA(image.Rect(0, 0, avatarSize, avatarSize))
		PainBG(avatar, GenerateBGColor())
		PaintAvatar(avatar, GenerateAvatarColor())

		SavePNG(avatar, fmt.Sprint("avatar_", i + 1))
	}
}

func SavePNG(avatar image.Image, name string) {
	file, err := os.Create(name + ".png")
	err  = png.Encode(file, avatar)

	if err != nil {
		panic(err)
	}
}

func PaintAvatar(avatar *image.RGBA, avatarColor color.RGBA) {
	for x := 0; x < avatarSize / 2; x += brushSize {
		for y := 0; y < avatarSize; y += brushSize {
			if rand.Float32() > 0.5 {
				for rectX := x; rectX <= x + brushSize; rectX++ {
					for rectY := y; rectY <= y + brushSize; rectY++ {
						avatar.SetRGBA(rectX, rectY, avatarColor)
						avatar.SetRGBA(avatarSize - rectX, rectY, avatarColor)
					}
				}
			}
		}
	}
}

func PainBG(avatar *image.RGBA, bgColor color.RGBA) {
	for y := 0; y < avatarSize; y++ {
		for x := 0; x < avatarSize; x++ {
			avatar.SetRGBA(x, y, bgColor)
		}
	}
}

func GenerateAvatarColor() (avatarColor color.RGBA) {
	avatarColor = GenerateColor(50, 100)
	return
}

func GenerateBGColor() (bgColor color.RGBA) {
	bgColor = GenerateColor(200, 255)
	return
}

func GenerateColor(min, max float32) (color color.RGBA) {
	color.A = 255

	color.R = uint8((rand.Float32() * (max - min)) + min)
	color.G = uint8((rand.Float32() * (max - min)) + min)
	color.B = uint8((rand.Float32() * (max - min)) + min)

	return
}
