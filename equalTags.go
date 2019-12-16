package dom

// EqualTags checks if the two Tags have the same values and are therefor the same
func EqualTags(tag1, tag2 *Tag) bool {
  if tag1.Type != tag2.Type {
    return false
  }
  if !equalMaps(tag1.Attributes, tag2.Attributes) {
    return false
  }

  return true
}
