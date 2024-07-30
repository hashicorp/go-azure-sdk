package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterStatusDetails string

const (
	PrinterStatusDetails_AlertRemovalOfBinaryChangeEntry           PrinterStatusDetails = "alertRemovalOfBinaryChangeEntry"
	PrinterStatusDetails_BanderAdded                               PrinterStatusDetails = "banderAdded"
	PrinterStatusDetails_BanderAlmostEmpty                         PrinterStatusDetails = "banderAlmostEmpty"
	PrinterStatusDetails_BanderAlmostFull                          PrinterStatusDetails = "banderAlmostFull"
	PrinterStatusDetails_BanderAtLimit                             PrinterStatusDetails = "banderAtLimit"
	PrinterStatusDetails_BanderClosed                              PrinterStatusDetails = "banderClosed"
	PrinterStatusDetails_BanderConfigurationChange                 PrinterStatusDetails = "banderConfigurationChange"
	PrinterStatusDetails_BanderCoverClosed                         PrinterStatusDetails = "banderCoverClosed"
	PrinterStatusDetails_BanderCoverOpen                           PrinterStatusDetails = "banderCoverOpen"
	PrinterStatusDetails_BanderEmpty                               PrinterStatusDetails = "banderEmpty"
	PrinterStatusDetails_BanderFull                                PrinterStatusDetails = "banderFull"
	PrinterStatusDetails_BanderInterlockClosed                     PrinterStatusDetails = "banderInterlockClosed"
	PrinterStatusDetails_BanderInterlockOpen                       PrinterStatusDetails = "banderInterlockOpen"
	PrinterStatusDetails_BanderJam                                 PrinterStatusDetails = "banderJam"
	PrinterStatusDetails_BanderLifeAlmostOver                      PrinterStatusDetails = "banderLifeAlmostOver"
	PrinterStatusDetails_BanderLifeOver                            PrinterStatusDetails = "banderLifeOver"
	PrinterStatusDetails_BanderMemoryExhausted                     PrinterStatusDetails = "banderMemoryExhausted"
	PrinterStatusDetails_BanderMissing                             PrinterStatusDetails = "banderMissing"
	PrinterStatusDetails_BanderMotorFailure                        PrinterStatusDetails = "banderMotorFailure"
	PrinterStatusDetails_BanderNearLimit                           PrinterStatusDetails = "banderNearLimit"
	PrinterStatusDetails_BanderOffline                             PrinterStatusDetails = "banderOffline"
	PrinterStatusDetails_BanderOpened                              PrinterStatusDetails = "banderOpened"
	PrinterStatusDetails_BanderOverTemperature                     PrinterStatusDetails = "banderOverTemperature"
	PrinterStatusDetails_BanderPowerSaver                          PrinterStatusDetails = "banderPowerSaver"
	PrinterStatusDetails_BanderRecoverableFailure                  PrinterStatusDetails = "banderRecoverableFailure"
	PrinterStatusDetails_BanderRecoverableStorage                  PrinterStatusDetails = "banderRecoverableStorage"
	PrinterStatusDetails_BanderRemoved                             PrinterStatusDetails = "banderRemoved"
	PrinterStatusDetails_BanderResourceAdded                       PrinterStatusDetails = "banderResourceAdded"
	PrinterStatusDetails_BanderResourceRemoved                     PrinterStatusDetails = "banderResourceRemoved"
	PrinterStatusDetails_BanderThermistorFailure                   PrinterStatusDetails = "banderThermistorFailure"
	PrinterStatusDetails_BanderTimingFailure                       PrinterStatusDetails = "banderTimingFailure"
	PrinterStatusDetails_BanderTurnedOff                           PrinterStatusDetails = "banderTurnedOff"
	PrinterStatusDetails_BanderTurnedOn                            PrinterStatusDetails = "banderTurnedOn"
	PrinterStatusDetails_BanderUnderTemperature                    PrinterStatusDetails = "banderUnderTemperature"
	PrinterStatusDetails_BanderUnrecoverableFailure                PrinterStatusDetails = "banderUnrecoverableFailure"
	PrinterStatusDetails_BanderUnrecoverableStorageError           PrinterStatusDetails = "banderUnrecoverableStorageError"
	PrinterStatusDetails_BanderWarmingUp                           PrinterStatusDetails = "banderWarmingUp"
	PrinterStatusDetails_BinderAdded                               PrinterStatusDetails = "binderAdded"
	PrinterStatusDetails_BinderAlmostEmpty                         PrinterStatusDetails = "binderAlmostEmpty"
	PrinterStatusDetails_BinderAlmostFull                          PrinterStatusDetails = "binderAlmostFull"
	PrinterStatusDetails_BinderAtLimit                             PrinterStatusDetails = "binderAtLimit"
	PrinterStatusDetails_BinderClosed                              PrinterStatusDetails = "binderClosed"
	PrinterStatusDetails_BinderConfigurationChange                 PrinterStatusDetails = "binderConfigurationChange"
	PrinterStatusDetails_BinderCoverClosed                         PrinterStatusDetails = "binderCoverClosed"
	PrinterStatusDetails_BinderCoverOpen                           PrinterStatusDetails = "binderCoverOpen"
	PrinterStatusDetails_BinderEmpty                               PrinterStatusDetails = "binderEmpty"
	PrinterStatusDetails_BinderFull                                PrinterStatusDetails = "binderFull"
	PrinterStatusDetails_BinderInterlockClosed                     PrinterStatusDetails = "binderInterlockClosed"
	PrinterStatusDetails_BinderInterlockOpen                       PrinterStatusDetails = "binderInterlockOpen"
	PrinterStatusDetails_BinderJam                                 PrinterStatusDetails = "binderJam"
	PrinterStatusDetails_BinderLifeAlmostOver                      PrinterStatusDetails = "binderLifeAlmostOver"
	PrinterStatusDetails_BinderLifeOver                            PrinterStatusDetails = "binderLifeOver"
	PrinterStatusDetails_BinderMemoryExhausted                     PrinterStatusDetails = "binderMemoryExhausted"
	PrinterStatusDetails_BinderMissing                             PrinterStatusDetails = "binderMissing"
	PrinterStatusDetails_BinderMotorFailure                        PrinterStatusDetails = "binderMotorFailure"
	PrinterStatusDetails_BinderNearLimit                           PrinterStatusDetails = "binderNearLimit"
	PrinterStatusDetails_BinderOffline                             PrinterStatusDetails = "binderOffline"
	PrinterStatusDetails_BinderOpened                              PrinterStatusDetails = "binderOpened"
	PrinterStatusDetails_BinderOverTemperature                     PrinterStatusDetails = "binderOverTemperature"
	PrinterStatusDetails_BinderPowerSaver                          PrinterStatusDetails = "binderPowerSaver"
	PrinterStatusDetails_BinderRecoverableFailure                  PrinterStatusDetails = "binderRecoverableFailure"
	PrinterStatusDetails_BinderRecoverableStorage                  PrinterStatusDetails = "binderRecoverableStorage"
	PrinterStatusDetails_BinderRemoved                             PrinterStatusDetails = "binderRemoved"
	PrinterStatusDetails_BinderResourceAdded                       PrinterStatusDetails = "binderResourceAdded"
	PrinterStatusDetails_BinderResourceRemoved                     PrinterStatusDetails = "binderResourceRemoved"
	PrinterStatusDetails_BinderThermistorFailure                   PrinterStatusDetails = "binderThermistorFailure"
	PrinterStatusDetails_BinderTimingFailure                       PrinterStatusDetails = "binderTimingFailure"
	PrinterStatusDetails_BinderTurnedOff                           PrinterStatusDetails = "binderTurnedOff"
	PrinterStatusDetails_BinderTurnedOn                            PrinterStatusDetails = "binderTurnedOn"
	PrinterStatusDetails_BinderUnderTemperature                    PrinterStatusDetails = "binderUnderTemperature"
	PrinterStatusDetails_BinderUnrecoverableFailure                PrinterStatusDetails = "binderUnrecoverableFailure"
	PrinterStatusDetails_BinderUnrecoverableStorageError           PrinterStatusDetails = "binderUnrecoverableStorageError"
	PrinterStatusDetails_BinderWarmingUp                           PrinterStatusDetails = "binderWarmingUp"
	PrinterStatusDetails_CameraFailure                             PrinterStatusDetails = "cameraFailure"
	PrinterStatusDetails_ChamberCooling                            PrinterStatusDetails = "chamberCooling"
	PrinterStatusDetails_ChamberFailure                            PrinterStatusDetails = "chamberFailure"
	PrinterStatusDetails_ChamberHeating                            PrinterStatusDetails = "chamberHeating"
	PrinterStatusDetails_ChamberTemperatureHigh                    PrinterStatusDetails = "chamberTemperatureHigh"
	PrinterStatusDetails_ChamberTemperatureLow                     PrinterStatusDetails = "chamberTemperatureLow"
	PrinterStatusDetails_CleanerLifeAlmostOver                     PrinterStatusDetails = "cleanerLifeAlmostOver"
	PrinterStatusDetails_CleanerLifeOver                           PrinterStatusDetails = "cleanerLifeOver"
	PrinterStatusDetails_ConfigurationChange                       PrinterStatusDetails = "configurationChange"
	PrinterStatusDetails_ConnectingToDevice                        PrinterStatusDetails = "connectingToDevice"
	PrinterStatusDetails_CoverOpen                                 PrinterStatusDetails = "coverOpen"
	PrinterStatusDetails_Deactivated                               PrinterStatusDetails = "deactivated"
	PrinterStatusDetails_Deleted                                   PrinterStatusDetails = "deleted"
	PrinterStatusDetails_DeveloperEmpty                            PrinterStatusDetails = "developerEmpty"
	PrinterStatusDetails_DeveloperLow                              PrinterStatusDetails = "developerLow"
	PrinterStatusDetails_DieCutterAdded                            PrinterStatusDetails = "dieCutterAdded"
	PrinterStatusDetails_DieCutterAlmostEmpty                      PrinterStatusDetails = "dieCutterAlmostEmpty"
	PrinterStatusDetails_DieCutterAlmostFull                       PrinterStatusDetails = "dieCutterAlmostFull"
	PrinterStatusDetails_DieCutterAtLimit                          PrinterStatusDetails = "dieCutterAtLimit"
	PrinterStatusDetails_DieCutterClosed                           PrinterStatusDetails = "dieCutterClosed"
	PrinterStatusDetails_DieCutterConfigurationChange              PrinterStatusDetails = "dieCutterConfigurationChange"
	PrinterStatusDetails_DieCutterCoverClosed                      PrinterStatusDetails = "dieCutterCoverClosed"
	PrinterStatusDetails_DieCutterCoverOpen                        PrinterStatusDetails = "dieCutterCoverOpen"
	PrinterStatusDetails_DieCutterEmpty                            PrinterStatusDetails = "dieCutterEmpty"
	PrinterStatusDetails_DieCutterFull                             PrinterStatusDetails = "dieCutterFull"
	PrinterStatusDetails_DieCutterInterlockClosed                  PrinterStatusDetails = "dieCutterInterlockClosed"
	PrinterStatusDetails_DieCutterInterlockOpen                    PrinterStatusDetails = "dieCutterInterlockOpen"
	PrinterStatusDetails_DieCutterJam                              PrinterStatusDetails = "dieCutterJam"
	PrinterStatusDetails_DieCutterLifeAlmostOver                   PrinterStatusDetails = "dieCutterLifeAlmostOver"
	PrinterStatusDetails_DieCutterLifeOver                         PrinterStatusDetails = "dieCutterLifeOver"
	PrinterStatusDetails_DieCutterMemoryExhausted                  PrinterStatusDetails = "dieCutterMemoryExhausted"
	PrinterStatusDetails_DieCutterMissing                          PrinterStatusDetails = "dieCutterMissing"
	PrinterStatusDetails_DieCutterMotorFailure                     PrinterStatusDetails = "dieCutterMotorFailure"
	PrinterStatusDetails_DieCutterNearLimit                        PrinterStatusDetails = "dieCutterNearLimit"
	PrinterStatusDetails_DieCutterOffline                          PrinterStatusDetails = "dieCutterOffline"
	PrinterStatusDetails_DieCutterOpened                           PrinterStatusDetails = "dieCutterOpened"
	PrinterStatusDetails_DieCutterOverTemperature                  PrinterStatusDetails = "dieCutterOverTemperature"
	PrinterStatusDetails_DieCutterPowerSaver                       PrinterStatusDetails = "dieCutterPowerSaver"
	PrinterStatusDetails_DieCutterRecoverableFailure               PrinterStatusDetails = "dieCutterRecoverableFailure"
	PrinterStatusDetails_DieCutterRecoverableStorage               PrinterStatusDetails = "dieCutterRecoverableStorage"
	PrinterStatusDetails_DieCutterRemoved                          PrinterStatusDetails = "dieCutterRemoved"
	PrinterStatusDetails_DieCutterResourceAdded                    PrinterStatusDetails = "dieCutterResourceAdded"
	PrinterStatusDetails_DieCutterResourceRemoved                  PrinterStatusDetails = "dieCutterResourceRemoved"
	PrinterStatusDetails_DieCutterThermistorFailure                PrinterStatusDetails = "dieCutterThermistorFailure"
	PrinterStatusDetails_DieCutterTimingFailure                    PrinterStatusDetails = "dieCutterTimingFailure"
	PrinterStatusDetails_DieCutterTurnedOff                        PrinterStatusDetails = "dieCutterTurnedOff"
	PrinterStatusDetails_DieCutterTurnedOn                         PrinterStatusDetails = "dieCutterTurnedOn"
	PrinterStatusDetails_DieCutterUnderTemperature                 PrinterStatusDetails = "dieCutterUnderTemperature"
	PrinterStatusDetails_DieCutterUnrecoverableFailure             PrinterStatusDetails = "dieCutterUnrecoverableFailure"
	PrinterStatusDetails_DieCutterUnrecoverableStorageError        PrinterStatusDetails = "dieCutterUnrecoverableStorageError"
	PrinterStatusDetails_DieCutterWarmingUp                        PrinterStatusDetails = "dieCutterWarmingUp"
	PrinterStatusDetails_DoorOpen                                  PrinterStatusDetails = "doorOpen"
	PrinterStatusDetails_ExtruderCooling                           PrinterStatusDetails = "extruderCooling"
	PrinterStatusDetails_ExtruderFailure                           PrinterStatusDetails = "extruderFailure"
	PrinterStatusDetails_ExtruderHeating                           PrinterStatusDetails = "extruderHeating"
	PrinterStatusDetails_ExtruderJam                               PrinterStatusDetails = "extruderJam"
	PrinterStatusDetails_ExtruderTemperatureHigh                   PrinterStatusDetails = "extruderTemperatureHigh"
	PrinterStatusDetails_ExtruderTemperatureLow                    PrinterStatusDetails = "extruderTemperatureLow"
	PrinterStatusDetails_FanFailure                                PrinterStatusDetails = "fanFailure"
	PrinterStatusDetails_FaxModemLifeAlmostOver                    PrinterStatusDetails = "faxModemLifeAlmostOver"
	PrinterStatusDetails_FaxModemLifeOver                          PrinterStatusDetails = "faxModemLifeOver"
	PrinterStatusDetails_FaxModemMissing                           PrinterStatusDetails = "faxModemMissing"
	PrinterStatusDetails_FaxModemTurnedOff                         PrinterStatusDetails = "faxModemTurnedOff"
	PrinterStatusDetails_FaxModemTurnedOn                          PrinterStatusDetails = "faxModemTurnedOn"
	PrinterStatusDetails_FolderAdded                               PrinterStatusDetails = "folderAdded"
	PrinterStatusDetails_FolderAlmostEmpty                         PrinterStatusDetails = "folderAlmostEmpty"
	PrinterStatusDetails_FolderAlmostFull                          PrinterStatusDetails = "folderAlmostFull"
	PrinterStatusDetails_FolderAtLimit                             PrinterStatusDetails = "folderAtLimit"
	PrinterStatusDetails_FolderClosed                              PrinterStatusDetails = "folderClosed"
	PrinterStatusDetails_FolderConfigurationChange                 PrinterStatusDetails = "folderConfigurationChange"
	PrinterStatusDetails_FolderCoverClosed                         PrinterStatusDetails = "folderCoverClosed"
	PrinterStatusDetails_FolderCoverOpen                           PrinterStatusDetails = "folderCoverOpen"
	PrinterStatusDetails_FolderEmpty                               PrinterStatusDetails = "folderEmpty"
	PrinterStatusDetails_FolderFull                                PrinterStatusDetails = "folderFull"
	PrinterStatusDetails_FolderInterlockClosed                     PrinterStatusDetails = "folderInterlockClosed"
	PrinterStatusDetails_FolderInterlockOpen                       PrinterStatusDetails = "folderInterlockOpen"
	PrinterStatusDetails_FolderJam                                 PrinterStatusDetails = "folderJam"
	PrinterStatusDetails_FolderLifeAlmostOver                      PrinterStatusDetails = "folderLifeAlmostOver"
	PrinterStatusDetails_FolderLifeOver                            PrinterStatusDetails = "folderLifeOver"
	PrinterStatusDetails_FolderMemoryExhausted                     PrinterStatusDetails = "folderMemoryExhausted"
	PrinterStatusDetails_FolderMissing                             PrinterStatusDetails = "folderMissing"
	PrinterStatusDetails_FolderMotorFailure                        PrinterStatusDetails = "folderMotorFailure"
	PrinterStatusDetails_FolderNearLimit                           PrinterStatusDetails = "folderNearLimit"
	PrinterStatusDetails_FolderOffline                             PrinterStatusDetails = "folderOffline"
	PrinterStatusDetails_FolderOpened                              PrinterStatusDetails = "folderOpened"
	PrinterStatusDetails_FolderOverTemperature                     PrinterStatusDetails = "folderOverTemperature"
	PrinterStatusDetails_FolderPowerSaver                          PrinterStatusDetails = "folderPowerSaver"
	PrinterStatusDetails_FolderRecoverableFailure                  PrinterStatusDetails = "folderRecoverableFailure"
	PrinterStatusDetails_FolderRecoverableStorage                  PrinterStatusDetails = "folderRecoverableStorage"
	PrinterStatusDetails_FolderRemoved                             PrinterStatusDetails = "folderRemoved"
	PrinterStatusDetails_FolderResourceAdded                       PrinterStatusDetails = "folderResourceAdded"
	PrinterStatusDetails_FolderResourceRemoved                     PrinterStatusDetails = "folderResourceRemoved"
	PrinterStatusDetails_FolderThermistorFailure                   PrinterStatusDetails = "folderThermistorFailure"
	PrinterStatusDetails_FolderTimingFailure                       PrinterStatusDetails = "folderTimingFailure"
	PrinterStatusDetails_FolderTurnedOff                           PrinterStatusDetails = "folderTurnedOff"
	PrinterStatusDetails_FolderTurnedOn                            PrinterStatusDetails = "folderTurnedOn"
	PrinterStatusDetails_FolderUnderTemperature                    PrinterStatusDetails = "folderUnderTemperature"
	PrinterStatusDetails_FolderUnrecoverableFailure                PrinterStatusDetails = "folderUnrecoverableFailure"
	PrinterStatusDetails_FolderUnrecoverableStorageError           PrinterStatusDetails = "folderUnrecoverableStorageError"
	PrinterStatusDetails_FolderWarmingUp                           PrinterStatusDetails = "folderWarmingUp"
	PrinterStatusDetails_FuserOverTemp                             PrinterStatusDetails = "fuserOverTemp"
	PrinterStatusDetails_FuserUnderTemp                            PrinterStatusDetails = "fuserUnderTemp"
	PrinterStatusDetails_Hibernate                                 PrinterStatusDetails = "hibernate"
	PrinterStatusDetails_HoldNewJobs                               PrinterStatusDetails = "holdNewJobs"
	PrinterStatusDetails_IdentifyPrinterRequested                  PrinterStatusDetails = "identifyPrinterRequested"
	PrinterStatusDetails_ImprinterAdded                            PrinterStatusDetails = "imprinterAdded"
	PrinterStatusDetails_ImprinterAlmostEmpty                      PrinterStatusDetails = "imprinterAlmostEmpty"
	PrinterStatusDetails_ImprinterAlmostFull                       PrinterStatusDetails = "imprinterAlmostFull"
	PrinterStatusDetails_ImprinterAtLimit                          PrinterStatusDetails = "imprinterAtLimit"
	PrinterStatusDetails_ImprinterClosed                           PrinterStatusDetails = "imprinterClosed"
	PrinterStatusDetails_ImprinterConfigurationChange              PrinterStatusDetails = "imprinterConfigurationChange"
	PrinterStatusDetails_ImprinterCoverClosed                      PrinterStatusDetails = "imprinterCoverClosed"
	PrinterStatusDetails_ImprinterCoverOpen                        PrinterStatusDetails = "imprinterCoverOpen"
	PrinterStatusDetails_ImprinterEmpty                            PrinterStatusDetails = "imprinterEmpty"
	PrinterStatusDetails_ImprinterFull                             PrinterStatusDetails = "imprinterFull"
	PrinterStatusDetails_ImprinterInterlockClosed                  PrinterStatusDetails = "imprinterInterlockClosed"
	PrinterStatusDetails_ImprinterInterlockOpen                    PrinterStatusDetails = "imprinterInterlockOpen"
	PrinterStatusDetails_ImprinterJam                              PrinterStatusDetails = "imprinterJam"
	PrinterStatusDetails_ImprinterLifeAlmostOver                   PrinterStatusDetails = "imprinterLifeAlmostOver"
	PrinterStatusDetails_ImprinterLifeOver                         PrinterStatusDetails = "imprinterLifeOver"
	PrinterStatusDetails_ImprinterMemoryExhausted                  PrinterStatusDetails = "imprinterMemoryExhausted"
	PrinterStatusDetails_ImprinterMissing                          PrinterStatusDetails = "imprinterMissing"
	PrinterStatusDetails_ImprinterMotorFailure                     PrinterStatusDetails = "imprinterMotorFailure"
	PrinterStatusDetails_ImprinterNearLimit                        PrinterStatusDetails = "imprinterNearLimit"
	PrinterStatusDetails_ImprinterOffline                          PrinterStatusDetails = "imprinterOffline"
	PrinterStatusDetails_ImprinterOpened                           PrinterStatusDetails = "imprinterOpened"
	PrinterStatusDetails_ImprinterOverTemperature                  PrinterStatusDetails = "imprinterOverTemperature"
	PrinterStatusDetails_ImprinterPowerSaver                       PrinterStatusDetails = "imprinterPowerSaver"
	PrinterStatusDetails_ImprinterRecoverableFailure               PrinterStatusDetails = "imprinterRecoverableFailure"
	PrinterStatusDetails_ImprinterRecoverableStorage               PrinterStatusDetails = "imprinterRecoverableStorage"
	PrinterStatusDetails_ImprinterRemoved                          PrinterStatusDetails = "imprinterRemoved"
	PrinterStatusDetails_ImprinterResourceAdded                    PrinterStatusDetails = "imprinterResourceAdded"
	PrinterStatusDetails_ImprinterResourceRemoved                  PrinterStatusDetails = "imprinterResourceRemoved"
	PrinterStatusDetails_ImprinterThermistorFailure                PrinterStatusDetails = "imprinterThermistorFailure"
	PrinterStatusDetails_ImprinterTimingFailure                    PrinterStatusDetails = "imprinterTimingFailure"
	PrinterStatusDetails_ImprinterTurnedOff                        PrinterStatusDetails = "imprinterTurnedOff"
	PrinterStatusDetails_ImprinterTurnedOn                         PrinterStatusDetails = "imprinterTurnedOn"
	PrinterStatusDetails_ImprinterUnderTemperature                 PrinterStatusDetails = "imprinterUnderTemperature"
	PrinterStatusDetails_ImprinterUnrecoverableFailure             PrinterStatusDetails = "imprinterUnrecoverableFailure"
	PrinterStatusDetails_ImprinterUnrecoverableStorageError        PrinterStatusDetails = "imprinterUnrecoverableStorageError"
	PrinterStatusDetails_ImprinterWarmingUp                        PrinterStatusDetails = "imprinterWarmingUp"
	PrinterStatusDetails_InputCannotFeedSizeSelected               PrinterStatusDetails = "inputCannotFeedSizeSelected"
	PrinterStatusDetails_InputManualInputRequest                   PrinterStatusDetails = "inputManualInputRequest"
	PrinterStatusDetails_InputMediaColorChange                     PrinterStatusDetails = "inputMediaColorChange"
	PrinterStatusDetails_InputMediaFormPartsChange                 PrinterStatusDetails = "inputMediaFormPartsChange"
	PrinterStatusDetails_InputMediaSizeChange                      PrinterStatusDetails = "inputMediaSizeChange"
	PrinterStatusDetails_InputMediaTrayFailure                     PrinterStatusDetails = "inputMediaTrayFailure"
	PrinterStatusDetails_InputMediaTrayFeedError                   PrinterStatusDetails = "inputMediaTrayFeedError"
	PrinterStatusDetails_InputMediaTrayJam                         PrinterStatusDetails = "inputMediaTrayJam"
	PrinterStatusDetails_InputMediaTypeChange                      PrinterStatusDetails = "inputMediaTypeChange"
	PrinterStatusDetails_InputMediaWeightChange                    PrinterStatusDetails = "inputMediaWeightChange"
	PrinterStatusDetails_InputPickRollerFailure                    PrinterStatusDetails = "inputPickRollerFailure"
	PrinterStatusDetails_InputPickRollerLifeOver                   PrinterStatusDetails = "inputPickRollerLifeOver"
	PrinterStatusDetails_InputPickRollerLifeWarn                   PrinterStatusDetails = "inputPickRollerLifeWarn"
	PrinterStatusDetails_InputPickRollerMissing                    PrinterStatusDetails = "inputPickRollerMissing"
	PrinterStatusDetails_InputTrayElevationFailure                 PrinterStatusDetails = "inputTrayElevationFailure"
	PrinterStatusDetails_InputTrayMissing                          PrinterStatusDetails = "inputTrayMissing"
	PrinterStatusDetails_InputTrayPositionFailure                  PrinterStatusDetails = "inputTrayPositionFailure"
	PrinterStatusDetails_InserterAdded                             PrinterStatusDetails = "inserterAdded"
	PrinterStatusDetails_InserterAlmostEmpty                       PrinterStatusDetails = "inserterAlmostEmpty"
	PrinterStatusDetails_InserterAlmostFull                        PrinterStatusDetails = "inserterAlmostFull"
	PrinterStatusDetails_InserterAtLimit                           PrinterStatusDetails = "inserterAtLimit"
	PrinterStatusDetails_InserterClosed                            PrinterStatusDetails = "inserterClosed"
	PrinterStatusDetails_InserterConfigurationChange               PrinterStatusDetails = "inserterConfigurationChange"
	PrinterStatusDetails_InserterCoverClosed                       PrinterStatusDetails = "inserterCoverClosed"
	PrinterStatusDetails_InserterCoverOpen                         PrinterStatusDetails = "inserterCoverOpen"
	PrinterStatusDetails_InserterEmpty                             PrinterStatusDetails = "inserterEmpty"
	PrinterStatusDetails_InserterFull                              PrinterStatusDetails = "inserterFull"
	PrinterStatusDetails_InserterInterlockClosed                   PrinterStatusDetails = "inserterInterlockClosed"
	PrinterStatusDetails_InserterInterlockOpen                     PrinterStatusDetails = "inserterInterlockOpen"
	PrinterStatusDetails_InserterJam                               PrinterStatusDetails = "inserterJam"
	PrinterStatusDetails_InserterLifeAlmostOver                    PrinterStatusDetails = "inserterLifeAlmostOver"
	PrinterStatusDetails_InserterLifeOver                          PrinterStatusDetails = "inserterLifeOver"
	PrinterStatusDetails_InserterMemoryExhausted                   PrinterStatusDetails = "inserterMemoryExhausted"
	PrinterStatusDetails_InserterMissing                           PrinterStatusDetails = "inserterMissing"
	PrinterStatusDetails_InserterMotorFailure                      PrinterStatusDetails = "inserterMotorFailure"
	PrinterStatusDetails_InserterNearLimit                         PrinterStatusDetails = "inserterNearLimit"
	PrinterStatusDetails_InserterOffline                           PrinterStatusDetails = "inserterOffline"
	PrinterStatusDetails_InserterOpened                            PrinterStatusDetails = "inserterOpened"
	PrinterStatusDetails_InserterOverTemperature                   PrinterStatusDetails = "inserterOverTemperature"
	PrinterStatusDetails_InserterPowerSaver                        PrinterStatusDetails = "inserterPowerSaver"
	PrinterStatusDetails_InserterRecoverableFailure                PrinterStatusDetails = "inserterRecoverableFailure"
	PrinterStatusDetails_InserterRecoverableStorage                PrinterStatusDetails = "inserterRecoverableStorage"
	PrinterStatusDetails_InserterRemoved                           PrinterStatusDetails = "inserterRemoved"
	PrinterStatusDetails_InserterResourceAdded                     PrinterStatusDetails = "inserterResourceAdded"
	PrinterStatusDetails_InserterResourceRemoved                   PrinterStatusDetails = "inserterResourceRemoved"
	PrinterStatusDetails_InserterThermistorFailure                 PrinterStatusDetails = "inserterThermistorFailure"
	PrinterStatusDetails_InserterTimingFailure                     PrinterStatusDetails = "inserterTimingFailure"
	PrinterStatusDetails_InserterTurnedOff                         PrinterStatusDetails = "inserterTurnedOff"
	PrinterStatusDetails_InserterTurnedOn                          PrinterStatusDetails = "inserterTurnedOn"
	PrinterStatusDetails_InserterUnderTemperature                  PrinterStatusDetails = "inserterUnderTemperature"
	PrinterStatusDetails_InserterUnrecoverableFailure              PrinterStatusDetails = "inserterUnrecoverableFailure"
	PrinterStatusDetails_InserterUnrecoverableStorageError         PrinterStatusDetails = "inserterUnrecoverableStorageError"
	PrinterStatusDetails_InserterWarmingUp                         PrinterStatusDetails = "inserterWarmingUp"
	PrinterStatusDetails_InterlockClosed                           PrinterStatusDetails = "interlockClosed"
	PrinterStatusDetails_InterlockOpen                             PrinterStatusDetails = "interlockOpen"
	PrinterStatusDetails_InterpreterCartridgeAdded                 PrinterStatusDetails = "interpreterCartridgeAdded"
	PrinterStatusDetails_InterpreterCartridgeDeleted               PrinterStatusDetails = "interpreterCartridgeDeleted"
	PrinterStatusDetails_InterpreterComplexPageEncountered         PrinterStatusDetails = "interpreterComplexPageEncountered"
	PrinterStatusDetails_InterpreterMemoryDecrease                 PrinterStatusDetails = "interpreterMemoryDecrease"
	PrinterStatusDetails_InterpreterMemoryIncrease                 PrinterStatusDetails = "interpreterMemoryIncrease"
	PrinterStatusDetails_InterpreterResourceAdded                  PrinterStatusDetails = "interpreterResourceAdded"
	PrinterStatusDetails_InterpreterResourceDeleted                PrinterStatusDetails = "interpreterResourceDeleted"
	PrinterStatusDetails_InterpreterResourceUnavailable            PrinterStatusDetails = "interpreterResourceUnavailable"
	PrinterStatusDetails_LampAtEol                                 PrinterStatusDetails = "lampAtEol"
	PrinterStatusDetails_LampFailure                               PrinterStatusDetails = "lampFailure"
	PrinterStatusDetails_LampNearEol                               PrinterStatusDetails = "lampNearEol"
	PrinterStatusDetails_LaserAtEol                                PrinterStatusDetails = "laserAtEol"
	PrinterStatusDetails_LaserFailure                              PrinterStatusDetails = "laserFailure"
	PrinterStatusDetails_LaserNearEol                              PrinterStatusDetails = "laserNearEol"
	PrinterStatusDetails_MakeEnvelopeAdded                         PrinterStatusDetails = "makeEnvelopeAdded"
	PrinterStatusDetails_MakeEnvelopeAlmostEmpty                   PrinterStatusDetails = "makeEnvelopeAlmostEmpty"
	PrinterStatusDetails_MakeEnvelopeAlmostFull                    PrinterStatusDetails = "makeEnvelopeAlmostFull"
	PrinterStatusDetails_MakeEnvelopeAtLimit                       PrinterStatusDetails = "makeEnvelopeAtLimit"
	PrinterStatusDetails_MakeEnvelopeClosed                        PrinterStatusDetails = "makeEnvelopeClosed"
	PrinterStatusDetails_MakeEnvelopeConfigurationChange           PrinterStatusDetails = "makeEnvelopeConfigurationChange"
	PrinterStatusDetails_MakeEnvelopeCoverClosed                   PrinterStatusDetails = "makeEnvelopeCoverClosed"
	PrinterStatusDetails_MakeEnvelopeCoverOpen                     PrinterStatusDetails = "makeEnvelopeCoverOpen"
	PrinterStatusDetails_MakeEnvelopeEmpty                         PrinterStatusDetails = "makeEnvelopeEmpty"
	PrinterStatusDetails_MakeEnvelopeFull                          PrinterStatusDetails = "makeEnvelopeFull"
	PrinterStatusDetails_MakeEnvelopeInterlockClosed               PrinterStatusDetails = "makeEnvelopeInterlockClosed"
	PrinterStatusDetails_MakeEnvelopeInterlockOpen                 PrinterStatusDetails = "makeEnvelopeInterlockOpen"
	PrinterStatusDetails_MakeEnvelopeJam                           PrinterStatusDetails = "makeEnvelopeJam"
	PrinterStatusDetails_MakeEnvelopeLifeAlmostOver                PrinterStatusDetails = "makeEnvelopeLifeAlmostOver"
	PrinterStatusDetails_MakeEnvelopeLifeOver                      PrinterStatusDetails = "makeEnvelopeLifeOver"
	PrinterStatusDetails_MakeEnvelopeMemoryExhausted               PrinterStatusDetails = "makeEnvelopeMemoryExhausted"
	PrinterStatusDetails_MakeEnvelopeMissing                       PrinterStatusDetails = "makeEnvelopeMissing"
	PrinterStatusDetails_MakeEnvelopeMotorFailure                  PrinterStatusDetails = "makeEnvelopeMotorFailure"
	PrinterStatusDetails_MakeEnvelopeNearLimit                     PrinterStatusDetails = "makeEnvelopeNearLimit"
	PrinterStatusDetails_MakeEnvelopeOffline                       PrinterStatusDetails = "makeEnvelopeOffline"
	PrinterStatusDetails_MakeEnvelopeOpened                        PrinterStatusDetails = "makeEnvelopeOpened"
	PrinterStatusDetails_MakeEnvelopeOverTemperature               PrinterStatusDetails = "makeEnvelopeOverTemperature"
	PrinterStatusDetails_MakeEnvelopePowerSaver                    PrinterStatusDetails = "makeEnvelopePowerSaver"
	PrinterStatusDetails_MakeEnvelopeRecoverableFailure            PrinterStatusDetails = "makeEnvelopeRecoverableFailure"
	PrinterStatusDetails_MakeEnvelopeRecoverableStorage            PrinterStatusDetails = "makeEnvelopeRecoverableStorage"
	PrinterStatusDetails_MakeEnvelopeRemoved                       PrinterStatusDetails = "makeEnvelopeRemoved"
	PrinterStatusDetails_MakeEnvelopeResourceAdded                 PrinterStatusDetails = "makeEnvelopeResourceAdded"
	PrinterStatusDetails_MakeEnvelopeResourceRemoved               PrinterStatusDetails = "makeEnvelopeResourceRemoved"
	PrinterStatusDetails_MakeEnvelopeThermistorFailure             PrinterStatusDetails = "makeEnvelopeThermistorFailure"
	PrinterStatusDetails_MakeEnvelopeTimingFailure                 PrinterStatusDetails = "makeEnvelopeTimingFailure"
	PrinterStatusDetails_MakeEnvelopeTurnedOff                     PrinterStatusDetails = "makeEnvelopeTurnedOff"
	PrinterStatusDetails_MakeEnvelopeTurnedOn                      PrinterStatusDetails = "makeEnvelopeTurnedOn"
	PrinterStatusDetails_MakeEnvelopeUnderTemperature              PrinterStatusDetails = "makeEnvelopeUnderTemperature"
	PrinterStatusDetails_MakeEnvelopeUnrecoverableFailure          PrinterStatusDetails = "makeEnvelopeUnrecoverableFailure"
	PrinterStatusDetails_MakeEnvelopeUnrecoverableStorageError     PrinterStatusDetails = "makeEnvelopeUnrecoverableStorageError"
	PrinterStatusDetails_MakeEnvelopeWarmingUp                     PrinterStatusDetails = "makeEnvelopeWarmingUp"
	PrinterStatusDetails_MarkerAdjustingPrintQuality               PrinterStatusDetails = "markerAdjustingPrintQuality"
	PrinterStatusDetails_MarkerCleanerMissing                      PrinterStatusDetails = "markerCleanerMissing"
	PrinterStatusDetails_MarkerDeveloperAlmostEmpty                PrinterStatusDetails = "markerDeveloperAlmostEmpty"
	PrinterStatusDetails_MarkerDeveloperEmpty                      PrinterStatusDetails = "markerDeveloperEmpty"
	PrinterStatusDetails_MarkerDeveloperMissing                    PrinterStatusDetails = "markerDeveloperMissing"
	PrinterStatusDetails_MarkerFuserMissing                        PrinterStatusDetails = "markerFuserMissing"
	PrinterStatusDetails_MarkerFuserThermistorFailure              PrinterStatusDetails = "markerFuserThermistorFailure"
	PrinterStatusDetails_MarkerFuserTimingFailure                  PrinterStatusDetails = "markerFuserTimingFailure"
	PrinterStatusDetails_MarkerInkAlmostEmpty                      PrinterStatusDetails = "markerInkAlmostEmpty"
	PrinterStatusDetails_MarkerInkEmpty                            PrinterStatusDetails = "markerInkEmpty"
	PrinterStatusDetails_MarkerInkMissing                          PrinterStatusDetails = "markerInkMissing"
	PrinterStatusDetails_MarkerOpcMissing                          PrinterStatusDetails = "markerOpcMissing"
	PrinterStatusDetails_MarkerPrintRibbonAlmostEmpty              PrinterStatusDetails = "markerPrintRibbonAlmostEmpty"
	PrinterStatusDetails_MarkerPrintRibbonEmpty                    PrinterStatusDetails = "markerPrintRibbonEmpty"
	PrinterStatusDetails_MarkerPrintRibbonMissing                  PrinterStatusDetails = "markerPrintRibbonMissing"
	PrinterStatusDetails_MarkerSupplyAlmostEmpty                   PrinterStatusDetails = "markerSupplyAlmostEmpty"
	PrinterStatusDetails_MarkerSupplyEmpty                         PrinterStatusDetails = "markerSupplyEmpty"
	PrinterStatusDetails_MarkerSupplyLow                           PrinterStatusDetails = "markerSupplyLow"
	PrinterStatusDetails_MarkerSupplyMissing                       PrinterStatusDetails = "markerSupplyMissing"
	PrinterStatusDetails_MarkerTonerCartridgeMissing               PrinterStatusDetails = "markerTonerCartridgeMissing"
	PrinterStatusDetails_MarkerTonerMissing                        PrinterStatusDetails = "markerTonerMissing"
	PrinterStatusDetails_MarkerWasteAlmostFull                     PrinterStatusDetails = "markerWasteAlmostFull"
	PrinterStatusDetails_MarkerWasteFull                           PrinterStatusDetails = "markerWasteFull"
	PrinterStatusDetails_MarkerWasteInkReceptacleAlmostFull        PrinterStatusDetails = "markerWasteInkReceptacleAlmostFull"
	PrinterStatusDetails_MarkerWasteInkReceptacleFull              PrinterStatusDetails = "markerWasteInkReceptacleFull"
	PrinterStatusDetails_MarkerWasteInkReceptacleMissing           PrinterStatusDetails = "markerWasteInkReceptacleMissing"
	PrinterStatusDetails_MarkerWasteMissing                        PrinterStatusDetails = "markerWasteMissing"
	PrinterStatusDetails_MarkerWasteTonerReceptacleAlmostFull      PrinterStatusDetails = "markerWasteTonerReceptacleAlmostFull"
	PrinterStatusDetails_MarkerWasteTonerReceptacleFull            PrinterStatusDetails = "markerWasteTonerReceptacleFull"
	PrinterStatusDetails_MarkerWasteTonerReceptacleMissing         PrinterStatusDetails = "markerWasteTonerReceptacleMissing"
	PrinterStatusDetails_MaterialEmpty                             PrinterStatusDetails = "materialEmpty"
	PrinterStatusDetails_MaterialLow                               PrinterStatusDetails = "materialLow"
	PrinterStatusDetails_MaterialNeeded                            PrinterStatusDetails = "materialNeeded"
	PrinterStatusDetails_MediaDrying                               PrinterStatusDetails = "mediaDrying"
	PrinterStatusDetails_MediaEmpty                                PrinterStatusDetails = "mediaEmpty"
	PrinterStatusDetails_MediaJam                                  PrinterStatusDetails = "mediaJam"
	PrinterStatusDetails_MediaLow                                  PrinterStatusDetails = "mediaLow"
	PrinterStatusDetails_MediaNeeded                               PrinterStatusDetails = "mediaNeeded"
	PrinterStatusDetails_MediaPathCannotDuplexMediaSelected        PrinterStatusDetails = "mediaPathCannotDuplexMediaSelected"
	PrinterStatusDetails_MediaPathFailure                          PrinterStatusDetails = "mediaPathFailure"
	PrinterStatusDetails_MediaPathInputEmpty                       PrinterStatusDetails = "mediaPathInputEmpty"
	PrinterStatusDetails_MediaPathInputFeedError                   PrinterStatusDetails = "mediaPathInputFeedError"
	PrinterStatusDetails_MediaPathInputJam                         PrinterStatusDetails = "mediaPathInputJam"
	PrinterStatusDetails_MediaPathInputRequest                     PrinterStatusDetails = "mediaPathInputRequest"
	PrinterStatusDetails_MediaPathJam                              PrinterStatusDetails = "mediaPathJam"
	PrinterStatusDetails_MediaPathMediaTrayAlmostFull              PrinterStatusDetails = "mediaPathMediaTrayAlmostFull"
	PrinterStatusDetails_MediaPathMediaTrayFull                    PrinterStatusDetails = "mediaPathMediaTrayFull"
	PrinterStatusDetails_MediaPathMediaTrayMissing                 PrinterStatusDetails = "mediaPathMediaTrayMissing"
	PrinterStatusDetails_MediaPathOutputFeedError                  PrinterStatusDetails = "mediaPathOutputFeedError"
	PrinterStatusDetails_MediaPathOutputFull                       PrinterStatusDetails = "mediaPathOutputFull"
	PrinterStatusDetails_MediaPathOutputJam                        PrinterStatusDetails = "mediaPathOutputJam"
	PrinterStatusDetails_MediaPathPickRollerFailure                PrinterStatusDetails = "mediaPathPickRollerFailure"
	PrinterStatusDetails_MediaPathPickRollerLifeOver               PrinterStatusDetails = "mediaPathPickRollerLifeOver"
	PrinterStatusDetails_MediaPathPickRollerLifeWarn               PrinterStatusDetails = "mediaPathPickRollerLifeWarn"
	PrinterStatusDetails_MediaPathPickRollerMissing                PrinterStatusDetails = "mediaPathPickRollerMissing"
	PrinterStatusDetails_MotorFailure                              PrinterStatusDetails = "motorFailure"
	PrinterStatusDetails_MovingToPaused                            PrinterStatusDetails = "movingToPaused"
	PrinterStatusDetails_None                                      PrinterStatusDetails = "none"
	PrinterStatusDetails_OpticalPhotoConductorLifeOver             PrinterStatusDetails = "opticalPhotoConductorLifeOver"
	PrinterStatusDetails_OpticalPhotoConductorNearEndOfLife        PrinterStatusDetails = "opticalPhotoConductorNearEndOfLife"
	PrinterStatusDetails_Other                                     PrinterStatusDetails = "other"
	PrinterStatusDetails_OutputAreaAlmostFull                      PrinterStatusDetails = "outputAreaAlmostFull"
	PrinterStatusDetails_OutputAreaFull                            PrinterStatusDetails = "outputAreaFull"
	PrinterStatusDetails_OutputMailboxSelectFailure                PrinterStatusDetails = "outputMailboxSelectFailure"
	PrinterStatusDetails_OutputMediaTrayFailure                    PrinterStatusDetails = "outputMediaTrayFailure"
	PrinterStatusDetails_OutputMediaTrayFeedError                  PrinterStatusDetails = "outputMediaTrayFeedError"
	PrinterStatusDetails_OutputMediaTrayJam                        PrinterStatusDetails = "outputMediaTrayJam"
	PrinterStatusDetails_OutputTrayMissing                         PrinterStatusDetails = "outputTrayMissing"
	PrinterStatusDetails_Paused                                    PrinterStatusDetails = "paused"
	PrinterStatusDetails_PerforaterAdded                           PrinterStatusDetails = "perforaterAdded"
	PrinterStatusDetails_PerforaterAlmostEmpty                     PrinterStatusDetails = "perforaterAlmostEmpty"
	PrinterStatusDetails_PerforaterAlmostFull                      PrinterStatusDetails = "perforaterAlmostFull"
	PrinterStatusDetails_PerforaterAtLimit                         PrinterStatusDetails = "perforaterAtLimit"
	PrinterStatusDetails_PerforaterClosed                          PrinterStatusDetails = "perforaterClosed"
	PrinterStatusDetails_PerforaterConfigurationChange             PrinterStatusDetails = "perforaterConfigurationChange"
	PrinterStatusDetails_PerforaterCoverClosed                     PrinterStatusDetails = "perforaterCoverClosed"
	PrinterStatusDetails_PerforaterCoverOpen                       PrinterStatusDetails = "perforaterCoverOpen"
	PrinterStatusDetails_PerforaterEmpty                           PrinterStatusDetails = "perforaterEmpty"
	PrinterStatusDetails_PerforaterFull                            PrinterStatusDetails = "perforaterFull"
	PrinterStatusDetails_PerforaterInterlockClosed                 PrinterStatusDetails = "perforaterInterlockClosed"
	PrinterStatusDetails_PerforaterInterlockOpen                   PrinterStatusDetails = "perforaterInterlockOpen"
	PrinterStatusDetails_PerforaterJam                             PrinterStatusDetails = "perforaterJam"
	PrinterStatusDetails_PerforaterLifeAlmostOver                  PrinterStatusDetails = "perforaterLifeAlmostOver"
	PrinterStatusDetails_PerforaterLifeOver                        PrinterStatusDetails = "perforaterLifeOver"
	PrinterStatusDetails_PerforaterMemoryExhausted                 PrinterStatusDetails = "perforaterMemoryExhausted"
	PrinterStatusDetails_PerforaterMissing                         PrinterStatusDetails = "perforaterMissing"
	PrinterStatusDetails_PerforaterMotorFailure                    PrinterStatusDetails = "perforaterMotorFailure"
	PrinterStatusDetails_PerforaterNearLimit                       PrinterStatusDetails = "perforaterNearLimit"
	PrinterStatusDetails_PerforaterOffline                         PrinterStatusDetails = "perforaterOffline"
	PrinterStatusDetails_PerforaterOpened                          PrinterStatusDetails = "perforaterOpened"
	PrinterStatusDetails_PerforaterOverTemperature                 PrinterStatusDetails = "perforaterOverTemperature"
	PrinterStatusDetails_PerforaterPowerSaver                      PrinterStatusDetails = "perforaterPowerSaver"
	PrinterStatusDetails_PerforaterRecoverableFailure              PrinterStatusDetails = "perforaterRecoverableFailure"
	PrinterStatusDetails_PerforaterRecoverableStorage              PrinterStatusDetails = "perforaterRecoverableStorage"
	PrinterStatusDetails_PerforaterRemoved                         PrinterStatusDetails = "perforaterRemoved"
	PrinterStatusDetails_PerforaterResourceAdded                   PrinterStatusDetails = "perforaterResourceAdded"
	PrinterStatusDetails_PerforaterResourceRemoved                 PrinterStatusDetails = "perforaterResourceRemoved"
	PrinterStatusDetails_PerforaterThermistorFailure               PrinterStatusDetails = "perforaterThermistorFailure"
	PrinterStatusDetails_PerforaterTimingFailure                   PrinterStatusDetails = "perforaterTimingFailure"
	PrinterStatusDetails_PerforaterTurnedOff                       PrinterStatusDetails = "perforaterTurnedOff"
	PrinterStatusDetails_PerforaterTurnedOn                        PrinterStatusDetails = "perforaterTurnedOn"
	PrinterStatusDetails_PerforaterUnderTemperature                PrinterStatusDetails = "perforaterUnderTemperature"
	PrinterStatusDetails_PerforaterUnrecoverableFailure            PrinterStatusDetails = "perforaterUnrecoverableFailure"
	PrinterStatusDetails_PerforaterUnrecoverableStorageError       PrinterStatusDetails = "perforaterUnrecoverableStorageError"
	PrinterStatusDetails_PerforaterWarmingUp                       PrinterStatusDetails = "perforaterWarmingUp"
	PrinterStatusDetails_PlatformCooling                           PrinterStatusDetails = "platformCooling"
	PrinterStatusDetails_PlatformFailure                           PrinterStatusDetails = "platformFailure"
	PrinterStatusDetails_PlatformHeating                           PrinterStatusDetails = "platformHeating"
	PrinterStatusDetails_PlatformTemperatureHigh                   PrinterStatusDetails = "platformTemperatureHigh"
	PrinterStatusDetails_PlatformTemperatureLow                    PrinterStatusDetails = "platformTemperatureLow"
	PrinterStatusDetails_PowerDown                                 PrinterStatusDetails = "powerDown"
	PrinterStatusDetails_PowerUp                                   PrinterStatusDetails = "powerUp"
	PrinterStatusDetails_PrinterManualReset                        PrinterStatusDetails = "printerManualReset"
	PrinterStatusDetails_PrinterNmsReset                           PrinterStatusDetails = "printerNmsReset"
	PrinterStatusDetails_PrinterReadyToPrint                       PrinterStatusDetails = "printerReadyToPrint"
	PrinterStatusDetails_PuncherAdded                              PrinterStatusDetails = "puncherAdded"
	PrinterStatusDetails_PuncherAlmostEmpty                        PrinterStatusDetails = "puncherAlmostEmpty"
	PrinterStatusDetails_PuncherAlmostFull                         PrinterStatusDetails = "puncherAlmostFull"
	PrinterStatusDetails_PuncherAtLimit                            PrinterStatusDetails = "puncherAtLimit"
	PrinterStatusDetails_PuncherClosed                             PrinterStatusDetails = "puncherClosed"
	PrinterStatusDetails_PuncherConfigurationChange                PrinterStatusDetails = "puncherConfigurationChange"
	PrinterStatusDetails_PuncherCoverClosed                        PrinterStatusDetails = "puncherCoverClosed"
	PrinterStatusDetails_PuncherCoverOpen                          PrinterStatusDetails = "puncherCoverOpen"
	PrinterStatusDetails_PuncherEmpty                              PrinterStatusDetails = "puncherEmpty"
	PrinterStatusDetails_PuncherFull                               PrinterStatusDetails = "puncherFull"
	PrinterStatusDetails_PuncherInterlockClosed                    PrinterStatusDetails = "puncherInterlockClosed"
	PrinterStatusDetails_PuncherInterlockOpen                      PrinterStatusDetails = "puncherInterlockOpen"
	PrinterStatusDetails_PuncherJam                                PrinterStatusDetails = "puncherJam"
	PrinterStatusDetails_PuncherLifeAlmostOver                     PrinterStatusDetails = "puncherLifeAlmostOver"
	PrinterStatusDetails_PuncherLifeOver                           PrinterStatusDetails = "puncherLifeOver"
	PrinterStatusDetails_PuncherMemoryExhausted                    PrinterStatusDetails = "puncherMemoryExhausted"
	PrinterStatusDetails_PuncherMissing                            PrinterStatusDetails = "puncherMissing"
	PrinterStatusDetails_PuncherMotorFailure                       PrinterStatusDetails = "puncherMotorFailure"
	PrinterStatusDetails_PuncherNearLimit                          PrinterStatusDetails = "puncherNearLimit"
	PrinterStatusDetails_PuncherOffline                            PrinterStatusDetails = "puncherOffline"
	PrinterStatusDetails_PuncherOpened                             PrinterStatusDetails = "puncherOpened"
	PrinterStatusDetails_PuncherOverTemperature                    PrinterStatusDetails = "puncherOverTemperature"
	PrinterStatusDetails_PuncherPowerSaver                         PrinterStatusDetails = "puncherPowerSaver"
	PrinterStatusDetails_PuncherRecoverableFailure                 PrinterStatusDetails = "puncherRecoverableFailure"
	PrinterStatusDetails_PuncherRecoverableStorage                 PrinterStatusDetails = "puncherRecoverableStorage"
	PrinterStatusDetails_PuncherRemoved                            PrinterStatusDetails = "puncherRemoved"
	PrinterStatusDetails_PuncherResourceAdded                      PrinterStatusDetails = "puncherResourceAdded"
	PrinterStatusDetails_PuncherResourceRemoved                    PrinterStatusDetails = "puncherResourceRemoved"
	PrinterStatusDetails_PuncherThermistorFailure                  PrinterStatusDetails = "puncherThermistorFailure"
	PrinterStatusDetails_PuncherTimingFailure                      PrinterStatusDetails = "puncherTimingFailure"
	PrinterStatusDetails_PuncherTurnedOff                          PrinterStatusDetails = "puncherTurnedOff"
	PrinterStatusDetails_PuncherTurnedOn                           PrinterStatusDetails = "puncherTurnedOn"
	PrinterStatusDetails_PuncherUnderTemperature                   PrinterStatusDetails = "puncherUnderTemperature"
	PrinterStatusDetails_PuncherUnrecoverableFailure               PrinterStatusDetails = "puncherUnrecoverableFailure"
	PrinterStatusDetails_PuncherUnrecoverableStorageError          PrinterStatusDetails = "puncherUnrecoverableStorageError"
	PrinterStatusDetails_PuncherWarmingUp                          PrinterStatusDetails = "puncherWarmingUp"
	PrinterStatusDetails_Resuming                                  PrinterStatusDetails = "resuming"
	PrinterStatusDetails_ScanMediaPathFailure                      PrinterStatusDetails = "scanMediaPathFailure"
	PrinterStatusDetails_ScanMediaPathInputEmpty                   PrinterStatusDetails = "scanMediaPathInputEmpty"
	PrinterStatusDetails_ScanMediaPathInputFeedError               PrinterStatusDetails = "scanMediaPathInputFeedError"
	PrinterStatusDetails_ScanMediaPathInputJam                     PrinterStatusDetails = "scanMediaPathInputJam"
	PrinterStatusDetails_ScanMediaPathInputRequest                 PrinterStatusDetails = "scanMediaPathInputRequest"
	PrinterStatusDetails_ScanMediaPathJam                          PrinterStatusDetails = "scanMediaPathJam"
	PrinterStatusDetails_ScanMediaPathOutputFeedError              PrinterStatusDetails = "scanMediaPathOutputFeedError"
	PrinterStatusDetails_ScanMediaPathOutputFull                   PrinterStatusDetails = "scanMediaPathOutputFull"
	PrinterStatusDetails_ScanMediaPathOutputJam                    PrinterStatusDetails = "scanMediaPathOutputJam"
	PrinterStatusDetails_ScanMediaPathPickRollerFailure            PrinterStatusDetails = "scanMediaPathPickRollerFailure"
	PrinterStatusDetails_ScanMediaPathPickRollerLifeOver           PrinterStatusDetails = "scanMediaPathPickRollerLifeOver"
	PrinterStatusDetails_ScanMediaPathPickRollerLifeWarn           PrinterStatusDetails = "scanMediaPathPickRollerLifeWarn"
	PrinterStatusDetails_ScanMediaPathPickRollerMissing            PrinterStatusDetails = "scanMediaPathPickRollerMissing"
	PrinterStatusDetails_ScanMediaPathTrayAlmostFull               PrinterStatusDetails = "scanMediaPathTrayAlmostFull"
	PrinterStatusDetails_ScanMediaPathTrayFull                     PrinterStatusDetails = "scanMediaPathTrayFull"
	PrinterStatusDetails_ScanMediaPathTrayMissing                  PrinterStatusDetails = "scanMediaPathTrayMissing"
	PrinterStatusDetails_ScannerLightFailure                       PrinterStatusDetails = "scannerLightFailure"
	PrinterStatusDetails_ScannerLightLifeAlmostOver                PrinterStatusDetails = "scannerLightLifeAlmostOver"
	PrinterStatusDetails_ScannerLightLifeOver                      PrinterStatusDetails = "scannerLightLifeOver"
	PrinterStatusDetails_ScannerLightMissing                       PrinterStatusDetails = "scannerLightMissing"
	PrinterStatusDetails_ScannerSensorFailure                      PrinterStatusDetails = "scannerSensorFailure"
	PrinterStatusDetails_ScannerSensorLifeAlmostOver               PrinterStatusDetails = "scannerSensorLifeAlmostOver"
	PrinterStatusDetails_ScannerSensorLifeOver                     PrinterStatusDetails = "scannerSensorLifeOver"
	PrinterStatusDetails_ScannerSensorMissing                      PrinterStatusDetails = "scannerSensorMissing"
	PrinterStatusDetails_SeparationCutterAdded                     PrinterStatusDetails = "separationCutterAdded"
	PrinterStatusDetails_SeparationCutterAlmostEmpty               PrinterStatusDetails = "separationCutterAlmostEmpty"
	PrinterStatusDetails_SeparationCutterAlmostFull                PrinterStatusDetails = "separationCutterAlmostFull"
	PrinterStatusDetails_SeparationCutterAtLimit                   PrinterStatusDetails = "separationCutterAtLimit"
	PrinterStatusDetails_SeparationCutterClosed                    PrinterStatusDetails = "separationCutterClosed"
	PrinterStatusDetails_SeparationCutterConfigurationChange       PrinterStatusDetails = "separationCutterConfigurationChange"
	PrinterStatusDetails_SeparationCutterCoverClosed               PrinterStatusDetails = "separationCutterCoverClosed"
	PrinterStatusDetails_SeparationCutterCoverOpen                 PrinterStatusDetails = "separationCutterCoverOpen"
	PrinterStatusDetails_SeparationCutterEmpty                     PrinterStatusDetails = "separationCutterEmpty"
	PrinterStatusDetails_SeparationCutterFull                      PrinterStatusDetails = "separationCutterFull"
	PrinterStatusDetails_SeparationCutterInterlockClosed           PrinterStatusDetails = "separationCutterInterlockClosed"
	PrinterStatusDetails_SeparationCutterInterlockOpen             PrinterStatusDetails = "separationCutterInterlockOpen"
	PrinterStatusDetails_SeparationCutterJam                       PrinterStatusDetails = "separationCutterJam"
	PrinterStatusDetails_SeparationCutterLifeAlmostOver            PrinterStatusDetails = "separationCutterLifeAlmostOver"
	PrinterStatusDetails_SeparationCutterLifeOver                  PrinterStatusDetails = "separationCutterLifeOver"
	PrinterStatusDetails_SeparationCutterMemoryExhausted           PrinterStatusDetails = "separationCutterMemoryExhausted"
	PrinterStatusDetails_SeparationCutterMissing                   PrinterStatusDetails = "separationCutterMissing"
	PrinterStatusDetails_SeparationCutterMotorFailure              PrinterStatusDetails = "separationCutterMotorFailure"
	PrinterStatusDetails_SeparationCutterNearLimit                 PrinterStatusDetails = "separationCutterNearLimit"
	PrinterStatusDetails_SeparationCutterOffline                   PrinterStatusDetails = "separationCutterOffline"
	PrinterStatusDetails_SeparationCutterOpened                    PrinterStatusDetails = "separationCutterOpened"
	PrinterStatusDetails_SeparationCutterOverTemperature           PrinterStatusDetails = "separationCutterOverTemperature"
	PrinterStatusDetails_SeparationCutterPowerSaver                PrinterStatusDetails = "separationCutterPowerSaver"
	PrinterStatusDetails_SeparationCutterRecoverableFailure        PrinterStatusDetails = "separationCutterRecoverableFailure"
	PrinterStatusDetails_SeparationCutterRecoverableStorage        PrinterStatusDetails = "separationCutterRecoverableStorage"
	PrinterStatusDetails_SeparationCutterRemoved                   PrinterStatusDetails = "separationCutterRemoved"
	PrinterStatusDetails_SeparationCutterResourceAdded             PrinterStatusDetails = "separationCutterResourceAdded"
	PrinterStatusDetails_SeparationCutterResourceRemoved           PrinterStatusDetails = "separationCutterResourceRemoved"
	PrinterStatusDetails_SeparationCutterThermistorFailure         PrinterStatusDetails = "separationCutterThermistorFailure"
	PrinterStatusDetails_SeparationCutterTimingFailure             PrinterStatusDetails = "separationCutterTimingFailure"
	PrinterStatusDetails_SeparationCutterTurnedOff                 PrinterStatusDetails = "separationCutterTurnedOff"
	PrinterStatusDetails_SeparationCutterTurnedOn                  PrinterStatusDetails = "separationCutterTurnedOn"
	PrinterStatusDetails_SeparationCutterUnderTemperature          PrinterStatusDetails = "separationCutterUnderTemperature"
	PrinterStatusDetails_SeparationCutterUnrecoverableFailure      PrinterStatusDetails = "separationCutterUnrecoverableFailure"
	PrinterStatusDetails_SeparationCutterUnrecoverableStorageError PrinterStatusDetails = "separationCutterUnrecoverableStorageError"
	PrinterStatusDetails_SeparationCutterWarmingUp                 PrinterStatusDetails = "separationCutterWarmingUp"
	PrinterStatusDetails_SheetRotatorAdded                         PrinterStatusDetails = "sheetRotatorAdded"
	PrinterStatusDetails_SheetRotatorAlmostEmpty                   PrinterStatusDetails = "sheetRotatorAlmostEmpty"
	PrinterStatusDetails_SheetRotatorAlmostFull                    PrinterStatusDetails = "sheetRotatorAlmostFull"
	PrinterStatusDetails_SheetRotatorAtLimit                       PrinterStatusDetails = "sheetRotatorAtLimit"
	PrinterStatusDetails_SheetRotatorClosed                        PrinterStatusDetails = "sheetRotatorClosed"
	PrinterStatusDetails_SheetRotatorConfigurationChange           PrinterStatusDetails = "sheetRotatorConfigurationChange"
	PrinterStatusDetails_SheetRotatorCoverClosed                   PrinterStatusDetails = "sheetRotatorCoverClosed"
	PrinterStatusDetails_SheetRotatorCoverOpen                     PrinterStatusDetails = "sheetRotatorCoverOpen"
	PrinterStatusDetails_SheetRotatorEmpty                         PrinterStatusDetails = "sheetRotatorEmpty"
	PrinterStatusDetails_SheetRotatorFull                          PrinterStatusDetails = "sheetRotatorFull"
	PrinterStatusDetails_SheetRotatorInterlockClosed               PrinterStatusDetails = "sheetRotatorInterlockClosed"
	PrinterStatusDetails_SheetRotatorInterlockOpen                 PrinterStatusDetails = "sheetRotatorInterlockOpen"
	PrinterStatusDetails_SheetRotatorJam                           PrinterStatusDetails = "sheetRotatorJam"
	PrinterStatusDetails_SheetRotatorLifeAlmostOver                PrinterStatusDetails = "sheetRotatorLifeAlmostOver"
	PrinterStatusDetails_SheetRotatorLifeOver                      PrinterStatusDetails = "sheetRotatorLifeOver"
	PrinterStatusDetails_SheetRotatorMemoryExhausted               PrinterStatusDetails = "sheetRotatorMemoryExhausted"
	PrinterStatusDetails_SheetRotatorMissing                       PrinterStatusDetails = "sheetRotatorMissing"
	PrinterStatusDetails_SheetRotatorMotorFailure                  PrinterStatusDetails = "sheetRotatorMotorFailure"
	PrinterStatusDetails_SheetRotatorNearLimit                     PrinterStatusDetails = "sheetRotatorNearLimit"
	PrinterStatusDetails_SheetRotatorOffline                       PrinterStatusDetails = "sheetRotatorOffline"
	PrinterStatusDetails_SheetRotatorOpened                        PrinterStatusDetails = "sheetRotatorOpened"
	PrinterStatusDetails_SheetRotatorOverTemperature               PrinterStatusDetails = "sheetRotatorOverTemperature"
	PrinterStatusDetails_SheetRotatorPowerSaver                    PrinterStatusDetails = "sheetRotatorPowerSaver"
	PrinterStatusDetails_SheetRotatorRecoverableFailure            PrinterStatusDetails = "sheetRotatorRecoverableFailure"
	PrinterStatusDetails_SheetRotatorRecoverableStorage            PrinterStatusDetails = "sheetRotatorRecoverableStorage"
	PrinterStatusDetails_SheetRotatorRemoved                       PrinterStatusDetails = "sheetRotatorRemoved"
	PrinterStatusDetails_SheetRotatorResourceAdded                 PrinterStatusDetails = "sheetRotatorResourceAdded"
	PrinterStatusDetails_SheetRotatorResourceRemoved               PrinterStatusDetails = "sheetRotatorResourceRemoved"
	PrinterStatusDetails_SheetRotatorThermistorFailure             PrinterStatusDetails = "sheetRotatorThermistorFailure"
	PrinterStatusDetails_SheetRotatorTimingFailure                 PrinterStatusDetails = "sheetRotatorTimingFailure"
	PrinterStatusDetails_SheetRotatorTurnedOff                     PrinterStatusDetails = "sheetRotatorTurnedOff"
	PrinterStatusDetails_SheetRotatorTurnedOn                      PrinterStatusDetails = "sheetRotatorTurnedOn"
	PrinterStatusDetails_SheetRotatorUnderTemperature              PrinterStatusDetails = "sheetRotatorUnderTemperature"
	PrinterStatusDetails_SheetRotatorUnrecoverableFailure          PrinterStatusDetails = "sheetRotatorUnrecoverableFailure"
	PrinterStatusDetails_SheetRotatorUnrecoverableStorageError     PrinterStatusDetails = "sheetRotatorUnrecoverableStorageError"
	PrinterStatusDetails_SheetRotatorWarmingUp                     PrinterStatusDetails = "sheetRotatorWarmingUp"
	PrinterStatusDetails_Shutdown                                  PrinterStatusDetails = "shutdown"
	PrinterStatusDetails_SlitterAdded                              PrinterStatusDetails = "slitterAdded"
	PrinterStatusDetails_SlitterAlmostEmpty                        PrinterStatusDetails = "slitterAlmostEmpty"
	PrinterStatusDetails_SlitterAlmostFull                         PrinterStatusDetails = "slitterAlmostFull"
	PrinterStatusDetails_SlitterAtLimit                            PrinterStatusDetails = "slitterAtLimit"
	PrinterStatusDetails_SlitterClosed                             PrinterStatusDetails = "slitterClosed"
	PrinterStatusDetails_SlitterConfigurationChange                PrinterStatusDetails = "slitterConfigurationChange"
	PrinterStatusDetails_SlitterCoverClosed                        PrinterStatusDetails = "slitterCoverClosed"
	PrinterStatusDetails_SlitterCoverOpen                          PrinterStatusDetails = "slitterCoverOpen"
	PrinterStatusDetails_SlitterEmpty                              PrinterStatusDetails = "slitterEmpty"
	PrinterStatusDetails_SlitterFull                               PrinterStatusDetails = "slitterFull"
	PrinterStatusDetails_SlitterInterlockClosed                    PrinterStatusDetails = "slitterInterlockClosed"
	PrinterStatusDetails_SlitterInterlockOpen                      PrinterStatusDetails = "slitterInterlockOpen"
	PrinterStatusDetails_SlitterJam                                PrinterStatusDetails = "slitterJam"
	PrinterStatusDetails_SlitterLifeAlmostOver                     PrinterStatusDetails = "slitterLifeAlmostOver"
	PrinterStatusDetails_SlitterLifeOver                           PrinterStatusDetails = "slitterLifeOver"
	PrinterStatusDetails_SlitterMemoryExhausted                    PrinterStatusDetails = "slitterMemoryExhausted"
	PrinterStatusDetails_SlitterMissing                            PrinterStatusDetails = "slitterMissing"
	PrinterStatusDetails_SlitterMotorFailure                       PrinterStatusDetails = "slitterMotorFailure"
	PrinterStatusDetails_SlitterNearLimit                          PrinterStatusDetails = "slitterNearLimit"
	PrinterStatusDetails_SlitterOffline                            PrinterStatusDetails = "slitterOffline"
	PrinterStatusDetails_SlitterOpened                             PrinterStatusDetails = "slitterOpened"
	PrinterStatusDetails_SlitterOverTemperature                    PrinterStatusDetails = "slitterOverTemperature"
	PrinterStatusDetails_SlitterPowerSaver                         PrinterStatusDetails = "slitterPowerSaver"
	PrinterStatusDetails_SlitterRecoverableFailure                 PrinterStatusDetails = "slitterRecoverableFailure"
	PrinterStatusDetails_SlitterRecoverableStorage                 PrinterStatusDetails = "slitterRecoverableStorage"
	PrinterStatusDetails_SlitterRemoved                            PrinterStatusDetails = "slitterRemoved"
	PrinterStatusDetails_SlitterResourceAdded                      PrinterStatusDetails = "slitterResourceAdded"
	PrinterStatusDetails_SlitterResourceRemoved                    PrinterStatusDetails = "slitterResourceRemoved"
	PrinterStatusDetails_SlitterThermistorFailure                  PrinterStatusDetails = "slitterThermistorFailure"
	PrinterStatusDetails_SlitterTimingFailure                      PrinterStatusDetails = "slitterTimingFailure"
	PrinterStatusDetails_SlitterTurnedOff                          PrinterStatusDetails = "slitterTurnedOff"
	PrinterStatusDetails_SlitterTurnedOn                           PrinterStatusDetails = "slitterTurnedOn"
	PrinterStatusDetails_SlitterUnderTemperature                   PrinterStatusDetails = "slitterUnderTemperature"
	PrinterStatusDetails_SlitterUnrecoverableFailure               PrinterStatusDetails = "slitterUnrecoverableFailure"
	PrinterStatusDetails_SlitterUnrecoverableStorageError          PrinterStatusDetails = "slitterUnrecoverableStorageError"
	PrinterStatusDetails_SlitterWarmingUp                          PrinterStatusDetails = "slitterWarmingUp"
	PrinterStatusDetails_SpoolAreaFull                             PrinterStatusDetails = "spoolAreaFull"
	PrinterStatusDetails_StackerAdded                              PrinterStatusDetails = "stackerAdded"
	PrinterStatusDetails_StackerAlmostEmpty                        PrinterStatusDetails = "stackerAlmostEmpty"
	PrinterStatusDetails_StackerAlmostFull                         PrinterStatusDetails = "stackerAlmostFull"
	PrinterStatusDetails_StackerAtLimit                            PrinterStatusDetails = "stackerAtLimit"
	PrinterStatusDetails_StackerClosed                             PrinterStatusDetails = "stackerClosed"
	PrinterStatusDetails_StackerConfigurationChange                PrinterStatusDetails = "stackerConfigurationChange"
	PrinterStatusDetails_StackerCoverClosed                        PrinterStatusDetails = "stackerCoverClosed"
	PrinterStatusDetails_StackerCoverOpen                          PrinterStatusDetails = "stackerCoverOpen"
	PrinterStatusDetails_StackerEmpty                              PrinterStatusDetails = "stackerEmpty"
	PrinterStatusDetails_StackerFull                               PrinterStatusDetails = "stackerFull"
	PrinterStatusDetails_StackerInterlockClosed                    PrinterStatusDetails = "stackerInterlockClosed"
	PrinterStatusDetails_StackerInterlockOpen                      PrinterStatusDetails = "stackerInterlockOpen"
	PrinterStatusDetails_StackerJam                                PrinterStatusDetails = "stackerJam"
	PrinterStatusDetails_StackerLifeAlmostOver                     PrinterStatusDetails = "stackerLifeAlmostOver"
	PrinterStatusDetails_StackerLifeOver                           PrinterStatusDetails = "stackerLifeOver"
	PrinterStatusDetails_StackerMemoryExhausted                    PrinterStatusDetails = "stackerMemoryExhausted"
	PrinterStatusDetails_StackerMissing                            PrinterStatusDetails = "stackerMissing"
	PrinterStatusDetails_StackerMotorFailure                       PrinterStatusDetails = "stackerMotorFailure"
	PrinterStatusDetails_StackerNearLimit                          PrinterStatusDetails = "stackerNearLimit"
	PrinterStatusDetails_StackerOffline                            PrinterStatusDetails = "stackerOffline"
	PrinterStatusDetails_StackerOpened                             PrinterStatusDetails = "stackerOpened"
	PrinterStatusDetails_StackerOverTemperature                    PrinterStatusDetails = "stackerOverTemperature"
	PrinterStatusDetails_StackerPowerSaver                         PrinterStatusDetails = "stackerPowerSaver"
	PrinterStatusDetails_StackerRecoverableFailure                 PrinterStatusDetails = "stackerRecoverableFailure"
	PrinterStatusDetails_StackerRecoverableStorage                 PrinterStatusDetails = "stackerRecoverableStorage"
	PrinterStatusDetails_StackerRemoved                            PrinterStatusDetails = "stackerRemoved"
	PrinterStatusDetails_StackerResourceAdded                      PrinterStatusDetails = "stackerResourceAdded"
	PrinterStatusDetails_StackerResourceRemoved                    PrinterStatusDetails = "stackerResourceRemoved"
	PrinterStatusDetails_StackerThermistorFailure                  PrinterStatusDetails = "stackerThermistorFailure"
	PrinterStatusDetails_StackerTimingFailure                      PrinterStatusDetails = "stackerTimingFailure"
	PrinterStatusDetails_StackerTurnedOff                          PrinterStatusDetails = "stackerTurnedOff"
	PrinterStatusDetails_StackerTurnedOn                           PrinterStatusDetails = "stackerTurnedOn"
	PrinterStatusDetails_StackerUnderTemperature                   PrinterStatusDetails = "stackerUnderTemperature"
	PrinterStatusDetails_StackerUnrecoverableFailure               PrinterStatusDetails = "stackerUnrecoverableFailure"
	PrinterStatusDetails_StackerUnrecoverableStorageError          PrinterStatusDetails = "stackerUnrecoverableStorageError"
	PrinterStatusDetails_StackerWarmingUp                          PrinterStatusDetails = "stackerWarmingUp"
	PrinterStatusDetails_Standby                                   PrinterStatusDetails = "standby"
	PrinterStatusDetails_StaplerAdded                              PrinterStatusDetails = "staplerAdded"
	PrinterStatusDetails_StaplerAlmostEmpty                        PrinterStatusDetails = "staplerAlmostEmpty"
	PrinterStatusDetails_StaplerAlmostFull                         PrinterStatusDetails = "staplerAlmostFull"
	PrinterStatusDetails_StaplerAtLimit                            PrinterStatusDetails = "staplerAtLimit"
	PrinterStatusDetails_StaplerClosed                             PrinterStatusDetails = "staplerClosed"
	PrinterStatusDetails_StaplerConfigurationChange                PrinterStatusDetails = "staplerConfigurationChange"
	PrinterStatusDetails_StaplerCoverClosed                        PrinterStatusDetails = "staplerCoverClosed"
	PrinterStatusDetails_StaplerCoverOpen                          PrinterStatusDetails = "staplerCoverOpen"
	PrinterStatusDetails_StaplerEmpty                              PrinterStatusDetails = "staplerEmpty"
	PrinterStatusDetails_StaplerFull                               PrinterStatusDetails = "staplerFull"
	PrinterStatusDetails_StaplerInterlockClosed                    PrinterStatusDetails = "staplerInterlockClosed"
	PrinterStatusDetails_StaplerInterlockOpen                      PrinterStatusDetails = "staplerInterlockOpen"
	PrinterStatusDetails_StaplerJam                                PrinterStatusDetails = "staplerJam"
	PrinterStatusDetails_StaplerLifeAlmostOver                     PrinterStatusDetails = "staplerLifeAlmostOver"
	PrinterStatusDetails_StaplerLifeOver                           PrinterStatusDetails = "staplerLifeOver"
	PrinterStatusDetails_StaplerMemoryExhausted                    PrinterStatusDetails = "staplerMemoryExhausted"
	PrinterStatusDetails_StaplerMissing                            PrinterStatusDetails = "staplerMissing"
	PrinterStatusDetails_StaplerMotorFailure                       PrinterStatusDetails = "staplerMotorFailure"
	PrinterStatusDetails_StaplerNearLimit                          PrinterStatusDetails = "staplerNearLimit"
	PrinterStatusDetails_StaplerOffline                            PrinterStatusDetails = "staplerOffline"
	PrinterStatusDetails_StaplerOpened                             PrinterStatusDetails = "staplerOpened"
	PrinterStatusDetails_StaplerOverTemperature                    PrinterStatusDetails = "staplerOverTemperature"
	PrinterStatusDetails_StaplerPowerSaver                         PrinterStatusDetails = "staplerPowerSaver"
	PrinterStatusDetails_StaplerRecoverableFailure                 PrinterStatusDetails = "staplerRecoverableFailure"
	PrinterStatusDetails_StaplerRecoverableStorage                 PrinterStatusDetails = "staplerRecoverableStorage"
	PrinterStatusDetails_StaplerRemoved                            PrinterStatusDetails = "staplerRemoved"
	PrinterStatusDetails_StaplerResourceAdded                      PrinterStatusDetails = "staplerResourceAdded"
	PrinterStatusDetails_StaplerResourceRemoved                    PrinterStatusDetails = "staplerResourceRemoved"
	PrinterStatusDetails_StaplerThermistorFailure                  PrinterStatusDetails = "staplerThermistorFailure"
	PrinterStatusDetails_StaplerTimingFailure                      PrinterStatusDetails = "staplerTimingFailure"
	PrinterStatusDetails_StaplerTurnedOff                          PrinterStatusDetails = "staplerTurnedOff"
	PrinterStatusDetails_StaplerTurnedOn                           PrinterStatusDetails = "staplerTurnedOn"
	PrinterStatusDetails_StaplerUnderTemperature                   PrinterStatusDetails = "staplerUnderTemperature"
	PrinterStatusDetails_StaplerUnrecoverableFailure               PrinterStatusDetails = "staplerUnrecoverableFailure"
	PrinterStatusDetails_StaplerUnrecoverableStorageError          PrinterStatusDetails = "staplerUnrecoverableStorageError"
	PrinterStatusDetails_StaplerWarmingUp                          PrinterStatusDetails = "staplerWarmingUp"
	PrinterStatusDetails_StitcherAdded                             PrinterStatusDetails = "stitcherAdded"
	PrinterStatusDetails_StitcherAlmostEmpty                       PrinterStatusDetails = "stitcherAlmostEmpty"
	PrinterStatusDetails_StitcherAlmostFull                        PrinterStatusDetails = "stitcherAlmostFull"
	PrinterStatusDetails_StitcherAtLimit                           PrinterStatusDetails = "stitcherAtLimit"
	PrinterStatusDetails_StitcherClosed                            PrinterStatusDetails = "stitcherClosed"
	PrinterStatusDetails_StitcherConfigurationChange               PrinterStatusDetails = "stitcherConfigurationChange"
	PrinterStatusDetails_StitcherCoverClosed                       PrinterStatusDetails = "stitcherCoverClosed"
	PrinterStatusDetails_StitcherCoverOpen                         PrinterStatusDetails = "stitcherCoverOpen"
	PrinterStatusDetails_StitcherEmpty                             PrinterStatusDetails = "stitcherEmpty"
	PrinterStatusDetails_StitcherFull                              PrinterStatusDetails = "stitcherFull"
	PrinterStatusDetails_StitcherInterlockClosed                   PrinterStatusDetails = "stitcherInterlockClosed"
	PrinterStatusDetails_StitcherInterlockOpen                     PrinterStatusDetails = "stitcherInterlockOpen"
	PrinterStatusDetails_StitcherJam                               PrinterStatusDetails = "stitcherJam"
	PrinterStatusDetails_StitcherLifeAlmostOver                    PrinterStatusDetails = "stitcherLifeAlmostOver"
	PrinterStatusDetails_StitcherLifeOver                          PrinterStatusDetails = "stitcherLifeOver"
	PrinterStatusDetails_StitcherMemoryExhausted                   PrinterStatusDetails = "stitcherMemoryExhausted"
	PrinterStatusDetails_StitcherMissing                           PrinterStatusDetails = "stitcherMissing"
	PrinterStatusDetails_StitcherMotorFailure                      PrinterStatusDetails = "stitcherMotorFailure"
	PrinterStatusDetails_StitcherNearLimit                         PrinterStatusDetails = "stitcherNearLimit"
	PrinterStatusDetails_StitcherOffline                           PrinterStatusDetails = "stitcherOffline"
	PrinterStatusDetails_StitcherOpened                            PrinterStatusDetails = "stitcherOpened"
	PrinterStatusDetails_StitcherOverTemperature                   PrinterStatusDetails = "stitcherOverTemperature"
	PrinterStatusDetails_StitcherPowerSaver                        PrinterStatusDetails = "stitcherPowerSaver"
	PrinterStatusDetails_StitcherRecoverableFailure                PrinterStatusDetails = "stitcherRecoverableFailure"
	PrinterStatusDetails_StitcherRecoverableStorage                PrinterStatusDetails = "stitcherRecoverableStorage"
	PrinterStatusDetails_StitcherRemoved                           PrinterStatusDetails = "stitcherRemoved"
	PrinterStatusDetails_StitcherResourceAdded                     PrinterStatusDetails = "stitcherResourceAdded"
	PrinterStatusDetails_StitcherResourceRemoved                   PrinterStatusDetails = "stitcherResourceRemoved"
	PrinterStatusDetails_StitcherThermistorFailure                 PrinterStatusDetails = "stitcherThermistorFailure"
	PrinterStatusDetails_StitcherTimingFailure                     PrinterStatusDetails = "stitcherTimingFailure"
	PrinterStatusDetails_StitcherTurnedOff                         PrinterStatusDetails = "stitcherTurnedOff"
	PrinterStatusDetails_StitcherTurnedOn                          PrinterStatusDetails = "stitcherTurnedOn"
	PrinterStatusDetails_StitcherUnderTemperature                  PrinterStatusDetails = "stitcherUnderTemperature"
	PrinterStatusDetails_StitcherUnrecoverableFailure              PrinterStatusDetails = "stitcherUnrecoverableFailure"
	PrinterStatusDetails_StitcherUnrecoverableStorageError         PrinterStatusDetails = "stitcherUnrecoverableStorageError"
	PrinterStatusDetails_StitcherWarmingUp                         PrinterStatusDetails = "stitcherWarmingUp"
	PrinterStatusDetails_StoppedPartially                          PrinterStatusDetails = "stoppedPartially"
	PrinterStatusDetails_Stopping                                  PrinterStatusDetails = "stopping"
	PrinterStatusDetails_SubunitAdded                              PrinterStatusDetails = "subunitAdded"
	PrinterStatusDetails_SubunitAlmostEmpty                        PrinterStatusDetails = "subunitAlmostEmpty"
	PrinterStatusDetails_SubunitAlmostFull                         PrinterStatusDetails = "subunitAlmostFull"
	PrinterStatusDetails_SubunitAtLimit                            PrinterStatusDetails = "subunitAtLimit"
	PrinterStatusDetails_SubunitClosed                             PrinterStatusDetails = "subunitClosed"
	PrinterStatusDetails_SubunitCoolingDown                        PrinterStatusDetails = "subunitCoolingDown"
	PrinterStatusDetails_SubunitEmpty                              PrinterStatusDetails = "subunitEmpty"
	PrinterStatusDetails_SubunitFull                               PrinterStatusDetails = "subunitFull"
	PrinterStatusDetails_SubunitLifeAlmostOver                     PrinterStatusDetails = "subunitLifeAlmostOver"
	PrinterStatusDetails_SubunitLifeOver                           PrinterStatusDetails = "subunitLifeOver"
	PrinterStatusDetails_SubunitMemoryExhausted                    PrinterStatusDetails = "subunitMemoryExhausted"
	PrinterStatusDetails_SubunitMissing                            PrinterStatusDetails = "subunitMissing"
	PrinterStatusDetails_SubunitMotorFailure                       PrinterStatusDetails = "subunitMotorFailure"
	PrinterStatusDetails_SubunitNearLimit                          PrinterStatusDetails = "subunitNearLimit"
	PrinterStatusDetails_SubunitOffline                            PrinterStatusDetails = "subunitOffline"
	PrinterStatusDetails_SubunitOpened                             PrinterStatusDetails = "subunitOpened"
	PrinterStatusDetails_SubunitOverTemperature                    PrinterStatusDetails = "subunitOverTemperature"
	PrinterStatusDetails_SubunitPowerSaver                         PrinterStatusDetails = "subunitPowerSaver"
	PrinterStatusDetails_SubunitRecoverableFailure                 PrinterStatusDetails = "subunitRecoverableFailure"
	PrinterStatusDetails_SubunitRecoverableStorage                 PrinterStatusDetails = "subunitRecoverableStorage"
	PrinterStatusDetails_SubunitRemoved                            PrinterStatusDetails = "subunitRemoved"
	PrinterStatusDetails_SubunitResourceAdded                      PrinterStatusDetails = "subunitResourceAdded"
	PrinterStatusDetails_SubunitResourceRemoved                    PrinterStatusDetails = "subunitResourceRemoved"
	PrinterStatusDetails_SubunitThermistorFailure                  PrinterStatusDetails = "subunitThermistorFailure"
	PrinterStatusDetails_SubunitTimingFailure                      PrinterStatusDetails = "subunitTimingFailure"
	PrinterStatusDetails_SubunitTurnedOff                          PrinterStatusDetails = "subunitTurnedOff"
	PrinterStatusDetails_SubunitTurnedOn                           PrinterStatusDetails = "subunitTurnedOn"
	PrinterStatusDetails_SubunitUnderTemperature                   PrinterStatusDetails = "subunitUnderTemperature"
	PrinterStatusDetails_SubunitUnrecoverableFailure               PrinterStatusDetails = "subunitUnrecoverableFailure"
	PrinterStatusDetails_SubunitUnrecoverableStorage               PrinterStatusDetails = "subunitUnrecoverableStorage"
	PrinterStatusDetails_SubunitWarmingUp                          PrinterStatusDetails = "subunitWarmingUp"
	PrinterStatusDetails_Suspend                                   PrinterStatusDetails = "suspend"
	PrinterStatusDetails_Testing                                   PrinterStatusDetails = "testing"
	PrinterStatusDetails_TimedOut                                  PrinterStatusDetails = "timedOut"
	PrinterStatusDetails_TonerEmpty                                PrinterStatusDetails = "tonerEmpty"
	PrinterStatusDetails_TonerLow                                  PrinterStatusDetails = "tonerLow"
	PrinterStatusDetails_TrimmerAdded                              PrinterStatusDetails = "trimmerAdded"
	PrinterStatusDetails_TrimmerAlmostEmpty                        PrinterStatusDetails = "trimmerAlmostEmpty"
	PrinterStatusDetails_TrimmerAlmostFull                         PrinterStatusDetails = "trimmerAlmostFull"
	PrinterStatusDetails_TrimmerAtLimit                            PrinterStatusDetails = "trimmerAtLimit"
	PrinterStatusDetails_TrimmerClosed                             PrinterStatusDetails = "trimmerClosed"
	PrinterStatusDetails_TrimmerConfigurationChange                PrinterStatusDetails = "trimmerConfigurationChange"
	PrinterStatusDetails_TrimmerCoverClosed                        PrinterStatusDetails = "trimmerCoverClosed"
	PrinterStatusDetails_TrimmerCoverOpen                          PrinterStatusDetails = "trimmerCoverOpen"
	PrinterStatusDetails_TrimmerEmpty                              PrinterStatusDetails = "trimmerEmpty"
	PrinterStatusDetails_TrimmerFull                               PrinterStatusDetails = "trimmerFull"
	PrinterStatusDetails_TrimmerInterlockClosed                    PrinterStatusDetails = "trimmerInterlockClosed"
	PrinterStatusDetails_TrimmerInterlockOpen                      PrinterStatusDetails = "trimmerInterlockOpen"
	PrinterStatusDetails_TrimmerJam                                PrinterStatusDetails = "trimmerJam"
	PrinterStatusDetails_TrimmerLifeAlmostOver                     PrinterStatusDetails = "trimmerLifeAlmostOver"
	PrinterStatusDetails_TrimmerLifeOver                           PrinterStatusDetails = "trimmerLifeOver"
	PrinterStatusDetails_TrimmerMemoryExhausted                    PrinterStatusDetails = "trimmerMemoryExhausted"
	PrinterStatusDetails_TrimmerMissing                            PrinterStatusDetails = "trimmerMissing"
	PrinterStatusDetails_TrimmerMotorFailure                       PrinterStatusDetails = "trimmerMotorFailure"
	PrinterStatusDetails_TrimmerNearLimit                          PrinterStatusDetails = "trimmerNearLimit"
	PrinterStatusDetails_TrimmerOffline                            PrinterStatusDetails = "trimmerOffline"
	PrinterStatusDetails_TrimmerOpened                             PrinterStatusDetails = "trimmerOpened"
	PrinterStatusDetails_TrimmerOverTemperature                    PrinterStatusDetails = "trimmerOverTemperature"
	PrinterStatusDetails_TrimmerPowerSaver                         PrinterStatusDetails = "trimmerPowerSaver"
	PrinterStatusDetails_TrimmerRecoverableFailure                 PrinterStatusDetails = "trimmerRecoverableFailure"
	PrinterStatusDetails_TrimmerRecoverableStorage                 PrinterStatusDetails = "trimmerRecoverableStorage"
	PrinterStatusDetails_TrimmerRemoved                            PrinterStatusDetails = "trimmerRemoved"
	PrinterStatusDetails_TrimmerResourceAdded                      PrinterStatusDetails = "trimmerResourceAdded"
	PrinterStatusDetails_TrimmerResourceRemoved                    PrinterStatusDetails = "trimmerResourceRemoved"
	PrinterStatusDetails_TrimmerThermistorFailure                  PrinterStatusDetails = "trimmerThermistorFailure"
	PrinterStatusDetails_TrimmerTimingFailure                      PrinterStatusDetails = "trimmerTimingFailure"
	PrinterStatusDetails_TrimmerTurnedOff                          PrinterStatusDetails = "trimmerTurnedOff"
	PrinterStatusDetails_TrimmerTurnedOn                           PrinterStatusDetails = "trimmerTurnedOn"
	PrinterStatusDetails_TrimmerUnderTemperature                   PrinterStatusDetails = "trimmerUnderTemperature"
	PrinterStatusDetails_TrimmerUnrecoverableFailure               PrinterStatusDetails = "trimmerUnrecoverableFailure"
	PrinterStatusDetails_TrimmerUnrecoverableStorageError          PrinterStatusDetails = "trimmerUnrecoverableStorageError"
	PrinterStatusDetails_TrimmerWarmingUp                          PrinterStatusDetails = "trimmerWarmingUp"
	PrinterStatusDetails_Unknown                                   PrinterStatusDetails = "unknown"
	PrinterStatusDetails_WrapperAdded                              PrinterStatusDetails = "wrapperAdded"
	PrinterStatusDetails_WrapperAlmostEmpty                        PrinterStatusDetails = "wrapperAlmostEmpty"
	PrinterStatusDetails_WrapperAlmostFull                         PrinterStatusDetails = "wrapperAlmostFull"
	PrinterStatusDetails_WrapperAtLimit                            PrinterStatusDetails = "wrapperAtLimit"
	PrinterStatusDetails_WrapperClosed                             PrinterStatusDetails = "wrapperClosed"
	PrinterStatusDetails_WrapperConfigurationChange                PrinterStatusDetails = "wrapperConfigurationChange"
	PrinterStatusDetails_WrapperCoverClosed                        PrinterStatusDetails = "wrapperCoverClosed"
	PrinterStatusDetails_WrapperCoverOpen                          PrinterStatusDetails = "wrapperCoverOpen"
	PrinterStatusDetails_WrapperEmpty                              PrinterStatusDetails = "wrapperEmpty"
	PrinterStatusDetails_WrapperFull                               PrinterStatusDetails = "wrapperFull"
	PrinterStatusDetails_WrapperInterlockClosed                    PrinterStatusDetails = "wrapperInterlockClosed"
	PrinterStatusDetails_WrapperInterlockOpen                      PrinterStatusDetails = "wrapperInterlockOpen"
	PrinterStatusDetails_WrapperJam                                PrinterStatusDetails = "wrapperJam"
	PrinterStatusDetails_WrapperLifeAlmostOver                     PrinterStatusDetails = "wrapperLifeAlmostOver"
	PrinterStatusDetails_WrapperLifeOver                           PrinterStatusDetails = "wrapperLifeOver"
	PrinterStatusDetails_WrapperMemoryExhausted                    PrinterStatusDetails = "wrapperMemoryExhausted"
	PrinterStatusDetails_WrapperMissing                            PrinterStatusDetails = "wrapperMissing"
	PrinterStatusDetails_WrapperMotorFailure                       PrinterStatusDetails = "wrapperMotorFailure"
	PrinterStatusDetails_WrapperNearLimit                          PrinterStatusDetails = "wrapperNearLimit"
	PrinterStatusDetails_WrapperOffline                            PrinterStatusDetails = "wrapperOffline"
	PrinterStatusDetails_WrapperOpened                             PrinterStatusDetails = "wrapperOpened"
	PrinterStatusDetails_WrapperOverTemperature                    PrinterStatusDetails = "wrapperOverTemperature"
	PrinterStatusDetails_WrapperPowerSaver                         PrinterStatusDetails = "wrapperPowerSaver"
	PrinterStatusDetails_WrapperRecoverableFailure                 PrinterStatusDetails = "wrapperRecoverableFailure"
	PrinterStatusDetails_WrapperRecoverableStorage                 PrinterStatusDetails = "wrapperRecoverableStorage"
	PrinterStatusDetails_WrapperRemoved                            PrinterStatusDetails = "wrapperRemoved"
	PrinterStatusDetails_WrapperResourceAdded                      PrinterStatusDetails = "wrapperResourceAdded"
	PrinterStatusDetails_WrapperResourceRemoved                    PrinterStatusDetails = "wrapperResourceRemoved"
	PrinterStatusDetails_WrapperThermistorFailure                  PrinterStatusDetails = "wrapperThermistorFailure"
	PrinterStatusDetails_WrapperTimingFailure                      PrinterStatusDetails = "wrapperTimingFailure"
	PrinterStatusDetails_WrapperTurnedOff                          PrinterStatusDetails = "wrapperTurnedOff"
	PrinterStatusDetails_WrapperTurnedOn                           PrinterStatusDetails = "wrapperTurnedOn"
	PrinterStatusDetails_WrapperUnderTemperature                   PrinterStatusDetails = "wrapperUnderTemperature"
	PrinterStatusDetails_WrapperUnrecoverableFailure               PrinterStatusDetails = "wrapperUnrecoverableFailure"
	PrinterStatusDetails_WrapperUnrecoverableStorageError          PrinterStatusDetails = "wrapperUnrecoverableStorageError"
	PrinterStatusDetails_WrapperWarmingUp                          PrinterStatusDetails = "wrapperWarmingUp"
)

