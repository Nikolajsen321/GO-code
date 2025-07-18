package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) { //Powershell bench test go test -bench="."
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}

}

// go mod init example.com/iteratio
