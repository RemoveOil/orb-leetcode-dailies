package removesubfolders

import "slices"

func removeSubfolders(folders []string) []string {

	slices.Sort(folders)
	var parents []string
	for len(folders) > 0 {
		parent := folders[0]
		folders = popChildren(parent, folders)
		parents = append(parents, parent)
	}
	return parents
}

func popChildren(parent string, folders []string) []string {
	parentWithSlash := parent + "/"
	for i := 0; i < len(folders); i++ {
		folder := folders[i] + "/"
		if len(folder) < len(parentWithSlash) || folder[0:len(parentWithSlash)] != parentWithSlash {
			return folders[i:]
		}
	}
	return nil
}
