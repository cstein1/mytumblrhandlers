package mytumblrhandlers

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestSaveToJSON(t *testing.T) {
	filename := "testfile.txt"
	t.Cleanup(func() {
		err := os.Remove(filename)
		if err != nil {
			fmt.Println("could not remove file")
			t.Fail()
		}
	})
	testString := "Throw Back Thursday?"
	a := APITokens{
		CallBackURL: testString,
	}
	a.SaveToJSON(filename)

	readString, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("could not read file")
		t.Fatal()
	}
	b := &APITokens{}
	json.Unmarshal(readString, b)
	if b.CallBackURL != testString {
		fmt.Println("file had wrong contents")
		fmt.Printf("%s", readString)
		fmt.Println("VS")
		fmt.Printf("%s\n", testString)
		t.Fatal()
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
		t.Fail()
	}
}
