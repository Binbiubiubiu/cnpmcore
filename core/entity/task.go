package entity

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/Binbiubiubiu/cnpmcore/common/enum"
	"github.com/Binbiubiubiu/cnpmcore/repository/model"
	"github.com/Binbiubiubiu/cnpmcore/util"
	"github.com/samber/lo"
)

var hostName, _ = os.Hostname()
var pID = os.Getpid()

type TaskUpdateCondition struct {
	TaskId   string
	Attempts uint32
}

type SyncPackageTaskOptions struct {
	AuthorId         string
	AuthorIp         string
	Tips             string
	SkipDependencies bool
	SyncDownloadData bool
	ForceSyncHistory bool
	RegistryId       string
}

type Task struct {
	Entity
	TaskId           string         `json:"task_id,omitempty"`
	Type             string         `json:"type,omitempty"`
	State            string         `json:"state,omitempty"`
	TargetName       string         `json:"target_name,omitempty"`
	AuthorId         string         `json:"author_id,omitempty"`
	AuthorIp         string         `json:"author_ip,omitempty"`
	Data             map[string]any `json:"data,omitempty"`
	LogPath          string         `json:"log_path,omitempty"`
	LogStorePosition string         `json:"log_store_position,omitempty"`
	Attempts         uint32         `json:"attempts,omitempty"`
	Error            string         `json:"error,omitempty"`
	BizId            string         `json:"biz_id,omitempty"`
}

func (t *Task) ResetLogPath() {
	t.LogPath = fmt.Sprintf("%s/%s-%s-%v.log", path.Dir(t.LogPath), time.Now().Format("021514"), t.TaskId, t.Attempts)
	t.LogStorePosition = ""
}

func (t *Task) SetExecuteWorker() {
	t.Data["taskBaseData"] = fmt.Sprintf("%s:%v", hostName, pID)
}

func (t *Task) UpdateSyncData(info SyncInfo) {
	t.Data["since"] = info.LastSince
	tc, ok := t.Data["task_count"].(uint)
	if ok {
		t.Data["task_count"] = tc + info.TaskCount
	}

	if info.TaskCount > 0 {
		t.Data["last_package"] = info.LastPackage
		t.Data["last_package_created"] = time.Now()
	}
}

func (t *Task) Start() TaskUpdateCondition {
	cond := TaskUpdateCondition{
		TaskId:   t.TaskId,
		Attempts: t.Attempts,
	}
	t.SetExecuteWorker()
	t.State = enum.TaskStateProcessing
	t.Attempts = 1
	return cond
}

func (e *Task) Into() *model.Task {
	var m model.Task
	m.ID = e.ID
	m.UpdateAt = e.UpdateAt
	m.CreateAt = e.CreateAt
	m.TaskId = e.TaskId
	m.Type = e.Type
	m.State = e.State
	m.TargetName = e.TargetName
	m.AuthorId = e.AuthorId
	m.AuthorIp = e.AuthorIp
	m.Data = e.Data
	m.LogPath = e.LogPath
	m.LogStorePosition = e.LogStorePosition
	m.Attempts = e.Attempts
	m.Error = e.Error
	m.BizId = e.BizId
	return &m
}

func (e *Task) From(m *model.Task) {
	e.ID = m.ID
	e.UpdateAt = m.UpdateAt
	e.CreateAt = m.CreateAt
	e.TaskId = m.TaskId
	e.Type = m.Type
	e.State = m.State
	e.TargetName = m.TargetName
	e.AuthorId = m.AuthorId
	e.AuthorIp = m.AuthorIp
	e.Data = m.Data
	e.LogPath = m.LogPath
	e.LogStorePosition = m.LogStorePosition
	e.Attempts = m.Attempts
	e.Error = m.Error
	e.BizId = m.BizId
}

func (e *Task) FromHistory(m *model.HistoryTask) {
	e.ID = m.ID
	e.UpdateAt = m.UpdateAt
	e.CreateAt = m.CreateAt
	e.TaskId = m.TaskId
	e.Type = m.Type
	e.State = m.State
	e.TargetName = m.TargetName
	e.AuthorId = m.AuthorId
	e.AuthorIp = m.AuthorIp
	e.Data = m.Data
	e.LogPath = m.LogPath
	e.LogStorePosition = m.LogStorePosition
	e.Attempts = m.Attempts
	e.Error = m.Error
}

