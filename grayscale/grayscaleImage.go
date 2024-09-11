package grayscale

import (
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

type GrayscaleImage struct {
	x, y int
	data []float64
}

type Image interface {
	WriteImage(io.Writer) error
	MapImage(func(float64) float64) error
	MapImageMulti(windower func(idxX, idxY int, image SingleImage) [][]float64, filter func([][]float64) float64)
}

func (image *GrayscaleImage) Get1DSize() (int, error) {
	valX, err := image.GetSizeX()
	if err != nil {
		return 0, err
	}
	valY, err := image.GetSizeY()
	if err != nil {
		return 0, err
	}
	return valX * valY, nil
}

func (image *GrayscaleImage) GetVal1D(index int) (val float64, idxX int, idxY int, err error) {
	if index >= len(image.data) {
		err = errors.New("index out of range")
		return 0, 0, 0, err
	}
	if index < 0 {
		err = errors.New("index cannot be negative")
		return 0, 0, 0, err
	}
	idxX, idxY = index%image.x, index/image.x
	val, err = image.GetVal(idxX, idxY)
	err = nil
	return
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
	ints = ints[2:]
	if len(ints) != xsize*ysize {
		return GrayscaleImage{}, errors.New("bad number of pixels")
	}
	ints = ints[:2]
	floats := make([]float64, len(ints))
	for i := range ints {
		floats[i] = float64(ints[i])
	}
	return GrayscaleImage{xsize, ysize, floats}, nil
}

func (image *GrayscaleImage) GetSizeX() (int, error) {
	if image.data == nil {
		return 0, errors.New("image not initialized")
	}
	return image.x, nil
}

func (image *GrayscaleImage) GetSizeY() (int, error) {
	if image.data == nil {
		return 0, errors.New("image not initialized")
	}
	return image.y, nil
}

func (image *GrayscaleImage) GetVal(x int, y int) (float64, error) {
	if image.data == nil {
		return 0, errors.New("image not initialized")
	}
	if x < 0 || y < 0 {
		return 0, errors.New("index negative")
	}
	if x >= image.x || y >= image.y {
		return 0, errors.New("index out of range")
	}

	return image.data[x+y*image.x], nil
}

// Ensure that GrayscaleImage satisfies the Image interface.
var _ SingleImage = &GrayscaleImage{}

func (image *GrayscaleImage) SetVal(x int, y int, val float64) error {
	if image.data == nil {
		return errors.New("image not initialized")
	}
	if x < 0 || y < 0 {
		return errors.New("index negative")
	}
	if x >= image.x || y >= image.y {
		return errors.New("index out of range")
	}
	image.data[x+y*image.x] = val
	return nil
}

func (image *GrayscaleImage) SetVal1D(idx int, val float64) error {
	if image.data == nil {
		return errors.New("image not initialized")
	}
	if idx < 0 || idx >= len(image.data) {
		return errors.New("index out of range")
	}
	image.data[idx] = val
	return nil
}

func (image *GrayscaleImage) GetMinMax() (float64, float64, error) {
	if image.data == nil {
		return 0, 0, errors.New("image not initialized")
	}
	var minVal = math.MaxFloat64
	var maxVal = math.MaxFloat64 * -1
	for _, v := range image.data {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}
	return minVal, maxVal, nil
}

func GetMinMax(data []float64) (minVal float64, maxVal float64, err error) {
	if data == nil {
		return 0, 0, errors.New("no data")
	}
	minVal = math.MaxFloat64
	maxVal = math.MaxFloat64 * -1
	for _, v := range data {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}
	return minVal, maxVal, nil
}

func (image *GrayscaleImage) EmptyCopy() SingleImage {
	valSize, err := image.Get1DSize()
	if err != nil {
		panic(err)
	}
	valX, err := image.GetSizeX()
	if err != nil {
		panic(err)
	}
	valY, err := image.GetSizeY()
	if err != nil {
		panic(err)
	}
	return &GrayscaleImage{data: make([]float64, valSize), x: valX, y: valY}
}
