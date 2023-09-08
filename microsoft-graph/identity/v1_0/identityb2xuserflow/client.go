package identityb2xuserflow

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityB2xUserFlowClient struct {
	Client *msgraph.Client
}

func NewIdentityB2xUserFlowClientWithBaseURI(api sdkEnv.Api) (*IdentityB2xUserFlowClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "identityb2xuserflow", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IdentityB2xUserFlowClient: %+v", err)
	}

	return &IdentityB2xUserFlowClient{
		Client: client,
	}, nil
}
