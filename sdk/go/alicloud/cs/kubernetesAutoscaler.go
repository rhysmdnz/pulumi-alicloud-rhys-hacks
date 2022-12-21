// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// cluster-autoscaler in Kubernetes Cluster.
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/go/alicloud/cs"
//	"github.com/pulumi/pulumi-alicloud/sdk/go/alicloud/ecs"
//	"github.com/pulumi/pulumi-alicloud/sdk/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/cs"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/ecs"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/ess"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			name := "autoscaler"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			defaultNetworks, err := vpc.GetNetworks(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultImages, err := ecs.GetImages(ctx, &ecs.GetImagesArgs{
//				Owners:     pulumi.StringRef("system"),
//				NameRegex:  pulumi.StringRef("^centos_7"),
//				MostRecent: pulumi.BoolRef(true),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultManagedKubernetesClusters, err := cs.GetManagedKubernetesClusters(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultInstanceTypes, err := ecs.GetInstanceTypes(ctx, &ecs.GetInstanceTypesArgs{
//				CpuCoreCount: pulumi.IntRef(2),
//				MemorySize:   pulumi.Float64Ref(4),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultSecurityGroup, err := ecs.NewSecurityGroup(ctx, "defaultSecurityGroup", &ecs.SecurityGroupArgs{
//				VpcId: pulumi.String(defaultNetworks.Vpcs[0].Id),
//			})
//			if err != nil {
//				return err
//			}
//			defaultScalingGroup, err := ess.NewScalingGroup(ctx, "defaultScalingGroup", &ess.ScalingGroupArgs{
//				ScalingGroupName: pulumi.String(name),
//				MinSize:          pulumi.Any(_var.Min_size),
//				MaxSize:          pulumi.Any(_var.Max_size),
//				VswitchIds: pulumi.StringArray{
//					pulumi.String(defaultNetworks.Vpcs[0].VswitchIds[0]),
//				},
//				RemovalPolicies: pulumi.StringArray{
//					pulumi.String("OldestInstance"),
//					pulumi.String("NewestInstance"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			defaultScalingConfiguration, err := ess.NewScalingConfiguration(ctx, "defaultScalingConfiguration", &ess.ScalingConfigurationArgs{
//				ImageId:            pulumi.String(defaultImages.Images[0].Id),
//				SecurityGroupId:    defaultSecurityGroup.ID(),
//				ScalingGroupId:     defaultScalingGroup.ID(),
//				InstanceType:       pulumi.String(defaultInstanceTypes.InstanceTypes[0].Id),
//				InternetChargeType: pulumi.String("PayByTraffic"),
//				ForceDelete:        pulumi.Bool(true),
//				Enable:             pulumi.Bool(true),
//				Active:             pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = cs.NewKubernetesAutoscaler(ctx, "defaultKubernetesAutoscaler", &cs.KubernetesAutoscalerArgs{
//				ClusterId: pulumi.String(defaultManagedKubernetesClusters.Clusters[0].Id),
//				Nodepools: cs.KubernetesAutoscalerNodepoolArray{
//					&cs.KubernetesAutoscalerNodepoolArgs{
//						Id:     defaultScalingGroup.ID(),
//						Labels: pulumi.String("a=b"),
//					},
//				},
//				Utilization:          pulumi.Any(_var.Utilization),
//				CoolDownDuration:     pulumi.Any(_var.Cool_down_duration),
//				DeferScaleInDuration: pulumi.Any(_var.Defer_scale_in_duration),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				alicloud_ess_scaling_group.Defalut,
//				defaultScalingConfiguration,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type KubernetesAutoscaler struct {
	pulumi.CustomResourceState

	// The id of kubernetes cluster.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The coolDownDuration option of cluster-autoscaler.
	CoolDownDuration pulumi.StringOutput `pulumi:"coolDownDuration"`
	// The deferScaleInDuration option of cluster-autoscaler.
	DeferScaleInDuration pulumi.StringOutput `pulumi:"deferScaleInDuration"`
	// * `nodepools.id` - (Required) The scaling group id of the groups configured for cluster-autoscaler.
	// * `nodepools.taints` - (Required) The taints for the nodes in scaling group.
	// * `nodepools.labels` - (Required) The labels for the nodes in scaling group.
	Nodepools KubernetesAutoscalerNodepoolArrayOutput `pulumi:"nodepools"`
	// Enable autoscaler access to alibabacloud service by ecs ramrole token. default: false
	UseEcsRamRoleToken pulumi.BoolPtrOutput `pulumi:"useEcsRamRoleToken"`
	// The utilization option of cluster-autoscaler.
	Utilization pulumi.StringOutput `pulumi:"utilization"`
}

