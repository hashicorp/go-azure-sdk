package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoverySearchExportOperationAdditionalOptions string

const (
	SecurityEdiscoverySearchExportOperationAdditionalOptions_AllDocumentVersions         SecurityEdiscoverySearchExportOperationAdditionalOptions = "allDocumentVersions"
	SecurityEdiscoverySearchExportOperationAdditionalOptions_CloudAttachments            SecurityEdiscoverySearchExportOperationAdditionalOptions = "cloudAttachments"
	SecurityEdiscoverySearchExportOperationAdditionalOptions_ListAttachments             SecurityEdiscoverySearchExportOperationAdditionalOptions = "listAttachments"
	SecurityEdiscoverySearchExportOperationAdditionalOptions_None                        SecurityEdiscoverySearchExportOperationAdditionalOptions = "none"
	SecurityEdiscoverySearchExportOperationAdditionalOptions_SubfolderContents           SecurityEdiscoverySearchExportOperationAdditionalOptions = "subfolderContents"
	SecurityEdiscoverySearchExportOperationAdditionalOptions_TeamsAndYammerConversations SecurityEdiscoverySearchExportOperationAdditionalOptions = "teamsAndYammerConversations"
)

func PossibleValuesForSecurityEdiscoverySearchExportOperationAdditionalOptions() []string {
	return []string{
		string(SecurityEdiscoverySearchExportOperationAdditionalOptions_AllDocumentVersions),
		string(SecurityEdiscoverySearchExportOperationAdditionalOptions_CloudAttachments),
		string(SecurityEdiscoverySearchExportOperationAdditionalOptions_ListAttachments),
		string(SecurityEdiscoverySearchExportOperationAdditionalOptions_None),
		string(SecurityEdiscoverySearchExportOperationAdditionalOptions_SubfolderContents),
		string(SecurityEdiscoverySearchExportOperationAdditionalOptions_TeamsAndYammerConversations),
	}
}

func (s *SecurityEdiscoverySearchExportOperationAdditionalOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoverySearchExportOperationAdditionalOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoverySearchExportOperationAdditionalOptions(input string) (*SecurityEdiscoverySearchExportOperationAdditionalOptions, error) {
	vals := map[string]SecurityEdiscoverySearchExportOperationAdditionalOptions{
		"alldocumentversions":         SecurityEdiscoverySearchExportOperationAdditionalOptions_AllDocumentVersions,
		"cloudattachments":            SecurityEdiscoverySearchExportOperationAdditionalOptions_CloudAttachments,
		"listattachments":             SecurityEdiscoverySearchExportOperationAdditionalOptions_ListAttachments,
		"none":                        SecurityEdiscoverySearchExportOperationAdditionalOptions_None,
		"subfoldercontents":           SecurityEdiscoverySearchExportOperationAdditionalOptions_SubfolderContents,
		"teamsandyammerconversations": SecurityEdiscoverySearchExportOperationAdditionalOptions_TeamsAndYammerConversations,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoverySearchExportOperationAdditionalOptions(input)
	return &out, nil
}
