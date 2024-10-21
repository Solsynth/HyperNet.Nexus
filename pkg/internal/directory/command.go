package directory

const (
	CommandMethodGet    = "get"
	CommandMethodPut    = "put"
	CommandMethodPatch  = "patch"
	CommandMethodPost   = "post"
	CommandMethodDelete = "delete"
)

type Command struct {
	// The unique identifier of the command, different method command can hold the same command id
	ID string `json:"id"`
	// The method of the command, such as get, post, others; inspired by RESTful design
	Method string `json:"method"`
	// The tags of the command will be used to invoke the pre-command middlewares and post-command middlewares
	Tags []string `json:"tags"`
	// The implementation of the command, the handler is the service that will be invoked
	Handler []*ServiceInstance `json:"handler"`

	RobinIndex uint `json:"robin_index"`
}
