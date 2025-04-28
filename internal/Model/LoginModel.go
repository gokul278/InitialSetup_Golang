package model

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type AdminLoginModel struct {
	UserId        int    `json:"refUserId" gorm:"column:refUserId"`
	CustUserId    string `json:"refUserCustId" gorm:"column:refUserCustId"`
	RId           int    `json:"refRTId" gorm:"column:refRTId"`
	UserFirstName string `json:"refUserFirstName" gorm:"column:refUserFirstName"`
	UserLastName  string `json:"refUserLastName" gorm:"column:refUserLastName"`
	RTName        string `json:"refRTName" gorm:"column:refRTName"`
	ADHashPass    string `json:"refADHashPass" gorm:"column:refADHashPass"`
	CODOPhoneNo1  string `json:"refCODOPhoneNo1" gorm:"column:refCODOPhoneNo1"`
	CODOEmail     string `json:"refCODOEmail" gorm:"column:refCODOEmail"`
}

type RefTransHistory struct {
	TransTypeId int    `json:"transTypeId" gorm:"column:transTypeId"`
	THData      string `json:"refTHData" gorm:"column:refTHData"`
	UserId      int    `json:"refUserId" gorm:"column:refUserId"`
	THActionBy  int    `json:"refTHActionBy" gorm:"column:refTHActionBy"`
}

func (RefTransHistory) TableName() string {
	return "aduit.refTransHistory"
}