// NewKubernetesAutoscaler registers a new resource with the given unique name, arguments, and options.
func NewKubernetesAutoscaler(ctx *pulumi.Context,
	name string, args *KubernetesAutoscalerArgs, opts ...pulumi.ResourceOption) (*KubernetesAutoscaler, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.CoolDownDuration == nil {
		return nil, errors.New("invalid value for required argument 'CoolDownDuration'")
	}
	if args.DeferScaleInDuration == nil {
		return nil, errors.New("invalid value for required argument 'DeferScaleInDuration'")
	}
	if args.Utilization == nil {
		return nil, errors.New("invalid value for required argument 'Utilization'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource KubernetesAutoscaler
	err := ctx.RegisterResource("alicloud:cs/kubernetesAutoscaler:KubernetesAutoscaler", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKubernetesAutoscaler gets an existing KubernetesAutoscaler resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKubernetesAutoscaler(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KubernetesAutoscalerState, opts ...pulumi.ResourceOption) (*KubernetesAutoscaler, error) {
	var resource KubernetesAutoscaler
	err := ctx.ReadResource("alicloud:cs/kubernetesAutoscaler:KubernetesAutoscaler", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KubernetesAutoscaler resources.
type kubernetesAutoscalerState struct {
	// The id of kubernetes cluster.
	ClusterId *string `pulumi:"clusterId"`
	// The coolDownDuration option of cluster-autoscaler.
	CoolDownDuration *string `pulumi:"coolDownDuration"`
	// The deferScaleInDuration option of cluster-autoscaler.
	DeferScaleInDuration *string `pulumi:"deferScaleInDuration"`
	// * `nodepools.id` - (Required) The scaling group id of the groups configured for cluster-autoscaler.
	// * `nodepools.taints` - (Required) The taints for the nodes in scaling group.
	// * `nodepools.labels` - (Required) The labels for the nodes in scaling group.
	Nodepools []KubernetesAutoscalerNodepool `pulumi:"nodepools"`
	// Enable autoscaler access to alibabacloud service by ecs ramrole token. default: false
	UseEcsRamRoleToken *bool `pulumi:"useEcsRamRoleToken"`
	// The utilization option of cluster-autoscaler.
	Utilization *string `pulumi:"utilization"`
}

type KubernetesAutoscalerState struct {
	// The id of kubernetes cluster.
	ClusterId pulumi.StringPtrInput
	// The coolDownDuration option of cluster-autoscaler.
	CoolDownDuration pulumi.StringPtrInput
	// The deferScaleInDuration option of cluster-autoscaler.
	DeferScaleInDuration pulumi.StringPtrInput
	// * `nodepools.id` - (Required) The scaling group id of the groups configured for cluster-autoscaler.
	// * `nodepools.taints` - (Required) The taints for the nodes in scaling group.
	// * `nodepools.labels` - (Required) The labels for the nodes in scaling group.
	Nodepools KubernetesAutoscalerNodepoolArrayInput
	// Enable autoscaler access to alibabacloud service by ecs ramrole token. default: false
	UseEcsRamRoleToken pulumi.BoolPtrInput
	// The utilization option of cluster-autoscaler.
	Utilization pulumi.StringPtrInput
}

func (KubernetesAutoscalerState) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesAutoscalerState)(nil)).Elem()
}

type kubernetesAutoscalerArgs struct {
	// The id of kubernetes cluster.
	ClusterId string `pulumi:"clusterId"`
	// The coolDownDuration option of cluster-autoscaler.
	CoolDownDuration string `pulumi:"coolDownDuration"`
	// The deferScaleInDuration option of cluster-autoscaler.
	DeferScaleInDuration string `pulumi:"deferScaleInDuration"`
	// * `nodepools.id` - (Required) The scaling group id of the groups configured for cluster-autoscaler.
	// * `nodepools.taints` - (Required) The taints for the nodes in scaling group.
	// * `nodepools.labels` - (Required) The labels for the nodes in scaling group.
	Nodepools []KubernetesAutoscalerNodepool `pulumi:"nodepools"`
	// Enable autoscaler access to alibabacloud service by ecs ramrole token. default: false
	UseEcsRamRoleToken *bool `pulumi:"useEcsRamRoleToken"`
	// The utilization option of cluster-autoscaler.
	Utilization string `pulumi:"utilization"`
}

// The set of arguments for constructing a KubernetesAutoscaler resource.
type KubernetesAutoscalerArgs struct {
	// The id of kubernetes cluster.
	ClusterId pulumi.StringInput
	// The coolDownDuration option of cluster-autoscaler.
	CoolDownDuration pulumi.StringInput
	// The deferScaleInDuration option of cluster-autoscaler.
	DeferScaleInDuration pulumi.StringInput
	// * `nodepools.id` - (Required) The scaling group id of the groups configured for cluster-autoscaler.
	// * `nodepools.taints` - (Required) The taints for the nodes in scaling group.
	// * `nodepools.labels` - (Required) The labels for the nodes in scaling group.
	Nodepools KubernetesAutoscalerNodepoolArrayInput
	// Enable autoscaler access to alibabacloud service by ecs ramrole token. default: false
	UseEcsRamRoleToken pulumi.BoolPtrInput
	// The utilization option of cluster-autoscaler.
	Utilization pulumi.StringInput
}

func (KubernetesAutoscalerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesAutoscalerArgs)(nil)).Elem()
}

type KubernetesAutoscalerInput interface {
	pulumi.Input

	ToKubernetesAutoscalerOutput() KubernetesAutoscalerOutput
	ToKubernetesAutoscalerOutputWithContext(ctx context.Context) KubernetesAutoscalerOutput
}

func (*KubernetesAutoscaler) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesAutoscaler)(nil)).Elem()
}

