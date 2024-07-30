package impactedresource

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateImpactedResourcePostponeRequest struct {
	PostponeUntilDateTime *string `json:"postponeUntilDateTime,omitempty"`
}
