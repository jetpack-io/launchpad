package jetconfig

import (
	"github.com/pkg/errors"
	"go.jetpack.io/launchpad/padcli/semver"
)

const (
	// These 1.0 and 1.1 legacy versions were present before we made the ConfigVersion
	// follow the semver versioning format.
	legacyVersionOneDotZero = "1.0"
	legacyVersionOneDotOne  = "1.1" // This was generated by jetpack-init for a couple days
)

// singleton (unenforced) for this package
var Versions = version{
	min:  "0.1.2", // update this when we drop support for upgrading these older versions
	prod: "0.1.2", // update this when ready to upgrade everyone to the new official version
	dev:  "0.1.2", // update this when developing a new version
}

type version struct {
	min  string
	prod string
	dev  string
}

func (v version) MinSupported() string {
	return v.min
}
func (v version) Prod() string {
	return v.prod
}
func (v version) Dev() string {
	return v.dev
}

func isVersionLessThanMinimumSupported(cfgVersion string) (bool, error) {
	if cfgVersion == legacyVersionOneDotZero || cfgVersion == legacyVersionOneDotOne {
		return true, nil
	}

	// This comparison ensures that this function is kept forward-compatible
	cmp, err := semver.Compare(cfgVersion, Versions.MinSupported())
	if err != nil {
		return false, errors.WithStack(err)
	}
	return cmp < 0, nil
}

func doesVersionNeedUpgrade(cfgVersion string) (bool, error) {
	// This comparison ensures that this function is kept forward-compatible
	cmp, err := semver.Compare(cfgVersion, Versions.Prod())
	if err != nil {
		return false, errors.WithStack(err)
	}

	// This means that cfgVersion < Versions.Prod
	return (cmp < 0), nil
}
