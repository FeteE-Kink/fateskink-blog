package constants

const (
	AuthorizationHeader = "Fateskink-Authorization"
	GinContextKey       = "FTKContextKey"
	ContextCurrentUser  = "CurrentUser"

	MaxStringLength   = 255
	MaxLongTextLength = 4294967295

	BadRequestErrorCode = 400
	BadRequestErrorMsg  = "Bad Request"

	NotFoundErrorCode = 404
	NotFoundErrorMsg  = "Not Found"

	UnauthorizedErrorCode = 401
	UnauthorizedErrorMsg  = "You need to sign in to perform this action"

	UnprocessableContentErrorCode = 422
	UnprocessableContentErrorMsg  = "Please check your input"

	DDMMYYYY_DateFormat       = "02-01-2006" // "Date-Month-Year"
	HUMAN_DDMMYYYY_DateFormat = "%d-%m-%y"

	YYYYMMDD_DateFormat       = "2006-01-02" // "Month-Date-Year"
	HUMAN_YYYYMMDD_DateFormat = "%y-%m-%d"

	DateTimeZoneFormat       = "2006-01-02 15:04:05 -0700"
	HUMAN_DateTimeZoneFormat = "%d-%m-%y %H:%M -%Z"

	YYYYMMDD_HHMMSS_DateTimeFormat = "2006-01-02 15:04:05"

	DDMMYYY_HHMM_DateFormat       = "2-1-2006 15:04"
	HUMAN_DDMMYYY_HHMM_DateFormat = "%d-%m-%y %H:%M"
	MMDD_DateFormatForChart       = "Jan 02"

	RequestTimeOut = 20
	Get            = "GET"
	Post           = "POST"

	MaximumLogMinutesPerDay = 840

	Charset     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	EmailFormat = `\A([^@\s]+)@((?:[-a-z0-9]+\.)+[a-z]{2,})\z`
)
