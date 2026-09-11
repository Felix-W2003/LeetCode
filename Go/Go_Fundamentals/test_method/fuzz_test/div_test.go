package main

import "testing"

//func TestDiv(t *testing.T) {
//	type args struct {
//		a int
//		b int
//	}
//	tests := []struct {
//		name string
//		args args
//		want int
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := Div(tt.args.a, tt.args.b); got != tt.want {
//				t.Errorf("Div() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func FuzzDiv(f *testing.F) {
	//f.Add()	添加测试数据
	f.Fuzz(func(t *testing.T, a, b int) {
		Div(a, b)
	})
}
