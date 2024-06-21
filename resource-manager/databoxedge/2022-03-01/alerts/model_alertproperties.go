package alerts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertProperties struct {
	AlertType           *string            `json:"alertType,omitempty"`
	AppearedAtDateTime  *string            `json:"appearedAtDateTime,omitempty"`
	DetailedInformation *map[string]string `json:"detailedInformation,omitempty"`
	ErrorDetails        *AlertErrorDetails `json:"errorDetails,omitempty"`
	Recommendation      *string            `json:"recommendation,omitempty"`
	Severity            *AlertSeverity     `json:"severity,omitempty"`
	Title               *string            `json:"title,omitempty"`
}
