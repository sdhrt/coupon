package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *sql.DB
}

type User struct {
	User_id       string
	Name          string
	Email         string
	Password_hash string
}

// Create_user method take email, name and password and registers it to the database
func (m *UserModel) Create_user(name, email, password string) (uuid.UUID, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.Nil, err
	}

	query := `
	INSERT INTO users(user_id, name, email, password_hash)
	VALUES ($1, $2, $3, $4)`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	id := uuid.New()
	args := []any{id.String(), name, email, hash}

	_, err = m.DB.ExecContext(ctx, query, args...)
	if err != nil {
		if strings.HasPrefix(err.Error(), "pq: duplicate key") {
			return uuid.Nil, errors.New("The email address is already associated with an account")
		}
		fmt.Println("users.go NewUser error: ", err.Error())
		return uuid.Nil, errors.New("Couldn't create new user")
	}
	return id, nil
}

// Validate user takes email and password, it hashes the password and then
// return the user object if the user credentials are valid
func (m *UserModel) Validate_user(email, password string) (User, error) {
	var user User
	query := `
	SELECT user_id, email, name, password_hash FROM users WHERE email=$1
	`
	args := []any{email}
	err := m.DB.QueryRow(query, args...).Scan(&user.User_id, &user.Email, &user.Name, &user.Password_hash)
	if err != nil {
		return User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password_hash), []byte(password))
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// Get access token returns a jwt token whose external claims are
//
//	user_id
//	email
//	name
//	These are fetched from the User object that has been passed
func (m *UserModel) Get_access_token(user User, jwt_secret string) (string, error) {
	claims := &jwt.MapClaims{
		"iss": "issuer",
		"exp": time.Now().Add(time.Hour).Unix(),
		"data": map[string]string{
			"user_id": user.User_id,
			"email":   user.Email,
			"name":    user.Name,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwt_secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
