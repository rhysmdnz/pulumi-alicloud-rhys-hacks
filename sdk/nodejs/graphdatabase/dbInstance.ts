// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a Graph Database Db Instance resource.
 *
 * For information about Graph Database Db Instance and how to use it, see [What is Db Instance](https://help.aliyun.com/document_detail/102865.html).
 *
 * > **NOTE:** Available in v1.136.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.graphdatabase.DbInstance("example", {
 *     dbInstanceCategory: "HA",
 *     dbInstanceDescription: "example_value",
 *     dbInstanceNetworkType: "vpc",
 *     dbInstanceStorageType: "cloud_ssd",
 *     dbNodeClass: "gdb.r.2xlarge",
 *     dbNodeStorage: Number.parseFloat("example_value"),
 *     dbVersion: "1.0",
 *     paymentType: "PayAsYouGo",
 * });
 * ```
 *
 * ## Import
 *
 * Graph Database Db Instance can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:graphdatabase/dbInstance:DbInstance example <id>
 * ```
 */
export class DbInstance extends pulumi.CustomResource {
    /**
     * Get an existing DbInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DbInstanceState, opts?: pulumi.CustomResourceOptions): DbInstance {
        return new DbInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:graphdatabase/dbInstance:DbInstance';

    /**
     * Returns true if the given object is an instance of DbInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DbInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DbInstance.__pulumiType;
    }

    /**
     * (Available in 1.196.0+)  The connection string of the instance.
     */
    public /*out*/ readonly connectionString!: pulumi.Output<string>;
    /**
     * The category of the db instance. Valid values: `HA`, `SINGLE`(Available in 1.173.0+).
     */
    public readonly dbInstanceCategory!: pulumi.Output<string>;
    /**
     * According to the practical example or notes.
     */
    public readonly dbInstanceDescription!: pulumi.Output<string | undefined>;
    /**
     * IP ADDRESS whitelist for the instance group list. See the following `Block dbInstanceIpArray`.
     */
    public readonly dbInstanceIpArrays!: pulumi.Output<outputs.graphdatabase.DbInstanceDbInstanceIpArray[]>;
    /**
     * The network type of the db instance. Valid values: `vpc`.
     */
    public readonly dbInstanceNetworkType!: pulumi.Output<string>;
    /**
     * Disk storage type. Valid values: `cloudEssd`, `cloudSsd`. Modification is not supported.
     */
    public readonly dbInstanceStorageType!: pulumi.Output<string>;
    /**
     * The class of the db node. Valid values: `gdb.r.xlarge`, `gdb.r.2xlarge`, `gdb.r.4xlarge`, `gdb.r.8xlarge`, `gdb.r.16xlarge`, `gdb.r.xlarge_basic`, `gdb.r.2xlarge_basic`, `gdb.r.4xlarge_basic`, `gdb.r.8xlarge_basic`, `gdb.r.16xlarge_basic`.
     */
    public readonly dbNodeClass!: pulumi.Output<string>;
    /**
     * Instance storage space, which is measured in GB.
     */
    public readonly dbNodeStorage!: pulumi.Output<number>;
    /**
     * Kernel Version. Valid values: `1.0` or `1.0-OpenCypher`. `1.0`: represented as gremlin, `1.0-OpenCypher`: said opencypher.
     */
    public readonly dbVersion!: pulumi.Output<string>;
    /**
     * The paymen type of the resource. Valid values: `PayAsYouGo`.
     */
    public readonly paymentType!: pulumi.Output<string>;
    /**
     * (Available in 1.196.0+) The connection port of the instance.
     */
    public /*out*/ readonly port!: pulumi.Output<string>;
    /**
     * Instance status. Value range: `Creating`, `Running`, `Deleting`, `Rebooting`, `DBInstanceClassChanging`, `NetAddressCreating` and `NetAddressDeleting`.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * ID of the VPC.
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * The ID of attaching vswitch to instance.
     */
    public readonly vswitchId!: pulumi.Output<string>;
    /**
     * The zone ID of the resource.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a DbInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DbInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DbInstanceArgs | DbInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DbInstanceState | undefined;
            resourceInputs["connectionString"] = state ? state.connectionString : undefined;
            resourceInputs["dbInstanceCategory"] = state ? state.dbInstanceCategory : undefined;
            resourceInputs["dbInstanceDescription"] = state ? state.dbInstanceDescription : undefined;
            resourceInputs["dbInstanceIpArrays"] = state ? state.dbInstanceIpArrays : undefined;
            resourceInputs["dbInstanceNetworkType"] = state ? state.dbInstanceNetworkType : undefined;
            resourceInputs["dbInstanceStorageType"] = state ? state.dbInstanceStorageType : undefined;
            resourceInputs["dbNodeClass"] = state ? state.dbNodeClass : undefined;
            resourceInputs["dbNodeStorage"] = state ? state.dbNodeStorage : undefined;
            resourceInputs["dbVersion"] = state ? state.dbVersion : undefined;
            resourceInputs["paymentType"] = state ? state.paymentType : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["vswitchId"] = state ? state.vswitchId : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as DbInstanceArgs | undefined;
            if ((!args || args.dbInstanceCategory === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbInstanceCategory'");
            }
            if ((!args || args.dbInstanceNetworkType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbInstanceNetworkType'");
            }
            if ((!args || args.dbInstanceStorageType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbInstanceStorageType'");
            }
            if ((!args || args.dbNodeClass === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbNodeClass'");
            }
            if ((!args || args.dbNodeStorage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbNodeStorage'");
            }
            if ((!args || args.dbVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbVersion'");
            }
            if ((!args || args.paymentType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'paymentType'");
            }
            resourceInputs["dbInstanceCategory"] = args ? args.dbInstanceCategory : undefined;
            resourceInputs["dbInstanceDescription"] = args ? args.dbInstanceDescription : undefined;
            resourceInputs["dbInstanceIpArrays"] = args ? args.dbInstanceIpArrays : undefined;
            resourceInputs["dbInstanceNetworkType"] = args ? args.dbInstanceNetworkType : undefined;
            resourceInputs["dbInstanceStorageType"] = args ? args.dbInstanceStorageType : undefined;
            resourceInputs["dbNodeClass"] = args ? args.dbNodeClass : undefined;
            resourceInputs["dbNodeStorage"] = args ? args.dbNodeStorage : undefined;
            resourceInputs["dbVersion"] = args ? args.dbVersion : undefined;
            resourceInputs["paymentType"] = args ? args.paymentType : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["vswitchId"] = args ? args.vswitchId : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
            resourceInputs["connectionString"] = undefined /*out*/;
            resourceInputs["port"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DbInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DbInstance resources.
 */
export interface DbInstanceState {
    /**
     * (Available in 1.196.0+)  The connection string of the instance.
     */
    connectionString?: pulumi.Input<string>;
    /**
     * The category of the db instance. Valid values: `HA`, `SINGLE`(Available in 1.173.0+).
     */
    dbInstanceCategory?: pulumi.Input<string>;
    /**
     * According to the practical example or notes.
     */
    dbInstanceDescription?: pulumi.Input<string>;
    /**
     * IP ADDRESS whitelist for the instance group list. See the following `Block dbInstanceIpArray`.
     */
    dbInstanceIpArrays?: pulumi.Input<pulumi.Input<inputs.graphdatabase.DbInstanceDbInstanceIpArray>[]>;
    /**
     * The network type of the db instance. Valid values: `vpc`.
     */
    dbInstanceNetworkType?: pulumi.Input<string>;
    /**
     * Disk storage type. Valid values: `cloudEssd`, `cloudSsd`. Modification is not supported.
     */
    dbInstanceStorageType?: pulumi.Input<string>;
    /**
     * The class of the db node. Valid values: `gdb.r.xlarge`, `gdb.r.2xlarge`, `gdb.r.4xlarge`, `gdb.r.8xlarge`, `gdb.r.16xlarge`, `gdb.r.xlarge_basic`, `gdb.r.2xlarge_basic`, `gdb.r.4xlarge_basic`, `gdb.r.8xlarge_basic`, `gdb.r.16xlarge_basic`.
     */
    dbNodeClass?: pulumi.Input<string>;
    /**
     * Instance storage space, which is measured in GB.
     */
    dbNodeStorage?: pulumi.Input<number>;
    /**
     * Kernel Version. Valid values: `1.0` or `1.0-OpenCypher`. `1.0`: represented as gremlin, `1.0-OpenCypher`: said opencypher.
     */
    dbVersion?: pulumi.Input<string>;
    /**
     * The paymen type of the resource. Valid values: `PayAsYouGo`.
     */
    paymentType?: pulumi.Input<string>;
    /**
     * (Available in 1.196.0+) The connection port of the instance.
     */
    port?: pulumi.Input<string>;
    /**
     * Instance status. Value range: `Creating`, `Running`, `Deleting`, `Rebooting`, `DBInstanceClassChanging`, `NetAddressCreating` and `NetAddressDeleting`.
     */
    status?: pulumi.Input<string>;
    /**
     * ID of the VPC.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The ID of attaching vswitch to instance.
     */
    vswitchId?: pulumi.Input<string>;
    /**
     * The zone ID of the resource.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DbInstance resource.
 */
export interface DbInstanceArgs {
    /**
     * The category of the db instance. Valid values: `HA`, `SINGLE`(Available in 1.173.0+).
     */
    dbInstanceCategory: pulumi.Input<string>;
    /**
     * According to the practical example or notes.
     */
    dbInstanceDescription?: pulumi.Input<string>;
    /**
     * IP ADDRESS whitelist for the instance group list. See the following `Block dbInstanceIpArray`.
     */
    dbInstanceIpArrays?: pulumi.Input<pulumi.Input<inputs.graphdatabase.DbInstanceDbInstanceIpArray>[]>;
    /**
     * The network type of the db instance. Valid values: `vpc`.
     */
    dbInstanceNetworkType: pulumi.Input<string>;
    /**
     * Disk storage type. Valid values: `cloudEssd`, `cloudSsd`. Modification is not supported.
     */
    dbInstanceStorageType: pulumi.Input<string>;
    /**
     * The class of the db node. Valid values: `gdb.r.xlarge`, `gdb.r.2xlarge`, `gdb.r.4xlarge`, `gdb.r.8xlarge`, `gdb.r.16xlarge`, `gdb.r.xlarge_basic`, `gdb.r.2xlarge_basic`, `gdb.r.4xlarge_basic`, `gdb.r.8xlarge_basic`, `gdb.r.16xlarge_basic`.
     */
    dbNodeClass: pulumi.Input<string>;
    /**
     * Instance storage space, which is measured in GB.
     */
    dbNodeStorage: pulumi.Input<number>;
    /**
     * Kernel Version. Valid values: `1.0` or `1.0-OpenCypher`. `1.0`: represented as gremlin, `1.0-OpenCypher`: said opencypher.
     */
    dbVersion: pulumi.Input<string>;
    /**
     * The paymen type of the resource. Valid values: `PayAsYouGo`.
     */
    paymentType: pulumi.Input<string>;
    /**
     * ID of the VPC.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The ID of attaching vswitch to instance.
     */
    vswitchId?: pulumi.Input<string>;
    /**
     * The zone ID of the resource.
     */
    zoneId?: pulumi.Input<string>;
}
