package mount

import (
	"github.com/chrislusf/seaweedfs/weed/pb/filer_pb"
	"github.com/hanwen/go-fuse/v2/fuse"
)

func (wfs *WFS) AcquireHandle(inode uint64, uid, gid uint32) (fileHandle *FileHandle, status fuse.Status) {
	var entry *filer_pb.Entry
	_, _, entry, inode, status = wfs.maybeReadEntry(inode, true)
	if status == fuse.OK {
		// need to AcquireFileHandle again to ensure correct handle counter
		fileHandle = wfs.fhmap.AcquireFileHandle(wfs, inode, entry)
	}
	return
}

func (wfs *WFS) ReleaseHandle(handleId FileHandleId) {
	wfs.fhmap.ReleaseByHandle(handleId)
}

func (wfs *WFS) GetHandle(handleId FileHandleId) *FileHandle {
	return wfs.fhmap.GetFileHandle(handleId)
}
