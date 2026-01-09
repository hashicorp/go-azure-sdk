package operationapprovalpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationApprovalPolicyClient struct {
	Client *msgraph.Client
}

func NewOperationApprovalPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*OperationApprovalPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "operationapprovalpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OperationApprovalPolicyClient: %+v", err)
	}

	return &OperationApprovalPolicyClient{
		Client: client,
	}, nil
}
