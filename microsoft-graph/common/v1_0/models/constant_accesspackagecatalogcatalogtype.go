package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageCatalogCatalogType string

const (
	AccessPackageCatalogCatalogTypeserviceDefault AccessPackageCatalogCatalogType = "ServiceDefault"
	AccessPackageCatalogCatalogTypeserviceManaged AccessPackageCatalogCatalogType = "ServiceManaged"
	AccessPackageCatalogCatalogTypeuserManaged    AccessPackageCatalogCatalogType = "UserManaged"
)

func PossibleValuesForAccessPackageCatalogCatalogType() []string {
	return []string{
		string(AccessPackageCatalogCatalogTypeserviceDefault),
		string(AccessPackageCatalogCatalogTypeserviceManaged),
		string(AccessPackageCatalogCatalogTypeuserManaged),
	}
}

func parseAccessPackageCatalogCatalogType(input string) (*AccessPackageCatalogCatalogType, error) {
	vals := map[string]AccessPackageCatalogCatalogType{
		"servicedefault": AccessPackageCatalogCatalogTypeserviceDefault,
		"servicemanaged": AccessPackageCatalogCatalogTypeserviceManaged,
		"usermanaged":    AccessPackageCatalogCatalogTypeuserManaged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageCatalogCatalogType(input)
	return &out, nil
}
