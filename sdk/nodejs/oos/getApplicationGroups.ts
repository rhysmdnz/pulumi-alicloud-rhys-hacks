// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Oos Application Groups of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.146.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const ids = alicloud.oos.getApplicationGroups({
 *     applicationName: "example_value",
 *     ids: [
 *         "my-ApplicationGroup-1",
 *         "my-ApplicationGroup-2",
 *     ],
 * });
 * export const oosApplicationGroupId1 = ids.then(ids => ids.groups?[0]?.id);
 * const nameRegex = alicloud.oos.getApplicationGroups({
 *     applicationName: "example_value",
 *     nameRegex: "^my-ApplicationGroup",
 * });
 * export const oosApplicationGroupId2 = nameRegex.then(nameRegex => nameRegex.groups?[0]?.id);
 * ```
 */
export function getApplicationGroups(args: GetApplicationGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationGroupsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("alicloud:oos/getApplicationGroups:getApplicationGroups", {
        "applicationName": args.applicationName,
        "deployRegionId": args.deployRegionId,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getApplicationGroups.
 */
export interface GetApplicationGroupsArgs {
    /**
     * The name of the Application.
     */
    applicationName: string;
    /**
     * The region ID of the deployment.
     */
    deployRegionId?: string;
    /**
     * A list of Application Group IDs. Its element value is same as Application Group Name.
     */
    ids?: string[];
    /**
     * A regex string to filter results by Application Group name.
     */
    nameRegex?: string;
    outputFile?: string;
}

/**
 * A collection of values returned by getApplicationGroups.
 */
export interface GetApplicationGroupsResult {
    readonly applicationName: string;
    readonly deployRegionId?: string;
    readonly groups: outputs.oos.GetApplicationGroupsGroup[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly names: string[];
    readonly outputFile?: string;
}

export function getApplicationGroupsOutput(args: GetApplicationGroupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApplicationGroupsResult> {
    return pulumi.output(args).apply(a => getApplicationGroups(a, opts))
}

/**
 * A collection of arguments for invoking getApplicationGroups.
 */
export interface GetApplicationGroupsOutputArgs {
    /**
     * The name of the Application.
     */
    applicationName: pulumi.Input<string>;
    /**
     * The region ID of the deployment.
     */
    deployRegionId?: pulumi.Input<string>;
    /**
     * A list of Application Group IDs. Its element value is same as Application Group Name.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string to filter results by Application Group name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
}