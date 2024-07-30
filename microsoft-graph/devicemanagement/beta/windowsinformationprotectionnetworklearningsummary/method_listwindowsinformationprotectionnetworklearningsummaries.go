package windowsinformationprotectionnetworklearningsummary

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

type ListWindowsInformationProtectionNetworkLearningSummariesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.WindowsInformationProtectionNetworkLearningSummary
}

type ListWindowsInformationProtectionNetworkLearningSummariesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.WindowsInformationProtectionNetworkLearningSummary
}

type ListWindowsInformationProtectionNetworkLearningSummariesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListWindowsInformationProtectionNetworkLearningSummariesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListWindowsInformationProtectionNetworkLearningSummaries ...
func (c WindowsInformationProtectionNetworkLearningSummaryClient) ListWindowsInformationProtectionNetworkLearningSummaries(ctx context.Context) (result ListWindowsInformationProtectionNetworkLearningSummariesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListWindowsInformationProtectionNetworkLearningSummariesCustomPager{},
		Path:       "/deviceManagement/windowsInformationProtectionNetworkLearningSummaries",
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
		Values *[]beta.WindowsInformationProtectionNetworkLearningSummary `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListWindowsInformationProtectionNetworkLearningSummariesComplete retrieves all the results into a single object
func (c WindowsInformationProtectionNetworkLearningSummaryClient) ListWindowsInformationProtectionNetworkLearningSummariesComplete(ctx context.Context) (ListWindowsInformationProtectionNetworkLearningSummariesCompleteResult, error) {
	return c.ListWindowsInformationProtectionNetworkLearningSummariesCompleteMatchingPredicate(ctx, WindowsInformationProtectionNetworkLearningSummaryOperationPredicate{})
}

// ListWindowsInformationProtectionNetworkLearningSummariesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WindowsInformationProtectionNetworkLearningSummaryClient) ListWindowsInformationProtectionNetworkLearningSummariesCompleteMatchingPredicate(ctx context.Context, predicate WindowsInformationProtectionNetworkLearningSummaryOperationPredicate) (result ListWindowsInformationProtectionNetworkLearningSummariesCompleteResult, err error) {
	items := make([]beta.WindowsInformationProtectionNetworkLearningSummary, 0)

	resp, err := c.ListWindowsInformationProtectionNetworkLearningSummaries(ctx)
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

	result = ListWindowsInformationProtectionNetworkLearningSummariesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
