// Code generated by github.com/twpayne/chezmoi/internal/generate-helps. DO NOT EDIT.

package cmd

type help struct {
	long    string
	example string
}

var helps = map[string]help{
	"add": {
		long: "" +
			"Description:\n" +
			"  Add *targets* to the source state. If any target is already in the source\n" +
			"  state, then its source state is replaced with its current state in the\n" +
			"  destination directory. The `add` command accepts additional flags:\n" +
			"\n" +
			"  `--autotemplate`\n" +
			"\n" +
			"  Automatically generate a template by replacing strings with variable names\n" +
			"  from the `data` section of the config file. Longer subsitutions occur before\n" +
			"  shorter ones. This implies the `--template` option.\n" +
			"\n" +
			"  `-e`, `--empty`\n" +
			"\n" +
			"  Set the `empty` attribute on added files.\n" +
			"\n" +
			"  `-f`, `--force`\n" +
			"\n" +
			"  Add *targets*, even if doing so would cause a source template to be\n" +
			"  overwritten.\n" +
			"\n" +
			"  `-x`, `--exact`\n" +
			"\n" +
			"  Set the `exact` attribute on added directories.\n" +
			"\n" +
			"  `-p`, `--prompt`\n" +
			"\n" +
			"  Interactively prompt before adding each file.\n" +
			"\n" +
			"  `-r`, `--recursive`\n" +
			"\n" +
			"  Recursively add all files, directories, and symlinks.\n" +
			"\n" +
			"  `-T`, `--template`\n" +
			"\n" +
			"  Set the `template` attribute on added files and symlinks.",
		example: "" +
			"  chezmoi add ~/.bashrc\n" +
			"  chezmoi add ~/.gitconfig --template\n" +
			"  chezmoi add ~/.vim --recursive\n" +
			"  chezmoi add ~/.oh-my-zsh --exact --recursive",
	},
	"apply": {
		long: "" +
			"Description:\n" +
			"  Ensure that *targets* are in the target state, updating them if necessary. If\n" +
			"  no targets are specified, the state of all targets are ensured.",
		example: "" +
			"  chezmoi apply\n" +
			"  chezmoi apply --dry-run --verbose\n" +
			"  chezmoi apply ~/.bashrc",
	},
	"archive": {
		long: "" +
			"Description:\n" +
			"  Write a tar archive of the target state to stdout. This can be piped into\n" +
			"  `tar` to inspect the target state.",
		example: "" +
			"  chezmoi archive | tar tvf -",
	},
	"cat": {
		long: "" +
			"Description:\n" +
			"  Write the target state of *targets*  to stdout. *targets* must be files or\n" +
			"  symlinks. For files, the target file contents are written. For symlinks, the\n" +
			"  target target is written.",
		example: "" +
			"  chezmoi cat ~/.bashrc",
	},
	"cd": {
		long: "" +
			"Description:\n" +
			"  Launch a shell in the source directory. chezmoi will launch the command set by\n" +
			"  the `cd.command` confiuration varaible. If this is not set, chezmoi will\n" +
			"  attempt to detect your shell and will finally fall back to an OS-specific\n" +
			"  default.",
		example: "" +
			"  chezmoi cd",
	},
	"chattr": {
		long: "" +
			"Description:\n" +
			"  Change the attributes of *targets*. *attributes* specifies which attributes to\n" +
			"  modify. Add attributes by specifying them or their abbreviations directly,\n" +
			"  optionally prefixed with a plus sign (`+`). Remove attributes by prefixing\n" +
			"  them or their attributes with the string `no` or a minus sign (`-`). The\n" +
			"  available attributes and their abbreviations are:\n" +
			"\n" +
			"    ATTRIBUTE  | ABBREVIATION\n" +
			"  -------------+---------------\n" +
			"    empty      | e\n" +
			"    encrypted  | none\n" +
			"    exact      | none\n" +
			"    executable | x\n" +
			"    private    | p\n" +
			"    template   | t\n" +
			"\n" +
			"  Multiple attributes modifications may be specified by separating them with a\n" +
			"  comma (`,`).",
		example: "" +
			"  chezmoi chattr template ~/.bashrc\n" +
			"  chezmoi chattr noempty ~/.profile\n" +
			"  chezmoi chattr private,template ~/.netrc",
	},
	"completion": {
		long: "" +
			"Description:\n" +
			"  Generate shell completion code for the specified shell (`bash`, `fish`, or\n" +
			"  `zsh`).\n" +
			"\n" +
			"  `--output`, `-o` *filename*\n" +
			"\n" +
			"  Write the shell completion code to *filename* instead of stdout.",
		example: "" +
			"  chezmoi completion bash\n" +
			"  chezmoi completion fish --output ~/.config/fish/completions/chezmoi.fish",
	},
	"data": {
		long: "" +
			"Description:\n" +
			"  Write the computed template data in JSON format to stdout. The `data` command\n" +
			"  accepts additional flags:\n" +
			"\n" +
			"  `-f`, `--format` *format*\n" +
			"\n" +
			"  Print the computed template data in the given format. The accepted formats are\n" +
			"  `json` (JSON), `toml` (TOML), and `yaml` (YAML).",
		example: "" +
			"  chezmoi data\n" +
			"  chezmoi data --format=yaml",
	},
	"diff": {
		long: "" +
			"Description:\n" +
			"  Print the difference between the target state and the destination state for\n" +
			"  *targets*. If no targets are specified, print the differences for all targets.\n" +
			"\n" +
			"  If a `diff.pager` command is set in the configuration file then the output\n" +
			"  will be piped into it.\n" +
			"\n" +
			"  `-f`, `--format` *format*\n" +
			"\n" +
			"  Print the diff in *format*. The format can be set with the `diff.format`\n" +
			"  variable in the configuration file. Valid formats are:\n" +
			"\n" +
			"  ##### `chezmoi`\n" +
			"\n" +
			"  A mix of unified diffs and pseudo shell commands, including scripts,\n" +
			"  equivalent to `chezmoi apply --dry-run --verbose`.\n" +
			"\n" +
			"  ##### `git`\n" +
			"\n" +
			"  A git format diff https://git-scm.com/docs/diff-format, excluding scripts. In\n" +
			"  version 2.0.0 of chezmoi, `git` format diffs will become the default and\n" +
			"  include scripts and the `chezmoi` format will be removed.\n" +
			"\n" +
			"  `--no-pager`\n" +
			"\n" +
			"  Do not use the pager.",
		example: "" +
			"  chezmoi diff\n" +
			"  chezmoi diff ~/.bashrc\n" +
			"  chezmoi diff --format=git",
	},
	"docs": {
		long: "" +
			"Description:\n" +
			"  Print the documentation page matching the regular expression *regexp*.\n" +
			"  Matching is case insensitive. If no pattern is given, print `REFERENCE.md`.",
		example: "" +
			"  chezmoi docs\n" +
			"  chezmoi docs faq\n" +
			"  chezmoi docs howto",
	},
	"doctor": {
		long: "" +
			"Description:\n" +
			"  Check for potential problems.",
		example: "" +
			"  chezmoi doctor",
	},
	"dump": {
		long: "" +
			"Description:\n" +
			"  Dump the target state in JSON format. If no targets are specified, then the\n" +
			"  entire target state. The `dump` command accepts additional arguments:\n" +
			"\n" +
			"  `-f`, `--format` *format*\n" +
			"\n" +
			"  Print the target state in the given format. The accepted formats are `json`\n" +
			"  (JSON) and `yaml` (YAML).",
		example: "" +
			"  chezmoi dump ~/.bashrc\n" +
			"  chezmoi dump --format=yaml",
	},
	"edit": {
		long: "" +
			"Description:\n" +
			"  Edit the source state of *targets*, which must be files or symlinks. If no\n" +
			"  targets are given the the source directory itself is opened with `$EDITOR`.\n" +
			"  The `edit` command accepts additional arguments:\n" +
			"\n" +
			"  `-a`, `--apply`\n" +
			"\n" +
			"  Apply target immediately after editing. Ignored if there are no targets.\n" +
			"\n" +
			"  `-d`, `--diff`\n" +
			"\n" +
			"  Print the difference between the target state and the actual state after\n" +
			"  editing.. Ignored if there are no targets.\n" +
			"\n" +
			"  `-p`, `--prompt`\n" +
			"\n" +
			"  Prompt before applying each target.. Ignored if there are no targets.",
		example: "" +
			"  chezmoi edit ~/.bashrc\n" +
			"  chezmoi edit ~/.bashrc --apply --prompt\n" +
			"  chezmoi edit",
	},
	"edit-config": {
		long: "" +
			"Description:\n" +
			"  Edit the configuration file.\n" +
			"\n" +
			"  `edit-config` examples\n" +
			"\n" +
			"    chezmoi edit-config",
	},
	"execute-template": {
		long: "" +
			"Description:\n" +
			"  Execute *templates*. This is useful for testing templates or for calling\n" +
			"  chezmoi from other scripts. *templates* are interpeted as literal templates,\n" +
			"  with no whitespace added to the output between arguments. If no templates are\n" +
			"  specified, the template is read from stdin.\n" +
			"\n" +
			"  `--init`, `-i`\n" +
			"\n" +
			"  Include simulated functions only available during `chezmoi init`.\n" +
			"\n" +
			"  '--output', '-o' *filename*\n" +
			"\n" +
			"  Write the output to *filename* instead of stdout.\n" +
			"\n" +
			"  `--promptString`, `-p` *pairs*\n" +
			"\n" +
			"  Simulate the `promptString` function with a function that returns values from\n" +
			"  *pairs*. *pairs* is a comma-separated list of *prompt*`=`*value* pairs. If\n" +
			"  `promptString` is called with a *prompt* that does not match any of *pairs*,\n" +
			"  then it returns *prompt* unchanged.\n" +
			"\n" +
			"  `execute-template` examples\n" +
			"\n" +
			"    chezmoi execute-template '{{ .chezmoi.sourceDir }}'\n" +
			"    chezmoi execute-template '{{ .chezmoi.os }}' / '{{ .chezmoi.arch }}'\n" +
			"    echo '{{ .chezmoi | toJson }}' | chezmoi execute-template\n" +
			"    chezmoi execute-template --init --promptString email=john@home.org <\n" +
			"  ~/.local/share/chezmoi/.chezmoi.toml.tmpl",
	},
	"forget": {
		long: "" +
			"Description:\n" +
			"  Remove *targets* from the source state, i.e. stop managing them.",
		example: "" +
			"  chezmoi forget ~/.bashrc",
	},
	"git": {
		long: "" +
			"Description:\n" +
			"  Run `git` *arguments* in the source directory. Note that flags in *arguments*\n" +
			"  must occur after `--` to prevent chezmoi from interpreting them.",
		example: "" +
			"  chezmoi git add .\n" +
			"  chezmoi git add dot_gitconfig\n" +
			"  chezmoi git -- commit -m \"Add .gitconfig\"",
	},
	"help": {
		long: "" +
			"Description:\n" +
			"  Print the help associated with *command*.",
	},
	"hg": {
		long: "" +
			"Description:\n" +
			"  Run `hg` *arguments* in the source directory. Note that flags in *arguments*\n" +
			"  must occur after `--` to prevent chezmoi from interpreting them.",
		example: "" +
			"  chezmoi hg -- pull --rebase --update",
	},
	"import": {
		long: "" +
			"Description:\n" +
			"  Import the source state from an archive file in to a directory in the source\n" +
			"  state. This is primarily used to make subdirectories of your home directory\n" +
			"  exactly match the contents of a downloaded archive. You will generally always\n" +
			"  want to set the `--destination`, `--exact`, and `--remove-destination` flags.\n" +
			"\n" +
			"  The only supported archive format is `.tar.gz`.\n" +
			"\n" +
			"  `--destination` *directory*\n" +
			"\n" +
			"  Set the destination (in the source state) where the archive will be imported.\n" +
			"\n" +
			"  `-x`, `--exact`\n" +
			"\n" +
			"  Set the `exact` attribute on all imported directories.\n" +
			"\n" +
			"  `-r`, `--remove-destination`\n" +
			"\n" +
			"  Remove destination (in the source state) before importing.\n" +
			"\n" +
			"  `--strip-components` *n*\n" +
			"\n" +
			"  Strip *n* leading components from paths.",
		example: "" +
			"  curl -s -L -o oh-my-zsh-master.tar.gz https://github.com/robbyrussell/oh-my-\n" +
			"zsh/archive/master.tar.gz\n" +
			"  chezmoi import --strip-components 1 --destination ~/.oh-my-zsh oh-my-zsh-master.tar.gz",
	},
	"init": {
		long: "" +
			"Description:\n" +
			"  Setup the source directory and update the destination directory to match the\n" +
			"  target state. If *repo* is given then it is checked out into the source\n" +
			"  directory, otherwise a new repository is initialized in the source directory.\n" +
			"  If a file called `.chezmoi.format.tmpl` exists, where `format` is one of the\n" +
			"  supported file formats (e.g. `json`, `toml`, or `yaml`) then a new\n" +
			"  configuration file is created using that file as a template. Finally, if the `--\n" +
			"  apply` flag is passed, `chezmoi apply` is run.",
		example: "" +
			"  chezmoi init https://github.com/user/dotfiles.git\n" +
			"  chezmoi init https://github.com/user/dotfiles.git --apply",
	},
	"manage": {
		long: "" +
			"Description:\n" +
			"  `manage` is an alias for `add` for symmetry with `unmanage`.",
	},
	"managed": {
		long: "" +
			"Description:\n" +
			"  List all managed entries in the destination directory in alphabetical order.\n" +
			"\n" +
			"  `-i`, `--include` *types*\n" +
			"\n" +
			"  Only list entries of type *types*. *types* is a comma-separated list of types\n" +
			"  of entry to include. Valid types are `dirs`, `files`, and `symlinks` which can\n" +
			"  be abbreviated to `d`, `f`, and `s` respectively. By default, `manage` will\n" +
			"  list entries of all types.",
		example: "" +
			"  chezmoi managed\n" +
			"  chezmoi managed --include=files\n" +
			"  chezmoi managed --include=files,symlinks\n" +
			"  chezmoi managed -i d\n" +
			"  chezmoi managed -i d,f",
	},
	"merge": {
		long: "" +
			"Description:\n" +
			"  Perform a three-way merge between the destination state, the source state, and\n" +
			"  the target state. The merge tool is defined by the `merge.command`\n" +
			"  configuration variable, and defaults to `vimdiff`. If multiple targets are\n" +
			"  specified the merge tool is invoked for each target. If the target state\n" +
			"  cannot be computed (for example if source is a template containing errors or\n" +
			"  an encrypted file that cannot be decrypted) a two-way merge is performed\n" +
			"  instead.",
		example: "" +
			"  chezmoi merge ~/.bashrc",
	},
	"purge": {
		long: "" +
			"Description:\n" +
			"  Remove chezmoi's configuration, state, and source directory, but leave the\n" +
			"  target state intact.\n" +
			"\n" +
			"  `-f`, `--force`\n" +
			"\n" +
			"  Remove without prompting.",
		example: "" +
			"  chezmoi purge\n" +
			"  chezmoi purge --force",
	},
	"remove": {
		long: "" +
			"Description:\n" +
			"  Remove *targets* from both the source state and the destination directory.\n" +
			"\n" +
			"  `-f`, `--force`\n" +
			"\n" +
			"  Remove without prompting.",
	},
	"rm": {
		long: "" +
			"Description:\n" +
			"  `rm` is an alias for `remove`.",
	},
	"secret": {
		long: "" +
			"Description:\n" +
			"  Run a secret manager's CLI, passing any extra arguments to the secret\n" +
			"  manager's CLI. This is primarily for verifying chezmoi's integration with your\n" +
			"  secret manager. Normally you would use template functions to retrieve secrets.\n" +
			"  Note that if you want to pass flags to the secret manager's CLU you will need\n" +
			"  to separate them with `--` to prevent chezmoi from interpreting them.\n" +
			"\n" +
			"  To get a full list of available commands run:\n" +
			"\n" +
			"    chezmoi secret help",
		example: "" +
			"  chezmoi secret bitwarden list items\n" +
			"  chezmoi secret keyring set --service service --user user\n" +
			"  chezmoi secret keyring get --service service --user user\n" +
			"  chezmoi secret lastpass ls\n" +
			"  chezmoi secret lastpass -- show --format=json id\n" +
			"  chezmoi secret onepassword list items\n" +
			"  chezmoi secret onepassword get item id\n" +
			"  chezmoi secret pass show id\n" +
			"  chezmoi secret vault -- kv get -format=json id",
	},
	"source": {
		long: "" +
			"Description:\n" +
			"  Execute the source version control system in the source directory with *args*.\n" +
			"  Note that any flags for the source version control system must be sepeated\n" +
			"  with a `--` to stop chezmoi from reading them.",
		example: "" +
			"  chezmoi source init\n" +
			"  chezmoi source add .\n" +
			"  chezmoi source commit -- -m \"Initial commit\"",
	},
	"source-path": {
		long: "" +
			"Description:\n" +
			"  Print the path to each target's source state. If no targets are specified then\n" +
			"  print the source directory.\n" +
			"\n" +
			"  `source-path` examples\n" +
			"\n" +
			"    chezmoi source-path\n" +
			"    chezmoi source-path ~/.bashrc",
	},
	"unmanage": {
		long: "" +
			"Description:\n" +
			"  `unmanage` is an alias for `forget` for symmetry with `manage`.",
	},
	"unmanaged": {
		long: "" +
			"Description:\n" +
			"  List all unmanaged files in the destination directory.",
		example: "" +
			"  chezmoi unmanaged",
	},
	"update": {
		long: "" +
			"Description:\n" +
			"  Pull changes from the source VCS and apply any changes.",
		example: "" +
			"  chezmoi update",
	},
	"upgrade": {
		long: "" +
			"Description:\n" +
			"  Upgrade chezmoi by downloading and installing the latest released version.\n" +
			"  This will call the GitHub API to determine if there is a new version of\n" +
			"  chezmoi available, and if so, download and attempt to install it in the same\n" +
			"  way as chezmoi was previously installed.\n" +
			"\n" +
			"  If chezmoi was installed with a package manager (`dpkg` or `rpm`) then\n" +
			"  `upgrade` will download a new package and install it, using `sudo` if it is\n" +
			"  installed. Otherwise, chezmoi will download the latest executable and replace\n" +
			"  the existing executable with the new version.\n" +
			"\n" +
			"  If the `CHEZMOI_GITHUB_API_TOKEN` environment variable is set, then its value\n" +
			"  will be used to authenticate requests to the GitHub API, otherwise\n" +
			"  unauthenticated requests are used which are subject to stricter rate limiting\n" +
			"  https://developer.github.com/v3/#rate-limiting. Unauthenticated requests should\n" +
			"  be sufficient for most cases.",
		example: "" +
			"  chezmoi upgrade",
	},
	"verify": {
		long: "" +
			"Description:\n" +
			"  Verify that all *targets* match their target state. chezmoi exits with code 0\n" +
			"  (success) if all targets match their target state, or 1 (failure) otherwise.\n" +
			"  If no targets are specified then all targets are checked.",
		example: "" +
			"  chezmoi verify\n" +
			"  chezmoi verify ~/.bashrc",
	},
}
