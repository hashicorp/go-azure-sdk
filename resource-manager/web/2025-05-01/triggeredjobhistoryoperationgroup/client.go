package triggeredjobhistoryoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TriggeredJobHistoryOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewTriggeredJobHistoryOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*TriggeredJobHistoryOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "triggeredjobhistoryoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TriggeredJobHistoryOperationGroupClient: %+v", err)
	}

	return &TriggeredJobHistoryOperationGroupClient{
		Client: client,
	}, nil
}
