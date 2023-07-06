package database

import (
	"fmt"
	"sort"

	"golang.org/x/exp/slices"
)

type Framework struct {
	ID          uint16 `form:"id,omitempty"`
	Name        string `form:"name,omitempty"`
	Description string `form:"description,omitempty"`
	IsPoop      bool   `form:"is_poop,omitempty"`
}

type FrameworkList []Framework

type DB struct {
	frameworks []Framework
}

var (
	frameworkID uint16 = 0
)

func NewDB() *DB {
	return &DB{
		frameworks: []Framework{
			newFramework("AngolaJS", "legacy code", true),
			newFramework("Cockout.js", "MVVM", false),
			newFramework("Eww.js", "ewwwwwwwwwwww", false),
			newFramework("R***t", "V****l D*M", true),
			newFramework("SolidPoopJS", "signals", false),
			newFramework("Swolte", "Rich Harris... what a mensch", false),
		},
	}
}

func (db *DB) CreateFramework(body Framework) *Framework {
	fw := newFramework(body.Name, body.Description, body.IsPoop)
	db.frameworks = append(db.frameworks, fw)
	return &fw
}

func (db *DB) DeleteFramework(id uint16) error {
	idx := slices.IndexFunc(db.frameworks, func(fw Framework) bool { return fw.ID == id })
	if idx < 0 {
		return fmt.Errorf("cannot find framework with id \"%d\"", id)
	}
	db.frameworks[idx] = db.frameworks[len(db.frameworks)-1]
	db.frameworks = db.frameworks[:len(db.frameworks)-1]
	return nil
}

func (db *DB) Framework(id uint16) (*Framework, error) {
	for _, fw := range db.frameworks {
		if fw.ID == id {
			return &fw, nil
		}
	}
	return nil, fmt.Errorf("cannot find framework with id \"%d\"", id)
}

func (db *DB) ListFrameworks() []Framework {
	// FIXME inefficient
	cpy := make([]Framework, len(db.frameworks))
	copy(cpy, db.frameworks)
	sort.Sort(FrameworkList(cpy))
	return cpy
}

func (db *DB) UpdateFramework(id uint16, body Framework) (*Framework, error) {
	for i := 0; i <= len(db.frameworks); i++ {
		fw := &db.frameworks[i]
		if fw.ID == id {
			fw.Name = body.Name
			fw.Description = body.Description
			fw.IsPoop = body.IsPoop
			return fw, nil
		}
	}
	return nil, fmt.Errorf("cannot find framework with id \"%d\"", id)
}

func newFramework(name, description string, isPoop bool) Framework {
	fw := Framework{
		ID:          frameworkID,
		Name:        name,
		Description: description,
		IsPoop:      isPoop,
	}
	frameworkID += 1
	return fw
}

func (l FrameworkList) Len() int {
	return len(l)
}

func (l FrameworkList) Less(i, j int) bool {
	return l[i].ID < l[j].ID
}

func (l FrameworkList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
