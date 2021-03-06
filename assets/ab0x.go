// Code generated by fileb0x at "2017-11-30 13:37:33.936530875 +0800 CST" from config file "ab0x.yaml" DO NOT EDIT.

package assets

import (
	"bytes"
	"compress/gzip"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"golang.org/x/net/context"
	"golang.org/x/net/webdav"
)

var (
	// CTX is a context for webdav vfs
	CTX = context.Background()

	// FS is a virtual memory file system
	FS = webdav.NewMemFS()

	// Handler is used to server files through a http handler
	Handler *webdav.Handler

	// HTTP is the http file system
	HTTP http.FileSystem = new(HTTPFS)
)

// HTTPFS implements http.FileSystem
type HTTPFS struct{}

// FileImagesFaviconIco is "/images/favicon.ico"
var FileImagesFaviconIco = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x9d\x0b\xb0\x1c\x45\xd5\xc7\xff\x73\xf7\x23\x9b\xdc\x90\x6c\x91\xcf\x1b\x48\x48\xb8\x31\x09\x31\xa0\x18\x05\x25\xf1\x81\x17\x41\xc4\x07\xc6\x52\x83\x28\x1a\x50\xaa\x0c\x8a\x2f\x54\x24\x85\xc6\xdc\x0d\x2f\x0b\x14\xb4\x4a\x4b\xa0\x7c\xa0\x49\x14\x91\x58\x28\x60\x69\x29\x14\xab\x05\x12\xac\x12\x24\x4a\x88\x26\xea\x26\x11\x50\x82\xa9\x35\x7a\x75\x93\xbb\x99\xb1\x7a\xe7\xcc\xde\x9e\xf5\xee\xee\xf4\xe9\x99\xe9\xdd\xd9\xfe\x55\x9d\xa2\x3a\xdc\x9d\x3e\xa7\xcf\x4c\x4f\x4f\xf7\xe9\xd3\x80\x03\x07\xc5\xa2\xf8\xef\x02\x2c\x9d\xea\x60\x36\x80\xa5\x00\x8a\x80\xf8\xc7\xfa\xbf\xfb\x38\x98\x35\x1d\x75\x09\xf0\x3c\xcf\x8a\x15\x2b\x56\xac\x58\xb1\x62\xc5\x8a\x15\x2b\x7d\x22\x16\x8b\xc5\x62\xb1\x24\xc8\x5c\x00\x1f\x02\xf0\x7d\x00\x3b\x00\xec\x07\x70\x00\xc0\x2e\x00\x25\x00\xd7\x01\x38\xcd\x9f\xad\xb2\x64\x88\x25\x00\xbe\x0b\x60\x1c\x80\x17\x41\x9e\x00\xf0\x0e\x7b\x1f\xf4\x3c\xc2\x7f\x97\x03\xa8\x46\xf4\x7b\xb3\xfc\x14\xc0\x31\xa6\x8d\xb0\xb0\x98\x02\xe0\x36\xa6\xdf\x65\x79\x12\xc0\x0a\xd3\xc6\x58\x94\x10\xcf\xfd\xe6\x18\x7c\x1f\xc8\x7f\x00\x5c\x68\xda\x28\x4b\x64\x8a\x31\xfa\x5e\x96\x1b\x01\xe4\x4c\x1b\x67\x69\xcb\x2a\x00\x6e\x42\xfe\x0f\xc6\x04\xb3\x4c\x1b\x69\x99\x94\x17\x03\x18\x4b\xd0\xf7\x81\xec\x04\x70\xa2\x69\x63\x2d\x21\x66\x03\xd8\x9b\x82\xef\x03\x39\x00\x60\xa5\x69\xa3\x2d\x75\xc4\x58\xff\xa1\x14\x7d\x1f\xc8\x61\x00\xeb\xec\x3c\x81\x71\xbe\x69\xc0\xf7\xb2\xdc\x0e\x60\xba\xe9\x46\xe8\x53\x2e\x63\xf8\xeb\x6f\x00\xae\x01\xf0\x6a\x00\xc7\x03\x58\x0e\xe0\xc3\x00\x1e\xd5\xb8\x07\x7e\x03\x60\xd8\x74\x63\xf4\x19\x6f\x00\x50\x53\xf4\xd3\xb7\x01\xcc\x6c\x71\x3d\xd1\x8f\x5f\x42\xdf\xfb\x9c\x7b\xe0\x19\x00\x23\x29\xb7\x41\xbf\x72\x02\x80\x8a\xa2\x7f\x36\x47\x7c\x57\x2f\xa7\x79\x3f\xce\x3d\x70\x08\xc0\x19\x29\xd8\xdf\xcf\xcc\xa2\x6f\x30\x15\xbf\x3c\x48\xe3\xc4\xa8\xcc\x65\x8e\x29\x1f\x06\x30\x2d\x41\xdb\xfb\x9d\x1c\x80\xfb\x14\x7d\xb2\x1b\xc0\x10\xa3\xae\xa9\x8a\x63\x4b\xd1\x67\xcc\x4b\xc0\x66\xcb\x04\x5f\x56\xf4\xfd\x18\x80\x93\x34\xeb\xfc\x78\x84\x71\x46\x95\xde\x1b\x96\xe4\xb8\x58\xd1\xf7\x2e\x80\x37\xc7\x54\xf7\x59\x14\x2f\xd2\xaa\x2e\xbb\x3e\x94\x2c\xa7\x03\x38\xa8\xe8\xff\x2b\x62\xd6\x61\x31\x80\xc7\x27\xa9\xe7\x86\x98\xeb\xb1\x84\x59\x44\xdf\x56\x49\x8c\xf5\x55\x99\x01\xe0\x2e\xa9\x9e\x9f\xd8\x35\xc1\x44\x11\xdf\xea\xdb\x14\x7d\xbf\x35\xe1\x31\xf8\x00\x80\xab\x01\xfc\x01\x40\x21\xc1\x7a\xfa\x1d\xf1\x5c\xfd\x50\xd1\xf7\x7b\x01\xcc\x49\x49\xbf\xc1\x94\xea\xe9\x57\x36\x30\xc6\xfa\x27\x9b\x56\xda\x12\x0b\x2f\x51\x9c\xdb\x15\x63\xfd\x73\x4d\x2b\x6d\x89\x8d\x9f\x2b\x3e\xfb\xa3\xa6\x15\xb6\xc4\xc6\xa9\x8a\xbe\xbf\xdd\xae\xc1\x67\x8a\xcf\x2b\xf8\xfe\xc9\x1e\x1b\x83\x8b\x31\xe3\xeb\x69\x6e\xe2\x66\x5a\x8f\xfc\x12\xc5\x91\xbc\xd6\xc6\x11\xd4\x51\x59\x87\xff\xa4\x69\x65\x23\xf2\x02\x00\x1b\x01\xfc\xbb\x83\x3d\x07\x69\x7f\x5a\xbf\xce\x25\xff\x1f\xc5\x55\x45\xf5\xff\xf1\xa6\x15\xee\x80\x78\x9e\x6f\x52\xb4\x29\x90\x4d\x00\x8e\x32\x6d\x40\xca\x1c\xa3\xd0\x3e\x87\xbb\xfc\xbd\xff\xdc\x16\xf3\xc5\x2a\x52\xa6\xb8\xe6\x7e\xe1\x79\x8a\xed\x33\xd5\xb4\xc2\x2d\x98\x17\x63\x3c\xf2\x18\x80\xf3\x4c\x1b\x94\x12\xc7\x2a\xb6\xcd\x0b\x4d\x2b\x3c\x09\x05\xc6\x9c\x75\x27\x71\x01\x5c\x4b\x73\xcf\x59\x66\x20\xc2\x18\x49\x96\x6b\x4d\x2b\xdc\x44\x0e\xc0\xdd\x31\xfb\x5e\x96\xbb\xdb\xc4\x2f\x66\x85\x07\x15\xfb\xc6\x85\xa6\x15\x96\xb8\x3e\x41\xdf\x07\xb2\xbd\x07\xc6\xbd\x3a\x7c\x5a\xb1\x3d\x44\x5f\xfb\xff\xa6\x95\x06\xb0\x3a\x05\xdf\x07\xb2\x1f\xc0\xd9\xa6\x0d\x4e\x88\xf9\x0a\xb9\x3a\x02\xd9\x65\x78\x4f\xde\x72\x46\x7c\x8a\x90\xef\x68\xc4\x1a\xd7\x28\x36\x2d\x8b\xdc\xca\x68\x0f\x53\x7b\xf2\xc4\x58\xff\x69\x86\xbe\x3f\xa6\xf1\x02\x37\xd6\x38\x90\x8d\x5d\xfc\x1d\xc4\x45\xb4\xc9\xdf\x19\x6d\x91\xf6\x9e\xbc\x41\x00\xbf\x66\xe8\xb9\xa3\x69\xde\x3a\xaf\xb9\x8f\xed\x61\xfa\x76\xca\x12\xe7\x30\xf6\xf7\x04\x92\xc6\x9e\x3c\x87\xea\x51\xd5\xad\x42\xb1\x84\x93\xf1\x31\x0d\x9b\xb3\x98\xab\x66\x8d\x46\x3e\x87\xa4\xf7\xe4\xad\x67\xe8\x54\xa3\x75\x9e\x76\x9c\xc5\xec\xfb\x3c\x8a\x43\x7f\x4f\x82\x36\x9b\xe0\xfc\x2e\xdc\x93\xf7\x36\xe6\x7d\x79\x69\xc4\xeb\xb7\x8a\x35\x8e\x2a\x5f\xa0\xb5\x94\xac\xa0\xbb\x27\xef\xfd\x31\xea\xb2\x8c\x99\x67\xe4\xeb\x8a\xf5\xcc\x60\xc4\x3f\xca\x92\xb5\x5c\x35\xba\xe3\xe4\x5b\x00\x1c\xa1\xa9\xc3\x10\x73\x5e\xff\x01\xc5\xfd\x87\x01\x03\x00\xae\xd2\x78\x07\x8a\xef\xe2\xe7\x6b\xda\xdc\x4d\xa8\xee\xc9\x6b\x96\x5f\x50\xae\x18\x0e\x53\x14\xe7\x26\x03\xd9\xc3\xdc\x7f\x28\xf3\x76\x00\xff\x62\xda\x7c\x20\xc6\x7d\x50\xdd\x82\xce\x38\x99\xbb\xa6\xfa\x09\x46\x5d\x63\xf4\xbe\x88\x83\x17\x91\xee\x1c\x9b\xdd\x0c\xe6\xaa\xd1\x19\x27\x8f\xd1\x33\xa5\x42\x4e\x31\x36\xcd\xa5\x71\x62\x9c\x0c\x31\x62\x63\x65\xf9\x5e\xc6\x62\xcc\x74\xc6\xc9\xdc\x35\xd5\x77\x47\xfc\x1e\x49\x2a\x26\xf9\x08\x8a\x29\xe2\xde\x03\xe2\xbb\x78\x41\x42\xba\x99\x40\x77\x9c\x7c\x17\x63\x4d\xf5\xa5\x00\xfe\xd2\xe6\x9a\x77\xa4\xd0\xd7\x5e\x4c\xdf\x36\x1c\x9b\xb3\x96\xab\x46\x77\x9c\xbc\xbd\xcd\x9c\x5c\x2b\xe6\x00\xf8\xe5\x24\xd7\x7a\x34\xc5\xbd\x61\xa7\x51\x3e\x2b\x8e\xcd\xe2\xde\xf9\x40\x4a\x7a\xa6\x85\xce\x38\x79\x17\x63\x0e\x3d\x0f\xe0\x1b\x4d\xcf\x55\xda\x39\x40\x16\x30\xf2\xe1\xc8\x72\x0b\xf3\xdb\xb4\x5b\x59\xa6\x31\x4e\xe6\xce\xa1\x7f\x84\xc6\x94\x2f\x4f\xc0\x9e\x28\x2c\xd1\xcc\x7d\xab\xf3\x5d\xdc\x8d\x0c\xd1\x99\x2e\x9c\xb6\xa8\x02\x78\x2f\xa3\x4e\xd3\xed\xb7\x01\xc0\x63\x1a\xe7\x5d\xec\xc9\x58\xac\xb1\xee\x38\xb9\xd7\xe6\xd0\x8f\xa6\x35\x65\xd1\x07\x3d\xc5\xb4\x79\x8c\xce\xbe\xc9\x12\x6b\x34\xc6\xc9\xf7\xf6\xe8\x1c\xba\x18\xc7\xfc\x8a\x69\x73\x16\x63\x8d\x75\xc6\xc9\xbb\x68\xff\x56\xaf\x31\x95\xf6\x10\x71\xfb\xbf\xac\xc5\x1a\x1f\x07\xe0\x11\x66\x5b\xfc\xb3\x87\xe7\xd0\xa3\xe4\xaf\x6b\x25\x59\x8b\x35\x1e\xa4\xb3\xe0\xb8\xfd\xe2\xfa\x1e\x9d\x43\x3f\xbb\x43\xfe\xba\x76\x92\xb5\x58\x63\x87\xf6\x5f\x73\xf6\x64\x0a\xd9\x02\xe0\x48\xd3\x46\x30\xd0\x99\x2b\xaf\xd1\xfa\x57\x96\x38\x07\xc0\x3f\x98\xed\xf1\x18\xed\xef\xec\x35\x66\x36\xe5\xaf\x53\x95\x4d\x19\x8b\x35\x3e\x81\xf2\xb8\x71\xda\x62\x1f\x9d\x21\x60\x1a\xd5\x31\x5a\x90\xbf\x8e\x3b\x57\x9e\xb5\x58\xe3\xa3\x28\x8f\x23\xa7\x2d\xc6\xe9\x4c\x61\x53\x2c\x04\xf0\x57\x00\x9f\x62\x8c\x4b\x74\xe6\xca\x9f\xca\x58\xac\xb1\xea\xda\x7e\xb3\x7c\x95\xd6\x02\xd2\xa4\x79\x9f\x31\x27\xe6\x5d\x67\xae\x9c\x3b\x4f\xda\xcd\x44\x5d\xdb\x9f\x4c\xfe\x48\x79\x8a\xd3\xa0\x55\x6e\x4c\x4e\xcc\xfb\x3c\xfa\x1d\xf7\xde\xef\xb5\x79\xd2\x4e\x9c\xda\x61\x6d\xbf\x9d\x1c\x02\xf0\xc1\x14\x74\xbc\xae\x8d\x0e\x9c\xb5\xfd\xe9\xcc\xbd\x2c\x81\xfc\xac\x47\xe7\x49\x5b\x31\x27\x86\x58\xe3\xa4\xd6\x54\x2f\x88\x78\x1f\x5e\xa2\x78\x5d\x87\x62\x04\xb9\xdf\xc5\x59\x8b\x35\x6e\x5e\xdb\x57\x95\x07\x12\x58\x13\x5c\xa1\xb8\xcf\x78\x3d\xa3\x8e\x95\x14\x33\xcc\xb1\x39\x8b\xb1\xc6\x97\x32\xf6\xa3\x07\xb2\x27\xc6\x9c\xc4\xf3\x14\xf3\xe0\x57\xe8\xfb\x96\xc3\x89\x1a\x31\x25\x59\x8c\x35\x7e\x8d\x66\xac\xb1\xee\x9a\xea\x20\xcd\x39\x45\xad\xb3\x46\xe7\xe3\xe9\x30\x8b\xf6\x10\x71\xfb\xbf\x3b\x32\x16\x6b\xbc\x08\xc0\x6f\x35\x9e\x89\xcf\x32\xd7\x54\xc5\x73\x74\xa7\x62\x7d\x97\xc5\x64\x73\x8e\xc6\xf7\xdc\x7b\x20\x8b\xb1\xc6\x3f\xd0\x68\x8f\x7b\x18\xf9\x6a\xaf\x51\xac\xe3\x5b\x09\xd8\x7d\xa1\x46\x5c\xd1\xbe\x8c\xc5\x1a\x8b\xe7\xf1\x4a\xcd\x58\xe3\x25\x11\xeb\x7a\x97\x62\x3d\x5b\x13\xfc\xee\x58\xa1\xb9\x07\xf7\x82\x84\xf4\x32\xc5\x2a\x8a\x0b\xe0\xb4\xc7\x7e\x00\x6f\xea\x70\xfd\xd5\x8a\x63\xfd\xbd\x94\x3f\x35\x49\xe6\xd2\xfc\x3f\xc7\xe6\x5a\x06\xcf\xcc\x5f\x06\xe0\xcf\x1a\xef\x83\x2d\x00\xce\x94\xd6\xd4\xc4\x78\xe9\x75\x8c\xf5\x08\x31\xc6\x3c\x25\x25\x9b\xa7\xd2\x3b\x86\x63\xef\x58\x8f\xc6\x52\xb5\xe3\x39\x00\xee\xd7\xb8\x07\x02\xe1\xae\x45\xbb\x86\xf2\xc3\x72\xe3\x8a\xb6\x66\xec\xdb\x10\x14\x6b\xac\x7a\x36\x69\x5c\x72\x95\x41\xbb\xb9\x71\x45\x6f\x35\xa8\x73\x92\xac\x61\xe6\x01\xe4\xca\x9d\x5d\xf0\x2c\x2d\xa6\xb9\x5f\x15\xbd\xef\x35\xac\x73\x92\xbc\x42\x23\xd6\x58\x45\xb6\x75\xd1\xfc\xca\x71\x8a\x39\x10\x0f\x77\x49\xce\xd6\xa4\x98\xcf\xcc\x0b\x18\x55\x9e\xe9\xc2\x79\x95\xf3\xed\x3b\x20\xc4\x34\x00\xb7\x25\xe0\x7b\xf1\x7e\x79\x95\x69\xe3\x26\x21\xa7\xd8\x07\xac\x33\xad\x70\x0a\x88\x77\xf3\x5a\x8d\x35\xd5\xc9\xe4\x7d\xa6\x8d\x6a\xc3\x66\x05\x3b\xbe\x68\x5a\xd9\x14\x79\x23\xad\xc7\xe9\xfa\xfe\x2b\xa6\x0d\xe9\xc0\xe7\xac\xff\x5b\xb2\x14\xc0\xef\x35\x7c\xff\x04\x5d\xa3\x9b\x51\x89\x97\xf8\x8c\x69\x65\x0d\x30\x8d\xe2\x85\x55\xe6\x0c\xf7\x01\xb8\x9c\xf6\x30\x76\x3b\xdb\x15\xec\x5a\x6d\x5a\x59\x43\x38\xb4\x7f\x48\xbc\x13\xbe\x46\x7d\x42\xf3\x1a\x8f\x78\x57\xfc\x08\xc0\x45\xb4\xf6\xdf\x0b\xeb\x67\xaf\x54\xec\xcf\xb2\xb4\xaf\x90\x43\x4e\x9a\xbb\x19\xa4\x6f\xc6\x85\x93\xec\xe7\x78\x99\x81\xd8\x72\x55\x06\x15\x63\x88\xb7\x99\x56\xb8\x4b\x70\x3a\xcc\xdf\x2d\xa5\xbc\x0e\xdd\x8c\x43\xe7\x98\x64\xe5\x1b\x26\x6d\x06\x5a\xc4\xcf\x1f\xab\x10\x27\x60\x92\x75\x8a\xbe\xdf\x9d\xb1\x7c\x53\x71\xe0\x34\xe5\x8c\x2b\x30\xd6\x73\x4d\xbc\x23\xde\xa2\x18\x9f\xe2\xc6\x10\x97\x98\x65\x8e\xa4\xf5\x75\xd5\xf1\xde\x99\x14\x6b\x9c\xe6\x9e\x3c\x4e\x1e\xfc\xeb\x53\xd4\xaf\x57\x39\x9a\xc6\x86\x51\x59\x0c\xe0\x59\x6a\xdf\x2a\xc5\xea\x25\xcd\x6c\x46\x1e\xfc\x7b\x14\xed\xb2\x74\xa6\xd0\x22\x26\xf9\xc6\x04\xdb\x7a\x0a\xc5\x70\xa8\xf8\xfe\x71\x46\xdc\xab\xa5\x3d\x03\x1d\xf2\x1e\x27\x75\xfe\xc7\x46\x45\xdf\x3f\xcb\xc8\xad\x6b\xe9\xcc\xd5\x11\xda\x7e\x67\xcc\xe7\x62\xae\x55\xf4\xfd\x38\x80\x33\x62\xac\xdf\xe2\x73\x9e\xc2\xb8\xfb\x40\x84\x58\xe3\x28\xac\x64\xc4\xfb\xa9\xee\x41\xb5\x74\xe6\x64\xc6\xb8\xfb\x30\x9d\xab\xcc\x8d\x15\x3b\x89\x11\x9f\x7a\x53\xcc\x76\x5b\x7c\x74\xf6\x24\x73\x72\x84\x0c\x01\xf8\x93\x62\x3d\xf7\xd9\x39\x9e\xc4\x10\xe3\xfa\x1b\x34\xee\x81\x47\x14\x72\x84\x4c\x61\x9c\x37\xb3\x33\x86\xf3\xae\x2c\x9d\x59\xad\x79\x2e\x66\xa7\x18\xb2\x29\x8c\x73\x55\x2a\xf4\xae\xb0\xa4\x83\x4e\xae\x9a\x71\x8a\xd7\x69\xce\x51\x31\x40\x67\xd4\x6e\x53\xbc\x5e\x2d\xa6\x71\xa6\x45\x0d\xdd\x73\x31\x0f\xd1\x7c\xce\x16\xda\x77\xc6\xcd\x1d\xbf\xd6\x74\x43\xf4\x31\x79\x00\xb7\x6a\xdc\x03\xba\xb2\xc9\x74\x03\x58\xea\xe8\xe4\xaa\xe1\xca\xd6\x8c\xe5\x84\xed\x75\x74\xce\xc5\x54\x95\xbd\x94\x2b\xcd\xd2\x5d\x2c\x02\xf0\xbb\x84\x7d\x9f\xe6\x9e\x73\x8b\x3a\xba\xb9\x6a\xda\x89\xa9\x3d\xe7\x16\x35\x1c\xcd\x73\x31\x5b\xc9\x95\xa6\x0d\xb3\x28\xb1\x4a\x23\xd7\x77\xb3\xdc\xdc\x05\x7b\xce\x2d\xea\x2c\x03\xb0\x43\xb3\xcf\xdf\x60\x7d\xdf\xd3\x4c\xa3\x7c\x7e\xaa\xe7\xe1\x6d\xb7\xeb\xf8\x99\x62\x3e\xcd\xfd\xee\x6e\xe3\xf3\x71\xca\xd3\xf1\xce\x8c\xe5\x77\xb7\x84\x59\x0c\xe0\x5c\x00\x1f\xa5\xbc\xa1\x17\xd1\x39\x06\x33\x4c\x2b\x66\xb1\x58\x2c\x96\x6c\xe3\x79\x9e\x15\x2b\x56\xac\x58\xb1\x62\xc5\x8a\x95\xfe\x92\x6e\xa3\x58\x08\x15\x5d\xe4\x42\xe5\x2a\x9c\x50\xb9\x82\xb0\x11\x65\x60\xb4\xa9\x3c\x2c\x97\x4b\x40\xa8\x82\x22\x90\x97\xcb\x40\xa8\x42\x97\x46\xc9\xa1\xb2\x54\x41\x4d\x94\x47\x64\xf5\x0a\x45\xb9\xc2\x0a\x86\xcb\x72\x85\x65\x8c\x54\xe4\x72\x09\xa3\x55\xf9\xef\x4b\xf0\x5c\xf9\x7a\x45\xc7\xf3\xe4\xfa\x84\x72\x45\x59\x9d\x9c\xe7\x95\x64\x75\xf2\x9e\xb7\x3b\xa4\x8e\xe7\x8d\xcb\xe5\x50\x63\x78\x15\xb9\xae\x7a\x39\xd4\x78\x5e\xf9\x7f\xca\xbe\x96\xc1\x55\xca\xbe\x33\x1a\x4d\x58\x72\x7c\x2d\x9c\x50\xb9\xda\x68\x42\xbf\x5c\x69\x34\xa1\x5f\x9e\xf0\x91\x5f\x9e\xf0\x91\x5f\x5f\xa9\xe1\x23\x5f\x9f\x62\xc3\x47\xe5\xba\xfe\x40\x50\x61\x45\x5c\x48\xf2\x51\x45\xfc\xb0\x86\x5c\xe0\xa3\xaa\xf8\x61\x15\xf9\x12\x55\x58\xd7\xb4\x8a\x42\xe0\x23\x57\x68\x5a\xc1\x70\x25\xa8\x40\x5c\xa8\x8c\x91\xc6\x5d\x27\x9c\x5b\xc2\x48\x2d\xb0\xa8\x84\x82\x57\xc4\xa8\x27\x57\x58\x84\xff\x77\x41\x85\xc2\x87\x25\xaa\x40\xfc\xb0\x9c\x13\x8a\xe7\x83\x0a\x47\x2b\x39\x52\xcc\xaf\x70\xa4\x9a\x97\x9a\xb0\x84\xe1\x5a\xc1\xff\x3b\xb2\xb8\xe0\x0d\xfb\x7f\xd7\xa8\xf0\x7e\x4f\xae\xd0\x21\x45\x25\x8b\xea\x7f\x27\x59\x24\x3f\x47\x45\xff\x42\xae\x6c\x91\xf4\x77\x0d\x4d\x8b\x72\x13\x7a\xb2\x45\xe2\x42\xae\x6c\x91\xf8\xd7\x11\xd9\x22\xcf\xab\x0e\xcb\x16\x79\x5e\xa5\x10\xb2\xc8\xab\xe4\x25\x8b\x44\xfb\xe4\xc2\x16\x95\x1c\x2f\xe4\xa3\xa2\x13\xb6\x08\x39\xd9\xa2\x61\x17\x79\x4f\xf6\x91\x2b\x7e\x2b\x59\x54\x13\x65\xc9\xa2\xfa\x73\x31\x71\x53\xe0\x74\xf1\xbf\xa4\x0a\xeb\x3f\xad\x06\x77\x21\xf5\x09\xb5\xa0\x5c\xf2\xcb\xee\xc4\x5d\xea\xab\x12\x28\x58\xa1\xfb\xb5\x48\xe5\x2a\xdd\xcf\x41\xb9\x46\xf7\x7b\x29\xe8\x79\xe8\x79\x68\x94\x8b\x4d\xe5\x92\xaf\x79\xa3\x4c\x8f\x69\x70\x3d\xef\xe0\xcc\x7a\xb9\xf1\x58\xd6\x6a\x9e\xac\x9f\xe7\x5e\x11\x2e\x7b\xa7\x88\x7b\xa8\x36\xd1\x31\x3d\x24\xee\x31\xa9\xb3\x78\xda\x2d\xd0\xc3\x15\x5c\x31\x17\x3c\x7c\x64\x92\x13\xee\x2c\xaa\xa8\xf7\x65\x0d\x44\xc7\x52\x94\x7a\xee\x32\x86\x3d\xb9\xa3\x2d\x0b\x1f\x85\xca\xf9\x9a\xdc\x31\xd7\x7d\x24\xf5\x9b\x35\xa0\x2c\xf7\x65\x6e\x53\x3f\xec\x35\xf5\xd3\x75\x1f\xc9\xe5\x72\xb8\x5f\x17\x3e\x0a\xbd\x58\xaa\xe1\xf7\x80\xa8\x30\xf4\x9e\xf0\x9a\xcb\xc5\xf0\x7b\x25\xf0\x91\x54\x41\xb8\x2b\x75\xc3\xef\x35\x33\xfc\x37\x00\x00\xff\xff\xf1\xf1\xf6\x0b\x3e\x08\x01\x00")

