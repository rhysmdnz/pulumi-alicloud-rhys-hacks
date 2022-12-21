// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package edas

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GetApplicationsApplication struct {
	// The ID of the application that you want to deploy.
	AppId string `pulumi:"appId"`
	// The name of your EDAS application. Only letters '-' '_' and numbers are allowed. The length cannot exceed 36 characters.
	AppName string `pulumi:"appName"`
	// The type of the package for the deployment of the application that you want to create. The valid values are: WAR and JAR. We strongly recommend you to set this parameter when creating the application.
	ApplicationType string `pulumi:"applicationType"`
	// The package ID of Enterprise Distributed Application Service (EDAS) Container.
	BuildPackageId int `pulumi:"buildPackageId"`
	// The ID of the cluster that you want to create the application.
	ClusterId string `pulumi:"clusterId"`
	// The type of the cluster that you want to create. Valid values: 1: Swarm cluster. 2: ECS cluster. 3: Kubernates cluster.
	ClusterType int `pulumi:"clusterType"`
	// The ID of the namespace the application belongs to.
	RegionId string `pulumi:"regionId"`
}

// GetApplicationsApplicationInput is an input type that accepts GetApplicationsApplicationArgs and GetApplicationsApplicationOutput values.
// You can construct a concrete instance of `GetApplicationsApplicationInput` via:
//
//	GetApplicationsApplicationArgs{...}
type GetApplicationsApplicationInput interface {
	pulumi.Input

	ToGetApplicationsApplicationOutput() GetApplicationsApplicationOutput
	ToGetApplicationsApplicationOutputWithContext(context.Context) GetApplicationsApplicationOutput
}

type GetApplicationsApplicationArgs struct {
	// The ID of the application that you want to deploy.
	AppId pulumi.StringInput `pulumi:"appId"`
	// The name of your EDAS application. Only letters '-' '_' and numbers are allowed. The length cannot exceed 36 characters.
	AppName pulumi.StringInput `pulumi:"appName"`
	// The type of the package for the deployment of the application that you want to create. The valid values are: WAR and JAR. We strongly recommend you to set this parameter when creating the application.
	ApplicationType pulumi.StringInput `pulumi:"applicationType"`
	// The package ID of Enterprise Distributed Application Service (EDAS) Container.
	BuildPackageId pulumi.IntInput `pulumi:"buildPackageId"`
	// The ID of the cluster that you want to create the application.
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// The type of the cluster that you want to create. Valid values: 1: Swarm cluster. 2: ECS cluster. 3: Kubernates cluster.
	ClusterType pulumi.IntInput `pulumi:"clusterType"`
	// The ID of the namespace the application belongs to.
	RegionId pulumi.StringInput `pulumi:"regionId"`
}

func (GetApplicationsApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApplicationsApplication)(nil)).Elem()
}

func (i GetApplicationsApplicationArgs) ToGetApplicationsApplicationOutput() GetApplicationsApplicationOutput {
	return i.ToGetApplicationsApplicationOutputWithContext(context.Background())
}

func (i GetApplicationsApplicationArgs) ToGetApplicationsApplicationOutputWithContext(ctx context.Context) GetApplicationsApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetApplicationsApplicationOutput)
}

// GetApplicationsApplicationArrayInput is an input type that accepts GetApplicationsApplicationArray and GetApplicationsApplicationArrayOutput values.
// You can construct a concrete instance of `GetApplicationsApplicationArrayInput` via:
//
//	GetApplicationsApplicationArray{ GetApplicationsApplicationArgs{...} }
type GetApplicationsApplicationArrayInput interface {
	pulumi.Input

	ToGetApplicationsApplicationArrayOutput() GetApplicationsApplicationArrayOutput
	ToGetApplicationsApplicationArrayOutputWithContext(context.Context) GetApplicationsApplicationArrayOutput
}

type GetApplicationsApplicationArray []GetApplicationsApplicationInput

func (GetApplicationsApplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetApplicationsApplication)(nil)).Elem()
}

func (i GetApplicationsApplicationArray) ToGetApplicationsApplicationArrayOutput() GetApplicationsApplicationArrayOutput {
	return i.ToGetApplicationsApplicationArrayOutputWithContext(context.Background())
}

