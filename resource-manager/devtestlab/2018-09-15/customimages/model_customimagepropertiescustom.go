package customimages

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomImagePropertiesCustom struct {
	ImageName *string           `json:"imageName,omitempty"`
	OsType    CustomImageOsType `json:"osType"`
	SysPrep   *bool             `json:"sysPrep,omitempty"`
}
