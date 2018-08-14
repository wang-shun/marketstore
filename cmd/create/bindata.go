// Code generated by go-bindata.
// sources:
// cmd/create/default.yml
// DO NOT EDIT!

package create

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

var _cmdCreateDefaultYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\xc1\x6e\xe3\x36\x10\xbd\xf3\x2b\x06\xd6\xa5\x0d\xa0\x58\x6e\x9b\x16\xe1\x2d\xbb\xdd\x6d\x0e\xd9\x6e\x80\xb4\x8b\xee\x49\x18\x49\x63\x99\x30\xc5\xd1\x0e\x47\x71\x5c\xec\xc7\x17\x94\x1d\x27\xb6\xe3\x22\xd1\x49\x7c\xef\xf1\x3d\x0e\x87\x64\x06\xf9\x6b\x3f\x93\xc1\xd5\xa0\x9c\xb7\x14\x48\x50\xa9\x81\x0e\x65\x49\x1a\x95\x85\xa0\xe6\x30\x77\xed\x20\xa8\x8e\xc3\xb9\x79\x9b\xaf\x30\x2b\x34\x4e\xa8\x56\x96\x35\xf0\x1c\x74\x41\xd0\xa0\x62\x85\x91\x4c\xa2\xcb\x1d\x6d\x47\xc2\x64\xe0\x5d\x54\x0a\xd0\xb3\x28\xe4\xe3\x8c\xf1\x97\x1e\x7a\x8e\xd4\x40\xb5\xde\x73\x81\x48\x72\x4f\x62\x36\xb3\xca\x24\xb5\x70\x71\x79\xf9\x73\x72\xe2\x16\x3c\xdd\x93\x87\x1f\x5c\x98\xf3\xf7\x15\x4a\xf8\x4e\x22\x2c\x3f\x1a\xcf\x6d\x39\x72\x16\x12\x67\x32\xf8\x36\x90\xac\xb1\xf2\x04\x39\xa0\xf7\xbc\x8a\xfb\x41\xca\x50\xd1\xa8\x72\xd4\x80\x2e\x84\x87\x76\x01\x08\xb5\x77\x14\x34\xed\x54\xa0\x3a\x6d\x93\xd9\x39\x59\x50\x19\xc8\x64\x26\x2a\xf7\x65\x2b\x58\x53\xd9\x93\x38\x6e\x2c\x14\x26\x33\x2b\xf4\xa5\xb0\xa2\x52\xe9\x82\x92\xdc\xa3\xb7\x70\x61\x32\x43\x21\x4d\x2f\xb1\x69\x76\x16\x5b\x48\xa8\xe3\x7b\xb2\x30\x47\x1f\x9f\xc1\x1e\xa3\x96\xcb\xc0\xab\xb0\xa3\xc0\x64\xa0\xae\xa3\x7f\x39\x90\x85\xc9\x55\x47\xe2\x6a\x9c\xfe\x49\xab\xf2\x2b\xcb\x72\x62\xde\xd0\x4d\x93\xc1\x87\x07\xec\x7a\x4f\xa0\xe2\xda\x96\x04\x3a\x6e\x06\x4f\xd1\x64\x26\x83\xbf\x43\x5e\x73\xd7\xa5\x6d\x50\x86\xcd\x92\xde\x72\x58\x46\x93\xad\x71\xb4\x26\x03\x80\x7c\x1b\x60\x81\x43\xe3\xe2\x12\xdb\xf6\x3c\xf2\x48\x01\x70\xb0\x30\x39\x9b\xce\x3e\xb9\x30\xfd\x7c\x7d\xf3\xfe\xcb\x64\x4b\x6c\x8e\xab\xdd\x8e\x00\x1a\x8a\xea\xc2\x78\x78\xe3\x13\x9a\xdc\x2f\x3e\xb9\xb0\x07\xcc\x8e\x91\xeb\xfd\xe1\xef\x07\x0b\x8b\x2a\x84\xdd\xd1\xaa\xce\xa6\x67\xa7\x96\x33\x77\x5e\x49\x2c\x04\x8c\x0d\x7e\x33\x19\x54\xed\x8a\x65\xf9\x42\xd1\x6d\x83\x0f\x73\xa2\x86\xe4\xc9\x3f\x60\x47\x16\xfe\x68\xf0\xe1\x23\x69\xbd\x20\x39\x91\x32\x1e\xbf\x32\x2a\xa6\xab\x30\xf9\xa9\x98\xfd\x96\x17\x97\x79\x31\x83\xa2\xb0\x45\x31\x39\xac\xc2\xa3\xd3\xa7\x90\xc7\x98\xbb\x04\xdf\x0d\x55\xac\xc5\x55\xbb\xa8\xe3\xb0\xf4\x51\x68\x7a\x76\x41\x2d\x78\xae\xd1\x2f\x38\xaa\xbd\x28\x8a\x62\x4f\xa4\xdc\xbb\xda\x42\x85\x12\xcb\x54\xdd\x1e\x89\xaa\xe2\xaa\x41\xa9\x6c\x85\x87\xde\xc2\xd8\xd5\x3d\x49\x5c\x60\x4f\xfb\xb9\x39\xe4\xf0\xa1\xe7\x7a\xb1\x87\x26\xdc\x05\xfd\xf5\x97\x23\xed\xe7\x9e\xc2\x91\x74\xee\x19\x5f\x12\x5f\xbb\xf6\xd8\xf7\x94\xf8\x86\x57\xaf\xd6\xbe\xf7\x3c\x5e\xcf\xd7\xa9\xbf\xb0\x1f\xba\xff\x97\x3f\xf5\xb2\x67\xbf\x6e\x39\x3c\xef\xe6\x63\x3f\x6f\x37\xd4\x33\xfc\xa5\x4e\x02\x60\xef\xca\x25\xad\x2d\xac\x79\x90\x72\x3b\x3a\xd0\xa4\xe7\xb0\x1c\xc4\x5b\x58\xa8\xf6\xd1\x4e\xa7\xd8\xbb\xf3\xc7\x70\xc7\x07\xf2\xb8\xee\x2a\xf6\xf1\x30\x69\x53\xc8\xd5\xd5\xed\xcd\x8b\xc4\xdd\xed\xd7\x83\xea\x2a\xa7\x1d\x9d\xb8\x15\xef\x46\xee\xe3\xc8\xbd\xe1\x5a\xcc\x0e\xae\xc5\x89\xe5\xe6\x70\xfe\xcf\xbb\xbf\x76\xc0\x58\x7f\x7a\x5a\xe7\x32\x86\x4f\xd2\xd3\x31\x31\xff\x05\x00\x00\xff\xff\x38\x2f\x55\x05\x75\x07\x00\x00")

func cmdCreateDefaultYmlBytes() ([]byte, error) {
	return bindataRead(
		_cmdCreateDefaultYml,
		"cmd/create/default.yml",
	)
}

func cmdCreateDefaultYml() (*asset, error) {
	bytes, err := cmdCreateDefaultYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/create/default.yml", size: 1909, mode: os.FileMode(420), modTime: time.Unix(1534302410, 0)}
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
	"cmd/create/default.yml": cmdCreateDefaultYml,
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
	"cmd": &bintree{nil, map[string]*bintree{
		"create": &bintree{nil, map[string]*bintree{
			"default.yml": &bintree{cmdCreateDefaultYml, map[string]*bintree{}},
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

