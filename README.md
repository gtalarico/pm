# PM - CLI Project Manager

[![CircleCI](https://circleci.com/gh/gtalarico/pm.svg?style=svg)](https://circleci.com/gh/gtalarico/pm)

[![Coverage Status](https://coveralls.io/repos/github/gtalarico/pm/badge.svg?branch=master)](https://coveralls.io/github/gtalarico/pm?branch=master)

Minimal CLI Project Manager

---

`pm` is a simple CLI utility to help you keep frequently used project paths at your finger tips.

![Demo Gif](demo.gif)

## Usage

Add a project
```
~ $ pm add myproject ~/code/repos/myproject
```

Go to project
```
~ $ pm go myproject
Starting new shell ...
~/code/repos/myproject $
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

## Installing

MacOs (brew)
```
$ brew install gtalarico/homebrew-tap/pm
```

Using Go

```
$ go get github.com/gtalarico/pm
```

## Configuration File

Project settings are stored in `~/.pm.json`
The first time you use `pm` one will be automatically created.
You may edit this file manually on your own.

## Disclaimer

This project is _very_ similar to [github.com/Angelmmiguel/pm](https://github.com/Angelmmiguel/pm),
which appears to be mature. You should probably use that instead.

This project was created as Go learning project. 

![](https://tutorialedge.net/images/golang.png)

## License

[MIT](https://opensource.org/licenses/MIT)

