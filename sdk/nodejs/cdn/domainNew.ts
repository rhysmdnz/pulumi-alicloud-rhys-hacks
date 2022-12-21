// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a CDN Accelerated Domain resource. This resource is based on CDN's new version OpenAPI.
 *
 * For information about Cdn Domain New and how to use it, see [Add a domain](https://www.alibabacloud.com/help/doc-detail/91176.html).
 *
 * > **NOTE:** Available in v1.34.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * // Create a new Domain.
 * const domain = new alicloud.cdn.DomainNew("domain", {
 *     cdnType: "web",
 *     domainName: "terraform.test.com",
 *     scope: "overseas",
 *     sources: [{
 *         content: "1.1.1.1",
 *         port: 80,
 *         priority: 20,
 *         type: "ipaddr",
 *         weight: 10,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * CDN domain can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:cdn/domainNew:DomainNew example xxxx.com
 * ```
 */
export class DomainNew extends pulumi.CustomResource {
    /**
     * Get an existing DomainNew resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainNewState, opts?: pulumi.CustomResourceOptions): DomainNew {
        return new DomainNew(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:cdn/domainNew:DomainNew';

    /**
     * Returns true if the given object is an instance of DomainNew.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DomainNew {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DomainNew.__pulumiType;
    }

    /**
     * Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
     */
    public readonly cdnType!: pulumi.Output<string>;
    /**
     * Certificate config of the accelerated domain. It's a list and consist of at most 1 item.
     */
    public readonly certificateConfig!: pulumi.Output<outputs.cdn.DomainNewCertificateConfig>;
    /**
     * (Available in v1.90.0+) The CNAME of the CDN domain.
     */
    public /*out*/ readonly cname!: pulumi.Output<string>;
    /**
     * Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * Resource group ID.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter's setting is valid Only for the international users and domestic L3 and above users .
     */
    public readonly scope!: pulumi.Output<string>;
    /**
     * The source address list of the accelerated domain. Defaults to null. See Block Sources.
     */
    public readonly sources!: pulumi.Output<outputs.cdn.DomainNewSource[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a DomainNew resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainNewArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainNewArgs | DomainNewState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DomainNewState | undefined;
            resourceInputs["cdnType"] = state ? state.cdnType : undefined;
            resourceInputs["certificateConfig"] = state ? state.certificateConfig : undefined;
            resourceInputs["cname"] = state ? state.cname : undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["scope"] = state ? state.scope : undefined;
            resourceInputs["sources"] = state ? state.sources : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as DomainNewArgs | undefined;
            if ((!args || args.cdnType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cdnType'");
            }
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            if ((!args || args.sources === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sources'");
            }
            resourceInputs["cdnType"] = args ? args.cdnType : undefined;
            resourceInputs["certificateConfig"] = args ? args.certificateConfig : undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["scope"] = args ? args.scope : undefined;
            resourceInputs["sources"] = args ? args.sources : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["cname"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DomainNew.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DomainNew resources.
 */
export interface DomainNewState {
    /**
     * Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
     */
    cdnType?: pulumi.Input<string>;
    /**
     * Certificate config of the accelerated domain. It's a list and consist of at most 1 item.
     */
    certificateConfig?: pulumi.Input<inputs.cdn.DomainNewCertificateConfig>;
    /**
     * (Available in v1.90.0+) The CNAME of the CDN domain.
     */
    cname?: pulumi.Input<string>;
    /**
     * Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     */
    domainName?: pulumi.Input<string>;
    /**
     * Resource group ID.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter's setting is valid Only for the international users and domestic L3 and above users .
     */
    scope?: pulumi.Input<string>;
    /**
     * The source address list of the accelerated domain. Defaults to null. See Block Sources.
     */
    sources?: pulumi.Input<pulumi.Input<inputs.cdn.DomainNewSource>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a DomainNew resource.
 */
export interface DomainNewArgs {
    /**
     * Cdn type of the accelerated domain. Valid values are `web`, `download`, `video`.
     */
    cdnType: pulumi.Input<string>;
    /**
     * Certificate config of the accelerated domain. It's a list and consist of at most 1 item.
     */
    certificateConfig?: pulumi.Input<inputs.cdn.DomainNewCertificateConfig>;
    /**
     * Name of the accelerated domain. This name without suffix can have a string of 1 to 63 characters, must contain only alphanumeric characters or "-", and must not begin or end with "-", and "-" must not in the 3th and 4th character positions at the same time. Suffix `.sh` and `.tel` are not supported.
     */
    domainName: pulumi.Input<string>;
    /**
     * Resource group ID.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * Scope of the accelerated domain. Valid values are `domestic`, `overseas`, `global`. Default value is `domestic`. This parameter's setting is valid Only for the international users and domestic L3 and above users .
     */
    scope?: pulumi.Input<string>;
    /**
     * The source address list of the accelerated domain. Defaults to null. See Block Sources.
     */
    sources: pulumi.Input<pulumi.Input<inputs.cdn.DomainNewSource>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}