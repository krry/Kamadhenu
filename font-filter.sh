#!/bin/sh

warn() {
  printf "\r\033[2K  [\033[0;31mFAIL\033[0m] $1\n"
}

# grab list of fonts I want
fonts=("${(@f)$(cat "./fonts.txt")}")

# find the font files
FONTDIR='/usr/local/share/figlet/fonts'

# set the dest dir
DESTDIR='./figlets/'
TOTAL=0
COUNT=0
# iterate through the list
for font in $fonts; do
  let TOTAL+=1
  if [ -e "$FONTDIR/$font.flf" ]; then
    # copying matched files to the destdir
    cp -n $FONTDIR/$font.flf $DESTDIR
    let COUNT+=1
  else
    # or alert that it was missed
    warn "$font.flf not found"
  fi
done

echo "$COUNT of $TOTAL fonts copied"

