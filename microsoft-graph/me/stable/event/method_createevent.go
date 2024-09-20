package event

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateEventOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.Event
}

type CreateEventOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateEventOperationOptions() CreateEventOperationOptions {
	return CreateEventOperationOptions{}
}

func (o CreateEventOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEventOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEventOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEvent - Create event. Create an event in the user's default calendar or specified calendar. By default, the
// allowNewTimeProposals property is set to true when an event is created, which means invitees can propose a different
// date/time for the event. See Propose new meeting times for more information on how to propose a time, and how to
// receive and accept a new time proposal. You can specify the time zone for each of the start and end times of the
// event as part of their values, because the start and end properties are of dateTimeTimeZone type. First find the
// supported time zones to make sure you set only time zones that have been configured for the user's mailbox server.
// When an event is sent, the server sends invitations to all the attendees. Setting the location in an event An
// Exchange administrator can set up a mailbox and an email address for a resource such as a meeting room, or equipment
// like a projector. Users can then invite the resource as an attendee to a meeting. On behalf of the resource, the
// server accepts or rejects the meeting request based on the free/busy schedule of the resource. If the server accepts
// a meeting for the resource, it creates an event for the meeting in the resource's calendar. If the meeting is
// rescheduled, the server automatically updates the event in the resource's calendar. Another advantage of setting up a
// mailbox for a resource is to control scheduling of the resource, for example, only executives or their delegates can
// book a private meeting room. If you're organizing an event that involves a meeting location: Additionally, if the
// meeting location has been set up as a resource, or if the event involves some equipment that has been set up as a
// resource
func (c EventClient) CreateEvent(ctx context.Context, input stable.Event, options CreateEventOperationOptions) (result CreateEventOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/me/events",
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

	var model stable.Event
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
