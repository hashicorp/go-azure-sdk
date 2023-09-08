package applicationsynchronization

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AcquireApplicationByIdSynchronizationAccessTokenRequest struct {
	Credentials *[]SynchronizationSecretKeyStringValuePair `json:"credentials,omitempty"`
}
