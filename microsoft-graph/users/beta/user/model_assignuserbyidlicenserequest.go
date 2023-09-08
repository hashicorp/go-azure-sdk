package user

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignUserByIdLicenseRequest struct {
	AddLicenses    *[]AssignedLicense `json:"addLicenses,omitempty"`
	RemoveLicenses *[]string          `json:"removeLicenses,omitempty"`
}
