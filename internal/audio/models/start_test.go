package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDownloadAndExtract(t *testing.T) {
	err := DownloadAndExtract(srcUrl)

	assert.Nil(t, err)
}
