package connectordefinitions

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataConnectorDefinitionKind string

const (
	DataConnectorDefinitionKindCustomizable DataConnectorDefinitionKind = "Customizable"
)

func PossibleValuesForDataConnectorDefinitionKind() []string {
	return []string{
		string(DataConnectorDefinitionKindCustomizable),
	}
}

func (s *DataConnectorDefinitionKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataConnectorDefinitionKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataConnectorDefinitionKind(input string) (*DataConnectorDefinitionKind, error) {
	vals := map[string]DataConnectorDefinitionKind{
		"customizable": DataConnectorDefinitionKindCustomizable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataConnectorDefinitionKind(input)
	return &out, nil
}

type ProviderPermissionsScope string

const (
	ProviderPermissionsScopeResourceGroup ProviderPermissionsScope = "ResourceGroup"
	ProviderPermissionsScopeSubscription  ProviderPermissionsScope = "Subscription"
	ProviderPermissionsScopeWorkspace     ProviderPermissionsScope = "Workspace"
)

func PossibleValuesForProviderPermissionsScope() []string {
	return []string{
		string(ProviderPermissionsScopeResourceGroup),
		string(ProviderPermissionsScopeSubscription),
		string(ProviderPermissionsScopeWorkspace),
	}
}

func (s *ProviderPermissionsScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProviderPermissionsScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProviderPermissionsScope(input string) (*ProviderPermissionsScope, error) {
	vals := map[string]ProviderPermissionsScope{
		"resourcegroup": ProviderPermissionsScopeResourceGroup,
		"subscription":  ProviderPermissionsScopeSubscription,
		"workspace":     ProviderPermissionsScopeWorkspace,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProviderPermissionsScope(input)
	return &out, nil
}
