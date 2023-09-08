package groupevent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupEventClient struct {
	Client *msgraph.Client
}

func NewGroupEventClientWithBaseURI(api sdkEnv.Api) (*GroupEventClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupevent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupEventClient: %+v", err)
	}

	return &GroupEventClient{
		Client: client,
	}, nil
}
