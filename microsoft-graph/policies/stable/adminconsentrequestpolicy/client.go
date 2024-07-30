package adminconsentrequestpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdminConsentRequestPolicyClient struct {
	Client *msgraph.Client
}

func NewAdminConsentRequestPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*AdminConsentRequestPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "adminconsentrequestpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AdminConsentRequestPolicyClient: %+v", err)
	}

	return &AdminConsentRequestPolicyClient{
		Client: client,
	}, nil
}
