package v2workspaceconnectionresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionAuthType string

const (
	ConnectionAuthTypeManagedIdentity  ConnectionAuthType = "ManagedIdentity"
	ConnectionAuthTypeNone             ConnectionAuthType = "None"
	ConnectionAuthTypePAT              ConnectionAuthType = "PAT"
	ConnectionAuthTypeSAS              ConnectionAuthType = "SAS"
	ConnectionAuthTypeUsernamePassword ConnectionAuthType = "UsernamePassword"
)

func PossibleValuesForConnectionAuthType() []string {
	return []string{
		string(ConnectionAuthTypeManagedIdentity),
		string(ConnectionAuthTypeNone),
		string(ConnectionAuthTypePAT),
		string(ConnectionAuthTypeSAS),
		string(ConnectionAuthTypeUsernamePassword),
	}
}

type ConnectionCategory string

const (
	ConnectionCategoryContainerRegistry ConnectionCategory = "ContainerRegistry"
	ConnectionCategoryGit               ConnectionCategory = "Git"
	ConnectionCategoryPythonFeed        ConnectionCategory = "PythonFeed"
)

func PossibleValuesForConnectionCategory() []string {
	return []string{
		string(ConnectionCategoryContainerRegistry),
		string(ConnectionCategoryGit),
		string(ConnectionCategoryPythonFeed),
	}
}

type ValueFormat string

const (
	ValueFormatJSON ValueFormat = "JSON"
)

func PossibleValuesForValueFormat() []string {
	return []string{
		string(ValueFormatJSON),
	}
}
