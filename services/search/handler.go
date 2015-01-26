package search

import (
	"net/http"

	. "github.com/philippfranke/mathub/shared"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) error {
	results, err := Search(r.URL.Query().Get("query"))
	if err != nil {
		return err
	}

	return WriteJSON(w, results)
}

type SearchItem struct {
	Type           string `json:"type,omitempty" db:"tablename"`
	AssignmentId   int64  `json:"assignment_id,omitempty" db:"assignment_id"`
	AssignmentName string `json:"assignment_name,omitempty" db:"assignment_name"`
	SolutionId     int64  `json:"solution_id,omitempty" db:"solution_id"`
	UniversityId   int64  `json:"university_id,omitempty" db:"university_id"`
	UniversityName string `json:"university_name,omitempty" db:"university_name"`
	LectureId      int64  `json:"lecture_id,omitempty" db:"lecture_id"`
	LectureName    string `json:"lecture_name,omitempty" db:"lecture_name"`
	UserId         int64  `json:"user_id,omitempty" db:"user_id"`
	UserName       string `json:"user_name,omitempty" db:"user_name"`
	Tex            string `json:"tex,omitempty" db:"tex"`
}

type SearchResults []SearchItem

func Search(query string) (SearchResults, error) {
	var results SearchResults

	searchSQL := `SELECT DISTINCT tablename,
                user_id,
                IFNULL(users.name,"") AS user_name,
                assignment_id,
                assignment_name,
                solution_id,
                lecture_id,
                lectures.name AS lecture_name,
                universities.id AS university_id,
                universities.name AS university_name,
                tex
FROM
  (SELECT "solutions" AS tablename,
          solutions.id AS solution_id,
          solutions.user_id,
          assignment_id,
          lecture_id,
          assignments.name AS assignment_name,
          solutions.tex
   FROM solutions
   LEFT JOIN assignments ON assignment_id = assignments.id
   UNION SELECT "assignments" AS tablename,
                0 AS solution_id,
                user_id,
                assignments.id AS assignment_id,
                lecture_id,
                name AS assignment_name,
                assignments.tex
   FROM assignments) searchTable
LEFT JOIN lectures ON lectures.Id = lecture_id
LEFT JOIN users ON users.Id = user_id
LEFT JOIN universities ON universities.id = lectures.university_id
WHERE tex LIKE CONCat("%", ?,"%");`

	err := DB.Select(&results, searchSQL, query)
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return SearchResults{}, nil
	}

	return results, nil
}
