package managements

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckNameAvailabilityRequestType string

const (
	CheckNameAvailabilityRequestTypeMicrosoftPointManagementManagementGroups CheckNameAvailabilityRequestType = "Microsoft.Management/managementGroups"
)

func PossibleValuesForCheckNameAvailabilityRequestType() []string {
	return []string{
		string(CheckNameAvailabilityRequestTypeMicrosoftPointManagementManagementGroups),
	}
}

func (s *CheckNameAvailabilityRequestType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCheckNameAvailabilityRequestType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCheckNameAvailabilityRequestType(input string) (*CheckNameAvailabilityRequestType, error) {
	vals := map[string]CheckNameAvailabilityRequestType{
		"microsoft.management/managementgroups": CheckNameAvailabilityRequestTypeMicrosoftPointManagementManagementGroups,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CheckNameAvailabilityRequestType(input)
	return &out, nil
}

type EntitySearchType string

const (
	EntitySearchTypeAllowedChildren             EntitySearchType = "AllowedChildren"
	EntitySearchTypeAllowedParents              EntitySearchType = "AllowedParents"
	EntitySearchTypeChildrenOnly                EntitySearchType = "ChildrenOnly"
	EntitySearchTypeParentAndFirstLevelChildren EntitySearchType = "ParentAndFirstLevelChildren"
	EntitySearchTypeParentOnly                  EntitySearchType = "ParentOnly"
)

func PossibleValuesForEntitySearchType() []string {
	return []string{
		string(EntitySearchTypeAllowedChildren),
		string(EntitySearchTypeAllowedParents),
		string(EntitySearchTypeChildrenOnly),
		string(EntitySearchTypeParentAndFirstLevelChildren),
		string(EntitySearchTypeParentOnly),
	}
}

func (s *EntitySearchType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEntitySearchType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEntitySearchType(input string) (*EntitySearchType, error) {
	vals := map[string]EntitySearchType{
		"allowedchildren":             EntitySearchTypeAllowedChildren,
		"allowedparents":              EntitySearchTypeAllowedParents,
		"childrenonly":                EntitySearchTypeChildrenOnly,
		"parentandfirstlevelchildren": EntitySearchTypeParentAndFirstLevelChildren,
		"parentonly":                  EntitySearchTypeParentOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EntitySearchType(input)
	return &out, nil
}

type EntityViewParameterType string

const (
	EntityViewParameterTypeAudit             EntityViewParameterType = "Audit"
	EntityViewParameterTypeFullHierarchy     EntityViewParameterType = "FullHierarchy"
	EntityViewParameterTypeGroupsOnly        EntityViewParameterType = "GroupsOnly"
	EntityViewParameterTypeSubscriptionsOnly EntityViewParameterType = "SubscriptionsOnly"
)

func PossibleValuesForEntityViewParameterType() []string {
	return []string{
		string(EntityViewParameterTypeAudit),
		string(EntityViewParameterTypeFullHierarchy),
		string(EntityViewParameterTypeGroupsOnly),
		string(EntityViewParameterTypeSubscriptionsOnly),
	}
}

func (s *EntityViewParameterType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEntityViewParameterType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEntityViewParameterType(input string) (*EntityViewParameterType, error) {
	vals := map[string]EntityViewParameterType{
		"audit":             EntityViewParameterTypeAudit,
		"fullhierarchy":     EntityViewParameterTypeFullHierarchy,
		"groupsonly":        EntityViewParameterTypeGroupsOnly,
		"subscriptionsonly": EntityViewParameterTypeSubscriptionsOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EntityViewParameterType(input)
	return &out, nil
}

type Permissions string

const (
	PermissionsDelete   Permissions = "delete"
	PermissionsEdit     Permissions = "edit"
	PermissionsNoaccess Permissions = "noaccess"
	PermissionsView     Permissions = "view"
)

func PossibleValuesForPermissions() []string {
	return []string{
		string(PermissionsDelete),
		string(PermissionsEdit),
		string(PermissionsNoaccess),
		string(PermissionsView),
	}
}

func (s *Permissions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePermissions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePermissions(input string) (*Permissions, error) {
	vals := map[string]Permissions{
		"delete":   PermissionsDelete,
		"edit":     PermissionsEdit,
		"noaccess": PermissionsNoaccess,
		"view":     PermissionsView,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Permissions(input)
	return &out, nil
}

type Reason string

const (
	ReasonAlreadyExists Reason = "AlreadyExists"
	ReasonInvalid       Reason = "Invalid"
)

func PossibleValuesForReason() []string {
	return []string{
		string(ReasonAlreadyExists),
		string(ReasonInvalid),
	}
}

func (s *Reason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReason(input string) (*Reason, error) {
	vals := map[string]Reason{
		"alreadyexists": ReasonAlreadyExists,
		"invalid":       ReasonInvalid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Reason(input)
	return &out, nil
}

type Status string

const (
	StatusCancelled                Status = "Cancelled"
	StatusCompleted                Status = "Completed"
	StatusFailed                   Status = "Failed"
	StatusNotStarted               Status = "NotStarted"
	StatusNotStartedButGroupsExist Status = "NotStartedButGroupsExist"
	StatusStarted                  Status = "Started"
)

func PossibleValuesForStatus() []string {
	return []string{
		string(StatusCancelled),
		string(StatusCompleted),
		string(StatusFailed),
		string(StatusNotStarted),
		string(StatusNotStartedButGroupsExist),
		string(StatusStarted),
	}
}

func (s *Status) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStatus(input string) (*Status, error) {
	vals := map[string]Status{
		"cancelled":                StatusCancelled,
		"completed":                StatusCompleted,
		"failed":                   StatusFailed,
		"notstarted":               StatusNotStarted,
		"notstartedbutgroupsexist": StatusNotStartedButGroupsExist,
		"started":                  StatusStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Status(input)
	return &out, nil
}
