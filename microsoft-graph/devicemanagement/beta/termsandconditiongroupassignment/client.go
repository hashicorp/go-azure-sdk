package termsandconditiongroupassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermsAndConditionGroupAssignmentClient struct {
	Client *msgraph.Client
}

func NewTermsAndConditionGroupAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*TermsAndConditionGroupAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "termsandconditiongroupassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TermsAndConditionGroupAssignmentClient: %+v", err)
	}

	return &TermsAndConditionGroupAssignmentClient{
		Client: client,
	}, nil
}
