package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationMetadataEntryKey string

const (
	SynchronizationMetadataEntryKeyConfigurationFields                      SynchronizationMetadataEntryKey = "ConfigurationFields"
	SynchronizationMetadataEntryKeyGalleryApplicationIdentifier             SynchronizationMetadataEntryKey = "GalleryApplicationIdentifier"
	SynchronizationMetadataEntryKeyGalleryApplicationKey                    SynchronizationMetadataEntryKey = "GalleryApplicationKey"
	SynchronizationMetadataEntryKeyIsOAuthEnabled                           SynchronizationMetadataEntryKey = "IsOAuthEnabled"
	SynchronizationMetadataEntryKeyIsSynchronizationAgentAssignmentRequired SynchronizationMetadataEntryKey = "IsSynchronizationAgentAssignmentRequired"
	SynchronizationMetadataEntryKeyIsSynchronizationAgentRequired           SynchronizationMetadataEntryKey = "IsSynchronizationAgentRequired"
	SynchronizationMetadataEntryKeyIsSynchronizationInPreview               SynchronizationMetadataEntryKey = "IsSynchronizationInPreview"
	SynchronizationMetadataEntryKeyOAuthSettings                            SynchronizationMetadataEntryKey = "OAuthSettings"
	SynchronizationMetadataEntryKeySynchronizationLearnMoreIbizaFwLink      SynchronizationMetadataEntryKey = "SynchronizationLearnMoreIbizaFwLink"
)

func PossibleValuesForSynchronizationMetadataEntryKey() []string {
	return []string{
		string(SynchronizationMetadataEntryKeyConfigurationFields),
		string(SynchronizationMetadataEntryKeyGalleryApplicationIdentifier),
		string(SynchronizationMetadataEntryKeyGalleryApplicationKey),
		string(SynchronizationMetadataEntryKeyIsOAuthEnabled),
		string(SynchronizationMetadataEntryKeyIsSynchronizationAgentAssignmentRequired),
		string(SynchronizationMetadataEntryKeyIsSynchronizationAgentRequired),
		string(SynchronizationMetadataEntryKeyIsSynchronizationInPreview),
		string(SynchronizationMetadataEntryKeyOAuthSettings),
		string(SynchronizationMetadataEntryKeySynchronizationLearnMoreIbizaFwLink),
	}
}

func parseSynchronizationMetadataEntryKey(input string) (*SynchronizationMetadataEntryKey, error) {
	vals := map[string]SynchronizationMetadataEntryKey{
		"configurationfields":                      SynchronizationMetadataEntryKeyConfigurationFields,
		"galleryapplicationidentifier":             SynchronizationMetadataEntryKeyGalleryApplicationIdentifier,
		"galleryapplicationkey":                    SynchronizationMetadataEntryKeyGalleryApplicationKey,
		"isoauthenabled":                           SynchronizationMetadataEntryKeyIsOAuthEnabled,
		"issynchronizationagentassignmentrequired": SynchronizationMetadataEntryKeyIsSynchronizationAgentAssignmentRequired,
		"issynchronizationagentrequired":           SynchronizationMetadataEntryKeyIsSynchronizationAgentRequired,
		"issynchronizationinpreview":               SynchronizationMetadataEntryKeyIsSynchronizationInPreview,
		"oauthsettings":                            SynchronizationMetadataEntryKeyOAuthSettings,
		"synchronizationlearnmoreibizafwlink":      SynchronizationMetadataEntryKeySynchronizationLearnMoreIbizaFwLink,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SynchronizationMetadataEntryKey(input)
	return &out, nil
}
