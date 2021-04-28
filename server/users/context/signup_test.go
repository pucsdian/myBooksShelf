package context

import (
	"testing"

	"github.com/shubhamjagdhane/myBooksShelf/server/users/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {

	hashPassword, err := util.HashPassword(util.RandomString(8))
	require.NoError(t, err)

	arg := Users{
		Username: util.RandomUsername(10),
		Password: hashPassword,
		Email:    util.RandomEmail(6),
	}

	newUser, err := testConn.CreateUser(arg)
	require.NoError(t, err)
	require.NotEmpty(t, newUser)
	require.Equal(t, arg.Username, newUser.Username)
	require.Equal(t, arg.Email, newUser.Email)
	require.Equal(t, arg.Password, newUser.Password)
	require.NotZero(t, newUser.ID)
	require.NotZero(t, newUser.CreatedAt)

}
