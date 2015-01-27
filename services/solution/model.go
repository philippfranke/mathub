package solution

type Solution struct {
	Id           int64  `json:"id"`
	AssignmentId int64  `json:"assignment_id" db:"assignment_id"`
	UserId       int64  `json:"user_id" db:"user_id"`
	Tex          string `json:"tex"`

	CommitHash string `json:"-" db:"commit_hash"`
}

type Solutions []Solution

func All(user_id string) (Solutions, error) {
	var solutions Solutions

	err := DB.Select(&solutions, "SELECT id, assignment_id, user_id, commit_hash, tex FROM solutions WHERE user_id = ?;", user_id)
	if err != nil {
		return nil, err
	}

	if len(solutions) == 0 {
		return []Solution{}, nil
	}

	return solutions, nil
}

func AllByAssignment(assignment string) (Solutions, error) {
	var solutions Solutions

	err := DB.Select(&solutions, "SELECT id, assignment_id, user_id, commit_hash, tex FROM solutions WHERE assignment_id = ?;", assignment)
	if err != nil {
		return nil, err
	}

	if len(solutions) == 0 {
		return []Solution{}, nil
	}

	return solutions, nil
}

func Get(id string) (Solution, error) {
	var solution Solution
	err := DB.Get(&solution, "SELECT id, assignment_id, user_id, commit_hash, tex FROM solutions WHERE id = ?;", id)
	if err != nil {
		return Solution{}, err
	}

	return solution, nil
}

func Create(solution Solution) (Solution, error) {
	res, err := DB.Exec("INSERT INTO solutions (id, assignment_id, user_id, commit_hash,tex) VALUES(NULL, ?, ?, 0,?);", solution.AssignmentId, solution.UserId, solution.Tex)
	if err != nil {
		return Solution{}, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return Solution{}, err
	}
	solution.Id = lastId

	return solution, nil
}

func GetLastId() (int64, error) {
	res, err := DB.Exec("SELECT null FROM solutions ORDER by id DESC LIMIT 1;")
	if err != nil {
		return 0, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastId, nil
}

func UpdateId(solution Solution) error {
	_, err := DB.Exec("UPDATE solutions SET commit_hash = ? WHERE id = ?;", solution.CommitHash, solution.Id)

	if err != nil {
		return err
	}

	return err
}

func Update(solution Solution) error {
	_, err := DB.Exec("UPDATE solutions SET assignment_id = ?, tex = ?, commit_hash = ? WHERE id = ?;", solution.AssignmentId, solution.Tex, solution.CommitHash, solution.Id)
	if err != nil {
		return err
	}

	return err
}

func Destroy(id string) error {
	_, err := DB.Exec("DELETE FROM solutions WHERE id = ?;", id)

	if err != nil {
		return err
	}

	return nil
}
