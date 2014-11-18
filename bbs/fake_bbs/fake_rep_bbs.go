// This file was generated by counterfeiter
package fake_bbs

import (
	"sync"
	"time"

	"github.com/cloudfoundry-incubator/runtime-schema/bbs"
	"github.com/cloudfoundry-incubator/runtime-schema/models"
	"github.com/tedsuo/ifrit"
)

type FakeRepBBS struct {
	NewCellHeartbeatStub        func(cellPresence models.CellPresence, interval time.Duration) ifrit.Runner
	newCellHeartbeatMutex       sync.RWMutex
	newCellHeartbeatArgsForCall []struct {
		cellPresence models.CellPresence
		interval     time.Duration
	}
	newCellHeartbeatReturns struct {
		result1 ifrit.Runner
	}
	WatchForDesiredTaskStub        func() (<-chan models.Task, chan<- bool, <-chan error)
	watchForDesiredTaskMutex       sync.RWMutex
	watchForDesiredTaskArgsForCall []struct{}
	watchForDesiredTaskReturns     struct {
		result1 <-chan models.Task
		result2 chan<- bool
		result3 <-chan error
	}
	ClaimTaskStub        func(taskGuid string, cellID string) error
	claimTaskMutex       sync.RWMutex
	claimTaskArgsForCall []struct {
		taskGuid string
		cellID   string
	}
	claimTaskReturns struct {
		result1 error
	}
	StartTaskStub        func(taskGuid string, cellID string) error
	startTaskMutex       sync.RWMutex
	startTaskArgsForCall []struct {
		taskGuid string
		cellID   string
	}
	startTaskReturns struct {
		result1 error
	}
	TaskByGuidStub           func(taskGuid string) (models.Task, error)
	getTaskByGuidMutex       sync.RWMutex
	getTaskByGuidArgsForCall []struct {
		taskGuid string
	}
	getTaskByGuidReturns struct {
		result1 models.Task
		result2 error
	}
	TasksByCellIDStub              func(cellID string) ([]models.Task, error)
	getAllTasksByCellIDMutex       sync.RWMutex
	getAllTasksByCellIDArgsForCall []struct {
		cellID string
	}
	getAllTasksByCellIDReturns struct {
		result1 []models.Task
		result2 error
	}
	CompleteTaskStub        func(taskGuid string, failed bool, failureReason string, result string) error
	completeTaskMutex       sync.RWMutex
	completeTaskArgsForCall []struct {
		taskGuid      string
		failed        bool
		failureReason string
		result        string
	}
	completeTaskReturns struct {
		result1 error
	}
	ActualLRPsByProcessGuidStub           func(string) ([]models.ActualLRP, error)
	getActualLRPsByProcessGuidMutex       sync.RWMutex
	getActualLRPsByProcessGuidArgsForCall []struct {
		arg1 string
	}
	getActualLRPsByProcessGuidReturns struct {
		result1 []models.ActualLRP
		result2 error
	}
	ActualLRPsByCellIDStub              func(cellID string) ([]models.ActualLRP, error)
	getAllActualLRPsByCellIDMutex       sync.RWMutex
	getAllActualLRPsByCellIDArgsForCall []struct {
		cellID string
	}
	getAllActualLRPsByCellIDReturns struct {
		result1 []models.ActualLRP
		result2 error
	}
	ReportActualLRPAsStartingStub        func(processGuid, instanceGuid, cellID, domain string, index int) (models.ActualLRP, error)
	reportActualLRPAsStartingMutex       sync.RWMutex
	reportActualLRPAsStartingArgsForCall []struct {
		processGuid  string
		instanceGuid string
		cellID       string
		domain       string
		index        int
	}
	reportActualLRPAsStartingReturns struct {
		result1 models.ActualLRP
		result2 error
	}
	ReportActualLRPAsRunningStub        func(lrp models.ActualLRP, cellId string) error
	reportActualLRPAsRunningMutex       sync.RWMutex
	reportActualLRPAsRunningArgsForCall []struct {
		lrp    models.ActualLRP
		cellId string
	}
	reportActualLRPAsRunningReturns struct {
		result1 error
	}
	RemoveActualLRPStub        func(lrp models.ActualLRP) error
	removeActualLRPMutex       sync.RWMutex
	removeActualLRPArgsForCall []struct {
		lrp models.ActualLRP
	}
	removeActualLRPReturns struct {
		result1 error
	}
	RemoveActualLRPForIndexStub        func(processGuid string, index int, instanceGuid string) error
	removeActualLRPForIndexMutex       sync.RWMutex
	removeActualLRPForIndexArgsForCall []struct {
		processGuid  string
		index        int
		instanceGuid string
	}
	removeActualLRPForIndexReturns struct {
		result1 error
	}
	WatchForStopLRPInstanceStub        func() (<-chan models.StopLRPInstance, chan<- bool, <-chan error)
	watchForStopLRPInstanceMutex       sync.RWMutex
	watchForStopLRPInstanceArgsForCall []struct{}
	watchForStopLRPInstanceReturns     struct {
		result1 <-chan models.StopLRPInstance
		result2 chan<- bool
		result3 <-chan error
	}
	ResolveStopLRPInstanceStub        func(stopInstance models.StopLRPInstance) error
	resolveStopLRPInstanceMutex       sync.RWMutex
	resolveStopLRPInstanceArgsForCall []struct {
		stopInstance models.StopLRPInstance
	}
	resolveStopLRPInstanceReturns struct {
		result1 error
	}
}

