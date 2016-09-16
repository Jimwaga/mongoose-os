// +build clubby_strict

package filesystem

// GENERATED FILE DO NOT EDIT
// This file is automatically generated with miot clubbygen.

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"cesanta.com/clubby"
	"cesanta.com/clubby/endpoint"
	"cesanta.com/clubby/frame"
	"cesanta.com/common/go/ourjson"
	"cesanta.com/common/go/ourtrace"
	"github.com/cesanta/errors"
	"golang.org/x/net/trace"

	"github.com/cesanta/ucl"
	"github.com/cesanta/validate-json/schema"
	"github.com/golang/glog"
)

var _ = bytes.MinRead
var _ = fmt.Errorf
var emptyMessage = ourjson.RawMessage{}
var _ = ourtrace.New
var _ = trace.New

const ServiceID = "http://mongoose-iot.com/fw/v1/Filesystem"

type GetArgs struct {
	Filename *string `json:"filename,omitempty"`
	Len      *int64  `json:"len,omitempty"`
	Offset   *int64  `json:"offset,omitempty"`
}

type GetResult struct {
	Data *string `json:"data,omitempty"`
	Left *int64  `json:"left,omitempty"`
}

type PutArgs struct {
	Append   *bool   `json:"append,omitempty"`
	Data     *string `json:"data,omitempty"`
	Filename *string `json:"filename,omitempty"`
}

type Service interface {
	Get(ctx context.Context, args *GetArgs) (*GetResult, error)
	List(ctx context.Context) ([]string, error)
	Put(ctx context.Context, args *PutArgs) error
}

type Instance interface {
	Call(context.Context, string, *frame.Command) (*frame.Response, error)
	TraceCall(context.Context, string, *frame.Command) (context.Context, trace.Trace, func(*error))
}

type _validators struct {
	// This comment prevents gofmt from aligning types in the struct.
	GetArgs *schema.Validator
	// This comment prevents gofmt from aligning types in the struct.
	GetResult *schema.Validator
	// This comment prevents gofmt from aligning types in the struct.
	ListResult *schema.Validator
	// This comment prevents gofmt from aligning types in the struct.
	PutArgs *schema.Validator
}

var (
	validators     *_validators
	validatorsOnce sync.Once
)

func initValidators() {
	validators = &_validators{}

	loader := schema.NewLoader()

	service, err := ucl.Parse(bytes.NewBuffer(_ServiceDefinition))
	if err != nil {
		panic(err)
	}
	// Patch up shortcuts to be proper schemas.
	for _, v := range service.(*ucl.Object).Find("methods").(*ucl.Object).Value {
		if s, ok := v.(*ucl.Object).Find("result").(*ucl.String); ok {
			for kk := range v.(*ucl.Object).Value {
				if kk.Value == "result" {
					v.(*ucl.Object).Value[kk] = &ucl.Object{
						Value: map[ucl.Key]ucl.Value{
							ucl.Key{Value: "type"}: s,
						},
					}
				}
			}
		}
		if v.(*ucl.Object).Find("args") == nil {
			continue
		}
		args := v.(*ucl.Object).Find("args").(*ucl.Object)
		for kk, vv := range args.Value {
			if s, ok := vv.(*ucl.String); ok {
				args.Value[kk] = &ucl.Object{
					Value: map[ucl.Key]ucl.Value{
						ucl.Key{Value: "type"}: s,
					},
				}
			}
		}
	}
	var s *ucl.Object
	_ = s // avoid unused var error
	s = &ucl.Object{
		Value: map[ucl.Key]ucl.Value{
			ucl.Key{Value: "properties"}: service.(*ucl.Object).Find("methods").(*ucl.Object).Find("Get").(*ucl.Object).Find("args"),
			ucl.Key{Value: "type"}:       &ucl.String{Value: "object"},
		},
	}
	if req, found := service.(*ucl.Object).Find("methods").(*ucl.Object).Find("Get").(*ucl.Object).Lookup("required_args"); found {
		s.Value[ucl.Key{Value: "required"}] = req
	}
	validators.GetArgs, err = schema.NewValidator(s, loader)
	if err != nil {
		panic(err)
	}
	validators.GetResult, err = schema.NewValidator(service.(*ucl.Object).Find("methods").(*ucl.Object).Find("Get").(*ucl.Object).Find("result"), loader)
	if err != nil {
		panic(err)
	}
	validators.ListResult, err = schema.NewValidator(service.(*ucl.Object).Find("methods").(*ucl.Object).Find("List").(*ucl.Object).Find("result"), loader)
	if err != nil {
		panic(err)
	}
	s = &ucl.Object{
		Value: map[ucl.Key]ucl.Value{
			ucl.Key{Value: "properties"}: service.(*ucl.Object).Find("methods").(*ucl.Object).Find("Put").(*ucl.Object).Find("args"),
			ucl.Key{Value: "type"}:       &ucl.String{Value: "object"},
		},
	}
	if req, found := service.(*ucl.Object).Find("methods").(*ucl.Object).Find("Put").(*ucl.Object).Lookup("required_args"); found {
		s.Value[ucl.Key{Value: "required"}] = req
	}
	validators.PutArgs, err = schema.NewValidator(s, loader)
	if err != nil {
		panic(err)
	}
}

