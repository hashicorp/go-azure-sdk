package computenodes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMExtensionInstanceView struct {
	Name        *string               `json:"name,omitempty"`
	Statuses    *[]InstanceViewStatus `json:"statuses,omitempty"`
	SubStatuses *[]InstanceViewStatus `json:"subStatuses,omitempty"`
}
