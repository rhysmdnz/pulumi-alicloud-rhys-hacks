// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an RDS database resource. A DB database deployed in a DB instance. A DB instance can own multiple databases.
//
// > **NOTE:** This resource does not support creating 'PPAS' database. You have to login RDS instance to create manually.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-alicloud/sdk/go/alicloud"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/rds"
//	"github.com/rhysmdnz/pulumi-alicloud/sdk/go/alicloud/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			creation := "Rds"
//			if param := cfg.Get("creation"); param != "" {
//				creation = param
//			}
//			name := "dbdatabasebasic"
//			if param := cfg.Get("name"); param != "" {
//				name = param
//			}
//			defaultZones, err := alicloud.GetZones(ctx, &GetZonesArgs{
//				AvailableResourceCreation: pulumi.StringRef(creation),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			defaultNetwork, err := vpc.NewNetwork(ctx, "defaultNetwork", &vpc.NetworkArgs{
//				CidrBlock: pulumi.String("172.16.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			defaultSwitch, err := vpc.NewSwitch(ctx, "defaultSwitch", &vpc.SwitchArgs{
//				VpcId:       defaultNetwork.ID(),
//				CidrBlock:   pulumi.String("172.16.0.0/24"),
//				ZoneId:      pulumi.String(defaultZones.Zones[0].Id),
//				VswitchName: pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			instance, err := rds.NewInstance(ctx, "instance", &rds.InstanceArgs{
//				Engine:          pulumi.String("MySQL"),
//				EngineVersion:   pulumi.String("5.6"),
//				InstanceType:    pulumi.String("rds.mysql.s1.small"),
//				InstanceStorage: pulumi.Int(10),
//				VswitchId:       defaultSwitch.ID(),
//				InstanceName:    pulumi.String(name),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = rds.NewDatabase(ctx, "defaultDatabase", &rds.DatabaseArgs{
//				InstanceId: instance.ID(),
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
// RDS database can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import alicloud:rds/database:Database example "rm-12345:tf_database"
//
// ```
type Database struct {
	pulumi.CustomResourceState

	// Character set. The value range is limited to the following:
	// - MySQL: [ utf8, gbk, latin1, utf8mb4 ] \(`utf8mb4` only supports versions 5.5 and 5.6\).
	// - SQLServer: [ Chinese_PRC_CI_AS, Chinese_PRC_CS_AS, SQL_Latin1_General_CP1_CI_AS, SQL_Latin1_General_CP1_CS_AS, Chinese_PRC_BIN ]
	// - PostgreSQL: [ KOI8U、UTF8、WIN866、WIN874、WIN1250、WIN1251、WIN1252、WIN1253、WIN1254、WIN1255、WIN1256、WIN1257、WIN1258、EUC_CN、EUC_KR、EUC_TW、EUC_JP、EUC_JIS_2004、KOI8R、MULE_INTERNAL、LATIN1、LATIN2、LATIN3、LATIN4、LATIN5、LATIN6、LATIN7、LATIN8、LATIN9、LATIN10、ISO_8859_5、ISO_8859_6、ISO_8859_7、ISO_8859_8、SQL_ASCII ]
	CharacterSet pulumi.StringPtrOutput `pulumi:"characterSet"`
	// Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Id of instance that can run database.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Name of the database requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter
	// and have no more than 64 characters.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Database
	err := ctx.RegisterResource("alicloud:rds/database:Database", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabase gets an existing Database resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	var resource Database
	err := ctx.ReadResource("alicloud:rds/database:Database", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Database resources.
type databaseState struct {
	// Character set. The value range is limited to the following:
	// - MySQL: [ utf8, gbk, latin1, utf8mb4 ] \(`utf8mb4` only supports versions 5.5 and 5.6\).
	// - SQLServer: [ Chinese_PRC_CI_AS, Chinese_PRC_CS_AS, SQL_Latin1_General_CP1_CI_AS, SQL_Latin1_General_CP1_CS_AS, Chinese_PRC_BIN ]
	// - PostgreSQL: [ KOI8U、UTF8、WIN866、WIN874、WIN1250、WIN1251、WIN1252、WIN1253、WIN1254、WIN1255、WIN1256、WIN1257、WIN1258、EUC_CN、EUC_KR、EUC_TW、EUC_JP、EUC_JIS_2004、KOI8R、MULE_INTERNAL、LATIN1、LATIN2、LATIN3、LATIN4、LATIN5、LATIN6、LATIN7、LATIN8、LATIN9、LATIN10、ISO_8859_5、ISO_8859_6、ISO_8859_7、ISO_8859_8、SQL_ASCII ]
	CharacterSet *string `pulumi:"characterSet"`
	// Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
	Description *string `pulumi:"description"`
	// The Id of instance that can run database.
	InstanceId *string `pulumi:"instanceId"`
	// Name of the database requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter
	// and have no more than 64 characters.
	Name *string `pulumi:"name"`
}

type DatabaseState struct {
	// Character set. The value range is limited to the following:
	// - MySQL: [ utf8, gbk, latin1, utf8mb4 ] \(`utf8mb4` only supports versions 5.5 and 5.6\).
	// - SQLServer: [ Chinese_PRC_CI_AS, Chinese_PRC_CS_AS, SQL_Latin1_General_CP1_CI_AS, SQL_Latin1_General_CP1_CS_AS, Chinese_PRC_BIN ]
	// - PostgreSQL: [ KOI8U、UTF8、WIN866、WIN874、WIN1250、WIN1251、WIN1252、WIN1253、WIN1254、WIN1255、WIN1256、WIN1257、WIN1258、EUC_CN、EUC_KR、EUC_TW、EUC_JP、EUC_JIS_2004、KOI8R、MULE_INTERNAL、LATIN1、LATIN2、LATIN3、LATIN4、LATIN5、LATIN6、LATIN7、LATIN8、LATIN9、LATIN10、ISO_8859_5、ISO_8859_6、ISO_8859_7、ISO_8859_8、SQL_ASCII ]
	CharacterSet pulumi.StringPtrInput
	// Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
	Description pulumi.StringPtrInput
	// The Id of instance that can run database.
	InstanceId pulumi.StringPtrInput
	// Name of the database requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter
	// and have no more than 64 characters.
	Name pulumi.StringPtrInput
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	// Character set. The value range is limited to the following:
	// - MySQL: [ utf8, gbk, latin1, utf8mb4 ] \(`utf8mb4` only supports versions 5.5 and 5.6\).
	// - SQLServer: [ Chinese_PRC_CI_AS, Chinese_PRC_CS_AS, SQL_Latin1_General_CP1_CI_AS, SQL_Latin1_General_CP1_CS_AS, Chinese_PRC_BIN ]
	// - PostgreSQL: [ KOI8U、UTF8、WIN866、WIN874、WIN1250、WIN1251、WIN1252、WIN1253、WIN1254、WIN1255、WIN1256、WIN1257、WIN1258、EUC_CN、EUC_KR、EUC_TW、EUC_JP、EUC_JIS_2004、KOI8R、MULE_INTERNAL、LATIN1、LATIN2、LATIN3、LATIN4、LATIN5、LATIN6、LATIN7、LATIN8、LATIN9、LATIN10、ISO_8859_5、ISO_8859_6、ISO_8859_7、ISO_8859_8、SQL_ASCII ]
	CharacterSet *string `pulumi:"characterSet"`
	// Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
	Description *string `pulumi:"description"`
	// The Id of instance that can run database.
	InstanceId string `pulumi:"instanceId"`
	// Name of the database requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter
	// and have no more than 64 characters.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// Character set. The value range is limited to the following:
	// - MySQL: [ utf8, gbk, latin1, utf8mb4 ] \(`utf8mb4` only supports versions 5.5 and 5.6\).
	// - SQLServer: [ Chinese_PRC_CI_AS, Chinese_PRC_CS_AS, SQL_Latin1_General_CP1_CI_AS, SQL_Latin1_General_CP1_CS_AS, Chinese_PRC_BIN ]
	// - PostgreSQL: [ KOI8U、UTF8、WIN866、WIN874、WIN1250、WIN1251、WIN1252、WIN1253、WIN1254、WIN1255、WIN1256、WIN1257、WIN1258、EUC_CN、EUC_KR、EUC_TW、EUC_JP、EUC_JIS_2004、KOI8R、MULE_INTERNAL、LATIN1、LATIN2、LATIN3、LATIN4、LATIN5、LATIN6、LATIN7、LATIN8、LATIN9、LATIN10、ISO_8859_5、ISO_8859_6、ISO_8859_7、ISO_8859_8、SQL_ASCII ]
	CharacterSet pulumi.StringPtrInput
	// Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
	Description pulumi.StringPtrInput
	// The Id of instance that can run database.
	InstanceId pulumi.StringInput
	// Name of the database requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter
	// and have no more than 64 characters.
	Name pulumi.StringPtrInput
}

func (DatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseArgs)(nil)).Elem()
}

type DatabaseInput interface {
	pulumi.Input

	ToDatabaseOutput() DatabaseOutput
	ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput
}

func (*Database) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (i *Database) ToDatabaseOutput() DatabaseOutput {
	return i.ToDatabaseOutputWithContext(context.Background())
}

func (i *Database) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseOutput)
}