func NewClient(i Instance, addr string) Service {
	validatorsOnce.Do(initValidators)
	return &_Client{i: i, addr: addr}
}

type _Client struct {
	i    Instance
	addr string
}

func (c *_Client) Get(pctx context.Context, args *GetArgs) (res *GetResult, err error) {
	cmd := &frame.Command{
		Cmd: "/v1/Filesystem.Get",
	}
	ctx, tr, finish := c.i.TraceCall(pctx, c.addr, cmd)
	defer finish(&err)
	_ = tr

	tr.LazyPrintf("args: %s", ourjson.LazyJSON(&args))
	cmd.Args = ourjson.DelayMarshaling(args)
	if args.Filename == nil {
		return nil, errors.Errorf("Filename is required")
	}
	b, err := cmd.Args.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal args as JSON: %+v", err)
	} else {
		v, err := ucl.Parse(bytes.NewReader(b))
		if err != nil {
			glog.Errorf("Failed to parse just serialized JSON value %q: %+v", string(b), err)
		} else {
			if err := validators.GetArgs.Validate(v); err != nil {
				glog.Warningf("Sending invalid args for Get: %+v", err)
				return nil, errors.Annotatef(err, "invalid args for Get")
			}
		}
	}
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if resp.Status != 0 {
		return nil, errors.Trace(&endpoint.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}

	tr.LazyPrintf("res: %s", ourjson.LazyJSON(&resp))

	bb, err := resp.Response.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal result as JSON: %+v", err)
	} else {
		rv, err := ucl.Parse(bytes.NewReader(bb))
		if err == nil {
			if err := validators.GetResult.Validate(rv); err != nil {
				glog.Warningf("Got invalid result for Get: %+v", err)
				return nil, errors.Annotatef(err, "invalid response for Get")
			}
		}
	}
	var r *GetResult
	err = resp.Response.UnmarshalInto(&r)
	if err != nil {
		return nil, errors.Annotatef(err, "unmarshaling response")
	}
	return r, nil
}

func (c *_Client) List(pctx context.Context) (res []string, err error) {
	cmd := &frame.Command{
		Cmd: "/v1/Filesystem.List",
	}
	ctx, tr, finish := c.i.TraceCall(pctx, c.addr, cmd)
	defer finish(&err)
	_ = tr
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if resp.Status != 0 {
		return nil, errors.Trace(&endpoint.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}

	tr.LazyPrintf("res: %s", ourjson.LazyJSON(&resp))

	bb, err := resp.Response.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal result as JSON: %+v", err)
	} else {
		rv, err := ucl.Parse(bytes.NewReader(bb))
		if err == nil {
			if err := validators.ListResult.Validate(rv); err != nil {
				glog.Warningf("Got invalid result for List: %+v", err)
				return nil, errors.Annotatef(err, "invalid response for List")
			}
		}
	}
	var r []string
	err = resp.Response.UnmarshalInto(&r)
	if err != nil {
		return nil, errors.Annotatef(err, "unmarshaling response")
	}
	return r, nil
}

func (c *_Client) Put(pctx context.Context, args *PutArgs) (err error) {
	cmd := &frame.Command{
		Cmd: "/v1/Filesystem.Put",
	}
	ctx, tr, finish := c.i.TraceCall(pctx, c.addr, cmd)
	defer finish(&err)
	_ = tr

	tr.LazyPrintf("args: %s", ourjson.LazyJSON(&args))
	cmd.Args = ourjson.DelayMarshaling(args)
	if args.Filename == nil {
		return errors.Errorf("Filename is required")
	}
	b, err := cmd.Args.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal args as JSON: %+v", err)
	} else {
		v, err := ucl.Parse(bytes.NewReader(b))
		if err != nil {
			glog.Errorf("Failed to parse just serialized JSON value %q: %+v", string(b), err)
		} else {
			if err := validators.PutArgs.Validate(v); err != nil {
				glog.Warningf("Sending invalid args for Put: %+v", err)
				return errors.Annotatef(err, "invalid args for Put")
			}
		}
	}
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return errors.Trace(err)
	}
	if resp.Status != 0 {
		return errors.Trace(&endpoint.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}
	return nil
}

func RegisterService(i *clubby.Instance, impl Service) error {
	validatorsOnce.Do(initValidators)
	s := &_Server{impl}
	i.RegisterCommandHandler("/v1/Filesystem.Get", s.Get)
	i.RegisterCommandHandler("/v1/Filesystem.List", s.List)
	i.RegisterCommandHandler("/v1/Filesystem.Put", s.Put)
	i.RegisterService(ServiceID, _ServiceDefinition)
	return nil
}

