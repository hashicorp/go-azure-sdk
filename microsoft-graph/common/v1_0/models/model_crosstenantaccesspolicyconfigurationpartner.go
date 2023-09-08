package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrossTenantAccessPolicyConfigurationPartner struct {
	AutomaticUserConsentSettings *InboundOutboundPolicyConfiguration   `json:"automaticUserConsentSettings,omitempty"`
	B2bCollaborationInbound      *CrossTenantAccessPolicyB2BSetting    `json:"b2bCollaborationInbound,omitempty"`
	B2bCollaborationOutbound     *CrossTenantAccessPolicyB2BSetting    `json:"b2bCollaborationOutbound,omitempty"`
	B2bDirectConnectInbound      *CrossTenantAccessPolicyB2BSetting    `json:"b2bDirectConnectInbound,omitempty"`
	B2bDirectConnectOutbound     *CrossTenantAccessPolicyB2BSetting    `json:"b2bDirectConnectOutbound,omitempty"`
	IdentitySynchronization      *CrossTenantIdentitySyncPolicyPartner `json:"identitySynchronization,omitempty"`
	InboundTrust                 *CrossTenantAccessPolicyInboundTrust  `json:"inboundTrust,omitempty"`
	IsServiceProvider            *bool                                 `json:"isServiceProvider,omitempty"`
	ODataType                    *string                               `json:"@odata.type,omitempty"`
	TenantId                     *string                               `json:"tenantId,omitempty"`
}
