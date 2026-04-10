package webtestlocationsapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationInsightsComponentWebTestLocationOperationPredicate struct {
	DisplayName *string
	Tag         *string
}

func (p ApplicationInsightsComponentWebTestLocationOperationPredicate) Matches(input ApplicationInsightsComponentWebTestLocation) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Tag != nil && (input.Tag == nil || *p.Tag != *input.Tag) {
		return false
	}

	return true
}
