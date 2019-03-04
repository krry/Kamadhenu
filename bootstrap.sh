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

sudo git clone https://github.com/krry/Kamadhenu.git || \
(printf "Cloning the git repository has gone awry." && exit 1)

cd "./Kamadhenu" || (printf "Oops, flubbed the dismount.\nFind \
'Kamadhenu' in %s and type 'brew bundle' to continue.\n" "$PWD" && exit 1)

printf "%s\n" "Building Kamadhenu's temple"

if type brew > /dev/null 2>&1; then
    brew update
    brew bundle
    if [ ! -e /usr/local/bin/Kamadhenu ] > /dev/null 2>&1 ; then
        printf "Symlinking Kamadhenu into %s/bin\n" "$BREWFIX"
        ln -s "$PWD/Kamadhenu" "$BREWFIX/bin/"
    fi
    printf "Herding cowsays..."
    printf "%s\n" "$PWD/cows/*.cow to $BREWFIX/share/cows/"
    cp -n "$PWD/cows/*.cow" "$BREWFIX/share/cows/"
    printf "Wrangling figlets..."
    printf "%s\n" "$PWD/figlet/fonts/*.flf to $BREWFIX/share/figlet/fonts"
    cp -n "$PWD/figlet/fonts/*.flf" "$BREWFIX/share/figlet/fonts"
    printf "Stuffing fortune cookies..."
    printf "%s\n" "$PWD/fortunes/* to $BREWFIX/share/games/fortunes"
    cp -n "$PWD/fortunes/*" "$BREWFIX/share/games/fortunes"
fi
printf "%s\n" "GREAT SUCCESS!"
printf "%s\n" "You may now call upon Kamadhenu, like so..." | cowsay -f fox | lolcat
sleep 3
figlet -f computer "Kamadhenu" | lolcat -a
sleep 3
Kamadhenu
