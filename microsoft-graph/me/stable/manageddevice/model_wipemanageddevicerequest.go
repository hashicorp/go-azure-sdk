package manageddevice

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WipeManagedDeviceRequest struct {
	KeepEnrollmentData  nullable.Type[bool]   `json:"keepEnrollmentData,omitempty"`
	KeepUserData        nullable.Type[bool]   `json:"keepUserData,omitempty"`
	MacOsUnlockCode     nullable.Type[string] `json:"macOsUnlockCode,omitempty"`
	PersistEsimDataPlan nullable.Type[bool]   `json:"persistEsimDataPlan,omitempty"`
}
