package intellij

import (
	"fmt"
	"strings"
	"sync"

	"github.com/davido912/cshex-utils/internal/teleport"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type Manager struct {
	tsh *teleport.Teleport
}

func NewManager(tsh *teleport.Teleport) *Manager {
	return &Manager{
		tsh: tsh,
	}
}

func (m *Manager) CreateDatasources(tshEnvs []string, projectRootPath string, dbNames ...string) {
	datasourceXmlConfigs := make([]DatasourceXmlConfig, 0)

	for _, tshEnv := range tshEnvs {
		log.Info().Msgf("Creating datasources for %s environment", tshEnv)
		m.tsh.SetEnv(tshEnv)
		dbListResult := m.tsh.ListDb(dbNames...)
		var wg sync.WaitGroup

		for _, db := range dbListResult {
			wg.Add(1)

			go func(dbName string) {
				defer wg.Done()
				m.tsh.DbConnect(dbName)
				dbConf := m.tsh.DbConfig(dbName)

				xmlConf := DatasourceXmlConfig{
					Name:       fmt.Sprintf("[%s] %s", string(m.tsh.GetEnv()), dbConf.Name),
					Uuid:       uuid.New(),
					Host:       dbConf.Host,
					Port:       dbConf.Port,
					Database:   strings.Replace(dbConf.Database, "_read", "", 1),
					CaCert:     dbConf.Ca,
					ClientCert: dbConf.Cert,
					ClientKey:  dbConf.Key,
				}

				datasourceXmlConfigs = append(datasourceXmlConfigs, xmlConf)

			}(db.Metadata.Name)
		}

		wg.Wait()
	}

	CreateDatasourcesLocalXml(datasourceXmlConfigs, projectRootPath)
	CreateDatasourcesXml(datasourceXmlConfigs, projectRootPath)
}
