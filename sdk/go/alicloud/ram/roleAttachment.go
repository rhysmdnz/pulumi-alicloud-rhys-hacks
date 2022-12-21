// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ram

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a RAM role attachment resource to bind role for several ECS instances.
type RoleAttachment struct {
	pulumi.CustomResourceState

	// The list of ECS instance's IDs.
	InstanceIds pulumi.StringArrayOutput `pulumi:"instanceIds"`
	// The name of role used to bind. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.
	RoleName pulumi.StringOutput `pulumi:"roleName"`
}

// NewRoleAttachment registers a new resource with the given unique name, arguments, and options.
func NewRoleAttachment(ctx *pulumi.Context,
	name string, args *RoleAttachmentArgs, opts ...pulumi.ResourceOption) (*RoleAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceIds == nil {
		return nil, errors.New("invalid value for required argument 'InstanceIds'")
	}
	if args.RoleName == nil {
		return nil, errors.New("invalid value for required argument 'RoleName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource RoleAttachment
	err := ctx.RegisterResource("alicloud:ram/roleAttachment:RoleAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoleAttachment gets an existing RoleAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoleAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleAttachmentState, opts ...pulumi.ResourceOption) (*RoleAttachment, error) {
	var resource RoleAttachment
	err := ctx.ReadResource("alicloud:ram/roleAttachment:RoleAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RoleAttachment resources.
type roleAttachmentState struct {
	// The list of ECS instance's IDs.
	InstanceIds []string `pulumi:"instanceIds"`
	// The name of role used to bind. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.
	RoleName *string `pulumi:"roleName"`
}

type RoleAttachmentState struct {
	// The list of ECS instance's IDs.
	InstanceIds pulumi.StringArrayInput
	// The name of role used to bind. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.
	RoleName pulumi.StringPtrInput
}

func (RoleAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleAttachmentState)(nil)).Elem()
}

type roleAttachmentArgs struct {
	// The list of ECS instance's IDs.
	InstanceIds []string `pulumi:"instanceIds"`
	// The name of role used to bind. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.
	RoleName string `pulumi:"roleName"`
}

// The set of arguments for constructing a RoleAttachment resource.
type RoleAttachmentArgs struct {
	// The list of ECS instance's IDs.
	InstanceIds pulumi.StringArrayInput
	// The name of role used to bind. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.
	RoleName pulumi.StringInput
}

func (RoleAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleAttachmentArgs)(nil)).Elem()
}

type RoleAttachmentInput interface {
	pulumi.Input

	ToRoleAttachmentOutput() RoleAttachmentOutput
	ToRoleAttachmentOutputWithContext(ctx context.Context) RoleAttachmentOutput
}

func (*RoleAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleAttachment)(nil)).Elem()
}

func (i *RoleAttachment) ToRoleAttachmentOutput() RoleAttachmentOutput {
	return i.ToRoleAttachmentOutputWithContext(context.Background())
}

func (i *RoleAttachment) ToRoleAttachmentOutputWithContext(ctx context.Context) RoleAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleAttachmentOutput)
}

// RoleAttachmentArrayInput is an input type that accepts RoleAttachmentArray and RoleAttachmentArrayOutput values.
// You can construct a concrete instance of `RoleAttachmentArrayInput` via:
//
//	RoleAttachmentArray{ RoleAttachmentArgs{...} }
type RoleAttachmentArrayInput interface {
	pulumi.Input

	ToRoleAttachmentArrayOutput() RoleAttachmentArrayOutput
	ToRoleAttachmentArrayOutputWithContext(context.Context) RoleAttachmentArrayOutput
}

type RoleAttachmentArray []RoleAttachmentInput

func (RoleAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RoleAttachment)(nil)).Elem()
}

func (i RoleAttachmentArray) ToRoleAttachmentArrayOutput() RoleAttachmentArrayOutput {
	return i.ToRoleAttachmentArrayOutputWithContext(context.Background())
}

