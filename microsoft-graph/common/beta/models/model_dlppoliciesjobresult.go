package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DlpPoliciesJobResult struct {
	AuditCorrelationId *string            `json:"auditCorrelationId,omitempty"`
	EvaluationDateTime *string            `json:"evaluationDateTime,omitempty"`
	MatchingRules      *[]MatchingDlpRule `json:"matchingRules,omitempty"`
	ODataType          *string            `json:"@odata.type,omitempty"`
}
