// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// InferenceClassificationOverride undocumented
type InferenceClassificationOverride struct {
	// Entity is the base model of InferenceClassificationOverride
	Entity
	// ClassifyAs undocumented
	ClassifyAs *InferenceClassificationType `json:"classifyAs,omitempty"`
	// SenderEmailAddress undocumented
	SenderEmailAddress *EmailAddress `json:"senderEmailAddress,omitempty"`
}