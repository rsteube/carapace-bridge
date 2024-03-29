package bridges

// https://fishshell.com/docs/current/commands.html
var fishBuiltins = []string{
	".",
	"abbr",        // manage fish abbreviations
	"alias",       // create a function
	"and",         // conditionally execute a command
	"argparse",    // parse options passed to a fish script or function
	"begin",       // start a new block of code
	"bg",          // send jobs to background
	"bind",        // handle fish key bindings
	"block",       // temporarily block delivery of events
	"breakpoint",  // launch debug mode
	"break",       // stop the current inner loop
	"builtin",     // run a builtin command
	"_",           // call fish’s translations
	"case",        // conditionally execute a block of commands
	"cd",          // change directory
	"cdh",         // change to a recently visited directory
	"commandline", // set or get the current command line buffer
	"command",     // run a program
	"complete",    // edit command-specific tab-completions
	"contains",    // test if a word is present in a list
	"continue",    // skip the remainder of the current iteration of the current inner loop
	"count",       // count the number of elements of a list
	"dirh",        // print directory history
	"dirs",        // print directory stack
	"disown",      // remove a process from the list of jobs
	"echo",        // display a line of text
	"else",        // execute command if a condition is not met
	"emit",        // emit a generic event
	"end",         // end a block of commands
	"eval",        // evaluate the specified commands
	"exec",        // execute command in current process
	"exit",        // exit the shell
	"false",       // return an unsuccessful result
	"fg",          // bring job to foreground
	"..fish",
	"fish_add_path",             // add to the path
	"fish_breakpoint_prompt",    // define the prompt when stopped at a breakpoint
	"fish_clipboard_copy",       // copy text to the system’s clipboard
	"fish_clipboard_paste",      // get text from the system’s clipboard
	"fish_command_not_found",    // what to do when a command wasn’t found
	"fish_config",               // start the web-based configuration interface
	"fish_default_key_bindings", // set emacs key bindings for fish
	"fish_delta",                // compare functions and completions to the default
	"fish_git_prompt",           // output git information for use in a prompt
	"fish_greeting",             // display a welcome message in interactive shells
	"fish_hg_prompt",            // output Mercurial information for use in a prompt
	"fish_indent",               // indenter and prettifier
	"fish_is_root_user",         // check if the current user is root
	"fish_key_reader",           // explore what characters keyboard keys send
	"fish_mode_prompt",          // define the appearance of the mode indicator
	"fish_opt",                  // create an option specification for the argparse command
	"fish_prompt",               // define the appearance of the command line prompt
	"fish_right_prompt",         // define the appearance of the right-side command line prompt
	"fish_status_to_signal",     // convert exit codes to human-friendly signals
	"fish_svn_prompt",           // output Subversion information for use in a prompt
	"fish",                      // the friendly interactive shell
	"fish_title",                // define the terminal’s title
	"fish_update_completions",   // update completions using manual pages
	"fish_vcs_prompt",           // output version control system information for use in a prompt
	"fish_vi_key_bindings",      // set vi key bindings for fish
	"for",                       // perform a set of commands multiple times
	"funced",                    // edit a function interactively
	"funcsave",                  // save the definition of a function to the user’s autoload directory
	"function",                  // create a function
	"functions",                 // print or erase functions
	"help",                      // display fish documentation
	"history",                   // show and manipulate command history
	"if",                        // conditionally execute a command
	"isatty",                    // test if a file descriptor is a terminal
	"jobs",                      // print currently running jobs
	"math",                      // perform mathematics calculations
	"nextd",                     // move forward through directory history
	"not",                       // negate the exit status of a job
	"open",                      // open file in its default application
	"or",                        // conditionally execute a command
	"path",                      // manipulate and check paths
	"popd",                      // move through directory stack
	"prevd",                     // move backward through directory history
	"printf",                    // display text according to a format string
	"prompt_hostname",           // print the hostname, shortened for use in the prompt
	"prompt_login",              // describe the login suitable for prompt
	"prompt_pwd",                // print pwd suitable for prompt
	"psub",                      // perform process substitution
	"pushd",                     // push directory to directory stack
	"pwd",                       // output the current working directory
	"random",                    // generate random number
	"read",                      // read line of input into variables
	"realpath",                  // convert a path to an absolute path without symlinks
	"return",                    // stop the current inner function
	"set_color",                 // set the terminal color
	"set",                       // display and change shell variables
	"source",                    // evaluate contents of file
	"status",                    // query fish runtime information
	"string-collect",            // join strings into one
	"string-escape",             // escape special characters
	"string-join0",              // join strings with zero bytes
	"string-join",               // join strings with delimiter
	"string-length",             // print string lengths
	"string-lower",              // convert strings to lowercase
	"string",                    // manipulate strings
	"string-match",              // match substrings
	"string-pad",                // pad strings to a fixed width
	"string-repeat",             // multiply a string
	"string-replace",            // replace substrings
	"string-shorten",            // shorten strings to a width, with an ellipsis
	"string-split0",             // split on zero bytes
	"string-split",              // split strings by delimiter
	"string-sub",                // extract substrings
	"string-trim",               // remove trailing whitespace
	"string-unescape",           // expand escape sequences
	"string-upper",              // convert strings to uppercase
	"suspend",                   // suspend the current shell
	"switch",                    // conditionally execute a block of commands
	"test",                      // perform tests on files and text
	"time",                      // measure how long a command or block takes
	"trap",                      // perform an action when the shell receives a signal
	"true",                      // return a successful result
	"type",                      // locate a command and describe its type
	"ulimit",                    // set or get resource usage limits
	"umask",                     // set or get the file creation mode mask
	"vared",                     // interactively edit the value of an environment variable
	"wait",                      // wait for jobs to complete
	"while",                     // perform a set of commands multiple times
}
