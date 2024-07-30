package troubleshootingevent

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

type ListTroubleshootingEventsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementTroubleshootingEvent
}

type ListTroubleshootingEventsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementTroubleshootingEvent
}

type ListTroubleshootingEventsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTroubleshootingEventsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTroubleshootingEvents ...
func (c TroubleshootingEventClient) ListTroubleshootingEvents(ctx context.Context) (result ListTroubleshootingEventsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTroubleshootingEventsCustomPager{},
		Path:       "/deviceManagement/troubleshootingEvents",
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
		Values *[]beta.DeviceManagementTroubleshootingEvent `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTroubleshootingEventsComplete retrieves all the results into a single object
func (c TroubleshootingEventClient) ListTroubleshootingEventsComplete(ctx context.Context) (ListTroubleshootingEventsCompleteResult, error) {
	return c.ListTroubleshootingEventsCompleteMatchingPredicate(ctx, DeviceManagementTroubleshootingEventOperationPredicate{})
}

// ListTroubleshootingEventsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TroubleshootingEventClient) ListTroubleshootingEventsCompleteMatchingPredicate(ctx context.Context, predicate DeviceManagementTroubleshootingEventOperationPredicate) (result ListTroubleshootingEventsCompleteResult, err error) {
	items := make([]beta.DeviceManagementTroubleshootingEvent, 0)

	resp, err := c.ListTroubleshootingEvents(ctx)
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

	result = ListTroubleshootingEventsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
