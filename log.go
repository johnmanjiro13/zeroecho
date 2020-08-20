package zeroecho

import (
	"fmt"
	"io"

	"github.com/labstack/gommon/log"
	"github.com/rs/zerolog"
)

type Logger struct {
	log    zerolog.Logger
	out    io.Writer
	prefix string
	level  log.Lvl
}

func New(out io.Writer, prefix string) *Logger {
	return &Logger{
		log:    zerolog.New(out).With().Str("prefix", prefix).Timestamp().Logger(),
		out:    out,
		prefix: prefix,
	}
}

func (l *Logger) Output() io.Writer {
	return l.out
}

func (l *Logger) SetOutput(newOut io.Writer) {
	l.out = newOut
	l.log.Output(newOut)
}

func (l *Logger) Prefix() string {
	return l.prefix
}

func (l *Logger) SetPrefix(p string) {
	withPrefix := zerolog.New(l.out).With().Str("prefix", p).Timestamp().Logger()
	l.log = withPrefix
	l.prefix = p
}

func (l *Logger) Level() log.Lvl {
	return l.level
}

func (l *Logger) SetLevel(lvl log.Lvl) {
	level := levels[lvl]
	l.level = lvl
	l.log.Level(level)
}

func (l *Logger) SetHeader(h string) {
	// not implemented.
}

func (l *Logger) Print(i ...interface{}) {
	l.log.Print(i)
}

func (l *Logger) Printf(format string, args ...interface{}) {
	l.log.Printf(format, args)
}

func (l *Logger) Printj(j log.JSON) {
	l.logJSON(l.log.Debug(), j)
}

func (l *Logger) Debug(i ...interface{}) {
	l.log.Debug().Msg(fmt.Sprint(i...))
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.log.Debug().Msgf(format, args...)
}

func (l *Logger) Debugj(j log.JSON) {
	l.logJSON(l.log.Debug(), j)
}

func (l *Logger) Info(i ...interface{}) {
	l.log.Info().Msg(fmt.Sprint(i...))
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.log.Info().Msgf(format, args...)
}

func (l *Logger) Infoj(j log.JSON) {
	l.logJSON(l.log.Info(), j)
}

func (l *Logger) Warn(i ...interface{}) {
	l.log.Warn().Msg(fmt.Sprint(i...))
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.log.Warn().Msgf(format, args...)
}

func (l *Logger) Warnj(j log.JSON) {
	l.logJSON(l.log.Warn(), j)
}

func (l *Logger) Error(i ...interface{}) {
	l.log.Error().Msg(fmt.Sprint(i...))
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.log.Error().Msgf(format, args...)
}

func (l *Logger) Errorj(j log.JSON) {
	l.logJSON(l.log.Error(), j)
}

func (l *Logger) Fatal(i ...interface{}) {
	l.log.Fatal().Msg(fmt.Sprint(i...))
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.log.Fatal().Msgf(format, args...)
}

func (l *Logger) Fatalj(j log.JSON) {
	l.logJSON(l.log.Fatal(), j)
}

func (l *Logger) Panic(i ...interface{}) {
	l.log.Panic().Msg(fmt.Sprint(i...))
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.log.Panic().Msgf(format, args...)
}

func (l *Logger) Panicj(j log.JSON) {
	l.logJSON(l.log.Panic(), j)
}

func (l *Logger) logJSON(event *zerolog.Event, j log.JSON) {
	for k, v := range j {
		event = event.Interface(k, v)
	}
	event.Msg("")
}
