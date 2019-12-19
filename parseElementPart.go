package dom

import (
  "strings"
)

const divider string = "=\""
const dividerLength int = 2

// parseElementPart parses one part of an elements Attributes
func parseElementPart(input string) (string, string) {
  dividerIndex := strings.Index(input, divider)
  if dividerIndex == -1 || dividerIndex == 0 {
    return "", ""
  }

  key := input[:dividerIndex]
  value := input[dividerIndex + dividerLength:len(input) - 1]

  return key, value
}
