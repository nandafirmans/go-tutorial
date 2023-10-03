package embed

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:embed version.txt
var vesion string

func TestEmbedString(t *testing.T) {
	fmt.Println("Version", vesion)
}

//go:embed cat.png
var catImage []byte

func TestEmbedByte(t *testing.T) {
	err := os.WriteFile("cat_copy.png", catImage, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestEmbedFS(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dirEntries, _ := path.ReadDir("files")
	for i, entry := range dirEntries {
		if entry.IsDir() == false {
			fmt.Println(i, entry.Name())

			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}

// NOTE: hasil embed di compile ke dalam binary, jadinya tidak bisa diubah lagi dan kita tidak perlu
// mengirimkan file-file tersebut lagi setelah binary di compile
