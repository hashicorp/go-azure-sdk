package alerts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertPropertiesDefinition struct {
	Category *AlertCategory `json:"category,omitempty"`
	Criteria *AlertCriteria `json:"criteria,omitempty"`
	Type     *AlertType     `json:"type,omitempty"`
}
