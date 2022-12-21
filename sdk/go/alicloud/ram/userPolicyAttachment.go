// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ram

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a RAM User Policy attachment resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/ram"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			user, err := ram.NewUser(ctx, "user", &ram.UserArgs{
//				DisplayName: pulumi.String("user_display_name"),
//				Mobile:      pulumi.String("86-18688888888"),
//				Email:       pulumi.String("hello.uuu@aaa.com"),
//				Comments:    pulumi.String("yoyoyo"),
//				Force:       pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			policy, err := ram.NewPolicy(ctx, "policy", &ram.PolicyArgs{
//				Document: pulumi.String(fmt.Sprintf(`  {
//	    "Statement": [
//	      {
//	        "Action": [
//	          "oss:ListObjects",
//	          "oss:GetObject"
//	        ],
//	        "Effect": "Allow",
//	        "Resource": [
//	          "acs:oss:*:*:mybucket",
//	          "acs:oss:*:*:mybucket/*"
//	        ]
//	      }
//	    ],
//	      "Version": "1"
//	  }
//
// `)),
//
//				Description: pulumi.String("this is a policy test"),
//				Force:       pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ram.NewUserPolicyAttachment(ctx, "attach", &ram.UserPolicyAttachmentArgs{
//				PolicyName: policy.Name,
//				PolicyType: policy.Type,
//				UserName:   user.Name,
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
// RAM User Policy attachment can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:ram/userPolicyAttachment:UserPolicyAttachment example user:my-policy:Custom:my-user
//
// ```
type UserPolicyAttachment struct {
	pulumi.CustomResourceState

	// Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
	PolicyName pulumi.StringOutput `pulumi:"policyName"`
	// Type of the RAM policy. It must be `Custom` or `System`.
	PolicyType pulumi.StringOutput `pulumi:"policyType"`
	// Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
	UserName pulumi.StringOutput `pulumi:"userName"`
}

// NewUserPolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewUserPolicyAttachment(ctx *pulumi.Context,
	name string, args *UserPolicyAttachmentArgs, opts ...pulumi.ResourceOption) (*UserPolicyAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyName == nil {
		return nil, errors.New("invalid value for required argument 'PolicyName'")
	}
	if args.PolicyType == nil {
		return nil, errors.New("invalid value for required argument 'PolicyType'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource UserPolicyAttachment
	err := ctx.RegisterResource("alicloud:ram/userPolicyAttachment:UserPolicyAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserPolicyAttachment gets an existing UserPolicyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserPolicyAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserPolicyAttachmentState, opts ...pulumi.ResourceOption) (*UserPolicyAttachment, error) {
	var resource UserPolicyAttachment
	err := ctx.ReadResource("alicloud:ram/userPolicyAttachment:UserPolicyAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserPolicyAttachment resources.
type userPolicyAttachmentState struct {
	// Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
	PolicyName *string `pulumi:"policyName"`
	// Type of the RAM policy. It must be `Custom` or `System`.
	PolicyType *string `pulumi:"policyType"`
	// Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
	UserName *string `pulumi:"userName"`
}

type UserPolicyAttachmentState struct {
	// Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
	PolicyName pulumi.StringPtrInput
	// Type of the RAM policy. It must be `Custom` or `System`.
	PolicyType pulumi.StringPtrInput
	// Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
	UserName pulumi.StringPtrInput
}

func (UserPolicyAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*userPolicyAttachmentState)(nil)).Elem()
}

type userPolicyAttachmentArgs struct {
	// Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
	PolicyName string `pulumi:"policyName"`
	// Type of the RAM policy. It must be `Custom` or `System`.
	PolicyType string `pulumi:"policyType"`
	// Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
	UserName string `pulumi:"userName"`
}

// The set of arguments for constructing a UserPolicyAttachment resource.
type UserPolicyAttachmentArgs struct {
	// Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
	PolicyName pulumi.StringInput
	// Type of the RAM policy. It must be `Custom` or `System`.
	PolicyType pulumi.StringInput
	// Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
	UserName pulumi.StringInput
}

func (UserPolicyAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userPolicyAttachmentArgs)(nil)).Elem()
}

type UserPolicyAttachmentInput interface {
	pulumi.Input

	ToUserPolicyAttachmentOutput() UserPolicyAttachmentOutput
	ToUserPolicyAttachmentOutputWithContext(ctx context.Context) UserPolicyAttachmentOutput
}

func (*UserPolicyAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPolicyAttachment)(nil)).Elem()
}

func (i *UserPolicyAttachment) ToUserPolicyAttachmentOutput() UserPolicyAttachmentOutput {
	return i.ToUserPolicyAttachmentOutputWithContext(context.Background())
}

func (i *UserPolicyAttachment) ToUserPolicyAttachmentOutputWithContext(ctx context.Context) UserPolicyAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPolicyAttachmentOutput)
}

// UserPolicyAttachmentArrayInput is an input type that accepts UserPolicyAttachmentArray and UserPolicyAttachmentArrayOutput values.
// You can construct a concrete instance of `UserPolicyAttachmentArrayInput` via:
//
//	UserPolicyAttachmentArray{ UserPolicyAttachmentArgs{...} }
type UserPolicyAttachmentArrayInput interface {
	pulumi.Input

	ToUserPolicyAttachmentArrayOutput() UserPolicyAttachmentArrayOutput
	ToUserPolicyAttachmentArrayOutputWithContext(context.Context) UserPolicyAttachmentArrayOutput
}

