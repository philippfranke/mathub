package comment

import "time"
import "strconv"

type Comment struct {
	Id         int64     `json:"id" db:"id"`
	RefType    string    `json:"ref_type" db:"ref_type"`
	RefID      int64     `json:"ref_id" db:"ref_id"`
	RefVersion int64     `json:"ref_version" db:"ref_version"`
	RefLine    int64     `json:"ref_line" db:"ref_line"`
	ParentID   int64     `json:"parent_id" db:"parent_id"`
	UserID     int64     `json:"user_id" db:"user_id"`
	Timestamp  time.Time `json:"timestamp" db:"timestamp"`
	Text       string    `json:"text" db:"text"`
	Name       string    `json:"username" db:"name"`
}
type CommentNode struct {
	Node     Comment       `json:"comment"`
	Children []CommentNode `json:"children,omitempty"`
}
type Comments []Comment
type CommentTree []CommentNode

func All(refType, refId string) (CommentTree, error) {
	var comments Comments

	err := DB.Select(&comments, "SELECT comments.id, ref_type, ref_id, ref_version, ref_line, parent_id, user_id, timestamp, text, users.name FROM comments, users WHERE ref_type = ? AND ref_id =? AND user_id=users.id ORDER BY timestamp;", refType, refId)
	if err != nil {
		return nil, err
	}

	if len(comments) == 0 {
		return CommentTree{}, nil
	}

	commentTree := BuildCommentTree(0, comments)

	return commentTree, nil
}

func BuildCommentTree(parentId int64, comments Comments) CommentTree {
	commentTree := CommentTree{}
	for _, curComment := range comments {
		if curComment.ParentID == parentId {

			node := CommentNode{curComment, BuildCommentTree(curComment.Id, comments)}
			commentTree = append(commentTree, node)
		}
	}

	return commentTree
}

func Get(id string) (Comment, error) {
	var comment Comment
	err := DB.Get(&comment, "SELECT comments.id, ref_type, ref_id, ref_version, ref_line, parent_id, user_id, timestamp, text, users.name  FROM comments, users WHERE id = ? AND user_id=users.id;", id)
	if err != nil {
		return Comment{}, err
	}

	return comment, nil
}

func Create(comment Comment) (Comment, error) {
	res, err := DB.Exec("INSERT INTO comments (id, ref_type, ref_id, ref_version, ref_line, parent_id, user_id, timestamp, text) VALUES(NULL, ?, ?, ?, ?, ?, ?, ?,?);", comment.RefType, comment.RefID, comment.RefVersion, comment.RefLine, comment.ParentID, comment.UserID, comment.Timestamp, comment.Text)
	if err != nil {
		return Comment{}, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return Comment{}, err
	}
	comment.Id = lastId

	return comment, nil
}

func Update(comment Comment) error {
	_, err := DB.Exec("UPDATE comments SET ref_type=? ,ref_id=?, ref_version=?, ref_line=?, parent_id=? ,user_id=? ,timestamp=?, text=? WHERE id = ?;", comment.RefType, comment.RefID, comment.RefVersion, comment.RefLine, comment.ParentID, comment.UserID, comment.Timestamp, comment.Text, comment.Id)
	if err != nil {
		return err
	}
	return nil
}

func Destroy(id string) error {

	comment, err := Get(id)
	if err != nil {
		return err
	}
	var comments Comments

	err = DB.Select(&comments, "SELECT id, ref_type, ref_id, ref_version, ref_line, parent_id, user_id, timestamp, text FROM comments WHERE ref_type = ? AND ref_id =? ORDER BY timestamp;", comment.RefType, comment.RefID)
	if err != nil {
		return err
	}

	if len(comments) == 0 {
		return nil
	}

	var idInt int64
	idInt, err = strconv.ParseInt(id, 10, 64)

	if err != nil {
		return err
	}

	commentIDs := FindChildren(idInt, comments)
	commentIDs = append(commentIDs, idInt)

	_, err = DB.Exec("DELETE FROM comments WHERE id IN ? ;", commentIDs)

	stmt, err := DB.Prepare(`DELETE FROM comments WHERE id=?;`)

	for _, curCommentID := range commentIDs {
		_, err = stmt.Exec(curCommentID)
		if err != nil {
			return err
		}
	}

	return nil
}

func FindChildren(parentId int64, comments Comments) []int64 {
	commentIDs := []int64{}
	for _, curComment := range comments {
		if curComment.ParentID == parentId {

			commentIDs = append(commentIDs, curComment.Id)
			commentIDs = append(commentIDs, FindChildren(curComment.Id, comments)...)
		}
	}

	return commentIDs
}
