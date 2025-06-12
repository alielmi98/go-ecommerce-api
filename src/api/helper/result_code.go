package helper

type ResultCode int

const (
	Success           ResultCode = 0
	ValidationError   ResultCode = 40001
	AuthError         ResultCode = 40101
	ForbiddenError    ResultCode = 40301
	NotFoundError     ResultCode = 40401
	LimiterError      ResultCode = 42901
	OtpLimiterError   ResultCode = 42902
	CustomRecovery    ResultCode = 50001
	InternalError     ResultCode = 50002
	InvalidInputError ResultCode = 50003
	DatabaseError     ResultCode = 50004
	UnknownError      ResultCode = 50005
	BadRequest        ResultCode = 40002
)
