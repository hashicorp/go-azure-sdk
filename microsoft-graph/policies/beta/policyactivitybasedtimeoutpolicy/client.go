package policyactivitybasedtimeoutpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyActivityBasedTimeoutPolicyClient struct {
	Client *msgraph.Client
}

func NewPolicyActivityBasedTimeoutPolicyClientWithBaseURI(api sdkEnv.Api) (*PolicyActivityBasedTimeoutPolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policyactivitybasedtimeoutpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyActivityBasedTimeoutPolicyClient: %+v", err)
	}

	return &PolicyActivityBasedTimeoutPolicyClient{
		Client: client,
	}, nil
}
