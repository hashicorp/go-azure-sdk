// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sharedgalleries

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
