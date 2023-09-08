package memobileappintentandstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeMobileAppIntentAndStateClient struct {
	Client *msgraph.Client
}

func NewMeMobileAppIntentAndStateClientWithBaseURI(api sdkEnv.Api) (*MeMobileAppIntentAndStateClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "memobileappintentandstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeMobileAppIntentAndStateClient: %+v", err)
	}

	return &MeMobileAppIntentAndStateClient{
		Client: client,
	}, nil
}
