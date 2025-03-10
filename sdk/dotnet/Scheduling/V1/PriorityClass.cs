// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Scheduling.V1
{
    /// <summary>
    /// PriorityClass defines mapping from a priority class name to the priority integer value. The value can be any valid integer.
    /// </summary>
    [KubernetesResourceType("kubernetes:scheduling.k8s.io/v1:PriorityClass")]
    public partial class PriorityClass : KubernetesResource
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Output("apiVersion")]
        public Output<string> ApiVersion { get; private set; } = null!;

        /// <summary>
        /// description is an arbitrary string that usually provides guidelines on when this priority class should be used.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.
        /// </summary>
        [Output("globalDefault")]
        public Output<bool> GlobalDefault { get; private set; } = null!;

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        /// </summary>
        [Output("metadata")]
        public Output<Pulumi.Kubernetes.Types.Outputs.Meta.V1.ObjectMeta> Metadata { get; private set; } = null!;

        /// <summary>
        /// preemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset.
        /// </summary>
        [Output("preemptionPolicy")]
        public Output<string> PreemptionPolicy { get; private set; } = null!;

        /// <summary>
        /// value represents the integer value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.
        /// </summary>
        [Output("value")]
        public Output<int> Value { get; private set; } = null!;


        /// <summary>
        /// Create a PriorityClass resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PriorityClass(string name, Pulumi.Kubernetes.Types.Inputs.Scheduling.V1.PriorityClassArgs? args = null, CustomResourceOptions? options = null)
            : base("kubernetes:scheduling.k8s.io/v1:PriorityClass", name, MakeArgs(args), MakeResourceOptions(options, ""))
        {
        }
        internal PriorityClass(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("kubernetes:scheduling.k8s.io/v1:PriorityClass", name, new DictionaryResourceArgs(dictionary), MakeResourceOptions(options, ""))
        {
        }

        private PriorityClass(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("kubernetes:scheduling.k8s.io/v1:PriorityClass", name, null, MakeResourceOptions(options, id))
        {
        }

        private static Pulumi.Kubernetes.Types.Inputs.Scheduling.V1.PriorityClassArgs? MakeArgs(Pulumi.Kubernetes.Types.Inputs.Scheduling.V1.PriorityClassArgs? args)
        {
            args ??= new Pulumi.Kubernetes.Types.Inputs.Scheduling.V1.PriorityClassArgs();
            args.ApiVersion = "scheduling.k8s.io/v1";
            args.Kind = "PriorityClass";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "kubernetes:scheduling.k8s.io/v1alpha1:PriorityClass"},
                    new global::Pulumi.Alias { Type = "kubernetes:scheduling.k8s.io/v1beta1:PriorityClass"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PriorityClass resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PriorityClass Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PriorityClass(name, id, options);
        }
    }
}
namespace Pulumi.Kubernetes.Types.Inputs.Scheduling.V1
{

    public class PriorityClassArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Input("apiVersion")]
        public Input<string>? ApiVersion { get; set; }

        /// <summary>
        /// description is an arbitrary string that usually provides guidelines on when this priority class should be used.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as `globalDefault`. However, if more than one PriorityClasses exists with their `globalDefault` field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.
        /// </summary>
        [Input("globalDefault")]
        public Input<bool>? GlobalDefault { get; set; }

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        /// </summary>
        [Input("metadata")]
        public Input<Pulumi.Kubernetes.Types.Inputs.Meta.V1.ObjectMetaArgs>? Metadata { get; set; }

        /// <summary>
        /// preemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset.
        /// </summary>
        [Input("preemptionPolicy")]
        public Input<string>? PreemptionPolicy { get; set; }

        /// <summary>
        /// value represents the integer value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.
        /// </summary>
        [Input("value", required: true)]
        public Input<int> Value { get; set; } = null!;

        public PriorityClassArgs()
        {
        }
        public static new PriorityClassArgs Empty => new PriorityClassArgs();
    }
}
