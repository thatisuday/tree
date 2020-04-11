package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/logrusorgru/aurora"

	"github.com/thatisuday/commando"
)

// get human-readable file size
func getSize(size int64) string {
	switch {

	// bytes
	case size < 1000:
		return fmt.Sprintf("%d bytes", size)

	// kb
	case size < (1000 * 1000):
		return fmt.Sprintf("%.1fkb", float64(size)/1000.0)

	// mb
	case size < (1000 * 1000 * 1000):
		return fmt.Sprintf("%.1fmb", float64(size)/(1000*1000))

	// gb
	case size < (1000 * 1000 * 1000 * 1000):
		return fmt.Sprintf("%.1fgb", float64(size)/(1000*1000*1000))
	}

	// default
	return ""
}

// this function returns the information of a file
func getFileInfo(isInfo bool, options map[string]commando.FlagValue, fileInfo os.FileInfo) string {

	// options
	var displaySize, displayMode, displayColor bool
	displayColor, _ = options["color"].GetBool()

	if !isInfo {
		displaySize, _ = options["size"].GetBool()
		displayMode, _ = options["mode"].GetBool()
	}

	/*--------*/

	// name of the file
	var info string

	if fileInfo.IsDir() && displayColor {
		info = aurora.Sprintf("%s", aurora.Green(fileInfo.Name()))
	} else if displayColor {
		info = aurora.Sprintf("%s", aurora.Yellow(fileInfo.Name()).Bold())
	} else {
		info = fileInfo.Name()
	}

	/*--------*/

	// size of the file
	if (displaySize || isInfo) && !fileInfo.IsDir() {
		value := fmt.Sprintf(" (%v)", getSize(fileInfo.Size()))

		if displayColor {
			info += aurora.Sprintf("%s", aurora.Cyan(value).Italic())
		} else {
			info += value
		}
	}

	// display file mode
	if displayMode || isInfo {

		// last 9 bits are important (rest should be zero)
		value := fmt.Sprintf(" (%o)", (fileInfo.Mode()<<21)>>21)

		if displayColor {
			info += aurora.Sprintf("%s", aurora.Magenta(value).Italic())
		} else {
			info += value
		}
	}

	return info
}

// this function prints the contents of a single directory
func printContent(isInfo bool, dirPath string, options map[string]commando.FlagValue, depth int, paddingPrefix string) {

	// read contents of the directory
	contents, err := ioutil.ReadDir(dirPath)

	/*----------------------*/

	// on error, return
	if err != nil {
		return
	}

	/*----------------------*/

	// for each item in the `contents`, print item information
OUTER:
	for index, fileInfo := range contents {

		// get list of directories to ignore
		ignore, _ := options["ignore"].GetString()
		ignoreList := strings.Split(ignore, ",")

		// continue on ignored files
		if fileInfo.IsDir() {
			for _, ignoreItem := range ignoreList {
				if fileInfo.Name() == ignoreItem {
					continue OUTER
				}
			}
		}

		/*--------*/

		// prefix
		prefix := paddingPrefix + "├──"

		// prefix for the last item
		if index == len(contents)-1 {
			prefix = paddingPrefix + "└──"
		}

		/*--------*/

		// new padding prefix value
		_paddingPrefix := paddingPrefix + "|  "

		if index == len(contents)-1 {
			_paddingPrefix = paddingPrefix + "   "
		}

		/*--------*/

		// get option values
		level, _ := options["level"].GetInt()

		// print information of the file
		fmt.Println(fmt.Sprintf("%s %s", prefix, getFileInfo(isInfo, options, fileInfo)))

		// recursion when `level` is greater than 1
		if fileInfo.IsDir() && level > 1 && depth != level-1 {
			printContent(isInfo, filepath.Join(dirPath, fileInfo.Name()), options, depth+1, _paddingPrefix)
		}
	}
}

// this function prints the contents of a relative or an absolute directory
func list(isInfo bool, dir string, options map[string]commando.FlagValue) {

	// get an absolute directory path
	var dirPath string
	if filepath.IsAbs(dir) {
		dirPath = dir
	} else {
		workdingDir, _ := os.Getwd()
		dirPath = filepath.Join(workdingDir, dir)
	}

	// print contents
	printContent(isInfo, dirPath, options, 0, "")
}
