package error

import (
	"github.com/mesh-master/foundation/pkg/z"
	"google.golang.org/grpc/status"
)

type error struct {
	level     z.ErrorSeverityLevel
	status    *status.Status
	createdAt int64
}

func (e error) Error() string {
	return e.status.String()
}
