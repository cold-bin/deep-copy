// @author cold bin
// @date 2022/9/4

package test

import (
	dc "github.com/cold-bin/deep-copy"
	"reflect"
	"testing"
)

func TestArrCopy(t *testing.T) {
	//x:=[4]int{1, 2, 3, 4}
	type args struct {
		src interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
		//missing type in composite literal
		//{name: "easy case", args: , want: [4]int{1, 2, 3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dc.Copy(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}
