package mutualfund

import (
	"fmt"
	"reflect"
)

type Indicator struct {
	Beta          float32
	JensenAlpha   float32
	SharpeRatio   float32
	TreynorRatio  float32
	TrackingError float32
}

func (i *Indicator) toArrayString() []string {
	arr := []string{}

	v := reflect.ValueOf(*i)

	for i := 0; i < v.NumField(); i++ {
		fieldvalue := v.Field(i).Interface()
		arr = append(arr, fmt.Sprintf("%f", fieldvalue))
	}

	return arr
}
