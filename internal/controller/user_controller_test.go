package controller

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/farhaniupr/dating-api/package/library"
	"github.com/farhaniupr/dating-api/resource/model"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func generateRandomEmail() string {
	rand.Seed(time.Now().UnixNano())
	domain := "example.com"
	localPart := fmt.Sprintf("user%d", rand.Intn(10000))
	return fmt.Sprintf("%s@%s", localPart, domain)
}

func generateRandomPassword() string {
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	pass := make([]byte, 10)
	for i := range pass {
		pass[i] = charset[rand.Intn(len(charset))]
	}
	return string(pass)
}

func generateRandomnName() string {
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	pass := make([]byte, 10)
	for i := range pass {
		pass[i] = charset[rand.Intn(len(charset))]
	}
	return string(pass)
}

func generatePhoneNumber() string {
	rand.Seed(time.Now().UnixNano())

	areaCode := rand.Intn(900) + 100
	exchangeCode := rand.Intn(900) + 100
	subscriberNumber := rand.Intn(10000)

	phoneNumber := fmt.Sprintf("%03d-%03d-%04d", areaCode, exchangeCode, subscriberNumber)
	return phoneNumber
}

func generateRandomUrlPhoto() string {
	rand.Seed(time.Now().UnixNano())

	url := fmt.Sprintf("https://example/data.jpg")
	return url
}

func generateRandomDateBirth() string {
	rand.Seed(time.Now().UnixNano())

	start := time.Date(1950, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2010, 12, 31, 23, 59, 59, 0, time.UTC)

	randomTime := start.Add(time.Duration(rand.Int63n(end.Unix()-start.Unix())) * time.Second)

	return fmt.Sprintf("%s", randomTime.Format("2006-01-02"))
}

func generateRandomGender() string {
	rand.Seed(time.Now().UnixNano())

	genders := []string{"male", "female"}

	randomIndex := rand.Intn(len(genders))

	randomGender := genders[randomIndex]

	// Print the random gender
	return fmt.Sprintf("%s", randomGender)
}

type MockUserService struct {
	mock.Mock
}

type MockCommonHelper struct {
	mock.Mock
}

func (m *MockUserService) DetailUser(ctx context.Context, phone string) (model.User, error) {
	args := m.Called(ctx, phone)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockUserService) ListUser(ctx context.Context, page, limit int) ([]model.User, int64, error) {
	args := m.Called(ctx, page, limit)
	return args.Get(0).([]model.User), args.Get(1).(int64), args.Error(2)
}

func (m *MockUserService) UpdateUser(db *gorm.DB, user model.User, phone string) (model.User, error) {
	args := m.Called(db, user, phone)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockUserService) Login(ctx context.Context, user model.User) (model.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockUserService) StoreUser(db *gorm.DB, user model.User) (model.User, error) {
	args := m.Called(db, user)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockUserService) SwiftRight(ctx context.Context, db *gorm.DB, jwtData map[string]interface{}, phoneTarget string) (model.User, error) {
	args := m.Called(ctx, db, jwtData, phoneTarget)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockUserService) FindDate(ctx context.Context, jwtData map[string]interface{}) (model.User, error) {
	args := m.Called(ctx, jwtData)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockUserService) BuyPremium(db *gorm.DB, jwtData map[string]interface{}) (model.User, error) {
	args := m.Called(db, jwtData)
	return args.Get(0).(model.User), args.Error(1)
}

