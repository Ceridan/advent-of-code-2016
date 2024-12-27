package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

const codeTmpl = `
package main

import (
	"fmt"
	"os"
	"strings"
)

func Part1(data []string) int {
	return 0
}

func Part2(data []string) int {
	return 0
}

func main() {
	input, err := os.ReadFile("days/{{.Day}}/input.txt")
	if err != nil {
		panic(err)
	}
	data := strings.Split(string(input), "\n")
	data = data[:len(data)-1]

	fmt.Printf("Day {{.Day}}, part 1: %v\n", Part1(data))
	fmt.Printf("Day {{.Day}}, part 2: %v\n", Part2(data))
}
`

const testTmpl = `
package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	expected := 0
	data := []string{}

	got := Part1(data)

	if got != expected {
		t.Errorf("Day {{.Day}}, part1 = %d; want: %d", got, expected)
	}
}
`

func generateFile(tmpl string, day string) (string, error) {
	t := template.Must(template.New("code").Parse(tmpl))
	var res strings.Builder
	err := t.Execute(&res, struct{ Day string }{Day: day})
	if err != nil {
		return "", err
	}
	return res.String(), nil
}

func downloadInput(day string) (string, error) {
	sessionCookie := os.Getenv("AOC_SESSION")
	d := day
	if d[0:1] == "0" {
		d = day[1:]
	}
	url := fmt.Sprintf("https://adventofcode.com/2016/day/%s/input", day)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to download input: status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run gen.go --day=XX")
		os.Exit(1)
	}
	day := strings.Split(os.Args[1], "=")[1]

	dirPath := filepath.Join(".", "days", day)
	if _, err := os.Stat(dirPath); os.IsExist(err) {
		fmt.Printf("%s already exists\n", dirPath)
		os.Exit(1)
	}

	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Generate code file
	codeData, err := generateFile(codeTmpl, day)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	codePath := filepath.Join(dirPath, fmt.Sprintf("day%s.go", day))
	err = os.WriteFile(codePath, []byte(codeData), 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Generate test file
	testData, err := generateFile(testTmpl, day)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	testPath := filepath.Join(dirPath, fmt.Sprintf("day%s_test.go", day))
	err = os.WriteFile(testPath, []byte(testData), 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Load input data
	input, err := downloadInput(day)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputPath := filepath.Join(dirPath, "input.txt")
	err = os.WriteFile(inputPath, []byte(input), 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	git := exec.Command("git", "add", ".")
	err = git.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
