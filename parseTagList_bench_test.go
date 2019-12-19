package dom

import (
  "testing"
  "strconv"
)

func benchParseTagList(parts int, b *testing.B) {
  input := ""

  for i := 1; i < parts; i++ {
    input += strconv.Itoa(i)
    input += "<p id=\"test\">"
  }

  b.ResetTimer()
  for n := 0; n < b.N; n++ {
    ParseTagList(input)
  }
}

func BenchmarkParseTagList_001Parts(b *testing.B) { benchParseTagList(0, b) }
func BenchmarkParseTagList_005Parts(b *testing.B) { benchParseTagList(5, b) }
func BenchmarkParseTagList_010Parts(b *testing.B) { benchParseTagList(10, b) }
func BenchmarkParseTagList_025Parts(b *testing.B) { benchParseTagList(25, b) }
func BenchmarkParseTagList_050Parts(b *testing.B) { benchParseTagList(50, b) }
func BenchmarkParseTagList_100Parts(b *testing.B) { benchParseTagList(100, b) }