func CreateTask(e *Task) *Task {
	e.TaskId = util.CreateObjectId()
	return e
}

func CreateTaskFromModel(m *model.Task) *Task {
	var e Task
	e.From(m)
	return &e
}

func CreateTaskFromHistoryModel(m *model.HistoryTask) *Task {
	var e Task
	e.FromHistory(m)
	return &e
}

func CreateSyncPackage(fullname string, options *SyncPackageTaskOptions) *Task {
	if options == nil {
		options = &SyncPackageTaskOptions{}
	}
	t := CreateTask(&Task{
		Type:     enum.TaskTypeSyncPackage,
		State:    enum.TaskStateWaiting,
		AuthorId: options.AuthorId,
		AuthorIp: options.AuthorIp,
		Data: map[string]any{
			"taskWorker":       "",
			"tips":             options.Tips,
			"registryId":       options.RegistryId,
			"skipDependencies": options.SkipDependencies,
			"syncDownloadData": options.SyncDownloadData,
			"forceSyncHistory": options.ForceSyncHistory,
		},
	})
	t.LogPath = fmt.Sprintf("/packages/%s/syncs/%s-%s.log", fullname, time.Now().Format("2006/01/021504"), t.TaskId)
	return t
}

func CreateChangesStream(targetName string, args ...string) *Task {
	var registryId, since string
	if len(args) > 0 {
		if lo.IsNotEmpty(args[0]) {
			registryId = args[0]
		}
		if lo.IsNotEmpty(args[1]) {
			since = args[1]
		}
	}
	t := CreateTask(&Task{
		Type:       enum.TaskTypeChangesStream,
		State:      enum.TaskStateWaiting,
		TargetName: targetName,
		AuthorId:   fmt.Sprintf("pid_%v", pID),
		AuthorIp:   hostName,
		Data: map[string]any{
			"taskWorker": "",
			"registryId": registryId,
			"since":      since,
		},
	})
	return t
}

func CreateCreateHookTask(hookEvent HookEvent) *Task {
	t := CreateTask(&Task{
		Type:       enum.TaskTypeCreateHook,
		State:      enum.TaskStateWaiting,
		TargetName: hookEvent.FullName,
		AuthorId:   fmt.Sprintf("pid_%v", pID),
		AuthorIp:   hostName,
		BizId:      fmt.Sprintf("CreateHook:%s", hookEvent.ChangeId),
		Data: map[string]any{
			"taskWorker": "",
			"hookEvent":  hookEvent,
		},
	})
	t.LogPath = fmt.Sprintf("/packages/%s/syncs/%s-%s.log", hookEvent.FullName, time.Now().Format("2006/01/021504"), t.TaskId)
	return t
}

func createTriggerHookTask(hookEvent HookEvent, hookId string) *Task {
	t := CreateTask(&Task{
		Type:       enum.TaskTypeTriggerHook,
		State:      enum.TaskStateWaiting,
		TargetName: hookEvent.FullName,
		AuthorId:   fmt.Sprintf("pid_%v", pID),
		AuthorIp:   hostName,
		BizId:      fmt.Sprintf("TriggerHook:%s:%s", hookEvent.ChangeId, hookId),
		Data: map[string]any{
			"taskWorker": "",
			"hookEvent":  hookEvent,
			"hookId":     hookId,
		},
	})
	t.LogPath = fmt.Sprintf("/packages/%s/syncs/%s-%s.log", hookEvent.FullName, time.Now().Format("2006/01/021504"), t.TaskId)
	return t
}

func CreateSyncBinary(targetName string, lastData map[string]any) *Task {
	lastData["taskWorker"] = ""
	t := CreateTask(&Task{
		Type:       enum.TaskTypeSyncBinary,
		State:      enum.TaskStateWaiting,
		TargetName: targetName,
		AuthorId:   fmt.Sprintf("pid_%v", pID),
		AuthorIp:   hostName,
		BizId:      fmt.Sprintf("SyncBinary:%s", targetName),
		Data: lo.Assign(
			map[string]any{"taskWorker": ""},
			lastData,
		),
	})
	t.LogPath = fmt.Sprintf("/packages/%s/syncs/%s-%s.log", targetName, time.Now().Format("2006/01/021504"), t.TaskId)
	return t
}

type SyncInfo struct {
	LastSince   string
	TaskCount   uint
	LastPackage string
}