func (i *KubernetesAutoscaler) ToKubernetesAutoscalerOutput() KubernetesAutoscalerOutput {
	return i.ToKubernetesAutoscalerOutputWithContext(context.Background())
}

func (i *KubernetesAutoscaler) ToKubernetesAutoscalerOutputWithContext(ctx context.Context) KubernetesAutoscalerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesAutoscalerOutput)
}

// KubernetesAutoscalerArrayInput is an input type that accepts KubernetesAutoscalerArray and KubernetesAutoscalerArrayOutput values.
// You can construct a concrete instance of `KubernetesAutoscalerArrayInput` via:
//
//	KubernetesAutoscalerArray{ KubernetesAutoscalerArgs{...} }
type KubernetesAutoscalerArrayInput interface {
	pulumi.Input

	ToKubernetesAutoscalerArrayOutput() KubernetesAutoscalerArrayOutput
	ToKubernetesAutoscalerArrayOutputWithContext(context.Context) KubernetesAutoscalerArrayOutput
}

type KubernetesAutoscalerArray []KubernetesAutoscalerInput

func (KubernetesAutoscalerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubernetesAutoscaler)(nil)).Elem()
}

func (i KubernetesAutoscalerArray) ToKubernetesAutoscalerArrayOutput() KubernetesAutoscalerArrayOutput {
	return i.ToKubernetesAutoscalerArrayOutputWithContext(context.Background())
}

func (i KubernetesAutoscalerArray) ToKubernetesAutoscalerArrayOutputWithContext(ctx context.Context) KubernetesAutoscalerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesAutoscalerArrayOutput)
}

// KubernetesAutoscalerMapInput is an input type that accepts KubernetesAutoscalerMap and KubernetesAutoscalerMapOutput values.
// You can construct a concrete instance of `KubernetesAutoscalerMapInput` via:
//
//	KubernetesAutoscalerMap{ "key": KubernetesAutoscalerArgs{...} }
type KubernetesAutoscalerMapInput interface {
	pulumi.Input

	ToKubernetesAutoscalerMapOutput() KubernetesAutoscalerMapOutput
	ToKubernetesAutoscalerMapOutputWithContext(context.Context) KubernetesAutoscalerMapOutput
}

type KubernetesAutoscalerMap map[string]KubernetesAutoscalerInput

func (KubernetesAutoscalerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubernetesAutoscaler)(nil)).Elem()
}

func (i KubernetesAutoscalerMap) ToKubernetesAutoscalerMapOutput() KubernetesAutoscalerMapOutput {
	return i.ToKubernetesAutoscalerMapOutputWithContext(context.Background())
}

func (i KubernetesAutoscalerMap) ToKubernetesAutoscalerMapOutputWithContext(ctx context.Context) KubernetesAutoscalerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesAutoscalerMapOutput)
}