func (i GetApplicationsApplicationArray) ToGetApplicationsApplicationArrayOutputWithContext(ctx context.Context) GetApplicationsApplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetApplicationsApplicationArrayOutput)
}

type GetApplicationsApplicationOutput struct{ *pulumi.OutputState }

func (GetApplicationsApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApplicationsApplication)(nil)).Elem()
}

func (o GetApplicationsApplicationOutput) ToGetApplicationsApplicationOutput() GetApplicationsApplicationOutput {
	return o
}

func (o GetApplicationsApplicationOutput) ToGetApplicationsApplicationOutputWithContext(ctx context.Context) GetApplicationsApplicationOutput {
	return o
}

// The ID of the application that you want to deploy.
func (o GetApplicationsApplicationOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationsApplication) string { return v.AppId }).(pulumi.StringOutput)
}

// The name of your EDAS application. Only letters '-' '_' and numbers are allowed. The length cannot exceed 36 characters.
func (o GetApplicationsApplicationOutput) AppName() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationsApplication) string { return v.AppName }).(pulumi.StringOutput)
}

// The type of the package for the deployment of the application that you want to create. The valid values are: WAR and JAR. We strongly recommend you to set this parameter when creating the application.
func (o GetApplicationsApplicationOutput) ApplicationType() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationsApplication) string { return v.ApplicationType }).(pulumi.StringOutput)
}

// The package ID of Enterprise Distributed Application Service (EDAS) Container.
func (o GetApplicationsApplicationOutput) BuildPackageId() pulumi.IntOutput {
	return o.ApplyT(func(v GetApplicationsApplication) int { return v.BuildPackageId }).(pulumi.IntOutput)
}

// The ID of the cluster that you want to create the application.
func (o GetApplicationsApplicationOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationsApplication) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The type of the cluster that you want to create. Valid values: 1: Swarm cluster. 2: ECS cluster. 3: Kubernates cluster.
func (o GetApplicationsApplicationOutput) ClusterType() pulumi.IntOutput {
	return o.ApplyT(func(v GetApplicationsApplication) int { return v.ClusterType }).(pulumi.IntOutput)
}

// The ID of the namespace the application belongs to.
func (o GetApplicationsApplicationOutput) RegionId() pulumi.StringOutput {
	return o.ApplyT(func(v GetApplicationsApplication) string { return v.RegionId }).(pulumi.StringOutput)
}

type GetApplicationsApplicationArrayOutput struct{ *pulumi.OutputState }

func (GetApplicationsApplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetApplicationsApplication)(nil)).Elem()
}

func (o GetApplicationsApplicationArrayOutput) ToGetApplicationsApplicationArrayOutput() GetApplicationsApplicationArrayOutput {
	return o
}

func (o GetApplicationsApplicationArrayOutput) ToGetApplicationsApplicationArrayOutputWithContext(ctx context.Context) GetApplicationsApplicationArrayOutput {
	return o
}

func (o GetApplicationsApplicationArrayOutput) Index(i pulumi.IntInput) GetApplicationsApplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetApplicationsApplication {
		return vs[0].([]GetApplicationsApplication)[vs[1].(int)]
	}).(GetApplicationsApplicationOutput)
}

type GetClustersCluster struct {
	// The ID of the cluster that you want to create the application.
	ClusterId string `pulumi:"clusterId"`
	// The name of the cluster.
	ClusterName string `pulumi:"clusterName"`
	// The type of the cluster, Valid values: 1: Swarm cluster. 2: ECS cluster. 3: Kubernates cluster.
	ClusterType int `pulumi:"clusterType"`
	// The total number of CPUs in the cluster.
	Cpu int `pulumi:"cpu"`
	// The number of used CPUs in the cluster.
	CpuUsed int `pulumi:"cpuUsed"`
	// Cluster's creation time.
	CreateTime int `pulumi:"createTime"`
	// The total amount of memory in the cluser. Unit: MB.
	Mem int `pulumi:"mem"`
	// The amount of used memory in the cluser. Unit: MB.
	MemUsed int `pulumi:"memUsed"`
	// The network type of the cluster. Valid values: 1: classic network. 2: VPC.
	NetworkMode int `pulumi:"networkMode"`
	// The number of the Elastic Compute Service (ECS) instances that are deployed to the cluster.
	NodeNum int `pulumi:"nodeNum"`
	// The ID of the namespace the application belongs to.
	RegionId string `pulumi:"regionId"`
	// The time when the cluster was last updated.
	UpdateTime int `pulumi:"updateTime"`
	// The ID of the Virtual Private Cloud (VPC) for the cluster.
	VpcId string `pulumi:"vpcId"`
}

