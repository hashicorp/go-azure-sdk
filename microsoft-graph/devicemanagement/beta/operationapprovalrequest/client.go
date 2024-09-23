package operationapprovalrequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationApprovalRequestClient struct {
	Client *msgraph.Client
}

func NewOperationApprovalRequestClientWithBaseURI(sdkApi sdkEnv.Api) (*OperationApprovalRequestClient, error) {
	client, err := msgraph.NewClient(sdkApi, "operationapprovalrequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OperationApprovalRequestClient: %+v", err)
	}

	return &OperationApprovalRequestClient{
		Client: client,
	}, nil
}
