package service

import (
	"context"
	"github.com/Muxi-Backend-Classroom/workbench/microservice/project/dao"
	pb "github.com/Muxi-Backend-Classroom/workbench/microservice/project/proto"
	"github.com/Muxi-Backend-Classroom/workbench/pkg/errno"
)

// ...获取垃圾箱列表
func (s *ProjectService) GetTrashbin(ctx context.Context, req *pb.GetTrashbinRequest, resp *pb.GetTrashbinResponse) error {
	doc, err := dao.ListTrashbinForDoc(req.Offset, req.Limit, req.ProjectId)
	if err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	file, err := dao.ListTrashbinForFile(req.Offset, req.Limit, req.ProjectId)
	if err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}

	docfolder, err := dao.ListTrashbinForMdFolder(req.Offset, req.Limit, req.ProjectId)
	if err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}
	filefolder, err := dao.ListTrashbinForFileFolder(req.Offset, req.Limit, req.ProjectId)
	if err != nil {
		return errno.ServerErr(errno.ErrDatabase, err.Error())
	}

	resList := make([]*pb.Trashbin, len(doc)+len(file)+len(docfolder)+len(filefolder))
	var i int
	for _, item := range doc {
		resList[i] = &pb.Trashbin{
			Id:   uint32(item.ID),
			Type: 1,
			Name: item.FileName,
		}
		i++
	}
	for _, item := range file {
		resList[i] = &pb.Trashbin{
			Id:   uint32(item.ID),
			Type: 2,
			Name: item.FileName,
		}
		i++
	}
	for _, item := range docfolder {
		resList[i] = &pb.Trashbin{
			Id:   uint32(item.ID),
			Type: 3,
			Name: item.Name,
		}
		i++
	}
	for _, item := range filefolder {
		resList[i] = &pb.Trashbin{
			Id:   uint32(item.ID),
			Type: 4,
			Name: item.Name,
		}
		i++
	}

	resp.List = resList
	return nil
}
