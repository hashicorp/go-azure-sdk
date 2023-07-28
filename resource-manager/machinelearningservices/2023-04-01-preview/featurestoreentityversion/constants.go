package featurestoreentityversion

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssetProvisioningState string

const (
	AssetProvisioningStateCanceled  AssetProvisioningState = "Canceled"
	AssetProvisioningStateCreating  AssetProvisioningState = "Creating"
	AssetProvisioningStateDeleting  AssetProvisioningState = "Deleting"
	AssetProvisioningStateFailed    AssetProvisioningState = "Failed"
	AssetProvisioningStateSucceeded AssetProvisioningState = "Succeeded"
	AssetProvisioningStateUpdating  AssetProvisioningState = "Updating"
)

func PossibleValuesForAssetProvisioningState() []string {
	return []string{
		string(AssetProvisioningStateCanceled),
		string(AssetProvisioningStateCreating),
		string(AssetProvisioningStateDeleting),
		string(AssetProvisioningStateFailed),
		string(AssetProvisioningStateSucceeded),
		string(AssetProvisioningStateUpdating),
	}
}

func (s *AssetProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAssetProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAssetProvisioningState(input string) (*AssetProvisioningState, error) {
	vals := map[string]AssetProvisioningState{
		"canceled":  AssetProvisioningStateCanceled,
		"creating":  AssetProvisioningStateCreating,
		"deleting":  AssetProvisioningStateDeleting,
		"failed":    AssetProvisioningStateFailed,
		"succeeded": AssetProvisioningStateSucceeded,
		"updating":  AssetProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssetProvisioningState(input)
	return &out, nil
}

type AutoDeleteCondition string

const (
	AutoDeleteConditionCreatedGreaterThan      AutoDeleteCondition = "CreatedGreaterThan"
	AutoDeleteConditionLastAccessedGreaterThan AutoDeleteCondition = "LastAccessedGreaterThan"
)

func PossibleValuesForAutoDeleteCondition() []string {
	return []string{
		string(AutoDeleteConditionCreatedGreaterThan),
		string(AutoDeleteConditionLastAccessedGreaterThan),
	}
}

func (s *AutoDeleteCondition) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutoDeleteCondition(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutoDeleteCondition(input string) (*AutoDeleteCondition, error) {
	vals := map[string]AutoDeleteCondition{
		"createdgreaterthan":      AutoDeleteConditionCreatedGreaterThan,
		"lastaccessedgreaterthan": AutoDeleteConditionLastAccessedGreaterThan,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutoDeleteCondition(input)
	return &out, nil
}

type FeatureDataType string

const (
	FeatureDataTypeBinary   FeatureDataType = "Binary"
	FeatureDataTypeBoolean  FeatureDataType = "Boolean"
	FeatureDataTypeDatetime FeatureDataType = "Datetime"
	FeatureDataTypeDouble   FeatureDataType = "Double"
	FeatureDataTypeFloat    FeatureDataType = "Float"
	FeatureDataTypeInteger  FeatureDataType = "Integer"
	FeatureDataTypeLong     FeatureDataType = "Long"
	FeatureDataTypeString   FeatureDataType = "String"
)

func PossibleValuesForFeatureDataType() []string {
	return []string{
		string(FeatureDataTypeBinary),
		string(FeatureDataTypeBoolean),
		string(FeatureDataTypeDatetime),
		string(FeatureDataTypeDouble),
		string(FeatureDataTypeFloat),
		string(FeatureDataTypeInteger),
		string(FeatureDataTypeLong),
		string(FeatureDataTypeString),
	}
}

func (s *FeatureDataType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFeatureDataType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFeatureDataType(input string) (*FeatureDataType, error) {
	vals := map[string]FeatureDataType{
		"binary":   FeatureDataTypeBinary,
		"boolean":  FeatureDataTypeBoolean,
		"datetime": FeatureDataTypeDatetime,
		"double":   FeatureDataTypeDouble,
		"float":    FeatureDataTypeFloat,
		"integer":  FeatureDataTypeInteger,
		"long":     FeatureDataTypeLong,
		"string":   FeatureDataTypeString,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FeatureDataType(input)
	return &out, nil
}

type ListViewType string

const (
	ListViewTypeActiveOnly   ListViewType = "ActiveOnly"
	ListViewTypeAll          ListViewType = "All"
	ListViewTypeArchivedOnly ListViewType = "ArchivedOnly"
)

func PossibleValuesForListViewType() []string {
	return []string{
		string(ListViewTypeActiveOnly),
		string(ListViewTypeAll),
		string(ListViewTypeArchivedOnly),
	}
}

func (s *ListViewType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseListViewType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseListViewType(input string) (*ListViewType, error) {
	vals := map[string]ListViewType{
		"activeonly":   ListViewTypeActiveOnly,
		"all":          ListViewTypeAll,
		"archivedonly": ListViewTypeArchivedOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ListViewType(input)
	return &out, nil
}
