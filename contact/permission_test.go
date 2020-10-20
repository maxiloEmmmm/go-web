package contact

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRole(t *testing.T) {
	mm, err := model.NewModelFromString(rbac)
	require.Nil(t, err)
	e, err := casbin.NewEnforcer(mm, fileadapter.NewAdapter("permission_test_policy.csv"))
	require.Nil(t, err)
	ok, err := e.Enforce("admin", "/sys/rbac", "POST")
	require.True(t, ok)
	ok, err = e.Enforce("admin", "/sys/rbac/1", "GET")
	require.True(t, ok)
	ok, err = e.Enforce("admin", "/sys/rbac", "GET")
	require.True(t, ok)
	ok, err = e.Enforce("admin", "/sys/rbac/1", "PATCH")
	require.True(t, ok)
	ok, err = e.Enforce("admin", "/sys/rbac/1", "DELETE")
	require.True(t, ok)
	ok, err = e.Enforce("admin1", "/sys/rbac/1", "DELETE")
	require.False(t, ok)
	ok, err = e.Enforce("haha", "/sys/rbac", "GET")
	require.True(t, ok)
}
