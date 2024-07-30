package oauth2permissiongrant

import "fmt"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

const defaultApiVersion = "stable"

func userAgent() string {
	return fmt.Sprintf("hashicorp/go-azure-sdk/oauth2permissiongrant/%s", defaultApiVersion)
}
