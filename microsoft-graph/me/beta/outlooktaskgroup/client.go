package outlooktaskgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutlookTaskGroupClient struct {
	Client *msgraph.Client
}

func NewOutlookTaskGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*OutlookTaskGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "outlooktaskgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OutlookTaskGroupClient: %+v", err)
	}

	return &OutlookTaskGroupClient{
		Client: client,
	}, nil
}