// FileImagesLogoPng is "/images/logo.png"
var FileImagesLogoPng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x57\x09\x34\x94\xeb\xff\x1f\x4b\xb6\x46\x51\xa2\x42\xc6\x28\x4a\x97\x59\x18\xc3\x18\x23\x33\x63\x32\xc9\x36\x88\x6c\x19\x66\x86\x31\xcd\x62\x66\x30\x64\x2f\x5b\x57\x84\x84\x70\x95\x6e\x5d\x17\xc9\xed\x97\xa5\xc4\x64\x44\x94\x50\x84\x2c\x59\x4a\xb9\x65\x19\x5b\xd6\xab\xff\xd1\xbd\xdd\xff\x72\x7e\xe7\xfe\xcf\xf9\x7d\xcf\xfb\x9e\xf3\xbe\xdf\xf3\xf9\x7c\xbe\xcb\xf3\x7d\xce\x79\x9e\x14\x47\xfb\xe3\x8a\x0a\xfb\x15\x00\x00\x80\x22\xd1\x06\x4f\x02\x00\x00\xb1\x5b\xaf\x9c\x0c\x00\x00\x68\xc2\x90\x9f\x03\x00\x80\x7d\x7c\x6b\x77\xbe\x33\x9b\xc6\x0f\x23\x73\xa9\x00\x2b\x0a\xdb\x8f\x0a\x22\x32\xc9\x01\x54\x12\x95\x4c\x09\x0f\x6e\xa3\xa2\x01\x00\x29\x1a\xdd\xc5\x9d\xef\x6e\x77\x12\xe5\xcf\x66\x1a\x92\xb7\x30\x86\x02\x26\x07\xb0\x65\x68\x4b\x01\x87\xec\xcf\xa0\xf2\x41\x7e\xd4\x00\x3a\xcb\x02\x3c\x5b\x27\x04\x83\xe8\x14\x0b\xb0\x1b\xc2\x0e\x6a\xc7\xc1\x51\x03\xe9\x36\x11\x5c\xaa\x73\x84\xbd\x8b\x7f\x04\xc3\xdf\x8c\x02\xb6\xc4\x80\xd0\x02\x94\x80\xc9\x61\x52\xf9\x64\x90\x80\x79\x96\xc5\x43\x09\x2c\xc0\xdf\x74\x51\x2c\x1e\x6a\xcb\x0d\x01\x83\xbe\x41\xf8\x0c\x0b\xf0\x9f\x49\xb9\xdb\x39\x82\x70\x6c\x2e\x15\x84\x30\x34\x32\xf0\x87\xc2\x60\x20\x13\x13\x43\x98\x31\xc2\xc4\x04\xf6\x03\x08\x0e\x85\xc1\x21\x50\x38\x04\x6a\x62\x00\x33\x46\x21\x4c\x50\x70\x24\xe8\x2f\x03\x63\x40\x68\x2e\x85\x86\x22\xe1\x09\x7f\xc5\xe2\x52\x68\x16\xe0\x40\x3e\x9f\x83\x82\x40\xc2\xc2\xc2\x0c\xc3\x8c\x0c\xd9\xdc\x00\x08\xcc\xcc\xcc\x6c\x4b\x03\x0e\x37\xe0\x52\x68\x06\xbc\x70\x16\x9f\x2c\x30\x60\xf1\x74\xbe\x2b\xe0\xa9\x3c\x7f\x2e\x9d\xc3\xa7\xb3\x59\xa0\xad\x7f\xb2\x1f\x3b\x84\x6f\x01\x06\x7f\x2f\x81\xc9\xb1\xb3\xfb\x5b\x98\xc5\xfb\xab\x51\xfe\x6c\x26\x44\x40\xe6\x40\x60\x86\x50\x08\x93\x09\xf9\x8e\xe6\xf1\x49\x54\xda\x3f\xa3\x79\x2e\xe1\x1c\x2a\x84\x44\xe5\xb1\x43\xb8\xfe\x54\x12\x95\xa6\xf3\x3f\x42\xfd\x33\x75\x0b\xc8\xb1\xb3\x43\x39\x70\xe9\x01\x74\x16\xf9\x2c\x9e\xed\x1f\xc2\xa4\xb2\xf8\x44\xbc\x05\x58\xc0\xe4\x18\x52\xe8\x14\x94\x11\x02\x61\x65\x6d\x8c\x43\x20\x90\x38\x82\x35\x1c\x06\xc3\xc2\x4c\x08\x04\x53\x63\xac\x95\x31\xc2\xc8\x08\x4f\xc0\x7e\xd7\xf8\x77\x5c\x24\x1e\x0b\x83\x5b\xe1\xa1\x38\x82\x09\x14\x06\xb3\x86\x5b\x11\x70\x46\x04\x63\x84\x99\x15\x16\x6e\x04\x83\x43\xf1\xdf\xb9\x44\x16\x8f\x4f\x66\xf9\x53\xbf\x73\xe9\x7f\x73\x71\x84\x7f\xe4\xa2\x70\x5c\x2a\x99\xcf\xe6\xba\xb0\xd9\x67\xbf\x4f\x80\x63\x20\x9b\xcf\xe6\x05\xb2\x39\x20\x9c\xb3\x09\xe8\xb0\x1b\x9d\x45\x61\x87\xf1\x8e\x6c\x2d\xcf\x5f\x99\x52\xb9\xf4\x50\x2a\x85\xc0\x65\x33\x41\xdf\xfa\x8b\xa2\xff\x9b\xf8\xff\x7f\xdd\x7f\x72\x29\xff\x49\xcf\x20\x18\x10\x1a\xf2\x7f\x86\xe5\xbb\x8b\x84\x27\x6c\x7d\xfe\x3d\xfa\x18\xd0\x7f\x6f\x1e\x2a\x8b\x62\x01\xe6\x82\x2d\x31\x97\x7b\xee\xde\x02\x00\xe4\x3b\x89\x78\x2b\x17\xc1\xc0\x94\x77\xb8\xf8\xb5\xb3\x8a\xa5\xf0\x0a\x9e\x3e\xe5\x74\x45\x09\x7b\x44\xe5\xf4\xbd\x3d\x59\x4c\x17\x3c\x25\xdf\xe9\x0b\xf1\xd0\x79\xfc\x76\x42\x49\x4c\xfb\x93\x7e\x9c\x76\xa2\xaa\xf3\xee\x16\xf8\x55\xb8\xbe\xe7\xf6\x84\xf9\x7f\x65\xd1\x95\x4f\xcf\x0f\x75\x3e\x5a\xa7\xf9\x74\xce\x2e\x47\xcf\x16\xac\xd7\x8f\x4c\x34\x2c\xe4\xad\x6e\x2c\xaf\xd7\xcf\xd6\x2f\x0f\x4d\x14\xf0\x3a\xcd\x6a\x74\x73\x9e\x67\x14\x47\xbc\xb4\x49\x3a\xa6\x17\x8b\xfd\xf3\xb9\xe1\x78\x58\xe5\x62\x10\xf6\xd2\x0e\xab\xe2\xfe\xbe\xca\x1f\xc7\x0e\x24\xda\x01\x63\x10\x1b\xaf\x38\x37\xc6\x64\x34\x7f\x82\xfe\xf6\xcb\x98\x45\x62\x99\x22\x6f\x49\xcc\x6e\x72\x53\xec\x44\x9e\x2d\xc7\xa9\xca\x5c\xdb\x67\x0f\xdf\x28\x4b\x1d\xfd\x3c\x4e\x6a\xda\x1d\xb2\x72\xef\x52\x47\xe2\xf3\xbc\x78\xc8\x07\x61\xa2\xff\x62\x6b\x87\xee\x9a\xfa\xe6\x07\xd0\x86\x1f\xe6\x7a\x16\xc8\x53\xe1\x8b\x69\x5d\x04\x4f\x99\x5b\xbe\x83\xb9\xf6\x52\x24\xbe\x95\xbc\xed\x4c\xa4\x06\xb0\x4b\xdb\x1e\xf2\x05\xb6\xd1\x93\x25\xde\x3e\x14\x28\xc8\xdd\x11\xb2\xd6\xf5\x9c\x73\x29\x60\x4c\xcd\x9d\x51\x5b\xb7\x23\x3c\x02\xda\x33\x30\xfb\x60\xaf\xd9\x08\x47\x5d\x73\x8d\xf2\xb2\x47\xfc\x5b\xee\xb6\x33\xed\x7a\xde\xfb\x85\xd9\xcf\x3a\xe5\x39\x63\x1d\x30\xb4\xf7\xca\xd7\x7e\x18\x06\xea\xe5\x5b\xb3\x38\x5d\x2e\xd9\x66\xa5\xd7\x64\x97\x38\x5b\x62\x3a\x6e\x8d\x6c\x62\x5f\x28\x53\x8d\x4f\x32\x17\xf4\x65\x8e\xe0\x23\x09\xec\xa6\xf4\x3c\xe1\x72\xfa\xd8\x3e\xfd\xcc\x15\x68\xab\x78\xe7\x0d\x71\x0c\x1c\x5d\x11\x37\x32\x32\xc0\x91\x59\x08\xeb\xdb\x7f\xa0\xb8\x78\x6e\xbb\x46\xde\xb2\x59\x41\xdf\x8f\xc0\xc1\x53\x2b\x06\x27\xe3\x91\x79\x42\xa1\x88\xa3\xb6\x27\xcb\x12\xfd\x09\xb4\x11\x30\xbd\xaf\x53\xe2\xc2\xea\xa1\x05\xf3\xbb\x6f\xda\x6e\x8a\x63\x2b\x7a\xe5\x9b\x94\x5b\x69\x4d\x9e\x36\x33\x4d\x4f\xab\x1b\xd2\x56\x5f\xb4\xe9\xeb\xc7\xc4\xaa\x57\x93\xbe\x6e\xae\x25\x5f\xf9\x97\xd1\x9d\xe2\xa7\x79\xd2\x6e\xbf\x96\x6c\x90\x23\x13\x3e\x06\x3f\x95\x13\x86\xdf\x7c\xfc\xaa\xcb\x54\xd2\x21\xb0\x27\xf6\x2a\xa6\x93\x73\x7c\xe7\x48\xa1\xa8\x40\x5a\xe8\x88\x99\x3a\x39\x26\x61\xa0\x68\xf9\xab\x7f\x97\x5e\xa6\xa5\xac\xfa\x95\xba\xd5\x76\x62\xe5\x93\x6a\x57\x27\x89\x50\x7f\x05\xf6\x99\xba\x67\xd0\x2e\x71\x8b\xb7\x06\x4e\x7d\x0c\x5d\x4a\xb5\x8f\x7f\xf3\xb3\xcf\x88\xcc\xcb\x82\x79\xb5\xa3\x16\x79\x5f\xc3\x32\xc7\xfa\x12\x8f\xb7\xf4\xf7\x76\x48\x32\x28\xf1\x17\x6b\x59\x8f\xd1\x83\xd5\x14\xc6\xc7\xb7\x55\x48\xbb\xeb\xc3\x01\xbb\x21\x03\x72\x31\x83\xc7\x9b\xff\x88\xa1\x56\x54\x54\x7e\xee\x4d\xd4\x84\x77\x6d\x3a\xb5\x7b\xfb\x56\xa1\x52\xaa\xbd\x56\x27\x0a\x8c\x64\xdd\x0b\x41\xd3\x51\xb9\xcf\x0a\xa4\x73\xca\x26\x17\x14\x9c\xac\x56\xfa\x37\x33\xca\x5f\xbe\xe5\xac\x67\x0a\xcd\xf5\x6f\x4c\x17\x4a\xdd\xff\x8a\x6b\x91\x5e\xa9\x1d\x70\xb3\xb8\x52\xdf\x0d\xd4\x32\xd9\x28\x4f\xe3\xa9\xd4\x04\x41\xd4\x7d\x68\xb2\x35\x01\x32\x6b\x21\x40\x9e\x5a\xff\xdd\xfd\x09\xe9\xed\x05\x0c\xa0\xd0\x0f\xd8\xd3\x99\xa8\xa3\x98\xfe\xe2\xd8\xd7\xab\xf0\x35\x44\xc9\xc0\x85\x98\x58\xef\xe9\x57\xf9\x8f\xa8\x5f\x90\xaf\x80\x1d\x76\x13\x77\x9a\x3d\x38\x93\x8f\x43\x2d\x47\x9c\xca\x92\xc4\x77\x6d\xc3\x2d\xfb\xa0\x9b\xa1\xbf\x04\xe5\x02\x4c\x67\x83\xe3\xe0\x59\xc5\xc1\x09\x3e\x79\x5c\x45\x0f\xdc\x13\xdb\x8e\x75\xba\x4a\x94\xf4\x4c\x51\x0f\xc7\xd4\x8d\x74\xea\xc5\x72\x6a\x74\xaa\xc2\xe9\x15\xe8\xef\xaa\x99\x42\xe8\xd9\x27\x8b\x55\xb6\xa7\x8b\x64\xe1\xeb\x5d\x48\x97\x31\x2d\xe6\x8a\xce\x6a\x7a\x73\x92\x41\xbd\xf7\x89\xb5\x47\x6a\xbc\x25\xb6\x61\xaf\x04\x56\xe4\x5b\x67\xfd\x34\xfb\xf2\x1c\xf8\x50\x51\xe9\xa5\xc1\xd6\xec\xf6\x50\xe1\xfe\x6b\xa3\x4e\x9e\x1f\xa3\x4f\xe8\x1e\x0d\xf0\x8a\xff\x03\x93\x30\x8f\x48\xa8\x65\xee\x08\x5a\x1f\x79\x23\x5e\x92\x93\x7e\x90\x76\x2f\xbb\xd4\x22\x75\x05\x16\x9f\xe0\x25\xf2\x55\x38\xc9\xa7\x45\x2c\x9d\x58\x45\x4f\xcc\xce\x6f\x6f\xcb\x36\xed\x76\x40\x14\xd8\xfe\x6e\xf7\x9a\xe0\x59\x04\x0a\x3f\xb1\xda\x78\x43\x4e\x8b\xea\x66\x22\x27\x2b\x39\x48\xc3\x65\x3c\x3c\x9f\xbb\x26\x0d\xe6\x7b\x84\xcf\x18\xcd\x2b\xa9\xc9\xad\xa7\x0e\x8f\xb7\x3b\x16\x98\xfb\x13\x7d\x42\xa9\x47\x9a\x74\x04\x12\x0b\xe6\xfd\x3e\x51\xc3\x55\x3a\x03\x93\x9e\xe4\x22\xa7\xec\xa9\x62\xa5\x0b\xb3\x1e\xa1\xd4\xea\xc7\x3a\x02\xd9\x05\x75\x07\xc1\x24\x72\x1c\x88\x94\x5f\x4f\x7d\x5e\xfe\x93\xf6\x96\x30\x26\xd8\xb1\x67\x06\x9c\x37\x58\xd7\x72\xed\x5c\xc1\xec\xa5\xe4\xae\x3a\x64\x1b\x59\x14\x93\x8b\x4d\xca\x96\xd3\xcb\x7a\x34\xa5\xa0\x52\xb4\xa0\x21\x4c\xcf\x8d\x4d\xcf\xfd\x78\x87\x31\xa5\xc9\x55\xf8\x60\x26\x17\x13\x71\x51\xc3\xd5\x9e\x73\xa4\x2e\xde\x66\xb8\x42\xb9\xea\xa6\xe2\x84\xf7\xe3\x02\x97\x69\xa9\x84\x11\xac\xc9\x4b\x85\xa4\x61\x96\x8b\xdc\xed\x03\xa2\x4a\x2f\x78\xf0\xc8\xed\xaf\x6a\x67\x92\xb1\xca\xef\xdc\x65\x6e\x6e\x1b\x89\xbd\xdd\xe0\xd3\x1f\xb0\x31\x72\x5a\xbb\x3c\xc2\x8b\xd1\xf0\x11\xb4\x21\x6b\x11\xf4\x4e\xcd\xb5\x24\x84\xb7\xdb\x88\x31\xf5\x6c\xf8\x87\x29\xa3\xa1\x46\xd9\x04\x50\x64\xce\xbd\xa3\x7b\x57\xca\xfc\x03\xdf\x0d\xbb\x6c\xde\xbe\x5c\x16\xf2\xae\x3c\xba\x97\x2c\x12\x62\xa3\xe7\x95\xfc\x9f\x26\x1c\x84\xd4\xcd\x27\x8c\xcd\xc1\xaa\x76\x0f\xac\xa6\x46\xdc\x9f\xa9\x6e\x1f\x09\xc2\x3d\x4a\x7e\x4b\x12\x9d\x68\x8d\x34\x0b\x07\x4e\xd6\xef\x38\x2f\x56\x6a\xd6\x7a\xa8\x18\x83\x0d\xd6\xd6\x23\x2e\xed\xac\xbd\xe7\x26\xbf\x0e\x14\x3e\xbe\xa5\x5d\xbe\xa6\x99\x3c\x7b\x0e\xd2\x11\x23\xd4\x6d\xac\xd6\x49\x1f\x95\x6d\xca\xdf\xe6\xb1\xfc\xe0\xe0\x5d\xd0\xa6\x66\x73\x7e\xd2\xa4\x22\x34\x0b\x13\x26\xe6\x14\x5e\x8e\xc7\x9c\xed\xdc\x38\x67\x7a\xb3\x25\x6b\x28\xfb\x52\x5d\x6c\xaf\xb6\x9c\x56\xc2\x9e\x5b\x6e\x95\x9c\x67\x53\x85\x8d\x18\xff\x7b\xed\x7c\xd3\x4f\x8e\xca\xc6\x0b\x54\xf2\x6f\xae\x6f\xdf\xfb\xce\x8c\x2a\xd4\x84\xcb\xac\x5f\x6a\x2d\x3c\x74\x50\xd8\xa6\x7e\x40\x54\x9a\x64\xbb\xb4\x4f\x76\x67\xdf\x64\x6a\x44\xed\xcf\x19\xe6\x09\x65\x9f\x9d\xb1\x4a\x18\x49\x35\x35\x79\x59\x32\x2e\xd1\xc2\xeb\x55\x6b\xa1\xe8\x75\x88\x92\x59\x64\x6b\x59\x16\xe1\x93\xa8\x29\xbb\xf9\xf0\x55\xe6\x54\x9e\x81\xc5\x6c\x80\xbc\x4a\xbb\x8b\x1f\x74\xae\x88\x78\x34\xcd\xe7\x45\x71\xd2\x76\x8f\xb5\xf5\x1d\xc7\x87\x0b\x55\x8e\xb5\xa8\x82\x43\xf4\xdc\x32\x92\x18\xc8\xcf\x1a\x9f\xc8\xd3\x8b\x3f\x57\x54\x38\x4b\xf4\xd4\x44\xe9\xd0\x76\xf6\x55\xc8\x2f\x87\xa7\x22\x1c\x6e\x65\x98\xd7\x3e\x9c\xd9\x15\x2f\xf6\x1e\xca\x1f\x51\xb0\xf0\xea\x7b\xf3\xba\x5b\xe7\x06\x26\x0c\x14\xdd\x5c\xe1\xd7\x90\x1d\x2f\xaf\x21\xef\x35\x94\x7f\x74\xdb\x81\x8a\x57\xb9\x12\x43\x4f\xaa\x52\x1e\x74\x4e\xdb\x1c\x78\x34\xb6\xfe\x09\x65\x72\x8c\x1c\x1c\x0e\x0a\x3e\x32\xf7\x3c\x0f\x4c\x1b\x1d\x46\x45\x33\x60\xcc\xd3\xf3\x06\x77\x7a\x15\x55\xde\x1d\xc5\xb9\x2f\x25\xf7\x14\x8e\x55\x4a\xd2\x14\xda\x0b\x75\x4a\x22\xdb\x1f\x2f\xee\xd5\x20\xe0\x34\xca\xd5\x89\xed\xc7\xf4\xc7\xc6\xdf\x5e\xb9\x66\x9e\xdd\x02\x1c\xba\x4a\xfa\xac\x9b\x6b\xe3\x9d\x23\xbb\x0b\xe9\x11\x15\x27\xc5\x6a\x07\xe8\xba\x5b\xee\xec\x7d\x91\xbf\x43\x14\x93\x26\x99\xd2\x53\xff\xb6\xb1\xfe\xc1\x89\xc0\x19\xa0\x9a\xfd\xfd\x0e\xe5\xa5\x55\x47\xd9\x6f\xf8\x5a\xa4\x96\x43\xeb\x50\xdc\xed\x8d\xb4\xae\x86\xd0\x28\x7c\x5b\xf3\x43\x79\x5d\x92\x5a\x4a\x4a\x45\xc6\x43\x3b\x70\x7f\x3c\x83\xa8\x0a\x36\x84\x55\x79\xdd\xc9\xf7\x05\x87\xd4\x65\x79\xfa\x8d\xef\x51\x36\xbf\x76\xd0\x3e\xe7\xb8\x5f\x35\x2b\xb2\x7e\xef\x2e\x73\x99\xcc\x41\xa9\x14\x4c\x7b\x61\xb4\xc8\xd5\xcc\x58\xe6\x9b\xa6\x6f\x65\xd3\x17\x3d\xf9\x8c\x0f\xbe\xc6\x6d\x60\x46\xc4\xb9\xfb\x2c\xce\x4e\x8e\xfc\xb6\xf8\x6e\x4d\xe3\x28\x62\xb2\xb1\x1a\xe2\xa0\xed\x6d\x77\xcf\x96\xcf\x93\x7e\xc3\x36\xb3\x30\x52\x99\x98\x78\xaa\xf1\x43\xee\x0b\xcd\xda\xc8\x43\x4b\x80\x9c\xee\xab\x21\x92\xd7\x50\x1e\xc8\x46\xfc\x27\x55\x71\x6e\xd7\x10\xb1\x58\x5a\xde\x30\xa6\x41\xed\x7e\x58\x5a\x88\xe1\x03\x0f\xa4\xf3\xe1\x58\x9b\x61\xbf\xeb\xcf\x74\x91\x9a\xea\x1f\xec\x9f\x4a\x0f\x02\x8f\x5a\x93\xb4\x38\xa5\xe7\x43\xce\xdf\xf4\xdb\xbb\x2d\xd9\x33\x57\x2a\x6e\xb1\xc4\xd8\xa5\x50\xe3\x7d\xe5\xc5\x32\x13\x52\x8d\x21\xe7\x11\xb3\xf2\xe9\xe0\x2f\x25\xd3\xbb\xfd\x94\xa2\xa1\x19\x0f\x97\xc2\x1b\xa2\x03\x80\x2d\x85\x47\x18\x96\x1d\x8b\xa7\x0f\x24\x68\x9d\xfc\x99\x15\x15\xf8\x42\xe3\xb2\xfd\x91\x29\x1c\xbd\x97\xa5\x02\xd4\x25\xd5\x68\x4f\x75\x6a\xe4\xde\x0f\x12\x56\x02\x83\x74\x76\x61\xd2\x1e\x48\xb8\x1c\x0a\x6c\xbe\xbf\xda\x1d\x59\x9f\xf2\xfe\xe5\xa1\x34\xd1\x19\xeb\x38\xfc\xa7\x53\x17\x5b\x9e\xd8\x0e\x9c\xac\x83\x4b\xb3\xdf\x53\x32\x0f\x15\x83\xcb\x02\x64\x4b\xb1\x3f\xd0\x80\xaf\x14\xf7\xfd\xde\xdd\xd8\x4b\xee\xbe\x1c\xbf\xfd\x8c\x68\x7b\xef\x5d\x47\xc4\xe1\xe1\x16\xe5\x50\x57\xba\xfe\x7c\x58\xaf\xc4\x7d\x43\x9d\x28\x03\xb4\x64\xca\xe4\x01\x2f\x97\xaa\x4a\x13\x0b\xf0\xb9\x91\x1a\x5d\x31\xa1\x94\xcb\xd7\xd9\x6c\x76\x50\x3c\x12\x2f\x89\x14\xcf\x99\x68\xde\xab\xe4\x98\xce\x55\x58\x93\xf5\x8a\x6e\x36\xa9\xaa\xb6\xc8\x4d\x4d\x54\xa0\xce\x8e\x5d\x6f\xf0\x4b\x56\xac\xbf\xa6\xdc\xf7\x39\x7b\xb0\x34\x79\x43\x2a\x9a\x5d\xe1\xb9\xe7\x2e\xcd\x78\xa1\xff\x36\x36\x0f\x71\x25\x87\xf1\x2b\x78\x7e\x27\xba\x47\xbc\x9b\xb7\x0b\x3d\xd1\x06\x8f\x7d\x73\xca\x8b\xc6\xac\x18\xfe\x55\xe5\xa7\x60\xf4\x9c\x94\x21\xeb\x62\xad\xe2\x92\xf5\xa7\xb7\xbf\x8c\xa7\xc8\xe2\x92\x63\xce\xab\xba\x12\x4e\x42\xf4\xcf\xd0\xa6\xfd\xf6\x38\xa0\x47\x1b\x3a\xe4\x9f\x8f\x2a\x3d\x65\x9f\x22\xd1\x14\xc7\xfa\x87\x5d\xd1\x73\xee\x32\x77\x61\x7a\x8d\x41\x99\x13\x1f\x82\x17\x94\x94\x07\xfb\x08\xd6\xef\xb5\x6c\xc2\xc6\xaf\xe2\xe3\xd2\x0b\xb3\xf2\xf5\x2e\xf5\x78\x8f\xbe\x91\x74\x72\xbf\xfe\x6c\xd9\xc3\xe1\x6a\x1c\xe3\xc0\x78\x8a\xfe\x4f\x82\xae\xa4\x95\xd3\xc6\x0e\xce\x7f\xc0\x42\x72\x35\x04\x01\x7a\xcc\xf0\xe4\x8b\xf0\xbd\xda\x4a\x2a\x15\x3e\x73\xfb\xf5\x82\xfe\x98\x0d\x70\xdb\xa8\x7a\xf2\x23\x2e\xd9\x7c\xa7\x93\x93\x4c\xba\x73\xe4\xdc\x0a\x71\x42\x5b\xff\x99\x2f\x83\x6b\x7b\x27\x13\x3e\x83\xc6\x5b\xef\xbb\x3e\x65\xcb\x4c\x15\x39\xe8\x6f\x21\xbc\x5b\xef\x78\x05\xdd\xba\x20\x3f\xe0\x94\x64\x5e\x38\xe9\x30\x1f\xbf\xcb\xd5\x35\xf4\x10\x65\x68\x86\xb8\xb2\x9b\xc8\xde\x84\xc7\xed\xff\xbc\x98\x72\x42\x2c\xde\x3e\x49\x75\x90\x31\x38\x4f\x57\x96\xb3\x2c\xb9\x77\xf4\x75\xed\xee\xd4\x83\xab\x52\x0a\xa6\xb3\x9e\xbf\x43\x69\x5a\x82\x93\x12\x0e\xcd\xae\x1e\x21\x94\x7d\x13\x45\x76\x4e\xeb\x98\x96\xec\xfd\xb2\x9c\xa6\xbe\x9c\x79\xc3\xd1\xf3\xe5\xb3\x80\x07\xc9\xe2\xb1\x9b\x8b\xe6\xac\x79\x9a\x73\x43\x69\x80\x7d\x73\xb1\xe3\x61\x15\x85\x84\x18\x4d\x50\x4e\xb7\x4e\xcd\xb5\xff\x75\x98\xd2\x8b\x2d\x5e\x13\x17\x01\xa4\x00\xcf\xd5\xf7\x27\x8a\xa0\xaf\xe9\x5b\x57\x26\xa2\xb5\x3d\xbe\x02\xeb\x1b\xff\x5f\x01\x00\x00\xff\xff\xb5\x11\xe1\x91\x9f\x0d\x00\x00")

