// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const foo = new alicloud.vpn.CustomerGateway("foo", {
 *     description: "vpnCgwDescriptionExample",
 *     ipAddress: "43.104.22.228",
 * });
 * ```
 *
 * ## Import
 *
 * VPN customer gateway can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:vpn/customerGateway:CustomerGateway example cgw-abc123456
 * ```
 */
export class CustomerGateway extends pulumi.CustomResource {
    /**
     * Get an existing CustomerGateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CustomerGatewayState, opts?: pulumi.CustomResourceOptions): CustomerGateway {
        return new CustomerGateway(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:vpn/customerGateway:CustomerGateway';

    /**
     * Returns true if the given object is an instance of CustomerGateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomerGateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomerGateway.__pulumiType;
    }

    /**
     * The autonomous system number of the gateway device in the data center. The `asn` is a 4-byte number. You can enter the number in two segments and separate the first 16 bits from the following 16 bits with a period (.). Enter the number in each segment in the decimal format.
     */
    public readonly asn!: pulumi.Output<string | undefined>;
    /**
     * The description of the VPN customer gateway instance.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The IP address of the customer gateway.
     */
    public readonly ipAddress!: pulumi.Output<string>;
    /**
     * The name of the VPN customer gateway. Defaults to null.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a CustomerGateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomerGatewayArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomerGatewayArgs | CustomerGatewayState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CustomerGatewayState | undefined;
            resourceInputs["asn"] = state ? state.asn : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["ipAddress"] = state ? state.ipAddress : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as CustomerGatewayArgs | undefined;
            if ((!args || args.ipAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipAddress'");
            }
            resourceInputs["asn"] = args ? args.asn : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["ipAddress"] = args ? args.ipAddress : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CustomerGateway.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CustomerGateway resources.
 */
export interface CustomerGatewayState {
    /**
     * The autonomous system number of the gateway device in the data center. The `asn` is a 4-byte number. You can enter the number in two segments and separate the first 16 bits from the following 16 bits with a period (.). Enter the number in each segment in the decimal format.
     */
    asn?: pulumi.Input<string>;
    /**
     * The description of the VPN customer gateway instance.
     */
    description?: pulumi.Input<string>;
    /**
     * The IP address of the customer gateway.
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * The name of the VPN customer gateway. Defaults to null.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CustomerGateway resource.
 */
export interface CustomerGatewayArgs {
    /**
     * The autonomous system number of the gateway device in the data center. The `asn` is a 4-byte number. You can enter the number in two segments and separate the first 16 bits from the following 16 bits with a period (.). Enter the number in each segment in the decimal format.
     */
    asn?: pulumi.Input<string>;
    /**
     * The description of the VPN customer gateway instance.
     */
    description?: pulumi.Input<string>;
    /**
     * The IP address of the customer gateway.
     */
    ipAddress: pulumi.Input<string>;
    /**
     * The name of the VPN customer gateway. Defaults to null.
     */
    name?: pulumi.Input<string>;
}