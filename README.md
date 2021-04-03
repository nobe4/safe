Safe
====

> Don't look at the credentials on the screen. - me, sometimes

Sharing one's terminal is a great way to explain, help, contribute.
It's possible to display secrets inadvertently, `cat`ing a configuration file.

This simple binary helps to prevent this from happening.

# Usage

## Shell Mode

Running `safe` in shell mode opens up a new shell which will filter out all output.

E.g.
```bash
$ safe
$ cat config.yml
node-dev:
  hostname: mynode-dev.local
  username: node-user
  password: XXXXXXXXXXXXXXXX
```

*Warning:* Dynamic program might not work as expected. For instance, `top`, `vim`, `nano` will fail to run properly. This is due to the underlying parsing of the input.

## Pipe Mode

Running `safe` in pipe mode reads from `stdin` and filter out the output to `stdout`.

E.g.
```bash
$ cat config.yml | safe
node-dev:
  hostname: mynode-dev.local
  username: node-user
  password: XXXXXXXXXXXXXXXX
```

# Installation

## Release

This is the prefered way.

Check the [lastest release](https://github.com/nobe4/safe/releases/latest).

*Note:* If you use the macOs binary, they are not currently signed. You need to `right-click` and `open` first ([ref](https://support.apple.com/guide/mac-help/open-a-mac-app-from-an-unidentified-developer-mh40616/mac)).

## Manual build

Currently you can `git pull` and run `make`.

## Go Get

```bash
$ go get github.com/nobe4/safe/tree/master/cmd/safe
```

*Note:* This method doesn't embed version information in the binary, prefer the [release](#release).

# Contribute

Feel free to open issues/PR.

Improvement ideas in no particular order:
- [ ] Write tests
- [ ] Support dynamic programs (`vim`, `top`, ...)
- [ ] Add signal handler when terminal closes
- [ ] Support windows
- [ ] Improve logger formatting
- [ ] Sign apple binary

# License

[MIT](./LICENSE)
