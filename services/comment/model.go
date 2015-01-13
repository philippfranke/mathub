package comment

import "time"

type Comment struct {
	Id        int64     `json:"id"`
	RefType   string    `json:"ref_type" db:"ref_type"`
	RefID     int64     `json:"ref_id" db:"ref_id"`
	ParentID  int64     `json:"parent_id" db:"parent_id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
	Text      string    `json:"text" db:"text"`
}

type Comments []Comment

func All(assignmentId string) (Comments, error) {
	var comments Comments

	err := DB.Select(&comments, "SELECT id, ref_type, ref_id, parent_id, user_id, timestamp, text FROM comments WHERE parent_id = ? ORDER BY timestamp;", assignmentId)
	if err != nil {
		return nil, err
	}

	if len(comments) == 0 {
		return Comments{}, nil
	}

	return comments, nil
}
