package v1alpha5

import (
	"testing"

	"github.com/weaveworks/eksctl/pkg/testutils"
)

func TestAPIs(t *testing.T) {
	testutils.RegisterAndRun(t)
}
