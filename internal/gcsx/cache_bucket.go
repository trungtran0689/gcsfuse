package gcsx

import (
	"path"

	"github.com/jacobsa/gcloud/gcs"
	"golang.org/x/net/context"
)

func NewCacheBucket(b gcs.Bucket) gcs.Bucket {
	return cacheBucket{b}
}

type cacheBucket struct {
	gcs.Bucket
}

func (b cacheBucket) CreateObject(
	ctx context.Context,
	req *gcs.CreateObjectRequest) (o *gcs.Object, err error) {

	if path.Ext(req.Name) == "m3u8" {
		req.CacheControl = "no-cache"
	}

	// Pass on the request.
	o, err = b.Bucket.CreateObject(ctx, req)
	return
}

func (b cacheBucket) ComposeObjects(
	ctx context.Context,
	req *gcs.ComposeObjectsRequest) (o *gcs.Object, err error) {
	// Guess a content type if necessary.

	// if path.Ext(req.DstName) == "m3u8" {
	// 	req.Metadata = "no-cache"
	// }

	// Pass on the request.
	o, err = b.Bucket.ComposeObjects(ctx, req)
	return
}
