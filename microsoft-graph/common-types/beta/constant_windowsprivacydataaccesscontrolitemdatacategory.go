package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPrivacyDataAccessControlItemDataCategory string

const (
	WindowsPrivacyDataAccessControlItemDataCategory_AccountInfo         WindowsPrivacyDataAccessControlItemDataCategory = "accountInfo"
	WindowsPrivacyDataAccessControlItemDataCategory_AppsRunInBackground WindowsPrivacyDataAccessControlItemDataCategory = "appsRunInBackground"
	WindowsPrivacyDataAccessControlItemDataCategory_Calendar            WindowsPrivacyDataAccessControlItemDataCategory = "calendar"
	WindowsPrivacyDataAccessControlItemDataCategory_CallHistory         WindowsPrivacyDataAccessControlItemDataCategory = "callHistory"
	WindowsPrivacyDataAccessControlItemDataCategory_Camera              WindowsPrivacyDataAccessControlItemDataCategory = "camera"
	WindowsPrivacyDataAccessControlItemDataCategory_Contacts            WindowsPrivacyDataAccessControlItemDataCategory = "contacts"
	WindowsPrivacyDataAccessControlItemDataCategory_DiagnosticsInfo     WindowsPrivacyDataAccessControlItemDataCategory = "diagnosticsInfo"
	WindowsPrivacyDataAccessControlItemDataCategory_Email               WindowsPrivacyDataAccessControlItemDataCategory = "email"
	WindowsPrivacyDataAccessControlItemDataCategory_Location            WindowsPrivacyDataAccessControlItemDataCategory = "location"
	WindowsPrivacyDataAccessControlItemDataCategory_Messaging           WindowsPrivacyDataAccessControlItemDataCategory = "messaging"
	WindowsPrivacyDataAccessControlItemDataCategory_Microphone          WindowsPrivacyDataAccessControlItemDataCategory = "microphone"
	WindowsPrivacyDataAccessControlItemDataCategory_Motion              WindowsPrivacyDataAccessControlItemDataCategory = "motion"
	WindowsPrivacyDataAccessControlItemDataCategory_NotConfigured       WindowsPrivacyDataAccessControlItemDataCategory = "notConfigured"
	WindowsPrivacyDataAccessControlItemDataCategory_Notifications       WindowsPrivacyDataAccessControlItemDataCategory = "notifications"
	WindowsPrivacyDataAccessControlItemDataCategory_Phone               WindowsPrivacyDataAccessControlItemDataCategory = "phone"
	WindowsPrivacyDataAccessControlItemDataCategory_Radios              WindowsPrivacyDataAccessControlItemDataCategory = "radios"
	WindowsPrivacyDataAccessControlItemDataCategory_SyncWithDevices     WindowsPrivacyDataAccessControlItemDataCategory = "syncWithDevices"
	WindowsPrivacyDataAccessControlItemDataCategory_Tasks               WindowsPrivacyDataAccessControlItemDataCategory = "tasks"
	WindowsPrivacyDataAccessControlItemDataCategory_TrustedDevices      WindowsPrivacyDataAccessControlItemDataCategory = "trustedDevices"
)

func PossibleValuesForWindowsPrivacyDataAccessControlItemDataCategory() []string {
	return []string{
		string(WindowsPrivacyDataAccessControlItemDataCategory_AccountInfo),
		string(WindowsPrivacyDataAccessControlItemDataCategory_AppsRunInBackground),
		string(WindowsPrivacyDataAccessControlItemDataCategory_Calendar),
		string(WindowsPrivacyDataAccessControlItemDataCategory_CallHistory),
		string(WindowsPrivacyDataAccessControlItemDataCategory_Camera),
		string(WindowsPrivacyDataAccessControlItemDataCategory_Contacts),
		string(WindowsPrivacyDataAccessControlItemDataCategory_DiagnosticsInfo),
		string(WindowsPrivacyDataAccessControlItemDataCategory_Email),
		string(WindowsPrivacyDataAccessControlItemDataCategory_Location),
		string(WindowsPrivacyDataAccessControlItemDataCategory_Messaging),
		string(WindowsPrivacyDataAccessControlItemDataCategory_Microphone),
		string(WindowsPrivacyDataAccessControlItemDataCategory_Motion),
		string(WindowsPrivacyDataAccessControlItemDataCategory_NotConfigured),
		string(WindowsPrivacyDataAccessControlItemDataCategory_Notifications),
		string(WindowsPrivacyDataAccessControlItemDataCategory_Phone),
		string(WindowsPrivacyDataAccessControlItemDataCategory_Radios),
		string(WindowsPrivacyDataAccessControlItemDataCategory_SyncWithDevices),
		string(WindowsPrivacyDataAccessControlItemDataCategory_Tasks),
		string(WindowsPrivacyDataAccessControlItemDataCategory_TrustedDevices),
	}
}

func (s *WindowsPrivacyDataAccessControlItemDataCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPrivacyDataAccessControlItemDataCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPrivacyDataAccessControlItemDataCategory(input string) (*WindowsPrivacyDataAccessControlItemDataCategory, error) {
	vals := map[string]WindowsPrivacyDataAccessControlItemDataCategory{
		"accountinfo":         WindowsPrivacyDataAccessControlItemDataCategory_AccountInfo,
		"appsruninbackground": WindowsPrivacyDataAccessControlItemDataCategory_AppsRunInBackground,
		"calendar":            WindowsPrivacyDataAccessControlItemDataCategory_Calendar,
		"callhistory":         WindowsPrivacyDataAccessControlItemDataCategory_CallHistory,
		"camera":              WindowsPrivacyDataAccessControlItemDataCategory_Camera,
		"contacts":            WindowsPrivacyDataAccessControlItemDataCategory_Contacts,
		"diagnosticsinfo":     WindowsPrivacyDataAccessControlItemDataCategory_DiagnosticsInfo,
		"email":               WindowsPrivacyDataAccessControlItemDataCategory_Email,
		"location":            WindowsPrivacyDataAccessControlItemDataCategory_Location,
		"messaging":           WindowsPrivacyDataAccessControlItemDataCategory_Messaging,
		"microphone":          WindowsPrivacyDataAccessControlItemDataCategory_Microphone,
		"motion":              WindowsPrivacyDataAccessControlItemDataCategory_Motion,
		"notconfigured":       WindowsPrivacyDataAccessControlItemDataCategory_NotConfigured,
		"notifications":       WindowsPrivacyDataAccessControlItemDataCategory_Notifications,
		"phone":               WindowsPrivacyDataAccessControlItemDataCategory_Phone,
		"radios":              WindowsPrivacyDataAccessControlItemDataCategory_Radios,
		"syncwithdevices":     WindowsPrivacyDataAccessControlItemDataCategory_SyncWithDevices,
		"tasks":               WindowsPrivacyDataAccessControlItemDataCategory_Tasks,
		"trusteddevices":      WindowsPrivacyDataAccessControlItemDataCategory_TrustedDevices,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPrivacyDataAccessControlItemDataCategory(input)
	return &out, nil
}
