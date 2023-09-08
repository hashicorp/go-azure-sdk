package policyconditionalaccesspolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyConditionalAccessPolicyClient struct {
	Client *msgraph.Client
}

func NewPolicyConditionalAccessPolicyClientWithBaseURI(api sdkEnv.Api) (*PolicyConditionalAccessPolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policyconditionalaccesspolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyConditionalAccessPolicyClient: %+v", err)
	}

	return &PolicyConditionalAccessPolicyClient{
		Client: client,
	}, nil
}
