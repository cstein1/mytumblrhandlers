package mytumblrhandlers

import (
	"encoding/json"
	"os"
	"testing"
)

func TestSaveToJSON(t *testing.T) {
	filename := "testfile.txt"
	t.Cleanup(func() {
		err := os.Remove(filename)
		if err != nil {
			t.Fatal("could not remove file")
		}
	})
	testString := "Throw Back Thursday?"
	a := APITokens{
		CallBackURL: testString,
	}
	a.SaveToJSON(filename)

	readString, err := os.ReadFile(filename)
	if err != nil {
		t.Fatal("could not read file")
	}
	b := &APITokens{}
	json.Unmarshal(readString, b)
	if b.CallBackURL != testString {
		t.Fatalf("file had wrong contents: \n%s\n\n VS \n\n%s\n", readString, testString)
	}
}

func TestLoadFromJSON(t *testing.T) {
	filename := "testfile.txt"
	t.Cleanup(func() {
		err := os.Remove(filename)
		if err != nil {
			t.Fail()
		}
	})
	testString := "Throw Back Thursday?"
	a := APITokens{
		CallBackURL: testString,
	}
	a.SaveToJSON(filename)
	b := &APITokens{}
	b.LoadFromJSON(filename)
	if b.CallBackURL != testString {
		t.Fatal("callback URL not equal to teststring")
	}
}
