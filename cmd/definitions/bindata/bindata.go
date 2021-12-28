// Code generated by go-bindata. DO NOT EDIT.
// sources:
// definitions/features.toml (1.578kB)
// definitions/fields.toml (714B)
// definitions/info_object_meta.toml (1.112kB)
// definitions/info_storage_meta.toml (1.139kB)
// definitions/operations.toml (10.788kB)
// definitions/pairs.toml (1.945kB)

// +build tools

package bindata

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _definitionsFeaturesToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x94\x41\x8b\x13\x41\x10\x85\xef\xf3\x2b\x1e\x7b\x51\x21\x09\x7a\x11\x15\x3c\x8b\x37\x41\x6f\x8b\x84\xca\x74\xc5\x29\xd3\xd3\x35\x54\xd7\x4c\xd8\x7f\x2f\x3d\x93\xd9\x0d\x21\x13\xc4\xdd\x83\xb7\x24\xdd\xf5\xea\xbd\xf7\x35\xb9\x8f\xaa\x99\xb7\x1d\x89\xfd\xac\x02\xe7\xda\xa4\x73\xd1\x84\xcf\xb8\xbb\xbb\xab\x9e\x4e\xb1\x67\xf2\xde\x18\x92\x11\x38\xcb\xaf\xc4\x01\x7b\x35\xf4\x99\x2d\xe3\xd8\x28\x82\xa6\x57\x8e\x23\x25\x47\x76\x93\xda\x31\x0e\xd6\x0d\xd7\x87\xbc\xa9\xaa\xaf\x7b\x78\x23\xf9\x5c\x89\x13\xed\x22\x87\x15\xbc\x61\x64\xb6\x41\x6a\xc6\x51\x62\x44\x52\x87\xb1\xf7\x96\x40\x09\x6c\xa6\x36\xae\x2b\xbf\xe7\xbe\xeb\xd4\x26\xf9\x22\xfc\xe3\x5c\xf5\x48\x19\x92\xdc\x34\xf4\x35\x07\x48\xc2\x97\xef\xdf\xd6\xef\xde\x7e\xdc\x54\x25\x51\x75\x3f\x88\x79\x4f\x71\x1b\xae\x47\x3e\x3b\x5e\xcc\x4c\x8f\x5e\xbd\x21\x47\x50\xce\x25\x7a\x43\x03\x23\x91\xcb\xc0\x28\xe3\xb3\xcf\x5d\x3f\xd5\x92\xe1\x8a\xce\x74\x90\xc0\xc8\xd2\xf6\x91\x9c\x03\xb4\x63\xa3\xe2\xa0\x64\x59\xe3\x4a\x4d\x41\xf2\xd8\x13\x5e\x97\x9e\x02\xef\xa9\x8f\x8e\x1d\x37\x34\x88\xda\x9b\x2b\xf5\x8d\x67\x8c\x28\x07\x86\x5c\x38\xa4\xf4\x70\x6e\x6f\x73\x7d\xe7\x32\x9a\x39\xd6\x53\x82\xa2\x36\xbb\x29\x8d\xd7\xc6\xe4\x5c\x2a\x5c\x9d\x3e\xaf\x10\x25\xfb\x0a\x81\x23\x97\x6f\x94\x02\xb2\x42\xd3\x3f\xe3\xd3\xdd\x6f\xae\x7d\xdb\xb2\x53\x20\xa7\x5b\x28\x2f\xae\x3e\x07\xeb\x24\x85\x47\xa9\xff\x18\xf1\x82\xd5\x17\xc0\x7d\xa9\x7c\x8e\xbe\x8c\x4f\xa8\xb3\x93\xbf\x04\xe8\x28\xe9\x70\x8b\x6e\x39\x7f\x0e\xd2\x39\x5f\xb9\x5e\xb4\x96\x08\x2d\x37\x64\x7d\x42\xad\x6d\x47\x2e\xbb\xc8\x68\x35\xf0\xa7\xd3\xbb\x1f\x15\x31\x08\xcd\xdb\x5a\xf6\x46\x43\x5e\x8d\x2f\x86\x62\xd4\x23\x8c\x29\x4c\xf7\xf6\xa6\x2d\x34\x86\x75\xf6\x87\x78\x9a\x9d\xca\x5e\xa0\x56\xfe\x0d\x6f\xfb\x92\x34\x6f\xa6\x0c\xf5\x86\x6d\xbe\xf3\x77\x40\x3e\xbc\x3f\xf1\xf8\x13\x00\x00\xff\xff\x02\x8a\xc0\x31\x2a\x06\x00\x00")

func definitionsFeaturesTomlBytes() ([]byte, error) {
	return bindataRead(
		_definitionsFeaturesToml,
		"definitions/features.toml",
	)
}

func definitionsFeaturesToml() (*asset, error) {
	bytes, err := definitionsFeaturesTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "definitions/features.toml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe8, 0xde, 0x64, 0x57, 0x6f, 0xe4, 0xa1, 0x28, 0x9a, 0xdf, 0xf5, 0xa, 0x78, 0xe3, 0xe7, 0xe5, 0x4c, 0x35, 0xce, 0xe9, 0xb8, 0x4b, 0x32, 0xb7, 0xb2, 0xdc, 0x74, 0xc5, 0x98, 0xd5, 0x9f, 0x28}}
	return a, nil
}

