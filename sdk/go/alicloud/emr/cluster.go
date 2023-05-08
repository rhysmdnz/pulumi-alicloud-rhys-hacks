// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package emr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ## Import
//
// Aliclioud E-MapReduce cluster can be imported using the id e.g.
//
// ```sh
//
//	$ pulumi import alicloud:emr/cluster:Cluster default C-B47FB8FE96C67XXXX
//
// ```
type Cluster struct {
	pulumi.CustomResourceState

	// Boot action parameters.
	BootstrapActions ClusterBootstrapActionArrayOutput `pulumi:"bootstrapActions"`
	// Charge Type for this group of hosts: PostPaid or PrePaid. If this is not specified, charge type will follow global chargeType value.
	ChargeType pulumi.StringPtrOutput `pulumi:"chargeType"`
	// EMR Cluster Type, e.g. HADOOP, KAFKA, DRUID, GATEWAY etc. You can find all valid EMR cluster type in emr web console. Supported 'GATEWAY' available in 1.61.0+.
	ClusterType pulumi.StringOutput `pulumi:"clusterType"`
	// The custom configurations of emr-cluster service.
	Configs ClusterConfigArrayOutput `pulumi:"configs"`
	// Cluster deposit type, HALF_MANAGED or FULL_MANAGED.
	DepositType pulumi.StringPtrOutput `pulumi:"depositType"`
	// High security cluster (true) or not. Default value is false.
	EasEnable pulumi.BoolPtrOutput `pulumi:"easEnable"`
	// EMR Version, e.g. EMR-3.22.0. You can find the all valid EMR Version in emr web console.
	EmrVer pulumi.StringOutput `pulumi:"emrVer"`
	// High Available for HDFS and YARN. If this is set true, MASTER group must have two nodes.
	HighAvailabilityEnable pulumi.BoolPtrOutput `pulumi:"highAvailabilityEnable"`
	// Groups of Host, You can specify MASTER as a group, CORE as a group (just like the above example).
	HostGroups ClusterHostGroupArrayOutput `pulumi:"hostGroups"`
	// Whether the MASTER node has a public IP address enabled. Default value is false.
	IsOpenPublicIp pulumi.BoolPtrOutput `pulumi:"isOpenPublicIp"`
	// Ssh key pair.
	KeyPairName pulumi.StringPtrOutput `pulumi:"keyPairName"`
	// Master ssh password.
	MasterPwd pulumi.StringPtrOutput `pulumi:"masterPwd"`
	// The configuration of emr-cluster service component metadata storage. If meta store type is ’user_rds’, this should be specified.
	MetaStoreConf ClusterMetaStoreConfPtrOutput `pulumi:"metaStoreConf"`
	// The type of emr-cluster service component metadata storage. ’dlf’ or ’local’ or ’user_rds’ .
	MetaStoreType pulumi.StringOutput `pulumi:"metaStoreType"`
	// The configurations of emr-cluster service modification after cluster created.
	ModifyClusterServiceConfig ClusterModifyClusterServiceConfigPtrOutput `pulumi:"modifyClusterServiceConfig"`
	// bootstrap action name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional software list.
	OptionSoftwareLists pulumi.StringArrayOutput `pulumi:"optionSoftwareLists"`
	// If charge type is PrePaid, this should be specified, unit is month. Supported value: 1、2、3、4、5、6、7、8、9、12、24、36.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// This specify the related cluster id, if this cluster is a Gateway.
	RelatedClusterId pulumi.StringPtrOutput `pulumi:"relatedClusterId"`
	// The Id of resource group which the emr-cluster belongs.
	ResourceGroupId pulumi.StringPtrOutput `pulumi:"resourceGroupId"`
	// Security Group ID for Cluster, you can also specify this key for each host group.
	SecurityGroupId pulumi.StringPtrOutput `pulumi:"securityGroupId"`
	// If this is set true, we can ssh into cluster. Default value is false.
	SshEnable pulumi.BoolPtrOutput `pulumi:"sshEnable"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Use local metadb. Default is false.
	UseLocalMetadb pulumi.BoolPtrOutput `pulumi:"useLocalMetadb"`
	// Alicloud EMR uses roles to perform actions on your behalf when provisioning cluster resources, running applications, dynamically scaling resources. EMR uses the following roles when interacting with other Alicloud services. Default value is AliyunEmrEcsDefaultRole.
	UserDefinedEmrEcsRole pulumi.StringPtrOutput `pulumi:"userDefinedEmrEcsRole"`
	// Global vswitch id, you can also specify it in host group.
	VswitchId pulumi.StringPtrOutput `pulumi:"vswitchId"`
	// Zone ID, e.g. cn-huhehaote-a
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterType == nil {
		return nil, errors.New("invalid value for required argument 'ClusterType'")
	}
	if args.EmrVer == nil {
		return nil, errors.New("invalid value for required argument 'EmrVer'")
	}
	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Cluster
	err := ctx.RegisterResource("alicloud:emr/cluster:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("alicloud:emr/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// Boot action parameters.
	BootstrapActions []ClusterBootstrapAction `pulumi:"bootstrapActions"`
	// Charge Type for this group of hosts: PostPaid or PrePaid. If this is not specified, charge type will follow global chargeType value.
	ChargeType *string `pulumi:"chargeType"`
	// EMR Cluster Type, e.g. HADOOP, KAFKA, DRUID, GATEWAY etc. You can find all valid EMR cluster type in emr web console. Supported 'GATEWAY' available in 1.61.0+.
	ClusterType *string `pulumi:"clusterType"`
	// The custom configurations of emr-cluster service.
	Configs []ClusterConfig `pulumi:"configs"`
	// Cluster deposit type, HALF_MANAGED or FULL_MANAGED.
	DepositType *string `pulumi:"depositType"`
	// High security cluster (true) or not. Default value is false.
	EasEnable *bool `pulumi:"easEnable"`
	// EMR Version, e.g. EMR-3.22.0. You can find the all valid EMR Version in emr web console.
	EmrVer *string `pulumi:"emrVer"`
	// High Available for HDFS and YARN. If this is set true, MASTER group must have two nodes.
	HighAvailabilityEnable *bool `pulumi:"highAvailabilityEnable"`
	// Groups of Host, You can specify MASTER as a group, CORE as a group (just like the above example).
	HostGroups []ClusterHostGroup `pulumi:"hostGroups"`
	// Whether the MASTER node has a public IP address enabled. Default value is false.
	IsOpenPublicIp *bool `pulumi:"isOpenPublicIp"`
	// Ssh key pair.
	KeyPairName *string `pulumi:"keyPairName"`
	// Master ssh password.
	MasterPwd *string `pulumi:"masterPwd"`
	// The configuration of emr-cluster service component metadata storage. If meta store type is ’user_rds’, this should be specified.
	MetaStoreConf *ClusterMetaStoreConf `pulumi:"metaStoreConf"`
	// The type of emr-cluster service component metadata storage. ’dlf’ or ’local’ or ’user_rds’ .
	MetaStoreType *string `pulumi:"metaStoreType"`
	// The configurations of emr-cluster service modification after cluster created.
	ModifyClusterServiceConfig *ClusterModifyClusterServiceConfig `pulumi:"modifyClusterServiceConfig"`
	// bootstrap action name.
	Name *string `pulumi:"name"`
	// Optional software list.
	OptionSoftwareLists []string `pulumi:"optionSoftwareLists"`
	// If charge type is PrePaid, this should be specified, unit is month. Supported value: 1、2、3、4、5、6、7、8、9、12、24、36.
	Period *int `pulumi:"period"`
	// This specify the related cluster id, if this cluster is a Gateway.
	RelatedClusterId *string `pulumi:"relatedClusterId"`
	// The Id of resource group which the emr-cluster belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// Security Group ID for Cluster, you can also specify this key for each host group.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// If this is set true, we can ssh into cluster. Default value is false.
	SshEnable *bool `pulumi:"sshEnable"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// Use local metadb. Default is false.
	UseLocalMetadb *bool `pulumi:"useLocalMetadb"`
	// Alicloud EMR uses roles to perform actions on your behalf when provisioning cluster resources, running applications, dynamically scaling resources. EMR uses the following roles when interacting with other Alicloud services. Default value is AliyunEmrEcsDefaultRole.
	UserDefinedEmrEcsRole *string `pulumi:"userDefinedEmrEcsRole"`
	// Global vswitch id, you can also specify it in host group.
	VswitchId *string `pulumi:"vswitchId"`
	// Zone ID, e.g. cn-huhehaote-a
	ZoneId *string `pulumi:"zoneId"`
}

