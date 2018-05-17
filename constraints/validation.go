// Copyright 2014 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package constraints

import (
	"fmt"
	"reflect"

	"github.com/juju/collections/set"
)

// Validator defines operations on constraints attributes which are
// used to ensure a constraints value is valid, as well as being able
// to handle overridden attributes.
type Validator interface {

	// RegisterConflicts is used to define cross-constraint override behaviour.
	// The red and blue attribute lists contain attribute names which conflict
	// with those in the other list.
	// When two constraints conflict:
	//  it is an error to set both constraints in the same constraints Value.
	//  when a constraints Value overrides another which specifies a conflicting
	//   attribute, the attribute in the overridden Value is cleared.
	RegisterConflicts(reds, blues []string)

	// RegisterUnsupported records attributes which are not supported by a constraints Value.
	RegisterUnsupported(unsupported []string)

	// RegisterVocabulary records allowed values for the specified constraint attribute.
	// allowedValues is expected to be a slice/array but is declared as interface{} so
	// that vocabs of different types can be passed in.
	RegisterVocabulary(attributeName string, allowedValues interface{})

	// Validate returns an error if the given constraints are not valid, and also
	// any unsupported attributes.
	Validate(cons Value) ([]string, error)

	// Merge merges cons into consFallback, with any conflicting attributes from cons
	// overriding those from consFallback.
	Merge(consFallback, cons Value) (Value, error)

	// UpdateVocabulary merges new attribute values with existing values.
	// This method does not overwrite or delete values, i.e.
	//     if existing values are {a, b}
	//     and new values are {c, d},
	//     then the merge result would be {a, b, c, d}.
	UpdateVocabulary(attributeName string, newValues interface{})
}

// NewValidator returns a new constraints Validator instance.
func NewValidator() Validator {
	return &validator{
		conflicts: make(map[string]set.Strings),
		vocab:     make(map[string][]interface{}),
	}
}

type validator struct {
	unsupported set.Strings
	conflicts   map[string]set.Strings
	vocab       map[string][]interface{}
}

// RegisterConflicts is defined on Validator.
func (v *validator) RegisterConflicts(reds, blues []string) {
	for _, red := range reds {
		v.conflicts[red] = set.NewStrings(blues...)
	}
	for _, blue := range blues {
		v.conflicts[blue] = set.NewStrings(reds...)
	}
}

// RegisterUnsupported is defined on Validator.
func (v *validator) RegisterUnsupported(unsupported []string) {
	v.unsupported = set.NewStrings(unsupported...)
}

// RegisterVocabulary is defined on Validator.
func (v *validator) RegisterVocabulary(attributeName string, allowedValues interface{}) {
	v.vocab[resolveAlias(attributeName)] = convertToSlice(allowedValues)
}

var checkIsCollection = func(coll interface{}) {
	k := reflect.TypeOf(coll).Kind()
	if k != reflect.Slice && k != reflect.Array {
		panic(fmt.Errorf("invalid vocab: %v of type %T is not a slice", coll, coll))
	}
}

var convertToSlice = func(coll interface{}) []interface{} {
	checkIsCollection(coll)
	var slice []interface{}
	val := reflect.ValueOf(coll)
	for i := 0; i < val.Len(); i++ {
		slice = append(slice, val.Index(i).Interface())
	}
	return slice
}

// UpdateVocabulary is defined on Validator.
func (v *validator) UpdateVocabulary(attributeName string, allowedValues interface{}) {
	attributeName = resolveAlias(attributeName)
	// If this attribute is not registered, delegate to RegisterVocabulary()
	currentValues, ok := v.vocab[attributeName]
	if !ok {
		v.RegisterVocabulary(attributeName, allowedValues)
	}

	unique := map[interface{}]bool{}
	writeUnique := func(all []interface{}) {
		for _, one := range all {
			unique[one] = true
		}
	}

	// merge existing values with new, ensuring uniqueness
	writeUnique(currentValues)
	newValues := convertToSlice(allowedValues)
	writeUnique(newValues)

	v.updateVocabularyFromMap(attributeName, unique)
}

func (v *validator) updateVocabularyFromMap(attributeName string, valuesMap map[interface{}]bool) {
	attributeName = resolveAlias(attributeName)
	var merged []interface{}
	for one, _ := range valuesMap {
		// TODO (anastasiamac) Because it's coming from the map, the order maybe affected
		// and can be unreliable. Not sure how to fix it yet...
		// How can we guarantee the order here?
		merged = append(merged, one)
	}
	v.RegisterVocabulary(attributeName, merged)
}

