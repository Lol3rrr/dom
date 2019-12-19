package dom

import (
  "strings"
)

// parseElementPart parses one part of an elements Attributes
func parseElementPart(input string) (string, string) {
  divider := "=\""
  dividerLength := 2
  dividerIndex := strings.Index(input, divider)
  if dividerIndex == -1 || dividerIndex == 0 {
    return "", ""
  }

  key := input[:dividerIndex]
  value := input[dividerIndex + dividerLength:len(input) - 1]

  return key, value
}
