package filestorage

import (
	"slices"
)

// File storage could be
// a directory if there's no data
// and it could be a file if there's a data.
type FileStorage struct {
	Name     string
	Children []*FileStorage
	Content  []byte
	IsFile   bool
}

type Data struct {
	Parent *string
	File   FileStorage
}

func NewDir(name string, parent *string) *Data {
	return &Data{
		Parent: parent,
		File: FileStorage{
			Name:     name,
			Children: nil,
			Content:  nil,
		},
	}
}

// Add adds new file to the root or
// create new directorty
func (fs *FileStorage) Add(new Data) {

	if fs == nil {
		return
	}

	// if there's no parent add to the root.
	if new.Parent == nil && len(fs.Children) == 0 {
		fs.Children = slices.Insert(fs.Children, 0, &new.File)
		return
	} else if new.Parent == nil {
		fs.Children = append(fs.Children, &new.File)
		return
	}

	if fs.Name == *new.Parent {
		if len(fs.Children) == 0 {
			fs.Children = slices.Insert(fs.Children, 0, &new.File)
		} else {
			fs.Children = append(fs.Children, &new.File)
		}
		return
	}

	for _, fc := range fs.Children {

		if fc.Name == *new.Parent && len(fc.Children) == 0 {
			fc.Children = slices.Insert(fc.Children, 0, &new.File)
			return
		} else if fc.Name == *new.Parent {
			fc.Children = append(fc.Children, &new.File)
			return
		}

		for _, nc := range fc.Children {
			nc.Add(new)
		}

	}

}
