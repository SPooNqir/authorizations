// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package authorizations

import (
	context "context"
	users "github.com/slavayssiere-spoon/users"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthorizationsClient is the client API for Authorizations service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorizationsClient interface {
	Get(ctx context.Context, in *users.User, opts ...grpc.CallOption) (*Authorizations, error)
	GetDescriptions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*IAMDescription, error)
	Add(ctx context.Context, in *Authorization, opts ...grpc.CallOption) (*Authorization, error)
	Remove(ctx context.Context, in *Authorization, opts ...grpc.CallOption) (*Authorization, error)
	AddSuperUser(ctx context.Context, in *users.User, opts ...grpc.CallOption) (*Authorizations, error)
	RemoveSuperUser(ctx context.Context, in *users.User, opts ...grpc.CallOption) (*Authorizations, error)
	GetSuperUsers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*users.Users, error)
	AddRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error)
	RemoveRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error)
	GetRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error)
	GetRoleForEmail(ctx context.Context, in *UserGroup, opts ...grpc.CallOption) (*Role, error)
	UpdateRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error)
	GetAllRoles(ctx context.Context, in *Roles, opts ...grpc.CallOption) (*Roles, error)
	AddAuthorizationToRole(ctx context.Context, in *AuthRole, opts ...grpc.CallOption) (*Role, error)
	RemoveAuthorizationToRole(ctx context.Context, in *AuthRole, opts ...grpc.CallOption) (*Role, error)
}

type authorizationsClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorizationsClient(cc grpc.ClientConnInterface) AuthorizationsClient {
	return &authorizationsClient{cc}
}

