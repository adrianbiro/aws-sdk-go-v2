// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// An LDAP attribute of an Active Directory computer account, in the form of a
// name:value pair.
type ActiveDirectoryComputerAttribute struct {

	// The name for the LDAP attribute.
	Name *string

	// The value for the LDAP attribute.
	Value *string
}

// The configuration for a Microsoft Active Directory (Microsoft AD) studio
// resource.
type ActiveDirectoryConfiguration struct {

	// A collection of custom attributes for an Active Directory computer.
	ComputerAttributes []ActiveDirectoryComputerAttribute

	// The directory ID of the AWS Directory Service for Microsoft AD to access using
	// this studio component.
	DirectoryId *string

	// The distinguished name (DN) and organizational unit (OU) of an Active Directory
	// computer.
	OrganizationalUnitDistinguishedName *string
}

// The configuration for a render farm that is associated with a studio resource.
type ComputeFarmConfiguration struct {

	// The name of an Active Directory user that is used on ComputeFarm worker
	// instances.
	ActiveDirectoryUser *string

	// The endpoint of the ComputeFarm that is accessed by the studio component
	// resource.
	Endpoint *string
}

// Represents a EULA resource.
type Eula struct {

	// The EULA content.
	Content *string

	// The Unix epoch timestamp in seconds for when the resource was created.
	CreatedAt *time.Time

	// The EULA ID.
	EulaId *string

	// The name for the EULA.
	Name *string

	// The Unix epoch timestamp in seconds for when the resource was updated.
	UpdatedAt *time.Time
}

type EulaAcceptance struct {

	// The Unix epoch timestamp in seconds for when the EULA was accepted.
	AcceptedAt *time.Time

	// The ID of the person who accepted the EULA.
	AcceptedBy *string

	// The ID of the acceptee.
	AccepteeId *string

	// The EULA acceptance ID.
	EulaAcceptanceId *string

	// The EULA ID.
	EulaId *string
}

type LaunchProfile struct {

	// The ARN of the resource.
	Arn *string

	// The Unix epoch timestamp in seconds for when the resource was created.
	CreatedAt *time.Time

	// The user ID of the user that created the launch profile.
	CreatedBy *string

	// A human-readable description of the launch profile.
	Description *string

	// Unique identifiers for a collection of EC2 subnets.
	Ec2SubnetIds []string

	// The launch profile ID.
	LaunchProfileId *string

	// The version number of the protocol that is used by the launch profile. The only
	// valid version is "2021-03-31".
	LaunchProfileProtocolVersions []string

	// A friendly name for the launch profile.
	Name *string

	// The current state.
	State LaunchProfileState

	// The status code.
	StatusCode LaunchProfileStatusCode

	// The status message for the launch profile.
	StatusMessage *string

	// A configuration for a streaming session.
	StreamConfiguration *StreamConfiguration

	// Unique identifiers for a collection of studio components that can be used with
	// this launch profile.
	StudioComponentIds []string

	// A collection of labels, in the form of key:value pairs, that apply to this
	// resource.
	Tags map[string]string

	// The Unix epoch timestamp in seconds for when the resource was updated.
	UpdatedAt *time.Time

	// The user ID of the user that most recently updated the resource.
	UpdatedBy *string
}

type LaunchProfileInitialization struct {

	// A LaunchProfileInitializationActiveDirectory resource.
	ActiveDirectory *LaunchProfileInitializationActiveDirectory

	// The EC2 security groups that control access to the studio component.
	Ec2SecurityGroupIds []string

	// The launch profile ID.
	LaunchProfileId *string

	// The version number of the protocol that is used by the launch profile. The only
	// valid version is "2021-03-31".
	LaunchProfileProtocolVersion *string

	// The launch purpose.
	LaunchPurpose *string

	// The name for the launch profile.
	Name *string

	// The platform of the launch platform, either WINDOWS or LINUX.
	Platform LaunchProfilePlatform

	// The system initializtion scripts.
	SystemInitializationScripts []LaunchProfileInitializationScript

	// The user initializtion scripts.
	UserInitializationScripts []LaunchProfileInitializationScript
}

