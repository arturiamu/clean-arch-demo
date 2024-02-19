package localstorage

import (
	"clean-arch-demo/auth"
	"clean-arch-demo/models"
	"context"
	"sync"
)

var _ auth.UserRepository = (*UserLocalStorage)(nil)

type UserLocalStorage struct {
	users map[uint]*models.User
	mutex *sync.Mutex
}

func NewUserLocalStorage() *UserLocalStorage {
	return &UserLocalStorage{
		users: make(map[uint]*models.User),
		mutex: new(sync.Mutex),
	}
}

func (u *UserLocalStorage) CreateUser(ctx context.Context, user *models.User) error {
	u.mutex.Lock()
	u.users[user.ID] = user
	u.mutex.Unlock()

	return nil
}

func (u *UserLocalStorage) GetUser(ctx context.Context, username, password string) (*models.User, error) {
	u.mutex.Lock()
	defer u.mutex.Unlock()

	for _, user := range u.users {
		if user.Username == username && user.Password == password {
			return user, nil
		}
	}

	return nil, auth.ErrUserNotFound
}
