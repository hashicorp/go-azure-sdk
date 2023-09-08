package useractivityhistoryitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserActivityHistoryItemClient struct {
	Client *msgraph.Client
}

func NewUserActivityHistoryItemClientWithBaseURI(api sdkEnv.Api) (*UserActivityHistoryItemClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useractivityhistoryitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserActivityHistoryItemClient: %+v", err)
	}

	return &UserActivityHistoryItemClient{
		Client: client,
	}, nil
}
