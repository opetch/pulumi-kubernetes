// Copyright 2016-2021, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package helm

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/yaml"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Chart is a component representing a collection of resources described by an arbitrary Helm
// Chart. The Chart can be fetched from any source that is accessible to the `helm` command
// line. Values in the `values.yml` file can be overridden using `ChartOpts.values` (equivalent
// to `--set` or having multiple `values.yml` files). Objects can be transformed arbitrarily by
// supplying callbacks to `ChartOpts.transformations`.
//
// `Chart` does not use Tiller. The Chart specified is copied and expanded locally; the semantics
// are equivalent to running `helm template` and then using Pulumi to manage the resulting YAML
// manifests. Any values that would be retrieved in-cluster are assigned fake values, and
// none of Tiller's server-side validity testing is executed.
//
// ## Example Usage
// ### Local Chart Directory
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/helm/v3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//	    pulumi.Run(func(ctx *pulumi.Context) error {
//	        _, err := helm.NewChart(ctx, "nginx-ingress", helm.ChartArgs{
//	            Path: pulumi.String("./nginx-ingress"),
//	        })
//	        if err != nil {
//	            return err
//	        }
//
//	        return nil
//	    })
//	}
//
// ```
// ### Remote Chart
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/helm/v3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//	    pulumi.Run(func(ctx *pulumi.Context) error {
//	        _, err := helm.NewChart(ctx, "nginx-ingress", helm.ChartArgs{
//	            Chart:   pulumi.String("nginx-ingress"),
//	            Version: pulumi.String("1.24.4"),
//	            FetchArgs: helm.FetchArgs{
//	                Repo: pulumi.String("https://charts.helm.sh/stable"),
//	            },
//	        })
//	        if err != nil {
//	            return err
//	        }
//
//	        return nil
//	    })
//	}
//
// ```
// ### Set Chart values
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/helm/v3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//	    pulumi.Run(func(ctx *pulumi.Context) error {
//	        _, err := helm.NewChart(ctx, "nginx-ingress", helm.ChartArgs{
//	            Chart:   pulumi.String("nginx-ingress"),
//	            Version: pulumi.String("1.24.4"),
//	            FetchArgs: helm.FetchArgs{
//	                Repo: pulumi.String("https://charts.helm.sh/stable"),
//	            },
//	            Values: pulumi.Map{
//	                "controller": pulumi.Map{
//	                    "metrics": pulumi.Map{
//	                        "enabled": pulumi.Bool(true),
//	                    },
//	                },
//	            },
//	        })
//	        if err != nil {
//	            return err
//	        }
//
//	        return nil
//	    })
//	}
//
// ```
// ### Deploy Chart into Namespace
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/helm/v3"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//	    pulumi.Run(func(ctx *pulumi.Context) error {
//	        _, err := helm.NewChart(ctx, "nginx-ingress", helm.ChartArgs{
//	            Chart:     pulumi.String("nginx-ingress"),
//	            Version:   pulumi.String("1.24.4"),
//	            Namespace: pulumi.String("test-namespace"),
//	            FetchArgs: helm.FetchArgs{
//	                Repo: pulumi.String("https://charts.helm.sh/stable"),
//	            },
//	        })
//	        if err != nil {
//	            return err
//	        }
//
//	        return nil
//	    })
//	}
//
// ```
// ### Chart with Transformations
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/helm/v3"
//	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/yaml"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//	    pulumi.Run(func(ctx *pulumi.Context) error {
//	        _, err := helm.NewChart(ctx, "nginx-ingress", helm.ChartArgs{
//	            Chart:   pulumi.String("nginx-ingress"),
//	            Version: pulumi.String("1.24.4"),
//	            FetchArgs: helm.FetchArgs{
//	                Repo: pulumi.String("https://charts.helm.sh/stable"),
//	            },
//	            Transformations: []yaml.Transformation{
//	                // Make every service private to the cluster, i.e., turn all services into ClusterIP
//	                // instead of LoadBalancer.
//	                func(state map[string]interface{}, opts ...pulumi.ResourceOption) {
//	                    if state["kind"] == "Service" {
//	                        spec := state["spec"].(map[string]interface{})
//	                        spec["type"] = "ClusterIP"
//	                    }
//	                },
//
//	                // Set a resource alias for a previous name.
//	                func(state map[string]interface{}, opts ...pulumi.ResourceOption) {
//	                    if state["kind"] == "Deployment" {
//	                        aliases := pulumi.Aliases([]pulumi.Alias{
//	                            {
//	                                Name: pulumi.String("oldName"),
//	                            },
//	                        })
//	                        opts = append(opts, aliases)
//	                    }
//	                },
//
//	                // Omit a resource from the Chart by transforming the specified resource definition
//	                // to an empty List.
//	                func(state map[string]interface{}, opts ...pulumi.ResourceOption) {
//	                    name := state["metadata"].(map[string]interface{})["name"]
//	                    if state["kind"] == "Pod" && name == "test" {
//	                        state["apiVersion"] = "v1"
//	                        state["kind"] = "List"
//	                    }
//	                },
//	            },
//	        })
//	        if err != nil {
//	            return err
//	        }
//
//	        return nil
//	    })
//	}
//
// ```
type Chart struct {
	pulumi.ResourceState

	Ready     pulumi.ResourceArrayOutput
	Resources pulumi.Output
}

