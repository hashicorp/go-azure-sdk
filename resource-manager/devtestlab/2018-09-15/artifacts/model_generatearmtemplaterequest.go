package artifacts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GenerateArmTemplateRequest struct {
	FileUploadOptions  *FileUploadOptions `json:"fileUploadOptions,omitempty"`
	Location           *string            `json:"location,omitempty"`
	Parameters         *[]ParameterInfo   `json:"parameters,omitempty"`
	VirtualMachineName *string            `json:"virtualMachineName,omitempty"`
}
