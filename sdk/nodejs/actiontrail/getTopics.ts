// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source provides a list of ALIKAFKA Topics in an Alibaba Cloud account according to the specified filters.
 *
 * > **NOTE:** Available in 1.56.0+
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const topicsDs = pulumi.output(alicloud.actiontrail.getTopics({
 *     instanceId: "xxx",
 *     nameRegex: "alikafkaTopicName",
 *     outputFile: "topics.txt",
 * }));
 *
 * export const firstTopicName = topicsDs.topics[0].topic;
 * ```
 */
export function getTopics(args: GetTopicsArgs, opts?: pulumi.InvokeOptions): Promise<GetTopicsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("alicloud:actiontrail/getTopics:getTopics", {
        "ids": args.ids,
        "instanceId": args.instanceId,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "pageNumber": args.pageNumber,
        "pageSize": args.pageSize,
        "topic": args.topic,
    }, opts);
}

/**
 * A collection of arguments for invoking getTopics.
 */
export interface GetTopicsArgs {
    /**
     * A list of ALIKAFKA Topics IDs, It is formatted to `<instance_id>:<topic>`.
     */
    ids?: string[];
    /**
     * ID of the instance.
     */
    instanceId: string;
    /**
     * A regex string to filter results by the topic name.
     */
    nameRegex?: string;
    outputFile?: string;
    pageNumber?: number;
    pageSize?: number;
    /**
     * A topic to filter results by the topic name.
     */
    topic?: string;
}

/**
 * A collection of values returned by getTopics.
 */
export interface GetTopicsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    /**
     * The instanceId of the instance.
     */
    readonly instanceId: string;
    readonly nameRegex?: string;
    /**
     * A list of topic names.
     */
    readonly names: string[];
    readonly outputFile?: string;
    readonly pageNumber?: number;
    readonly pageSize?: number;
    /**
     * The name of the topic.
     */
    readonly topic?: string;
    /**
     * A list of topics. Each element contains the following attributes:
     */
    readonly topics: outputs.actiontrail.GetTopicsTopic[];
    readonly totalCount: number;
}

export function getTopicsOutput(args: GetTopicsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTopicsResult> {
    return pulumi.output(args).apply(a => getTopics(a, opts))
}

/**
 * A collection of arguments for invoking getTopics.
 */
export interface GetTopicsOutputArgs {
    /**
     * A list of ALIKAFKA Topics IDs, It is formatted to `<instance_id>:<topic>`.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID of the instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * A regex string to filter results by the topic name.
     */
    nameRegex?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    pageNumber?: pulumi.Input<number>;
    pageSize?: pulumi.Input<number>;
    /**
     * A topic to filter results by the topic name.
     */
    topic?: pulumi.Input<string>;
}