package teleport

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type Teleport struct {
	env Environment
}

func New() *Teleport {
	return &Teleport{}
}

func (t *Teleport) GetEnv() Environment {
	return t.env
}

func (t *Teleport) SetEnv(_env string) {
	env, err := NewEnv(_env)
	if err != nil {
		panic(err)
	}
	cmd := exec.Command("tsh", "login", string(env))
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	t.env = env
}

func (t *Teleport) ListDb(dbNames ...string) DatabaseListResult {
	cmd := exec.Command("tsh", "db", "ls", "--format=json")
	buf := bytes.NewBuffer(make([]byte, 0))
	cmd.Stdout = buf
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	var databaseListResult DatabaseListResult
	err = json.Unmarshal(buf.Bytes(), &databaseListResult)
	if err != nil {
		panic(err)
	}

	if len(dbNames) == 0 {
		return databaseListResult
	}

	filteredDbList := make(DatabaseListResult, 0)
	_regex := strings.Join(dbNames, "|")
	for _, db := range databaseListResult {
		matched, _ := regexp.MatchString(_regex, db.Metadata.Name)
		if matched {
			filteredDbList = append(filteredDbList, db)
		}
	}
	return filteredDbList
}

func (t *Teleport) DbConnect(dbName string) {
	cmd := exec.Command("tsh", "db", "login",
		"--db-user", "teleport_admin",
		"--db-name", strings.ReplaceAll(dbName, "-", "_"), dbName,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		panic(err)
	}

}

func (t *Teleport) DbConfig(dbName string) DatabaseConfigResult {
	cmd := exec.Command("tsh", "db", "config", "--format=json", dbName)
	buf := bytes.NewBuffer(make([]byte, 0))
	cmd.Stdout = buf
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	var DatabaseConfigResult DatabaseConfigResult
	err = json.Unmarshal(buf.Bytes(), &DatabaseConfigResult)
	if err != nil {
		panic(err)
	}
	return DatabaseConfigResult
}
