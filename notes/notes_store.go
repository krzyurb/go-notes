package notes

import "github.com/jinzhu/gorm"

type Note struct {
	gorm.Model
	Title string
	Content string
	Public bool
}

type NotesStore struct {
	Db *gorm.DB
}

type INotesStore interface {
	GetAll() []Note
	GetByID(id string) Note
}

func (ns *NotesStore) GetAll() (notes []Note) {
	ns.Db.Find(&notes)
	return notes
}

func (ns *NotesStore) GetByID(id string) (note Note) {
	ns.Db.Find(&note, id)
	return note
}