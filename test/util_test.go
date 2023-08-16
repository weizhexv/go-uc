package test

import (
	"go-uc/internal/tool/maputil"
	"testing"
)

func TestMapUtils(t *testing.T) {
	mp := make(map[string]any)
	mp["a"] = 1
	mp["b"] = "false"
	mp["c"] = "fake"

	val, ok := maputil.GetInt64(mp, "a")
	t.Logf("val: %d, ok: %v", val, ok)
	str, ok := maputil.GetString(mp, "b")
	t.Logf("val: %s, ok: %v", str, ok)

	val, ok = maputil.GetInt64(mp, "c")
	t.Logf("val: %d, ok: %v", val, ok)
	bol, ok := maputil.GetBool(mp, "c")
	t.Logf("bol: %v, ok: %v", bol, ok)
}
