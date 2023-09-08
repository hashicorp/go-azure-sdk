package meactivityhistoryitemactivity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeActivityHistoryItemActivityClient struct {
	Client *msgraph.Client
}

func NewMeActivityHistoryItemActivityClientWithBaseURI(api sdkEnv.Api) (*MeActivityHistoryItemActivityClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meactivityhistoryitemactivity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeActivityHistoryItemActivityClient: %+v", err)
	}

	return &MeActivityHistoryItemActivityClient{
		Client: client,
	}, nil
}
