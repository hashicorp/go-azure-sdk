package accessreviewdefinitioninstance

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

type ResetAccessReviewDefinitionInstanceDecisionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ResetAccessReviewDefinitionInstanceDecisionsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultResetAccessReviewDefinitionInstanceDecisionsOperationOptions() ResetAccessReviewDefinitionInstanceDecisionsOperationOptions {
	return ResetAccessReviewDefinitionInstanceDecisionsOperationOptions{}
}

func (o ResetAccessReviewDefinitionInstanceDecisionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ResetAccessReviewDefinitionInstanceDecisionsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ResetAccessReviewDefinitionInstanceDecisionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ResetAccessReviewDefinitionInstanceDecisions - Invoke action resetDecisions. Resets decisions of all
// accessReviewInstanceDecisionItem objects on an accessReviewInstance to notReviewed.
func (c AccessReviewDefinitionInstanceClient) ResetAccessReviewDefinitionInstanceDecisions(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceId, options ResetAccessReviewDefinitionInstanceDecisionsOperationOptions) (result ResetAccessReviewDefinitionInstanceDecisionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/resetDecisions", id.ID()),
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

	return
}
