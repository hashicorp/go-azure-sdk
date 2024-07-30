package calendargroupcalendareventexceptionoccurrenceinstanceextension

import "fmt"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

const defaultApiVersion = "beta"

func userAgent() string {
	return fmt.Sprintf("hashicorp/go-azure-sdk/calendargroupcalendareventexceptionoccurrenceinstanceextension/%s", defaultApiVersion)
}
