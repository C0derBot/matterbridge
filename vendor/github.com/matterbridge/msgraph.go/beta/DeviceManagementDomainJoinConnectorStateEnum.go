// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementDomainJoinConnectorState undocumented
type DeviceManagementDomainJoinConnectorState int

const (
	// DeviceManagementDomainJoinConnectorStateVActive undocumented
	DeviceManagementDomainJoinConnectorStateVActive DeviceManagementDomainJoinConnectorState = 0
	// DeviceManagementDomainJoinConnectorStateVError undocumented
	DeviceManagementDomainJoinConnectorStateVError DeviceManagementDomainJoinConnectorState = 1
	// DeviceManagementDomainJoinConnectorStateVInactive undocumented
	DeviceManagementDomainJoinConnectorStateVInactive DeviceManagementDomainJoinConnectorState = 2
)

// DeviceManagementDomainJoinConnectorStatePActive returns a pointer to DeviceManagementDomainJoinConnectorStateVActive
func DeviceManagementDomainJoinConnectorStatePActive() *DeviceManagementDomainJoinConnectorState {
	v := DeviceManagementDomainJoinConnectorStateVActive
	return &v
}

// DeviceManagementDomainJoinConnectorStatePError returns a pointer to DeviceManagementDomainJoinConnectorStateVError
func DeviceManagementDomainJoinConnectorStatePError() *DeviceManagementDomainJoinConnectorState {
	v := DeviceManagementDomainJoinConnectorStateVError
	return &v
}

// DeviceManagementDomainJoinConnectorStatePInactive returns a pointer to DeviceManagementDomainJoinConnectorStateVInactive
func DeviceManagementDomainJoinConnectorStatePInactive() *DeviceManagementDomainJoinConnectorState {
	v := DeviceManagementDomainJoinConnectorStateVInactive
	return &v
}