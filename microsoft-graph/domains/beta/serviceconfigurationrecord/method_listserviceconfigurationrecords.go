package serviceconfigurationrecord

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
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

type ListServiceConfigurationRecordsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListServiceConfigurationRecordsOperationOptions() ListServiceConfigurationRecordsOperationOptions {
	return ListServiceConfigurationRecordsOperationOptions{}
}

func (o ListServiceConfigurationRecordsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListServiceConfigurationRecordsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListServiceConfigurationRecordsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
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

// ListServiceConfigurationRecords - List serviceConfigurationRecords. Retrieves a list of domainDnsRecord objects
// needed to enable services for the domain. Use the returned list to add records to the zone file of the domain. This
// can be done through the domain registrar or DNS server configuration.
func (c ServiceConfigurationRecordClient) ListServiceConfigurationRecords(ctx context.Context, id beta.DomainId, options ListServiceConfigurationRecordsOperationOptions) (result ListServiceConfigurationRecordsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListServiceConfigurationRecordsCustomPager{},
		Path:          fmt.Sprintf("%s/serviceConfigurationRecords", id.ID()),
		RetryFunc:     options.RetryFunc,
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]beta.DomainDnsRecord, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := beta.UnmarshalDomainDnsRecordImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for beta.DomainDnsRecord (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ListServiceConfigurationRecordsComplete retrieves all the results into a single object
func (c ServiceConfigurationRecordClient) ListServiceConfigurationRecordsComplete(ctx context.Context, id beta.DomainId, options ListServiceConfigurationRecordsOperationOptions) (ListServiceConfigurationRecordsCompleteResult, error) {
	return c.ListServiceConfigurationRecordsCompleteMatchingPredicate(ctx, id, options, DomainDnsRecordOperationPredicate{})
}

// ListServiceConfigurationRecordsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ServiceConfigurationRecordClient) ListServiceConfigurationRecordsCompleteMatchingPredicate(ctx context.Context, id beta.DomainId, options ListServiceConfigurationRecordsOperationOptions, predicate DomainDnsRecordOperationPredicate) (result ListServiceConfigurationRecordsCompleteResult, err error) {
	items := make([]beta.DomainDnsRecord, 0)

	resp, err := c.ListServiceConfigurationRecords(ctx, id, options)
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
