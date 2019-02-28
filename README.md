
# KAMADHENU

An udderly absurd soothsayer for your shell.

## Installation

``` sh
git clone https://github.com/krry/kamadhenu.git
cd kamadhenu && . ./install.sh
```

### Bless every terminal
Begin each shell sesh with Kamadhenu.
```echo kamadhenu >> .$(basename $SHELL)rc```

## Dependencies
- [fortune](https://github.com/bmc/fortunes/)
- [cowsay](https://linux.die.net/man/1/cowsay)
- [lolcat](https://github.com/busyloop/lolcat)
- [cowsay-files](https://github.com/krry/cowsay-files)

I believe each `cowsay`, `fortune`, `figlet`, and `lolcat` are available for install via `brew`, `pip`, `snap`, and `apt-get`, and various other *nix pacmen. You can do eet.


## TODO
- [ ] offer settings mode, a settings.file, and CLI flags
  - cow face settings
  - choose default cow
  - choose default font
  - monochrome mode - no lolcat
  - no cowsay mode
  - no fortune mode
- [ ] gracefully degrade if dependencies not present
- [x] add cowsay and figlet fonts to the repo (ideally as submodules)
  - [x] categorize figlet fonts into non-readable, favored, others
  - [x] categorize cowsays into mono, color, and large
- [ ] moar easter eggs - make yourself at home
- add more serendipitous transformations on user input
- interactive mode
  - editable cowsay bubble
  - allow live typing in figlet?
- programmatically generate the temple
- expand coloration beyond lolcat
- expand compatibility to various Linux distros
