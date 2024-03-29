// Code generated by ogen, DO NOT EDIT.

package ogent

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ogen-go/ogen/uri"
)

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
	}
	if prefix := s.cfg.Prefix; len(prefix) > 0 {
		if strings.HasPrefix(elem, prefix) {
			// Cut prefix from the path.
			elem = strings.TrimPrefix(elem, prefix)
		} else {
			// Prefix doesn't match.
			s.notFound(w, r)
			return
		}
	}
	if len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [1]string{}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'c': // Prefix: "cats"
				if l := len("cats"); len(elem) >= l && elem[0:l] == "cats" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handleListCatRequest([0]string{}, w, r)
					case "POST":
						s.handleCreateCatRequest([0]string{}, w, r)
					default:
						s.notAllowed(w, r, "GET,POST")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch r.Method {
						case "DELETE":
							s.handleDeleteCatRequest([1]string{
								args[0],
							}, w, r)
						case "GET":
							s.handleReadCatRequest([1]string{
								args[0],
							}, w, r)
						case "PATCH":
							s.handleUpdateCatRequest([1]string{
								args[0],
							}, w, r)
						default:
							s.notAllowed(w, r, "DELETE,GET,PATCH")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'o': // Prefix: "owner"
							if l := len("owner"); len(elem) >= l && elem[0:l] == "owner" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleReadCatOwnerRequest([1]string{
										args[0],
									}, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}
						case 't': // Prefix: "toilets"
							if l := len("toilets"); len(elem) >= l && elem[0:l] == "toilets" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleListCatToiletsRequest([1]string{
										args[0],
									}, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}
						}
					}
				}
			case 't': // Prefix: "toilets"
				if l := len("toilets"); len(elem) >= l && elem[0:l] == "toilets" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handleListToiletRequest([0]string{}, w, r)
					case "POST":
						s.handleCreateToiletRequest([0]string{}, w, r)
					default:
						s.notAllowed(w, r, "GET,POST")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch r.Method {
						case "DELETE":
							s.handleDeleteToiletRequest([1]string{
								args[0],
							}, w, r)
						case "GET":
							s.handleReadToiletRequest([1]string{
								args[0],
							}, w, r)
						case "PATCH":
							s.handleUpdateToiletRequest([1]string{
								args[0],
							}, w, r)
						default:
							s.notAllowed(w, r, "DELETE,GET,PATCH")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/cat"
						if l := len("/cat"); len(elem) >= l && elem[0:l] == "/cat" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleReadToiletCatRequest([1]string{
									args[0],
								}, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					}
				}
			case 'u': // Prefix: "users"
				if l := len("users"); len(elem) >= l && elem[0:l] == "users" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handleListUserRequest([0]string{}, w, r)
					case "POST":
						s.handleCreateUserRequest([0]string{}, w, r)
					default:
						s.notAllowed(w, r, "GET,POST")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch r.Method {
						case "DELETE":
							s.handleDeleteUserRequest([1]string{
								args[0],
							}, w, r)
						case "GET":
							s.handleReadUserRequest([1]string{
								args[0],
							}, w, r)
						case "PATCH":
							s.handleUpdateUserRequest([1]string{
								args[0],
							}, w, r)
						default:
							s.notAllowed(w, r, "DELETE,GET,PATCH")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/cats"
						if l := len("/cats"); len(elem) >= l && elem[0:l] == "/cats" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleListUserCatsRequest([1]string{
									args[0],
								}, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					}
				}
			}
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	operationID string
	pathPattern string
	count       int
	args        [1]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// PathPattern returns OpenAPI path.
func (r Route) PathPattern() string {
	return r.pathPattern
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
//
// Note: this method does not unescape path or handle reserved characters in path properly. Use FindPath instead.
func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

// FindPath finds Route for given method and URL.
func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'c': // Prefix: "cats"
				if l := len("cats"); len(elem) >= l && elem[0:l] == "cats" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = "ListCat"
						r.operationID = "listCat"
						r.pathPattern = "/cats"
						r.args = args
						r.count = 0
						return r, true
					case "POST":
						r.name = "CreateCat"
						r.operationID = "createCat"
						r.pathPattern = "/cats"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch method {
						case "DELETE":
							r.name = "DeleteCat"
							r.operationID = "deleteCat"
							r.pathPattern = "/cats/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "GET":
							r.name = "ReadCat"
							r.operationID = "readCat"
							r.pathPattern = "/cats/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "PATCH":
							r.name = "UpdateCat"
							r.operationID = "updateCat"
							r.pathPattern = "/cats/{id}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/"
						if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'o': // Prefix: "owner"
							if l := len("owner"); len(elem) >= l && elem[0:l] == "owner" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "GET":
									// Leaf: ReadCatOwner
									r.name = "ReadCatOwner"
									r.operationID = "readCatOwner"
									r.pathPattern = "/cats/{id}/owner"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}
						case 't': // Prefix: "toilets"
							if l := len("toilets"); len(elem) >= l && elem[0:l] == "toilets" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "GET":
									// Leaf: ListCatToilets
									r.name = "ListCatToilets"
									r.operationID = "listCatToilets"
									r.pathPattern = "/cats/{id}/toilets"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}
						}
					}
				}
			case 't': // Prefix: "toilets"
				if l := len("toilets"); len(elem) >= l && elem[0:l] == "toilets" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = "ListToilet"
						r.operationID = "listToilet"
						r.pathPattern = "/toilets"
						r.args = args
						r.count = 0
						return r, true
					case "POST":
						r.name = "CreateToilet"
						r.operationID = "createToilet"
						r.pathPattern = "/toilets"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch method {
						case "DELETE":
							r.name = "DeleteToilet"
							r.operationID = "deleteToilet"
							r.pathPattern = "/toilets/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "GET":
							r.name = "ReadToilet"
							r.operationID = "readToilet"
							r.pathPattern = "/toilets/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "PATCH":
							r.name = "UpdateToilet"
							r.operationID = "updateToilet"
							r.pathPattern = "/toilets/{id}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/cat"
						if l := len("/cat"); len(elem) >= l && elem[0:l] == "/cat" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: ReadToiletCat
								r.name = "ReadToiletCat"
								r.operationID = "readToiletCat"
								r.pathPattern = "/toilets/{id}/cat"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
					}
				}
			case 'u': // Prefix: "users"
				if l := len("users"); len(elem) >= l && elem[0:l] == "users" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = "ListUser"
						r.operationID = "listUser"
						r.pathPattern = "/users"
						r.args = args
						r.count = 0
						return r, true
					case "POST":
						r.name = "CreateUser"
						r.operationID = "createUser"
						r.pathPattern = "/users"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch method {
						case "DELETE":
							r.name = "DeleteUser"
							r.operationID = "deleteUser"
							r.pathPattern = "/users/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "GET":
							r.name = "ReadUser"
							r.operationID = "readUser"
							r.pathPattern = "/users/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "PATCH":
							r.name = "UpdateUser"
							r.operationID = "updateUser"
							r.pathPattern = "/users/{id}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/cats"
						if l := len("/cats"); len(elem) >= l && elem[0:l] == "/cats" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: ListUserCats
								r.name = "ListUserCats"
								r.operationID = "listUserCats"
								r.pathPattern = "/users/{id}/cats"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
					}
				}
			}
		}
	}
	return r, false
}
