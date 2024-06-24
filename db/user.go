package db

type User struct {
	ID   int
	Name string
}

func CreateUser(name string) error {
	stmt, err := DB.Prepare("INSERT INTO users(name) VALUES(?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(name)
	return err
}

func GetUser(id int) (*User, error) {
	var user User
	err := DB.QueryRow("SELECT id, name FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUsers() ([]*User, error) {
	var users []*User
	rows, err := DB.Query("SELECT id, name FROM users")
	if err != nil {
		return []*User{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return []*User{}, err
		}
		users = append(users, &user)
	}
	return users, nil
}