func (i RoleAttachmentArray) ToRoleAttachmentArrayOutputWithContext(ctx context.Context) RoleAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleAttachmentArrayOutput)
}

// RoleAttachmentMapInput is an input type that accepts RoleAttachmentMap and RoleAttachmentMapOutput values.
// You can construct a concrete instance of `RoleAttachmentMapInput` via:
//
//	RoleAttachmentMap{ "key": RoleAttachmentArgs{...} }
type RoleAttachmentMapInput interface {
	pulumi.Input

	ToRoleAttachmentMapOutput() RoleAttachmentMapOutput
	ToRoleAttachmentMapOutputWithContext(context.Context) RoleAttachmentMapOutput
}

type RoleAttachmentMap map[string]RoleAttachmentInput

func (RoleAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RoleAttachment)(nil)).Elem()
}

func (i RoleAttachmentMap) ToRoleAttachmentMapOutput() RoleAttachmentMapOutput {
	return i.ToRoleAttachmentMapOutputWithContext(context.Background())
}

func (i RoleAttachmentMap) ToRoleAttachmentMapOutputWithContext(ctx context.Context) RoleAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleAttachmentMapOutput)
}

type RoleAttachmentOutput struct{ *pulumi.OutputState }

func (RoleAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleAttachment)(nil)).Elem()
}

func (o RoleAttachmentOutput) ToRoleAttachmentOutput() RoleAttachmentOutput {
	return o
}

func (o RoleAttachmentOutput) ToRoleAttachmentOutputWithContext(ctx context.Context) RoleAttachmentOutput {
	return o
}

// The list of ECS instance's IDs.
func (o RoleAttachmentOutput) InstanceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RoleAttachment) pulumi.StringArrayOutput { return v.InstanceIds }).(pulumi.StringArrayOutput)
}

// The name of role used to bind. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.
func (o RoleAttachmentOutput) RoleName() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleAttachment) pulumi.StringOutput { return v.RoleName }).(pulumi.StringOutput)
}

type RoleAttachmentArrayOutput struct{ *pulumi.OutputState }

func (RoleAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RoleAttachment)(nil)).Elem()
}

func (o RoleAttachmentArrayOutput) ToRoleAttachmentArrayOutput() RoleAttachmentArrayOutput {
	return o
}

func (o RoleAttachmentArrayOutput) ToRoleAttachmentArrayOutputWithContext(ctx context.Context) RoleAttachmentArrayOutput {
	return o
}

func (o RoleAttachmentArrayOutput) Index(i pulumi.IntInput) RoleAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RoleAttachment {
		return vs[0].([]*RoleAttachment)[vs[1].(int)]
	}).(RoleAttachmentOutput)
}

type RoleAttachmentMapOutput struct{ *pulumi.OutputState }

func (RoleAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RoleAttachment)(nil)).Elem()
}

func (o RoleAttachmentMapOutput) ToRoleAttachmentMapOutput() RoleAttachmentMapOutput {
	return o
}

func (o RoleAttachmentMapOutput) ToRoleAttachmentMapOutputWithContext(ctx context.Context) RoleAttachmentMapOutput {
	return o
}

func (o RoleAttachmentMapOutput) MapIndex(k pulumi.StringInput) RoleAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RoleAttachment {
		return vs[0].(map[string]*RoleAttachment)[vs[1].(string)]
	}).(RoleAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RoleAttachmentInput)(nil)).Elem(), &RoleAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleAttachmentArrayInput)(nil)).Elem(), RoleAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleAttachmentMapInput)(nil)).Elem(), RoleAttachmentMap{})
	pulumi.RegisterOutputType(RoleAttachmentOutput{})
	pulumi.RegisterOutputType(RoleAttachmentArrayOutput{})
	pulumi.RegisterOutputType(RoleAttachmentMapOutput{})
}