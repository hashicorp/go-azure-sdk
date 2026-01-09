package termsandconditiongroupassignmenttermsandcondition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermsAndConditionGroupAssignmentTermsAndConditionClient struct {
	Client *msgraph.Client
}

func NewTermsAndConditionGroupAssignmentTermsAndConditionClientWithBaseURI(sdkApi sdkEnv.Api) (*TermsAndConditionGroupAssignmentTermsAndConditionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "termsandconditiongroupassignmenttermsandcondition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TermsAndConditionGroupAssignmentTermsAndConditionClient: %+v", err)
	}

	return &TermsAndConditionGroupAssignmentTermsAndConditionClient{
		Client: client,
	}, nil
}
