package server

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	api "omdb/api/v1"
	external "omdb/external"
	"strconv"
)

// We define a server struct that implements the server interface. 🥳🥳🥳
type grpcServer struct {
	api.UnimplementedOMDBServiceServer
}

//// We implement the SayHello method of the server interface. 🥳🥳🥳
//func (s *Server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
//	return &HelloReply{Message: "Hello, " + in.GetName()}, nil
//}

func (s *grpcServer) GetMovieByID(ctx context.Context, req *api.GetMovieByIDRequest) (*api.GetMovieByIDResponse, error) {

	_, resp := external.GetMovieByID(req.Id)
	//if err == ErrIDNotFound {
	//	return nil, status.Error(codes.NotFound, "id was not found")
	//}
	//if err != nil {
	//	return nil, status.Error(codes.Internal, err.Error())
	//}

	//s := string(`{"operation": "get", "key": "example"}`)
	data := external.MovieDetail{}
	json.Unmarshal([]byte(resp), &data)
	fmt.Printf("Actors: %s", data.Actors)

	movieByIDResp := &api.GetMovieByIDResponse{
		Id:        data.ImdbID,
		Title:     data.Title,
		Year:      "",
		Rated:     "",
		Genre:     "",
		Plot:      "",
		Director:  "",
		Actors:    nil,
		Language:  "",
		Country:   "",
		Type:      "",
		PosterUrl: "",
	}

	return movieByIDResp, nil
}

func (s *grpcServer) SearchMovies(ctx context.Context, req *api.SearchMoviesRequest) (*api.SearchMoviesResponse, error) {
	_, resp := external.SearchMovie(req.GetQuery(), req.GetType(), int(req.GetPage()))
	//if err == ErrIDNotFound {
	//	return nil, status.Error(codes.NotFound, "id was not found")
	//}
	//if err != nil {
	//	return nil, status.Error(codes.Internal, err.Error())
	//}

	data := external.Movies{}
	json.Unmarshal([]byte(resp), &data)
	fmt.Printf("datas[0].TotalResults: %s", data.TotalResults)

	total, err := strconv.ParseUint(data.TotalResults, 10, 64)
	if err != nil {
		fmt.Printf("SearchMovies : error converting total results : %s\n", err)
		total = 0
	}
	searchMovieResp := &api.SearchMoviesResponse{
		Movies:       nil,
		TotalResults: total,
	}

	return searchMovieResp, nil
}

//func (s *grpcServer) List(ctx context.Context, req *api.ListRequest) (*api.Activities, error) {
//	activities, err := s.Activities.List(int(req.Offset))
//	if err != nil {
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//	return &api.Activities{Activities: activities}, nil
//}

func NewGRPCServer() *grpc.Server {
	//var acc *Activities
	//var err error
	//if acc, err = NewActivities(); err != nil {
	//	log.Fatal(err)
	//}
	gsrv := grpc.NewServer()
	srv := grpcServer{
		//	Activities: acc,
	}
	api.RegisterOMDBServiceServer(gsrv, &srv)
	return gsrv
}
