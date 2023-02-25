package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/Binbiubiubiu/cnpmcore/core/entity"
	"github.com/Binbiubiubiu/cnpmcore/global"
	"github.com/Binbiubiubiu/cnpmcore/repository/model"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type TaskRepository struct{}

func (r *TaskRepository) SaveTask(task *entity.Task) (err error) {
	isCreate := lo.IsEmpty(task.ID)
	if isCreate {
		m := &model.Task{}
		tx := global.DB.First(m, task.ID)
		if err = tx.Error; err != nil {
			return
		}
		task.CreateAt = m.CreateAt
	}
	err = global.DB.Save(task).Error
	if isCreate && err != nil && task.BizId != "" {
		err = fmt.Errorf("[TaskRepository] insert Task failed:%s", err.Error())
		global.ZLog.Warn(err.Error())
		m := &model.Task{}
		err = global.DB.Where(&model.Task{BizId: task.BizId}).First(m).Error
		if err == nil {
			task.ID = m.ID
			task.TaskId = m.TaskId
			return
		}
	}
	return
}

func (r *TaskRepository) IdempotentSaveTask(task *entity.Task, condition *entity.TaskUpdateCondition) (bool, error) {
	if lo.IsEmpty(task.ID) {
		return false, errors.New("task have no save")
	}
	tx := global.DB.Where(&model.Task{TaskId: condition.TaskId, Attempts: condition.Attempts}).Updates(task)
	if err := tx.Error; err != nil {
		return false, err
	}
	return tx.RowsAffected == 1, nil
}

func (r *TaskRepository) SaveTaskToHistory(task *entity.Task) (err error) {
	m := &model.Task{}
	err = global.DB.First(m, task.ID).Error
	if err != nil {
		return
	}
	history := &model.HistoryTask{}
	err = global.DB.Where(&model.HistoryTask{TaskId: task.TaskId}).First(history).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	err = global.DB.Model(history).Save(task).Error
	if err != nil {
		return
	}
	return global.DB.Delete(m).Error
}

// TODO: 需要适配
func (r *TaskRepository) FindTask(taskId string) (*entity.Task, error) {
	m := &model.Task{}
	err := global.DB.Where(&model.Task{TaskId: taskId}).First(m).Error
	if err == nil {
		return entity.CreateTaskFromModel(m), nil
	}
	err = global.DB.Where(&model.HistoryTask{TaskId: taskId}).First(m).Error
	if err == nil {
		return entity.CreateTaskFromModel(m), nil
	}
	return nil, err
}

func (r *TaskRepository) FindTaskByBizId(bizId string) (*entity.Task, error) {
	m := &model.Task{}
	tx := global.DB.Where(&model.Task{BizId: bizId}).First(m)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return entity.CreateTaskFromModel(m), nil
}

func (r *TaskRepository) FindTasks(taskIds *[]string) ([]*entity.Task, error) {
	var list []*model.HistoryTask
	tx := global.DB.Where("taskId in ?", taskIds).Find(&list)
	if err := tx.Error; err != nil {
		return nil, err
	}

	return lo.Map(list, func(m *model.HistoryTask, i int) *entity.Task {
		return entity.CreateTaskFromHistoryModel(m)
	}), nil
}

func (r *TaskRepository) FindTasksByCondition(where *struct {
	TargetName string
	State      string
	Type       string
}) ([]*entity.Task, error) {
	var list []*model.Task
	tx := global.DB.Where(where).Find(&list)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return lo.Map(list, func(m *model.Task, i int) *entity.Task {
		return entity.CreateTaskFromModel(m)
	}), nil
}

func (r *TaskRepository) FindTaskByTargetName(targetName string, ty string, args ...string) (*entity.Task, error) {
	where := &model.Task{TargetName: targetName, Type: ty}
	if len(args) > 0 && lo.IsNotEmpty(args[0]) {
		where.State = args[0]
	}
	m := &model.Task{}
	tx := global.DB.Where(where).First(m)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return entity.CreateTaskFromModel(m), nil
}

func (r *TaskRepository) FindTimeoutTasks(taskState string, timeout int) ([]*entity.Task, error) {
	timeoutDate := time.Now()
	timeoutDate.Add(-time.Duration(timeout))
	var list []*model.Task
	tx := global.DB.Where("state = ? and update_at < ?", taskState, timeoutDate).Limit(1000).Find(&list)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return lo.Map(list, func(m *model.Task, i int) *entity.Task {
		return entity.CreateTaskFromModel(m)
	}), nil
}

func (r *TaskRepository) FindTaskByAuthorIpAndType(authorIp string, ty string) ([]*entity.Task, error) {
	var list []*model.Task
	tx := global.DB.Where(&model.Task{
		Type:     ty,
		AuthorIp: authorIp,
	}).Limit(1000).Find(&list)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return lo.Map(list, func(m *model.Task, i int) *entity.Task {
		return entity.CreateTaskFromModel(m)
	}), nil
}
