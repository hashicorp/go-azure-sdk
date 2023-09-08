package group

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignGroupByIdLicenseRequest struct {
	AddLicenses    *[]AssignedLicense `json:"addLicenses,omitempty"`
	RemoveLicenses *[]string          `json:"removeLicenses,omitempty"`
}
