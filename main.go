package main

import (
	"bufio"
	"embed"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/krry/kamadhenu/internal/colors"
	"github.com/krry/kamadhenu/internal/config"
	"github.com/krry/kamadhenu/internal/cowsay"
	"github.com/krry/kamadhenu/internal/fortune"
	"github.com/krry/kamadhenu/internal/typewriter"
)

//go:embed cows/*.cow
var cowFiles embed.FS

//go:embed fortunes/*
var fortuneFiles embed.FS

const version = "2.0.0"

func main() {
	rand.Seed(time.Now().UnixNano())

	// Parse flags
	helpFlag := flag.Bool("h", false, "Show help")
	versionFlag := flag.Bool("v", false, "Show version")
	noColorFlag := flag.Bool("no-color", false, "Disable rainbow colors")
	monoFlag := flag.Bool("mono", false, "Disable rainbow colors (alias for --no-color)")
	fastFlag := flag.Bool("fast", false, "Skip typewriter effect")
	unicornFlag := flag.Bool("unicorn", false, "Chaotic rainbow mode")
	
	flag.Usage = showHelp
	flag.Parse()

	if *versionFlag {
		fmt.Printf("Kamadhenu v%s\n", version)
		return
	}

	if *helpFlag {
		showHelp()
		return
	}

	// Initialize components
	fortune.Init(fortuneFiles)
	cowsay.Init(cowFiles)
	
	config := Config{
		NoColor: *noColorFlag || *monoFlag,
		Fast:    *fastFlag,
		Unicorn: *unicornFlag,
	}
	
	// Set rainbow mode (only if colors enabled)
	if !config.NoColor {
		if *unicornFlag {
			colors.SetMode(colors.ModeUnicorn)
		} else {
			colors.SetMode(colors.ModeNormal)
		}
	}

	args := flag.Args()
	
	if len(args) == 0 {
		// Default: single fortune
		showFortune(config)
	} else if args[0] == "summon" {
		// Interactive mode
		runInteractive(config)
	} else if args[0] == "sources" {
		// Source management
		runSources(args[1:])
	} else if args[0] == "alias" {
		// Easter egg: install shell alias
		runAliasInstall()
	} else {
		// Single fortune or custom text
		arg := args[0]
		
		if n, err := strconv.Atoi(arg); err == nil && n > 0 {
			// Repeat N times
			runRepeat(n, config)
		} else {
			// Custom text
			runCustomText(arg, config)
		}
	}
}

type Config struct {
	NoColor bool
	Fast    bool
	Unicorn bool
}

func showHelp() {
	help := `
Kamadhenu - An udderly absurd fortune teller

USAGE:
  kamadhenu [options]              Show one random fortune
  kamadhenu summon                 Enter interactive mode
  kamadhenu [options] "text"       Display custom text
  kamadhenu [options] <N>          Show N random fortunes
  kamadhenu sources <cmd>          Manage fortune sources
  kamadhenu alias                  Install 'kama' shell alias

OPTIONS:
  -h, --help             Show this help and exit
  -v, --version          Show version and exit
  --fast                 Skip typewriter effect (instant output)
  --mono, --no-color     Disable rainbow gradient colors
  --unicorn              Chaotic rainbow mode (wild color combos)

EXAMPLES:
  kamadhenu                         # One fortune
  kamadhenu summon                  # Interactive mode
  kamadhenu "sudo make me a sandwich"
  kamadhenu 5                       # 5 fortunes
  kamadhenu --fast --mono           # Quick, plain output
  kamadhenu --mono summon           # Interactive, monochrome
  kamadhenu --unicorn "chaos!"      # Wild rainbow colors

NOTE: Flags must come BEFORE subcommands (e.g., --mono summon, not summon --mono)

INTERACTIVE MODE (kamadhenu summon):
  <Enter>       Random fortune
  "some text"   Custom text
  <number>      N fortunes
  Ctrl-C        Exit

SOURCES:
  kamadhenu sources list           List configured fortune sources
  kamadhenu sources add <path>     Add local file or URL
  kamadhenu sources rm <path>      Remove a source
  kamadhenu sources reset          Reset to embedded fortunes

  Supported formats: Fortune (% delimited), JSON, plain text, paragraphs

STATS:
  Fortunes: 13,443 • Cows: 109 • Dependencies: 0

See 'man kamadhenu' for more details.
`
	fmt.Println(help)
}

func runInteractive(config Config) {
	// Show intro
	showIntro(config)
	
	reader := bufio.NewReader(os.Stdin)
	
	for {
		showPrompt(config)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		if input == "" {
			// Just show a fortune
			showFortune(config)
		} else if n, err := strconv.Atoi(input); err == nil && n > 0 {
			runRepeat(n, config)
		} else {
			runCustomText(input, config)
		}
	}
}

func showIntro(config Config) {
	intro := `
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
_.--._._._.--._..-.-.___.--._.__..__._._._.-..
^^+~*^^**^*~*^*^*~^*^*^*^*~^*~*^+^+^^^+^^+^^+^
`
	printOutput(intro, config)
}