// GetClustersClusterInput is an input type that accepts GetClustersClusterArgs and GetClustersClusterOutput values.
// You can construct a concrete instance of `GetClustersClusterInput` via:
//
//	GetClustersClusterArgs{...}
type GetClustersClusterInput interface {
	pulumi.Input

	ToGetClustersClusterOutput() GetClustersClusterOutput
	ToGetClustersClusterOutputWithContext(context.Context) GetClustersClusterOutput
}

type GetClustersClusterArgs struct {
	// The ID of the cluster that you want to create the application.
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// The name of the cluster.
	ClusterName pulumi.StringInput `pulumi:"clusterName"`
	// The type of the cluster, Valid values: 1: Swarm cluster. 2: ECS cluster. 3: Kubernates cluster.
	ClusterType pulumi.IntInput `pulumi:"clusterType"`
	// The total number of CPUs in the cluster.
	Cpu pulumi.IntInput `pulumi:"cpu"`
	// The number of used CPUs in the cluster.
	CpuUsed pulumi.IntInput `pulumi:"cpuUsed"`
	// Cluster's creation time.
	CreateTime pulumi.IntInput `pulumi:"createTime"`
	// The total amount of memory in the cluser. Unit: MB.
	Mem pulumi.IntInput `pulumi:"mem"`
	// The amount of used memory in the cluser. Unit: MB.
	MemUsed pulumi.IntInput `pulumi:"memUsed"`
	// The network type of the cluster. Valid values: 1: classic network. 2: VPC.
	NetworkMode pulumi.IntInput `pulumi:"networkMode"`
	// The number of the Elastic Compute Service (ECS) instances that are deployed to the cluster.
	NodeNum pulumi.IntInput `pulumi:"nodeNum"`
	// The ID of the namespace the application belongs to.
	RegionId pulumi.StringInput `pulumi:"regionId"`
	// The time when the cluster was last updated.
	UpdateTime pulumi.IntInput `pulumi:"updateTime"`
	// The ID of the Virtual Private Cloud (VPC) for the cluster.
	VpcId pulumi.StringInput `pulumi:"vpcId"`
}

func (GetClustersClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClustersCluster)(nil)).Elem()
}

func (i GetClustersClusterArgs) ToGetClustersClusterOutput() GetClustersClusterOutput {
	return i.ToGetClustersClusterOutputWithContext(context.Background())
}

func (i GetClustersClusterArgs) ToGetClustersClusterOutputWithContext(ctx context.Context) GetClustersClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetClustersClusterOutput)
}

// GetClustersClusterArrayInput is an input type that accepts GetClustersClusterArray and GetClustersClusterArrayOutput values.
// You can construct a concrete instance of `GetClustersClusterArrayInput` via:
//
//	GetClustersClusterArray{ GetClustersClusterArgs{...} }
type GetClustersClusterArrayInput interface {
	pulumi.Input

	ToGetClustersClusterArrayOutput() GetClustersClusterArrayOutput
	ToGetClustersClusterArrayOutputWithContext(context.Context) GetClustersClusterArrayOutput
}

type GetClustersClusterArray []GetClustersClusterInput

func (GetClustersClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetClustersCluster)(nil)).Elem()
}

func (i GetClustersClusterArray) ToGetClustersClusterArrayOutput() GetClustersClusterArrayOutput {
	return i.ToGetClustersClusterArrayOutputWithContext(context.Background())
}

func (i GetClustersClusterArray) ToGetClustersClusterArrayOutputWithContext(ctx context.Context) GetClustersClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetClustersClusterArrayOutput)
}

type GetClustersClusterOutput struct{ *pulumi.OutputState }

