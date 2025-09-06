package opt

// Custom types for parameters
type Script string
type Module string
type LibPath string
type Encoding string

// Boolean flag types with constants
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

// Flags represents the configuration options for perl commands
type Flags struct {
	Script      Script          // Perl script code (-e)
	Module      Module          // Module to use (-M)
	LibPath     LibPath         // Library path (-I)
	Encoding    Encoding        // Input/output encoding
	InPlace     InPlaceFlag     // Edit files in place (-i)
	Print       PrintFlag       // Print lines (-p)
	Loop        LoopFlag        // Loop over input (-n)
	AutoSplit   AutoSplitFlag   // Auto-split lines (-a)
	CheckSyntax CheckSyntaxFlag // Check syntax only (-c)
	Warnings    WarningsFlag    // Enable warnings (-w)
	Strict      StrictFlag      // Use strict mode
	Debug       DebugFlag       // Debug mode (-d)
	Taint       TaintFlag       // Taint mode (-T)
}

// Configure methods for the opt system
func (s Script) Configure(flags *Flags)          { flags.Script = s }
func (m Module) Configure(flags *Flags)          { flags.Module = m }
func (l LibPath) Configure(flags *Flags)         { flags.LibPath = l }
func (e Encoding) Configure(flags *Flags)        { flags.Encoding = e }
func (i InPlaceFlag) Configure(flags *Flags)     { flags.InPlace = i }
func (p PrintFlag) Configure(flags *Flags)       { flags.Print = p }
func (l LoopFlag) Configure(flags *Flags)        { flags.Loop = l }
func (a AutoSplitFlag) Configure(flags *Flags)   { flags.AutoSplit = a }
func (c CheckSyntaxFlag) Configure(flags *Flags) { flags.CheckSyntax = c }
func (w WarningsFlag) Configure(flags *Flags)    { flags.Warnings = w }
func (s StrictFlag) Configure(flags *Flags)      { flags.Strict = s }
func (d DebugFlag) Configure(flags *Flags)       { flags.Debug = d }
func (t TaintFlag) Configure(flags *Flags)       { flags.Taint = t }
