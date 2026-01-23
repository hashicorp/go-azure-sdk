package virtualmachinescalesetvmruncommands

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineRunCommandScriptSource struct {
	CommandId                *string                    `json:"commandId,omitempty"`
	GalleryScriptReferenceId *string                    `json:"galleryScriptReferenceId,omitempty"`
	Script                   *string                    `json:"script,omitempty"`
	ScriptShell              *ScriptShellTypes          `json:"scriptShell,omitempty"`
	ScriptUri                *string                    `json:"scriptUri,omitempty"`
	ScriptUriManagedIdentity *RunCommandManagedIdentity `json:"scriptUriManagedIdentity,omitempty"`
}
