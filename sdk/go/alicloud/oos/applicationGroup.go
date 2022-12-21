// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oos

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a OOS Application Group resource.
//
// For information about OOS Application Group and how to use it, see [What is Application Group](https://www.alibabacloud.com/help/en/doc-detail/120556.html).
//
// > **NOTE:** Available in v1.146.0+.
//
// ## Example Usage
//
// # Basic Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/go/alicloud/resourcemanager"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/oos"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/resourcemanager"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultResourceGroups, err := resourcemanager.GetResourceGroups(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			defaultApplication, err := oos.NewApplication(ctx, "defaultApplication", &oos.ApplicationArgs{
//				ResourceGroupId: pulumi.String(defaultResourceGroups.Groups[0].Id),
//				ApplicationName: pulumi.String("example_value"),
//				Description:     pulumi.String("example_value"),
//				Tags: pulumi.AnyMap{
//					"Created": pulumi.Any("TF"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = oos.NewApplicationGroup(ctx, "defaultApplicationGroup", &oos.ApplicationGroupArgs{
//				ApplicationGroupName: pulumi.Any(_var.Name),
//				ApplicationName:      defaultApplication.ID(),
//				DeployRegionId:       pulumi.String("example_value"),
//				Description:          pulumi.String("example_value"),
//				ImportTagKey:         pulumi.String("example_value"),
//				ImportTagValue:       pulumi.String("example_value"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// OOS Application Group can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:oos/applicationGroup:ApplicationGroup example <application_name>:<application_group_name>
//
// ```
type ApplicationGroup struct {
	pulumi.CustomResourceState

	// The name of the Application group.
	ApplicationGroupName pulumi.StringOutput `pulumi:"applicationGroupName"`
	// The name of the Application.
	ApplicationName pulumi.StringOutput `pulumi:"applicationName"`
	// The region ID of the deployment.
	DeployRegionId pulumi.StringOutput `pulumi:"deployRegionId"`
	// Application group description information.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The tag key must be passed in at the same time as the tag value (import_tag_value) or none, not just one. If both `importTagKey` and `importTagValue` are left empty, the default is app-{ApplicationName} (application name).
	ImportTagKey pulumi.StringOutput `pulumi:"importTagKey"`
	// The tag value must be passed in at the same time as the tag key (import_tag_key) or none, not just one. If both `importTagKey` and `importTagValue` are left empty, the default is application group name.
	// .
	ImportTagValue pulumi.StringOutput `pulumi:"importTagValue"`
}

