package gah

type Backend interface {
	ComparePasswords(hashedPwd string, plainPwd []byte) bool
	GetUserByEmail(email string) (UserStruct, error)
	GetUserByID(_id interface{}) (UserStruct, error)
	CreateUser(email string, password string) UserStruct
	InsertHashedLoginToken(guid string) string
	GetUserByToken(guid string, token string) (UserStruct, error)
}
