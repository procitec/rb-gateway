
	"gopkg.in/libgit2/git2go.v22"
	err = idx.AddByPath(repoFile2)
	checkFatal(t, err)
		"+%s"+
		"diff --git a/%s b/%s\n"+
		"new file mode 100644\n"+
		"index 0000000000000000000000000000000000000000..5716ca5987cbf97d6bb54920bea6adde242d87e6\n"+
		"--- /dev/null\n"+
		"+++ b/%s\n"+
		"@@ -0,0 +1 @@\n"+
		repoFile, repoFile, fileId, repoFile, repoFileContent,
		repoFile2, repoFile2, repoFile2, repoFileContent2)