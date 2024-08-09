package ptr

import (
	"errors"
	"fmt"
	"unsafe"
)

func GetElement(arr []int, idx int) (int, error) {
	if arr == nil {
		return -1, errors.New("arr is nill")
	}
	if idx < 0 || idx >= len(arr) {
		return -1, fmt.Errorf("idx is not correct: %d", idx)
	}
	ptr := unsafe.Pointer(&arr[0])
	step := unsafe.Sizeof(arr[0])
	ptr = unsafe.Pointer(uintptr(ptr) + step*uintptr(idx))

	result := *(*int)(ptr)
	return result, nil
}
