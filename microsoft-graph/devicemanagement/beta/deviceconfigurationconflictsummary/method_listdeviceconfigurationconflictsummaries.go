package deviceconfigurationconflictsummary

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

type ListDeviceConfigurationConflictSummariesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceConfigurationConflictSummary
}

type ListDeviceConfigurationConflictSummariesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceConfigurationConflictSummary
}

type ListDeviceConfigurationConflictSummariesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceConfigurationConflictSummariesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceConfigurationConflictSummaries ...
func (c DeviceConfigurationConflictSummaryClient) ListDeviceConfigurationConflictSummaries(ctx context.Context) (result ListDeviceConfigurationConflictSummariesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceConfigurationConflictSummariesCustomPager{},
		Path:       "/deviceManagement/deviceConfigurationConflictSummary",
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
		Values *[]beta.DeviceConfigurationConflictSummary `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceConfigurationConflictSummariesComplete retrieves all the results into a single object
func (c DeviceConfigurationConflictSummaryClient) ListDeviceConfigurationConflictSummariesComplete(ctx context.Context) (ListDeviceConfigurationConflictSummariesCompleteResult, error) {
	return c.ListDeviceConfigurationConflictSummariesCompleteMatchingPredicate(ctx, DeviceConfigurationConflictSummaryOperationPredicate{})
}

// ListDeviceConfigurationConflictSummariesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceConfigurationConflictSummaryClient) ListDeviceConfigurationConflictSummariesCompleteMatchingPredicate(ctx context.Context, predicate DeviceConfigurationConflictSummaryOperationPredicate) (result ListDeviceConfigurationConflictSummariesCompleteResult, err error) {
	items := make([]beta.DeviceConfigurationConflictSummary, 0)

	resp, err := c.ListDeviceConfigurationConflictSummaries(ctx)
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

	result = ListDeviceConfigurationConflictSummariesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
