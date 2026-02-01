package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPathValidationIntegration(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}

	t.Run("malformed paths don't affect store", func(t *testing.T) {
		// Try to record wins with malformed paths
		malformedRequests := []string{
			"/players/",
			"/players/Alice/extra",
			"/other/path",
		}

		for _, path := range malformedRequests {
			request, _ := http.NewRequest(http.MethodPost, path, nil)
			response := httptest.NewRecorder()

			server.ServeHTTP(response, request)
			assertStatus(t, response.Code, http.StatusNotFound)
		}

		// Verify no wins were recorded
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest("Alice"))
		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")
}

func TestMultiplePlayersIntegration(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}

	players := []string{"Alice", "Bob", "Charlie"}

	// Record different number of wins for each player
	for i, player := range players {
		for j := 0; j <= i; j++ {
			server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
		}
	}

	// Verify scores
	testCases := []struct {
		player        string
		expectedScore string
	}{
		{"Alice", "1"},
		{"Bob", "2"},
		{"Charlie", "3"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("returns correct score for %s", tc.player), func(t *testing.T) {
			response := httptest.NewRecorder()
			server.ServeHTTP(response, newGetScoreRequest(tc.player))

			assertStatus(t, response.Code, http.StatusOK)
			assertResponseBody(t, response.Body.String(), tc.expectedScore)
		})
	}
}

func TestPlayerNotFound(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest("NonExistentPlayer"))

	assertStatus(t, response.Code, http.StatusNotFound)
}
