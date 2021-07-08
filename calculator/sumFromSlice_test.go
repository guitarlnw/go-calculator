package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SumFromSliceTestSuite struct {
	suite.Suite
	resExpected int
}

func (suite *SumFromSliceTestSuite) TestWhenSentSliceIntoFuncThenReturnSumData() {
	suite.resExpected = 10
	mockData := []int{2, 2, 5, 1}
	assert.Equal(suite.T(), suite.resExpected, SumFromSlice(mockData))
}

func (suite *SumFromSliceTestSuite) TestWhenSentSliceToBeEmtryIntoFuncThenReturnZero() {
	suite.resExpected = 0
	mockData := []int{}
	assert.Equal(suite.T(), suite.resExpected, SumFromSlice(mockData))
}

func TestSumFromSliceTestSuite(t *testing.T) {
	suite.Run(t, new(SumFromSliceTestSuite))
}
