package database

type userData struct {
	login string
	password string
	logStatus bool
}
var user userData
func SetUser(l string, p string) {
	user.login = l
	user.password = p
	user.logStatus = true

}
func IfLogged() bool {
	return user.logStatus
}