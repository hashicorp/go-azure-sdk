package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationMetadataEntryKey string

const (
	SynchronizationMetadataEntryKey_ConfigurationFields                      SynchronizationMetadataEntryKey = "ConfigurationFields"
	SynchronizationMetadataEntryKey_GalleryApplicationIdentifier             SynchronizationMetadataEntryKey = "GalleryApplicationIdentifier"
	SynchronizationMetadataEntryKey_GalleryApplicationKey                    SynchronizationMetadataEntryKey = "GalleryApplicationKey"
	SynchronizationMetadataEntryKey_IsOAuthEnabled                           SynchronizationMetadataEntryKey = "IsOAuthEnabled"
	SynchronizationMetadataEntryKey_IsSynchronizationAgentAssignmentRequired SynchronizationMetadataEntryKey = "IsSynchronizationAgentAssignmentRequired"
	SynchronizationMetadataEntryKey_IsSynchronizationAgentRequired           SynchronizationMetadataEntryKey = "IsSynchronizationAgentRequired"
	SynchronizationMetadataEntryKey_IsSynchronizationInPreview               SynchronizationMetadataEntryKey = "IsSynchronizationInPreview"
	SynchronizationMetadataEntryKey_OAuthSettings                            SynchronizationMetadataEntryKey = "OAuthSettings"
	SynchronizationMetadataEntryKey_SynchronizationLearnMoreIbizaFwLink      SynchronizationMetadataEntryKey = "SynchronizationLearnMoreIbizaFwLink"
)

func PossibleValuesForSynchronizationMetadataEntryKey() []string {
	return []string{
		string(SynchronizationMetadataEntryKey_ConfigurationFields),
		string(SynchronizationMetadataEntryKey_GalleryApplicationIdentifier),
		string(SynchronizationMetadataEntryKey_GalleryApplicationKey),
		string(SynchronizationMetadataEntryKey_IsOAuthEnabled),
		string(SynchronizationMetadataEntryKey_IsSynchronizationAgentAssignmentRequired),
		string(SynchronizationMetadataEntryKey_IsSynchronizationAgentRequired),
		string(SynchronizationMetadataEntryKey_IsSynchronizationInPreview),
		string(SynchronizationMetadataEntryKey_OAuthSettings),
		string(SynchronizationMetadataEntryKey_SynchronizationLearnMoreIbizaFwLink),
	}
}

func (s *SynchronizationMetadataEntryKey) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSynchronizationMetadataEntryKey(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSynchronizationMetadataEntryKey(input string) (*SynchronizationMetadataEntryKey, error) {
	vals := map[string]SynchronizationMetadataEntryKey{
		"configurationfields":                      SynchronizationMetadataEntryKey_ConfigurationFields,
		"galleryapplicationidentifier":             SynchronizationMetadataEntryKey_GalleryApplicationIdentifier,
		"galleryapplicationkey":                    SynchronizationMetadataEntryKey_GalleryApplicationKey,
		"isoauthenabled":                           SynchronizationMetadataEntryKey_IsOAuthEnabled,
		"issynchronizationagentassignmentrequired": SynchronizationMetadataEntryKey_IsSynchronizationAgentAssignmentRequired,
		"issynchronizationagentrequired":           SynchronizationMetadataEntryKey_IsSynchronizationAgentRequired,
		"issynchronizationinpreview":               SynchronizationMetadataEntryKey_IsSynchronizationInPreview,
		"oauthsettings":                            SynchronizationMetadataEntryKey_OAuthSettings,
		"synchronizationlearnmoreibizafwlink":      SynchronizationMetadataEntryKey_SynchronizationLearnMoreIbizaFwLink,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SynchronizationMetadataEntryKey(input)
	return &out, nil
}
