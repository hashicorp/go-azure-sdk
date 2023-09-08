package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkAppConfigurationSchemaItemDataType string

const (
	AndroidForWorkAppConfigurationSchemaItemDataTypebool        AndroidForWorkAppConfigurationSchemaItemDataType = "Bool"
	AndroidForWorkAppConfigurationSchemaItemDataTypebundle      AndroidForWorkAppConfigurationSchemaItemDataType = "Bundle"
	AndroidForWorkAppConfigurationSchemaItemDataTypebundleArray AndroidForWorkAppConfigurationSchemaItemDataType = "BundleArray"
	AndroidForWorkAppConfigurationSchemaItemDataTypechoice      AndroidForWorkAppConfigurationSchemaItemDataType = "Choice"
	AndroidForWorkAppConfigurationSchemaItemDataTypehidden      AndroidForWorkAppConfigurationSchemaItemDataType = "Hidden"
	AndroidForWorkAppConfigurationSchemaItemDataTypeinteger     AndroidForWorkAppConfigurationSchemaItemDataType = "Integer"
	AndroidForWorkAppConfigurationSchemaItemDataTypemultiselect AndroidForWorkAppConfigurationSchemaItemDataType = "Multiselect"
	AndroidForWorkAppConfigurationSchemaItemDataTypestring      AndroidForWorkAppConfigurationSchemaItemDataType = "String"
)

func PossibleValuesForAndroidForWorkAppConfigurationSchemaItemDataType() []string {
	return []string{
		string(AndroidForWorkAppConfigurationSchemaItemDataTypebool),
		string(AndroidForWorkAppConfigurationSchemaItemDataTypebundle),
		string(AndroidForWorkAppConfigurationSchemaItemDataTypebundleArray),
		string(AndroidForWorkAppConfigurationSchemaItemDataTypechoice),
		string(AndroidForWorkAppConfigurationSchemaItemDataTypehidden),
		string(AndroidForWorkAppConfigurationSchemaItemDataTypeinteger),
		string(AndroidForWorkAppConfigurationSchemaItemDataTypemultiselect),
		string(AndroidForWorkAppConfigurationSchemaItemDataTypestring),
	}
}

func parseAndroidForWorkAppConfigurationSchemaItemDataType(input string) (*AndroidForWorkAppConfigurationSchemaItemDataType, error) {
	vals := map[string]AndroidForWorkAppConfigurationSchemaItemDataType{
		"bool":        AndroidForWorkAppConfigurationSchemaItemDataTypebool,
		"bundle":      AndroidForWorkAppConfigurationSchemaItemDataTypebundle,
		"bundlearray": AndroidForWorkAppConfigurationSchemaItemDataTypebundleArray,
		"choice":      AndroidForWorkAppConfigurationSchemaItemDataTypechoice,
		"hidden":      AndroidForWorkAppConfigurationSchemaItemDataTypehidden,
		"integer":     AndroidForWorkAppConfigurationSchemaItemDataTypeinteger,
		"multiselect": AndroidForWorkAppConfigurationSchemaItemDataTypemultiselect,
		"string":      AndroidForWorkAppConfigurationSchemaItemDataTypestring,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkAppConfigurationSchemaItemDataType(input)
	return &out, nil
}
