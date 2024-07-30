package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadTechnique string

const (
	PayloadTechnique_AttachmentMalware    PayloadTechnique = "attachmentMalware"
	PayloadTechnique_CredentialHarvesting PayloadTechnique = "credentialHarvesting"
	PayloadTechnique_DriveByUrl           PayloadTechnique = "driveByUrl"
	PayloadTechnique_LinkInAttachment     PayloadTechnique = "linkInAttachment"
	PayloadTechnique_LinkToMalwareFile    PayloadTechnique = "linkToMalwareFile"
	PayloadTechnique_OAuthConsentGrant    PayloadTechnique = "oAuthConsentGrant"
	PayloadTechnique_Unknown              PayloadTechnique = "unknown"
)

func PossibleValuesForPayloadTechnique() []string {
	return []string{
		string(PayloadTechnique_AttachmentMalware),
		string(PayloadTechnique_CredentialHarvesting),
		string(PayloadTechnique_DriveByUrl),
		string(PayloadTechnique_LinkInAttachment),
		string(PayloadTechnique_LinkToMalwareFile),
		string(PayloadTechnique_OAuthConsentGrant),
		string(PayloadTechnique_Unknown),
	}
}

func (s *PayloadTechnique) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePayloadTechnique(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePayloadTechnique(input string) (*PayloadTechnique, error) {
	vals := map[string]PayloadTechnique{
		"attachmentmalware":    PayloadTechnique_AttachmentMalware,
		"credentialharvesting": PayloadTechnique_CredentialHarvesting,
		"drivebyurl":           PayloadTechnique_DriveByUrl,
		"linkinattachment":     PayloadTechnique_LinkInAttachment,
		"linktomalwarefile":    PayloadTechnique_LinkToMalwareFile,
		"oauthconsentgrant":    PayloadTechnique_OAuthConsentGrant,
		"unknown":              PayloadTechnique_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PayloadTechnique(input)
	return &out, nil
}
