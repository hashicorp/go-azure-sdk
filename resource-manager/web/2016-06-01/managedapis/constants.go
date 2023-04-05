package managedapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiType string

const (
	ApiTypeNotSpecified ApiType = "NotSpecified"
	ApiTypeRest         ApiType = "Rest"
	ApiTypeSoap         ApiType = "Soap"
)

func PossibleValuesForApiType() []string {
	return []string{
		string(ApiTypeNotSpecified),
		string(ApiTypeRest),
		string(ApiTypeSoap),
	}
}

type ConnectionParameterType string

const (
	ConnectionParameterTypeArray        ConnectionParameterType = "array"
	ConnectionParameterTypeBool         ConnectionParameterType = "bool"
	ConnectionParameterTypeConnection   ConnectionParameterType = "connection"
	ConnectionParameterTypeInt          ConnectionParameterType = "int"
	ConnectionParameterTypeOauthSetting ConnectionParameterType = "oauthSetting"
	ConnectionParameterTypeObject       ConnectionParameterType = "object"
	ConnectionParameterTypeSecureobject ConnectionParameterType = "secureobject"
	ConnectionParameterTypeSecurestring ConnectionParameterType = "securestring"
	ConnectionParameterTypeString       ConnectionParameterType = "string"
)

func PossibleValuesForConnectionParameterType() []string {
	return []string{
		string(ConnectionParameterTypeArray),
		string(ConnectionParameterTypeBool),
		string(ConnectionParameterTypeConnection),
		string(ConnectionParameterTypeInt),
		string(ConnectionParameterTypeOauthSetting),
		string(ConnectionParameterTypeObject),
		string(ConnectionParameterTypeSecureobject),
		string(ConnectionParameterTypeSecurestring),
		string(ConnectionParameterTypeString),
	}
}

type WsdlImportMethod string

const (
	WsdlImportMethodNotSpecified    WsdlImportMethod = "NotSpecified"
	WsdlImportMethodSoapPassThrough WsdlImportMethod = "SoapPassThrough"
	WsdlImportMethodSoapToRest      WsdlImportMethod = "SoapToRest"
)

func PossibleValuesForWsdlImportMethod() []string {
	return []string{
		string(WsdlImportMethodNotSpecified),
		string(WsdlImportMethodSoapPassThrough),
		string(WsdlImportMethodSoapToRest),
	}
}
