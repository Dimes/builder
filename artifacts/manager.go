package artifacts

import (
	"builder/model"
	"io"
)

// Manager implementations can read/write artifacts to backing data store
type Manager interface {
	Setup() error // Idempotently creates any necessary structures for the manager, e.g. Dynamo tables
	OpenReader(artifact *model.Artifact) (io.ReadCloser, error)
	OpenWriter(artifact *model.Artifact) (io.WriteCloser, error)
}

// GCSManager stores artifacts in GCS
type GCSManager struct {
}
