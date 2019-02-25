printf "%s\n" "Building Kamadhenu's temple"

if type brew; then
    brew update
    brew bundle
    if [ ! -e /usr/local/bin/kamadhenu ] ; then
        printf "%s\n" "Making `kamadhenu` accessible from $PATH"
        printf "%s\n" "Symlinking in $HOMEBREW_PREFIX/bin"
        ln -s "$PWD/kamadhenu" $HOMEBREW_PREFIX/bin/
    fi
    KMDIR=$HOMEBREW_PREFIX/share/kamadhenu
    echo "Herding cowsays and wrangling figlets..."
    cp $PWD/cows $KMDIR
    cp $PWD/figlets $KMDIR
elif type snap; then
    snap install figlet cowsay lolcat fortune
fi
