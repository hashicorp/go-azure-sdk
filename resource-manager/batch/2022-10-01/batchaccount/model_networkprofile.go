package batchaccount

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkProfile struct {
	AccountAccess        *EndpointAccessProfile `json:"accountAccess"`
	NodeManagementAccess *EndpointAccessProfile `json:"nodeManagementAccess"`
}
