package dppfeaturesupport

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeatureSupportStatus string

const (
	FeatureSupportStatusAlphaPreview       FeatureSupportStatus = "AlphaPreview"
	FeatureSupportStatusGenerallyAvailable FeatureSupportStatus = "GenerallyAvailable"
	FeatureSupportStatusInvalid            FeatureSupportStatus = "Invalid"
	FeatureSupportStatusNotSupported       FeatureSupportStatus = "NotSupported"
	FeatureSupportStatusPrivatePreview     FeatureSupportStatus = "PrivatePreview"
	FeatureSupportStatusPublicPreview      FeatureSupportStatus = "PublicPreview"
)

func PossibleValuesForFeatureSupportStatus() []string {
	return []string{
		string(FeatureSupportStatusAlphaPreview),
		string(FeatureSupportStatusGenerallyAvailable),
		string(FeatureSupportStatusInvalid),
		string(FeatureSupportStatusNotSupported),
		string(FeatureSupportStatusPrivatePreview),
		string(FeatureSupportStatusPublicPreview),
	}
}

func (s *FeatureSupportStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFeatureSupportStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFeatureSupportStatus(input string) (*FeatureSupportStatus, error) {
	vals := map[string]FeatureSupportStatus{
		"alphapreview":       FeatureSupportStatusAlphaPreview,
		"generallyavailable": FeatureSupportStatusGenerallyAvailable,
		"invalid":            FeatureSupportStatusInvalid,
		"notsupported":       FeatureSupportStatusNotSupported,
		"privatepreview":     FeatureSupportStatusPrivatePreview,
		"publicpreview":      FeatureSupportStatusPublicPreview,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FeatureSupportStatus(input)
	return &out, nil
}

type FeatureType string

const (
	FeatureTypeDataSourceType FeatureType = "DataSourceType"
	FeatureTypeInvalid        FeatureType = "Invalid"
)

func PossibleValuesForFeatureType() []string {
	return []string{
		string(FeatureTypeDataSourceType),
		string(FeatureTypeInvalid),
	}
}

func (s *FeatureType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFeatureType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFeatureType(input string) (*FeatureType, error) {
	vals := map[string]FeatureType{
		"datasourcetype": FeatureTypeDataSourceType,
		"invalid":        FeatureTypeInvalid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FeatureType(input)
	return &out, nil
}
