// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// CreateCat implements createCat operation.
	//
	// Creates a new Cat and persists it to storage.
	//
	// POST /cats
	CreateCat(ctx context.Context, req *CreateCatReq) (CreateCatRes, error)
	// CreateToilet implements createToilet operation.
	//
	// Creates a new Toilet and persists it to storage.
	//
	// POST /toilets
	CreateToilet(ctx context.Context, req *CreateToiletReq) (CreateToiletRes, error)
	// CreateUser implements createUser operation.
	//
	// Creates a new User and persists it to storage.
	//
	// POST /users
	CreateUser(ctx context.Context, req *CreateUserReq) (CreateUserRes, error)
	// DeleteCat implements deleteCat operation.
	//
	// Deletes the Cat with the requested ID.
	//
	// DELETE /cats/{id}
	DeleteCat(ctx context.Context, params DeleteCatParams) (DeleteCatRes, error)
	// DeleteToilet implements deleteToilet operation.
	//
	// Deletes the Toilet with the requested ID.
	//
	// DELETE /toilets/{id}
	DeleteToilet(ctx context.Context, params DeleteToiletParams) (DeleteToiletRes, error)
	// DeleteUser implements deleteUser operation.
	//
	// Deletes the User with the requested ID.
	//
	// DELETE /users/{id}
	DeleteUser(ctx context.Context, params DeleteUserParams) (DeleteUserRes, error)
	// ListCat implements listCat operation.
	//
	// List Cats.
	//
	// GET /cats
	ListCat(ctx context.Context, params ListCatParams) (ListCatRes, error)
	// ListCatToilets implements listCatToilets operation.
	//
	// List attached Toilets.
	//
	// GET /cats/{id}/toilets
	ListCatToilets(ctx context.Context, params ListCatToiletsParams) (ListCatToiletsRes, error)
	// ListToilet implements listToilet operation.
	//
	// List Toilets.
	//
	// GET /toilets
	ListToilet(ctx context.Context, params ListToiletParams) (ListToiletRes, error)
	// ListUser implements listUser operation.
	//
	// List Users.
	//
	// GET /users
	ListUser(ctx context.Context, params ListUserParams) (ListUserRes, error)
	// ListUserCats implements listUserCats operation.
	//
	// List attached Cats.
	//
	// GET /users/{id}/cats
	ListUserCats(ctx context.Context, params ListUserCatsParams) (ListUserCatsRes, error)
	// ReadCat implements readCat operation.
	//
	// Finds the Cat with the requested ID and returns it.
	//
	// GET /cats/{id}
	ReadCat(ctx context.Context, params ReadCatParams) (ReadCatRes, error)
	// ReadCatOwner implements readCatOwner operation.
	//
	// Find the attached User of the Cat with the given ID.
	//
	// GET /cats/{id}/owner
	ReadCatOwner(ctx context.Context, params ReadCatOwnerParams) (ReadCatOwnerRes, error)
	// ReadToilet implements readToilet operation.
	//
	// Finds the Toilet with the requested ID and returns it.
	//
	// GET /toilets/{id}
	ReadToilet(ctx context.Context, params ReadToiletParams) (ReadToiletRes, error)
	// ReadToiletCat implements readToiletCat operation.
	//
	// Find the attached Cat of the Toilet with the given ID.
	//
	// GET /toilets/{id}/cat
	ReadToiletCat(ctx context.Context, params ReadToiletCatParams) (ReadToiletCatRes, error)
	// ReadUser implements readUser operation.
	//
	// Finds the User with the requested ID and returns it.
	//
	// GET /users/{id}
	ReadUser(ctx context.Context, params ReadUserParams) (ReadUserRes, error)
	// UpdateCat implements updateCat operation.
	//
	// Updates a Cat and persists changes to storage.
	//
	// PATCH /cats/{id}
	UpdateCat(ctx context.Context, req *UpdateCatReq, params UpdateCatParams) (UpdateCatRes, error)
	// UpdateToilet implements updateToilet operation.
	//
	// Updates a Toilet and persists changes to storage.
	//
	// PATCH /toilets/{id}
	UpdateToilet(ctx context.Context, req *UpdateToiletReq, params UpdateToiletParams) (UpdateToiletRes, error)
	// UpdateUser implements updateUser operation.
	//
	// Updates a User and persists changes to storage.
	//
	// PATCH /users/{id}
	UpdateUser(ctx context.Context, req *UpdateUserReq, params UpdateUserParams) (UpdateUserRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
