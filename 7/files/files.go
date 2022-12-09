package files

import "fmt"

// import "fmt"

type Dir struct {
	Name   string // Name of Directory
	Parent *Dir   // Pointer to father
	Size   int
}

func InsertFile(d *Dir, size int) {
	d.Size += size
	if d.Parent != nil {
		InsertFile(d.Parent, size)
	}
}

func PrintSizes(d *Dir) {
	if d.Parent != nil {
		PrintSizes(d.Parent)
	}
	fmt.Print("<", d.Size)
}

func InsertDir(d *Dir, name string) *Dir {
	dir := Dir{name, d, 0}
	return &dir
}

func FindDir(dir []*Dir, name string) *Dir {
	for _, d := range dir {
		if d.Name == name {
			return d
		}
	}
	return nil
}
