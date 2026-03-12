package publiccertificates

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PublicCertificateLocation string

const (
	PublicCertificateLocationCurrentUserMy  PublicCertificateLocation = "CurrentUserMy"
	PublicCertificateLocationLocalMachineMy PublicCertificateLocation = "LocalMachineMy"
	PublicCertificateLocationUnknown        PublicCertificateLocation = "Unknown"
)

func PossibleValuesForPublicCertificateLocation() []string {
	return []string{
		string(PublicCertificateLocationCurrentUserMy),
		string(PublicCertificateLocationLocalMachineMy),
		string(PublicCertificateLocationUnknown),
	}
}

func (s *PublicCertificateLocation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePublicCertificateLocation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePublicCertificateLocation(input string) (*PublicCertificateLocation, error) {
	vals := map[string]PublicCertificateLocation{
		"currentusermy":  PublicCertificateLocationCurrentUserMy,
		"localmachinemy": PublicCertificateLocationLocalMachineMy,
		"unknown":        PublicCertificateLocationUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PublicCertificateLocation(input)
	return &out, nil
}
