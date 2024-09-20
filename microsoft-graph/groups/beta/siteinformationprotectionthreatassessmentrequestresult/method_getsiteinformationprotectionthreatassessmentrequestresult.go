package siteinformationprotectionthreatassessmentrequestresult

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetSiteInformationProtectionThreatAssessmentRequestResultOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ThreatAssessmentResult
}

type GetSiteInformationProtectionThreatAssessmentRequestResultOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetSiteInformationProtectionThreatAssessmentRequestResultOperationOptions() GetSiteInformationProtectionThreatAssessmentRequestResultOperationOptions {
	return GetSiteInformationProtectionThreatAssessmentRequestResultOperationOptions{}
}

func (o GetSiteInformationProtectionThreatAssessmentRequestResultOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetSiteInformationProtectionThreatAssessmentRequestResultOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetSiteInformationProtectionThreatAssessmentRequestResultOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetSiteInformationProtectionThreatAssessmentRequestResult - Get results from groups. A collection of threat
// assessment results. Read-only. By default, a GET /threatAssessmentRequests/{id} does not return this property unless
// you apply $expand on it.
func (c SiteInformationProtectionThreatAssessmentRequestResultClient) GetSiteInformationProtectionThreatAssessmentRequestResult(ctx context.Context, id beta.GroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultId, options GetSiteInformationProtectionThreatAssessmentRequestResultOperationOptions) (result GetSiteInformationProtectionThreatAssessmentRequestResultOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var model beta.ThreatAssessmentResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
