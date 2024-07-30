package autopilotevent

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

type ListAutopilotEventsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementAutopilotEvent
}

type ListAutopilotEventsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementAutopilotEvent
}

type ListAutopilotEventsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAutopilotEventsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAutopilotEvents ...
func (c AutopilotEventClient) ListAutopilotEvents(ctx context.Context) (result ListAutopilotEventsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAutopilotEventsCustomPager{},
		Path:       "/deviceManagement/autopilotEvents",
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
		Values *[]beta.DeviceManagementAutopilotEvent `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAutopilotEventsComplete retrieves all the results into a single object
func (c AutopilotEventClient) ListAutopilotEventsComplete(ctx context.Context) (ListAutopilotEventsCompleteResult, error) {
	return c.ListAutopilotEventsCompleteMatchingPredicate(ctx, DeviceManagementAutopilotEventOperationPredicate{})
}

// ListAutopilotEventsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AutopilotEventClient) ListAutopilotEventsCompleteMatchingPredicate(ctx context.Context, predicate DeviceManagementAutopilotEventOperationPredicate) (result ListAutopilotEventsCompleteResult, err error) {
	items := make([]beta.DeviceManagementAutopilotEvent, 0)

	resp, err := c.ListAutopilotEvents(ctx)
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

	result = ListAutopilotEventsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