func showPrompt(config Config) {
	prompt := `

_.-^-.__._.-..-.__.__._.-^-.__.-.._..-..__.-._
+^+=^=+^+=+^+^+^==^+^+^+=+^+=^+^+=^=+^+=+^+^+^
       You may put words in my mouth,
         or <ENTER> if you dare.
;,:.,:;,.,.:;.,.;:..,;.,:.;..;.,.;:..,;.,:.;..

$ `
	printOutput(prompt, config)
}

func showFortune(config Config) {
	text := fortune.Random()
	cow := cowsay.Random()
	output := cowsay.Say(text, cow)
	
	if config.Fast {
		printOutput(output, config)
	} else {
		typewriter.Print(output, config.NoColor)
	}
}

func runRepeat(n int, config Config) {
	for i := 0; i < n; i++ {
		showFortune(config)
		
		if i < n-1 {
			fmt.Printf("\n%d remaining\n", n-i-1)
			fmt.Print("Press Enter to continue (or Ctrl-C to exit)... ")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		}
	}
}

func runCustomText(text string, config Config) {
	cow := cowsay.Random()
	output := cowsay.Say(text, cow)
	
	if config.Fast {
		printOutput(output, config)
	} else {
		typewriter.Print(output, config.NoColor)
	}
	
	// Footer
	footer := "\n" + strings.Repeat(" ", 20) + "Kamadhenu\n"
	printOutput(footer, config)
}

func printOutput(text string, config Config) {
	if config.NoColor {
		fmt.Print(text)
	} else {
		colors.Rainbow(text)
	}
}

func runSources(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: kamadhenu sources <list|add|rm|reset>")
		return
	}
	
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}
	
	switch args[0] {
	case "list":
		if len(cfg.Sources) == 0 {
			fmt.Println("No custom sources configured. Using embedded fortunes.")
			fmt.Printf("Config: %s\n", config.ConfigPath())
		} else {
			fmt.Println("Custom fortune sources:")
			for i, s := range cfg.Sources {
				fmt.Printf("  %d. %s\n", i+1, s)
			}
			fmt.Printf("\nConfig: %s\n", config.ConfigPath())
		}
		
	case "add":
		if len(args) < 2 {
			fmt.Println("Usage: kamadhenu sources add <filepath|url>")
			return
		}
		source := args[1]
		if cfg.AddSource(source) {
			if err := cfg.Save(); err != nil {
				fmt.Printf("Error saving config: %v\n", err)
				return
			}
			fmt.Printf("✓ Added source: %s\n", source)
		} else {
			fmt.Printf("Source already exists: %s\n", source)
		}
		
	case "rm", "remove":
		if len(args) < 2 {
			fmt.Println("Usage: kamadhenu sources rm <filepath|url>")
			return
		}
		source := args[1]
		if cfg.RemoveSource(source) {
			if err := cfg.Save(); err != nil {
				fmt.Printf("Error saving config: %v\n", err)
				return
			}
			fmt.Printf("✓ Removed source: %s\n", source)
		} else {
			fmt.Printf("Source not found: %s\n", source)
		}
		
	case "reset":
		cfg.Reset()
		if err := cfg.Save(); err != nil {
			fmt.Printf("Error saving config: %v\n", err)
			return
		}
		fmt.Println("✓ Reset to embedded fortunes")
		
	default:
		fmt.Printf("Unknown command: %s\n", args[0])
		fmt.Println("Usage: kamadhenu sources <list|add|rm|reset>")
	}
}

func runAliasInstall() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	aliasLine := "alias kama='kamadhenu'"
	
	// Detect shell
	shell := os.Getenv("SHELL")
	var rcFiles []string
	
	if strings.Contains(shell, "fish") {
		rcFiles = []string{filepath.Join(home, ".config/fish/config.fish")}
		aliasLine = "alias kama kamadhenu"
	} else if strings.Contains(shell, "zsh") {
		rcFiles = []string{filepath.Join(home, ".zshrc")}
	} else {
		// Bash or unknown - try both
		rcFiles = []string{
			filepath.Join(home, ".bashrc"),
			filepath.Join(home, ".bash_profile"),
		}
	}
	
	// Check if alias already exists
	for _, rcFile := range rcFiles {
		data, err := os.ReadFile(rcFile)
		if err != nil {
			continue
		}
		if strings.Contains(string(data), "alias kama") {
			fmt.Println("✓ Alias 'kama' already configured!")
			fmt.Printf("  %s\n", rcFile)
			return
		}
	}
	
	// Add to first existing rc file
	for _, rcFile := range rcFiles {
		if _, err := os.Stat(rcFile); err == nil {
			f, err := os.OpenFile(rcFile, os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Printf("Error opening %s: %v\n", rcFile, err)
				continue
			}
			defer f.Close()
			
			fmt.Fprintf(f, "\n# Kamadhenu alias (installed by kamadhenu alias)\n%s\n", aliasLine)
			fmt.Println("✓ Alias 'kama' installed!")
			fmt.Printf("  %s\n", rcFile)
			fmt.Println("\nRestart your shell or run:")
			fmt.Printf("  source %s\n", rcFile)
			return
		}
	}
	
	fmt.Println("Could not find shell config file.")
	fmt.Printf("Add this to your shell config manually:\n  %s\n", aliasLine)
}
