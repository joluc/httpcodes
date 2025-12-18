package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
)

//go:embed codes.json
var codesJSON []byte

const version = "0.1.0"

// StatusCode represents HTTP status code data from https://github.com/MattIPv4/status-codes
type StatusCode struct {
	Code        interface{} `json:"code"`
	Message     string      `json:"message"`
	Description string      `json:"description"`
}

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(lipgloss.Color("#8B5CF6")).
			Padding(0, 2).
			MarginBottom(1)

	tableStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#8B5CF6")).
			Padding(1, 2)

	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#8B5CF6")).
			Align(lipgloss.Center).
			MarginBottom(1)

	codeStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#A855F7"))

	descriptionStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFFFFF")).
				MarginTop(1)

	sourceStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#9CA3AF")).
			Italic(true).
			MarginTop(1)
)

func main() {
	versionFlag := flag.Bool("version", false, "Show version information")
	flag.BoolVar(versionFlag, "v", false, "Show version information (shorthand)")
	flag.Usage = showUsage
	flag.Parse()

	if *versionFlag {
		showVersion()
		os.Exit(0)
	}

	if flag.NArg() != 1 {
		showUsage()
		os.Exit(1)
	}

	codes, err := loadCodes()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	code := flag.Arg(0)
	if statusCode, found := codes[code]; found {
		displayStatusCode(statusCode)
	} else {
		fmt.Printf("Status code '%s' not found.\n", code)
		os.Exit(1)
	}
}

func showUsage() {
	title := codeStyle.Render("HTTP Status Code Lookup")

	fmt.Printf("\n%s\n", title)
	fmt.Printf("\nUsage: httpcodes [flags] <status_code>\n")
	fmt.Printf("\nFlags:\n")
	fmt.Println("  -v, --version    Show version information")
	fmt.Println("  -h, --help       Show this help message")
	fmt.Printf("\nExamples:\n")
	fmt.Println("  httpcodes 200")
	fmt.Println("  httpcodes 404")
	fmt.Println("  httpcodes 500")
}

func showVersion() {
	fmt.Printf("httpcodes version %s\n", version)
}

func loadCodes() (map[string]StatusCode, error) {
	var codes map[string]StatusCode
	if err := json.Unmarshal(codesJSON, &codes); err != nil {
		return nil, fmt.Errorf("could not parse embedded JSON: %w", err)
	}

	return codes, nil
}

func displayStatusCode(sc StatusCode) {

	statusLine := codeStyle.Render(fmt.Sprintf("HTTP %v: %s", sc.Code, sc.Message))

	description := descriptionStyle.Render(sc.Description)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		statusLine,
		description,
	)

	table := tableStyle.Render(content)

	fmt.Println(table)
	fmt.Printf("%s\n", sourceStyle.Render("Source: github.com/MattIPv4/status-codes"))
}