func (c *authorizationsClient) Get(ctx context.Context, in *users.User, opts ...grpc.CallOption) (*Authorizations, error) {
	out := new(Authorizations)
	err := c.cc.Invoke(ctx, "/authorizations.authorizations/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationsClient) GetDescriptions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*IAMDescription, error) {
	out := new(IAMDescription)
	err := c.cc.Invoke(ctx, "/authorizations.authorizations/GetDescriptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationsClient) Add(ctx context.Context, in *Authorization, opts ...grpc.CallOption) (*Authorization, error) {
	out := new(Authorization)
	err := c.cc.Invoke(ctx, "/authorizations.authorizations/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationsClient) Remove(ctx context.Context, in *Authorization, opts ...grpc.CallOption) (*Authorization, error) {
	out := new(Authorization)
	err := c.cc.Invoke(ctx, "/authorizations.authorizations/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationsClient) AddSuperUser(ctx context.Context, in *users.User, opts ...grpc.CallOption) (*Authorizations, error) {
	out := new(Authorizations)
	err := c.cc.Invoke(ctx, "/authorizations.authorizations/AddSuperUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationsClient) RemoveSuperUser(ctx context.Context, in *users.User, opts ...grpc.CallOption) (*Authorizations, error) {
	out := new(Authorizations)
	err := c.cc.Invoke(ctx, "/authorizations.authorizations/RemoveSuperUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationsClient) GetSuperUsers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*users.Users, error) {
	out := new(users.Users)
	err := c.cc.Invoke(ctx, "/authorizations.authorizations/GetSuperUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationsClient) AddRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, "/authorizations.authorizations/AddRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationsClient) RemoveRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, "/authorizations.authorizations/RemoveRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationsClient) GetRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, "/authorizations.authorizations/GetRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationsClient) GetRoleForEmail(ctx context.Context, in *UserGroup, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, "/authorizations.authorizations/GetRoleForEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationsClient) UpdateRole(ctx context.Context, in *Role, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, "/authorizations.authorizations/UpdateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationsClient) GetAllRoles(ctx context.Context, in *Roles, opts ...grpc.CallOption) (*Roles, error) {
	out := new(Roles)
	err := c.cc.Invoke(ctx, "/authorizations.authorizations/GetAllRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationsClient) AddAuthorizationToRole(ctx context.Context, in *AuthRole, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, "/authorizations.authorizations/AddAuthorizationToRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationsClient) RemoveAuthorizationToRole(ctx context.Context, in *AuthRole, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, "/authorizations.authorizations/RemoveAuthorizationToRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationsServer is the server API for Authorizations service.
// All implementations must embed UnimplementedAuthorizationsServer
// for forward compatibility
type AuthorizationsServer interface {
	Get(context.Context, *users.User) (*Authorizations, error)
	GetDescriptions(context.Context, *emptypb.Empty) (*IAMDescription, error)
	Add(context.Context, *Authorization) (*Authorization, error)
	Remove(context.Context, *Authorization) (*Authorization, error)
	AddSuperUser(context.Context, *users.User) (*Authorizations, error)
	RemoveSuperUser(context.Context, *users.User) (*Authorizations, error)
	GetSuperUsers(context.Context, *emptypb.Empty) (*users.Users, error)
	AddRole(context.Context, *Role) (*Role, error)
	RemoveRole(context.Context, *Role) (*Role, error)
	GetRole(context.Context, *Role) (*Role, error)
	GetRoleForEmail(context.Context, *UserGroup) (*Role, error)
	UpdateRole(context.Context, *Role) (*Role, error)
	GetAllRoles(context.Context, *Roles) (*Roles, error)
	AddAuthorizationToRole(context.Context, *AuthRole) (*Role, error)
	RemoveAuthorizationToRole(context.Context, *AuthRole) (*Role, error)
	mustEmbedUnimplementedAuthorizationsServer()
}

// UnimplementedAuthorizationsServer must be embedded to have forward compatible implementations.
type UnimplementedAuthorizationsServer struct {
}

func (UnimplementedAuthorizationsServer) Get(context.Context, *users.User) (*Authorizations, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAuthorizationsServer) GetDescriptions(context.Context, *emptypb.Empty) (*IAMDescription, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDescriptions not implemented")
}
func (UnimplementedAuthorizationsServer) Add(context.Context, *Authorization) (*Authorization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedAuthorizationsServer) Remove(context.Context, *Authorization) (*Authorization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedAuthorizationsServer) AddSuperUser(context.Context, *users.User) (*Authorizations, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSuperUser not implemented")
}
func (UnimplementedAuthorizationsServer) RemoveSuperUser(context.Context, *users.User) (*Authorizations, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSuperUser not implemented")
}
func (UnimplementedAuthorizationsServer) GetSuperUsers(context.Context, *emptypb.Empty) (*users.Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSuperUsers not implemented")
}
func (UnimplementedAuthorizationsServer) AddRole(context.Context, *Role) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRole not implemented")
}
func (UnimplementedAuthorizationsServer) RemoveRole(context.Context, *Role) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveRole not implemented")
}
func (UnimplementedAuthorizationsServer) GetRole(context.Context, *Role) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRole not implemented")
}
func (UnimplementedAuthorizationsServer) GetRoleForEmail(context.Context, *UserGroup) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoleForEmail not implemented")
}
func (UnimplementedAuthorizationsServer) UpdateRole(context.Context, *Role) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (UnimplementedAuthorizationsServer) GetAllRoles(context.Context, *Roles) (*Roles, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRoles not implemented")
}
func (UnimplementedAuthorizationsServer) AddAuthorizationToRole(context.Context, *AuthRole) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAuthorizationToRole not implemented")
}
func (UnimplementedAuthorizationsServer) RemoveAuthorizationToRole(context.Context, *AuthRole) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveAuthorizationToRole not implemented")
}
func (UnimplementedAuthorizationsServer) mustEmbedUnimplementedAuthorizationsServer() {}

// UnsafeAuthorizationsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthorizationsServer will
// result in compilation errors.
type UnsafeAuthorizationsServer interface {
	mustEmbedUnimplementedAuthorizationsServer()
}

func RegisterAuthorizationsServer(s grpc.ServiceRegistrar, srv AuthorizationsServer) {
	s.RegisterService(&Authorizations_ServiceDesc, srv)
}

func _Authorizations_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(users.User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorizations.authorizations/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationsServer).Get(ctx, req.(*users.User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorizations_GetDescriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationsServer).GetDescriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorizations.authorizations/GetDescriptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationsServer).GetDescriptions(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorizations_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Authorization)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationsServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorizations.authorizations/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationsServer).Add(ctx, req.(*Authorization))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorizations_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Authorization)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationsServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorizations.authorizations/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationsServer).Remove(ctx, req.(*Authorization))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorizations_AddSuperUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(users.User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationsServer).AddSuperUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorizations.authorizations/AddSuperUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationsServer).AddSuperUser(ctx, req.(*users.User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorizations_RemoveSuperUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(users.User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationsServer).RemoveSuperUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorizations.authorizations/RemoveSuperUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationsServer).RemoveSuperUser(ctx, req.(*users.User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorizations_GetSuperUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationsServer).GetSuperUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorizations.authorizations/GetSuperUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationsServer).GetSuperUsers(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorizations_AddRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Role)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationsServer).AddRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorizations.authorizations/AddRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationsServer).AddRole(ctx, req.(*Role))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorizations_RemoveRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Role)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationsServer).RemoveRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorizations.authorizations/RemoveRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationsServer).RemoveRole(ctx, req.(*Role))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorizations_GetRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Role)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationsServer).GetRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorizations.authorizations/GetRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationsServer).GetRole(ctx, req.(*Role))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorizations_GetRoleForEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationsServer).GetRoleForEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorizations.authorizations/GetRoleForEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationsServer).GetRoleForEmail(ctx, req.(*UserGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorizations_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Role)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationsServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorizations.authorizations/UpdateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationsServer).UpdateRole(ctx, req.(*Role))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorizations_GetAllRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Roles)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationsServer).GetAllRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorizations.authorizations/GetAllRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationsServer).GetAllRoles(ctx, req.(*Roles))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorizations_AddAuthorizationToRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRole)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationsServer).AddAuthorizationToRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorizations.authorizations/AddAuthorizationToRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationsServer).AddAuthorizationToRole(ctx, req.(*AuthRole))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorizations_RemoveAuthorizationToRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRole)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationsServer).RemoveAuthorizationToRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorizations.authorizations/RemoveAuthorizationToRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationsServer).RemoveAuthorizationToRole(ctx, req.(*AuthRole))
	}
	return interceptor(ctx, in, info, handler)
}

