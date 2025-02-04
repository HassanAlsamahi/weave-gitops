package kube_test

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	wego "github.com/weaveworks/weave-gitops/api/v1alpha1"
	"k8s.io/apimachinery/pkg/types"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"github.com/weaveworks/weave-gitops/pkg/kube"
	"github.com/weaveworks/weave-gitops/pkg/runner/runnerfakes"
)

var (
	runner     *runnerfakes.FakeRunner
	kubeClient *kube.KubeClient
)

var _ = BeforeEach(func() {
	runner = &runnerfakes.FakeRunner{}

	kubeClient = kube.New(runner)
})

var _ = Describe("Apply", func() {
	It("applies manifests", func() {
		runner.RunWithStdinStub = func(s1 string, s2 []string, b []byte) ([]byte, error) {
			return []byte("out"), nil
		}

		err := kubeClient.Apply(context.Background(), []byte("manifests"), "wego-system")
		Expect(err).ShouldNot(HaveOccurred())

		cmd, args, manifests := runner.RunWithStdinArgsForCall(0)
		Expect(cmd).To(Equal("kubectl"))

		Expect(strings.Join(args, " ")).To(Equal("apply --namespace wego-system -f -"))
		Expect(manifests).To(Equal([]byte("manifests")))
	})
})

var _ = Describe("Delete", func() {
	It("delete manifests", func() {
		runner.RunWithStdinStub = func(s1 string, s2 []string, b []byte) ([]byte, error) {
			return []byte("out"), nil
		}

		err := kubeClient.Delete(context.Background(), []byte("manifests"))
		Expect(err).ShouldNot(HaveOccurred())

		cmd, args, manifests := runner.RunWithStdinArgsForCall(0)
		Expect(cmd).To(Equal("kubectl"))

		Expect(strings.Join(args, " ")).To(Equal("delete -f -"))
		Expect(manifests).To(Equal([]byte("manifests")))
	})
})

var _ = Describe("GetClusterStatus", func() {
	It("returns wego is installed", func() {
		runner.RunStub = func(cmd string, args ...string) ([]byte, error) {
			if strings.Join(args, " ") == "get crd apps.wego.weave.works" {
				return []byte("out"), nil
			}

			return []byte("error"), fmt.Errorf("error")
		}

		status := kubeClient.GetClusterStatus(context.Background())
		Expect(status).To(Equal(kube.WeGOInstalled))
	})

	It("returns flux is installed", func() {
		runner.RunStub = func(cmd string, args ...string) ([]byte, error) {
			if strings.Join(args, " ") == "get namespace flux-system" {
				return []byte("out"), nil
			}

			return []byte("error"), fmt.Errorf("error")
		}

		status := kubeClient.GetClusterStatus(context.Background())
		Expect(status).To(Equal(kube.FluxInstalled))
	})

	It("returns unmodified cluster", func() {
		runner.RunStub = func(cmd string, args ...string) ([]byte, error) {
			if strings.Join(args, " ") == "get deployment coredns -n kube-system" {
				return []byte("out"), nil
			}

			return []byte("error"), fmt.Errorf("error")
		}

		status := kubeClient.GetClusterStatus(context.Background())
		Expect(status).To(Equal(kube.Unmodified))
	})

	It("returns unknown when it cant talk to the cluster", func() {
		runner.RunStub = func(cmd string, args ...string) ([]byte, error) {
			return []byte("was refused - did you specify the right host or port?"), fmt.Errorf("error")
		}

		status := kubeClient.GetClusterStatus(context.Background())
		Expect(status).To(Equal(kube.Unknown))
	})
})

var _ = Describe("GetClusterName", func() {
	It("returns cluster name", func() {
		runner.RunStub = func(cmd string, args ...string) ([]byte, error) {
			return []byte("cluster-name\n"), nil
		}

		out, err := kubeClient.GetClusterName(context.Background())
		Expect(err).ShouldNot(HaveOccurred())
		Expect(string(out)).To(Equal("cluster-name"))

		cmd, args := runner.RunArgsForCall(0)
		Expect(cmd).To(Equal("kubectl"))

		Expect(strings.Join(args, " ")).To(Equal("config current-context"))
	})
})

