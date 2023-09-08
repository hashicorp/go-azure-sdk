package useractivityhistoryitemactivity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserActivityHistoryItemActivityClient struct {
	Client *msgraph.Client
}

func NewUserActivityHistoryItemActivityClientWithBaseURI(api sdkEnv.Api) (*UserActivityHistoryItemActivityClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useractivityhistoryitemactivity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserActivityHistoryItemActivityClient: %+v", err)
	}

	return &UserActivityHistoryItemActivityClient{
		Client: client,
	}, nil
}
