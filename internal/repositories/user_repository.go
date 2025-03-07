package repositories

import "github.com/assaabriiii/gin-cookie-auth/internal/models"

// In-memory storage for users
var users = make(map[string]models.User)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindByUsername(username string) (*models.User, bool) {
	user, exists := users[username]
	return &user, exists
}

func (r *UserRepository) Save(user *models.User) {
	users[user.Username] = *user
}
