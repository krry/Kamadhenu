#!/usr/bin/env bash
printf "Testing for homebrew...\n"
if type brew > /dev/null 2>&1; then
    printf "Homebrew found. Commencing installation...\n"
    BREWFIX="$(brew --prefix)"
    printf "You may have to sudo now.\n"
else
    printf "Sorry, you'll need Homebrew for Kamadhenu to install itself.\n"
    printf "See the REAMDE.md for manual installation instructions.\n"
    exit 2
fi

sudo git clone https://github.com/krry/kamadhenu/ "$BREWFIX/kamadhenu" || \
printf "Cloning the git repository has gone awry." && exit 1

cd "$BREWFIX/kamadhenu" || printf "Oops, flubbed the dismount.\nFind \
'kamadhenu' in %s and type 'brew bundle' to continue.\n" "$BREWFIX" && exit 1

printf "%s\n" "Building Kamadhenu's temple"

if type brew > /dev/null 2>&1; then
    brew update
    brew bundle
    if [ ! -e /usr/local/bin/kamadhenu ] > /dev/null 2>&1 ; then
        printf "Symlinking kamadhenu into %s/bin" "$BREWFIX"
        ln -s "$PWD/kamadhenu" "$BREWFIX/bin/"
    fi
    printf "Herding cowsays..."
    cp -n "$PWD/cows/*.cow" "$BREWFIX/share/cows/"
    printf "Wrangling figlets..."
    cp -n "$PWD/figlet/fonts/*.flf" "$BREWFIX/share/figlet/fonts"
    printf "Stuffing fortune cookies..."
    cp -n "$PWD/fortunes/*" "$BREWFIX/share/games/fortunes"
fi
printf "%s\n" "GREAT SUCCESS!"
printf "%s\n" "You may now call upon Kamadhenu, like so..." | cowsay -f fox | lolcat
sleep 2
figlet -f computer "Kamadhenu" | lolcat -a
sleep 1
kamadhenu
