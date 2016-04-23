// Code generated by go-bindata.
// sources:
// db/drop_all_tables.sql
// db/migrations/0001_initial.sql
// db/migrations/0002_use_tz_in_timestamps.sql
// db/migrations/0003_package_channels_blacklist.sql
// DO NOT EDIT!

package api

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _dbDrop_all_tablesSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x91\xcd\x8a\x83\x30\x10\xc7\xef\x3e\x45\xde\xc3\x93\xab\x2e\x08\x0b\xbb\xac\x1e\x7a\x1b\xc6\x74\xd0\xd0\x98\x84\x64\x14\x7c\xfb\x06\x69\x2f\xa5\x74\x3c\x05\xc2\x6f\x66\xfe\x1f\xcd\xff\xef\x9f\x1a\xaa\xaf\x9f\x56\x75\xdf\xaa\xbd\x74\xfd\xd0\x2b\x26\x5c\x54\x5d\xf5\x75\xd5\xb4\x65\xf1\x16\x59\x13\xc5\x24\x30\x18\x82\x35\x1a\xd9\x78\x27\x90\x01\xf5\x0d\x27\x12\x28\xed\x23\xf9\x04\xa8\x4f\x6c\xd4\x33\x3a\x47\x56\xa0\xa6\xe8\xd7\x20\xd9\x30\x2e\x31\x3a\x2d\xa9\x7b\x62\x90\x1f\x5e\xcf\x2e\x85\xf3\x21\xbd\x1c\x80\xd9\x24\xf6\x71\x17\xa6\x68\x23\xc7\xc0\x7b\x90\xf4\x1f\xa0\xd4\x68\x8e\x7e\x33\x2c\xdd\xbc\x22\xe3\x88\x89\x60\x31\x53\x3c\xac\xa5\xcf\xc5\xc3\xa3\x2e\x18\x6d\xfe\xb0\xd9\x58\x59\xdc\x03\x00\x00\xff\xff\x1c\x61\x2a\xd3\x9a\x02\x00\x00")

func dbDrop_all_tablesSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbDrop_all_tablesSql,
		"db/drop_all_tables.sql",
	)
}

