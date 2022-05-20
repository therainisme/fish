package petpet

import "path/filepath"

const petpetExecPath = "./function/petpet/exec/"

func getQQJPGPath(qq string) string {
	return filepath.Join(petpetExecPath, "img", qq+".jpg")
}

func getQQGIFPath(qq string) string {
	return filepath.Join(petpetExecPath, "img", qq+".gif")
}
