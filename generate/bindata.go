// Package main Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// tpl/curd.tmpl
package main

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _tplCurdTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x58\x6d\x6f\xdb\xc8\x11\xfe\x5c\xfd\x8a\x39\x22\x08\x48\x81\x59\x5f\xef\x82\x7c\x30\xa0\x02\x8a\xed\x38\x2a\x82\xd8\x8d\x93\xb6\x40\x60\xe4\x56\xe4\x90\xde\x84\xda\x25\x56\x4b\x27\xa9\x8e\xff\xbd\xd8\x17\x8a\xcb\x37\x49\x6d\xae\xf5\x17\x4b\xbb\x33\xcf\xcc\x3e\xf3\xb2\xb3\xda\xed\x20\xc5\x8c\x71\x84\x20\xa9\x64\x1a\x40\x5d\xcf\x76\xbb\xb3\x39\x2c\xd3\x14\xd4\x03\xc2\x9a\x6e\x11\x1e\x90\xa6\x28\x21\x13\xd2\xac\xe5\xc8\x51\x52\x85\x29\x64\xac\x40\x98\x9f\x19\x25\x78\x52\x7e\xc9\xe1\x7c\x61\x55\x9e\x90\x0b\xc1\x33\x96\x93\x5b\x9a\x7c\xa1\x39\x5a\x60\x50\xb8\x29\x0b\xaa\x10\x02\x8b\x19\xc0\x13\xbd\x33\x63\x9b\x52\x48\x05\xe1\x0c\x00\x20\x11\x5c\xd1\x44\x41\x90\x33\xf5\x50\xad\x49\x22\x36\x67\x1b\xfa\x8d\x15\xe2\x6a\xb3\xd9\x6c\xce\x72\xf1\xec\x2b\xae\xcf\x9c\x58\x60\x74\x7c\xd9\x5c\x88\xbc\xc0\xb3\xaa\x62\xe9\xc8\x26\xe3\xcf\x72\xc1\x59\xa2\x3f\xd9\xed\x5c\x7c\x52\x42\x14\x87\xec\xe9\x7d\x87\xb5\x55\x32\x11\xfc\xb1\xfd\xc6\x78\xbe\x0d\x66\xd1\xcc\x70\xb0\x11\x29\x16\x9a\x85\x80\x96\xa5\xa1\x73\x96\x55\x3c\x01\x2b\xb7\x4a\x43\x96\xba\xcf\x91\xfb\x0f\x3b\x83\x24\x51\x55\x92\x03\x4b\x67\x8d\x0a\xe3\xaa\x2b\xcf\xb8\x72\xc2\x8f\xb4\x88\xe1\x93\x36\xe3\xbc\x21\x4b\x25\x58\xc8\xd2\xc8\xc7\x7a\xa4\x85\x0f\xf6\xeb\x2f\x03\xb8\x5f\x7f\xe9\x59\xd7\x4b\x21\xe3\xea\xc5\x73\x23\x1b\x45\x3e\x80\x5b\xf4\x01\x5e\x3c\x9f\xf4\xe8\x96\xca\x2d\xae\xb8\x0a\x59\x1a\xc3\x9f\x7f\x8e\xe1\xc5\xf3\x49\xef\x74\xa8\xba\xd8\x7a\x85\x7c\xf8\xb0\xba\x74\xf8\x95\x43\x37\xeb\x06\xda\x9c\xf6\x4f\x0e\xac\xda\x43\x71\xfc\xba\x2c\x59\x98\x14\x0c\xb9\x82\xf9\x85\xf9\x1f\x83\x28\x15\xcc\x97\x25\xbb\x29\x15\x13\x3c\x32\x9f\x1d\xb4\x58\x7f\xd6\xc8\x4f\x97\x25\xdb\x59\xf1\x73\xb0\xea\xf5\xcc\x08\xb0\xcc\xa8\xff\xb4\x00\xce\x0a\xa7\xe4\xd6\x0b\xe4\xa1\x28\x15\x79\xc5\xb0\x48\xb7\x11\xfc\x05\x7e\xf6\x04\xf4\x1f\xcd\x34\xf8\x32\xd1\x66\xad\xd4\xae\xee\x09\x90\x3b\x54\x76\xcb\xc3\xfa\x18\x5c\x48\xa4\x0a\x83\xfb\xa8\x23\x2e\xd6\x9f\x9d\x04\xb1\x02\xb0\x00\x9a\xcd\xfa\x36\xff\x3b\x93\x1f\xca\xf4\xb0\x49\x2b\x60\x4d\x36\xdb\x16\xbb\xf6\x63\x2b\xd6\x9f\x75\x40\xd4\xf7\x12\x61\x4f\xba\x8e\x6d\x95\x34\x29\x6c\x01\x61\x43\xcb\x8f\x36\xe6\xf7\x1f\xef\xed\x07\x5f\x73\x54\xa7\xb3\xa6\xff\x1c\x11\xfe\x89\xf7\x7b\xce\xe3\xc1\x9e\xf5\xf7\xa2\x93\x26\xad\x65\x4f\xba\x6b\xee\x35\xdd\xc2\x5a\x88\x62\xe2\x0c\x66\xab\x49\xc5\x90\x66\x30\xf7\xa1\xa2\x96\xf5\xcc\xaa\x36\x67\x8e\x1c\x3c\xcd\x88\xb6\xb0\x00\x25\x2b\x6c\x56\x9c\x99\x05\x6c\xe8\x17\x0c\x7b\xd6\x62\x93\x83\x16\x2e\xb2\x71\xd3\x7d\xfa\x53\x0c\x66\x4d\x27\x9f\xa4\x3c\x47\x70\x16\x5b\xd6\xf6\xc8\x1f\xcd\xd6\xbd\x6f\xb5\xd6\x87\xd0\xfd\x8c\x8b\x14\xb7\x1a\x24\x63\x85\x42\xf9\x56\xa4\xba\xb7\xbf\x35\xab\x75\x43\xd7\x15\x4f\x4b\xa1\xbb\x13\xe3\x0a\x65\x46\x13\x74\x66\xde\xb0\xad\x0a\xe7\xae\x51\x93\x6b\xc6\x5f\x63\x51\x5a\x27\x6d\x5c\x26\x36\x6d\x40\x27\x36\xaf\x71\x0a\xf3\x12\x0b\x1c\x55\x6b\x1c\xbd\xa8\x64\xfa\xb2\x62\x85\xbe\xcd\x3a\x61\x5d\x96\x6c\x24\xaf\x76\x3b\xc7\xdd\x13\xae\x29\x70\x64\xd4\xdd\x62\x32\x2c\x91\xb7\x74\xa3\xef\x38\x98\x77\xbe\x2e\x4b\xe6\x83\x21\x4f\x1b\xed\x7a\x9f\x25\x6f\xf1\xab\xe7\x56\x18\xc1\xdc\xf7\xd2\x3a\x93\xac\x4d\x87\xf2\x36\x5c\x45\x1f\xf5\x30\x59\x13\x7d\x36\xd2\xf5\x72\x01\x4f\xfb\x7e\xb6\x80\x9e\x97\xae\x9c\x93\x75\x9b\xd3\xc9\xba\xe3\x60\xf4\x4e\x54\x0a\xc3\x52\x62\xc6\xbe\xb9\xee\x1d\x83\x84\x9c\x71\xb2\x32\x7b\x32\x82\xb9\xfe\x66\xbf\x5c\x4b\x51\x95\xee\x54\x2c\x83\x9f\xdc\xfd\xa9\xb3\xfe\xd6\x60\x84\xc1\x59\x10\x83\xc5\x8b\xbc\x58\x38\x0b\x8b\xe6\xb6\x26\x77\x46\xf3\xaf\x82\xf1\x8e\x8a\x57\xdc\x66\x1a\x91\xc4\x98\x0c\x9b\xed\x93\x69\xb3\x6a\x79\x0c\x81\xa5\xaa\x1d\x64\x82\x78\x9c\xd6\xa8\xc7\xa0\x4f\x61\x3e\xcd\xa0\x33\x64\x88\xe9\x33\x15\x43\x49\xd5\xc3\x9e\x57\x5a\xb2\x7d\xb1\xed\x5b\x46\xc9\x2c\xa9\xe7\x0b\x30\x20\xcd\x81\xa9\x7a\x88\x3a\x12\xe4\xfa\xea\x7d\x18\x68\xef\xbb\xf5\xf1\x9a\xf2\xb4\xc0\x90\x96\x8c\xe8\x82\x8d\xc6\xb4\xce\xce\x59\x7a\x50\xf3\x1a\x07\x8a\xb7\x37\x77\xc7\xec\xd9\x42\x1f\x28\x2e\xdf\x5f\xbc\x3e\xc1\xa6\xed\x21\x7d\xed\xcb\xab\x37\x57\xef\xaf\x4e\x50\xb7\xed\xc2\xce\x38\x07\x33\x42\x47\xd9\xb5\x4f\xb3\xd7\x74\xe4\x26\xc6\x7a\x5f\x62\x82\xec\x11\xa5\x93\x78\xd7\x7c\xf5\x20\x68\xc9\xee\x6c\x8b\x39\x5f\x40\x29\x75\xc3\x6c\x92\x27\x58\x96\x2c\xf0\x65\xcd\x1c\xb9\x52\xb8\x69\x65\x83\x84\xd8\x6b\x8a\x04\x5e\xd2\x59\x0f\x4c\x77\xeb\xda\xa8\xeb\x61\x43\x9b\xfb\x0d\xe9\x95\xe9\xe8\x03\x2d\xbb\xec\xca\x68\x1a\xdc\x69\x4f\xdc\xc5\xb7\xac\x44\xd0\xe9\x1e\x3e\x60\x51\x42\xbf\x23\xc7\x90\xb4\x62\xdd\x96\xe9\x32\xa2\x87\xb7\xcc\xb4\xb1\x43\x80\x4c\x73\x35\x1f\xa9\xc7\xf6\xb2\x39\xea\x54\xd5\x8a\x75\x91\xac\xfe\x0d\xc7\x3e\xe4\x0f\xfa\xa5\xcb\xed\xa8\x57\x45\x23\xd4\x45\xf9\x5b\x85\xf2\x7b\x17\xea\x92\x2a\x7a\xd4\x19\x3d\x74\xf4\xfc\x69\xef\x6d\x6f\x50\xb4\xd5\xf1\x12\x33\x21\x0f\xfb\x77\xf0\x84\xd7\x78\xfc\x80\x39\x1e\x39\x9f\xcb\xc3\xe6\xba\xec\xa7\xe2\xf1\x61\x7f\x50\x19\x6d\xbe\xba\xfe\xfc\xb4\x2f\xb2\x5b\x96\xec\xbc\xfb\x9a\x30\xc0\x51\x3d\x70\x29\x4c\xb4\x85\x56\x5b\x33\x6a\x26\x9f\xd1\xf3\xfa\x97\x9a\x16\x20\xef\x70\x2b\x2a\x99\xe0\x2d\xcd\x31\x34\x3c\x6d\x15\x95\x66\x98\x8a\x61\xcb\xfe\x85\x60\xba\x7d\xe8\x05\x29\x06\xef\x02\xd8\xdf\x90\x9a\xc3\xf3\x45\xaf\x7d\xd4\x35\x31\x4c\x86\xdd\x99\x9e\x65\x90\x10\x5b\xc3\x64\x9f\x85\x83\x17\xce\xfe\x3a\xec\x8b\x9a\xa3\xc5\xc6\x64\x17\xb7\x3b\x1b\x25\x85\xe0\xb6\x9e\x74\x17\x63\x25\x92\x0b\xbd\x12\x46\xb3\xa1\xe7\x4e\xe0\x26\xcb\xb6\xa8\x2c\x05\x11\x79\xc3\x36\x4c\x85\x9a\x84\x9e\xfb\x26\x91\x1b\xd0\x65\x51\xfc\xd3\x78\x44\x96\x65\x79\x21\xb8\xc2\x6f\xaa\x67\xe2\x91\x4a\x48\x75\x79\x78\x2c\xc2\xc2\xe2\x1c\x24\xc6\xd4\xd4\x24\x31\x06\x72\x31\xd4\x70\xfc\x18\xf8\x43\x04\xb9\xec\xd3\x30\x71\xcb\x16\xb9\x10\x15\x57\x23\x47\xda\x83\x0c\xeb\x62\x2c\x09\xdd\x3c\x7c\x34\x0d\x2b\xc9\xcc\x7c\x39\x68\xe6\xfa\x6f\x95\xba\x11\x77\x75\x49\xde\xeb\x9b\xa0\xae\xe1\xb7\x4a\xb2\xf3\x80\xa5\xc1\x6f\xad\x4b\x5e\xeb\x30\x7e\xaf\xf8\xdf\x69\xc1\xd2\x97\x8c\xa7\x1f\x24\x0b\x2b\xc9\xbc\x90\x30\x77\xb3\x0d\x72\xf5\x1a\x87\xe7\x8e\xb5\x83\x64\x95\x46\xb3\xb1\x20\x75\xfa\xd4\x68\xa0\x46\x45\xbd\x08\x45\xb3\x61\x70\x06\x8e\x59\xdd\x1b\x8e\xa1\x51\x21\x57\xdf\x30\x39\x10\xa1\x4e\x65\xbb\x38\x9c\x16\x34\xf7\xf6\x39\x1a\xb4\xb5\x48\xbf\x4f\x47\xed\x96\x7e\x2f\x04\x4d\xbb\xaf\x93\x93\x82\x15\x6a\x60\x2f\x54\x93\x6d\xc5\x39\x1a\x79\x9c\x3d\x6b\x86\x28\xf3\xa3\x47\x33\x33\xf5\x1e\x4c\x7a\xec\x4f\xba\xbf\x61\x98\x47\xef\xef\xbf\x43\x7f\xb9\xf9\x5d\xc2\x0c\x60\xc4\x72\x64\xd6\xf4\x14\x7e\x3f\x52\x8f\xa6\x1d\xdc\xa1\x1a\x53\x30\x07\x23\x8e\x18\x32\x26\x31\x55\xa9\xfa\x60\xfe\x4c\xdf\x4f\x41\x6f\xe6\x39\x9c\x80\xad\xe0\x68\x03\xad\x87\x05\x62\x0f\x44\x1f\xf1\x50\x83\x1b\xfa\x62\xe7\x92\x53\x9c\x31\x92\x13\xc5\x30\x9e\xcd\x61\x2b\x76\x2c\x97\xdd\x23\xff\xd4\x06\xf4\xbf\xec\x3f\x4f\xbb\x0d\xe8\x47\xab\xe7\x0f\xec\x64\x06\x61\x31\x16\x2a\xff\x1c\x61\x20\x1d\xff\x41\x0c\x01\x17\x0a\x32\x51\xf1\x34\xf0\xe2\x05\x58\x6c\x71\x62\x28\xd0\x46\xdc\x7b\xa9\x37\x09\x9c\x58\xb5\xd0\xaf\x5c\x0b\x36\xa8\x5c\xb7\xfc\x1f\x57\x2e\xfc\xc1\xd5\x0b\x83\xbb\xd6\xab\xe2\xde\xa1\xf6\x25\xe1\xbd\x13\x8e\xcf\x42\xad\xf0\x89\xd3\x90\x0d\xf4\xc1\x92\x3e\xec\xd8\x81\xaa\x1e\xf1\x6c\xba\xb2\xbb\xae\xd5\x3f\x5c\xe6\xd7\x78\xc2\xac\xfb\xff\xaf\xf1\x63\x03\x31\xf9\xc7\x03\x4a\x0c\xfb\x3f\xef\x90\xd5\x65\xe8\xca\x74\x7c\xe2\x68\xde\x34\x87\xdb\xab\x93\x9a\x68\xf4\xe3\x94\x9b\xd4\x78\xc5\xe4\x76\x64\xf8\x6b\x42\xd1\xfe\xbe\xb4\xff\xf4\xef\x00\x00\x00\xff\xff\x95\x51\x8d\x7d\xc6\x1b\x00\x00")

func tplCurdTmplBytes() ([]byte, error) {
	return bindataRead(
		_tplCurdTmpl,
		"tpl/curd.tmpl",
	)
}

func tplCurdTmpl() (*asset, error) {
	bytes, err := tplCurdTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tpl/curd.tmpl", size: 7110, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"tpl/curd.tmpl": tplCurdTmpl,
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
	"tpl": &bintree{nil, map[string]*bintree{
		"curd.tmpl": &bintree{tplCurdTmpl, map[string]*bintree{}},
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