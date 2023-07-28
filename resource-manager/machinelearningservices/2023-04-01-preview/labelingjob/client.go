package labelingjob

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LabelingJobClient struct {
	Client *resourcemanager.Client
}

func NewLabelingJobClientWithBaseURI(api environments.Api) (*LabelingJobClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "labelingjob", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LabelingJobClient: %+v", err)
	}

	return &LabelingJobClient{
		Client: client,
	}, nil
}
