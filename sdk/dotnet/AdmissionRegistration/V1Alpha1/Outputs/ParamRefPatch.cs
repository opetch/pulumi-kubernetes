// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.AdmissionRegistration.V1Alpha1
{

    /// <summary>
    /// ParamRef references a parameter resource
    /// </summary>
    [OutputType]
    public sealed class ParamRefPatch
    {
        /// <summary>
        /// Name of the resource being referenced.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Namespace of the referenced resource. Should be empty for the cluster-scoped resources
        /// </summary>
        public readonly string Namespace;

        [OutputConstructor]
        private ParamRefPatch(
            string name,

            string @namespace)
        {
            Name = name;
            Namespace = @namespace;
        }
    }
}
