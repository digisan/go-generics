package generics

import (
	"os"
	"strings"

	spt "github.com/digisan/go-generics/support"
)

var mAliasTyp = map[string]string{
	"object": "interface{}",
	"double": "float64",
}

// type -> package name
var mTypPkg = map[string]string{
	"int":         "i32",
	"int8":        "i8",
	"int16":       "i16",
	"int32":       "i32",
	"rune":        "i32",
	"int64":       "i64",
	"float32":     "f32",
	"float64":     "f64",
	"bool":        "boolean",
	"uint":        "u32",
	"uint8":       "u8",
	"byte":        "u8",
	"uint16":      "u16",
	"uint32":      "u32",
	"uint64":      "u64",
	"complex64":   "c64",
	"complex128":  "c128",
	"string":      "str",
	"interface{}": "obj",
	"image.Point": "pt",
}

// T1FuncGen :
func T1FuncGen(Tx, pkgdir string) {

	pkgname, ok := mTypPkg[Tx]
	if !ok {
		panic(Tx + " is not supported for T<xxx>")
	}

	pkgdir = strings.TrimSuffix(pkgdir, "/") + "/"
	if !spt.DirExists(pkgdir) {
		if err := os.MkdirAll(pkgdir, os.ModePerm); err != nil {
			panic(err.Error())
		}
	}

	outgofile := pkgdir + pkgname + "/auto.go"

	// if !spt.FileExists(outgofile) || spt.FileIsEmpty(outgofile) {
	// 	spt.MustWriteFile(outgofile, []byte("package "+pkgname))
	// }
	if empty, err := spt.FileIsEmpty(outgofile); err != nil || empty {
		spt.MustWriteFile(outgofile, []byte("package "+pkgname))
	}

	src, err := spt.FileLineScan("./T1.template", func(line string) (bool, string) {
		line = strings.TrimRight(line, " \t")
		return true, strings.ReplaceAll(line, "xxx", Tx)
	}, "")

	if err != nil {
		panic(err.Error())
	}

	spt.MustAppendFile(outgofile, []byte(""), true)
	spt.MustAppendFile(outgofile, []byte(src), true)
}
