package services_test

import (
	"sampleprj/features/users"
	"sampleprj/features/users/mocks"
	"sampleprj/features/users/services"
	eMock "sampleprj/helper/enkrip/mocks"

	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	repo := mocks.NewRepository(t)
	enkrip := eMock.NewHashInterface(t)
	// s := services.New(&MockRepository{}, &MockEnkrip{})
	s := services.New(repo, enkrip)
	var inputData = users.User{Nama: "jerry", HP: "12345", Password: "altamantul123"}
	var repoData = users.User{Nama: "jerry", HP: "12345", Password: "some string"}
	var successReturnData = users.User{ID: uint(1), Nama: "jerry", HP: "12345"}
	var falseData = users.User{}
	t.Run("Success Case", func(t *testing.T) {
		enkrip.On("HashPassword", inputData.Password).Return("some string", nil).Once()
		repo.On("Insert", repoData).Return(successReturnData, nil).Once()
		res, err := s.Register(inputData)

		enkrip.AssertExpectations(t)
		repo.AssertExpectations(t)

		assert.Nil(t, err)
		assert.Equal(t, uint(1), res.ID)
		assert.Equal(t, "", res.Password)

	})

	t.Run("Failed Case", func(t *testing.T) {
		s = services.New(&FalseMockRepository{}, &MockEnkrip{})
		res, err := s.Register(falseData)
		assert.Error(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.Equal(t, "", res.Nama)
	})
}

type MockRepository struct{}

func (mr *MockRepository) Insert(newUser users.User) (users.User, error) {
	return users.User{ID: uint(1), Nama: "jerry", HP: "12345"}, nil
}
func (mr *MockRepository) Login(hp string) (users.User, error) {
	return users.User{}, nil
}

type FalseMockRepository struct{}

func (fmr *FalseMockRepository) Insert(newUser users.User) (users.User, error) {
	return users.User{}, errors.New("something happend")
}
func (fmr *FalseMockRepository) Login(hp string) (users.User, error) {
	return users.User{}, nil
}

type MockEnkrip struct{}

func (me *MockEnkrip) Compare(hashed string, input string) error {
	return nil
}
func (me *MockEnkrip) HashPassword(input string) (string, error) {
	return "some string", nil
}
