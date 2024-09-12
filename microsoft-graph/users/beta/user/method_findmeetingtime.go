package user

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

type FindMeetingTimeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.MeetingTimeSuggestionsResult
}

// FindMeetingTime - Invoke action findMeetingTimes. Suggest meeting times and locations based on organizer and attendee
// availability, and time or location constraints specified as parameters. If findMeetingTimes cannot return any meeting
// suggestions, the response would indicate a reason in the emptySuggestionsReason property. Based on this value, you
// can better adjust the parameters and call findMeetingTimes again. The algorithm used to suggest meeting times and
// locations undergoes fine-tuning from time to time. In scenarios like test environments where the input parameters and
// calendar data remain static, expect that the suggested results may differ over time.
func (c UserClient) FindMeetingTime(ctx context.Context, id beta.UserId, input FindMeetingTimeRequest) (result FindMeetingTimeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/findMeetingTimes", id.ID()),
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

	var model beta.MeetingTimeSuggestionsResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
