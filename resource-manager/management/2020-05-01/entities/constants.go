package entities

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InheritedPermissions string

const (
	InheritedPermissionsDelete   InheritedPermissions = "delete"
	InheritedPermissionsEdit     InheritedPermissions = "edit"
	InheritedPermissionsNoaccess InheritedPermissions = "noaccess"
	InheritedPermissionsView     InheritedPermissions = "view"
)

func PossibleValuesForInheritedPermissions() []string {
	return []string{
		string(InheritedPermissionsDelete),
		string(InheritedPermissionsEdit),
		string(InheritedPermissionsNoaccess),
		string(InheritedPermissionsView),
	}
}

func (s *InheritedPermissions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInheritedPermissions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInheritedPermissions(input string) (*InheritedPermissions, error) {
	vals := map[string]InheritedPermissions{
		"delete":   InheritedPermissionsDelete,
		"edit":     InheritedPermissionsEdit,
		"noaccess": InheritedPermissionsNoaccess,
		"view":     InheritedPermissionsView,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InheritedPermissions(input)
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

type Search string

const (
	SearchAllowedChildren             Search = "AllowedChildren"
	SearchAllowedParents              Search = "AllowedParents"
	SearchChildrenOnly                Search = "ChildrenOnly"
	SearchParentAndFirstLevelChildren Search = "ParentAndFirstLevelChildren"
	SearchParentOnly                  Search = "ParentOnly"
)

func PossibleValuesForSearch() []string {
	return []string{
		string(SearchAllowedChildren),
		string(SearchAllowedParents),
		string(SearchChildrenOnly),
		string(SearchParentAndFirstLevelChildren),
		string(SearchParentOnly),
	}
}

func (s *Search) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSearch(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSearch(input string) (*Search, error) {
	vals := map[string]Search{
		"allowedchildren":             SearchAllowedChildren,
		"allowedparents":              SearchAllowedParents,
		"childrenonly":                SearchChildrenOnly,
		"parentandfirstlevelchildren": SearchParentAndFirstLevelChildren,
		"parentonly":                  SearchParentOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Search(input)
	return &out, nil
}

type View string

const (
	ViewAudit             View = "Audit"
	ViewFullHierarchy     View = "FullHierarchy"
	ViewGroupsOnly        View = "GroupsOnly"
	ViewSubscriptionsOnly View = "SubscriptionsOnly"
)

func PossibleValuesForView() []string {
	return []string{
		string(ViewAudit),
		string(ViewFullHierarchy),
		string(ViewGroupsOnly),
		string(ViewSubscriptionsOnly),
	}
}

func (s *View) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseView(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseView(input string) (*View, error) {
	vals := map[string]View{
		"audit":             ViewAudit,
		"fullhierarchy":     ViewFullHierarchy,
		"groupsonly":        ViewGroupsOnly,
		"subscriptionsonly": ViewSubscriptionsOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := View(input)
	return &out, nil
}