var _ = Describe("FixInvalidClusterName", func() {
	DescribeTable("checks to verify that cluster names are sanitized",
		func(invalid string, valid string, expected bool) {
			runner.RunStub = func(cmd string, args ...string) ([]byte, error) {
				return []byte(invalid), nil
			}

			out, err := kubeClient.GetClusterName(context.Background())
			Expect(err).ShouldNot(HaveOccurred())
			Expect(out == valid).Should(Equal(expected))
		},
		Entry("underscore to dash", "cluster_name\n", "cluster-name", true),
		Entry("remove invalid @", "Cluster@name\n", "clustername", true),
		Entry("remove front -'s", "--cluster-name\n", "cluster-name", true),
		Entry("remove back -'s", "cluster-name-\n", "cluster-name", true),
		Entry("remove (str)$", "cluster-name$\n", "cluster-name", true),
		Entry("remove $(str)", "$cluster-name\n", "cluster-name", true),
		Entry("remove embed @", "cluster-name@1\n", "cluster-name1", true),
		Entry("remove invalid chars", "1@#$%^&*(_+w2\n", "1-w2", true),
	)
})

var _ = Describe("FluxPresent", func() {
	It("looks for flux-system namespace", func() {
		_, err := kubeClient.FluxPresent(context.Background())
		Expect(err).ShouldNot(HaveOccurred())

		cmd, args := runner.RunArgsForCall(0)
		Expect(cmd).To(Equal("kubectl"))

		Expect(strings.Join(args, " ")).To(Equal("get namespace flux-system"))
	})

	It("returns true if flux is present in the cluster", func() {
		runner.RunStub = func(cmd string, args ...string) ([]byte, error) {
			return []byte("namespace"), nil
		}

		present, err := kubeClient.FluxPresent(context.Background())
		Expect(err).ShouldNot(HaveOccurred())
		Expect(present).To(Equal(true))
	})

	It("returns false if flux is not present in the cluster", func() {
		runner.RunStub = func(cmd string, args ...string) ([]byte, error) {
			return []byte("not found"), fmt.Errorf("error")
		}

		present, err := kubeClient.FluxPresent(context.Background())
		Expect(err).ShouldNot(HaveOccurred())
		Expect(present).To(Equal(false))
	})
})

var _ = Describe("GetApplication", func() {
	It("gets an application by name", func() {
		res, err := json.Marshal(&wego.Application{
			Spec: wego.ApplicationSpec{
				Path: "some/path",
				URL:  "example.com/some-org/some-repo",
			},
		})
		Expect(err).ShouldNot(HaveOccurred())

		runner.RunStub = func(cmd string, args ...string) ([]byte, error) {
			return res, nil
		}

		out, err := kubeClient.GetApplication(context.Background(), types.NamespacedName{Name: "my-app", Namespace: kube.WeGONamespace})

		Expect(err).ShouldNot(HaveOccurred())
		Expect(out.Spec.Path).To(Equal("some/path"))
	})

	It("get all applications", func() {

		appsList := &wego.ApplicationList{Items: []wego.Application{
			{
				Spec: wego.ApplicationSpec{
					Path: "some/path0",
					URL:  "example.com/some-org/some-repo0",
				},
			},
			{
				Spec: wego.ApplicationSpec{
					Path: "some/path1",
					URL:  "example.com/some-org/some-repo1",
				},
			},
		}}

		res, err := json.Marshal(appsList)
		Expect(err).ShouldNot(HaveOccurred())

		runner.RunStub = func(cmd string, args ...string) ([]byte, error) {
			return res, nil
		}

		apps, err := kubeClient.GetApplications(context.Background(), "wego-system")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(2).To(Equal(len(apps)))

		for i, a := range appsList.Items {
			app := (apps)[i]
			Expect(a.Name).To(Equal(app.Name))
			Expect(a.Spec.Path).To(Equal(app.Spec.Path))
			Expect(a.Spec.URL).To(Equal(app.Spec.URL))
		}

	})
})
