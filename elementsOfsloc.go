package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func countSLOC() { // ::: - -
	filenameOfThisFile1 := "/Users/quasar/Jap2-main/main.go"
	totalLines1, nonEmptyLines1 := reportSLOCstats(filenameOfThisFile1) // another locally-defined func; returns, and creates, local values of predetermined type

	filenameOfThisFile2 := "/Users/quasar/Jap2-main/constants.go"
	totalLines2, nonEmptyLines2 := reportSLOCstats(filenameOfThisFile2) // another locally-defined func; returns, and creates, local values of predetermined type

	filenameOfThisFile5 := "/Users/quasar/Jap2-main/elementsOfsloc.go"
	totalLines5, nonEmptyLines5 := reportSLOCstats(filenameOfThisFile5) // another locally-defined func; returns, and creates, local values of predetermined type

	filenameOfThisFile6 := "/Users/quasar/Jap2-main/functions.go"
	totalLines6, nonEmptyLines6 := reportSLOCstats(filenameOfThisFile6) // another locally-defined func; returns, and creates, local values of predetermined type

	filenameOfThisFile7 := "/Users/quasar/Jap2-main/globalVariables.go"
	totalLines7, nonEmptyLines7 := reportSLOCstats(filenameOfThisFile7) // another locally-defined func; returns, and creates, local values of predetermined type

	filenameOfThisFile8 := "/Users/quasar/Jap2-main/locateCard.go"
	totalLines8, nonEmptyLines8 := reportSLOCstats(filenameOfThisFile8) // another locally-defined func; returns, and creates, local values of predetermined type

	filenameOfThisFile9 := "/Users/quasar/Jap2-main/memoryFunctions.go"
	totalLines9, nonEmptyLines9 := reportSLOCstats(filenameOfThisFile9) // another locally-defined func; returns, and creates, local values of predetermined type

	filenameOfThisFile10 := "/Users/quasar/Jap2-main/objectsAndMethods.go"
	totalLines10, nonEmptyLines10 := reportSLOCstats(filenameOfThisFile10) // another locally-defined func; returns, and creates, local values of predetermined type

	filenameOfThisFile11 := "/Users/quasar/Jap2-main/pick_a_card.go"
	totalLines11, nonEmptyLines11 := reportSLOCstats(filenameOfThisFile11) // another locally-defined func; returns, and creates, local values of predetermined type

	filenameOfThisFile14 := "/Users/quasar/Jap2-main/pick_a_card_new.go"
	totalLines14, nonEmptyLines14 := reportSLOCstats(filenameOfThisFile14) // another locally-defined func; returns, and creates, local values of predetermined type

	filenameOfThisFile15 := "/Users/quasar/Jap2-main/prompts&directions.go"
	totalLines15, nonEmptyLines15 := reportSLOCstats(filenameOfThisFile15) // another locally-defined func; returns, and creates, local values of predetermined type

	filenameOfThisFile16 := "/Users/quasar/Jap2-main/statsFunctions.go"
	totalLines16, nonEmptyLines16 := reportSLOCstats(filenameOfThisFile16) // another locally-defined func; returns, and creates, local values of predetermined type

	totalLines := totalLines1 + totalLines2 + totalLines5 + totalLines6 + totalLines7 + totalLines8 + totalLines9 +
		totalLines10 + totalLines14 + totalLines15 + totalLines16 + totalLines11

	nonEmptyLines := nonEmptyLines1 + nonEmptyLines2 + nonEmptyLines5 + nonEmptyLines6 + nonEmptyLines7 + nonEmptyLines8 + nonEmptyLines9 +
		nonEmptyLines10 + nonEmptyLines14 + nonEmptyLines15 + nonEmptyLines16 + nonEmptyLines11

	fmt.Printf("Total lines of Code: %d\n", totalLines)
	fmt.Printf("Total lines of executable Code: %d\n", nonEmptyLines)

}

func reportSLOCstats(filepath string) (int, int) { // ::: - -
	// Patterns to identify comments, blank lines, and strings
	singleLineCommentPattern := `^\s*//`
	multiLineCommentPattern := `(?s)/\*.*?\*/`
	blankLinePattern := `^\s*$`
	stringLiteralPattern := `(?s)"(?:\\.|[^\\"])*?"|` + "`(?:\\.|[^`])*?`"

	// Compile regular expressions
	singleLineCommentRe := regexp.MustCompile(singleLineCommentPattern)
	multiLineCommentRe := regexp.MustCompile(multiLineCommentPattern)
	blankLineRe := regexp.MustCompile(blankLinePattern)
	stringLiteralRe := regexp.MustCompile(stringLiteralPattern)

	// Open the file
	file, err := os.Open(filepath)
	if err != nil {
		// return 0, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalLines := 0
	sloc := 0
	inMultiLineComment := false
	inMultiLineString := false

	for scanner.Scan() {
		line := scanner.Text()
		totalLines++

		// Check for blank lines
		if blankLineRe.MatchString(line) {
			continue
		}

		// Check for single-line comments
		if singleLineCommentRe.MatchString(line) {
			continue
		}

		// Check for multi-line comments
		if inMultiLineComment {
			if strings.Contains(line, "*/") {
				inMultiLineComment = false
				line = multiLineCommentRe.ReplaceAllString(line, "")
				if blankLineRe.MatchString(line) || singleLineCommentRe.MatchString(line) {
					continue
				}
			} else {
				continue
			}
		}

		if strings.Contains(line, "/*") {
			inMultiLineComment = true
			line = multiLineCommentRe.ReplaceAllString(line, "")
			if blankLineRe.MatchString(line) || singleLineCommentRe.MatchString(line) {
				continue
			}
		}

		// Check for multi-line strings
		if inMultiLineString {
			if strings.Count(line, "`")%2 != 0 || strings.Count(line, "\"")%2 != 0 {
				inMultiLineString = false
				line = stringLiteralRe.ReplaceAllString(line, "")
				if blankLineRe.MatchString(line) || singleLineCommentRe.MatchString(line) {
					continue
				}
			} else {
				continue
			}
		}

		if strings.Count(line, "`")%2 != 0 || strings.Count(line, "\"")%2 != 0 {
			inMultiLineString = true
			line = stringLiteralRe.ReplaceAllString(line, "")
			if blankLineRe.MatchString(line) || singleLineCommentRe.MatchString(line) {
				continue
			}
		}

		sloc++
	}

	if err := scanner.Err(); err != nil {
		// return 0, 0, err
	}

	return totalLines, sloc
}
