package openai

const (
	RoleSystem    = "system"
	RoleAssistant = "assistant"
	RoleUser      = "user"
	RoleTool      = "tool"
)

const (
	ToolTypeFunction = "function"
)

const (
	FinishReasonToolCalls = "tool_calls"
	FinishReasonStop      = "stop"
)

const (
	ContentPartTypeText     = "text"
	ContentPartTypeImageUrl = "image_url"
)

const (
	ImageUrlDetailLow  = "low"
	ImageUrlDetailHigh = "high"
	ImageUrlDetailAuto = "auto"
)

const (
	ImageSize256x256   = "256x256"
	ImageSize512x512   = "512x512"
	ImageSize1024x1024 = "1024x1024"
	ImageSize1792x1024 = "1792x1024"
	ImageSize1024x1792 = "1024x1792"
)

const (
	EncodingFormatFloat  = "float"
	EncodingFormatBase64 = "base64"
)

const (
	ResponseFormatJSONObject = "json_object"
	ResponseFormatText       = "text"
	ResponseFormatUrl        = "url"
	ResponseFormatB64JSON    = "b64_json"
)

const (
	PurposeFineTune   = "fine-tune"
	PurposeAssistants = "assistants"
)

const (
	DefPropertyTypeObject  = "object"
	DefPropertyTypeArray   = "array"
	DefPropertyTypeNull    = "null"
	DefPropertyTypeString  = "string"
	DefPropertyTypeNumber  = "number"
	DefPropertyTypeInteger = "integer"
	DefPropertyTypeBoolean = "boolean"
)
