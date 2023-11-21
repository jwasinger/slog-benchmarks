package slog_benchmarks

import (
	"io"
	//"os"
	"testing"

	"github.com/ethereum/go-ethereum/log"
)

func BenchmarkJson(b *testing.B) {
	log.SetDefault(log.NewLogger(log.JSONHandler(io.Discard)))

	for i := 0; i < b.N; i++ {
		log.Info("a message", "foo", "bar", "baz", "bat")		
	}
}

func BenchmarkGloggerDisabled(b *testing.B) {
	glogHandler := log.NewGlogHandler(log.NewTerminalHandler(io.Discard, false))
	glogHandler.Verbosity(log.LevelError)
	log.SetDefault(log.NewLogger(glogHandler))

	for i := 0; i < b.N; i++ {
		log.Info("foo", "bar", "baz", "bat")		
	}
	
}

func BenchmarkGloggerDiscard(b *testing.B) {
	glogHandler := log.NewGlogHandler(log.DiscardHandler())
	glogHandler.Verbosity(log.LevelInfo)
	log.SetDefault(log.NewLogger(glogHandler))

	for i := 0; i < b.N; i++ {
		log.Info("foo", "bar", "baz", "bat")		
	}
	
}

func BenchmarkGloggerTerminal(b *testing.B) {
	glogHandler := log.NewGlogHandler(log.NewTerminalHandler(io.Discard, false))
	glogHandler.Verbosity(log.LevelInfo)
	log.SetDefault(log.NewLogger(glogHandler))

	for i := 0; i < b.N; i++ {
		log.Info("foo", "bar", "baz", "bat")		
	}
	
}

func BenchmarkGloggerLogfmt(b *testing.B) {
	glogHandler := log.NewGlogHandler(log.LogfmtHandler(io.Discard))
	glogHandler.Verbosity(log.LevelInfo)
	log.SetDefault(log.NewLogger(glogHandler))

	for i := 0; i < b.N; i++ {
		log.Info("foo", "bar", "baz", "bat")		
	}
	
}

func BenchmarkGloggerJson(b *testing.B) {
	glogHandler := log.NewGlogHandler(log.NewTerminalHandler(io.Discard, false))
	glogHandler.Verbosity(log.LevelInfo)
	log.SetDefault(log.NewLogger(glogHandler))

	for i := 0; i < b.N; i++ {
		log.Info("foo", "bar", "baz", "bat")		
	}
	
}
