package dom

import (
  "testing"
  "reflect"
)

func TestParseTagList(t *testing.T) {
  tables := []struct{
    InputDom string
    Output TagList
  }{
    {
      InputDom: "<p>",
      Output: TagList{
        Tags: []Tag{
          Tag{
            Type: "p",
            Attributes: map[string]string{},
          },
        },
      },
    },
    {
      InputDom: "<p class=\"yikes\">",
      Output: TagList{
        Tags: []Tag{
          Tag{
            Type: "p",
            Attributes: map[string]string{
              "class": "yikes",
            },
          },
        },
      },
    },
    {
      InputDom: "<p class=\"yikes\"/>",
      Output: TagList{
        Tags: []Tag{
          Tag{
            Type: "p",
            Attributes: map[string]string{
              "class": "yikes",
            },
          },
        },
      },
    },
  }

  for _, table := range tables {
    inputDom := table.InputDom
    output := table.Output

    result := ParseTagList(inputDom)

    if !reflect.DeepEqual(result.Tags, output.Tags) {
      t.Errorf("Did not parse the Tag Lists correctly \n")
      t.Errorf("Returned %v \n", result.Tags)
      t.Errorf("Expected %v \n", output.Tags)
      continue
    }
  }
}
