package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

func TestFilter(t *testing.T) {

	chain := []string{"snap1", "snap7", "snap10"}
	snapshots := []string{"snap1", "snap2", "snap3", "snap4", "snap5"}
	snapshots = Filter(snapshots, func(i string) bool {
		return Contains(chain, i)
	})

	fmt.Println(snapshots)
}

func TestParseAddresses(t *testing.T) {
	tests := []struct {
		name    string
		address string
		want    string
		want1   string
		want2   string
		wantErr bool
	}{
		{"correct address", "localhost:1234", "localhost:1234", "localhost:1235", "localhost:1236", false},
		{"correct address 2", "https://www.test.com:1234", "https://www.test.com:1234", "https://www.test.com:1235", "https://www.test.com:1236", false},
		{"bad address", "https://www.test.com/1234", "", "", "", true},
		{"correct address 3", "https://www.test.com:1234/status", "https://www.test.com:1234", "https://www.test.com:1235", "https://www.test.com:1236", false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := ParseAddresses(tt.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAddresses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseAddresses() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ParseAddresses() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("ParseAddresses() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestUUID(t *testing.T) {
	uuidPattern := regexp.MustCompile(`^[0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-4[0-9A-Fa-f]{3}-[89ABab][0-9A-Fa-f]{3}-[0-9A-Fa-f]{12}$`)
	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("test No. %d", i+1), func(t *testing.T) {
			got := UUID()
			match := uuidPattern.MatchString(got)
			if !match {
				t.Errorf("UUID() returned an invalid UUID: %s", got)
			}
		})
	}
}

func TestValidVolumeName(t *testing.T) {
	tests := []struct {
		scenario string
		name     string
		want     bool
	}{
		{"correct name", "My-Volume1", true},
		{"name too long", "This-Volume-Name-is-more-than-the-limit-accepted-adding-random-characters-to-reach-it", false},
		{"invalid name", "My=Volume1", false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.scenario, func(t *testing.T) {
			if got := ValidVolumeName(tt.name); got != tt.want {
				t.Errorf("ValidVolumeName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFileActualSize(t *testing.T) {
	data := []byte("test")
	fileName := "/tmp/data"
	minFileSize := int64(4096)
	err := ioutil.WriteFile(fileName, data, 0644)
	defer os.RemoveAll(fileName)
	if err != nil {
		t.Fatalf("error writing file %s", err)
	}
	if got := GetFileActualSize(fileName); got != minFileSize {
		t.Errorf("GetFileActualSize() = %v, want %v", got, minFileSize)
	}
}
