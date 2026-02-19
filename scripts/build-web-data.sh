#!/bin/bash
# Build JSON data from fortunes and cows for web app

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
WEB_DIR="$PROJECT_ROOT/web/public"

mkdir -p "$WEB_DIR"

echo "🐄 Building Kamadhenu web data..."

# Parse fortunes into JSON array
echo "[" > "$WEB_DIR/fortunes.json"
first=true
for fortune_file in "$PROJECT_ROOT/fortunes"/*.dat; do
  while IFS= read -r line || [ -n "$line" ]; do
    if [ "$line" = "%" ]; then
      continue
    fi
    if [ -n "$fortune" ]; then
      if [ "$first" = true ]; then
        first=false
      else
        echo "," >> "$WEB_DIR/fortunes.json"
      fi
      # Escape quotes and newlines for JSON
      escaped=$(echo "$fortune" | sed 's/"/\\"/g' | sed ':a;N;$!ba;s/\n/\\n/g')
      echo "  \"$escaped\"" >> "$WEB_DIR/fortunes.json"
      fortune=""
    fi
    fortune="${fortune}${line}"$'\n'
  done < "$fortune_file"
done
echo "]" >> "$WEB_DIR/fortunes.json"

echo "✅ Fortunes parsed: $(wc -l < "$WEB_DIR/fortunes.json") lines"

# Parse cows into JSON
echo "{" > "$WEB_DIR/cows.json"
first=true
for cow_file in "$PROJECT_ROOT/cows"/*.cow; do
  cow_name=$(basename "$cow_file" .cow)
  
  if [ "$first" = true ]; then
    first=false
  else
    echo "," >> "$WEB_DIR/cows.json"
  fi
  
  # Extract the cow art (everything between <<EOC and EOC)
  cow_art=$(awk '/\$the_cow = <<EOC;/,/^EOC/ {if (!/\$the_cow = <<EOC;/ && !/^EOC/) print}' "$cow_file")
  
  # Escape for JSON
  escaped=$(echo "$cow_art" | sed 's/"/\\"/g' | sed 's/\\/\\\\/g' | sed ':a;N;$!ba;s/\n/\\n/g')
  
  echo "  \"$cow_name\": \"$escaped\"" >> "$WEB_DIR/cows.json"
done
echo "}" >> "$WEB_DIR/cows.json"

echo "✅ Cows parsed: $(ls "$PROJECT_ROOT/cows"/*.cow | wc -l) cows"
echo "🎉 Build complete! Data in $WEB_DIR"
