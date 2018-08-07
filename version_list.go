package main

type VersionList []*Version

// implement Len() Less() Swap() for sort

func (v VersionList) Len() int {
	return len(v)
}

func (v VersionList) Less(i, j int) bool {
	return v[i].LessThan(v[j])
}

func (v VersionList) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
