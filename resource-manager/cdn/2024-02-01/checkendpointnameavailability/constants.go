package checkendpointnameavailability

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutoGeneratedDomainNameLabelScope string

const (
	AutoGeneratedDomainNameLabelScopeNoReuse            AutoGeneratedDomainNameLabelScope = "NoReuse"
	AutoGeneratedDomainNameLabelScopeResourceGroupReuse AutoGeneratedDomainNameLabelScope = "ResourceGroupReuse"
	AutoGeneratedDomainNameLabelScopeSubscriptionReuse  AutoGeneratedDomainNameLabelScope = "SubscriptionReuse"
	AutoGeneratedDomainNameLabelScopeTenantReuse        AutoGeneratedDomainNameLabelScope = "TenantReuse"
)

func PossibleValuesForAutoGeneratedDomainNameLabelScope() []string {
	return []string{
		string(AutoGeneratedDomainNameLabelScopeNoReuse),
		string(AutoGeneratedDomainNameLabelScopeResourceGroupReuse),
		string(AutoGeneratedDomainNameLabelScopeSubscriptionReuse),
		string(AutoGeneratedDomainNameLabelScopeTenantReuse),
	}
}

func (s *AutoGeneratedDomainNameLabelScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutoGeneratedDomainNameLabelScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutoGeneratedDomainNameLabelScope(input string) (*AutoGeneratedDomainNameLabelScope, error) {
	vals := map[string]AutoGeneratedDomainNameLabelScope{
		"noreuse":            AutoGeneratedDomainNameLabelScopeNoReuse,
		"resourcegroupreuse": AutoGeneratedDomainNameLabelScopeResourceGroupReuse,
		"subscriptionreuse":  AutoGeneratedDomainNameLabelScopeSubscriptionReuse,
		"tenantreuse":        AutoGeneratedDomainNameLabelScopeTenantReuse,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutoGeneratedDomainNameLabelScope(input)
	return &out, nil
}

type ResourceType string

const (
	ResourceTypeMicrosoftPointCdnProfilesAfdEndpoints ResourceType = "Microsoft.Cdn/Profiles/AfdEndpoints"
	ResourceTypeMicrosoftPointCdnProfilesEndpoints    ResourceType = "Microsoft.Cdn/Profiles/Endpoints"
)

func PossibleValuesForResourceType() []string {
	return []string{
		string(ResourceTypeMicrosoftPointCdnProfilesAfdEndpoints),
		string(ResourceTypeMicrosoftPointCdnProfilesEndpoints),
	}
}

func (s *ResourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResourceType(input string) (*ResourceType, error) {
	vals := map[string]ResourceType{
		"microsoft.cdn/profiles/afdendpoints": ResourceTypeMicrosoftPointCdnProfilesAfdEndpoints,
		"microsoft.cdn/profiles/endpoints":    ResourceTypeMicrosoftPointCdnProfilesEndpoints,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceType(input)
	return &out, nil
}
