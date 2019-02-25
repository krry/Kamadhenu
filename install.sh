printf "%s\n" "Building Kamadhenu's temple"

if type brew > /dev/null 2>&1; then
    brew update
    brew bundle
    if [ ! -e /usr/local/bin/kamadhenu ] > /dev/null 2>&1 ; then
        printf "%s\n" "Symlinking in $HOMEBREW_PREFIX/bin"
        ln -s "$PWD/kamadhenu" $HOMEBREW_PREFIX/bin/
    fi
    printf "%s\n" "You may now consult Kamadhenu"
                \ "Just type kamadhenu"
                \ "kamadhenu"
    # KMDIR=$HOMEBREW_PREFIX/share/kamadhenu
    # echo "Herding cowsays and wrangling figlets..."
    # cp $PWD/cows $KMDIR
    # cp $PWD/figlets $KMDIR
# elif type snap; then
#     snap install figlet cowsay lolcat fortune
fi
