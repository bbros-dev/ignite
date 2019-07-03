package imgmd

import (
	"path"

	api "github.com/weaveworks/ignite/pkg/apis/ignite/v1alpha1"
	"github.com/weaveworks/ignite/pkg/client"
	"github.com/weaveworks/ignite/pkg/constants"
	"github.com/weaveworks/ignite/pkg/metadata"
)

type ImageMetadata struct {
	*api.Image
}

var _ metadata.Metadata = &ImageMetadata{}

func NewImageMetadata(id string, name *string, object *api.Image) (*ImageMetadata, error) {
	if object == nil {
		object = &api.Image{}
	}

	md := &ImageMetadata{
		Image: object,
	}

	metadata.InitName(md, name)

	if err := metadata.NewID(md, id); err != nil {
		return nil, err
	}

	return md, nil
}

func (md *ImageMetadata) Type() api.PoolDeviceType {
	return api.PoolDeviceTypeImage
}

func (md *ImageMetadata) TypePath() string {
	return constants.IMAGE_DIR
}

func (md *ImageMetadata) ObjectPath() string {
	return path.Join(md.TypePath(), md.GetUID())
}

func (md *ImageMetadata) Load() (err error) {
	md.Image, err = client.Images().Get(md.GetUID())
	return
}

func (md *ImageMetadata) Save() error {
	return client.Images().Set(md.Image)
}
