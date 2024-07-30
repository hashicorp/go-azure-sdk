package synchronization

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AcquireApplicationSynchronizationAccessTokenRequest struct {
	Credentials *[]SynchronizationSecretKeyStringValuePair `json:"credentials,omitempty"`
}
