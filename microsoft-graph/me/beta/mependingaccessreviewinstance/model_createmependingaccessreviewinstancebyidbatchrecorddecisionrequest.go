package mependingaccessreviewinstance

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMePendingAccessReviewInstanceByIdBatchRecordDecisionRequest struct {
	Decision      *string `json:"decision,omitempty"`
	Justification *string `json:"justification,omitempty"`
	PrincipalId   *string `json:"principalId,omitempty"`
	ResourceId    *string `json:"resourceId,omitempty"`
}
