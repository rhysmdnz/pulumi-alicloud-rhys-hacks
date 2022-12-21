// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vod.Outputs
{

    [OutputType]
    public sealed class GetDomainsDomainSourceResult
    {
        public readonly string SourceContent;
        public readonly string SourcePort;
        public readonly string SourcePriority;
        public readonly string SourceType;

        [OutputConstructor]
        private GetDomainsDomainSourceResult(
            string sourceContent,

            string sourcePort,

            string sourcePriority,

            string sourceType)
        {
            SourceContent = sourceContent;
            SourcePort = sourcePort;
            SourcePriority = sourcePriority;
            SourceType = sourceType;
        }
    }
}