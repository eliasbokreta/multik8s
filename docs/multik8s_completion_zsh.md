## multik8s completion zsh

Generate the autocompletion script for zsh

### Synopsis

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(multik8s completion zsh); compdef _multik8s multik8s

To load completions for every new session, execute once:

#### Linux:

	multik8s completion zsh > "${fpath[1]}/_multik8s"

#### macOS:

	multik8s completion zsh > $(brew --prefix)/share/zsh/site-functions/_multik8s

You will need to start a new shell for this setup to take effect.


```
multik8s completion zsh [flags]
```

### Options

```
  -h, --help              help for zsh
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
  -n, --namespace string   Kubernetes namespace (should be the exact name) (default "default")
  -p, --podname string     Kubernetes pod name (works as a wildcard)
```

### SEE ALSO

* [multik8s completion](multik8s_completion.md)	 - Generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 11-Jul-2022