type LaunchProfileInitializationActiveDirectory struct {

	// A collection of custom attributes for an Active Directory computer.
	ComputerAttributes []ActiveDirectoryComputerAttribute

	// The directory ID of the AWS Directory Service for Microsoft AD to access using
	// this launch profile.
	DirectoryId *string

	// The directory name.
	DirectoryName *string

	// The DNS IP address.
	DnsIpAddresses []string

	// The name for the organizational unit distinguished name.
	OrganizationalUnitDistinguishedName *string

	// The unique identifier for a studio component resource.
	StudioComponentId *string

	// The name for the studio component.
	StudioComponentName *string
}

type LaunchProfileInitializationScript struct {

	// The initialization script.
	Script *string

	// The unique identifier for a studio component resource.
	StudioComponentId *string

	// The name for the studio component.
	StudioComponentName *string
}

type LaunchProfileMembership struct {

	// The ID of the identity store.
	IdentityStoreId *string

	// The persona.
	Persona LaunchProfilePersona

	// The principal ID.
	PrincipalId *string
}

// The configuration for a license service that is associated with a studio
// resource.
type LicenseServiceConfiguration struct {

	// The endpoint of the license service that is accessed by the studio component
	// resource.
	Endpoint *string
}

type NewLaunchProfileMember struct {

	// The persona.
	//
	// This member is required.
	Persona LaunchProfilePersona

	// The principal ID.
	//
	// This member is required.
	PrincipalId *string
}

type NewStudioMember struct {

	// The persona.
	//
	// This member is required.
	Persona StudioPersona

	// The principal ID.
	//
	// This member is required.
	PrincipalId *string
}

// A parameter for a studio component script, in the form of a key:value pair.
type ScriptParameterKeyValue struct {

	// A script parameter key.
	Key *string

	// A script parameter value.
	Value *string
}

// The configuration for a shared file storage system that is associated with a
// studio resource.
type SharedFileSystemConfiguration struct {

	// The endpoint of the shared file system that is accessed by the studio component
	// resource.
	Endpoint *string

	// The unique identifier for a file system.
	FileSystemId *string

	// The mount location for a shared file system on a Linux virtual workstation.
	LinuxMountPoint *string

	// The name of the file share.
	ShareName *string

	// The mount location for a shared file system on a Windows virtual workstation.
	WindowsMountDrive *string
}

// A configuration for a streaming session.
type StreamConfiguration struct {

	// Enable or disable the use of the system clipboard to copy and paste between the
	// streaming session and streaming client.
	ClipboardMode StreamingClipboardMode

	// The EC2 instance types that users can select from when launching a streaming
	// session with this launch profile.
	Ec2InstanceTypes []StreamingInstanceType

	// The length of time, in minutes, that a streaming session can run. After this
	// point, Nimble Studio automatically terminates the session.
	MaxSessionLengthInMinutes int32

	// The streaming images that users can select from when launching a streaming
	// session with this launch profile.
	StreamingImageIds []string
}

type StreamConfigurationCreate struct {

	// Enable or disable the use of the system clipboard to copy and paste between the
	// streaming session and streaming client.
	//
	// This member is required.
	ClipboardMode StreamingClipboardMode

	// The EC2 instance types that users can select from when launching a streaming
	// session with this launch profile.
	//
	// This member is required.
	Ec2InstanceTypes []StreamingInstanceType

	// The streaming images that users can select from when launching a streaming
	// session with this launch profile.
	//
	// This member is required.
	StreamingImageIds []string

	// The length of time, in minutes, that a streaming session can run. After this
	// point, Nimble Studio automatically terminates the session.
	MaxSessionLengthInMinutes int32
}

