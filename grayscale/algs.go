package grayscale

import (
	"io"
	"strconv"
)

func (image *GrayscaleImage) MapImage(filter func(float64) float64) {
	imageTemp := MapImageI(filter, image)
	val, err := image.Get1DSize()
	if err != nil {
		panic(err)
	}
	for i := range val {
		val, _, _, err := imageTemp.GetVal1D(i)
		if err != nil {
			panic(err)
		}
		err = image.SetVal1D(i, val)
		if err != nil {
			panic(err)
		}
	}

}

func (image *GrayscaleImage) MapImageMulti(windower func(idxX, idxY int, image SingleImage) [][]float64, filter func([][]float64) float64) {
	imageTemp := MapImageMultiI(windower, filter, image)
	val, err := image.Get1DSize()
	if err != nil {
		panic(err)
	}
	for i := range val {
		val, _, _, err := imageTemp.GetVal1D(i)
		if err != nil {
			panic(err)
		}
		err = image.SetVal1D(i, val)
		if err != nil {
			panic(err)
		}
	}
}

func (image *GrayscaleImage) WriteImage(writer io.Writer) error {
	data := "P2\n"
	data += strconv.Itoa(image.x) + " " + strconv.Itoa(image.y) + "\n255\n"
	for v := range image.data {
		data += strconv.Itoa(int(image.data[v])) + " "
	}
	data = data[:len(data)]
	_, err := writer.Write([]byte(data))
	return err
}
