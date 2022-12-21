// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Hbase
{
    public static class GetInstanceTypes
    {
        /// <summary>
        /// This data source provides availability instance_types for HBase that can be accessed by an Alibaba Cloud account within the region configured in the provider.
        /// 
        /// &gt; **NOTE:** Available in v1.106.0+.
        /// </summary>
        public static Task<GetInstanceTypesResult> InvokeAsync(GetInstanceTypesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceTypesResult>("alicloud:hbase/getInstanceTypes:getInstanceTypes", args ?? new GetInstanceTypesArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides availability instance_types for HBase that can be accessed by an Alibaba Cloud account within the region configured in the provider.
        /// 
        /// &gt; **NOTE:** Available in v1.106.0+.
        /// </summary>
        public static Output<GetInstanceTypesResult> Invoke(GetInstanceTypesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstanceTypesResult>("alicloud:hbase/getInstanceTypes:getInstanceTypes", args ?? new GetInstanceTypesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceTypesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The charge type of create hbase cluster instance, `PrePaid` or `PostPaid`.
        /// </summary>
        [Input("chargeType")]
        public string? ChargeType { get; set; }

        /// <summary>
        /// The disk type, `cloud_ssd`, `cloud_essd_pl1`, `cloud_efficiency`, `local_hdd_pro`, `local_ssd_pro`.
        /// </summary>
        [Input("diskType")]
        public string? DiskType { get; set; }

        /// <summary>
        /// The engine name, `singlehbase`, `hbase`, `hbaseue`, `bds`.
        /// </summary>
        [Input("engine")]
        public string? Engine { get; set; }

        /// <summary>
        /// The hbase instance type of create hbase cluster instance.
        /// </summary>
        [Input("instanceType")]
        public string? InstanceType { get; set; }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The dest region id, default client region.
        /// </summary>
        [Input("regionId")]
        public string? RegionId { get; set; }

        /// <summary>
        /// The engine version, singlehbase/hbase=1.1/2.0, bds=1.0.
        /// </summary>
        [Input("version")]
        public string? Version { get; set; }

        /// <summary>
        /// The zone id, belong to regionId.
        /// </summary>
        [Input("zoneId")]
        public string? ZoneId { get; set; }

        public GetInstanceTypesArgs()
        {
        }
        public static new GetInstanceTypesArgs Empty => new GetInstanceTypesArgs();
    }

    public sealed class GetInstanceTypesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The charge type of create hbase cluster instance, `PrePaid` or `PostPaid`.
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// The disk type, `cloud_ssd`, `cloud_essd_pl1`, `cloud_efficiency`, `local_hdd_pro`, `local_ssd_pro`.
        /// </summary>
        [Input("diskType")]
        public Input<string>? DiskType { get; set; }

        /// <summary>
        /// The engine name, `singlehbase`, `hbase`, `hbaseue`, `bds`.
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        /// <summary>
        /// The hbase instance type of create hbase cluster instance.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The dest region id, default client region.
        /// </summary>
        [Input("regionId")]
        public Input<string>? RegionId { get; set; }

        /// <summary>
        /// The engine version, singlehbase/hbase=1.1/2.0, bds=1.0.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        /// <summary>
        /// The zone id, belong to regionId.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public GetInstanceTypesInvokeArgs()
        {
        }
        public static new GetInstanceTypesInvokeArgs Empty => new GetInstanceTypesInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstanceTypesResult
    {
        public readonly string? ChargeType;
        /// <summary>
        /// (Available in 1.115.0+) A list of core instance types. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceTypesCoreInstanceTypeResult> CoreInstanceTypes;
        public readonly string? DiskType;
        /// <summary>
        /// Name of the engine.
        /// </summary>
        public readonly string? Engine;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of instance types type IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// Name of the instance type.
        /// </summary>
        public readonly string? InstanceType;
        /// <summary>
        /// (Available in 1.115.0+) A list of master instance types. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceTypesMasterInstanceTypeResult> MasterInstanceTypes;
        public readonly string? OutputFile;
        public readonly string? RegionId;
        /// <summary>
        /// (Deprecated) A list of instance types. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstanceTypesTypeResult> Types;
        /// <summary>
        /// The version of the engine.
        /// </summary>
        public readonly string? Version;
        public readonly string? ZoneId;

        [OutputConstructor]
        private GetInstanceTypesResult(
            string? chargeType,

            ImmutableArray<Outputs.GetInstanceTypesCoreInstanceTypeResult> coreInstanceTypes,

            string? diskType,

            string? engine,

            string id,

            ImmutableArray<string> ids,

            string? instanceType,

            ImmutableArray<Outputs.GetInstanceTypesMasterInstanceTypeResult> masterInstanceTypes,

            string? outputFile,

            string? regionId,

            ImmutableArray<Outputs.GetInstanceTypesTypeResult> types,

            string? version,

            string? zoneId)
        {
            ChargeType = chargeType;
            CoreInstanceTypes = coreInstanceTypes;
            DiskType = diskType;
            Engine = engine;
            Id = id;
            Ids = ids;
            InstanceType = instanceType;
            MasterInstanceTypes = masterInstanceTypes;
            OutputFile = outputFile;
            RegionId = regionId;
            Types = types;
            Version = version;
            ZoneId = zoneId;
        }
    }
}