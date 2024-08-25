package main

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type GrayscaleImage struct {
	x, y int
	data []int
}

type Image interface {
	getSizeX() (int, error)
	getSizeY() (int, error)
	getVal(int, int) (int, error)
	writeImage(io.Writer) error
	mapImage(func(int, []int) int, []int)
	mapImageMulti(func([][]int, int, int, []int) int, int, int, []int)
}

func (image GrayscaleImage) mapImageMulti(filter func([][]int, int, int, []int) int, xSize int, ySize int, opts []int) {
	a := make([][]int, ySize)
	for i := range a {
		a[i] = make([]int, xSize)
	}
	for _, i := range image.data {
		x, y := i%image.x, i/image.x
		for i := range a {
			for j := range a[i] {
				valUint, err := image.getVal(x, y)
				val := int(valUint)
				if err != nil {
					val = -1
				}
				a[i][j] = val
			}
		}
		image.getVal(x, y)
		image.data[i] = filter(a, xSize, ySize, opts)
	}
}

func (image GrayscaleImage) mapImage(filter func(int, []int) int, opts []int) {
	for i, v := range image.data {
		image.data[i] = filter(v, opts)
	}
}

func (image GrayscaleImage) writeImage(writer io.Writer) error {
	data := "P2\n"
	data += string(image.x) + " " + string(image.y) + "\n255\n"
	for v := range image.data {
		data += string(image.data[v]) + " "
	}
	data = data[:-1]
	_, err := writer.Write([]byte(data))
	return err
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func split(r rune) bool {
	s := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ \n"
	return strings.ContainsRune(s, r)
}

func NewGrayscaleImage(data string) (GrayscaleImage, error) {
	if data[:2] != "P2" {
		return GrayscaleImage{}, errors.New(fmt.Sprint("GrayscaleImage data must start with 'P2'"))
	}
	data = data[2:]
	dataSplit := strings.FieldsFunc(data, split)
	ints := make([]int, len(dataSplit))
	for v := range dataSplit {
		i, err := strconv.Atoi(dataSplit[v])
		if err != nil {
			return GrayscaleImage{}, err
		}
		ints = append(ints, i)
	}
	xsize := ints[0]
	ysize := ints[1]
	if len(ints)-3 != xsize*ysize {
		return GrayscaleImage{}, errors.New("bad number of pixels")
	}
	return GrayscaleImage{ints[0], ints[1], ints[3:]}, nil
}

func (image GrayscaleImage) getSizeX() (int, error) {
	if image.data == nil {
		return 0, errors.New("image not initialized")
	}
	return image.x, nil
}

func (image GrayscaleImage) getSizeY() (int, error) {
	if image.data == nil {
		return 0, errors.New("image not initialized")
	}
	return image.y, nil
}

func (image GrayscaleImage) getVal(x int, y int) (int, error) {
	if image.data == nil {
		return 0, errors.New("image not initialized")
	}
	if x >= image.x || y >= image.y {
		return 0, errors.New("index out of range")
	}

	return image.data[x+y*image.x], nil
}

// Ensure that GrayscaleImage satisfies the Image interface.
var _ Image = GrayscaleImage{}
