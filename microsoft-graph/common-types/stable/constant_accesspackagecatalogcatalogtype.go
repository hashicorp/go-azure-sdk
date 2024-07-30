package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageCatalogCatalogType string

const (
	AccessPackageCatalogCatalogType_ServiceDefault AccessPackageCatalogCatalogType = "serviceDefault"
	AccessPackageCatalogCatalogType_ServiceManaged AccessPackageCatalogCatalogType = "serviceManaged"
	AccessPackageCatalogCatalogType_UserManaged    AccessPackageCatalogCatalogType = "userManaged"
)

func PossibleValuesForAccessPackageCatalogCatalogType() []string {
	return []string{
		string(AccessPackageCatalogCatalogType_ServiceDefault),
		string(AccessPackageCatalogCatalogType_ServiceManaged),
		string(AccessPackageCatalogCatalogType_UserManaged),
	}
}

func (s *AccessPackageCatalogCatalogType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessPackageCatalogCatalogType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessPackageCatalogCatalogType(input string) (*AccessPackageCatalogCatalogType, error) {
	vals := map[string]AccessPackageCatalogCatalogType{
		"servicedefault": AccessPackageCatalogCatalogType_ServiceDefault,
		"servicemanaged": AccessPackageCatalogCatalogType_ServiceManaged,
		"usermanaged":    AccessPackageCatalogCatalogType_UserManaged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageCatalogCatalogType(input)
	return &out, nil
}
