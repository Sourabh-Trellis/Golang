package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ExampleSuite struct {
	suite.Suite
	Length  float64
	Breadth float64
}

func TestExampleSuit(t *testing.T) {
	suite.Run(t, &ExampleSuite{})
}

// #------------------------------------------------------
// TEST 1
// #------------------------------------------------------
func (es *ExampleSuite) TestTrue() {
	es.T().Log("In TestTrue")
	fmt.Println(es.Length)
	fmt.Println(es.Breadth)
	es.True(true)
}

// #------------------------------------------------------
// TEST 2
// #------------------------------------------------------
func (es *ExampleSuite) TestFalse() {
	es.T().Log("in TestFalse")
	fmt.Println(es.Length)
	fmt.Println(es.Breadth)
	es.True(true)
}
func (es *ExampleSuite) SetupSuite() {
	es.T().Log("Setup test Suite")
	es.Length = 22.22
	es.Breadth = 33.33
}

func (es *ExampleSuite) TearDownSuite() {
	es.T().Log("Teardown Suite")
}

func (es *ExampleSuite) SetupTest() {
	es.T().Log("Setup Test")
}

func (es *ExampleSuite) TearDownTest() {
	es.T().Log("Teardown Test")
}

func (es *ExampleSuite) BeforeTest(suiteName,methodName string) {
	if methodName == "TestTrue" {
		es.T().Log("before test TestTrue")
	} else if methodName == "TestFalse" {
		es.T().Log("before test TestFalse")
	}

}
