// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CR.Outputs
{

    [OutputType]
    public sealed class GetReposRepoResult
    {
        /// <summary>
        /// The repository domain list.
        /// </summary>
        public readonly Outputs.GetReposRepoDomainListResult DomainList;
        /// <summary>
        /// Name of container registry namespace.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Name of container registry namespace where the repositories are located in.
        /// </summary>
        public readonly string Namespace;
        /// <summary>
        /// `PUBLIC` or `PRIVATE`, repository's visibility.
        /// </summary>
        public readonly string RepoType;
        /// <summary>
        /// The repository general information.
        /// </summary>
        public readonly string Summary;
        /// <summary>
        /// A list of image tags belong to this repository. Each contains several attributes, see `Block Tag`.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetReposRepoTagResult> Tags;

        [OutputConstructor]
        private GetReposRepoResult(
            Outputs.GetReposRepoDomainListResult domainList,

            string name,

            string @namespace,

            string repoType,

            string summary,

            ImmutableArray<Outputs.GetReposRepoTagResult> tags)
        {
            DomainList = domainList;
            Name = name;
            Namespace = @namespace;
            RepoType = repoType;
            Summary = summary;
            Tags = tags;
        }
    }
}