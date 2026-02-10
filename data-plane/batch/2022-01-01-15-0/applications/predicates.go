package applications

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationSummaryOperationPredicate struct {
	DisplayName *string
	Id          *string
}

func (p ApplicationSummaryOperationPredicate) Matches(input ApplicationSummary) bool {

	if p.DisplayName != nil && *p.DisplayName != input.DisplayName {
		return false
	}

	if p.Id != nil && *p.Id != input.Id {
		return false
	}

	return true
}
