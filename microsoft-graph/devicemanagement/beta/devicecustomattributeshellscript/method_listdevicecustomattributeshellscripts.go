package devicecustomattributeshellscript

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

type ListDeviceCustomAttributeShellScriptsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceCustomAttributeShellScript
}

type ListDeviceCustomAttributeShellScriptsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceCustomAttributeShellScript
}

type ListDeviceCustomAttributeShellScriptsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDeviceCustomAttributeShellScriptsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDeviceCustomAttributeShellScripts ...
func (c DeviceCustomAttributeShellScriptClient) ListDeviceCustomAttributeShellScripts(ctx context.Context) (result ListDeviceCustomAttributeShellScriptsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDeviceCustomAttributeShellScriptsCustomPager{},
		Path:       "/deviceManagement/deviceCustomAttributeShellScripts",
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
		Values *[]beta.DeviceCustomAttributeShellScript `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeviceCustomAttributeShellScriptsComplete retrieves all the results into a single object
func (c DeviceCustomAttributeShellScriptClient) ListDeviceCustomAttributeShellScriptsComplete(ctx context.Context) (ListDeviceCustomAttributeShellScriptsCompleteResult, error) {
	return c.ListDeviceCustomAttributeShellScriptsCompleteMatchingPredicate(ctx, DeviceCustomAttributeShellScriptOperationPredicate{})
}

// ListDeviceCustomAttributeShellScriptsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeviceCustomAttributeShellScriptClient) ListDeviceCustomAttributeShellScriptsCompleteMatchingPredicate(ctx context.Context, predicate DeviceCustomAttributeShellScriptOperationPredicate) (result ListDeviceCustomAttributeShellScriptsCompleteResult, err error) {
	items := make([]beta.DeviceCustomAttributeShellScript, 0)

	resp, err := c.ListDeviceCustomAttributeShellScripts(ctx)
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

	result = ListDeviceCustomAttributeShellScriptsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
