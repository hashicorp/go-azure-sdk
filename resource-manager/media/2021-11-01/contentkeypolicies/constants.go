package contentkeypolicies

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentKeyPolicyFairPlayRentalAndLeaseKeyType string

const (
	ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeDualExpiry          ContentKeyPolicyFairPlayRentalAndLeaseKeyType = "DualExpiry"
	ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePersistentLimited   ContentKeyPolicyFairPlayRentalAndLeaseKeyType = "PersistentLimited"
	ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePersistentUnlimited ContentKeyPolicyFairPlayRentalAndLeaseKeyType = "PersistentUnlimited"
	ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeUndefined           ContentKeyPolicyFairPlayRentalAndLeaseKeyType = "Undefined"
	ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeUnknown             ContentKeyPolicyFairPlayRentalAndLeaseKeyType = "Unknown"
)

func PossibleValuesForContentKeyPolicyFairPlayRentalAndLeaseKeyType() []string {
	return []string{
		string(ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeDualExpiry),
		string(ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePersistentLimited),
		string(ContentKeyPolicyFairPlayRentalAndLeaseKeyTypePersistentUnlimited),
		string(ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeUndefined),
		string(ContentKeyPolicyFairPlayRentalAndLeaseKeyTypeUnknown),
	}
}

func (s *ContentKeyPolicyFairPlayRentalAndLeaseKeyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForContentKeyPolicyFairPlayRentalAndLeaseKeyType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ContentKeyPolicyFairPlayRentalAndLeaseKeyType(decoded)
	return nil
}

type ContentKeyPolicyPlayReadyContentType string

const (
	ContentKeyPolicyPlayReadyContentTypeUltraVioletDownload  ContentKeyPolicyPlayReadyContentType = "UltraVioletDownload"
	ContentKeyPolicyPlayReadyContentTypeUltraVioletStreaming ContentKeyPolicyPlayReadyContentType = "UltraVioletStreaming"
	ContentKeyPolicyPlayReadyContentTypeUnknown              ContentKeyPolicyPlayReadyContentType = "Unknown"
	ContentKeyPolicyPlayReadyContentTypeUnspecified          ContentKeyPolicyPlayReadyContentType = "Unspecified"
)

func PossibleValuesForContentKeyPolicyPlayReadyContentType() []string {
	return []string{
		string(ContentKeyPolicyPlayReadyContentTypeUltraVioletDownload),
		string(ContentKeyPolicyPlayReadyContentTypeUltraVioletStreaming),
		string(ContentKeyPolicyPlayReadyContentTypeUnknown),
		string(ContentKeyPolicyPlayReadyContentTypeUnspecified),
	}
}

func (s *ContentKeyPolicyPlayReadyContentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForContentKeyPolicyPlayReadyContentType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ContentKeyPolicyPlayReadyContentType(decoded)
	return nil
}

type ContentKeyPolicyPlayReadyLicenseType string

const (
	ContentKeyPolicyPlayReadyLicenseTypeNonPersistent ContentKeyPolicyPlayReadyLicenseType = "NonPersistent"
	ContentKeyPolicyPlayReadyLicenseTypePersistent    ContentKeyPolicyPlayReadyLicenseType = "Persistent"
	ContentKeyPolicyPlayReadyLicenseTypeUnknown       ContentKeyPolicyPlayReadyLicenseType = "Unknown"
)

func PossibleValuesForContentKeyPolicyPlayReadyLicenseType() []string {
	return []string{
		string(ContentKeyPolicyPlayReadyLicenseTypeNonPersistent),
		string(ContentKeyPolicyPlayReadyLicenseTypePersistent),
		string(ContentKeyPolicyPlayReadyLicenseTypeUnknown),
	}
}

func (s *ContentKeyPolicyPlayReadyLicenseType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForContentKeyPolicyPlayReadyLicenseType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ContentKeyPolicyPlayReadyLicenseType(decoded)
	return nil
}

type ContentKeyPolicyPlayReadyUnknownOutputPassingOption string

const (
	ContentKeyPolicyPlayReadyUnknownOutputPassingOptionAllowed                      ContentKeyPolicyPlayReadyUnknownOutputPassingOption = "Allowed"
	ContentKeyPolicyPlayReadyUnknownOutputPassingOptionAllowedWithVideoConstriction ContentKeyPolicyPlayReadyUnknownOutputPassingOption = "AllowedWithVideoConstriction"
	ContentKeyPolicyPlayReadyUnknownOutputPassingOptionNotAllowed                   ContentKeyPolicyPlayReadyUnknownOutputPassingOption = "NotAllowed"
	ContentKeyPolicyPlayReadyUnknownOutputPassingOptionUnknown                      ContentKeyPolicyPlayReadyUnknownOutputPassingOption = "Unknown"
)

func PossibleValuesForContentKeyPolicyPlayReadyUnknownOutputPassingOption() []string {
	return []string{
		string(ContentKeyPolicyPlayReadyUnknownOutputPassingOptionAllowed),
		string(ContentKeyPolicyPlayReadyUnknownOutputPassingOptionAllowedWithVideoConstriction),
		string(ContentKeyPolicyPlayReadyUnknownOutputPassingOptionNotAllowed),
		string(ContentKeyPolicyPlayReadyUnknownOutputPassingOptionUnknown),
	}
}

func (s *ContentKeyPolicyPlayReadyUnknownOutputPassingOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForContentKeyPolicyPlayReadyUnknownOutputPassingOption() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ContentKeyPolicyPlayReadyUnknownOutputPassingOption(decoded)
	return nil
}

type ContentKeyPolicyRestrictionTokenType string

const (
	ContentKeyPolicyRestrictionTokenTypeJwt     ContentKeyPolicyRestrictionTokenType = "Jwt"
	ContentKeyPolicyRestrictionTokenTypeSwt     ContentKeyPolicyRestrictionTokenType = "Swt"
	ContentKeyPolicyRestrictionTokenTypeUnknown ContentKeyPolicyRestrictionTokenType = "Unknown"
)

func PossibleValuesForContentKeyPolicyRestrictionTokenType() []string {
	return []string{
		string(ContentKeyPolicyRestrictionTokenTypeJwt),
		string(ContentKeyPolicyRestrictionTokenTypeSwt),
		string(ContentKeyPolicyRestrictionTokenTypeUnknown),
	}
}

func (s *ContentKeyPolicyRestrictionTokenType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForContentKeyPolicyRestrictionTokenType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ContentKeyPolicyRestrictionTokenType(decoded)
	return nil
}
