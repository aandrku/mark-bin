package assert

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Equal(t *testing.T, want, got any) {
	t.Helper()

	if diff := cmp.Diff(want, got, cmpopts.EquateErrors(), cmpopts.EquateComparable()); diff != "" {
		t.Errorf("assertion failed: (-want +got):\n%s", diff)
	}
}

func NilError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("expected nil err, got: %v", err)
	}
}
