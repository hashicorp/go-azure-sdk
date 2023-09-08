package groupeventinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupEventInstanceClient struct {
	Client *msgraph.Client
}

func NewGroupEventInstanceClientWithBaseURI(api sdkEnv.Api) (*GroupEventInstanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupeventinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupEventInstanceClient: %+v", err)
	}

	return &GroupEventInstanceClient{
		Client: client,
	}, nil
}
