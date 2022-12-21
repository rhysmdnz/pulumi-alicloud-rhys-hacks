// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides CEN Regional Route Entries available to the user.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const entry = pulumi.output(alicloud.cen.getRegionRouteEntries({
 *     instanceId: "cen-id1",
 *     regionId: "cn-beijing",
 * }));
 *
 * export const firstRegionRouteEntriesRouteEntryCidrBlock = entry.entries[0].cidrBlock;
 * ```
 */
export function getRegionRouteEntries(args: GetRegionRouteEntriesArgs, opts?: pulumi.InvokeOptions): Promise<GetRegionRouteEntriesResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("alicloud:cen/getRegionRouteEntries:getRegionRouteEntries", {
        "instanceId": args.instanceId,
        "outputFile": args.outputFile,
        "regionId": args.regionId,
    }, opts);
}

/**
 * A collection of arguments for invoking getRegionRouteEntries.
 */
export interface GetRegionRouteEntriesArgs {
    /**
     * ID of the CEN instance.
     */
    instanceId: string;
    outputFile?: string;
    /**
     * ID of the region.
     */
    regionId: string;
}

/**
 * A collection of values returned by getRegionRouteEntries.
 */
export interface GetRegionRouteEntriesResult {
    /**
     * A list of CEN Route Entries. Each element contains the following attributes:
     */
    readonly entries: outputs.cen.GetRegionRouteEntriesEntry[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly outputFile?: string;
    readonly regionId: string;
}

export function getRegionRouteEntriesOutput(args: GetRegionRouteEntriesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRegionRouteEntriesResult> {
    return pulumi.output(args).apply(a => getRegionRouteEntries(a, opts))
}

/**
 * A collection of arguments for invoking getRegionRouteEntries.
 */
export interface GetRegionRouteEntriesOutputArgs {
    /**
     * ID of the CEN instance.
     */
    instanceId: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * ID of the region.
     */
    regionId: pulumi.Input<string>;
}