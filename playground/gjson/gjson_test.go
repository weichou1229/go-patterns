package gjson

import (
	"github.com/tidwall/gjson"
	"testing"
)

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
const json2 = `{
  "programmers": [
    {
      "firstName": "Janet", 
      "lastName": "McLaughlin", 
    }, {
      "firstName": "Elliotte", 
      "lastName": "Hunter", 
    }, {
      "firstName": "Jason", 
      "lastName": "Harold", 
    }
  ]
}`

func TestGjson(t *testing.T) {
	println(gjson.Get(json, "name.last").String())
	println(gjson.Get(json, "name").String())
	println(gjson.Get(json, "age").Int())
	println(gjson.Get(json, "age").String())
}

func TestGjson2(t *testing.T) {
	println(gjson.Get(json2, "programmers.#.lastName").String())
}
