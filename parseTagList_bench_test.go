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

func BenchmarkParseTagList_0001Parts(b *testing.B) { benchParseTagList(0, b) }
func BenchmarkParseTagList_0005Parts(b *testing.B) { benchParseTagList(5, b) }
func BenchmarkParseTagList_0010Parts(b *testing.B) { benchParseTagList(10, b) }
func BenchmarkParseTagList_0025Parts(b *testing.B) { benchParseTagList(25, b) }
func BenchmarkParseTagList_0050Parts(b *testing.B) { benchParseTagList(50, b) }
func BenchmarkParseTagList_0100Parts(b *testing.B) { benchParseTagList(100, b) }
func BenchmarkParseTagList_0500Parts(b *testing.B) { benchParseTagList(500, b) }
func BenchmarkParseTagList_1000Parts(b *testing.B) { benchParseTagList(1000, b) }
