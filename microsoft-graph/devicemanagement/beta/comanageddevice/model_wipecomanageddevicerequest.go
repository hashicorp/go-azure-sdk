package comanageddevice

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WipeComanagedDeviceRequest struct {
	KeepEnrollmentData   nullable.Type[bool]        `json:"keepEnrollmentData,omitempty"`
	KeepUserData         nullable.Type[bool]        `json:"keepUserData,omitempty"`
	MacOsUnlockCode      nullable.Type[string]      `json:"macOsUnlockCode,omitempty"`
	ObliterationBehavior *beta.ObliterationBehavior `json:"obliterationBehavior,omitempty"`
	PersistEsimDataPlan  nullable.Type[bool]        `json:"persistEsimDataPlan,omitempty"`
	UseProtectedWipe     nullable.Type[bool]        `json:"useProtectedWipe,omitempty"`
}