var _definitionsFieldsToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x52\xb1\x4e\xc3\x30\x14\xdc\xf3\x15\x55\xc7\x0e\x9e\x10\x1b\x0b\x62\x61\x40\x54\x65\x60\x88\x3a\xbc\x24\xd7\xf4\x41\x13\xbb\xcf\x2f\xa2\xf0\xf5\x28\x09\x95\x5f\xa8\xd7\xbb\xf3\xdd\xe9\x9e\xcb\x8a\xf7\x85\x7e\x07\xac\x1e\x56\xeb\xcd\xe3\xc9\xd7\x9f\xcf\x0a\x21\xf5\xb2\x2e\x8a\xb2\xe2\x26\xd1\x51\x85\xfb\xf6\x0f\x8e\x09\x2f\xf7\x89\x69\xa2\xe6\x1e\x40\x24\xc1\x10\x99\xdd\x71\x09\x2c\x48\x84\x72\x07\xf7\x34\x08\x29\xfb\x7e\x14\x70\xdf\xe0\x92\x78\xee\x75\x44\x3b\x28\x99\xd2\x6f\xea\x85\x5a\xbc\x40\x69\x64\xfb\x85\xfe\xfe\x6e\xc2\xa8\x43\xae\x96\x37\x36\xaf\xd5\x07\xea\xc9\xdf\x1f\x0e\x11\x9a\xb1\xf1\x7c\xa3\xb7\x63\xf9\x90\xcb\x08\xc4\x62\xc6\x72\xce\x6d\x89\x65\x66\xc4\xa4\x6c\xb6\x24\x7a\x85\x17\xeb\x1a\x46\x8f\xd9\x08\xfe\x67\x63\x5b\x99\xe1\xd9\xbb\x1d\xa8\xc1\x8c\xe3\x6c\x5e\x1d\x55\x83\xdb\xe1\x3c\x20\x4e\x51\x91\x7f\x90\x59\x20\x4a\x9d\x2b\x10\x95\x6f\x0f\x22\xb6\x45\x54\x6f\x2f\x7d\x95\x8c\xd4\x20\xa7\x9c\xe7\xd7\xa2\xf7\xbb\xb0\xce\x72\x25\x69\x91\xf9\x64\xbf\x01\x00\x00\xff\xff\x5c\xa2\x94\x99\xca\x02\x00\x00")

func definitionsFieldsTomlBytes() ([]byte, error) {
	return bindataRead(
		_definitionsFieldsToml,
		"definitions/fields.toml",
	)
}

func definitionsFieldsToml() (*asset, error) {
	bytes, err := definitionsFieldsTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "definitions/fields.toml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xfe, 0xa5, 0x29, 0x70, 0xa5, 0x18, 0x85, 0x7b, 0xfe, 0x73, 0xbf, 0x8c, 0xef, 0x61, 0xf6, 0xa1, 0xa2, 0x5, 0x4c, 0xf, 0xe9, 0x27, 0xee, 0x26, 0x1b, 0x41, 0xe8, 0x33, 0x87, 0xaa, 0x98, 0x7f}}
	return a, nil
}

var _definitionsInfo_object_metaToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\xbd\x6e\xdb\x4c\x10\xec\xef\x29\x16\x6a\xdc\x7c\x54\xf5\x25\x9d\x8b\x00\x2a\x12\xc0\x82\x02\xd8\x42\x0a\xc1\xc5\x89\xb7\xa4\x36\xe2\xfd\x64\x77\x69\x9b\x08\xf2\xee\xc1\x1d\xc5\x48\x96\x98\x22\x1d\x6f\x7f\x67\x67\x86\x3b\x9b\x12\x06\x57\xc5\xa6\x11\xd4\x67\xa3\x43\x42\xb8\x87\x05\x05\xfd\xf8\xff\xc2\x38\x94\x9a\x29\x29\xc5\x90\xa3\x9f\x4a\xf1\xa6\xd4\x02\x09\xe8\x01\x61\xec\x84\xd8\x94\xd7\x38\x0e\xe2\xfe\x3b\xd6\xba\x5c\x18\xb3\xab\x63\x50\x0c\x5a\x75\x18\x5a\x3d\xdc\x6c\x38\x17\x78\xf7\xe1\x9c\x15\x65\x0a\xed\x65\x3a\x67\xe6\xf2\xa8\xb6\x9d\x8b\x93\xbb\x8d\xe2\x5b\x8a\xac\x70\x0f\xca\x3d\x5e\x1f\xb7\x58\x98\x2f\xab\xe9\xaa\x3e\xd0\x8f\x1e\xe1\x88\x03\x50\x00\xd1\xc8\xb6\xc5\xa5\xc9\x15\x8f\x9f\x37\xdb\x87\x15\xec\x11\x6c\x00\xbb\x97\xd8\xf5\x8a\x90\xac\x1e\xa0\x8e\x3e\x59\xa5\x7d\x87\xf0\x4a\x7a\x28\x93\xd4\x72\x9b\xf9\x49\xc8\x56\x29\xb4\x20\x83\x28\xfa\xca\x61\x43\x01\x1d\x34\xd4\x8d\xdd\xf2\x1f\x44\x06\x3b\xad\x26\x87\x41\xa9\x21\x64\x90\x84\x75\xfe\x72\xb0\x1f\x40\x90\x5f\xa8\xc6\xa5\xc9\x80\xcd\xae\xb3\xa2\x95\x8f\xae\xe4\xcf\x17\x2b\x79\x5c\x3e\x91\xc7\x52\x43\xe1\x58\x8d\x38\x6e\x39\xb9\x62\xe1\x81\xc2\xf1\x69\x84\x7c\xa2\x42\x06\x9f\x07\x4c\x87\x34\x91\xa1\xbc\x2f\x34\xf6\xd1\x5d\x68\xb3\x29\x89\x75\x74\x78\xcd\xb8\xd9\xf9\xbe\x53\x4a\x96\xb5\x9a\xd3\xe7\x0a\xcb\x7a\x2a\x3e\xeb\x92\x5f\x40\x2e\xdb\xad\x7c\x5e\x80\xc8\x1c\xfe\xbb\xe4\x5f\xb3\x6e\x24\x80\xa4\x07\xe4\xd1\xc3\xef\x34\x8d\x63\x90\xb1\xb3\x4a\x2f\xa7\xa0\xc6\x57\xcb\x4e\x26\x63\xdc\x09\x7c\x8b\x7c\x5c\x11\x83\xc3\xfc\x03\x08\xc4\x00\xbd\x20\xdf\x09\x50\x48\xbd\x2e\xcd\xb8\xe9\x6c\x9e\x6d\xa0\x37\x10\x1d\xba\x3f\x52\x9e\x8c\xe1\x51\xad\xb3\x6a\xdf\xfd\x2a\xc8\x8d\xad\xf1\xe7\xaf\x1b\x8e\x1e\x4b\xd3\xfa\xd4\x53\x00\xa1\x9c\x3c\x06\x93\xc7\xa6\x91\x85\xa7\x0c\x6b\x66\x89\xb7\x69\x37\x92\xf6\xfc\x17\x39\xb6\x82\x7c\xbd\x28\x0f\x9b\x5b\xf3\x3b\x00\x00\xff\xff\x9a\xe9\x67\x93\x58\x04\x00\x00")

