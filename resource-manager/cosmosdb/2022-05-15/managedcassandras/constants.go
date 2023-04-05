package managedcassandras

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethod string

const (
	AuthenticationMethodCassandra AuthenticationMethod = "Cassandra"
	AuthenticationMethodNone      AuthenticationMethod = "None"
)

func PossibleValuesForAuthenticationMethod() []string {
	return []string{
		string(AuthenticationMethodCassandra),
		string(AuthenticationMethodNone),
	}
}

type ConnectionState string

const (
	ConnectionStateDatacenterToDatacenterNetworkError           ConnectionState = "DatacenterToDatacenterNetworkError"
	ConnectionStateInternalError                                ConnectionState = "InternalError"
	ConnectionStateInternalOperatorToDataCenterCertificateError ConnectionState = "InternalOperatorToDataCenterCertificateError"
	ConnectionStateOK                                           ConnectionState = "OK"
	ConnectionStateOperatorToDataCenterNetworkError             ConnectionState = "OperatorToDataCenterNetworkError"
	ConnectionStateUnknown                                      ConnectionState = "Unknown"
)

func PossibleValuesForConnectionState() []string {
	return []string{
		string(ConnectionStateDatacenterToDatacenterNetworkError),
		string(ConnectionStateInternalError),
		string(ConnectionStateInternalOperatorToDataCenterCertificateError),
		string(ConnectionStateOK),
		string(ConnectionStateOperatorToDataCenterNetworkError),
		string(ConnectionStateUnknown),
	}
}

type ManagedCassandraProvisioningState string

const (
	ManagedCassandraProvisioningStateCanceled  ManagedCassandraProvisioningState = "Canceled"
	ManagedCassandraProvisioningStateCreating  ManagedCassandraProvisioningState = "Creating"
	ManagedCassandraProvisioningStateDeleting  ManagedCassandraProvisioningState = "Deleting"
	ManagedCassandraProvisioningStateFailed    ManagedCassandraProvisioningState = "Failed"
	ManagedCassandraProvisioningStateSucceeded ManagedCassandraProvisioningState = "Succeeded"
	ManagedCassandraProvisioningStateUpdating  ManagedCassandraProvisioningState = "Updating"
)

func PossibleValuesForManagedCassandraProvisioningState() []string {
	return []string{
		string(ManagedCassandraProvisioningStateCanceled),
		string(ManagedCassandraProvisioningStateCreating),
		string(ManagedCassandraProvisioningStateDeleting),
		string(ManagedCassandraProvisioningStateFailed),
		string(ManagedCassandraProvisioningStateSucceeded),
		string(ManagedCassandraProvisioningStateUpdating),
	}
}

type NodeState string

const (
	NodeStateJoining NodeState = "Joining"
	NodeStateLeaving NodeState = "Leaving"
	NodeStateMoving  NodeState = "Moving"
	NodeStateNormal  NodeState = "Normal"
	NodeStateStopped NodeState = "Stopped"
)

func PossibleValuesForNodeState() []string {
	return []string{
		string(NodeStateJoining),
		string(NodeStateLeaving),
		string(NodeStateMoving),
		string(NodeStateNormal),
		string(NodeStateStopped),
	}
}
