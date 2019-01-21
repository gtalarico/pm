# Project Manager

[![CircleCI](https://circleci.com/gh/gtalarico/pm.svg?style=svg)](https://circleci.com/gh/gtalarico/pm)

Minimal CLI Project Manager

---

## Installing

MacOs (brew)
```
$ brew install gtalarico/homebrew-tap/pm
```

Using Go

```
$ go get github.com/gtalarico/pm
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

Project settings is stored in ~/.pm.json
The first time you use `pm` one will be automatically created.
You may edit this file on your own as needed.

## Disclaimer

This project is _very_ similar to [github.com/Angelmmiguel/pm](https://github.com/Angelmmiguel/pm),
which is much more mature. You should absolutely use that instead.

I built this project primarily to learn go.

## License

[MIT](https://opensource.org/licenses/MIT)

