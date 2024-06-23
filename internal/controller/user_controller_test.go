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
	"github.com/farhaniupr/dating-api/resource/constants"
	"github.com/farhaniupr/dating-api/resource/model"
	"github.com/go-playground/validator"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func dateValidation(fl validator.FieldLevel) bool {
	_, err := time.Parse(constants.LayoutDate, fl.Field().String())
	return err == nil
}

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

	return "https://example/data.jpg"
}

func generateRandomDateBirth() string {
	rand.Seed(time.Now().UnixNano())

	start := time.Date(1950, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2010, 12, 31, 23, 59, 59, 0, time.UTC)

	randomTime := start.Add(time.Duration(rand.Int63n(end.Unix()-start.Unix())) * time.Second)

	return randomTime.Format("2006-01-02")
}

func generateRandomGender() string {
	rand.Seed(time.Now().UnixNano())

	genders := []string{"male", "female"}

	randomIndex := rand.Intn(len(genders))

	randomGender := genders[randomIndex]

	// Print the random gender
	return randomGender
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

func (m *MockUserService) UpdateUser(db interface{}, user model.User, phone string) (model.User, error) {
	args := m.Called(db, user, phone)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockUserService) Login(ctx context.Context, user model.User) (model.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockUserService) StoreUser(db interface{}, user model.User) (model.User, error) {
	args := m.Called(db, user)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockUserService) SwiftRight(ctx context.Context, db interface{}, jwtData map[string]interface{}, phoneTarget string) (model.User, error) {
	args := m.Called(ctx, db, jwtData, phoneTarget)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockUserService) FindDate(ctx context.Context, jwtData map[string]interface{}) (model.User, error) {
	args := m.Called(ctx, jwtData)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockUserService) BuyPremium(db interface{}, jwtData map[string]interface{}) (model.User, error) {
	args := m.Called(db, jwtData)
	return args.Get(0).(model.User), args.Error(1)
}

func TestLogin(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/user/login", strings.NewReader(`{"phone":"216-253-6879","password":"mypassword"}`))
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
		badReq := httptest.NewRequest(http.MethodPost, "/user/login", strings.NewReader(`{email:"test@example.com","password":"password123"}`))
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

	validator := validator.New()
	validator.RegisterValidation("date", dateValidation)

	validator.RegisterValidation("date", dateValidation)
	mockContext.Validator = &CustomValidator{Validator: validator}

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
			userInput: fmt.Sprintf(`{"phone":1123213,"name":2132131,"url_photo":"%s","date_birth":"%s","gender":"%s","email":"%s","password":"%s"}`,
				generateRandomUrlPhoto(), generateRandomDateBirth(), generateRandomGender(), generateRandomEmail(), generateRandomPassword()),
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: "Bad Request",
			setupMock:        func() {},
		},
		{
			name:             "validation error",
			userInput:        fmt.Sprintf(`{"phone":"%s","name":"%s","url_photo":"%s","date_birth":"%s","gender":"%s"}`, generatePhoneNumber(), generateRandomnName(), generateRandomUrlPhoto(), generateRandomDateBirth(), generateRandomGender()),
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: "required",
			setupMock: func() {
				mockContext.Validator = &CustomValidator{Validator: validator}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewReader([]byte(tt.userInput)))
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

func TestUserController_SwiftRight(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/user/swift-right?phone_target=123456789", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("phone_target")
	c.SetParamValues("123456789")
	db := &gorm.DB{}
	c.Set(constants.DBTransaction, db)
	dataJwt := map[string]interface{}{"user_id": "123"}
	c.Set("data_jwt", dataJwt)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserService := new(MockUserService)

	userController := UserController{
		userService: mockUserService,
	}

	// Setting context data

	tests := []struct {
		name       string
		ctx        context.Context
		setupMocks func()
		wantErr    bool
		wantStatus int
	}{
		{
			name:       "Success Liked and Find New Date",
			ctx:        req.Context(),
			wantErr:    false,
			wantStatus: 200,
			setupMocks: func() {
				mockUserService.On("SwiftRight", c.Request().Context(),
					db,
					dataJwt,
					"123456789").Return(model.User{}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()
			err := userController.SwiftRight(c)

			if (err != nil) != tt.wantErr {
				t.Errorf("UserController.SwiftRight() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				httpError, ok := err.(*echo.HTTPError)
				if !ok {
					t.Errorf("Expected echo.HTTPError, got %T", err)
				}
				assert.Equal(t, tt.wantStatus, http.StatusUnprocessableEntity)
				assert.Equal(t, "Internal Server Error", httpError.Message)
			} else {
				assert.Equal(t, tt.wantStatus, rec.Code)
				assert.Contains(t, rec.Body.String(), "Success Liked and Find New Date")
			}
		})
	}
}

func TestUserController_SwiftLeft(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/user/find-date", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	db := &gorm.DB{}
	c.Set(constants.DBTransaction, db)
	dataJwt := map[string]interface{}{"phone": "216-253-6879"}
	c.Set("data_jwt", dataJwt)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserService := new(MockUserService)

	userController := UserController{
		userService: mockUserService,
	}

	tests := []struct {
		name       string
		ctx        context.Context
		setupMocks func()
		wantErr    bool
		wantStatus int
	}{

		// ctx context.Context, jwtUser map[string]interface{}
		{
			name:       "Data User Date",
			ctx:        req.Context(),
			wantErr:    false,
			wantStatus: 200,
			setupMocks: func() {
				mockUserService.On("FindDate",
					c.Request().Context(),
					dataJwt).Return(model.User{}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()
			err := userController.Finddate(c)

			if (err != nil) != tt.wantErr {
				t.Errorf("UserController.FindDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				httpError, ok := err.(*echo.HTTPError)
				if !ok {
					t.Errorf("Expected echo.HTTPError, got %T", err)
				}
				assert.Equal(t, tt.wantStatus, http.StatusUnprocessableEntity)
				assert.Equal(t, "Internal Server Error", httpError.Message)
			} else {
				assert.Equal(t, tt.wantStatus, rec.Code)
				assert.Contains(t, rec.Body.String(), "Data User Date")
			}
		})
	}
}

func TestUserController_BuyPremium(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/user/buy-premium", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	db := &gorm.DB{}
	c.Set(constants.DBTransaction, db)
	dataJwt := map[string]interface{}{"phone": "216-253-6879"}
	c.Set("data_jwt", dataJwt)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserService := new(MockUserService)

	userController := UserController{
		userService: mockUserService,
	}

	tests := []struct {
		name       string
		ctx        context.Context
		setupMocks func()
		wantErr    bool
		wantStatus int
	}{

		{
			name:       "Success Upgrade Premium",
			ctx:        req.Context(),
			wantErr:    false,
			wantStatus: 200,
			setupMocks: func() {
				mockUserService.On("BuyPremium",
					db,
					dataJwt).Return(model.User{}, nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()
			err := userController.BuyPremium(c)

			if (err != nil) != tt.wantErr {
				t.Errorf("UserController.BuyPremium() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				httpError, ok := err.(*echo.HTTPError)
				if !ok {
					t.Errorf("Expected echo.HTTPError, got %T", err)
				}
				assert.Equal(t, tt.wantStatus, http.StatusUnprocessableEntity)
				assert.Equal(t, "Internal Server Error", httpError.Message)
			} else {
				assert.Equal(t, tt.wantStatus, rec.Code)
				assert.Contains(t, rec.Body.String(), "Success Upgrade Premium")
			}
		})
	}
}
