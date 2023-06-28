package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	zv1 "github.com/zalando-incubator/stackset-controller/pkg/apis/zalando.org/v1"
	autoscaling "k8s.io/api/autoscaling/v2"
	autoscalingv2beta1 "k8s.io/api/autoscaling/v2beta1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func generateAutoscalerStub(minReplicas, maxReplicas int32) StackContainer {
	return StackContainer{
		Stack: &zv1.Stack{
			ObjectMeta: metav1.ObjectMeta{
				Name: "stackset-v1",
			},
			Spec: zv1.StackSpec{
				Autoscaler: &zv1.Autoscaler{
					MinReplicas: &minReplicas,
					MaxReplicas: maxReplicas,
					Metrics:     []zv1.AutoscalerMetrics{},
				},
			},
		},
		stacksetName:        "stackset",
		actualTrafficWeight: 100.0,
	}
}

func generateAutoscalerCPU(minReplicas, maxReplicas, utilization int32, containerName string) StackContainer {
	container := generateAutoscalerStub(minReplicas, maxReplicas)
	container.Stack.Spec.Autoscaler.Metrics = append(
		container.Stack.Spec.Autoscaler.Metrics, zv1.AutoscalerMetrics{
			Type:               zv1.CPUAutoscalerMetric,
			AverageUtilization: &utilization,
			Container:          containerName,
		})
	return container
}

func generateAutoscalerMemory(minReplicas, maxReplicas, utilization int32, containerName string) StackContainer {
	container := generateAutoscalerStub(minReplicas, maxReplicas)
	container.Stack.Spec.Autoscaler.Metrics = append(
		container.Stack.Spec.Autoscaler.Metrics, zv1.AutoscalerMetrics{
			Type:               zv1.MemoryAutoscalerMetric,
			AverageUtilization: &utilization,
			Container:          containerName,
		})
	return container
}

func generateAutoscalerSQS(minReplicas, maxReplicas, utilization int32, queueName, queueRegion string) StackContainer {
	container := generateAutoscalerStub(minReplicas, maxReplicas)
	container.Stack.Spec.Autoscaler.Metrics = append(
		container.Stack.Spec.Autoscaler.Metrics, zv1.AutoscalerMetrics{
			Type: zv1.AmazonSQSAutoscalerMetric,
			Queue: &zv1.MetricsQueue{
				Name:   queueName,
				Region: queueRegion,
			},
			Average: resource.NewQuantity(int64(utilization), resource.DecimalSI),
		},
	)
	return container
}
func generateAutoscalerZMON(minReplicas, maxReplicas, utilization int32, checkID, key, application, duration string, aggregators []zv1.ZMONMetricAggregatorType) StackContainer {
	container := generateAutoscalerStub(minReplicas, maxReplicas)
	container.Stack.Spec.Autoscaler.Metrics = append(
		container.Stack.Spec.Autoscaler.Metrics, zv1.AutoscalerMetrics{
			Type: zv1.ZMONAutoscalerMetric,
			ZMON: &zv1.MetricsZMON{
				CheckID:     checkID,
				Key:         key,
				Duration:    duration,
				Aggregators: aggregators,
				Tags: map[string]string{
					"application": application,
				},
			},
			Average: resource.NewQuantity(int64(utilization), resource.DecimalSI),
		},
	)
	return container
}

func generateAutoscalerScalingSchedule(minReplicas, maxReplicas, average int32, name string) StackContainer {
	container := generateAutoscalerStub(minReplicas, maxReplicas)
	container.Stack.Spec.Autoscaler.Metrics = append(
		container.Stack.Spec.Autoscaler.Metrics, zv1.AutoscalerMetrics{
			Type: zv1.ScalingScheduleMetric,
			ScalingSchedule: &zv1.MetricsScalingSchedule{
				Name: name,
			},
			Average: resource.NewQuantity(int64(average), resource.DecimalSI),
		},
	)
	return container
}

