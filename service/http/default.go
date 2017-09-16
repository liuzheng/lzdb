package http

// HTTP Code 2xx
const (
	HTTPCodeOK = 200 + iota
	HTTPCodeCreated // 201
	HTTPCodeAccepted // 202
	HTTPCodeNonAuthoritativeInformation // 203
	HTTPCodeNoContent  // 204
	HTTPCodeResetContent //  205
	HTTPCodePartialContent // 206
)

// HTTP Code 3xx
const (
	HTTPCodeMultipleChoices = 300 + iota
	HTTPCodeMovedPermanently // 301
	HTTPCodeFound // 302
	HTTPCodeSeeOther // 303
	HTTPCodeNotModified // 304
	HTTPCodeUseProxy // 305
	HTTPCodeSwitchProxy // 306
	HTTPCodeTemporaryRedirect // 307
	HTTPCodePermanentRedirect // 308

)

// HTTP Code 4xx
const (
	HTTPCodeBadRequest = 400 + iota
	HTTPCodeUnauthorized  // 401
	HTTPCodePaymentRequired // 402
	HTTPCodeForbidden // 403
	HTTPCodeNotFound // 404
	HTTPCodeMethodNotAllowed // 405
	HTTPCodeNotAcceptable // 406
	HTTPCodeProxyAuthenticationRequired // 407
	HTTPCodeRequestTimeout // 408
	HTTPCodeConflict // 409
	HTTPCodeGone // 410
	HTTPCodeLengthRequired // 411
	HTTPCodePreconditionFailed // 412
	HTTPCodeRequestEntityTooLarge // 413
	HTTPCodeRequestURITooLong // 414
	HTTPCodeUnsupportedMediaType // 415
	HTTPCodeRequestedRangeNotSatisfiable // 416
	HTTPCodeExpectationFailed // 417
	HTTPCodeImateapot // 418
	_
	_
	HTTPCodeMisdirectedRequest // 421
	HTTPCodeUnprocessableEntity // 422
	HTTPCodeLocked // 423
	HTTPCodeFailedDependency // 424
	_
	HTTPCodeUpgradeRequired // 426
	_
	HTTPCodePreconditionRequired // 428
	HTTPCodeTooManyRequests // 429
	_
	HTTPCodeRequestHeaderFieldsTooLarge // 431
	HTTPCodeUnavailableForLegalReasons = 451
)


// HTTP Code 5xx
const (
	HTTPCodeInternalServerError = 500 + iota // 500
	HTTPCodeNotImplemented  // 501
	HTTPCodeBadGateway // 502
	HTTPCodeServiceUnavailable  // 503
	HTTPCodeGatewayTimeout  // 504
	HTTPCodeHTTPVersionNotSupported  // 505
	HTTPCodeVariantAlsoNegotiates  // 506
	HTTPCodeInsufficientStorage  // 507
	HTTPCodeLoopDetected  // 508
	_
	HTTPCodeNotExtended  // 510
	HTTPCodeNetworkAuthenticationRequired  // 511
)

