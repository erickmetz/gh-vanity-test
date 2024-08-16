package githubvanitytest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGithubVanityOutput(t *testing.T) {
	str := PerformGhVanityTestString()
	assert.Equal(t, str, "This is a Golang Vanity URL test that points to GitHub")
}
