package accessreviewdefinitioninstancestage

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateAccessReviewDefinitionInstanceStageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateAccessReviewDefinitionInstanceStageOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateAccessReviewDefinitionInstanceStageOperationOptions() UpdateAccessReviewDefinitionInstanceStageOperationOptions {
	return UpdateAccessReviewDefinitionInstanceStageOperationOptions{}
}

func (o UpdateAccessReviewDefinitionInstanceStageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateAccessReviewDefinitionInstanceStageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateAccessReviewDefinitionInstanceStageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateAccessReviewDefinitionInstanceStage - Update accessReviewStage. Update the properties of an accessReviewStage
// object. Only the reviewers and fallbackReviewers properties can be updated. You can only add reviewers to the
// fallbackReviewers property but can't remove existing fallbackReviewers. To update an accessReviewStage, its status
// must be NotStarted, Initializing, or InProgress.
func (c AccessReviewDefinitionInstanceStageClient) UpdateAccessReviewDefinitionInstanceStage(ctx context.Context, id stable.IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageId, input stable.AccessReviewStage, options UpdateAccessReviewDefinitionInstanceStageOperationOptions) (result UpdateAccessReviewDefinitionInstanceStageOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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

	return
}
