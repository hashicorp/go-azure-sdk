package synchronization

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AcquireSynchronizationAccessTokenRequest struct {
	Credentials *[]stable.SynchronizationSecretKeyStringValuePair `json:"credentials,omitempty"`
}
