package user

import (
	"testing"

	"github.com/darchlabs/backoffice/internal/api/context"
	"github.com/darchlabs/backoffice/internal/storage"
	userdb "github.com/darchlabs/backoffice/internal/storage/user"
	"github.com/gofiber/fiber/v2"
	"github.com/jaekwon/testify/require"
)

func Test_PostSignupHandler_Invoke(t *testing.T) {
	testCases := []struct {
		name                string
		inputReq            *postSignupHandlerRequest
		expectedErrMsg      string
		expectedStatus      int
		userInsertQueryFunc userInsertQuery
	}{
		{
			name: "should return a 201 status code without error",
			inputReq: &postSignupHandlerRequest{
				Email:    "jdoe@gmail.com",
				Name:     "jon doe",
				Password: "securePassword",
			},
			expectedStatus: fiber.StatusCreated,
			userInsertQueryFunc: func(_ storage.QueryContext, _ *userdb.Record) error {
				return nil
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			h := PostSignupHandler{
				userInsertQuery: tc.userInsertQueryFunc,
			}

			ctx := &context.Ctx{}

			_, status, err := h.invoke(ctx, tc.inputReq)
			if tc.expectedErrMsg != "" {
				require.Equal(t, err.Error(), tc.expectedErrMsg)
				require.Equal(t, status, tc.expectedStatus)
				return
			}

			require.Equal(t, status, tc.expectedStatus)
		})
	}
}
