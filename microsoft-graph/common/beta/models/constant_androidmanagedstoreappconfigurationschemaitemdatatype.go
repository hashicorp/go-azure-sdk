package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAppConfigurationSchemaItemDataType string

const (
	AndroidManagedStoreAppConfigurationSchemaItemDataTypebool        AndroidManagedStoreAppConfigurationSchemaItemDataType = "Bool"
	AndroidManagedStoreAppConfigurationSchemaItemDataTypebundle      AndroidManagedStoreAppConfigurationSchemaItemDataType = "Bundle"
	AndroidManagedStoreAppConfigurationSchemaItemDataTypebundleArray AndroidManagedStoreAppConfigurationSchemaItemDataType = "BundleArray"
	AndroidManagedStoreAppConfigurationSchemaItemDataTypechoice      AndroidManagedStoreAppConfigurationSchemaItemDataType = "Choice"
	AndroidManagedStoreAppConfigurationSchemaItemDataTypehidden      AndroidManagedStoreAppConfigurationSchemaItemDataType = "Hidden"
	AndroidManagedStoreAppConfigurationSchemaItemDataTypeinteger     AndroidManagedStoreAppConfigurationSchemaItemDataType = "Integer"
	AndroidManagedStoreAppConfigurationSchemaItemDataTypemultiselect AndroidManagedStoreAppConfigurationSchemaItemDataType = "Multiselect"
	AndroidManagedStoreAppConfigurationSchemaItemDataTypestring      AndroidManagedStoreAppConfigurationSchemaItemDataType = "String"
)

func PossibleValuesForAndroidManagedStoreAppConfigurationSchemaItemDataType() []string {
	return []string{
		string(AndroidManagedStoreAppConfigurationSchemaItemDataTypebool),
		string(AndroidManagedStoreAppConfigurationSchemaItemDataTypebundle),
		string(AndroidManagedStoreAppConfigurationSchemaItemDataTypebundleArray),
		string(AndroidManagedStoreAppConfigurationSchemaItemDataTypechoice),
		string(AndroidManagedStoreAppConfigurationSchemaItemDataTypehidden),
		string(AndroidManagedStoreAppConfigurationSchemaItemDataTypeinteger),
		string(AndroidManagedStoreAppConfigurationSchemaItemDataTypemultiselect),
		string(AndroidManagedStoreAppConfigurationSchemaItemDataTypestring),
	}
}

func parseAndroidManagedStoreAppConfigurationSchemaItemDataType(input string) (*AndroidManagedStoreAppConfigurationSchemaItemDataType, error) {
	vals := map[string]AndroidManagedStoreAppConfigurationSchemaItemDataType{
		"bool":        AndroidManagedStoreAppConfigurationSchemaItemDataTypebool,
		"bundle":      AndroidManagedStoreAppConfigurationSchemaItemDataTypebundle,
		"bundlearray": AndroidManagedStoreAppConfigurationSchemaItemDataTypebundleArray,
		"choice":      AndroidManagedStoreAppConfigurationSchemaItemDataTypechoice,
		"hidden":      AndroidManagedStoreAppConfigurationSchemaItemDataTypehidden,
		"integer":     AndroidManagedStoreAppConfigurationSchemaItemDataTypeinteger,
		"multiselect": AndroidManagedStoreAppConfigurationSchemaItemDataTypemultiselect,
		"string":      AndroidManagedStoreAppConfigurationSchemaItemDataTypestring,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedStoreAppConfigurationSchemaItemDataType(input)
	return &out, nil
}