type StreamingImage struct {

	// The ARN of the resource.
	Arn *string

	// A human-readable description of the streaming image.
	Description *string

	// The ID of an EC2 machine image with which to create the streaming image.
	Ec2ImageId *string

	// The encryption configuration.
	EncryptionConfiguration *StreamingImageEncryptionConfiguration

	// The list of EULAs that must be accepted before a Streaming Session can be
	// started using this streaming image.
	EulaIds []string

	// A friendly name for a streaming image resource.
	Name *string

	// The owner of the streaming image, either the studioId that contains the
	// streaming image, or 'amazon' for images that are provided by Amazon Nimble
	// Studio.
	Owner *string

	// The platform of the streaming image, either WINDOWS or LINUX.
	Platform *string

	// The current state.
	State StreamingImageState

	// The status code.
	StatusCode StreamingImageStatusCode

	// The status message for the streaming image.
	StatusMessage *string

	// The ID of the streaming image.
	StreamingImageId *string

	// A collection of labels, in the form of key:value pairs, that apply to this
	// resource.
	Tags map[string]string
}

type StreamingImageEncryptionConfiguration struct {

	// The type of KMS key that is used to encrypt studio data.
	//
	// This member is required.
	KeyType StreamingImageEncryptionConfigurationKeyType

	// The ARN for a KMS key that is used to encrypt studio data.
	KeyArn *string
}

type StreamingSession struct {

	// The ARN of the resource.
	Arn *string

	// The Unix epoch timestamp in seconds for when the resource was created.
	CreatedAt *time.Time

	// The user ID of the user that created the streaming session.
	CreatedBy *string

	// The EC2 Instance type used for the streaming session.
	Ec2InstanceType *string

	// The ID of the launch profile used to control access from the streaming session.
	LaunchProfileId *string

	// The session ID.
	SessionId *string

	// The current state.
	State StreamingSessionState

	// The status code.
	StatusCode StreamingSessionStatusCode

	// The status message for the streaming session.
	StatusMessage *string

	// The ID of the streaming image.
	StreamingImageId *string

	// A collection of labels, in the form of key:value pairs, that apply to this
	// resource.
	Tags map[string]string

	// The time the streaming session will automatically terminate if not terminated by
	// the user.
	TerminateAt *time.Time

	// The Unix epoch timestamp in seconds for when the resource was updated.
	UpdatedAt *time.Time

	// The user ID of the user that most recently updated the resource.
	UpdatedBy *string
}

type StreamingSessionStream struct {

	// The Unix epoch timestamp in seconds for when the resource was created.
	CreatedAt *time.Time

	// The user ID of the user that created the streaming session stream.
	CreatedBy *string

	// The Unix epoch timestamp in seconds for when the resource expires.
	ExpiresAt *time.Time

	// The current state.
	State StreamingSessionStreamState

	// The streaming session stream status code.
	StatusCode StreamingSessionStreamStatusCode

	// The stream ID.
	StreamId *string

	// The URL to connect to this stream using the DCV client.
	Url *string
}

type Studio struct {

	// The IAM role that studio admins assume when logging in to the Nimble Studio
	// portal.
	AdminRoleArn *string

	// The Amazon Resource Name (ARN) that is assigned to a studio resource and
	// uniquely identifies it. ARNs are unique across all Regions.
	Arn *string

	// The Unix epoch timestamp in seconds for when the resource was created.
	CreatedAt *time.Time

	// A friendly name for the studio.
	DisplayName *string

	// The AWS region where the studio resource is located.
	HomeRegion *string

	// The AWS SSO application client ID used to integrate with AWS SSO to enable AWS
	// SSO users to log in to Nimble portal.
	SsoClientId *string

	// The current state of the studio resource.
	State StudioState

	// Status codes that provide additional detail on the studio state.
	StatusCode StudioStatusCode

	// Additional detail on the studio state.
	StatusMessage *string

	// Configuration of the encryption method that is used for the studio.
	StudioEncryptionConfiguration *StudioEncryptionConfiguration

	// The unique identifier for a studio resource. In Nimble Studio, all other
	// resources are contained in a studio resource.
	StudioId *string

	// The name of the studio, as included in the URL when accessing it in the Nimble
	// Studio portal.
	StudioName *string

	// The address of the web page for the studio.
	StudioUrl *string

	// A collection of labels, in the form of key:value pairs, that apply to this
	// resource.
	Tags map[string]string

	// The Unix epoch timestamp in seconds for when the resource was updated.
	UpdatedAt *time.Time

	// The IAM role that studio users assume when logging in to the Nimble Studio
	// portal.
	UserRoleArn *string
}

