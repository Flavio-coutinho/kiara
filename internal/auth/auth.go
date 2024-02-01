// auth.go

package auth

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	PasswordHash string
	Admin     string 
}

type AuthManager struct {
	users map[string]User
	tokens map[string]Token
}

type Token struct {
	TokenString string
	Username    string
	Expiration  time.Time
}

func NewAuthManager() *AuthManager {
	return &AuthManager{
		users: make(map[string]User),
		tokens: make(map[string]Token),
	}
}

func (am *AuthManager) AddUser(username, password, admin string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	am.users[username] = User{
		Username: username,
		PasswordHash: string(hashedPassword),
		Admin:     admin,
	}

	return nil
}

func (am *AuthManager) Authenticate(username, password string) (string, error) {
	user, exists := am.users[username]
	if !exists {
		return "", fmt.Errorf("user not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", fmt.Errorf("authentication failed")
	}

	tokenString := generateTokenString()
	expiration := time.Now().Add(time.Hour * 1) 
	token := Token{
		TokenString: tokenString,
		Username:    username,
		Expiration:  expiration,
	}
	am.tokens[tokenString] = token

	return tokenString, nil
}

func (am *AuthManager) Authorize(tokenString, action string) bool {
	token, exists := am.tokens[tokenString]
	if !exists {
		return false 
	}

	
	return token.Expiration.After(time.Now()) && (token.Username == "admin" || (token.Username == "user" && action == "read"))
}

func generateTokenString() string {
	return "unique-token"
}