func generateAutoscalerClusterScalingSchedule(minReplicas, maxReplicas, average int32, name string) StackContainer {
	container := generateAutoscalerStub(minReplicas, maxReplicas)
	container.Stack.Spec.Autoscaler.Metrics = append(
		container.Stack.Spec.Autoscaler.Metrics, zv1.AutoscalerMetrics{
			Type: zv1.ClusterScalingScheduleMetric,
			ClusterScalingSchedule: &zv1.MetricsClusterScalingSchedule{
				Name: name,
			},
			Average: resource.NewQuantity(int64(average), resource.DecimalSI),
		},
	)
	return container
}

func generateAutoscalerPodJson(minReplicas, maxReplicas, utilization, port int32, name, path, key string) StackContainer {
	container := generateAutoscalerStub(minReplicas, maxReplicas)
	container.Stack.Spec.Autoscaler.Metrics = append(
		container.Stack.Spec.Autoscaler.Metrics, zv1.AutoscalerMetrics{
			Type: zv1.PodJSONAutoscalerMetric,
			Endpoint: &zv1.MetricsEndpoint{
				Path: path,
				Name: name,
				Key:  key,
				Port: port,
			},
			Average: resource.NewQuantity(int64(utilization), resource.DecimalSI),
		},
	)
	return container
}
func generateAutoscalerIngress(minReplicas, maxReplicas, utilization int32) StackContainer {
	container := generateAutoscalerStub(minReplicas, maxReplicas)
	container.Stack.Spec.Autoscaler.Metrics = append(
		container.Stack.Spec.Autoscaler.Metrics, zv1.AutoscalerMetrics{
			Type:    zv1.IngressAutoscalerMetric,
			Average: resource.NewQuantity(int64(utilization), resource.DecimalSI),
		},
	)
	return container
}

func generateAutoscalerRouteGroup(minReplicas, maxReplicas, utilization int32) StackContainer {
	container := generateAutoscalerStub(minReplicas, maxReplicas)
	container.Stack.Spec.Autoscaler.Metrics = append(
		container.Stack.Spec.Autoscaler.Metrics, zv1.AutoscalerMetrics{
			Type:    zv1.RouteGroupAutoscalerMetric,
			Average: resource.NewQuantity(int64(utilization), resource.DecimalSI),
		},
	)
	return container
}

func generateAutoscalerExternalRPS(minReplicas, maxReplicas, utilization int32, weight float64, hosts []string) StackContainer {
	container := generateAutoscalerStub(minReplicas, maxReplicas)
	container.actualTrafficWeight = weight
	container.Stack.Spec.Autoscaler.Metrics = append(
		container.Stack.Spec.Autoscaler.Metrics, zv1.AutoscalerMetrics{
			Type: zv1.ExternalRPSMetric,
			RequestsPerSecond: &zv1.MetricsRequestsPerSecond{
				Name:      "a-rps-metric",
				Hostnames: hosts,
			},
			Average: resource.NewQuantity(int64(utilization), resource.DecimalSI),
		},
	)
	return container
}

func TestStackSetController_ReconcileAutoscalersCPU(t *testing.T) {
	ssc := generateAutoscalerCPU(1, 10, 80, "")
	hpa, err := ssc.GenerateHPA()
	require.NoError(t, err, "failed to create an HPA")
	require.NotNil(t, hpa, "hpa not generated")
	require.Equal(t, int32(1), *hpa.Spec.MinReplicas, "min replicas not generated correctly")
	require.Equal(t, int32(10), hpa.Spec.MaxReplicas, "max replicas generated incorrectly")
	require.Len(t, hpa.Spec.Metrics, 1, "expected HPA to have 1 metric. instead got %d", len(hpa.Spec.Metrics))
	cpuMetric := hpa.Spec.Metrics[0]
	require.Equal(t, cpuMetric.Type, autoscaling.ResourceMetricSourceType)
	require.Equal(t, cpuMetric.Resource.Name, corev1.ResourceCPU)
	require.Equal(t, *cpuMetric.Resource.Target.AverageUtilization, int32(80))

	ssc = generateAutoscalerCPU(1, 10, 80, "container-x")
	hpa, err = ssc.GenerateHPA()
	require.NoError(t, err, "failed to create an HPA")
	require.NotNil(t, hpa, "hpa not generated")
	require.Equal(t, int32(1), *hpa.Spec.MinReplicas, "min replicas not generated correctly")
	require.Equal(t, int32(10), hpa.Spec.MaxReplicas, "max replicas generated incorrectly")
	require.Len(t, hpa.Spec.Metrics, 1, "expected HPA to have 1 metric. instead got %d", len(hpa.Spec.Metrics))
	cpuMetric = hpa.Spec.Metrics[0]
	require.Equal(t, autoscaling.ContainerResourceMetricSourceType, cpuMetric.Type)
	require.Equal(t, corev1.ResourceCPU, cpuMetric.ContainerResource.Name)
	require.Equal(t, int32(80), *cpuMetric.ContainerResource.Target.AverageUtilization)
	require.Equal(t, "container-x", cpuMetric.ContainerResource.Container)
}

