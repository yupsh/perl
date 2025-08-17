package command

type Script string
type Module string
type LibPath string
type Encoding string

type InPlaceFlag bool

const (
	InPlace   InPlaceFlag = true
	NoInPlace InPlaceFlag = false
)

type PrintFlag bool

const (
	Print   PrintFlag = true
	NoPrint PrintFlag = false
)

type LoopFlag bool

const (
	Loop   LoopFlag = true
	NoLoop LoopFlag = false
)

type AutoSplitFlag bool

const (
	AutoSplit   AutoSplitFlag = true
	NoAutoSplit AutoSplitFlag = false
)

type CheckSyntaxFlag bool

const (
	CheckSyntax   CheckSyntaxFlag = true
	NoCheckSyntax CheckSyntaxFlag = false
)

type WarningsFlag bool

const (
	Warnings   WarningsFlag = true
	NoWarnings WarningsFlag = false
)

type StrictFlag bool

const (
	Strict   StrictFlag = true
	NoStrict StrictFlag = false
)

type DebugFlag bool

const (
	Debug   DebugFlag = true
	NoDebug DebugFlag = false
)

type TaintFlag bool

const (
	Taint   TaintFlag = true
	NoTaint TaintFlag = false
)

type flags struct {
	Script      Script
	Module      Module
	LibPath     LibPath
	Encoding    Encoding
	InPlace     InPlaceFlag
	Print       PrintFlag
	Loop        LoopFlag
	AutoSplit   AutoSplitFlag
	CheckSyntax CheckSyntaxFlag
	Warnings    WarningsFlag
	Strict      StrictFlag
	Debug       DebugFlag
	Taint       TaintFlag
}

func (s Script) Configure(flags *flags)          { flags.Script = s }
func (m Module) Configure(flags *flags)          { flags.Module = m }
func (l LibPath) Configure(flags *flags)         { flags.LibPath = l }
func (e Encoding) Configure(flags *flags)        { flags.Encoding = e }
func (i InPlaceFlag) Configure(flags *flags)     { flags.InPlace = i }
func (p PrintFlag) Configure(flags *flags)       { flags.Print = p }
func (l LoopFlag) Configure(flags *flags)        { flags.Loop = l }
func (a AutoSplitFlag) Configure(flags *flags)   { flags.AutoSplit = a }
func (c CheckSyntaxFlag) Configure(flags *flags) { flags.CheckSyntax = c }
func (w WarningsFlag) Configure(flags *flags)    { flags.Warnings = w }
func (s StrictFlag) Configure(flags *flags)      { flags.Strict = s }
func (d DebugFlag) Configure(flags *flags)       { flags.Debug = d }
func (t TaintFlag) Configure(flags *flags)       { flags.Taint = t }