func (GetClustersClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClustersCluster)(nil)).Elem()
}

func (o GetClustersClusterOutput) ToGetClustersClusterOutput() GetClustersClusterOutput {
	return o
}

func (o GetClustersClusterOutput) ToGetClustersClusterOutputWithContext(ctx context.Context) GetClustersClusterOutput {
	return o
}

// The ID of the cluster that you want to create the application.
func (o GetClustersClusterOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetClustersCluster) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The name of the cluster.
func (o GetClustersClusterOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v GetClustersCluster) string { return v.ClusterName }).(pulumi.StringOutput)
}

// The type of the cluster, Valid values: 1: Swarm cluster. 2: ECS cluster. 3: Kubernates cluster.
func (o GetClustersClusterOutput) ClusterType() pulumi.IntOutput {
	return o.ApplyT(func(v GetClustersCluster) int { return v.ClusterType }).(pulumi.IntOutput)
}

// The total number of CPUs in the cluster.
func (o GetClustersClusterOutput) Cpu() pulumi.IntOutput {
	return o.ApplyT(func(v GetClustersCluster) int { return v.Cpu }).(pulumi.IntOutput)
}

// The number of used CPUs in the cluster.
func (o GetClustersClusterOutput) CpuUsed() pulumi.IntOutput {
	return o.ApplyT(func(v GetClustersCluster) int { return v.CpuUsed }).(pulumi.IntOutput)
}

// Cluster's creation time.
func (o GetClustersClusterOutput) CreateTime() pulumi.IntOutput {
	return o.ApplyT(func(v GetClustersCluster) int { return v.CreateTime }).(pulumi.IntOutput)
}

// The total amount of memory in the cluser. Unit: MB.
func (o GetClustersClusterOutput) Mem() pulumi.IntOutput {
	return o.ApplyT(func(v GetClustersCluster) int { return v.Mem }).(pulumi.IntOutput)
}

// The amount of used memory in the cluser. Unit: MB.
func (o GetClustersClusterOutput) MemUsed() pulumi.IntOutput {
	return o.ApplyT(func(v GetClustersCluster) int { return v.MemUsed }).(pulumi.IntOutput)
}

// The network type of the cluster. Valid values: 1: classic network. 2: VPC.
func (o GetClustersClusterOutput) NetworkMode() pulumi.IntOutput {
	return o.ApplyT(func(v GetClustersCluster) int { return v.NetworkMode }).(pulumi.IntOutput)
}

// The number of the Elastic Compute Service (ECS) instances that are deployed to the cluster.
func (o GetClustersClusterOutput) NodeNum() pulumi.IntOutput {
	return o.ApplyT(func(v GetClustersCluster) int { return v.NodeNum }).(pulumi.IntOutput)
}

// The ID of the namespace the application belongs to.
func (o GetClustersClusterOutput) RegionId() pulumi.StringOutput {
	return o.ApplyT(func(v GetClustersCluster) string { return v.RegionId }).(pulumi.StringOutput)
}

// The time when the cluster was last updated.
func (o GetClustersClusterOutput) UpdateTime() pulumi.IntOutput {
	return o.ApplyT(func(v GetClustersCluster) int { return v.UpdateTime }).(pulumi.IntOutput)
}

// The ID of the Virtual Private Cloud (VPC) for the cluster.
func (o GetClustersClusterOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v GetClustersCluster) string { return v.VpcId }).(pulumi.StringOutput)
}

type GetClustersClusterArrayOutput struct{ *pulumi.OutputState }

func (GetClustersClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetClustersCluster)(nil)).Elem()
}

func (o GetClustersClusterArrayOutput) ToGetClustersClusterArrayOutput() GetClustersClusterArrayOutput {
	return o
}

func (o GetClustersClusterArrayOutput) ToGetClustersClusterArrayOutputWithContext(ctx context.Context) GetClustersClusterArrayOutput {
	return o
}

func (o GetClustersClusterArrayOutput) Index(i pulumi.IntInput) GetClustersClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetClustersCluster {
		return vs[0].([]GetClustersCluster)[vs[1].(int)]
	}).(GetClustersClusterOutput)
}

