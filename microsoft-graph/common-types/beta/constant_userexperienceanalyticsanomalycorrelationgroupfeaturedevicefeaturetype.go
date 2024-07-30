package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType string

const (
	UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType_Application  UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType = "application"
	UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType_Driver       UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType = "driver"
	UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType_Manufacturer UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType = "manufacturer"
	UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType_Model        UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType = "model"
	UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType_OsVersion    UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType = "osVersion"
)

func PossibleValuesForUserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType() []string {
	return []string{
		string(UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType_Application),
		string(UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType_Driver),
		string(UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType_Manufacturer),
		string(UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType_Model),
		string(UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType_OsVersion),
	}
}

func (s *UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType(input string) (*UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType, error) {
	vals := map[string]UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType{
		"application":  UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType_Application,
		"driver":       UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType_Driver,
		"manufacturer": UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType_Manufacturer,
		"model":        UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType_Model,
		"osversion":    UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType_OsVersion,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType(input)
	return &out, nil
}
