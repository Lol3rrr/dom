package dom

import (
  "testing"
)

func TestParseElementPart(t *testing.T) {
  tables := []struct{
    Input string
    OutputKey string
    OutputValue string
    OutputBool bool
  }{
    {
      Input: "class=\"test\"",
      OutputKey: "class",
      OutputValue: "test",
      OutputBool: true,
    },
    {
      Input: "id=\"yikes\"",
      OutputKey: "id",
      OutputValue: "yikes",
      OutputBool: true,
    },
    {
      Input: "id=\"\"",
      OutputKey: "id",
      OutputValue: "",
      OutputBool: true,
    },
    {
      Input: "id=",
      OutputKey: "",
      OutputValue: "",
      OutputBool: false,
    },
    {
      Input: "=\"test\"",
      OutputKey: "",
      OutputValue: "",
      OutputBool: false,
    },
  }

  for _, table := range tables {
    input := table.Input
    outputKey := table.OutputKey
    outputValue := table.OutputValue
    outputBool := table.OutputBool

    resultKey, resultValue, resultBool := parseElementPart(input)

    if resultBool != outputBool {
      t.Errorf("Did not return the Bool correctly('%s'), expected '%v' but returned '%v' \n", input, outputBool, resultBool)
      continue
    }
    if resultKey != outputKey {
      t.Errorf("Did not parse the Key right('%s'), expected '%s' but returned '%s' \n", input, outputKey, resultKey)
      continue
    }
    if resultValue != outputValue {
      t.Errorf("Did not parse the Value right('%s'), expected '%s' but returned '%s' \n", input, outputValue, resultValue)
      continue
    }
  }
}