type GetDeployGroupsGroup struct {
	// ID of the EDAS application.
	AppId string `pulumi:"appId"`
	// The version of the deployment package for the application.
	AppVersionId string `pulumi:"appVersionId"`
	// The ID of the cluster that you want to create the application.
	ClusterId string `pulumi:"clusterId"`
	// The time when the instance group was created.
	CreateTime int `pulumi:"createTime"`
	// The ID of the instance group.
	GroupId string `pulumi:"groupId"`
	// The name of the instance group. The length cannot exceed 64 characters.
	GroupName string `pulumi:"groupName"`
	// The type of the instance group. Valid values: 0: Default group. 1: Phased release is disabled for traffic management. 2: Phased release is enabled for traffic management.
	GroupType int `pulumi:"groupType"`
	// The version of the deployment package for the instance group that was created.
	PackageVersionId string `pulumi:"packageVersionId"`
	// The time when the instance group was updated.
	UpdateTime int `pulumi:"updateTime"`
}

// GetDeployGroupsGroupInput is an input type that accepts GetDeployGroupsGroupArgs and GetDeployGroupsGroupOutput values.
// You can construct a concrete instance of `GetDeployGroupsGroupInput` via:
//
//	GetDeployGroupsGroupArgs{...}
type GetDeployGroupsGroupInput interface {
	pulumi.Input

	ToGetDeployGroupsGroupOutput() GetDeployGroupsGroupOutput
	ToGetDeployGroupsGroupOutputWithContext(context.Context) GetDeployGroupsGroupOutput
}

type GetDeployGroupsGroupArgs struct {
	// ID of the EDAS application.
	AppId pulumi.StringInput `pulumi:"appId"`
	// The version of the deployment package for the application.
	AppVersionId pulumi.StringInput `pulumi:"appVersionId"`
	// The ID of the cluster that you want to create the application.
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// The time when the instance group was created.
	CreateTime pulumi.IntInput `pulumi:"createTime"`
	// The ID of the instance group.
	GroupId pulumi.StringInput `pulumi:"groupId"`
	// The name of the instance group. The length cannot exceed 64 characters.
	GroupName pulumi.StringInput `pulumi:"groupName"`
	// The type of the instance group. Valid values: 0: Default group. 1: Phased release is disabled for traffic management. 2: Phased release is enabled for traffic management.
	GroupType pulumi.IntInput `pulumi:"groupType"`
	// The version of the deployment package for the instance group that was created.
	PackageVersionId pulumi.StringInput `pulumi:"packageVersionId"`
	// The time when the instance group was updated.
	UpdateTime pulumi.IntInput `pulumi:"updateTime"`
}

func (GetDeployGroupsGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDeployGroupsGroup)(nil)).Elem()
}

func (i GetDeployGroupsGroupArgs) ToGetDeployGroupsGroupOutput() GetDeployGroupsGroupOutput {
	return i.ToGetDeployGroupsGroupOutputWithContext(context.Background())
}

func (i GetDeployGroupsGroupArgs) ToGetDeployGroupsGroupOutputWithContext(ctx context.Context) GetDeployGroupsGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDeployGroupsGroupOutput)
}

// GetDeployGroupsGroupArrayInput is an input type that accepts GetDeployGroupsGroupArray and GetDeployGroupsGroupArrayOutput values.
// You can construct a concrete instance of `GetDeployGroupsGroupArrayInput` via:
//
//	GetDeployGroupsGroupArray{ GetDeployGroupsGroupArgs{...} }
type GetDeployGroupsGroupArrayInput interface {
	pulumi.Input

	ToGetDeployGroupsGroupArrayOutput() GetDeployGroupsGroupArrayOutput
	ToGetDeployGroupsGroupArrayOutputWithContext(context.Context) GetDeployGroupsGroupArrayOutput
}

type GetDeployGroupsGroupArray []GetDeployGroupsGroupInput

func (GetDeployGroupsGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDeployGroupsGroup)(nil)).Elem()
}

func (i GetDeployGroupsGroupArray) ToGetDeployGroupsGroupArrayOutput() GetDeployGroupsGroupArrayOutput {
	return i.ToGetDeployGroupsGroupArrayOutputWithContext(context.Background())
}

