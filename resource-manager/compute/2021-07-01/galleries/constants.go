package galleries

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GallerySharingPermissionTypes string

const (
	GallerySharingPermissionTypesGroups  GallerySharingPermissionTypes = "Groups"
	GallerySharingPermissionTypesPrivate GallerySharingPermissionTypes = "Private"
)

func PossibleValuesForGallerySharingPermissionTypes() []string {
	return []string{
		string(GallerySharingPermissionTypesGroups),
		string(GallerySharingPermissionTypesPrivate),
	}
}

func (s *GallerySharingPermissionTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGallerySharingPermissionTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGallerySharingPermissionTypes(input string) (*GallerySharingPermissionTypes, error) {
	vals := map[string]GallerySharingPermissionTypes{
		"groups":  GallerySharingPermissionTypesGroups,
		"private": GallerySharingPermissionTypesPrivate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GallerySharingPermissionTypes(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateMigrating ProvisioningState = "Migrating"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateMigrating),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"creating":  ProvisioningStateCreating,
		"deleting":  ProvisioningStateDeleting,
		"failed":    ProvisioningStateFailed,
		"migrating": ProvisioningStateMigrating,
		"succeeded": ProvisioningStateSucceeded,
		"updating":  ProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}

type SelectPermissions string

const (
	SelectPermissionsPermissions SelectPermissions = "Permissions"
)

func PossibleValuesForSelectPermissions() []string {
	return []string{
		string(SelectPermissionsPermissions),
	}
}

func (s *SelectPermissions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSelectPermissions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSelectPermissions(input string) (*SelectPermissions, error) {
	vals := map[string]SelectPermissions{
		"permissions": SelectPermissionsPermissions,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SelectPermissions(input)
	return &out, nil
}

type SharingProfileGroupTypes string

const (
	SharingProfileGroupTypesAADTenants    SharingProfileGroupTypes = "AADTenants"
	SharingProfileGroupTypesSubscriptions SharingProfileGroupTypes = "Subscriptions"
)

func PossibleValuesForSharingProfileGroupTypes() []string {
	return []string{
		string(SharingProfileGroupTypesAADTenants),
		string(SharingProfileGroupTypesSubscriptions),
	}
}

func (s *SharingProfileGroupTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSharingProfileGroupTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSharingProfileGroupTypes(input string) (*SharingProfileGroupTypes, error) {
	vals := map[string]SharingProfileGroupTypes{
		"aadtenants":    SharingProfileGroupTypesAADTenants,
		"subscriptions": SharingProfileGroupTypesSubscriptions,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharingProfileGroupTypes(input)
	return &out, nil
}