func (fake *FakeRepBBS) NewCellHeartbeat(cellPresence models.CellPresence, interval time.Duration) ifrit.Runner {
	fake.newCellHeartbeatMutex.Lock()
	fake.newCellHeartbeatArgsForCall = append(fake.newCellHeartbeatArgsForCall, struct {
		cellPresence models.CellPresence
		interval     time.Duration
	}{cellPresence, interval})
	fake.newCellHeartbeatMutex.Unlock()
	if fake.NewCellHeartbeatStub != nil {
		return fake.NewCellHeartbeatStub(cellPresence, interval)
	} else {
		return fake.newCellHeartbeatReturns.result1
	}
}

func (fake *FakeRepBBS) NewCellHeartbeatCallCount() int {
	fake.newCellHeartbeatMutex.RLock()
	defer fake.newCellHeartbeatMutex.RUnlock()
	return len(fake.newCellHeartbeatArgsForCall)
}

func (fake *FakeRepBBS) NewCellHeartbeatArgsForCall(i int) (models.CellPresence, time.Duration) {
	fake.newCellHeartbeatMutex.RLock()
	defer fake.newCellHeartbeatMutex.RUnlock()
	return fake.newCellHeartbeatArgsForCall[i].cellPresence, fake.newCellHeartbeatArgsForCall[i].interval
}

func (fake *FakeRepBBS) NewCellHeartbeatReturns(result1 ifrit.Runner) {
	fake.NewCellHeartbeatStub = nil
	fake.newCellHeartbeatReturns = struct {
		result1 ifrit.Runner
	}{result1}
}

func (fake *FakeRepBBS) WatchForDesiredTask() (<-chan models.Task, chan<- bool, <-chan error) {
	fake.watchForDesiredTaskMutex.Lock()
	fake.watchForDesiredTaskArgsForCall = append(fake.watchForDesiredTaskArgsForCall, struct{}{})
	fake.watchForDesiredTaskMutex.Unlock()
	if fake.WatchForDesiredTaskStub != nil {
		return fake.WatchForDesiredTaskStub()
	} else {
		return fake.watchForDesiredTaskReturns.result1, fake.watchForDesiredTaskReturns.result2, fake.watchForDesiredTaskReturns.result3
	}
}

func (fake *FakeRepBBS) WatchForDesiredTaskCallCount() int {
	fake.watchForDesiredTaskMutex.RLock()
	defer fake.watchForDesiredTaskMutex.RUnlock()
	return len(fake.watchForDesiredTaskArgsForCall)
}

