package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPrint struct {
	ODataType           *string         `json:"@odata.type,omitempty"`
	RecentPrinterShares *[]PrinterShare `json:"recentPrinterShares,omitempty"`
}
