package intentuserstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntentUserStateClient struct {
	Client *msgraph.Client
}

func NewIntentUserStateClientWithBaseURI(sdkApi sdkEnv.Api) (*IntentUserStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "intentuserstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IntentUserStateClient: %+v", err)
	}

	return &IntentUserStateClient{
		Client: client,
	}, nil
}
