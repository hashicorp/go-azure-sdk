package schedule

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LabelCategory struct {
	Classes     *map[string]LabelClass `json:"classes,omitempty"`
	DisplayName *string                `json:"displayName,omitempty"`
	MultiSelect *MultiSelect           `json:"multiSelect,omitempty"`
}
