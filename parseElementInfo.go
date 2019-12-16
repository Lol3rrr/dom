package dom

import (
  "strings"
)

// parseElementInfo parses the Element and returns its type and all its attributes
func parseElementInfo(rawElementInfo string) (string, map[string]string) {
  if len(rawElementInfo) < 1 {
    return "", map[string]string{}
  }

  typeEnd := strings.Index(rawElementInfo, " ")
  if typeEnd < 0 {
    typeEnd = len(rawElementInfo)
  }

  elemType := rawElementInfo[:typeEnd]
  elemAttributes := getElementParts(rawElementInfo)

  return elemType, elemAttributes
}
