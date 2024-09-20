package assessments

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssessmentsClient struct {
	Client *resourcemanager.Client
}

func NewAssessmentsClientWithBaseURI(sdkApi sdkEnv.Api) (*AssessmentsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "assessments", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AssessmentsClient: %+v", err)
	}

	return &AssessmentsClient{
		Client: client,
	}, nil
}
