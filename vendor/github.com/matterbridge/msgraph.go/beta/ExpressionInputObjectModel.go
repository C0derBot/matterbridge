// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ExpressionInputObject undocumented
type ExpressionInputObject struct {
	// Object is the base model of ExpressionInputObject
	Object
	// Definition undocumented
	Definition *ObjectDefinition `json:"definition,omitempty"`
	// Properties undocumented
	Properties []StringKeyObjectValuePair `json:"properties,omitempty"`
}