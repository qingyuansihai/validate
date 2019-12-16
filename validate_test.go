package validate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	v := New(map[string][]string{
		"age":  {"12"},
		"name": {"inhere"},
	})
	v.StringRules(MS{
		"age":  "required|strInt",
		"name": "required|string:3|strLen:4,6",
	})

	assert.True(t, v.Validate())
}

func TestRule_Apply(t *testing.T) {
	is := assert.New(t)
	mp := M{
		"name": "inhere",
		"code": "2363",
	}

	v := Map(mp)
	v.ConfigRules(MS{
		"name": `regex:\w+`,
	})
	v.AddRule("name", "stringLength", 3)
	v.StringRule("code", `required|regex:\d{4,6}`)

	is.True(v.Validate())
}

func TestStructUseDefault(t *testing.T) {
	is := assert.New(t)

	type user struct {
		Name string `validate:"required|default:tom" filter:"trim|upper"`
		Age  int    `validate:"uint|default:23"`
	}

	u := &user{Age: 90}
	v := New(u)
	is.True(v.Validate())
	is.Equal("tom", u.Name)

	u = &user{Name: "inhere"}
	v = New(u)
	is.True(v.Validate())
	// fmt.Println(v.SafeData())

	// check/filter default value
	u = &user{Age: 90}
	v = New(u)
	v.CheckDefault = true

	is.True(v.Validate())
	is.Equal("TOM", u.Name)
}
