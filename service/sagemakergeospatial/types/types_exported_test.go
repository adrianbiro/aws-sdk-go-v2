// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/sagemakergeospatial/types"
)

func ExampleAreaOfInterest_outputUsage() {
	var union types.AreaOfInterest
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.AreaOfInterestMemberAreaOfInterestGeometry:
		_ = v.Value // Value is types.AreaOfInterestGeometry

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ types.AreaOfInterestGeometry

func ExampleAreaOfInterestGeometry_outputUsage() {
	var union types.AreaOfInterestGeometry
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.AreaOfInterestGeometryMemberMultiPolygonGeometry:
		_ = v.Value // Value is types.MultiPolygonGeometryInput

	case *types.AreaOfInterestGeometryMemberPolygonGeometry:
		_ = v.Value // Value is types.PolygonGeometryInput

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.PolygonGeometryInput
var _ *types.MultiPolygonGeometryInput

func ExampleEojDataSourceConfigInput_outputUsage() {
	var union types.EojDataSourceConfigInput
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.EojDataSourceConfigInputMemberS3Data:
		_ = v.Value // Value is types.S3DataInput

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.S3DataInput

func ExampleJobConfigInput_outputUsage() {
	var union types.JobConfigInput
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.JobConfigInputMemberBandMathConfig:
		_ = v.Value // Value is types.BandMathConfigInput

	case *types.JobConfigInputMemberCloudMaskingConfig:
		_ = v.Value // Value is types.CloudMaskingConfigInput

	case *types.JobConfigInputMemberCloudRemovalConfig:
		_ = v.Value // Value is types.CloudRemovalConfigInput

	case *types.JobConfigInputMemberGeoMosaicConfig:
		_ = v.Value // Value is types.GeoMosaicConfigInput

	case *types.JobConfigInputMemberLandCoverSegmentationConfig:
		_ = v.Value // Value is types.LandCoverSegmentationConfigInput

	case *types.JobConfigInputMemberResamplingConfig:
		_ = v.Value // Value is types.ResamplingConfigInput

	case *types.JobConfigInputMemberStackConfig:
		_ = v.Value // Value is types.StackConfigInput

	case *types.JobConfigInputMemberTemporalStatisticsConfig:
		_ = v.Value // Value is types.TemporalStatisticsConfigInput

	case *types.JobConfigInputMemberZonalStatisticsConfig:
		_ = v.Value // Value is types.ZonalStatisticsConfigInput

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.CloudMaskingConfigInput
var _ *types.ZonalStatisticsConfigInput
var _ *types.GeoMosaicConfigInput
var _ *types.ResamplingConfigInput
var _ *types.BandMathConfigInput
var _ *types.LandCoverSegmentationConfigInput
var _ *types.TemporalStatisticsConfigInput
var _ *types.CloudRemovalConfigInput
var _ *types.StackConfigInput

func ExampleProperty_outputUsage() {
	var union types.Property
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.PropertyMemberEoCloudCover:
		_ = v.Value // Value is types.EoCloudCoverInput

	case *types.PropertyMemberLandsatCloudCoverLand:
		_ = v.Value // Value is types.LandsatCloudCoverLandInput

	case *types.PropertyMemberPlatform:
		_ = v.Value // Value is types.PlatformInput

	case *types.PropertyMemberViewOffNadir:
		_ = v.Value // Value is types.ViewOffNadirInput

	case *types.PropertyMemberViewSunAzimuth:
		_ = v.Value // Value is types.ViewSunAzimuthInput

	case *types.PropertyMemberViewSunElevation:
		_ = v.Value // Value is types.ViewSunElevationInput

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.LandsatCloudCoverLandInput
var _ *types.ViewSunElevationInput
var _ *types.ViewOffNadirInput
var _ *types.EoCloudCoverInput
var _ *types.PlatformInput
var _ *types.ViewSunAzimuthInput

func ExampleVectorEnrichmentJobConfig_outputUsage() {
	var union types.VectorEnrichmentJobConfig
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.VectorEnrichmentJobConfigMemberMapMatchingConfig:
		_ = v.Value // Value is types.MapMatchingConfig

	case *types.VectorEnrichmentJobConfigMemberReverseGeocodingConfig:
		_ = v.Value // Value is types.ReverseGeocodingConfig

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.MapMatchingConfig
var _ *types.ReverseGeocodingConfig

func ExampleVectorEnrichmentJobDataSourceConfigInput_outputUsage() {
	var union types.VectorEnrichmentJobDataSourceConfigInput
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.VectorEnrichmentJobDataSourceConfigInputMemberS3Data:
		_ = v.Value // Value is types.VectorEnrichmentJobS3Data

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.VectorEnrichmentJobS3Data
