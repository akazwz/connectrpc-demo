// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: file/v1/file.proto

package filev1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	http "net/http"
	v1 "opendrive/gen/file/v1"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// FileServiceName is the fully-qualified name of the FileService service.
	FileServiceName = "file.v1.FileService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// FileServiceUploadFileProcedure is the fully-qualified name of the FileService's UploadFile RPC.
	FileServiceUploadFileProcedure = "/file.v1.FileService/UploadFile"
	// FileServiceDownloadFileProcedure is the fully-qualified name of the FileService's DownloadFile
	// RPC.
	FileServiceDownloadFileProcedure = "/file.v1.FileService/DownloadFile"
	// FileServiceListFilesProcedure is the fully-qualified name of the FileService's ListFiles RPC.
	FileServiceListFilesProcedure = "/file.v1.FileService/ListFiles"
	// FileServiceGetFileInfoProcedure is the fully-qualified name of the FileService's GetFileInfo RPC.
	FileServiceGetFileInfoProcedure = "/file.v1.FileService/GetFileInfo"
	// FileServiceDeleteFileProcedure is the fully-qualified name of the FileService's DeleteFile RPC.
	FileServiceDeleteFileProcedure = "/file.v1.FileService/DeleteFile"
	// FileServiceShareFileProcedure is the fully-qualified name of the FileService's ShareFile RPC.
	FileServiceShareFileProcedure = "/file.v1.FileService/ShareFile"
	// FileServiceCreateFolderProcedure is the fully-qualified name of the FileService's CreateFolder
	// RPC.
	FileServiceCreateFolderProcedure = "/file.v1.FileService/CreateFolder"
)

// FileServiceClient is a client for the file.v1.FileService service.
type FileServiceClient interface {
	// 上传文件
	UploadFile(context.Context, *connect.Request[v1.UploadFileRequest]) (*connect.Response[v1.UploadFileResponse], error)
	// 下载文件
	DownloadFile(context.Context, *connect.Request[v1.DownloadFileRequest]) (*connect.ServerStreamForClient[v1.DownloadFileResponse], error)
	// 列出用户文件
	ListFiles(context.Context, *connect.Request[v1.ListFilesRequest]) (*connect.Response[v1.ListFilesResponse], error)
	// 获取文件信息
	GetFileInfo(context.Context, *connect.Request[v1.GetFileInfoRequest]) (*connect.Response[v1.GetFileInfoResponse], error)
	// 删除文件
	DeleteFile(context.Context, *connect.Request[v1.DeleteFileRequest]) (*connect.Response[v1.DeleteFileResponse], error)
	// 分享文件
	ShareFile(context.Context, *connect.Request[v1.ShareFileRequest]) (*connect.Response[v1.ShareFileResponse], error)
	// 创建文件夹
	CreateFolder(context.Context, *connect.Request[v1.CreateFolderRequest]) (*connect.Response[v1.CreateFolderResponse], error)
}

// NewFileServiceClient constructs a client for the file.v1.FileService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewFileServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) FileServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	fileServiceMethods := v1.File_file_v1_file_proto.Services().ByName("FileService").Methods()
	return &fileServiceClient{
		uploadFile: connect.NewClient[v1.UploadFileRequest, v1.UploadFileResponse](
			httpClient,
			baseURL+FileServiceUploadFileProcedure,
			connect.WithSchema(fileServiceMethods.ByName("UploadFile")),
			connect.WithClientOptions(opts...),
		),
		downloadFile: connect.NewClient[v1.DownloadFileRequest, v1.DownloadFileResponse](
			httpClient,
			baseURL+FileServiceDownloadFileProcedure,
			connect.WithSchema(fileServiceMethods.ByName("DownloadFile")),
			connect.WithClientOptions(opts...),
		),
		listFiles: connect.NewClient[v1.ListFilesRequest, v1.ListFilesResponse](
			httpClient,
			baseURL+FileServiceListFilesProcedure,
			connect.WithSchema(fileServiceMethods.ByName("ListFiles")),
			connect.WithClientOptions(opts...),
		),
		getFileInfo: connect.NewClient[v1.GetFileInfoRequest, v1.GetFileInfoResponse](
			httpClient,
			baseURL+FileServiceGetFileInfoProcedure,
			connect.WithSchema(fileServiceMethods.ByName("GetFileInfo")),
			connect.WithClientOptions(opts...),
		),
		deleteFile: connect.NewClient[v1.DeleteFileRequest, v1.DeleteFileResponse](
			httpClient,
			baseURL+FileServiceDeleteFileProcedure,
			connect.WithSchema(fileServiceMethods.ByName("DeleteFile")),
			connect.WithClientOptions(opts...),
		),
		shareFile: connect.NewClient[v1.ShareFileRequest, v1.ShareFileResponse](
			httpClient,
			baseURL+FileServiceShareFileProcedure,
			connect.WithSchema(fileServiceMethods.ByName("ShareFile")),
			connect.WithClientOptions(opts...),
		),
		createFolder: connect.NewClient[v1.CreateFolderRequest, v1.CreateFolderResponse](
			httpClient,
			baseURL+FileServiceCreateFolderProcedure,
			connect.WithSchema(fileServiceMethods.ByName("CreateFolder")),
			connect.WithClientOptions(opts...),
		),
	}
}

