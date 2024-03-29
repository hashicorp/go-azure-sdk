package hypervmachines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharePointServer struct {
	IsEnterprise *bool   `json:"isEnterprise,omitempty"`
	ProductName  *string `json:"productName,omitempty"`
	Status       *string `json:"status,omitempty"`
	Version      *string `json:"version,omitempty"`
}
