package labelingjob

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LabelClass struct {
	DisplayName *string                `json:"displayName,omitempty"`
	Subclasses  *map[string]LabelClass `json:"subclasses,omitempty"`
}
