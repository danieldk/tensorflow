package tensorflow

import "testing"

func TestSession(t *testing.T) {
	opts := NewSessionOptions()
	defer opts.Close()

	graph := NewGraph()
	defer graph.Close()

	sess, err := NewSession(graph, opts)
	if err != nil {
		t.Error(err)
	}
	defer sess.Close()
}
