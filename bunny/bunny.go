package bunny

import (
	"os"
	"image"
	"image/draw"
	"image/jpeg"
	"code.google.com/p/graphics-go/graphics"
	"io"
	"strconv"
)

var source image.Image

func init() {
	source = load("sources/bunny.jpg")
}

func load(path string) image.Image {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	return img
}

func New(x string, y string) *Bunny {
	xi, _ := strconv.Atoi(x)
	yi, _ := strconv.Atoi(y)
	return &Bunny{x:xi,y:yi}
}

type Bunny struct {
	y int
	x int
	img draw.Image
}

func (b *Bunny) Thumbnail() error {
	b.img = image.NewRGBA(image.Rect(0, 0, b.x, b.y))
	err := graphics.Thumbnail(b.img, source)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bunny) Write(w io.Writer) error {
	err := jpeg.Encode(w, b.img, nil)
	if err != nil {
		return err
	}
	return nil
}