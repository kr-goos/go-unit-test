package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// ExampleTestSuite 구조체 정의
type ExampleTestSuite struct {
	suite.Suite
	value int
}

// SetupSuite: 모든 테스트 전에 한 번 실행됩니다.
func (suite *ExampleTestSuite) SetupSuite() {
	// 초기 설정 작업
}

// TearDownSuite: 모든 테스트 후에 한 번 실행됩니다.
func (suite *ExampleTestSuite) TearDownSuite() {
	// 정리 작업
}

// SetupTest: 각 테스트 전에 실행됩니다.
func (suite *ExampleTestSuite) SetupTest() {
	suite.value = 5
}

// TearDownTest: 각 테스트 후에 실행됩니다.
func (suite *ExampleTestSuite) TearDownTest() {
	// 각 테스트 후에 실행될 정리 작업
}

// 실제 테스트 함수
func (suite *ExampleTestSuite) TestExample() {
	assert.Equal(suite.T(), 5, suite.value)
}

// 테스트 스위트를 실행
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