type UserPolicyAttachmentArray []UserPolicyAttachmentInput

func (UserPolicyAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserPolicyAttachment)(nil)).Elem()
}

func (i UserPolicyAttachmentArray) ToUserPolicyAttachmentArrayOutput() UserPolicyAttachmentArrayOutput {
	return i.ToUserPolicyAttachmentArrayOutputWithContext(context.Background())
}

func (i UserPolicyAttachmentArray) ToUserPolicyAttachmentArrayOutputWithContext(ctx context.Context) UserPolicyAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPolicyAttachmentArrayOutput)
}

// UserPolicyAttachmentMapInput is an input type that accepts UserPolicyAttachmentMap and UserPolicyAttachmentMapOutput values.
// You can construct a concrete instance of `UserPolicyAttachmentMapInput` via:
//
//	UserPolicyAttachmentMap{ "key": UserPolicyAttachmentArgs{...} }
type UserPolicyAttachmentMapInput interface {
	pulumi.Input

	ToUserPolicyAttachmentMapOutput() UserPolicyAttachmentMapOutput
	ToUserPolicyAttachmentMapOutputWithContext(context.Context) UserPolicyAttachmentMapOutput
}

type UserPolicyAttachmentMap map[string]UserPolicyAttachmentInput

func (UserPolicyAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserPolicyAttachment)(nil)).Elem()
}

func (i UserPolicyAttachmentMap) ToUserPolicyAttachmentMapOutput() UserPolicyAttachmentMapOutput {
	return i.ToUserPolicyAttachmentMapOutputWithContext(context.Background())
}

func (i UserPolicyAttachmentMap) ToUserPolicyAttachmentMapOutputWithContext(ctx context.Context) UserPolicyAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPolicyAttachmentMapOutput)
}

type UserPolicyAttachmentOutput struct{ *pulumi.OutputState }

func (UserPolicyAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPolicyAttachment)(nil)).Elem()
}

func (o UserPolicyAttachmentOutput) ToUserPolicyAttachmentOutput() UserPolicyAttachmentOutput {
	return o
}

func (o UserPolicyAttachmentOutput) ToUserPolicyAttachmentOutputWithContext(ctx context.Context) UserPolicyAttachmentOutput {
	return o
}

// Name of the RAM policy. This name can have a string of 1 to 128 characters, must contain only alphanumeric characters or hyphen "-", and must not begin with a hyphen.
func (o UserPolicyAttachmentOutput) PolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *UserPolicyAttachment) pulumi.StringOutput { return v.PolicyName }).(pulumi.StringOutput)
}

// Type of the RAM policy. It must be `Custom` or `System`.
func (o UserPolicyAttachmentOutput) PolicyType() pulumi.StringOutput {
	return o.ApplyT(func(v *UserPolicyAttachment) pulumi.StringOutput { return v.PolicyType }).(pulumi.StringOutput)
}

// Name of the RAM user. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin with a hyphen.
func (o UserPolicyAttachmentOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *UserPolicyAttachment) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

type UserPolicyAttachmentArrayOutput struct{ *pulumi.OutputState }

func (UserPolicyAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserPolicyAttachment)(nil)).Elem()
}

func (o UserPolicyAttachmentArrayOutput) ToUserPolicyAttachmentArrayOutput() UserPolicyAttachmentArrayOutput {
	return o
}

func (o UserPolicyAttachmentArrayOutput) ToUserPolicyAttachmentArrayOutputWithContext(ctx context.Context) UserPolicyAttachmentArrayOutput {
	return o
}

func (o UserPolicyAttachmentArrayOutput) Index(i pulumi.IntInput) UserPolicyAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserPolicyAttachment {
		return vs[0].([]*UserPolicyAttachment)[vs[1].(int)]
	}).(UserPolicyAttachmentOutput)
}

type UserPolicyAttachmentMapOutput struct{ *pulumi.OutputState }

func (UserPolicyAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserPolicyAttachment)(nil)).Elem()
}

func (o UserPolicyAttachmentMapOutput) ToUserPolicyAttachmentMapOutput() UserPolicyAttachmentMapOutput {
	return o
}

func (o UserPolicyAttachmentMapOutput) ToUserPolicyAttachmentMapOutputWithContext(ctx context.Context) UserPolicyAttachmentMapOutput {
	return o
}

func (o UserPolicyAttachmentMapOutput) MapIndex(k pulumi.StringInput) UserPolicyAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserPolicyAttachment {
		return vs[0].(map[string]*UserPolicyAttachment)[vs[1].(string)]
	}).(UserPolicyAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserPolicyAttachmentInput)(nil)).Elem(), &UserPolicyAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserPolicyAttachmentArrayInput)(nil)).Elem(), UserPolicyAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserPolicyAttachmentMapInput)(nil)).Elem(), UserPolicyAttachmentMap{})
	pulumi.RegisterOutputType(UserPolicyAttachmentOutput{})
	pulumi.RegisterOutputType(UserPolicyAttachmentArrayOutput{})
	pulumi.RegisterOutputType(UserPolicyAttachmentMapOutput{})
}