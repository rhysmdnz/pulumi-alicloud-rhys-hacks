// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the apps of the current Alibaba Cloud user.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const dataApigatway = pulumi.output(alicloud.apigateway.getApps({
 *     outputFile: "outapps",
 * }));
 *
 * export const firstAppId = dataApigatway.apps[0].id;
 * ```
 */
export function getApps(args?: GetAppsArgs, opts?: pulumi.InvokeOptions): Promise<GetAppsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("alicloud:apigateway/getApps:getApps", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getApps.
 */
export interface GetAppsArgs {
    /**
     * A list of app IDs.
     */
    ids?: string[];
    /**
     * A regex string to filter apps by name.
     */
    nameRegex?: string;
    outputFile?: string;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getApps.
 */
export interface GetAppsResult {
    /**
     * A list of apps. Each element contains the following attributes:
     */
    readonly apps: outputs.apigateway.GetAppsApp[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of app IDs.
     */
    readonly ids: string[];
    readonly nameRegex?: string;
    /**
     * A list of app names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    readonly tags?: {[key: string]: any};
}

export function getAppsOutput(args?: GetAppsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAppsResult> {
    return pulumi.output(args).apply(a => getApps(a, opts))
}

/**
 * A collection of arguments for invoking getApps.
 */
export interface GetAppsOutputArgs {
    /**
     * A list of app IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter apps by name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: any}>;
}