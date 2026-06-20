package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ParseFastaFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	seqStarted := false
	sb := strings.Builder{}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(strings.TrimSpace(line), ">") {
			if seqStarted {
				return "", fmt.Errorf("multiple seqs found in %s. expected a single seq", path)
			} else {
				seqStarted = true
				continue
			}
		}

		sb.WriteString(line)
	}

	if scanner.Err() != nil {
		return "", scanner.Err()
	}

	return sb.String(), nil
}
