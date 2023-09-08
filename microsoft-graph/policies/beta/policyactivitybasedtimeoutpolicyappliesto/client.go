package policyactivitybasedtimeoutpolicyappliesto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyActivityBasedTimeoutPolicyAppliesToClient struct {
	Client *msgraph.Client
}

func NewPolicyActivityBasedTimeoutPolicyAppliesToClientWithBaseURI(api sdkEnv.Api) (*PolicyActivityBasedTimeoutPolicyAppliesToClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policyactivitybasedtimeoutpolicyappliesto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyActivityBasedTimeoutPolicyAppliesToClient: %+v", err)
	}

	return &PolicyActivityBasedTimeoutPolicyAppliesToClient{
		Client: client,
	}, nil
}
