package lint

import (
	"testing"
)

func Test_filesFromArgs(t *testing.T) {
	tests := []struct {
		Filename string
		Want     string
	}{
		{
			"testdata/01.tf",
			`{
  "provider": [
    {
      "google": [
        {
          "project": "my-project-id",
          "region": "us-central1"
        }
      ]
    }
  ]
}`,
		},
		{
			"testdata/02.tf",
			`{
  "provider": [
    {
      "google": [
        {
          "project": "my-project-id",
          "region": "us-central1"
        }
      ]
    },
    {
      "aws": [
        {
          "region": "us-east-1",
          "version": "~\u003e 2.0"
        }
      ]
    },
    {
      "google": [
        {
          "alias": "west",
          "project": "my-project-id",
          "region": "us-west1"
        }
      ]
    }
  ]
}`,
		},
		{
			"testdata/03.tf",
			`{
  "module": [
    {
      "foo": [
        {
          "bar": "baz",
          "data": [
            {
              "att1": "val1",
              "att2": "val2"
            }
          ]
        }
      ]
    }
  ]
}`,
		},
	}

	for _, test := range tests {
		t.Run(test.Filename, func(t *testing.T) {
			files, err := filesFromArgs([]string{test.Filename})
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if len(files) != 1 {
				t.Errorf("unexpected files length: got %d, want 1", len(files))
			}
			got := string(files[0].Data)
			// fmt.Println(got)
			if got != test.Want {
				t.Errorf("wrong result: got: %#v, want: %#v", got, test.Want)
			}
		})
	}
}
