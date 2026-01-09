package dailyprintusage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DailyPrintUsageClient struct {
	Client *msgraph.Client
}

func NewDailyPrintUsageClientWithBaseURI(sdkApi sdkEnv.Api) (*DailyPrintUsageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "dailyprintusage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DailyPrintUsageClient: %+v", err)
	}

	return &DailyPrintUsageClient{
		Client: client,
	}, nil
}