func (i GetDeployGroupsGroupArray) ToGetDeployGroupsGroupArrayOutputWithContext(ctx context.Context) GetDeployGroupsGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDeployGroupsGroupArrayOutput)
}

type GetDeployGroupsGroupOutput struct{ *pulumi.OutputState }

func (GetDeployGroupsGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDeployGroupsGroup)(nil)).Elem()
}

func (o GetDeployGroupsGroupOutput) ToGetDeployGroupsGroupOutput() GetDeployGroupsGroupOutput {
	return o
}

func (o GetDeployGroupsGroupOutput) ToGetDeployGroupsGroupOutputWithContext(ctx context.Context) GetDeployGroupsGroupOutput {
	return o
}

// ID of the EDAS application.
func (o GetDeployGroupsGroupOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeployGroupsGroup) string { return v.AppId }).(pulumi.StringOutput)
}

// The version of the deployment package for the application.
func (o GetDeployGroupsGroupOutput) AppVersionId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeployGroupsGroup) string { return v.AppVersionId }).(pulumi.StringOutput)
}

// The ID of the cluster that you want to create the application.
func (o GetDeployGroupsGroupOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeployGroupsGroup) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The time when the instance group was created.
func (o GetDeployGroupsGroupOutput) CreateTime() pulumi.IntOutput {
	return o.ApplyT(func(v GetDeployGroupsGroup) int { return v.CreateTime }).(pulumi.IntOutput)
}

// The ID of the instance group.
func (o GetDeployGroupsGroupOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeployGroupsGroup) string { return v.GroupId }).(pulumi.StringOutput)
}

// The name of the instance group. The length cannot exceed 64 characters.
func (o GetDeployGroupsGroupOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeployGroupsGroup) string { return v.GroupName }).(pulumi.StringOutput)
}

// The type of the instance group. Valid values: 0: Default group. 1: Phased release is disabled for traffic management. 2: Phased release is enabled for traffic management.
func (o GetDeployGroupsGroupOutput) GroupType() pulumi.IntOutput {
	return o.ApplyT(func(v GetDeployGroupsGroup) int { return v.GroupType }).(pulumi.IntOutput)
}

// The version of the deployment package for the instance group that was created.
func (o GetDeployGroupsGroupOutput) PackageVersionId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeployGroupsGroup) string { return v.PackageVersionId }).(pulumi.StringOutput)
}

// The time when the instance group was updated.
func (o GetDeployGroupsGroupOutput) UpdateTime() pulumi.IntOutput {
	return o.ApplyT(func(v GetDeployGroupsGroup) int { return v.UpdateTime }).(pulumi.IntOutput)
}

type GetDeployGroupsGroupArrayOutput struct{ *pulumi.OutputState }

func (GetDeployGroupsGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDeployGroupsGroup)(nil)).Elem()
}

func (o GetDeployGroupsGroupArrayOutput) ToGetDeployGroupsGroupArrayOutput() GetDeployGroupsGroupArrayOutput {
	return o
}

func (o GetDeployGroupsGroupArrayOutput) ToGetDeployGroupsGroupArrayOutputWithContext(ctx context.Context) GetDeployGroupsGroupArrayOutput {
	return o
}

func (o GetDeployGroupsGroupArrayOutput) Index(i pulumi.IntInput) GetDeployGroupsGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetDeployGroupsGroup {
		return vs[0].([]GetDeployGroupsGroup)[vs[1].(int)]
	}).(GetDeployGroupsGroupOutput)
}

type GetNamespacesNamespace struct {
	// The ID of the physical region to which the namespace belongs.
	BelongRegion string `pulumi:"belongRegion"`
	// Indicates whether remote debugging is allowed in this region.
	DebugEnable bool `pulumi:"debugEnable"`
	// The description of the namespace.
	Description string `pulumi:"description"`
	// The ID of the resource.
	Id string `pulumi:"id"`
	// The unique ID of the namespace generated by Enterprise Distributed Application Service (EDAS).
	NamespaceId string `pulumi:"namespaceId"`
	// The ID of the namespace. **Note:** The ID cannot be changed after the namespace is created. The ID is in the format of `Physical region ID:Logical region identifier`.
	NamespaceLogicalId string `pulumi:"namespaceLogicalId"`
	// The name of the namespace.
	NamespaceName string `pulumi:"namespaceName"`
	// The ID of the Alibaba Cloud account to which the namespace belongs.
	UserId string `pulumi:"userId"`
}

