package notes

import (
	"github.com/go-playground/assert/v2"
	db2 "notesapp/db"
	"notesapp/loaders"
	"notesapp/notes"
	"testing"
)

func TestNotesStore(t *testing.T){
	db := db2.BuildConnection(loaders.BuildConfig())
	store := notes.NotesStore{Db: db}

	t.Run("get all notes", func (t *testing.T) {
		note := store.GetByID("AAA")
		assert.Equal(t, note, nil)
	})
}