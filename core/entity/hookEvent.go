package entity

type PublishChnagePayload struct {
	DistTag string
	Version string
}

type UnpublishChangePayload struct {
	DistTag string
	Version string
}

type DistTagChangePayload struct {
	DistTag string
}

type PackageOwnerPayload struct {
	Maintainer string
}

type DeprecatedChangePayload struct {
	Deprecated string
}

type HookEvent struct {
	ChangeId string
	Event    string
	FullName string
	Type     string
	Version  string
	Change   map[string]any
	Time     uint
}

func CreatePublishEvent(fullname string, changeId string, version string, distTag string)
