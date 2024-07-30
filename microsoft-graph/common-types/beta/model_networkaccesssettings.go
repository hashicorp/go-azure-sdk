package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessSettings struct {
	ConditionalAccess *NetworkaccessConditionalAccessSettings `json:"conditionalAccess,omitempty"`
	CrossTenantAccess *NetworkaccessCrossTenantAccessSettings `json:"crossTenantAccess,omitempty"`
	EnrichedAuditLogs *NetworkaccessEnrichedAuditLogs         `json:"enrichedAuditLogs,omitempty"`
	ForwardingOptions *NetworkaccessForwardingOptions         `json:"forwardingOptions,omitempty"`
	Id                *string                                 `json:"id,omitempty"`
	ODataType         *string                                 `json:"@odata.type,omitempty"`
}
