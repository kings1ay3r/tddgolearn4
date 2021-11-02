package main

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://testsdf.test"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://testsdf.test",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://testsdf.test":        false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		go CheckWebsites(slowStubWebsiteChecker, urls)

	}
}