func TestStackSetController_ReconcileAutoscalersMemory(t *testing.T) {
	ssc := generateAutoscalerMemory(1, 10, 80, "")
	hpa, err := ssc.GenerateHPA()
	require.NoError(t, err, "failed to create an HPA")
	require.NotNil(t, hpa, "hpa not generated")
	require.Equal(t, int32(1), *hpa.Spec.MinReplicas, "min replicas not generated correctly")
	require.Equal(t, int32(10), hpa.Spec.MaxReplicas, "max replicas generated incorrectly")
	require.Len(t, hpa.Spec.Metrics, 1, "expected HPA to have 1 metric. instead got %d", len(hpa.Spec.Metrics))
	memoryMetric := hpa.Spec.Metrics[0]
	require.Equal(t, memoryMetric.Type, autoscaling.ResourceMetricSourceType)
	require.Equal(t, memoryMetric.Resource.Name, corev1.ResourceMemory)
	require.Equal(t, *memoryMetric.Resource.Target.AverageUtilization, int32(80))

	ssc = generateAutoscalerMemory(1, 10, 80, "container-x")
	hpa, err = ssc.GenerateHPA()
	require.NoError(t, err, "failed to create an HPA")
	require.NotNil(t, hpa, "hpa not generated")
	require.Equal(t, int32(1), *hpa.Spec.MinReplicas, "min replicas not generated correctly")
	require.Equal(t, int32(10), hpa.Spec.MaxReplicas, "max replicas generated incorrectly")
	require.Len(t, hpa.Spec.Metrics, 1, "expected HPA to have 1 metric. instead got %d", len(hpa.Spec.Metrics))
	memoryMetric = hpa.Spec.Metrics[0]
	require.Equal(t, autoscaling.ContainerResourceMetricSourceType, memoryMetric.Type)
	require.Equal(t, corev1.ResourceMemory, memoryMetric.ContainerResource.Name)
	require.Equal(t, int32(80), *memoryMetric.ContainerResource.Target.AverageUtilization)
	require.Equal(t, "container-x", memoryMetric.ContainerResource.Container)
}

func TestStackSetController_ReconcileAutoscalersSQS(t *testing.T) {
	ssc := generateAutoscalerSQS(1, 10, 80, "test-queue", "test-region")
	hpa, err := ssc.GenerateHPA()
	require.NoError(t, err, "failed to create an HPA")
	require.NotNil(t, hpa, "hpa not generated")
	require.Equal(t, int32(1), *hpa.Spec.MinReplicas, "min replicas not generated correctly")
	require.Equal(t, int32(10), hpa.Spec.MaxReplicas, "max replicas generated incorrectly")
	require.Len(t, hpa.Spec.Metrics, 1, "expected HPA to have 1 metric. instead got %d", len(hpa.Spec.Metrics))
	externalMetric := hpa.Spec.Metrics[0]
	require.Equal(t, externalMetric.Type, autoscaling.ExternalMetricSourceType)
	require.Equal(t, externalMetric.External.Metric.Name, "sqs-queue-length")
	require.Equal(t, externalMetric.External.Metric.Selector.MatchLabels["queue-name"], "test-queue")
	require.Equal(t, externalMetric.External.Metric.Selector.MatchLabels["queue-name"], "test-queue")
	require.Equal(t, externalMetric.External.Target.AverageValue.Value(), int64(80))
}

