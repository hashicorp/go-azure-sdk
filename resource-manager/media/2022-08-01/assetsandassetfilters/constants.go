package assetsandassetfilters

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssetContainerPermission string

const (
	AssetContainerPermissionRead            AssetContainerPermission = "Read"
	AssetContainerPermissionReadWrite       AssetContainerPermission = "ReadWrite"
	AssetContainerPermissionReadWriteDelete AssetContainerPermission = "ReadWriteDelete"
)

func PossibleValuesForAssetContainerPermission() []string {
	return []string{
		string(AssetContainerPermissionRead),
		string(AssetContainerPermissionReadWrite),
		string(AssetContainerPermissionReadWriteDelete),
	}
}

func (s *AssetContainerPermission) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForAssetContainerPermission() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = AssetContainerPermission(decoded)
	return nil
}

type AssetStorageEncryptionFormat string

const (
	AssetStorageEncryptionFormatMediaStorageClientEncryption AssetStorageEncryptionFormat = "MediaStorageClientEncryption"
	AssetStorageEncryptionFormatNone                         AssetStorageEncryptionFormat = "None"
)

func PossibleValuesForAssetStorageEncryptionFormat() []string {
	return []string{
		string(AssetStorageEncryptionFormatMediaStorageClientEncryption),
		string(AssetStorageEncryptionFormatNone),
	}
}

func (s *AssetStorageEncryptionFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForAssetStorageEncryptionFormat() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = AssetStorageEncryptionFormat(decoded)
	return nil
}

type FilterTrackPropertyCompareOperation string

const (
	FilterTrackPropertyCompareOperationEqual    FilterTrackPropertyCompareOperation = "Equal"
	FilterTrackPropertyCompareOperationNotEqual FilterTrackPropertyCompareOperation = "NotEqual"
)

func PossibleValuesForFilterTrackPropertyCompareOperation() []string {
	return []string{
		string(FilterTrackPropertyCompareOperationEqual),
		string(FilterTrackPropertyCompareOperationNotEqual),
	}
}

func (s *FilterTrackPropertyCompareOperation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForFilterTrackPropertyCompareOperation() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = FilterTrackPropertyCompareOperation(decoded)
	return nil
}

type FilterTrackPropertyType string

const (
	FilterTrackPropertyTypeBitrate  FilterTrackPropertyType = "Bitrate"
	FilterTrackPropertyTypeFourCC   FilterTrackPropertyType = "FourCC"
	FilterTrackPropertyTypeLanguage FilterTrackPropertyType = "Language"
	FilterTrackPropertyTypeName     FilterTrackPropertyType = "Name"
	FilterTrackPropertyTypeType     FilterTrackPropertyType = "Type"
	FilterTrackPropertyTypeUnknown  FilterTrackPropertyType = "Unknown"
)

func PossibleValuesForFilterTrackPropertyType() []string {
	return []string{
		string(FilterTrackPropertyTypeBitrate),
		string(FilterTrackPropertyTypeFourCC),
		string(FilterTrackPropertyTypeLanguage),
		string(FilterTrackPropertyTypeName),
		string(FilterTrackPropertyTypeType),
		string(FilterTrackPropertyTypeUnknown),
	}
}

func (s *FilterTrackPropertyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForFilterTrackPropertyType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = FilterTrackPropertyType(decoded)
	return nil
}

type ProvisioningState string

const (
	ProvisioningStateFailed     ProvisioningState = "Failed"
	ProvisioningStateInProgress ProvisioningState = "InProgress"
	ProvisioningStateSucceeded  ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateFailed),
		string(ProvisioningStateInProgress),
		string(ProvisioningStateSucceeded),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForProvisioningState() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ProvisioningState(decoded)
	return nil
}

type Visibility string

const (
	VisibilityHidden  Visibility = "Hidden"
	VisibilityVisible Visibility = "Visible"
)

func PossibleValuesForVisibility() []string {
	return []string{
		string(VisibilityHidden),
		string(VisibilityVisible),
	}
}

func (s *Visibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForVisibility() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = Visibility(decoded)
	return nil
}
