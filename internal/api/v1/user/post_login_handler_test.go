package user

import (
	"testing"

	"github.com/darchlabs/backoffice/internal/api/context"
	"github.com/darchlabs/backoffice/internal/storage"
	"github.com/darchlabs/backoffice/internal/storage/auth"
	"github.com/darchlabs/backoffice/internal/storage/user"
	userdb "github.com/darchlabs/backoffice/internal/storage/user"
	"github.com/jaekwon/testify/require"
)

func Test_PostLoginHandler_Invoke(t *testing.T) {
	testCases := []struct {
		name                   string
		inputReq               *postLoginHandlerRequest
		secretKey              string
		expectedErrMsg         string
		expectedStatus         int
		userSelectByEmailQuery userSelectByEmailQuery
		authUpsertQuery        authUpsertQuery
	}{
		{
			name:      "should return a 201 status code without error",
			secretKey: "this-is-secre-key",
			inputReq: &postLoginHandlerRequest{
				Email:    "jdoe@gmail.com",
				Password: "password124",
			},
			expectedStatus: 201,
			userSelectByEmailQuery: func(storage.Transaction, string) (*user.Record, error) {
				return &userdb.Record{
					ID:             "test-id",
					Email:          "jdoe@gmail.com",
					Name:           "jon doe",
					HashedPassword: "$2a$10$g67l3Y4ldTU1D/9qJOq9N.T2yLfmij/YMt0SBkC1iTmQ4UJZwvIB2",
				}, nil
			},
			authUpsertQuery: func(storage.QueryContext, *auth.Record) error {
				return nil
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			h := &PostLoginHandler{
				secretKey:              tc.secretKey,
				userSelectByEmailQuery: tc.userSelectByEmailQuery,
				authUpsertQuery:        tc.authUpsertQuery,
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
