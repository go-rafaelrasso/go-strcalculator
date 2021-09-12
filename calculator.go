package calculator

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	defaultDelimiter = ","
	delimiterSupport = "//"
)

func Add(input string) (int, error) {
	if len(strings.TrimSpace(input)) == 0 {
		return 0, nil
	}

	value := input
	delimiter := defaultDelimiter

	if strings.HasPrefix(value, delimiterSupport) {
		delimiter = input[2:3]
		value = strings.TrimLeft(input, input[0:3])
	}

	count := 0
	values := strings.Split(value, delimiter)

	negatives := []int{}

	for _, value := range values {
		converted, err := strconv.Atoi(strings.TrimSpace(value))
		if err != nil {
			log.Fatalf("Impossible convert %s to int", value)
		} else {
			if converted < 0 {
				negatives = append(negatives, converted)
			} else if converted >= 0 && converted <= 1000 {
				count += converted
			}

		}
	}

	if len(negatives) > 0 {
		return 0, errors.New(fmt.Sprintf("Negatives not allowed! %v", negatives))
	}

	return count, nil
}
