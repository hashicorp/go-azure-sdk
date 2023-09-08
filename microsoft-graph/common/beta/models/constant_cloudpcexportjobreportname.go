package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCExportJobReportName string

const (
	CloudPCExportJobReportNamedailyAggregatedRemoteConnectionReports      CloudPCExportJobReportName = "DailyAggregatedRemoteConnectionReports"
	CloudPCExportJobReportNameinaccessibleCloudPcReports                  CloudPCExportJobReportName = "InaccessibleCloudPCReports"
	CloudPCExportJobReportNamenoLicenseAvailableConnectivityFailureReport CloudPCExportJobReportName = "NoLicenseAvailableConnectivityFailureReport"
	CloudPCExportJobReportNameremoteConnectionHistoricalReports           CloudPCExportJobReportName = "RemoteConnectionHistoricalReports"
	CloudPCExportJobReportNameremoteConnectionQualityReports              CloudPCExportJobReportName = "RemoteConnectionQualityReports"
	CloudPCExportJobReportNamesharedUseLicenseUsageRealTimeReport         CloudPCExportJobReportName = "SharedUseLicenseUsageRealTimeReport"
	CloudPCExportJobReportNamesharedUseLicenseUsageReport                 CloudPCExportJobReportName = "SharedUseLicenseUsageReport"
	CloudPCExportJobReportNametotalAggregatedRemoteConnectionReports      CloudPCExportJobReportName = "TotalAggregatedRemoteConnectionReports"
)

func PossibleValuesForCloudPCExportJobReportName() []string {
	return []string{
		string(CloudPCExportJobReportNamedailyAggregatedRemoteConnectionReports),
		string(CloudPCExportJobReportNameinaccessibleCloudPcReports),
		string(CloudPCExportJobReportNamenoLicenseAvailableConnectivityFailureReport),
		string(CloudPCExportJobReportNameremoteConnectionHistoricalReports),
		string(CloudPCExportJobReportNameremoteConnectionQualityReports),
		string(CloudPCExportJobReportNamesharedUseLicenseUsageRealTimeReport),
		string(CloudPCExportJobReportNamesharedUseLicenseUsageReport),
		string(CloudPCExportJobReportNametotalAggregatedRemoteConnectionReports),
	}
}

func parseCloudPCExportJobReportName(input string) (*CloudPCExportJobReportName, error) {
	vals := map[string]CloudPCExportJobReportName{
		"dailyaggregatedremoteconnectionreports":      CloudPCExportJobReportNamedailyAggregatedRemoteConnectionReports,
		"inaccessiblecloudpcreports":                  CloudPCExportJobReportNameinaccessibleCloudPcReports,
		"nolicenseavailableconnectivityfailurereport": CloudPCExportJobReportNamenoLicenseAvailableConnectivityFailureReport,
		"remoteconnectionhistoricalreports":           CloudPCExportJobReportNameremoteConnectionHistoricalReports,
		"remoteconnectionqualityreports":              CloudPCExportJobReportNameremoteConnectionQualityReports,
		"shareduselicenseusagerealtimereport":         CloudPCExportJobReportNamesharedUseLicenseUsageRealTimeReport,
		"shareduselicenseusagereport":                 CloudPCExportJobReportNamesharedUseLicenseUsageReport,
		"totalaggregatedremoteconnectionreports":      CloudPCExportJobReportNametotalAggregatedRemoteConnectionReports,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCExportJobReportName(input)
	return &out, nil
}
