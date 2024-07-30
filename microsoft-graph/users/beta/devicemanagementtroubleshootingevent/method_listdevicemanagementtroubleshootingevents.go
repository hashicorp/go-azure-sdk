package devicemanagementtroubleshootingevent

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

type ListDeviceManagementTroubleshootingEventsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementTroubleshootingEvent
}

type ListDeviceManagementTroubleshootingEventsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementTroubleshootingEvent
}

type ListDeviceManagementTroubleshootingEventsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceManagementTroubleshootingEventsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceManagementTroubleshootingEvents ...
func (c DeviceManagementTroubleshootingEventClient) ListDeviceManagementTroubleshootingEvents(ctx context.Context, id UserId) (result ListDeviceManagementTroubleshootingEventsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceManagementTroubleshootingEventsCustomPager{},
		Path:       fmt.Sprintf("%s/deviceManagementTroubleshootingEvents", id.ID()),
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

// ListDeviceManagementTroubleshootingEventsComplete retrieves all the results into a single object
func (c DeviceManagementTroubleshootingEventClient) ListDeviceManagementTroubleshootingEventsComplete(ctx context.Context, id UserId) (ListDeviceManagementTroubleshootingEventsCompleteResult, error) {
	return c.ListDeviceManagementTroubleshootingEventsCompleteMatchingPredicate(ctx, id, DeviceManagementTroubleshootingEventOperationPredicate{})
}

// ListDeviceManagementTroubleshootingEventsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceManagementTroubleshootingEventClient) ListDeviceManagementTroubleshootingEventsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate DeviceManagementTroubleshootingEventOperationPredicate) (result ListDeviceManagementTroubleshootingEventsCompleteResult, err error) {
	items := make([]beta.DeviceManagementTroubleshootingEvent, 0)

	resp, err := c.ListDeviceManagementTroubleshootingEvents(ctx, id)
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

	result = ListDeviceManagementTroubleshootingEventsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
