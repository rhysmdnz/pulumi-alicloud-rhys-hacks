// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VPC Ipv4 Gateway resource.
//
// For information about VPC Ipv4 Gateway and how to use it, see [What is Ipv4 Gateway](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/createipv4gateway).
//
// > **NOTE:** Available in v1.181.0+.
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
//	"github.com/pulumi/pulumi-alicloud/sdk/go/alicloud/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := vpc.GetNetworks(ctx, &vpc.GetNetworksArgs{
//				NameRegex: pulumi.StringRef("default-NoDeleting"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vpc.NewIpv4Gateway(ctx, "example", &vpc.Ipv4GatewayArgs{
//				Ipv4GatewayName: pulumi.String("example_value"),
//				VpcId:           pulumi.String(_default.Ids[0]),
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
// VPC Ipv4 Gateway can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:vpc/ipv4Gateway:Ipv4Gateway example <id>
//
// ```
type Ipv4Gateway struct {
	pulumi.CustomResourceState

	// The dry run.
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// Whether the IPv4 gateway is active or not. Valid values are `true` and `false`.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The description of the IPv4 gateway. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
	Ipv4GatewayDescription pulumi.StringPtrOutput `pulumi:"ipv4GatewayDescription"`
	// The name of the IPv4 gateway. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	Ipv4GatewayName pulumi.StringPtrOutput `pulumi:"ipv4GatewayName"`
	// The status of the resource.
	Status pulumi.StringOutput `pulumi:"status"`
	// The ID of the virtual private cloud (VPC) where you want to create the IPv4 gateway. You can create only one IPv4 gateway in a VPC.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewIpv4Gateway registers a new resource with the given unique name, arguments, and options.
func NewIpv4Gateway(ctx *pulumi.Context,
	name string, args *Ipv4GatewayArgs, opts ...pulumi.ResourceOption) (*Ipv4Gateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Ipv4Gateway
	err := ctx.RegisterResource("alicloud:vpc/ipv4Gateway:Ipv4Gateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpv4Gateway gets an existing Ipv4Gateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpv4Gateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Ipv4GatewayState, opts ...pulumi.ResourceOption) (*Ipv4Gateway, error) {
	var resource Ipv4Gateway
	err := ctx.ReadResource("alicloud:vpc/ipv4Gateway:Ipv4Gateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Ipv4Gateway resources.
type ipv4GatewayState struct {
	// The dry run.
	DryRun *bool `pulumi:"dryRun"`
	// Whether the IPv4 gateway is active or not. Valid values are `true` and `false`.
	Enabled *bool `pulumi:"enabled"`
	// The description of the IPv4 gateway. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
	Ipv4GatewayDescription *string `pulumi:"ipv4GatewayDescription"`
	// The name of the IPv4 gateway. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	Ipv4GatewayName *string `pulumi:"ipv4GatewayName"`
	// The status of the resource.
	Status *string `pulumi:"status"`
	// The ID of the virtual private cloud (VPC) where you want to create the IPv4 gateway. You can create only one IPv4 gateway in a VPC.
	VpcId *string `pulumi:"vpcId"`
}

type Ipv4GatewayState struct {
	// The dry run.
	DryRun pulumi.BoolPtrInput
	// Whether the IPv4 gateway is active or not. Valid values are `true` and `false`.
	Enabled pulumi.BoolPtrInput
	// The description of the IPv4 gateway. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
	Ipv4GatewayDescription pulumi.StringPtrInput
	// The name of the IPv4 gateway. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	Ipv4GatewayName pulumi.StringPtrInput
	// The status of the resource.
	Status pulumi.StringPtrInput
	// The ID of the virtual private cloud (VPC) where you want to create the IPv4 gateway. You can create only one IPv4 gateway in a VPC.
	VpcId pulumi.StringPtrInput
}

func (Ipv4GatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipv4GatewayState)(nil)).Elem()
}

type ipv4GatewayArgs struct {
	// The dry run.
	DryRun *bool `pulumi:"dryRun"`
	// Whether the IPv4 gateway is active or not. Valid values are `true` and `false`.
	Enabled *bool `pulumi:"enabled"`
	// The description of the IPv4 gateway. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
	Ipv4GatewayDescription *string `pulumi:"ipv4GatewayDescription"`
	// The name of the IPv4 gateway. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	Ipv4GatewayName *string `pulumi:"ipv4GatewayName"`
	// The ID of the virtual private cloud (VPC) where you want to create the IPv4 gateway. You can create only one IPv4 gateway in a VPC.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a Ipv4Gateway resource.
type Ipv4GatewayArgs struct {
	// The dry run.
	DryRun pulumi.BoolPtrInput
	// Whether the IPv4 gateway is active or not. Valid values are `true` and `false`.
	Enabled pulumi.BoolPtrInput
	// The description of the IPv4 gateway. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
	Ipv4GatewayDescription pulumi.StringPtrInput
	// The name of the IPv4 gateway. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
	Ipv4GatewayName pulumi.StringPtrInput
	// The ID of the virtual private cloud (VPC) where you want to create the IPv4 gateway. You can create only one IPv4 gateway in a VPC.
	VpcId pulumi.StringInput
}

func (Ipv4GatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipv4GatewayArgs)(nil)).Elem()
}

type Ipv4GatewayInput interface {
	pulumi.Input

	ToIpv4GatewayOutput() Ipv4GatewayOutput
	ToIpv4GatewayOutputWithContext(ctx context.Context) Ipv4GatewayOutput
}

func (*Ipv4Gateway) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipv4Gateway)(nil)).Elem()
}

func (i *Ipv4Gateway) ToIpv4GatewayOutput() Ipv4GatewayOutput {
	return i.ToIpv4GatewayOutputWithContext(context.Background())
}

