#!/usr/bin/env fish

# Checking for Homebrew
echo ''
if which brew
    echo "Homebrew found. Commencing installation..."
    echo "Cloning the Kamadhenu repository..."
else
    echo "Sorry, you'll need Homebrew for Kamadhenu to install itself."
    echo "See the REAMDE.md for manual installation instructions."
    exit 2
end
echo ''

# CONSTANTS
set BREWFIX "(brew --prefix)"
set COW_DIR $BREWFIX/share/cows/
set FIGLET_DIR $BREWFIX/share/figlet/fonts
set FORTUNE_DIR $BREWFIX/share/games/fortunes

# Cloning
if git clone --quiet https://github.com/krry/Kamadhenu.git && echo ''
    echo "Repository cloned into $PWD"
    if cd "./Kamadhenu"
        echo "Building Kamadhenu's temple..."
        echo "Updating brew and unbundling brewfile..."
    else
        echo "Oops, flubbed the dismount. Couldn't find the cloned files."
        echo "find 'Kamadhenu' in $PWD and type 'brew bundle' to continue."
        exit 1
    end
else
    echo ''
    echo "Cloning the git repository has gone awry."
    exit 1
end
echo ''

# Brewing
if brew update && brew bundle && echo ''
    echo "Brewed and ready."
else
    echo "There was an issue with the brew."
end
echo ''

# Symlinking
if [ -L $BREWFIX/bin/Kamadhenu ]
    rm $BREWFIX/bin/Kamadhenu
end
echo "Symlinking Kamadhenu into $BREWFIX/bin"
echo ''
if ln -s "$PWD/Kamadhenu" "$BREWFIX/bin/"
    echo "Kamadhenu symlinked"
else
    echo "Failed to symlink. You'll want to see to that."
end
echo ''

# Copying Data
echo "Herding cowsays" "(oo) "
if cp -n "$PWD/cows/"*.cow "$COW_DIR"
    echo "Cowsays herded into $COW_DIR"
else
    echo "Couldn't copy cows into $COW_DIR"
end
echo ''
echo "Wrangling figlets" "<..> "
if cp -n "$PWD/figlet/fonts/"*.flf "$FIGLET_DIR"
    echo "figlets wrangled into $FIGLET_DIR"
else
    echo "Couldn't copy Figlet fonts into $FIGLET_DIR"
end
echo ''
echo "Stuffing fortune cookies" "\$\$ "
if cp -n "$PWD/fortunes/"* "$FORTUNE_DIR"
    echo "Fortunes stuffed into $FORTUNE_DIR"
else
    echo "Couldn't copy Fortunes into $FORTUNE_DIR"
end
echo ''
if [ $WARNED ]
    echo "Check the warnings above to fix the flubs."
else
    echo "Installation complete."
    echo ''
    echo "GREAT SUCCESS!"
    sleep 2
    echo ''
    echo ''
    echo "You may now call upon Kamadhenu, like so..." | cowsay -f fox | lolcat
    echo ''
    echo ''
    sleep 3
    echo ''
    echo ''
    figlet -f computer "Kamadhenu" | lolcat -a
    echo ''
    echo ''
    sleep 3
    Kamadhenu
end