func PossibleValuesForPrinterStatusDetails() []string {
	return []string{
		string(PrinterStatusDetails_AlertRemovalOfBinaryChangeEntry),
		string(PrinterStatusDetails_BanderAdded),
		string(PrinterStatusDetails_BanderAlmostEmpty),
		string(PrinterStatusDetails_BanderAlmostFull),
		string(PrinterStatusDetails_BanderAtLimit),
		string(PrinterStatusDetails_BanderClosed),
		string(PrinterStatusDetails_BanderConfigurationChange),
		string(PrinterStatusDetails_BanderCoverClosed),
		string(PrinterStatusDetails_BanderCoverOpen),
		string(PrinterStatusDetails_BanderEmpty),
		string(PrinterStatusDetails_BanderFull),
		string(PrinterStatusDetails_BanderInterlockClosed),
		string(PrinterStatusDetails_BanderInterlockOpen),
		string(PrinterStatusDetails_BanderJam),
		string(PrinterStatusDetails_BanderLifeAlmostOver),
		string(PrinterStatusDetails_BanderLifeOver),
		string(PrinterStatusDetails_BanderMemoryExhausted),
		string(PrinterStatusDetails_BanderMissing),
		string(PrinterStatusDetails_BanderMotorFailure),
		string(PrinterStatusDetails_BanderNearLimit),
		string(PrinterStatusDetails_BanderOffline),
		string(PrinterStatusDetails_BanderOpened),
		string(PrinterStatusDetails_BanderOverTemperature),
		string(PrinterStatusDetails_BanderPowerSaver),
		string(PrinterStatusDetails_BanderRecoverableFailure),
		string(PrinterStatusDetails_BanderRecoverableStorage),
		string(PrinterStatusDetails_BanderRemoved),
		string(PrinterStatusDetails_BanderResourceAdded),
		string(PrinterStatusDetails_BanderResourceRemoved),
		string(PrinterStatusDetails_BanderThermistorFailure),
		string(PrinterStatusDetails_BanderTimingFailure),
		string(PrinterStatusDetails_BanderTurnedOff),
		string(PrinterStatusDetails_BanderTurnedOn),
		string(PrinterStatusDetails_BanderUnderTemperature),
		string(PrinterStatusDetails_BanderUnrecoverableFailure),
		string(PrinterStatusDetails_BanderUnrecoverableStorageError),
		string(PrinterStatusDetails_BanderWarmingUp),
		string(PrinterStatusDetails_BinderAdded),
		string(PrinterStatusDetails_BinderAlmostEmpty),
		string(PrinterStatusDetails_BinderAlmostFull),
		string(PrinterStatusDetails_BinderAtLimit),
		string(PrinterStatusDetails_BinderClosed),
		string(PrinterStatusDetails_BinderConfigurationChange),
		string(PrinterStatusDetails_BinderCoverClosed),
		string(PrinterStatusDetails_BinderCoverOpen),
		string(PrinterStatusDetails_BinderEmpty),
		string(PrinterStatusDetails_BinderFull),
		string(PrinterStatusDetails_BinderInterlockClosed),
		string(PrinterStatusDetails_BinderInterlockOpen),
		string(PrinterStatusDetails_BinderJam),
		string(PrinterStatusDetails_BinderLifeAlmostOver),
		string(PrinterStatusDetails_BinderLifeOver),
		string(PrinterStatusDetails_BinderMemoryExhausted),
		string(PrinterStatusDetails_BinderMissing),
		string(PrinterStatusDetails_BinderMotorFailure),
		string(PrinterStatusDetails_BinderNearLimit),
		string(PrinterStatusDetails_BinderOffline),
		string(PrinterStatusDetails_BinderOpened),
		string(PrinterStatusDetails_BinderOverTemperature),
		string(PrinterStatusDetails_BinderPowerSaver),
		string(PrinterStatusDetails_BinderRecoverableFailure),
		string(PrinterStatusDetails_BinderRecoverableStorage),
		string(PrinterStatusDetails_BinderRemoved),
		string(PrinterStatusDetails_BinderResourceAdded),
		string(PrinterStatusDetails_BinderResourceRemoved),
		string(PrinterStatusDetails_BinderThermistorFailure),
		string(PrinterStatusDetails_BinderTimingFailure),
		string(PrinterStatusDetails_BinderTurnedOff),
		string(PrinterStatusDetails_BinderTurnedOn),
		string(PrinterStatusDetails_BinderUnderTemperature),
		string(PrinterStatusDetails_BinderUnrecoverableFailure),
		string(PrinterStatusDetails_BinderUnrecoverableStorageError),
		string(PrinterStatusDetails_BinderWarmingUp),
		string(PrinterStatusDetails_CameraFailure),
		string(PrinterStatusDetails_ChamberCooling),
		string(PrinterStatusDetails_ChamberFailure),
		string(PrinterStatusDetails_ChamberHeating),
		string(PrinterStatusDetails_ChamberTemperatureHigh),
		string(PrinterStatusDetails_ChamberTemperatureLow),
		string(PrinterStatusDetails_CleanerLifeAlmostOver),
		string(PrinterStatusDetails_CleanerLifeOver),
		string(PrinterStatusDetails_ConfigurationChange),
		string(PrinterStatusDetails_ConnectingToDevice),
		string(PrinterStatusDetails_CoverOpen),
		string(PrinterStatusDetails_Deactivated),
		string(PrinterStatusDetails_Deleted),
		string(PrinterStatusDetails_DeveloperEmpty),
		string(PrinterStatusDetails_DeveloperLow),
		string(PrinterStatusDetails_DieCutterAdded),
		string(PrinterStatusDetails_DieCutterAlmostEmpty),
		string(PrinterStatusDetails_DieCutterAlmostFull),
		string(PrinterStatusDetails_DieCutterAtLimit),
		string(PrinterStatusDetails_DieCutterClosed),
		string(PrinterStatusDetails_DieCutterConfigurationChange),
		string(PrinterStatusDetails_DieCutterCoverClosed),
		string(PrinterStatusDetails_DieCutterCoverOpen),
		string(PrinterStatusDetails_DieCutterEmpty),
		string(PrinterStatusDetails_DieCutterFull),
		string(PrinterStatusDetails_DieCutterInterlockClosed),
		string(PrinterStatusDetails_DieCutterInterlockOpen),
		string(PrinterStatusDetails_DieCutterJam),
		string(PrinterStatusDetails_DieCutterLifeAlmostOver),
		string(PrinterStatusDetails_DieCutterLifeOver),
		string(PrinterStatusDetails_DieCutterMemoryExhausted),
		string(PrinterStatusDetails_DieCutterMissing),
		string(PrinterStatusDetails_DieCutterMotorFailure),
		string(PrinterStatusDetails_DieCutterNearLimit),
		string(PrinterStatusDetails_DieCutterOffline),
		string(PrinterStatusDetails_DieCutterOpened),
		string(PrinterStatusDetails_DieCutterOverTemperature),
		string(PrinterStatusDetails_DieCutterPowerSaver),
		string(PrinterStatusDetails_DieCutterRecoverableFailure),
		string(PrinterStatusDetails_DieCutterRecoverableStorage),
		string(PrinterStatusDetails_DieCutterRemoved),
		string(PrinterStatusDetails_DieCutterResourceAdded),
		string(PrinterStatusDetails_DieCutterResourceRemoved),
		string(PrinterStatusDetails_DieCutterThermistorFailure),
		string(PrinterStatusDetails_DieCutterTimingFailure),
		string(PrinterStatusDetails_DieCutterTurnedOff),
		string(PrinterStatusDetails_DieCutterTurnedOn),
		string(PrinterStatusDetails_DieCutterUnderTemperature),
		string(PrinterStatusDetails_DieCutterUnrecoverableFailure),
		string(PrinterStatusDetails_DieCutterUnrecoverableStorageError),
		string(PrinterStatusDetails_DieCutterWarmingUp),
		string(PrinterStatusDetails_DoorOpen),
		string(PrinterStatusDetails_ExtruderCooling),
		string(PrinterStatusDetails_ExtruderFailure),
		string(PrinterStatusDetails_ExtruderHeating),
		string(PrinterStatusDetails_ExtruderJam),
		string(PrinterStatusDetails_ExtruderTemperatureHigh),
		string(PrinterStatusDetails_ExtruderTemperatureLow),
		string(PrinterStatusDetails_FanFailure),
		string(PrinterStatusDetails_FaxModemLifeAlmostOver),
		string(PrinterStatusDetails_FaxModemLifeOver),
		string(PrinterStatusDetails_FaxModemMissing),
		string(PrinterStatusDetails_FaxModemTurnedOff),
		string(PrinterStatusDetails_FaxModemTurnedOn),
		string(PrinterStatusDetails_FolderAdded),
		string(PrinterStatusDetails_FolderAlmostEmpty),
		string(PrinterStatusDetails_FolderAlmostFull),
		string(PrinterStatusDetails_FolderAtLimit),
		string(PrinterStatusDetails_FolderClosed),
		string(PrinterStatusDetails_FolderConfigurationChange),
		string(PrinterStatusDetails_FolderCoverClosed),
		string(PrinterStatusDetails_FolderCoverOpen),
		string(PrinterStatusDetails_FolderEmpty),
		string(PrinterStatusDetails_FolderFull),
		string(PrinterStatusDetails_FolderInterlockClosed),
		string(PrinterStatusDetails_FolderInterlockOpen),
		string(PrinterStatusDetails_FolderJam),
		string(PrinterStatusDetails_FolderLifeAlmostOver),
		string(PrinterStatusDetails_FolderLifeOver),
		string(PrinterStatusDetails_FolderMemoryExhausted),
		string(PrinterStatusDetails_FolderMissing),
		string(PrinterStatusDetails_FolderMotorFailure),
		string(PrinterStatusDetails_FolderNearLimit),
		string(PrinterStatusDetails_FolderOffline),
		string(PrinterStatusDetails_FolderOpened),
		string(PrinterStatusDetails_FolderOverTemperature),
		string(PrinterStatusDetails_FolderPowerSaver),
		string(PrinterStatusDetails_FolderRecoverableFailure),
		string(PrinterStatusDetails_FolderRecoverableStorage),
		string(PrinterStatusDetails_FolderRemoved),
		string(PrinterStatusDetails_FolderResourceAdded),
		string(PrinterStatusDetails_FolderResourceRemoved),
		string(PrinterStatusDetails_FolderThermistorFailure),
		string(PrinterStatusDetails_FolderTimingFailure),
		string(PrinterStatusDetails_FolderTurnedOff),
		string(PrinterStatusDetails_FolderTurnedOn),
		string(PrinterStatusDetails_FolderUnderTemperature),
		string(PrinterStatusDetails_FolderUnrecoverableFailure),
		string(PrinterStatusDetails_FolderUnrecoverableStorageError),
		string(PrinterStatusDetails_FolderWarmingUp),
		string(PrinterStatusDetails_FuserOverTemp),
		string(PrinterStatusDetails_FuserUnderTemp),
		string(PrinterStatusDetails_Hibernate),
		string(PrinterStatusDetails_HoldNewJobs),
		string(PrinterStatusDetails_IdentifyPrinterRequested),
		string(PrinterStatusDetails_ImprinterAdded),
		string(PrinterStatusDetails_ImprinterAlmostEmpty),
		string(PrinterStatusDetails_ImprinterAlmostFull),
		string(PrinterStatusDetails_ImprinterAtLimit),
		string(PrinterStatusDetails_ImprinterClosed),
		string(PrinterStatusDetails_ImprinterConfigurationChange),
		string(PrinterStatusDetails_ImprinterCoverClosed),
		string(PrinterStatusDetails_ImprinterCoverOpen),
		string(PrinterStatusDetails_ImprinterEmpty),
		string(PrinterStatusDetails_ImprinterFull),
		string(PrinterStatusDetails_ImprinterInterlockClosed),
		string(PrinterStatusDetails_ImprinterInterlockOpen),
		string(PrinterStatusDetails_ImprinterJam),
		string(PrinterStatusDetails_ImprinterLifeAlmostOver),
		string(PrinterStatusDetails_ImprinterLifeOver),
		string(PrinterStatusDetails_ImprinterMemoryExhausted),
		string(PrinterStatusDetails_ImprinterMissing),
		string(PrinterStatusDetails_ImprinterMotorFailure),
		string(PrinterStatusDetails_ImprinterNearLimit),
		string(PrinterStatusDetails_ImprinterOffline),
		string(PrinterStatusDetails_ImprinterOpened),
		string(PrinterStatusDetails_ImprinterOverTemperature),
		string(PrinterStatusDetails_ImprinterPowerSaver),
		string(PrinterStatusDetails_ImprinterRecoverableFailure),
		string(PrinterStatusDetails_ImprinterRecoverableStorage),
		string(PrinterStatusDetails_ImprinterRemoved),
		string(PrinterStatusDetails_ImprinterResourceAdded),
		string(PrinterStatusDetails_ImprinterResourceRemoved),
		string(PrinterStatusDetails_ImprinterThermistorFailure),
		string(PrinterStatusDetails_ImprinterTimingFailure),
		string(PrinterStatusDetails_ImprinterTurnedOff),
		string(PrinterStatusDetails_ImprinterTurnedOn),
		string(PrinterStatusDetails_ImprinterUnderTemperature),
		string(PrinterStatusDetails_ImprinterUnrecoverableFailure),
		string(PrinterStatusDetails_ImprinterUnrecoverableStorageError),
		string(PrinterStatusDetails_ImprinterWarmingUp),
		string(PrinterStatusDetails_InputCannotFeedSizeSelected),
		string(PrinterStatusDetails_InputManualInputRequest),
		string(PrinterStatusDetails_InputMediaColorChange),
		string(PrinterStatusDetails_InputMediaFormPartsChange),
		string(PrinterStatusDetails_InputMediaSizeChange),
		string(PrinterStatusDetails_InputMediaTrayFailure),
		string(PrinterStatusDetails_InputMediaTrayFeedError),
		string(PrinterStatusDetails_InputMediaTrayJam),
		string(PrinterStatusDetails_InputMediaTypeChange),
		string(PrinterStatusDetails_InputMediaWeightChange),
		string(PrinterStatusDetails_InputPickRollerFailure),
		string(PrinterStatusDetails_InputPickRollerLifeOver),
		string(PrinterStatusDetails_InputPickRollerLifeWarn),
		string(PrinterStatusDetails_InputPickRollerMissing),
		string(PrinterStatusDetails_InputTrayElevationFailure),
		string(PrinterStatusDetails_InputTrayMissing),
		string(PrinterStatusDetails_InputTrayPositionFailure),
		string(PrinterStatusDetails_InserterAdded),
		string(PrinterStatusDetails_InserterAlmostEmpty),
		string(PrinterStatusDetails_InserterAlmostFull),
		string(PrinterStatusDetails_InserterAtLimit),
		string(PrinterStatusDetails_InserterClosed),
		string(PrinterStatusDetails_InserterConfigurationChange),
		string(PrinterStatusDetails_InserterCoverClosed),
		string(PrinterStatusDetails_InserterCoverOpen),
		string(PrinterStatusDetails_InserterEmpty),
		string(PrinterStatusDetails_InserterFull),
		string(PrinterStatusDetails_InserterInterlockClosed),
		string(PrinterStatusDetails_InserterInterlockOpen),
		string(PrinterStatusDetails_InserterJam),
		string(PrinterStatusDetails_InserterLifeAlmostOver),
		string(PrinterStatusDetails_InserterLifeOver),
		string(PrinterStatusDetails_InserterMemoryExhausted),
		string(PrinterStatusDetails_InserterMissing),
		string(PrinterStatusDetails_InserterMotorFailure),
		string(PrinterStatusDetails_InserterNearLimit),
		string(PrinterStatusDetails_InserterOffline),
		string(PrinterStatusDetails_InserterOpened),
		string(PrinterStatusDetails_InserterOverTemperature),
		string(PrinterStatusDetails_InserterPowerSaver),
		string(PrinterStatusDetails_InserterRecoverableFailure),
		string(PrinterStatusDetails_InserterRecoverableStorage),
		string(PrinterStatusDetails_InserterRemoved),
		string(PrinterStatusDetails_InserterResourceAdded),
		string(PrinterStatusDetails_InserterResourceRemoved),
		string(PrinterStatusDetails_InserterThermistorFailure),
		string(PrinterStatusDetails_InserterTimingFailure),
		string(PrinterStatusDetails_InserterTurnedOff),
		string(PrinterStatusDetails_InserterTurnedOn),
		string(PrinterStatusDetails_InserterUnderTemperature),
		string(PrinterStatusDetails_InserterUnrecoverableFailure),
		string(PrinterStatusDetails_InserterUnrecoverableStorageError),
		string(PrinterStatusDetails_InserterWarmingUp),
		string(PrinterStatusDetails_InterlockClosed),
		string(PrinterStatusDetails_InterlockOpen),
		string(PrinterStatusDetails_InterpreterCartridgeAdded),
		string(PrinterStatusDetails_InterpreterCartridgeDeleted),
		string(PrinterStatusDetails_InterpreterComplexPageEncountered),
		string(PrinterStatusDetails_InterpreterMemoryDecrease),
		string(PrinterStatusDetails_InterpreterMemoryIncrease),
		string(PrinterStatusDetails_InterpreterResourceAdded),
		string(PrinterStatusDetails_InterpreterResourceDeleted),
		string(PrinterStatusDetails_InterpreterResourceUnavailable),
		string(PrinterStatusDetails_LampAtEol),
		string(PrinterStatusDetails_LampFailure),
		string(PrinterStatusDetails_LampNearEol),
		string(PrinterStatusDetails_LaserAtEol),
		string(PrinterStatusDetails_LaserFailure),
		string(PrinterStatusDetails_LaserNearEol),
		string(PrinterStatusDetails_MakeEnvelopeAdded),
		string(PrinterStatusDetails_MakeEnvelopeAlmostEmpty),
		string(PrinterStatusDetails_MakeEnvelopeAlmostFull),
		string(PrinterStatusDetails_MakeEnvelopeAtLimit),
		string(PrinterStatusDetails_MakeEnvelopeClosed),
		string(PrinterStatusDetails_MakeEnvelopeConfigurationChange),
		string(PrinterStatusDetails_MakeEnvelopeCoverClosed),
		string(PrinterStatusDetails_MakeEnvelopeCoverOpen),
		string(PrinterStatusDetails_MakeEnvelopeEmpty),
		string(PrinterStatusDetails_MakeEnvelopeFull),
		string(PrinterStatusDetails_MakeEnvelopeInterlockClosed),
		string(PrinterStatusDetails_MakeEnvelopeInterlockOpen),
		string(PrinterStatusDetails_MakeEnvelopeJam),
		string(PrinterStatusDetails_MakeEnvelopeLifeAlmostOver),
		string(PrinterStatusDetails_MakeEnvelopeLifeOver),
		string(PrinterStatusDetails_MakeEnvelopeMemoryExhausted),
		string(PrinterStatusDetails_MakeEnvelopeMissing),
		string(PrinterStatusDetails_MakeEnvelopeMotorFailure),
		string(PrinterStatusDetails_MakeEnvelopeNearLimit),
		string(PrinterStatusDetails_MakeEnvelopeOffline),
		string(PrinterStatusDetails_MakeEnvelopeOpened),
		string(PrinterStatusDetails_MakeEnvelopeOverTemperature),
		string(PrinterStatusDetails_MakeEnvelopePowerSaver),
		string(PrinterStatusDetails_MakeEnvelopeRecoverableFailure),
		string(PrinterStatusDetails_MakeEnvelopeRecoverableStorage),
		string(PrinterStatusDetails_MakeEnvelopeRemoved),
		string(PrinterStatusDetails_MakeEnvelopeResourceAdded),
		string(PrinterStatusDetails_MakeEnvelopeResourceRemoved),
		string(PrinterStatusDetails_MakeEnvelopeThermistorFailure),
		string(PrinterStatusDetails_MakeEnvelopeTimingFailure),
		string(PrinterStatusDetails_MakeEnvelopeTurnedOff),
		string(PrinterStatusDetails_MakeEnvelopeTurnedOn),
		string(PrinterStatusDetails_MakeEnvelopeUnderTemperature),
		string(PrinterStatusDetails_MakeEnvelopeUnrecoverableFailure),
		string(PrinterStatusDetails_MakeEnvelopeUnrecoverableStorageError),
		string(PrinterStatusDetails_MakeEnvelopeWarmingUp),
		string(PrinterStatusDetails_MarkerAdjustingPrintQuality),
		string(PrinterStatusDetails_MarkerCleanerMissing),
		string(PrinterStatusDetails_MarkerDeveloperAlmostEmpty),
		string(PrinterStatusDetails_MarkerDeveloperEmpty),
		string(PrinterStatusDetails_MarkerDeveloperMissing),
		string(PrinterStatusDetails_MarkerFuserMissing),
		string(PrinterStatusDetails_MarkerFuserThermistorFailure),
		string(PrinterStatusDetails_MarkerFuserTimingFailure),
		string(PrinterStatusDetails_MarkerInkAlmostEmpty),
		string(PrinterStatusDetails_MarkerInkEmpty),
		string(PrinterStatusDetails_MarkerInkMissing),
		string(PrinterStatusDetails_MarkerOpcMissing),
		string(PrinterStatusDetails_MarkerPrintRibbonAlmostEmpty),
		string(PrinterStatusDetails_MarkerPrintRibbonEmpty),
		string(PrinterStatusDetails_MarkerPrintRibbonMissing),
		string(PrinterStatusDetails_MarkerSupplyAlmostEmpty),
		string(PrinterStatusDetails_MarkerSupplyEmpty),
		string(PrinterStatusDetails_MarkerSupplyLow),
		string(PrinterStatusDetails_MarkerSupplyMissing),
		string(PrinterStatusDetails_MarkerTonerCartridgeMissing),
		string(PrinterStatusDetails_MarkerTonerMissing),
		string(PrinterStatusDetails_MarkerWasteAlmostFull),
		string(PrinterStatusDetails_MarkerWasteFull),
		string(PrinterStatusDetails_MarkerWasteInkReceptacleAlmostFull),
		string(PrinterStatusDetails_MarkerWasteInkReceptacleFull),
		string(PrinterStatusDetails_MarkerWasteInkReceptacleMissing),
		string(PrinterStatusDetails_MarkerWasteMissing),
		string(PrinterStatusDetails_MarkerWasteTonerReceptacleAlmostFull),
		string(PrinterStatusDetails_MarkerWasteTonerReceptacleFull),
		string(PrinterStatusDetails_MarkerWasteTonerReceptacleMissing),
		string(PrinterStatusDetails_MaterialEmpty),
		string(PrinterStatusDetails_MaterialLow),
		string(PrinterStatusDetails_MaterialNeeded),
		string(PrinterStatusDetails_MediaDrying),
		string(PrinterStatusDetails_MediaEmpty),
		string(PrinterStatusDetails_MediaJam),
		string(PrinterStatusDetails_MediaLow),
		string(PrinterStatusDetails_MediaNeeded),
		string(PrinterStatusDetails_MediaPathCannotDuplexMediaSelected),
		string(PrinterStatusDetails_MediaPathFailure),
		string(PrinterStatusDetails_MediaPathInputEmpty),
		string(PrinterStatusDetails_MediaPathInputFeedError),
		string(PrinterStatusDetails_MediaPathInputJam),
		string(PrinterStatusDetails_MediaPathInputRequest),
		string(PrinterStatusDetails_MediaPathJam),
		string(PrinterStatusDetails_MediaPathMediaTrayAlmostFull),
		string(PrinterStatusDetails_MediaPathMediaTrayFull),
		string(PrinterStatusDetails_MediaPathMediaTrayMissing),
		string(PrinterStatusDetails_MediaPathOutputFeedError),
		string(PrinterStatusDetails_MediaPathOutputFull),
		string(PrinterStatusDetails_MediaPathOutputJam),
		string(PrinterStatusDetails_MediaPathPickRollerFailure),
		string(PrinterStatusDetails_MediaPathPickRollerLifeOver),
		string(PrinterStatusDetails_MediaPathPickRollerLifeWarn),
		string(PrinterStatusDetails_MediaPathPickRollerMissing),
		string(PrinterStatusDetails_MotorFailure),
		string(PrinterStatusDetails_MovingToPaused),
		string(PrinterStatusDetails_None),
		string(PrinterStatusDetails_OpticalPhotoConductorLifeOver),
		string(PrinterStatusDetails_OpticalPhotoConductorNearEndOfLife),
		string(PrinterStatusDetails_Other),
		string(PrinterStatusDetails_OutputAreaAlmostFull),
		string(PrinterStatusDetails_OutputAreaFull),
		string(PrinterStatusDetails_OutputMailboxSelectFailure),
		string(PrinterStatusDetails_OutputMediaTrayFailure),
		string(PrinterStatusDetails_OutputMediaTrayFeedError),
		string(PrinterStatusDetails_OutputMediaTrayJam),
		string(PrinterStatusDetails_OutputTrayMissing),
		string(PrinterStatusDetails_Paused),
		string(PrinterStatusDetails_PerforaterAdded),
		string(PrinterStatusDetails_PerforaterAlmostEmpty),
		string(PrinterStatusDetails_PerforaterAlmostFull),
		string(PrinterStatusDetails_PerforaterAtLimit),
		string(PrinterStatusDetails_PerforaterClosed),
		string(PrinterStatusDetails_PerforaterConfigurationChange),
		string(PrinterStatusDetails_PerforaterCoverClosed),
		string(PrinterStatusDetails_PerforaterCoverOpen),
		string(PrinterStatusDetails_PerforaterEmpty),
		string(PrinterStatusDetails_PerforaterFull),
		string(PrinterStatusDetails_PerforaterInterlockClosed),
		string(PrinterStatusDetails_PerforaterInterlockOpen),
		string(PrinterStatusDetails_PerforaterJam),
		string(PrinterStatusDetails_PerforaterLifeAlmostOver),
		string(PrinterStatusDetails_PerforaterLifeOver),
		string(PrinterStatusDetails_PerforaterMemoryExhausted),
		string(PrinterStatusDetails_PerforaterMissing),
		string(PrinterStatusDetails_PerforaterMotorFailure),
		string(PrinterStatusDetails_PerforaterNearLimit),
		string(PrinterStatusDetails_PerforaterOffline),
		string(PrinterStatusDetails_PerforaterOpened),
		string(PrinterStatusDetails_PerforaterOverTemperature),
		string(PrinterStatusDetails_PerforaterPowerSaver),
		string(PrinterStatusDetails_PerforaterRecoverableFailure),
		string(PrinterStatusDetails_PerforaterRecoverableStorage),
		string(PrinterStatusDetails_PerforaterRemoved),
		string(PrinterStatusDetails_PerforaterResourceAdded),
		string(PrinterStatusDetails_PerforaterResourceRemoved),
		string(PrinterStatusDetails_PerforaterThermistorFailure),
		string(PrinterStatusDetails_PerforaterTimingFailure),
		string(PrinterStatusDetails_PerforaterTurnedOff),
		string(PrinterStatusDetails_PerforaterTurnedOn),
		string(PrinterStatusDetails_PerforaterUnderTemperature),
		string(PrinterStatusDetails_PerforaterUnrecoverableFailure),
		string(PrinterStatusDetails_PerforaterUnrecoverableStorageError),
		string(PrinterStatusDetails_PerforaterWarmingUp),
		string(PrinterStatusDetails_PlatformCooling),
		string(PrinterStatusDetails_PlatformFailure),
		string(PrinterStatusDetails_PlatformHeating),
		string(PrinterStatusDetails_PlatformTemperatureHigh),
		string(PrinterStatusDetails_PlatformTemperatureLow),
		string(PrinterStatusDetails_PowerDown),
		string(PrinterStatusDetails_PowerUp),
		string(PrinterStatusDetails_PrinterManualReset),
		string(PrinterStatusDetails_PrinterNmsReset),
		string(PrinterStatusDetails_PrinterReadyToPrint),
		string(PrinterStatusDetails_PuncherAdded),
		string(PrinterStatusDetails_PuncherAlmostEmpty),
		string(PrinterStatusDetails_PuncherAlmostFull),
		string(PrinterStatusDetails_PuncherAtLimit),
		string(PrinterStatusDetails_PuncherClosed),
		string(PrinterStatusDetails_PuncherConfigurationChange),
		string(PrinterStatusDetails_PuncherCoverClosed),
		string(PrinterStatusDetails_PuncherCoverOpen),
		string(PrinterStatusDetails_PuncherEmpty),
		string(PrinterStatusDetails_PuncherFull),
		string(PrinterStatusDetails_PuncherInterlockClosed),
		string(PrinterStatusDetails_PuncherInterlockOpen),
		string(PrinterStatusDetails_PuncherJam),
		string(PrinterStatusDetails_PuncherLifeAlmostOver),
		string(PrinterStatusDetails_PuncherLifeOver),
		string(PrinterStatusDetails_PuncherMemoryExhausted),
		string(PrinterStatusDetails_PuncherMissing),
		string(PrinterStatusDetails_PuncherMotorFailure),
		string(PrinterStatusDetails_PuncherNearLimit),
		string(PrinterStatusDetails_PuncherOffline),
		string(PrinterStatusDetails_PuncherOpened),
		string(PrinterStatusDetails_PuncherOverTemperature),
		string(PrinterStatusDetails_PuncherPowerSaver),
		string(PrinterStatusDetails_PuncherRecoverableFailure),
		string(PrinterStatusDetails_PuncherRecoverableStorage),
		string(PrinterStatusDetails_PuncherRemoved),
		string(PrinterStatusDetails_PuncherResourceAdded),
		string(PrinterStatusDetails_PuncherResourceRemoved),
		string(PrinterStatusDetails_PuncherThermistorFailure),
		string(PrinterStatusDetails_PuncherTimingFailure),
		string(PrinterStatusDetails_PuncherTurnedOff),
		string(PrinterStatusDetails_PuncherTurnedOn),
		string(PrinterStatusDetails_PuncherUnderTemperature),
		string(PrinterStatusDetails_PuncherUnrecoverableFailure),
		string(PrinterStatusDetails_PuncherUnrecoverableStorageError),
		string(PrinterStatusDetails_PuncherWarmingUp),
		string(PrinterStatusDetails_Resuming),
		string(PrinterStatusDetails_ScanMediaPathFailure),
		string(PrinterStatusDetails_ScanMediaPathInputEmpty),
		string(PrinterStatusDetails_ScanMediaPathInputFeedError),
		string(PrinterStatusDetails_ScanMediaPathInputJam),
		string(PrinterStatusDetails_ScanMediaPathInputRequest),
		string(PrinterStatusDetails_ScanMediaPathJam),
		string(PrinterStatusDetails_ScanMediaPathOutputFeedError),
		string(PrinterStatusDetails_ScanMediaPathOutputFull),
		string(PrinterStatusDetails_ScanMediaPathOutputJam),
		string(PrinterStatusDetails_ScanMediaPathPickRollerFailure),
		string(PrinterStatusDetails_ScanMediaPathPickRollerLifeOver),
		string(PrinterStatusDetails_ScanMediaPathPickRollerLifeWarn),
		string(PrinterStatusDetails_ScanMediaPathPickRollerMissing),
		string(PrinterStatusDetails_ScanMediaPathTrayAlmostFull),
		string(PrinterStatusDetails_ScanMediaPathTrayFull),
		string(PrinterStatusDetails_ScanMediaPathTrayMissing),
		string(PrinterStatusDetails_ScannerLightFailure),
		string(PrinterStatusDetails_ScannerLightLifeAlmostOver),
		string(PrinterStatusDetails_ScannerLightLifeOver),
		string(PrinterStatusDetails_ScannerLightMissing),
		string(PrinterStatusDetails_ScannerSensorFailure),
		string(PrinterStatusDetails_ScannerSensorLifeAlmostOver),
		string(PrinterStatusDetails_ScannerSensorLifeOver),
		string(PrinterStatusDetails_ScannerSensorMissing),
		string(PrinterStatusDetails_SeparationCutterAdded),
		string(PrinterStatusDetails_SeparationCutterAlmostEmpty),
		string(PrinterStatusDetails_SeparationCutterAlmostFull),
		string(PrinterStatusDetails_SeparationCutterAtLimit),
		string(PrinterStatusDetails_SeparationCutterClosed),
		string(PrinterStatusDetails_SeparationCutterConfigurationChange),
		string(PrinterStatusDetails_SeparationCutterCoverClosed),
		string(PrinterStatusDetails_SeparationCutterCoverOpen),
		string(PrinterStatusDetails_SeparationCutterEmpty),
		string(PrinterStatusDetails_SeparationCutterFull),
		string(PrinterStatusDetails_SeparationCutterInterlockClosed),
		string(PrinterStatusDetails_SeparationCutterInterlockOpen),
		string(PrinterStatusDetails_SeparationCutterJam),
		string(PrinterStatusDetails_SeparationCutterLifeAlmostOver),
		string(PrinterStatusDetails_SeparationCutterLifeOver),
		string(PrinterStatusDetails_SeparationCutterMemoryExhausted),
		string(PrinterStatusDetails_SeparationCutterMissing),
		string(PrinterStatusDetails_SeparationCutterMotorFailure),
		string(PrinterStatusDetails_SeparationCutterNearLimit),
		string(PrinterStatusDetails_SeparationCutterOffline),
		string(PrinterStatusDetails_SeparationCutterOpened),
		string(PrinterStatusDetails_SeparationCutterOverTemperature),
		string(PrinterStatusDetails_SeparationCutterPowerSaver),
		string(PrinterStatusDetails_SeparationCutterRecoverableFailure),
		string(PrinterStatusDetails_SeparationCutterRecoverableStorage),
		string(PrinterStatusDetails_SeparationCutterRemoved),
		string(PrinterStatusDetails_SeparationCutterResourceAdded),
		string(PrinterStatusDetails_SeparationCutterResourceRemoved),
		string(PrinterStatusDetails_SeparationCutterThermistorFailure),
		string(PrinterStatusDetails_SeparationCutterTimingFailure),
		string(PrinterStatusDetails_SeparationCutterTurnedOff),
		string(PrinterStatusDetails_SeparationCutterTurnedOn),
		string(PrinterStatusDetails_SeparationCutterUnderTemperature),
		string(PrinterStatusDetails_SeparationCutterUnrecoverableFailure),
		string(PrinterStatusDetails_SeparationCutterUnrecoverableStorageError),
		string(PrinterStatusDetails_SeparationCutterWarmingUp),
		string(PrinterStatusDetails_SheetRotatorAdded),
		string(PrinterStatusDetails_SheetRotatorAlmostEmpty),
		string(PrinterStatusDetails_SheetRotatorAlmostFull),
		string(PrinterStatusDetails_SheetRotatorAtLimit),
		string(PrinterStatusDetails_SheetRotatorClosed),
		string(PrinterStatusDetails_SheetRotatorConfigurationChange),
		string(PrinterStatusDetails_SheetRotatorCoverClosed),
		string(PrinterStatusDetails_SheetRotatorCoverOpen),
		string(PrinterStatusDetails_SheetRotatorEmpty),
		string(PrinterStatusDetails_SheetRotatorFull),
		string(PrinterStatusDetails_SheetRotatorInterlockClosed),
		string(PrinterStatusDetails_SheetRotatorInterlockOpen),
		string(PrinterStatusDetails_SheetRotatorJam),
		string(PrinterStatusDetails_SheetRotatorLifeAlmostOver),
		string(PrinterStatusDetails_SheetRotatorLifeOver),
		string(PrinterStatusDetails_SheetRotatorMemoryExhausted),
		string(PrinterStatusDetails_SheetRotatorMissing),
		string(PrinterStatusDetails_SheetRotatorMotorFailure),
		string(PrinterStatusDetails_SheetRotatorNearLimit),
		string(PrinterStatusDetails_SheetRotatorOffline),
		string(PrinterStatusDetails_SheetRotatorOpened),
		string(PrinterStatusDetails_SheetRotatorOverTemperature),
		string(PrinterStatusDetails_SheetRotatorPowerSaver),
		string(PrinterStatusDetails_SheetRotatorRecoverableFailure),
		string(PrinterStatusDetails_SheetRotatorRecoverableStorage),
		string(PrinterStatusDetails_SheetRotatorRemoved),
		string(PrinterStatusDetails_SheetRotatorResourceAdded),
		string(PrinterStatusDetails_SheetRotatorResourceRemoved),
		string(PrinterStatusDetails_SheetRotatorThermistorFailure),
		string(PrinterStatusDetails_SheetRotatorTimingFailure),
		string(PrinterStatusDetails_SheetRotatorTurnedOff),
		string(PrinterStatusDetails_SheetRotatorTurnedOn),
		string(PrinterStatusDetails_SheetRotatorUnderTemperature),
		string(PrinterStatusDetails_SheetRotatorUnrecoverableFailure),
		string(PrinterStatusDetails_SheetRotatorUnrecoverableStorageError),
		string(PrinterStatusDetails_SheetRotatorWarmingUp),
		string(PrinterStatusDetails_Shutdown),
		string(PrinterStatusDetails_SlitterAdded),
		string(PrinterStatusDetails_SlitterAlmostEmpty),
		string(PrinterStatusDetails_SlitterAlmostFull),
		string(PrinterStatusDetails_SlitterAtLimit),
		string(PrinterStatusDetails_SlitterClosed),
		string(PrinterStatusDetails_SlitterConfigurationChange),
		string(PrinterStatusDetails_SlitterCoverClosed),
		string(PrinterStatusDetails_SlitterCoverOpen),
		string(PrinterStatusDetails_SlitterEmpty),
		string(PrinterStatusDetails_SlitterFull),
		string(PrinterStatusDetails_SlitterInterlockClosed),
		string(PrinterStatusDetails_SlitterInterlockOpen),
		string(PrinterStatusDetails_SlitterJam),
		string(PrinterStatusDetails_SlitterLifeAlmostOver),
		string(PrinterStatusDetails_SlitterLifeOver),
		string(PrinterStatusDetails_SlitterMemoryExhausted),
		string(PrinterStatusDetails_SlitterMissing),
		string(PrinterStatusDetails_SlitterMotorFailure),
		string(PrinterStatusDetails_SlitterNearLimit),
		string(PrinterStatusDetails_SlitterOffline),
		string(PrinterStatusDetails_SlitterOpened),
		string(PrinterStatusDetails_SlitterOverTemperature),
		string(PrinterStatusDetails_SlitterPowerSaver),
		string(PrinterStatusDetails_SlitterRecoverableFailure),
		string(PrinterStatusDetails_SlitterRecoverableStorage),
		string(PrinterStatusDetails_SlitterRemoved),
		string(PrinterStatusDetails_SlitterResourceAdded),
		string(PrinterStatusDetails_SlitterResourceRemoved),
		string(PrinterStatusDetails_SlitterThermistorFailure),
		string(PrinterStatusDetails_SlitterTimingFailure),
		string(PrinterStatusDetails_SlitterTurnedOff),
		string(PrinterStatusDetails_SlitterTurnedOn),
		string(PrinterStatusDetails_SlitterUnderTemperature),
		string(PrinterStatusDetails_SlitterUnrecoverableFailure),
		string(PrinterStatusDetails_SlitterUnrecoverableStorageError),
		string(PrinterStatusDetails_SlitterWarmingUp),
		string(PrinterStatusDetails_SpoolAreaFull),
		string(PrinterStatusDetails_StackerAdded),
		string(PrinterStatusDetails_StackerAlmostEmpty),
		string(PrinterStatusDetails_StackerAlmostFull),
		string(PrinterStatusDetails_StackerAtLimit),
		string(PrinterStatusDetails_StackerClosed),
		string(PrinterStatusDetails_StackerConfigurationChange),
		string(PrinterStatusDetails_StackerCoverClosed),
		string(PrinterStatusDetails_StackerCoverOpen),
		string(PrinterStatusDetails_StackerEmpty),
		string(PrinterStatusDetails_StackerFull),
		string(PrinterStatusDetails_StackerInterlockClosed),
		string(PrinterStatusDetails_StackerInterlockOpen),
		string(PrinterStatusDetails_StackerJam),
		string(PrinterStatusDetails_StackerLifeAlmostOver),
		string(PrinterStatusDetails_StackerLifeOver),
		string(PrinterStatusDetails_StackerMemoryExhausted),
		string(PrinterStatusDetails_StackerMissing),
		string(PrinterStatusDetails_StackerMotorFailure),
		string(PrinterStatusDetails_StackerNearLimit),
		string(PrinterStatusDetails_StackerOffline),
		string(PrinterStatusDetails_StackerOpened),
		string(PrinterStatusDetails_StackerOverTemperature),
		string(PrinterStatusDetails_StackerPowerSaver),
		string(PrinterStatusDetails_StackerRecoverableFailure),
		string(PrinterStatusDetails_StackerRecoverableStorage),
		string(PrinterStatusDetails_StackerRemoved),
		string(PrinterStatusDetails_StackerResourceAdded),
		string(PrinterStatusDetails_StackerResourceRemoved),
		string(PrinterStatusDetails_StackerThermistorFailure),
		string(PrinterStatusDetails_StackerTimingFailure),
		string(PrinterStatusDetails_StackerTurnedOff),
		string(PrinterStatusDetails_StackerTurnedOn),
		string(PrinterStatusDetails_StackerUnderTemperature),
		string(PrinterStatusDetails_StackerUnrecoverableFailure),
		string(PrinterStatusDetails_StackerUnrecoverableStorageError),
		string(PrinterStatusDetails_StackerWarmingUp),
		string(PrinterStatusDetails_Standby),
		string(PrinterStatusDetails_StaplerAdded),
		string(PrinterStatusDetails_StaplerAlmostEmpty),
		string(PrinterStatusDetails_StaplerAlmostFull),
		string(PrinterStatusDetails_StaplerAtLimit),
		string(PrinterStatusDetails_StaplerClosed),
		string(PrinterStatusDetails_StaplerConfigurationChange),
		string(PrinterStatusDetails_StaplerCoverClosed),
		string(PrinterStatusDetails_StaplerCoverOpen),
		string(PrinterStatusDetails_StaplerEmpty),
		string(PrinterStatusDetails_StaplerFull),
		string(PrinterStatusDetails_StaplerInterlockClosed),
		string(PrinterStatusDetails_StaplerInterlockOpen),
		string(PrinterStatusDetails_StaplerJam),
		string(PrinterStatusDetails_StaplerLifeAlmostOver),
		string(PrinterStatusDetails_StaplerLifeOver),
		string(PrinterStatusDetails_StaplerMemoryExhausted),
		string(PrinterStatusDetails_StaplerMissing),
		string(PrinterStatusDetails_StaplerMotorFailure),
		string(PrinterStatusDetails_StaplerNearLimit),
		string(PrinterStatusDetails_StaplerOffline),
		string(PrinterStatusDetails_StaplerOpened),
		string(PrinterStatusDetails_StaplerOverTemperature),
		string(PrinterStatusDetails_StaplerPowerSaver),
		string(PrinterStatusDetails_StaplerRecoverableFailure),
		string(PrinterStatusDetails_StaplerRecoverableStorage),
		string(PrinterStatusDetails_StaplerRemoved),
		string(PrinterStatusDetails_StaplerResourceAdded),
		string(PrinterStatusDetails_StaplerResourceRemoved),
		string(PrinterStatusDetails_StaplerThermistorFailure),
		string(PrinterStatusDetails_StaplerTimingFailure),
		string(PrinterStatusDetails_StaplerTurnedOff),
		string(PrinterStatusDetails_StaplerTurnedOn),
		string(PrinterStatusDetails_StaplerUnderTemperature),
		string(PrinterStatusDetails_StaplerUnrecoverableFailure),
		string(PrinterStatusDetails_StaplerUnrecoverableStorageError),
		string(PrinterStatusDetails_StaplerWarmingUp),
		string(PrinterStatusDetails_StitcherAdded),
		string(PrinterStatusDetails_StitcherAlmostEmpty),
		string(PrinterStatusDetails_StitcherAlmostFull),
		string(PrinterStatusDetails_StitcherAtLimit),
		string(PrinterStatusDetails_StitcherClosed),
		string(PrinterStatusDetails_StitcherConfigurationChange),
		string(PrinterStatusDetails_StitcherCoverClosed),
		string(PrinterStatusDetails_StitcherCoverOpen),
		string(PrinterStatusDetails_StitcherEmpty),
		string(PrinterStatusDetails_StitcherFull),
		string(PrinterStatusDetails_StitcherInterlockClosed),
		string(PrinterStatusDetails_StitcherInterlockOpen),
		string(PrinterStatusDetails_StitcherJam),
		string(PrinterStatusDetails_StitcherLifeAlmostOver),
		string(PrinterStatusDetails_StitcherLifeOver),
		string(PrinterStatusDetails_StitcherMemoryExhausted),
		string(PrinterStatusDetails_StitcherMissing),
		string(PrinterStatusDetails_StitcherMotorFailure),
		string(PrinterStatusDetails_StitcherNearLimit),
		string(PrinterStatusDetails_StitcherOffline),
		string(PrinterStatusDetails_StitcherOpened),
		string(PrinterStatusDetails_StitcherOverTemperature),
		string(PrinterStatusDetails_StitcherPowerSaver),
		string(PrinterStatusDetails_StitcherRecoverableFailure),
		string(PrinterStatusDetails_StitcherRecoverableStorage),
		string(PrinterStatusDetails_StitcherRemoved),
		string(PrinterStatusDetails_StitcherResourceAdded),
		string(PrinterStatusDetails_StitcherResourceRemoved),
		string(PrinterStatusDetails_StitcherThermistorFailure),
		string(PrinterStatusDetails_StitcherTimingFailure),
		string(PrinterStatusDetails_StitcherTurnedOff),
		string(PrinterStatusDetails_StitcherTurnedOn),
		string(PrinterStatusDetails_StitcherUnderTemperature),
		string(PrinterStatusDetails_StitcherUnrecoverableFailure),
		string(PrinterStatusDetails_StitcherUnrecoverableStorageError),
		string(PrinterStatusDetails_StitcherWarmingUp),
		string(PrinterStatusDetails_StoppedPartially),
		string(PrinterStatusDetails_Stopping),
		string(PrinterStatusDetails_SubunitAdded),
		string(PrinterStatusDetails_SubunitAlmostEmpty),
		string(PrinterStatusDetails_SubunitAlmostFull),
		string(PrinterStatusDetails_SubunitAtLimit),
		string(PrinterStatusDetails_SubunitClosed),
		string(PrinterStatusDetails_SubunitCoolingDown),
		string(PrinterStatusDetails_SubunitEmpty),
		string(PrinterStatusDetails_SubunitFull),
		string(PrinterStatusDetails_SubunitLifeAlmostOver),
		string(PrinterStatusDetails_SubunitLifeOver),
		string(PrinterStatusDetails_SubunitMemoryExhausted),
		string(PrinterStatusDetails_SubunitMissing),
		string(PrinterStatusDetails_SubunitMotorFailure),
		string(PrinterStatusDetails_SubunitNearLimit),
		string(PrinterStatusDetails_SubunitOffline),
		string(PrinterStatusDetails_SubunitOpened),
		string(PrinterStatusDetails_SubunitOverTemperature),
		string(PrinterStatusDetails_SubunitPowerSaver),
		string(PrinterStatusDetails_SubunitRecoverableFailure),
		string(PrinterStatusDetails_SubunitRecoverableStorage),
		string(PrinterStatusDetails_SubunitRemoved),
		string(PrinterStatusDetails_SubunitResourceAdded),
		string(PrinterStatusDetails_SubunitResourceRemoved),
		string(PrinterStatusDetails_SubunitThermistorFailure),
		string(PrinterStatusDetails_SubunitTimingFailure),
		string(PrinterStatusDetails_SubunitTurnedOff),
		string(PrinterStatusDetails_SubunitTurnedOn),
		string(PrinterStatusDetails_SubunitUnderTemperature),
		string(PrinterStatusDetails_SubunitUnrecoverableFailure),
		string(PrinterStatusDetails_SubunitUnrecoverableStorage),
		string(PrinterStatusDetails_SubunitWarmingUp),
		string(PrinterStatusDetails_Suspend),
		string(PrinterStatusDetails_Testing),
		string(PrinterStatusDetails_TimedOut),
		string(PrinterStatusDetails_TonerEmpty),
		string(PrinterStatusDetails_TonerLow),
		string(PrinterStatusDetails_TrimmerAdded),
		string(PrinterStatusDetails_TrimmerAlmostEmpty),
		string(PrinterStatusDetails_TrimmerAlmostFull),
		string(PrinterStatusDetails_TrimmerAtLimit),
		string(PrinterStatusDetails_TrimmerClosed),
		string(PrinterStatusDetails_TrimmerConfigurationChange),
		string(PrinterStatusDetails_TrimmerCoverClosed),
		string(PrinterStatusDetails_TrimmerCoverOpen),
		string(PrinterStatusDetails_TrimmerEmpty),
		string(PrinterStatusDetails_TrimmerFull),
		string(PrinterStatusDetails_TrimmerInterlockClosed),
		string(PrinterStatusDetails_TrimmerInterlockOpen),
		string(PrinterStatusDetails_TrimmerJam),
		string(PrinterStatusDetails_TrimmerLifeAlmostOver),
		string(PrinterStatusDetails_TrimmerLifeOver),
		string(PrinterStatusDetails_TrimmerMemoryExhausted),
		string(PrinterStatusDetails_TrimmerMissing),
		string(PrinterStatusDetails_TrimmerMotorFailure),
		string(PrinterStatusDetails_TrimmerNearLimit),
		string(PrinterStatusDetails_TrimmerOffline),
		string(PrinterStatusDetails_TrimmerOpened),
		string(PrinterStatusDetails_TrimmerOverTemperature),
		string(PrinterStatusDetails_TrimmerPowerSaver),
		string(PrinterStatusDetails_TrimmerRecoverableFailure),
		string(PrinterStatusDetails_TrimmerRecoverableStorage),
		string(PrinterStatusDetails_TrimmerRemoved),
		string(PrinterStatusDetails_TrimmerResourceAdded),
		string(PrinterStatusDetails_TrimmerResourceRemoved),
		string(PrinterStatusDetails_TrimmerThermistorFailure),
		string(PrinterStatusDetails_TrimmerTimingFailure),
		string(PrinterStatusDetails_TrimmerTurnedOff),
		string(PrinterStatusDetails_TrimmerTurnedOn),
		string(PrinterStatusDetails_TrimmerUnderTemperature),
		string(PrinterStatusDetails_TrimmerUnrecoverableFailure),
		string(PrinterStatusDetails_TrimmerUnrecoverableStorageError),
		string(PrinterStatusDetails_TrimmerWarmingUp),
		string(PrinterStatusDetails_Unknown),
		string(PrinterStatusDetails_WrapperAdded),
		string(PrinterStatusDetails_WrapperAlmostEmpty),
		string(PrinterStatusDetails_WrapperAlmostFull),
		string(PrinterStatusDetails_WrapperAtLimit),
		string(PrinterStatusDetails_WrapperClosed),
		string(PrinterStatusDetails_WrapperConfigurationChange),
		string(PrinterStatusDetails_WrapperCoverClosed),
		string(PrinterStatusDetails_WrapperCoverOpen),
		string(PrinterStatusDetails_WrapperEmpty),
		string(PrinterStatusDetails_WrapperFull),
		string(PrinterStatusDetails_WrapperInterlockClosed),
		string(PrinterStatusDetails_WrapperInterlockOpen),
		string(PrinterStatusDetails_WrapperJam),
		string(PrinterStatusDetails_WrapperLifeAlmostOver),
		string(PrinterStatusDetails_WrapperLifeOver),
		string(PrinterStatusDetails_WrapperMemoryExhausted),
		string(PrinterStatusDetails_WrapperMissing),
		string(PrinterStatusDetails_WrapperMotorFailure),
		string(PrinterStatusDetails_WrapperNearLimit),
		string(PrinterStatusDetails_WrapperOffline),
		string(PrinterStatusDetails_WrapperOpened),
		string(PrinterStatusDetails_WrapperOverTemperature),
		string(PrinterStatusDetails_WrapperPowerSaver),
		string(PrinterStatusDetails_WrapperRecoverableFailure),
		string(PrinterStatusDetails_WrapperRecoverableStorage),
		string(PrinterStatusDetails_WrapperRemoved),
		string(PrinterStatusDetails_WrapperResourceAdded),
		string(PrinterStatusDetails_WrapperResourceRemoved),
		string(PrinterStatusDetails_WrapperThermistorFailure),
		string(PrinterStatusDetails_WrapperTimingFailure),
		string(PrinterStatusDetails_WrapperTurnedOff),
		string(PrinterStatusDetails_WrapperTurnedOn),
		string(PrinterStatusDetails_WrapperUnderTemperature),
		string(PrinterStatusDetails_WrapperUnrecoverableFailure),
		string(PrinterStatusDetails_WrapperUnrecoverableStorageError),
		string(PrinterStatusDetails_WrapperWarmingUp),
	}
}

