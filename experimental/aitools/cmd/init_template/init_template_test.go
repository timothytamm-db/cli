package init_template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildJobConfigMapOmitsEmptyCatalog(t *testing.T) {
	configMap := buildJobConfigMap("test_project", "")
	_, exists := configMap["default_catalog"]
	assert.False(t, exists, "default_catalog should not be in map when catalog is empty")
}

func TestBuildJobConfigMapIncludesExplicitCatalog(t *testing.T) {
	configMap := buildJobConfigMap("test_project", "my_catalog")
	assert.Equal(t, "my_catalog", configMap["default_catalog"])
}
