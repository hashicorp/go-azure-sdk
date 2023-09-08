package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftManagedDesktopType string

const (
	MicrosoftManagedDesktopTypenotManaged      MicrosoftManagedDesktopType = "NotManaged"
	MicrosoftManagedDesktopTypepremiumManaged  MicrosoftManagedDesktopType = "PremiumManaged"
	MicrosoftManagedDesktopTypestandardManaged MicrosoftManagedDesktopType = "StandardManaged"
	MicrosoftManagedDesktopTypestarterManaged  MicrosoftManagedDesktopType = "StarterManaged"
)

func PossibleValuesForMicrosoftManagedDesktopType() []string {
	return []string{
		string(MicrosoftManagedDesktopTypenotManaged),
		string(MicrosoftManagedDesktopTypepremiumManaged),
		string(MicrosoftManagedDesktopTypestandardManaged),
		string(MicrosoftManagedDesktopTypestarterManaged),
	}
}

func parseMicrosoftManagedDesktopType(input string) (*MicrosoftManagedDesktopType, error) {
	vals := map[string]MicrosoftManagedDesktopType{
		"notmanaged":      MicrosoftManagedDesktopTypenotManaged,
		"premiummanaged":  MicrosoftManagedDesktopTypepremiumManaged,
		"standardmanaged": MicrosoftManagedDesktopTypestandardManaged,
		"startermanaged":  MicrosoftManagedDesktopTypestarterManaged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftManagedDesktopType(input)
	return &out, nil
}
