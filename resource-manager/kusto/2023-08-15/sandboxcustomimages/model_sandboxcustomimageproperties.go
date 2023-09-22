package sandboxcustomimages

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SandboxCustomImageProperties struct {
	Language                Language           `json:"language"`
	LanguageVersion         string             `json:"languageVersion"`
	ProvisioningState       *ProvisioningState `json:"provisioningState,omitempty"`
	RequirementsFileContent *string            `json:"requirementsFileContent,omitempty"`
}
