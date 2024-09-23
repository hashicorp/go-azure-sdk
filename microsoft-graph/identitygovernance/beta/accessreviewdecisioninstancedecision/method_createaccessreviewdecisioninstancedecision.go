package accessreviewdecisioninstancedecision

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

type CreateAccessReviewDecisionInstanceDecisionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AccessReviewInstanceDecisionItem
}

type CreateAccessReviewDecisionInstanceDecisionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateAccessReviewDecisionInstanceDecisionOperationOptions() CreateAccessReviewDecisionInstanceDecisionOperationOptions {
	return CreateAccessReviewDecisionInstanceDecisionOperationOptions{}
}

func (o CreateAccessReviewDecisionInstanceDecisionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAccessReviewDecisionInstanceDecisionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAccessReviewDecisionInstanceDecisionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAccessReviewDecisionInstanceDecision - Create new navigation property to decisions for identityGovernance
func (c AccessReviewDecisionInstanceDecisionClient) CreateAccessReviewDecisionInstanceDecision(ctx context.Context, id beta.IdentityGovernanceAccessReviewDecisionId, input beta.AccessReviewInstanceDecisionItem, options CreateAccessReviewDecisionInstanceDecisionOperationOptions) (result CreateAccessReviewDecisionInstanceDecisionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/instance/decisions", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
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

	var model beta.AccessReviewInstanceDecisionItem
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