func (fake *FakeRepBBS) WatchForDesiredTaskReturns(result1 <-chan models.Task, result2 chan<- bool, result3 <-chan error) {
	fake.WatchForDesiredTaskStub = nil
	fake.watchForDesiredTaskReturns = struct {
		result1 <-chan models.Task
		result2 chan<- bool
		result3 <-chan error
	}{result1, result2, result3}
}

func (fake *FakeRepBBS) ClaimTask(taskGuid string, cellID string) error {
	fake.claimTaskMutex.Lock()
	fake.claimTaskArgsForCall = append(fake.claimTaskArgsForCall, struct {
		taskGuid string
		cellID   string
	}{taskGuid, cellID})
	fake.claimTaskMutex.Unlock()
	if fake.ClaimTaskStub != nil {
		return fake.ClaimTaskStub(taskGuid, cellID)
	} else {
		return fake.claimTaskReturns.result1
	}
}

func (fake *FakeRepBBS) ClaimTaskCallCount() int {
	fake.claimTaskMutex.RLock()
	defer fake.claimTaskMutex.RUnlock()
	return len(fake.claimTaskArgsForCall)
}

func (fake *FakeRepBBS) ClaimTaskArgsForCall(i int) (string, string) {
	fake.claimTaskMutex.RLock()
	defer fake.claimTaskMutex.RUnlock()
	return fake.claimTaskArgsForCall[i].taskGuid, fake.claimTaskArgsForCall[i].cellID
}

func (fake *FakeRepBBS) ClaimTaskReturns(result1 error) {
	fake.ClaimTaskStub = nil
	fake.claimTaskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepBBS) StartTask(taskGuid string, cellID string) error {
	fake.startTaskMutex.Lock()
	fake.startTaskArgsForCall = append(fake.startTaskArgsForCall, struct {
		taskGuid string
		cellID   string
	}{taskGuid, cellID})
	fake.startTaskMutex.Unlock()
	if fake.StartTaskStub != nil {
		return fake.StartTaskStub(taskGuid, cellID)
	} else {
		return fake.startTaskReturns.result1
	}
}

func (fake *FakeRepBBS) StartTaskCallCount() int {
	fake.startTaskMutex.RLock()
	defer fake.startTaskMutex.RUnlock()
	return len(fake.startTaskArgsForCall)
}

func (fake *FakeRepBBS) StartTaskArgsForCall(i int) (string, string) {
	fake.startTaskMutex.RLock()
	defer fake.startTaskMutex.RUnlock()
	return fake.startTaskArgsForCall[i].taskGuid, fake.startTaskArgsForCall[i].cellID
}

func (fake *FakeRepBBS) StartTaskReturns(result1 error) {
	fake.StartTaskStub = nil
	fake.startTaskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepBBS) TaskByGuid(taskGuid string) (models.Task, error) {
	fake.getTaskByGuidMutex.Lock()
	fake.getTaskByGuidArgsForCall = append(fake.getTaskByGuidArgsForCall, struct {
		taskGuid string
	}{taskGuid})
	fake.getTaskByGuidMutex.Unlock()
	if fake.TaskByGuidStub != nil {
		return fake.TaskByGuidStub(taskGuid)
	} else {
		return fake.getTaskByGuidReturns.result1, fake.getTaskByGuidReturns.result2
	}
}

func (fake *FakeRepBBS) TaskByGuidCallCount() int {
	fake.getTaskByGuidMutex.RLock()
	defer fake.getTaskByGuidMutex.RUnlock()
	return len(fake.getTaskByGuidArgsForCall)
}

func (fake *FakeRepBBS) TaskByGuidArgsForCall(i int) string {
	fake.getTaskByGuidMutex.RLock()
	defer fake.getTaskByGuidMutex.RUnlock()
	return fake.getTaskByGuidArgsForCall[i].taskGuid
}

