package logger

import (
	"sync"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoggerNil(t *testing.T) {
	err := godotenv.Load("../.env")
	require.NoError(t, err)

	logger := GetLogger()
	require.NotNil(t, logger, "logger cannot be nil")

	var wg sync.WaitGroup
	for i := 0; i < 25; i++ {
		wg.Add(1)
		go func(i int) {
			l := GetLogger()
			assert.NotNil(t, l, "logger cannot be nil")
			l.Info().Int("thread index", i).Send()
			wg.Done()
		}(i)
	}
	wg.Wait()
}
