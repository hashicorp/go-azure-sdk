package onlinemeetingregistrationcustomquestion

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateOnlineMeetingRegistrationCustomQuestionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateOnlineMeetingRegistrationCustomQuestionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateOnlineMeetingRegistrationCustomQuestionOperationOptions() UpdateOnlineMeetingRegistrationCustomQuestionOperationOptions {
	return UpdateOnlineMeetingRegistrationCustomQuestionOperationOptions{}
}

func (o UpdateOnlineMeetingRegistrationCustomQuestionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateOnlineMeetingRegistrationCustomQuestionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateOnlineMeetingRegistrationCustomQuestionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateOnlineMeetingRegistrationCustomQuestion - Update meetingRegistrationQuestion. Update a custom registration
// question associated with a meetingRegistration object on behalf of the organizer.
func (c OnlineMeetingRegistrationCustomQuestionClient) UpdateOnlineMeetingRegistrationCustomQuestion(ctx context.Context, id beta.MeOnlineMeetingIdRegistrationCustomQuestionId, input beta.MeetingRegistrationQuestion, options UpdateOnlineMeetingRegistrationCustomQuestionOperationOptions) (result UpdateOnlineMeetingRegistrationCustomQuestionOperationResponse, err error) {
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
