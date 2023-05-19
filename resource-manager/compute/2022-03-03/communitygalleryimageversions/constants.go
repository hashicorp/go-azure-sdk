package communitygalleryimageversions

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedGalleryHostCaching string

const (
	SharedGalleryHostCachingNone      SharedGalleryHostCaching = "None"
	SharedGalleryHostCachingReadOnly  SharedGalleryHostCaching = "ReadOnly"
	SharedGalleryHostCachingReadWrite SharedGalleryHostCaching = "ReadWrite"
)

func PossibleValuesForSharedGalleryHostCaching() []string {
	return []string{
		string(SharedGalleryHostCachingNone),
		string(SharedGalleryHostCachingReadOnly),
		string(SharedGalleryHostCachingReadWrite),
	}
}

func parseSharedGalleryHostCaching(input string) (*SharedGalleryHostCaching, error) {
	vals := map[string]SharedGalleryHostCaching{
		"none":      SharedGalleryHostCachingNone,
		"readonly":  SharedGalleryHostCachingReadOnly,
		"readwrite": SharedGalleryHostCachingReadWrite,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharedGalleryHostCaching(input)
	return &out, nil
}