type _Server struct {
	impl Service
}

func (s *_Server) Get(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	b, err := cmd.Args.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal args as JSON: %+v", err)
	} else {
		if v, err := ucl.Parse(bytes.NewReader(b)); err != nil {
			glog.Errorf("Failed to parse valid JSON value %q: %+v", string(b), err)
		} else {
			if err := validators.GetArgs.Validate(v); err != nil {
				glog.Warningf("Got invalid args for Get: %+v", err)
				return nil, errors.Annotatef(err, "invalid args for Get")
			}
		}
	}
	var args GetArgs
	if len(cmd.Args) > 0 {
		if err := cmd.Args.UnmarshalInto(&args); err != nil {
			return nil, errors.Annotatef(err, "unmarshaling args")
		}
	}
	if args.Filename == nil {
		return nil, errors.Errorf("Filename is required")
	}
	r, err := s.impl.Get(ctx, &args)
	if err != nil {
		return nil, errors.Trace(err)
	}
	bb, err := json.Marshal(r)
	if err == nil {
		v, err := ucl.Parse(bytes.NewBuffer(bb))
		if err != nil {
			glog.Errorf("Failed to parse just serialized JSON value %q: %+v", string(bb), err)
		} else {
			if err := validators.GetResult.Validate(v); err != nil {
				glog.Warningf("Returned invalid response for Get: %+v", err)
				return nil, errors.Annotatef(err, "server generated invalid responce for Get")
			}
		}
	}
	return r, nil
}

func (s *_Server) List(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	r, err := s.impl.List(ctx)
	if err != nil {
		return nil, errors.Trace(err)
	}
	bb, err := json.Marshal(r)
	if err == nil {
		v, err := ucl.Parse(bytes.NewBuffer(bb))
		if err != nil {
			glog.Errorf("Failed to parse just serialized JSON value %q: %+v", string(bb), err)
		} else {
			if err := validators.ListResult.Validate(v); err != nil {
				glog.Warningf("Returned invalid response for List: %+v", err)
				return nil, errors.Annotatef(err, "server generated invalid responce for List")
			}
		}
	}
	return r, nil
}

func (s *_Server) Put(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	b, err := cmd.Args.MarshalJSON()
	if err != nil {
		glog.Errorf("Failed to marshal args as JSON: %+v", err)
	} else {
		if v, err := ucl.Parse(bytes.NewReader(b)); err != nil {
			glog.Errorf("Failed to parse valid JSON value %q: %+v", string(b), err)
		} else {
			if err := validators.PutArgs.Validate(v); err != nil {
				glog.Warningf("Got invalid args for Put: %+v", err)
				return nil, errors.Annotatef(err, "invalid args for Put")
			}
		}
	}
	var args PutArgs
	if len(cmd.Args) > 0 {
		if err := cmd.Args.UnmarshalInto(&args); err != nil {
			return nil, errors.Annotatef(err, "unmarshaling args")
		}
	}
	if args.Filename == nil {
		return nil, errors.Errorf("Filename is required")
	}
	return nil, s.impl.Put(ctx, &args)
}

var _ServiceDefinition = json.RawMessage([]byte(`{
  "methods": {
    "Get": {
      "args": {
        "filename": {
          "doc": "Name of the file to read.",
          "type": "string"
        },
        "len": {
          "doc": "Length of chunk to read. If omitted, all available data until the EOF\nwill be read. If (offset + len) is larger than the file size, no\nerror will be returned, and only available data until the EOF will be\nread.\n",
          "type": "integer"
        },
        "offset": {
          "doc": "Offset from the beginning of the file to start reading from.\nIf omitted, 0 is assumed. If the given offset is larger than the file\nsize, no error is returned, and the returned data will be null.\n",
          "type": "integer"
        }
      },
      "doc": "Read a file or a part of file from the device's filesystem.",
      "required_args": [
        "filename"
      ],
      "result": {
        "properties": {
          "data": {
            "doc": "Chunk of data read from the file.",
            "type": "string"
          },
          "left": {
            "doc": "Number of bytes left past the read chunk of data.",
            "type": "integer"
          }
        },
        "type": "object"
      }
    },
    "List": {
      "doc": "List files at the device's filesystem.",
      "result": {
        "items": {
          "doc": "Filename",
          "type": "string"
        },
        "type": "array"
      }
    },
    "Put": {
      "args": {
        "append": {
          "doc": "If true, and if the file with the given filename already exists, the\ndata will be appended to it. Otherwise, the file will be overwritten\nor created.\n",
          "type": "boolean"
        },
        "data": {
          "doc": "Data to write or append.",
          "type": "string"
        },
        "filename": {
          "doc": "Name of the file to write or append to.",
          "type": "string"
        }
      },
      "doc": "Write or append data to file.",
      "required_args": [
        "filename"
      ]
    }
  },
  "name": "/v1/Filesystem",
  "namespace": "http://mongoose-iot.com/fw"
}`))
