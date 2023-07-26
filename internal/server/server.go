package server

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	api "omdb/api/v1"
	external "omdb/external"
	"strconv"
	"strings"
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

	err, resp := external.GetMovieByID(req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	data := external.MovieDetail{}
	json.Unmarshal(resp, &data)

	// handle error
	if data.ImdbID == "" && data.Title == "" {
		errData := external.ErrorBody{}
		json.Unmarshal(resp, &errData)
		return nil, status.Error(codes.NotFound, errData.Error)
	}

	actors := strings.Split(data.Actors, ", ")
	movieByIDResp := &api.GetMovieByIDResponse{
		Id:        data.ImdbID,
		Title:     data.Title,
		Year:      data.Year,
		Rated:     data.Rated,
		Genre:     data.Genre,
		Plot:      data.Plot,
		Director:  data.Director,
		Actors:    actors,
		Language:  data.Language,
		Country:   data.Country,
		Type:      data.Type,
		PosterUrl: data.Poster,
	}

	return movieByIDResp, nil
}

func (s *grpcServer) SearchMovies(ctx context.Context, req *api.SearchMoviesRequest) (*api.SearchMoviesResponse, error) {
	err, resp := external.SearchMovie(req.GetQuery(), req.GetType(), int(req.GetPage()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	data := external.Movies{}
	json.Unmarshal(resp, &data)
	fmt.Printf("datas[0].TotalResults: %s", data.TotalResults)

	// handle error
	if data.TotalResults == "" && data.Response == "" {
		errData := external.ErrorBody{}
		json.Unmarshal(resp, &errData)
		return nil, status.Error(codes.NotFound, errData.Error)
	}

	total, err := strconv.ParseUint(data.TotalResults, 10, 64)
	if err != nil {
		fmt.Printf("SearchMovies : error converting total results : %s\n", err)
		total = 0
	}

	var movies []*api.MovieResult

	for _, item := range data.Search {
		movie := &api.MovieResult{
			Id:        item.ImdbID,
			Title:     item.Title,
			Year:      item.Year,
			Type:      item.Type,
			PosterUrl: item.Poster,
		}
		movies = append(movies, movie)
	}

	searchMovieResp := &api.SearchMoviesResponse{
		Movies:       movies,
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
