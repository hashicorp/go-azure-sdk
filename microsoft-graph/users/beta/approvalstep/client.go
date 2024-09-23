package approvalstep

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApprovalStepClient struct {
	Client *msgraph.Client
}

func NewApprovalStepClientWithBaseURI(sdkApi sdkEnv.Api) (*ApprovalStepClient, error) {
	client, err := msgraph.NewClient(sdkApi, "approvalstep", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApprovalStepClient: %+v", err)
	}

	return &ApprovalStepClient{
		Client: client,
	}, nil
}