// GetNamespacesNamespaceInput is an input type that accepts GetNamespacesNamespaceArgs and GetNamespacesNamespaceOutput values.
// You can construct a concrete instance of `GetNamespacesNamespaceInput` via:
//
//	GetNamespacesNamespaceArgs{...}
type GetNamespacesNamespaceInput interface {
	pulumi.Input

	ToGetNamespacesNamespaceOutput() GetNamespacesNamespaceOutput
	ToGetNamespacesNamespaceOutputWithContext(context.Context) GetNamespacesNamespaceOutput
}

type GetNamespacesNamespaceArgs struct {
	// The ID of the physical region to which the namespace belongs.
	BelongRegion pulumi.StringInput `pulumi:"belongRegion"`
	// Indicates whether remote debugging is allowed in this region.
	DebugEnable pulumi.BoolInput `pulumi:"debugEnable"`
	// The description of the namespace.
	Description pulumi.StringInput `pulumi:"description"`
	// The ID of the resource.
	Id pulumi.StringInput `pulumi:"id"`
	// The unique ID of the namespace generated by Enterprise Distributed Application Service (EDAS).
	NamespaceId pulumi.StringInput `pulumi:"namespaceId"`
	// The ID of the namespace. **Note:** The ID cannot be changed after the namespace is created. The ID is in the format of `Physical region ID:Logical region identifier`.
	NamespaceLogicalId pulumi.StringInput `pulumi:"namespaceLogicalId"`
	// The name of the namespace.
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	// The ID of the Alibaba Cloud account to which the namespace belongs.
	UserId pulumi.StringInput `pulumi:"userId"`
}

func (GetNamespacesNamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNamespacesNamespace)(nil)).Elem()
}

func (i GetNamespacesNamespaceArgs) ToGetNamespacesNamespaceOutput() GetNamespacesNamespaceOutput {
	return i.ToGetNamespacesNamespaceOutputWithContext(context.Background())
}

func (i GetNamespacesNamespaceArgs) ToGetNamespacesNamespaceOutputWithContext(ctx context.Context) GetNamespacesNamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetNamespacesNamespaceOutput)
}

// GetNamespacesNamespaceArrayInput is an input type that accepts GetNamespacesNamespaceArray and GetNamespacesNamespaceArrayOutput values.
// You can construct a concrete instance of `GetNamespacesNamespaceArrayInput` via:
//
//	GetNamespacesNamespaceArray{ GetNamespacesNamespaceArgs{...} }
type GetNamespacesNamespaceArrayInput interface {
	pulumi.Input

	ToGetNamespacesNamespaceArrayOutput() GetNamespacesNamespaceArrayOutput
	ToGetNamespacesNamespaceArrayOutputWithContext(context.Context) GetNamespacesNamespaceArrayOutput
}

type GetNamespacesNamespaceArray []GetNamespacesNamespaceInput

func (GetNamespacesNamespaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetNamespacesNamespace)(nil)).Elem()
}

func (i GetNamespacesNamespaceArray) ToGetNamespacesNamespaceArrayOutput() GetNamespacesNamespaceArrayOutput {
	return i.ToGetNamespacesNamespaceArrayOutputWithContext(context.Background())
}

func (i GetNamespacesNamespaceArray) ToGetNamespacesNamespaceArrayOutputWithContext(ctx context.Context) GetNamespacesNamespaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetNamespacesNamespaceArrayOutput)
}

type GetNamespacesNamespaceOutput struct{ *pulumi.OutputState }

func (GetNamespacesNamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNamespacesNamespace)(nil)).Elem()
}

func (o GetNamespacesNamespaceOutput) ToGetNamespacesNamespaceOutput() GetNamespacesNamespaceOutput {
	return o
}

func (o GetNamespacesNamespaceOutput) ToGetNamespacesNamespaceOutputWithContext(ctx context.Context) GetNamespacesNamespaceOutput {
	return o
}

