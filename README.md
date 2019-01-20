# Project Manager

Minimal CLI Project Manager

---

## Installing


```
$ brew tap-pin gtalarico/pm
```

# Usage

Add a project
```
$ pm add myproject ~/code/repos/myproject
```

Go to project
```
$ pm go myproject
Starting new shell...
```

List projects
```
$ pm list
myproject
```

Remove project
```
$ pm remove myproject
```

## Configuration File

Project settings is stored in ~/.go-projects.
The first time you use `pm` one will be automatically created.
You may edit this file on your own as needed.

## About

This project is almost identical to [github.com/Angelmmiguel/pm](https://github.com/Angelmmiguel/pm), which is much more mature.

You should absolutely use that instead.

## Release

This project uses [goreleaser](https://goreleaser.com/) to distribute its binaries.

To use goreleaser, you must export the variable `GITHUB_TOKEN` with a valid token
that can write to this repository.

## License

[MIT](https://opensource.org/licenses/MIT)


# Todo

- [ ] Add homebrew tab
- [ ] Add tests + circleci
- [ ] Add Action cmd
- [ ] Resolve fish, bash, zsh autocomplete
- [ ] Refactor
