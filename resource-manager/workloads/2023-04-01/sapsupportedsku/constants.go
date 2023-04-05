package sapsupportedsku

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SAPDatabaseType string

const (
	SAPDatabaseTypeDBTwo SAPDatabaseType = "DB2"
	SAPDatabaseTypeHANA  SAPDatabaseType = "HANA"
)

func PossibleValuesForSAPDatabaseType() []string {
	return []string{
		string(SAPDatabaseTypeDBTwo),
		string(SAPDatabaseTypeHANA),
	}
}

func (s *SAPDatabaseType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSAPDatabaseType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SAPDatabaseType(decoded)
	return nil
}

type SAPDeploymentType string

const (
	SAPDeploymentTypeSingleServer SAPDeploymentType = "SingleServer"
	SAPDeploymentTypeThreeTier    SAPDeploymentType = "ThreeTier"
)

func PossibleValuesForSAPDeploymentType() []string {
	return []string{
		string(SAPDeploymentTypeSingleServer),
		string(SAPDeploymentTypeThreeTier),
	}
}

func (s *SAPDeploymentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSAPDeploymentType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SAPDeploymentType(decoded)
	return nil
}

type SAPEnvironmentType string

const (
	SAPEnvironmentTypeNonProd SAPEnvironmentType = "NonProd"
	SAPEnvironmentTypeProd    SAPEnvironmentType = "Prod"
)

func PossibleValuesForSAPEnvironmentType() []string {
	return []string{
		string(SAPEnvironmentTypeNonProd),
		string(SAPEnvironmentTypeProd),
	}
}

func (s *SAPEnvironmentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSAPEnvironmentType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SAPEnvironmentType(decoded)
	return nil
}

type SAPHighAvailabilityType string

const (
	SAPHighAvailabilityTypeAvailabilitySet  SAPHighAvailabilityType = "AvailabilitySet"
	SAPHighAvailabilityTypeAvailabilityZone SAPHighAvailabilityType = "AvailabilityZone"
)

func PossibleValuesForSAPHighAvailabilityType() []string {
	return []string{
		string(SAPHighAvailabilityTypeAvailabilitySet),
		string(SAPHighAvailabilityTypeAvailabilityZone),
	}
}

func (s *SAPHighAvailabilityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSAPHighAvailabilityType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SAPHighAvailabilityType(decoded)
	return nil
}

type SAPProductType string

const (
	SAPProductTypeECC       SAPProductType = "ECC"
	SAPProductTypeOther     SAPProductType = "Other"
	SAPProductTypeSFourHANA SAPProductType = "S4HANA"
)

func PossibleValuesForSAPProductType() []string {
	return []string{
		string(SAPProductTypeECC),
		string(SAPProductTypeOther),
		string(SAPProductTypeSFourHANA),
	}
}

func (s *SAPProductType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForSAPProductType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = SAPProductType(decoded)
	return nil
}
