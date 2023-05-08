// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package drds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Distributed Relational Database Service (DRDS) is a lightweight (stateless), flexible, stable, and efficient middleware product independently developed by Alibaba Group to resolve scalability issues with single-host relational databases.
// With its compatibility with MySQL protocols and syntaxes, DRDS enables database/table sharding, smooth scaling, configuration upgrade/downgrade,
// transparent read/write splitting, and distributed transactions, providing O&M capabilities for distributed databases throughout their entire lifecycle.
//
// For information about DRDS and how to use it, see [What is DRDS](https://www.alibabacloud.com/help/product/29657.htm).
//
// > **NOTE:** At present, DRDS instance only can be supported in the regions: cn-shenzhen, cn-beijing, cn-hangzhou, cn-hongkong, cn-qingdao, ap-southeast-1.
//
// > **NOTE:** Currently, this resource only support `Domestic Site Account`.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/drds"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := drds.NewInstance(ctx, "default", &drds.InstanceArgs{
//				Description:        pulumi.String("drds instance"),
//				InstanceChargeType: pulumi.String("PostPaid"),
//				InstanceSeries:     pulumi.String("drds.sn1.4c8g"),
//				Specification:      pulumi.String("drds.sn1.4c8g.8C16G"),
//				VswitchId:          pulumi.String("vsw-bp1jlu3swk8rq2yoi40ey"),
//				ZoneId:             pulumi.String("cn-hangzhou-e"),
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
// Distributed Relational Database Service (DRDS) can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:drds/instance:Instance example drds-abc123456
//
// ```
type Instance struct {
	pulumi.CustomResourceState

	// (Available in 1.196.0+) The connection string of the DRDS instance.
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	// Description of the DRDS instance, This description can have a string of 2 to 256 characters.
	Description pulumi.StringOutput `pulumi:"description"`
	// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`.
	InstanceChargeType pulumi.StringPtrOutput `pulumi:"instanceChargeType"`
	// The parameter of the instance series. **NOTE:**  `drds.sn1.4c8g`,`drds.sn1.8c16g`,`drds.sn1.16c32g`,`drds.sn1.32c64g` are no longer supported. Valid values:
	// - `drds.sn2.4c16g` Starter Edition.
	// - `drds.sn2.8c32g` Standard Edition.
	// - `drds.sn2.16c64g` Enterprise Edition.
	InstanceSeries pulumi.StringOutput `pulumi:"instanceSeries"`
	// The MySQL version supported by the instance, with the following range of values. `5`: Fully compatible with MySQL 5.x (default) `8`: Fully compatible with MySQL 8.0. This parameter takes effect when the primary instance is created, and the read-only instance has the same MySQL version as the primary instance by default.
	MysqlVersion pulumi.IntOutput `pulumi:"mysqlVersion"`
	// (Available in 1.196.0+) The connection port of the DRDS instance.
	Port pulumi.StringOutput `pulumi:"port"`
	// User-defined DRDS instance specification. Value range:
	// - `drds.sn1.4c8g` for DRDS instance Starter version;
	// - value range : `drds.sn1.4c8g.8c16g`, `drds.sn1.4c8g.16c32g`, `drds.sn1.4c8g.32c64g`, `drds.sn1.4c8g.64c128g`
	// - `drds.sn1.8c16g` for DRDS instance Standard edition;
	// - value range : `drds.sn1.8c16g.16c32g`, `drds.sn1.8c16g.32c64g`, `drds.sn1.8c16g.64c128g`
	// - `drds.sn1.16c32g` for DRDS instance Enterprise Edition;
	// - value range : `drds.sn1.16c32g.32c64g`, `drds.sn1.16c32g.64c128g`
	// - `drds.sn1.32c64g` for DRDS instance Extreme Edition;
	// - value range : `drds.sn1.32c64g.128c256g`
	Specification pulumi.StringOutput `pulumi:"specification"`
	// The id of the VPC.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The VSwitch ID to launch in.
	VswitchId pulumi.StringOutput `pulumi:"vswitchId"`
	// The Zone to launch the DRDS instance.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.InstanceSeries == nil {
		return nil, errors.New("invalid value for required argument 'InstanceSeries'")
	}
	if args.Specification == nil {
		return nil, errors.New("invalid value for required argument 'Specification'")
	}
	if args.VswitchId == nil {
		return nil, errors.New("invalid value for required argument 'VswitchId'")
	}
	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Instance
	err := ctx.RegisterResource("alicloud:drds/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("alicloud:drds/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// (Available in 1.196.0+) The connection string of the DRDS instance.
	ConnectionString *string `pulumi:"connectionString"`
	// Description of the DRDS instance, This description can have a string of 2 to 256 characters.
	Description *string `pulumi:"description"`
	// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// The parameter of the instance series. **NOTE:**  `drds.sn1.4c8g`,`drds.sn1.8c16g`,`drds.sn1.16c32g`,`drds.sn1.32c64g` are no longer supported. Valid values:
	// - `drds.sn2.4c16g` Starter Edition.
	// - `drds.sn2.8c32g` Standard Edition.
	// - `drds.sn2.16c64g` Enterprise Edition.
	InstanceSeries *string `pulumi:"instanceSeries"`
	// The MySQL version supported by the instance, with the following range of values. `5`: Fully compatible with MySQL 5.x (default) `8`: Fully compatible with MySQL 8.0. This parameter takes effect when the primary instance is created, and the read-only instance has the same MySQL version as the primary instance by default.
	MysqlVersion *int `pulumi:"mysqlVersion"`
	// (Available in 1.196.0+) The connection port of the DRDS instance.
	Port *string `pulumi:"port"`
	// User-defined DRDS instance specification. Value range:
	// - `drds.sn1.4c8g` for DRDS instance Starter version;
	// - value range : `drds.sn1.4c8g.8c16g`, `drds.sn1.4c8g.16c32g`, `drds.sn1.4c8g.32c64g`, `drds.sn1.4c8g.64c128g`
	// - `drds.sn1.8c16g` for DRDS instance Standard edition;
	// - value range : `drds.sn1.8c16g.16c32g`, `drds.sn1.8c16g.32c64g`, `drds.sn1.8c16g.64c128g`
	// - `drds.sn1.16c32g` for DRDS instance Enterprise Edition;
	// - value range : `drds.sn1.16c32g.32c64g`, `drds.sn1.16c32g.64c128g`
	// - `drds.sn1.32c64g` for DRDS instance Extreme Edition;
	// - value range : `drds.sn1.32c64g.128c256g`
	Specification *string `pulumi:"specification"`
	// The id of the VPC.
	VpcId *string `pulumi:"vpcId"`
	// The VSwitch ID to launch in.
	VswitchId *string `pulumi:"vswitchId"`
	// The Zone to launch the DRDS instance.
	ZoneId *string `pulumi:"zoneId"`
}