func (fake *FakeRepBBS) TaskByGuidReturns(result1 models.Task, result2 error) {
	fake.TaskByGuidStub = nil
	fake.getTaskByGuidReturns = struct {
		result1 models.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeRepBBS) TasksByCellID(cellID string) ([]models.Task, error) {
	fake.getAllTasksByCellIDMutex.Lock()
	fake.getAllTasksByCellIDArgsForCall = append(fake.getAllTasksByCellIDArgsForCall, struct {
		cellID string
	}{cellID})
	fake.getAllTasksByCellIDMutex.Unlock()
	if fake.TasksByCellIDStub != nil {
		return fake.TasksByCellIDStub(cellID)
	} else {
		return fake.getAllTasksByCellIDReturns.result1, fake.getAllTasksByCellIDReturns.result2
	}
}

func (fake *FakeRepBBS) TasksByCellIDCallCount() int {
	fake.getAllTasksByCellIDMutex.RLock()
	defer fake.getAllTasksByCellIDMutex.RUnlock()
	return len(fake.getAllTasksByCellIDArgsForCall)
}

func (fake *FakeRepBBS) TasksByCellIDArgsForCall(i int) string {
	fake.getAllTasksByCellIDMutex.RLock()
	defer fake.getAllTasksByCellIDMutex.RUnlock()
	return fake.getAllTasksByCellIDArgsForCall[i].cellID
}

func (fake *FakeRepBBS) TasksByCellIDReturns(result1 []models.Task, result2 error) {
	fake.TasksByCellIDStub = nil
	fake.getAllTasksByCellIDReturns = struct {
		result1 []models.Task
		result2 error
	}{result1, result2}
}

func (fake *FakeRepBBS) CompleteTask(taskGuid string, failed bool, failureReason string, result string) error {
	fake.completeTaskMutex.Lock()
	fake.completeTaskArgsForCall = append(fake.completeTaskArgsForCall, struct {
		taskGuid      string
		failed        bool
		failureReason string
		result        string
	}{taskGuid, failed, failureReason, result})
	fake.completeTaskMutex.Unlock()
	if fake.CompleteTaskStub != nil {
		return fake.CompleteTaskStub(taskGuid, failed, failureReason, result)
	} else {
		return fake.completeTaskReturns.result1
	}
}

func (fake *FakeRepBBS) CompleteTaskCallCount() int {
	fake.completeTaskMutex.RLock()
	defer fake.completeTaskMutex.RUnlock()
	return len(fake.completeTaskArgsForCall)
}

func (fake *FakeRepBBS) CompleteTaskArgsForCall(i int) (string, bool, string, string) {
	fake.completeTaskMutex.RLock()
	defer fake.completeTaskMutex.RUnlock()
	return fake.completeTaskArgsForCall[i].taskGuid, fake.completeTaskArgsForCall[i].failed, fake.completeTaskArgsForCall[i].failureReason, fake.completeTaskArgsForCall[i].result
}

func (fake *FakeRepBBS) CompleteTaskReturns(result1 error) {
	fake.CompleteTaskStub = nil
	fake.completeTaskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepBBS) ActualLRPsByProcessGuid(arg1 string) ([]models.ActualLRP, error) {
	fake.getActualLRPsByProcessGuidMutex.Lock()
	fake.getActualLRPsByProcessGuidArgsForCall = append(fake.getActualLRPsByProcessGuidArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.getActualLRPsByProcessGuidMutex.Unlock()
	if fake.ActualLRPsByProcessGuidStub != nil {
		return fake.ActualLRPsByProcessGuidStub(arg1)
	} else {
		return fake.getActualLRPsByProcessGuidReturns.result1, fake.getActualLRPsByProcessGuidReturns.result2
	}
}

func (fake *FakeRepBBS) ActualLRPsByProcessGuidCallCount() int {
	fake.getActualLRPsByProcessGuidMutex.RLock()
	defer fake.getActualLRPsByProcessGuidMutex.RUnlock()
	return len(fake.getActualLRPsByProcessGuidArgsForCall)
}

func (fake *FakeRepBBS) ActualLRPsByProcessGuidArgsForCall(i int) string {
	fake.getActualLRPsByProcessGuidMutex.RLock()
	defer fake.getActualLRPsByProcessGuidMutex.RUnlock()
	return fake.getActualLRPsByProcessGuidArgsForCall[i].arg1
}

func (fake *FakeRepBBS) ActualLRPsByProcessGuidReturns(result1 []models.ActualLRP, result2 error) {
	fake.ActualLRPsByProcessGuidStub = nil
	fake.getActualLRPsByProcessGuidReturns = struct {
		result1 []models.ActualLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeRepBBS) ActualLRPsByCellID(cellID string) ([]models.ActualLRP, error) {
	fake.getAllActualLRPsByCellIDMutex.Lock()
	fake.getAllActualLRPsByCellIDArgsForCall = append(fake.getAllActualLRPsByCellIDArgsForCall, struct {
		cellID string
	}{cellID})
	fake.getAllActualLRPsByCellIDMutex.Unlock()
	if fake.ActualLRPsByCellIDStub != nil {
		return fake.ActualLRPsByCellIDStub(cellID)
	} else {
		return fake.getAllActualLRPsByCellIDReturns.result1, fake.getAllActualLRPsByCellIDReturns.result2
	}
}

func (fake *FakeRepBBS) ActualLRPsByCellIDCallCount() int {
	fake.getAllActualLRPsByCellIDMutex.RLock()
	defer fake.getAllActualLRPsByCellIDMutex.RUnlock()
	return len(fake.getAllActualLRPsByCellIDArgsForCall)
}

func (fake *FakeRepBBS) ActualLRPsByCellIDArgsForCall(i int) string {
	fake.getAllActualLRPsByCellIDMutex.RLock()
	defer fake.getAllActualLRPsByCellIDMutex.RUnlock()
	return fake.getAllActualLRPsByCellIDArgsForCall[i].cellID
}

func (fake *FakeRepBBS) ActualLRPsByCellIDReturns(result1 []models.ActualLRP, result2 error) {
	fake.ActualLRPsByCellIDStub = nil
	fake.getAllActualLRPsByCellIDReturns = struct {
		result1 []models.ActualLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeRepBBS) ReportActualLRPAsStarting(processGuid string, instanceGuid string, cellID string, domain string, index int) (models.ActualLRP, error) {
	fake.reportActualLRPAsStartingMutex.Lock()
	fake.reportActualLRPAsStartingArgsForCall = append(fake.reportActualLRPAsStartingArgsForCall, struct {
		processGuid  string
		instanceGuid string
		cellID       string
		domain       string
		index        int
	}{processGuid, instanceGuid, cellID, domain, index})
	fake.reportActualLRPAsStartingMutex.Unlock()
	if fake.ReportActualLRPAsStartingStub != nil {
		return fake.ReportActualLRPAsStartingStub(processGuid, instanceGuid, cellID, domain, index)
	} else {
		return fake.reportActualLRPAsStartingReturns.result1, fake.reportActualLRPAsStartingReturns.result2
	}
}

func (fake *FakeRepBBS) ReportActualLRPAsStartingCallCount() int {
	fake.reportActualLRPAsStartingMutex.RLock()
	defer fake.reportActualLRPAsStartingMutex.RUnlock()
	return len(fake.reportActualLRPAsStartingArgsForCall)
}

func (fake *FakeRepBBS) ReportActualLRPAsStartingArgsForCall(i int) (string, string, string, string, int) {
	fake.reportActualLRPAsStartingMutex.RLock()
	defer fake.reportActualLRPAsStartingMutex.RUnlock()
	return fake.reportActualLRPAsStartingArgsForCall[i].processGuid, fake.reportActualLRPAsStartingArgsForCall[i].instanceGuid, fake.reportActualLRPAsStartingArgsForCall[i].cellID, fake.reportActualLRPAsStartingArgsForCall[i].domain, fake.reportActualLRPAsStartingArgsForCall[i].index
}

func (fake *FakeRepBBS) ReportActualLRPAsStartingReturns(result1 models.ActualLRP, result2 error) {
	fake.ReportActualLRPAsStartingStub = nil
	fake.reportActualLRPAsStartingReturns = struct {
		result1 models.ActualLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeRepBBS) ReportActualLRPAsRunning(lrp models.ActualLRP, cellId string) error {
	fake.reportActualLRPAsRunningMutex.Lock()
	fake.reportActualLRPAsRunningArgsForCall = append(fake.reportActualLRPAsRunningArgsForCall, struct {
		lrp    models.ActualLRP
		cellId string
	}{lrp, cellId})
	fake.reportActualLRPAsRunningMutex.Unlock()
	if fake.ReportActualLRPAsRunningStub != nil {
		return fake.ReportActualLRPAsRunningStub(lrp, cellId)
	} else {
		return fake.reportActualLRPAsRunningReturns.result1
	}
}

func (fake *FakeRepBBS) ReportActualLRPAsRunningCallCount() int {
	fake.reportActualLRPAsRunningMutex.RLock()
	defer fake.reportActualLRPAsRunningMutex.RUnlock()
	return len(fake.reportActualLRPAsRunningArgsForCall)
}

func (fake *FakeRepBBS) ReportActualLRPAsRunningArgsForCall(i int) (models.ActualLRP, string) {
	fake.reportActualLRPAsRunningMutex.RLock()
	defer fake.reportActualLRPAsRunningMutex.RUnlock()
	return fake.reportActualLRPAsRunningArgsForCall[i].lrp, fake.reportActualLRPAsRunningArgsForCall[i].cellId
}

func (fake *FakeRepBBS) ReportActualLRPAsRunningReturns(result1 error) {
	fake.ReportActualLRPAsRunningStub = nil
	fake.reportActualLRPAsRunningReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepBBS) RemoveActualLRP(lrp models.ActualLRP) error {
	fake.removeActualLRPMutex.Lock()
	fake.removeActualLRPArgsForCall = append(fake.removeActualLRPArgsForCall, struct {
		lrp models.ActualLRP
	}{lrp})
	fake.removeActualLRPMutex.Unlock()
	if fake.RemoveActualLRPStub != nil {
		return fake.RemoveActualLRPStub(lrp)
	} else {
		return fake.removeActualLRPReturns.result1
	}
}

func (fake *FakeRepBBS) RemoveActualLRPCallCount() int {
	fake.removeActualLRPMutex.RLock()
	defer fake.removeActualLRPMutex.RUnlock()
	return len(fake.removeActualLRPArgsForCall)
}

func (fake *FakeRepBBS) RemoveActualLRPArgsForCall(i int) models.ActualLRP {
	fake.removeActualLRPMutex.RLock()
	defer fake.removeActualLRPMutex.RUnlock()
	return fake.removeActualLRPArgsForCall[i].lrp
}

func (fake *FakeRepBBS) RemoveActualLRPReturns(result1 error) {
	fake.RemoveActualLRPStub = nil
	fake.removeActualLRPReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepBBS) RemoveActualLRPForIndex(processGuid string, index int, instanceGuid string) error {
	fake.removeActualLRPForIndexMutex.Lock()
	fake.removeActualLRPForIndexArgsForCall = append(fake.removeActualLRPForIndexArgsForCall, struct {
		processGuid  string
		index        int
		instanceGuid string
	}{processGuid, index, instanceGuid})
	fake.removeActualLRPForIndexMutex.Unlock()
	if fake.RemoveActualLRPForIndexStub != nil {
		return fake.RemoveActualLRPForIndexStub(processGuid, index, instanceGuid)
	} else {
		return fake.removeActualLRPForIndexReturns.result1
	}
}

func (fake *FakeRepBBS) RemoveActualLRPForIndexCallCount() int {
	fake.removeActualLRPForIndexMutex.RLock()
	defer fake.removeActualLRPForIndexMutex.RUnlock()
	return len(fake.removeActualLRPForIndexArgsForCall)
}

func (fake *FakeRepBBS) RemoveActualLRPForIndexArgsForCall(i int) (string, int, string) {
	fake.removeActualLRPForIndexMutex.RLock()
	defer fake.removeActualLRPForIndexMutex.RUnlock()
	return fake.removeActualLRPForIndexArgsForCall[i].processGuid, fake.removeActualLRPForIndexArgsForCall[i].index, fake.removeActualLRPForIndexArgsForCall[i].instanceGuid
}

func (fake *FakeRepBBS) RemoveActualLRPForIndexReturns(result1 error) {
	fake.RemoveActualLRPForIndexStub = nil
	fake.removeActualLRPForIndexReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepBBS) WatchForStopLRPInstance() (<-chan models.StopLRPInstance, chan<- bool, <-chan error) {
	fake.watchForStopLRPInstanceMutex.Lock()
	fake.watchForStopLRPInstanceArgsForCall = append(fake.watchForStopLRPInstanceArgsForCall, struct{}{})
	fake.watchForStopLRPInstanceMutex.Unlock()
	if fake.WatchForStopLRPInstanceStub != nil {
		return fake.WatchForStopLRPInstanceStub()
	} else {
		return fake.watchForStopLRPInstanceReturns.result1, fake.watchForStopLRPInstanceReturns.result2, fake.watchForStopLRPInstanceReturns.result3
	}
}

func (fake *FakeRepBBS) WatchForStopLRPInstanceCallCount() int {
	fake.watchForStopLRPInstanceMutex.RLock()
	defer fake.watchForStopLRPInstanceMutex.RUnlock()
	return len(fake.watchForStopLRPInstanceArgsForCall)
}

func (fake *FakeRepBBS) WatchForStopLRPInstanceReturns(result1 <-chan models.StopLRPInstance, result2 chan<- bool, result3 <-chan error) {
	fake.WatchForStopLRPInstanceStub = nil
	fake.watchForStopLRPInstanceReturns = struct {
		result1 <-chan models.StopLRPInstance
		result2 chan<- bool
		result3 <-chan error
	}{result1, result2, result3}
}

func (fake *FakeRepBBS) ResolveStopLRPInstance(stopInstance models.StopLRPInstance) error {
	fake.resolveStopLRPInstanceMutex.Lock()
	fake.resolveStopLRPInstanceArgsForCall = append(fake.resolveStopLRPInstanceArgsForCall, struct {
		stopInstance models.StopLRPInstance
	}{stopInstance})
	fake.resolveStopLRPInstanceMutex.Unlock()
	if fake.ResolveStopLRPInstanceStub != nil {
		return fake.ResolveStopLRPInstanceStub(stopInstance)
	} else {
		return fake.resolveStopLRPInstanceReturns.result1
	}
}

func (fake *FakeRepBBS) ResolveStopLRPInstanceCallCount() int {
	fake.resolveStopLRPInstanceMutex.RLock()
	defer fake.resolveStopLRPInstanceMutex.RUnlock()
	return len(fake.resolveStopLRPInstanceArgsForCall)
}

func (fake *FakeRepBBS) ResolveStopLRPInstanceArgsForCall(i int) models.StopLRPInstance {
	fake.resolveStopLRPInstanceMutex.RLock()
	defer fake.resolveStopLRPInstanceMutex.RUnlock()
	return fake.resolveStopLRPInstanceArgsForCall[i].stopInstance
}

func (fake *FakeRepBBS) ResolveStopLRPInstanceReturns(result1 error) {
	fake.ResolveStopLRPInstanceStub = nil
	fake.resolveStopLRPInstanceReturns = struct {
		result1 error
	}{result1}
}

var _ bbs.RepBBS = new(FakeRepBBS)
