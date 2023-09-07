// Copyright 2023 - Hoby Smith - hoby@thoughtrealm.com. All rights reserved.
// Use of this source code is governed by an MIT license that can be found in
// the LICENSE file found in the project's main folder.

package helpers

import (
	"github.com/kirsle/configdir"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

const (
	TCGLobalFolderName = "timeconverter"
	TCConfigFileName   = "timeconverter.yaml"
)

func (hi *HelpersInfo) SaveLocalDefault() {
	configPath, err := os.Getwd()
	if err != nil {
		OP.Print(OutputMode_Force, "Unable to retrieve the current working directory. Error:", err)
		return
	}
	hi.writeConfigData(configPath, getConfigFileName())
}

func (hi *HelpersInfo) SaveGlobalDefault() {
	configPath := configdir.LocalConfig(TCGLobalFolderName)
	hi.writeConfigData(configPath, getConfigFileName())
}

func (hi *HelpersInfo) writeConfigData(configPath, configFile string) {
	configYaml := YamlConfig{ConfigInfo: CmdHelpers}

	configBytes, err := yaml.Marshal(configYaml)
	if err != nil {
		OP.Print(OutputMode_Force, "Unable to marshal config data. Error:", err)
		return
	}

	err = configdir.MakePath(configPath)
	if err != nil {
		OP.Print(OutputMode_Force, "Unable to create config path:", configFile)
		OP.Print(OutputMode_Force, "Error:", err)
	}

	configFilePath := filepath.Join(configPath, configFile)

	err = os.WriteFile(configFilePath, configBytes, 0666)
	if err != nil {
		OP.Print(OutputMode_Force, "Unable to write config data to:", configFilePath)
		OP.Print(OutputMode_Force, "Error:", err)
	}
}

func (hi *HelpersInfo) LoadConfigIfExists() {
	configFileName := getConfigFileName()

	globalConfigPath := configdir.LocalConfig(TCGLobalFolderName)
	globalConfigFilePath := filepath.Join(globalConfigPath, configFileName)

	localConfigPath, err := os.Getwd()
	if err != nil {
		OP.Print(OutputMode_Force, "Unable to retrieve the current working directory. Error:", err)
		return
	}
	localConfigFilePath := filepath.Join(localConfigPath, configFileName)

	var newHelperInfo *HelpersInfo
	var exists bool
	if exists, err = FileExists(localConfigFilePath); exists == true && err == nil {
		if ArgWasProvidedByUser([]string{"--set-global-default"}) && !CommandContains([]string{"clear", "show"}) {
			OP.Print(OutputMode_Force,
				"Warning: --set-global-default was provided when local defaults exist. "+
					"Local defaults will not be loaded. "+
					"Be sure to provide all required flags when setting global defaults and local defaults also exist.",
			)

			return
		}

		_, newHelperInfo, err = hi.readConfigFile(localConfigFilePath)
		if err != nil {
			OP.Print(OutputMode_Force, "Unable to read the local config file. Error:", err)
			return
		}
	} else if exists, err = FileExists(globalConfigFilePath); exists == true && err == nil {
		_, newHelperInfo, err = hi.readConfigFile(globalConfigFilePath)
		if err != nil {
			OP.Print(OutputMode_Force, "Unable to read the global config file. Error:", err)
			return
		}
	} else {
		return
	}

	hi.mergeHelperInfo(newHelperInfo)
}

func (hi *HelpersInfo) readConfigFile(configFilePath string) (configData []byte, newHelperInfo *HelpersInfo, err error) {
	if exists, err := FileExists(configFilePath); exists == false || err != nil {
		return nil, nil, err
	}

	fileBytes, err := os.ReadFile(configFilePath)
	if err != nil {
		OP.Print(OutputMode_Force, "Unable to read config file:", configFilePath)
		OP.Print(OutputMode_Force, "Error:", err)
		return nil, nil, err
	}

	var yamlConfig YamlConfig
	err = yaml.Unmarshal(fileBytes, &yamlConfig)
	if err != nil {
		OP.Print(OutputMode_Force, "Unable to parse config data:", err)
		return nil, nil, err
	}

	return fileBytes, yamlConfig.ConfigInfo, nil
}

func (hi *HelpersInfo) GetLocalConfigDataIfExists() (exists bool, contentBytes []byte, err error) {
	configFileName := getConfigFileName()

	localConfigPath, err := os.Getwd()
	if err != nil {
		OP.Print(OutputMode_Force, "Unable to retrieve the current working directory. Error:", err)
		return
	}
	localConfigFilePath := filepath.Join(localConfigPath, configFileName)

	exists, err = FileExists(localConfigFilePath)
	if exists == false || err != nil {
		return exists, nil, err
	}

	contentBytes, _, err = hi.readConfigFile(localConfigFilePath)
	return true, contentBytes, err
}

func (hi *HelpersInfo) GetGlobalConfigDataIfExists() (exists bool, contentBytes []byte, err error) {
	configFileName := getConfigFileName()

	globalConfigPath := configdir.LocalConfig(TCGLobalFolderName)
	globalConfigFilePath := filepath.Join(globalConfigPath, configFileName)

	exists, err = FileExists(globalConfigFilePath)
	if exists == false || err != nil {
		return exists, nil, err
	}

	contentBytes, _, err = hi.readConfigFile(globalConfigFilePath)
	return true, contentBytes, err
}

func (hi *HelpersInfo) mergeHelperInfo(newHelperInfo *HelpersInfo) {
	if !ArgWasProvidedByUser([]string{"--input-format", "-i"}) {
		CmdHelpers.InputFormatName = newHelperInfo.InputFormatName
	}

	if !ArgWasProvidedByUser([]string{"--input-layout", "l"}) {
		CmdHelpers.InputLayout = newHelperInfo.InputLayout
	}

	if !ArgWasProvidedByUser([]string{"--output-format", "-o"}) {
		CmdHelpers.OutputFormatName = newHelperInfo.OutputFormatName
	}

	if !ArgWasProvidedByUser([]string{"--output-layout", "-r"}) {
		CmdHelpers.OutputLayout = newHelperInfo.OutputLayout
	}

	if !ArgWasProvidedByUser([]string{"--output-target", "-t"}) {
		CmdHelpers.OutputTargetName = newHelperInfo.OutputTargetName
	}

	if !ArgWasProvidedByUser([]string{"--output-value-only", "-v"}) {
		CmdHelpers.OutputTargetName = newHelperInfo.OutputTargetName
	}

	if !ArgWasProvidedByUser([]string{"--output-timezone", "-z"}) {
		CmdHelpers.OutputTimeZone = newHelperInfo.OutputTimeZone
	}
}

func ArgWasProvidedByUser(argNames []string) bool {
	for _, arg := range os.Args {
		for _, argName := range argNames {
			if strings.ToLower(arg) == strings.ToLower(argName) {
				return true
			}
		}
	}

	return false
}

func CommandContains(names []string) bool {
	if len(os.Args) <= 1 {
		return false
	}

	cmdText := strings.ToLower(os.Args[1])
	for _, name := range names {
		if strings.Contains(cmdText, strings.ToLower(name)) {
			return true
		}
	}

	return false
}

func (hi *HelpersInfo) ClearLocalConfig() error {
	localConfigPath, err := os.Getwd()
	if err != nil {
		return err
	}
	localConfigFilePath := filepath.Join(localConfigPath, getConfigFileName())

	return os.Remove(localConfigFilePath)
}

func (hi *HelpersInfo) ClearGlobalConfig() error {
	globalConfigPath := configdir.LocalConfig(TCGLobalFolderName)
	globalConfigFilePath := filepath.Join(globalConfigPath, getConfigFileName())

	return os.Remove(globalConfigFilePath)
}
