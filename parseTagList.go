package dom

import (
  "strings"
)

// ParseTagList parses the input and returns a Taglist for all tags found
func ParseTagList(tagInput string) (TagList) {
  result := TagList{
    Tags: make([]Tag, 0),
  }

  startIndex := strings.Index(tagInput, "<")
  for startIndex > -1 {
    endIndex := strings.Index(tagInput, ">")
    if endIndex <= 0 {
      break
    }

    contentEnd := endIndex
    if tagInput[endIndex - 1] == '/' {
      contentEnd--
    }

    content := tagInput[startIndex + 1: contentEnd]
    tagType, tagAtr := parseElementInfo(content)
    tag := Tag{
      Type: tagType,
      Attributes: tagAtr,
    }
    result.Tags = append(result.Tags, tag)

    tagInput = tagInput[endIndex + 1:]
    startIndex = strings.Index(tagInput, "<")
  }

  return result
}