func TestStackSetController_ReconcileAutoscalersPodJson(t *testing.T) {
	ssc := generateAutoscalerPodJson(1, 10, 80, 8080, "current-load", "/metrics", "$.current-load.counter")
	hpa, err := ssc.GenerateHPA()
	require.NoError(t, err, "failed to create an HPA")
	require.NotNil(t, hpa, "hpa not generated")
	require.Equal(t, int32(1), *hpa.Spec.MinReplicas, "min replicas not generated correctly")
	require.Equal(t, int32(10), hpa.Spec.MaxReplicas, "max replicas generated incorrectly")
	require.Len(t, hpa.Spec.Metrics, 1, "expected HPA to have 1 metric. instead got %d", len(hpa.Spec.Metrics))
	podMetrics := hpa.Spec.Metrics[0]
	require.Equal(t, podMetrics.Type, autoscaling.PodsMetricSourceType)
	require.Equal(t, hpa.Annotations["metric-config.pods.current-load.json-path/json-key"], "$.current-load.counter")
	require.Equal(t, hpa.Annotations["metric-config.pods.current-load.json-path/path"], "/metrics")
	require.Equal(t, hpa.Annotations["metric-config.pods.current-load.json-path/port"], "8080")
	require.Equal(t, podMetrics.Pods.Target.AverageValue.Value(), int64(80))
	require.Equal(t, podMetrics.Pods.Metric.Name, "current-load")
}

func TestStackSetController_ReconcileAutoscalersIngress(t *testing.T) {
	ssc := generateAutoscalerIngress(1, 10, 80)
	hpa, err := ssc.GenerateHPA()
	require.NoError(t, err, "failed to create an HPA")
	require.NotNil(t, hpa, "hpa not generated")
	require.Equal(t, int32(1), *hpa.Spec.MinReplicas, "min replicas not generated correctly")
	require.Equal(t, int32(10), hpa.Spec.MaxReplicas, "max replicas generated incorrectly")
	require.Len(t, hpa.Spec.Metrics, 1, "expected HPA to have 1 metric. instead got %d", len(hpa.Spec.Metrics))
	ingressMetrics := hpa.Spec.Metrics[0]
	require.Equal(t, autoscaling.ObjectMetricSourceType, ingressMetrics.Type)
	require.Equal(t, int64(80), ingressMetrics.Object.Target.AverageValue.Value())
	require.Equal(t, ingressMetrics.Object.Metric.Name, "requests-per-second,stackset-v1")
	// require.Equal(t, ingressMetrics.Object.Metric.Selector.MatchLabels, map[string]string{"backend": "stackset-v1"})
}

func TestStackSetController_ReconcileAutoscalersRouteGroup(t *testing.T) {
	ssc := generateAutoscalerRouteGroup(1, 10, 80)
	hpa, err := ssc.GenerateHPA()
	require.NoError(t, err, "failed to create an HPA")
	require.NotNil(t, hpa, "hpa not generated")
	require.Equal(t, int32(1), *hpa.Spec.MinReplicas, "min replicas not generated correctly")
	require.Equal(t, int32(10), hpa.Spec.MaxReplicas, "max replicas generated incorrectly")
	require.Len(t, hpa.Spec.Metrics, 1, "expected HPA to have 1 metric. instead got %d", len(hpa.Spec.Metrics))
	metrics := hpa.Spec.Metrics[0]
	require.Equal(t, autoscaling.ObjectMetricSourceType, metrics.Type)
	require.Equal(t, int64(80), metrics.Object.Target.AverageValue.Value())
	require.Equal(t, metrics.Object.Metric.Name, "requests-per-second")
	require.Equal(t, metrics.Object.Metric.Selector.MatchLabels, map[string]string{"backend": "stackset-v1"})
}