func dbDrop_all_tablesSql() (*asset, error) {
	bytes, err := dbDrop_all_tablesSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/drop_all_tables.sql", size: 666, mode: os.FileMode(420), modTime: time.Unix(1461201771, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dbMigrations0001_initialSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x5c\x6d\x57\xdb\xb8\xb6\xfe\x7c\xf8\x15\x9a\x75\x3f\x40\x57\x31\x58\xb6\x6c\x4b\x9c\x7b\xce\x5a\x34\x84\xc2\x94\x42\x79\x6b\x2f\xf3\x25\x4b\x96\xe4\xc4\xc5\xb1\x33\xb6\x03\x85\x5f\x7f\xb7\x1c\xc7\x38\x24\x01\x27\x90\x96\xe9\x9a\x81\x26\xb6\x5e\xf6\xf3\xec\x17\xed\x2d\xcb\x18\x06\x7a\xdf\x0f\xbb\x29\xcf\x15\xba\x1c\xac\xad\xc1\xf7\x56\x92\xaa\xb3\x24\x8a\x54\x8a\x32\xd1\x53\x7d\xbe\xb6\xd6\x3a\x6b\xef\x5e\xb4\xd1\xc5\xee\x87\xa3\x36\xca\x15\xef\xa3\x8d\xb5\x7f\x85\x12\x5d\x5e\x1e\xee\xa1\x2f\x67\x87\x9f\x77\xcf\xae\xd0\xa7\xf6\x15\xda\x6b\xef\xef\x5e\x1e\x5d\xa0\xe1\x30\x94\x9d\xae\x8a\x95\x1e\xb8\x73\x43\x36\xde\x6d\xae\xfd\x2b\xe6\x7d\x85\x6e\x78\x2a\x7a\x3c\xdd\xb0\x9c\x77\xe8\xf8\xe4\x02\x1d\x5f\x1e\x1d\xa1\xd6\x41\xbb\xf5\x09\x6d\x14\x0d\xfe\xf7\xbf\x68\x7d\xfd\x1d\xba\x3c\x3e\x3c\xbd\x6c\x43\x2f\x91\x2a\x18\x43\x76\xf2\x0c\xe5\x61\x5f\x65\x39\xef\x0f\xd0\xb7\xc3\x8b\x83\x93\xcb\x0b\x74\x71\xf8\xb9\x8d\xfe\x3a\x39\x6e\x57\x33\x6f\xc4\xc9\xed\xc6\x3b\xc4\xf3\xa2\x35\xba\x4f\x62\x85\xd6\x87\xb9\x58\x7f\x98\x6e\xed\xdd\xbf\x1f\x41\x1a\x66\x2a\xcd\x96\xc2\xa4\x7b\x3e\x8b\xab\x6a\xf4\x18\x5b\xa6\x00\x5d\x5e\xf5\x75\xcc\xe9\xbe\x65\x93\x51\xcf\x95\xd0\x01\x83\x6a\x95\x76\xc6\xd8\x2b\x09\xce\xda\xfb\xed\xb3\xf6\x71\xab\x7d\x5e\xea\x3c\x94\xef\xa6\xb9\xe3\x83\x41\x14\x0a\x9e\x87\x49\xfc\x72\xab\x98\xc5\x40\x8d\x39\x68\x2e\x55\x26\xd2\x70\x50\x4c\x97\xab\x1f\xf9\xaf\xa7\x04\x1a\x8f\x14\x3a\x12\x75\x13\x95\x5d\x67\x70\x35\xe0\xe2\x9a\x77\xd5\x52\x3c\xe5\x77\x03\x85\xc2\x38\x9f\xe2\xa7\xb8\xf1\x5f\x64\xea\x46\x37\x60\xc7\x9a\x99\xf3\xf6\xe7\xaf\xed\xb3\x3a\x9e\x61\x1a\xd5\x6c\xd4\x9d\x61\xa4\xd0\xa0\x62\x39\x08\x23\x35\xa1\x18\x6c\x9a\x73\xd8\xcf\xc2\xfb\x9a\xf5\x17\xad\x7a\x3c\xeb\x55\x97\x5c\xb2\x32\xb3\xad\x99\xde\x93\xaa\x9a\x30\x51\x50\x0c\x3a\x39\x86\x99\x8e\xda\xa0\x97\xd6\xee\x79\x6b\x77\xaf\x5d\xe9\x70\xa3\x64\x70\x13\x4d\x0e\x3e\x52\x26\x44\xc7\x8b\x93\xbd\x93\x1d\x94\xaa\x9b\x50\xdd\xa2\x20\x54\x91\x04\x44\xa0\x81\x0c\xf1\x58\x22\x60\xad\x9b\xf7\x32\xd0\x13\xca\x7b\x21\xdc\xe1\x7e\xa4\xd0\x1f\x7f\xfc\x31\x69\x07\x02\x42\x6c\x92\x75\xb8\x58\xda\x6b\xd4\x8d\x8a\xf3\x09\xda\x11\x90\xdc\x4b\x93\xbe\x1e\x78\x6c\x06\x8f\xee\x67\x3d\x0e\xaa\x7f\xa4\x99\x58\x29\x09\xa2\xc8\x3e\xc8\xec\x27\x49\xa4\x78\x0c\x57\xc3\xac\x23\x55\x94\xf3\xda\x25\x19\x66\x1a\x4d\x67\xc0\xef\xa2\x84\xcb\x8e\x0f\xc6\x9c\x04\x41\xad\x45\x5f\xe5\x5c\xf2\x9c\x77\xb2\xb0\x1b\xf3\x7c\x98\xaa\x4e\x9a\xf1\x09\xb3\x9b\x6c\x75\x3f\xc3\xc0\xb8\x8c\xc2\x78\xea\xfa\x4a\xec\xa7\x74\xc7\x27\x6d\xa7\x72\xd9\x99\x76\x33\xed\xe2\x20\x74\x1c\xab\x68\xe5\x0b\xa4\xe6\x24\x89\x92\x74\x66\xfb\x7f\x84\xc3\x3d\x66\xff\x79\xd2\xcf\xdb\x95\x24\x93\x11\x77\x96\xaf\x4e\x68\xa5\x9b\x26\xc3\xc1\x72\x2b\xfc\x0b\xd7\xa7\x3a\x7d\x29\x64\x55\xc9\x30\xef\x84\x71\x67\x90\x26\xdd\x54\x65\xd9\xd8\x7b\x2a\x19\x02\x1e\x65\x6a\xc2\x48\x13\x80\x76\xd7\x19\x0e\xc0\x67\x54\xd6\x81\xb0\x0c\x4e\x28\xa7\xfa\xe5\xe9\x70\x56\xb7\x8c\x07\xaa\xd3\x4f\xa4\x6a\xda\x01\x3c\x3a\x14\xaa\xd3\x4b\x86\x69\x73\xe1\xb4\xcd\x14\x26\x33\xe6\x89\x14\x4e\x5b\xde\x1d\xa8\x34\x4c\x24\xa0\xce\x55\x7a\xc3\xa3\x7a\x54\x9a\x22\x73\x4e\x97\x8a\xde\xf2\x7e\x9f\xff\xa8\x18\x81\xb6\x65\x7b\xbd\x3c\xaa\x2e\xe4\xad\x73\x46\x9d\xd3\xab\x5c\x3b\x27\x88\x2e\x20\x81\xae\x9a\x08\xfb\xa8\xc7\x6a\x53\xb5\x57\x75\xc1\x32\x58\xcd\x72\xc1\x2a\x8e\xbd\x96\x0b\x86\x31\x10\x10\x8b\x71\xf2\x53\xf7\xa8\xba\x2f\x96\xd4\x42\x8b\x8a\xc7\x70\x00\x9d\x55\xbe\xda\xe8\x36\x5f\xe0\x0e\xfc\x93\x0f\xc7\xc1\x03\x52\xf9\x10\x0c\xb2\x26\xf2\x54\xe0\x36\x27\x03\xf1\x53\x51\x3a\x14\x93\xab\x74\x13\x71\x1e\x25\xdc\xf3\x53\xbe\x95\xd8\x5f\x49\x46\xe9\x69\x70\x21\xe2\x59\xde\x81\x3a\x51\x5c\x77\x82\x24\x1d\xfb\xd7\x2b\xce\x58\x4c\x50\xfa\x18\x14\xaa\xf1\xf3\x90\x1e\x75\x9a\x64\x48\xe7\xc2\xa3\xeb\x0b\x85\xe1\x8a\xfe\x47\xb6\x3b\xcb\xf7\x1e\x4c\x7d\x9e\xe3\xbd\xaa\x17\x17\x8b\xdb\x2c\x1f\x1e\xaf\x7a\x73\x5d\xb8\xee\x78\x1b\x35\x84\x0b\xb8\x73\xe9\x1d\x1d\xc8\x76\xf3\x24\xbd\x7b\xca\x4b\xa6\x4c\x67\x4a\x31\x2b\xb1\xd8\xdf\x47\x73\x73\x93\xce\xa2\x1a\xe8\x14\x85\xe0\x13\xfc\x8f\x2b\xc8\x89\x25\x52\xa7\x25\x2a\x1b\x46\xf9\xac\x3b\xf5\x64\xa6\x9e\x90\x3f\x11\xa6\x46\x85\xc9\x13\x52\xac\x26\x93\xd7\x05\x59\x32\x9c\xae\x7c\x70\xb1\xb4\xab\x34\x85\xd0\x24\x74\x1a\xf4\xa8\xae\x78\xc3\xc6\xf1\xa0\x54\x3d\xd6\x54\x6a\x53\x1b\xae\xae\x7e\x3d\xda\x8c\x4d\x1a\xa8\x34\x6f\xc2\xfc\x49\xf7\x5c\x89\x62\x04\x04\xe2\x6c\x96\x69\x65\x20\x74\xaa\x25\x9a\x71\x6f\xfe\x92\xf6\x06\xbc\x6f\xa9\xb4\xe9\xa1\xef\x3c\x83\x6b\x6e\x67\xe3\x9d\x88\xc3\x38\xcc\xb5\x16\x75\x2d\x5d\x5c\xd8\x53\x01\xd7\x8e\x5c\xec\x4c\xe9\xdd\x08\xbd\xef\x88\x36\x8a\xd2\x7e\xbb\xf8\xfd\x6e\xed\xf0\xf8\xbc\x7d\x76\x81\x0e\x8f\x2f\x4e\xaa\x1d\xac\x4d\xa4\xf3\x97\x77\xe8\xeb\xee\xd1\x25\x4c\xbf\xb1\x2e\x29\xb3\x89\x25\x85\xc1\x2c\x4c\x0c\x42\xb0\x34\x38\xe1\x81\xe1\x4b\x49\x6d\x8f\xdb\xbe\x65\xb3\xf5\x4d\xb4\x2e\x47\x13\xae\x83\x40\xf5\x71\xcb\xdd\xd4\xf1\xae\xe7\x26\x1a\xed\x61\x3e\x6c\x8c\x3d\xcc\x54\x48\xa5\x87\xa2\xbe\x8d\x2d\x66\x49\xe2\x79\xd4\xa1\x96\x30\x95\x13\x70\xe6\x72\xa5\x1c\x07\xdb\x01\x2e\xa6\x6b\x22\x56\x49\x4e\xbb\x08\x43\xc5\xc6\xcc\x84\x6c\x75\x57\xd1\xbf\x37\xd1\x28\xfa\x6d\xa2\x5a\xac\x7b\x10\xd0\xde\x44\x26\xcc\x7c\x38\xd6\x48\xaa\x06\x49\x0a\x4e\x02\xf4\xa2\x22\xa8\x20\x39\x4c\xc3\xb8\xab\xbf\x8f\xb2\x0a\x94\xe5\x6a\xb0\xf5\x98\x92\xc5\xa7\xc5\x30\xed\x65\x31\x62\x8a\x7a\x3c\x43\x90\xa9\x08\x48\x54\x8a\xa9\xe5\xc8\xbe\xe1\x73\x59\x28\xbf\xc2\x7c\x56\x0d\x66\x06\x58\x20\xd3\x92\x30\x41\x9e\x20\x31\x4c\x53\x4d\xe6\xd8\xb4\x4b\xef\x7c\xe9\x9c\xb8\x04\xb9\x97\xdc\xc6\x7a\x3f\x49\xb3\x18\xe9\xd4\x31\x7f\xb5\x19\x48\x9d\xc6\x6a\x53\x81\xa7\x69\x78\x03\xd0\xb2\xa1\xd0\x8c\x06\xc3\x28\xba\x7b\xe9\x54\xd4\x34\x47\x73\x15\x14\x46\xd1\x78\xf0\x2d\x54\x4e\x2e\x92\xfe\x20\x52\x45\x3c\xd2\x6b\x95\xd2\x59\x2c\xf2\xef\x2a\x5f\xdf\x1a\x1b\xae\x7e\xfa\x72\x32\x11\xc0\x26\xe4\x7a\x14\xd8\x46\xae\x3b\x21\xd4\x2c\x37\x53\xcc\xb5\x28\xe6\xae\x21\x31\xf8\x0b\xf1\xa5\x32\x18\x37\xb9\xc1\x3c\xdf\x73\x95\xe3\x4a\xe1\x78\xda\xc3\x46\x93\xeb\x4f\x47\x61\x3c\xfc\x81\x20\xa1\x47\x7d\x08\xe0\xc0\x97\x5e\x33\x40\x2f\x30\xd3\x20\x4a\xee\xfa\x00\x20\x5b\xc4\x27\xeb\x18\xc6\x8a\xa8\xc4\xb3\x7c\x4e\x04\xa3\xc4\x70\x14\xf3\x0d\x82\xb1\x32\x7c\x4f\xd8\x86\x6f\x2b\x1f\x93\xc0\xe3\x96\xab\x03\x80\xe6\xd7\x73\xdd\x2d\x7b\xcb\xd4\x53\xf7\xf2\x7c\x90\xed\x6c\x6f\x03\xb5\xfd\x24\xd6\x51\x50\x67\xa0\xda\x1b\xba\x49\xd2\x85\x6c\x7e\x10\x66\x5b\x70\x73\x7b\xe4\x98\xc6\xf8\xae\xde\x7c\x35\x92\x6c\x0b\x4a\xca\x6d\xde\x97\x2e\x31\x86\x59\xba\x5d\x0e\xbc\xad\x47\x1e\x75\xd8\xea\xde\xc3\x97\x62\xe1\x41\xeb\xd8\x21\xcc\xf5\x88\x43\xf5\xfd\x88\x7c\xba\xf5\x76\xd5\x87\xa3\xf4\xeb\xe1\x1e\xfb\xd3\x0f\xee\x3e\x27\x7f\xaa\x0f\xce\xdd\xa7\xee\x7f\xf4\x7d\xcb\xc4\x8e\x61\x32\xc3\x32\x91\x69\xee\x60\x6b\xc7\xf6\xb6\x1c\x20\xc2\x2e\x7a\x37\x52\xc6\x73\x94\xd9\xb6\xe7\xdb\x81\xa7\x8c\x20\xb0\x98\x41\x3c\x45\x0d\x6e\x3a\x96\x11\x98\xd4\x26\xd2\x72\x7c\xe9\x3b\x35\xca\xc8\xaa\x28\x23\x4f\x51\xe6\x98\x98\x32\x6c\xe9\xfb\x41\x7a\x7d\xcd\xdf\x7f\xd8\xbe\xbf\xb8\xf1\x4e\xbe\x7c\xeb\x86\xf2\xea\xfd\x35\x39\x8f\xf7\xce\xa7\x29\x73\x77\xb0\xb3\x63\xb1\x2d\x6c\x52\xcb\x75\x5f\x8d\x32\x61\x71\xdb\x85\x55\xc6\xf0\x19\x05\xca\x4c\x5b\x19\xdc\x77\x3c\xc3\x74\x85\xe9\x38\xdc\x53\x5c\x58\x25\x65\xd4\xa4\x5b\xe6\x2a\x28\x2b\x07\x9e\x4b\x99\xe7\x79\xd8\x23\x98\xe8\xfb\xfe\xdf\x76\x70\x7a\x76\xf0\x85\xfe\xf8\x60\x9f\xed\x5f\x7e\x6f\xc9\xdd\xc0\xbe\x3d\xbd\x6a\x59\xed\x19\x56\x66\xb2\x1d\xd3\xdd\xa2\x36\x03\x70\xaf\x46\x19\xb1\x1d\x90\x98\x59\x86\xe0\x92\x1a\xc4\xa5\x1c\xac\xcc\xf7\x0d\xe5\x33\x69\x2a\x93\x29\xc1\xc9\x98\x32\xec\xac\x88\xb2\xd1\xc0\xf3\x29\xa3\x2e\x88\xe9\x15\xa0\xaf\x8f\x09\xef\x27\x9f\xae\xbe\xfe\x75\xf9\xd1\xfa\x96\x9c\xcb\xd3\x03\xfc\xe5\xe0\xcb\x7d\xea\xec\x4e\x52\xe6\x20\x6c\xef\x38\x60\x65\xe6\x96\xa6\x1b\xbf\x1e\x65\x16\x05\xf7\x63\x8e\x6f\x38\x98\xc2\x20\xd2\xf5\x0c\xc6\x18\x8c\x04\xf1\x83\x02\x69\x4a\x32\x73\x4c\x19\xd8\xf8\x6a\x28\x1b\x0d\x3c\x97\x32\xea\x5a\x04\xd2\xaa\xc2\xca\xac\xa8\x97\x5c\xde\xdc\xc4\xc9\x95\xed\xb0\x41\x68\xed\xc7\xfc\x7c\xfb\x47\xd6\xcd\xc3\x9a\x63\x62\x13\xfe\x47\x96\xbd\x83\xf1\x0e\x36\xb7\xa8\xe5\x30\xea\x2c\x4b\xd9\x38\x8f\x78\x58\x9d\x4c\xd7\x74\x09\x97\x7a\xf9\x80\x35\x84\x99\x04\x06\x71\x21\xb4\x49\xe2\x3a\x0e\xb3\x25\xf6\x8b\x18\x92\x15\x4f\xd3\xf4\xa7\xff\xc1\x04\x2c\xd0\xad\xab\x14\x33\x64\x3a\xda\x0b\x6c\xb2\x05\xcb\x85\x45\x70\x63\xf9\xa0\x5d\xa3\x70\xfa\x1c\x0e\x6c\x51\x9f\x0a\xe8\xef\x98\x8e\xf6\x16\x62\x1b\x54\x81\xfe\x39\xe6\xa6\xb2\x85\x4b\xb0\x28\xec\xcc\x57\x39\x2f\x50\x04\xc2\x0b\x6c\x7b\x3e\x0a\x62\xdb\xe4\xa7\xa3\xe0\xd4\xe3\xa6\xad\xb5\xa1\x57\x64\xe2\x71\x6c\x50\x29\x88\x61\x3b\xa6\xe7\x73\xc5\xb0\x52\x05\xb5\x3c\x1a\xf4\x46\x30\x70\xe0\xfb\xf4\x09\x65\x38\x9e\xb3\x10\x8c\x46\x2e\xf4\x08\x46\x59\xb2\x55\x28\x18\xb7\xa4\xf2\x3d\xd3\x00\x4a\x14\x04\x7b\xcb\x35\xa8\x63\x43\x5e\x12\x48\xe9\xfa\x36\xf1\x7c\xbf\xb0\xa4\xf3\xca\xa6\xf6\x21\xd7\x81\x24\x5b\x0e\x47\x0f\x65\x45\x34\x84\x4c\x3e\xd5\x49\x4e\xb1\x17\xb8\x59\x3c\x30\x19\xff\x2e\x2f\xad\xef\x42\xa3\x94\x47\x21\xdf\x3e\xbf\x93\xb1\xba\x5b\x2f\xd6\x3d\x04\x45\xcd\x10\x32\xd9\xf5\x51\x5a\xed\x9a\xb5\x0b\x73\x38\x62\xa6\x6b\x2d\xc2\x51\x23\x9f\x79\x86\x23\x3b\x50\xd8\x84\xae\x86\x94\x9e\x0d\x73\x11\x06\x69\x97\xc5\x0d\x8b\x72\x81\x19\x17\x81\x90\x42\xcf\xf5\xa1\xb4\xd7\x2f\x69\xd2\x4f\x8a\xe2\x47\xab\x1e\x12\x62\x88\x4b\x99\xca\x36\x8b\x22\x81\xe7\xa2\x87\xfc\x61\x37\x43\xd9\x40\x89\x30\x08\x85\xbe\x7c\x97\x0c\x53\xc8\x7d\xe3\x20\xec\x0e\xd3\x22\x6b\x5d\x9f\x24\x72\x35\x74\x7a\xb6\x45\x16\x32\xb9\x46\xae\xfb\x0c\x9d\x8e\x4f\xb1\xe9\x52\xd3\x50\xb6\xcb\x0d\x42\x3d\xa6\x33\x0d\xf8\x14\xc0\x92\xac\xa8\xe9\x53\x56\xb8\xff\xee\xd8\x71\x2e\x52\x58\x3c\xb2\xaa\xbe\x92\x50\x16\x44\xc9\x40\xa7\xd6\xe8\x36\x49\xaf\x8b\x42\x2f\xcc\xc6\x3c\x4b\x14\xa4\xea\xef\x21\xdc\x8d\xee\x5e\x68\x94\x7a\xf1\xb1\x1b\xb0\x48\x18\x5e\x28\x8a\x36\x0a\x1d\x8f\xe3\xcf\xc4\x51\x88\x8a\x4c\xdf\xf2\xb1\xab\x2c\x65\x38\x5e\x00\x0a\xf1\x3c\xc7\xa0\x96\x17\x18\x34\x30\x7d\x8c\x61\x2c\x5f\x16\xfe\x32\x48\xb2\x3c\x1c\xd5\x5c\xfa\x6b\x91\x04\xd0\xd6\x07\x9a\x7f\x56\xe6\x67\xba\x77\xf7\x97\x73\xf6\xd7\xed\xfd\xde\xd1\xdd\x85\xbc\x3e\xf8\x7e\xb2\x7d\xd5\x0d\x3e\x7d\x8d\xad\xb3\xee\xe5\xe7\xe4\x5a\xfc\xe7\x81\xc8\x09\x3e\xd7\xd7\x6b\x3f\x33\xb3\x79\xdb\x02\x46\x8a\xbb\x4d\x6a\x97\x66\x98\xa5\xc3\x2d\xe1\x07\xd0\x17\x0c\xc9\x20\x8a\x0a\x83\x52\x88\xe1\xae\x0c\x18\x0e\x04\xac\x07\xc2\x9d\x83\xf9\xf4\xf2\x63\xdc\xff\xe2\xe0\xde\xc0\xbb\xbf\x7b\xff\xfe\x7d\xe2\x04\x1f\x0e\x6f\xdb\xd1\x61\x7c\xb1\xdb\xcf\xbc\xed\xf8\x7b\x7c\xfd\x63\x98\xc7\xdb\xa7\x87\x0b\x63\xae\xd2\x71\xec\x8e\x16\xfd\xa5\x96\x99\xd9\x90\x2d\xc6\x84\x43\x24\x36\x60\xa5\x03\x8b\x01\xbd\x1a\x94\x4b\xcb\x70\x84\x6b\x4b\x30\x2c\x57\x8c\xb2\x94\x19\x90\xcf\x5b\x37\x94\x7d\xbc\xba\xff\xe1\x1d\xfe\x78\x7f\x11\x7d\xff\xdb\x3f\xce\xa4\x97\xc4\xae\x93\x24\xdf\xfe\xfe\x70\x2f\x5a\xe9\xd1\xfe\x11\xb9\x6d\xf5\x4e\x97\x50\xf3\x28\x9d\x66\xe0\x03\x45\xe9\xd9\xac\x78\x68\x86\xd9\x23\x54\x06\x4e\x20\x0c\x6c\x71\x07\x96\x37\xf0\x17\xee\x81\xd9\x40\xe0\x16\x02\xbb\x10\x41\x08\x9d\x83\x99\x1d\x5c\x66\xe4\xb6\x17\xde\x07\x77\x37\x3e\xe9\x77\xa3\xf7\xdf\xf8\x31\xff\x66\x7f\x3d\x3a\xbd\xba\xcd\xbe\xe1\x8f\x07\xc7\x07\x7f\x1e\x0f\xc4\x7e\x97\x2c\x86\xb9\x96\x0f\x43\x72\x67\x91\xc2\xb7\x1a\x65\xff\xcd\x30\x33\x21\x89\x47\x84\x63\xa8\x80\xdb\x3a\x3c\x30\xbd\xa4\x83\x9e\xa9\x94\x3e\x75\x2c\x25\xe9\x3c\x3d\xe3\x73\x76\x7f\xda\x3a\xfa\xf8\xbd\xaf\x82\xab\xb8\xbd\xcd\xf7\xc5\xa0\xf5\xfd\x08\x1f\x67\xdd\xe1\x41\xef\xf4\xe3\x77\xf3\xb2\xd5\x77\x02\xf3\x33\xa3\xcd\x31\x3f\x4a\x68\x19\xb6\x3d\x8f\x2e\x92\x7b\xe8\xad\x9a\x73\xae\xf7\x74\x26\xb6\x64\xf0\x2b\x6c\xd6\xf8\x2e\x01\xde\x4d\x08\x78\x01\x81\x30\xe0\x5a\x1e\x04\x13\xdb\x37\x7c\xe5\x99\x60\x28\x54\x10\x3a\x4a\x5d\xa6\xa6\xd7\x57\xff\x84\xe0\xaf\x77\x23\xeb\x73\xc3\x42\x9c\xf5\x92\x5b\xa4\x7f\x44\x92\x44\xf5\x23\xbe\xb0\xc0\xec\xbc\x7b\xe9\x4e\x4e\x01\x6e\xb4\x4d\x36\x4c\x23\x60\xbf\x3c\x3e\xb8\x89\xe6\x1d\xa7\xab\x2d\x9b\x98\x39\x3a\xf2\x19\x4e\x40\x2d\x03\xe2\xa7\x63\x30\x09\x89\x5b\xa0\x82\x80\x62\x30\x1c\xc1\x02\x90\x8f\x4c\x54\x49\xa9\x4a\x0b\xf1\xb7\x92\xb4\x5b\x94\x39\x7a\xdf\xb0\x83\xa1\xea\x29\x2c\xb7\xfa\xd0\x88\xcb\x15\x80\xc2\x90\xd1\x79\x01\x27\x1a\x94\xfd\x72\x50\xf6\x18\x94\xfd\x2b\x41\xc1\x9c\xc4\xd7\xa1\xc0\x09\x98\xf7\x72\x50\x64\x0c\x8a\x2c\x0b\xaa\xf6\xb0\x65\xec\x5b\xc5\x69\x8b\xc7\x18\x36\xd1\xc3\x59\xb3\xba\xa3\x05\xca\xb6\x7c\xc2\xb5\x92\xc4\x7c\x3c\xeb\x9f\xb9\xae\x04\x8a\x52\xc7\x34\x5b\x2d\xd3\x6c\x2c\x6f\xb1\x95\xd4\x80\xb4\xd7\xc5\x25\x7c\xa8\x7d\x14\xa7\xcf\x18\x5f\xad\xf6\x01\x5c\x8c\xed\xef\x2f\x82\xab\x91\x85\x3f\x93\x31\xfb\x82\x73\x97\xfa\xe2\x19\xd7\xd7\x85\x87\x44\xed\x96\x85\x86\x99\x71\x0b\xe6\x63\x58\xe3\xab\x65\xb1\x36\xda\x9b\x86\x4a\x44\xdf\x05\xae\x40\x61\x3f\xb9\x6e\x6b\x4a\x5b\x23\xdd\x3c\x43\x9b\x67\x7a\x04\x92\x41\x28\xd4\x4c\xe8\x49\x7c\x48\x3e\x99\xab\x24\x64\x4e\x20\x8f\xf4\xb9\x63\x5a\xf8\x31\x6d\x50\x41\xe4\x06\x9e\x47\x9b\xbe\xfb\xdb\xd3\x06\x75\x83\x49\xb1\xcd\x9f\xb3\xb6\x53\x6e\xec\xa9\x9b\xe2\xd3\x6e\x51\x83\xd5\x0b\xb3\x8a\xb2\x8a\x80\xb7\x48\x57\xa3\xc8\xf6\x88\xae\xfa\x53\xe7\x4d\x14\x0e\x6a\x01\x65\x7c\xab\xb0\x1f\x48\x97\xf4\x7f\x53\x45\x4d\x83\xfe\x56\xad\xff\x54\xb6\xdc\xa0\xbf\x5d\xeb\x6f\x2f\xd1\x9f\xd4\xfa\x4f\x65\xae\x0d\xfa\x3b\xb5\xfe\x53\x15\x4e\x83\xfe\x6e\xad\xbf\xbb\x44\x7f\xaf\xd6\x7f\x6a\x5b\xb5\x41\x7f\x5a\xeb\x4f\x97\xe8\xcf\x6a\xfd\xa7\x72\xc1\x26\xf6\x63\xd6\x0d\x68\x6a\x0f\xaf\xc9\x08\x13\x26\x38\xd7\x06\x27\xce\x48\x3e\x24\x65\x65\x02\x55\x9e\x60\x29\x32\x95\xea\x73\xf5\x61\xb4\x49\x5e\x7d\x9d\x34\xfe\xc6\x0e\xd8\x64\x65\xfb\x89\xc2\x2f\x16\x3d\x56\x22\xbc\xb5\xb4\xf0\xcd\x93\xde\x37\xc8\x7c\xf3\xe4\x16\xda\x35\x5a\xd9\x7f\xa2\xf0\xce\xaf\x17\x7e\x79\xb3\x71\x7f\xbd\xf0\xb8\x2e\xbc\xbd\x88\xf0\xde\xaf\x17\x9e\x2c\xcd\x3c\x5d\xc8\x61\x9b\xa4\x65\x2f\xb3\x79\x6f\x11\xe1\xd9\xaf\x17\x7e\xc2\xe6\xad\x85\x16\xa9\x85\x6a\xd3\xd5\x48\xbf\xbc\xd1\xe3\xc5\xd6\xd8\x25\xa4\xaf\xce\xc1\x8e\x25\xae\x84\x98\x77\x9a\x15\x19\xa8\x7a\x0b\x69\xdd\x46\xc5\x7b\x51\xeb\xc5\x23\x14\xb2\xf0\x9e\x4c\x53\x57\x6d\x5a\xf0\x4c\x64\x27\xaf\x8b\xd4\xad\x90\x3a\x85\x12\x7f\x5f\xa4\xd8\xaa\xa0\x96\x27\x12\x7f\x5f\xa8\xb4\x82\x4a\x7e\x73\xfb\xb5\x48\x05\xd5\x1a\xd5\xd9\x6f\x17\xea\xbc\x57\x69\xa6\x90\xd7\x35\xd6\x88\x85\x07\x7d\xff\x82\x94\xee\x39\x34\xce\x92\x68\x50\x7d\x07\xe5\x0d\xe2\x72\x17\xc5\xc5\xde\xb2\x96\xbc\x25\xd1\x20\xf2\xc6\xd5\x64\xbd\x1c\x18\xc2\xa6\x7e\xb7\x21\x89\xe5\x1b\xc5\x38\x0e\x18\x56\x53\x8c\xb6\xfb\x96\x6d\xd1\x59\x16\xce\x3f\x25\x64\x34\x07\xe6\xbd\x65\x3d\x79\xcb\xc2\xf9\xc7\x04\x8d\x97\x20\xfb\x99\x51\x63\xce\x31\x05\xeb\x15\x8e\x29\x78\xd4\x94\x2e\x03\x29\x19\x27\x74\x74\x0c\x91\x52\x19\x18\xcc\xf1\xb9\x6b\xfb\xbe\x62\xc2\xd7\xa2\xce\x9a\x5e\x5f\xdf\x8d\x93\xbc\xa7\xff\xc6\xd8\xd4\xfd\x4d\x14\x28\x15\xe9\xe3\x6e\x4a\x1f\x5d\x48\x55\x3f\xb9\x51\xa8\x5f\x00\xff\x55\xa7\x14\x54\xe0\x63\xea\x0a\x66\x48\x47\xf8\x06\x91\x81\x65\x30\x1b\xca\x4e\x50\x88\x85\x15\x91\x42\x10\x7f\xfd\xe1\xd9\xf7\xce\xf6\x76\x94\x08\x1e\xf5\x92\x2c\xdf\x01\x85\x8e\xce\x62\x4b\x40\xd2\xb9\xc1\xe3\x43\xdf\xd5\x87\x46\x54\xae\x00\x94\xcf\x2d\xca\x03\x98\xd2\x77\x7c\x06\xd6\x66\xea\x73\x2e\x2e\x37\x94\xcf\x99\xe9\x0a\xe5\x49\x1e\x2c\x00\x0a\x8f\x41\xe1\x65\x41\xbd\xf0\xc1\x37\xf7\x04\x15\x30\x9d\x21\x2d\x7d\xca\xc8\x73\x24\x38\x8e\x0b\xae\x81\x31\xa4\xe2\xd2\x55\x8c\x15\xfb\x6a\xb5\x07\xfa\x0f\x07\xc9\x9b\xda\x73\x23\x4b\x78\xee\xa8\x28\xe6\xb6\xc5\x39\x33\x6c\xc7\xb1\x0c\xc2\x02\x61\x80\x2a\x84\xe1\x3a\xc4\xf6\x85\x08\x4c\xd3\x65\x0f\x82\x42\x40\x91\xa3\x07\x92\x17\x3d\x35\x7e\x47\x2e\xcb\x87\x41\x80\x6e\xc3\x28\x42\x3e\x78\x4f\x74\xcb\xef\x32\x04\xfe\xa4\xbd\xe4\x67\x3e\x92\x6c\x4a\x5b\x23\xdd\x94\x11\xab\xfa\x8b\x84\xfa\xdd\xc0\xb5\xb5\xbd\xb3\x93\x2f\xe5\x7b\xcc\x87\xfb\xa8\xfd\x7f\x87\xe7\x17\xe5\x9f\x5f\x2b\xdf\x88\xfd\xf7\xec\x26\xa3\xf7\x50\x9f\x6e\x53\x0f\x4b\x4f\xb7\x1c\x3b\xdc\xd3\xad\x26\xcf\xbf\x3d\xd3\xb6\xb4\xf6\xa7\x5b\x95\xb6\xf3\x74\xa3\xea\xd1\x59\xb3\x66\xe3\x3f\x69\xd2\xb0\x75\x73\x92\xe6\x2d\xa4\x4f\xf7\xaa\xbd\x6c\xd9\xa0\xe1\x73\x1a\x1d\x6f\x1c\x54\xcd\xfe\x3f\x00\x00\xff\xff\x52\xec\x3d\x41\xe4\x52\x00\x00")

