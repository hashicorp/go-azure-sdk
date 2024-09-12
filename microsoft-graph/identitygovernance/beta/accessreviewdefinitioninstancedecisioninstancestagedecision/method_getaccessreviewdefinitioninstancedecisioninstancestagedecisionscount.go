package accessreviewdefinitioninstancedecisioninstancestagedecision

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

type GetAccessReviewDefinitionInstanceDecisionInstanceStageDecisionsCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetAccessReviewDefinitionInstanceDecisionInstanceStageDecisionsCountOperationOptions struct {
	Filter *string
	Search *string
}

func DefaultGetAccessReviewDefinitionInstanceDecisionInstanceStageDecisionsCountOperationOptions() GetAccessReviewDefinitionInstanceDecisionInstanceStageDecisionsCountOperationOptions {
	return GetAccessReviewDefinitionInstanceDecisionInstanceStageDecisionsCountOperationOptions{}
}

func (o GetAccessReviewDefinitionInstanceDecisionInstanceStageDecisionsCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAccessReviewDefinitionInstanceDecisionInstanceStageDecisionsCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	return &out
}

func (o GetAccessReviewDefinitionInstanceDecisionInstanceStageDecisionsCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAccessReviewDefinitionInstanceDecisionInstanceStageDecisionsCount - Get the number of the resource
func (c AccessReviewDefinitionInstanceDecisionInstanceStageDecisionClient) GetAccessReviewDefinitionInstanceDecisionInstanceStageDecisionsCount(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageId, options GetAccessReviewDefinitionInstanceDecisionInstanceStageDecisionsCountOperationOptions) (result GetAccessReviewDefinitionInstanceDecisionInstanceStageDecisionsCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/decisions/$count", id.ID()),
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
