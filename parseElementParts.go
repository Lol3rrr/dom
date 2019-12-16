package dom

import (
  "strings"
)

// parseElementParts parses the tags for an attribute
func parseElementParts(input string) map[string]string {
  result := map[string]string{}

  parts := strings.Split(input, " ")
  for _, part := range parts {
    key, value := parseElementPart(part)

    if len(key) == 0 || len(value) == 0 {
      continue
    }

    result[key] = value
  }

  return result
}
