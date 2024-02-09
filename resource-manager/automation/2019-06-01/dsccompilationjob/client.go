package dsccompilationjob

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DscCompilationJobClient struct {
	Client *resourcemanager.Client
}

func NewDscCompilationJobClientWithBaseURI(sdkApi sdkEnv.Api) (*DscCompilationJobClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "dsccompilationjob", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DscCompilationJobClient: %+v", err)
	}

	return &DscCompilationJobClient{
		Client: client,
	}, nil
}
