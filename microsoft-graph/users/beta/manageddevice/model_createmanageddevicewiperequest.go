package manageddevice

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateManagedDeviceWipeRequest struct {
	KeepEnrollmentData   *bool                                               `json:"keepEnrollmentData,omitempty"`
	KeepUserData         *bool                                               `json:"keepUserData,omitempty"`
	MacOsUnlockCode      *string                                             `json:"macOsUnlockCode,omitempty"`
	ObliterationBehavior *CreateManagedDeviceWipeRequestObliterationBehavior `json:"obliterationBehavior,omitempty"`
	PersistEsimDataPlan  *bool                                               `json:"persistEsimDataPlan,omitempty"`
	UseProtectedWipe     *bool                                               `json:"useProtectedWipe,omitempty"`
}
