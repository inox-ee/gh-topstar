# inox-ee/gh-topstar

## How to install

```shell-session
# use go get if your go version is less than 1.16
$ go install github.com/inox-ee/gh-topstar
```

## How to uninstall

```shell-session
$ rm -rf $GOPATH/pkg/mod/inox-ee
$ rm $GOPATH/bin/gh-topstar
```

or please tell me uninstall command

## How to Use

Ex)

```shell-session
$ gh-topstar
🔍 > go migration
Search:
         https://github.com/search?o=desc&q=go+migration&s=stars&type=Repositories
Result:
          1 : 7.3k  https://github.com/golang-migrate/migrate
          2 : 2.3k  https://github.com/mattes/migrate
          3 : 2.3k  https://github.com/rubenv/sql-migrate
          4 : 1.9k  https://github.com/pressly/goose
          5 : 678   https://github.com/go-gormigrate/gormigrate
          6 : 581   https://github.com/scylladb/gocqlx
          7 : 483   https://github.com/rsc/c2go
          8 : 391   https://github.com/minamijoyo/tfmigrate
          9 : 352   https://github.com/alphagov/govuk-aws
         10 : 271   https://github.com/go-pg/migrations
```

## 参考

- Go CLI
  - [urfave/cli: A simple, fast, and fun package for building command line apps in Go](https://github.com/urfave/cli)
- Release
  - [Go で書いた CLI ツールのリリースは GoReleaser と GitHub Actions で個人的には決まり | tellme.tokyo](https://tellme.tokyo/post/2020/02/04/release-go-cli-tool/)
