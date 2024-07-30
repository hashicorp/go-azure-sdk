package devicehealthscript

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

type ListDeviceHealthScriptsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceHealthScript
}

type ListDeviceHealthScriptsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceHealthScript
}

type ListDeviceHealthScriptsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceHealthScriptsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceHealthScripts ...
func (c DeviceHealthScriptClient) ListDeviceHealthScripts(ctx context.Context) (result ListDeviceHealthScriptsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceHealthScriptsCustomPager{},
		Path:       "/deviceManagement/deviceHealthScripts",
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
		Values *[]beta.DeviceHealthScript `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceHealthScriptsComplete retrieves all the results into a single object
func (c DeviceHealthScriptClient) ListDeviceHealthScriptsComplete(ctx context.Context) (ListDeviceHealthScriptsCompleteResult, error) {
	return c.ListDeviceHealthScriptsCompleteMatchingPredicate(ctx, DeviceHealthScriptOperationPredicate{})
}

// ListDeviceHealthScriptsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceHealthScriptClient) ListDeviceHealthScriptsCompleteMatchingPredicate(ctx context.Context, predicate DeviceHealthScriptOperationPredicate) (result ListDeviceHealthScriptsCompleteResult, err error) {
	items := make([]beta.DeviceHealthScript, 0)

	resp, err := c.ListDeviceHealthScripts(ctx)
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

	result = ListDeviceHealthScriptsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
