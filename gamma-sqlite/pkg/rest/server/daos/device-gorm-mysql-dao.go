package daos

import (
	"errors"
	"github.com/azar-intelops/microservice-demo/gamma-sqlite/pkg/rest/server/daos/clients/sqls"
	"github.com/azar-intelops/microservice-demo/gamma-sqlite/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type DeviceDao struct {
	db *gorm.DB
}

func NewDeviceDao() (*DeviceDao, error) {
	sqlClient, err := sqls.InitGormMySQLDB()
	if err != nil {
		return nil, err
	}
	err = sqlClient.DB.AutoMigrate(models.Device{})
	if err != nil {
		return nil, err
	}
	return &DeviceDao{
		db: sqlClient.DB,
	}, nil
}

func (deviceDao *DeviceDao) CreateDevice(m *models.Device) (*models.Device, error) {
	if err := deviceDao.db.Create(&m).Error; err != nil {
		log.Debugf("failed to create device: %v", err)
		return nil, err
	}

	log.Debugf("device created")
	return m, nil
}

func (deviceDao *DeviceDao) UpdateDevice(id int64, m *models.Device) (*models.Device, error) {
	if id == 0 {
		return nil, errors.New("invalid device ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	var device *models.Device
	if err := deviceDao.db.Where("id = ?", id).First(&device).Error; err != nil {
		log.Debugf("failed to find device for update: %v", err)
		return nil, err
	}

	if err := deviceDao.db.Save(&m).Error; err != nil {
		log.Debugf("failed to update device: %v", err)
		return nil, err
	}
	log.Debugf("device updated")
	return m, nil
}

func (deviceDao *DeviceDao) DeleteDevice(id int64) error {
	var m *models.Device
	if err := deviceDao.db.Where("id = ?", id).Delete(&m).Error; err != nil {
		log.Debugf("failed to delete device: %v", err)
		return err
	}

	log.Debugf("device deleted")
	return nil
}

func (deviceDao *DeviceDao) ListDevices() ([]*models.Device, error) {
	var devices []*models.Device
	if err := deviceDao.db.Find(&devices).Error; err != nil {
		log.Debugf("failed to list devices: %v", err)
		return nil, err
	}

	log.Debugf("device listed")
	return devices, nil
}

func (deviceDao *DeviceDao) GetDevice(id int64) (*models.Device, error) {
	var m *models.Device
	if err := deviceDao.db.Where("id = ?", id).First(&m).Error; err != nil {
		log.Debugf("failed to get device: %v", err)
		return nil, err
	}
	log.Debugf("device retrieved")
	return m, nil
}
