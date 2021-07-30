package contact

import (
	"context"
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/gin-gonic/gin"
	go_tool "github.com/maxiloEmmmm/go-tool"
	"github.com/maxiloEmmmm/go-web/ent"
	"github.com/maxiloEmmmm/go-web/ent/casbinrule"
	"strings"
)

const rbac = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act, eft

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow)) && !some(where (p.eft == deny))

[matchers]
m = g(r.sub, p.sub) && keyMatch2(r.obj, p.obj) && r.act == p.act
`

type EntPolicyAdapterOption struct {
	Db     *sql.DB
	Driver string
	Ctx    context.Context
}

type entPolicyAdapter struct {
	client *ent.Client
	ctx    context.Context
}

func NewEntPolicyAdapter(option *EntPolicyAdapterOption) (*entPolicyAdapter, error) {
	client := ent.NewClient(ent.Driver(entsql.OpenDB(option.Driver, option.Db)))
	err := client.Schema.Create(
		option.Ctx,
	)
	if err != nil {
		return nil, err
	}
	return &entPolicyAdapter{
		client: client,
		ctx:    option.Ctx,
	}, nil
}

func (adapter *entPolicyAdapter) LoadPolicy(model model.Model) error {
	rules := adapter.client.CasbinRule.Query().AllX(adapter.ctx)

	for _, rule := range rules {
		vs := make([]string, 0, 5)
		for _, v := range []string{rule.V0, rule.V1, rule.V2, rule.V3, rule.V4, rule.V5} {
			if v != "" {
				vs = append(vs, v)
			}
		}
		persist.LoadPolicyLine(go_tool.StringJoin(rule.PType, ",", strings.Join(vs, ",")), model)
	}

	return nil
}

func (adapter *entPolicyAdapter) setRule(create *ent.CasbinRuleCreate, pType string, rule []string) {
	create.SetPType(pType)
	rLen := len(rule)
	if rLen > 0 {
		create.SetV0(rule[0])
	}
	if rLen > 1 {
		create.SetV1(rule[1])
	}
	if rLen > 2 {
		create.SetV2(rule[2])
	}
	if rLen > 3 {
		create.SetV3(rule[3])
	}
	if rLen > 4 {
		create.SetV4(rule[4])
	}
	if rLen > 5 {
		create.SetV5(rule[5])
	}
}

func (adapter *entPolicyAdapter) SavePolicy(m model.Model) error {
	adapter.client.CasbinRule.Delete().ExecX(adapter.ctx)
	for _, kind := range []model.AssertionMap{m["p"], m["g"]} {
		for pType, ast := range kind {
			for _, rule := range ast.Policy {
				pipe := adapter.client.CasbinRule.Create()
				adapter.setRule(pipe, pType, rule)
				_, err := pipe.Save(adapter.ctx)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// AddPolicy adds a policy rule to the storage.
func (adapter *entPolicyAdapter) AddPolicy(sec string, pType string, rule []string) error {
	pipe := adapter.client.CasbinRule.Create()
	adapter.setRule(pipe, pType, rule)
	_, err := pipe.Save(adapter.ctx)
	return err
}

// RemovePolicy removes a policy rule from the storage.
func (adapter *entPolicyAdapter) RemovePolicy(sec string, pType string, rule []string) error {
	pipe := adapter.client.CasbinRule.Delete().Where(casbinrule.PType(pType))

	rLen := len(rule)
	if rLen > 0 {
		pipe.Where(casbinrule.V0(rule[0]))
	}
	if rLen > 1 {
		pipe.Where(casbinrule.V1(rule[1]))
	}
	if rLen > 2 {
		pipe.Where(casbinrule.V2(rule[2]))
	}
	if rLen > 3 {
		pipe.Where(casbinrule.V3(rule[3]))
	}
	if rLen > 4 {
		pipe.Where(casbinrule.V4(rule[4]))
	}
	if rLen > 5 {
		pipe.Where(casbinrule.V5(rule[5]))
	}
	_, err := pipe.Exec(adapter.ctx)
	return err
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (adapter *entPolicyAdapter) RemoveFilteredPolicy(sec string, pType string, fieldIndex int, fieldValues ...string) error {
	query := adapter.client.CasbinRule.Delete().
		Where(casbinrule.PType(pType))

	if fieldIndex <= 0 && 0 < fieldIndex+len(fieldValues) {
		query.Where(casbinrule.V0(fieldValues[0-fieldIndex]))
	}
	if fieldIndex <= 1 && 1 < fieldIndex+len(fieldValues) {
		query.Where(casbinrule.V1(fieldValues[1-fieldIndex]))
	}
	if fieldIndex <= 2 && 2 < fieldIndex+len(fieldValues) {
		query.Where(casbinrule.V2(fieldValues[2-fieldIndex]))
	}
	if fieldIndex <= 3 && 3 < fieldIndex+len(fieldValues) {
		query.Where(casbinrule.V3(fieldValues[3-fieldIndex]))
	}
	if fieldIndex <= 4 && 4 < fieldIndex+len(fieldValues) {
		query.Where(casbinrule.V4(fieldValues[4-fieldIndex]))
	}
	if fieldIndex <= 5 && 5 < fieldIndex+len(fieldValues) {
		query.Where(casbinrule.V5(fieldValues[5-fieldIndex]))
	}
	_, err := query.Exec(adapter.ctx)
	return err
}

var Permission *casbin.Enforcer

func InitPermission(m string, adapter interface{}) {
	if m == "" {
		m = rbac
	}

	mm, err := model.NewModelFromString(m)
	if err != nil {
		Error.Log("casbin", err.Error())
	}
	Permission, err = casbin.NewEnforcer(mm, adapter)
	if err != nil {
		Error.Log("casbin.enforcer", err.Error())
	}

	if Config.App.Mode != gin.ReleaseMode {
		Permission.EnableLog(true)
	}

	if err := Permission.LoadPolicy(); err != nil {
		Error.Log("casbin.load.policy", err.Error())
	}
}

const DefaultRoleUser = "_"

type PermissionRulePut struct {
	Add    []string
	Remove []string
}

func PermissionRoute(r gin.IRouter, prefix string) *gin.RouterGroup {
	if prefix == "" {
		prefix = "casbin"
	}
	cr := r.Group(prefix)

	cr.GET("/rule", GinHelpHandle(func(c *GinHelp) {
		c.Resource(Permission.GetPolicy())
	}))
	cr.POST("/rule", GinHelpHandle(func(c *GinHelp) {
		body := &RuleCreate{}
		c.InValidBind(&body)
		_, err := Permission.AddPolicy(strings.Split(body.Rule, ","))
		c.AssetsInValid("add.policy", err)
		c.ResourceCreate(nil)
	}))
	cr.PUT("/rule", GinHelpHandle(func(c *GinHelp) {
		body := &PermissionRulePut{}
		c.InValidBind(body)

		for _, add := range body.Add {
			rules := strings.Split(add, ",")
			Permission.AddPolicy(rules)
		}
		for _, remove := range body.Remove {
			Permission.RemovePolicy(strings.Split(remove, ","))
		}
		c.Resource(nil)
	}))
	cr.DELETE("/rule", GinHelpHandle(func(c *GinHelp) {
		// rule转义后 如: x,/x/:w/1,GET,allow 或 x%2C%2Fx%2F%3Aw%2F1%2CGET%2Callow 无法匹配
		// 所以用query 而不是param
		_, err := Permission.RemovePolicy(strings.Split(c.Query("rule"), ","))
		c.AssetsInValid("remove.policy", err)
		c.ResourceDelete()
	}))
	cr.GET("/role", GinHelpHandle(func(c *GinHelp) {
		c.Resource(Permission.GetAllRoles())
	}))
	cr.GET("/role/:role/user", GinHelpHandle(func(c *GinHelp) {
		uri := &struct {
			Role string `uri:"role"`
		}{}
		c.InValidBindUri(&uri)

		users, err := Permission.GetUsersForRole(uri.Role)
		c.AssetsInValid("get", err)
		c.Resource(go_tool.ArrayFilter(&users, func(d interface{}) bool {
			return d.(string) != DefaultRoleUser
		}))
	}))
	cr.POST("/role/:role/user/:user", GinHelpHandle(func(c *GinHelp) {
		uri := &struct {
			Role string `uri:"role"`
			User string `uri:"user"`
		}{}
		c.InValidBindUri(&uri)

		_, err := Permission.AddRoleForUser(uri.User, uri.Role)
		c.AssetsInValid("add", err)
		c.ResourceCreate(nil)
	}))
	cr.DELETE("/role/:role/user/:user", GinHelpHandle(func(c *GinHelp) {
		uri := &struct {
			Role string `uri:"role"`
			User string `uri:"user"`
		}{}
		c.InValidBindUri(&uri)

		_, err := Permission.RemoveGroupingPolicy(uri.User, uri.Role)
		c.AssetsInValid("remove", err)
		c.ResourceDelete()
	}))
	cr.POST("/role", GinHelpHandle(func(c *GinHelp) {
		body := &RoleCreate{}
		c.InValidBind(&body)

		// 加前缀 区分策略中的sub
		role := go_tool.StringJoin("role_", body.Role)
		has, err := Permission.HasRoleForUser(DefaultRoleUser, role)
		c.AssetsInValid("has.role", err)

		if !has {
			_, err = Permission.AddGroupingPolicy(DefaultRoleUser, role)
			c.AssetsInValid("add.group.policy", err)
		}
		c.ResourceCreate(nil)
	}))
	cr.DELETE("/role/:role", GinHelpHandle(func(c *GinHelp) {
		role := c.Param("role")
		has, err := Permission.HasRoleForUser("_", role)
		c.AssetsInValid("has.role", err)

		if has {
			_, err = Permission.DeleteRole(role)
			c.AssetsInValid("delete.role", err)
		}
		c.ResourceDelete()
	}))
	cr.GET("/user", GinHelpHandle(func(c *GinHelp) {
		c.Resource(Permission.GetAllSubjects())
	}))

	return cr
}