func definitionsInfo_object_metaTomlBytes() ([]byte, error) {
	return bindataRead(
		_definitionsInfo_object_metaToml,
		"definitions/info_object_meta.toml",
	)
}

func definitionsInfo_object_metaToml() (*asset, error) {
	bytes, err := definitionsInfo_object_metaTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "definitions/info_object_meta.toml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xbe, 0xfc, 0xc9, 0x91, 0x3, 0x38, 0x82, 0x77, 0xaa, 0x97, 0xb1, 0x0, 0x25, 0xe, 0xb3, 0x46, 0xd3, 0x3, 0x6b, 0x3, 0xd8, 0xf3, 0xa, 0x4c, 0x8a, 0x82, 0x90, 0x2f, 0xa0, 0x14, 0x15, 0x2a}}
	return a, nil
}

var _definitionsInfo_storage_metaToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\xd2\xbd\x72\xab\x30\x10\x05\xe0\x5e\x4f\xa1\xa1\xe7\x56\x77\xd2\xf9\x11\x5c\xa5\xf4\xa4\x58\xa3\xc5\xd9\x09\xfa\x99\xd5\x12\x43\x32\x79\xf7\x8c\x84\x21\x63\x4c\x1c\xdb\x94\x08\x9d\xef\x2c\x48\x3b\x08\x01\x9d\x29\x5d\x6b\xf7\xc8\xa5\x85\x8e\x6c\x6b\x5f\x94\xf4\x01\xf5\x46\x17\xe4\xa4\x50\x06\x63\xc5\x14\x84\xbc\x4b\x6b\x5b\xe8\xf4\x10\xd3\x43\x2c\x6a\x72\xe3\x8a\x0f\xc8\x90\x76\xfe\x2b\x94\x1a\xf5\x48\x1f\xb8\x68\x3f\xfd\xbf\xa6\xa7\x58\xa2\x03\xf2\x55\x5e\xbc\x40\xf3\x68\x49\x0e\x4f\x55\x8b\x35\x95\x0f\xfd\xbd\x7c\xda\x36\xa8\xb5\x67\x9d\x84\x19\x5a\xa3\x54\xaf\xeb\xd4\x4c\xcc\xd8\xc6\x57\xf9\xe1\x07\x8b\xc2\xe4\x0e\xe9\x9d\xf5\xef\xb8\xae\x31\x09\xb3\x42\xdb\x36\x42\x01\x58\xee\xbd\x43\x59\x4e\xc1\xd3\x2d\x4a\xbf\x7f\xc2\x7e\x2d\x79\x64\xfc\x0c\xe6\x6f\x30\x58\x93\x43\xa3\xf7\xbd\x8e\xe2\x19\x0e\xc8\x8b\x3e\xb9\x9b\xfc\x61\xdb\x0d\xbe\x03\x8b\x97\x27\x82\x5d\xf0\x2c\x7a\xa3\x85\x5b\x54\x6a\x17\xfb\x28\x68\x4b\x8b\x02\x06\x04\xce\xea\x91\x6b\xa8\xf0\xf3\xeb\x62\x88\xe7\x1c\xda\x9e\x32\xb9\x16\xa3\x1e\xa8\x69\x9e\x91\xcc\xc3\x1c\x3d\xbf\x95\x86\xf8\xef\x81\x8e\x4c\xb2\xf2\xc6\x64\xe2\xec\x34\xbf\x03\x00\x00\xff\xff\x41\xde\xf8\x57\x73\x04\x00\x00")

func definitionsInfo_storage_metaTomlBytes() ([]byte, error) {
	return bindataRead(
		_definitionsInfo_storage_metaToml,
		"definitions/info_storage_meta.toml",
	)
}

func definitionsInfo_storage_metaToml() (*asset, error) {
	bytes, err := definitionsInfo_storage_metaTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "definitions/info_storage_meta.toml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9c, 0xdc, 0x18, 0x27, 0xc0, 0x7d, 0xb9, 0x57, 0x8b, 0xe9, 0x90, 0xfd, 0xf8, 0xb3, 0x2, 0x40, 0x16, 0xce, 0x94, 0xf9, 0x18, 0x51, 0x52, 0x7a, 0xff, 0x6b, 0x10, 0x2d, 0x50, 0x8, 0x2e, 0xb0}}
	return a, nil
}

