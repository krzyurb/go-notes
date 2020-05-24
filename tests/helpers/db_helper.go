package helpers

import (
	"github.com/bxcodec/faker/v3"
	"github.com/jinzhu/gorm"
	"notesapp/db"
	"notesapp/loaders"
	"notesapp/notes"
)

type DbHelper struct {
	Db *gorm.DB
}

func (dh *DbHelper) Clear () {
	dh.Db.Delete(notes.Note{})
}

func (dh *DbHelper) CreateNote () *notes.Note {
	note := notes.Note{
		Title: faker.Sentence(),
		IsPublic: false,
		Content: faker.Paragraph(),
	}
	dh.Db.Create(&note)
	return &note
}

func BuildDbHelper() DbHelper {
	return DbHelper{Db: db.BuildConnection(loaders.BuildConfig())}
}