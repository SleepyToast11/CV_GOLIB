package grayscale

func MapImageI(filter func(float64) float64, image SingleImage) SingleImage {
	var imageRet SingleImage = image.EmptyCopy()
	val, err := image.Get1DSize()
	if err != nil {
		panic(err)
	}
	for i := range val {
		val, _, _, err := image.GetVal1D(i)
		if err != nil {
			panic(err)
		}
		err = imageRet.SetVal1D(i, filter(val))
		if err != nil {
			panic(err)
		}
	}
	return imageRet
}

// window is measured by how many pixel from center: 0 is 1x1, 1 is 3x3, 2 is 5x5
func MapImageMultiI(windower func(int, int, SingleImage) [][]float64, filter func([][]float64) float64, image SingleImage) SingleImage {
	imageRet := image.EmptyCopy()
	val, err := image.Get1DSize()
	if err != nil {
		panic(err)
	}
	for i := range val {
		_, x, y, err := image.GetVal1D(i)
		if err != nil {
			panic(err)
		}
		window := windower(x, y, image)
		err = image.SetVal1D(i, filter(window))
		if err != nil {
			panic(err)
		}
	}
	return imageRet
}
