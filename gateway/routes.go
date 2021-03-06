package main

import (
	"net/http"

	"github.com/philippfranke/mathub/services/assignment"
	"github.com/philippfranke/mathub/services/comment"
	"github.com/philippfranke/mathub/services/lecture"
	"github.com/philippfranke/mathub/services/search"
	"github.com/philippfranke/mathub/services/solution"
	"github.com/philippfranke/mathub/services/university"
	"github.com/philippfranke/mathub/services/user"
	"github.com/philippfranke/mathub/services/version"
)

// Route represents an entrypoint
type Route struct {
	path    string
	handler http.Handler
}

// Routes represents all api entrypoints
type Routes map[string]*Route

// MarshalJSON returns a JSON representation of Route
func (r *Route) MarshalJSON() ([]byte, error) {
	return []byte(`"` + r.path + `"`), nil
}

// Routes list all api entrypoints
func Entrypoints() Routes {
	return Routes{
		"universities_url":         &Route{"/unis", university.Router()},
		"university_url":           &Route{"/unis/{uni}", university.Router()},
		"users_url":                &Route{"/users", user.Router()},
		"user_url":                 &Route{"/users/{user}", user.Router()},
		"userLogin_url":            &Route{"/login", user.Router()},
		"lectures_url":             &Route{"/unis/{uni}/lectures", lecture.Router()},
		"lecture_url":              &Route{"/unis/{uni}/lectures/{lecture}", lecture.Router()},
		"versions_url":             &Route{"/unis/{uni}/lectures/{lecture}/{ref_type}/{ref_id}/versions", version.Router(*dataPath)},
		"version_url":              &Route{"/unis/{uni}/lectures/{lecture}/{ref_type}/{ref_id}/versions/{version}", version.Router(*dataPath)},
		"versions_user_url":        &Route{"/users/{user}/{ref_type}/{ref_id}/versions", version.Router(*dataPath)},
		"version_user_url":         &Route{"/users/{user}/{ref_type}/{ref_id}/versions/{version}", version.Router(*dataPath)},
		"assignments_url":          &Route{"/unis/{uni}/lectures/{lecture}/assignments", assignment.Router(*dataPath)},
		"assignment_url":           &Route{"/unis/{uni}/lectures/{lecture}/assignments/{assignment}", assignment.Router(*dataPath)},
		"solutions":                &Route{"/users/{uni}/solutions", solution.Router(*dataPath)},
		"solution_url":             &Route{"/users/{uni}/solutions/{solution}", solution.Router(*dataPath)},
		"commentTree_url":          &Route{"/comments/{refType}/{refId}", comment.Router()},
		"commentCreate_url":        &Route{"/comments", comment.Router()},
		"comment_url":              &Route{"/comments/{comment}", comment.Router()},
		"search_url":               &Route{"/search", search.Router()},
		"assignment_solutions_url": &Route{"/assignments/{assignment}/solutions", solution.Router(*dataPath)},
	}
}
