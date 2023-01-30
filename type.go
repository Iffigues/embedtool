package main

type EmbedIO struct {
	Input []byte
	Output []byte
}

func (e *EmbedIO)Write(p []byte) (n int, err error) {
	return
}

func (e *EmbedIO)Read(p []byte) (n int, err error) {
	return
}
