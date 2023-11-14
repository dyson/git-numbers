package gitnumbers

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

func Run(args []string) int {
	cmd := exec.Command("git", "-c", "color.status=always", "-c", "status.short=false", "status")
	gitStatusBytes, err := cmd.CombinedOutput()
	gitStatus := string(gitStatusBytes)
	if err != nil {
		return cmd.ProcessState.ExitCode()
	}

	gitStatusAnnotated, indexedFiles := parseGitStatus(gitStatus)

	if len(args) == 1 {
		fmt.Print(gitStatusAnnotated)
	} else {
		cmd := exec.Command("git", parseArgs(os.Args[1:], indexedFiles)...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return cmd.ProcessState.ExitCode()
		}
	}

	return 0
}

func parseArgs(args []string, indexedFiles map[int]string) []string {
	var parsedArgs []string

	re := regexp.MustCompile(`(\d+)-(\d+)`)

	for _, arg := range args {
		if i, err := strconv.Atoi(arg); err == nil {
			parsedArgs = append(parsedArgs, indexedFiles[i])
			continue
		}

		rangeArg := re.FindStringSubmatch(arg)

		if rangeArg != nil {
			start, _ := strconv.Atoi(rangeArg[1])
			end, _ := strconv.Atoi(rangeArg[2])

			for i := start; i <= end; i++ {
				parsedArgs = append(parsedArgs, indexedFiles[i])
			}

			continue
		}

		parsedArgs = append(parsedArgs, arg)
	}

	return parsedArgs
}
