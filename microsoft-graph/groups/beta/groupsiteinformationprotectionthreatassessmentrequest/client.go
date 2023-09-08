package groupsiteinformationprotectionthreatassessmentrequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteInformationProtectionThreatAssessmentRequestClient struct {
	Client *msgraph.Client
}

func NewGroupSiteInformationProtectionThreatAssessmentRequestClientWithBaseURI(api sdkEnv.Api) (*GroupSiteInformationProtectionThreatAssessmentRequestClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteinformationprotectionthreatassessmentrequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteInformationProtectionThreatAssessmentRequestClient: %+v", err)
	}

	return &GroupSiteInformationProtectionThreatAssessmentRequestClient{
		Client: client,
	}, nil
}