var _definitionsOperationsToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x4b\x73\xdb\x38\x12\xbe\xeb\x57\x74\x69\x0e\xb9\xd8\xaa\x39\x4f\x55\x0e\x19\x3b\x5b\x93\xaa\xbc\x76\x9c\xd9\x39\xa4\x5c\x16\x44\x36\x45\x6c\x48\x80\x01\x40\x39\xda\x5f\xbf\xd5\x68\x90\x04\x5f\x12\xed\x68\xb2\x87\x9d\x4b\x4a\x26\xd1\x0f\x34\xbe\x6e\x7c\xdd\xcc\x67\x51\x55\xa8\x52\x34\xf7\xab\x14\x6d\x62\x64\xe5\xa4\x56\xf0\x12\xd6\xd2\x82\xcb\x11\xa4\x72\x68\x32\x91\x20\x64\xda\xc0\x2b\xbf\x1a\x0c\x16\xc2\x61\x0a\xba\x42\x23\x48\xc0\x6e\xd6\xab\x55\xab\x6b\xa3\xab\x4d\x62\x50\x38\x7c\xe0\x47\xf7\xab\x4a\x18\x51\x5a\x78\x09\x9f\xd7\x95\x70\xf9\xfa\x7e\x65\xd0\xd6\x85\xe3\x47\x7a\x3d\x32\xbf\x5e\xaf\x1e\x65\x51\x00\xeb\x01\xa1\x80\x55\x81\xde\xfd\x1b\x13\xb7\x59\xad\x7e\xfa\x09\x7e\xc5\x5c\x1c\xa4\x36\xab\xd5\x35\xdc\xf8\x85\xc1\xc1\xbb\xdf\x3e\xfc\xf1\xf6\x76\x24\x2c\x76\x05\x06\x05\xf0\x28\x5d\x0e\x95\xb6\xd2\x5b\xfc\x19\x84\x4a\xc1\xca\xff\x20\xfc\xbc\x99\xd1\xf6\xfe\xc3\x27\x30\xe8\x6a\xa3\x48\x23\x1a\xa3\x0d\x08\x8e\x52\xd0\x89\xdf\xa4\x75\x9b\x15\xc0\x35\xdc\xa1\x39\xc8\x04\x5b\x4f\x72\x4c\xbe\x78\x1b\x29\x16\xe8\x30\x96\x92\x19\x0b\xda\xcd\x8a\xb6\xdd\x8f\xe3\xa3\x91\x51\x18\x07\x41\xf2\x11\x0a\x71\x49\xb4\x72\xa8\x1c\x38\x3d\x0e\xd6\x3a\x8e\xbf\x5e\x5f\xc1\xda\xd0\x3f\xb4\xdd\xc1\x49\xa8\xf5\xfd\x8a\x85\x1e\x4a\x9d\x22\x19\x61\x55\xa3\xf3\xd5\x65\x29\xdd\x49\xc7\x78\x89\xdf\x74\x26\x95\xb4\x79\xe4\x58\x65\x74\x82\xd6\x0e\x3d\x3b\x61\x7c\x57\xe8\xe4\xcb\x52\x9c\xfe\x4a\x8b\xe7\x60\x1a\x34\x45\x28\xf5\x4f\x2e\x00\x52\x50\xf8\x08\x5e\xd9\x34\x4c\x1b\x5c\xb1\x7b\xdf\x03\x2b\x97\x0b\x07\xb9\xb0\xa0\x84\x93\x07\x04\x5b\x57\x95\x36\xce\x6f\x7e\xab\x0f\x68\x3c\x6e\xb6\x90\x6a\xb4\xea\x85\x83\xf7\xaf\x5f\xdf\x12\x34\x18\x87\x43\xdd\x16\xb4\x01\xa5\xa7\x6c\x34\x0a\x72\x71\xc0\x73\xc6\x9e\x83\xf5\xe8\x30\x18\xea\xe1\x2c\xa6\x00\xe5\x17\xf4\x80\xce\xc1\x3e\x89\xef\x2b\x58\xef\x64\x7a\x1e\xe6\x5e\xd3\x10\x1e\xba\xdc\x49\x75\xd2\xa7\xb0\x84\x1d\xb1\x04\x43\x9f\x7f\xb3\x89\xb7\x93\xa9\x5d\x66\xbc\x90\xd6\x9d\xb2\x4c\xef\x1b\xb3\x3b\x2c\xb4\xda\x53\x48\x5c\x2e\xed\x8c\xf5\x7e\x08\x76\xf2\x84\x1b\x89\xae\xe4\xd2\x5c\xbb\xd1\xd5\x71\xd3\x09\x71\xdc\xaa\x63\x2f\x9d\xac\x49\x68\xf3\xa9\x75\x27\x92\x48\x57\x47\x0a\xdd\x07\x46\x8a\x36\x50\xd6\x85\x93\x55\x57\xb2\xa5\xf2\xe6\x2d\xe3\x73\xe2\x0a\x20\x0d\x5a\x15\x47\xd6\xa5\x15\x7a\x14\xfa\x27\xf4\x47\x93\x93\x31\xc8\x6f\x3f\xbc\x7f\xf1\xa9\x4d\x8f\x06\xd9\xec\x0b\x28\xad\xae\xb1\xac\xdc\x11\x52\x69\x30\x71\xda\x1c\xc9\x2f\xff\x36\x93\x05\x5a\x30\x98\xd4\xc6\xca\x03\x16\x47\xd6\xfb\x87\x45\xd3\xaa\x93\x65\x55\x60\x49\x60\x3d\xa1\x50\xf8\xe2\x5d\x1d\x63\x5d\xb0\x3b\xd2\x56\x4b\x8b\x45\xc6\x7a\x6f\x58\x01\x59\x65\xe4\x77\x0a\x42\xda\x85\x1a\xb2\x7d\x6d\x0c\x87\xf0\x9d\x4e\xf1\x8d\x3a\x88\x42\xa6\xdb\x4d\x13\x9e\xd3\x45\x27\xb5\xae\x5f\x18\x2e\x5d\x75\x46\x06\xfe\xfa\xca\x13\x99\x8c\xaa\xcf\x35\xbc\x02\x5b\x27\x74\x09\x65\x75\x00\x9f\xae\xf8\x9a\x00\x9b\xeb\xba\x48\x61\x47\xc5\x86\x8e\xd0\xe1\x15\x3c\xe6\x32\xc9\xa1\x44\xa1\xec\x40\xed\x0b\xdb\xd6\x24\x32\x5d\xa2\x13\xa9\x70\x22\xd2\xe2\x51\x2b\x4a\xa4\x10\x5b\x93\xb4\x40\xe4\x12\x48\x27\xb9\x30\xd7\x6e\x9b\x43\xdf\xb4\x82\xd1\x35\x96\xca\xb1\x9a\xf1\xfd\x94\x4a\x33\x59\x20\x66\xae\xbc\xd5\xe7\x0c\x5d\x92\x2f\xf5\xf0\x1f\xb4\xd8\x7b\x17\xc4\xc8\x3f\xff\x73\x7c\xbf\x5e\xc1\xba\x36\xc5\x7c\x45\xf0\x62\x90\x19\x5d\x82\x80\xbd\x3c\xa0\x82\xda\x14\x04\x28\x12\x1f\xe7\xbf\xb7\xfd\xc4\x6b\x75\x0c\x05\xb6\xda\x72\x86\x45\x60\x78\x0e\x10\x0c\x7e\xad\xa5\x91\x6a\xcf\x3b\xa4\x57\xb5\x29\x1a\x50\x94\x84\xee\x65\x21\x7f\xa7\x0f\xb8\x69\x65\x28\xde\xf4\xe3\xc9\xf5\x97\x84\xba\xab\xeb\x6c\xb1\x25\xab\x5c\x5a\x4b\xfe\xf5\xec\x62\xcb\x86\xa7\x6a\xe3\xc9\x92\x3a\x2f\x36\xae\x9e\xef\x78\xed\xf7\x54\x4f\xaf\xe2\xef\xea\x39\x99\x32\x8c\x80\x36\x63\x82\xbe\x1f\x54\x3e\x99\x25\x08\xe3\x16\xe7\x4b\x23\x30\xd7\x20\x44\x1a\xa3\xea\xda\x3e\xbd\x54\xa3\xd0\x2a\x9c\x6b\x66\x3b\x47\x9f\x5c\xd3\x46\xa1\xe9\x08\x76\xb4\x8f\x85\x24\xbb\x73\xf4\x0c\xd1\x96\x2a\xc5\x6f\x23\xaa\x7d\x05\x6b\x12\x1e\xd3\x4d\xff\x74\x22\xde\x01\x34\xe7\x3c\x6d\xd6\xc5\x1e\x42\x5d\x15\x5a\xa4\x81\x54\x29\xeb\x4c\x9d\xb8\x8e\x55\x4e\x6c\x80\xa4\x26\x18\xf9\x8c\x6f\x9e\x94\x9f\xf1\xcb\x13\x73\xaf\x76\xc8\xcb\x67\x03\xd9\x0f\x59\x35\x41\xcd\x1b\x7f\x2a\xb1\x5f\x8a\xf3\x8f\x62\x8f\x13\x10\x0f\xa9\xd8\x14\x03\x23\x54\xaa\x4b\x3e\xf6\x4d\x6b\x21\xc2\x3d\xfd\x7d\x29\xc8\x93\xae\x33\x13\x1c\xef\xf5\x33\xf1\xde\xba\xce\x48\x67\xcf\x17\x81\xdc\x56\x98\xc8\x4c\x26\xa0\xb3\xcc\xe2\x59\xa0\xf3\xaa\xf3\x4d\x25\x79\x40\x8e\x19\x14\x93\x14\x6a\xbd\x5e\x4d\x1e\xdd\xef\xb4\x7e\xb3\x5a\xdd\x62\x65\x30\xa1\x03\xfc\x85\xae\x41\xb8\x73\xda\x88\x3d\xfe\xf6\xe9\xd3\xc7\x3b\xb9\x57\x68\x40\x2a\xeb\x50\xa4\x4d\x04\x82\x21\x8a\x81\xff\x79\xee\xe0\x4e\xf2\xaf\xca\xe8\x83\x4c\xe9\xec\x1e\xc5\xb1\x29\xe1\x89\x50\xe0\x55\x47\xc7\x30\xe1\xe9\x3f\x6b\x34\x47\xf2\x91\x7c\xfd\x1d\x45\x3a\xf4\x34\x10\x8b\x71\x4c\xc8\x40\x29\xa4\x72\x42\xaa\x28\xb1\x2d\x6f\xbd\x21\x24\x5c\xa9\x1b\x25\x1d\x5c\x17\x10\xe0\xa0\x89\x63\x27\x54\x82\xfd\xd3\x56\xa2\x1c\x4e\xc5\x48\x02\x3d\x17\x8e\x0d\xf2\x2d\x39\x6d\x30\xdc\xa0\x62\xa1\xb1\xbe\xe2\x3d\xce\x54\x96\x3d\x3a\x10\xe0\x39\xc9\x58\xb1\x07\x4e\x43\xd7\x9e\xb7\x23\xaa\x5c\x27\x6a\x9a\x28\x8a\xb1\x59\x0b\xb5\x4a\xd1\x70\x81\xeb\xcc\xf7\x8d\x49\x36\x15\x64\x97\x95\xaf\xc1\x81\xf3\x79\x07\x0d\x71\xf8\xc7\x00\xaf\x84\x34\x21\x73\xbb\x74\x9c\x87\x79\x73\x56\x2d\xf5\xf5\x7c\x7c\x96\xf8\xde\xf2\x72\x4f\x74\xd3\xe6\xf7\xb3\xc9\xaf\x41\xe6\xb1\x45\x71\x92\xec\xf2\xb2\x07\x3a\x81\x1e\xb9\x6d\xdd\x91\x16\x64\x8a\x65\xa5\xa9\xa2\x05\xc3\x1d\x41\x6b\xb6\x58\x3c\x8a\xa3\x6d\x4a\xaa\x92\x05\xd7\x54\x5e\x1e\x14\x85\xd2\xab\xf0\x80\xa6\x65\xc5\x7c\x7b\xbe\xd7\xee\x35\xd5\xdb\x6d\xbc\xbe\xbf\xad\xe9\xf9\x62\x4b\x53\x43\xea\x47\xa7\xd8\xd0\xbd\x69\xd8\x05\xfb\x49\x6d\x0c\x45\xa1\x05\x5f\x23\x35\x00\x1a\x3d\x5e\xdf\xaf\x0a\x9d\x88\x02\x5e\x82\x33\x35\x0e\xcc\x31\xc2\x4f\x41\x86\x6f\x78\x06\x4c\xef\x8a\x93\xf3\x08\x0a\x6e\x72\x92\x74\x77\xc9\x74\x97\x3a\xf8\x32\xd0\x00\x21\xc5\x4c\xd4\x85\x83\xed\x5b\x69\x7d\xf3\xe1\xbb\x8e\xc1\xe2\x0e\x10\xed\xb2\x5b\x69\xb6\xfe\x63\x86\xae\x9d\x0f\x3c\x1f\x81\xe7\xf8\xff\x92\xc6\xd5\xa2\xa0\x25\xb1\xb2\xfe\x91\x6d\xef\x9c\x70\x5b\xaa\xee\x05\xa5\x20\x6b\xde\x4e\x9d\x95\x41\x31\x33\xea\xa7\x37\xde\x38\x75\x58\x2f\x2c\x84\xb3\x99\x68\xf8\x1f\x7b\xc1\x1e\xdc\xa7\x44\x21\xf5\x43\x22\x8a\x62\x27\x92\x2f\x13\xd7\x6b\xdf\x9f\xa6\xe2\x9f\xe7\x26\x73\x15\xa1\x87\x94\x25\x04\x26\xfa\x78\x44\xf1\x16\xea\x08\xa2\x92\x90\xf8\xfc\x9d\x26\x34\x31\x99\xb1\xa8\x52\x2f\xf4\xea\xe3\x9b\x20\x34\x5c\x26\x92\x04\x2b\x07\x5d\x13\x0a\xe4\x3c\xd1\x9e\x60\x9b\x5c\x9f\x3a\x1d\xeb\xc4\x19\x68\xf7\xf7\xbe\x90\xbf\x91\x5a\x10\x1e\xcb\x84\x16\xba\x82\xa4\xca\x34\xe8\x2c\x9a\x73\x8f\x41\x4e\x52\x67\x77\x94\x4b\x35\x6c\x92\x6f\xbc\x88\xef\x5a\x53\x99\x65\xe8\xf3\xbe\x85\x7d\x60\xb0\x04\xee\xee\x6d\x2f\x2e\xb3\xdf\xde\x64\x16\xd2\x54\xaa\x7d\xec\x0b\xd1\x72\xe1\x92\x7c\x22\xa2\x9e\x25\x4e\x8e\xad\xe2\xcf\x67\x5d\x7c\x4f\x63\x77\x26\xbe\x4c\x45\x7d\xc7\xcb\x1f\x0f\x7c\x0e\x8d\x23\xfa\xa7\x5f\x37\x28\x1a\xb5\x45\x63\xa1\x12\xd6\xfa\x6a\xbe\x95\x7a\x43\x5c\x0b\x29\xdf\x97\x8e\x1e\xa6\xa4\x4f\xcc\x20\xe2\x55\x92\x05\x9f\x3d\x84\x98\x34\x3d\x9a\x46\x74\x5f\x71\x79\xce\xd3\xf3\x20\x03\xe9\x82\x1f\x9b\x61\x94\xfe\x5f\xbf\xac\x0d\xa6\x33\x8c\xb0\xa7\x8d\x67\x9e\x35\x9a\xe1\x7b\x0f\x89\xea\x07\xa3\x06\xbf\xd6\x68\x87\x17\xff\x43\xee\x5c\xf5\x60\x7d\x03\xb3\x8c\x0b\xde\x35\x17\xff\xd9\x76\x56\xd4\x2e\x47\xe5\x64\xe2\xdf\x6e\x66\xac\x52\x76\x7f\xa5\x0e\xc5\xff\xcd\x6f\xce\xdc\x6e\x7e\xe3\xed\xa4\xd6\x4f\x12\x77\x47\xa8\x2d\x95\x13\xaf\x0a\x7c\x9d\x40\x47\x19\x49\x79\xdc\x39\xd2\xc6\xc1\x4e\xdf\x89\xf8\xad\x92\x66\x58\x92\x0d\x7e\x8d\x2f\xbc\x73\xde\x87\x5a\x35\xdf\xec\x0e\x2b\xcc\xc5\xbc\x6f\xee\xf0\x4b\xec\x62\x49\x4b\x35\x49\xd3\x7f\xcc\x59\xb4\xdd\xe8\xd3\x11\xfc\xae\x9b\x23\x3d\x0b\xc4\x93\xa6\xa7\x42\x38\x9e\x56\x9e\x6d\x88\xbb\x26\xfb\x7f\x17\xc5\x59\x4c\x7f\xcf\xbc\xf2\xbb\xf7\xa3\xc7\xd3\xcd\x8b\xed\xee\x42\xa3\xc4\xcb\xec\xf1\x32\x5b\xfa\x0b\x27\xb7\x97\xd9\x26\x0f\x7b\xcf\xed\xb7\x90\x6a\xf1\xff\x34\xa2\xb5\xeb\x56\x26\x1a\x9b\xd2\x83\x49\xee\xe8\x84\xd9\x8f\x66\x87\xd3\x04\xfc\xcf\x5e\xaa\x92\xc6\x88\x70\xf7\xd8\xe1\x96\xb4\x6f\x7d\xf0\xb6\x6c\x60\x1b\xa8\xf4\x2e\xcc\x80\x89\x5e\x10\xff\xd9\x59\x5d\xd4\x0e\x43\x73\x7a\x0d\x6f\xb2\x4e\x42\xe9\x86\xb6\x5c\x85\xae\xe4\x2d\xd9\x0c\x7d\xc0\x9c\x33\xdd\x07\x59\xd6\xc6\xae\xb0\x9e\x5f\x3c\xef\xe9\x9e\x4a\x4b\x0d\xf2\xb1\x8c\xc4\xc7\xa6\xc2\x4c\xc4\x73\x8b\xde\xd2\x1e\x1d\xf4\xa5\x6b\xd2\x8d\x91\x49\xda\xd7\x12\xb3\x0d\x55\x9c\xfa\x08\x18\xf8\xe3\x63\x8e\xbd\x0f\xa3\x9e\xc6\x79\x03\x4d\xf1\x6e\x69\x1b\x13\xb1\xd8\xc3\xe8\x44\xc8\x14\x33\xa5\xd0\x7f\x13\xe3\x0b\x9d\xf7\x75\xec\xda\xcd\xa0\xf1\x3f\x70\x4f\xef\xf1\x05\x19\x0a\x57\x1b\x64\xaf\x1a\x8f\x9a\xe6\xb4\xcf\x29\xbb\xf1\x92\x6d\x18\x20\x2a\xff\x1f\x22\x7d\x3d\x69\x34\x75\xd3\xa5\x03\x86\xb9\xfa\x7f\x03\x00\x00\xff\xff\xdc\xce\x65\x5d\x24\x2a\x00\x00")

