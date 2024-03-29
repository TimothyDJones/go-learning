# https://gochronicles.dev/posts/algorithms/bubble-sort/
# To run test:
#   go test -v

package sorting

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	nums := []int{23, 1, 53, 543, 98, 4}
	fmt.Println("nums: ", nums)

	got := BubbleSort(nums)
	want := []int{1, 4, 23, 53, 98, 543}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v, given: %v", got, want, nums)
	}
}
