package main

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/fumiama/imgsz"
)

//go:generate go run main.go

const wifvesdir = "./wives/"
const jsonfile = "wife.json"

func main() {
	ent, err := os.ReadDir(wifvesdir)
	if err != nil {
		panic(err)
	}
	cards := make([]string, 0, len(ent))
	for _, en := range ent {
		if en.IsDir() {
			continue
		}
		name := en.Name()
		fn := wifvesdir + name
		f, err := os.Open(fn)
		if err != nil {
			continue
		}
		_, format, err := imgsz.DecodeSize(f)
		_ = f.Close()
		if err != nil {
			continue
		}
		i := strings.LastIndex(name, ".")
		if i <= 0 {
			continue
		}
		name = name[:i] + "." + format
		nfn := wifvesdir + name
		if fn != nfn {
			err = os.Rename(fn, nfn)
			if err != nil {
				continue
			}
		}
		cards = append(cards, name)
	}
	f, err := os.Create(jsonfile)
	if err != nil {
		panic(err)
	}
	err = json.NewEncoder(f).Encode(cards)
	if err != nil {
		panic(err)
	}
}