func definitionsOperationsTomlBytes() ([]byte, error) {
	return bindataRead(
		_definitionsOperationsToml,
		"definitions/operations.toml",
	)
}

func definitionsOperationsToml() (*asset, error) {
	bytes, err := definitionsOperationsTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "definitions/operations.toml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x59, 0x6f, 0x3d, 0x71, 0xf1, 0xf3, 0xd6, 0x7d, 0x33, 0xd8, 0xdd, 0x37, 0x69, 0x27, 0xec, 0x38, 0xfc, 0x4b, 0xac, 0x6d, 0x7a, 0xb8, 0x96, 0x1f, 0xbb, 0x5c, 0x23, 0x39, 0xcb, 0x52, 0xf4, 0x1b}}
	return a, nil
}

var _definitionsPairsToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x55\x4d\x6f\xe3\x38\x0c\xbd\xfb\x57\x10\x39\xed\x02\x45\xf6\xb2\xbb\x87\x05\xf6\x56\x0c\x3a\x40\x07\x9d\xe9\xc7\xa9\x28\x12\x59\xa2\x63\x4e\x64\xd1\x43\xd1\x49\x3d\xbf\x7e\x20\x39\x8e\x83\xb6\x49\x9b\xb9\x19\x12\xdf\x7b\x24\xf5\x48\x3f\x5a\x0e\x8a\x41\x17\x8d\xfb\xe7\xa9\xd0\xbe\x45\xf8\x1f\x66\x51\x85\xc2\x6a\x56\x14\xfb\xeb\x74\xf3\x54\x38\xac\x4c\xe7\xd5\x94\x3e\x85\xa9\x74\x78\x04\x42\xa1\x33\x4a\x1c\x16\xca\x6b\x0c\xaf\x89\x1d\x46\x2b\xd4\xa6\x90\x7c\xdc\xa2\xa5\xaa\x07\xad\x11\x0e\xe1\x90\xe1\x50\xb1\x80\xa7\xa8\x99\x5d\xd0\x61\x50\x32\xfe\xc3\xac\x35\x6f\x41\x19\x5a\xe1\x0d\x39\x84\x89\x21\x13\x47\x94\x0d\x59\x84\xf4\xa9\x2c\x66\x85\x49\x06\x83\x6b\x99\x82\xfe\xae\xc8\x88\x3f\x25\xf1\xdc\x92\xe0\x24\xa0\xd4\xe0\xfc\xb2\x93\x5c\xf9\x2b\x9d\xd9\xac\x18\xa5\xb6\x35\x86\xdc\xaa\x4e\x3c\x08\x6a\x27\x01\x1d\x94\x3d\x08\x1a\x5b\xc3\x96\xbc\x87\x81\x7c\x5e\x14\x97\xd8\x0a\x5a\xa3\xe8\xfe\x83\x87\x88\xb0\xfc\xd6\xa1\xf4\x77\xb4\x0a\x57\xf7\xf7\x5f\x6f\xd1\xb8\x25\x50\x88\x8a\xc6\x01\x57\xb0\xbc\x4d\x1c\x4b\x30\xc1\xc1\x72\x20\x59\x02\xc5\x2c\xe7\xd0\x7a\x23\xe8\xc0\xc8\xaa\x6b\x30\xe8\xbc\x48\x69\x15\x8f\x14\x14\xc5\x62\xab\x2c\x53\x39\x9f\xa7\xc3\x14\x93\xde\x6f\xd1\xb0\x3b\x28\xf8\x9a\xa2\x7e\x61\x97\x9b\xe1\xd9\xe6\xba\xcf\xb2\xca\x08\x3a\xd1\xe4\x60\x1a\x3c\x8b\x73\x87\x85\x04\x4c\x04\x5c\x7e\x47\xfb\x32\xf3\x9b\x7c\x38\xe4\xfe\x82\x6b\xba\x82\x9a\x42\xb6\x2c\x57\x55\xc4\x03\x27\x51\xd0\x7f\xff\x3e\x9a\xc4\x10\x9d\x4b\xd2\x9a\x22\x08\xfe\xe8\x30\xea\xc5\x3e\xb3\xfc\xbe\x11\x71\x9d\xfc\x96\x43\x76\x90\x12\x2b\x16\x4c\x2e\x70\x49\xb6\xe9\xbc\x52\x6b\x44\x17\xe4\xde\x1a\x6d\xe2\x85\x35\xde\x97\xc6\xae\xa7\xeb\xaa\x0b\xf6\x8f\xc7\xa7\xb2\x57\xfc\x73\xf6\xd6\xbc\x1f\xc9\x7a\x5b\x1b\x05\x65\xc7\x80\x1b\x94\x1e\x92\x99\x61\x3b\x24\x03\xce\xa8\x81\x4a\xb8\x81\xc8\x9d\xd8\xdc\xd7\x48\x3f\xf1\xa3\x2d\x49\xb1\xef\x35\x84\x83\xef\x07\x35\x4f\x0d\x29\x3a\xd8\xad\xae\xac\x9e\x14\xb7\x2c\xeb\x85\x23\x79\xd7\x0e\x07\xa3\x96\x1c\x91\x70\xe0\x48\x8e\xb8\xec\x62\x57\x31\xb7\x38\x8c\xee\x90\x4f\x99\x6a\xf7\x46\x69\x83\xfb\x77\x72\x24\xf3\x62\x4c\x03\xee\xae\x6e\x1e\xae\x2f\x53\xa0\x09\x60\xca\xc8\xbe\x53\x84\xd6\x68\x7d\x10\x34\x72\xed\x5e\x22\x51\xfd\x05\x54\x41\x60\x85\x88\xfa\x26\xdd\x43\xa0\x67\x88\xda\xfb\xa1\x69\x83\x83\xf7\xdd\xda\x55\x10\xe7\xc5\x27\x16\xa8\xe2\xcb\x0b\xc8\x15\x04\xc7\xdb\x08\xad\x37\x5a\xb1\x34\x17\xb9\x11\x25\xd6\x66\x43\x2c\x69\x1f\x38\xac\x28\x6d\x9d\x88\xad\x11\xa3\xe8\xfb\x71\x19\x58\x63\x6b\x5c\xa4\xe6\x0b\x7f\x7c\x47\xe7\xcd\x9f\x90\xb0\x43\x9e\x98\xe9\xf1\x9f\x84\xc1\xb2\xa3\xb0\x3a\xfb\xf7\x92\x5c\x31\x82\x4f\xe8\x08\xc6\x96\x43\x1c\x8a\x49\x82\x8e\x62\xcb\x91\xce\xde\x53\x23\xd1\x64\xc9\x89\xe8\x98\xfe\xaf\x00\x00\x00\xff\xff\xc7\x09\x29\x6c\x99\x07\x00\x00")

