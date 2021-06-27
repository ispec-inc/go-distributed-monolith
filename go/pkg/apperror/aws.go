package apperror

import (
	"github.com/ispec-inc/monorepo/go/pkg/aws"
)

func NewAWS(err aws.Error) error {
	var code Code
	switch err.Code() {
	case aws.ErrorCodeFileError:
		code = CodeInvalid
	case aws.ErrorCodeNetworkError:
		code = CodeError
	}
	return WithCode(code, err.OrgErr())
}
