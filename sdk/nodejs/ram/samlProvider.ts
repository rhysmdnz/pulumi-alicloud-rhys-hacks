// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a RAM SAML Provider resource.
 *
 * For information about RAM SAML Provider and how to use it, see [What is SAML Provider](https://www.alibabacloud.com/help/doc-detail/186846.htm).
 *
 * > **NOTE:** Available in v1.114.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const example = new alicloud.ram.SamlProvider("example", {
 *     description: "For Terraform Test",
 *     encodedsamlMetadataDocument: "your encodedsaml metadata document",
 *     samlProviderName: "tf-testAcc",
 * });
 * ```
 *
 * ## Import
 *
 * RAM SAML Provider can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:ram/samlProvider:SamlProvider example <saml_provider_name>
 * ```
 */
export class SamlProvider extends pulumi.CustomResource {
    /**
     * Get an existing SamlProvider resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SamlProviderState, opts?: pulumi.CustomResourceOptions): SamlProvider {
        return new SamlProvider(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:ram/samlProvider:SamlProvider';

    /**
     * Returns true if the given object is an instance of SamlProvider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SamlProvider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SamlProvider.__pulumiType;
    }

    /**
     * The Alibaba Cloud Resource Name (ARN) of the IdP.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The description of SAML Provider.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The metadata file, which is Base64 encoded. The file is provided by an IdP that supports SAML 2.0.
     */
    public readonly encodedsamlMetadataDocument!: pulumi.Output<string | undefined>;
    /**
     * The name of SAML Provider.
     */
    public readonly samlProviderName!: pulumi.Output<string>;
    /**
     * The update time.
     */
    public /*out*/ readonly updateDate!: pulumi.Output<string>;

    /**
     * Create a SamlProvider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SamlProviderArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SamlProviderArgs | SamlProviderState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SamlProviderState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["encodedsamlMetadataDocument"] = state ? state.encodedsamlMetadataDocument : undefined;
            resourceInputs["samlProviderName"] = state ? state.samlProviderName : undefined;
            resourceInputs["updateDate"] = state ? state.updateDate : undefined;
        } else {
            const args = argsOrState as SamlProviderArgs | undefined;
            if ((!args || args.samlProviderName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'samlProviderName'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["encodedsamlMetadataDocument"] = args ? args.encodedsamlMetadataDocument : undefined;
            resourceInputs["samlProviderName"] = args ? args.samlProviderName : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["updateDate"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SamlProvider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SamlProvider resources.
 */
export interface SamlProviderState {
    /**
     * The Alibaba Cloud Resource Name (ARN) of the IdP.
     */
    arn?: pulumi.Input<string>;
    /**
     * The description of SAML Provider.
     */
    description?: pulumi.Input<string>;
    /**
     * The metadata file, which is Base64 encoded. The file is provided by an IdP that supports SAML 2.0.
     */
    encodedsamlMetadataDocument?: pulumi.Input<string>;
    /**
     * The name of SAML Provider.
     */
    samlProviderName?: pulumi.Input<string>;
    /**
     * The update time.
     */
    updateDate?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SamlProvider resource.
 */
export interface SamlProviderArgs {
    /**
     * The description of SAML Provider.
     */
    description?: pulumi.Input<string>;
    /**
     * The metadata file, which is Base64 encoded. The file is provided by an IdP that supports SAML 2.0.
     */
    encodedsamlMetadataDocument?: pulumi.Input<string>;
    /**
     * The name of SAML Provider.
     */
    samlProviderName: pulumi.Input<string>;
}