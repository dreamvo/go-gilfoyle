/*
 * Gilfoyle server
 *
 * Cloud-native media hosting & streaming server for businesses.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type EntMediaFile struct {
	// CreatedAt holds the value of the \"created_at\" field.
	CreatedAt string `json:"created_at,omitempty"`
	// DurationSeconds holds the value of the \"duration_seconds\" field.
	DurationSeconds float32 `json:"duration_seconds,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph. The values are being populated by the MediaFileQuery when eager-loading is set.
	Edges *EntMediaFileEdges `json:"edges,omitempty"`
	// EncoderPreset holds the value of the \"encoder_preset\" field.
	EncoderPreset string `json:"encoder_preset,omitempty"`
	// Framerate holds the value of the \"framerate\" field.
	Framerate int32 `json:"framerate,omitempty"`
	// ID of the ent.
	Id string `json:"id,omitempty"`
	// MediaType holds the value of the \"media_type\" field.
	MediaType string `json:"media_type,omitempty"`
	// ScaledWidth holds the value of the \"scaled_width\" field.
	ScaledWidth int32 `json:"scaled_width,omitempty"`
	// UpdatedAt holds the value of the \"updated_at\" field.
	UpdatedAt string `json:"updated_at,omitempty"`
	// VideoBitrate holds the value of the \"video_bitrate\" field.
	VideoBitrate int32 `json:"video_bitrate,omitempty"`
}