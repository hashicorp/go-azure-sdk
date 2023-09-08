package policytokenissuancepolicyappliesto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyTokenIssuancePolicyAppliesToClient struct {
	Client *msgraph.Client
}

func NewPolicyTokenIssuancePolicyAppliesToClientWithBaseURI(api sdkEnv.Api) (*PolicyTokenIssuancePolicyAppliesToClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policytokenissuancepolicyappliesto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyTokenIssuancePolicyAppliesToClient: %+v", err)
	}

	return &PolicyTokenIssuancePolicyAppliesToClient{
		Client: client,
	}, nil
}
