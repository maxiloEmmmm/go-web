package test

import (
	"bytes"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "golang.org/x/tools/go/packages"
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestGen(t *testing.T) {
	cmd := exec.Command("go", "run", "../generate", "./ent/schema")
	dir, _ := os.Getwd()
	cmd.Dir = dir
	stderr := bytes.NewBuffer(nil)
	stdout := bytes.NewBuffer(nil)
	cmd.Stderr = stderr
	cmd.Stdout = stdout
	require.NoError(t, cmd.Run(), stderr.String(), stdout.String())
}

func TestMapArray(t *testing.T) {
	x := make(map[string][]string, 0)
	assert.Len(t, x["?"], 0)
}

func TestCasbin(t *testing.T) {
	m, err := model.NewModelFromString(`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
`)
	if err != nil {
		log.Fatalf("error: model: %s", err)
	}

	e, err := casbin.NewEnforcer(m)
	if err != nil {
		log.Fatalf("error: enforcer: %s", err)
	}
	rule := []string{"role1", "obj", "act"}
	e.AddPolicy(rule)
	//that's ok
	rule[0] = "role2"
	//adapter data ok
	//e.model["p"]["p"].Policy is invaild, role1,obj,act not exist role2,obj,act exist
	pass, _ := e.Enforce("role1", "obj", "act")
	require.Equal(t, pass, false)
}