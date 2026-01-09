package accessreviewdecisioninstancedecision

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateAccessReviewDecisionInstanceDecisionRecordAllDecisionRequest struct {
	Decision      nullable.Type[string] `json:"decision,omitempty"`
	Justification nullable.Type[string] `json:"justification,omitempty"`
	PrincipalId   nullable.Type[string] `json:"principalId,omitempty"`
	ResourceId    nullable.Type[string] `json:"resourceId,omitempty"`
}
