package longhorn_test

import (
	"testing"

	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
)

func TestCrd(t *testing.T) {
	crdclient, err := apiextensionsclient.NewForConfig(nil)
	if err != nil {
		t.Fatal("instantiating apiextensions client failed")
	}

	NewLonghornVolumeCustomResourceDefinition(nil)

}
