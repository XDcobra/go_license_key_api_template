package Redis

type RedisPayloadModel struct {
	// Redis Key
	Key string
	// Redis key value
	Value string
}

type RedisGetResponseModel struct {
	// Redis Key
	Key string
	// Redis key values
	Values []string
	// any errors if available
	Errors string
}

type RedisPostResponseModel struct {
	// Redis Key
	Key string
	// Redis key value
	Value string
	// any errors if available
	Errors string
}
