package period

import (
	"testing"
	"time"
)

func TestBetween(t *testing.T) {
	d := time.Date(1984, time.March, 4, 0, 0, 0, 0, &time.Location{})
	n := time.Date(2022, time.November, 15, 0, 0, 0, 0, &time.Location{})
	p := Between(d, n)

	var (
		expectedYear  = 38
		expectedMonth = 8
		expectedDay   = 11
	)

	if p.Years != expectedYear || p.Months != expectedMonth || p.Days != expectedDay {
		t.Fatalf("expected period is different from got period\n"+
			"expected: %d year, %d month and %d day\ngot: %d year, %d month and %d day\n",
			expectedYear, expectedMonth, expectedDay, p.Years, p.Months, p.Days)
	}
}

func BenchmarkBetween(b *testing.B) {
	d := time.Date(1984, time.March, 4, 0, 0, 0, 0, &time.Location{})
	n := time.Date(2022, time.November, 15, 0, 0, 0, 0, &time.Location{})

	for i := 0; i < b.N; i++ {
		Between(d, n)
	}
}