package employeeexperiencelearningcourseactivity

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetEmployeeExperienceLearningCourseActivityOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.LearningCourseActivity
}

type GetEmployeeExperienceLearningCourseActivityOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetEmployeeExperienceLearningCourseActivityOperationOptions() GetEmployeeExperienceLearningCourseActivityOperationOptions {
	return GetEmployeeExperienceLearningCourseActivityOperationOptions{}
}

func (o GetEmployeeExperienceLearningCourseActivityOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEmployeeExperienceLearningCourseActivityOperationOptions) ToOData() *odata.Query {
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

func (o GetEmployeeExperienceLearningCourseActivityOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEmployeeExperienceLearningCourseActivity - Get learningCourseActivity. Get the specified learningCourseActivity
// object using either an ID or an externalCourseActivityId of the learning provider, or a courseActivityId of a user.
func (c EmployeeExperienceLearningCourseActivityClient) GetEmployeeExperienceLearningCourseActivity(ctx context.Context, id stable.UserIdEmployeeExperienceLearningCourseActivityId, options GetEmployeeExperienceLearningCourseActivityOperationOptions) (result GetEmployeeExperienceLearningCourseActivityOperationResponse, err error) {
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := stable.UnmarshalLearningCourseActivityImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
