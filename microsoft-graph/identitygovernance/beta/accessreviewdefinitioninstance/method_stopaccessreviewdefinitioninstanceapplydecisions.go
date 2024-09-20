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

type StopAccessReviewDefinitionInstanceApplyDecisionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type StopAccessReviewDefinitionInstanceApplyDecisionsOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultStopAccessReviewDefinitionInstanceApplyDecisionsOperationOptions() StopAccessReviewDefinitionInstanceApplyDecisionsOperationOptions {
	return StopAccessReviewDefinitionInstanceApplyDecisionsOperationOptions{}
}

func (o StopAccessReviewDefinitionInstanceApplyDecisionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o StopAccessReviewDefinitionInstanceApplyDecisionsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o StopAccessReviewDefinitionInstanceApplyDecisionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// StopAccessReviewDefinitionInstanceApplyDecisions - Invoke action stopApplyDecisions
func (c AccessReviewDefinitionInstanceClient) StopAccessReviewDefinitionInstanceApplyDecisions(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceId, options StopAccessReviewDefinitionInstanceApplyDecisionsOperationOptions) (result StopAccessReviewDefinitionInstanceApplyDecisionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/stopApplyDecisions", id.ID()),
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