// Authorizations_ServiceDesc is the grpc.ServiceDesc for Authorizations service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Authorizations_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authorizations.authorizations",
	HandlerType: (*AuthorizationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Authorizations_Get_Handler,
		},
		{
			MethodName: "GetDescriptions",
			Handler:    _Authorizations_GetDescriptions_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Authorizations_Add_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Authorizations_Remove_Handler,
		},
		{
			MethodName: "AddSuperUser",
			Handler:    _Authorizations_AddSuperUser_Handler,
		},
		{
			MethodName: "RemoveSuperUser",
			Handler:    _Authorizations_RemoveSuperUser_Handler,
		},
		{
			MethodName: "GetSuperUsers",
			Handler:    _Authorizations_GetSuperUsers_Handler,
		},
		{
			MethodName: "AddRole",
			Handler:    _Authorizations_AddRole_Handler,
		},
		{
			MethodName: "RemoveRole",
			Handler:    _Authorizations_RemoveRole_Handler,
		},
		{
			MethodName: "GetRole",
			Handler:    _Authorizations_GetRole_Handler,
		},
		{
			MethodName: "GetRoleForEmail",
			Handler:    _Authorizations_GetRoleForEmail_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _Authorizations_UpdateRole_Handler,
		},
		{
			MethodName: "GetAllRoles",
			Handler:    _Authorizations_GetAllRoles_Handler,
		},
		{
			MethodName: "AddAuthorizationToRole",
			Handler:    _Authorizations_AddAuthorizationToRole_Handler,
		},
		{
			MethodName: "RemoveAuthorizationToRole",
			Handler:    _Authorizations_RemoveAuthorizationToRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authorizations.proto",
}