# Project Manager

Minimal CLI Project Manager

---

## Installing


```
$ brew install goreleaser
```

# Usage

```
$ pm list
```

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

## About

This project is almost identical to [github.com/Angelmmiguel/pm](https://github.com/Angelmmiguel/pm), which is much more mature.

You should absolutely use that instead.

## Release

This project uses [goreleaser](https://goreleaser.com/) to distribute its binaries.

To use goreleaser, you must export the variable `GITHUB_TOKEN` with a valid token
that can write to this repository.

## License

[MIT](https://opensource.org/licenses/MIT)
