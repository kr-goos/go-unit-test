package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Service 인터페이스 정의
type Service interface {
	GetData(id string) (string, error)
}

// MockService 구조체 정의
type MockService struct {
	mock.Mock
}

// GetData 메서드 구현
func (m *MockService) GetData(id string) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)
}

func TestGetData(t *testing.T) {
	// MockService 인스턴스 생성
	mockService := new(MockService)

	// GetData 메서드 호출에 대한 기대와 반환 값 설정
	mockService.On("GetData", "123").Return("Mocked Data", nil)

	// 테스트 대상 함수 호출
	data, err := mockService.GetData("123")

	// 결과 검증
	assert.NoError(t, err)
	assert.Equal(t, "Mocked Data", data)

	// 설정된 기대가 충족되었는지 확인
	mockService.AssertExpectations(t)
}
