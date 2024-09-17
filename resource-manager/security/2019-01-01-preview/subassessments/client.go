package subassessments

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubAssessmentsClient struct {
	Client *resourcemanager.Client
}

func NewSubAssessmentsClientWithBaseURI(sdkApi sdkEnv.Api) (*SubAssessmentsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "subassessments", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SubAssessmentsClient: %+v", err)
	}

	return &SubAssessmentsClient{
		Client: client,
	}, nil
}
