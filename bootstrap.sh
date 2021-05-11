#!/usr/bin/env bash

# CONSTANTS
BREWFIX="$(brew --prefix)"
COW_DIR=${BREWFIX}/share/cows/
FIGLET_DIR=${BREWFIX}/share/figlet/fonts
FORTUNE_DIR=${BREWFIX}/share/games/fortunes
KM_DIR=${PWD}/Kamadhenu

# UTILS
hr_msg ()  {
    echo ''
    printf -v _hr "%*s" "$(tput cols)" ""
    echo -en "${_hr// /${2--}}"
    echo -e "\r\033[$((($(tput cols)-${#1}-2)/2))C ${1-hi} "
}

message () {
    echo ''
    printf "\r  [ \033[00;34m(:\033[0m ] %s\n" "$1"
}

success () {
    echo ''
    printf "\r\033[2K  [ \033[00;32mOK\033[0m ] %s\n" "$1"
}

warning () {
    echo ''
    printf "\r  [ \033[00;33mUH\033[0m ] %s\n" "$1"
    WARNED=true
}

failure () {
    echo ''
    printf "\r\033[2K  [\033[0;31mNO\033[0m] %s\n" "$1"
    exit
}

# RUNTIME
check_for_brew () {
    # Checking for Homebrew
    if type brew > /dev/null 2>&1; then
        message "Homebrew found. Commencing installation..."
        message "Cloning the Kamadhenu repository..."
        clone_kama
    else
        failure "Sorry, you'll need Homebrew for Kamadhenu to install itself."
        hr_msg "See the REAMDE.md for manual installation instructions."
        exit 2
    fi
}

clone_kama () {
    # Cloning
    git clone --quiet https://github.com/krry/Kamadhenu.git
    git_clone_status=$?

    if (($git_clone_status < 2)); then
        success "Repository cloned into $PWD"
        cd "./Kamadhenu" && cd_status=$?
        if ((cd_status < 2)); then
            message "Updating brew and unbundling Brewfile..."
            echo ''
            brew_up
        else
            failure "Oops, flubbed the dismount. Couldn't find the cloned files."
            failure "Find 'Kamadhenu' in $PWD and type 'brew bundle' to continue."
            exit 1
        fi
    else
        failure "Cloning the git repository has gone awry."
        exit 1
    fi
}

brew_up () {
    brew update && brew bundle
    brew_status=$?

    if (($brew_status < 2)); then
        success "Brewed and ready."
        build_temple
    else
        failure "There was an issue with the brew."
    fi
}

build_temple () {
    message "Adding Kamadhenu into $BREWFIX/bin"

    cp -f "$PWD/Kamadhenu" "$BREWFIX/bin"
    symlink_status=$?

    if (($symlink_status < 2)); then
        success "Kamadhenu now lives in your homebrew bin"
    else
        warning "Failed to symlink. Kamadhenu wants to live in $BREWFIX/bin."
    fi
    herd_cowsays
}

herd_cowsays () {
    hr_msg "Herding cowsays" "(oo) "

    cp -n "${PWD}/cows/"*.cow "$COW_DIR"
    cowsay_status=$?

    if (($cowsay_status < 2)); then
        success "Cowsays herded into $COW_DIR"
    else
        warning "Couldn't copy all the cows into $COW_DIR"
    fi
    wrangle_figlets
}

wrangle_figlets () {
    hr_msg "Wrangling figlets" "<..> "

    cp -n "${PWD}/figlet/fonts/"*.flf "$FIGLET_DIR"
    figlet_status=$?

    if (($figlet_status < 2)); then
        success "Figlets wrangled into $FIGLET_DIR"
    else
        warning "Couldn't copy all the FIGlet fonts into $FIGLET_DIR"
    fi
    stuff_fortunes
}

stuff_fortunes () {
    hr_msg "Stuffing fortune cookies" "\$\$ "

    cp -n "${PWD}/fortunes/"* "$FORTUNE_DIR"
    fortune_status=$?

    if (($fortune_status < 2)); then
        success "Fortunes stuffed into $FORTUNE_DIR"
    else
        warning "Couldn't copy all the Fortunes into $FORTUNE_DIR"
    fi
    finish_install
}

remove_tmp_dir () {
    rm -rf $KM_DIR
}

finish_install () {
    if [ $WARNED ]; then
        warning "Check the warnings above to fix the flubs."
    else
        success "Cleaning up install files"
        success "Installation complete."
        hr_msg "GREAT SUCCESS!" "$"
        sleep 2
        echo ''
        printf "You may now call upon Kamadhenu by name." | cowsay -f fox | lolcat
        echo ''
        sleep 3
        echo ''
        figlet -f jazmine "Kamadhenu" | lolcat -a
        echo ''
        sleep 3
        Kamadhenu
        remove_tmp_dir
    fi
}

check_for_brew
