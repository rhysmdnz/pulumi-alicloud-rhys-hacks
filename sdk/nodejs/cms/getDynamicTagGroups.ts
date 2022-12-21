// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Cms Dynamic Tag Groups of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.142.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const config = new pulumi.Config();
 * const name = config.get("name") || "example_value";
 * const defaultAlarmContactGroup = new alicloud.cms.AlarmContactGroup("defaultAlarmContactGroup", {
 *     alarmContactGroupName: name,
 *     describe: "example_value",
 *     enableSubscribed: true,
 * });
 * const defaultDynamicTagGroup = new alicloud.cms.DynamicTagGroup("defaultDynamicTagGroup", {
 *     contactGroupLists: [defaultAlarmContactGroup.id],
 *     tagKey: "your_tag_key",
 *     matchExpresses: [{
 *         tagValue: "your_tag_value",
 *         tagValueMatchFunction: "all",
 *     }],
 * });
 * const ids = alicloud.cms.getDynamicTagGroupsOutput({
 *     ids: [defaultDynamicTagGroup.id],
 * });
 * export const cmsDynamicTagGroupId1 = ids.apply(ids => ids.groups?[0]?.id);
 * ```
 */
export function getDynamicTagGroups(args?: GetDynamicTagGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetDynamicTagGroupsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("alicloud:cms/getDynamicTagGroups:getDynamicTagGroups", {
        "ids": args.ids,
        "outputFile": args.outputFile,
        "status": args.status,
        "tagKey": args.tagKey,
    }, opts);
}

/**
 * A collection of arguments for invoking getDynamicTagGroups.
 */
export interface GetDynamicTagGroupsArgs {
    /**
     * A list of Dynamic Tag Group IDs.
     */
    ids?: string[];
    outputFile?: string;
    /**
     * The status of the resource. Valid values: `RUNNING`, `FINISH`.
     */
    status?: string;
    /**
     * The tag key of the tag.
     */
    tagKey?: string;
}

/**
 * A collection of values returned by getDynamicTagGroups.
 */
export interface GetDynamicTagGroupsResult {
    readonly groups: outputs.cms.GetDynamicTagGroupsGroup[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly outputFile?: string;
    readonly status?: string;
    readonly tagKey?: string;
}

export function getDynamicTagGroupsOutput(args?: GetDynamicTagGroupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDynamicTagGroupsResult> {
    return pulumi.output(args).apply(a => getDynamicTagGroups(a, opts))
}

/**
 * A collection of arguments for invoking getDynamicTagGroups.
 */
export interface GetDynamicTagGroupsOutputArgs {
    /**
     * A list of Dynamic Tag Group IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    outputFile?: pulumi.Input<string>;
    /**
     * The status of the resource. Valid values: `RUNNING`, `FINISH`.
     */
    status?: pulumi.Input<string>;
    /**
     * The tag key of the tag.
     */
    tagKey?: pulumi.Input<string>;
}