package usecase

import (
	"os"
	"time"

	"github.com/MingPV/NotificationService/internal/entities"
	"github.com/MingPV/NotificationService/internal/user/repository"
	"github.com/MingPV/NotificationService/pkg/apperror"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// NotificationService struct
type NotificationService struct {
	repo repository.UserRepository
}

// Init NotificationService
func NewNotificationService(repo repository.UserRepository) UserUseCase {
	return &NotificationService{repo: repo}
}

// NotificationService Methods - 1 Register user (hash password)
func (s *NotificationService) Register(user *entities.User) error {
	existingUser, _ := s.repo.FindByEmail(user.Email)
	if existingUser != nil {
		return apperror.ErrAlreadyExists
	}

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPwd)

	return s.repo.Save(user)
}

// NotificationService Methods - 2 Login user (check email + password)
func (s *NotificationService) Login(email string, password string) (string, *entities.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil || user == nil {
		return "", nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", nil, err
	}

	// Generate JWT token
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // 3 days
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", nil, err
	}

	return tokenString, user, nil
}

// NotificationService Methods - 3 Get user by id
func (s *NotificationService) FindUserByID(id string) (*entities.User, error) {
	return s.repo.FindByID(id)
}

// NotificationService Methods - 4 Get all users
func (s *NotificationService) FindAllUsers() ([]*entities.User, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// NotificationService Methods - 5 Get user by email
func (s *NotificationService) GetUserByEmail(email string) (*entities.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// NotificationService Methods - 6 Patch
func (s *NotificationService) PatchUser(id string, user *entities.User) (*entities.User, error) {
	if err := s.repo.Patch(id, user); err != nil {
		return nil, err
	}
	updatedUser, _ := s.repo.FindByID(id)

	return updatedUser, nil
}

// NotificationService Methods - 7 Delete
func (s *NotificationService) DeleteUser(id string) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
