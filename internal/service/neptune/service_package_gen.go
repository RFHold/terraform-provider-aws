// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package neptune

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	neptune_sdkv2 "github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceEngineVersion,
			TypeName: "aws_neptune_engine_version",
			Name:     "Engine Version",
		},
		{
			Factory:  dataSourceOrderableDBInstance,
			TypeName: "aws_neptune_orderable_db_instance",
			Name:     "Orderable DB Instance",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceCluster,
			TypeName: "aws_neptune_cluster",
			Name:     "Cluster",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceClusterEndpoint,
			TypeName: "aws_neptune_cluster_endpoint",
			Name:     "Cluster Endpoint",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceClusterInstance,
			TypeName: "aws_neptune_cluster_instance",
			Name:     "Cluster Instance",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceClusterParameterGroup,
			TypeName: "aws_neptune_cluster_parameter_group",
			Name:     "Cluster Parameter Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceClusterSnapshot,
			TypeName: "aws_neptune_cluster_snapshot",
			Name:     "Cluster Snapshot",
		},
		{
			Factory:  resourceEventSubscription,
			TypeName: "aws_neptune_event_subscription",
			Name:     "Event Subscription",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceGlobalCluster,
			TypeName: "aws_neptune_global_cluster",
			Name:     "Global Cluster",
		},
		{
			Factory:  resourceParameterGroup,
			TypeName: "aws_neptune_parameter_group",
			Name:     "Parameter Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  resourceSubnetGroup,
			TypeName: "aws_neptune_subnet_group",
			Name:     "Subnet Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Neptune
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*neptune_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return neptune_sdkv2.NewFromConfig(cfg,
		neptune_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
