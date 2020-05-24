package notes

import (
	"github.com/go-playground/assert/v2"
	"notesapp/notes"
	"notesapp/tests/helpers"
	"testing"
)

func TestNotesStore(t *testing.T){
	dbHelper := helpers.BuildDbHelper()
	store := notes.NotesStore{Db: dbHelper.Db}

	t.Run("responds with empty list if no notes", func (t *testing.T) {
		dbHelper.Clear()

		result := store.GetAll()
		assert.Equal(t, len(result), 0)
	})

	t.Run("responds with all notes", func (t *testing.T) {
		dbHelper.Clear()

		dbHelper.CreateNote()
		dbHelper.CreateNote()
		
		result := store.GetAll()

		assert.Equal(t, len(result), 2)
	})
}