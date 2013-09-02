package bunny

import (
	"os"
	"image"
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

// Write bunny in requested resolution from cache, generate if missing
func Write(w io.Writer, x string, y string) error {
	xi, _ := strconv.Atoi(x)
	yi, _ := strconv.Atoi(y)

	img, err := generate(xi, yi)
	if err != nil {
		return err
	}

	err = jpeg.Encode(w, img, nil)
	if err != nil {
		return err
	}
	return nil
}

func generate(x int, y int) (image.Image, error) {
	dst := image.NewRGBA(image.Rect(0, 0, x, y))
	src := sources[rand.Intn(len(sources))]
	err := graphics.Thumbnail(dst, src)
	if err != nil {
		return nil, err
	}
	return dst, nil
}