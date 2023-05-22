package sharedgalleries

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PirSharedGalleryResourceOperationPredicate struct {
	Location *string
	Name     *string
}

func (p PirSharedGalleryResourceOperationPredicate) Matches(input PirSharedGalleryResource) bool {

	if p.Location != nil && (input.Location == nil && *p.Location != *input.Location) {
		return false
	}

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	return true
}
