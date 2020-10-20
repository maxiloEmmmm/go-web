package contact

import (
	"context"
	"database/sql"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	entsql "github.com/facebook/ent/dialect/sql"
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
	Db  *sql.DB
	Ctx context.Context
}

type entPolicyAdapter struct {
	client *ent.Client
	ctx    context.Context
}

func NewEntPolicyAdapter(option *EntPolicyAdapterOption) (*entPolicyAdapter, error) {
	client := ent.NewClient(ent.Driver(entsql.OpenDB(DbEngine, option.Db)))
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

func (adapter *entPolicyAdapter) SavePolicy(m model.Model) error {
	adapter.client.CasbinRule.Delete().ExecX(adapter.ctx)
	for _, kind := range []model.AssertionMap{m["p"], m["g"]} {
		for pType, ast := range kind {
			for _, rule := range ast.Policy {
				_, err := adapter.client.CasbinRule.Create().
					SetPType(pType).
					SetV0(rule[0]).
					SetV1(rule[1]).
					SetV2(rule[2]).
					SetV3(rule[3]).
					SetV4(rule[4]).
					SetV5(rule[5]).Save(adapter.ctx)
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
	_, err := adapter.client.CasbinRule.Create().
		SetPType(pType).
		SetV0(rule[0]).
		SetV1(rule[1]).
		SetV2(rule[2]).
		SetV3(rule[3]).
		SetV4(rule[4]).
		SetV5(rule[5]).Save(adapter.ctx)
	return err
}

// RemovePolicy removes a policy rule from the storage.
func (adapter *entPolicyAdapter) RemovePolicy(sec string, pType string, rule []string) error {
	_, err := adapter.client.CasbinRule.Delete().
		Where(casbinrule.And(
			casbinrule.PType(pType),
			casbinrule.V0(rule[0]),
			casbinrule.V1(rule[1]),
			casbinrule.V2(rule[2]),
			casbinrule.V3(rule[3]),
			casbinrule.V4(rule[4]),
			casbinrule.V5(rule[5]),
		)).Exec(adapter.ctx)
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

	if err := Permission.LoadPolicy(); err != nil {
		Error.Log("casbin.load.policy", err.Error())
	}
}

func PermissionRoute(desc string, path string, handlers ...gin.HandlerFunc) (string, []gin.HandlerFunc) {

	return path, handlers
}
