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
	//return []userinfo
	var userlist []UserInfo
	rows, err := engine.Rows(new(UserInfo))
	checkErr(err)
	//select * from user
	defer rows.Close()
	bean := new(UserInfo)
	for rows.Next() {
		err = rows.Scan(bean)
		//add row to list
		userlist = append(userlist, *bean)
	}
	return userlist
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	u := new(UserInfo)
	_, err := engine.Where("id = ?", id).Get(u)
	checkErr(err)
	return u
}
