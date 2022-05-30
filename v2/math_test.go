package v2

import (
	"reflect"
	"testing"
)

func TestHasOverlapped(t *testing.T) {
	type args struct {
		sn [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				sn: [][]int{{2, 5}, {5, 6}, {8, 10}, {10, 11}},
			},
			want: false,
		},
		{
			args: args{
				sn: [][]int{{8, 10}, {10, 11}, {2, 5}, {5, 9}},
			},
			want: true,
		},
		{
			args: args{
				sn: [][]int{{8, 10}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasOverlapped(tt.args.sn...); got != tt.want {
				t.Errorf("HasOverlapped() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpanJoin(t *testing.T) {
	type args struct {
		s1      []int
		s2      []int
		ocJoint bool
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				s1:      []int{3, 6},
				s2:      []int{4, 9},
				ocJoint: false,
			},
			want:  []int{3, 9},
			want1: true,
		},
		{
			args: args{
				s1:      []int{3, 6},
				s2:      []int{7, 9},
				ocJoint: false,
			},
			want:  nil,
			want1: false,
		},
		{
			args: args{
				s1:      []int{3, 6},
				s2:      []int{6, 9},
				ocJoint: true,
			},
			want:  []int{3, 9},
			want1: true,
		},
		{
			args: args{
				s1:      []int{3, 6},
				s2:      []int{6, 9},
				ocJoint: false,
			},
			want:  nil,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := SpanJoin(tt.args.s1, tt.args.s2, tt.args.ocJoint)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SpanJoin() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SpanJoin() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
