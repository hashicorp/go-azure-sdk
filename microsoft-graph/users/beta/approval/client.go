package approval

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApprovalClient struct {
	Client *msgraph.Client
}

func NewApprovalClientWithBaseURI(sdkApi sdkEnv.Api) (*ApprovalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "approval", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApprovalClient: %+v", err)
	}

	return &ApprovalClient{
		Client: client,
	}, nil
}
