package vcs

import (
	"testing"
)

func TestVCSLookup(t *testing.T) {
	// TODO: Expand to make sure it detected the right vcs.
	urlList := map[string]struct {
		work bool
		t    VcsType
	}{
		"https://github.com/masterminds":                                   {work: false, t: GitType},
		"https://github.com/Masterminds/VCSTestRepo":                       {work: true, t: GitType},
		"https://bitbucket.org/mattfarina/testhgrepo":                      {work: true, t: HgType},
		"https://launchpad.net/govcstestbzrrepo/trunk":                     {work: true, t: BzrType},
		"https://launchpad.net/~mattfarina/+junk/mygovcstestbzrrepo":       {work: true, t: BzrType},
		"https://launchpad.net/~mattfarina/+junk/mygovcstestbzrrepo/trunk": {work: true, t: BzrType},
		"https://git.launchpad.net/govcstestgitrepo":                       {work: true, t: GitType},
		"https://git.launchpad.net/~mattfarina/+git/mygovcstestgitrepo":    {work: true, t: GitType},
	}

	for u, c := range urlList {
		ty, err := detectVcsFromUrl(u)
		if err == nil && c.work == false {
			t.Errorf("Error detecting VCS from URL(%s)", u)
		}

		if err == ErrCannotDetectVCS && c.work == true {
			t.Errorf("Error detecting VCS from URL(%s)", u)
		}

		if err != nil && c.work == true {
			t.Errorf("Error detecting VCS from URL(%s): %s", u, err)
		}

		if c.work == true && ty != c.t {
			t.Errorf("Incorrect VCS type returned(%s)", u)
		}
	}
}
