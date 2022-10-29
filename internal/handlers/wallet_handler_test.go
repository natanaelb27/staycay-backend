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

func TestHandler_TopUpWalletHandler(t *testing.T) {
	type args struct {
		method   string
		endPoint string
	}
	tests := []struct {
		name          string
		walletService *mocks.WalletService
		request       *dtos.TopUpWalletRequestDTO
		args          args
		mock          func(*mocks.WalletService)
		wants         int
	}{
		{
			name:          "Failed - Bad Request",
			walletService: mocks.NewWalletService(t),
			request:       &dtos.TopUpWalletRequestDTO{},
			args: args{
				method:   http.MethodPost,
				endPoint: "/topup",
			},
			mock: func(m *mocks.WalletService) {
			},
			wants: http.StatusBadRequest,
		},
		{
			name:          "Failed - Top Up Amount too Small",
			walletService: mocks.NewWalletService(t),
			request: &dtos.TopUpWalletRequestDTO{
				Amount: 5000,
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/topup",
			},
			mock: func(m *mocks.WalletService) {
				m.On("TopUpWallet", MockUser.UserID, 5000).Return(nil, e.ErrTopUpTooSmallError)
			},
			wants: http.StatusBadRequest,
		},
		{
			name:          "Failed - Internal Server Error",
			walletService: mocks.NewWalletService(t),
			request: &dtos.TopUpWalletRequestDTO{
				Amount: 50000,
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/topup",
			},
			mock: func(m *mocks.WalletService) {
				m.On("TopUpWallet", MockUser.UserID, 50000).Return(nil, errors.New("unknown error"))
			},
			wants: http.StatusInternalServerError,
		},
		{
			name:          "Success",
			walletService: mocks.NewWalletService(t),
			request: &dtos.TopUpWalletRequestDTO{
				Amount: 50000,
			},
			args: args{
				method:   http.MethodPost,
				endPoint: "/topup",
			},
			mock: func(m *mocks.WalletService) {
				m.On("TopUpWallet", MockUser.UserID, 50000).Return(&models.Wallet{}, nil)
			},
			wants: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(&HandlerConfig{
				WalletService: tt.walletService,
			})

			tt.mock(tt.walletService)

			r := gin.Default()

			r.POST(tt.args.endPoint, MiddlewareMockUser, h.TopUpWalletHandler)

			jsonValue, _ := json.Marshal(tt.request)

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
