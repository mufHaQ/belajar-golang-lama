package ring

import (
	"container/ring"
	"fmt"
	"strconv"
)

func TestRing(length int, text string) {
	data := ring.New(length)

	for i := 1; i <= data.Len(); i++ {
		data.Value = text + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}
