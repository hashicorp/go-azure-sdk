package elasticpools

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ElasticPoolLicenseType string

const (
	ElasticPoolLicenseTypeBasePrice       ElasticPoolLicenseType = "BasePrice"
	ElasticPoolLicenseTypeLicenseIncluded ElasticPoolLicenseType = "LicenseIncluded"
)

func PossibleValuesForElasticPoolLicenseType() []string {
	return []string{
		string(ElasticPoolLicenseTypeBasePrice),
		string(ElasticPoolLicenseTypeLicenseIncluded),
	}
}

func (s *ElasticPoolLicenseType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseElasticPoolLicenseType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseElasticPoolLicenseType(input string) (*ElasticPoolLicenseType, error) {
	vals := map[string]ElasticPoolLicenseType{
		"baseprice":       ElasticPoolLicenseTypeBasePrice,
		"licenseincluded": ElasticPoolLicenseTypeLicenseIncluded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ElasticPoolLicenseType(input)
	return &out, nil
}

type ElasticPoolState string

const (
	ElasticPoolStateCreating ElasticPoolState = "Creating"
	ElasticPoolStateDisabled ElasticPoolState = "Disabled"
	ElasticPoolStateReady    ElasticPoolState = "Ready"
)

func PossibleValuesForElasticPoolState() []string {
	return []string{
		string(ElasticPoolStateCreating),
		string(ElasticPoolStateDisabled),
		string(ElasticPoolStateReady),
	}
}

func (s *ElasticPoolState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseElasticPoolState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseElasticPoolState(input string) (*ElasticPoolState, error) {
	vals := map[string]ElasticPoolState{
		"creating": ElasticPoolStateCreating,
		"disabled": ElasticPoolStateDisabled,
		"ready":    ElasticPoolStateReady,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ElasticPoolState(input)
	return &out, nil
}
