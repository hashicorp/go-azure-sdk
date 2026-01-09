package activityhistoryitemactivity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActivityHistoryItemActivityClient struct {
	Client *msgraph.Client
}

func NewActivityHistoryItemActivityClientWithBaseURI(sdkApi sdkEnv.Api) (*ActivityHistoryItemActivityClient, error) {
	client, err := msgraph.NewClient(sdkApi, "activityhistoryitemactivity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ActivityHistoryItemActivityClient: %+v", err)
	}

	return &ActivityHistoryItemActivityClient{
		Client: client,
	}, nil
}
