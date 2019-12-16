package dom

import (
  "testing"
  "strconv"
)

func benchEqualMaps(attributeLength int, b *testing.B) {
  map1 := map[string]string{}
  map2 := map[string]string{}

  for i := 0; i < attributeLength; i++ {
    strKey := strconv.Itoa(i)
    map1[strKey] = "test"
    map2[strKey] = "test"
  }

  b.ResetTimer()
  for n := 0; n < b.N; n++ {
    equalMaps(map1, map2)
  }
}

func BenchmarkEqualMaps_000Entrys(b *testing.B) { benchEqualMaps(0, b) }
func BenchmarkEqualMaps_005Entrys(b *testing.B) { benchEqualMaps(5, b) }
func BenchmarkEqualMaps_010Entrys(b *testing.B) { benchEqualMaps(10, b) }
func BenchmarkEqualMaps_025Entrys(b *testing.B) { benchEqualMaps(25, b) }
func BenchmarkEqualMaps_050Entrys(b *testing.B) { benchEqualMaps(50, b) }
func BenchmarkEqualMaps_100Entrys(b *testing.B) { benchEqualMaps(100, b) }
