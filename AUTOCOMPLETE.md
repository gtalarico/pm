## Autocomplete

### Bash

```bash
# .bashrc
eval "$(_PM_COMPLETE=source pm)"
````

### zsh

```bash
# .zshrc
eval "$(_PM_COMPLETE=source_zsh pm)"
```

### fish

```bash
# .config/fish/completions/pm.fish
complete --command pm --arguments '(pm list)' --no-files
```

TODO: Update bbash and zsh with lazy as per [this](http://click.palletsprojects.com/en/7.x/bashcomplete/#activation-script)
