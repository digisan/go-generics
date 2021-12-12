package generics

import "testing"

func TestT2FuncGen(t *testing.T) {
	type args struct {
		Tx     string
		Ty     string
		pkgdir string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "int-string",
			args: args{
				Tx:     "int",
				Ty:     "string",
				pkgdir: "./",
			},
		},
		{
			name: "int-int",
			args: args{
				Tx:     "int",
				Ty:     "int",
				pkgdir: "./",
			},
		},
		{
			name: "string-string",
			args: args{
				Tx:     "string",
				Ty:     "string",
				pkgdir: "./",
			},
		},
		{
			name: "float64-float64",
			args: args{
				Tx:     "float64",
				Ty:     "float64",
				pkgdir: "./",
			},
		},
		{
			name: "string-int",
			args: args{
				Tx:     "string",
				Ty:     "int",
				pkgdir: "./",
			},
		},
		{
			name: "string-bool",
			args: args{
				Tx:     "string",
				Ty:     "bool",
				pkgdir: "./",
			},
		},
		{
			name: "interface{}-interface{}",
			args: args{
				Tx:     "interface{}",
				Ty:     "interface{}",
				pkgdir: "./",
			},
		},
		{
			name: "image.Point",
			args: args{
				Tx:     "image.Point",
				Ty:     "image.Point",
				pkgdir: "./",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T2FuncGen(tt.args.Tx, tt.args.Ty, tt.args.pkgdir)
		})
	}
}