type KubernetesAutoscalerOutput struct{ *pulumi.OutputState }

func (KubernetesAutoscalerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubernetesAutoscaler)(nil)).Elem()
}

func (o KubernetesAutoscalerOutput) ToKubernetesAutoscalerOutput() KubernetesAutoscalerOutput {
	return o
}

func (o KubernetesAutoscalerOutput) ToKubernetesAutoscalerOutputWithContext(ctx context.Context) KubernetesAutoscalerOutput {
	return o
}

// The id of kubernetes cluster.
func (o KubernetesAutoscalerOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesAutoscaler) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The coolDownDuration option of cluster-autoscaler.
func (o KubernetesAutoscalerOutput) CoolDownDuration() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesAutoscaler) pulumi.StringOutput { return v.CoolDownDuration }).(pulumi.StringOutput)
}

// The deferScaleInDuration option of cluster-autoscaler.
func (o KubernetesAutoscalerOutput) DeferScaleInDuration() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesAutoscaler) pulumi.StringOutput { return v.DeferScaleInDuration }).(pulumi.StringOutput)
}

// * `nodepools.id` - (Required) The scaling group id of the groups configured for cluster-autoscaler.
// * `nodepools.taints` - (Required) The taints for the nodes in scaling group.
// * `nodepools.labels` - (Required) The labels for the nodes in scaling group.
func (o KubernetesAutoscalerOutput) Nodepools() KubernetesAutoscalerNodepoolArrayOutput {
	return o.ApplyT(func(v *KubernetesAutoscaler) KubernetesAutoscalerNodepoolArrayOutput { return v.Nodepools }).(KubernetesAutoscalerNodepoolArrayOutput)
}

// Enable autoscaler access to alibabacloud service by ecs ramrole token. default: false
func (o KubernetesAutoscalerOutput) UseEcsRamRoleToken() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KubernetesAutoscaler) pulumi.BoolPtrOutput { return v.UseEcsRamRoleToken }).(pulumi.BoolPtrOutput)
}

// The utilization option of cluster-autoscaler.
func (o KubernetesAutoscalerOutput) Utilization() pulumi.StringOutput {
	return o.ApplyT(func(v *KubernetesAutoscaler) pulumi.StringOutput { return v.Utilization }).(pulumi.StringOutput)
}

type KubernetesAutoscalerArrayOutput struct{ *pulumi.OutputState }

func (KubernetesAutoscalerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubernetesAutoscaler)(nil)).Elem()
}

func (o KubernetesAutoscalerArrayOutput) ToKubernetesAutoscalerArrayOutput() KubernetesAutoscalerArrayOutput {
	return o
}

func (o KubernetesAutoscalerArrayOutput) ToKubernetesAutoscalerArrayOutputWithContext(ctx context.Context) KubernetesAutoscalerArrayOutput {
	return o
}

func (o KubernetesAutoscalerArrayOutput) Index(i pulumi.IntInput) KubernetesAutoscalerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KubernetesAutoscaler {
		return vs[0].([]*KubernetesAutoscaler)[vs[1].(int)]
	}).(KubernetesAutoscalerOutput)
}

type KubernetesAutoscalerMapOutput struct{ *pulumi.OutputState }

func (KubernetesAutoscalerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubernetesAutoscaler)(nil)).Elem()
}

func (o KubernetesAutoscalerMapOutput) ToKubernetesAutoscalerMapOutput() KubernetesAutoscalerMapOutput {
	return o
}

func (o KubernetesAutoscalerMapOutput) ToKubernetesAutoscalerMapOutputWithContext(ctx context.Context) KubernetesAutoscalerMapOutput {
	return o
}

func (o KubernetesAutoscalerMapOutput) MapIndex(k pulumi.StringInput) KubernetesAutoscalerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KubernetesAutoscaler {
		return vs[0].(map[string]*KubernetesAutoscaler)[vs[1].(string)]
	}).(KubernetesAutoscalerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesAutoscalerInput)(nil)).Elem(), &KubernetesAutoscaler{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesAutoscalerArrayInput)(nil)).Elem(), KubernetesAutoscalerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesAutoscalerMapInput)(nil)).Elem(), KubernetesAutoscalerMap{})
	pulumi.RegisterOutputType(KubernetesAutoscalerOutput{})
	pulumi.RegisterOutputType(KubernetesAutoscalerArrayOutput{})
	pulumi.RegisterOutputType(KubernetesAutoscalerMapOutput{})
}