package version

import (
	"github.com/hashicorp/packer/helper/version"
	packerVersion "github.com/hashicorp/packer/version"
)

var BreakpointPluginVersion *version.PluginVersion

func init() {
	BreakpointPluginVersion = version.InitializePluginVersion(
		packerVersion.Version, packerVersion.VersionPrerelease)
}
