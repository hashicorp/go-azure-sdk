package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesDurationDrivenRolloutSettings struct {
	DurationBetweenOffers      *string `json:"durationBetweenOffers,omitempty"`
	DurationUntilDeploymentEnd *string `json:"durationUntilDeploymentEnd,omitempty"`
	ODataType                  *string `json:"@odata.type,omitempty"`
}
