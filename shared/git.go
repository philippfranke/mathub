package shared

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Repo struct {
	dir      string
	DataPath string
}

func NewRepo(path string) *Repo {
	return &Repo{}
}

func (r *Repo) Open(path string) error {
	r.dir = filepath.Join(r.DataPath, path)
	return nil
}

func (r *Repo) Create(path string) error {
	r.dir = filepath.Join(r.DataPath, path)

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

func (r *Repo) Update(filename, tex string) error {
	path := filepath.Join(r.dir, filename)
	file, err := os.OpenFile(path, os.O_WRONLY, 0666)
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

func (r *Repo) Status() (string, error) {
	gitStatus := exec.Command("git", "status", "-s")
	gitStatus.Dir = r.dir
	out, err := gitStatus.CombinedOutput()
	log.Printf("gitStatus: %s", out)
	return string(out), err
}

func (r *Repo) Destroy(filename string) error {
	path := filepath.Join(r.dir, filename)
	gitRm := exec.Command("git", "rm", path)
	gitRm.Dir = r.dir
	out, err := gitRm.CombinedOutput()
	log.Printf("gitRm: %s", out)
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
	return strings.TrimSuffix(string(out), "\n")
}

func (r *Repo) ShowFile(filename, hash string) string {
	file := fmt.Sprintf("%s:%s", hash, filename)
	gitShow := exec.Command("git", "show", file)
	gitShow.Dir = r.dir
	out, _ := gitShow.CombinedOutput()
	log.Printf("gitShow: %s", out)
	return string(out)
}
