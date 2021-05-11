# KAMADHENU

An udderly absurd MOTD soothsayer for your *nix `$SHELL`

![Vaporware](./img/kama-ghost.png)
![TOC]

## Install

``` shell
curl -fsSL https://raw.githubusercontent.com/krry/Kamadhenu/main/bootstrap.sh | bash
```

Then follow the bouncing bytes...

Fresh out of the box,\
you get a bare minimum of\
**`275,254,848` spurts of insight**\
delivered one. at. a. time.

## Use

All you have to do is call her name.

``` shell
Kamadhenu
```

Who dares put words in the mouth of Kamadhenu?

``` shell
Kamadhenu Do you love me?
echo ' '
Kamadhenu Could you learn to love me?
```

Kamadhenu may not help, but you can still ask.

``` shell
Kamadhenu help
```

Pick a number between 1 and `Ctrl-C`

``` shell
Kamadhenu 2021
```

Start each shell sesh fresh with Kamadhenu

``` shell
echo Kamadhenu >> .$(basename $SHELL)rc
```

Oh, dear, is it Terminal?

### Coming soon (maybe)

Pipe `stdout` to Kamadhenu (like a cat)

``` shell
# echo "You ever drunk Bailey's out of a shoe?" | Kamadhenu
```

## Dependencies

- [fortune](https://github.com/bmc/fortunes/)
- [cowsay](https://linux.die.net/man/1/cowsay)
- [lolcat](https://github.com/busyloop/lolcat)
- [figlet](http://www.figlet.org/)
- and `coreutils`, specifically for GNU `shuf` which Macs don't come with.

## Troubleshooting

If you encounter a hiccup, hold your breath. Anything else, check here.

### No Homebrew, dog? That's otay

The dependencies are listed in the `Brewfile` and right nearby.
If you gather those, you're most of the way there. All that's left is to copy those artisanally curated, hand-fingered cows, fonts, and fortunes into their respective homes on your machine.

### Flubbed with the dismount? Failed to symlink?

The likely culprits are file permissions. Make sure you launch the bootstrap script from the safety of a locally owned, openly permissioned dir.

If nothing happens when you fire `Kamadhenu` into the command line, the symlink probably didn't make it into your `PATH`. Nobody knows your `PATH` like you do, so give it an `echo $PATH` and find your way.

#### Missing cows and/or figlets? Fortunes failing?

These need to be where `cowsay`, `figlet`, and `fortune` would look for them. Once again, you have a keen advantage in determining this.

First determine your `brew --prefix`.

On Macs with Intel chips we tend to use `/usr/local`, and for the Apple Silicon we tend to use `/opt/homebrew`.

- cows go in: `/usr/local/share/cows`
- figlets in: `/usr/local/share/figlet/fonts`
- fortunes => `/usr/local/share/games/fortunes`

## Many Gratitudes

If you especially enjoy Kamadhenu, fork away, or come and play!

I began this project because I didn't know jack about shell scripting. And I still don't!

Now that it's alive, it might be fun to glitter the flock out of it.

And who doesn't need yet another MOTD, amirite?

![Zuckle at the zipple of wisdom](./img/Kamadhenu.jpg)
