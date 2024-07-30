package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider string

const (
	DefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_AzureActiveDirectory    DefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider = "azureActiveDirectory"
	DefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_DefaultConfiguredIdp    DefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider = "defaultConfiguredIdp"
	DefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_EmailOneTimePasscode    DefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider = "emailOneTimePasscode"
	DefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_ExternalFederation      DefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider = "externalFederation"
	DefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_MicrosoftAccount        DefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider = "microsoftAccount"
	DefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_SocialIdentityProviders DefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider = "socialIdentityProviders"
)

func PossibleValuesForDefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider() []string {
	return []string{
		string(DefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_AzureActiveDirectory),
		string(DefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_DefaultConfiguredIdp),
		string(DefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_EmailOneTimePasscode),
		string(DefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_ExternalFederation),
		string(DefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_MicrosoftAccount),
		string(DefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_SocialIdentityProviders),
	}
}

func (s *DefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider(input string) (*DefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider, error) {
	vals := map[string]DefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider{
		"azureactivedirectory":    DefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_AzureActiveDirectory,
		"defaultconfiguredidp":    DefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_DefaultConfiguredIdp,
		"emailonetimepasscode":    DefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_EmailOneTimePasscode,
		"externalfederation":      DefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_ExternalFederation,
		"microsoftaccount":        DefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_MicrosoftAccount,
		"socialidentityproviders": DefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider_SocialIdentityProviders,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultInvitationRedemptionIdentityProviderConfigurationFallbackIdentityProvider(input)
	return &out, nil
}
