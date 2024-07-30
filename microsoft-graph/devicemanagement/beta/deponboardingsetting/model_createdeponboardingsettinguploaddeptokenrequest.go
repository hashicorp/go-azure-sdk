package deponboardingsetting

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateDepOnboardingSettingUploadDepTokenRequest struct {
	AppleId  *string `json:"appleId,omitempty"`
	DepToken *string `json:"depToken,omitempty"`
}
