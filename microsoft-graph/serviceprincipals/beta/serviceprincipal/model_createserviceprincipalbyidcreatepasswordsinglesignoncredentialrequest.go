package serviceprincipal

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateServicePrincipalByIdCreatePasswordSingleSignOnCredentialRequest struct {
	Credentials *[]Credential `json:"credentials,omitempty"`
	Id          *string       `json:"id,omitempty"`
}