func TestStackSetController_ReconcileAutoscalersZMON(t *testing.T) {
	ssc := generateAutoscalerZMON(1, 10, 80, "1234", "key", "app", "10m", []zv1.ZMONMetricAggregatorType{"avg", "max"})
	hpa, err := ssc.GenerateHPA()
	require.NoError(t, err, "failed to create an HPA")
	require.NotNil(t, hpa, "hpa not generated")
	require.Equal(t, int32(1), *hpa.Spec.MinReplicas, "min replicas not generated correctly")
	require.Equal(t, int32(10), hpa.Spec.MaxReplicas, "max replicas generated incorrectly")
	require.Len(t, hpa.Spec.Metrics, 1, "expected HPA to have 1 metric. instead got %d", len(hpa.Spec.Metrics))
	externalMetric := hpa.Spec.Metrics[0]
	require.Equal(t, externalMetric.Type, autoscaling.ExternalMetricSourceType)
	require.Equal(t, externalMetric.External.Metric.Name, zmonCheckMetricName)
	require.Equal(t, externalMetric.External.Metric.Selector.MatchLabels[zmonCheckCheckIDTag], "1234")
	require.Equal(t, externalMetric.External.Metric.Selector.MatchLabels[zmonCheckDurationTag], "10m")
	require.Equal(t, externalMetric.External.Metric.Selector.MatchLabels[zmonCheckAggregatorsTag], "avg,max")
	require.Equal(t, hpa.Annotations[zmonCheckKeyAnnotation], "key")
	require.Equal(t, hpa.Annotations[zmonCheckTagAnnotationPrefix+"application"], "app")
	require.Equal(t, externalMetric.External.Target.AverageValue.Value(), int64(80))
}

func TestStackSetController_ReconcileAutoscalersScalingSchedule(t *testing.T) {
	average := 80
	name := "scaling-schedule-name"

	validateHpa := func(t *testing.T, kind string, ssc StackContainer) {
		hpa, err := ssc.GenerateHPA()
		require.NoError(t, err, "failed to create an HPA")
		require.NotNil(t, hpa, "hpa not generated")
		require.Equal(t, int32(1), *hpa.Spec.MinReplicas, "min replicas not generated correctly")
		require.Equal(t, int32(10), hpa.Spec.MaxReplicas, "max replicas generated incorrectly")
		require.Len(t, hpa.Spec.Metrics, 1, "expected HPA to have 1 metric. instead got %d", len(hpa.Spec.Metrics))
		objectMetric := hpa.Spec.Metrics[0]
		require.Equal(t, autoscaling.ObjectMetricSourceType, objectMetric.Type)
		require.Equal(t, name, objectMetric.Object.Metric.Name)
		require.Equal(t, name, objectMetric.Object.DescribedObject.Name)
		require.Equal(t, kind, objectMetric.Object.DescribedObject.Kind)
		require.Equal(t, scalingScheduleAPIVersion, objectMetric.Object.DescribedObject.APIVersion)
		require.Equal(t, int64(average), objectMetric.Object.Target.AverageValue.Value())
	}

	ssc := generateAutoscalerScalingSchedule(1, 10, int32(average), name)
	t.Run("generate ScalingSchedule HPA", func(t *testing.T) {
		validateHpa(t, "ScalingSchedule", ssc)
	})

	ssc = generateAutoscalerClusterScalingSchedule(1, 10, int32(average), name)
	t.Run("generate ClusterScalingSchedule HPA", func(t *testing.T) {
		validateHpa(t, "ClusterScalingSchedule", ssc)
	})

}

