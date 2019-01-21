## Autocomplete

### Bash

```bash
#
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
function __fish_pm_projects
    command pm list 2>/dev/null
end

complete --command "pm" -n '__fish_seen_subcommand_from list' --no-files
complete --command "pm" -n '__fish_seen_subcommand_from go remove' --arguments '(__fish_pm_projects)' --no-files
```

TODO: Update bbash and zsh with lazy as per [this](http://click.palletsprojects.com/en/7.x/bashcomplete/#activation-script)
