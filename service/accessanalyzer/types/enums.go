// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AccessPreviewStatus string

// Enum values for AccessPreviewStatus
const (
	AccessPreviewStatusCompleted AccessPreviewStatus = "COMPLETED"
	AccessPreviewStatusCreating  AccessPreviewStatus = "CREATING"
	AccessPreviewStatusFailed    AccessPreviewStatus = "FAILED"
)

// Values returns all known values for AccessPreviewStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AccessPreviewStatus) Values() []AccessPreviewStatus {
	return []AccessPreviewStatus{
		"COMPLETED",
		"CREATING",
		"FAILED",
	}
}

type AccessPreviewStatusReasonCode string

// Enum values for AccessPreviewStatusReasonCode
const (
	AccessPreviewStatusReasonCodeInternalError        AccessPreviewStatusReasonCode = "INTERNAL_ERROR"
	AccessPreviewStatusReasonCodeInvalidConfiguration AccessPreviewStatusReasonCode = "INVALID_CONFIGURATION"
)

// Values returns all known values for AccessPreviewStatusReasonCode. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (AccessPreviewStatusReasonCode) Values() []AccessPreviewStatusReasonCode {
	return []AccessPreviewStatusReasonCode{
		"INTERNAL_ERROR",
		"INVALID_CONFIGURATION",
	}
}

type AclPermission string

// Enum values for AclPermission
const (
	AclPermissionRead        AclPermission = "READ"
	AclPermissionWrite       AclPermission = "WRITE"
	AclPermissionReadAcp     AclPermission = "READ_ACP"
	AclPermissionWriteAcp    AclPermission = "WRITE_ACP"
	AclPermissionFullControl AclPermission = "FULL_CONTROL"
)

// Values returns all known values for AclPermission. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AclPermission) Values() []AclPermission {
	return []AclPermission{
		"READ",
		"WRITE",
		"READ_ACP",
		"WRITE_ACP",
		"FULL_CONTROL",
	}
}

type AnalyzerStatus string

// Enum values for AnalyzerStatus
const (
	AnalyzerStatusActive   AnalyzerStatus = "ACTIVE"
	AnalyzerStatusCreating AnalyzerStatus = "CREATING"
	AnalyzerStatusDisabled AnalyzerStatus = "DISABLED"
	AnalyzerStatusFailed   AnalyzerStatus = "FAILED"
)

// Values returns all known values for AnalyzerStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AnalyzerStatus) Values() []AnalyzerStatus {
	return []AnalyzerStatus{
		"ACTIVE",
		"CREATING",
		"DISABLED",
		"FAILED",
	}
}

type FindingChangeType string

// Enum values for FindingChangeType
const (
	FindingChangeTypeChanged   FindingChangeType = "CHANGED"
	FindingChangeTypeNew       FindingChangeType = "NEW"
	FindingChangeTypeUnchanged FindingChangeType = "UNCHANGED"
)

// Values returns all known values for FindingChangeType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FindingChangeType) Values() []FindingChangeType {
	return []FindingChangeType{
		"CHANGED",
		"NEW",
		"UNCHANGED",
	}
}

type FindingSourceType string

// Enum values for FindingSourceType
const (
	FindingSourceTypePolicy               FindingSourceType = "POLICY"
	FindingSourceTypeBucketAcl            FindingSourceType = "BUCKET_ACL"
	FindingSourceTypeS3AccessPoint        FindingSourceType = "S3_ACCESS_POINT"
	FindingSourceTypeS3AccessPointAccount FindingSourceType = "S3_ACCESS_POINT_ACCOUNT"
)

// Values returns all known values for FindingSourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FindingSourceType) Values() []FindingSourceType {
	return []FindingSourceType{
		"POLICY",
		"BUCKET_ACL",
		"S3_ACCESS_POINT",
		"S3_ACCESS_POINT_ACCOUNT",
	}
}

type FindingStatus string

