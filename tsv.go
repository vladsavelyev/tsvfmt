package main

import (
	"bufio"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	delimPtr := flag.String("d", "", "column delimiter. If not specified, " +
		"a comma will be used for files with extension .csv, a tab otherwise")
	maxPreviewLinesPtr := flag.Int("l", 100, "The number of lines " +
		"to read in to estimate the size of a column [default: 100]" )
	maxColWidthPtr := flag.Int("max", 200, "Maximum width of a column. " +
		"Default: 200. Set to 0 to make unlimited.")
	flag.Parse()
	args := flag.Args()
	fPath := "-"
	if len(args) > 0 {
		fPath = args[0]
	}

	if *delimPtr == "" {
		*delimPtr = "\t"
		if strings.HasSuffix(fPath, ".csv") ||
			strings.HasSuffix(fPath, ".csv.gz") {
			*delimPtr = ","
		}
	}

	r, f := readGzFile(fPath)
	defer f.Close()
	scanner := bufio.NewScanner(r)
	tabView(scanner, *delimPtr, *maxPreviewLinesPtr, *maxColWidthPtr)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func readGzFile(fPath string) (io.Reader, *os.File) {
	if fPath == "" || fPath == "-" {
		// Reading from stdin
		f := os.Stdin
		r := bufio.NewReader(f)
		return r, f
	} else {
		f, err := os.Open(fPath)
		if err != nil {
			log.Fatal(err)
		}
		var r io.Reader = nil
		if strings.HasSuffix(fPath, ".gz") {
			gr, err := gzip.NewReader(f)
			if err != nil {
				log.Fatal(err)
			}
			r = gr
		} else {
			r = bufio.NewReader(f)
		}
		return r, f
	}
}

func tabView(scanner *bufio.Scanner, delim string, maxPreviewLines int, maxColWidth int) {
	var colWidths []int
	var colTypes []string
	var previewBuf []string
	var previewCnt = 0
	var inPreview = true

	for scanner.Scan() {
		line := scanner.Text()
		if inPreview {
			if strings.HasPrefix(line, "##") {
				previewBuf = append(previewBuf, line)
				continue
			}
			entries := strings.Split(strings.TrimSpace(line), delim)
			for colIdx, entry := range entries {
				colWidth := maxColWidth
				if len(colWidths) <= colIdx {
					if len(entry) < maxColWidth {
						colWidth = len(entry)
					}
					colWidths = append(colWidths, colWidth)
					colTypes = append(colTypes, "str")
				} else if len(entry) > colWidths[colIdx] {
					colWidths[colIdx] = len(entry)
					if colWidths[colIdx] > maxColWidth {
						colWidths[colIdx] = maxColWidth
					}
				}
				if _, err := strconv.Atoi(entry); err == nil {
					colTypes[colIdx] = "int"
				}
			}

			previewBuf = append(previewBuf, line)
			previewCnt += 1
			if previewCnt >= maxPreviewLines {
				for _, line := range previewBuf {
					writeCols(line, colWidths, colTypes, delim)
				}
				inPreview = false
				previewBuf = nil
			}
		} else {
			writeCols(line, colWidths, colTypes, delim)
		}
	}
	if previewBuf != nil {
		for _, line := range previewBuf {
			writeCols(line, colWidths, colTypes, delim)
		}
	}
}

func writeCols(line string, colWidths []int, colTypes []string, delim string) {
	var entries = strings.Split(strings.TrimSpace(line), delim)
	if len(entries) == 1 {
		fmt.Println(entries[0])
		return
	}
	for len(entries) < len(colWidths) {
		entries = append(entries, "")
	}
	for colIdx, entry := range entries {
		entry := strings.TrimSpace(entry)
		if len(entry) > colWidths[colIdx] {
			// if too big, show as much as possible, and indicate the
			// truncation with '$'
			entry = fmt.Sprintf("%s$", entry[0:colWidths[colIdx] - 1])
		} else if colTypes[colIdx] == "int" {
			// numbers are right-justified
			entryInt, err := strconv.Atoi(entry)
			if err == nil {
				entry = fmt.Sprintf("%*d", colWidths[colIdx], entryInt)
			} else {
				entry = fmt.Sprintf("%*s", colWidths[colIdx], entry)
			}
		} else {
			// text is right-justified
			entry = fmt.Sprintf("%-*s", colWidths[colIdx], entry)
		}
		fmt.Print(entry)
		if colIdx < len(entries) - 1 {
			fmt.Print("  ")  // separating with 2 spaces
		}
	}
	fmt.Print("\n")
}
