package domainverificationdnsrecord

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDomainByIdVerificationDnsRecordsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.DomainDnsRecordCollectionResponse
}

type ListDomainByIdVerificationDnsRecordsCompleteResult struct {
	Items []models.DomainDnsRecordCollectionResponse
}

// ListDomainByIdVerificationDnsRecords ...
func (c DomainVerificationDnsRecordClient) ListDomainByIdVerificationDnsRecords(ctx context.Context, id DomainId) (result ListDomainByIdVerificationDnsRecordsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/verificationDnsRecords", id.ID()),
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
		Values *[]models.DomainDnsRecordCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDomainByIdVerificationDnsRecordsComplete retrieves all the results into a single object
func (c DomainVerificationDnsRecordClient) ListDomainByIdVerificationDnsRecordsComplete(ctx context.Context, id DomainId) (ListDomainByIdVerificationDnsRecordsCompleteResult, error) {
	return c.ListDomainByIdVerificationDnsRecordsCompleteMatchingPredicate(ctx, id, models.DomainDnsRecordCollectionResponseOperationPredicate{})
}

// ListDomainByIdVerificationDnsRecordsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DomainVerificationDnsRecordClient) ListDomainByIdVerificationDnsRecordsCompleteMatchingPredicate(ctx context.Context, id DomainId, predicate models.DomainDnsRecordCollectionResponseOperationPredicate) (result ListDomainByIdVerificationDnsRecordsCompleteResult, err error) {
	items := make([]models.DomainDnsRecordCollectionResponse, 0)

	resp, err := c.ListDomainByIdVerificationDnsRecords(ctx, id)
	if err != nil {
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

	result = ListDomainByIdVerificationDnsRecordsCompleteResult{
		Items: items,
	}
	return
}
