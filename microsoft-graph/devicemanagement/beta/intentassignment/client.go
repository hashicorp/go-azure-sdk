package intentassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntentAssignmentClient struct {
	Client *msgraph.Client
}

func NewIntentAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*IntentAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "intentassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IntentAssignmentClient: %+v", err)
	}

	return &IntentAssignmentClient{
		Client: client,
	}, nil
}
