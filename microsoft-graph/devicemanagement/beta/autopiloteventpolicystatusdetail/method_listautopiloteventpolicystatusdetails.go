package autopiloteventpolicystatusdetail

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

type ListAutopilotEventPolicyStatusDetailsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementAutopilotPolicyStatusDetail
}

type ListAutopilotEventPolicyStatusDetailsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementAutopilotPolicyStatusDetail
}

type ListAutopilotEventPolicyStatusDetailsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAutopilotEventPolicyStatusDetailsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAutopilotEventPolicyStatusDetails ...
func (c AutopilotEventPolicyStatusDetailClient) ListAutopilotEventPolicyStatusDetails(ctx context.Context, id DeviceManagementAutopilotEventId) (result ListAutopilotEventPolicyStatusDetailsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAutopilotEventPolicyStatusDetailsCustomPager{},
		Path:       fmt.Sprintf("%s/policyStatusDetails", id.ID()),
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
		Values *[]beta.DeviceManagementAutopilotPolicyStatusDetail `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAutopilotEventPolicyStatusDetailsComplete retrieves all the results into a single object
func (c AutopilotEventPolicyStatusDetailClient) ListAutopilotEventPolicyStatusDetailsComplete(ctx context.Context, id DeviceManagementAutopilotEventId) (ListAutopilotEventPolicyStatusDetailsCompleteResult, error) {
	return c.ListAutopilotEventPolicyStatusDetailsCompleteMatchingPredicate(ctx, id, DeviceManagementAutopilotPolicyStatusDetailOperationPredicate{})
}

// ListAutopilotEventPolicyStatusDetailsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AutopilotEventPolicyStatusDetailClient) ListAutopilotEventPolicyStatusDetailsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementAutopilotEventId, predicate DeviceManagementAutopilotPolicyStatusDetailOperationPredicate) (result ListAutopilotEventPolicyStatusDetailsCompleteResult, err error) {
	items := make([]beta.DeviceManagementAutopilotPolicyStatusDetail, 0)

	resp, err := c.ListAutopilotEventPolicyStatusDetails(ctx, id)
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

	result = ListAutopilotEventPolicyStatusDetailsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
