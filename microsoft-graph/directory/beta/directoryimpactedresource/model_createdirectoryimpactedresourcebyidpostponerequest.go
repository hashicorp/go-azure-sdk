package directoryimpactedresource

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateDirectoryImpactedResourceByIdPostponeRequest struct {
	PostponeUntilDateTime *string `json:"postponeUntilDateTime,omitempty"`
}
