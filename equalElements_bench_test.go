package dom

import (
  "testing"
)

func benchEqualElements(at1, at2 map[string]string, child1, child2 []*Element, b *testing.B) {
  elem1 := &Element{
    Type: "typeTest",
    Attributes: at1,
    Inner: "innerTest",
    Children: child1,
  }
  elem2 := &Element{
    Type: "typeTest",
    Attributes: at2,
    Inner: "innerTest",
    Children: child2,
  }

  b.ResetTimer()
  for n := 0; n < b.N; n++ {
    EqualElements(elem1, elem2)
  }
}

func BenchmarkEqualElements_0Attributes_0Children(b *testing.B) {
  atr := map[string]string{}
  child := []*Element{}
  benchEqualElements(atr, atr, child, child, b)
}

func BenchmarkEqualElements_1Attributes_0Children(b *testing.B) {
  atr := map[string]string{
    "first": "firstTest",
  }
  child := []*Element{}
  benchEqualElements(atr, atr, child, child, b)
}
func BenchmarkEqualElements_2Attributes_0Children(b *testing.B) {
  atr := map[string]string{
    "first": "firstTest",
    "second": "secondTest",
  }
  child := []*Element{}
  benchEqualElements(atr, atr, child, child, b)
}
func BenchmarkEqualElements_3Attributes_0Children(b *testing.B) {
  atr := map[string]string{
    "first": "firstTest",
    "second": "secondTest",
    "third": "thirdTest",
  }
  child := []*Element{}
  benchEqualElements(atr, atr, child, child, b)
}
func BenchmarkEqualElements_4Attributes_0Children(b *testing.B) {
  atr := map[string]string{
    "first": "firstTest",
    "second": "secondTest",
    "third": "thirdTest",
    "forth": "forthTest",
  }
  child := []*Element{}
  benchEqualElements(atr, atr, child, child, b)
}
func BenchmarkEqualElements_5Attributes_0Children(b *testing.B) {
  atr := map[string]string{
    "first": "firstTest",
    "second": "secondTest",
    "third": "thirdTest",
    "forth": "forthTest",
    "fifth": "fifthTest",
  }
  child := []*Element{}
  benchEqualElements(atr, atr, child, child, b)
}

func BenchmarkEqualElements_0Attributes_1Children(b *testing.B) {
  atr := map[string]string{}
  child := []*Element{
    &Element{Type: "1Test"},
  }
  benchEqualElements(atr, atr, child, child, b)
}
func BenchmarkEqualElements_0Attributes_2Children(b *testing.B) {
  atr := map[string]string{}
  child := []*Element{
    &Element{Type: "1Test"},
    &Element{Type: "2Test"},
  }
  benchEqualElements(atr, atr, child, child, b)
}
func BenchmarkEqualElements_0Attributes_3Children(b *testing.B) {
  atr := map[string]string{}
  child := []*Element{
    &Element{Type: "1Test"},
    &Element{Type: "2Test"},
    &Element{Type: "3Test"},
  }
  benchEqualElements(atr, atr, child, child, b)
}
func BenchmarkEqualElements_0Attributes_4Children(b *testing.B) {
  atr := map[string]string{}
  child := []*Element{
    &Element{Type: "1Test"},
    &Element{Type: "2Test"},
    &Element{Type: "3Test"},
    &Element{Type: "4Test"},
  }
  benchEqualElements(atr, atr, child, child, b)
}
func BenchmarkEqualElements_0Attributes_5Children(b *testing.B) {
  atr := map[string]string{}
  child := []*Element{
    &Element{Type: "1Test"},
    &Element{Type: "2Test"},
    &Element{Type: "3Test"},
    &Element{Type: "4Test"},
    &Element{Type: "5Test"},
  }
  benchEqualElements(atr, atr, child, child, b)
}
