package dom

// equalMaps checks if the two Maps have the same values and are therefor the same
func equalMaps(map1, map2 map[string]string) bool {
  if len(map1) != len(map2) {
    return false
  }

  for key, value := range map1 {
    map2Value, ok := map2[key]
    if !ok {
      return false
    }

    if value != map2Value {
      return false
    }
  }

  return true
}
