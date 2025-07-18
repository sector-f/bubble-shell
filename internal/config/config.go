package config

import (
	"context"

	"github.com/DomBlack/bubble-shell/pkg/config/keymap"
	"github.com/DomBlack/bubble-shell/pkg/config/styles"
)

// Config is the configuration for the shell
type Config struct {
	// The file to store the history in (if not absolute will be relative to $HOME)
	//
	// If blank no history will be stored
	HistoryFile string

	KeyMap keymap.KeyMap // The key map to use
	Styles styles.Styles // The styles to use

	MaxStackFrames  int  // The maximum number of stack frames to show in errors
	ShowStackFrames bool // Whether or not to show stack frames on error

	// Packages to filter from the stack trace rendering of errors
	//
	// By default the shell will filter out packages related to running
	// the shell itself as well as the Go runtime package.
	//
	// If empty no filtering will be done
	PackagesToFilterFromStack []string

	// InlineShell will cause the shell to be rendered inline
	// rather than taking over the whole terminal
	InlineShell bool

	// RootContext is the context that will be used for the
	// commands when they run
	RootContext context.Context

	// PromptFunc is a function that returns the prompt to be used
	PromptFunc func() string
}

// Default returns a default configuration for the shell
func Default() *Config {
	return &Config{
		HistoryFile:     ".bubble-shell-history",
		RootContext:     context.Background(),
		PromptFunc:      func() string { return "> " },
		KeyMap:          keymap.Default,
		Styles:          styles.Default,
		MaxStackFrames:  8,
		ShowStackFrames: true,
		PackagesToFilterFromStack: []string{
			"runtime",
			"testing",
			"github.com/spf13/cobra",
			"github.com/cockroachdb/errors",
			"github.com/charmbracelet/bubbletea",
			"github.com/DomBlack/bubble-shell",
		},
	}
}
