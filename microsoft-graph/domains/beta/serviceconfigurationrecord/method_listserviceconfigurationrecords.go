package serviceconfigurationrecord

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

type ListServiceConfigurationRecordsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DomainDnsRecord
}

type ListServiceConfigurationRecordsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DomainDnsRecord
}

type ListServiceConfigurationRecordsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListServiceConfigurationRecordsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListServiceConfigurationRecords ...
func (c ServiceConfigurationRecordClient) ListServiceConfigurationRecords(ctx context.Context, id DomainId) (result ListServiceConfigurationRecordsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListServiceConfigurationRecordsCustomPager{},
		Path:       fmt.Sprintf("%s/serviceConfigurationRecords", id.ID()),
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
		Values *[]beta.DomainDnsRecord `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListServiceConfigurationRecordsComplete retrieves all the results into a single object
func (c ServiceConfigurationRecordClient) ListServiceConfigurationRecordsComplete(ctx context.Context, id DomainId) (ListServiceConfigurationRecordsCompleteResult, error) {
	return c.ListServiceConfigurationRecordsCompleteMatchingPredicate(ctx, id, DomainDnsRecordOperationPredicate{})
}

// ListServiceConfigurationRecordsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ServiceConfigurationRecordClient) ListServiceConfigurationRecordsCompleteMatchingPredicate(ctx context.Context, id DomainId, predicate DomainDnsRecordOperationPredicate) (result ListServiceConfigurationRecordsCompleteResult, err error) {
	items := make([]beta.DomainDnsRecord, 0)

	resp, err := c.ListServiceConfigurationRecords(ctx, id)
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

	result = ListServiceConfigurationRecordsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
