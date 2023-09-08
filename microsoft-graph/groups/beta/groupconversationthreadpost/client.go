package groupconversationthreadpost

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupConversationThreadPostClient struct {
	Client *msgraph.Client
}

func NewGroupConversationThreadPostClientWithBaseURI(api sdkEnv.Api) (*GroupConversationThreadPostClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupconversationthreadpost", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupConversationThreadPostClient: %+v", err)
	}

	return &GroupConversationThreadPostClient{
		Client: client,
	}, nil
}