func dbMigrations0001_initialSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbMigrations0001_initialSql,
		"db/migrations/0001_initial.sql",
	)
}

func dbMigrations0001_initialSql() (*asset, error) {
	bytes, err := dbMigrations0001_initialSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/migrations/0001_initial.sql", size: 21220, mode: os.FileMode(420), modTime: time.Unix(1461196056, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dbMigrations0002_use_tz_in_timestampsSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x95\xdf\x4e\xf3\x20\x18\x87\xcf\x7b\x15\xef\xd9\x0e\xbe\xec\x0a\xbe\x78\x50\x37\xa2\x4b\xf6\x2f\x96\xc6\xe8\x09\x21\x0c\x37\xb2\x0e\x08\xd0\x19\xef\x5e\xdc\x34\x69\x8c\x43\x4a\x15\x8f\x96\xb5\xf0\xfc\xa0\x2f\x3c\xef\x78\x0c\xff\x0e\x62\x6b\xa8\xe3\x50\xeb\xa2\x28\xe7\x18\xdd\xc1\xb4\xc4\xe5\x75\x59\x21\x60\xca\x70\xa3\x9a\x86\x1b\xa8\x10\x06\x3c\x5b\xa0\xc7\xd5\x12\xc1\x15\x8c\x6a\x3c\x19\xfd\x7f\x1f\xef\x47\xcf\x11\x38\x4e\x0f\x70\x7e\x30\x59\xcd\xeb\xc5\x12\x98\xe1\x1e\xbc\x21\xce\x02\x7e\x58\xa3\xd3\xfc\x0a\x97\x8b\x35\xdc\xcf\xf0\xed\xe9\x2f\x9c\x78\x75\x35\x5b\xde\x74\x87\x97\xb8\xf3\xf6\x8b\xac\xd6\x72\x63\x73\x85\x51\xad\x1b\xc1\xa8\x13\x4a\xe6\x8a\xd4\x94\xed\xe9\x96\xe7\x8a\x7b\xab\xb3\xb2\x84\xb2\x9c\x7b\x64\x3b\x2a\x25\x6f\x72\xc5\x6d\x8d\x6a\x75\xb6\x33\x23\xa4\x75\x54\xb2\x6c\x15\xfc\xc8\x23\x7f\x70\x58\xbf\xcf\x6e\xa8\x75\x84\xed\x38\xdb\x93\x27\x65\x48\xab\x37\x9e\x1d\xb7\x8e\x0b\x53\x7f\x68\x4d\x67\x1c\xf1\x02\x94\x3d\xbe\xcd\x85\xa9\xd1\x6b\xf2\x3f\xae\xb5\x64\x27\xac\x53\xe6\x25\x57\x99\xf8\x91\x4b\x97\xcd\x99\x5e\x25\x47\xe1\x7e\x75\x73\x45\xb7\x7b\x4d\xd5\xb3\x2c\x06\xf5\xa3\x55\xdd\xc9\x18\xd8\x6e\x02\xac\xc4\x0b\x1a\x20\x26\x34\x8b\x00\x2d\xb9\x17\x84\x98\xfd\x55\x1f\xa0\xf5\x37\x79\x00\x96\x22\xea\x08\x5c\xaa\x87\x07\xa1\xe3\x35\x3b\x3c\x26\xca\x9c\x31\x31\xe9\x32\x0c\xd0\x7b\xbb\x2e\x74\x61\x13\x54\xf6\x09\xf7\x1a\x00\x00\xff\xff\x56\x46\xd6\x50\x67\x0b\x00\x00")