func TestStackSetController_ReconcileAutoscalersExternalRPS(t *testing.T) {
	validateHpa := func(tt *testing.T, expectedHosts string, weight float64, average int32, ssc StackContainer) {
		hpa, err := ssc.GenerateHPA()
		require.NoError(tt, err, "failed to create an HPA")
		require.NotNil(tt, hpa, "hpa not generated")
		require.Equal(tt, int32(1), *hpa.Spec.MinReplicas, "min replicas not generated correctly")
		require.Equal(tt, int32(10), hpa.Spec.MaxReplicas, "max replicas generated incorrectly")
		require.Len(tt, hpa.Spec.Metrics, 1, "expected HPA to have 1 metric. instead got %d", len(hpa.Spec.Metrics))
		externalMetric := hpa.Spec.Metrics[0]
		require.Equal(tt, autoscaling.ExternalMetricSourceType, externalMetric.Type)
		require.Equal(tt, "a-rps-metric", externalMetric.External.Metric.Name)
		require.Equal(tt, "requests-per-second", externalMetric.External.Metric.Selector.MatchLabels["type"])
		require.Equal(tt, autoscaling.AverageValueMetricType, externalMetric.External.Target.Type)
		require.Equal(tt, int64(average), externalMetric.External.Target.AverageValue.Value())
		require.Equal(tt, expectedHosts, hpa.Annotations["metric-config.a-rps-metric.requests-per-second/hostnames"])
		require.Equal(tt, fmt.Sprintf("%d", int(weight)), hpa.Annotations["metric-config.a-rps-metric.requests-per-second/weight"])
	}

	for _, tc := range []struct {
		description   string
		average       int32
		hosts         []string
		expectedHosts string
		weight        float64
	}{
		{
			description:   "No weight; single host",
			average:       80,
			hosts:         []string{"foo.bar.baz"},
			expectedHosts: "foo.bar.baz",
			weight:        0.0,
		},
		{
			description:   "No weight; multiple hosts",
			average:       80,
			hosts:         []string{"foo.bar.baz", "foo.bar.bazzy"},
			expectedHosts: "foo.bar.baz,foo.bar.bazzy",
			weight:        0.0,
		},
		{
			description:   "With half weight; single host",
			average:       40,
			hosts:         []string{"foo.bar.baz"},
			expectedHosts: "foo.bar.baz",
			weight:        50.0,
		},
		{
			description:   "With full weight; single host",
			average:       80,
			hosts:         []string{"foo.bar.baz"},
			expectedHosts: "foo.bar.baz",
			weight:        100.0,
		},
	} {
		t.Run(tc.description, func(tt *testing.T) {
			ssc := generateAutoscalerExternalRPS(1, 10, tc.average, tc.weight, tc.hosts)
			validateHpa(tt, tc.expectedHosts, tc.weight, tc.average, ssc)
		})
	}
}

func TestCPUMetricValid(t *testing.T) {
	var utilization int32 = 80
	metrics := zv1.AutoscalerMetrics{Type: "cpu", AverageUtilization: &utilization}
	metric, err := cpuMetric(metrics)
	require.NoError(t, err, "could not create hpa metric")
	require.Equal(t, metric.Resource.Name, corev1.ResourceCPU)
}

func TestExternalRPSMetricInvalid(t *testing.T) {
	for _, tc := range []struct {
		desc string
		m    zv1.AutoscalerMetrics
	}{
		{
			desc: "No average value",
			m: zv1.AutoscalerMetrics{
				Type: "RequestPerSecond",
				RequestsPerSecond: &zv1.MetricsRequestsPerSecond{
					Name:      "a-test",
					Hostnames: []string{"foo.bar.baz"},
				},
				Average: nil,
			},
		},
		{
			desc: "No RequestsPerSecond value",
			m: zv1.AutoscalerMetrics{
				Type:    "RequestPerSecond",
				Average: resource.NewQuantity(80, resource.DecimalSI),
			},
		},
		{
			desc: "No RequestsPerSecond.Hostnames value",
			m: zv1.AutoscalerMetrics{
				Type: "RequestPerSecond",
				RequestsPerSecond: &zv1.MetricsRequestsPerSecond{
					Name: "a-test",
				},
				Average: resource.NewQuantity(80, resource.DecimalSI),
			},
		},
		{
			desc: "No RequestsPerSecond.Name value",
			m: zv1.AutoscalerMetrics{
				Type: "RequestPerSecond",
				RequestsPerSecond: &zv1.MetricsRequestsPerSecond{
					Hostnames: []string{"foo.bar.baz"},
				},
				Average: resource.NewQuantity(80, resource.DecimalSI),
			},
		},
	} {
		t.Run(tc.desc, func(tt *testing.T) {
			_, _, err := externalRPSMetric(tc.m, 100.0)
			require.Error(tt, err)
		})
	}
}

