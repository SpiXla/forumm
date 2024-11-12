package database

func IsEmailExist(email string) bool {
	query := `SELECT COUNT(*) FROM users WHERE email = ?`
	var count int
	err := Db.QueryRow(query, email).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func IsUnameExist(username string) bool {
	query := `SELECT COUNT(*) FROM users WHERE username = ?`
	var count int
	err := Db.QueryRow(query, username).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func CheckPswd(email string) (string, error) {
	var pswd string
	query := `SELECT password FROM users WHERE email = ?`
	err := Db.QueryRow(query, email).Scan(&pswd)
	if err != nil {
		return "", err
	}
	return pswd, nil
}

func CheckUname(email string) string {
	var uname string
	query := `SELECT username FROM users WHERE email = ?`
	err := Db.QueryRow(query, email).Scan(&uname)
	if err != nil {
		return ""
	}
	return uname
}
