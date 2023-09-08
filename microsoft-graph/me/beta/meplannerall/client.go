package meplannerall

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MePlannerAllClient struct {
	Client *msgraph.Client
}

func NewMePlannerAllClientWithBaseURI(api sdkEnv.Api) (*MePlannerAllClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meplannerall", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MePlannerAllClient: %+v", err)
	}

	return &MePlannerAllClient{
		Client: client,
	}, nil
}
