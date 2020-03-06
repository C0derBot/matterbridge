// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// RiskLevel undocumented
type RiskLevel int

const (
	// RiskLevelVLow undocumented
	RiskLevelVLow RiskLevel = 0
	// RiskLevelVMedium undocumented
	RiskLevelVMedium RiskLevel = 1
	// RiskLevelVHigh undocumented
	RiskLevelVHigh RiskLevel = 2
	// RiskLevelVHidden undocumented
	RiskLevelVHidden RiskLevel = 3
	// RiskLevelVNone undocumented
	RiskLevelVNone RiskLevel = 4
	// RiskLevelVUnknownFutureValue undocumented
	RiskLevelVUnknownFutureValue RiskLevel = 5
)

// RiskLevelPLow returns a pointer to RiskLevelVLow
func RiskLevelPLow() *RiskLevel {
	v := RiskLevelVLow
	return &v
}

// RiskLevelPMedium returns a pointer to RiskLevelVMedium
func RiskLevelPMedium() *RiskLevel {
	v := RiskLevelVMedium
	return &v
}

// RiskLevelPHigh returns a pointer to RiskLevelVHigh
func RiskLevelPHigh() *RiskLevel {
	v := RiskLevelVHigh
	return &v
}

// RiskLevelPHidden returns a pointer to RiskLevelVHidden
func RiskLevelPHidden() *RiskLevel {
	v := RiskLevelVHidden
	return &v
}

// RiskLevelPNone returns a pointer to RiskLevelVNone
func RiskLevelPNone() *RiskLevel {
	v := RiskLevelVNone
	return &v
}

// RiskLevelPUnknownFutureValue returns a pointer to RiskLevelVUnknownFutureValue
func RiskLevelPUnknownFutureValue() *RiskLevel {
	v := RiskLevelVUnknownFutureValue
	return &v
}