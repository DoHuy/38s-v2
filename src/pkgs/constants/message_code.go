package constants
type ErrorCode int
const (
	// Application error
	ApplicationError       ErrorCode = 500000
	NilPointerError        ErrorCode = 500002
	ValidationNotSupported ErrorCode = 500003
	IdentityIsMissing      ErrorCode = 500004

)

