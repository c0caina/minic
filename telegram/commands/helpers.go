package commands

import (
	"fmt"
	"regexp"
)



func parseArg(command string, args *map[string]string){
	for openingWords := range *args {

		var closingWords string
		for k := range *args {
			if k != openingWords {
				closingWords += fmt.Sprintf("| %s", k)
			}
		}
		
		re := regexp.MustCompile(fmt.Sprintf(" %s(.*?)(?:$%s)", openingWords, closingWords))
		result := re.FindStringSubmatch(command)
		if len(result) != 0 {
			(*args)[openingWords] = result[1]
		}
	}
}