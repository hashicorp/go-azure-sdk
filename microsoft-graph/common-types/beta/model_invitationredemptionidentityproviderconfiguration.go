package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvitationRedemptionIdentityProviderConfiguration struct {
	FallbackIdentityProvider               *InvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider               `json:"fallbackIdentityProvider,omitempty"`
	ODataType                              *string                                                                                  `json:"@odata.type,omitempty"`
	PrimaryIdentityProviderPrecedenceOrder *InvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder `json:"primaryIdentityProviderPrecedenceOrder,omitempty"`
}
