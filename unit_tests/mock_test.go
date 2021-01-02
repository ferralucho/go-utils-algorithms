// services/users_service_test.go
package services

/*
type dbClientMock struct {
	getUserFn func(id int64) (*domain.User, error)
}

func (c *dbClientMock) GetUser(id int64) (*domain.User, error) {
	return c.getUserFn(id)
}
func TestGetUser(t *testing.T) {
	dbMock := dbClientMock{}
	clients.Database = dbMock
	t.Run("ErrorFromDatabase", func(t *testing.T) {
		dbMock.getUserFn = func(id int64) (*domain.User, error) {
			return nil, errors.New("error connecting to our cool database")
		}
		user, err := GetUser(123)
		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualValues(t, "error connecting to our cool database", err.Error())
	})
	t.Run("UserNotFound", func(t *testing.T) {
		dbMock.getUserFn = func(id int64) (*domain.User, error) {
			return nil, nil
		}
		user, err := GetUser(123)
		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualValues(t, "user 123 not found", err.Error())
	})
	t.Run("NoError", func(t *testing.T) {
		dbMock.getUserFn = func(id int64) (*domain.User, error) {
			user := domain.User{
				Id:        id,
				FirstName: "Cool",
				LastName:  "Mock",
			}
			return &user, nil
		}
		user, err := GetUser(123)
		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.EqualValues(t, 123, user.Id)
		assert.EqualValues(t, "Cool", user.FirstName)
		assert.EqualValues(t, "Mock", user.LastName)
	})
}
*/