func TestCPUMetricInValid(t *testing.T) {
	metrics := zv1.AutoscalerMetrics{Type: "cpu", AverageUtilization: nil}
	_, err := cpuMetric(metrics)
	require.Error(t, err, "created metric even when utilization not specified")
}
func TestMemoryMetricValid(t *testing.T) {
	var utilization int32 = 80
	metrics := zv1.AutoscalerMetrics{Type: "memory", AverageUtilization: &utilization}
	metric, err := memoryMetric(metrics)
	require.NoError(t, err, "could not create hpa metric")
	require.Equal(t, metric.Resource.Name, corev1.ResourceMemory)
}

func TestMemoryMetricInValid(t *testing.T) {
	metrics := zv1.AutoscalerMetrics{Type: "memory", AverageUtilization: nil}
	_, err := memoryMetric(metrics)
	require.Error(t, err, "created metric even when utilization not specified")
}

func TestPodJsonMetricInvalid(t *testing.T) {
	endpoints := []zv1.MetricsEndpoint{
		{
			Path: "/metrics",
			Port: 8080,
			Key:  "$.metrics_key",
		},
		{
			Path: "/metrics",
			Port: 8080,
			Name: "metric-name",
		},
		{
			Path: "/metrics",
			Key:  "$.metrics_key",
			Name: "metric-name",
		},
		{
			Port: 8080,
			Key:  "$.metrics_key",
			Name: "metric-name",
		},
	}
	for _, e := range endpoints {
		metrics := zv1.AutoscalerMetrics{Type: zv1.PodJSONAutoscalerMetric, Endpoint: &e}
		_, _, err := podJsonMetric(metrics)
		require.Error(t, err, "created metric with invalid configuration")
	}
}

func TestZMONMetricInvalid(t *testing.T) {
	onemilli := resource.MustParse("1m")
	for _, tc := range []struct {
		name    string
		metrics zv1.AutoscalerMetrics
	}{
		{
			name:    "missing average",
			metrics: zv1.AutoscalerMetrics{Type: zv1.ZMONAutoscalerMetric, Average: nil},
		},
		{
			name:    "missing zmon definition",
			metrics: zv1.AutoscalerMetrics{Type: zv1.ZMONAutoscalerMetric, Average: &onemilli},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, _, err := zmonMetric(tc.metrics, "stack-name", "namespace")
			require.Errorf(t, err, "created metric with invalid configuration")
		})
	}
}

func TestScalingScheduleMetricInvalid(t *testing.T) {
	onemilli := resource.MustParse("1m")
	for _, tc := range []struct {
		name    string
		metrics zv1.AutoscalerMetrics
	}{
		{
			name:    "missing average",
			metrics: zv1.AutoscalerMetrics{Type: zv1.ScalingScheduleMetric, Average: nil},
		},
		{
			name:    "missing average",
			metrics: zv1.AutoscalerMetrics{Type: zv1.ClusterScalingScheduleMetric, Average: nil},
		},
		{
			name:    "missing ScalingSchedule",
			metrics: zv1.AutoscalerMetrics{Type: zv1.ScalingScheduleMetric, Average: &onemilli},
		},
		{
			name:    "missing ClusterScalingSchedule",
			metrics: zv1.AutoscalerMetrics{Type: zv1.ClusterScalingScheduleMetric, Average: &onemilli},
		},
		{
			name: "missing ScalingSchedule name",
			metrics: zv1.AutoscalerMetrics{
				Type:            zv1.ScalingScheduleMetric,
				Average:         &onemilli,
				ScalingSchedule: &zv1.MetricsScalingSchedule{Name: ""}},
		},
		{
			name: "missing ClusterScalingSchedule name",
			metrics: zv1.AutoscalerMetrics{
				Type:                   zv1.ClusterScalingScheduleMetric,
				Average:                &onemilli,
				ClusterScalingSchedule: &zv1.MetricsClusterScalingSchedule{Name: ""}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var err error
			switch tc.metrics.Type {
			case zv1.ClusterScalingScheduleMetric:
				_, err = clusterScalingScheduleMetric(tc.metrics, "stack-name", "namespace")
			case zv1.ScalingScheduleMetric:
				_, err = scalingScheduleMetric(tc.metrics, "stack-name", "namespace")
			}
			require.Errorf(t, err, "created metric with invalid configuration")
		})
	}
}

