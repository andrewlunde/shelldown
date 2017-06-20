package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// RootCmd is the root cobra command of this program
var RootCmd = &cobra.Command{
	Use:   "shelldown [markdown-file-paths]",
	Short: "Generate shell scripts for specially formatted markdown files",
	RunE:  rootCmd,
}

func main() {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}

func rootCmd(cmd *cobra.Command, args []string) error {

	if len(args) == 0 {
		return errors.New("must include files to script")
	}

	// Loop through all the files provided
	for _, filepath := range args {

		// Make sure the file exists
		//fmt.Printf("debug opening file %v\n", filepath)
		if _, err := os.Stat(filepath); os.IsNotExist(err) { //file doesn't exist
			return fmt.Errorf("Could not load file %v", filepath)
		}

		// Read file into a string object
		raw, err := readLines(filepath)
		if err != nil {
			return err
		}

		// Get the core code block
		shellCode := getShellCodeBlock(raw)

		// If the codeblock is empty just move on to the next .md file
		if shellCode == nil {
			fmt.Printf("file %v doesn't contain a valid shelldown header, Skipping...\n", filepath)
			continue
		}

		//replace all the holders in the codeblock
		shellCode, err = setCodeBlockHolders(shellCode, raw)
		if err != nil {
			return err
		}

		//set autogen text
		shellCode = writeAutoGenText(shellCode)

		//write the codeblock lines
		shellPath := filepath + ".sh"
		writeLines(shellCode, shellPath)
	}
	return nil
}

/////////////////////////////////////////////////////////////////////////////

const (

	//Start of the shell code stored within the markdown file
	scriptStart = "<!---"

	//End of the shell code stored within the markdown file
	scriptEnd = "-->"

	//Marker representing to grab a markdown codeblock for the shell script
	markerGet = "shelldown"

	//Marker representing to set a markdown codeblock for the shell script
	markerSet = "#shelldown"

	autoGenText = "#This script was generated with shelldown, see github.com/rigelrozanski/shelldown"
)

var (
	//regex scripts for determining placeholders in the shell script template
	regexIndex = regexp.MustCompile("([-?\\d]+)")
	regexSet   = regexp.MustCompile("(" + markerSet + "\\[\\d+\\]\\[-?\\d+\\])")
)

//Get the core code block for the shell script
func getShellCodeBlock(raw []string) []string {
	if !strings.HasPrefix(raw[0], scriptStart) {
		//fmt.Printf("debug %v %v\n", raw[0], scriptStart)
		return nil
	}

	//determine end line
	endCB := 0
	for i, line := range raw {
		if strings.HasPrefix(line, scriptEnd) && i > 0 {
			endCB = i //set to the previous line
			break
		}
	}

	//if no valid end codeblock was found
	if endCB == 0 {
		return nil
	}

	return raw[1:endCB]
}

func setCodeBlockHolders(shellCode, raw []string) ([]string, error) {
	var out []string
	for _, line := range shellCode {
		appendLine := line
		if strings.Contains(line, markerGet) {

			//get the index and the full holder (ex. shelldown[7][0])
			reg1 := regexSet.FindAllString(line, 1)
			if len(reg1) != 1 {
				return nil, errors.New("bad regexSet length")
			}

			holderStr := reg1[0]
			reg2 := regexIndex.FindAllString(holderStr, 2)
			if len(reg2) != 2 { //there should be 2 index files (ex. #shelldown[7][9])
				return nil, errors.New("bad regexIndex length")
			}

			index1, err := strconv.Atoi(reg2[0])
			if err != nil {
				return nil, err
			}

			index2, err := strconv.Atoi(reg2[1])
			if err != nil {
				return nil, err
			}

			holderValArr, err := getMarkdownHolder(raw, index1)
			if err != nil {
				return nil, err
			}

			//fmt.Printf("debug holder %v, index1 %v, index2 %v\n", holderValArr, index1, index2)
			if index2 == -1 { //special case to print all of the lines
				holderVal := strings.Join(holderValArr, "\n")
				appendLine = strings.Replace(line, holderStr, holderVal, -1)
			} else {
				appendLine = strings.Replace(line, holderStr, holderValArr[index2], -1)
			}
		}
		out = append(out, appendLine)
	}
	return out, nil
}

//Get the core code block for the shell script
func getMarkdownHolder(raw []string, index int) ([]string, error) {

	startCB, endCB := -1, -1
	markerStr := fmt.Sprintf("%v[%v]", markerGet, index)

	for i, line := range raw {
		if startCB == -1 &&
			strings.Contains(line, markerStr) &&
			strings.HasPrefix(line, "```") {
			startCB = i + 1
			break
		}
	}

	for i := startCB + 1; i < len(raw); i++ {
		if strings.HasPrefix(raw[i], "```") {
			endCB = i
			break
		}
	}

	if startCB == -1 || endCB == -1 {
		return nil, fmt.Errorf("invalid codeblock for index %v", index)
	}

	//fmt.Printf("debug start %v, end %v, index %v\n", startCB, endCB, index)
	return raw[startCB:endCB], nil
}

func writeAutoGenText(shellcode []string) []string {
	return append(
		[]string{shellcode[0], autoGenText},
		shellcode[1:]...,
	)
}

/////////////////////////////////////////////////////////////////////////////

// TODO add to github.com/rigelrozanski/common
// Credit: https://stackoverflow.com/questions/5884154/read-text-file-into-string-array-and-write

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
