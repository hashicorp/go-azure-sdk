package windowsautopilotdeviceidentity

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

type ListWindowsAutopilotDeviceIdentitiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.WindowsAutopilotDeviceIdentity
}

type ListWindowsAutopilotDeviceIdentitiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.WindowsAutopilotDeviceIdentity
}

type ListWindowsAutopilotDeviceIdentitiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListWindowsAutopilotDeviceIdentitiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListWindowsAutopilotDeviceIdentities ...
func (c WindowsAutopilotDeviceIdentityClient) ListWindowsAutopilotDeviceIdentities(ctx context.Context) (result ListWindowsAutopilotDeviceIdentitiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListWindowsAutopilotDeviceIdentitiesCustomPager{},
		Path:       "/deviceManagement/windowsAutopilotDeviceIdentities",
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
		Values *[]beta.WindowsAutopilotDeviceIdentity `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListWindowsAutopilotDeviceIdentitiesComplete retrieves all the results into a single object
func (c WindowsAutopilotDeviceIdentityClient) ListWindowsAutopilotDeviceIdentitiesComplete(ctx context.Context) (ListWindowsAutopilotDeviceIdentitiesCompleteResult, error) {
	return c.ListWindowsAutopilotDeviceIdentitiesCompleteMatchingPredicate(ctx, WindowsAutopilotDeviceIdentityOperationPredicate{})
}

// ListWindowsAutopilotDeviceIdentitiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WindowsAutopilotDeviceIdentityClient) ListWindowsAutopilotDeviceIdentitiesCompleteMatchingPredicate(ctx context.Context, predicate WindowsAutopilotDeviceIdentityOperationPredicate) (result ListWindowsAutopilotDeviceIdentitiesCompleteResult, err error) {
	items := make([]beta.WindowsAutopilotDeviceIdentity, 0)

	resp, err := c.ListWindowsAutopilotDeviceIdentities(ctx)
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

	result = ListWindowsAutopilotDeviceIdentitiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
