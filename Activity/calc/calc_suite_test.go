package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AddTestSuite struct {
	suite.Suite
	//	VariablesAtStart int
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *AddTestSuite) TestAdd() {
	assert.Equal(suite.T(), 3, Add(1, 2))
	assert.Equal(suite.T(), 1, Add(1, 0))
	assert.Equal(suite.T(), 0, Add(2, -2))
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSuite(t *testing.T) {
	suite.Run(t, new(AddTestSuite))
}
