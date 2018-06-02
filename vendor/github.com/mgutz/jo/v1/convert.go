package jo

import (
	"fmt"

	"github.com/mgutz/to"
)

// Bool returns bool value from path.
func (n *Object) Bool(path string) (bool, error) {
	o, err := n.Get(path)
	if err != nil {
		return false, err
	}
	return to.Bool(o)
}

// AsBool returns bool value from path else false.
func (n *Object) AsBool(path string) bool {
	v, err := n.Bool(path)
	if err != nil {
		return false
	}
	return v
}

// MayBool should get value from path or return val.
func (n *Object) MayBool(path string, val bool) bool {
	v, err := n.Bool(path)
	if err != nil {
		return val
	}
	return v
}

// MustBool gets string value from path or panics.
func (n *Object) MustBool(path string) bool {
	v, err := n.Bool(path)
	if err != nil {
		panic(fmt.Sprintf(mustFormat, "bool"))
	}
	return v
}

// Float returns a float64 value from path.
func (n *Object) Float(path string) (float64, error) {
	o, err := n.Get(path)
	if err != nil {
		return 0, err
	}
	return to.Float64(o)
}

// AsFloat returns float64 value from path else 0.
func (n *Object) AsFloat(path string) float64 {
	v, err := n.Float(path)
	if err != nil {
		return 0
	}
	return v
}

// MayFloat should get value from path or return val.
func (n *Object) MayFloat(path string, val float64) float64 {
	v, err := n.Float(path)
	if err != nil {
		return val
	}
	return v
}

// MustFloat gets string value from path or panics.
func (n *Object) MustFloat(path string) float64 {
	v, err := n.Float(path)
	if err != nil {
		panic(fmt.Sprintf(mustFormat, "float64"))
	}
	return v
}

// Int returns an integer value from path.
func (n *Object) Int(path string) (int, error) {
	o, err := n.Get(path)
	if err != nil {
		return 0, err
	}

	n64, err := to.Int64(o)
	if err != nil {
		return 0, err
	}
	return int(n64), nil
}

// AsInt returns int value from path else 0.
func (n *Object) AsInt(path string) int {
	v, err := n.Int(path)
	if err != nil {
		return 0
	}
	return v
}

// MayInt should get value from path or return val.
func (n *Object) MayInt(path string, val int) int {
	v, err := n.Int(path)
	if err != nil {
		return val
	}
	return v
}

// MustInt gets string value from path or panics.
func (n *Object) MustInt(path string) int {
	v, err := n.Int(path)
	if err != nil {
		panic(fmt.Sprintf(mustFormat, "int"))
	}
	return v
}

// Int64 returns an integer value from path.
func (n *Object) Int64(path string) (int64, error) {
	o, err := n.Get(path)
	if err != nil {
		return 0, err
	}
	return to.Int64(o)
}

// AsInt64 returns int64 value from path else 0.
func (n *Object) AsInt64(path string) int64 {
	v, err := n.Int64(path)
	if err != nil {
		return 0
	}
	return v
}

// MayInt64 should get value from path or return val.
func (n *Object) MayInt64(path string, val int64) int64 {
	v, err := n.Int64(path)
	if err != nil {
		return val
	}
	return v
}

// MustInt64 gets string value from path or panics.
func (n *Object) MustInt64(path string) int64 {
	v, err := n.Int64(path)
	if err != nil {
		panic(fmt.Sprintf(mustFormat, "int"))
	}
	return v
}

// String returns a string value from path.
func (n *Object) String(path string) (string, error) {
	o, err := n.Get(path)
	if err != nil {
		return "", err
	}
	return to.String(o), nil
}

// AsString returns string value from path else "".
func (n *Object) AsString(path string) string {
	v, err := n.String(path)
	if err != nil {
		return ""
	}
	return v
}

// MayString should get value from path or return s.
func (n *Object) MayString(path string, s string) string {
	v, err := n.String(path)
	if err != nil {
		return s
	}
	return v
}

// MustString gets string value from path or panics.
func (n *Object) MustString(path string) string {
	v, err := n.String(path)
	if err != nil {
		panic(fmt.Sprintf(mustFormat, "string"))
	}
	return v
}

////////// Collections

// Map returns the map at path.
func (n *Object) Map(path string) (map[string]interface{}, error) {
	o, err := n.Get(path)
	if err != nil {
		return nil, err
	}
	switch rv := o.(type) {
	case map[string]interface{}:
		return rv, nil
	default:
		return nil, fmt.Errorf("%s is not a map: %q", path, o)
	}
}

// AsMap returns map  from path else nil.
func (n *Object) AsMap(path string) map[string]interface{} {
	v, err := n.Map(path)
	if err != nil {
		return nil
	}
	return v
}

// MayMap should get value from path or return val.
func (n *Object) MayMap(path string, val map[string]interface{}) map[string]interface{} {
	v, err := n.Map(path)
	if err != nil {
		return val
	}
	return v
}

// MustMap gets string value from path or panics.
func (n *Object) MustMap(path string) map[string]interface{} {
	v, err := n.Map(path)
	if err != nil {
		panic(fmt.Sprintf(mustFormat, "map[string]interface{}"))
	}
	return v
}

// Array returns slice of interface{} from path.
func (n *Object) Array(path string) ([]interface{}, error) {
	o, err := n.Get(path)
	if err != nil {
		return nil, err
	}
	switch rv := o.(type) {
	case []interface{}:
		return rv, nil
	default:
		return nil, fmt.Errorf("%s is not n Array: %q", path, o)
	}
}

// StringArray returns a string slice from path. Null values are converted
// to "".
func (n *Object) StringArray(path string) ([]string, error) {
	arr, err := n.Array(path)
	if err != nil {
		return nil, err
	}
	result := make([]string, len(arr))
	for i, val := range arr {
		if val == nil {
			result[i] = ""
		} else if s, ok := val.(string); ok {
			result[i] = s
		} else {
			return nil, fmt.Errorf("Array has non-string values")
		}
	}
	return result, nil
}

// AsSlice returns slice of interface{} from path.
func (n *Object) AsSlice(path string) []*Object {
	o, err := n.Get(path)
	if err != nil {
		return nil
	}

	switch rv := o.(type) {
	case []interface{}:
		results := []*Object{}
		for _, iface := range rv {
			results = append(results, &Object{iface})
		}
		return results
	default:
		return nil
	}
}

// // MayArray should get value from path or return val.
// func (n *Object) MayArray(path string, val []interface{}) []Object {
// 	v, err := n.Array(path)
// 	if err != nil {
// 		return val
// 	}
// 	return v
// }

// MustArray gets array value from path or panics.
func (n *Object) MustArray(path string) []interface{} {
	v, err := n.Array(path)
	if err != nil {
		panic(fmt.Sprintf(mustFormat, "[]interface{}"))
	}
	return v
}

// At returns a sub map at path.
func (n *Object) At(path string) (*Object, error) {
	o, err := n.Get(path)
	if err != nil {
		return nil, err
	}
	return &Object{o}, nil
}

// MustAt gets a sub map at path or panics.
func (n *Object) MustAt(path string) *Object {
	m, err := n.At(path)
	if err != nil {
		panic(fmt.Sprintf(mustFormat, "*Object"))
	}
	return m
}
