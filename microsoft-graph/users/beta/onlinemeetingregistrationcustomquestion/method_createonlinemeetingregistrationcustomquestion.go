package onlinemeetingregistrationcustomquestion

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

type CreateOnlineMeetingRegistrationCustomQuestionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.MeetingRegistrationQuestion
}

type CreateOnlineMeetingRegistrationCustomQuestionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateOnlineMeetingRegistrationCustomQuestionOperationOptions() CreateOnlineMeetingRegistrationCustomQuestionOperationOptions {
	return CreateOnlineMeetingRegistrationCustomQuestionOperationOptions{}
}

func (o CreateOnlineMeetingRegistrationCustomQuestionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateOnlineMeetingRegistrationCustomQuestionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateOnlineMeetingRegistrationCustomQuestionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateOnlineMeetingRegistrationCustomQuestion - Create new navigation property to customQuestions for users
func (c OnlineMeetingRegistrationCustomQuestionClient) CreateOnlineMeetingRegistrationCustomQuestion(ctx context.Context, id beta.UserIdOnlineMeetingId, input beta.MeetingRegistrationQuestion, options CreateOnlineMeetingRegistrationCustomQuestionOperationOptions) (result CreateOnlineMeetingRegistrationCustomQuestionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/registration/customQuestions", id.ID()),
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

	var model beta.MeetingRegistrationQuestion
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
