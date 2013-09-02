package bunny

import (
	"os"
	"image"
	"image/draw"
	"image/jpeg"
	"code.google.com/p/graphics-go/graphics"
	"io"
	"strconv"
	"path/filepath"
	"math/rand"
)

var sources []image.Image

func init() {
	filepath.Walk("sources", load)
}

func load(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	sources = append(sources, img)
	return nil
}

func randomSource() image.Image {
	return sources[rand.Intn(len(sources))]
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
	err := graphics.Thumbnail(b.img, randomSource())
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