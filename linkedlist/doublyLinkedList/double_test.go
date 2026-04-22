package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func checkAndLog(t *testing.T, v string, wantSeq int, gotSeq *int) {
	t.Helper()

	t.Log(gotSeq, v)
	require.Equalf(t, wantSeq, *gotSeq, "expect seq to match: %s", v)

	*gotSeq = *gotSeq + 1
}

func TestXxx(t *testing.T) {
	seq := 1

	defer checkAndLog(t, "a", 10, &seq)
	t.Run("a", func(t *testing.T) {
		// defer checkAndLog(t, "a1", 3, &seq)
		defer checkAndLog(t, "a1", 1, &seq)

		t.Cleanup(func() {
			checkAndLog(t, "a1-2", 1, &seq)
			t.Cleanup(func() {
				checkAndLog(t, "a1-3", 2, &seq)
			})
		})
	})
	defer checkAndLog(t, "b", 9, &seq)
	t.Run("b", func(t *testing.T) {
		// defer checkAndLog(t, "b1", 4, &seq)
		defer checkAndLog(t, "b1", 3, &seq)
	})
	defer checkAndLog(t, "c", 8, &seq)
	t.Run("c", func(t *testing.T) {
		defer checkAndLog(t, "c1", 7, &seq)

		defer checkAndLog(t, "d", 6, &seq)
		t.Run("d", func(t *testing.T) {
			defer checkAndLog(t, "d1", 5, &seq)
		})
	})
}

// cleanup
// a1-2
// a1-3

// a1
// b1
// d1
// d
// c1
// c
// b
// a
