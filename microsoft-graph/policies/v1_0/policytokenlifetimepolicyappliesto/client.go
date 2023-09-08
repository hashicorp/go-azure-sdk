package policytokenlifetimepolicyappliesto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyTokenLifetimePolicyAppliesToClient struct {
	Client *msgraph.Client
}

func NewPolicyTokenLifetimePolicyAppliesToClientWithBaseURI(api sdkEnv.Api) (*PolicyTokenLifetimePolicyAppliesToClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policytokenlifetimepolicyappliesto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyTokenLifetimePolicyAppliesToClient: %+v", err)
	}

	return &PolicyTokenLifetimePolicyAppliesToClient{
		Client: client,
	}, nil
}
