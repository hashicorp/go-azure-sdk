package integrationruntime

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedIntegrationRuntimeStatusTypeProperties struct {
	CreateTime    *string                                   `json:"createTime,omitempty"`
	LastOperation *ManagedIntegrationRuntimeOperationResult `json:"lastOperation,omitempty"`
	Nodes         *[]ManagedIntegrationRuntimeNode          `json:"nodes,omitempty"`
	OtherErrors   *[]ManagedIntegrationRuntimeError         `json:"otherErrors,omitempty"`
}
