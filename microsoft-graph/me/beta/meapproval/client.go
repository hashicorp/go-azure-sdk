package meapproval

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeApprovalClient struct {
	Client *msgraph.Client
}

func NewMeApprovalClientWithBaseURI(api sdkEnv.Api) (*MeApprovalClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meapproval", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeApprovalClient: %+v", err)
	}

	return &MeApprovalClient{
		Client: client,
	}, nil
}