// checkConflicts returns an error if the constraints Value contains conflicting attributes.
func (v *validator) checkConflicts(cons Value) error {
	attrValues := cons.attributesWithValues()
	attrSet := make(set.Strings)
	for attrTag := range attrValues {
		attrSet.Add(attrTag)
	}
	for _, attrTag := range attrSet.SortedValues() {
		conflicts, ok := v.conflicts[attrTag]
		if !ok {
			continue
		}
		for _, conflict := range conflicts.SortedValues() {
			if attrSet.Contains(conflict) {
				return fmt.Errorf("ambiguous constraints: %q overlaps with %q", attrTag, conflict)
			}
		}
	}
	return nil
}

// checkUnsupported returns any unsupported attributes.
func (v *validator) checkUnsupported(cons Value) []string {
	return cons.hasAny(v.unsupported.Values()...)
}

// checkValidValues returns an error if the constraints value contains an
// attribute value which is not allowed by the vocab which may have been
// registered for it.
func (v *validator) checkValidValues(cons Value) error {
	for attrTag, attrValue := range cons.attributesWithValues() {
		k := reflect.TypeOf(attrValue).Kind()
		if k == reflect.Slice || k == reflect.Array {
			// For slices we check that all values are valid.
			val := reflect.ValueOf(attrValue)
			for i := 0; i < val.Len(); i++ {
				if err := v.checkInVocab(attrTag, val.Index(i).Interface()); err != nil {
					return err
				}
			}
		} else {
			if err := v.checkInVocab(attrTag, attrValue); err != nil {
				return err
			}
		}
	}
	return nil
}

// checkInVocab returns an error if the attribute value is not allowed by the
// vocab which may have been registered for it.
func (v *validator) checkInVocab(attributeName string, attributeValue interface{}) error {
	validValues, ok := v.vocab[resolveAlias(attributeName)]
	if !ok {
		return nil
	}
	for _, validValue := range validValues {
		if coerce(validValue) == coerce(attributeValue) {
			return nil
		}
	}
	return fmt.Errorf(
		"invalid constraint value: %v=%v\nvalid values are: %v", attributeName, attributeValue, validValues)
}

// coerce returns v in a format that allows constraint values to be easily
// compared. Its main purpose is to cast all numeric values to float64 (since
// the numbers we compare are generated from json serialization).
func coerce(v interface{}) interface{} {
	switch val := v.(type) {
	case string:
		return v
	// Yes, these are all the same, however we can't put them into a single
	// case, or the value becomes interface{}, which can't be converted to a
	// float64.
	case int:
		return float64(val)
	case int8:
		return float64(val)
	case int16:
		return float64(val)
	case int32:
		return float64(val)
	case int64:
		return float64(val)
	case uint:
		return float64(val)
	case uint8:
		return float64(val)
	case uint16:
		return float64(val)
	case uint32:
		return float64(val)
	case uint64:
		return float64(val)
	case float32:
		return float64(val)
	case float64:
		return val
	}
	return v
}

// withFallbacks returns a copy of v with nil values taken from vFallback.
func withFallbacks(v Value, vFallback Value) Value {
	vAttr := v.attributesWithValues()
	fbAttr := vFallback.attributesWithValues()
	for k, v := range fbAttr {
		if _, ok := vAttr[k]; !ok {
			vAttr[k] = v
		}
	}
	return fromAttributes(vAttr)
}

// Validate is defined on Validator.
func (v *validator) Validate(cons Value) ([]string, error) {
	unsupported := v.checkUnsupported(cons)
	if err := v.checkConflicts(cons); err != nil {
		return unsupported, err
	}
	if err := v.checkValidValues(cons); err != nil {
		return unsupported, err
	}
	return unsupported, nil
}

// Merge is defined on Validator.
func (v *validator) Merge(consFallback, cons Value) (Value, error) {
	// First ensure both constraints are valid. We don't care if there
	// are constraint attributes that are unsupported.
	if _, err := v.Validate(consFallback); err != nil {
		return Value{}, err
	}
	if _, err := v.Validate(cons); err != nil {
		return Value{}, err
	}
	// Gather any attributes from consFallback which conflict with those on cons.
	attrValues := cons.attributesWithValues()
	var fallbackConflicts []string
	for attrTag := range attrValues {
		fallbackConflicts = append(fallbackConflicts, v.conflicts[attrTag].Values()...)
	}
	// Null out the conflicting consFallback attribute values because
	// cons takes priority. We can't error here because we
	// know that aConflicts contains valid attr names.
	consFallbackMinusConflicts := consFallback.without(fallbackConflicts...)
	// The result is cons with fallbacks coming from any
	// non conflicting consFallback attributes.
	return withFallbacks(cons, consFallbackMinusConflicts), nil
}
