package applicationwhitelistings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectionMode struct {
	Exe        *Exe        `json:"exe,omitempty"`
	Executable *Executable `json:"executable,omitempty"`
	Msi        *Msi        `json:"msi,omitempty"`
	Script     *Script     `json:"script,omitempty"`
}
