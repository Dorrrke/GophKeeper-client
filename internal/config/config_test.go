package config

import (
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfit(t *testing.T) {
	os.Args = []string{}
	type want struct {
		cfg Config
	}
	type test struct {
		name     string
		flags    []string
		envSetup func()
		want
	}

	tests := []test{
		{
			name:  "Test ReadConfig function #1; Default call",
			flags: []string{"test", "-a", "localhost:9191", "-d", "storage/test.db"},
			want: want{
				cfg: Config{
					ServerAddr: "localhost:9191",
					DBPath:     "storage/test.db",
				},
			},
		},
		{
			name:  "Test ReadConfig function #2; Call with env",
			flags: []string{"test", "-a", ":9081", "-d", "storage/test.db"},
			envSetup: func() {
				t.Setenv("SERVER_ADDR", "12.12.12.12:4545")
				t.Setenv("DATA_BASE_PATH", "db_test.db")
			},
			want: want{
				cfg: Config{
					ServerAddr: "12.12.12.12:4545",
					DBPath:     "db_test.db",
				},
			},
		},
		{
			name:  "Test ReadConfig function #3; Call without flags and env",
			flags: []string{""},
			want: want{
				cfg: Config{
					ServerAddr: "localhost:8080",
					DBPath:     "internal/storage/gophkeeper.db",
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			os.Args = tc.flags
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			if tc.envSetup != nil {
				tc.envSetup()
				defer os.Unsetenv("SERVER_ADDR")
				defer os.Unsetenv("DATA_BASE_PATH")
			}
			testCfg := ReadConfig()
			assert.Equal(t, tc.want.cfg, *testCfg)
		})
	}
}
