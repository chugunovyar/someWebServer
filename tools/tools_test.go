package tools

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestConvertTimeToTimestamp(t *testing.T) {
	resultTime := ConvertTimeToTimestamp("2006-01-02 15:04:05")
	t.Log(resultTime.String())
	expectedTime, _ := time.Parse(format, "2006-01-02 15:04:05")
	assert.Equal(t, resultTime, expectedTime)
}

func TestGetEnv(t *testing.T) {
	err := os.Setenv("FOO_STRING", "bar")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	assert.Equal(t, "bar", GetEnv("FOO_STRING", "bar"))
}

func TestGetEnvAsInt(t *testing.T) {
	err := os.Setenv("FOO_INT", "42")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	assert.Equal(t, int64(42), GetEnvAsInt("FOO_INT", 42))
}