// NewApplicationGroup registers a new resource with the given unique name, arguments, and options.
func NewApplicationGroup(ctx *pulumi.Context,
	name string, args *ApplicationGroupArgs, opts ...pulumi.ResourceOption) (*ApplicationGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationGroupName'")
	}
	if args.ApplicationName == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationName'")
	}
	if args.DeployRegionId == nil {
		return nil, errors.New("invalid value for required argument 'DeployRegionId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ApplicationGroup
	err := ctx.RegisterResource("alicloud:oos/applicationGroup:ApplicationGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationGroup gets an existing ApplicationGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationGroupState, opts ...pulumi.ResourceOption) (*ApplicationGroup, error) {
	var resource ApplicationGroup
	err := ctx.ReadResource("alicloud:oos/applicationGroup:ApplicationGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationGroup resources.
type applicationGroupState struct {
	// The name of the Application group.
	ApplicationGroupName *string `pulumi:"applicationGroupName"`
	// The name of the Application.
	ApplicationName *string `pulumi:"applicationName"`
	// The region ID of the deployment.
	DeployRegionId *string `pulumi:"deployRegionId"`
	// Application group description information.
	Description *string `pulumi:"description"`
	// The tag key must be passed in at the same time as the tag value (import_tag_value) or none, not just one. If both `importTagKey` and `importTagValue` are left empty, the default is app-{ApplicationName} (application name).
	ImportTagKey *string `pulumi:"importTagKey"`
	// The tag value must be passed in at the same time as the tag key (import_tag_key) or none, not just one. If both `importTagKey` and `importTagValue` are left empty, the default is application group name.
	// .
	ImportTagValue *string `pulumi:"importTagValue"`
}

type ApplicationGroupState struct {
	// The name of the Application group.
	ApplicationGroupName pulumi.StringPtrInput
	// The name of the Application.
	ApplicationName pulumi.StringPtrInput
	// The region ID of the deployment.
	DeployRegionId pulumi.StringPtrInput
	// Application group description information.
	Description pulumi.StringPtrInput
	// The tag key must be passed in at the same time as the tag value (import_tag_value) or none, not just one. If both `importTagKey` and `importTagValue` are left empty, the default is app-{ApplicationName} (application name).
	ImportTagKey pulumi.StringPtrInput
	// The tag value must be passed in at the same time as the tag key (import_tag_key) or none, not just one. If both `importTagKey` and `importTagValue` are left empty, the default is application group name.
	// .
	ImportTagValue pulumi.StringPtrInput
}

func (ApplicationGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationGroupState)(nil)).Elem()
}

type applicationGroupArgs struct {
	// The name of the Application group.
	ApplicationGroupName string `pulumi:"applicationGroupName"`
	// The name of the Application.
	ApplicationName string `pulumi:"applicationName"`
	// The region ID of the deployment.
	DeployRegionId string `pulumi:"deployRegionId"`
	// Application group description information.
	Description *string `pulumi:"description"`
	// The tag key must be passed in at the same time as the tag value (import_tag_value) or none, not just one. If both `importTagKey` and `importTagValue` are left empty, the default is app-{ApplicationName} (application name).
	ImportTagKey *string `pulumi:"importTagKey"`
	// The tag value must be passed in at the same time as the tag key (import_tag_key) or none, not just one. If both `importTagKey` and `importTagValue` are left empty, the default is application group name.
	// .
	ImportTagValue *string `pulumi:"importTagValue"`
}

// The set of arguments for constructing a ApplicationGroup resource.
type ApplicationGroupArgs struct {
	// The name of the Application group.
	ApplicationGroupName pulumi.StringInput
	// The name of the Application.
	ApplicationName pulumi.StringInput
	// The region ID of the deployment.
	DeployRegionId pulumi.StringInput
	// Application group description information.
	Description pulumi.StringPtrInput
	// The tag key must be passed in at the same time as the tag value (import_tag_value) or none, not just one. If both `importTagKey` and `importTagValue` are left empty, the default is app-{ApplicationName} (application name).
	ImportTagKey pulumi.StringPtrInput
	// The tag value must be passed in at the same time as the tag key (import_tag_key) or none, not just one. If both `importTagKey` and `importTagValue` are left empty, the default is application group name.
	// .
	ImportTagValue pulumi.StringPtrInput
}

func (ApplicationGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationGroupArgs)(nil)).Elem()
}

type ApplicationGroupInput interface {
	pulumi.Input

	ToApplicationGroupOutput() ApplicationGroupOutput
	ToApplicationGroupOutputWithContext(ctx context.Context) ApplicationGroupOutput
}

func (*ApplicationGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGroup)(nil)).Elem()
}

func (i *ApplicationGroup) ToApplicationGroupOutput() ApplicationGroupOutput {
	return i.ToApplicationGroupOutputWithContext(context.Background())
}

func (i *ApplicationGroup) ToApplicationGroupOutputWithContext(ctx context.Context) ApplicationGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGroupOutput)
}

// ApplicationGroupArrayInput is an input type that accepts ApplicationGroupArray and ApplicationGroupArrayOutput values.
// You can construct a concrete instance of `ApplicationGroupArrayInput` via:
//
//	ApplicationGroupArray{ ApplicationGroupArgs{...} }
type ApplicationGroupArrayInput interface {
	pulumi.Input

	ToApplicationGroupArrayOutput() ApplicationGroupArrayOutput
	ToApplicationGroupArrayOutputWithContext(context.Context) ApplicationGroupArrayOutput
}

type ApplicationGroupArray []ApplicationGroupInput

func (ApplicationGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationGroup)(nil)).Elem()
}

func (i ApplicationGroupArray) ToApplicationGroupArrayOutput() ApplicationGroupArrayOutput {
	return i.ToApplicationGroupArrayOutputWithContext(context.Background())
}

func (i ApplicationGroupArray) ToApplicationGroupArrayOutputWithContext(ctx context.Context) ApplicationGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGroupArrayOutput)
}

