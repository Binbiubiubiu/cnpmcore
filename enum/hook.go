package enum

const (
	HookTypePackage = "package"
	HookTypeScope   = "scope"
	HookTypeOwner   = "owner"
)

const (
	HookEventTypeStar         = "package:star"
	HookEventTypeUnstar       = "package:unstar"
	HookEventTypePublish      = "package:publish"
	HookEventTypeUnpublish    = "package:unpublish"
	HookEventTypeOwner        = "package:owner"
	HookEventTypeOwnerRm      = "package:owner-rm"
	HookEventTypeDistTag      = "package:dist-tag"
	HookEventTypeDistTagRm    = "package:dist-tag-rm"
	HookEventTypeDeprecated   = "package:deprecated"
	HookEventTypeUndeprecated = "package:undeprecated"
	HookEventTypeChange       = "package:change"
)
