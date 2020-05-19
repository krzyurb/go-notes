package routes

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"notesapp/loaders"
	"testing"
)

func TestGetNotesRoute(t *testing.T) {
	router := loaders.BuildRouter()

	t.Run("get all notes", func (t *testing.T){
		request := httptest.NewRequest("GET","/notes", nil)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, response.Code, 200)
	})
}

