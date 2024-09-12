package onlinemeetingregistrationregistrant

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListOnlineMeetingRegistrationRegistrantsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MeetingRegistrantBase
}

type ListOnlineMeetingRegistrationRegistrantsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MeetingRegistrantBase
}

type ListOnlineMeetingRegistrationRegistrantsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListOnlineMeetingRegistrationRegistrantsOperationOptions() ListOnlineMeetingRegistrationRegistrantsOperationOptions {
	return ListOnlineMeetingRegistrationRegistrantsOperationOptions{}
}

func (o ListOnlineMeetingRegistrationRegistrantsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOnlineMeetingRegistrationRegistrantsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListOnlineMeetingRegistrationRegistrantsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListOnlineMeetingRegistrationRegistrantsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOnlineMeetingRegistrationRegistrantsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOnlineMeetingRegistrationRegistrants - Get registrants from users. Registrants of the online meeting.
func (c OnlineMeetingRegistrationRegistrantClient) ListOnlineMeetingRegistrationRegistrants(ctx context.Context, id beta.UserIdOnlineMeetingId, options ListOnlineMeetingRegistrationRegistrantsOperationOptions) (result ListOnlineMeetingRegistrationRegistrantsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListOnlineMeetingRegistrationRegistrantsCustomPager{},
		Path:          fmt.Sprintf("%s/registration/registrants", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]beta.MeetingRegistrantBase, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalMeetingRegistrantBaseImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.MeetingRegistrantBase (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListOnlineMeetingRegistrationRegistrantsComplete retrieves all the results into a single object
func (c OnlineMeetingRegistrationRegistrantClient) ListOnlineMeetingRegistrationRegistrantsComplete(ctx context.Context, id beta.UserIdOnlineMeetingId, options ListOnlineMeetingRegistrationRegistrantsOperationOptions) (ListOnlineMeetingRegistrationRegistrantsCompleteResult, error) {
	return c.ListOnlineMeetingRegistrationRegistrantsCompleteMatchingPredicate(ctx, id, options, MeetingRegistrantBaseOperationPredicate{})
}

// ListOnlineMeetingRegistrationRegistrantsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OnlineMeetingRegistrationRegistrantClient) ListOnlineMeetingRegistrationRegistrantsCompleteMatchingPredicate(ctx context.Context, id beta.UserIdOnlineMeetingId, options ListOnlineMeetingRegistrationRegistrantsOperationOptions, predicate MeetingRegistrantBaseOperationPredicate) (result ListOnlineMeetingRegistrationRegistrantsCompleteResult, err error) {
	items := make([]beta.MeetingRegistrantBase, 0)

	resp, err := c.ListOnlineMeetingRegistrationRegistrants(ctx, id, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListOnlineMeetingRegistrationRegistrantsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
