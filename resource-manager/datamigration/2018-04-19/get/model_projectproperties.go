package get

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProjectProperties struct {
	CreationTime         *string                   `json:"creationTime,omitempty"`
	DatabasesInfo        *[]DatabaseInfo           `json:"databasesInfo,omitempty"`
	ProvisioningState    *ProjectProvisioningState `json:"provisioningState,omitempty"`
	SourceConnectionInfo *ConnectionInfo           `json:"sourceConnectionInfo,omitempty"`
	SourcePlatform       ProjectSourcePlatform     `json:"sourcePlatform"`
	TargetConnectionInfo *ConnectionInfo           `json:"targetConnectionInfo,omitempty"`
	TargetPlatform       ProjectTargetPlatform     `json:"targetPlatform"`
}
