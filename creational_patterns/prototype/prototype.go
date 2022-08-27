package main

import "fmt"

/*
Prototype is a creational design pattern that allows
cloning objects, even complex ones, without coupling
to their specific classes.

All prototype classes should have a common interface
that makes it possible to copy objects even if their
concrete classes are unknown. Prototype objects can
produce full copies since objects of the same class
can access each other’s private fields.
*/

type Node interface {
	print(prefix string)
	clone() Node
}

type File struct {
	name string
}

func (f *File) print(prefix string) {
	fmt.Println(prefix + f.name)
}

func (f *File) clone() Node {
	c := &File{
		name: f.name + "_clone",
	}

	return c
}

type Folder struct {
	Files []Node
	name  string
}

func (f *Folder) print(prefix string) {
	fmt.Println(prefix + f.name)
	for _, file := range f.Files {
		file.print(prefix + prefix)
	}
}

func (f *Folder) clone() Node {
	folder := new(Folder)
	folder.name = f.name + "_clone"

	var files []Node

	for _, file := range f.Files {
		n := file.clone()
		files = append(files, n)
	}

	folder.Files = files

	return folder
}

func main() {
	file1 := &File{name: "file1"}
	file2 := &File{name: "file2"}

	folder1 := &Folder{name: "folder1", Files: []Node{file1, file2}}

	folder2 := &Folder{name: "folder2", Files: []Node{file1, folder1}}

	folder2.print("   ")

	cloneFolder2 := folder2.clone()
	cloneFolder2.print("   ")
}
