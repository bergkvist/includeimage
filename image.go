package includeimage

import (
	"encoding/json"

	"github.com/docker/distribution/reference"
	"github.com/moby/buildkit/client/llb"
	"github.com/moby/buildkit/client/llb/imagemetaresolver"
	"github.com/moby/buildkit/frontend/dockerfile/dockerfile2llb"
	"github.com/moby/buildkit/util/appcontext"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
)

func LoadImage(imageName string) (string, dockerfile2llb.Image, error) {
	var img dockerfile2llb.Image

	ref, err := reference.ParseNormalizedNamed(imageName)
	if err != nil {
		return "", img, err
	}
	ref = reference.TagNameOnly(ref)

	// TODO: Get current platform
	// TODO: Change LogName
	_, dt, err := imagemetaresolver.Default().ResolveImageConfig(appcontext.Context(), ref.String(), llb.ResolveImageConfigOpt{
		Platform:    &specs.Platform{Architecture: "amd64", OS: "linux"},
		LogName:     "test",
		ResolveMode: "default",
	})
	if err != nil {
		return ref.String(), img, err
	}

	if err = json.Unmarshal(dt, &img); err != nil {
		return ref.String(), img, err
	}

	return ref.String(), img, nil
}
