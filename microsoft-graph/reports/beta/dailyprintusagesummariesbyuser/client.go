package dailyprintusagesummariesbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DailyPrintUsageSummariesByUserClient struct {
	Client *msgraph.Client
}

func NewDailyPrintUsageSummariesByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*DailyPrintUsageSummariesByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "dailyprintusagesummariesbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DailyPrintUsageSummariesByUserClient: %+v", err)
	}

	return &DailyPrintUsageSummariesByUserClient{
		Client: client,
	}, nil
}
