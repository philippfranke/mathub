package version

type Version struct {
	CommitHash    string `db:"commit_hash"`
	ReferenceType string `db:"ref_type"`
	ReferenceId   int64  `db:"ref_id"`
	UserId        int64  `db:"user_id"`
	Number        int64  `db:"version"`
}

func Create(v Version) error {
	_, err := DB.Exec("INSERT INTO versions (SELECT ?,?,?,0, IFNULL((max(version)+1),1)  FROM versions where ref_Type=? and ref_id=?)", v.CommitHash, v.ReferenceType, v.ReferenceId, v.ReferenceType, v.ReferenceId)

	if err != nil {
		return err
	}

	return nil
}