// The ID of the physical region to which the namespace belongs.
func (o GetNamespacesNamespaceOutput) BelongRegion() pulumi.StringOutput {
	return o.ApplyT(func(v GetNamespacesNamespace) string { return v.BelongRegion }).(pulumi.StringOutput)
}

// Indicates whether remote debugging is allowed in this region.
func (o GetNamespacesNamespaceOutput) DebugEnable() pulumi.BoolOutput {
	return o.ApplyT(func(v GetNamespacesNamespace) bool { return v.DebugEnable }).(pulumi.BoolOutput)
}

// The description of the namespace.
func (o GetNamespacesNamespaceOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetNamespacesNamespace) string { return v.Description }).(pulumi.StringOutput)
}

// The ID of the resource.
func (o GetNamespacesNamespaceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNamespacesNamespace) string { return v.Id }).(pulumi.StringOutput)
}

// The unique ID of the namespace generated by Enterprise Distributed Application Service (EDAS).
func (o GetNamespacesNamespaceOutput) NamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetNamespacesNamespace) string { return v.NamespaceId }).(pulumi.StringOutput)
}

// The ID of the namespace. **Note:** The ID cannot be changed after the namespace is created. The ID is in the format of `Physical region ID:Logical region identifier`.
func (o GetNamespacesNamespaceOutput) NamespaceLogicalId() pulumi.StringOutput {
	return o.ApplyT(func(v GetNamespacesNamespace) string { return v.NamespaceLogicalId }).(pulumi.StringOutput)
}

// The name of the namespace.
func (o GetNamespacesNamespaceOutput) NamespaceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetNamespacesNamespace) string { return v.NamespaceName }).(pulumi.StringOutput)
}

// The ID of the Alibaba Cloud account to which the namespace belongs.
func (o GetNamespacesNamespaceOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v GetNamespacesNamespace) string { return v.UserId }).(pulumi.StringOutput)
}

type GetNamespacesNamespaceArrayOutput struct{ *pulumi.OutputState }

func (GetNamespacesNamespaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetNamespacesNamespace)(nil)).Elem()
}

func (o GetNamespacesNamespaceArrayOutput) ToGetNamespacesNamespaceArrayOutput() GetNamespacesNamespaceArrayOutput {
	return o
}

func (o GetNamespacesNamespaceArrayOutput) ToGetNamespacesNamespaceArrayOutputWithContext(ctx context.Context) GetNamespacesNamespaceArrayOutput {
	return o
}

func (o GetNamespacesNamespaceArrayOutput) Index(i pulumi.IntInput) GetNamespacesNamespaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetNamespacesNamespace {
		return vs[0].([]GetNamespacesNamespace)[vs[1].(int)]
	}).(GetNamespacesNamespaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetApplicationsApplicationInput)(nil)).Elem(), GetApplicationsApplicationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetApplicationsApplicationArrayInput)(nil)).Elem(), GetApplicationsApplicationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetClustersClusterInput)(nil)).Elem(), GetClustersClusterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetClustersClusterArrayInput)(nil)).Elem(), GetClustersClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetDeployGroupsGroupInput)(nil)).Elem(), GetDeployGroupsGroupArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetDeployGroupsGroupArrayInput)(nil)).Elem(), GetDeployGroupsGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetNamespacesNamespaceInput)(nil)).Elem(), GetNamespacesNamespaceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetNamespacesNamespaceArrayInput)(nil)).Elem(), GetNamespacesNamespaceArray{})
	pulumi.RegisterOutputType(GetApplicationsApplicationOutput{})
	pulumi.RegisterOutputType(GetApplicationsApplicationArrayOutput{})
	pulumi.RegisterOutputType(GetClustersClusterOutput{})
	pulumi.RegisterOutputType(GetClustersClusterArrayOutput{})
	pulumi.RegisterOutputType(GetDeployGroupsGroupOutput{})
	pulumi.RegisterOutputType(GetDeployGroupsGroupArrayOutput{})
	pulumi.RegisterOutputType(GetNamespacesNamespaceOutput{})
	pulumi.RegisterOutputType(GetNamespacesNamespaceArrayOutput{})
}