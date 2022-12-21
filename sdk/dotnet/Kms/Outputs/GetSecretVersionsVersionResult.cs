// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Kms.Outputs
{

    [OutputType]
    public sealed class GetSecretVersionsVersionResult
    {
        /// <summary>
        /// The secret value. Secrets Manager decrypts the stored secret value in ciphertext and returns it. (Returned when `enable_details` is true).
        /// </summary>
        public readonly string SecretData;
        /// <summary>
        /// The type of the secret value. (Returned when `enable_details` is true).
        /// </summary>
        public readonly string SecretDataType;
        /// <summary>
        /// The name of the secret.
        /// </summary>
        public readonly string SecretName;
        /// <summary>
        /// The version number of the secret value.
        /// </summary>
        public readonly string VersionId;
        /// <summary>
        /// Stage labels that mark the secret version.
        /// </summary>
        public readonly ImmutableArray<string> VersionStages;

        [OutputConstructor]
        private GetSecretVersionsVersionResult(
            string secretData,

            string secretDataType,

            string secretName,

            string versionId,

            ImmutableArray<string> versionStages)
        {
            SecretData = secretData;
            SecretDataType = secretDataType;
            SecretName = secretName;
            VersionId = versionId;
            VersionStages = versionStages;
        }
    }
}