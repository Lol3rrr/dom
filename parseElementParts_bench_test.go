package dom

import (
  "testing"
  "strconv"
)

func benchParseElementParts(parts int, b *testing.B) {
  input := ""

  for i := 1; i < parts; i++ {
    input += strconv.Itoa(i)
    input += "=\"test\" "
  }

  b.ResetTimer()
  for n := 0; n < b.N; n++ {
    parseElementParts(input)
  }
}

func BenchmarkParseElementParts_001Parts(b *testing.B) { benchParseElementParts(0, b) }
func BenchmarkParseElementParts_005Parts(b *testing.B) { benchParseElementParts(5, b) }
func BenchmarkParseElementParts_010Parts(b *testing.B) { benchParseElementParts(10, b) }
func BenchmarkParseElementParts_025Parts(b *testing.B) { benchParseElementParts(25, b) }
func BenchmarkParseElementParts_050Parts(b *testing.B) { benchParseElementParts(50, b) }
func BenchmarkParseElementParts_100Parts(b *testing.B) { benchParseElementParts(100, b) }
