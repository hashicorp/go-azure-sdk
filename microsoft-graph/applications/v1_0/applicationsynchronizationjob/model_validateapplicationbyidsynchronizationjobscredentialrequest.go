package applicationsynchronizationjob

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidateApplicationByIdSynchronizationJobsCredentialRequest struct {
	ApplicationIdentifier *string                                    `json:"applicationIdentifier,omitempty"`
	Credentials           *[]SynchronizationSecretKeyStringValuePair `json:"credentials,omitempty"`
	TemplateId            *string                                    `json:"templateId,omitempty"`
	UseSavedCredentials   *bool                                      `json:"useSavedCredentials,omitempty"`
}
