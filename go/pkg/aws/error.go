package aws

const (
	ErrorCodeFileError    = ErrorCode("file error")
	ErrorCodeNetworkError = ErrorCode("network error")
)

type ErrorCode string

func (c ErrorCode) String() string {
	return string(c)
}

type Error interface {
	Code() ErrorCode
	Error() string
	OrgErr() error
}

type awsError struct {
	code ErrorCode
	err  error
}

func (e awsError) Code() ErrorCode {
	return e.code
}

func (e awsError) Error() string {
	return e.err.Error()
}

func (e awsError) OrgErr() error {
	return e.err
}

func newFileError(err error) awsError {
	return awsError{
		code: ErrorCodeFileError,
		err:  err,
	}
}

func newNetworkError(err error) awsError {
	return awsError{
		code: ErrorCodeNetworkError,
		err:  err,
	}
}
