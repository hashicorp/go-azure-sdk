package policy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyClient struct {
	Client *msgraph.Client
}

func NewPolicyClientWithBaseURI(api sdkEnv.Api) (*PolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyClient: %+v", err)
	}

	return &PolicyClient{
		Client: client,
	}, nil
}
