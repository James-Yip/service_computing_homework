package entities

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {
	_, err := engine.Insert(u)
	checkErr(err)
	return nil
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	var infoList []UserInfo

	rows, err := engine.Rows(new(UserInfo))
	defer rows.Close()
	checkErr(err)
	userInfo := new(UserInfo)
	for rows.Next() {
		err = rows.Scan(userInfo)
		infoList = append(infoList, *userInfo)
	}
	return infoList
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	userInfo := new(UserInfo)
	_, err := engine.Id(id).Get(userInfo)
	checkErr(err)
	return userInfo
}
