package project

import (
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"

    pb "github.com/dinhtp/lets-go-pbtype/project"
)

func validateOne(r *pb.OneProjectRequest) error {
    if r.GetId() == "" {
        return status.Error(codes.InvalidArgument, "Project ID is required")
    }

    return nil
}
