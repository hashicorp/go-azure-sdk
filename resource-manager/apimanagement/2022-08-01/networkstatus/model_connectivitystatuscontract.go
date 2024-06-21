package networkstatus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectivityStatusContract struct {
	Error            *string                `json:"error,omitempty"`
	IsOptional       bool                   `json:"isOptional"`
	LastStatusChange string                 `json:"lastStatusChange"`
	LastUpdated      string                 `json:"lastUpdated"`
	Name             string                 `json:"name"`
	ResourceType     string                 `json:"resourceType"`
	Status           ConnectivityStatusType `json:"status"`
}
