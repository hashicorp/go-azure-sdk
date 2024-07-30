package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceStartupHistoryRestartCategory string

const (
	UserExperienceAnalyticsDeviceStartupHistoryRestartCategory_BlueScreen            UserExperienceAnalyticsDeviceStartupHistoryRestartCategory = "blueScreen"
	UserExperienceAnalyticsDeviceStartupHistoryRestartCategory_BootError             UserExperienceAnalyticsDeviceStartupHistoryRestartCategory = "bootError"
	UserExperienceAnalyticsDeviceStartupHistoryRestartCategory_LongPowerButtonPress  UserExperienceAnalyticsDeviceStartupHistoryRestartCategory = "longPowerButtonPress"
	UserExperienceAnalyticsDeviceStartupHistoryRestartCategory_RestartWithUpdate     UserExperienceAnalyticsDeviceStartupHistoryRestartCategory = "restartWithUpdate"
	UserExperienceAnalyticsDeviceStartupHistoryRestartCategory_RestartWithoutUpdate  UserExperienceAnalyticsDeviceStartupHistoryRestartCategory = "restartWithoutUpdate"
	UserExperienceAnalyticsDeviceStartupHistoryRestartCategory_ShutdownWithUpdate    UserExperienceAnalyticsDeviceStartupHistoryRestartCategory = "shutdownWithUpdate"
	UserExperienceAnalyticsDeviceStartupHistoryRestartCategory_ShutdownWithoutUpdate UserExperienceAnalyticsDeviceStartupHistoryRestartCategory = "shutdownWithoutUpdate"
	UserExperienceAnalyticsDeviceStartupHistoryRestartCategory_Unknown               UserExperienceAnalyticsDeviceStartupHistoryRestartCategory = "unknown"
	UserExperienceAnalyticsDeviceStartupHistoryRestartCategory_Update                UserExperienceAnalyticsDeviceStartupHistoryRestartCategory = "update"
)

func PossibleValuesForUserExperienceAnalyticsDeviceStartupHistoryRestartCategory() []string {
	return []string{
		string(UserExperienceAnalyticsDeviceStartupHistoryRestartCategory_BlueScreen),
		string(UserExperienceAnalyticsDeviceStartupHistoryRestartCategory_BootError),
		string(UserExperienceAnalyticsDeviceStartupHistoryRestartCategory_LongPowerButtonPress),
		string(UserExperienceAnalyticsDeviceStartupHistoryRestartCategory_RestartWithUpdate),
		string(UserExperienceAnalyticsDeviceStartupHistoryRestartCategory_RestartWithoutUpdate),
		string(UserExperienceAnalyticsDeviceStartupHistoryRestartCategory_ShutdownWithUpdate),
		string(UserExperienceAnalyticsDeviceStartupHistoryRestartCategory_ShutdownWithoutUpdate),
		string(UserExperienceAnalyticsDeviceStartupHistoryRestartCategory_Unknown),
		string(UserExperienceAnalyticsDeviceStartupHistoryRestartCategory_Update),
	}
}

func (s *UserExperienceAnalyticsDeviceStartupHistoryRestartCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsDeviceStartupHistoryRestartCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsDeviceStartupHistoryRestartCategory(input string) (*UserExperienceAnalyticsDeviceStartupHistoryRestartCategory, error) {
	vals := map[string]UserExperienceAnalyticsDeviceStartupHistoryRestartCategory{
		"bluescreen":            UserExperienceAnalyticsDeviceStartupHistoryRestartCategory_BlueScreen,
		"booterror":             UserExperienceAnalyticsDeviceStartupHistoryRestartCategory_BootError,
		"longpowerbuttonpress":  UserExperienceAnalyticsDeviceStartupHistoryRestartCategory_LongPowerButtonPress,
		"restartwithupdate":     UserExperienceAnalyticsDeviceStartupHistoryRestartCategory_RestartWithUpdate,
		"restartwithoutupdate":  UserExperienceAnalyticsDeviceStartupHistoryRestartCategory_RestartWithoutUpdate,
		"shutdownwithupdate":    UserExperienceAnalyticsDeviceStartupHistoryRestartCategory_ShutdownWithUpdate,
		"shutdownwithoutupdate": UserExperienceAnalyticsDeviceStartupHistoryRestartCategory_ShutdownWithoutUpdate,
		"unknown":               UserExperienceAnalyticsDeviceStartupHistoryRestartCategory_Unknown,
		"update":                UserExperienceAnalyticsDeviceStartupHistoryRestartCategory_Update,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsDeviceStartupHistoryRestartCategory(input)
	return &out, nil
}
