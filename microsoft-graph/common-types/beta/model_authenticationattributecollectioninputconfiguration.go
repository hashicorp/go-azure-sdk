package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationAttributeCollectionInputConfiguration struct {
	Attribute        *string                                                       `json:"attribute,omitempty"`
	DefaultValue     *string                                                       `json:"defaultValue,omitempty"`
	Editable         *bool                                                         `json:"editable,omitempty"`
	Hidden           *bool                                                         `json:"hidden,omitempty"`
	InputType        *AuthenticationAttributeCollectionInputConfigurationInputType `json:"inputType,omitempty"`
	Label            *string                                                       `json:"label,omitempty"`
	ODataType        *string                                                       `json:"@odata.type,omitempty"`
	Options          *[]AuthenticationAttributeCollectionOptionConfiguration       `json:"options,omitempty"`
	Required         *bool                                                         `json:"required,omitempty"`
	ValidationRegEx  *string                                                       `json:"validationRegEx,omitempty"`
	WriteToDirectory *bool                                                         `json:"writeToDirectory,omitempty"`
}
