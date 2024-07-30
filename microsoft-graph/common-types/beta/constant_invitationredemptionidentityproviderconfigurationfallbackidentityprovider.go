package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider string

const (
	InvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_AzureActiveDirectory    InvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider = "azureActiveDirectory"
	InvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_DefaultConfiguredIdp    InvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider = "defaultConfiguredIdp"
	InvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_EmailOneTimePasscode    InvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider = "emailOneTimePasscode"
	InvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_ExternalFederation      InvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider = "externalFederation"
	InvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_MicrosoftAccount        InvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider = "microsoftAccount"
	InvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_SocialIdentityProviders InvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider = "socialIdentityProviders"
)

func PossibleValuesForInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider() []string {
	return []string{
		string(InvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_AzureActiveDirectory),
		string(InvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_DefaultConfiguredIdp),
		string(InvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_EmailOneTimePasscode),
		string(InvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_ExternalFederation),
		string(InvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_MicrosoftAccount),
		string(InvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_SocialIdentityProviders),
	}
}

func (s *InvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider(input string) (*InvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider, error) {
	vals := map[string]InvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider{
		"azureactivedirectory":    InvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_AzureActiveDirectory,
		"defaultconfiguredidp":    InvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_DefaultConfiguredIdp,
		"emailonetimepasscode":    InvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_EmailOneTimePasscode,
		"externalfederation":      InvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_ExternalFederation,
		"microsoftaccount":        InvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_MicrosoftAccount,
		"socialidentityproviders": InvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_SocialIdentityProviders,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider(input)
	return &out, nil
}
