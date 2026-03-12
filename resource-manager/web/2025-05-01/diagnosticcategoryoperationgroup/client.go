package diagnosticcategoryoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiagnosticCategoryOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewDiagnosticCategoryOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*DiagnosticCategoryOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "diagnosticcategoryoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DiagnosticCategoryOperationGroupClient: %+v", err)
	}

	return &DiagnosticCategoryOperationGroupClient{
		Client: client,
	}, nil
}