func dbMigrations0002_use_tz_in_timestampsSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbMigrations0002_use_tz_in_timestampsSql,
		"db/migrations/0002_use_tz_in_timestamps.sql",
	)
}

func dbMigrations0002_use_tz_in_timestampsSql() (*asset, error) {
	bytes, err := dbMigrations0002_use_tz_in_timestampsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/migrations/0002_use_tz_in_timestamps.sql", size: 2919, mode: os.FileMode(420), modTime: time.Unix(1461196056, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dbMigrations0003_package_channels_blacklistSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\xcf\xcf\x4e\x84\x30\x10\xc7\xf1\xf3\xce\x53\xfc\x8e\x10\x77\x9f\x60\x4f\xb5\x9d\x4d\x1a\x6b\xd9\xf4\x4f\xe2\x9e\x36\x15\x08\x36\x22\x12\x25\xf1\xf5\xe5\x00\xc1\x8b\x5c\xdb\xc9\x67\xe6\x7b\x3a\xe1\xe1\x23\x77\x5f\x69\x6a\x11\x47\x22\xe9\x58\x04\x46\x10\x8f\x86\x31\xa6\xfa\x3d\x75\xed\xbd\x7e\x4b\xc3\xd0\xf6\xf7\xd7\x7e\x7e\xe8\xf3\xf7\x84\x82\x0e\xeb\x67\x6e\x10\xa3\x56\xb0\x55\x80\x8d\xc6\xc0\xf1\x85\x1d\x5b\xc9\x7e\x05\x50\xe4\xa6\x44\x65\xa1\xd8\xf0\xac\x4b\xe1\xa5\x50\x7c\xa4\xc3\x2a\xef\x21\xcb\xcc\xff\xc8\xd5\xe9\x67\xe1\x6e\x78\xe2\x1b\x8a\xed\xac\x23\x36\xbd\xa4\xf2\x4c\xf4\x37\x56\x7d\xfe\x0c\x44\xca\x55\xd7\x25\x56\x5f\xc0\x2f\xda\x07\xbf\x93\xbd\xec\x3c\xd3\x6f\x00\x00\x00\xff\xff\x24\x79\x55\x5d\x37\x01\x00\x00")

