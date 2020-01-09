package spinnakerpb

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// TODO(eac): we can get away with overriding the default json behavior for now because there's no expectation
//            that encoding/json produces roundtrippable jsonpb JSON
func (m Stage) MarshalJSON() ([]byte, error) { return serializeOneOf(reflect.ValueOf(m.Stage)) }
func (m Notification) MarshalJSON() ([]byte, error) {
	return serializeOneOf(reflect.ValueOf(m.Notification))
}
func (m Trigger) MarshalJSON() ([]byte, error) { return serializeOneOf(reflect.ValueOf(m.Trigger)) }

// This is grim. Spinnaker pipeline JSON serialization flattens type unions (notifications, webhooks, stages)
// and switches on the `type` field. jsonpb encodes the type union into an object, keyed on the type of the union.
// serializeOneOf reflects on a oneOf type and extracts the inner object (which always follows the pattern:
// <UnionType>_<SpecializedType>) to produce a flattened JSON structure.
func serializeOneOf(obj reflect.Value) ([]byte, error) {
	t := obj.Type()
	if t.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("serializeOneOf[%s]: must pass a pointer", t.String())
	}

	rt := t.Elem()
	n := rt.Name()
	p := strings.SplitN(n, "_", 2)

	if len(p) != 2 {
		return nil, fmt.Errorf("serializeOneOf[%s]: %s cannot be parsed", t.String(), n)
	}

	// TODO(eac): assert expected prefix name?

	fn := p[1]
	f, ok := rt.FieldByName(fn)
	if !ok {
		return nil, fmt.Errorf("serializeOneOf[%s]: field %s not found on type %s", t.String(), fn, n)
	}

	val := obj
	if t.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	rv := val.FieldByIndex(f.Index)
	return json.Marshal(rv.Interface())
}
