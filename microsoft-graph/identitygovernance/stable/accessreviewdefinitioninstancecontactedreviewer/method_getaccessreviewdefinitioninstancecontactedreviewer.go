package accessreviewdefinitioninstancecontactedreviewer

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAccessReviewDefinitionInstanceContactedReviewerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AccessReviewReviewer
}

type GetAccessReviewDefinitionInstanceContactedReviewerOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetAccessReviewDefinitionInstanceContactedReviewerOperationOptions() GetAccessReviewDefinitionInstanceContactedReviewerOperationOptions {
	return GetAccessReviewDefinitionInstanceContactedReviewerOperationOptions{}
}

func (o GetAccessReviewDefinitionInstanceContactedReviewerOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAccessReviewDefinitionInstanceContactedReviewerOperationOptions) ToOData() *odata.Query {
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

func (o GetAccessReviewDefinitionInstanceContactedReviewerOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAccessReviewDefinitionInstanceContactedReviewer - Get contactedReviewers from identityGovernance. Returns the
// collection of reviewers who were contacted to complete this review. While the reviewers and fallbackReviewers
// properties of the accessReviewScheduleDefinition might specify group owners or managers as reviewers,
// contactedReviewers returns their individual identities. Supports $select. Read-only.
func (c AccessReviewDefinitionInstanceContactedReviewerClient) GetAccessReviewDefinitionInstanceContactedReviewer(ctx context.Context, id stable.IdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerId, options GetAccessReviewDefinitionInstanceContactedReviewerOperationOptions) (result GetAccessReviewDefinitionInstanceContactedReviewerOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
		RetryFunc:     options.RetryFunc,
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

	var model stable.AccessReviewReviewer
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
