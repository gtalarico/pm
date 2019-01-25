## Autocomplete [WIP]

### fish

```fish
# .config/fish/completions/pm.fish
function __fish_pm_projects
    command pm list 2>/dev/null
end

complete --command "pm" -n '__fish_seen_subcommand_from list' --no-files
complete --command "pm" -n '__fish_seen_subcommand_from go remove' --arguments '(__fish_pm_projects)' --no-files

function __fish_pm_needs_command
    set cmd (commandline -opc)
    if [ (count $cmd) -eq 1 ]
        return 0
    end
    return 1
end

function __fish_pm_using_command
    set cmd (commandline -opc)
    if [ (count $cmd) -gt 1 ]
        if [ $argv[1] = $cmd[2] ]
            return 0
        end
    end
    return 1
end

set -l pm_looking -c pm -n '__fish_pm_needs_command'

# Main commands
complete $pm_looking -xa list -d 'list projects'
complete $pm_looking -xa go -d 'go to project'
complete $pm_looking -xa remove -d 'remove project'
complete $pm_looking -xa add -d 'add projects'
```

### TODO
- [ ] Remove explicit command descriptions from fish? (hidden docs: similar to click _complete)
- [ ] Add bash
- [ ] Add zsh
