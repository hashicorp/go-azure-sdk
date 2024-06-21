package domainservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HealthAlert struct {
	Id            *string `json:"id,omitempty"`
	Issue         *string `json:"issue,omitempty"`
	LastDetected  *string `json:"lastDetected,omitempty"`
	Name          *string `json:"name,omitempty"`
	Raised        *string `json:"raised,omitempty"`
	ResolutionUri *string `json:"resolutionUri,omitempty"`
	Severity      *string `json:"severity,omitempty"`
}
