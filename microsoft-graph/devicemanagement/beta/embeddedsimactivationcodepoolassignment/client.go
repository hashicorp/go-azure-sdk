package embeddedsimactivationcodepoolassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmbeddedSIMActivationCodePoolAssignmentClient struct {
	Client *msgraph.Client
}

func NewEmbeddedSIMActivationCodePoolAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EmbeddedSIMActivationCodePoolAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "embeddedsimactivationcodepoolassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EmbeddedSIMActivationCodePoolAssignmentClient: %+v", err)
	}

	return &EmbeddedSIMActivationCodePoolAssignmentClient{
		Client: client,
	}, nil
}