// Enum values for FindingStatus
const (
	FindingStatusActive   FindingStatus = "ACTIVE"
	FindingStatusArchived FindingStatus = "ARCHIVED"
	FindingStatusResolved FindingStatus = "RESOLVED"
)

// Values returns all known values for FindingStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FindingStatus) Values() []FindingStatus {
	return []FindingStatus{
		"ACTIVE",
		"ARCHIVED",
		"RESOLVED",
	}
}

type FindingStatusUpdate string

// Enum values for FindingStatusUpdate
const (
	FindingStatusUpdateActive   FindingStatusUpdate = "ACTIVE"
	FindingStatusUpdateArchived FindingStatusUpdate = "ARCHIVED"
)

// Values returns all known values for FindingStatusUpdate. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FindingStatusUpdate) Values() []FindingStatusUpdate {
	return []FindingStatusUpdate{
		"ACTIVE",
		"ARCHIVED",
	}
}

type JobErrorCode string

// Enum values for JobErrorCode
const (
	JobErrorCodeAuthorizationError        JobErrorCode = "AUTHORIZATION_ERROR"
	JobErrorCodeResourceNotFoundError     JobErrorCode = "RESOURCE_NOT_FOUND_ERROR"
	JobErrorCodeServiceQuotaExceededError JobErrorCode = "SERVICE_QUOTA_EXCEEDED_ERROR"
	JobErrorCodeServiceError              JobErrorCode = "SERVICE_ERROR"
)

// Values returns all known values for JobErrorCode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (JobErrorCode) Values() []JobErrorCode {
	return []JobErrorCode{
		"AUTHORIZATION_ERROR",
		"RESOURCE_NOT_FOUND_ERROR",
		"SERVICE_QUOTA_EXCEEDED_ERROR",
		"SERVICE_ERROR",
	}
}

type JobStatus string

// Enum values for JobStatus
const (
	JobStatusInProgress JobStatus = "IN_PROGRESS"
	JobStatusSucceeded  JobStatus = "SUCCEEDED"
	JobStatusFailed     JobStatus = "FAILED"
	JobStatusCanceled   JobStatus = "CANCELED"
)

// Values returns all known values for JobStatus. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (JobStatus) Values() []JobStatus {
	return []JobStatus{
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
		"CANCELED",
	}
}

type KmsGrantOperation string

// Enum values for KmsGrantOperation
const (
	KmsGrantOperationCreateGrant                         KmsGrantOperation = "CreateGrant"
	KmsGrantOperationDecrypt                             KmsGrantOperation = "Decrypt"
	KmsGrantOperationDescribeKey                         KmsGrantOperation = "DescribeKey"
	KmsGrantOperationEncrypt                             KmsGrantOperation = "Encrypt"
	KmsGrantOperationGenerateDataKey                     KmsGrantOperation = "GenerateDataKey"
	KmsGrantOperationGenerateDataKeyPair                 KmsGrantOperation = "GenerateDataKeyPair"
	KmsGrantOperationGenerateDataKeyPairWithoutPlaintext KmsGrantOperation = "GenerateDataKeyPairWithoutPlaintext"
	KmsGrantOperationGenerateDataKeyWithoutPlaintext     KmsGrantOperation = "GenerateDataKeyWithoutPlaintext"
	KmsGrantOperationGetPublicKey                        KmsGrantOperation = "GetPublicKey"
	KmsGrantOperationReencryptFrom                       KmsGrantOperation = "ReEncryptFrom"
	KmsGrantOperationReencryptTo                         KmsGrantOperation = "ReEncryptTo"
	KmsGrantOperationRetireGrant                         KmsGrantOperation = "RetireGrant"
	KmsGrantOperationSign                                KmsGrantOperation = "Sign"
	KmsGrantOperationVerify                              KmsGrantOperation = "Verify"
)