type InstanceState struct {
	// (Available in 1.196.0+) The connection string of the DRDS instance.
	ConnectionString pulumi.StringPtrInput
	// Description of the DRDS instance, This description can have a string of 2 to 256 characters.
	Description pulumi.StringPtrInput
	// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`.
	InstanceChargeType pulumi.StringPtrInput
	// The parameter of the instance series. **NOTE:**  `drds.sn1.4c8g`,`drds.sn1.8c16g`,`drds.sn1.16c32g`,`drds.sn1.32c64g` are no longer supported. Valid values:
	// - `drds.sn2.4c16g` Starter Edition.
	// - `drds.sn2.8c32g` Standard Edition.
	// - `drds.sn2.16c64g` Enterprise Edition.
	InstanceSeries pulumi.StringPtrInput
	// The MySQL version supported by the instance, with the following range of values. `5`: Fully compatible with MySQL 5.x (default) `8`: Fully compatible with MySQL 8.0. This parameter takes effect when the primary instance is created, and the read-only instance has the same MySQL version as the primary instance by default.
	MysqlVersion pulumi.IntPtrInput
	// (Available in 1.196.0+) The connection port of the DRDS instance.
	Port pulumi.StringPtrInput
	// User-defined DRDS instance specification. Value range:
	// - `drds.sn1.4c8g` for DRDS instance Starter version;
	// - value range : `drds.sn1.4c8g.8c16g`, `drds.sn1.4c8g.16c32g`, `drds.sn1.4c8g.32c64g`, `drds.sn1.4c8g.64c128g`
	// - `drds.sn1.8c16g` for DRDS instance Standard edition;
	// - value range : `drds.sn1.8c16g.16c32g`, `drds.sn1.8c16g.32c64g`, `drds.sn1.8c16g.64c128g`
	// - `drds.sn1.16c32g` for DRDS instance Enterprise Edition;
	// - value range : `drds.sn1.16c32g.32c64g`, `drds.sn1.16c32g.64c128g`
	// - `drds.sn1.32c64g` for DRDS instance Extreme Edition;
	// - value range : `drds.sn1.32c64g.128c256g`
	Specification pulumi.StringPtrInput
	// The id of the VPC.
	VpcId pulumi.StringPtrInput
	// The VSwitch ID to launch in.
	VswitchId pulumi.StringPtrInput
	// The Zone to launch the DRDS instance.
	ZoneId pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Description of the DRDS instance, This description can have a string of 2 to 256 characters.
	Description string `pulumi:"description"`
	// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// The parameter of the instance series. **NOTE:**  `drds.sn1.4c8g`,`drds.sn1.8c16g`,`drds.sn1.16c32g`,`drds.sn1.32c64g` are no longer supported. Valid values:
	// - `drds.sn2.4c16g` Starter Edition.
	// - `drds.sn2.8c32g` Standard Edition.
	// - `drds.sn2.16c64g` Enterprise Edition.
	InstanceSeries string `pulumi:"instanceSeries"`
	// The MySQL version supported by the instance, with the following range of values. `5`: Fully compatible with MySQL 5.x (default) `8`: Fully compatible with MySQL 8.0. This parameter takes effect when the primary instance is created, and the read-only instance has the same MySQL version as the primary instance by default.
	MysqlVersion *int `pulumi:"mysqlVersion"`
	// User-defined DRDS instance specification. Value range:
	// - `drds.sn1.4c8g` for DRDS instance Starter version;
	// - value range : `drds.sn1.4c8g.8c16g`, `drds.sn1.4c8g.16c32g`, `drds.sn1.4c8g.32c64g`, `drds.sn1.4c8g.64c128g`
	// - `drds.sn1.8c16g` for DRDS instance Standard edition;
	// - value range : `drds.sn1.8c16g.16c32g`, `drds.sn1.8c16g.32c64g`, `drds.sn1.8c16g.64c128g`
	// - `drds.sn1.16c32g` for DRDS instance Enterprise Edition;
	// - value range : `drds.sn1.16c32g.32c64g`, `drds.sn1.16c32g.64c128g`
	// - `drds.sn1.32c64g` for DRDS instance Extreme Edition;
	// - value range : `drds.sn1.32c64g.128c256g`
	Specification string `pulumi:"specification"`
	// The id of the VPC.
	VpcId *string `pulumi:"vpcId"`
	// The VSwitch ID to launch in.
	VswitchId string `pulumi:"vswitchId"`
	// The Zone to launch the DRDS instance.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Description of the DRDS instance, This description can have a string of 2 to 256 characters.
	Description pulumi.StringInput
	// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`.
	InstanceChargeType pulumi.StringPtrInput
	// The parameter of the instance series. **NOTE:**  `drds.sn1.4c8g`,`drds.sn1.8c16g`,`drds.sn1.16c32g`,`drds.sn1.32c64g` are no longer supported. Valid values:
	// - `drds.sn2.4c16g` Starter Edition.
	// - `drds.sn2.8c32g` Standard Edition.
	// - `drds.sn2.16c64g` Enterprise Edition.
	InstanceSeries pulumi.StringInput
	// The MySQL version supported by the instance, with the following range of values. `5`: Fully compatible with MySQL 5.x (default) `8`: Fully compatible with MySQL 8.0. This parameter takes effect when the primary instance is created, and the read-only instance has the same MySQL version as the primary instance by default.
	MysqlVersion pulumi.IntPtrInput
	// User-defined DRDS instance specification. Value range:
	// - `drds.sn1.4c8g` for DRDS instance Starter version;
	// - value range : `drds.sn1.4c8g.8c16g`, `drds.sn1.4c8g.16c32g`, `drds.sn1.4c8g.32c64g`, `drds.sn1.4c8g.64c128g`
	// - `drds.sn1.8c16g` for DRDS instance Standard edition;
	// - value range : `drds.sn1.8c16g.16c32g`, `drds.sn1.8c16g.32c64g`, `drds.sn1.8c16g.64c128g`
	// - `drds.sn1.16c32g` for DRDS instance Enterprise Edition;
	// - value range : `drds.sn1.16c32g.32c64g`, `drds.sn1.16c32g.64c128g`
	// - `drds.sn1.32c64g` for DRDS instance Extreme Edition;
	// - value range : `drds.sn1.32c64g.128c256g`
	Specification pulumi.StringInput
	// The id of the VPC.
	VpcId pulumi.StringPtrInput
	// The VSwitch ID to launch in.
	VswitchId pulumi.StringInput
	// The Zone to launch the DRDS instance.
	ZoneId pulumi.StringInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//	InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//	InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

// (Available in 1.196.0+) The connection string of the DRDS instance.
func (o InstanceOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ConnectionString }).(pulumi.StringOutput)
}

