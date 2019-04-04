#!/usr/bin/env bash

# Toolset
hr_msg ()  {
    printf -v _hr "%*s" "$(tput cols)" ""
    echo -en "${_hr// /${2--}}"
    echo -e "\r\033[$((($(tput cols)-${#1}-2)/2))C ${1-hi} "
}

message () {
  printf "\r  [ \033[00;34m(:\033[0m ] %s\n" "$1"
}

success () {
    printf "\r\033[2K  [ \033[00;32mOK\033[0m ] %s\n" "$1"
}

warning () {
  printf "\r  [ \033[00;33mUH\033[0m ] %s\n" "$1"
  WARNED=true
}

failure () {
  printf "\r\033[2K  [\033[0;31mNO\033[0m] %s\n" "$1"
  exit
}

# Checking for Homebrew
echo ''
if type brew > /dev/null 2>&1; then
    message "Homebrew found. Commencing installation..."
    message "Cloning the Kamadhenu repository..."
else
    printf "Sorry, you'll need Homebrew for Kamadhenu to install itself."
    printf "See the REAMDE.md for manual installation instructions."
    exit 2


fi
echo ''

# CONSTANTS
BREWFIX="$(brew --prefix)"
COW_DIR=${BREWFIX}/share/cows/
FIGLET_DIR=${BREWFIX}/share/figlet/fonts
FORTUNE_DIR=${BREWFIX}/share/games/fortunes

# Cloning
if git clone --quiet https://github.com/krry/Kamadhenu.git && echo ''; then
    success "Repository cloned into $PWD"
    if cd "./Kamadhenu" ; then
        message "Building Kamadhenu's temple..."
        message "Updating brew and unbundling brewfile..."
    else
        failure "Oops, flubbed the dismount. Couldn't find the cloned files."
        failure "Find 'Kamadhenu' in $PWD and type 'brew bundle' to continue."
        exit 1
    fi
else
    echo ''
    failure "Cloning the git repository has gone awry."
    exit 1
fi
echo ''

# Brewing
if brew update && brew bundle && echo ''; then
    success "Brewed and ready."
else
    warning "There was an issue with the brew."
fi
echo ''

# Symlinking
if [ -L /usr/local/bin/Kamadhenu ] > /dev/null 2>&1 ; then
    rm /usr/local/bin/Kamadhenu;
fi
message "Symlinking Kamadhenu into $BREWFIX/bin"
echo ''
if ln -s "$PWD/Kamadhenu" "$BREWFIX/bin/" ; then
    success "Kamadhenu symlinked"
else
    warning "Failed to symlink. You'll want to see to that."
fi
echo ''

# Copying Data
hr_msg "Herding cowsays" "(oo) "
if cp -n "${PWD}/cows/"*.cow "$COW_DIR" ; then
    success "Cowsays herded into $COW_DIR"
else
    warning "Couldn't copy cows into $COW_DIR"
fi
echo ''
hr_msg "Wrangling figlets" "<..> "
if cp -n "${PWD}/figlet/fonts/"*.flf "$FIGLET_DIR" ; then
    success "Figlets wrangled into $FIGLET_DIR"
else
    warning "Couldn't copy FIGlet fonts into $FIGLET_DIR"
fi
echo ''
hr_msg "Stuffing fortune cookies" "\$\$ "
if cp -n "${PWD}/fortunes/"* "$FORTUNE_DIR" ; then
    success "Fortunes stuffed into $FORTUNE_DIR"
else
    warning "Couldn't copy Fortunes into $FORTUNE_DIR"
fi
echo ''
if [ $WARNED ]; then
    warning "Check the warnings above to fix the flubs."
else
    success "Installation complete."
    echo ''
    hr_msg "GREAT SUCCESS!" "$"
    sleep 2
    echo ''
    echo ''
    printf "You may now call upon Kamadhenu, like so..." | cowsay -f fox | lolcat
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
fi
