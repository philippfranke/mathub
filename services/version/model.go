package version

type Version struct {
	CommitHash    string `db:"commit_hash"`
	ReferenceType string `json:"-,omitempty" db:"ref_type"`
	ReferenceId   int64  `json:"-,omitempty" db:"ref_id"`
	UserId        int64  `db:"user_id"`
	Number        int64  `db:"version"`
	Tex           string `json:"tex,omitempty"`
}

type Versions []Version

func All(refId, refType string) (Versions, error) {
	var versions Versions
	err := DB.Select(&versions, "SELECT commit_hash, ref_id, ref_type, user_id, version FROM versions WHERE ref_Type=? and ref_id=?", refType, refId)
	if err != nil {
		return nil, err
	}

	if len(versions) == 0 {
		return Versions{}, nil
	}

	return versions, nil
}

func Get(refId, refType, number string) (Version, error) {
	var version Version
	err := DB.Get(&version, "SELECT commit_hash, ref_id, ref_type, user_id, version FROM versions WHERE ref_id=? and ref_type=? and version=?", refId, refType, number)
	if err != nil {
		return Version{}, err
	}
	return version, nil
}

func Create(v Version) error {
	_, err := DB.Exec("INSERT INTO versions (SELECT ?,?,?,0, IFNULL((max(version)+1),1)  FROM versions where ref_Type=? and ref_id=?)", v.CommitHash, v.ReferenceType, v.ReferenceId, v.ReferenceType, v.ReferenceId)

	if err != nil {
		return err
	}

	return nil
}
