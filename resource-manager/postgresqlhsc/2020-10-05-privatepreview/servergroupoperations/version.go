package servergroupoperations

import "fmt"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

const defaultApiVersion = "2020-10-05-privatepreview"

func userAgent() string {
	return fmt.Sprintf("hashicorp/go-azure-sdk/servergroupoperations/%s", defaultApiVersion)
}
