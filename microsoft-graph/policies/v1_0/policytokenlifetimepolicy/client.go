package policytokenlifetimepolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyTokenLifetimePolicyClient struct {
	Client *msgraph.Client
}

func NewPolicyTokenLifetimePolicyClientWithBaseURI(api sdkEnv.Api) (*PolicyTokenLifetimePolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policytokenlifetimepolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyTokenLifetimePolicyClient: %+v", err)
	}

	return &PolicyTokenLifetimePolicyClient{
		Client: client,
	}, nil
}
