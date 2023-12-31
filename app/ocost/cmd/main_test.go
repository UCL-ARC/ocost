package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func sampleResponseJSONString() string {
	return `{"code":200,"status":"success","data":[{"kube-system":{"name":"kube-system","properties":{"cluster":"default-cluster","namespace":"kube-system"},"window":{"start":"2023-08-14T13:20:35Z","end":"2023-09-14T13:20:35Z"},"start":"2023-09-14T08:00:00Z","end":"2023-09-14T13:20:00Z","minutes":320,"cpuCores":0.2,"cpuCoreRequestAverage":0.2,"cpuCoreUsageAverage":0.00467,"cpuCoreHours":1.06667,"cpuCost":0.03372,"cpuCostAdjustment":0,"cpuEfficiency":0.02334,"gpuCount":0,"gpuHours":0,"gpuCost":0,"gpuCostAdjustment":0,"networkTransferBytes":55171601.8491,"networkReceiveBytes":35166640.36687,"networkCost":0,"networkCrossZoneCost":0,"networkCrossRegionCost":0,"networkInternetCost":0,"networkCostAdjustment":0,"loadBalancerCost":0,"loadBalancerCostAdjustment":0,"pvBytes":0,"pvByteHours":0,"pvCost":0,"pvs":null,"pvCostAdjustment":0,"ramBytes":146800640,"ramByteRequestAverage":146800640,"ramByteUsageAverage":73379840,"ramByteHours":782936746.66667,"ramCost":0.00309,"ramCostAdjustment":0,"ramEfficiency":0.49986,"externalCost":0,"sharedCost":0,"totalCost":0.03681,"totalEfficiency":0.06334,"proportionalAssetResourceCosts":{},"lbAllocations":null,"sharedCostBreakdown":{}},"ocost":{"name":"ocost","properties":{"cluster":"default-cluster","namespace":"ocost"},"window":{"start":"2023-08-14T13:20:35Z","end":"2023-09-14T13:20:35Z"},"start":"2023-09-13T13:20:35Z","end":"2023-09-14T13:20:35Z","minutes":1440,"cpuCores":0.00378,"cpuCoreRequestAverage":0.00378,"cpuCoreUsageAverage":0.00081,"cpuCoreHours":0.09083,"cpuCost":0.00287,"cpuCostAdjustment":0,"cpuEfficiency":0.21288,"gpuCount":0,"gpuHours":0,"gpuCost":0,"gpuCostAdjustment":0,"networkTransferBytes":563154180.28639,"networkReceiveBytes":620612437.91113,"networkCost":0,"networkCrossZoneCost":0,"networkCrossRegionCost":0,"networkInternetCost":0,"networkCostAdjustment":0,"loadBalancerCost":0,"loadBalancerCostAdjustment":0,"pvBytes":1908874353.77778,"pvByteHours":45812984490.66666,"pvCost":0,"pvs":{"cluster=default-cluster:name=pvc-4f790e38-cdc7-49bd-8789-cd676df215c2":{"byteHours":45812984490.666664,"cost":0}},"pvCostAdjustment":0,"ramBytes":21827128.88889,"ramByteRequestAverage":21827128.88889,"ramByteUsageAverage":91401978.15898,"ramByteHours":523851093.33333,"ramCost":0.00207,"ramCostAdjustment":0,"ramEfficiency":4.18754,"externalCost":0,"sharedCost":0,"totalCost":0.00494,"totalEfficiency":1.87658,"proportionalAssetResourceCosts":{},"lbAllocations":null,"sharedCostBreakdown":{}}}]}`
}

func TestUnmarshalJson(t *testing.T) {
	s := sampleResponseJSONString()
	data := OpenCostAPIResponse{}

	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		t.Fatal(err)
	}
	if data.Code != 200 {
		t.Fatal("Status code was not as expected")
	}
	if data.Data[0]["kube-system"].Name != "kube-system" {
		t.Fatal("Failed to get the namespace name")
	}
}

func TestRemoveDuplicates(t *testing.T) {
	sampleSlice := []string{"a", "b", "b"}
	expectedSlice := []string{"a", "b"}
	actualSlice := removeDuplicates(sampleSlice)

	for i, expected := range expectedSlice {
		if actualSlice[i] != expected {
			t.Fatalf("Actual: %v != Expected: %v", actualSlice, expectedSlice)
		}
	}
}

type TestGetJSONStruct struct {
	Message string `json:"message"`
}

func TestGetJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message":"hello"}`))
	}))
	defer server.Close()

	data := TestGetJSONStruct{}
	err := getJson(server.URL, &data)
	if err != nil {
		t.Fatalf("Failed to get JSON: %#v", err)
	}
	if data.Message != "hello" {
		t.Fatalf("Value was unset")
	}
}
