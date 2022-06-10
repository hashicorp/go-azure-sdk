package accounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HiveMetastoreProperties struct {
	DatabaseName                    *string                          `json:"databaseName,omitempty"`
	NestedResourceProvisioningState *NestedResourceProvisioningState `json:"nestedResourceProvisioningState,omitempty"`
	Password                        *string                          `json:"password,omitempty"`
	RuntimeVersion                  *string                          `json:"runtimeVersion,omitempty"`
	ServerUri                       *string                          `json:"serverUri,omitempty"`
	UserName                        *string                          `json:"userName,omitempty"`
}
