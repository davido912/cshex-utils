package intellij

import (
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
)

type DatasourceXmlConfig struct {
	Name       string
	Uuid       uuid.UUID
	Host       string
	Port       int
	Database   string
	CaCert     string
	ClientCert string
	ClientKey  string
}

func DatasourceLocalXmlDefinition(config DatasourceXmlConfig) string {
	return fmt.Sprintf(`<data-source name="%s" uuid="%s">
	<database-info product="PostgreSQL" version="12.14" jdbc-version="4.2" driver-name="PostgreSQL JDBC Driver" driver-version="42.6.0" dbms="POSTGRES" exact-version="12.14" exact-driver-version="42.6">
		<identifier-quote-string>&quot;</identifier-quote-string>
	</database-info>
	<case-sensitivity plain-identifiers="lower" quoted-identifiers="exact" />
	<secret-storage>master_key</secret-storage>
	<user-name>teleport_admin</user-name>
	<schema-mapping>
		<introspection-scope>
			<node negative="1">
				<node kind="database" qname="@">
					<node kind="schema" qname="@" />
				</node>
			</node>
		</introspection-scope>
	</schema-mapping>
	<ssl-config>
		<ca-cert>%s</ca-cert>
		<client-cert>%s</client-cert>
		<client-key>%s</client-key>
		<enabled>true</enabled>
		<mode>REQUIRE</mode>
	</ssl-config>
</data-source>
	`, config.Name, config.Uuid.String(), config.CaCert, config.ClientCert, config.ClientKey)
}

func DatasourceXmlDefinition(config DatasourceXmlConfig) string {
	return fmt.Sprintf(`<data-source source="LOCAL" name="%s" uuid="%s">
	<driver-ref>postgresql</driver-ref>
	<synchronize>true</synchronize>
	<jdbc-driver>org.postgresql.Driver</jdbc-driver>
	<jdbc-url>jdbc:postgresql://%s:%d/%s</jdbc-url>
	<jdbc-additional-properties>
		<property name="com.intellij.clouds.kubernetes.db.host.port" />
		<property name="com.intellij.clouds.kubernetes.db.enabled" value="false" />
		<property name="com.intellij.clouds.kubernetes.db.resource.type" value="Deployment" />
		<property name="com.intellij.clouds.kubernetes.db.container.port" />
	</jdbc-additional-properties>
	<working-dir>$ProjectFileDir$</working-dir>
</data-source>
	`, config.Name, config.Uuid.String(), config.Host, config.Port, config.Database)
}

func CreateDatasourcesLocalXml(configs []DatasourceXmlConfig, projectRootPath string) {
	var stringBuilder strings.Builder

	for _, c := range configs {
		stringBuilder.WriteString(fmt.Sprintf("%s\n", DatasourceLocalXmlDefinition(c)))
	}

	localXmlContent := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<project version="4">
  <component name="dataSourceStorageLocal" created-in="IU-233.13135.103">
    %s
  </component>
</project>`, stringBuilder.String())

	fname := "dataSources.local.xml"
	err := os.WriteFile(strings.Join([]string{projectRootPath, fname}, "/"), []byte(localXmlContent), 0644)
	if err != nil {
		panic(err)
	}
}

func CreateDatasourcesXml(configs []DatasourceXmlConfig, projectRootPath string) {
	var stringBuilder strings.Builder

	for _, c := range configs {
		stringBuilder.WriteString(fmt.Sprintf("%s\n", DatasourceXmlDefinition(c)))
	}

	xmlContent := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
	<project version="4">
		<component name="DataSourceManagerImpl" format="xml" multifile-model="true">
			%s
		</component>
	
	</project>`, stringBuilder.String())

	fname := "dataSources.xml"
	err := os.WriteFile(strings.Join([]string{projectRootPath, fname}, "/"), []byte(xmlContent), 0644)
	if err != nil {
		panic(err)
	}

}
