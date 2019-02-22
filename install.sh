# installing dependencies
echo "installing Kamadhenu dependencies"

if type brew; then
  brew update && brew install figlet cowsay lolcat fortune coreutils
elif type snap; then
  snap install figlet cowsay lolcat fortune coreutils
fi

# find cows and set COWPATH
if [[ "$(basename $PWD)" == "kamadhenu" ]]; then
  export COWPATH=$PWD/cowsay-files/cows/
  if [[ "$SHELL" == *"zsh"* ]]; then
    echo COWPATH=$PWD/cowsay-files/cows/ >> ~/.zshrc
  elif [[ "$SHELL" == *"bash"* ]]; then
    echo COWPATH=$PWD/cowsay-files/cows/ >> ~/.bashrc
  fi
fi

