// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a WAF Domain resource to create domain in the Web Application Firewall.
 *
 * For information about WAF and how to use it, see [What is Alibaba Cloud WAF](https://www.alibabacloud.com/help/doc-detail/28517.htm).
 *
 * > **NOTE:** Available in 1.82.0+ .
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const domain = new alicloud.waf.Domain("domain", {
 *     clusterType: "PhysicalCluster",
 *     domainName: "alicloud-provider.cn",
 *     http2Ports: ["443"],
 *     httpPorts: ["80"],
 *     httpToUserIp: "Off",
 *     httpsPorts: ["443"],
 *     httpsRedirect: "Off",
 *     instanceId: "waf-123455",
 *     isAccessProduct: "On",
 *     loadBalancing: "IpHash",
 *     logHeaders: [{
 *         key: "foo",
 *         value: "http",
 *     }],
 *     sourceIps: ["1.1.1.1"],
 * });
 * ```
 *
 * ## Import
 *
 * WAF domain can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import alicloud:waf/domain:Domain domain waf-132435:www.domain.com
 * ```
 */
export class Domain extends pulumi.CustomResource {
    /**
     * Get an existing Domain resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DomainState, opts?: pulumi.CustomResourceOptions): Domain {
        return new Domain(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'alicloud:waf/domain:Domain';

    /**
     * Returns true if the given object is an instance of Domain.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Domain {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Domain.__pulumiType;
    }

    /**
     * The type of the WAF cluster. Valid values: `PhysicalCluster` and `VirtualCluster`. Default to `PhysicalCluster`.
     */
    public readonly clusterType!: pulumi.Output<string | undefined>;
    /**
     * The CNAME record assigned by the WAF instance to the specified domain.
     */
    public /*out*/ readonly cname!: pulumi.Output<string>;
    /**
     * The connection timeout for WAF exclusive clusters. Unit: seconds.
     */
    public readonly connectionTime!: pulumi.Output<number | undefined>;
    /**
     * Field `domain` has been deprecated from version 1.94.0. Use `domainName` instead.
     *
     * @deprecated Field 'domain' has been deprecated from version 1.94.0. Use 'domain_name' instead.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * The domain that you want to add to WAF. The `domainName` is required when the value of the `domain`  is Empty.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * List of the HTTP 2.0 ports.
     */
    public readonly http2Ports!: pulumi.Output<string[] | undefined>;
    /**
     * List of the HTTP ports.
     */
    public readonly httpPorts!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies whether to enable the HTTP back-to-origin feature. After this feature is enabled, the WAF instance can use HTTP to forward HTTPS requests to the origin server. 
     * By default, port 80 is used to forward the requests to the origin server. Valid values: `On` and `Off`. Default to `Off`.
     */
    public readonly httpToUserIp!: pulumi.Output<string | undefined>;
    /**
     * List of the HTTPS ports.
     */
    public readonly httpsPorts!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies whether to redirect HTTP requests as HTTPS requests. Valid values: "On" and `Off`. Default to `Off`.
     */
    public readonly httpsRedirect!: pulumi.Output<string | undefined>;
    /**
     * The ID of the WAF instance.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Specifies whether to configure a Layer-7 proxy, such as Anti-DDoS Pro or CDN, to filter the inbound traffic before it is forwarded to WAF. Valid values: `On` and `Off`. Default to `Off`.
     */
    public readonly isAccessProduct!: pulumi.Output<string>;
    /**
     * The load balancing algorithm that is used to forward requests to the origin. Valid values: `IpHash` and `RoundRobin`. Default to `IpHash`.
     */
    public readonly loadBalancing!: pulumi.Output<string | undefined>;
    /**
     * The key-value pair that is used to mark the traffic that flows through WAF to the domain. Each item contains two field:
     * * key: The key of label
     * * value: The value of label
     */
    public readonly logHeaders!: pulumi.Output<outputs.waf.DomainLogHeader[] | undefined>;
    /**
     * The read timeout of a WAF exclusive cluster. Unit: seconds.
     */
    public readonly readTime!: pulumi.Output<number | undefined>;
    /**
     * The ID of the resource group to which the queried domain belongs in Resource Management. By default, no value is specified, indicating that the domain belongs to the default resource group.
     */
    public readonly resourceGroupId!: pulumi.Output<string>;
    /**
     * List of the IP address or domain of the origin server to which the specified domain points.
     */
    public readonly sourceIps!: pulumi.Output<string[] | undefined>;
    /**
     * The timeout period for a WAF exclusive cluster write connection. Unit: seconds.
     */
    public readonly writeTime!: pulumi.Output<number | undefined>;

    /**
     * Create a Domain resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DomainArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DomainArgs | DomainState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DomainState | undefined;
            resourceInputs["clusterType"] = state ? state.clusterType : undefined;
            resourceInputs["cname"] = state ? state.cname : undefined;
            resourceInputs["connectionTime"] = state ? state.connectionTime : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["domainName"] = state ? state.domainName : undefined;
            resourceInputs["http2Ports"] = state ? state.http2Ports : undefined;
            resourceInputs["httpPorts"] = state ? state.httpPorts : undefined;
            resourceInputs["httpToUserIp"] = state ? state.httpToUserIp : undefined;
            resourceInputs["httpsPorts"] = state ? state.httpsPorts : undefined;
            resourceInputs["httpsRedirect"] = state ? state.httpsRedirect : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["isAccessProduct"] = state ? state.isAccessProduct : undefined;
            resourceInputs["loadBalancing"] = state ? state.loadBalancing : undefined;
            resourceInputs["logHeaders"] = state ? state.logHeaders : undefined;
            resourceInputs["readTime"] = state ? state.readTime : undefined;
            resourceInputs["resourceGroupId"] = state ? state.resourceGroupId : undefined;
            resourceInputs["sourceIps"] = state ? state.sourceIps : undefined;
            resourceInputs["writeTime"] = state ? state.writeTime : undefined;
        } else {
            const args = argsOrState as DomainArgs | undefined;
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.isAccessProduct === undefined) && !opts.urn) {
                throw new Error("Missing required property 'isAccessProduct'");
            }
            resourceInputs["clusterType"] = args ? args.clusterType : undefined;
            resourceInputs["connectionTime"] = args ? args.connectionTime : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["http2Ports"] = args ? args.http2Ports : undefined;
            resourceInputs["httpPorts"] = args ? args.httpPorts : undefined;
            resourceInputs["httpToUserIp"] = args ? args.httpToUserIp : undefined;
            resourceInputs["httpsPorts"] = args ? args.httpsPorts : undefined;
            resourceInputs["httpsRedirect"] = args ? args.httpsRedirect : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["isAccessProduct"] = args ? args.isAccessProduct : undefined;
            resourceInputs["loadBalancing"] = args ? args.loadBalancing : undefined;
            resourceInputs["logHeaders"] = args ? args.logHeaders : undefined;
            resourceInputs["readTime"] = args ? args.readTime : undefined;
            resourceInputs["resourceGroupId"] = args ? args.resourceGroupId : undefined;
            resourceInputs["sourceIps"] = args ? args.sourceIps : undefined;
            resourceInputs["writeTime"] = args ? args.writeTime : undefined;
            resourceInputs["cname"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Domain.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Domain resources.
 */
export interface DomainState {
    /**
     * The type of the WAF cluster. Valid values: `PhysicalCluster` and `VirtualCluster`. Default to `PhysicalCluster`.
     */
    clusterType?: pulumi.Input<string>;
    /**
     * The CNAME record assigned by the WAF instance to the specified domain.
     */
    cname?: pulumi.Input<string>;
    /**
     * The connection timeout for WAF exclusive clusters. Unit: seconds.
     */
    connectionTime?: pulumi.Input<number>;
    /**
     * Field `domain` has been deprecated from version 1.94.0. Use `domainName` instead.
     *
     * @deprecated Field 'domain' has been deprecated from version 1.94.0. Use 'domain_name' instead.
     */
    domain?: pulumi.Input<string>;
    /**
     * The domain that you want to add to WAF. The `domainName` is required when the value of the `domain`  is Empty.
     */
    domainName?: pulumi.Input<string>;
    /**
     * List of the HTTP 2.0 ports.
     */
    http2Ports?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of the HTTP ports.
     */
    httpPorts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies whether to enable the HTTP back-to-origin feature. After this feature is enabled, the WAF instance can use HTTP to forward HTTPS requests to the origin server. 
     * By default, port 80 is used to forward the requests to the origin server. Valid values: `On` and `Off`. Default to `Off`.
     */
    httpToUserIp?: pulumi.Input<string>;
    /**
     * List of the HTTPS ports.
     */
    httpsPorts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies whether to redirect HTTP requests as HTTPS requests. Valid values: "On" and `Off`. Default to `Off`.
     */
    httpsRedirect?: pulumi.Input<string>;
    /**
     * The ID of the WAF instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Specifies whether to configure a Layer-7 proxy, such as Anti-DDoS Pro or CDN, to filter the inbound traffic before it is forwarded to WAF. Valid values: `On` and `Off`. Default to `Off`.
     */
    isAccessProduct?: pulumi.Input<string>;
    /**
     * The load balancing algorithm that is used to forward requests to the origin. Valid values: `IpHash` and `RoundRobin`. Default to `IpHash`.
     */
    loadBalancing?: pulumi.Input<string>;
    /**
     * The key-value pair that is used to mark the traffic that flows through WAF to the domain. Each item contains two field:
     * * key: The key of label
     * * value: The value of label
     */
    logHeaders?: pulumi.Input<pulumi.Input<inputs.waf.DomainLogHeader>[]>;
    /**
     * The read timeout of a WAF exclusive cluster. Unit: seconds.
     */
    readTime?: pulumi.Input<number>;
    /**
     * The ID of the resource group to which the queried domain belongs in Resource Management. By default, no value is specified, indicating that the domain belongs to the default resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * List of the IP address or domain of the origin server to which the specified domain points.
     */
    sourceIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The timeout period for a WAF exclusive cluster write connection. Unit: seconds.
     */
    writeTime?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Domain resource.
 */
export interface DomainArgs {
    /**
     * The type of the WAF cluster. Valid values: `PhysicalCluster` and `VirtualCluster`. Default to `PhysicalCluster`.
     */
    clusterType?: pulumi.Input<string>;
    /**
     * The connection timeout for WAF exclusive clusters. Unit: seconds.
     */
    connectionTime?: pulumi.Input<number>;
    /**
     * Field `domain` has been deprecated from version 1.94.0. Use `domainName` instead.
     *
     * @deprecated Field 'domain' has been deprecated from version 1.94.0. Use 'domain_name' instead.
     */
    domain?: pulumi.Input<string>;
    /**
     * The domain that you want to add to WAF. The `domainName` is required when the value of the `domain`  is Empty.
     */
    domainName?: pulumi.Input<string>;
    /**
     * List of the HTTP 2.0 ports.
     */
    http2Ports?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of the HTTP ports.
     */
    httpPorts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies whether to enable the HTTP back-to-origin feature. After this feature is enabled, the WAF instance can use HTTP to forward HTTPS requests to the origin server. 
     * By default, port 80 is used to forward the requests to the origin server. Valid values: `On` and `Off`. Default to `Off`.
     */
    httpToUserIp?: pulumi.Input<string>;
    /**
     * List of the HTTPS ports.
     */
    httpsPorts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies whether to redirect HTTP requests as HTTPS requests. Valid values: "On" and `Off`. Default to `Off`.
     */
    httpsRedirect?: pulumi.Input<string>;
    /**
     * The ID of the WAF instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Specifies whether to configure a Layer-7 proxy, such as Anti-DDoS Pro or CDN, to filter the inbound traffic before it is forwarded to WAF. Valid values: `On` and `Off`. Default to `Off`.
     */
    isAccessProduct: pulumi.Input<string>;
    /**
     * The load balancing algorithm that is used to forward requests to the origin. Valid values: `IpHash` and `RoundRobin`. Default to `IpHash`.
     */
    loadBalancing?: pulumi.Input<string>;
    /**
     * The key-value pair that is used to mark the traffic that flows through WAF to the domain. Each item contains two field:
     * * key: The key of label
     * * value: The value of label
     */
    logHeaders?: pulumi.Input<pulumi.Input<inputs.waf.DomainLogHeader>[]>;
    /**
     * The read timeout of a WAF exclusive cluster. Unit: seconds.
     */
    readTime?: pulumi.Input<number>;
    /**
     * The ID of the resource group to which the queried domain belongs in Resource Management. By default, no value is specified, indicating that the domain belongs to the default resource group.
     */
    resourceGroupId?: pulumi.Input<string>;
    /**
     * List of the IP address or domain of the origin server to which the specified domain points.
     */
    sourceIps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The timeout period for a WAF exclusive cluster write connection. Unit: seconds.
     */
    writeTime?: pulumi.Input<number>;
}
