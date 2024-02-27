package logging

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

const defaultLevel = "info"

// New creates new instance of custom setup logging.
func New(conf Config) (*slog.Logger, error) {
	conf = conf.setDefault()

	lvl, err := conf.level()
	if err != nil {
		return nil, err
	}

	opts := &slog.HandlerOptions{Level: lvl}
	h := slog.NewJSONHandler(os.Stderr, opts)
	l := slog.New(h)
	return l, nil
}

// Default creates new instance of logging with defaults.
func Default() *slog.Logger {
	var c Config
	l, _ := New(c.setDefault())
	return l
}

var levels = map[string]slog.Level{
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
	"debug": slog.LevelDebug,
}

type Config struct {
	Level  string
	Format string
}

func (c Config) level() (slog.Level, error) {
	v, ok := levels[c.Level]
	if !ok {
		return slog.LevelInfo, fmt.Errorf("un-supported level: %s", c.Level)
	}
	return v, nil
}

func (c Config) setDefault() Config {
	c.Level = strings.ToLower(strings.TrimSpace(c.Level))
	if c.Level == "" {
		c.Level = defaultLevel
	}
	return c
}
