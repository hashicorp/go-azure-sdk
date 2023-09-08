package grouprejectedsender

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupRejectedSenderClient struct {
	Client *msgraph.Client
}

func NewGroupRejectedSenderClientWithBaseURI(api sdkEnv.Api) (*GroupRejectedSenderClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouprejectedsender", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupRejectedSenderClient: %+v", err)
	}

	return &GroupRejectedSenderClient{
		Client: client,
	}, nil
}
