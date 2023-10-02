package slog_benchmarks

import (
	"io"
	//"os"
	"testing"

	"github.com/ethereum/go-ethereum/log"
)

func BenchmarkJson(b *testing.B) {
        log.Root().SetHandler(log.LvlFilterHandler(log.LvlTrace, log.StreamHandler(io.Discard, log.JSONFormat())))

        for i := 0; i < b.N; i++ {
                log.Info("a message", "foo", "bar", "baz", "bat")
        }
}

func BenchmarkGloggerNotEnabled(b *testing.B) {
	glogHandler := log.NewGlogHandler(log.StreamHandler(io.Discard, log.TerminalFormat(false)))
	glogHandler.Verbosity(log.LvlError)
	log.Root().SetHandler(glogHandler)

	for i := 0; i < b.N; i++ {
		log.Info("foo", "bar", "baz", "bat")		
	}
	
}

func BenchmarkGloggerEnabledDiscardHandler(b *testing.B) {
	glogHandler := log.NewGlogHandler(log.DiscardHandler())
	glogHandler.Verbosity(log.LvlInfo)
	log.Root().SetHandler(glogHandler)

	for i := 0; i < b.N; i++ {
		log.Info("foo", "bar", "baz", "bat")		
	}
	
}

func BenchmarkGloggerEnabledStreamHandler(b *testing.B) {
	glogHandler := log.NewGlogHandler(log.StreamHandler(io.Discard, log.TerminalFormat(false)))
	glogHandler.Verbosity(log.LvlInfo)
	log.Root().SetHandler(glogHandler)

	for i := 0; i < b.N; i++ {
		log.Info("foo", "bar", "baz", "bat")		
	}
	
}