// A network that is used by a studio’s users and workflows, including render farm,
// Active Directory, licensing, and file system.
type StudioComponent struct {

	// The ARN of the resource.
	Arn *string

	// The configuration of the studio component, based on component type.
	Configuration *StudioComponentConfiguration

	// The Unix epoch timestamp in seconds for when the resource was created.
	CreatedAt *time.Time

	// The user ID of the user that created the studio component.
	CreatedBy *string

	// A human-readable description for the studio component resource.
	Description *string

	// The EC2 security groups that control access to the studio component.
	Ec2SecurityGroupIds []string

	// Initialization scripts for studio components.
	InitializationScripts []StudioComponentInitializationScript

	// A friendly name for the studio component resource.
	Name *string

	// Parameters for the studio component scripts.
	ScriptParameters []ScriptParameterKeyValue

	// The current state.
	State StudioComponentState

	// The status code.
	StatusCode StudioComponentStatusCode

	// The status message for the studio component.
	StatusMessage *string

	// The unique identifier for a studio component resource.
	StudioComponentId *string

	// The specific subtype of a studio component.
	Subtype StudioComponentSubtype

	// A collection of labels, in the form of key:value pairs, that apply to this
	// resource.
	Tags map[string]string

	// The type of the studio component.
	Type StudioComponentType

	// The Unix epoch timestamp in seconds for when the resource was updated.
	UpdatedAt *time.Time

	// The user ID of the user that most recently updated the resource.
	UpdatedBy *string
}

// The configuration of the studio component, based on component type.
type StudioComponentConfiguration struct {

	// The configuration for a Microsoft Active Directory (Microsoft AD) studio
	// resource.
	ActiveDirectoryConfiguration *ActiveDirectoryConfiguration

	// The configuration for a render farm that is associated with a studio resource.
	ComputeFarmConfiguration *ComputeFarmConfiguration

	// The configuration for a license service that is associated with a studio
	// resource.
	LicenseServiceConfiguration *LicenseServiceConfiguration

	// The configuration for a shared file storage system that is associated with a
	// studio resource.
	SharedFileSystemConfiguration *SharedFileSystemConfiguration
}

// Initialization scripts for studio components.
type StudioComponentInitializationScript struct {

	// The version number of the protocol that is used by the launch profile. The only
	// valid version is "2021-03-31".
	LaunchProfileProtocolVersion *string

	// The platform of the initialization script, either WINDOWS or LINUX.
	Platform LaunchProfilePlatform

	// The method to use when running the initialization script.
	RunContext StudioComponentInitializationScriptRunContext

	// The initialization script.
	Script *string
}

//
type StudioComponentSummary struct {

	// The Unix epoch timestamp in seconds for when the resource was created.
	CreatedAt *time.Time

	// The user ID of the user that created the studio component.
	CreatedBy *string

	// The description.
	Description *string

	// The name for the studio component.
	Name *string

	// The unique identifier for a studio component resource.
	StudioComponentId *string

	// The specific subtype of a studio component.
	Subtype StudioComponentSubtype

	// The type of the studio component.
	Type StudioComponentType

	// The Unix epoch timestamp in seconds for when the resource was updated.
	UpdatedAt *time.Time

	// The user ID of the user that most recently updated the resource.
	UpdatedBy *string
}

// Configuration of the encryption method that is used for the studio.
type StudioEncryptionConfiguration struct {

	// The type of KMS key that is used to encrypt studio data.
	//
	// This member is required.
	KeyType StudioEncryptionConfigurationKeyType

	// The ARN for a KMS key that is used to encrypt studio data.
	KeyArn *string
}

type StudioMembership struct {

	// The ID of the identity store.
	IdentityStoreId *string

	// The persona.
	Persona StudioPersona

	// The principal ID.
	PrincipalId *string
}
