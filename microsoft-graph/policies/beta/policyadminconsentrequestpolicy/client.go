package policyadminconsentrequestpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyAdminConsentRequestPolicyClient struct {
	Client *msgraph.Client
}

func NewPolicyAdminConsentRequestPolicyClientWithBaseURI(api sdkEnv.Api) (*PolicyAdminConsentRequestPolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policyadminconsentrequestpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyAdminConsentRequestPolicyClient: %+v", err)
	}

	return &PolicyAdminConsentRequestPolicyClient{
		Client: client,
	}, nil
}
