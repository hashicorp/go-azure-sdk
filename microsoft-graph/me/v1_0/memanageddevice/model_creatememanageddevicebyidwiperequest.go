package memanageddevice

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMeManagedDeviceByIdWipeRequest struct {
	KeepEnrollmentData  *bool   `json:"keepEnrollmentData,omitempty"`
	KeepUserData        *bool   `json:"keepUserData,omitempty"`
	MacOsUnlockCode     *string `json:"macOsUnlockCode,omitempty"`
	PersistEsimDataPlan *bool   `json:"persistEsimDataPlan,omitempty"`
}