func TestLogin(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{"phone":"216-253-6879","password":"mypassword"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userService := new(MockUserService)
	userController := UserController{
		userService: userService,
	}

	t.Run("Successful Login", func(t *testing.T) {
		userService.On("Login", mock.Anything, mock.Anything).Return(model.User{Phone: "216-253-6879"}, nil)

		assert.NoError(t, userController.Login(c))
		assert.Equal(t, http.StatusOK, rec.Code)
		userService.AssertExpectations(t)
	})

	t.Run("Failed Login - Account Not Found", func(t *testing.T) {
		userService.On("Login", mock.Anything, mock.Anything).Return(nil, errors.New("account not found"))

		assert.NoError(t, userController.Login(c))
		assert.Equal(t, http.StatusOK, rec.Code)
		userService.AssertExpectations(t)
	})

	t.Run("Failed Login - Wrong Password", func(t *testing.T) {
		userService.On("Login", mock.Anything, mock.Anything).Return(nil, errors.New("password is wrong"))

		assert.NoError(t, userController.Login(c))
		assert.Equal(t, http.StatusOK, rec.Code)
		userService.AssertExpectations(t)
	})

	t.Run("Bad JSON Format", func(t *testing.T) {
		badReq := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(`{email:"test@example.com","password":"password123"}`))
		badReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		badRec := httptest.NewRecorder()
		badC := e.NewContext(badReq, badRec)

		assert.NoError(t, userController.Login(badC))
		assert.Equal(t, http.StatusBadRequest, badRec.Code)
	})
}

func TestRegister(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserService := new(MockUserService)
	mockContext := echo.New()

	userController := UserController{
		env:         library.Env{},
		userService: mockUserService,
	}

	tests := []struct {
		name             string
		userInput        string
		expectedStatus   int
		expectedResponse string
		setupMock        func()
	}{
		{
			name: "successful registration",
			userInput: fmt.Sprintf(`{"phone":"%s","name":"%s","url_photo":"%s","date_birth":"%s","gender":"%s","email":"%s","password":"%s"}`,
				generatePhoneNumber(), generateRandomnName(), generateRandomUrlPhoto(), generateRandomDateBirth(), generateRandomGender(), generateRandomEmail(), generateRandomPassword()),
			expectedStatus:   http.StatusCreated,
			expectedResponse: "Register Success",
			setupMock: func() {
				mockUserService.On("StoreUser", mock.Anything, mock.Anything).Return(model.User{Name: "", Phone: "123123-2131", Email: "test@test.com"}, nil)
			},
		},
		{
			name: "binding error",
			userInput: fmt.Sprintf(`{"phone":"%s","name":"%s","url_photo":"%s","date_birth":"%s","gender":"%s","email":"%s","password":"%s"}`,
				generatePhoneNumber(), generateRandomnName(), generateRandomUrlPhoto(), generateRandomDateBirth(), generateRandomGender(), generateRandomEmail(), generateRandomPassword()),
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: "Bad Request",
			setupMock:        func() {},
		},
		{
			name: "validation error",
			userInput: fmt.Sprintf(`{"phone":"%s","name":"%s","url_photo":"%s","date_birth":"%s","gender":"%s","email":"%s","password":"%s"}`,
				generatePhoneNumber(), generateRandomnName(), generateRandomUrlPhoto(), generateRandomDateBirth(), generateRandomGender(), generateRandomEmail(), generateRandomPassword()),
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: "BadRequest",
			setupMock: func() {
				mockContext.Validator = new(library.CustomValidator)
			},
		},
		{
			name: "user service error",
			userInput: fmt.Sprintf(`{"phone":"%s","name":"%s","url_photo":"%s","date_birth":"%s","gender":"%s","email":"%s","password":"%s"}`,
				generatePhoneNumber(), generateRandomnName(), generateRandomUrlPhoto(), generateRandomDateBirth(), generateRandomGender(), generateRandomEmail(), generateRandomPassword()),
			expectedStatus:   http.StatusInternalServerError,
			expectedResponse: "Internal Server Error",
			setupMock: func() {
				mockUserService.On("StoreUser", mock.Anything, mock.Anything).Return(model.User{}, errors.New("internal error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader([]byte(tt.userInput)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := mockContext.NewContext(req, rec)
			c.SetPath("/register")

			tt.setupMock()

			if assert.NoError(t, userController.Register(c)) {
				assert.Equal(t, tt.expectedStatus, rec.Code)
				assert.Contains(t, rec.Body.String(), tt.expectedResponse)
			}

			mockUserService.AssertExpectations(t)
		})
	}
}