// DatabaseArrayInput is an input type that accepts DatabaseArray and DatabaseArrayOutput values.
// You can construct a concrete instance of `DatabaseArrayInput` via:
//
//	DatabaseArray{ DatabaseArgs{...} }
type DatabaseArrayInput interface {
	pulumi.Input

	ToDatabaseArrayOutput() DatabaseArrayOutput
	ToDatabaseArrayOutputWithContext(context.Context) DatabaseArrayOutput
}

type DatabaseArray []DatabaseInput

func (DatabaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Database)(nil)).Elem()
}

func (i DatabaseArray) ToDatabaseArrayOutput() DatabaseArrayOutput {
	return i.ToDatabaseArrayOutputWithContext(context.Background())
}

func (i DatabaseArray) ToDatabaseArrayOutputWithContext(ctx context.Context) DatabaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseArrayOutput)
}

// DatabaseMapInput is an input type that accepts DatabaseMap and DatabaseMapOutput values.
// You can construct a concrete instance of `DatabaseMapInput` via:
//
//	DatabaseMap{ "key": DatabaseArgs{...} }
type DatabaseMapInput interface {
	pulumi.Input

	ToDatabaseMapOutput() DatabaseMapOutput
	ToDatabaseMapOutputWithContext(context.Context) DatabaseMapOutput
}

