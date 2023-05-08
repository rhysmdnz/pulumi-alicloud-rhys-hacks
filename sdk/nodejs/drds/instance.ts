// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Distributed Relational Database Service (DRDS) is a lightweight (stateless), flexible, stable, and efficient middleware product independently developed by Alibaba Group to resolve scalability issues with single-host relational databases.
 * With its compatibility with MySQL protocols and syntaxes, DRDS enables database/table sharding, smooth scaling, configuration upgrade/downgrade,
 * transparent read/write splitting, and distributed transactions, providing O&M capabilities for distributed databases throughout their entire lifecycle.
 *
 * For information about DRDS and how to use it, see [What is DRDS](https://www.alibabacloud.com/help/product/29657.htm).
 *
 * > **NOTE:** At present, DRDS instance only can be supported in the regions: cn-shenzhen, cn-beijing, cn-hangzhou, cn-hongkong, cn-qingdao, ap-southeast-1.
 *
 * > **NOTE:** Currently, this resource only support `Domestic Site Account`.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultInstance = new alicloud.drds.Instance("default", {
 *     description: "drds instance",
 *     instanceChargeType: "PostPaid",
 *     instanceSeries: "drds.sn1.4c8g",
 *     specification: "drds.sn1.4c8g.8C16G",
 *     vswitchId: "vsw-bp1jlu3swk8rq2yoi40ey",
 *     zoneId: "cn-hangzhou-e",
 * });
 * ```
 *
 * ## Import
 *
 * Distributed Relational Database Service (DRDS) can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:drds/instance:Instance example drds-abc123456
 * ```
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:drds/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * (Available in 1.196.0+) The connection string of the DRDS instance.
     */
    public /*out*/ readonly connectionString!: pulumi.Output<string>;
    /**
     * Description of the DRDS instance, This description can have a string of 2 to 256 characters.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`.
     */
    public readonly instanceChargeType!: pulumi.Output<string | undefined>;
    /**
     * The parameter of the instance series. **NOTE:**  `drds.sn1.4c8g`,`drds.sn1.8c16g`,`drds.sn1.16c32g`,`drds.sn1.32c64g` are no longer supported. Valid values:
     * - `drds.sn2.4c16g` Starter Edition.
     * - `drds.sn2.8c32g` Standard Edition.
     * - `drds.sn2.16c64g` Enterprise Edition.
     */
    public readonly instanceSeries!: pulumi.Output<string>;
    /**
     * The MySQL version supported by the instance, with the following range of values. `5`: Fully compatible with MySQL 5.x (default) `8`: Fully compatible with MySQL 8.0. This parameter takes effect when the primary instance is created, and the read-only instance has the same MySQL version as the primary instance by default.
     */
    public readonly mysqlVersion!: pulumi.Output<number>;
    /**
     * (Available in 1.196.0+) The connection port of the DRDS instance.
     */
    public /*out*/ readonly port!: pulumi.Output<string>;
    /**
     * User-defined DRDS instance specification. Value range:
     * - `drds.sn1.4c8g` for DRDS instance Starter version;
     * - value range : `drds.sn1.4c8g.8c16g`, `drds.sn1.4c8g.16c32g`, `drds.sn1.4c8g.32c64g`, `drds.sn1.4c8g.64c128g`
     * - `drds.sn1.8c16g` for DRDS instance Standard edition;
     * - value range : `drds.sn1.8c16g.16c32g`, `drds.sn1.8c16g.32c64g`, `drds.sn1.8c16g.64c128g`
     * - `drds.sn1.16c32g` for DRDS instance Enterprise Edition;
     * - value range : `drds.sn1.16c32g.32c64g`, `drds.sn1.16c32g.64c128g`
     * - `drds.sn1.32c64g` for DRDS instance Extreme Edition;
     * - value range : `drds.sn1.32c64g.128c256g`
     */
    public readonly specification!: pulumi.Output<string>;
    /**
     * The id of the VPC.
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * The VSwitch ID to launch in.
     */
    public readonly vswitchId!: pulumi.Output<string>;
    /**
     * The Zone to launch the DRDS instance.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            resourceInputs["connectionString"] = state ? state.connectionString : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["instanceChargeType"] = state ? state.instanceChargeType : undefined;
            resourceInputs["instanceSeries"] = state ? state.instanceSeries : undefined;
            resourceInputs["mysqlVersion"] = state ? state.mysqlVersion : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["specification"] = state ? state.specification : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["vswitchId"] = state ? state.vswitchId : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.instanceSeries === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceSeries'");
            }
            if ((!args || args.specification === undefined) && !opts.urn) {
                throw new Error("Missing required property 'specification'");
            }
            if ((!args || args.vswitchId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vswitchId'");
            }
            if ((!args || args.zoneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["instanceChargeType"] = args ? args.instanceChargeType : undefined;
            resourceInputs["instanceSeries"] = args ? args.instanceSeries : undefined;
            resourceInputs["mysqlVersion"] = args ? args.mysqlVersion : undefined;
            resourceInputs["specification"] = args ? args.specification : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["vswitchId"] = args ? args.vswitchId : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
            resourceInputs["connectionString"] = undefined /*out*/;
            resourceInputs["port"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * (Available in 1.196.0+) The connection string of the DRDS instance.
     */
    connectionString?: pulumi.Input<string>;
    /**
     * Description of the DRDS instance, This description can have a string of 2 to 256 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * The parameter of the instance series. **NOTE:**  `drds.sn1.4c8g`,`drds.sn1.8c16g`,`drds.sn1.16c32g`,`drds.sn1.32c64g` are no longer supported. Valid values:
     * - `drds.sn2.4c16g` Starter Edition.
     * - `drds.sn2.8c32g` Standard Edition.
     * - `drds.sn2.16c64g` Enterprise Edition.
     */
    instanceSeries?: pulumi.Input<string>;
    /**
     * The MySQL version supported by the instance, with the following range of values. `5`: Fully compatible with MySQL 5.x (default) `8`: Fully compatible with MySQL 8.0. This parameter takes effect when the primary instance is created, and the read-only instance has the same MySQL version as the primary instance by default.
     */
    mysqlVersion?: pulumi.Input<number>;
    /**
     * (Available in 1.196.0+) The connection port of the DRDS instance.
     */
    port?: pulumi.Input<string>;
    /**
     * User-defined DRDS instance specification. Value range:
     * - `drds.sn1.4c8g` for DRDS instance Starter version;
     * - value range : `drds.sn1.4c8g.8c16g`, `drds.sn1.4c8g.16c32g`, `drds.sn1.4c8g.32c64g`, `drds.sn1.4c8g.64c128g`
     * - `drds.sn1.8c16g` for DRDS instance Standard edition;
     * - value range : `drds.sn1.8c16g.16c32g`, `drds.sn1.8c16g.32c64g`, `drds.sn1.8c16g.64c128g`
     * - `drds.sn1.16c32g` for DRDS instance Enterprise Edition;
     * - value range : `drds.sn1.16c32g.32c64g`, `drds.sn1.16c32g.64c128g`
     * - `drds.sn1.32c64g` for DRDS instance Extreme Edition;
     * - value range : `drds.sn1.32c64g.128c256g`
     */
    specification?: pulumi.Input<string>;
    /**
     * The id of the VPC.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The VSwitch ID to launch in.
     */
    vswitchId?: pulumi.Input<string>;
    /**
     * The Zone to launch the DRDS instance.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Description of the DRDS instance, This description can have a string of 2 to 256 characters.
     */
    description: pulumi.Input<string>;
    /**
     * Valid values are `PrePaid`, `PostPaid`, Default to `PostPaid`.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * The parameter of the instance series. **NOTE:**  `drds.sn1.4c8g`,`drds.sn1.8c16g`,`drds.sn1.16c32g`,`drds.sn1.32c64g` are no longer supported. Valid values:
     * - `drds.sn2.4c16g` Starter Edition.
     * - `drds.sn2.8c32g` Standard Edition.
     * - `drds.sn2.16c64g` Enterprise Edition.
     */
    instanceSeries: pulumi.Input<string>;
    /**
     * The MySQL version supported by the instance, with the following range of values. `5`: Fully compatible with MySQL 5.x (default) `8`: Fully compatible with MySQL 8.0. This parameter takes effect when the primary instance is created, and the read-only instance has the same MySQL version as the primary instance by default.
     */
    mysqlVersion?: pulumi.Input<number>;
    /**
     * User-defined DRDS instance specification. Value range:
     * - `drds.sn1.4c8g` for DRDS instance Starter version;
     * - value range : `drds.sn1.4c8g.8c16g`, `drds.sn1.4c8g.16c32g`, `drds.sn1.4c8g.32c64g`, `drds.sn1.4c8g.64c128g`
     * - `drds.sn1.8c16g` for DRDS instance Standard edition;
     * - value range : `drds.sn1.8c16g.16c32g`, `drds.sn1.8c16g.32c64g`, `drds.sn1.8c16g.64c128g`
     * - `drds.sn1.16c32g` for DRDS instance Enterprise Edition;
     * - value range : `drds.sn1.16c32g.32c64g`, `drds.sn1.16c32g.64c128g`
     * - `drds.sn1.32c64g` for DRDS instance Extreme Edition;
     * - value range : `drds.sn1.32c64g.128c256g`
     */
    specification: pulumi.Input<string>;
    /**
     * The id of the VPC.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The VSwitch ID to launch in.
     */
    vswitchId: pulumi.Input<string>;
    /**
     * The Zone to launch the DRDS instance.
     */
    zoneId: pulumi.Input<string>;
}
