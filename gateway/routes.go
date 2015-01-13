package main

import (
	"net/http"

	"github.com/philippfranke/mathub/services/assignment"
	"github.com/philippfranke/mathub/services/lecture"
	"github.com/philippfranke/mathub/services/university"
	"github.com/philippfranke/mathub/services/user"
)

// Route represents an entrypoint
type Route struct {
	path    string
	handler http.Handler
}

// MarshalJSON returns a JSON representation of Route
func (r *Route) MarshalJSON() ([]byte, error) {
	return []byte(`"` + r.path + `"`), nil
}

// Routes list all api entrypoints
var Routes = map[string]*Route{
	"universities_url": &Route{"/unis", university.Router()},
	"university_url":   &Route{"/unis/{uni}", university.Router()},
	"users_url":        &Route{"/users", user.Router()},
	"user_url":         &Route{"/users/{user}", user.Router()},
	"lectures_url":     &Route{"/unis/{uni}/lectures", lecture.Router()},

	"lecture_url":     &Route{"/unis/{uni}/lectures/{lecture}", lecture.Router()},
	"assignments_url": &Route{"/unis/{uni}/lectures/{lecture}/assignments", assignment.Router(*dataPath)},
	"assignment_url":  &Route{"/unis/{uni}/lectures/{lecture}/assignments/{assignment}", assignment.Router(*dataPath)},
}
