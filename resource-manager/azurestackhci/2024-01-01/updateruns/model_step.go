package updateruns

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Step struct {
	Description        *string `json:"description,omitempty"`
	EndTimeUtc         *string `json:"endTimeUtc,omitempty"`
	ErrorMessage       *string `json:"errorMessage,omitempty"`
	LastUpdatedTimeUtc *string `json:"lastUpdatedTimeUtc,omitempty"`
	Name               *string `json:"name,omitempty"`
	StartTimeUtc       *string `json:"startTimeUtc,omitempty"`
	Status             *string `json:"status,omitempty"`
	Steps              *[]Step `json:"steps,omitempty"`
}
