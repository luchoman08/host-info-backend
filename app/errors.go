package app

// ErrSSLabsFatalError means than a non recoverable error was retrieved
// by the sslabs API
type ErrSSLabsFatalError interface {
}

// ErrWhoIsNotFound means than the who is info cant be retrieved
type ErrWhoIsNotFound interface {
}

// ErrSSLabsMaxRetryExceed means than the info from sslabs API is not ready
// and max retry configured was exceed, but the info is loading
type ErrSSLabsMaxRetryExceed interface {
}
