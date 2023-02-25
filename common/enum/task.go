package enum

const (
	TaskTypeSyncPackage   = "sync_package"
	TaskTypeChangesStream = "changes_stream"
	TaskTypeSyncBinary    = "sync_binary"
	TaskTypeCreateHook    = "create_hook"
	TaskTypeTriggerHook   = "trigger_hook"
)

const (
	TaskStateWaiting    = "waiting"
	TaskStateProcessing = "processing"
	TaskStateSuccess    = "success"
	TaskStateFail       = "fail"
	TaskStateTimeout    = "timeout"
)