// fileServiceClient implements FileServiceClient.
type fileServiceClient struct {
	uploadFile   *connect.Client[v1.UploadFileRequest, v1.UploadFileResponse]
	downloadFile *connect.Client[v1.DownloadFileRequest, v1.DownloadFileResponse]
	listFiles    *connect.Client[v1.ListFilesRequest, v1.ListFilesResponse]
	getFileInfo  *connect.Client[v1.GetFileInfoRequest, v1.GetFileInfoResponse]
	deleteFile   *connect.Client[v1.DeleteFileRequest, v1.DeleteFileResponse]
	shareFile    *connect.Client[v1.ShareFileRequest, v1.ShareFileResponse]
	createFolder *connect.Client[v1.CreateFolderRequest, v1.CreateFolderResponse]
}

// UploadFile calls file.v1.FileService.UploadFile.
func (c *fileServiceClient) UploadFile(ctx context.Context, req *connect.Request[v1.UploadFileRequest]) (*connect.Response[v1.UploadFileResponse], error) {
	return c.uploadFile.CallUnary(ctx, req)
}

// DownloadFile calls file.v1.FileService.DownloadFile.
func (c *fileServiceClient) DownloadFile(ctx context.Context, req *connect.Request[v1.DownloadFileRequest]) (*connect.ServerStreamForClient[v1.DownloadFileResponse], error) {
	return c.downloadFile.CallServerStream(ctx, req)
}

// ListFiles calls file.v1.FileService.ListFiles.
func (c *fileServiceClient) ListFiles(ctx context.Context, req *connect.Request[v1.ListFilesRequest]) (*connect.Response[v1.ListFilesResponse], error) {
	return c.listFiles.CallUnary(ctx, req)
}

// GetFileInfo calls file.v1.FileService.GetFileInfo.
func (c *fileServiceClient) GetFileInfo(ctx context.Context, req *connect.Request[v1.GetFileInfoRequest]) (*connect.Response[v1.GetFileInfoResponse], error) {
	return c.getFileInfo.CallUnary(ctx, req)
}

// DeleteFile calls file.v1.FileService.DeleteFile.
func (c *fileServiceClient) DeleteFile(ctx context.Context, req *connect.Request[v1.DeleteFileRequest]) (*connect.Response[v1.DeleteFileResponse], error) {
	return c.deleteFile.CallUnary(ctx, req)
}

// ShareFile calls file.v1.FileService.ShareFile.
func (c *fileServiceClient) ShareFile(ctx context.Context, req *connect.Request[v1.ShareFileRequest]) (*connect.Response[v1.ShareFileResponse], error) {
	return c.shareFile.CallUnary(ctx, req)
}

// CreateFolder calls file.v1.FileService.CreateFolder.
func (c *fileServiceClient) CreateFolder(ctx context.Context, req *connect.Request[v1.CreateFolderRequest]) (*connect.Response[v1.CreateFolderResponse], error) {
	return c.createFolder.CallUnary(ctx, req)
}

// FileServiceHandler is an implementation of the file.v1.FileService service.
type FileServiceHandler interface {
	// 上传文件
	UploadFile(context.Context, *connect.Request[v1.UploadFileRequest]) (*connect.Response[v1.UploadFileResponse], error)
	// 下载文件
	DownloadFile(context.Context, *connect.Request[v1.DownloadFileRequest], *connect.ServerStream[v1.DownloadFileResponse]) error
	// 列出用户文件
	ListFiles(context.Context, *connect.Request[v1.ListFilesRequest]) (*connect.Response[v1.ListFilesResponse], error)
	// 获取文件信息
	GetFileInfo(context.Context, *connect.Request[v1.GetFileInfoRequest]) (*connect.Response[v1.GetFileInfoResponse], error)
	// 删除文件
	DeleteFile(context.Context, *connect.Request[v1.DeleteFileRequest]) (*connect.Response[v1.DeleteFileResponse], error)
	// 分享文件
	ShareFile(context.Context, *connect.Request[v1.ShareFileRequest]) (*connect.Response[v1.ShareFileResponse], error)
	// 创建文件夹
	CreateFolder(context.Context, *connect.Request[v1.CreateFolderRequest]) (*connect.Response[v1.CreateFolderResponse], error)
}

// NewFileServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewFileServiceHandler(svc FileServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	fileServiceMethods := v1.File_file_v1_file_proto.Services().ByName("FileService").Methods()
	fileServiceUploadFileHandler := connect.NewUnaryHandler(
		FileServiceUploadFileProcedure,
		svc.UploadFile,
		connect.WithSchema(fileServiceMethods.ByName("UploadFile")),
		connect.WithHandlerOptions(opts...),
	)
	fileServiceDownloadFileHandler := connect.NewServerStreamHandler(
		FileServiceDownloadFileProcedure,
		svc.DownloadFile,
		connect.WithSchema(fileServiceMethods.ByName("DownloadFile")),
		connect.WithHandlerOptions(opts...),
	)
	fileServiceListFilesHandler := connect.NewUnaryHandler(
		FileServiceListFilesProcedure,
		svc.ListFiles,
		connect.WithSchema(fileServiceMethods.ByName("ListFiles")),
		connect.WithHandlerOptions(opts...),
	)
	fileServiceGetFileInfoHandler := connect.NewUnaryHandler(
		FileServiceGetFileInfoProcedure,
		svc.GetFileInfo,
		connect.WithSchema(fileServiceMethods.ByName("GetFileInfo")),
		connect.WithHandlerOptions(opts...),
	)
	fileServiceDeleteFileHandler := connect.NewUnaryHandler(
		FileServiceDeleteFileProcedure,
		svc.DeleteFile,
		connect.WithSchema(fileServiceMethods.ByName("DeleteFile")),
		connect.WithHandlerOptions(opts...),
	)
	fileServiceShareFileHandler := connect.NewUnaryHandler(
		FileServiceShareFileProcedure,
		svc.ShareFile,
		connect.WithSchema(fileServiceMethods.ByName("ShareFile")),
		connect.WithHandlerOptions(opts...),
	)
	fileServiceCreateFolderHandler := connect.NewUnaryHandler(
		FileServiceCreateFolderProcedure,
		svc.CreateFolder,
		connect.WithSchema(fileServiceMethods.ByName("CreateFolder")),
		connect.WithHandlerOptions(opts...),
	)
	return "/file.v1.FileService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case FileServiceUploadFileProcedure:
			fileServiceUploadFileHandler.ServeHTTP(w, r)
		case FileServiceDownloadFileProcedure:
			fileServiceDownloadFileHandler.ServeHTTP(w, r)
		case FileServiceListFilesProcedure:
			fileServiceListFilesHandler.ServeHTTP(w, r)
		case FileServiceGetFileInfoProcedure:
			fileServiceGetFileInfoHandler.ServeHTTP(w, r)
		case FileServiceDeleteFileProcedure:
			fileServiceDeleteFileHandler.ServeHTTP(w, r)
		case FileServiceShareFileProcedure:
			fileServiceShareFileHandler.ServeHTTP(w, r)
		case FileServiceCreateFolderProcedure:
			fileServiceCreateFolderHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedFileServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedFileServiceHandler struct{}

func (UnimplementedFileServiceHandler) UploadFile(context.Context, *connect.Request[v1.UploadFileRequest]) (*connect.Response[v1.UploadFileResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("file.v1.FileService.UploadFile is not implemented"))
}

func (UnimplementedFileServiceHandler) DownloadFile(context.Context, *connect.Request[v1.DownloadFileRequest], *connect.ServerStream[v1.DownloadFileResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("file.v1.FileService.DownloadFile is not implemented"))
}

func (UnimplementedFileServiceHandler) ListFiles(context.Context, *connect.Request[v1.ListFilesRequest]) (*connect.Response[v1.ListFilesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("file.v1.FileService.ListFiles is not implemented"))
}

func (UnimplementedFileServiceHandler) GetFileInfo(context.Context, *connect.Request[v1.GetFileInfoRequest]) (*connect.Response[v1.GetFileInfoResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("file.v1.FileService.GetFileInfo is not implemented"))
}

func (UnimplementedFileServiceHandler) DeleteFile(context.Context, *connect.Request[v1.DeleteFileRequest]) (*connect.Response[v1.DeleteFileResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("file.v1.FileService.DeleteFile is not implemented"))
}

func (UnimplementedFileServiceHandler) ShareFile(context.Context, *connect.Request[v1.ShareFileRequest]) (*connect.Response[v1.ShareFileResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("file.v1.FileService.ShareFile is not implemented"))
}

func (UnimplementedFileServiceHandler) CreateFolder(context.Context, *connect.Request[v1.CreateFolderRequest]) (*connect.Response[v1.CreateFolderResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("file.v1.FileService.CreateFolder is not implemented"))
}
