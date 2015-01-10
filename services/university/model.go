package university

type University struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Universities []University

func All() (Universities, error) {
	var unis Universities

	err := DB.Select(&unis, "SELECT id, name FROM universities ORDER BY name;")
	if err != nil {
		return nil, err
	}

	if len(unis) == 0 {
		return []University{}, nil
	}

	return unis, nil
}

func Get(id string) (University, error) {
	var uni University
	err := DB.Get(&uni, "SELECT id, name FROM universities WHERE id = ?;", id)
	if err != nil {
		return University{}, err
	}

	return uni, nil
}

func Create(uni University) (University, error) {
	res, err := DB.Exec("INSERT INTO universities (id, name) VALUES(NULL, ?);", uni.Name)
	if err != nil {
		return University{}, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return University{}, err
	}
	uni.Id = lastId

	return uni, nil
}

func Update(uni University) error {
	_, err := DB.Exec("UPDATE universities SET name = ? WHERE id = ?;", uni.Name, uni.Id)
	if err != nil {
		return err
	}
	return nil
}

func Destroy(id string) error {
	_, err := DB.Exec("DELETE FROM universities WHERE id = ?;", id)

	if err != nil {
		return err
	}

	return nil
}
