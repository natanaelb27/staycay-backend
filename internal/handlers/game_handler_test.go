package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/dtos"
	e "git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/errors"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/mocks"
	"git.garena.com/sea-labs-id/batch-02/natanael-benediktus/final-project-backend/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandler_PlayGameHandler(t *testing.T) {
	type args struct {
		method   string
		endPoint string
	}
	tests := []struct {
		name        string
		gameService *mocks.GameService
		request     *dtos.PlayGameRequestDTO
		args        args
		mock        func(*mocks.GameService)
		wants       int
	}{
		{
			name:        "Failed - Bad Request",
			gameService: mocks.NewGameService(t),
			request:     &dtos.PlayGameRequestDTO{},
			args: args{
				method:   http.MethodPost,
				endPoint: "/game",
			},
			mock: func(m *mocks.GameService) {
			},
			wants: http.StatusBadRequest,
		},
		{
			name:        "Failed - Insufficient Game Chance",
			gameService: mocks.NewGameService(t),
			request: &dtos.PlayGameRequestDTO{
				UserGuess: "heads",
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/game",
			},
			mock: func(m *mocks.GameService) {
				m.On("PlayGame", "heads", MockUser.UserID).Return("", 0, e.ErrInsufficientGameChanceError)
			},
			wants: http.StatusBadRequest,
		},
		{
			name:        "Failed - Internal Server Error",
			gameService: mocks.NewGameService(t),
			request: &dtos.PlayGameRequestDTO{
				UserGuess: "heads",
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/game",
			},
			mock: func(m *mocks.GameService) {
				m.On("PlayGame", "heads", MockUser.UserID).Return("", 0, errors.New("unknown error"))
			},
			wants: http.StatusInternalServerError,
		},
		{
			name:        "Success",
			gameService: mocks.NewGameService(t),
			request: &dtos.PlayGameRequestDTO{
				UserGuess: "heads",
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/game",
			},
			mock: func(m *mocks.GameService) {
				m.On("PlayGame", "heads", MockUser.UserID).Return("heads", 50000, nil)
			},
			wants: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(&HandlerConfig{
				GameService: tt.gameService,
			})

			tt.mock(tt.gameService)

			r := gin.Default()

			jsonValue, _ := json.Marshal(tt.request)

			r.POST(tt.args.endPoint, MiddlewareMockUser, h.PlayGameHandler)

			req, _ := http.NewRequest(
				tt.args.method,
				tt.args.endPoint,
				bytes.NewBuffer(jsonValue),
			)

			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wants, w.Code)
		})
	}
}

func TestHandler_GetGameHandler(t *testing.T) {
	type args struct {
		method   string
		endPoint string
	}
	tests := []struct {
		name        string
		gameService *mocks.GameService
		args        args
		mock        func(*mocks.GameService)
		wants       int
	}{
		{
			name:        "Failed - Internal Server Error",
			gameService: mocks.NewGameService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/game",
			},
			mock: func(m *mocks.GameService) {
				m.On("GetGame", MockUser.UserID).Return(nil, errors.New("unknown error"))
			},
			wants: http.StatusInternalServerError,
		},
		{
			name:        "Success",
			gameService: mocks.NewGameService(t),
			args: args{
				method:   http.MethodGet,
				endPoint: "/game",
			},
			mock: func(m *mocks.GameService) {
				m.On("GetGame", MockUser.UserID).Return(&models.GameChance{}, nil)
			},
			wants: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(&HandlerConfig{
				GameService: tt.gameService,
			})

			tt.mock(tt.gameService)

			r := gin.Default()

			r.GET(tt.args.endPoint, MiddlewareMockUser, h.GetGameHandler)

			req, _ := http.NewRequest(
				tt.args.method,
				tt.args.endPoint,
				nil,
			)

			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.wants, w.Code)
		})
	}
}
