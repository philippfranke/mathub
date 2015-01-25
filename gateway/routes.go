package main

import (
	"net/http"

	"github.com/philippfranke/mathub/services/assignment"
	"github.com/philippfranke/mathub/services/comment"
	"github.com/philippfranke/mathub/services/lecture"
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
		"universities_url":  &Route{"/unis", university.Router()},
		"university_url":    &Route{"/unis/{uni}", university.Router()},
		"users_url":         &Route{"/users", user.Router()},
		"user_url":          &Route{"/users/{user}", user.Router()},
		"userLogin_url":     &Route{"/users/login", user.Router()},
		"lectures_url":      &Route{"/unis/{uni}/lectures", lecture.Router()},
		"lecture_url":       &Route{"/unis/{uni}/lectures/{lecture}", lecture.Router()},
		"assignments_url":   &Route{"/unis/{uni}/lectures/{lecture}/assignments", assignment.Router(*dataPath)},
		"assignment_url":    &Route{"/unis/{uni}/lectures/{lecture}/assignments/{assignment}", assignment.Router(*dataPath)},
		"commentTree_url":   &Route{"/comments/{refType}/{refId}", comment.Router()},
		"commentCreate_url": &Route{"/comments", comment.Router()},
		"comment_url":       &Route{"/comments/{comment}", comment.Router()},
		"versions_url":      &Route{"/unis/{uni}/lectures/{lecture}/assignments/{assignment}/versions", version.Router(*dataPath)},
		"version_url":       &Route{"/unis/{uni}/lectures/{lecture}/assignments/{assignment}/versions/{version}", version.Router(*dataPath)},
	}
}
