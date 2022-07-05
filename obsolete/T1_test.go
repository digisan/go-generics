package generics

import "testing"

func TestT1FuncGen(t *testing.T) {
	type args struct {
		Tx     string
		pkgdir string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "int",
			args: args{
				Tx:     "int",
				pkgdir: "./",
			},
		},
		{
			name: "string",
			args: args{
				Tx:     "string",
				pkgdir: "./",
			},
		},
		{
			name: "bool",
			args: args{
				Tx:     "bool",
				pkgdir: "./",
			},
		},
		{
			name: "byte",
			args: args{
				Tx:     "byte",
				pkgdir: "./",
			},
		},
		{
			name: "rune",
			args: args{
				Tx:     "rune",
				pkgdir: "./",
			},
		},
		{
			name: "float64",
			args: args{
				Tx:     "float64",
				pkgdir: "./",
			},
		},
		{
			name: "interface{}",
			args: args{
				Tx:     "interface{}",
				pkgdir: "./",
			},
		},
		{
			name: "image.Point",
			args: args{
				Tx:     "image.Point",
				pkgdir: "./",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			T1FuncGen(tt.args.Tx, tt.args.pkgdir)
		})
	}
}
