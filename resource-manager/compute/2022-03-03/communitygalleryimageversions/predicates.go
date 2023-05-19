package communitygalleryimageversions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommunityGalleryImageVersionOperationPredicate struct {
	Location *string
	Name     *string
	Type     *string
}

func (p CommunityGalleryImageVersionOperationPredicate) Matches(input CommunityGalleryImageVersion) bool {

	if p.Location != nil && (input.Location == nil && *p.Location != *input.Location) {
		return false
	}

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil && *p.Type != *input.Type) {
		return false
	}

	return true
}
