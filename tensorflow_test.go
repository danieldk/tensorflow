package tensorflow

import "testing"

func TestSession(t *testing.T) {
	opts := NewSessionOptions()
	defer opts.Close()

	sess, err := NewSession(opts)
	if err != nil {
		t.Error(err)
	}
	defer sess.Close()
}
