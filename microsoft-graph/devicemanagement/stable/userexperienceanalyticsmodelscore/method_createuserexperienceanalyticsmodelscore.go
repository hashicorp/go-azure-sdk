package userexperienceanalyticsmodelscore

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserExperienceAnalyticsModelScoreOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UserExperienceAnalyticsModelScores
}

type CreateUserExperienceAnalyticsModelScoreOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateUserExperienceAnalyticsModelScoreOperationOptions() CreateUserExperienceAnalyticsModelScoreOperationOptions {
	return CreateUserExperienceAnalyticsModelScoreOperationOptions{}
}

func (o CreateUserExperienceAnalyticsModelScoreOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateUserExperienceAnalyticsModelScoreOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateUserExperienceAnalyticsModelScoreOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateUserExperienceAnalyticsModelScore - Create new navigation property to userExperienceAnalyticsModelScores for
// deviceManagement
func (c UserExperienceAnalyticsModelScoreClient) CreateUserExperienceAnalyticsModelScore(ctx context.Context, input stable.UserExperienceAnalyticsModelScores, options CreateUserExperienceAnalyticsModelScoreOperationOptions) (result CreateUserExperienceAnalyticsModelScoreOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/userExperienceAnalyticsModelScores",
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

	var model stable.UserExperienceAnalyticsModelScores
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
