package oexec

// Output is the struct that hold a command's output
type Output struct {
	// Stdout will have byte output of exec.Command
	Stdout []byte
	// Stderr will have error object of exec.Command
	Stderr error
}