// ApplicationGroupMapInput is an input type that accepts ApplicationGroupMap and ApplicationGroupMapOutput values.
// You can construct a concrete instance of `ApplicationGroupMapInput` via:
//
//	ApplicationGroupMap{ "key": ApplicationGroupArgs{...} }
type ApplicationGroupMapInput interface {
	pulumi.Input

	ToApplicationGroupMapOutput() ApplicationGroupMapOutput
	ToApplicationGroupMapOutputWithContext(context.Context) ApplicationGroupMapOutput
}

type ApplicationGroupMap map[string]ApplicationGroupInput

func (ApplicationGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationGroup)(nil)).Elem()
}

func (i ApplicationGroupMap) ToApplicationGroupMapOutput() ApplicationGroupMapOutput {
	return i.ToApplicationGroupMapOutputWithContext(context.Background())
}

func (i ApplicationGroupMap) ToApplicationGroupMapOutputWithContext(ctx context.Context) ApplicationGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGroupMapOutput)
}

type ApplicationGroupOutput struct{ *pulumi.OutputState }

func (ApplicationGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGroup)(nil)).Elem()
}

func (o ApplicationGroupOutput) ToApplicationGroupOutput() ApplicationGroupOutput {
	return o
}

func (o ApplicationGroupOutput) ToApplicationGroupOutputWithContext(ctx context.Context) ApplicationGroupOutput {
	return o
}

// The name of the Application group.
func (o ApplicationGroupOutput) ApplicationGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.ApplicationGroupName }).(pulumi.StringOutput)
}

// The name of the Application.
func (o ApplicationGroupOutput) ApplicationName() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.ApplicationName }).(pulumi.StringOutput)
}

// The region ID of the deployment.
func (o ApplicationGroupOutput) DeployRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.DeployRegionId }).(pulumi.StringOutput)
}

// Application group description information.
func (o ApplicationGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The tag key must be passed in at the same time as the tag value (import_tag_value) or none, not just one. If both `importTagKey` and `importTagValue` are left empty, the default is app-{ApplicationName} (application name).
func (o ApplicationGroupOutput) ImportTagKey() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.ImportTagKey }).(pulumi.StringOutput)
}

// The tag value must be passed in at the same time as the tag key (import_tag_key) or none, not just one. If both `importTagKey` and `importTagValue` are left empty, the default is application group name.
// .
func (o ApplicationGroupOutput) ImportTagValue() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGroup) pulumi.StringOutput { return v.ImportTagValue }).(pulumi.StringOutput)
}

type ApplicationGroupArrayOutput struct{ *pulumi.OutputState }

func (ApplicationGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApplicationGroup)(nil)).Elem()
}

func (o ApplicationGroupArrayOutput) ToApplicationGroupArrayOutput() ApplicationGroupArrayOutput {
	return o
}

func (o ApplicationGroupArrayOutput) ToApplicationGroupArrayOutputWithContext(ctx context.Context) ApplicationGroupArrayOutput {
	return o
}

func (o ApplicationGroupArrayOutput) Index(i pulumi.IntInput) ApplicationGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApplicationGroup {
		return vs[0].([]*ApplicationGroup)[vs[1].(int)]
	}).(ApplicationGroupOutput)
}

type ApplicationGroupMapOutput struct{ *pulumi.OutputState }

func (ApplicationGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApplicationGroup)(nil)).Elem()
}

func (o ApplicationGroupMapOutput) ToApplicationGroupMapOutput() ApplicationGroupMapOutput {
	return o
}

func (o ApplicationGroupMapOutput) ToApplicationGroupMapOutputWithContext(ctx context.Context) ApplicationGroupMapOutput {
	return o
}

func (o ApplicationGroupMapOutput) MapIndex(k pulumi.StringInput) ApplicationGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApplicationGroup {
		return vs[0].(map[string]*ApplicationGroup)[vs[1].(string)]
	}).(ApplicationGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGroupInput)(nil)).Elem(), &ApplicationGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGroupArrayInput)(nil)).Elem(), ApplicationGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationGroupMapInput)(nil)).Elem(), ApplicationGroupMap{})
	pulumi.RegisterOutputType(ApplicationGroupOutput{})
	pulumi.RegisterOutputType(ApplicationGroupArrayOutput{})
	pulumi.RegisterOutputType(ApplicationGroupMapOutput{})
}