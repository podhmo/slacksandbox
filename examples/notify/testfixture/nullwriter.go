package testfixture

// NullWriter :
type NullWriter struct{}

func (w NullWriter) Write(b []byte) (n int, err error) {
	return len(b), nil
}
