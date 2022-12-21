// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Yundun
{
    public static class GetDBAuditInstance
    {
        public static Task<GetDBAuditInstanceResult> InvokeAsync(GetDBAuditInstanceArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDBAuditInstanceResult>("alicloud:yundun/getDBAuditInstance:getDBAuditInstance", args ?? new GetDBAuditInstanceArgs(), options.WithDefaults());

        public static Output<GetDBAuditInstanceResult> Invoke(GetDBAuditInstanceInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDBAuditInstanceResult>("alicloud:yundun/getDBAuditInstance:getDBAuditInstance", args ?? new GetDBAuditInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDBAuditInstanceArgs : global::Pulumi.InvokeArgs
    {
        [Input("descriptionRegex")]
        public string? DescriptionRegex { get; set; }

        [Input("ids")]
        private List<string>? _ids;
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetDBAuditInstanceArgs()
        {
        }
        public static new GetDBAuditInstanceArgs Empty => new GetDBAuditInstanceArgs();
    }

    public sealed class GetDBAuditInstanceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("descriptionRegex")]
        public Input<string>? DescriptionRegex { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public GetDBAuditInstanceInvokeArgs()
        {
        }
        public static new GetDBAuditInstanceInvokeArgs Empty => new GetDBAuditInstanceInvokeArgs();
    }


    [OutputType]
    public sealed class GetDBAuditInstanceResult
    {
        public readonly string? DescriptionRegex;
        public readonly ImmutableArray<string> Descriptions;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly ImmutableArray<Outputs.GetDBAuditInstanceInstanceResult> Instances;
        public readonly string? OutputFile;
        public readonly ImmutableDictionary<string, object>? Tags;

        [OutputConstructor]
        private GetDBAuditInstanceResult(
            string? descriptionRegex,

            ImmutableArray<string> descriptions,

            string id,

            ImmutableArray<string> ids,

            ImmutableArray<Outputs.GetDBAuditInstanceInstanceResult> instances,

            string? outputFile,

            ImmutableDictionary<string, object>? tags)
        {
            DescriptionRegex = descriptionRegex;
            Descriptions = descriptions;
            Id = id;
            Ids = ids;
            Instances = instances;
            OutputFile = outputFile;
            Tags = tags;
        }
    }
}