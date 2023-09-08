package meactivityhistoryitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeActivityHistoryItemClient struct {
	Client *msgraph.Client
}

func NewMeActivityHistoryItemClientWithBaseURI(api sdkEnv.Api) (*MeActivityHistoryItemClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meactivityhistoryitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeActivityHistoryItemClient: %+v", err)
	}

	return &MeActivityHistoryItemClient{
		Client: client,
	}, nil
}
