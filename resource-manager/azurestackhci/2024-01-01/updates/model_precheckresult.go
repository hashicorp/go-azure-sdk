package updates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrecheckResult struct {
	AdditionalData     *string             `json:"additionalData,omitempty"`
	Description        *string             `json:"description,omitempty"`
	DisplayName        *string             `json:"displayName,omitempty"`
	HealthCheckSource  *string             `json:"healthCheckSource,omitempty"`
	Name               *string             `json:"name,omitempty"`
	Remediation        *string             `json:"remediation,omitempty"`
	Severity           *Severity           `json:"severity,omitempty"`
	Status             *Status             `json:"status,omitempty"`
	Tags               *PrecheckResultTags `json:"tags,omitempty"`
	TargetResourceID   *string             `json:"targetResourceID,omitempty"`
	TargetResourceName *string             `json:"targetResourceName,omitempty"`
	Timestamp          *string             `json:"timestamp,omitempty"`
	Title              *string             `json:"title,omitempty"`
}
