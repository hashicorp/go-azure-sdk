package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TiIndicator struct {
	Action                      *TiIndicatorAction       `json:"action,omitempty"`
	ActivityGroupNames          *[]string                `json:"activityGroupNames,omitempty"`
	AdditionalInformation       *string                  `json:"additionalInformation,omitempty"`
	AzureTenantId               *string                  `json:"azureTenantId,omitempty"`
	Confidence                  *int64                   `json:"confidence,omitempty"`
	Description                 *string                  `json:"description,omitempty"`
	DiamondModel                *TiIndicatorDiamondModel `json:"diamondModel,omitempty"`
	DomainName                  *string                  `json:"domainName,omitempty"`
	EmailEncoding               *string                  `json:"emailEncoding,omitempty"`
	EmailLanguage               *string                  `json:"emailLanguage,omitempty"`
	EmailRecipient              *string                  `json:"emailRecipient,omitempty"`
	EmailSenderAddress          *string                  `json:"emailSenderAddress,omitempty"`
	EmailSenderName             *string                  `json:"emailSenderName,omitempty"`
	EmailSourceDomain           *string                  `json:"emailSourceDomain,omitempty"`
	EmailSourceIpAddress        *string                  `json:"emailSourceIpAddress,omitempty"`
	EmailSubject                *string                  `json:"emailSubject,omitempty"`
	EmailXMailer                *string                  `json:"emailXMailer,omitempty"`
	ExpirationDateTime          *string                  `json:"expirationDateTime,omitempty"`
	ExternalId                  *string                  `json:"externalId,omitempty"`
	FileCompileDateTime         *string                  `json:"fileCompileDateTime,omitempty"`
	FileCreatedDateTime         *string                  `json:"fileCreatedDateTime,omitempty"`
	FileHashType                *TiIndicatorFileHashType `json:"fileHashType,omitempty"`
	FileHashValue               *string                  `json:"fileHashValue,omitempty"`
	FileMutexName               *string                  `json:"fileMutexName,omitempty"`
	FileName                    *string                  `json:"fileName,omitempty"`
	FilePacker                  *string                  `json:"filePacker,omitempty"`
	FilePath                    *string                  `json:"filePath,omitempty"`
	FileSize                    *int64                   `json:"fileSize,omitempty"`
	FileType                    *string                  `json:"fileType,omitempty"`
	Id                          *string                  `json:"id,omitempty"`
	IngestedDateTime            *string                  `json:"ingestedDateTime,omitempty"`
	IsActive                    *bool                    `json:"isActive,omitempty"`
	KillChain                   *[]string                `json:"killChain,omitempty"`
	KnownFalsePositives         *string                  `json:"knownFalsePositives,omitempty"`
	LastReportedDateTime        *string                  `json:"lastReportedDateTime,omitempty"`
	MalwareFamilyNames          *[]string                `json:"malwareFamilyNames,omitempty"`
	NetworkCidrBlock            *string                  `json:"networkCidrBlock,omitempty"`
	NetworkDestinationAsn       *int64                   `json:"networkDestinationAsn,omitempty"`
	NetworkDestinationCidrBlock *string                  `json:"networkDestinationCidrBlock,omitempty"`
	NetworkDestinationIPv4      *string                  `json:"networkDestinationIPv4,omitempty"`
	NetworkDestinationIPv6      *string                  `json:"networkDestinationIPv6,omitempty"`
	NetworkDestinationPort      *int64                   `json:"networkDestinationPort,omitempty"`
	NetworkIPv4                 *string                  `json:"networkIPv4,omitempty"`
	NetworkIPv6                 *string                  `json:"networkIPv6,omitempty"`
	NetworkPort                 *int64                   `json:"networkPort,omitempty"`
	NetworkProtocol             *int64                   `json:"networkProtocol,omitempty"`
	NetworkSourceAsn            *int64                   `json:"networkSourceAsn,omitempty"`
	NetworkSourceCidrBlock      *string                  `json:"networkSourceCidrBlock,omitempty"`
	NetworkSourceIPv4           *string                  `json:"networkSourceIPv4,omitempty"`
	NetworkSourceIPv6           *string                  `json:"networkSourceIPv6,omitempty"`
	NetworkSourcePort           *int64                   `json:"networkSourcePort,omitempty"`
	ODataType                   *string                  `json:"@odata.type,omitempty"`
	PassiveOnly                 *bool                    `json:"passiveOnly,omitempty"`
	Severity                    *int64                   `json:"severity,omitempty"`
	Tags                        *[]string                `json:"tags,omitempty"`
	TargetProduct               *string                  `json:"targetProduct,omitempty"`
	ThreatType                  *string                  `json:"threatType,omitempty"`
	TlpLevel                    *TiIndicatorTlpLevel     `json:"tlpLevel,omitempty"`
	Url                         *string                  `json:"url,omitempty"`
	UserAgent                   *string                  `json:"userAgent,omitempty"`
}
