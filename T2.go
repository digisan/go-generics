package generics

import (
	"os"
	"strings"

	spt "github.com/digisan/go-generics/support"
)

var (
	mPkgRename = map[string]string{
		"strboolean": "sb",
		"stri32":     "si32",
		"i32str":     "i32s",
		"strf64":     "sf64",
	}
)

// T2FuncGen :
func T2FuncGen(Tx, Ty, pkgdir string) {

	pkgname1, ok := mTypPkg[Tx]
	if !ok {
		panic(Tx + " is not supported for T<xxx>")
	}
	pkgname2, ok := mTypPkg[Ty]
	if !ok {
		panic(Ty + " is not supported for T<yyy>")
	}
	if Tx == Ty {
		pkgname2 = ""
	} else {
		pkgname2 = strings.TrimPrefix(pkgname2, "t")
	}

	pkgdir = strings.TrimSuffix(pkgdir, "/") + "/"
	if !spt.DirExists(pkgdir) {
		if err := os.MkdirAll(pkgdir, os.ModePerm); err != nil {
			panic(err.Error())
		}
	}

	pkgname := pkgname1 + pkgname2
	if rn, ok := mPkgRename[pkgname]; ok {
		pkgname = rn
	}

	outgofile := pkgdir + pkgname + "/auto.go"

	// if !spt.FileExists(outgofile) || spt.FileIsEmpty(outgofile) {
	// 	io.MustWriteFile(outgofile, []byte("package "+pkgname))
	// }
	if empty, err := spt.FileIsEmpty(outgofile); err != nil || empty {
		spt.MustWriteFile(outgofile, []byte("package "+pkgname))
	}

	flagXeqY, flagXneY := true, true

	src, err := spt.FileLineScan("./T2.template", func(line string) (bool, string) {
		line = strings.TrimRight(line, " \t")

		switch {
		case strings.HasSuffix(line, `[S@x==y]`):
			flagXeqY, flagXneY = true, false
		case strings.HasSuffix(line, `[E@x==y]`):
			flagXeqY, flagXneY = true, true
		case strings.HasSuffix(line, `[S@x!=y]`):
			flagXeqY, flagXneY = false, true
		case strings.HasSuffix(line, `[E@x!=y]`):
			flagXeqY, flagXneY = true, true
		}

		if strings.HasPrefix(strings.TrimLeft(line, " \t"), "// [") {
			return false, ""
		}

		line = strings.ReplaceAll(line, "yyy", Ty)
		line = strings.ReplaceAll(line, "xxx", Tx)

		if (Tx == Ty && flagXeqY) || (Tx != Ty && flagXneY) {
			return true, line
		}
		return false, ""

	}, "")

	if err != nil {
		panic(err.Error())
	}

	spt.MustAppendFile(outgofile, []byte(""), true)
	spt.MustAppendFile(outgofile, []byte(src), true)
}