type ClusterState struct {
	// Boot action parameters.
	BootstrapActions ClusterBootstrapActionArrayInput
	// Charge Type for this group of hosts: PostPaid or PrePaid. If this is not specified, charge type will follow global chargeType value.
	ChargeType pulumi.StringPtrInput
	// EMR Cluster Type, e.g. HADOOP, KAFKA, DRUID, GATEWAY etc. You can find all valid EMR cluster type in emr web console. Supported 'GATEWAY' available in 1.61.0+.
	ClusterType pulumi.StringPtrInput
	// The custom configurations of emr-cluster service.
	Configs ClusterConfigArrayInput
	// Cluster deposit type, HALF_MANAGED or FULL_MANAGED.
	DepositType pulumi.StringPtrInput
	// High security cluster (true) or not. Default value is false.
	EasEnable pulumi.BoolPtrInput
	// EMR Version, e.g. EMR-3.22.0. You can find the all valid EMR Version in emr web console.
	EmrVer pulumi.StringPtrInput
	// High Available for HDFS and YARN. If this is set true, MASTER group must have two nodes.
	HighAvailabilityEnable pulumi.BoolPtrInput
	// Groups of Host, You can specify MASTER as a group, CORE as a group (just like the above example).
	HostGroups ClusterHostGroupArrayInput
	// Whether the MASTER node has a public IP address enabled. Default value is false.
	IsOpenPublicIp pulumi.BoolPtrInput
	// Ssh key pair.
	KeyPairName pulumi.StringPtrInput
	// Master ssh password.
	MasterPwd pulumi.StringPtrInput
	// The configuration of emr-cluster service component metadata storage. If meta store type is ’user_rds’, this should be specified.
	MetaStoreConf ClusterMetaStoreConfPtrInput
	// The type of emr-cluster service component metadata storage. ’dlf’ or ’local’ or ’user_rds’ .
	MetaStoreType pulumi.StringPtrInput
	// The configurations of emr-cluster service modification after cluster created.
	ModifyClusterServiceConfig ClusterModifyClusterServiceConfigPtrInput
	// bootstrap action name.
	Name pulumi.StringPtrInput
	// Optional software list.
	OptionSoftwareLists pulumi.StringArrayInput
	// If charge type is PrePaid, this should be specified, unit is month. Supported value: 1、2、3、4、5、6、7、8、9、12、24、36.
	Period pulumi.IntPtrInput
	// This specify the related cluster id, if this cluster is a Gateway.
	RelatedClusterId pulumi.StringPtrInput
	// The Id of resource group which the emr-cluster belongs.
	ResourceGroupId pulumi.StringPtrInput
	// Security Group ID for Cluster, you can also specify this key for each host group.
	SecurityGroupId pulumi.StringPtrInput
	// If this is set true, we can ssh into cluster. Default value is false.
	SshEnable pulumi.BoolPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// Use local metadb. Default is false.
	UseLocalMetadb pulumi.BoolPtrInput
	// Alicloud EMR uses roles to perform actions on your behalf when provisioning cluster resources, running applications, dynamically scaling resources. EMR uses the following roles when interacting with other Alicloud services. Default value is AliyunEmrEcsDefaultRole.
	UserDefinedEmrEcsRole pulumi.StringPtrInput
	// Global vswitch id, you can also specify it in host group.
	VswitchId pulumi.StringPtrInput
	// Zone ID, e.g. cn-huhehaote-a
	ZoneId pulumi.StringPtrInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// Boot action parameters.
	BootstrapActions []ClusterBootstrapAction `pulumi:"bootstrapActions"`
	// Charge Type for this group of hosts: PostPaid or PrePaid. If this is not specified, charge type will follow global chargeType value.
	ChargeType *string `pulumi:"chargeType"`
	// EMR Cluster Type, e.g. HADOOP, KAFKA, DRUID, GATEWAY etc. You can find all valid EMR cluster type in emr web console. Supported 'GATEWAY' available in 1.61.0+.
	ClusterType string `pulumi:"clusterType"`
	// The custom configurations of emr-cluster service.
	Configs []ClusterConfig `pulumi:"configs"`
	// Cluster deposit type, HALF_MANAGED or FULL_MANAGED.
	DepositType *string `pulumi:"depositType"`
	// High security cluster (true) or not. Default value is false.
	EasEnable *bool `pulumi:"easEnable"`
	// EMR Version, e.g. EMR-3.22.0. You can find the all valid EMR Version in emr web console.
	EmrVer string `pulumi:"emrVer"`
	// High Available for HDFS and YARN. If this is set true, MASTER group must have two nodes.
	HighAvailabilityEnable *bool `pulumi:"highAvailabilityEnable"`
	// Groups of Host, You can specify MASTER as a group, CORE as a group (just like the above example).
	HostGroups []ClusterHostGroup `pulumi:"hostGroups"`
	// Whether the MASTER node has a public IP address enabled. Default value is false.
	IsOpenPublicIp *bool `pulumi:"isOpenPublicIp"`
	// Ssh key pair.
	KeyPairName *string `pulumi:"keyPairName"`
	// Master ssh password.
	MasterPwd *string `pulumi:"masterPwd"`
	// The configuration of emr-cluster service component metadata storage. If meta store type is ’user_rds’, this should be specified.
	MetaStoreConf *ClusterMetaStoreConf `pulumi:"metaStoreConf"`
	// The type of emr-cluster service component metadata storage. ’dlf’ or ’local’ or ’user_rds’ .
	MetaStoreType *string `pulumi:"metaStoreType"`
	// The configurations of emr-cluster service modification after cluster created.
	ModifyClusterServiceConfig *ClusterModifyClusterServiceConfig `pulumi:"modifyClusterServiceConfig"`
	// bootstrap action name.
	Name *string `pulumi:"name"`
	// Optional software list.
	OptionSoftwareLists []string `pulumi:"optionSoftwareLists"`
	// If charge type is PrePaid, this should be specified, unit is month. Supported value: 1、2、3、4、5、6、7、8、9、12、24、36.
	Period *int `pulumi:"period"`
	// This specify the related cluster id, if this cluster is a Gateway.
	RelatedClusterId *string `pulumi:"relatedClusterId"`
	// The Id of resource group which the emr-cluster belongs.
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// Security Group ID for Cluster, you can also specify this key for each host group.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// If this is set true, we can ssh into cluster. Default value is false.
	SshEnable *bool `pulumi:"sshEnable"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// Use local metadb. Default is false.
	UseLocalMetadb *bool `pulumi:"useLocalMetadb"`
	// Alicloud EMR uses roles to perform actions on your behalf when provisioning cluster resources, running applications, dynamically scaling resources. EMR uses the following roles when interacting with other Alicloud services. Default value is AliyunEmrEcsDefaultRole.
	UserDefinedEmrEcsRole *string `pulumi:"userDefinedEmrEcsRole"`
	// Global vswitch id, you can also specify it in host group.
	VswitchId *string `pulumi:"vswitchId"`
	// Zone ID, e.g. cn-huhehaote-a
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// Boot action parameters.
	BootstrapActions ClusterBootstrapActionArrayInput
	// Charge Type for this group of hosts: PostPaid or PrePaid. If this is not specified, charge type will follow global chargeType value.
	ChargeType pulumi.StringPtrInput
	// EMR Cluster Type, e.g. HADOOP, KAFKA, DRUID, GATEWAY etc. You can find all valid EMR cluster type in emr web console. Supported 'GATEWAY' available in 1.61.0+.
	ClusterType pulumi.StringInput
	// The custom configurations of emr-cluster service.
	Configs ClusterConfigArrayInput
	// Cluster deposit type, HALF_MANAGED or FULL_MANAGED.
	DepositType pulumi.StringPtrInput
	// High security cluster (true) or not. Default value is false.
	EasEnable pulumi.BoolPtrInput
	// EMR Version, e.g. EMR-3.22.0. You can find the all valid EMR Version in emr web console.
	EmrVer pulumi.StringInput
	// High Available for HDFS and YARN. If this is set true, MASTER group must have two nodes.
	HighAvailabilityEnable pulumi.BoolPtrInput
	// Groups of Host, You can specify MASTER as a group, CORE as a group (just like the above example).
	HostGroups ClusterHostGroupArrayInput
	// Whether the MASTER node has a public IP address enabled. Default value is false.
	IsOpenPublicIp pulumi.BoolPtrInput
	// Ssh key pair.
	KeyPairName pulumi.StringPtrInput
	// Master ssh password.
	MasterPwd pulumi.StringPtrInput
	// The configuration of emr-cluster service component metadata storage. If meta store type is ’user_rds’, this should be specified.
	MetaStoreConf ClusterMetaStoreConfPtrInput
	// The type of emr-cluster service component metadata storage. ’dlf’ or ’local’ or ’user_rds’ .
	MetaStoreType pulumi.StringPtrInput
	// The configurations of emr-cluster service modification after cluster created.
	ModifyClusterServiceConfig ClusterModifyClusterServiceConfigPtrInput
	// bootstrap action name.
	Name pulumi.StringPtrInput
	// Optional software list.
	OptionSoftwareLists pulumi.StringArrayInput
	// If charge type is PrePaid, this should be specified, unit is month. Supported value: 1、2、3、4、5、6、7、8、9、12、24、36.
	Period pulumi.IntPtrInput
	// This specify the related cluster id, if this cluster is a Gateway.
	RelatedClusterId pulumi.StringPtrInput
	// The Id of resource group which the emr-cluster belongs.
	ResourceGroupId pulumi.StringPtrInput
	// Security Group ID for Cluster, you can also specify this key for each host group.
	SecurityGroupId pulumi.StringPtrInput
	// If this is set true, we can ssh into cluster. Default value is false.
	SshEnable pulumi.BoolPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// Use local metadb. Default is false.
	UseLocalMetadb pulumi.BoolPtrInput
	// Alicloud EMR uses roles to perform actions on your behalf when provisioning cluster resources, running applications, dynamically scaling resources. EMR uses the following roles when interacting with other Alicloud services. Default value is AliyunEmrEcsDefaultRole.
	UserDefinedEmrEcsRole pulumi.StringPtrInput
	// Global vswitch id, you can also specify it in host group.
	VswitchId pulumi.StringPtrInput
	// Zone ID, e.g. cn-huhehaote-a
	ZoneId pulumi.StringInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

// ClusterArrayInput is an input type that accepts ClusterArray and ClusterArrayOutput values.
// You can construct a concrete instance of `ClusterArrayInput` via:
//
//	ClusterArray{ ClusterArgs{...} }
type ClusterArrayInput interface {
	pulumi.Input

	ToClusterArrayOutput() ClusterArrayOutput
	ToClusterArrayOutputWithContext(context.Context) ClusterArrayOutput
}

type ClusterArray []ClusterInput

func (ClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cluster)(nil)).Elem()
}

func (i ClusterArray) ToClusterArrayOutput() ClusterArrayOutput {
	return i.ToClusterArrayOutputWithContext(context.Background())
}

func (i ClusterArray) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterArrayOutput)
}

// ClusterMapInput is an input type that accepts ClusterMap and ClusterMapOutput values.
// You can construct a concrete instance of `ClusterMapInput` via:
//
//	ClusterMap{ "key": ClusterArgs{...} }
type ClusterMapInput interface {
	pulumi.Input

	ToClusterMapOutput() ClusterMapOutput
	ToClusterMapOutputWithContext(context.Context) ClusterMapOutput
}

type ClusterMap map[string]ClusterInput

func (ClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cluster)(nil)).Elem()
}

func (i ClusterMap) ToClusterMapOutput() ClusterMapOutput {
	return i.ToClusterMapOutputWithContext(context.Background())
}

func (i ClusterMap) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterMapOutput)
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

// Boot action parameters.
func (o ClusterOutput) BootstrapActions() ClusterBootstrapActionArrayOutput {
	return o.ApplyT(func(v *Cluster) ClusterBootstrapActionArrayOutput { return v.BootstrapActions }).(ClusterBootstrapActionArrayOutput)
}

// Charge Type for this group of hosts: PostPaid or PrePaid. If this is not specified, charge type will follow global chargeType value.
func (o ClusterOutput) ChargeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.ChargeType }).(pulumi.StringPtrOutput)
}

// EMR Cluster Type, e.g. HADOOP, KAFKA, DRUID, GATEWAY etc. You can find all valid EMR cluster type in emr web console. Supported 'GATEWAY' available in 1.61.0+.
func (o ClusterOutput) ClusterType() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ClusterType }).(pulumi.StringOutput)
}

// The custom configurations of emr-cluster service.
func (o ClusterOutput) Configs() ClusterConfigArrayOutput {
	return o.ApplyT(func(v *Cluster) ClusterConfigArrayOutput { return v.Configs }).(ClusterConfigArrayOutput)
}

// Cluster deposit type, HALF_MANAGED or FULL_MANAGED.
func (o ClusterOutput) DepositType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.DepositType }).(pulumi.StringPtrOutput)
}

// High security cluster (true) or not. Default value is false.
func (o ClusterOutput) EasEnable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolPtrOutput { return v.EasEnable }).(pulumi.BoolPtrOutput)
}

// EMR Version, e.g. EMR-3.22.0. You can find the all valid EMR Version in emr web console.
func (o ClusterOutput) EmrVer() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.EmrVer }).(pulumi.StringOutput)
}

// High Available for HDFS and YARN. If this is set true, MASTER group must have two nodes.
func (o ClusterOutput) HighAvailabilityEnable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolPtrOutput { return v.HighAvailabilityEnable }).(pulumi.BoolPtrOutput)
}

// Groups of Host, You can specify MASTER as a group, CORE as a group (just like the above example).
func (o ClusterOutput) HostGroups() ClusterHostGroupArrayOutput {
	return o.ApplyT(func(v *Cluster) ClusterHostGroupArrayOutput { return v.HostGroups }).(ClusterHostGroupArrayOutput)
}

// Whether the MASTER node has a public IP address enabled. Default value is false.
func (o ClusterOutput) IsOpenPublicIp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolPtrOutput { return v.IsOpenPublicIp }).(pulumi.BoolPtrOutput)
}

// Ssh key pair.
func (o ClusterOutput) KeyPairName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.KeyPairName }).(pulumi.StringPtrOutput)
}

// Master ssh password.
func (o ClusterOutput) MasterPwd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.MasterPwd }).(pulumi.StringPtrOutput)
}

// The configuration of emr-cluster service component metadata storage. If meta store type is ’user_rds’, this should be specified.
func (o ClusterOutput) MetaStoreConf() ClusterMetaStoreConfPtrOutput {
	return o.ApplyT(func(v *Cluster) ClusterMetaStoreConfPtrOutput { return v.MetaStoreConf }).(ClusterMetaStoreConfPtrOutput)
}

// The type of emr-cluster service component metadata storage. ’dlf’ or ’local’ or ’user_rds’ .
func (o ClusterOutput) MetaStoreType() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.MetaStoreType }).(pulumi.StringOutput)
}

// The configurations of emr-cluster service modification after cluster created.
func (o ClusterOutput) ModifyClusterServiceConfig() ClusterModifyClusterServiceConfigPtrOutput {
	return o.ApplyT(func(v *Cluster) ClusterModifyClusterServiceConfigPtrOutput { return v.ModifyClusterServiceConfig }).(ClusterModifyClusterServiceConfigPtrOutput)
}

// bootstrap action name.
func (o ClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional software list.
func (o ClusterOutput) OptionSoftwareLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringArrayOutput { return v.OptionSoftwareLists }).(pulumi.StringArrayOutput)
}

// If charge type is PrePaid, this should be specified, unit is month. Supported value: 1、2、3、4、5、6、7、8、9、12、24、36.
func (o ClusterOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

// This specify the related cluster id, if this cluster is a Gateway.
func (o ClusterOutput) RelatedClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.RelatedClusterId }).(pulumi.StringPtrOutput)
}

// The Id of resource group which the emr-cluster belongs.
func (o ClusterOutput) ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.ResourceGroupId }).(pulumi.StringPtrOutput)
}

// Security Group ID for Cluster, you can also specify this key for each host group.
func (o ClusterOutput) SecurityGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.SecurityGroupId }).(pulumi.StringPtrOutput)
}

// If this is set true, we can ssh into cluster. Default value is false.
func (o ClusterOutput) SshEnable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolPtrOutput { return v.SshEnable }).(pulumi.BoolPtrOutput)
}

// A mapping of tags to assign to the resource.
func (o ClusterOutput) Tags() pulumi.MapOutput {
	return o.ApplyT(func(v *Cluster) pulumi.MapOutput { return v.Tags }).(pulumi.MapOutput)
}

// Use local metadb. Default is false.
func (o ClusterOutput) UseLocalMetadb() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolPtrOutput { return v.UseLocalMetadb }).(pulumi.BoolPtrOutput)
}

// Alicloud EMR uses roles to perform actions on your behalf when provisioning cluster resources, running applications, dynamically scaling resources. EMR uses the following roles when interacting with other Alicloud services. Default value is AliyunEmrEcsDefaultRole.
func (o ClusterOutput) UserDefinedEmrEcsRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.UserDefinedEmrEcsRole }).(pulumi.StringPtrOutput)
}

// Global vswitch id, you can also specify it in host group.
func (o ClusterOutput) VswitchId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.VswitchId }).(pulumi.StringPtrOutput)
}

// Zone ID, e.g. cn-huhehaote-a
func (o ClusterOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

type ClusterArrayOutput struct{ *pulumi.OutputState }

func (ClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Cluster)(nil)).Elem()
}

func (o ClusterArrayOutput) ToClusterArrayOutput() ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) ToClusterArrayOutputWithContext(ctx context.Context) ClusterArrayOutput {
	return o
}

func (o ClusterArrayOutput) Index(i pulumi.IntInput) ClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Cluster {
		return vs[0].([]*Cluster)[vs[1].(int)]
	}).(ClusterOutput)
}

type ClusterMapOutput struct{ *pulumi.OutputState }

func (ClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Cluster)(nil)).Elem()
}

func (o ClusterMapOutput) ToClusterMapOutput() ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) ToClusterMapOutputWithContext(ctx context.Context) ClusterMapOutput {
	return o
}

func (o ClusterMapOutput) MapIndex(k pulumi.StringInput) ClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Cluster {
		return vs[0].(map[string]*Cluster)[vs[1].(string)]
	}).(ClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterInput)(nil)).Elem(), &Cluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterArrayInput)(nil)).Elem(), ClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterMapInput)(nil)).Elem(), ClusterMap{})
	pulumi.RegisterOutputType(ClusterOutput{})
	pulumi.RegisterOutputType(ClusterArrayOutput{})
	pulumi.RegisterOutputType(ClusterMapOutput{})
}
