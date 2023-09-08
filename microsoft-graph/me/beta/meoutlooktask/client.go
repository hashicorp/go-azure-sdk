package meoutlooktask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOutlookTaskClient struct {
	Client *msgraph.Client
}

func NewMeOutlookTaskClientWithBaseURI(api sdkEnv.Api) (*MeOutlookTaskClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meoutlooktask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOutlookTaskClient: %+v", err)
	}

	return &MeOutlookTaskClient{
		Client: client,
	}, nil
}
