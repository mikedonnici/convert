package convertv1

import (
	"context"
	"log"

	"connectrpc.com/connect"

	"github.com/regrowag/ses/go/pkg/convert"
)

// Server is the convert server with datastore attached
type Server struct{}

// NewServer returns a pointer to a server
func NewServer() *Server {
	return &Server{}
}

// ConvertValue converts values with common simple and complex units
func (s *Server) ConvertValue(ctx context.Context, req *connect.Request[ConvertValueRequest]) (*connect.Response[ConvertValueResponse], error) {
	log.Println("ConvertValue()")

	v, err := convert.ValueFromTo(req.Msg.GetValue(), req.Msg.GetFromUnit(), req.Msg.GetToUnit())
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&ConvertValueResponse{
		Value: v,
		Unit:  req.Msg.GetToUnit(),
	}), nil
}

// ConvertCropYield converts crop yields that may use crop-specific units like bushels and bales.
func (s *Server) ConvertCropYield(ctx context.Context, req *connect.Request[ConvertCropYieldRequest]) (*connect.Response[ConvertCropYieldResponse], error) {
	log.Println("ConvertCropYield()")

	v, err := convert.CropRate(req.Msg.GetCrop(), req.Msg.GetValue(), req.Msg.GetFromUnit(), req.Msg.GetToUnit())
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&ConvertCropYieldResponse{
		Value: v,
		Unit:  req.Msg.GetToUnit(),
		Crop:  req.Msg.GetCrop(),
	}), nil
}
