package contentkeypolicies

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentKeyPolicyPlayReadyLicense struct {
	AllowTestDevices       bool                                        `json:"allowTestDevices"`
	BeginDate              *string                                     `json:"beginDate,omitempty"`
	ContentKeyLocation     ContentKeyPolicyPlayReadyContentKeyLocation `json:"contentKeyLocation"`
	ContentType            ContentKeyPolicyPlayReadyContentType        `json:"contentType"`
	ExpirationDate         *string                                     `json:"expirationDate,omitempty"`
	GracePeriod            *string                                     `json:"gracePeriod,omitempty"`
	LicenseType            ContentKeyPolicyPlayReadyLicenseType        `json:"licenseType"`
	PlayRight              *ContentKeyPolicyPlayReadyPlayRight         `json:"playRight,omitempty"`
	RelativeBeginDate      *string                                     `json:"relativeBeginDate,omitempty"`
	RelativeExpirationDate *string                                     `json:"relativeExpirationDate,omitempty"`
}

var _ json.Unmarshaler = &ContentKeyPolicyPlayReadyLicense{}

func (s *ContentKeyPolicyPlayReadyLicense) UnmarshalJSON(bytes []byte) error {
	type alias ContentKeyPolicyPlayReadyLicense
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into ContentKeyPolicyPlayReadyLicense: %+v", err)
	}

	s.AllowTestDevices = decoded.AllowTestDevices
	s.BeginDate = decoded.BeginDate
	s.ContentType = decoded.ContentType
	s.ExpirationDate = decoded.ExpirationDate
	s.GracePeriod = decoded.GracePeriod
	s.LicenseType = decoded.LicenseType
	s.PlayRight = decoded.PlayRight
	s.RelativeBeginDate = decoded.RelativeBeginDate
	s.RelativeExpirationDate = decoded.RelativeExpirationDate

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ContentKeyPolicyPlayReadyLicense into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["contentKeyLocation"]; ok {
		impl, err := unmarshalContentKeyPolicyPlayReadyContentKeyLocationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ContentKeyLocation' for 'ContentKeyPolicyPlayReadyLicense': %+v", err)
		}
		s.ContentKeyLocation = impl
	}
	return nil
}
