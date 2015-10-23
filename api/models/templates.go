// Code generated by go-bindata.
// sources:
// models/templates/app.tmpl
// models/templates/service/mysql.tmpl
// models/templates/service/papertrail.tmpl
// models/templates/service/postgres.tmpl
// models/templates/service/redis.tmpl
// models/templates/service/webhook.tmpl
// DO NOT EDIT!

package models

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _templatesAppTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe4\x5a\x6d\x6f\xdb\x38\x12\xfe\xde\x5f\x41\x08\xfd\x14\x38\xce\x4b\xb1\x7b\x77\xc2\xdd\x01\x89\x93\x6c\x73\x9b\x76\x0d\xdb\xed\x02\x57\x14\x07\x45\xa6\x63\x9d\x65\x52\x47\x52\xd9\xa6\x81\xff\xfb\x0d\x5f\x24\x91\x14\x25\x39\x69\x5a\x1c\x70\xe8\xbe\xd4\xe2\x68\x38\x1c\x3e\x9c\x79\x86\xa3\xc7\x47\xb4\xc4\xab\x8c\x60\x14\x25\x45\x11\xa1\xdd\xee\x15\x42\x8f\xf0\x2f\x42\xd1\xd9\xef\xf3\x05\xde\x16\x79\x22\xf0\x15\x65\xdb\x44\x7c\xc4\x8c\x67\x94\x44\x28\x46\xd1\xe9\xf1\xc9\xf1\xe1\xf1\x5f\xe0\x9f\x68\xa4\xc5\x27\x94\x2c\x33\x01\xe3\x3c\x8a\x8d\x0a\x50\xf5\x88\x84\xd1\x81\xa2\xdb\x24\x4f\x48\x8a\xd9\x61\xda\x88\xa2\xb1\x9e\xb3\x25\x5c\x30\x9a\x62\xce\xbb\x64\xa3\x73\xd0\xb5\x99\xe4\x25\x17\x98\xc9\x09\x51\x74\x45\xe2\xf8\xf2\x3f\x65\x92\x4b\x03\x3e\xc9\x27\x33\xbc\x82\xbf\x46\x95\x14\xda\x8d\x50\x14\xa1\xcf\x48\x2b\xd9\x19\xc3\xa7\x09\x4b\xb6\x18\x04\xb8\x5c\x59\xbf\xe5\x85\x94\xdd\xc3\x6a\x47\xae\x32\xd9\xb2\xd6\x3c\x82\x87\x8b\x87\x02\x2b\x8f\xce\x05\xcb\xc8\x9d\xf1\xa6\x1a\xba\xc0\xab\xa4\xcc\x85\x1a\x75\x9f\xf3\x94\x65\x85\xa8\xf6\x22\x32\x43\xbb\x51\x3d\x53\x51\x06\x66\x01\xd1\xf7\xe5\xf6\x16\x2c\x08\x4c\xa2\xf6\xf4\xb8\x6b\x1a\xe9\xc5\xe9\x07\xc4\xd7\x09\xc3\x1c\xd1\x15\xc2\x49\xba\x46\x66\xb5\xed\xf9\x2f\xc9\x7d\xc6\x28\xd9\x62\x22\xc2\x76\x74\x2f\xb6\x67\xad\xc1\xa5\xfe\x8a\x1f\xbe\xf7\x14\x33\x9c\xe3\x84\xe3\x1f\xb0\x6f\x33\x5c\x50\x9e\x09\xca\x42\x6b\xfa\xb6\xc9\xe6\xb4\x64\x29\x46\x29\x5d\x62\xc4\x9a\x69\x5a\x26\xcc\xcb\x5b\x82\x05\xef\x98\xff\x26\xe3\xe2\xaf\x10\x18\xe0\xa4\x4d\x4e\xe3\x58\x0b\xc7\xf1\xf5\xf2\xef\xcf\xb1\xe9\xe3\x74\x82\xb8\x9e\x0f\xad\x28\x43\x62\x9d\x71\x24\xe3\x50\xcb\xaa\x2a\xf4\x38\x56\x59\xfb\x29\x4f\x1f\x17\x7d\xe8\xa5\xe4\x9e\x7e\x81\x95\xab\xad\x44\xf7\x46\xdf\xa8\x13\x37\x6d\x13\xa6\x93\x0e\xa7\x34\xfe\x00\x19\xe9\x8c\xe7\xfa\x22\xe8\x03\x27\x56\xcd\x30\x57\xfb\x68\xef\x4f\x34\x81\xb0\x42\xb7\x0b\x5a\x64\xe9\x8c\xe6\x21\x9c\x56\x46\x5e\x9f\xbd\x8b\x63\x25\x63\x59\x32\x65\xb4\xc0\x4c\x64\xd8\xdd\x74\x99\x01\x38\x2f\xb7\x58\xca\x4f\x69\x9e\xa5\x0f\x17\x34\x2d\x5b\x67\xda\xdb\x1f\x99\x19\x4e\x0f\x21\x39\x9c\xfc\xc9\x9a\x44\x43\x4b\xc0\x2e\x99\xf7\x3f\x39\x43\xc8\xd3\xa7\xc4\x2f\x57\x2b\x9c\xaa\xdd\x3d\xcb\x73\xfa\x87\xa7\xcd\x98\x9e\x91\x34\x2b\x92\x5c\x67\x80\x39\x66\xf7\x59\x8a\x55\xf8\x07\x48\x6c\x6f\x97\xc9\x38\xd9\x26\x5f\x29\x49\xfe\xe0\xe3\x94\x6e\x55\xf0\x0f\xe8\x39\x4b\x0d\x4e\xe0\x3d\x2e\x78\xdc\x2c\x1c\xde\xf0\xc4\x77\xce\x6f\x7b\xd4\xd1\x0c\x69\x45\xac\xa5\xf1\x47\x91\xfb\x58\x7a\x52\xfb\xda\xf5\x81\xef\x01\x2d\xf9\xf0\x1e\x72\x93\xf2\xc1\x72\x9b\x11\x38\x7d\x2c\x81\x73\xdb\xf2\x45\x34\xb0\x41\x4a\x66\x9f\x4d\x52\x82\xce\x46\x49\xc7\xb6\xb6\xc2\x72\x59\x74\x20\x7f\x56\xc0\xd4\x0f\xd0\x6e\xc0\x6d\xf6\xaf\x46\x72\xd7\x4e\x64\x0d\xb4\x7b\x60\x7d\xa3\xb6\x3a\x8e\xaf\x4a\xa2\xad\xda\x0b\xdd\x13\x08\x85\x6d\x24\xcf\xdf\x9c\x97\xe9\x06\x8b\x86\x53\xfc\x83\x66\x06\x1a\x87\xb0\x52\xf8\x5f\xaa\x62\x09\xfc\xbd\xa1\x18\xca\x8c\x19\xbe\x53\xa7\x19\x16\xdf\xc6\x19\x28\x36\xa9\xca\xd7\xaa\x95\x9a\xc8\x74\xe4\xa8\xad\x19\x97\x64\x2e\x47\x2b\xc5\xc2\xe0\xf7\xf8\x6b\x56\x44\x7a\x92\x4e\xf8\xbd\x4d\xc8\x32\x57\x64\xa3\x3a\x09\xf8\x0b\x90\x0f\x02\x67\xc5\x91\x7b\x87\xb7\x90\x07\xe6\xd9\x57\xe5\xce\x93\xd3\x3f\xbb\xc3\x55\x40\xd1\x46\xff\x82\xc5\x99\xd0\xa8\x68\x45\x1d\x89\x09\x46\x5a\x27\x2c\x9a\x95\x44\x64\x1a\xc3\x04\x3c\xfe\x6f\xee\x4e\xb0\x80\x31\x5a\x2a\x6c\xbd\x39\x8e\xba\xa1\x10\xe6\x63\xac\x8e\x87\x83\x94\xec\x09\xa2\x5c\xc7\x11\x8f\xbf\x39\xa2\x1c\xa7\x25\xcb\xc4\x43\xd4\xa1\x8a\xcb\x23\x54\x0f\x56\xe1\xfb\xb7\x52\x14\xa5\x18\x26\xc8\xd4\xc8\x0d\x5a\xea\x0a\xd6\x39\x1c\x0b\x01\x39\xcc\x4b\xe2\x1f\x93\xbc\x34\x7b\x69\xe0\x55\xcb\x35\xee\x7e\x55\xfd\x77\xf7\x0a\x26\xc4\x64\xa9\xf4\x5a\x65\x42\x88\x97\xeb\xaa\xe1\x11\xb1\x84\xdc\x61\xf4\x7a\x83\xe2\xbf\xa1\xf1\x25\x11\x4c\x45\x2f\x5e\xad\x41\x73\x76\x90\x2b\x0b\x38\x92\x52\x6e\xb7\x6b\x42\x76\x1f\x83\x0f\xbf\xd3\xf0\xf9\x91\x9e\xdf\x98\xdb\x6f\x78\x45\xcd\x3d\xa3\xb1\x32\xba\x36\xb5\x99\x11\x8f\xe5\x22\x60\x60\x42\xb7\x5b\x38\x54\x76\xe6\xed\xe2\x65\x4e\xd6\x07\x55\xa9\x7e\x55\x2a\x33\x5a\x40\x9f\x2d\x1d\x22\x86\x15\x68\x02\xa6\xc0\x0b\x19\xc3\xcb\x09\x2d\x9d\x58\xdf\xd8\xe3\xf1\x7c\xc7\x9e\x93\xee\x89\x17\x6b\x8c\x88\x7a\x55\x72\xfc\x8c\x00\x8a\xe1\xf0\xaa\x98\xa3\x58\xbf\x80\x71\xe3\x47\x24\x28\x02\x30\x02\xa5\x94\xab\xd9\x60\x5c\x20\x56\x12\x02\x5e\x40\x94\xa0\x07\x38\x6a\x28\x35\xf5\xce\xd0\x6a\xae\xb7\xc9\x1d\x7e\xb2\x5b\xbf\xc1\x7d\x3a\xe6\xb5\x66\xec\xf1\x9b\x4c\x99\x3f\xfd\x1c\x9e\x12\xc6\xde\x9d\x4b\xef\xcc\xce\xde\x49\xaf\x40\xa0\x01\x80\xe2\x41\x2b\x2c\xe8\xbf\xfc\xc2\xf7\x3d\x0d\x75\x00\x51\xc0\x8f\x7e\x85\x41\x9e\x35\x81\x23\x10\x34\x2a\x11\x1d\x25\x46\x03\xfa\xad\xa8\xdb\x31\x83\x93\xc9\xcd\x28\x54\x16\x82\xe1\x64\x5b\xdd\x2d\x04\x93\x78\x34\x87\x6a\xb4\x3e\x03\x27\x4d\xe4\x32\xeb\xcf\x56\x68\xfc\x36\xe1\x53\x6d\x89\x15\x86\x6e\xe8\x1d\xff\xc0\x9d\x42\x3c\x40\x93\x95\x44\xed\xea\x0e\x1a\xd1\xb0\x3c\x4d\x0a\x8e\x1c\xea\x11\xa6\x7a\x1e\xe1\x70\x69\x9e\xb4\xcd\x23\xe8\x96\x54\x0f\xbf\xdb\x93\xdd\xf5\x92\xf0\x10\x0d\xdf\x8b\x88\x7b\x14\x7a\x63\x76\x71\x5a\x8a\x19\x4e\x29\x5b\xc2\xee\x7f\x0e\xbe\x65\xd1\xc6\x4f\x5d\xcc\x28\x61\x24\x06\xfe\x1e\x57\x5a\x0f\xe0\x0f\x57\xf0\x38\x6a\x73\x30\x58\x5e\xba\x51\xce\x54\xd9\xe1\xf0\xc0\xf0\xa4\x16\x1f\xf5\x19\x29\xf2\x24\x1c\x66\xf5\xca\x97\x71\xd3\xba\xdc\xb4\xb3\x54\x5d\x85\xf4\x42\x4a\xcb\x48\x1a\x38\x88\x2b\x08\x0e\x99\x2a\x6b\xac\x70\x6d\x76\xaf\xe4\x4a\x2b\xf8\xfb\xde\xa9\xe3\x24\x60\x0d\x8c\x1a\xa7\xd4\x58\xf7\x53\xfc\x53\x82\x44\xc3\x87\xaa\xa4\xf9\x1a\x72\x58\xb6\x82\x92\xdb\xca\x9a\x3d\x47\xae\x2b\xcd\x86\x03\xe3\xe5\x64\x91\xf0\xcd\x85\x34\x22\x13\x81\xaa\xbf\x00\x53\xf9\x6f\x0a\x23\x4e\x81\x30\xaa\x2b\x40\x75\x82\x3e\x07\x8a\x7b\x2d\x2e\xab\xf5\xb9\x37\x87\x25\x6c\x9d\xa4\x93\xf1\xf1\x7e\xc5\x84\x99\x78\x41\x37\x98\x0c\xf2\xe5\x4e\xae\xdc\x6c\x5f\xa8\xf0\xe8\x87\xba\xdc\x95\xda\x85\x51\xbb\x04\xb1\xaf\xb1\x6a\x45\xd5\x33\x4f\xd4\xbb\xbf\xab\xc5\xed\xe7\xde\x2b\x75\x71\x53\xa5\x09\xfc\xe0\x8b\x48\x8f\x1b\x66\x27\xe1\x06\xf4\xf0\x5f\x02\x1e\x01\x1c\x1a\xc3\x7b\xab\xc1\x20\x56\xda\x69\xd4\x03\x89\xeb\xf9\x5e\x4c\x54\xba\xfe\x27\xc0\xe0\x5c\x6e\xb7\xaf\xb2\x6d\x51\x9f\x0f\x06\x88\x73\x98\x3b\xfe\x48\xfc\xb5\x0f\x75\x9f\x99\xad\x20\xe0\xa3\xb9\x2e\x4c\xeb\x62\xa6\x39\xfa\x9e\xec\x0d\x4d\x96\xe7\xa6\xb6\x0a\x5c\xbc\x34\xc1\x09\xd8\x39\x11\x09\xc4\x3e\x36\xa5\x4c\xf0\x26\x4c\xd5\xb2\xbe\x67\x62\xcf\x33\x76\x6c\x1c\x57\x73\x56\x59\x4e\x2d\xcd\x76\x52\xdb\x6f\xfa\xc1\x38\xec\x42\x3b\x62\x7b\x56\xd9\x1b\xf3\x9e\x6a\xd6\xb6\xdf\x2d\x8b\xad\x72\x28\x21\x04\x3b\x38\x6e\x21\x55\x2f\xda\x0e\xfe\xaf\xab\x17\x9d\xe8\xdf\xbc\x73\x69\xae\x26\x3c\xaf\xfb\x05\x63\xa5\x65\x6c\x12\x8c\x71\x9a\x7c\xcb\x38\x6d\x22\x0f\xe4\x2a\x4b\x65\xf1\xbd\x6f\x49\xf9\x34\xad\x6e\xd1\xf9\x4c\xf7\xed\x55\x83\x36\xee\xb9\x26\x61\xf7\x04\x8e\x8d\x65\xf5\x5b\xca\x43\xdd\x97\xe1\x36\x82\x4c\xe4\x64\x89\xbf\x48\xb5\x33\xa8\xee\xe8\x96\x6b\x44\xee\xdf\xcb\x08\x7b\xe5\x7b\xc1\xe4\x29\x7b\x59\xcd\xfb\x5c\xcf\x3c\xd1\x11\x4f\xb4\xce\xc3\xef\x93\x0d\xfc\x7e\x96\xbd\x00\x9a\xea\x49\x5e\x14\x54\x03\x47\xcd\x2d\x70\x9f\x83\xbf\xa8\x12\xad\xdd\xe5\x79\xc9\x73\x8d\x5d\x33\xfb\x79\x5f\x3a\xd3\xc9\x06\x3a\xe4\x5f\xbc\x9f\xeb\x44\xea\xf5\xa6\x7f\x04\xf4\xdb\x15\xfe\xb3\x74\xbe\x48\x5a\xf1\x6f\x0b\x9e\xb5\x5d\x01\x1f\xcf\xcd\x6d\xed\x2f\x8c\x96\x45\x67\x8d\xa6\x5b\x9a\x8e\xe8\x60\x9d\xa6\xc4\xdc\xdb\xa0\x16\x47\x40\x7b\xb0\xa7\xc3\xdb\xda\x91\x5e\xda\x8f\x1c\x8b\xae\xc9\x1d\xd3\x75\xa6\x73\xa5\x30\x84\x13\x23\xe6\x5f\x06\x4c\xb2\x25\xbb\x96\x0e\x89\x8e\xc7\xea\xcf\xd1\x71\xfb\xc6\xe0\xba\x80\xc5\x0b\x9a\x52\x55\x8e\x8a\xb4\x68\x8b\x5c\x31\xba\x95\x33\xbe\x04\x8c\x5a\xca\x17\xf4\xa5\x54\xbb\x05\xfe\xc8\x75\x60\x8b\x57\x0d\xb3\x2a\xbb\xc2\xfc\x58\xa4\xd7\x4b\xc7\x48\xd9\x30\x0e\x55\xdd\x61\x88\x76\xa3\x32\x4f\xb8\xc8\xd2\x86\xc3\x42\xac\x8d\x63\x9b\xd2\xee\x71\x9b\xd0\x34\xf5\x1b\xbe\x6c\x9e\x39\x58\x03\xfe\x4b\xb0\xba\xc9\xb9\x60\xc0\x83\x61\x2a\x7d\xf1\xa4\xdf\xbc\x24\xc9\x6d\x8e\xe5\x2a\x05\x2b\xf1\xc8\x6e\xde\xfc\x7c\xdc\xa1\xc7\xee\x45\x00\x96\x96\x39\xee\x7e\x89\x51\xce\xff\x49\x09\xae\x26\x68\x86\xde\xe2\x24\x17\xeb\xc9\x1a\xa7\x1b\xbf\xd8\xd2\x43\x0f\x8b\x35\x9c\x8b\x35\xcd\x97\xea\x02\xcc\xed\x31\x29\xfe\x74\xaf\xae\x53\x7e\xf2\x4a\x12\x76\x17\x6e\xf5\x69\x5e\x1f\x2d\x26\x53\xe7\x8e\xa9\x2b\x05\x54\x78\xbb\xca\x18\x17\xf2\x47\x95\x16\x42\xcd\x40\xcb\x6d\x6f\x9c\xe7\x1f\xc8\x3a\xb8\x98\x57\x01\xcc\xaa\x4f\x31\x30\x69\x17\x34\xfb\xc6\x02\xbd\xe2\xeb\x95\x5e\xef\x37\x31\x6d\xff\xc8\x06\x7a\xcf\x76\x08\x51\x4e\x6d\x8b\xd8\x90\xfe\x5e\xc1\x44\x81\x41\x76\x35\x52\xbc\x87\x45\xb5\xe8\xb7\x58\x63\x60\xe0\x69\x6f\x59\x36\xe0\xb3\xf9\xfc\xe6\xff\xcf\x67\xa1\x4f\x13\xe6\x37\x16\xf2\xbc\x88\xfb\xfc\x9a\xce\xdf\x1d\xe7\xf7\xe7\x97\xce\x15\x37\xe7\x13\x4a\x37\x19\x9e\x43\x5c\x57\xb7\xda\xbc\x8e\xb2\x9f\x1e\xfd\x8e\x40\xb2\x52\x57\x20\xf2\x36\xcd\xd1\x61\xed\x7c\x75\x6d\x03\x2b\xf3\x1f\x83\x99\x5d\x54\x22\x50\x11\xf7\xf3\x26\xeb\x16\xfd\x09\x97\xc8\x5e\xa7\xdc\xb9\xa7\x09\x37\x81\xfc\xaf\x94\x3a\xda\x3f\x7b\x7d\x9d\xd4\xd9\xef\xf0\xba\x30\x4d\xfb\xc2\xa7\x1f\xde\xf7\x40\x2e\x32\xbc\x2e\x4d\x7f\xa7\xc4\xfd\x5c\xc9\x9f\xc7\xfa\x78\xa9\x0d\x7b\x9c\x72\xef\x4b\x26\x4f\x66\xaf\xe6\x45\xf8\x5a\xd3\xea\x16\xb5\x2a\xc4\xf6\x47\x4c\xe1\xae\x96\x93\x91\x5d\xfc\xda\xfb\xdd\xfe\x12\xaa\xe7\x9b\xb2\x97\xff\x5c\xac\x73\x93\xd5\x28\xd6\x44\x2b\x87\x13\x74\x5b\x13\x2d\x4d\xed\x6f\xf1\x41\xb8\xeb\xd5\xf1\x0e\xc3\x77\x32\x3b\xb3\x2a\x18\x72\x49\x90\x83\x94\x6d\x58\xdb\xcc\xd7\xf5\x7b\x26\xd6\x7b\xe8\x4a\x4f\x07\x8d\x07\x91\xb3\x52\xac\x29\xcb\xbe\xe2\x60\x95\xd1\x7a\x2b\xd0\xc5\x73\x7a\x78\xa1\x69\x0e\x02\x6a\xfa\xe3\x6d\x10\xbd\xd5\xdf\xf4\x68\x6f\xf3\xd9\xfe\xf8\xa6\xfd\x31\x8c\x1b\x6c\xe6\x6f\xe2\xd8\x7c\xf1\x65\xa2\xcd\x05\xce\xb1\xc4\x49\x1d\x93\x61\x85\xf2\x52\x78\x20\x1a\xa9\xde\x9e\xbc\x3f\x66\x3a\x3f\xc2\x61\xbf\xb7\xc9\x11\x90\xcd\x3b\xef\xd0\x54\x5d\x93\x88\x3f\xc0\x0e\x6f\xe5\x4d\x40\x55\x8d\x57\x1f\x99\xd9\xa9\xa7\x91\x97\xdf\x87\x8e\x42\xa5\x7b\xab\xb4\x0c\xb9\xcd\xf2\xda\x7f\x03\x00\x00\xff\xff\x67\x26\x37\xdd\x00\x30\x00\x00")

func templatesAppTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesAppTmpl,
		"templates/app.tmpl",
	)
}

func templatesAppTmpl() (*asset, error) {
	bytes, err := templatesAppTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/app.tmpl", size: 12288, mode: os.FileMode(420), modTime: time.Unix(1445555807, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templatesServiceMysqlTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x56\x5b\x6f\xda\x30\x14\x7e\xe7\x57\x58\x7e\xda\x24\xc6\x28\x95\x2a\x2d\x9a\x26\x51\xa0\x55\xa4\xb5\x43\x85\x75\x0f\x53\x1f\x8c\x7d\x40\xd6\x82\xed\xd9\x4e\x27\x56\xf1\xdf\x67\x87\x24\xd8\x84\xd0\x8b\xb4\x5e\x50\xc8\x39\xdf\xf9\xce\xdd\x7e\x7a\x42\x0c\x96\x5c\x00\xc2\x06\xf4\x23\xa7\x80\xd1\x76\xdb\x41\xe8\xc9\xfd\x23\x84\x87\x3f\x66\x73\x58\xab\x8c\x58\xb8\x92\x7a\x4d\xec\x3d\x68\xc3\xa5\xc0\x28\x41\x78\xd0\x3f\xeb\x7f\xe8\x7f\x72\x7f\xb8\xbb\x53\x9f\x12\x4d\xd6\x60\x9d\x0e\x4e\x4a\x13\xde\x48\x96\x49\xea\x2c\xb0\x99\x95\x9a\xac\x20\x90\x39\xe9\x7c\xa3\xa0\x30\x77\x9b\xaf\x17\xa0\x4b\x53\x85\x68\x0c\x4b\x92\x67\xb6\x90\x9e\xf5\x63\x89\xa1\x9a\x2b\x5b\xb9\x52\x53\x20\xb3\xe3\x40\x86\xff\x05\xf4\xee\xfa\xf2\x3d\x2e\x51\xdb\x0a\x8e\xc7\xc4\x92\x05\x31\x6d\x7e\xcc\xac\xe6\x62\xd5\xe6\x07\x51\xea\x94\x23\xa5\x2a\x62\x25\x07\x12\x2e\x21\x4d\x17\x52\x61\x2c\x11\x14\x0a\xd2\xb7\xb8\xc1\x16\x3d\x3b\xe8\xad\x39\xd5\xf2\x94\x3b\x15\x0f\xa2\x19\x31\x06\x2d\xa5\x0e\x3c\x93\x0c\x4c\xd3\xb5\xa9\x53\xfc\x23\x35\x7b\x85\x5b\x31\xe7\xcc\x35\x12\x68\xa4\x2a\x3b\x0d\x86\x59\xbe\x10\x60\xcd\x11\x02\x87\xfe\xca\x8d\xfd\xec\xda\x2e\x49\x26\xa3\x41\x92\xec\x74\x93\x24\x65\x5f\xda\x38\x1d\xe8\x7e\x3a\x42\xa6\xb4\xda\xa0\xfb\xee\x3a\xbb\xa8\xc2\x7f\x28\x77\x19\x6b\x5e\x51\x34\xc8\xef\x15\x3d\x1e\xe7\x3e\x44\xe7\xbc\x8f\xef\x74\x78\xb5\xe5\x4e\x60\x1f\x7f\xcb\xad\xca\xa3\x4c\xe2\xa9\xd4\xf6\xfc\xbc\x7f\x31\xa7\x6a\xc8\x98\xf6\x22\x67\x80\x64\x39\xec\x1e\xaf\x44\x92\x5c\x83\x1d\x5a\xeb\xbe\xff\xdc\x77\x08\xee\x22\x3c\x11\x4c\x49\x2e\x6c\xcf\x23\xc1\x18\x8c\x1e\xd0\x36\x6c\x8d\xbd\x6d\xff\xf8\x36\xdb\x05\xf2\xc0\xf0\x44\x3c\xde\x6c\xcc\xef\x2c\x9c\xcc\xc8\xf2\x1d\x2c\x7d\x22\x6a\xf9\x51\x74\xd8\xb9\xc7\xd0\xb5\xfc\x28\x3a\x6c\x93\x63\xe8\x5a\xee\xd1\x51\x15\xee\xc0\xc8\x5c\x53\x88\xea\x30\x03\x9a\x6b\x6e\x37\xd7\x5a\xe6\xea\xb9\x16\x88\x95\x83\x46\x98\x6a\xa9\x40\x5b\x0e\xf1\xb4\x38\x49\xa1\x7a\xd0\x27\x6b\x1f\x07\xaa\x16\x79\x37\x54\x8f\x18\x52\xb1\x2a\xca\xeb\x8a\x14\xe8\x20\x1f\x6c\xaa\x1c\xa5\x95\x54\x66\xde\xa0\xa5\xca\xd7\xee\x4a\xcb\x75\x59\x70\xec\xeb\xef\xdf\xcd\xe5\xe1\x9b\x11\x67\x3a\xf5\xa1\xba\x5d\xdd\x2b\x7e\x3f\x9e\x5d\xe0\x32\x57\xbb\x9f\x87\xc8\x27\x37\x1b\x29\x8b\x72\xec\xa7\x25\x00\x6c\x5b\x56\xc7\x73\x39\xbd\x1b\xbb\x8f\xf1\x65\xa8\xfc\xa2\x9c\x46\x90\x57\xe4\xb6\x00\xa5\xcc\x44\xb1\x54\x5b\xee\x64\x3c\xf5\x84\x3c\x1b\xcc\x7e\x96\x5e\x12\xc9\xb1\x23\xb7\xf6\xac\x21\xdc\x3b\x54\xa6\xa1\x22\x1b\xf9\x53\x23\xc2\x46\x07\x57\x2b\x2e\x65\x20\x2c\x5f\x72\xd0\x31\xb1\x8f\x67\x66\x09\xfd\x75\xbb\x1b\xa4\x03\xf8\x6d\x3d\x7e\xcd\x79\xef\xb6\x16\xaa\x81\x0a\xeb\x7e\x00\x9c\x88\x95\xbb\xec\xd4\xf5\x8c\xeb\x78\x43\x8c\xbb\xba\xc4\x7b\xa0\x39\xfc\x2d\x90\x78\xf9\x34\x37\x4e\x04\x8b\x47\x27\x92\xe4\x8b\x8c\xd3\x6c\x33\xa4\x6e\x9f\x18\xbe\xc8\x0a\x67\x97\x24\x33\x87\x4d\xb7\xab\x5d\xd5\x2a\x2b\x35\x88\xe5\xee\xcc\x88\x66\xbe\x98\xf6\x30\x49\xd1\xca\x71\x3b\xed\xa1\xd9\xa6\x9d\xea\x73\xdb\x71\x17\x45\x10\xcc\xdf\x0d\xff\x05\x00\x00\xff\xff\xc3\x37\xc1\xe6\x33\x0a\x00\x00")

func templatesServiceMysqlTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesServiceMysqlTmpl,
		"templates/service/mysql.tmpl",
	)
}

func templatesServiceMysqlTmpl() (*asset, error) {
	bytes, err := templatesServiceMysqlTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/service/mysql.tmpl", size: 2611, mode: os.FileMode(420), modTime: time.Unix(1444684514, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templatesServicePapertrailTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x56\x6d\x6b\x1b\x39\x10\xfe\xde\x5f\x21\x44\xa1\x50\x5c\x27\x4e\x28\xc7\x09\xee\x83\x2f\x4d\x5a\x5f\x9d\xc4\xac\xd3\x1e\x5c\x09\x87\xb2\x3b\x76\x74\x5e\x4b\x8b\xa4\x4d\xf3\x82\xff\xfb\x8d\xa4\x5d\xef\x6b\x62\x13\x38\x0e\xea\x86\x95\x1e\xcd\x8c\x9e\x19\x3d\x33\x4f\x4f\x24\x81\x85\x90\x40\xa8\x01\x7d\x27\x62\xa0\x64\xb3\x79\x43\xc8\x13\xfe\x08\xa1\xe3\x3f\xe7\x57\xb0\xce\x52\x6e\xe1\x4c\xe9\x35\xb7\xdf\x41\x1b\xa1\x24\x25\x8c\xd0\xa3\xc3\xd1\xe1\x87\xc3\x5f\xf1\x1f\x1d\x04\xf8\x8c\x6b\xbe\x06\x8b\x18\xca\x0a\x13\xb8\xfa\x4d\xa7\xb5\x4f\x5c\xb8\x7a\xc8\xc0\x5b\x98\x5b\x2d\xe4\xb2\x38\xed\xb7\x3e\x81\x89\xb5\xc8\x6c\xe9\x63\xc6\x33\xd0\x56\x73\x91\x92\x6f\xd1\x74\x40\x60\xb8\x1c\x92\x77\xa9\x5a\x9a\xd1\x30\xdb\xee\xf1\x2c\x1b\xc6\x6a\xcd\x46\xa3\xa3\xe3\x8f\xef\x68\x61\x6e\xe3\xff\x6e\x8a\xd8\x22\x30\x2a\xd7\x31\x34\x42\x9b\xf2\xf5\x4d\xc2\x4f\xef\x21\xce\x9d\xcb\x48\xa5\xd0\x13\x2a\xf3\x44\x30\x36\x19\x9f\x33\xe6\x31\xb5\x88\x67\x5a\xb9\x30\x44\xc3\x70\x20\xcf\x98\x7c\x0d\x0e\x3f\x53\xa9\x88\x1f\x3e\xa9\x18\xbf\xa5\x6d\xe1\x10\x59\xb2\x1a\x48\x3d\xfa\x80\xbc\x8e\x7e\xa9\x39\xf1\xa0\xb9\xc5\x2c\x14\xe7\x7f\x34\xb6\x48\xcb\x9e\x87\x9f\x2e\x16\x10\x5b\x1f\x7b\x9a\xaa\x9f\x2d\x6b\x45\xe8\x42\xc6\x22\xe3\x3e\x3d\xe8\xa0\xa8\x00\x34\x4f\x68\xea\x99\x19\xf2\x35\x7f\x54\x92\xff\x34\x8e\x5f\x4a\xae\x4b\x3a\x1b\x76\xc6\xb1\x0d\xd1\xe3\x39\x63\x0d\xab\x2e\x8e\x27\x5a\xf0\x4d\xe3\xbb\xbe\xdb\xb0\x8c\x89\xb7\xb7\x2e\xf8\x03\xda\x5c\x76\x4c\x06\xae\x9b\x1c\xb4\x19\x08\xc8\x87\x0b\xac\x47\x67\x26\x24\xfa\x24\x55\x79\x12\x0a\x19\x03\xfe\x8a\x75\x6f\x84\xe9\x30\x43\x77\xa4\xcb\x63\xf6\x49\x99\x07\xbe\x94\xb6\xbe\xc0\x8b\x63\xbb\xd3\xe7\x61\x15\xf5\xbd\xfb\xa4\x4c\x24\x9b\xc8\x3b\xb5\x82\xb3\x5c\x86\x03\xbd\xe8\xeb\x67\x9c\x94\x4f\xe7\x25\x37\xef\x9f\x31\xd9\xb3\xda\x53\x42\xff\x3d\x0d\xab\x90\x6c\xf6\x19\x6c\x04\xb1\xd2\x49\x37\xef\x7d\xd8\xf9\x2d\xd7\xc9\x04\x15\x8d\x5b\xa5\x77\x9f\x08\xf2\x75\x03\x28\x6c\xc0\xd7\xbb\xf1\x53\x61\x6c\xc0\xbe\x10\x8e\x53\x3b\x76\x82\x20\x0b\x53\xb5\xfc\xac\x55\x9e\xed\x0b\xde\x15\x87\x47\xcf\x72\x8b\xd0\xd3\x3b\xac\x51\xf3\xea\xc2\xe8\x2f\x80\xff\x25\xd5\xb1\x7b\xe5\x8b\xf2\x95\xd7\x72\xc2\xe3\xd5\xeb\x2f\x88\x02\x79\x26\x19\xfb\x43\x89\x42\xe9\xe8\xc0\xfd\xcf\xb5\x64\x28\x8f\xac\xe5\x14\x37\x9f\xdc\xf1\xc5\xb6\x79\x44\xb0\xf4\x4d\x6d\x33\x20\xb4\x67\x7b\x1c\xc7\x2a\x97\x76\x92\x14\x08\xe3\xa2\x3d\x70\x38\x0f\x23\x25\xce\xdf\xc2\xab\x9a\x83\x1d\xbc\x77\x9a\x7c\xdd\x52\xd5\x82\xfb\xce\xda\xcb\x5a\x5c\xff\xaa\x90\xe5\xea\x36\x93\x38\x32\x68\x2e\x97\x40\xde\xae\x06\xe4\xed\x1d\x61\xbf\x91\xe1\x38\xba\x30\x61\x6e\x28\x78\x43\x50\x9e\x61\x5b\x44\x10\xae\x9f\x63\x7f\x76\x7d\xbe\xd5\x1f\xb7\x63\x80\xbf\x58\x50\x68\xc6\x7c\x1d\xce\x3d\xeb\xe5\xb9\x66\x07\xa8\xfa\x2d\xe9\x34\xd2\x53\xc9\x6f\x52\x48\xdc\x8e\xd5\x39\xb4\x3a\x68\xcd\xf4\x58\x87\xf9\x02\x03\xc5\x3b\x6c\x36\xed\x66\x5b\x0a\xa5\x67\x9a\x6c\xb3\x8f\x8a\x30\xb6\xd6\x2d\xfc\xa8\x8d\x26\xf8\x7a\x96\x80\xea\x80\x37\x41\xb3\x9d\x1e\xe9\x5a\x00\x06\x2c\x97\x33\x65\xc4\x76\xb2\xb9\x8a\x26\xe7\x7f\x7f\xb9\x8c\x26\x7f\x5d\x5e\xd4\x8b\xb2\x62\xb1\x4e\x39\xc8\xa4\x22\xb8\xeb\xfa\xf9\x99\xa5\xa4\x75\xab\xfc\x7b\x8d\x2e\x5f\xb8\x4c\x52\x6f\x97\x0a\x99\xc0\xfd\xf0\xb6\x58\x68\xa4\xa2\x1c\x97\xba\xdc\xf4\xcd\x55\xfd\xf4\xd0\x13\x95\x40\x77\x22\x9a\x1f\xff\x9e\xc7\x2b\xf0\x72\x10\x2b\xec\x5d\xf7\x9d\x71\xe8\xf8\x2b\x3c\xb8\xed\xd0\xe0\x0e\xaa\x69\x70\xf8\x28\x32\xfa\xec\x60\x11\xe1\x2b\x13\x61\x28\x90\xe8\xfa\x9f\xa6\xf0\xd2\x2b\xdc\x53\xb9\xf7\x7b\xf4\x91\x76\x9f\x41\x63\xa6\xbc\xcc\x6d\x96\xdb\x3a\x7b\xaf\x78\x1d\x53\x21\x57\x6d\xfe\xbf\xf3\x34\xf7\x21\x6e\xeb\x73\x8f\xaa\xe8\xcc\xd9\xa5\x95\x4a\x6a\x1c\xa4\x7d\x97\x37\xee\x57\x19\xfb\x37\x00\x00\xff\xff\x32\xc4\xd0\xe1\x13\x0c\x00\x00")

func templatesServicePapertrailTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesServicePapertrailTmpl,
		"templates/service/papertrail.tmpl",
	)
}

func templatesServicePapertrailTmpl() (*asset, error) {
	bytes, err := templatesServicePapertrailTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/service/papertrail.tmpl", size: 3091, mode: os.FileMode(420), modTime: time.Unix(1444684514, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templatesServicePostgresTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x56\x5b\x6f\xe2\x38\x14\x7e\xe7\x57\x58\x7e\xda\x95\x58\x96\xb2\x17\x69\xa3\xd5\x48\x14\x68\x15\x69\xa6\x83\x0a\xd3\x79\x18\xf5\xc1\xd8\x07\x64\x4d\xb0\x2d\xdb\xe9\xa8\x53\xf1\xdf\xe7\x38\x24\x21\x26\x04\x7a\x47\x21\xe7\x7c\xdf\xb9\x1f\xfb\xe5\x85\x08\x58\x4b\x05\x84\x3a\xb0\x4f\x92\x03\x25\xbb\x5d\x8f\x90\x17\xfc\x27\x84\x8e\xbf\x2e\x96\xb0\x35\x19\xf3\x70\xa3\xed\x96\xf9\x07\xb0\x4e\x6a\x45\x49\x42\xe8\x68\x78\x35\xfc\x63\xf8\x1f\xfe\xd1\xfe\x5e\x7d\xce\x2c\xdb\x82\x47\x1d\x9a\x94\x14\x81\x24\xcb\x34\x47\x06\xb1\xf0\xda\xb2\x0d\x34\x64\x28\x5d\x3e\x1b\x28\xe8\xee\xf2\xed\x0a\x6c\x49\x55\x88\xa6\xb0\x66\x79\xe6\x0b\xe9\xd5\x30\x96\x38\x6e\xa5\xf1\x95\x2b\xb5\x09\xe2\xf6\x36\x88\x93\x3f\x81\xfc\x76\x7b\xfd\x3b\x2d\x51\xbb\x0a\x4e\xa7\xcc\xb3\x15\x73\x5d\x7e\x2c\xbc\x95\x6a\xd3\xe5\x07\x33\xe6\x9c\x23\xa5\x2a\x11\xa5\x0d\xa2\x30\x21\x6d\x17\x52\xe5\x3c\x53\x1c\x0a\xa3\xef\x71\x43\xac\x06\x7e\x34\xd8\x4a\x6e\xf5\x39\x77\x2a\x3b\x84\x67\xcc\x39\xb2\xd6\xb6\xe1\x99\x16\xe0\xda\xae\xcd\x51\xf1\x87\xb6\xe2\x0d\x6e\xc5\x36\x17\xd8\x48\x60\x89\xa9\x78\x5a\x16\x16\xf9\x4a\x81\x77\x27\x0c\x20\xfa\xa3\x74\xfe\x7f\x6c\xbb\x24\x99\x4d\x46\x49\xb2\xd7\x4d\x92\x54\x7c\xe8\xb2\x89\xa0\x87\xf9\x84\xb8\x92\xb5\x65\xee\x0b\x76\x76\x51\x85\xf7\xe4\xd9\x68\xe7\x37\x16\xf3\x74\x39\xe0\xbc\xb2\xd3\xf2\xe0\xc1\xf0\xd3\xc1\x1e\xe2\xc4\x08\x42\x90\xe7\x63\xac\x99\x7b\x0d\x7e\xfa\x39\xf7\x26\x8f\xd2\x49\xe7\xda\xfa\x7f\xfe\xfe\x6b\xb4\xe4\x66\x2c\x84\x0d\x22\x24\x60\x59\x0e\xfb\xc7\x1b\x95\x24\xb7\xe0\xc7\xde\xe3\xf7\x6f\x87\x36\xa1\x7d\x42\x67\x4a\x18\x2d\x95\x1f\x04\x24\x38\x47\xc9\x23\xd9\x35\xfb\xe3\xc0\x1d\x1e\xdf\xc7\x5d\x20\x8f\x88\x67\xea\x69\x5e\x66\xbb\x39\xa1\x11\xf9\x3d\xac\x43\x2e\x6a\x79\x17\x41\xb3\x89\x4f\x11\xd4\xf2\x2e\x82\x66\xd3\x9c\x22\xa8\xe5\x81\x20\x2a\xc7\x3d\x38\x9d\x5b\x0e\x51\x41\x16\xc0\x73\x2b\xfd\xf3\xad\xd5\xb9\xb9\xd4\x0b\xb1\x72\xa3\x23\xe6\x56\x1b\xb0\x5e\x42\x3c\x3b\x28\x29\x54\x8f\x1a\xa6\x6a\x5d\x52\x6d\xf6\x7e\x13\x11\x19\x49\xd5\xa6\x28\x35\x16\xac\xa1\x43\x42\xbc\xa9\x41\xab\x5e\x73\x9d\x05\x4e\xcf\x4d\xa8\xe3\x8d\xd5\xdb\xb2\xf8\x34\xf4\x42\x78\xb7\xd4\xc7\x6f\x26\x52\xd8\x34\x44\x8b\xcb\x7b\x50\xfc\xfe\x79\xf5\x2f\x2d\xd3\xb5\xff\x79\x8c\x7c\xc2\x39\x49\x45\x94\xe6\x30\x39\x0d\xc0\xae\x63\x97\x5c\x4a\xeb\xfd\x14\x3f\xa6\xd7\x4d\xe5\x57\xa5\x35\x82\xbc\x2d\xbd\x05\x2e\x15\x2e\x0a\xa7\xda\x7c\x67\x43\xaa\x07\xe6\x62\x3c\x87\xd1\x7a\x4d\x30\xa7\x8e\xe1\xda\xb3\x96\xf0\xe0\x50\x99\x89\xca\xd8\x24\x9c\x24\x11\x36\x3a\xcc\x3a\x71\xa9\x00\xe5\xe5\x5a\x82\x8d\x0d\x87\x78\x16\x9e\xf1\xef\x77\xfb\x71\x3a\x82\xdf\xd5\x43\xd8\x9e\xfd\x7e\x67\xad\x5a\xa8\x66\xe9\x8f\x80\x33\xb5\xc1\x0b\x10\x3d\xbd\xec\x51\xfe\x89\x39\xbc\xd1\xc4\x0b\xa1\xbd\x05\x3a\x20\xf1\x22\x6a\x6f\x9f\x08\x16\x0f\x50\x24\xc9\x57\x99\xe4\xd9\xf3\x98\xe3\x62\x71\x72\x95\x15\xfe\xae\x59\xe6\x8e\xfb\x6e\x5f\xbe\xaa\x5b\x36\xe6\x88\x08\x4f\x91\x68\xf2\x8b\x99\x6f\xe6\x29\xda\x3d\xb8\xdc\x1e\xdb\x9d\xda\xab\x3e\x77\x3d\xbc\x3f\x82\x12\xe1\xca\xf8\x2b\x00\x00\xff\xff\xa1\x5a\xfe\x19\x4a\x0a\x00\x00")

func templatesServicePostgresTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesServicePostgresTmpl,
		"templates/service/postgres.tmpl",
	)
}

func templatesServicePostgresTmpl() (*asset, error) {
	bytes, err := templatesServicePostgresTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/service/postgres.tmpl", size: 2634, mode: os.FileMode(420), modTime: time.Unix(1444684514, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templatesServiceRedisTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x56\xd1\x6f\xda\x3e\x10\x7e\xe7\xaf\xb0\xfc\xf4\xfb\x49\x8c\x41\x27\x75\x6a\x34\x4d\x42\x0c\xaa\x48\x5b\x87\x80\x76\x0f\x55\x1f\x8c\x7d\x50\x6b\x89\x1d\xd9\x4e\xb7\xaa\xe2\x7f\xdf\xd9\x09\x10\x27\x1d\x6d\xa5\xb5\x80\x22\xdf\xdd\xf7\xdd\x7d\x77\xb6\xf3\xf4\x44\x04\x6c\xa4\x02\x42\x2d\x98\x07\xc9\x81\x92\xdd\xae\x47\xc8\x13\x7e\x09\xa1\xe3\x1f\xcb\x15\xe4\x45\xc6\x1c\xcc\xb4\xc9\x99\xbb\x01\x63\xa5\x56\x94\x24\x84\x9e\x0d\x47\xc3\x77\xc3\x0b\xfc\xd0\x7e\xe5\x3e\x67\x86\xe5\xe0\xd0\x87\x26\x35\x04\xae\x7e\x61\x8e\xad\x99\x85\xc6\x1a\xae\xae\x1e\x0b\x08\x30\x4b\x67\xa4\xda\xd6\x10\x55\x00\x6c\x58\x99\xb9\x60\x1d\xc6\x06\xcb\x8d\x2c\xdc\x3e\x83\xda\x91\x88\x9a\x81\x48\x25\xe0\x37\xad\x03\x76\xfb\x48\x9a\x2a\xeb\x98\xe2\x10\x38\xbb\x59\x9c\x4c\x02\x8d\x9c\xf1\x7b\x18\xb8\xb3\x41\x2e\xb9\xd1\x7f\x4b\x08\x1d\x57\xf7\x40\x1c\x22\x12\xbd\xc1\x54\x2a\x4e\xe2\x34\x29\xb1\xf8\x4e\x52\x73\x66\xed\x2f\x6d\xc4\x1b\x64\x89\xab\xbf\x56\x88\x2b\xc8\x7f\x48\xb0\x06\x62\x20\xd7\x0f\x20\xfe\xef\x12\x2d\xcb\xb5\x02\x67\x9f\x2f\xfc\xab\xb4\xee\x13\xb6\x39\x49\xa6\x93\xb3\x24\xa9\x7c\x93\x24\x15\x9f\x4f\xd4\x79\x33\x9f\x10\x5b\xa3\x76\xe8\x6e\x0a\xfe\x3c\xd5\x91\x05\xe3\x3d\xc5\x69\x86\x03\x72\xaf\x81\x4f\xbf\x97\xae\x28\xa3\x62\xe8\x5c\x1b\x77\xfe\xe1\xe3\xc5\x8a\x17\x63\x21\x8c\x37\x21\x00\xcb\x4a\xa8\x1e\x67\x2a\x49\x2e\xc1\x8d\x9d\x6f\xe6\x2d\xa1\x0b\x28\x32\xc9\x99\xa7\xba\x34\xba\x2c\x68\x1f\x31\x8c\xcc\x99\x79\x9c\x2a\x31\xd7\x52\xb9\x81\x07\x02\x6b\x29\xb9\x23\xbb\x66\xcf\x8e\x54\xfe\xf1\x9f\x50\x05\xa0\x16\xcf\x54\x3d\x2c\x40\x48\xdb\xdc\x3a\x11\xd1\x02\x36\x5e\xa6\x83\xdd\x47\x47\x3a\x2d\xc0\xea\xd2\x70\x88\x94\x5a\x02\x2f\x8d\x74\x8f\x55\x2e\x2f\x34\x29\x76\x6e\xb4\x6a\x6e\x74\x01\xc6\x49\x88\x47\x0a\x2d\xc1\xb5\xd5\x49\xe3\xeb\x20\xfb\xd3\xa5\xdf\x74\x8f\x18\x52\xb5\x0d\x8a\xa3\x6e\x0d\x1f\xe2\x8b\x4d\x0b\xa4\x74\x9a\xeb\xcc\x03\x3a\x1e\x64\x9c\x19\x9d\xd7\x3d\xa0\xbe\x25\x7e\x6d\xa5\xdb\x2b\x13\x29\x4c\xea\x4b\xa5\xa3\xe1\x20\xfc\xbf\x1f\x9d\xd3\x5a\xab\xea\xef\x2e\xca\x09\xa7\x37\x15\x91\xc6\x7e\x9e\x1b\x01\xbb\xce\xc0\x4f\xfc\xf1\x50\x6d\x9c\x17\x85\xcd\x98\x75\x32\x04\xec\xf7\xda\x1b\xe4\x6d\x29\xbb\xa8\x94\x0d\x28\x64\xdb\x82\x39\xec\xfb\x54\xd8\xa8\x9c\xfd\x69\x70\xb2\xa4\xce\xdc\xbe\xb2\xa4\xee\xbc\xbf\xa6\xae\x71\xe9\x34\xde\x2d\x92\xcf\x98\xcc\xf0\x10\x33\x53\xc5\xd6\x19\xf8\x36\x6c\x58\x66\xa1\xdf\x76\xfe\x26\x95\x36\xf5\x55\x74\x5d\x6c\x0d\x13\x3e\x1f\x67\xca\xd8\x35\x24\x75\xa5\xc5\xe1\xdc\x3f\x88\x10\x5d\x08\xc7\xba\x9f\x6d\xe7\x15\x5e\x69\x51\x70\xa7\xdf\x2d\x80\xa9\xda\xe2\x8d\x7a\x98\xfd\xb8\x29\x57\x65\x1e\xe2\x27\x59\x69\xeb\x8b\x92\x8e\x62\x97\x78\x86\x9b\x96\xb6\xbc\xf1\x40\x1c\x33\x0c\x8d\x59\x3a\xc6\x7f\x86\xe4\x5b\xf9\xc5\xdb\x2e\x8c\xc7\x6d\x73\x40\xa2\x7d\x8f\x07\xcb\x5d\x77\x50\x7a\xfb\xdf\x5d\x0f\x5f\x21\x40\x09\xff\xd6\xf0\x27\x00\x00\xff\xff\x64\xe5\x91\xd9\x4d\x08\x00\x00")

func templatesServiceRedisTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesServiceRedisTmpl,
		"templates/service/redis.tmpl",
	)
}

func templatesServiceRedisTmpl() (*asset, error) {
	bytes, err := templatesServiceRedisTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/service/redis.tmpl", size: 2125, mode: os.FileMode(420), modTime: time.Unix(1442329647, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _templatesServiceWebhookTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x92\x4d\x6f\xb2\x40\x10\xc7\xef\x7c\x8a\xcd\x5e\xbc\x28\xa0\xcf\x73\x29\x37\xd3\x97\x53\x63\x8d\xa0\x9e\x71\x19\x75\x23\x30\x9b\x65\x30\x6d\x88\xdf\xbd\xbb\x4b\x51\x48\xeb\xa5\x4d\x48\xc8\xce\xcb\x6f\x66\xfe\x33\x4d\xc3\x32\xd8\xcb\x12\x18\xaf\x40\x9f\xa5\x00\xce\x2e\x17\xaf\xf1\x18\xe3\xf3\x6d\x9c\x40\xa1\xf2\x94\xe0\x05\x75\x91\xd2\x06\x74\x25\xb1\xe4\x2c\x62\x7c\x16\x4e\xc3\x49\xf8\x60\x3e\x3e\xb6\xc1\xcb\x54\xa7\x05\x90\x89\xe0\x11\xb3\xe9\xc6\xb6\xd6\xf9\xf5\x61\x9e\xc9\x87\x02\x97\x1b\x93\x96\xe5\xc1\xe5\x39\xc7\x13\x54\x42\x4b\x45\x1d\x7b\x0b\xbb\x23\xe2\x89\xad\x57\xaf\x63\x06\xfe\xc1\x67\xa3\x23\x91\xaa\xa2\x20\x38\x68\x99\xf9\x02\xcb\x33\xbe\x9b\x5f\x11\xe8\x54\x9c\x26\x36\x38\x98\xce\xfe\xfd\x1f\x71\x87\xbc\xb4\x64\xfe\x58\x57\x84\x45\x82\x4a\x8a\x5f\xf5\x31\xa4\x2d\x90\xe4\x5e\x8a\xd4\xba\xff\xcc\xf4\xbe\xb8\x7c\x05\x15\xd6\x5a\x40\x4f\xb6\x7e\xa1\xea\xc7\x22\xed\x60\x51\x14\x2f\xe2\xb8\xde\xdd\x0a\x5c\xab\x76\x9b\x32\xb1\x53\x3f\xbc\xd9\x97\x1a\x15\x68\x92\xd0\xe7\x1a\x7b\xdc\xee\x3e\xc1\x13\xd8\xa4\xc6\xf6\xb5\xe7\xd1\x50\xc2\x4e\x89\xb6\x17\x6b\x9a\x6b\x37\xd4\x2d\xfc\xbb\x46\x83\x24\x53\x9e\x50\x60\xee\x86\xb0\x2b\xe5\x3d\xe7\x73\x99\x29\x94\x25\x0d\x89\xf6\x86\x9c\x5c\x9d\x6c\x7d\xf1\xde\x6a\x52\x35\xdd\xbf\xb8\x4d\x9a\xd7\xc0\xef\xe0\x1c\xc6\x33\xb7\xde\x30\x28\x33\x7b\xf5\x9f\x01\x00\x00\xff\xff\x54\xf9\xe1\x63\x0d\x03\x00\x00")

func templatesServiceWebhookTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesServiceWebhookTmpl,
		"templates/service/webhook.tmpl",
	)
}

func templatesServiceWebhookTmpl() (*asset, error) {
	bytes, err := templatesServiceWebhookTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/service/webhook.tmpl", size: 781, mode: os.FileMode(420), modTime: time.Unix(1445572155, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"templates/app.tmpl": templatesAppTmpl,
	"templates/service/mysql.tmpl": templatesServiceMysqlTmpl,
	"templates/service/papertrail.tmpl": templatesServicePapertrailTmpl,
	"templates/service/postgres.tmpl": templatesServicePostgresTmpl,
	"templates/service/redis.tmpl": templatesServiceRedisTmpl,
	"templates/service/webhook.tmpl": templatesServiceWebhookTmpl,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"app.tmpl": &bintree{templatesAppTmpl, map[string]*bintree{
		}},
		"service": &bintree{nil, map[string]*bintree{
			"mysql.tmpl": &bintree{templatesServiceMysqlTmpl, map[string]*bintree{
			}},
			"papertrail.tmpl": &bintree{templatesServicePapertrailTmpl, map[string]*bintree{
			}},
			"postgres.tmpl": &bintree{templatesServicePostgresTmpl, map[string]*bintree{
			}},
			"redis.tmpl": &bintree{templatesServiceRedisTmpl, map[string]*bintree{
			}},
			"webhook.tmpl": &bintree{templatesServiceWebhookTmpl, map[string]*bintree{
			}},
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
