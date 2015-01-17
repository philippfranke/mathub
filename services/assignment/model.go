package assignment

import "time"

type Assignment struct {
	Id        int64     `json:"id"`
	LectureId int64     `json:"lecture_id" db:"lecture_id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	DueDate   time.Time `json:"due_date" db:"due_date"`
	Tex       string    `json:"tex"`

	CommitHash string `json:"-" db:"commit_hash"`
}

type Assignments []Assignment

func All(assignment string) (Assignments, error) {
	var assignments Assignments

	err := DB.Select(&assignments, "SELECT id, lecture_id, user_id, due_date, commit_hash, tex FROM assignments WHERE lecture_id = ? ORDER BY due_date;", assignment)
	if err != nil {
		return nil, err
	}

	if len(assignments) == 0 {
		return []Assignment{}, nil
	}

	return assignments, nil
}

func Get(id string) (Assignment, error) {
	var assignment Assignment
	err := DB.Get(&assignment, "SELECT id, lecture_id, user_id, due_date, commit_hash, tex FROM assignments WHERE id = ?;", id)
	if err != nil {
		return Assignment{}, err
	}

	return assignment, nil
}

func Create(assignment Assignment) (Assignment, error) {
	res, err := DB.Exec("INSERT INTO assignments (id, lecture_id, user_id, due_date, commit_hash,tex) VALUES(NULL, ?, 0, 0,0,?);", assignment.LectureId, assignment.Tex)
	if err != nil {
		return Assignment{}, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return Assignment{}, err
	}
	assignment.Id = lastId

	return assignment, nil
}

func GetLastId() (int64, error) {
	res, err := DB.Exec("SELECT null FROM assignments ORDER by id DESC LIMIT 1;")
	if err != nil {
		return 0, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastId, nil
}

func UpdateId(assignment Assignment) error {
	_, err := DB.Exec("UPDATE assignments SET commit_hash = ?, tex= WHERE id = ?;", assignment.CommitHash, assignment.Id)
	if err != nil {
		return err
	}
	return nil
}

func Update(assignment Assignment) error {
	_, err := DB.Exec("UPDATE assignments SET tex = ?, commit_hash = ? WHERE id = ?;", assignment.Tex, assignment.CommitHash, assignment.Id)
	if err != nil {
		return err
	}
	return nil
}

func Destroy(id string) error {
	_, err := DB.Exec("DELETE FROM assignments WHERE id = ?;", id)

	if err != nil {
		return err
	}

	return nil
}
