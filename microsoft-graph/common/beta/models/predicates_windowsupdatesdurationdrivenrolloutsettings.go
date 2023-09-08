package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesDurationDrivenRolloutSettingsOperationPredicate struct {
	DurationBetweenOffers      *string
	DurationUntilDeploymentEnd *string
	ODataType                  *string
}

func (p WindowsUpdatesDurationDrivenRolloutSettingsOperationPredicate) Matches(input WindowsUpdatesDurationDrivenRolloutSettings) bool {

	if p.DurationBetweenOffers != nil && (input.DurationBetweenOffers == nil || *p.DurationBetweenOffers != *input.DurationBetweenOffers) {
		return false
	}

	if p.DurationUntilDeploymentEnd != nil && (input.DurationUntilDeploymentEnd == nil || *p.DurationUntilDeploymentEnd != *input.DurationUntilDeploymentEnd) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
