# Kamadhenu 🐄🌈

> An udderly absurd fortune teller in your shell

Kamadhenu is a delightful mashup of `fortune` and `cowsay` with rainbow gradients, packaged as a **single standalone binary** with zero dependencies.

## Features

- 🎲 **13,443 fortunes** embedded (or use your own!)
- 🐄 **109 ASCII cows** embedded  
- 🌈 **Rainbow gradients** with subtle randomness (organic feel)
- 🦄 **Unicorn mode** - chaotic rainbow colors (wild combos)
- ⌨️  **Typewriter effect** (optional)
- 🎮 **Interactive mode** or one-shot usage
- 📦 **Zero dependencies** - single binary, works out of the box
- 📝 **Custom sources** - add your own fortune files or URLs
- 🎁 **Easter egg** - `kamadhenu alias` installs `kama` shortcut

## Installation

### One-Liner (Easiest)

```bash
curl -fsSL https://raw.githubusercontent.com/krry/Kamadhenu/main/install.sh | bash
```

This auto-detects your OS/architecture and installs the binary + man page.

### Manual Download

```bash
# macOS (Apple Silicon)
curl -L https://github.com/krry/Kamadhenu/releases/latest/download/kamadhenu-darwin-arm64 -o kamadhenu
curl -L https://github.com/krry/Kamadhenu/releases/latest/download/kamadhenu.1 -o kamadhenu.1
chmod +x kamadhenu
sudo install -m 755 kamadhenu /usr/local/bin/
sudo install -m 644 kamadhenu.1 /usr/local/share/man/man1/

# macOS (Intel)
curl -L https://github.com/krry/Kamadhenu/releases/latest/download/kamadhenu-darwin-amd64 -o kamadhenu
curl -L https://github.com/krry/Kamadhenu/releases/latest/download/kamadhenu.1 -o kamadhenu.1
chmod +x kamadhenu
sudo install -m 755 kamadhenu /usr/local/bin/
sudo install -m 644 kamadhenu.1 /usr/local/share/man/man1/

# Linux
curl -L https://github.com/krry/Kamadhenu/releases/latest/download/kamadhenu-linux-amd64 -o kamadhenu
curl -L https://github.com/krry/Kamadhenu/releases/latest/download/kamadhenu.1 -o kamadhenu.1
chmod +x kamadhenu
sudo install -m 755 kamadhenu /usr/local/bin/
sudo install -m 644 kamadhenu.1 /usr/local/share/man/man1/
```

### Short Alias (Optional)

Add to your `~/.bashrc`, `~/.zshrc`, or `~/.config/fish/config.fish`:

```bash
# bash/zsh
alias kama='kamadhenu'

# fish
alias kama kamadhenu
```

Then use: `kama -h`, `kama "hello"`, etc.

### Homebrew (Coming Soon)

```bash
brew tap krry/kamadhenu
brew install kamadhenu
```

### Build from Source

```bash
git clone https://github.com/krry/Kamadhenu.git
cd Kamadhenu
make build
sudo make install  # Installs binary + man page
```

## Usage

```bash
# One random fortune (default)
kamadhenu

# Interactive mode
kamadhenu summon

# Interactive mode, monochrome (flags before subcommands!)
kamadhenu --mono summon

# Quick fortune (no animation)
kamadhenu --fast

# Custom text
kamadhenu "Hello, world"

# Show N fortunes in sequence
kamadhenu 5

# Unicorn mode (chaotic rainbow colors)
kamadhenu --unicorn

# Monochrome output
kamadhenu --mono

# Fast + plain output
kamadhenu --fast --mono
```

## Options

```
-h            Show help
-v            Show version
--mono        Disable rainbow colors (monochrome)
--no-color    Alias for --mono
--fast        Skip typewriter effect
--unicorn     Chaotic rainbow mode (wild color combos)
```

## Rainbow Modes

**Normal (default):**
- Random starting hue each run (rotates the rainbow)
- Smooth gradient across the classic spectrum
- Classic red→orange→yellow→green→blue→purple

**Unicorn mode (`--unicorn`):**
- Random RGB phase shifts each run
- Creates unexpected color palettes (lime/purple, cyan/yellow, etc.)
- Brightness boosted for readability on dark terminals
- Every run is a unique surprise