func dbMigrations0003_package_channels_blacklistSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbMigrations0003_package_channels_blacklistSql,
		"db/migrations/0003_package_channels_blacklist.sql",
	)
}

func dbMigrations0003_package_channels_blacklistSql() (*asset, error) {
	bytes, err := dbMigrations0003_package_channels_blacklistSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/migrations/0003_package_channels_blacklist.sql", size: 311, mode: os.FileMode(420), modTime: time.Unix(1461197415, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"db/drop_all_tables.sql": dbDrop_all_tablesSql,
	"db/migrations/0001_initial.sql": dbMigrations0001_initialSql,
	"db/migrations/0002_use_tz_in_timestamps.sql": dbMigrations0002_use_tz_in_timestampsSql,
	"db/migrations/0003_package_channels_blacklist.sql": dbMigrations0003_package_channels_blacklistSql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"db": &bintree{nil, map[string]*bintree{
		"drop_all_tables.sql": &bintree{dbDrop_all_tablesSql, map[string]*bintree{}},
		"migrations": &bintree{nil, map[string]*bintree{
			"0001_initial.sql": &bintree{dbMigrations0001_initialSql, map[string]*bintree{}},
			"0002_use_tz_in_timestamps.sql": &bintree{dbMigrations0002_use_tz_in_timestampsSql, map[string]*bintree{}},
			"0003_package_channels_blacklist.sql": &bintree{dbMigrations0003_package_channels_blacklistSql, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

