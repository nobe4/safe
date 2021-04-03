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

TODO: release

Currently you can `git pull` and run `make`.

# Contribute

Feel free to open issues/PR.

Improvement ideas in no particular order:
- [ ] Write tests
- [ ] Support dynamic programs (`vim`, `top`, ...)
- [ ] Add signal handler when terminal closes
- [ ] Support windows
- [ ] Improve logger formatting

# License

[MIT](./LICENSE)