func (i *Ipv4Gateway) ToIpv4GatewayOutputWithContext(ctx context.Context) Ipv4GatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv4GatewayOutput)
}

// Ipv4GatewayArrayInput is an input type that accepts Ipv4GatewayArray and Ipv4GatewayArrayOutput values.
// You can construct a concrete instance of `Ipv4GatewayArrayInput` via:
//
//	Ipv4GatewayArray{ Ipv4GatewayArgs{...} }
type Ipv4GatewayArrayInput interface {
	pulumi.Input

	ToIpv4GatewayArrayOutput() Ipv4GatewayArrayOutput
	ToIpv4GatewayArrayOutputWithContext(context.Context) Ipv4GatewayArrayOutput
}

type Ipv4GatewayArray []Ipv4GatewayInput

func (Ipv4GatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ipv4Gateway)(nil)).Elem()
}

func (i Ipv4GatewayArray) ToIpv4GatewayArrayOutput() Ipv4GatewayArrayOutput {
	return i.ToIpv4GatewayArrayOutputWithContext(context.Background())
}

func (i Ipv4GatewayArray) ToIpv4GatewayArrayOutputWithContext(ctx context.Context) Ipv4GatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv4GatewayArrayOutput)
}

// Ipv4GatewayMapInput is an input type that accepts Ipv4GatewayMap and Ipv4GatewayMapOutput values.
// You can construct a concrete instance of `Ipv4GatewayMapInput` via:
//
//	Ipv4GatewayMap{ "key": Ipv4GatewayArgs{...} }
type Ipv4GatewayMapInput interface {
	pulumi.Input

	ToIpv4GatewayMapOutput() Ipv4GatewayMapOutput
	ToIpv4GatewayMapOutputWithContext(context.Context) Ipv4GatewayMapOutput
}

type Ipv4GatewayMap map[string]Ipv4GatewayInput

func (Ipv4GatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ipv4Gateway)(nil)).Elem()
}

func (i Ipv4GatewayMap) ToIpv4GatewayMapOutput() Ipv4GatewayMapOutput {
	return i.ToIpv4GatewayMapOutputWithContext(context.Background())
}

func (i Ipv4GatewayMap) ToIpv4GatewayMapOutputWithContext(ctx context.Context) Ipv4GatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Ipv4GatewayMapOutput)
}

type Ipv4GatewayOutput struct{ *pulumi.OutputState }

func (Ipv4GatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ipv4Gateway)(nil)).Elem()
}

func (o Ipv4GatewayOutput) ToIpv4GatewayOutput() Ipv4GatewayOutput {
	return o
}

func (o Ipv4GatewayOutput) ToIpv4GatewayOutputWithContext(ctx context.Context) Ipv4GatewayOutput {
	return o
}

// The dry run.
func (o Ipv4GatewayOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Ipv4Gateway) pulumi.BoolPtrOutput { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// Whether the IPv4 gateway is active or not. Valid values are `true` and `false`.
func (o Ipv4GatewayOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Ipv4Gateway) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// The description of the IPv4 gateway. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
func (o Ipv4GatewayOutput) Ipv4GatewayDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ipv4Gateway) pulumi.StringPtrOutput { return v.Ipv4GatewayDescription }).(pulumi.StringPtrOutput)
}

// The name of the IPv4 gateway. The name must be `2` to `128` characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter.
func (o Ipv4GatewayOutput) Ipv4GatewayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ipv4Gateway) pulumi.StringPtrOutput { return v.Ipv4GatewayName }).(pulumi.StringPtrOutput)
}

// The status of the resource.
func (o Ipv4GatewayOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv4Gateway) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The ID of the virtual private cloud (VPC) where you want to create the IPv4 gateway. You can create only one IPv4 gateway in a VPC.
func (o Ipv4GatewayOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Ipv4Gateway) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type Ipv4GatewayArrayOutput struct{ *pulumi.OutputState }

func (Ipv4GatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Ipv4Gateway)(nil)).Elem()
}

func (o Ipv4GatewayArrayOutput) ToIpv4GatewayArrayOutput() Ipv4GatewayArrayOutput {
	return o
}

func (o Ipv4GatewayArrayOutput) ToIpv4GatewayArrayOutputWithContext(ctx context.Context) Ipv4GatewayArrayOutput {
	return o
}

func (o Ipv4GatewayArrayOutput) Index(i pulumi.IntInput) Ipv4GatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Ipv4Gateway {
		return vs[0].([]*Ipv4Gateway)[vs[1].(int)]
	}).(Ipv4GatewayOutput)
}

type Ipv4GatewayMapOutput struct{ *pulumi.OutputState }

func (Ipv4GatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Ipv4Gateway)(nil)).Elem()
}

func (o Ipv4GatewayMapOutput) ToIpv4GatewayMapOutput() Ipv4GatewayMapOutput {
	return o
}

func (o Ipv4GatewayMapOutput) ToIpv4GatewayMapOutputWithContext(ctx context.Context) Ipv4GatewayMapOutput {
	return o
}

func (o Ipv4GatewayMapOutput) MapIndex(k pulumi.StringInput) Ipv4GatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Ipv4Gateway {
		return vs[0].(map[string]*Ipv4Gateway)[vs[1].(string)]
	}).(Ipv4GatewayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv4GatewayInput)(nil)).Elem(), &Ipv4Gateway{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv4GatewayArrayInput)(nil)).Elem(), Ipv4GatewayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*Ipv4GatewayMapInput)(nil)).Elem(), Ipv4GatewayMap{})
	pulumi.RegisterOutputType(Ipv4GatewayOutput{})
	pulumi.RegisterOutputType(Ipv4GatewayArrayOutput{})
	pulumi.RegisterOutputType(Ipv4GatewayMapOutput{})
}
