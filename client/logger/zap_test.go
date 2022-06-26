package logger

import (
	"fmt"
	"micobianParty/config"
	"os/exec"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func init() {
	path, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	if err = config.Confs.Load(strings.TrimSpace(string(path)) + "/config.test.yaml"); err != nil {
		fmt.Println(err.Error())
	}
}

func TestAdd(t *testing.T) {
	logger := logs{}
	testCases := []struct {
		desc  string
		key   string
		field interface{}
	}{
		{
			desc:  "a",
			key:   "key",
			field: "some value",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			fields := logger.Add(tc.key, tc.field)
			if len(fields.fields) == 0 {
				assert.Equal(t, "0", fields.fields)
			}
		})
	}
}

func TestAppend(t *testing.T) {
	logger := logs{}
	testCases := []struct {
		desc   string
		fields []zapcore.Field
	}{
		{
			desc:   "a",
			fields: []zapcore.Field{},
		},
		{
			desc: "b",
			fields: []zapcore.Field{
				{Key: "some", Interface: "thing"},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			fields := logger.Append(tc.fields...)
			if len(fields.fields) != len(tc.fields) {
				assert.Equal(t, tc.fields, fields.fields)
			}
		})
	}
}

func TestCommit(t *testing.T) {
	logger := logs{}

	testCases := []struct {
		desc string
		msg  string
		lvl  zapcore.Level
	}{
		{
			desc: "a",
			msg:  "hello commit",
			lvl:  zapcore.ErrorLevel,
		},
		{
			desc: "b",
			msg:  "hello commit",
			lvl:  zapcore.InfoLevel,
		},
		{
			desc: "c",
			msg:  "hello commit",
			lvl:  zapcore.WarnLevel,
		},
		{
			desc: "d",
			msg:  "hello commit",
			lvl:  zapcore.DebugLevel,
		},
		{
			desc: "e",
			msg:  "hello commit",
			lvl:  zapcore.FatalLevel,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			logger.level = tc.lvl
			logger.logger = zap.L()

			logger.Commit(tc.msg)
		})
	}
}

func TestLevel(t *testing.T) {
	logger := logs{}

	testCases := []struct {
		desc string
		lvl  zapcore.Level
	}{
		{
			desc: "a",
			lvl:  zapcore.DebugLevel,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			logger.Level(tC.lvl)
		})
	}
}

func TestDevelopment(t *testing.T) {
	logger := logs{}

	testCases := []struct {
		desc string
		lvl  zapcore.Level
	}{
		{
			desc: "a",
			lvl:  zapcore.DebugLevel,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			l := logger.Development()
			if len(l.fields) == 0 {
				assert.Equal(t, "must be filled development keys", l.fields)
			}
		})
	}
}

func TestPrepare(t *testing.T) {

	testCases := []struct {
		desc string
		z    *zap.Logger
	}{
		{
			desc: "a",
			z:    zap.L(),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			Prepare(tC.z)
		})
	}
}
func TestGetZapLogger(t *testing.T) {
	testCases := []struct {
		desc  string
		debug bool
	}{
		{
			desc:  "a",
			debug: true,
		},
		{
			desc:  "b",
			debug: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			once = sync.Once{}
			GetZapLogger(tC.debug)
		})
	}
}
