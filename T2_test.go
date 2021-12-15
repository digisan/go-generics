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
		// {
		// 	name: "int-string",
		// 	args: args{
		// 		Tx:     "int",
		// 		Ty:     "string",
		// 		pkgdir: "./",
		// 	},
		// },
		// {
		// 	name: "rune-string",
		// 	args: args{
		// 		Tx:     "rune",
		// 		Ty:     "string",
		// 		pkgdir: "./",
		// 	},
		// },
		// {
		// 	name: "string-rune",
		// 	args: args{
		// 		Tx:     "string",
		// 		Ty:     "rune",
		// 		pkgdir: "./",
		// 	},
		// },
		// {
		// 	name: "int-int",
		// 	args: args{
		// 		Tx:     "int",
		// 		Ty:     "int",
		// 		pkgdir: "./",
		// 	},
		// },
		// {
		// 	name: "string-string",
		// 	args: args{
		// 		Tx:     "string",
		// 		Ty:     "string",
		// 		pkgdir: "./",
		// 	},
		// },
		// {
		// 	name: "float64-float64",
		// 	args: args{
		// 		Tx:     "float64",
		// 		Ty:     "float64",
		// 		pkgdir: "./",
		// 	},
		// },
		// {
		// 	name: "string-int",
		// 	args: args{
		// 		Tx:     "string",
		// 		Ty:     "int",
		// 		pkgdir: "./",
		// 	},
		// },
		// {
		// 	name: "string-float64",
		// 	args: args{
		// 		Tx:     "string",
		// 		Ty:     "float64",
		// 		pkgdir: "./",
		// 	},
		// },
		// {
		// 	name: "string-bool",
		// 	args: args{
		// 		Tx:     "string",
		// 		Ty:     "bool",
		// 		pkgdir: "./",
		// 	},
		// },
		// {
		// 	name: "interface{}-interface{}",
		// 	args: args{
		// 		Tx:     "interface{}",
		// 		Ty:     "interface{}",
		// 		pkgdir: "./",
		// 	},
		// },
		// {
		// 	name: "image.Point-image.Point",
		// 	args: args{
		// 		Tx:     "image.Point",
		// 		Ty:     "image.Point",
		// 		pkgdir: "./",
		// 	},
		// },
		// {
		// 	name: "string-interface{}",
		// 	args: args{
		// 		Tx:     "string",
		// 		Ty:     "interface{}",
		// 		pkgdir: "./",
		// 	},
		// },
		// {
		// 	name: "string-byte",
		// 	args: args{
		// 		Tx:     "string",
		// 		Ty:     "byte",
		// 		pkgdir: "./",
		// 	},
		// },
		// {
		// 	name: "byte-string",
		// 	args: args{
		// 		Tx:     "byte",
		// 		Ty:     "string",
		// 		pkgdir: "./",
		// 	},
		// },
		// {
		// 	name: "byte-int",
		// 	args: args{
		// 		Tx:     "byte",
		// 		Ty:     "int",
		// 		pkgdir: "./",
		// 	},
		// },
		// {
		// 	name: "int-byte",
		// 	args: args{
		// 		Tx:     "int",
		// 		Ty:     "byte",
		// 		pkgdir: "./",
		// 	},
		// },
		// {
		// 	name: "byte-float64",
		// 	args: args{
		// 		Tx:     "byte",
		// 		Ty:     "float64",
		// 		pkgdir: "./",
		// 	},
		// },
		// {
		// 	name: "byte-rune",
		// 	args: args{
		// 		Tx:     "byte",
		// 		Ty:     "rune",
		// 		pkgdir: "./",
		// 	},
		// },
		{
			name: "float64-int",
			args: args{
				Tx:     "float64",
				Ty:     "int",
				pkgdir: "./",
			},
		},
		{
			name: "float64-uint",
			args: args{
				Tx:     "float64",
				Ty:     "uint",
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