// FilePemGgzPem is "/pem/ggz.pem"
var FilePemGgzPem = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xd4\xb9\xce\xab\xc8\x16\x05\xe0\x9c\xa7\xb8\x39\xba\x02\x8c\xf1\x10\x56\x15\x83\x0b\xcc\x64\x66\x32\x26\x97\x01\x33\xfc\x18\x5c\xc0\xd3\xb7\xce\x69\xb5\xd4\x52\xb7\x7a\x87\x2b\xf8\x92\xbd\xb4\xfe\xff\xeb\xa0\xa2\x61\xeb\x7f\x48\x79\xf8\x58\xc5\x08\xf8\xca\xef\x94\x31\x31\x46\xdc\x8e\x10\xa8\x2a\x02\x28\x86\x80\x60\x3d\x9c\xac\x2c\x23\xed\xa5\x98\x5f\x93\x09\x78\x0d\x79\x3f\x9a\x87\x73\x51\x76\x15\x88\x68\x00\x4c\xc8\xaf\xda\x0e\x52\x48\xac\x90\x81\xc0\xf4\x95\x4e\xa5\x85\xb6\xbe\x93\xee\x2a\xdd\x3b\x55\x28\x35\x42\xef\x9d\xf5\xcd\x7d\x50\xa9\x94\x5f\xad\x1d\x88\xa6\x9f\x50\xb3\x31\x77\xb3\x49\x33\x95\xf2\xbb\xe9\x83\xdd\x6c\x02\xca\xfc\x15\xfe\x5d\xfd\x6f\x14\x21\xe0\x61\x2a\xbb\x89\x6e\x0c\x29\x66\x5e\xdf\xc2\x02\xae\x02\xa1\x0b\x64\x42\x14\x07\xc8\x08\x01\x77\x40\x84\x28\x10\xd8\x63\xf0\x23\x13\xd7\x2e\x2f\x1b\x07\xcf\x5a\x7e\x67\x3b\xd7\xf9\x99\x16\xbe\xae\x7f\x5e\xc7\x97\x8d\x9d\x23\x43\xae\xd5\x23\x0b\x63\x73\x89\x05\x83\xab\xba\xeb\x61\xbe\x97\x7a\x1a\xb6\x9b\xdf\xd3\xa1\xb4\x2e\x91\x89\x83\x5e\x37\xae\xef\xcb\x59\xf1\xcc\x2b\x7c\xc3\xe5\x1d\x5f\x0e\xb3\xf1\x9c\xcf\x1f\x67\x46\x8c\xda\x0b\x7d\xcb\x7e\xa5\xe9\x72\xc4\xeb\x6c\x2b\xa7\xd6\xea\x72\x6c\x0e\x5f\xae\xc1\xe9\x59\x9d\xd9\x68\x70\xf5\xe1\x1c\x47\xe1\x92\x1e\xac\xe0\x9b\xdf\xf8\xa5\x91\xd9\xf8\xfd\x74\x1e\x61\x35\x1d\x99\x7c\xce\x6b\xb2\x5d\xa7\x8f\x23\xd0\xda\x3b\x73\x49\x9d\xa8\x88\xcd\x4a\xdb\x61\xb7\xbe\xe2\xba\x2f\x1b\x86\xd2\xf3\xd4\x4b\x06\xea\xfa\x82\x93\x6c\x5b\xfb\x50\xee\x7b\xff\x7e\xbb\xf2\x0a\x76\x87\x51\xe5\xb4\x46\x25\x69\xc9\x52\x3f\xb7\xcc\x33\x35\x6e\x16\x05\xfb\x94\xe7\x78\x99\xa3\xc7\xfb\xd3\x39\x8b\xa0\xb3\xec\xd3\x68\x3f\x72\xfb\x12\xc6\x31\x82\xa4\xde\xfb\x7a\x07\x8d\xc5\xf7\xf3\xcc\xa8\x90\x84\x03\x86\x91\xe7\x35\xec\x50\xf7\xa1\x96\x35\x63\x33\xa4\xef\xe6\x14\x41\x63\x11\x4e\x99\xe8\x2e\xc5\xbc\xf0\xee\x0d\x8d\x35\xbd\x72\x76\x72\x41\x80\x2a\x00\x64\x16\x32\x15\x40\x19\x99\x26\x72\xf8\xe0\x7d\xe0\xde\x38\x08\x02\x0a\xa8\x02\xb9\x1d\x94\xbf\xde\x79\x73\x8f\x8a\x4a\xdc\xc0\x8a\xcb\x89\x3f\xd2\xeb\x8f\xb2\xb0\xf3\x86\x0c\x7d\x2d\xed\xc7\xb2\x88\x96\xaf\x50\x46\x26\xbf\x01\xe7\x4f\xc0\x95\x01\x31\x94\x7f\x6b\x1f\x38\x62\x08\x5c\x18\x0b\xd5\xd0\x6d\x49\xa1\xdd\x8e\x2b\x64\x93\x47\xc7\x24\xdb\xc3\xd8\x71\x5b\xdf\x9d\xf1\xfa\xb1\x7b\xb9\x96\xad\x9a\x6d\x8c\x5b\x2d\xca\xed\x15\x98\x1c\x88\xea\x55\xfe\x18\xa7\xb0\x31\xa9\x77\x5e\xc7\x48\xa4\x6b\x63\x6f\xa7\xbe\x09\x4e\xfd\xee\x4b\x4c\x60\xec\x64\xb9\xeb\x4f\xc3\x19\x3f\x13\xde\xb5\x47\x74\x6e\x5e\x56\x91\x8e\x43\xfd\x6a\x43\x5f\xc4\x3d\xf5\xf5\x69\x3b\x38\x1c\x55\xc3\xa7\x88\x6f\x6d\x19\xb3\x66\x9a\x07\x59\x90\xb4\x8e\x37\x33\x87\x68\xd5\x84\x4a\x95\xf3\x77\x54\x24\xb6\x76\x38\xf0\x13\xaf\x36\x6e\x5a\x91\xbd\x1e\x9b\xa2\x3e\xc2\x22\x42\xb0\x36\x03\xcc\x75\x29\x3b\x6f\x1e\x0a\x94\x38\x9f\x58\x49\x0c\x4e\x5b\x64\x8f\x0c\xba\xa8\x29\xbb\xde\x56\x37\x7b\x60\x8e\x34\x53\xd6\xdb\x88\x86\x47\x96\xb7\x47\xec\x68\xd4\xc8\x63\x28\xf4\xa3\x54\x0e\x3d\xd5\x6f\x44\x3f\x14\x8f\x34\x92\xbb\x17\xb4\x4f\x0a\x27\xd9\xa2\xc4\x18\x88\xd4\x33\x3f\x7a\xca\x29\x96\xd2\x72\xd0\x85\x30\xbe\xe9\xcb\x18\x94\xc3\x4f\x91\xe4\x51\xcc\xc3\x66\xe9\x2b\xbb\xf2\x25\x5b\x5b\x17\xef\x6d\xb5\xf3\x5b\x55\xc6\xcb\x44\x58\x47\x7a\xde\x9c\xcb\x2c\x33\xc5\x94\x88\xcc\xef\x75\x50\x2c\xf9\x9f\x8b\xf1\x47\x00\x00\x00\xff\xff\xef\x91\x8c\x5a\x4e\x04\x00\x00")

