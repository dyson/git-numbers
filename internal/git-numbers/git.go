package gitnumbers

import (
	"fmt"
	"regexp"
	"strings"

	r "github.com/dyson/git-numbers/pkg/regexp"
)

func parseGitStatus(status string) (string, map[int]string) {
	re := regexp.MustCompile(`\t(\x1b\[[0-9;]+m)(.+:\s+)?(.+)(\x1b\[m)`)

	index := 0
	indexedFiles := map[int]string{}

	annotatedStatus := r.ReplaceAllSubmatchFunc(re, status, func(groups []string) string {
		colourStart := groups[1]
		description := groups[2]
		files := groups[3]
		colourEnd := groups[4]

		fs := strings.Split(files, " -> ")
		for k, f := range fs {
			index += 1
			indexedFiles[index] = f
			fs[k] = fmt.Sprintf("[%d] %s", index, f)
		}
		files = strings.Join(fs, " -> ")

		return fmt.Sprintf("\t%s%s%s%s", colourStart, description, files, colourEnd)
	})

	return annotatedStatus, indexedFiles
}
