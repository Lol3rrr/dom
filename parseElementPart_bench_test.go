package dom

import (
  "testing"
)

func benchParseElementPart(attributeCount int, b *testing.B) {
  input := "a"

  for i := 1; i < attributeCount; i++ {
    input += "a"
  }
  input += "=\"test\""

  b.ResetTimer()
  for n := 0; n < b.N; n++ {
    parseElementPart(input)
  }
}

func BenchmarkParseElementPart_001KeyLength(b *testing.B) { benchParseElementPart(0, b) }
func BenchmarkParseElementPart_005KeyLength(b *testing.B) { benchParseElementPart(5, b) }
func BenchmarkParseElementPart_010KeyLength(b *testing.B) { benchParseElementPart(10, b) }
func BenchmarkParseElementPart_025KeyLength(b *testing.B) { benchParseElementPart(25, b) }
func BenchmarkParseElementPart_050KeyLength(b *testing.B) { benchParseElementPart(50, b) }
func BenchmarkParseElementPart_100KeyLength(b *testing.B) { benchParseElementPart(100, b) }
