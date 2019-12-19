package dom

import (
  "testing"
  "strconv"
)

func benchParseElementInfo(attributeCount int, b *testing.B) {
  input := "p"

  for i := 0; i < attributeCount; i++ {
    input += " "
    input += strconv.Itoa(i)
    input += "=\"test\""
  }

  b.ResetTimer()
  for n := 0; n < b.N; n++ {
    parseElementInfo(input)
  }
}

func BenchmarkParseElementInfo_000Attributes(b *testing.B) { benchParseElementInfo(0, b) }
func BenchmarkParseElementInfo_005Attributes(b *testing.B) { benchParseElementInfo(5, b) }
func BenchmarkParseElementInfo_010Attributes(b *testing.B) { benchParseElementInfo(10, b) }
func BenchmarkParseElementInfo_025Attributes(b *testing.B) { benchParseElementInfo(25, b) }
func BenchmarkParseElementInfo_050Attributes(b *testing.B) { benchParseElementInfo(50, b) }
func BenchmarkParseElementInfo_100Attributes(b *testing.B) { benchParseElementInfo(100, b) }
