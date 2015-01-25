package user

type User struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email" db:"email"`
	PasswordHash string `json:"password_hash,omitempty" db:"password_hash"`
}

type Users []User

func All() (Users, error) {
	var users Users

	err := DB.Select(&users, "SELECT id, name, email FROM users ORDER BY name;")
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return []User{}, nil
	}

	return users, nil
}

func Get(id string) (User, error) {
	var user User
	err := DB.Get(&user, "SELECT id, name, email FROM users WHERE id = ? ;", id)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func GetByEmail(email string) (User, error) {
	var user User
	err := DB.Get(&user, "SELECT id, name, email, password_hash FROM users WHERE email = ? ;", email)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func Create(user User) (User, error) {
	res, err := DB.Exec("INSERT INTO users (id, name, email, password_hash) VALUES(NULL, ?, ?, ?);", user.Name, user.Email, user.PasswordHash)
	if err != nil {
		return User{}, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return User{}, err
	}
	user.Id = lastId
	user.PasswordHash = ""

	return user, nil
}

func Update(user User) error {
	_, err := DB.Exec("UPDATE users SET name = ? AND email = ? AND password_hash = ? WHERE id = ?;", user.Name, user.Email, user.PasswordHash, user.Id)
	if err != nil {
		return err
	}
	return nil
}

func Destroy(id string) error {
	_, err := DB.Exec("DELETE FROM users WHERE id = ?;", id)

	if err != nil {
		return err
	}

	return nil
}
