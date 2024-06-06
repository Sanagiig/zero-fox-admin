// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	sysdept "zero-fox-admin/api/admin/internal/handler/sys/dept"
	sysdict_item "zero-fox-admin/api/admin/internal/handler/sys/dict_item"
	sysdict_type "zero-fox-admin/api/admin/internal/handler/sys/dict_type"
	syslog "zero-fox-admin/api/admin/internal/handler/sys/log"
	sysmenu "zero-fox-admin/api/admin/internal/handler/sys/menu"
	syspost "zero-fox-admin/api/admin/internal/handler/sys/post"
	sysrole "zero-fox-admin/api/admin/internal/handler/sys/role"
	sysupload "zero-fox-admin/api/admin/internal/handler/sys/upload"
	sysuser "zero-fox-admin/api/admin/internal/handler/sys/user"
	"zero-fox-admin/api/admin/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckUrl},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/addDept",
					Handler: sysdept.AddDeptHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/deleteDept",
					Handler: sysdept.DeleteDeptHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/queryDeptDetail",
					Handler: sysdept.QueryDeptDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/queryDeptList",
					Handler: sysdept.QueryDeptListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/updateDept",
					Handler: sysdept.UpdateDeptHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/updateDeptStatus",
					Handler: sysdept.UpdateDeptStatusHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/sys/dept"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckUrl},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/addDictItem",
					Handler: sysdict_item.AddDictItemHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/deleteDictItem",
					Handler: sysdict_item.DeleteDictItemHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/queryDictItemDetail",
					Handler: sysdict_item.QueryDictItemDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/queryDictItemList",
					Handler: sysdict_item.QueryDictItemListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/updateDictItem",
					Handler: sysdict_item.UpdateDictItemHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/updateDictItemStatus",
					Handler: sysdict_item.UpdateDictItemStatusHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/sys/dictItem"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckUrl},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/addDictType",
					Handler: sysdict_type.AddDictTypeHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/deleteDictType",
					Handler: sysdict_type.DeleteDictTypeHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/queryDictTypeDetail",
					Handler: sysdict_type.QueryDictTypeDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/queryDictTypeList",
					Handler: sysdict_type.QueryDictTypeListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/updateDictType",
					Handler: sysdict_type.UpdateDictTypeHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/updateDictTypeStatus",
					Handler: sysdict_type.UpdateDictTypeStatusHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/sys/dictType"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckUrl},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/deleteLoginLog",
					Handler: syslog.DeleteLoginLogHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/queryLoginLogDetail",
					Handler: syslog.QueryLoginLogDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/queryLoginLogList",
					Handler: syslog.QueryLoginLogListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/queryStatisticsLoginLog",
					Handler: syslog.QueryStatisticsLoginLogHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/sys/log"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckUrl},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/deleteOperateLog",
					Handler: syslog.DeleteOperateLogHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/queryOperateLogDetail",
					Handler: syslog.QueryOperateLogDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/queryOperateLogList",
					Handler: syslog.QueryOperateLogListHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/sys/log"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckUrl},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/addMenu",
					Handler: sysmenu.AddMenuHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/deleteMenu",
					Handler: sysmenu.DeleteMenuHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/queryMenuDetail",
					Handler: sysmenu.QueryMenuDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/queryMenuList",
					Handler: sysmenu.QueryMenuListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/updateMenu",
					Handler: sysmenu.UpdateMenuHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/updateMenuStatus",
					Handler: sysmenu.UpdateMenuStatusHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/sys/menu"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckUrl},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/addPost",
					Handler: syspost.AddPostHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/deletePost",
					Handler: syspost.DeletePostHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/queryPostDetail",
					Handler: syspost.QueryPostDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/queryPostList",
					Handler: syspost.QueryPostListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/updatePost",
					Handler: syspost.UpdatePostHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/updatePostStatus",
					Handler: syspost.UpdatePostStatusHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/sys/post"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckUrl},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/addRole",
					Handler: sysrole.AddRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/cancelAuthorization",
					Handler: sysrole.CancelAuthorizationHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/deleteRole",
					Handler: sysrole.DeleteRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/queryRoleDetail",
					Handler: sysrole.QueryRoleDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/queryRoleList",
					Handler: sysrole.QueryRoleListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/queryRoleMenuList",
					Handler: sysrole.QueryRoleMenuListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/queryRoleUserList",
					Handler: sysrole.QueryRoleUserListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/updateRole",
					Handler: sysrole.UpdateRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/updateRoleMenuList",
					Handler: sysrole.UpdateRoleMenuListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/updateRoleStatus",
					Handler: sysrole.UpdateRoleStatusHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/sys/role"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckUrl},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/upload",
					Handler: sysupload.UploadHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/sys"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckUrl},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/addUser",
					Handler: sysuser.AddUserHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/deleteUser",
					Handler: sysuser.DeleteUserHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/info",
					Handler: sysuser.UserInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/queryDeptAndPostList",
					Handler: sysuser.QueryDeptAndPostListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/queryUserDetail",
					Handler: sysuser.QueryUserDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/queryUserList",
					Handler: sysuser.QueryUserListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/queryUserRoleList",
					Handler: sysuser.QueryUserRoleListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/reSetPassword",
					Handler: sysuser.ReSetPasswordHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/updateUser",
					Handler: sysuser.UpdateUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/updateUserRoleList",
					Handler: sysuser.UpdateUserRoleListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/updateUserStatus",
					Handler: sysuser.UpdateUserStatusHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/sys/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/user/login",
				Handler: sysuser.UserLoginHandler(serverCtx),
			},
		},
	)
}