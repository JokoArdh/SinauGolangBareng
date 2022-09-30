package golangembed

//jngan lupa embed nya harus di import dulu ya

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

func TestEmbed(t *testing.T) {
	fmt.Println(version)
}

//cara mengeload adata binary di golang
//go:embed device.png
var gambar []byte

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("gambar device roouter.png", gambar, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//cara load embed multiple file
//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestMultiple(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}

//memnggunakan path matcher semua yang ada di files di load
//go:embed files/*txt
var path embed.FS

func TestPath(t *testing.T) {
	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
		}
	}
}
