// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides Sag Acls available to the user.
 *
 * > **NOTE:** Available in 1.60.0+
 *
 * > **NOTE:** Only the following regions support create Cloud Connect Network. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const defaultAcls = alicloud_sag_acls_default.id.apply(id => alicloud.sag.getAcls({
 *     ids: [id],
 *     nameRegex: "^tf-testAcc.*",
 * }));
 * const defaultAcl = new alicloud.rocketmq.Acl("default", {});
 * ```
 */
export function getAcls(args?: GetAclsArgs, opts?: pulumi.InvokeOptions): Promise<GetAclsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("alicloud:sag/getAcls:getAcls", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getAcls.
 */
export interface GetAclsArgs {
    /**
     * A list of Sag Acl IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter Sag Acl instances by name.
     */
    nameRegex?: string;
    outputFile?: string;
}

/**
 * A collection of values returned by getAcls.
 */
export interface GetAclsResult {
    /**
     * A list of Sag Acls. Each element contains the following attributes:
     */
    readonly acls: outputs.sag.GetAclsAcl[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of Sag Acl IDs.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of Sag Acls names.
     */
    readonly names: string[];
    readonly outputFile?: string;
}

export function getAclsOutput(args?: GetAclsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAclsResult> {
    return pulumi.output(args).apply(a => getAcls(a, opts))
}

/**
 * A collection of arguments for invoking getAcls.
 */
export interface GetAclsOutputArgs {
    /**
     * A list of Sag Acl IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter Sag Acl instances by name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
}