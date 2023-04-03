package alertruleincidents

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncidentListResult struct {
	Value *[]Incident `json:"value,omitempty"`
}
