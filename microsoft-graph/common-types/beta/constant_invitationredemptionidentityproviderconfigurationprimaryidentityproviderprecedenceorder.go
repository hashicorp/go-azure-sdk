package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder string

const (
	InvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder_AzureActiveDirectory    InvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder = "azureActiveDirectory"
	InvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder_DefaultConfiguredIdp    InvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder = "defaultConfiguredIdp"
	InvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder_EmailOneTimePasscode    InvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder = "emailOneTimePasscode"
	InvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder_ExternalFederation      InvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder = "externalFederation"
	InvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder_MicrosoftAccount        InvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder = "microsoftAccount"
	InvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder_SocialIdentityProviders InvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder = "socialIdentityProviders"
)

func PossibleValuesForInvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder() []string {
	return []string{
		string(InvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder_AzureActiveDirectory),
		string(InvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder_DefaultConfiguredIdp),
		string(InvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder_EmailOneTimePasscode),
		string(InvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder_ExternalFederation),
		string(InvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder_MicrosoftAccount),
		string(InvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder_SocialIdentityProviders),
	}
}

func (s *InvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder(input string) (*InvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder, error) {
	vals := map[string]InvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder{
		"azureactivedirectory":    InvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder_AzureActiveDirectory,
		"defaultconfiguredidp":    InvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder_DefaultConfiguredIdp,
		"emailonetimepasscode":    InvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder_EmailOneTimePasscode,
		"externalfederation":      InvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder_ExternalFederation,
		"microsoftaccount":        InvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder_MicrosoftAccount,
		"socialidentityproviders": InvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder_SocialIdentityProviders,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InvitationRedemptionIdentityProviderConfigurationPrimaryIdentityProviderPrecedenceOrder(input)
	return &out, nil
}
