package servicetasks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckOCIDriverTaskOutput struct {
	InstalledDriver  *OracleOCIDriverInfo   `json:"installedDriver,omitempty"`
	ValidationErrors *[]ReportableException `json:"validationErrors,omitempty"`
}
