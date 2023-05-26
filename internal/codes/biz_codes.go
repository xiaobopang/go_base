package codes

// http status, bisiness code, message
var (
	CodeOK          = New(200, 200, "success")
	CodePartSuccess = New(202, 202, "part success")

	CodePermissionDenied = New(401, 401, "authentication failed")
	CodeNotAuthorized    = New(403, 403, "resource is not authorized")
	CodeNotFound         = New(404, 404, "resource does not exist")
	CodeValidationFailed = New(400, 400, "")
	CodeNotAvailable     = New(400, 400, "not available")

	CodeInternal = New(500, 500, "server error")
)
