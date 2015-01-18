package lecture

type Lecture struct {
	Id           int64  `json:"id"`
	UniversityId int64  `json:"university_id" db:"university_id"`
	Name         string `json:"name"`
}

type Lectures []Lecture

func All(uni string) (Lectures, error) {
	var lectures Lectures

	err := DB.Select(&lectures, "SELECT id, university_id, name FROM lectures WHERE university_id = ? ORDER BY name;", uni)
	if err != nil {
		return nil, err
	}

	if len(lectures) == 0 {
		return []Lecture{}, nil
	}

	return lectures, nil
}

func Get(id string) (Lecture, error) {
	var lecture Lecture
	err := DB.Get(&lecture, "SELECT id, university_id, name FROM lectures WHERE id = ?;", id)
	if err != nil {
		return Lecture{}, err
	}

	return lecture, nil
}

func Create(lecture Lecture) (Lecture, error) {
	res, err := DB.Exec("INSERT INTO lectures (id, university_id, name) VALUES(NULL, ?, ?);", lecture.UniversityId, lecture.Name)
	if err != nil {
		return Lecture{}, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return Lecture{}, err
	}
	lecture.Id = lastId

	return lecture, nil
}

func Update(lecture Lecture) error {
	_, err := DB.Exec("UPDATE lectures SET university_id = ?, name = ?  WHERE id = ?;", lecture.UniversityId, lecture.Name, lecture.Id)
	if err != nil {
		return err
	}
	return nil
}

func Destroy(id string) error {
	_, err := DB.Exec("DELETE FROM lectures WHERE id = ?;", id)

	if err != nil {
		return err
	}

	return nil
}