func init() {
	if CTX.Err() != nil {
		log.Fatal(CTX.Err())
	}

	var err error

	err = FS.Mkdir(CTX, "/images/", 0777)
	if err != nil && err != os.ErrExist {
		log.Fatal(err)
	}

	err = FS.Mkdir(CTX, "/pem/", 0777)
	if err != nil && err != os.ErrExist {
		log.Fatal(err)
	}

	var f webdav.File

	var rb *bytes.Reader
	var r *gzip.Reader

	rb = bytes.NewReader(FileImagesFaviconIco)
	r, err = gzip.NewReader(rb)
	if err != nil {
		log.Fatal(err)
	}

	err = r.Close()
	if err != nil {
		log.Fatal(err)
	}

	f, err = FS.OpenFile(CTX, "/images/favicon.ico", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}

	rb = bytes.NewReader(FileImagesLogoPng)
	r, err = gzip.NewReader(rb)
	if err != nil {
		log.Fatal(err)
	}

	err = r.Close()
	if err != nil {
		log.Fatal(err)
	}

	f, err = FS.OpenFile(CTX, "/images/logo.png", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}

	rb = bytes.NewReader(FilePemGgzPem)
	r, err = gzip.NewReader(rb)
	if err != nil {
		log.Fatal(err)
	}

	err = r.Close()
	if err != nil {
		log.Fatal(err)
	}

	f, err = FS.OpenFile(CTX, "/pem/ggz.pem", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}

	Handler = &webdav.Handler{
		FileSystem: FS,
		LockSystem: webdav.NewMemLS(),
	}

}

// Open a file
func (hfs *HTTPFS) Open(path string) (http.File, error) {
	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// ReadFile is adapTed from ioutil
func ReadFile(path string) ([]byte, error) {
	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))

	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(f)
	return buf.Bytes(), err
}

// WriteFile is adapTed from ioutil
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := FS.OpenFile(CTX, filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// WalkDirs looks for files in the given dir and returns a list of files in it
// usage for all files in the b0x: WalkDirs("", false)
func WalkDirs(name string, includeDirsInList bool, files ...string) ([]string, error) {
	f, err := FS.OpenFile(CTX, name, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	fileInfos, err := f.Readdir(0)
	f.Close()
	if err != nil {
		return nil, err
	}

	for _, info := range fileInfos {
		filename := path.Join(name, info.Name())

		if includeDirsInList || !info.IsDir() {
			files = append(files, filename)
		}

		if info.IsDir() {
			files, err = WalkDirs(filename, includeDirsInList, files...)
			if err != nil {
				return nil, err
			}
		}
	}

	return files, nil
}
