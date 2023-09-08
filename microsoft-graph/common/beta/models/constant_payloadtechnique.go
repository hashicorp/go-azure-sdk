package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadTechnique string

const (
	PayloadTechniqueattachmentMalware    PayloadTechnique = "AttachmentMalware"
	PayloadTechniquecredentialHarvesting PayloadTechnique = "CredentialHarvesting"
	PayloadTechniquedriveByUrl           PayloadTechnique = "DriveByUrl"
	PayloadTechniquelinkInAttachment     PayloadTechnique = "LinkInAttachment"
	PayloadTechniquelinkToMalwareFile    PayloadTechnique = "LinkToMalwareFile"
	PayloadTechniqueoAuthConsentGrant    PayloadTechnique = "OAuthConsentGrant"
	PayloadTechniqueunknown              PayloadTechnique = "Unknown"
)

func PossibleValuesForPayloadTechnique() []string {
	return []string{
		string(PayloadTechniqueattachmentMalware),
		string(PayloadTechniquecredentialHarvesting),
		string(PayloadTechniquedriveByUrl),
		string(PayloadTechniquelinkInAttachment),
		string(PayloadTechniquelinkToMalwareFile),
		string(PayloadTechniqueoAuthConsentGrant),
		string(PayloadTechniqueunknown),
	}
}

func parsePayloadTechnique(input string) (*PayloadTechnique, error) {
	vals := map[string]PayloadTechnique{
		"attachmentmalware":    PayloadTechniqueattachmentMalware,
		"credentialharvesting": PayloadTechniquecredentialHarvesting,
		"drivebyurl":           PayloadTechniquedriveByUrl,
		"linkinattachment":     PayloadTechniquelinkInAttachment,
		"linktomalwarefile":    PayloadTechniquelinkToMalwareFile,
		"oauthconsentgrant":    PayloadTechniqueoAuthConsentGrant,
		"unknown":              PayloadTechniqueunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PayloadTechnique(input)
	return &out, nil
}
