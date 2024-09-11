package grayscale

type SingleImage interface {
	GetVal(int, int) (float64, error)
	SetVal(int, int, float64) error
	SetVal1D(int, float64) error
	GetVal1D(int) (val float64, idxX int, idxY int, err error)
	GetMinMax() (float64, float64, error)
	GetSizeX() (int, error)
	GetSizeY() (int, error)
	Get1DSize() (int, error)
	EmptyCopy() SingleImage
}
