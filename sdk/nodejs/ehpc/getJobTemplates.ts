// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides the Ehpc Job Templates of the current Alibaba Cloud user.
 *
 * > **NOTE:** Available in v1.133.0+.
 *
 * ## Example Usage
 *
 * Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const _default = new alicloud.ehpc.JobTemplate("default", {
 *     jobTemplateName: "example_value",
 *     commandLine: "./LammpsTest/lammps.pbs",
 * });
 * const ids = alicloud.ehpc.getJobTemplatesOutput({
 *     ids: [_default.id],
 * });
 * export const ehpcJobTemplateId1 = ids.apply(ids => ids.id);
 * ```
 */
export function getJobTemplates(args?: GetJobTemplatesArgs, opts?: pulumi.InvokeOptions): Promise<GetJobTemplatesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("alicloud:ehpc/getJobTemplates:getJobTemplates", {
        "ids": args.ids,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking getJobTemplates.
 */
export interface GetJobTemplatesArgs {
    /**
     * A list of Job Template IDs.
     */
    ids?: string[];
    outputFile?: string;
}

/**
 * A collection of values returned by getJobTemplates.
 */
export interface GetJobTemplatesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly outputFile?: string;
    readonly templates: outputs.ehpc.GetJobTemplatesTemplate[];
}

export function getJobTemplatesOutput(args?: GetJobTemplatesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetJobTemplatesResult> {
    return pulumi.output(args).apply(a => getJobTemplates(a, opts))
}

/**
 * A collection of arguments for invoking getJobTemplates.
 */
export interface GetJobTemplatesOutputArgs {
    /**
     * A list of Job Template IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    outputFile?: pulumi.Input<string>;
}