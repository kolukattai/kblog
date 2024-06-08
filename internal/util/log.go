package util

import (
	"os"

	"github.com/charmbracelet/log"
)

func Error(err string) {
	log.Error(err)
	os.Exit(1)
}
