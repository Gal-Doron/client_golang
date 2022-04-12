// Copyright 2018 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gocollector

import (
	"runtime"
	"runtime/debug"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func goRuntimeMemStats() memStatsMetrics {
	return memStatsMetrics{
		{
			desc: prometheus.NewDesc(
				memstatNamespace("alloc_bytes"),
				"Number of bytes allocated and still in use.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.Alloc) },
			valType: prometheus.GaugeValue,
		}, {
			desc: prometheus.NewDesc(
				memstatNamespace("alloc_bytes_total"),
				"Total number of bytes allocated, even if freed.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.TotalAlloc) },
			valType: prometheus.CounterValue,
		}, {
			desc: prometheus.NewDesc(
				memstatNamespace("sys_bytes"),
				"Number of bytes obtained from system.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.Sys) },
			valType: prometheus.GaugeValue,
		}, {
			desc: prometheus.NewDesc(
				memstatNamespace("lookups_total"),
				"Total number of pointer lookups.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.Lookups) },
			valType: prometheus.CounterValue,
		}, {
			desc: prometheus.NewDesc(
				memstatNamespace("mallocs_total"),
				"Total number of mallocs.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.Mallocs) },
			valType: prometheus.CounterValue,
		}, {
			desc: prometheus.NewDesc(
				memstatNamespace("frees_total"),
				"Total number of frees.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.Frees) },
			valType: prometheus.CounterValue,
		}, {
			desc: prometheus.NewDesc(
				memstatNamespace("heap_alloc_bytes"),
				"Number of heap bytes allocated and still in use.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.HeapAlloc) },
			valType: prometheus.GaugeValue,
		}, {
			desc: prometheus.NewDesc(
				memstatNamespace("heap_sys_bytes"),
				"Number of heap bytes obtained from system.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.HeapSys) },
			valType: prometheus.GaugeValue,
		}, {
			desc: prometheus.NewDesc(
				memstatNamespace("heap_idle_bytes"),
				"Number of heap bytes waiting to be used.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.HeapIdle) },
			valType: prometheus.GaugeValue,
		}, {
			desc: prometheus.NewDesc(
				memstatNamespace("heap_inuse_bytes"),
				"Number of heap bytes that are in use.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.HeapInuse) },
			valType: prometheus.GaugeValue,
		}, {
			desc: prometheus.NewDesc(
				memstatNamespace("heap_released_bytes"),
				"Number of heap bytes released to OS.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.HeapReleased) },
			valType: prometheus.GaugeValue,
		}, {
			desc: prometheus.NewDesc(
				memstatNamespace("heap_objects"),
				"Number of allocated objects.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.HeapObjects) },
			valType: prometheus.GaugeValue,
		}, {
			desc: prometheus.NewDesc(
				memstatNamespace("stack_inuse_bytes"),
				"Number of bytes in use by the stack allocator.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.StackInuse) },
			valType: prometheus.GaugeValue,
		}, {
			desc: prometheus.NewDesc(
				memstatNamespace("stack_sys_bytes"),
				"Number of bytes obtained from system for stack allocator.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.StackSys) },
			valType: prometheus.GaugeValue,
		}, {
			desc: prometheus.NewDesc(
				memstatNamespace("mspan_inuse_bytes"),
				"Number of bytes in use by mspan structures.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.MSpanInuse) },
			valType: prometheus.GaugeValue,
		}, {
			desc: prometheus.NewDesc(
				memstatNamespace("mspan_sys_bytes"),
				"Number of bytes used for mspan structures obtained from system.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.MSpanSys) },
			valType: prometheus.GaugeValue,
		}, {
			desc: prometheus.NewDesc(
				memstatNamespace("mcache_inuse_bytes"),
				"Number of bytes in use by mcache structures.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.MCacheInuse) },
			valType: prometheus.GaugeValue,
		}, {
			desc: prometheus.NewDesc(
				memstatNamespace("mcache_sys_bytes"),
				"Number of bytes used for mcache structures obtained from system.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.MCacheSys) },
			valType: prometheus.GaugeValue,
		}, {
			desc: prometheus.NewDesc(
				memstatNamespace("buck_hash_sys_bytes"),
				"Number of bytes used by the profiling bucket hash table.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.BuckHashSys) },
			valType: prometheus.GaugeValue,
		}, {
			desc: prometheus.NewDesc(
				memstatNamespace("gc_sys_bytes"),
				"Number of bytes used for garbage collection system metadata.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.GCSys) },
			valType: prometheus.GaugeValue,
		}, {
			desc: prometheus.NewDesc(
				memstatNamespace("other_sys_bytes"),
				"Number of bytes used for other system allocations.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.OtherSys) },
			valType: prometheus.GaugeValue,
		}, {
			desc: prometheus.NewDesc(
				memstatNamespace("next_gc_bytes"),
				"Number of heap bytes when next garbage collection will take place.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return float64(ms.NextGC) },
			valType: prometheus.GaugeValue,
		}, {
			desc: prometheus.NewDesc(
				memstatNamespace("gc_cpu_fraction"),
				"The fraction of this program's available CPU time used by the GC since the program started.",
				nil, nil,
			),
			eval:    func(ms *runtime.MemStats) float64 { return ms.GCCPUFraction },
			valType: prometheus.GaugeValue,
		},
	}
}

type baseGoCollector struct {
	goroutinesDesc *prometheus.Desc
	threadsDesc    *prometheus.Desc
	gcDesc         *prometheus.Desc
	gcLastTimeDesc *prometheus.Desc
	goInfoDesc     *prometheus.Desc
}

func newBaseGoCollector() baseGoCollector {
	return baseGoCollector{
		goroutinesDesc: prometheus.NewDesc(
			"go_goroutines",
			"Number of goroutines that currently exist.",
			nil, nil),
		threadsDesc: prometheus.NewDesc(
			"go_threads",
			"Number of OS threads created.",
			nil, nil),
		gcDesc: prometheus.NewDesc(
			"go_gc_duration_seconds",
			"A summary of the pause duration of garbage collection cycles.",
			nil, nil),
		gcLastTimeDesc: prometheus.NewDesc(
			memstatNamespace("last_gc_time_seconds"),
			"Number of seconds since 1970 of last garbage collection.",
			nil, nil),
		goInfoDesc: prometheus.NewDesc(
			"go_info",
			"Information about the Go environment.",
			nil, prometheus.Labels{"version": runtime.Version()}),
	}
}

// Describe returns all descriptions of the collector.
func (c *baseGoCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.goroutinesDesc
	ch <- c.threadsDesc
	ch <- c.gcDesc
	ch <- c.gcLastTimeDesc
	ch <- c.goInfoDesc
}

// Collect returns the current state of all metrics of the collector.
func (c *baseGoCollector) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(c.goroutinesDesc, prometheus.GaugeValue, float64(runtime.NumGoroutine()))
	n, _ := runtime.ThreadCreateProfile(nil)
	ch <- prometheus.MustNewConstMetric(c.threadsDesc, prometheus.GaugeValue, float64(n))

	var stats debug.GCStats
	stats.PauseQuantiles = make([]time.Duration, 5)
	debug.ReadGCStats(&stats)

	quantiles := make(map[float64]float64)
	for idx, pq := range stats.PauseQuantiles[1:] {
		quantiles[float64(idx+1)/float64(len(stats.PauseQuantiles)-1)] = pq.Seconds()
	}
	quantiles[0.0] = stats.PauseQuantiles[0].Seconds()
	ch <- prometheus.MustNewConstSummary(c.gcDesc, uint64(stats.NumGC), stats.PauseTotal.Seconds(), quantiles)
	ch <- prometheus.MustNewConstMetric(c.gcLastTimeDesc, prometheus.GaugeValue, float64(stats.LastGC.UnixNano())/1e9)

	ch <- prometheus.MustNewConstMetric(c.goInfoDesc, prometheus.GaugeValue, 1)
}

func memstatNamespace(s string) string {
	return "go_memstats_" + s
}

// memStatsMetrics provide description, evaluator, runtime/metrics name, and
// value type for memstat metrics.
type memStatsMetrics []struct {
	desc    *prometheus.Desc
	eval    func(*runtime.MemStats) float64
	valType prometheus.ValueType
}
