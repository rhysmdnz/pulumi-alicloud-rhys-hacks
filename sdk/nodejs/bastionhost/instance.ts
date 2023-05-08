// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Yundun_bastionhost instance can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:bastionhost/instance:Instance example bastionhost-exampe123456
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
    public static readonly __pulumiType = 'alicloud:bastionhost/instance:Instance';

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
     * The AD auth server of the Instance. See the following `Block adAuthServer`.
     */
    public readonly adAuthServers!: pulumi.Output<outputs.bastionhost.InstanceAdAuthServer[]>;
    /**
     * The bandwidth of Cloud Bastionhost instance. Valid values: 0 to 500. Unit: Mbit/s.
     */
    public readonly bandwidth!: pulumi.Output<string>;
    /**
     * Description of the instance. This name can have a string of 1 to 63 characters.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Whether to Enable the public internet access to a specified Bastionhost instance. The valid values: `true`, `false`.
     */
    public readonly enablePublicAccess!: pulumi.Output<boolean>;
    /**
     * The LDAP auth server of the Instance. See the following `Block ldapAuthServer`.
     */
    public readonly ldapAuthServers!: pulumi.Output<outputs.bastionhost.InstanceLdapAuthServer[]>;
    /**
     * The package type of Cloud Bastionhost instance. You can query more supported types through the [DescribePricingModule](https://help.aliyun.com/document_detail/96469.html).
     */
    public readonly licenseCode!: pulumi.Output<string>;
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * The plan code of Cloud Bastionhost instance. Valid values:
     * - `cloudbastion`: Basic Edition.
     * - `cloudbastionHa`: HA Edition.
     */
    public readonly planCode!: pulumi.Output<string>;
    /**
     * The public IP address that you want to add to the whitelist.
     */
    public readonly publicWhiteLists!: pulumi.Output<string[] | undefined>;
    /**
     * Automatic renewal period. Valid values: `1` to `9`, `12`, `24`, `36`. **NOTE:** The `renewPeriod` is required under the condition that `renewalStatus` is `AutoRenewal`. From version 1.193.0, `renewPeriod` can be modified.
     */
    public readonly renewPeriod!: pulumi.Output<number | undefined>;
    /**
     * The unit of the auto-renewal period. Valid values:  **NOTE:** The `renewalPeriodUnit` is required under the condition that `renewalStatus` is `AutoRenewal`.
     * - `M`: months.
     * - `Y`: years.
     */
    public readonly renewalPeriodUnit!: pulumi.Output<string>;
    /**
     * Automatic renewal status. Valid values: `AutoRenewal`, `ManualRenewal`, `NotRenewal`. From version 1.193.0, `renewalStatus` can be modified.
     */
    public readonly renewalStatus!: pulumi.Output<string>;
    /**
     * The Id of resource group which the Bastionhost Instance belongs. If not set, the resource is created in the default resource group.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    public readonly securityGroupIds!: pulumi.Output<string[]>;
    /**
     * The storage of Cloud Bastionhost instance. Valid values: 0 to 500. Unit: TB.
     */
    public readonly storage!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * VSwitch ID configured to Bastionhost.
     */
    public readonly vswitchId!: pulumi.Output<string>;

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
            resourceInputs["adAuthServers"] = state ? state.adAuthServers : undefined;
            resourceInputs["bandwidth"] = state ? state.bandwidth : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enablePublicAccess"] = state ? state.enablePublicAccess : undefined;
            resourceInputs["ldapAuthServers"] = state ? state.ldapAuthServers : undefined;
            resourceInputs["licenseCode"] = state ? state.licenseCode : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["planCode"] = state ? state.planCode : undefined;
            resourceInputs["publicWhiteLists"] = state ? state.publicWhiteLists : undefined;
            resourceInputs["renewPeriod"] = state ? state.renewPeriod : undefined;
            resourceInputs["renewalPeriodUnit"] = state ? state.renewalPeriodUnit : undefined;
            resourceInputs["renewalStatus"] = state ? state.renewalStatus : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            resourceInputs["storage"] = state ? state.storage : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["vswitchId"] = state ? state.vswitchId : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.bandwidth === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bandwidth'");
            }
            if ((!args || args.description === undefined) && !opts.urn) {
                throw new Error("Missing required property 'description'");
            }
            if ((!args || args.licenseCode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'licenseCode'");
            }
            if ((!args || args.planCode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'planCode'");
            }
            if ((!args || args.securityGroupIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroupIds'");
            }
            if ((!args || args.storage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storage'");
            }
            if ((!args || args.vswitchId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vswitchId'");
            }
            resourceInputs["adAuthServers"] = args ? args.adAuthServers : undefined;
            resourceInputs["bandwidth"] = args ? args.bandwidth : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enablePublicAccess"] = args ? args.enablePublicAccess : undefined;
            resourceInputs["ldapAuthServers"] = args ? args.ldapAuthServers : undefined;
            resourceInputs["licenseCode"] = args ? args.licenseCode : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["planCode"] = args ? args.planCode : undefined;
            resourceInputs["publicWhiteLists"] = args ? args.publicWhiteLists : undefined;
            resourceInputs["renewPeriod"] = args ? args.renewPeriod : undefined;
            resourceInputs["renewalPeriodUnit"] = args ? args.renewalPeriodUnit : undefined;
            resourceInputs["renewalStatus"] = args ? args.renewalStatus : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["storage"] = args ? args.storage : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["vswitchId"] = args ? args.vswitchId : undefined;
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
     * The AD auth server of the Instance. See the following `Block adAuthServer`.
     */
    adAuthServers?: pulumi.Input<pulumi.Input<inputs.bastionhost.InstanceAdAuthServer>[]>;
    /**
     * The bandwidth of Cloud Bastionhost instance. Valid values: 0 to 500. Unit: Mbit/s.
     */
    bandwidth?: pulumi.Input<string>;
    /**
     * Description of the instance. This name can have a string of 1 to 63 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether to Enable the public internet access to a specified Bastionhost instance. The valid values: `true`, `false`.
     */
    enablePublicAccess?: pulumi.Input<boolean>;
    /**
     * The LDAP auth server of the Instance. See the following `Block ldapAuthServer`.
     */
    ldapAuthServers?: pulumi.Input<pulumi.Input<inputs.bastionhost.InstanceLdapAuthServer>[]>;
    /**
     * The package type of Cloud Bastionhost instance. You can query more supported types through the [DescribePricingModule](https://help.aliyun.com/document_detail/96469.html).
     */
    licenseCode?: pulumi.Input<string>;
    period?: pulumi.Input<number>;
    /**
     * The plan code of Cloud Bastionhost instance. Valid values:
     * - `cloudbastion`: Basic Edition.
     * - `cloudbastionHa`: HA Edition.
     */
    planCode?: pulumi.Input<string>;
    /**
     * The public IP address that you want to add to the whitelist.
     */
    publicWhiteLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Automatic renewal period. Valid values: `1` to `9`, `12`, `24`, `36`. **NOTE:** The `renewPeriod` is required under the condition that `renewalStatus` is `AutoRenewal`. From version 1.193.0, `renewPeriod` can be modified.
     */
    renewPeriod?: pulumi.Input<number>;
    /**
     * The unit of the auto-renewal period. Valid values:  **NOTE:** The `renewalPeriodUnit` is required under the condition that `renewalStatus` is `AutoRenewal`.
     * - `M`: months.
     * - `Y`: years.
     */
    renewalPeriodUnit?: pulumi.Input<string>;
    /**
     * Automatic renewal status. Valid values: `AutoRenewal`, `ManualRenewal`, `NotRenewal`. From version 1.193.0, `renewalStatus` can be modified.
     */
    renewalStatus?: pulumi.Input<string>;
    /**
     * The Id of resource group which the Bastionhost Instance belongs. If not set, the resource is created in the default resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The storage of Cloud Bastionhost instance. Valid values: 0 to 500. Unit: TB.
     */
    storage?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * VSwitch ID configured to Bastionhost.
     */
    vswitchId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * The AD auth server of the Instance. See the following `Block adAuthServer`.
     */
    adAuthServers?: pulumi.Input<pulumi.Input<inputs.bastionhost.InstanceAdAuthServer>[]>;
    /**
     * The bandwidth of Cloud Bastionhost instance. Valid values: 0 to 500. Unit: Mbit/s.
     */
    bandwidth: pulumi.Input<string>;
    /**
     * Description of the instance. This name can have a string of 1 to 63 characters.
     */
    description: pulumi.Input<string>;
    /**
     * Whether to Enable the public internet access to a specified Bastionhost instance. The valid values: `true`, `false`.
     */
    enablePublicAccess?: pulumi.Input<boolean>;
    /**
     * The LDAP auth server of the Instance. See the following `Block ldapAuthServer`.
     */
    ldapAuthServers?: pulumi.Input<pulumi.Input<inputs.bastionhost.InstanceLdapAuthServer>[]>;
    /**
     * The package type of Cloud Bastionhost instance. You can query more supported types through the [DescribePricingModule](https://help.aliyun.com/document_detail/96469.html).
     */
    licenseCode: pulumi.Input<string>;
    period?: pulumi.Input<number>;
    /**
     * The plan code of Cloud Bastionhost instance. Valid values:
     * - `cloudbastion`: Basic Edition.
     * - `cloudbastionHa`: HA Edition.
     */
    planCode: pulumi.Input<string>;
    /**
     * The public IP address that you want to add to the whitelist.
     */
    publicWhiteLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Automatic renewal period. Valid values: `1` to `9`, `12`, `24`, `36`. **NOTE:** The `renewPeriod` is required under the condition that `renewalStatus` is `AutoRenewal`. From version 1.193.0, `renewPeriod` can be modified.
     */
    renewPeriod?: pulumi.Input<number>;
    /**
     * The unit of the auto-renewal period. Valid values:  **NOTE:** The `renewalPeriodUnit` is required under the condition that `renewalStatus` is `AutoRenewal`.
     * - `M`: months.
     * - `Y`: years.
     */
    renewalPeriodUnit?: pulumi.Input<string>;
    /**
     * Automatic renewal status. Valid values: `AutoRenewal`, `ManualRenewal`, `NotRenewal`. From version 1.193.0, `renewalStatus` can be modified.
     */
    renewalStatus?: pulumi.Input<string>;
    /**
     * The Id of resource group which the Bastionhost Instance belongs. If not set, the resource is created in the default resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    securityGroupIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The storage of Cloud Bastionhost instance. Valid values: 0 to 500. Unit: TB.
     */
    storage: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * VSwitch ID configured to Bastionhost.
     */
    vswitchId: pulumi.Input<string>;
}
