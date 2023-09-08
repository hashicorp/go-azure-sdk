package userapprovalstep

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserApprovalStepClient struct {
	Client *msgraph.Client
}

func NewUserApprovalStepClientWithBaseURI(api sdkEnv.Api) (*UserApprovalStepClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userapprovalstep", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserApprovalStepClient: %+v", err)
	}

	return &UserApprovalStepClient{
		Client: client,
	}, nil
}
