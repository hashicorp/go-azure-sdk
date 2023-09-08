package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessEnrichedAuditLogs struct {
	Exchange   *NetworkaccessEnrichedAuditLogsSettings `json:"exchange,omitempty"`
	Id         *string                                 `json:"id,omitempty"`
	ODataType  *string                                 `json:"@odata.type,omitempty"`
	Sharepoint *NetworkaccessEnrichedAuditLogsSettings `json:"sharepoint,omitempty"`
	Teams      *NetworkaccessEnrichedAuditLogsSettings `json:"teams,omitempty"`
}