func TestIngressMetricInvalid(t *testing.T) {
	metrics := zv1.AutoscalerMetrics{Type: zv1.IngressAutoscalerMetric, Average: nil}
	_, err := ingressMetric(metrics, "stack-name", "test-stack")
	require.Errorf(t, err, "created metric with invalid configuration")
}

func TestSortingMetrics(t *testing.T) {
	container := generateAutoscalerStub(1, 10)
	metrics := []zv1.AutoscalerMetrics{
		{Type: zv1.CPUAutoscalerMetric, AverageUtilization: pint32(50)},
		{Type: zv1.IngressAutoscalerMetric, Average: resource.NewQuantity(10, resource.DecimalSI)},
		{Type: zv1.PodJSONAutoscalerMetric, Average: resource.NewQuantity(10, resource.DecimalSI), Endpoint: &zv1.MetricsEndpoint{Name: "abc", Path: "/metrics", Port: 1222, Key: "test.abc"}},
		{Type: zv1.AmazonSQSAutoscalerMetric, Average: resource.NewQuantity(10, resource.DecimalSI), Queue: &zv1.MetricsQueue{Name: "test", Region: "region"}},
	}
	container.Stack.Spec.Autoscaler.Metrics = metrics
	hpa, err := container.GenerateHPA()
	require.NoError(t, err, "failed to create an HPA")
	require.NotNil(t, hpa, "hpa not generated")
	require.Len(t, hpa.Spec.Metrics, 4)
	require.EqualValues(t, autoscaling.ExternalMetricSourceType, hpa.Spec.Metrics[0].Type)
	require.EqualValues(t, autoscaling.ObjectMetricSourceType, hpa.Spec.Metrics[1].Type)
	require.EqualValues(t, autoscaling.PodsMetricSourceType, hpa.Spec.Metrics[2].Type)
	require.EqualValues(t, autoscaling.ResourceMetricSourceType, hpa.Spec.Metrics[3].Type)
}

func pint32(val int) *int32 {
	return &[]int32{int32(val)}[0]
}

func generateHPA(minReplicas, maxReplicas int32) StackContainer {
	return StackContainer{
		Stack: &zv1.Stack{
			ObjectMeta: metav1.ObjectMeta{
				Name: "stackset-v1",
			},
			Spec: zv1.StackSpec{
				HorizontalPodAutoscaler: &zv1.HorizontalPodAutoscaler{
					MinReplicas: &minReplicas,
					MaxReplicas: maxReplicas,
					Metrics:     []autoscalingv2beta1.MetricSpec{},
				},
			},
		},
		stacksetName: "stackset",
	}
}

func TestStackSetController_ReconcileHPA(t *testing.T) {
	ssc := generateHPA(1, 10)
	hpa, err := ssc.GenerateHPA()
	require.NoError(t, err, "failed to create an HPA")
	require.NotNil(t, hpa, "hpa not generated")
	require.Equal(t, int32(1), *hpa.Spec.MinReplicas, "min replicas not generated correctly")
	require.Equal(t, int32(10), hpa.Spec.MaxReplicas, "max replicas generated incorrectly")
}
