package notebooks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotebookMetadata struct {
	Kernelspec   *NotebookKernelSpec   `json:"kernelspec,omitempty"`
	LanguageInfo *NotebookLanguageInfo `json:"language_info,omitempty"`
}
