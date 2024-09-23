package accessreviewdefinitioninstancestagedecisioninstance

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAccessReviewDefinitionInstanceStageDecisionInstanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AccessReviewInstance
}

type GetAccessReviewDefinitionInstanceStageDecisionInstanceOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetAccessReviewDefinitionInstanceStageDecisionInstanceOperationOptions() GetAccessReviewDefinitionInstanceStageDecisionInstanceOperationOptions {
	return GetAccessReviewDefinitionInstanceStageDecisionInstanceOperationOptions{}
}

func (o GetAccessReviewDefinitionInstanceStageDecisionInstanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAccessReviewDefinitionInstanceStageDecisionInstanceOperationOptions) ToOData() *odata.Query {
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

func (o GetAccessReviewDefinitionInstanceStageDecisionInstanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAccessReviewDefinitionInstanceStageDecisionInstance - Get instance from identityGovernance. There's exactly one
// accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the
// recurrence of the access review the decision is made on.
func (c AccessReviewDefinitionInstanceStageDecisionInstanceClient) GetAccessReviewDefinitionInstanceStageDecisionInstance(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionId, options GetAccessReviewDefinitionInstanceStageDecisionInstanceOperationOptions) (result GetAccessReviewDefinitionInstanceStageDecisionInstanceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/instance", id.ID()),
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

	var model beta.AccessReviewInstance
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