func (s *PrinterStatusDetails) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterStatusDetails(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterStatusDetails(input string) (*PrinterStatusDetails, error) {
	vals := map[string]PrinterStatusDetails{
		"alertremovalofbinarychangeentry":           PrinterStatusDetails_AlertRemovalOfBinaryChangeEntry,
		"banderadded":                               PrinterStatusDetails_BanderAdded,
		"banderalmostempty":                         PrinterStatusDetails_BanderAlmostEmpty,
		"banderalmostfull":                          PrinterStatusDetails_BanderAlmostFull,
		"banderatlimit":                             PrinterStatusDetails_BanderAtLimit,
		"banderclosed":                              PrinterStatusDetails_BanderClosed,
		"banderconfigurationchange":                 PrinterStatusDetails_BanderConfigurationChange,
		"bandercoverclosed":                         PrinterStatusDetails_BanderCoverClosed,
		"bandercoveropen":                           PrinterStatusDetails_BanderCoverOpen,
		"banderempty":                               PrinterStatusDetails_BanderEmpty,
		"banderfull":                                PrinterStatusDetails_BanderFull,
		"banderinterlockclosed":                     PrinterStatusDetails_BanderInterlockClosed,
		"banderinterlockopen":                       PrinterStatusDetails_BanderInterlockOpen,
		"banderjam":                                 PrinterStatusDetails_BanderJam,
		"banderlifealmostover":                      PrinterStatusDetails_BanderLifeAlmostOver,
		"banderlifeover":                            PrinterStatusDetails_BanderLifeOver,
		"bandermemoryexhausted":                     PrinterStatusDetails_BanderMemoryExhausted,
		"bandermissing":                             PrinterStatusDetails_BanderMissing,
		"bandermotorfailure":                        PrinterStatusDetails_BanderMotorFailure,
		"bandernearlimit":                           PrinterStatusDetails_BanderNearLimit,
		"banderoffline":                             PrinterStatusDetails_BanderOffline,
		"banderopened":                              PrinterStatusDetails_BanderOpened,
		"banderovertemperature":                     PrinterStatusDetails_BanderOverTemperature,
		"banderpowersaver":                          PrinterStatusDetails_BanderPowerSaver,
		"banderrecoverablefailure":                  PrinterStatusDetails_BanderRecoverableFailure,
		"banderrecoverablestorage":                  PrinterStatusDetails_BanderRecoverableStorage,
		"banderremoved":                             PrinterStatusDetails_BanderRemoved,
		"banderresourceadded":                       PrinterStatusDetails_BanderResourceAdded,
		"banderresourceremoved":                     PrinterStatusDetails_BanderResourceRemoved,
		"banderthermistorfailure":                   PrinterStatusDetails_BanderThermistorFailure,
		"bandertimingfailure":                       PrinterStatusDetails_BanderTimingFailure,
		"banderturnedoff":                           PrinterStatusDetails_BanderTurnedOff,
		"banderturnedon":                            PrinterStatusDetails_BanderTurnedOn,
		"banderundertemperature":                    PrinterStatusDetails_BanderUnderTemperature,
		"banderunrecoverablefailure":                PrinterStatusDetails_BanderUnrecoverableFailure,
		"banderunrecoverablestorageerror":           PrinterStatusDetails_BanderUnrecoverableStorageError,
		"banderwarmingup":                           PrinterStatusDetails_BanderWarmingUp,
		"binderadded":                               PrinterStatusDetails_BinderAdded,
		"binderalmostempty":                         PrinterStatusDetails_BinderAlmostEmpty,
		"binderalmostfull":                          PrinterStatusDetails_BinderAlmostFull,
		"binderatlimit":                             PrinterStatusDetails_BinderAtLimit,
		"binderclosed":                              PrinterStatusDetails_BinderClosed,
		"binderconfigurationchange":                 PrinterStatusDetails_BinderConfigurationChange,
		"bindercoverclosed":                         PrinterStatusDetails_BinderCoverClosed,
		"bindercoveropen":                           PrinterStatusDetails_BinderCoverOpen,
		"binderempty":                               PrinterStatusDetails_BinderEmpty,
		"binderfull":                                PrinterStatusDetails_BinderFull,
		"binderinterlockclosed":                     PrinterStatusDetails_BinderInterlockClosed,
		"binderinterlockopen":                       PrinterStatusDetails_BinderInterlockOpen,
		"binderjam":                                 PrinterStatusDetails_BinderJam,
		"binderlifealmostover":                      PrinterStatusDetails_BinderLifeAlmostOver,
		"binderlifeover":                            PrinterStatusDetails_BinderLifeOver,
		"bindermemoryexhausted":                     PrinterStatusDetails_BinderMemoryExhausted,
		"bindermissing":                             PrinterStatusDetails_BinderMissing,
		"bindermotorfailure":                        PrinterStatusDetails_BinderMotorFailure,
		"bindernearlimit":                           PrinterStatusDetails_BinderNearLimit,
		"binderoffline":                             PrinterStatusDetails_BinderOffline,
		"binderopened":                              PrinterStatusDetails_BinderOpened,
		"binderovertemperature":                     PrinterStatusDetails_BinderOverTemperature,
		"binderpowersaver":                          PrinterStatusDetails_BinderPowerSaver,
		"binderrecoverablefailure":                  PrinterStatusDetails_BinderRecoverableFailure,
		"binderrecoverablestorage":                  PrinterStatusDetails_BinderRecoverableStorage,
		"binderremoved":                             PrinterStatusDetails_BinderRemoved,
		"binderresourceadded":                       PrinterStatusDetails_BinderResourceAdded,
		"binderresourceremoved":                     PrinterStatusDetails_BinderResourceRemoved,
		"binderthermistorfailure":                   PrinterStatusDetails_BinderThermistorFailure,
		"bindertimingfailure":                       PrinterStatusDetails_BinderTimingFailure,
		"binderturnedoff":                           PrinterStatusDetails_BinderTurnedOff,
		"binderturnedon":                            PrinterStatusDetails_BinderTurnedOn,
		"binderundertemperature":                    PrinterStatusDetails_BinderUnderTemperature,
		"binderunrecoverablefailure":                PrinterStatusDetails_BinderUnrecoverableFailure,
		"binderunrecoverablestorageerror":           PrinterStatusDetails_BinderUnrecoverableStorageError,
		"binderwarmingup":                           PrinterStatusDetails_BinderWarmingUp,
		"camerafailure":                             PrinterStatusDetails_CameraFailure,
		"chambercooling":                            PrinterStatusDetails_ChamberCooling,
		"chamberfailure":                            PrinterStatusDetails_ChamberFailure,
		"chamberheating":                            PrinterStatusDetails_ChamberHeating,
		"chambertemperaturehigh":                    PrinterStatusDetails_ChamberTemperatureHigh,
		"chambertemperaturelow":                     PrinterStatusDetails_ChamberTemperatureLow,
		"cleanerlifealmostover":                     PrinterStatusDetails_CleanerLifeAlmostOver,
		"cleanerlifeover":                           PrinterStatusDetails_CleanerLifeOver,
		"configurationchange":                       PrinterStatusDetails_ConfigurationChange,
		"connectingtodevice":                        PrinterStatusDetails_ConnectingToDevice,
		"coveropen":                                 PrinterStatusDetails_CoverOpen,
		"deactivated":                               PrinterStatusDetails_Deactivated,
		"deleted":                                   PrinterStatusDetails_Deleted,
		"developerempty":                            PrinterStatusDetails_DeveloperEmpty,
		"developerlow":                              PrinterStatusDetails_DeveloperLow,
		"diecutteradded":                            PrinterStatusDetails_DieCutterAdded,
		"diecutteralmostempty":                      PrinterStatusDetails_DieCutterAlmostEmpty,
		"diecutteralmostfull":                       PrinterStatusDetails_DieCutterAlmostFull,
		"diecutteratlimit":                          PrinterStatusDetails_DieCutterAtLimit,
		"diecutterclosed":                           PrinterStatusDetails_DieCutterClosed,
		"diecutterconfigurationchange":              PrinterStatusDetails_DieCutterConfigurationChange,
		"diecuttercoverclosed":                      PrinterStatusDetails_DieCutterCoverClosed,
		"diecuttercoveropen":                        PrinterStatusDetails_DieCutterCoverOpen,
		"diecutterempty":                            PrinterStatusDetails_DieCutterEmpty,
		"diecutterfull":                             PrinterStatusDetails_DieCutterFull,
		"diecutterinterlockclosed":                  PrinterStatusDetails_DieCutterInterlockClosed,
		"diecutterinterlockopen":                    PrinterStatusDetails_DieCutterInterlockOpen,
		"diecutterjam":                              PrinterStatusDetails_DieCutterJam,
		"diecutterlifealmostover":                   PrinterStatusDetails_DieCutterLifeAlmostOver,
		"diecutterlifeover":                         PrinterStatusDetails_DieCutterLifeOver,
		"diecuttermemoryexhausted":                  PrinterStatusDetails_DieCutterMemoryExhausted,
		"diecuttermissing":                          PrinterStatusDetails_DieCutterMissing,
		"diecuttermotorfailure":                     PrinterStatusDetails_DieCutterMotorFailure,
		"diecutternearlimit":                        PrinterStatusDetails_DieCutterNearLimit,
		"diecutteroffline":                          PrinterStatusDetails_DieCutterOffline,
		"diecutteropened":                           PrinterStatusDetails_DieCutterOpened,
		"diecutterovertemperature":                  PrinterStatusDetails_DieCutterOverTemperature,
		"diecutterpowersaver":                       PrinterStatusDetails_DieCutterPowerSaver,
		"diecutterrecoverablefailure":               PrinterStatusDetails_DieCutterRecoverableFailure,
		"diecutterrecoverablestorage":               PrinterStatusDetails_DieCutterRecoverableStorage,
		"diecutterremoved":                          PrinterStatusDetails_DieCutterRemoved,
		"diecutterresourceadded":                    PrinterStatusDetails_DieCutterResourceAdded,
		"diecutterresourceremoved":                  PrinterStatusDetails_DieCutterResourceRemoved,
		"diecutterthermistorfailure":                PrinterStatusDetails_DieCutterThermistorFailure,
		"diecuttertimingfailure":                    PrinterStatusDetails_DieCutterTimingFailure,
		"diecutterturnedoff":                        PrinterStatusDetails_DieCutterTurnedOff,
		"diecutterturnedon":                         PrinterStatusDetails_DieCutterTurnedOn,
		"diecutterundertemperature":                 PrinterStatusDetails_DieCutterUnderTemperature,
		"diecutterunrecoverablefailure":             PrinterStatusDetails_DieCutterUnrecoverableFailure,
		"diecutterunrecoverablestorageerror":        PrinterStatusDetails_DieCutterUnrecoverableStorageError,
		"diecutterwarmingup":                        PrinterStatusDetails_DieCutterWarmingUp,
		"dooropen":                                  PrinterStatusDetails_DoorOpen,
		"extrudercooling":                           PrinterStatusDetails_ExtruderCooling,
		"extruderfailure":                           PrinterStatusDetails_ExtruderFailure,
		"extruderheating":                           PrinterStatusDetails_ExtruderHeating,
		"extruderjam":                               PrinterStatusDetails_ExtruderJam,
		"extrudertemperaturehigh":                   PrinterStatusDetails_ExtruderTemperatureHigh,
		"extrudertemperaturelow":                    PrinterStatusDetails_ExtruderTemperatureLow,
		"fanfailure":                                PrinterStatusDetails_FanFailure,
		"faxmodemlifealmostover":                    PrinterStatusDetails_FaxModemLifeAlmostOver,
		"faxmodemlifeover":                          PrinterStatusDetails_FaxModemLifeOver,
		"faxmodemmissing":                           PrinterStatusDetails_FaxModemMissing,
		"faxmodemturnedoff":                         PrinterStatusDetails_FaxModemTurnedOff,
		"faxmodemturnedon":                          PrinterStatusDetails_FaxModemTurnedOn,
		"folderadded":                               PrinterStatusDetails_FolderAdded,
		"folderalmostempty":                         PrinterStatusDetails_FolderAlmostEmpty,
		"folderalmostfull":                          PrinterStatusDetails_FolderAlmostFull,
		"folderatlimit":                             PrinterStatusDetails_FolderAtLimit,
		"folderclosed":                              PrinterStatusDetails_FolderClosed,
		"folderconfigurationchange":                 PrinterStatusDetails_FolderConfigurationChange,
		"foldercoverclosed":                         PrinterStatusDetails_FolderCoverClosed,
		"foldercoveropen":                           PrinterStatusDetails_FolderCoverOpen,
		"folderempty":                               PrinterStatusDetails_FolderEmpty,
		"folderfull":                                PrinterStatusDetails_FolderFull,
		"folderinterlockclosed":                     PrinterStatusDetails_FolderInterlockClosed,
		"folderinterlockopen":                       PrinterStatusDetails_FolderInterlockOpen,
		"folderjam":                                 PrinterStatusDetails_FolderJam,
		"folderlifealmostover":                      PrinterStatusDetails_FolderLifeAlmostOver,
		"folderlifeover":                            PrinterStatusDetails_FolderLifeOver,
		"foldermemoryexhausted":                     PrinterStatusDetails_FolderMemoryExhausted,
		"foldermissing":                             PrinterStatusDetails_FolderMissing,
		"foldermotorfailure":                        PrinterStatusDetails_FolderMotorFailure,
		"foldernearlimit":                           PrinterStatusDetails_FolderNearLimit,
		"folderoffline":                             PrinterStatusDetails_FolderOffline,
		"folderopened":                              PrinterStatusDetails_FolderOpened,
		"folderovertemperature":                     PrinterStatusDetails_FolderOverTemperature,
		"folderpowersaver":                          PrinterStatusDetails_FolderPowerSaver,
		"folderrecoverablefailure":                  PrinterStatusDetails_FolderRecoverableFailure,
		"folderrecoverablestorage":                  PrinterStatusDetails_FolderRecoverableStorage,
		"folderremoved":                             PrinterStatusDetails_FolderRemoved,
		"folderresourceadded":                       PrinterStatusDetails_FolderResourceAdded,
		"folderresourceremoved":                     PrinterStatusDetails_FolderResourceRemoved,
		"folderthermistorfailure":                   PrinterStatusDetails_FolderThermistorFailure,
		"foldertimingfailure":                       PrinterStatusDetails_FolderTimingFailure,
		"folderturnedoff":                           PrinterStatusDetails_FolderTurnedOff,
		"folderturnedon":                            PrinterStatusDetails_FolderTurnedOn,
		"folderundertemperature":                    PrinterStatusDetails_FolderUnderTemperature,
		"folderunrecoverablefailure":                PrinterStatusDetails_FolderUnrecoverableFailure,
		"folderunrecoverablestorageerror":           PrinterStatusDetails_FolderUnrecoverableStorageError,
		"folderwarmingup":                           PrinterStatusDetails_FolderWarmingUp,
		"fuserovertemp":                             PrinterStatusDetails_FuserOverTemp,
		"fuserundertemp":                            PrinterStatusDetails_FuserUnderTemp,
		"hibernate":                                 PrinterStatusDetails_Hibernate,
		"holdnewjobs":                               PrinterStatusDetails_HoldNewJobs,
		"identifyprinterrequested":                  PrinterStatusDetails_IdentifyPrinterRequested,
		"imprinteradded":                            PrinterStatusDetails_ImprinterAdded,
		"imprinteralmostempty":                      PrinterStatusDetails_ImprinterAlmostEmpty,
		"imprinteralmostfull":                       PrinterStatusDetails_ImprinterAlmostFull,
		"imprinteratlimit":                          PrinterStatusDetails_ImprinterAtLimit,
		"imprinterclosed":                           PrinterStatusDetails_ImprinterClosed,
		"imprinterconfigurationchange":              PrinterStatusDetails_ImprinterConfigurationChange,
		"imprintercoverclosed":                      PrinterStatusDetails_ImprinterCoverClosed,
		"imprintercoveropen":                        PrinterStatusDetails_ImprinterCoverOpen,
		"imprinterempty":                            PrinterStatusDetails_ImprinterEmpty,
		"imprinterfull":                             PrinterStatusDetails_ImprinterFull,
		"imprinterinterlockclosed":                  PrinterStatusDetails_ImprinterInterlockClosed,
		"imprinterinterlockopen":                    PrinterStatusDetails_ImprinterInterlockOpen,
		"imprinterjam":                              PrinterStatusDetails_ImprinterJam,
		"imprinterlifealmostover":                   PrinterStatusDetails_ImprinterLifeAlmostOver,
		"imprinterlifeover":                         PrinterStatusDetails_ImprinterLifeOver,
		"imprintermemoryexhausted":                  PrinterStatusDetails_ImprinterMemoryExhausted,
		"imprintermissing":                          PrinterStatusDetails_ImprinterMissing,
		"imprintermotorfailure":                     PrinterStatusDetails_ImprinterMotorFailure,
		"imprinternearlimit":                        PrinterStatusDetails_ImprinterNearLimit,
		"imprinteroffline":                          PrinterStatusDetails_ImprinterOffline,
		"imprinteropened":                           PrinterStatusDetails_ImprinterOpened,
		"imprinterovertemperature":                  PrinterStatusDetails_ImprinterOverTemperature,
		"imprinterpowersaver":                       PrinterStatusDetails_ImprinterPowerSaver,
		"imprinterrecoverablefailure":               PrinterStatusDetails_ImprinterRecoverableFailure,
		"imprinterrecoverablestorage":               PrinterStatusDetails_ImprinterRecoverableStorage,
		"imprinterremoved":                          PrinterStatusDetails_ImprinterRemoved,
		"imprinterresourceadded":                    PrinterStatusDetails_ImprinterResourceAdded,
		"imprinterresourceremoved":                  PrinterStatusDetails_ImprinterResourceRemoved,
		"imprinterthermistorfailure":                PrinterStatusDetails_ImprinterThermistorFailure,
		"imprintertimingfailure":                    PrinterStatusDetails_ImprinterTimingFailure,
		"imprinterturnedoff":                        PrinterStatusDetails_ImprinterTurnedOff,
		"imprinterturnedon":                         PrinterStatusDetails_ImprinterTurnedOn,
		"imprinterundertemperature":                 PrinterStatusDetails_ImprinterUnderTemperature,
		"imprinterunrecoverablefailure":             PrinterStatusDetails_ImprinterUnrecoverableFailure,
		"imprinterunrecoverablestorageerror":        PrinterStatusDetails_ImprinterUnrecoverableStorageError,
		"imprinterwarmingup":                        PrinterStatusDetails_ImprinterWarmingUp,
		"inputcannotfeedsizeselected":               PrinterStatusDetails_InputCannotFeedSizeSelected,
		"inputmanualinputrequest":                   PrinterStatusDetails_InputManualInputRequest,
		"inputmediacolorchange":                     PrinterStatusDetails_InputMediaColorChange,
		"inputmediaformpartschange":                 PrinterStatusDetails_InputMediaFormPartsChange,
		"inputmediasizechange":                      PrinterStatusDetails_InputMediaSizeChange,
		"inputmediatrayfailure":                     PrinterStatusDetails_InputMediaTrayFailure,
		"inputmediatrayfeederror":                   PrinterStatusDetails_InputMediaTrayFeedError,
		"inputmediatrayjam":                         PrinterStatusDetails_InputMediaTrayJam,
		"inputmediatypechange":                      PrinterStatusDetails_InputMediaTypeChange,
		"inputmediaweightchange":                    PrinterStatusDetails_InputMediaWeightChange,
		"inputpickrollerfailure":                    PrinterStatusDetails_InputPickRollerFailure,
		"inputpickrollerlifeover":                   PrinterStatusDetails_InputPickRollerLifeOver,
		"inputpickrollerlifewarn":                   PrinterStatusDetails_InputPickRollerLifeWarn,
		"inputpickrollermissing":                    PrinterStatusDetails_InputPickRollerMissing,
		"inputtrayelevationfailure":                 PrinterStatusDetails_InputTrayElevationFailure,
		"inputtraymissing":                          PrinterStatusDetails_InputTrayMissing,
		"inputtraypositionfailure":                  PrinterStatusDetails_InputTrayPositionFailure,
		"inserteradded":                             PrinterStatusDetails_InserterAdded,
		"inserteralmostempty":                       PrinterStatusDetails_InserterAlmostEmpty,
		"inserteralmostfull":                        PrinterStatusDetails_InserterAlmostFull,
		"inserteratlimit":                           PrinterStatusDetails_InserterAtLimit,
		"inserterclosed":                            PrinterStatusDetails_InserterClosed,
		"inserterconfigurationchange":               PrinterStatusDetails_InserterConfigurationChange,
		"insertercoverclosed":                       PrinterStatusDetails_InserterCoverClosed,
		"insertercoveropen":                         PrinterStatusDetails_InserterCoverOpen,
		"inserterempty":                             PrinterStatusDetails_InserterEmpty,
		"inserterfull":                              PrinterStatusDetails_InserterFull,
		"inserterinterlockclosed":                   PrinterStatusDetails_InserterInterlockClosed,
		"inserterinterlockopen":                     PrinterStatusDetails_InserterInterlockOpen,
		"inserterjam":                               PrinterStatusDetails_InserterJam,
		"inserterlifealmostover":                    PrinterStatusDetails_InserterLifeAlmostOver,
		"inserterlifeover":                          PrinterStatusDetails_InserterLifeOver,
		"insertermemoryexhausted":                   PrinterStatusDetails_InserterMemoryExhausted,
		"insertermissing":                           PrinterStatusDetails_InserterMissing,
		"insertermotorfailure":                      PrinterStatusDetails_InserterMotorFailure,
		"inserternearlimit":                         PrinterStatusDetails_InserterNearLimit,
		"inserteroffline":                           PrinterStatusDetails_InserterOffline,
		"inserteropened":                            PrinterStatusDetails_InserterOpened,
		"inserterovertemperature":                   PrinterStatusDetails_InserterOverTemperature,
		"inserterpowersaver":                        PrinterStatusDetails_InserterPowerSaver,
		"inserterrecoverablefailure":                PrinterStatusDetails_InserterRecoverableFailure,
		"inserterrecoverablestorage":                PrinterStatusDetails_InserterRecoverableStorage,
		"inserterremoved":                           PrinterStatusDetails_InserterRemoved,
		"inserterresourceadded":                     PrinterStatusDetails_InserterResourceAdded,
		"inserterresourceremoved":                   PrinterStatusDetails_InserterResourceRemoved,
		"inserterthermistorfailure":                 PrinterStatusDetails_InserterThermistorFailure,
		"insertertimingfailure":                     PrinterStatusDetails_InserterTimingFailure,
		"inserterturnedoff":                         PrinterStatusDetails_InserterTurnedOff,
		"inserterturnedon":                          PrinterStatusDetails_InserterTurnedOn,
		"inserterundertemperature":                  PrinterStatusDetails_InserterUnderTemperature,
		"inserterunrecoverablefailure":              PrinterStatusDetails_InserterUnrecoverableFailure,
		"inserterunrecoverablestorageerror":         PrinterStatusDetails_InserterUnrecoverableStorageError,
		"inserterwarmingup":                         PrinterStatusDetails_InserterWarmingUp,
		"interlockclosed":                           PrinterStatusDetails_InterlockClosed,
		"interlockopen":                             PrinterStatusDetails_InterlockOpen,
		"interpretercartridgeadded":                 PrinterStatusDetails_InterpreterCartridgeAdded,
		"interpretercartridgedeleted":               PrinterStatusDetails_InterpreterCartridgeDeleted,
		"interpretercomplexpageencountered":         PrinterStatusDetails_InterpreterComplexPageEncountered,
		"interpretermemorydecrease":                 PrinterStatusDetails_InterpreterMemoryDecrease,
		"interpretermemoryincrease":                 PrinterStatusDetails_InterpreterMemoryIncrease,
		"interpreterresourceadded":                  PrinterStatusDetails_InterpreterResourceAdded,
		"interpreterresourcedeleted":                PrinterStatusDetails_InterpreterResourceDeleted,
		"interpreterresourceunavailable":            PrinterStatusDetails_InterpreterResourceUnavailable,
		"lampateol":                                 PrinterStatusDetails_LampAtEol,
		"lampfailure":                               PrinterStatusDetails_LampFailure,
		"lampneareol":                               PrinterStatusDetails_LampNearEol,
		"laserateol":                                PrinterStatusDetails_LaserAtEol,
		"laserfailure":                              PrinterStatusDetails_LaserFailure,
		"laserneareol":                              PrinterStatusDetails_LaserNearEol,
		"makeenvelopeadded":                         PrinterStatusDetails_MakeEnvelopeAdded,
		"makeenvelopealmostempty":                   PrinterStatusDetails_MakeEnvelopeAlmostEmpty,
		"makeenvelopealmostfull":                    PrinterStatusDetails_MakeEnvelopeAlmostFull,
		"makeenvelopeatlimit":                       PrinterStatusDetails_MakeEnvelopeAtLimit,
		"makeenvelopeclosed":                        PrinterStatusDetails_MakeEnvelopeClosed,
		"makeenvelopeconfigurationchange":           PrinterStatusDetails_MakeEnvelopeConfigurationChange,
		"makeenvelopecoverclosed":                   PrinterStatusDetails_MakeEnvelopeCoverClosed,
		"makeenvelopecoveropen":                     PrinterStatusDetails_MakeEnvelopeCoverOpen,
		"makeenvelopeempty":                         PrinterStatusDetails_MakeEnvelopeEmpty,
		"makeenvelopefull":                          PrinterStatusDetails_MakeEnvelopeFull,
		"makeenvelopeinterlockclosed":               PrinterStatusDetails_MakeEnvelopeInterlockClosed,
		"makeenvelopeinterlockopen":                 PrinterStatusDetails_MakeEnvelopeInterlockOpen,
		"makeenvelopejam":                           PrinterStatusDetails_MakeEnvelopeJam,
		"makeenvelopelifealmostover":                PrinterStatusDetails_MakeEnvelopeLifeAlmostOver,
		"makeenvelopelifeover":                      PrinterStatusDetails_MakeEnvelopeLifeOver,
		"makeenvelopememoryexhausted":               PrinterStatusDetails_MakeEnvelopeMemoryExhausted,
		"makeenvelopemissing":                       PrinterStatusDetails_MakeEnvelopeMissing,
		"makeenvelopemotorfailure":                  PrinterStatusDetails_MakeEnvelopeMotorFailure,
		"makeenvelopenearlimit":                     PrinterStatusDetails_MakeEnvelopeNearLimit,
		"makeenvelopeoffline":                       PrinterStatusDetails_MakeEnvelopeOffline,
		"makeenvelopeopened":                        PrinterStatusDetails_MakeEnvelopeOpened,
		"makeenvelopeovertemperature":               PrinterStatusDetails_MakeEnvelopeOverTemperature,
		"makeenvelopepowersaver":                    PrinterStatusDetails_MakeEnvelopePowerSaver,
		"makeenveloperecoverablefailure":            PrinterStatusDetails_MakeEnvelopeRecoverableFailure,
		"makeenveloperecoverablestorage":            PrinterStatusDetails_MakeEnvelopeRecoverableStorage,
		"makeenveloperemoved":                       PrinterStatusDetails_MakeEnvelopeRemoved,
		"makeenveloperesourceadded":                 PrinterStatusDetails_MakeEnvelopeResourceAdded,
		"makeenveloperesourceremoved":               PrinterStatusDetails_MakeEnvelopeResourceRemoved,
		"makeenvelopethermistorfailure":             PrinterStatusDetails_MakeEnvelopeThermistorFailure,
		"makeenvelopetimingfailure":                 PrinterStatusDetails_MakeEnvelopeTimingFailure,
		"makeenvelopeturnedoff":                     PrinterStatusDetails_MakeEnvelopeTurnedOff,
		"makeenvelopeturnedon":                      PrinterStatusDetails_MakeEnvelopeTurnedOn,
		"makeenvelopeundertemperature":              PrinterStatusDetails_MakeEnvelopeUnderTemperature,
		"makeenvelopeunrecoverablefailure":          PrinterStatusDetails_MakeEnvelopeUnrecoverableFailure,
		"makeenvelopeunrecoverablestorageerror":     PrinterStatusDetails_MakeEnvelopeUnrecoverableStorageError,
		"makeenvelopewarmingup":                     PrinterStatusDetails_MakeEnvelopeWarmingUp,
		"markeradjustingprintquality":               PrinterStatusDetails_MarkerAdjustingPrintQuality,
		"markercleanermissing":                      PrinterStatusDetails_MarkerCleanerMissing,
		"markerdeveloperalmostempty":                PrinterStatusDetails_MarkerDeveloperAlmostEmpty,
		"markerdeveloperempty":                      PrinterStatusDetails_MarkerDeveloperEmpty,
		"markerdevelopermissing":                    PrinterStatusDetails_MarkerDeveloperMissing,
		"markerfusermissing":                        PrinterStatusDetails_MarkerFuserMissing,
		"markerfuserthermistorfailure":              PrinterStatusDetails_MarkerFuserThermistorFailure,
		"markerfusertimingfailure":                  PrinterStatusDetails_MarkerFuserTimingFailure,
		"markerinkalmostempty":                      PrinterStatusDetails_MarkerInkAlmostEmpty,
		"markerinkempty":                            PrinterStatusDetails_MarkerInkEmpty,
		"markerinkmissing":                          PrinterStatusDetails_MarkerInkMissing,
		"markeropcmissing":                          PrinterStatusDetails_MarkerOpcMissing,
		"markerprintribbonalmostempty":              PrinterStatusDetails_MarkerPrintRibbonAlmostEmpty,
		"markerprintribbonempty":                    PrinterStatusDetails_MarkerPrintRibbonEmpty,
		"markerprintribbonmissing":                  PrinterStatusDetails_MarkerPrintRibbonMissing,
		"markersupplyalmostempty":                   PrinterStatusDetails_MarkerSupplyAlmostEmpty,
		"markersupplyempty":                         PrinterStatusDetails_MarkerSupplyEmpty,
		"markersupplylow":                           PrinterStatusDetails_MarkerSupplyLow,
		"markersupplymissing":                       PrinterStatusDetails_MarkerSupplyMissing,
		"markertonercartridgemissing":               PrinterStatusDetails_MarkerTonerCartridgeMissing,
		"markertonermissing":                        PrinterStatusDetails_MarkerTonerMissing,
		"markerwastealmostfull":                     PrinterStatusDetails_MarkerWasteAlmostFull,
		"markerwastefull":                           PrinterStatusDetails_MarkerWasteFull,
		"markerwasteinkreceptaclealmostfull":        PrinterStatusDetails_MarkerWasteInkReceptacleAlmostFull,
		"markerwasteinkreceptaclefull":              PrinterStatusDetails_MarkerWasteInkReceptacleFull,
		"markerwasteinkreceptaclemissing":           PrinterStatusDetails_MarkerWasteInkReceptacleMissing,
		"markerwastemissing":                        PrinterStatusDetails_MarkerWasteMissing,
		"markerwastetonerreceptaclealmostfull":      PrinterStatusDetails_MarkerWasteTonerReceptacleAlmostFull,
		"markerwastetonerreceptaclefull":            PrinterStatusDetails_MarkerWasteTonerReceptacleFull,
		"markerwastetonerreceptaclemissing":         PrinterStatusDetails_MarkerWasteTonerReceptacleMissing,
		"materialempty":                             PrinterStatusDetails_MaterialEmpty,
		"materiallow":                               PrinterStatusDetails_MaterialLow,
		"materialneeded":                            PrinterStatusDetails_MaterialNeeded,
		"mediadrying":                               PrinterStatusDetails_MediaDrying,
		"mediaempty":                                PrinterStatusDetails_MediaEmpty,
		"mediajam":                                  PrinterStatusDetails_MediaJam,
		"medialow":                                  PrinterStatusDetails_MediaLow,
		"medianeeded":                               PrinterStatusDetails_MediaNeeded,
		"mediapathcannotduplexmediaselected":        PrinterStatusDetails_MediaPathCannotDuplexMediaSelected,
		"mediapathfailure":                          PrinterStatusDetails_MediaPathFailure,
		"mediapathinputempty":                       PrinterStatusDetails_MediaPathInputEmpty,
		"mediapathinputfeederror":                   PrinterStatusDetails_MediaPathInputFeedError,
		"mediapathinputjam":                         PrinterStatusDetails_MediaPathInputJam,
		"mediapathinputrequest":                     PrinterStatusDetails_MediaPathInputRequest,
		"mediapathjam":                              PrinterStatusDetails_MediaPathJam,
		"mediapathmediatrayalmostfull":              PrinterStatusDetails_MediaPathMediaTrayAlmostFull,
		"mediapathmediatrayfull":                    PrinterStatusDetails_MediaPathMediaTrayFull,
		"mediapathmediatraymissing":                 PrinterStatusDetails_MediaPathMediaTrayMissing,
		"mediapathoutputfeederror":                  PrinterStatusDetails_MediaPathOutputFeedError,
		"mediapathoutputfull":                       PrinterStatusDetails_MediaPathOutputFull,
		"mediapathoutputjam":                        PrinterStatusDetails_MediaPathOutputJam,
		"mediapathpickrollerfailure":                PrinterStatusDetails_MediaPathPickRollerFailure,
		"mediapathpickrollerlifeover":               PrinterStatusDetails_MediaPathPickRollerLifeOver,
		"mediapathpickrollerlifewarn":               PrinterStatusDetails_MediaPathPickRollerLifeWarn,
		"mediapathpickrollermissing":                PrinterStatusDetails_MediaPathPickRollerMissing,
		"motorfailure":                              PrinterStatusDetails_MotorFailure,
		"movingtopaused":                            PrinterStatusDetails_MovingToPaused,
		"none":                                      PrinterStatusDetails_None,
		"opticalphotoconductorlifeover":             PrinterStatusDetails_OpticalPhotoConductorLifeOver,
		"opticalphotoconductornearendoflife":        PrinterStatusDetails_OpticalPhotoConductorNearEndOfLife,
		"other":                                     PrinterStatusDetails_Other,
		"outputareaalmostfull":                      PrinterStatusDetails_OutputAreaAlmostFull,
		"outputareafull":                            PrinterStatusDetails_OutputAreaFull,
		"outputmailboxselectfailure":                PrinterStatusDetails_OutputMailboxSelectFailure,
		"outputmediatrayfailure":                    PrinterStatusDetails_OutputMediaTrayFailure,
		"outputmediatrayfeederror":                  PrinterStatusDetails_OutputMediaTrayFeedError,
		"outputmediatrayjam":                        PrinterStatusDetails_OutputMediaTrayJam,
		"outputtraymissing":                         PrinterStatusDetails_OutputTrayMissing,
		"paused":                                    PrinterStatusDetails_Paused,
		"perforateradded":                           PrinterStatusDetails_PerforaterAdded,
		"perforateralmostempty":                     PrinterStatusDetails_PerforaterAlmostEmpty,
		"perforateralmostfull":                      PrinterStatusDetails_PerforaterAlmostFull,
		"perforateratlimit":                         PrinterStatusDetails_PerforaterAtLimit,
		"perforaterclosed":                          PrinterStatusDetails_PerforaterClosed,
		"perforaterconfigurationchange":             PrinterStatusDetails_PerforaterConfigurationChange,
		"perforatercoverclosed":                     PrinterStatusDetails_PerforaterCoverClosed,
		"perforatercoveropen":                       PrinterStatusDetails_PerforaterCoverOpen,
		"perforaterempty":                           PrinterStatusDetails_PerforaterEmpty,
		"perforaterfull":                            PrinterStatusDetails_PerforaterFull,
		"perforaterinterlockclosed":                 PrinterStatusDetails_PerforaterInterlockClosed,
		"perforaterinterlockopen":                   PrinterStatusDetails_PerforaterInterlockOpen,
		"perforaterjam":                             PrinterStatusDetails_PerforaterJam,
		"perforaterlifealmostover":                  PrinterStatusDetails_PerforaterLifeAlmostOver,
		"perforaterlifeover":                        PrinterStatusDetails_PerforaterLifeOver,
		"perforatermemoryexhausted":                 PrinterStatusDetails_PerforaterMemoryExhausted,
		"perforatermissing":                         PrinterStatusDetails_PerforaterMissing,
		"perforatermotorfailure":                    PrinterStatusDetails_PerforaterMotorFailure,
		"perforaternearlimit":                       PrinterStatusDetails_PerforaterNearLimit,
		"perforateroffline":                         PrinterStatusDetails_PerforaterOffline,
		"perforateropened":                          PrinterStatusDetails_PerforaterOpened,
		"perforaterovertemperature":                 PrinterStatusDetails_PerforaterOverTemperature,
		"perforaterpowersaver":                      PrinterStatusDetails_PerforaterPowerSaver,
		"perforaterrecoverablefailure":              PrinterStatusDetails_PerforaterRecoverableFailure,
		"perforaterrecoverablestorage":              PrinterStatusDetails_PerforaterRecoverableStorage,
		"perforaterremoved":                         PrinterStatusDetails_PerforaterRemoved,
		"perforaterresourceadded":                   PrinterStatusDetails_PerforaterResourceAdded,
		"perforaterresourceremoved":                 PrinterStatusDetails_PerforaterResourceRemoved,
		"perforaterthermistorfailure":               PrinterStatusDetails_PerforaterThermistorFailure,
		"perforatertimingfailure":                   PrinterStatusDetails_PerforaterTimingFailure,
		"perforaterturnedoff":                       PrinterStatusDetails_PerforaterTurnedOff,
		"perforaterturnedon":                        PrinterStatusDetails_PerforaterTurnedOn,
		"perforaterundertemperature":                PrinterStatusDetails_PerforaterUnderTemperature,
		"perforaterunrecoverablefailure":            PrinterStatusDetails_PerforaterUnrecoverableFailure,
		"perforaterunrecoverablestorageerror":       PrinterStatusDetails_PerforaterUnrecoverableStorageError,
		"perforaterwarmingup":                       PrinterStatusDetails_PerforaterWarmingUp,
		"platformcooling":                           PrinterStatusDetails_PlatformCooling,
		"platformfailure":                           PrinterStatusDetails_PlatformFailure,
		"platformheating":                           PrinterStatusDetails_PlatformHeating,
		"platformtemperaturehigh":                   PrinterStatusDetails_PlatformTemperatureHigh,
		"platformtemperaturelow":                    PrinterStatusDetails_PlatformTemperatureLow,
		"powerdown":                                 PrinterStatusDetails_PowerDown,
		"powerup":                                   PrinterStatusDetails_PowerUp,
		"printermanualreset":                        PrinterStatusDetails_PrinterManualReset,
		"printernmsreset":                           PrinterStatusDetails_PrinterNmsReset,
		"printerreadytoprint":                       PrinterStatusDetails_PrinterReadyToPrint,
		"puncheradded":                              PrinterStatusDetails_PuncherAdded,
		"puncheralmostempty":                        PrinterStatusDetails_PuncherAlmostEmpty,
		"puncheralmostfull":                         PrinterStatusDetails_PuncherAlmostFull,
		"puncheratlimit":                            PrinterStatusDetails_PuncherAtLimit,
		"puncherclosed":                             PrinterStatusDetails_PuncherClosed,
		"puncherconfigurationchange":                PrinterStatusDetails_PuncherConfigurationChange,
		"punchercoverclosed":                        PrinterStatusDetails_PuncherCoverClosed,
		"punchercoveropen":                          PrinterStatusDetails_PuncherCoverOpen,
		"puncherempty":                              PrinterStatusDetails_PuncherEmpty,
		"puncherfull":                               PrinterStatusDetails_PuncherFull,
		"puncherinterlockclosed":                    PrinterStatusDetails_PuncherInterlockClosed,
		"puncherinterlockopen":                      PrinterStatusDetails_PuncherInterlockOpen,
		"puncherjam":                                PrinterStatusDetails_PuncherJam,
		"puncherlifealmostover":                     PrinterStatusDetails_PuncherLifeAlmostOver,
		"puncherlifeover":                           PrinterStatusDetails_PuncherLifeOver,
		"punchermemoryexhausted":                    PrinterStatusDetails_PuncherMemoryExhausted,
		"punchermissing":                            PrinterStatusDetails_PuncherMissing,
		"punchermotorfailure":                       PrinterStatusDetails_PuncherMotorFailure,
		"punchernearlimit":                          PrinterStatusDetails_PuncherNearLimit,
		"puncheroffline":                            PrinterStatusDetails_PuncherOffline,
		"puncheropened":                             PrinterStatusDetails_PuncherOpened,
		"puncherovertemperature":                    PrinterStatusDetails_PuncherOverTemperature,
		"puncherpowersaver":                         PrinterStatusDetails_PuncherPowerSaver,
		"puncherrecoverablefailure":                 PrinterStatusDetails_PuncherRecoverableFailure,
		"puncherrecoverablestorage":                 PrinterStatusDetails_PuncherRecoverableStorage,
		"puncherremoved":                            PrinterStatusDetails_PuncherRemoved,
		"puncherresourceadded":                      PrinterStatusDetails_PuncherResourceAdded,
		"puncherresourceremoved":                    PrinterStatusDetails_PuncherResourceRemoved,
		"puncherthermistorfailure":                  PrinterStatusDetails_PuncherThermistorFailure,
		"punchertimingfailure":                      PrinterStatusDetails_PuncherTimingFailure,
		"puncherturnedoff":                          PrinterStatusDetails_PuncherTurnedOff,
		"puncherturnedon":                           PrinterStatusDetails_PuncherTurnedOn,
		"puncherundertemperature":                   PrinterStatusDetails_PuncherUnderTemperature,
		"puncherunrecoverablefailure":               PrinterStatusDetails_PuncherUnrecoverableFailure,
		"puncherunrecoverablestorageerror":          PrinterStatusDetails_PuncherUnrecoverableStorageError,
		"puncherwarmingup":                          PrinterStatusDetails_PuncherWarmingUp,
		"resuming":                                  PrinterStatusDetails_Resuming,
		"scanmediapathfailure":                      PrinterStatusDetails_ScanMediaPathFailure,
		"scanmediapathinputempty":                   PrinterStatusDetails_ScanMediaPathInputEmpty,
		"scanmediapathinputfeederror":               PrinterStatusDetails_ScanMediaPathInputFeedError,
		"scanmediapathinputjam":                     PrinterStatusDetails_ScanMediaPathInputJam,
		"scanmediapathinputrequest":                 PrinterStatusDetails_ScanMediaPathInputRequest,
		"scanmediapathjam":                          PrinterStatusDetails_ScanMediaPathJam,
		"scanmediapathoutputfeederror":              PrinterStatusDetails_ScanMediaPathOutputFeedError,
		"scanmediapathoutputfull":                   PrinterStatusDetails_ScanMediaPathOutputFull,
		"scanmediapathoutputjam":                    PrinterStatusDetails_ScanMediaPathOutputJam,
		"scanmediapathpickrollerfailure":            PrinterStatusDetails_ScanMediaPathPickRollerFailure,
		"scanmediapathpickrollerlifeover":           PrinterStatusDetails_ScanMediaPathPickRollerLifeOver,
		"scanmediapathpickrollerlifewarn":           PrinterStatusDetails_ScanMediaPathPickRollerLifeWarn,
		"scanmediapathpickrollermissing":            PrinterStatusDetails_ScanMediaPathPickRollerMissing,
		"scanmediapathtrayalmostfull":               PrinterStatusDetails_ScanMediaPathTrayAlmostFull,
		"scanmediapathtrayfull":                     PrinterStatusDetails_ScanMediaPathTrayFull,
		"scanmediapathtraymissing":                  PrinterStatusDetails_ScanMediaPathTrayMissing,
		"scannerlightfailure":                       PrinterStatusDetails_ScannerLightFailure,
		"scannerlightlifealmostover":                PrinterStatusDetails_ScannerLightLifeAlmostOver,
		"scannerlightlifeover":                      PrinterStatusDetails_ScannerLightLifeOver,
		"scannerlightmissing":                       PrinterStatusDetails_ScannerLightMissing,
		"scannersensorfailure":                      PrinterStatusDetails_ScannerSensorFailure,
		"scannersensorlifealmostover":               PrinterStatusDetails_ScannerSensorLifeAlmostOver,
		"scannersensorlifeover":                     PrinterStatusDetails_ScannerSensorLifeOver,
		"scannersensormissing":                      PrinterStatusDetails_ScannerSensorMissing,
		"separationcutteradded":                     PrinterStatusDetails_SeparationCutterAdded,
		"separationcutteralmostempty":               PrinterStatusDetails_SeparationCutterAlmostEmpty,
		"separationcutteralmostfull":                PrinterStatusDetails_SeparationCutterAlmostFull,
		"separationcutteratlimit":                   PrinterStatusDetails_SeparationCutterAtLimit,
		"separationcutterclosed":                    PrinterStatusDetails_SeparationCutterClosed,
		"separationcutterconfigurationchange":       PrinterStatusDetails_SeparationCutterConfigurationChange,
		"separationcuttercoverclosed":               PrinterStatusDetails_SeparationCutterCoverClosed,
		"separationcuttercoveropen":                 PrinterStatusDetails_SeparationCutterCoverOpen,
		"separationcutterempty":                     PrinterStatusDetails_SeparationCutterEmpty,
		"separationcutterfull":                      PrinterStatusDetails_SeparationCutterFull,
		"separationcutterinterlockclosed":           PrinterStatusDetails_SeparationCutterInterlockClosed,
		"separationcutterinterlockopen":             PrinterStatusDetails_SeparationCutterInterlockOpen,
		"separationcutterjam":                       PrinterStatusDetails_SeparationCutterJam,
		"separationcutterlifealmostover":            PrinterStatusDetails_SeparationCutterLifeAlmostOver,
		"separationcutterlifeover":                  PrinterStatusDetails_SeparationCutterLifeOver,
		"separationcuttermemoryexhausted":           PrinterStatusDetails_SeparationCutterMemoryExhausted,
		"separationcuttermissing":                   PrinterStatusDetails_SeparationCutterMissing,
		"separationcuttermotorfailure":              PrinterStatusDetails_SeparationCutterMotorFailure,
		"separationcutternearlimit":                 PrinterStatusDetails_SeparationCutterNearLimit,
		"separationcutteroffline":                   PrinterStatusDetails_SeparationCutterOffline,
		"separationcutteropened":                    PrinterStatusDetails_SeparationCutterOpened,
		"separationcutterovertemperature":           PrinterStatusDetails_SeparationCutterOverTemperature,
		"separationcutterpowersaver":                PrinterStatusDetails_SeparationCutterPowerSaver,
		"separationcutterrecoverablefailure":        PrinterStatusDetails_SeparationCutterRecoverableFailure,
		"separationcutterrecoverablestorage":        PrinterStatusDetails_SeparationCutterRecoverableStorage,
		"separationcutterremoved":                   PrinterStatusDetails_SeparationCutterRemoved,
		"separationcutterresourceadded":             PrinterStatusDetails_SeparationCutterResourceAdded,
		"separationcutterresourceremoved":           PrinterStatusDetails_SeparationCutterResourceRemoved,
		"separationcutterthermistorfailure":         PrinterStatusDetails_SeparationCutterThermistorFailure,
		"separationcuttertimingfailure":             PrinterStatusDetails_SeparationCutterTimingFailure,
		"separationcutterturnedoff":                 PrinterStatusDetails_SeparationCutterTurnedOff,
		"separationcutterturnedon":                  PrinterStatusDetails_SeparationCutterTurnedOn,
		"separationcutterundertemperature":          PrinterStatusDetails_SeparationCutterUnderTemperature,
		"separationcutterunrecoverablefailure":      PrinterStatusDetails_SeparationCutterUnrecoverableFailure,
		"separationcutterunrecoverablestorageerror": PrinterStatusDetails_SeparationCutterUnrecoverableStorageError,
		"separationcutterwarmingup":                 PrinterStatusDetails_SeparationCutterWarmingUp,
		"sheetrotatoradded":                         PrinterStatusDetails_SheetRotatorAdded,
		"sheetrotatoralmostempty":                   PrinterStatusDetails_SheetRotatorAlmostEmpty,
		"sheetrotatoralmostfull":                    PrinterStatusDetails_SheetRotatorAlmostFull,
		"sheetrotatoratlimit":                       PrinterStatusDetails_SheetRotatorAtLimit,
		"sheetrotatorclosed":                        PrinterStatusDetails_SheetRotatorClosed,
		"sheetrotatorconfigurationchange":           PrinterStatusDetails_SheetRotatorConfigurationChange,
		"sheetrotatorcoverclosed":                   PrinterStatusDetails_SheetRotatorCoverClosed,
		"sheetrotatorcoveropen":                     PrinterStatusDetails_SheetRotatorCoverOpen,
		"sheetrotatorempty":                         PrinterStatusDetails_SheetRotatorEmpty,
		"sheetrotatorfull":                          PrinterStatusDetails_SheetRotatorFull,
		"sheetrotatorinterlockclosed":               PrinterStatusDetails_SheetRotatorInterlockClosed,
		"sheetrotatorinterlockopen":                 PrinterStatusDetails_SheetRotatorInterlockOpen,
		"sheetrotatorjam":                           PrinterStatusDetails_SheetRotatorJam,
		"sheetrotatorlifealmostover":                PrinterStatusDetails_SheetRotatorLifeAlmostOver,
		"sheetrotatorlifeover":                      PrinterStatusDetails_SheetRotatorLifeOver,
		"sheetrotatormemoryexhausted":               PrinterStatusDetails_SheetRotatorMemoryExhausted,
		"sheetrotatormissing":                       PrinterStatusDetails_SheetRotatorMissing,
		"sheetrotatormotorfailure":                  PrinterStatusDetails_SheetRotatorMotorFailure,
		"sheetrotatornearlimit":                     PrinterStatusDetails_SheetRotatorNearLimit,
		"sheetrotatoroffline":                       PrinterStatusDetails_SheetRotatorOffline,
		"sheetrotatoropened":                        PrinterStatusDetails_SheetRotatorOpened,
		"sheetrotatorovertemperature":               PrinterStatusDetails_SheetRotatorOverTemperature,
		"sheetrotatorpowersaver":                    PrinterStatusDetails_SheetRotatorPowerSaver,
		"sheetrotatorrecoverablefailure":            PrinterStatusDetails_SheetRotatorRecoverableFailure,
		"sheetrotatorrecoverablestorage":            PrinterStatusDetails_SheetRotatorRecoverableStorage,
		"sheetrotatorremoved":                       PrinterStatusDetails_SheetRotatorRemoved,
		"sheetrotatorresourceadded":                 PrinterStatusDetails_SheetRotatorResourceAdded,
		"sheetrotatorresourceremoved":               PrinterStatusDetails_SheetRotatorResourceRemoved,
		"sheetrotatorthermistorfailure":             PrinterStatusDetails_SheetRotatorThermistorFailure,
		"sheetrotatortimingfailure":                 PrinterStatusDetails_SheetRotatorTimingFailure,
		"sheetrotatorturnedoff":                     PrinterStatusDetails_SheetRotatorTurnedOff,
		"sheetrotatorturnedon":                      PrinterStatusDetails_SheetRotatorTurnedOn,
		"sheetrotatorundertemperature":              PrinterStatusDetails_SheetRotatorUnderTemperature,
		"sheetrotatorunrecoverablefailure":          PrinterStatusDetails_SheetRotatorUnrecoverableFailure,
		"sheetrotatorunrecoverablestorageerror":     PrinterStatusDetails_SheetRotatorUnrecoverableStorageError,
		"sheetrotatorwarmingup":                     PrinterStatusDetails_SheetRotatorWarmingUp,
		"shutdown":                                  PrinterStatusDetails_Shutdown,
		"slitteradded":                              PrinterStatusDetails_SlitterAdded,
		"slitteralmostempty":                        PrinterStatusDetails_SlitterAlmostEmpty,
		"slitteralmostfull":                         PrinterStatusDetails_SlitterAlmostFull,
		"slitteratlimit":                            PrinterStatusDetails_SlitterAtLimit,
		"slitterclosed":                             PrinterStatusDetails_SlitterClosed,
		"slitterconfigurationchange":                PrinterStatusDetails_SlitterConfigurationChange,
		"slittercoverclosed":                        PrinterStatusDetails_SlitterCoverClosed,
		"slittercoveropen":                          PrinterStatusDetails_SlitterCoverOpen,
		"slitterempty":                              PrinterStatusDetails_SlitterEmpty,
		"slitterfull":                               PrinterStatusDetails_SlitterFull,
		"slitterinterlockclosed":                    PrinterStatusDetails_SlitterInterlockClosed,
		"slitterinterlockopen":                      PrinterStatusDetails_SlitterInterlockOpen,
		"slitterjam":                                PrinterStatusDetails_SlitterJam,
		"slitterlifealmostover":                     PrinterStatusDetails_SlitterLifeAlmostOver,
		"slitterlifeover":                           PrinterStatusDetails_SlitterLifeOver,
		"slittermemoryexhausted":                    PrinterStatusDetails_SlitterMemoryExhausted,
		"slittermissing":                            PrinterStatusDetails_SlitterMissing,
		"slittermotorfailure":                       PrinterStatusDetails_SlitterMotorFailure,
		"slitternearlimit":                          PrinterStatusDetails_SlitterNearLimit,
		"slitteroffline":                            PrinterStatusDetails_SlitterOffline,
		"slitteropened":                             PrinterStatusDetails_SlitterOpened,
		"slitterovertemperature":                    PrinterStatusDetails_SlitterOverTemperature,
		"slitterpowersaver":                         PrinterStatusDetails_SlitterPowerSaver,
		"slitterrecoverablefailure":                 PrinterStatusDetails_SlitterRecoverableFailure,
		"slitterrecoverablestorage":                 PrinterStatusDetails_SlitterRecoverableStorage,
		"slitterremoved":                            PrinterStatusDetails_SlitterRemoved,
		"slitterresourceadded":                      PrinterStatusDetails_SlitterResourceAdded,
		"slitterresourceremoved":                    PrinterStatusDetails_SlitterResourceRemoved,
		"slitterthermistorfailure":                  PrinterStatusDetails_SlitterThermistorFailure,
		"slittertimingfailure":                      PrinterStatusDetails_SlitterTimingFailure,
		"slitterturnedoff":                          PrinterStatusDetails_SlitterTurnedOff,
		"slitterturnedon":                           PrinterStatusDetails_SlitterTurnedOn,
		"slitterundertemperature":                   PrinterStatusDetails_SlitterUnderTemperature,
		"slitterunrecoverablefailure":               PrinterStatusDetails_SlitterUnrecoverableFailure,
		"slitterunrecoverablestorageerror":          PrinterStatusDetails_SlitterUnrecoverableStorageError,
		"slitterwarmingup":                          PrinterStatusDetails_SlitterWarmingUp,
		"spoolareafull":                             PrinterStatusDetails_SpoolAreaFull,
		"stackeradded":                              PrinterStatusDetails_StackerAdded,
		"stackeralmostempty":                        PrinterStatusDetails_StackerAlmostEmpty,
		"stackeralmostfull":                         PrinterStatusDetails_StackerAlmostFull,
		"stackeratlimit":                            PrinterStatusDetails_StackerAtLimit,
		"stackerclosed":                             PrinterStatusDetails_StackerClosed,
		"stackerconfigurationchange":                PrinterStatusDetails_StackerConfigurationChange,
		"stackercoverclosed":                        PrinterStatusDetails_StackerCoverClosed,
		"stackercoveropen":                          PrinterStatusDetails_StackerCoverOpen,
		"stackerempty":                              PrinterStatusDetails_StackerEmpty,
		"stackerfull":                               PrinterStatusDetails_StackerFull,
		"stackerinterlockclosed":                    PrinterStatusDetails_StackerInterlockClosed,
		"stackerinterlockopen":                      PrinterStatusDetails_StackerInterlockOpen,
		"stackerjam":                                PrinterStatusDetails_StackerJam,
		"stackerlifealmostover":                     PrinterStatusDetails_StackerLifeAlmostOver,
		"stackerlifeover":                           PrinterStatusDetails_StackerLifeOver,
		"stackermemoryexhausted":                    PrinterStatusDetails_StackerMemoryExhausted,
		"stackermissing":                            PrinterStatusDetails_StackerMissing,
		"stackermotorfailure":                       PrinterStatusDetails_StackerMotorFailure,
		"stackernearlimit":                          PrinterStatusDetails_StackerNearLimit,
		"stackeroffline":                            PrinterStatusDetails_StackerOffline,
		"stackeropened":                             PrinterStatusDetails_StackerOpened,
		"stackerovertemperature":                    PrinterStatusDetails_StackerOverTemperature,
		"stackerpowersaver":                         PrinterStatusDetails_StackerPowerSaver,
		"stackerrecoverablefailure":                 PrinterStatusDetails_StackerRecoverableFailure,
		"stackerrecoverablestorage":                 PrinterStatusDetails_StackerRecoverableStorage,
		"stackerremoved":                            PrinterStatusDetails_StackerRemoved,
		"stackerresourceadded":                      PrinterStatusDetails_StackerResourceAdded,
		"stackerresourceremoved":                    PrinterStatusDetails_StackerResourceRemoved,
		"stackerthermistorfailure":                  PrinterStatusDetails_StackerThermistorFailure,
		"stackertimingfailure":                      PrinterStatusDetails_StackerTimingFailure,
		"stackerturnedoff":                          PrinterStatusDetails_StackerTurnedOff,
		"stackerturnedon":                           PrinterStatusDetails_StackerTurnedOn,
		"stackerundertemperature":                   PrinterStatusDetails_StackerUnderTemperature,
		"stackerunrecoverablefailure":               PrinterStatusDetails_StackerUnrecoverableFailure,
		"stackerunrecoverablestorageerror":          PrinterStatusDetails_StackerUnrecoverableStorageError,
		"stackerwarmingup":                          PrinterStatusDetails_StackerWarmingUp,
		"standby":                                   PrinterStatusDetails_Standby,
		"stapleradded":                              PrinterStatusDetails_StaplerAdded,
		"stapleralmostempty":                        PrinterStatusDetails_StaplerAlmostEmpty,
		"stapleralmostfull":                         PrinterStatusDetails_StaplerAlmostFull,
		"stapleratlimit":                            PrinterStatusDetails_StaplerAtLimit,
		"staplerclosed":                             PrinterStatusDetails_StaplerClosed,
		"staplerconfigurationchange":                PrinterStatusDetails_StaplerConfigurationChange,
		"staplercoverclosed":                        PrinterStatusDetails_StaplerCoverClosed,
		"staplercoveropen":                          PrinterStatusDetails_StaplerCoverOpen,
		"staplerempty":                              PrinterStatusDetails_StaplerEmpty,
		"staplerfull":                               PrinterStatusDetails_StaplerFull,
		"staplerinterlockclosed":                    PrinterStatusDetails_StaplerInterlockClosed,
		"staplerinterlockopen":                      PrinterStatusDetails_StaplerInterlockOpen,
		"staplerjam":                                PrinterStatusDetails_StaplerJam,
		"staplerlifealmostover":                     PrinterStatusDetails_StaplerLifeAlmostOver,
		"staplerlifeover":                           PrinterStatusDetails_StaplerLifeOver,
		"staplermemoryexhausted":                    PrinterStatusDetails_StaplerMemoryExhausted,
		"staplermissing":                            PrinterStatusDetails_StaplerMissing,
		"staplermotorfailure":                       PrinterStatusDetails_StaplerMotorFailure,
		"staplernearlimit":                          PrinterStatusDetails_StaplerNearLimit,
		"stapleroffline":                            PrinterStatusDetails_StaplerOffline,
		"stapleropened":                             PrinterStatusDetails_StaplerOpened,
		"staplerovertemperature":                    PrinterStatusDetails_StaplerOverTemperature,
		"staplerpowersaver":                         PrinterStatusDetails_StaplerPowerSaver,
		"staplerrecoverablefailure":                 PrinterStatusDetails_StaplerRecoverableFailure,
		"staplerrecoverablestorage":                 PrinterStatusDetails_StaplerRecoverableStorage,
		"staplerremoved":                            PrinterStatusDetails_StaplerRemoved,
		"staplerresourceadded":                      PrinterStatusDetails_StaplerResourceAdded,
		"staplerresourceremoved":                    PrinterStatusDetails_StaplerResourceRemoved,
		"staplerthermistorfailure":                  PrinterStatusDetails_StaplerThermistorFailure,
		"staplertimingfailure":                      PrinterStatusDetails_StaplerTimingFailure,
		"staplerturnedoff":                          PrinterStatusDetails_StaplerTurnedOff,
		"staplerturnedon":                           PrinterStatusDetails_StaplerTurnedOn,
		"staplerundertemperature":                   PrinterStatusDetails_StaplerUnderTemperature,
		"staplerunrecoverablefailure":               PrinterStatusDetails_StaplerUnrecoverableFailure,
		"staplerunrecoverablestorageerror":          PrinterStatusDetails_StaplerUnrecoverableStorageError,
		"staplerwarmingup":                          PrinterStatusDetails_StaplerWarmingUp,
		"stitcheradded":                             PrinterStatusDetails_StitcherAdded,
		"stitcheralmostempty":                       PrinterStatusDetails_StitcherAlmostEmpty,
		"stitcheralmostfull":                        PrinterStatusDetails_StitcherAlmostFull,
		"stitcheratlimit":                           PrinterStatusDetails_StitcherAtLimit,
		"stitcherclosed":                            PrinterStatusDetails_StitcherClosed,
		"stitcherconfigurationchange":               PrinterStatusDetails_StitcherConfigurationChange,
		"stitchercoverclosed":                       PrinterStatusDetails_StitcherCoverClosed,
		"stitchercoveropen":                         PrinterStatusDetails_StitcherCoverOpen,
		"stitcherempty":                             PrinterStatusDetails_StitcherEmpty,
		"stitcherfull":                              PrinterStatusDetails_StitcherFull,
		"stitcherinterlockclosed":                   PrinterStatusDetails_StitcherInterlockClosed,
		"stitcherinterlockopen":                     PrinterStatusDetails_StitcherInterlockOpen,
		"stitcherjam":                               PrinterStatusDetails_StitcherJam,
		"stitcherlifealmostover":                    PrinterStatusDetails_StitcherLifeAlmostOver,
		"stitcherlifeover":                          PrinterStatusDetails_StitcherLifeOver,
		"stitchermemoryexhausted":                   PrinterStatusDetails_StitcherMemoryExhausted,
		"stitchermissing":                           PrinterStatusDetails_StitcherMissing,
		"stitchermotorfailure":                      PrinterStatusDetails_StitcherMotorFailure,
		"stitchernearlimit":                         PrinterStatusDetails_StitcherNearLimit,
		"stitcheroffline":                           PrinterStatusDetails_StitcherOffline,
		"stitcheropened":                            PrinterStatusDetails_StitcherOpened,
		"stitcherovertemperature":                   PrinterStatusDetails_StitcherOverTemperature,
		"stitcherpowersaver":                        PrinterStatusDetails_StitcherPowerSaver,
		"stitcherrecoverablefailure":                PrinterStatusDetails_StitcherRecoverableFailure,
		"stitcherrecoverablestorage":                PrinterStatusDetails_StitcherRecoverableStorage,
		"stitcherremoved":                           PrinterStatusDetails_StitcherRemoved,
		"stitcherresourceadded":                     PrinterStatusDetails_StitcherResourceAdded,
		"stitcherresourceremoved":                   PrinterStatusDetails_StitcherResourceRemoved,
		"stitcherthermistorfailure":                 PrinterStatusDetails_StitcherThermistorFailure,
		"stitchertimingfailure":                     PrinterStatusDetails_StitcherTimingFailure,
		"stitcherturnedoff":                         PrinterStatusDetails_StitcherTurnedOff,
		"stitcherturnedon":                          PrinterStatusDetails_StitcherTurnedOn,
		"stitcherundertemperature":                  PrinterStatusDetails_StitcherUnderTemperature,
		"stitcherunrecoverablefailure":              PrinterStatusDetails_StitcherUnrecoverableFailure,
		"stitcherunrecoverablestorageerror":         PrinterStatusDetails_StitcherUnrecoverableStorageError,
		"stitcherwarmingup":                         PrinterStatusDetails_StitcherWarmingUp,
		"stoppedpartially":                          PrinterStatusDetails_StoppedPartially,
		"stopping":                                  PrinterStatusDetails_Stopping,
		"subunitadded":                              PrinterStatusDetails_SubunitAdded,
		"subunitalmostempty":                        PrinterStatusDetails_SubunitAlmostEmpty,
		"subunitalmostfull":                         PrinterStatusDetails_SubunitAlmostFull,
		"subunitatlimit":                            PrinterStatusDetails_SubunitAtLimit,
		"subunitclosed":                             PrinterStatusDetails_SubunitClosed,
		"subunitcoolingdown":                        PrinterStatusDetails_SubunitCoolingDown,
		"subunitempty":                              PrinterStatusDetails_SubunitEmpty,
		"subunitfull":                               PrinterStatusDetails_SubunitFull,
		"subunitlifealmostover":                     PrinterStatusDetails_SubunitLifeAlmostOver,
		"subunitlifeover":                           PrinterStatusDetails_SubunitLifeOver,
		"subunitmemoryexhausted":                    PrinterStatusDetails_SubunitMemoryExhausted,
		"subunitmissing":                            PrinterStatusDetails_SubunitMissing,
		"subunitmotorfailure":                       PrinterStatusDetails_SubunitMotorFailure,
		"subunitnearlimit":                          PrinterStatusDetails_SubunitNearLimit,
		"subunitoffline":                            PrinterStatusDetails_SubunitOffline,
		"subunitopened":                             PrinterStatusDetails_SubunitOpened,
		"subunitovertemperature":                    PrinterStatusDetails_SubunitOverTemperature,
		"subunitpowersaver":                         PrinterStatusDetails_SubunitPowerSaver,
		"subunitrecoverablefailure":                 PrinterStatusDetails_SubunitRecoverableFailure,
		"subunitrecoverablestorage":                 PrinterStatusDetails_SubunitRecoverableStorage,
		"subunitremoved":                            PrinterStatusDetails_SubunitRemoved,
		"subunitresourceadded":                      PrinterStatusDetails_SubunitResourceAdded,
		"subunitresourceremoved":                    PrinterStatusDetails_SubunitResourceRemoved,
		"subunitthermistorfailure":                  PrinterStatusDetails_SubunitThermistorFailure,
		"subunittimingfailure":                      PrinterStatusDetails_SubunitTimingFailure,
		"subunitturnedoff":                          PrinterStatusDetails_SubunitTurnedOff,
		"subunitturnedon":                           PrinterStatusDetails_SubunitTurnedOn,
		"subunitundertemperature":                   PrinterStatusDetails_SubunitUnderTemperature,
		"subunitunrecoverablefailure":               PrinterStatusDetails_SubunitUnrecoverableFailure,
		"subunitunrecoverablestorage":               PrinterStatusDetails_SubunitUnrecoverableStorage,
		"subunitwarmingup":                          PrinterStatusDetails_SubunitWarmingUp,
		"suspend":                                   PrinterStatusDetails_Suspend,
		"testing":                                   PrinterStatusDetails_Testing,
		"timedout":                                  PrinterStatusDetails_TimedOut,
		"tonerempty":                                PrinterStatusDetails_TonerEmpty,
		"tonerlow":                                  PrinterStatusDetails_TonerLow,
		"trimmeradded":                              PrinterStatusDetails_TrimmerAdded,
		"trimmeralmostempty":                        PrinterStatusDetails_TrimmerAlmostEmpty,
		"trimmeralmostfull":                         PrinterStatusDetails_TrimmerAlmostFull,
		"trimmeratlimit":                            PrinterStatusDetails_TrimmerAtLimit,
		"trimmerclosed":                             PrinterStatusDetails_TrimmerClosed,
		"trimmerconfigurationchange":                PrinterStatusDetails_TrimmerConfigurationChange,
		"trimmercoverclosed":                        PrinterStatusDetails_TrimmerCoverClosed,
		"trimmercoveropen":                          PrinterStatusDetails_TrimmerCoverOpen,
		"trimmerempty":                              PrinterStatusDetails_TrimmerEmpty,
		"trimmerfull":                               PrinterStatusDetails_TrimmerFull,
		"trimmerinterlockclosed":                    PrinterStatusDetails_TrimmerInterlockClosed,
		"trimmerinterlockopen":                      PrinterStatusDetails_TrimmerInterlockOpen,
		"trimmerjam":                                PrinterStatusDetails_TrimmerJam,
		"trimmerlifealmostover":                     PrinterStatusDetails_TrimmerLifeAlmostOver,
		"trimmerlifeover":                           PrinterStatusDetails_TrimmerLifeOver,
		"trimmermemoryexhausted":                    PrinterStatusDetails_TrimmerMemoryExhausted,
		"trimmermissing":                            PrinterStatusDetails_TrimmerMissing,
		"trimmermotorfailure":                       PrinterStatusDetails_TrimmerMotorFailure,
		"trimmernearlimit":                          PrinterStatusDetails_TrimmerNearLimit,
		"trimmeroffline":                            PrinterStatusDetails_TrimmerOffline,
		"trimmeropened":                             PrinterStatusDetails_TrimmerOpened,
		"trimmerovertemperature":                    PrinterStatusDetails_TrimmerOverTemperature,
		"trimmerpowersaver":                         PrinterStatusDetails_TrimmerPowerSaver,
		"trimmerrecoverablefailure":                 PrinterStatusDetails_TrimmerRecoverableFailure,
		"trimmerrecoverablestorage":                 PrinterStatusDetails_TrimmerRecoverableStorage,
		"trimmerremoved":                            PrinterStatusDetails_TrimmerRemoved,
		"trimmerresourceadded":                      PrinterStatusDetails_TrimmerResourceAdded,
		"trimmerresourceremoved":                    PrinterStatusDetails_TrimmerResourceRemoved,
		"trimmerthermistorfailure":                  PrinterStatusDetails_TrimmerThermistorFailure,
		"trimmertimingfailure":                      PrinterStatusDetails_TrimmerTimingFailure,
		"trimmerturnedoff":                          PrinterStatusDetails_TrimmerTurnedOff,
		"trimmerturnedon":                           PrinterStatusDetails_TrimmerTurnedOn,
		"trimmerundertemperature":                   PrinterStatusDetails_TrimmerUnderTemperature,
		"trimmerunrecoverablefailure":               PrinterStatusDetails_TrimmerUnrecoverableFailure,
		"trimmerunrecoverablestorageerror":          PrinterStatusDetails_TrimmerUnrecoverableStorageError,
		"trimmerwarmingup":                          PrinterStatusDetails_TrimmerWarmingUp,
		"unknown":                                   PrinterStatusDetails_Unknown,
		"wrapperadded":                              PrinterStatusDetails_WrapperAdded,
		"wrapperalmostempty":                        PrinterStatusDetails_WrapperAlmostEmpty,
		"wrapperalmostfull":                         PrinterStatusDetails_WrapperAlmostFull,
		"wrapperatlimit":                            PrinterStatusDetails_WrapperAtLimit,
		"wrapperclosed":                             PrinterStatusDetails_WrapperClosed,
		"wrapperconfigurationchange":                PrinterStatusDetails_WrapperConfigurationChange,
		"wrappercoverclosed":                        PrinterStatusDetails_WrapperCoverClosed,
		"wrappercoveropen":                          PrinterStatusDetails_WrapperCoverOpen,
		"wrapperempty":                              PrinterStatusDetails_WrapperEmpty,
		"wrapperfull":                               PrinterStatusDetails_WrapperFull,
		"wrapperinterlockclosed":                    PrinterStatusDetails_WrapperInterlockClosed,
		"wrapperinterlockopen":                      PrinterStatusDetails_WrapperInterlockOpen,
		"wrapperjam":                                PrinterStatusDetails_WrapperJam,
		"wrapperlifealmostover":                     PrinterStatusDetails_WrapperLifeAlmostOver,
		"wrapperlifeover":                           PrinterStatusDetails_WrapperLifeOver,
		"wrappermemoryexhausted":                    PrinterStatusDetails_WrapperMemoryExhausted,
		"wrappermissing":                            PrinterStatusDetails_WrapperMissing,
		"wrappermotorfailure":                       PrinterStatusDetails_WrapperMotorFailure,
		"wrappernearlimit":                          PrinterStatusDetails_WrapperNearLimit,
		"wrapperoffline":                            PrinterStatusDetails_WrapperOffline,
		"wrapperopened":                             PrinterStatusDetails_WrapperOpened,
		"wrapperovertemperature":                    PrinterStatusDetails_WrapperOverTemperature,
		"wrapperpowersaver":                         PrinterStatusDetails_WrapperPowerSaver,
		"wrapperrecoverablefailure":                 PrinterStatusDetails_WrapperRecoverableFailure,
		"wrapperrecoverablestorage":                 PrinterStatusDetails_WrapperRecoverableStorage,
		"wrapperremoved":                            PrinterStatusDetails_WrapperRemoved,
		"wrapperresourceadded":                      PrinterStatusDetails_WrapperResourceAdded,
		"wrapperresourceremoved":                    PrinterStatusDetails_WrapperResourceRemoved,
		"wrapperthermistorfailure":                  PrinterStatusDetails_WrapperThermistorFailure,
		"wrappertimingfailure":                      PrinterStatusDetails_WrapperTimingFailure,
		"wrapperturnedoff":                          PrinterStatusDetails_WrapperTurnedOff,
		"wrapperturnedon":                           PrinterStatusDetails_WrapperTurnedOn,
		"wrapperundertemperature":                   PrinterStatusDetails_WrapperUnderTemperature,
		"wrapperunrecoverablefailure":               PrinterStatusDetails_WrapperUnrecoverableFailure,
		"wrapperunrecoverablestorageerror":          PrinterStatusDetails_WrapperUnrecoverableStorageError,
		"wrapperwarmingup":                          PrinterStatusDetails_WrapperWarmingUp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterStatusDetails(input)
	return &out, nil
}
