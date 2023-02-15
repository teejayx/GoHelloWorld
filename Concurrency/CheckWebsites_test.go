package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
	
	 
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"waat://furhurterwe.geds",
		"www.fb.com",
	}

	want := map[string]bool {
		"http://google.com": 		true,
		"waat://furhurterwe.geds":	false,
		"www.fb.com": 				true,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)


	if !reflect.DeepEqual(want, got){
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}