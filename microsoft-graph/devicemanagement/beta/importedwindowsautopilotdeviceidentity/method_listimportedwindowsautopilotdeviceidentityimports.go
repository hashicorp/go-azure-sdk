package importedwindowsautopilotdeviceidentity

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

type ListImportedWindowsAutopilotDeviceIdentityImportsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ImportedWindowsAutopilotDeviceIdentity
}

type ListImportedWindowsAutopilotDeviceIdentityImportsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ImportedWindowsAutopilotDeviceIdentity
}

type ListImportedWindowsAutopilotDeviceIdentityImportsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListImportedWindowsAutopilotDeviceIdentityImportsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListImportedWindowsAutopilotDeviceIdentityImports ...
func (c ImportedWindowsAutopilotDeviceIdentityClient) ListImportedWindowsAutopilotDeviceIdentityImports(ctx context.Context, input ListImportedWindowsAutopilotDeviceIdentityImportsRequest) (result ListImportedWindowsAutopilotDeviceIdentityImportsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &ListImportedWindowsAutopilotDeviceIdentityImportsCustomPager{},
		Path:       "/deviceManagement/importedWindowsAutopilotDeviceIdentities/import",
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
		Values *[]beta.ImportedWindowsAutopilotDeviceIdentity `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListImportedWindowsAutopilotDeviceIdentityImportsComplete retrieves all the results into a single object
func (c ImportedWindowsAutopilotDeviceIdentityClient) ListImportedWindowsAutopilotDeviceIdentityImportsComplete(ctx context.Context, input ListImportedWindowsAutopilotDeviceIdentityImportsRequest) (ListImportedWindowsAutopilotDeviceIdentityImportsCompleteResult, error) {
	return c.ListImportedWindowsAutopilotDeviceIdentityImportsCompleteMatchingPredicate(ctx, input, ImportedWindowsAutopilotDeviceIdentityOperationPredicate{})
}

// ListImportedWindowsAutopilotDeviceIdentityImportsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ImportedWindowsAutopilotDeviceIdentityClient) ListImportedWindowsAutopilotDeviceIdentityImportsCompleteMatchingPredicate(ctx context.Context, input ListImportedWindowsAutopilotDeviceIdentityImportsRequest, predicate ImportedWindowsAutopilotDeviceIdentityOperationPredicate) (result ListImportedWindowsAutopilotDeviceIdentityImportsCompleteResult, err error) {
	items := make([]beta.ImportedWindowsAutopilotDeviceIdentity, 0)

	resp, err := c.ListImportedWindowsAutopilotDeviceIdentityImports(ctx, input)
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

	result = ListImportedWindowsAutopilotDeviceIdentityImportsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