// Values returns all known values for KmsGrantOperation. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (KmsGrantOperation) Values() []KmsGrantOperation {
	return []KmsGrantOperation{
		"CreateGrant",
		"Decrypt",
		"DescribeKey",
		"Encrypt",
		"GenerateDataKey",
		"GenerateDataKeyPair",
		"GenerateDataKeyPairWithoutPlaintext",
		"GenerateDataKeyWithoutPlaintext",
		"GetPublicKey",
		"ReEncryptFrom",
		"ReEncryptTo",
		"RetireGrant",
		"Sign",
		"Verify",
	}
}

type Locale string

// Enum values for Locale
const (
	LocaleDe   Locale = "DE"
	LocaleEn   Locale = "EN"
	LocaleEs   Locale = "ES"
	LocaleFr   Locale = "FR"
	LocaleIt   Locale = "IT"
	LocaleJa   Locale = "JA"
	LocaleKo   Locale = "KO"
	LocalePtBr Locale = "PT_BR"
	LocaleZhCn Locale = "ZH_CN"
	LocaleZhTw Locale = "ZH_TW"
)

// Values returns all known values for Locale. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Locale) Values() []Locale {
	return []Locale{
		"DE",
		"EN",
		"ES",
		"FR",
		"IT",
		"JA",
		"KO",
		"PT_BR",
		"ZH_CN",
		"ZH_TW",
	}
}

type OrderBy string

// Enum values for OrderBy
const (
	OrderByAsc  OrderBy = "ASC"
	OrderByDesc OrderBy = "DESC"
)

// Values returns all known values for OrderBy. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (OrderBy) Values() []OrderBy {
	return []OrderBy{
		"ASC",
		"DESC",
	}
}

type PolicyType string

// Enum values for PolicyType
const (
	PolicyTypeIdentityPolicy       PolicyType = "IDENTITY_POLICY"
	PolicyTypeResourcePolicy       PolicyType = "RESOURCE_POLICY"
	PolicyTypeServiceControlPolicy PolicyType = "SERVICE_CONTROL_POLICY"
)

// Values returns all known values for PolicyType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (PolicyType) Values() []PolicyType {
	return []PolicyType{
		"IDENTITY_POLICY",
		"RESOURCE_POLICY",
		"SERVICE_CONTROL_POLICY",
	}
}

type ReasonCode string

// Enum values for ReasonCode
const (
	ReasonCodeAwsServiceAccessDisabled           ReasonCode = "AWS_SERVICE_ACCESS_DISABLED"
	ReasonCodeDelegatedAdministratorDeregistered ReasonCode = "DELEGATED_ADMINISTRATOR_DEREGISTERED"
	ReasonCodeOrganizationDeleted                ReasonCode = "ORGANIZATION_DELETED"
	ReasonCodeServiceLinkedRoleCreationFailed    ReasonCode = "SERVICE_LINKED_ROLE_CREATION_FAILED"
)

// Values returns all known values for ReasonCode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ReasonCode) Values() []ReasonCode {
	return []ReasonCode{
		"AWS_SERVICE_ACCESS_DISABLED",
		"DELEGATED_ADMINISTRATOR_DEREGISTERED",
		"ORGANIZATION_DELETED",
		"SERVICE_LINKED_ROLE_CREATION_FAILED",
	}
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeAwsS3Bucket             ResourceType = "AWS::S3::Bucket"
	ResourceTypeAwsIamRole              ResourceType = "AWS::IAM::Role"
	ResourceTypeAwsSqsQueue             ResourceType = "AWS::SQS::Queue"
	ResourceTypeAwsLambdaFunction       ResourceType = "AWS::Lambda::Function"
	ResourceTypeAwsLambdaLayerversion   ResourceType = "AWS::Lambda::LayerVersion"
	ResourceTypeAwsKmsKey               ResourceType = "AWS::KMS::Key"
	ResourceTypeAwsSecretsmanagerSecret ResourceType = "AWS::SecretsManager::Secret"
	ResourceTypeAwsEfsFilesystem        ResourceType = "AWS::EFS::FileSystem"
	ResourceTypeAwsEc2Snapshot          ResourceType = "AWS::EC2::Snapshot"
	ResourceTypeAwsEcrRepository        ResourceType = "AWS::ECR::Repository"
	ResourceTypeAwsRdsDbsnapshot        ResourceType = "AWS::RDS::DBSnapshot"
	ResourceTypeAwsRdsDbclustersnapshot ResourceType = "AWS::RDS::DBClusterSnapshot"
	ResourceTypeAwsSnsTopic             ResourceType = "AWS::SNS::Topic"
)