func definitionsPairsTomlBytes() ([]byte, error) {
	return bindataRead(
		_definitionsPairsToml,
		"definitions/pairs.toml",
	)
}

func definitionsPairsToml() (*asset, error) {
	bytes, err := definitionsPairsTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "definitions/pairs.toml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x99, 0x57, 0xd2, 0x6e, 0xe8, 0x5c, 0xba, 0xa5, 0xa5, 0x47, 0x60, 0xf0, 0xb0, 0x40, 0x97, 0x6f, 0x92, 0x26, 0xbb, 0xdc, 0x7a, 0x75, 0xa5, 0xbc, 0xc3, 0x8a, 0xd2, 0x7a, 0xc7, 0x34, 0x57, 0xdc}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"definitions/features.toml":          definitionsFeaturesToml,
	"definitions/fields.toml":            definitionsFieldsToml,
	"definitions/info_object_meta.toml":  definitionsInfo_object_metaToml,
	"definitions/info_storage_meta.toml": definitionsInfo_storage_metaToml,
	"definitions/operations.toml":        definitionsOperationsToml,
	"definitions/pairs.toml":             definitionsPairsToml,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"definitions": {nil, map[string]*bintree{
		"features.toml":          {definitionsFeaturesToml, map[string]*bintree{}},
		"fields.toml":            {definitionsFieldsToml, map[string]*bintree{}},
		"info_object_meta.toml":  {definitionsInfo_object_metaToml, map[string]*bintree{}},
		"info_storage_meta.toml": {definitionsInfo_storage_metaToml, map[string]*bintree{}},
		"operations.toml":        {definitionsOperationsToml, map[string]*bintree{}},
		"pairs.toml":             {definitionsPairsToml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
