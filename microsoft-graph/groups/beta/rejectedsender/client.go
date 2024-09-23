package rejectedsender

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RejectedSenderClient struct {
	Client *msgraph.Client
}

func NewRejectedSenderClientWithBaseURI(sdkApi sdkEnv.Api) (*RejectedSenderClient, error) {
	client, err := msgraph.NewClient(sdkApi, "rejectedsender", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RejectedSenderClient: %+v", err)
	}

	return &RejectedSenderClient{
		Client: client,
	}, nil
}
