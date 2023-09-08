package serviceprincipalsynchronization

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AcquireServicePrincipalByIdSynchronizationAccessTokenRequest struct {
	Credentials *[]SynchronizationSecretKeyStringValuePair `json:"credentials,omitempty"`
}
