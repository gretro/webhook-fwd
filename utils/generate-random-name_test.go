package utils_test

import (
	"regexp"
	"testing"

	"github.com/gretro/webhook-fwd/utils"
	"github.com/stretchr/testify/suite"
)

type GenerateRandomNameTestSuite struct {
	suite.Suite
}

func TestGenerateRandomName(t *testing.T) {
	suite.Run(t, new(GenerateRandomNameTestSuite))
}

var regex = regexp.MustCompile("^([a-zA-Z0-9]+)-([a-zA-Z0-9]+)$")

func (suite *GenerateRandomNameTestSuite) Test_ShouldGenerateRandomName() {
	result := utils.GenerateRandomName()

	suite.Regexp(regex, result)
}