// Values returns all known values for ResourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"AWS::S3::Bucket",
		"AWS::IAM::Role",
		"AWS::SQS::Queue",
		"AWS::Lambda::Function",
		"AWS::Lambda::LayerVersion",
		"AWS::KMS::Key",
		"AWS::SecretsManager::Secret",
		"AWS::EFS::FileSystem",
		"AWS::EC2::Snapshot",
		"AWS::ECR::Repository",
		"AWS::RDS::DBSnapshot",
		"AWS::RDS::DBClusterSnapshot",
		"AWS::SNS::Topic",
	}
}

type Type string

// Enum values for Type
const (
	TypeAccount      Type = "ACCOUNT"
	TypeOrganization Type = "ORGANIZATION"
)

// Values returns all known values for Type. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Type) Values() []Type {
	return []Type{
		"ACCOUNT",
		"ORGANIZATION",
	}
}

type ValidatePolicyFindingType string

// Enum values for ValidatePolicyFindingType
const (
	ValidatePolicyFindingTypeError           ValidatePolicyFindingType = "ERROR"
	ValidatePolicyFindingTypeSecurityWarning ValidatePolicyFindingType = "SECURITY_WARNING"
	ValidatePolicyFindingTypeSuggestion      ValidatePolicyFindingType = "SUGGESTION"
	ValidatePolicyFindingTypeWarning         ValidatePolicyFindingType = "WARNING"
)

// Values returns all known values for ValidatePolicyFindingType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidatePolicyFindingType) Values() []ValidatePolicyFindingType {
	return []ValidatePolicyFindingType{
		"ERROR",
		"SECURITY_WARNING",
		"SUGGESTION",
		"WARNING",
	}
}

type ValidatePolicyResourceType string

// Enum values for ValidatePolicyResourceType
const (
	ValidatePolicyResourceTypeS3Bucket                  ValidatePolicyResourceType = "AWS::S3::Bucket"
	ValidatePolicyResourceTypeS3AccessPoint             ValidatePolicyResourceType = "AWS::S3::AccessPoint"
	ValidatePolicyResourceTypeS3MultiRegionAccessPoint  ValidatePolicyResourceType = "AWS::S3::MultiRegionAccessPoint"
	ValidatePolicyResourceTypeS3ObjectLambdaAccessPoint ValidatePolicyResourceType = "AWS::S3ObjectLambda::AccessPoint"
	ValidatePolicyResourceTypeRoleTrust                 ValidatePolicyResourceType = "AWS::IAM::AssumeRolePolicyDocument"
)

// Values returns all known values for ValidatePolicyResourceType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidatePolicyResourceType) Values() []ValidatePolicyResourceType {
	return []ValidatePolicyResourceType{
		"AWS::S3::Bucket",
		"AWS::S3::AccessPoint",
		"AWS::S3::MultiRegionAccessPoint",
		"AWS::S3ObjectLambda::AccessPoint",
		"AWS::IAM::AssumeRolePolicyDocument",
	}
}

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonUnknownOperation      ValidationExceptionReason = "unknownOperation"
	ValidationExceptionReasonCannotParse           ValidationExceptionReason = "cannotParse"
	ValidationExceptionReasonFieldValidationFailed ValidationExceptionReason = "fieldValidationFailed"
	ValidationExceptionReasonOther                 ValidationExceptionReason = "other"
)

// Values returns all known values for ValidationExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidationExceptionReason) Values() []ValidationExceptionReason {
	return []ValidationExceptionReason{
		"unknownOperation",
		"cannotParse",
		"fieldValidationFailed",
		"other",
	}
}
