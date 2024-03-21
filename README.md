# password

[![Go Report Card](https://goreportcard.com/badge/github.com/jimrazmus/password)](https://goreportcard.com/report/github.com/jimrazmus/password)
[![Go Documentation](https://pkg.go.dev/github.com/jimrazmus/password)]
![Build](https://github.com/jimrazmus/password/workflows/build/badge.svg?branch=master)
[![CodeCov](https://codecov.io/gh/jimrazmus/password/branch/master/graph/badge.svg)](https://codecov.io/gh/jimrazmus/password)
[![License](http://img.shields.io/:license-mit-blue.svg?style=flat-square)](http://badges.mit-license.org)

## Summary

Password hashing and checking using bcrypt.

## Installation

`go get github.com/jimrazmus/password`

## Usage

```golang
import github.com/jimrazmus/password

bcrypted, err := HashPassword("your_password")
if err != nil {
    log.Fatal(err)
}

// store 'bcrypted' somewhere for later use

// ... later ...

if !CheckPassword("your_password", bcrypted) {
    log.Fatal("Incorrect password")
}
```

## Contributing

Please follow the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) specification for your commit messages. Commit type options include: feat, fix, build, chore, ci, docs, style, refactor, perf, and test.

Open a pull request when you have completed your work and want it reviewed for inclusion.

## Author

Jim Razmus II

## License

This project is licensed under the MIT License - see the [LICENSE.txt](LICENSE.txt) file for details.