type DatabaseMap map[string]DatabaseInput

func (DatabaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Database)(nil)).Elem()
}

func (i DatabaseMap) ToDatabaseMapOutput() DatabaseMapOutput {
	return i.ToDatabaseMapOutputWithContext(context.Background())
}

func (i DatabaseMap) ToDatabaseMapOutputWithContext(ctx context.Context) DatabaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseMapOutput)
}

type DatabaseOutput struct{ *pulumi.OutputState }

func (DatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (o DatabaseOutput) ToDatabaseOutput() DatabaseOutput {
	return o
}

func (o DatabaseOutput) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return o
}

// Character set. The value range is limited to the following:
// - MySQL: [ utf8, gbk, latin1, utf8mb4 ] \(`utf8mb4` only supports versions 5.5 and 5.6\).
// - SQLServer: [ Chinese_PRC_CI_AS, Chinese_PRC_CS_AS, SQL_Latin1_General_CP1_CI_AS, SQL_Latin1_General_CP1_CS_AS, Chinese_PRC_BIN ]
// - PostgreSQL: [ KOI8U、UTF8、WIN866、WIN874、WIN1250、WIN1251、WIN1252、WIN1253、WIN1254、WIN1255、WIN1256、WIN1257、WIN1258、EUC_CN、EUC_KR、EUC_TW、EUC_JP、EUC_JIS_2004、KOI8R、MULE_INTERNAL、LATIN1、LATIN2、LATIN3、LATIN4、LATIN5、LATIN6、LATIN7、LATIN8、LATIN9、LATIN10、ISO_8859_5、ISO_8859_6、ISO_8859_7、ISO_8859_8、SQL_ASCII ]
func (o DatabaseOutput) CharacterSet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.CharacterSet }).(pulumi.StringPtrOutput)
}

// Database description. It cannot begin with https://. It must start with a Chinese character or English letter. It can include Chinese and English characters, underlines (_), hyphens (-), and numbers. The length may be 2-256 characters.
func (o DatabaseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The Id of instance that can run database.
func (o DatabaseOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Name of the database requiring a uniqueness check. It may consist of lower case letters, numbers, and underlines, and must start with a letter
// and have no more than 64 characters.
func (o DatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type DatabaseArrayOutput struct{ *pulumi.OutputState }

func (DatabaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Database)(nil)).Elem()
}

func (o DatabaseArrayOutput) ToDatabaseArrayOutput() DatabaseArrayOutput {
	return o
}

func (o DatabaseArrayOutput) ToDatabaseArrayOutputWithContext(ctx context.Context) DatabaseArrayOutput {
	return o
}

func (o DatabaseArrayOutput) Index(i pulumi.IntInput) DatabaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Database {
		return vs[0].([]*Database)[vs[1].(int)]
	}).(DatabaseOutput)
}

type DatabaseMapOutput struct{ *pulumi.OutputState }

func (DatabaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Database)(nil)).Elem()
}

func (o DatabaseMapOutput) ToDatabaseMapOutput() DatabaseMapOutput {
	return o
}

func (o DatabaseMapOutput) ToDatabaseMapOutputWithContext(ctx context.Context) DatabaseMapOutput {
	return o
}

func (o DatabaseMapOutput) MapIndex(k pulumi.StringInput) DatabaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Database {
		return vs[0].(map[string]*Database)[vs[1].(string)]
	}).(DatabaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseInput)(nil)).Elem(), &Database{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseArrayInput)(nil)).Elem(), DatabaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseMapInput)(nil)).Elem(), DatabaseMap{})
	pulumi.RegisterOutputType(DatabaseOutput{})
	pulumi.RegisterOutputType(DatabaseArrayOutput{})
	pulumi.RegisterOutputType(DatabaseMapOutput{})
}