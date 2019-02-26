printf "%s\n" "Building Kamadhenu's temple"

if type brew > /dev/null 2>&1; then
    brew update
    brew bundle
    if [ ! -e /usr/local/bin/kamadhenu ] > /dev/null 2>&1 ; then
        printf "%s\n" "Symlink kamadhenu to $HOMEBREW_PREFIX/bin"
        ln -s "$PWD/kamadhenu" $HOMEBREW_PREFIX/bin/
    fi
    echo "Herding cowsays and wrangling figlets..."
    cp -n $PWD/cows/*.cow $HOMEBREW_PREFIX/share/cows/
    cp -n $PWD/figlet/fonts/*.flf $HOMEBREW_PREFIX/share/figlet/fonts
fi

(echo "You may now consult Kamadhenu\n"; echo "Just type kamadhenu") | cowsay -s
sleep 2
figlet -f basic "Kamadhenu" | lolcat -a
sleep 1
kamadhenu