Try it:
```bash
kamadhenu --unicorn "Chaos reigns!"
```

## Custom Fortune Sources

Add your own quotes, wisdom, or humor:

```bash
# List current sources
kamadhenu sources list

# Add a local file
kamadhenu sources add ~/my-quotes.txt

# Add a URL (cached for 24h)
kamadhenu sources add https://example.com/quotes.json

# Remove a source
kamadhenu sources rm ~/my-quotes.txt

# Reset to embedded fortunes
kamadhenu sources reset
```

**Supported formats:**
- **Fortune format:** Text separated by `\n%\n`
- **JSON:** `[{"quote": "...", "author": "..."}]` or single object
- **Plain text:** One quote per line (# for comments)
- **Paragraphs:** Quotes separated by double newlines

**Examples:**

Fortune format (`my-quotes.txt`):
```
The best time to plant a tree was 20 years ago.
The second best time is now.
%
Be yourself; everyone else is already taken.
		— Oscar Wilde
```

JSON format (from URL or file):
```json
[
  {"quote": "Life is what happens...", "author": "John Lennon"},
  {"text": "The only way out is through."}
]
```

Plain text:
```
Every line is a separate fortune
# Comments start with #
This is another fortune
```

## Easter Egg: Shell Alias

Install the `kama` alias for quick access:

```bash
kamadhenu alias
# Adds 'alias kama=kamadhenu' to your shell config

# Then use it:
kama
kama summon
kama --unicorn "wheee"
```

## Examples

### Interactive Mode
```bash
$ kamadhenu
                    {}}
                   {{}{}
                  {{}}{}}
                 &>&&&&?&@
               &&@&&&@*&@#&#@
            &&&&#&*&&&&@&&&#&&&@
         &#&&@&&$&&&&&&&&&&&@&#&*&@
      &&*$&&$&7&&&$&>&&&&?&&&<&#&$&&#@
   &&&&&&$&&&&@&!&&&&&&&&&&&&#&*&#&&@>*&&
 &&!*&#+&&&&&!&<@&?&&&&?&&@&&/&#@&&&#&<&@&&&
&*  Prostrate yourself before the Almighty  &@
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

          Now, let's cast your lots.
```

### Custom Text
```bash
$ kamadhenu "Life is absurd"
 _______________
< Life is absurd >
 ---------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```

## Development

### Project Structure

```
kamadhenu/
├── main.go                  # CLI entry point
├── internal/
│   ├── fortune/             # Fortune loading & selection
│   ├── cowsay/              # Cow rendering & word wrap
│   ├── colors/              # Rainbow gradient (lolcat replacement)
│   └── typewriter/          # Line-by-line output effect
├── cows/                    # 109 .cow files (embedded)
├── fortunes/                # Fortune files (embedded)
└── README.md
```

### Build for All Platforms

```bash
# macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o dist/kamadhenu-darwin-arm64

# macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -o dist/kamadhenu-darwin-amd64

# Linux (amd64)
GOOS=linux GOARCH=amd64 go build -o dist/kamadhenu-linux-amd64

# Linux (arm64)
GOOS=linux GOARCH=arm64 go build -o dist/kamadhenu-linux-arm64

# Windows
GOOS=windows GOARCH=amd64 go build -o dist/kamadhenu-windows-amd64.exe
```

### Adding New Cows or Fortunes

1. Drop `.cow` files into `cows/`
2. Add fortune text to `fortunes/` (use `%` as delimiter)
3. Rebuild: `go build -o kamadhenu`

The `//go:embed` directives automatically include all files.

## Tech Stack

- **Language:** Go 1.22+
- **Dependencies:** 
  - `golang.org/x/term` (terminal width detection)
- **Embedded Data:** All cows and fortunes baked into the binary via `embed.FS`

## Distribution Roadmap

- [x] GitHub Releases (binaries for macOS/Linux)
- [ ] Homebrew Tap (`brew tap krry/kamadhenu`)
- [ ] Homebrew Core (`brew install kamadhenu`)
- [ ] AUR (Arch Linux)
- [ ] Snap/Flatpak

## Credits

- Inspired by the original shell script mashup
- Built on the legacy of `fortune`, `cowsay`, and `lolcat`
- Cow files from various community sources

## License

MIT

---

**Prostrate yourself before the Almighty.** 🐄✨