// NewChart registers a new resource with the given unique name, arguments, and options.
func NewChart(ctx *pulumi.Context,
	name string, args ChartArgs, opts ...pulumi.ResourceOption) (*Chart, error) {

	// Register the resulting resource state.
	chart := &Chart{}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:helm.sh/v2:Chart"),
		},
	})

	var resourceOrInvokeOptions []pulumi.ResourceOrInvokeOption
	for _, o := range opts {
		if asResOrInv, ok := o.(pulumi.ResourceOrInvokeOption); ok {
			resourceOrInvokeOptions = append(resourceOrInvokeOptions, asResOrInv)
		}
	}
	opts = append(opts, aliases)
	err := ctx.RegisterComponentResource("kubernetes:helm.sh/v3:Chart", name, chart, opts...)
	if err != nil {
		return nil, err
	}

	// Honor the resource name prefix if specified.
	if args.ResourcePrefix != "" {
		name = args.ResourcePrefix + "-" + name
	}

	resourceOrInvokeOptions = append(resourceOrInvokeOptions, pulumi.Parent(chart))
	resources := args.ToChartArgsOutput().ApplyT(func(args chartArgs) (map[string]pulumi.Resource, error) {
		return parseChart(ctx, name, args, resourceOrInvokeOptions...)
	})
	chart.Resources = resources

	// Finally, register all of the resources found.
	// Note: Go requires that we "pull" on our futures in order to get them scheduled for execution. Here, we use
	// the engine's RegisterResourceOutputs to wait for the resolution of all resources that this Helm chart created.
	err = ctx.RegisterResourceOutputs(chart, pulumi.Map{"resources": resources})
	if err != nil {
		return nil, errors.Wrap(err, "registering child resources")
	}

	chart.Ready = resources.ApplyT(func(x interface{}) []pulumi.Resource {
		resources := x.(map[string]pulumi.Resource)
		var outputs []pulumi.Resource
		for _, r := range resources {
			outputs = append(outputs, r)
		}
		return outputs
	}).(pulumi.ResourceArrayOutput)

	return chart, nil
}

func parseChart(ctx *pulumi.Context, name string, args chartArgs, opts ...pulumi.ResourceOrInvokeOption,
) (map[string]pulumi.Resource, error) {
	type jsonOptsArgs struct {
		chartArgs

		ReleaseName string `json:"release_name,omitempty"`
	}
	jsonOpts := jsonOptsArgs{
		chartArgs:   args,
		ReleaseName: name,
	}

	b, err := json.Marshal(jsonOpts)
	if err != nil {
		return nil, err
	}

	var invokeOpts []pulumi.InvokeOption
	var resourceOpts []pulumi.ResourceOption
	for _, o := range opts {
		invokeOpts = append(invokeOpts, o)
		resourceOpts = append(resourceOpts, o)
	}
	objs, err := helmTemplate(ctx, string(b), invokeOpts...)
	if err != nil {
		return nil, err
	}

	transformations := args.Transformations
	if args.SkipAwait {
		transformations = yaml.AddSkipAwaitTransformation(transformations)
	}

	resources, err := yaml.ParseYamlObjects(ctx, objs, transformations, args.ResourcePrefix, resourceOpts...)
	if err != nil {
		return nil, err
	}
	return resources, nil
}

// helmTemplate invokes the function to fetch and template a Helm Chart and decompose it into object structures.
func helmTemplate(
	ctx *pulumi.Context, jsonOpts string, opts ...pulumi.InvokeOption,
) ([]map[string]interface{}, error) {
	args := struct {
		JsonOpts string `pulumi:"jsonOpts"`
	}{JsonOpts: jsonOpts}
	var ret struct {
		Result []map[string]interface{} `pulumi:"result"`
	}

	if err := ctx.Invoke("kubernetes:helm:template", &args, &ret, opts...); err != nil {
		return nil, errors.Wrap(err, "failed to invoke helm template")
	}
	return ret.Result, nil
}

// GetResource returns a resource defined by a built-in Kubernetes group/version/kind, name and namespace.
// For example, GetResource("v1/Pod", "foo", "") would return a Pod called "foo" from the "default" namespace.
func (c *Chart) GetResource(gvk, name, namespace string) pulumi.AnyOutput {
	id := name
	if len(namespace) > 0 && namespace != "default" {
		id = fmt.Sprintf("%s/%s", namespace, name)
	}
	key := fmt.Sprintf("%s::%s", gvk, id)
	return c.Resources.ApplyT(func(x interface{}) interface{} {
		resources := x.(map[string]pulumi.Resource)
		return resources[key]
	}).(pulumi.AnyOutput)
}