// Description of the DRDS instance, This description can have a string of 2 to 256 characters.
func (o InstanceOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`.
func (o InstanceOutput) InstanceChargeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.InstanceChargeType }).(pulumi.StringPtrOutput)
}

// The parameter of the instance series. **NOTE:**  `drds.sn1.4c8g`,`drds.sn1.8c16g`,`drds.sn1.16c32g`,`drds.sn1.32c64g` are no longer supported. Valid values:
// - `drds.sn2.4c16g` Starter Edition.
// - `drds.sn2.8c32g` Standard Edition.
// - `drds.sn2.16c64g` Enterprise Edition.
func (o InstanceOutput) InstanceSeries() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.InstanceSeries }).(pulumi.StringOutput)
}

// The MySQL version supported by the instance, with the following range of values. `5`: Fully compatible with MySQL 5.x (default) `8`: Fully compatible with MySQL 8.0. This parameter takes effect when the primary instance is created, and the read-only instance has the same MySQL version as the primary instance by default.
func (o InstanceOutput) MysqlVersion() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.MysqlVersion }).(pulumi.IntOutput)
}

// (Available in 1.196.0+) The connection port of the DRDS instance.
func (o InstanceOutput) Port() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Port }).(pulumi.StringOutput)
}

// User-defined DRDS instance specification. Value range:
// - `drds.sn1.4c8g` for DRDS instance Starter version;
// - value range : `drds.sn1.4c8g.8c16g`, `drds.sn1.4c8g.16c32g`, `drds.sn1.4c8g.32c64g`, `drds.sn1.4c8g.64c128g`
// - `drds.sn1.8c16g` for DRDS instance Standard edition;
// - value range : `drds.sn1.8c16g.16c32g`, `drds.sn1.8c16g.32c64g`, `drds.sn1.8c16g.64c128g`
// - `drds.sn1.16c32g` for DRDS instance Enterprise Edition;
// - value range : `drds.sn1.16c32g.32c64g`, `drds.sn1.16c32g.64c128g`
// - `drds.sn1.32c64g` for DRDS instance Extreme Edition;
// - value range : `drds.sn1.32c64g.128c256g`
func (o InstanceOutput) Specification() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Specification }).(pulumi.StringOutput)
}

// The id of the VPC.
func (o InstanceOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// The VSwitch ID to launch in.
func (o InstanceOutput) VswitchId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.VswitchId }).(pulumi.StringOutput)
}

// The Zone to launch the DRDS instance.
func (o InstanceOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].([]*Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].(map[string]*Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceArrayInput)(nil)).Elem(), InstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMapInput)(nil)).Elem(), InstanceMap{})
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}
