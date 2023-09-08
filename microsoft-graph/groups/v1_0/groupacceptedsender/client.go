package groupacceptedsender

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupAcceptedSenderClient struct {
	Client *msgraph.Client
}

func NewGroupAcceptedSenderClientWithBaseURI(api sdkEnv.Api) (*GroupAcceptedSenderClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupacceptedsender", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupAcceptedSenderClient: %+v", err)
	}

	return &GroupAcceptedSenderClient{
		Client: client,
	}, nil
}
