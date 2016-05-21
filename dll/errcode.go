package dll

const (
	ErrCode_InternalServer = 500

	ErrCode_MissParamToken = iota + 40010
	ErrCode_TokenExpired
	ErrCode_MissParamUid
	ErrCode_UidNotObjectId
	//
	ErrCode_UserMissParamKey
	ErrCode_UserKeyNotFound
	ErrCode_UserKeyUsed
	ErrCode_UserMissParamPhone
	ErrCode_UserPhoneNotMatch
	ErrCode_UserMissParamPassword
	//
	ErrCode_EventMissParamTitle
	ErrCode_EventMissParamDetail
	ErrCode_EventMissParamAddr
	ErrCode_EventMissParamStart
	ErrCode_EventMissParamPrice
	ErrCode_EventMissParamTotal
	ErrCode_EventDetailLenNotEnough
	ErrCode_UserAlreadySignUp
	ErrCode_UserNotSignUp
	ErrCode_EnrollmentFull
	ErrCode_EventNotCreateUser
	//
	ErrCode_MissParamEid
	ErrCode_EidNotObjectId

)
