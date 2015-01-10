package assignment

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const BasePath = "/Users/philiqq/Projects/go/src/github.com/philippfranke/mathub/repos"

type Repo struct {
	uni     string
	lecture string
	dir     string
}

func (r *Repo) Create() error {
	r.dir = filepath.Join(BasePath, r.uni, r.lecture)

	if err := os.MkdirAll(r.dir, 0755); err != nil {
		return err
	}

	gitDir := filepath.Join(r.dir, ".git")
	_, err := os.Stat(gitDir)
	if os.IsNotExist(err) {
		gitInit := exec.Command("git", "init", r.dir)
		gitInit.Dir = r.dir
		out, err := gitInit.CombinedOutput()
		log.Printf("GitInit: %s", out)
		return err
	}
	return nil
}

func (r *Repo) Add(filename, tex string) error {
	path := filepath.Join(r.dir, filename)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	file.WriteString(tex)

	gitAdd := exec.Command("git", "add", path)
	gitAdd.Dir = r.dir
	out, err := gitAdd.CombinedOutput()
	log.Printf("gitAdd: %s", out)
	return err
}

func (r *Repo) Commit(message string, author string) error {
	gitCommit := exec.Command("git", "commit", "-m", message, "--author", author)
	gitCommit.Dir = r.dir
	out, err := gitCommit.CombinedOutput()
	log.Printf("gitCommit: %s", out)
	return err
}

func (r *Repo) LastHash() string {
	gitRev := exec.Command("git", "rev-parse", "--verify", "HEAD")
	gitRev.Dir = r.dir
	out, _ := gitRev.CombinedOutput()
	log.Printf("gitRev: %s", out)
	return string(out)
}
