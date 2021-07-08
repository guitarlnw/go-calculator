package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GetPiTestSuite struct {
	suite.Suite
	resExpected float64
}

func (suite *GetPiTestSuite) SetupTest() {
	suite.resExpected = 3.141592653589793
}

func (suite *GetPiTestSuite) TestGetPiData() {
	assert.Equal(suite.T(), suite.resExpected, GetPI())
}

func TestGetPiTestSuite(t *testing.T) {
	suite.Run(t, new(GetPiTestSuite))
}
