// Code generated by gocode.Generate; DO NOT EDIT.

package filetypes

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/encoding/gocode/gocodec"
)

var cuegenvalFileInfo = cuegenMake("FileInfo", &FileInfo{})

// Validate validates x.
func (x *FileInfo) Validate() error {
	return cuegenCodec.Validate(cuegenvalFileInfo, x)
}

var cuegenCodec, cuegenInstance = func() (*gocodec.Codec, *cue.Instance) {
	var r *cue.Runtime
	r = &cue.Runtime{}
	instances, err := r.Unmarshal(cuegenInstanceData)
	if err != nil {
		panic(err)
	}
	if len(instances) != 1 {
		panic("expected encoding of exactly one instance")
	}
	return gocodec.New(r, nil), instances[0]
}()

// cuegenMake is called in the init phase to initialize CUE values for
// validation functions.
func cuegenMake(name string, x interface{}) cue.Value {
	f, err := cuegenInstance.LookupField(name)
	if err != nil {
		panic(fmt.Errorf("could not find type %q in instance", name))
	}
	v := f.Value
	if x != nil {
		w, err := cuegenCodec.ExtractType(x)
		if err != nil {
			panic(err)
		}
		v = v.Unify(w)
	}
	return v
}

// Data size: 1122 bytes.
var cuegenInstanceData = []byte("\x01\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xccW_k\xe46\x10\xb76)T\xe2Z\xee\v\x14T?\x1c\u05c5\xeeC_\n\vG^\xae\aG\xa1\x94\xbe\x86#(^\xed\xc6=[2\xb6\\\x12\x9a}h{\xbd\xf6S\u07d6\xd1H\x96%{\x13\x92\x96r~I\xf67\x9a?\xfa\u034c4\xfa\xec\xf0\u05c2,\x0e\u007fg\xe4\xf0{\x96}{\xf8\ud110'\xa5\xea\x8cP\x85|)\x8c\x00\x9c\x9c\x90\u04df\xb46d\x91\x91\xd3\x1f\x85\xb9\"O2\xf2\u026b\xb2\x92\x1d9\xbc\u03f2\xec\x8b\u00df\vB>?\u007fS\xf4r\xb5-+\xa7\xf9>#\x87wY\xf6\xfc\xf0\xc7\t!\x9f\x06\xfc]F\x16\xe4\xf4\aQK0tjA\x96e\u0647\xa7\xdfC$\x84,\b\xa1\u6991\u076a\xe8%\xf9\xf0\xf4\x9bF\x14o\xc5N\xf2\u02fe\xac6\x8c\x81k\xbe^\xf3_\x19\x05\xabJ\xd4r\xcd\xdd\u05d9\xb6T;F\xa5*\xf4\xa6T\xbbA\xf0\x9d\x03\x18-\x95\x91m\xd3J#L\xa9\xd5\u065a\xbf\x8e\x00F\xb7\xba\xad\xcf\x06E\xce\xf9+\xdd\u058c\x1a\xb1\xeb\u03acWz\x8en\u07ac\a\u007f{\xb6g\xde\x05\u0106\u07d7/\xf2\x9c\xc5\xe6A\xe8\x94\xc0lX;\x8a\x1e<YGF^\x1b\xf4\x18\xf6\x93\x03\x983j\xc3D\xe5|#\x8c\xc8!\b\n\xff\xa1\x06\x8aG\xa2\xa2\x97\x13[E/Q\xd8\x15W\xb2\x8e5\x11B\xf1\u03ddV\x13e\x00\x83\xb8\x9a\x95W\xb8\xe0F\xd4S9\x80(\xde\xe9T\x88_^\xe8\r\x04\x98\xe4l\xcds\x00\a\x9a(\xad\x84\xb5\xb8\u04c0\xef\xad\u0366\xd5fb6\xb7\xa8s\u068a\xe6*\u06b1E<\x8f\xbb\x84\u019dcQo&4>*X_9.\\\xd1O\xa3\xc5o\x89D\xf3[~1g\x1c\x14C\x16\xc6y|\xa0\xa1\xa0\x8e\xe6t#\x95h\xcaG\xd9r\xba96\xc6K\xb9\x15}e\xa0\xd6m\xef>\x8b[w\x99\u007f\r\x86\x1c!{\xdb\u07ef\xd5V\xbb\x1e\u01d2\xf6\xdf\u04b4\xbd\xe4\xb7|+\xaaN2\xda\u02adl\xa5*d\xb7\x9e\n\x8b\x9b\xa2B\xc1\x8c\xe6FnKUB\xbc\xb0\xe2R\xeb\n\xb6\f\xbfE\x85*\x88\x15Zu\xa6\x15\xa52a\xdd[)\x1b\xb7\xa9n\xed\xb0R\x15\xban*i\xeca\u4c3a\u046d\xf1\x11 \u0599V\x8a\xda\a\x85\xd8F\x17]\xd8\"b\u0098\xb6\xbc\xec\rn\xc0\u017et\xc1\x03El\xcfj\xbd\x91XU\xa5jzwT\x8c\u0636\xe5\x16R\xb7\xb4\xfd\xeerFW\xab\x15V\x1fM\b\xa7.\x9e\x84\xb0\xb1\x86\xb7\xe9Kz8[(]B\xc7t+\xac$\xefk\xef\xf5\xae\x8dT\x1drn\x97\xe7+[G^9-\xa4%\x96wd\x06\xcb\aT\xed\x01\xf2H\xd5\aj\xdaBFuy\rI\xbd\x9f\xebq\x83\xdcOv\x92\xe0\x87\x91\r-r,\\\xdd7\u0197\xc6\xff\xed[\xfe\"\xaa\xff\xb8(\xffU\xac\xdbR\x89\xeaX\xb0\x1b\xb9\xfd\xe8\x1b\bN\xd3XuP\xf4\xe7\x8as6l+\\\xdd~\x05r\xc7o-\x97\x8c\x0e\xa7\x8f\x0f3\xe4ud&\xdc\xe0\x91#H\x853\x83j\xf6\x0eM\xdc'\x8a\xd1\xfa\xc4M\x98x\x92\xfd\x1cY\xaeg\xa3:\xb6\xfc\ue640\x8e\x89\x1fi\x85+\u007f\xde\xc9H\x81Y\xc4.>\x87\t\xf7\xc50(\xfa\x12\xf1F\xf3|\xcd/\xfc\x8f\xe9\xf45\\w~\f\xe3\xb7<\xb7\xe5k\xff\xf3\x93Jr\x8b\xa5\x05\x17\xdfg\xcf#\xf1W\xfcY\x8a0\x9a\xdcv\xa9\xbd\xf8\xdeK\xa5\xf1\r8\x91Fwa*\x8do\u0164\u0247LX\x02\xe6hr\xd4L\xb6<\x1f\xf8\xd1\xf4\xa1\x97\xc9`\x18\x92\x81\xbcC\x06` \u013fv\xbeN\xa6\x14\xdfYQv\xe6\xb3\xe2\u0458\xf9\xbb\x03\x8f\x99\x9eg8\xe5.\x99h\xed\xe3d(#?)\xc5\u0324\xbd\x9c>.\xc2+\a\xc7\xe3x\x1a\x1b66\x9e\xc2f9\x98\xa5`vWiw\xbb\x97W2V@o\xb9\uf085\x11\x83s\x18\xc2\xed/\x16\xa6\a\x87\xc2/6\f\x06)\n\xc7~@\vH.\x9a\xb5\x8b\a\xb3vm\xb5q\xfebX\xcd\xc3\xe6\u068c,\xc3\xf1\a\xe8N\xbb-Xt\xa7\x01\xc3c\u02bb\xb3\xbf\xe0\xc0)+\x19N\x97t\xb4\u03b7Z\xaf\xdc;m\xf4^\xf4c\xfe\x9e\u0173\xcf\xc3O\xae\xf0l9\xd2L\xc7\x1f%\xf1\x84tD\xfd\xc8#\xe4\x1e]\x96e\xff\x04\x00\x00\xff\xff\xbe\xea\xa5\xf0\xdb\x10\x00\x00")
