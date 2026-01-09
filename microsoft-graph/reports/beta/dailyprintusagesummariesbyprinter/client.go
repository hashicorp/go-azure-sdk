package dailyprintusagesummariesbyprinter

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DailyPrintUsageSummariesByPrinterClient struct {
	Client *msgraph.Client
}

func NewDailyPrintUsageSummariesByPrinterClientWithBaseURI(sdkApi sdkEnv.Api) (*DailyPrintUsageSummariesByPrinterClient, error) {
	client, err := msgraph.NewClient(sdkApi, "dailyprintusagesummariesbyprinter", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DailyPrintUsageSummariesByPrinterClient: %+v", err)
	}

	return &DailyPrintUsageSummariesByPrinterClient{
		Client: client,
	}, nil
}
