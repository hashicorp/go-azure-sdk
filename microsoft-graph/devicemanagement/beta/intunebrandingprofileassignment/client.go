package intunebrandingprofileassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntuneBrandingProfileAssignmentClient struct {
	Client *msgraph.Client
}

func NewIntuneBrandingProfileAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*IntuneBrandingProfileAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "intunebrandingprofileassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IntuneBrandingProfileAssignmentClient: %+v", err)
	}

	return &IntuneBrandingProfileAssignmentClient{
		Client: client,
	}, nil
}
