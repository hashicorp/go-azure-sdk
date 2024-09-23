package outlooktask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutlookTaskClient struct {
	Client *msgraph.Client
}

func NewOutlookTaskClientWithBaseURI(sdkApi sdkEnv.Api) (*OutlookTaskClient, error) {
	client, err := msgraph.NewClient(sdkApi, "outlooktask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OutlookTaskClient: %+v", err)
	}

	return &OutlookTaskClient{
		Client: client,
	}, nil
}
