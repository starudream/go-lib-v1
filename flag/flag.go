package flag

import (
	"github.com/starudream/go-lib/internal/cobra"
)

var (
	AppendActiveHelp    = cobra.AppendActiveHelp
	GetActiveHelpConfig = cobra.GetActiveHelpConfig

	ArbitraryArgs = cobra.ArbitraryArgs
	ExactArgs     = cobra.ExactArgs
	MatchAll      = cobra.MatchAll
	MaximumNArgs  = cobra.MaximumNArgs
	MinimumNArgs  = cobra.MinimumNArgs
	NoArgs        = cobra.NoArgs
	OnlyValidArgs = cobra.OnlyValidArgs
	RangeArgs     = cobra.RangeArgs

	AddTemplateFunc  = cobra.AddTemplateFunc
	AddTemplateFuncs = cobra.AddTemplateFuncs

	CheckErr     = cobra.CheckErr
	OnFinalize   = cobra.OnFinalize
	OnInitialize = cobra.OnInitialize
)

type (
	Command = cobra.Command

	FParseErrWhitelist = cobra.FParseErrWhitelist

	CompletionOptions = cobra.CompletionOptions

	PositionalArgs = cobra.PositionalArgs
)
