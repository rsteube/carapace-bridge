package bridges

// print -roC1 -- ${(k)builtins}
var zshBuiltins = []string{
	"-",
	".",
	":",
	"[",
	"alias",
	"_arguments",
	"autoload",
	"bg",
	"bindkey",
	"break",
	"builtin",
	"bye",
	"cd",
	"chdir",
	"command",
	"command_names",
	"compadd",
	"comparguments",
	"compcall",
	"compctl",
	"compdef",
	"compdescribe",
	"compfiles",
	"compgroups",
	"complete",
	"complete_debug",
	"complete_help",
	"complete_help_generic",
	"completers",
	"complete_tag",
	"comp_locale",
	"compquote",
	"compset",
	"comptags",
	"comptry",
	"compvalues",
	"continue",
	"declare",
	"dirs",
	"disable",
	"disown",
	"echo",
	"echotc",
	"echoti",
	"emulate",
	"enable",
	"eval",
	"exec",
	"exit",
	"export",
	"false",
	"fc",
	"fg",
	"float",
	"functions",
	"getln",
	"getopts",
	"hash",
	"history",
	"integer",
	"jobs",
	"kill",
	"let",
	"limit",
	"local",
	"log",
	"logout",
	"noglob",
	"popd",
	"print",
	"printf",
	"private",
	"pushd",
	"pushln",
	"pwd",
	"r",
	"read",
	"readonly",
	"rehash",
	"return",
	"sched",
	"set",
	"setopt",
	"shift",
	"source",
	"strftime",
	"suspend",
	"test",
	"times",
	"trap",
	"true",
	"ttyctl",
	"type",
	"typeset",
	"ulimit",
	"umask",
	"unalias",
	"unfunction",
	"unhash",
	"unlimit",
	"unset",
	"unsetopt",
	"vared",
	"wait",
	"whence",
	"where",
	"which",
	"zcompile",
	"zformat",
	"zle",
	"zmodload",
	"zparseopts",
	"zregexparse",
	"zstat",
	"zstyle",
}
