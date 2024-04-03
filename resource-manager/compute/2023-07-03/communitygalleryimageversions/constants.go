package communitygalleryimageversions

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

func (s *SharedGalleryHostCaching) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSharedGalleryHostCaching(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
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
