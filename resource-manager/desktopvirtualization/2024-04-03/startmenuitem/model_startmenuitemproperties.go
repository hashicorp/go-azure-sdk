package startmenuitem

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StartMenuItemProperties struct {
	AppAlias             *string `json:"appAlias,omitempty"`
	CommandLineArguments *string `json:"commandLineArguments,omitempty"`
	FilePath             *string `json:"filePath,omitempty"`
	IconIndex            *int64  `json:"iconIndex,omitempty"`
	IconPath             *string `json:"iconPath,omitempty"`
}
