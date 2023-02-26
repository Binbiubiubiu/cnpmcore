package entity

import (
	"time"

	"github.com/Binbiubiubiu/cnpmcore/enum"
)

type HookEvent struct {
	ChangeId string         `json:"change_id,omitempty"`
	Event    string         `json:"event,omitempty"`
	FullName string         `json:"full_name,omitempty"`
	Type     string         `json:"type,omitempty"`
	Version  string         `json:"version,omitempty"`
	Change   map[string]any `json:"change,omitempty"`
	Time     *time.Time     `json:"time,omitempty"`
}

func CreateHookEvent(h *HookEvent) *HookEvent {
	h.Type = "package"
	h.Version = "1.0.0"
	t := time.Now()
	h.Time = &t
	return h
}

func CreatePublishEvent(fullname string, changeId string, version string, distTag string) *HookEvent {
	return CreateHookEvent(&HookEvent{
		Event:    enum.HookEventTypePublish,
		ChangeId: changeId,
		FullName: fullname,
		Change: map[string]any{
			"dist-tag": distTag,
			"version":  version,
		},
	})
}

func CreateUnpublishEvent(fullname string, changeId string, version string, distTag string) *HookEvent {
	return CreateHookEvent(&HookEvent{
		Event:    enum.HookEventTypeUnpublish,
		ChangeId: changeId,
		FullName: fullname,
		Change: map[string]any{
			"dist-tag": distTag,
			"version":  version,
		},
	})
}

func CreateOwnerEvent(fullname string, changeId string, maintainer string) *HookEvent {
	return CreateHookEvent(&HookEvent{
		Event:    enum.HookEventTypeUnpublish,
		ChangeId: changeId,
		FullName: fullname,
		Change: map[string]any{
			"maintainer": maintainer,
		},
	})
}

func CreateOwnerRmEvent(fullname string, changeId string, maintainer string) *HookEvent {
	return CreateHookEvent(&HookEvent{
		Event:    enum.HookEventTypeOwnerRm,
		ChangeId: changeId,
		FullName: fullname,
		Change: map[string]any{
			"maintainer": maintainer,
		},
	})
}

func CreateDistTagEvent(fullname string, changeId string, distTag string) *HookEvent {
	return CreateHookEvent(&HookEvent{
		Event:    enum.HookEventTypeDistTag,
		ChangeId: changeId,
		FullName: fullname,
		Change: map[string]any{
			"dist-tag": distTag,
		},
	})
}

func CreateDistTagRmEvent(fullname string, changeId string, distTag string) *HookEvent {
	return CreateHookEvent(&HookEvent{
		Event:    enum.HookEventTypeDistTag,
		ChangeId: changeId,
		FullName: fullname,
		Change: map[string]any{
			"dist-tag": distTag,
		},
	})
}

func CreateDeprecatedEvent(fullname string, changeId string, deprecated string) *HookEvent {
	return CreateHookEvent(&HookEvent{
		Event:    enum.HookEventTypeDistTag,
		ChangeId: changeId,
		FullName: fullname,
		Change: map[string]any{
			"deprecated": deprecated,
		},
	})
}

func CreateUndeprecatedEvent(fullname string, changeId string, deprecated string) *HookEvent {
	return CreateHookEvent(&HookEvent{
		Event:    enum.HookEventTypeDistTag,
		ChangeId: changeId,
		FullName: fullname,
		Change: map[string]any{
			"deprecated": deprecated,
		},
	})
}
