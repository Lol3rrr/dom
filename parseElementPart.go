package dom

import (
  "strings"
)

const divider string = "=\""
const dividerLength int = 2

// parseElementPart parses one part of an elements Attributes
func parseElementPart(input string) (string, string, bool) {
  dividerIndex := strings.Index(input, divider)
  if dividerIndex == -1 || dividerIndex == 0 {
    return "", "", false
  }

  return input[:dividerIndex], input[dividerIndex + dividerLength:len(input) - 1], true
}
