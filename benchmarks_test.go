package slog_benchmarks

import (
	"io"
	"testing"
	"time"
	"math/big"

	"github.com/ethereum/go-ethereum/log"
)

func lazy() interface{} {
	return time.Now()
}

func BenchmarkGloggerDisabled(b *testing.B) {
	glogHandler := log.NewGlogHandler(log.StreamHandler(io.Discard, log.TerminalFormat(false)))
	glogHandler.Verbosity(log.LvlError)
	log.Root().SetHandler(glogHandler)

	num, _ := new(big.Int).SetString("123412312901879817298712498719248791283798124", 10)

	for i := 0; i < b.N; i++ {
		log.Info("a message", "big.Int", num, "lazy(time)", log.Lazy{lazy})
	}
}

func BenchmarkGloggerTerminal(b *testing.B) {
	glogHandler := log.NewGlogHandler(log.StreamHandler(io.Discard, log.TerminalFormat(false)))
	glogHandler.Verbosity(log.LvlInfo)
	log.Root().SetHandler(glogHandler)

	num, _ := new(big.Int).SetString("123412312901879817298712498719248791283798124", 10)

	for i := 0; i < b.N; i++ {
		log.Info("a message", "big.Int", num, "lazy(time)", log.Lazy{lazy})
	}
}

func BenchmarkGloggerLogfmt(b *testing.B) {
	glogHandler := log.NewGlogHandler(log.StreamHandler(io.Discard, log.LogfmtFormat()))
	glogHandler.Verbosity(log.LvlInfo)
	log.Root().SetHandler(glogHandler)

	num, _ := new(big.Int).SetString("123412312901879817298712498719248791283798124", 10)

	for i := 0; i < b.N; i++ {
		log.Info("a message", "big.Int", num, "lazy(time)", log.Lazy{lazy})
	}
}

func BenchmarkGloggerJson(b *testing.B) {
	glogHandler := log.NewGlogHandler(log.StreamHandler(io.Discard, log.JSONFormat()))
	glogHandler.Verbosity(log.LvlInfo)
	log.Root().SetHandler(glogHandler)

	num, _ := new(big.Int).SetString("123412312901879817298712498719248791283798124", 10)

	for i := 0; i < b.N; i++ {
		log.Info("a message", "big.Int", num, "lazy(time)", log.Lazy{lazy})
	}
}
