package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterStatusDetails string

const (
	PrinterStatusDetailsalertRemovalOfBinaryChangeEntry           PrinterStatusDetails = "AlertRemovalOfBinaryChangeEntry"
	PrinterStatusDetailsbanderAdded                               PrinterStatusDetails = "BanderAdded"
	PrinterStatusDetailsbanderAlmostEmpty                         PrinterStatusDetails = "BanderAlmostEmpty"
	PrinterStatusDetailsbanderAlmostFull                          PrinterStatusDetails = "BanderAlmostFull"
	PrinterStatusDetailsbanderAtLimit                             PrinterStatusDetails = "BanderAtLimit"
	PrinterStatusDetailsbanderClosed                              PrinterStatusDetails = "BanderClosed"
	PrinterStatusDetailsbanderConfigurationChange                 PrinterStatusDetails = "BanderConfigurationChange"
	PrinterStatusDetailsbanderCoverClosed                         PrinterStatusDetails = "BanderCoverClosed"
	PrinterStatusDetailsbanderCoverOpen                           PrinterStatusDetails = "BanderCoverOpen"
	PrinterStatusDetailsbanderEmpty                               PrinterStatusDetails = "BanderEmpty"
	PrinterStatusDetailsbanderFull                                PrinterStatusDetails = "BanderFull"
	PrinterStatusDetailsbanderInterlockClosed                     PrinterStatusDetails = "BanderInterlockClosed"
	PrinterStatusDetailsbanderInterlockOpen                       PrinterStatusDetails = "BanderInterlockOpen"
	PrinterStatusDetailsbanderJam                                 PrinterStatusDetails = "BanderJam"
	PrinterStatusDetailsbanderLifeAlmostOver                      PrinterStatusDetails = "BanderLifeAlmostOver"
	PrinterStatusDetailsbanderLifeOver                            PrinterStatusDetails = "BanderLifeOver"
	PrinterStatusDetailsbanderMemoryExhausted                     PrinterStatusDetails = "BanderMemoryExhausted"
	PrinterStatusDetailsbanderMissing                             PrinterStatusDetails = "BanderMissing"
	PrinterStatusDetailsbanderMotorFailure                        PrinterStatusDetails = "BanderMotorFailure"
	PrinterStatusDetailsbanderNearLimit                           PrinterStatusDetails = "BanderNearLimit"
	PrinterStatusDetailsbanderOffline                             PrinterStatusDetails = "BanderOffline"
	PrinterStatusDetailsbanderOpened                              PrinterStatusDetails = "BanderOpened"
	PrinterStatusDetailsbanderOverTemperature                     PrinterStatusDetails = "BanderOverTemperature"
	PrinterStatusDetailsbanderPowerSaver                          PrinterStatusDetails = "BanderPowerSaver"
	PrinterStatusDetailsbanderRecoverableFailure                  PrinterStatusDetails = "BanderRecoverableFailure"
	PrinterStatusDetailsbanderRecoverableStorage                  PrinterStatusDetails = "BanderRecoverableStorage"
	PrinterStatusDetailsbanderRemoved                             PrinterStatusDetails = "BanderRemoved"
	PrinterStatusDetailsbanderResourceAdded                       PrinterStatusDetails = "BanderResourceAdded"
	PrinterStatusDetailsbanderResourceRemoved                     PrinterStatusDetails = "BanderResourceRemoved"
	PrinterStatusDetailsbanderThermistorFailure                   PrinterStatusDetails = "BanderThermistorFailure"
	PrinterStatusDetailsbanderTimingFailure                       PrinterStatusDetails = "BanderTimingFailure"
	PrinterStatusDetailsbanderTurnedOff                           PrinterStatusDetails = "BanderTurnedOff"
	PrinterStatusDetailsbanderTurnedOn                            PrinterStatusDetails = "BanderTurnedOn"
	PrinterStatusDetailsbanderUnderTemperature                    PrinterStatusDetails = "BanderUnderTemperature"
	PrinterStatusDetailsbanderUnrecoverableFailure                PrinterStatusDetails = "BanderUnrecoverableFailure"
	PrinterStatusDetailsbanderUnrecoverableStorageError           PrinterStatusDetails = "BanderUnrecoverableStorageError"
	PrinterStatusDetailsbanderWarmingUp                           PrinterStatusDetails = "BanderWarmingUp"
	PrinterStatusDetailsbinderAdded                               PrinterStatusDetails = "BinderAdded"
	PrinterStatusDetailsbinderAlmostEmpty                         PrinterStatusDetails = "BinderAlmostEmpty"
	PrinterStatusDetailsbinderAlmostFull                          PrinterStatusDetails = "BinderAlmostFull"
	PrinterStatusDetailsbinderAtLimit                             PrinterStatusDetails = "BinderAtLimit"
	PrinterStatusDetailsbinderClosed                              PrinterStatusDetails = "BinderClosed"
	PrinterStatusDetailsbinderConfigurationChange                 PrinterStatusDetails = "BinderConfigurationChange"
	PrinterStatusDetailsbinderCoverClosed                         PrinterStatusDetails = "BinderCoverClosed"
	PrinterStatusDetailsbinderCoverOpen                           PrinterStatusDetails = "BinderCoverOpen"
	PrinterStatusDetailsbinderEmpty                               PrinterStatusDetails = "BinderEmpty"
	PrinterStatusDetailsbinderFull                                PrinterStatusDetails = "BinderFull"
	PrinterStatusDetailsbinderInterlockClosed                     PrinterStatusDetails = "BinderInterlockClosed"
	PrinterStatusDetailsbinderInterlockOpen                       PrinterStatusDetails = "BinderInterlockOpen"
	PrinterStatusDetailsbinderJam                                 PrinterStatusDetails = "BinderJam"
	PrinterStatusDetailsbinderLifeAlmostOver                      PrinterStatusDetails = "BinderLifeAlmostOver"
	PrinterStatusDetailsbinderLifeOver                            PrinterStatusDetails = "BinderLifeOver"
	PrinterStatusDetailsbinderMemoryExhausted                     PrinterStatusDetails = "BinderMemoryExhausted"
	PrinterStatusDetailsbinderMissing                             PrinterStatusDetails = "BinderMissing"
	PrinterStatusDetailsbinderMotorFailure                        PrinterStatusDetails = "BinderMotorFailure"
	PrinterStatusDetailsbinderNearLimit                           PrinterStatusDetails = "BinderNearLimit"
	PrinterStatusDetailsbinderOffline                             PrinterStatusDetails = "BinderOffline"
	PrinterStatusDetailsbinderOpened                              PrinterStatusDetails = "BinderOpened"
	PrinterStatusDetailsbinderOverTemperature                     PrinterStatusDetails = "BinderOverTemperature"
	PrinterStatusDetailsbinderPowerSaver                          PrinterStatusDetails = "BinderPowerSaver"
	PrinterStatusDetailsbinderRecoverableFailure                  PrinterStatusDetails = "BinderRecoverableFailure"
	PrinterStatusDetailsbinderRecoverableStorage                  PrinterStatusDetails = "BinderRecoverableStorage"
	PrinterStatusDetailsbinderRemoved                             PrinterStatusDetails = "BinderRemoved"
	PrinterStatusDetailsbinderResourceAdded                       PrinterStatusDetails = "BinderResourceAdded"
	PrinterStatusDetailsbinderResourceRemoved                     PrinterStatusDetails = "BinderResourceRemoved"
	PrinterStatusDetailsbinderThermistorFailure                   PrinterStatusDetails = "BinderThermistorFailure"
	PrinterStatusDetailsbinderTimingFailure                       PrinterStatusDetails = "BinderTimingFailure"
	PrinterStatusDetailsbinderTurnedOff                           PrinterStatusDetails = "BinderTurnedOff"
	PrinterStatusDetailsbinderTurnedOn                            PrinterStatusDetails = "BinderTurnedOn"
	PrinterStatusDetailsbinderUnderTemperature                    PrinterStatusDetails = "BinderUnderTemperature"
	PrinterStatusDetailsbinderUnrecoverableFailure                PrinterStatusDetails = "BinderUnrecoverableFailure"
	PrinterStatusDetailsbinderUnrecoverableStorageError           PrinterStatusDetails = "BinderUnrecoverableStorageError"
	PrinterStatusDetailsbinderWarmingUp                           PrinterStatusDetails = "BinderWarmingUp"
	PrinterStatusDetailscameraFailure                             PrinterStatusDetails = "CameraFailure"
	PrinterStatusDetailschamberCooling                            PrinterStatusDetails = "ChamberCooling"
	PrinterStatusDetailschamberFailure                            PrinterStatusDetails = "ChamberFailure"
	PrinterStatusDetailschamberHeating                            PrinterStatusDetails = "ChamberHeating"
	PrinterStatusDetailschamberTemperatureHigh                    PrinterStatusDetails = "ChamberTemperatureHigh"
	PrinterStatusDetailschamberTemperatureLow                     PrinterStatusDetails = "ChamberTemperatureLow"
	PrinterStatusDetailscleanerLifeAlmostOver                     PrinterStatusDetails = "CleanerLifeAlmostOver"
	PrinterStatusDetailscleanerLifeOver                           PrinterStatusDetails = "CleanerLifeOver"
	PrinterStatusDetailsconfigurationChange                       PrinterStatusDetails = "ConfigurationChange"
	PrinterStatusDetailsconnectingToDevice                        PrinterStatusDetails = "ConnectingToDevice"
	PrinterStatusDetailscoverOpen                                 PrinterStatusDetails = "CoverOpen"
	PrinterStatusDetailsdeactivated                               PrinterStatusDetails = "Deactivated"
	PrinterStatusDetailsdeleted                                   PrinterStatusDetails = "Deleted"
	PrinterStatusDetailsdeveloperEmpty                            PrinterStatusDetails = "DeveloperEmpty"
	PrinterStatusDetailsdeveloperLow                              PrinterStatusDetails = "DeveloperLow"
	PrinterStatusDetailsdieCutterAdded                            PrinterStatusDetails = "DieCutterAdded"
	PrinterStatusDetailsdieCutterAlmostEmpty                      PrinterStatusDetails = "DieCutterAlmostEmpty"
	PrinterStatusDetailsdieCutterAlmostFull                       PrinterStatusDetails = "DieCutterAlmostFull"
	PrinterStatusDetailsdieCutterAtLimit                          PrinterStatusDetails = "DieCutterAtLimit"
	PrinterStatusDetailsdieCutterClosed                           PrinterStatusDetails = "DieCutterClosed"
	PrinterStatusDetailsdieCutterConfigurationChange              PrinterStatusDetails = "DieCutterConfigurationChange"
	PrinterStatusDetailsdieCutterCoverClosed                      PrinterStatusDetails = "DieCutterCoverClosed"
	PrinterStatusDetailsdieCutterCoverOpen                        PrinterStatusDetails = "DieCutterCoverOpen"
	PrinterStatusDetailsdieCutterEmpty                            PrinterStatusDetails = "DieCutterEmpty"
	PrinterStatusDetailsdieCutterFull                             PrinterStatusDetails = "DieCutterFull"
	PrinterStatusDetailsdieCutterInterlockClosed                  PrinterStatusDetails = "DieCutterInterlockClosed"
	PrinterStatusDetailsdieCutterInterlockOpen                    PrinterStatusDetails = "DieCutterInterlockOpen"
	PrinterStatusDetailsdieCutterJam                              PrinterStatusDetails = "DieCutterJam"
	PrinterStatusDetailsdieCutterLifeAlmostOver                   PrinterStatusDetails = "DieCutterLifeAlmostOver"
	PrinterStatusDetailsdieCutterLifeOver                         PrinterStatusDetails = "DieCutterLifeOver"
	PrinterStatusDetailsdieCutterMemoryExhausted                  PrinterStatusDetails = "DieCutterMemoryExhausted"
	PrinterStatusDetailsdieCutterMissing                          PrinterStatusDetails = "DieCutterMissing"
	PrinterStatusDetailsdieCutterMotorFailure                     PrinterStatusDetails = "DieCutterMotorFailure"
	PrinterStatusDetailsdieCutterNearLimit                        PrinterStatusDetails = "DieCutterNearLimit"
	PrinterStatusDetailsdieCutterOffline                          PrinterStatusDetails = "DieCutterOffline"
	PrinterStatusDetailsdieCutterOpened                           PrinterStatusDetails = "DieCutterOpened"
	PrinterStatusDetailsdieCutterOverTemperature                  PrinterStatusDetails = "DieCutterOverTemperature"
	PrinterStatusDetailsdieCutterPowerSaver                       PrinterStatusDetails = "DieCutterPowerSaver"
	PrinterStatusDetailsdieCutterRecoverableFailure               PrinterStatusDetails = "DieCutterRecoverableFailure"
	PrinterStatusDetailsdieCutterRecoverableStorage               PrinterStatusDetails = "DieCutterRecoverableStorage"
	PrinterStatusDetailsdieCutterRemoved                          PrinterStatusDetails = "DieCutterRemoved"
	PrinterStatusDetailsdieCutterResourceAdded                    PrinterStatusDetails = "DieCutterResourceAdded"
	PrinterStatusDetailsdieCutterResourceRemoved                  PrinterStatusDetails = "DieCutterResourceRemoved"
	PrinterStatusDetailsdieCutterThermistorFailure                PrinterStatusDetails = "DieCutterThermistorFailure"
	PrinterStatusDetailsdieCutterTimingFailure                    PrinterStatusDetails = "DieCutterTimingFailure"
	PrinterStatusDetailsdieCutterTurnedOff                        PrinterStatusDetails = "DieCutterTurnedOff"
	PrinterStatusDetailsdieCutterTurnedOn                         PrinterStatusDetails = "DieCutterTurnedOn"
	PrinterStatusDetailsdieCutterUnderTemperature                 PrinterStatusDetails = "DieCutterUnderTemperature"
	PrinterStatusDetailsdieCutterUnrecoverableFailure             PrinterStatusDetails = "DieCutterUnrecoverableFailure"
	PrinterStatusDetailsdieCutterUnrecoverableStorageError        PrinterStatusDetails = "DieCutterUnrecoverableStorageError"
	PrinterStatusDetailsdieCutterWarmingUp                        PrinterStatusDetails = "DieCutterWarmingUp"
	PrinterStatusDetailsdoorOpen                                  PrinterStatusDetails = "DoorOpen"
	PrinterStatusDetailsextruderCooling                           PrinterStatusDetails = "ExtruderCooling"
	PrinterStatusDetailsextruderFailure                           PrinterStatusDetails = "ExtruderFailure"
	PrinterStatusDetailsextruderHeating                           PrinterStatusDetails = "ExtruderHeating"
	PrinterStatusDetailsextruderJam                               PrinterStatusDetails = "ExtruderJam"
	PrinterStatusDetailsextruderTemperatureHigh                   PrinterStatusDetails = "ExtruderTemperatureHigh"
	PrinterStatusDetailsextruderTemperatureLow                    PrinterStatusDetails = "ExtruderTemperatureLow"
	PrinterStatusDetailsfanFailure                                PrinterStatusDetails = "FanFailure"
	PrinterStatusDetailsfaxModemLifeAlmostOver                    PrinterStatusDetails = "FaxModemLifeAlmostOver"
	PrinterStatusDetailsfaxModemLifeOver                          PrinterStatusDetails = "FaxModemLifeOver"
	PrinterStatusDetailsfaxModemMissing                           PrinterStatusDetails = "FaxModemMissing"
	PrinterStatusDetailsfaxModemTurnedOff                         PrinterStatusDetails = "FaxModemTurnedOff"
	PrinterStatusDetailsfaxModemTurnedOn                          PrinterStatusDetails = "FaxModemTurnedOn"
	PrinterStatusDetailsfolderAdded                               PrinterStatusDetails = "FolderAdded"
	PrinterStatusDetailsfolderAlmostEmpty                         PrinterStatusDetails = "FolderAlmostEmpty"
	PrinterStatusDetailsfolderAlmostFull                          PrinterStatusDetails = "FolderAlmostFull"
	PrinterStatusDetailsfolderAtLimit                             PrinterStatusDetails = "FolderAtLimit"
	PrinterStatusDetailsfolderClosed                              PrinterStatusDetails = "FolderClosed"
	PrinterStatusDetailsfolderConfigurationChange                 PrinterStatusDetails = "FolderConfigurationChange"
	PrinterStatusDetailsfolderCoverClosed                         PrinterStatusDetails = "FolderCoverClosed"
	PrinterStatusDetailsfolderCoverOpen                           PrinterStatusDetails = "FolderCoverOpen"
	PrinterStatusDetailsfolderEmpty                               PrinterStatusDetails = "FolderEmpty"
	PrinterStatusDetailsfolderFull                                PrinterStatusDetails = "FolderFull"
	PrinterStatusDetailsfolderInterlockClosed                     PrinterStatusDetails = "FolderInterlockClosed"
	PrinterStatusDetailsfolderInterlockOpen                       PrinterStatusDetails = "FolderInterlockOpen"
	PrinterStatusDetailsfolderJam                                 PrinterStatusDetails = "FolderJam"
	PrinterStatusDetailsfolderLifeAlmostOver                      PrinterStatusDetails = "FolderLifeAlmostOver"
	PrinterStatusDetailsfolderLifeOver                            PrinterStatusDetails = "FolderLifeOver"
	PrinterStatusDetailsfolderMemoryExhausted                     PrinterStatusDetails = "FolderMemoryExhausted"
	PrinterStatusDetailsfolderMissing                             PrinterStatusDetails = "FolderMissing"
	PrinterStatusDetailsfolderMotorFailure                        PrinterStatusDetails = "FolderMotorFailure"
	PrinterStatusDetailsfolderNearLimit                           PrinterStatusDetails = "FolderNearLimit"
	PrinterStatusDetailsfolderOffline                             PrinterStatusDetails = "FolderOffline"
	PrinterStatusDetailsfolderOpened                              PrinterStatusDetails = "FolderOpened"
	PrinterStatusDetailsfolderOverTemperature                     PrinterStatusDetails = "FolderOverTemperature"
	PrinterStatusDetailsfolderPowerSaver                          PrinterStatusDetails = "FolderPowerSaver"
	PrinterStatusDetailsfolderRecoverableFailure                  PrinterStatusDetails = "FolderRecoverableFailure"
	PrinterStatusDetailsfolderRecoverableStorage                  PrinterStatusDetails = "FolderRecoverableStorage"
	PrinterStatusDetailsfolderRemoved                             PrinterStatusDetails = "FolderRemoved"
	PrinterStatusDetailsfolderResourceAdded                       PrinterStatusDetails = "FolderResourceAdded"
	PrinterStatusDetailsfolderResourceRemoved                     PrinterStatusDetails = "FolderResourceRemoved"
	PrinterStatusDetailsfolderThermistorFailure                   PrinterStatusDetails = "FolderThermistorFailure"
	PrinterStatusDetailsfolderTimingFailure                       PrinterStatusDetails = "FolderTimingFailure"
	PrinterStatusDetailsfolderTurnedOff                           PrinterStatusDetails = "FolderTurnedOff"
	PrinterStatusDetailsfolderTurnedOn                            PrinterStatusDetails = "FolderTurnedOn"
	PrinterStatusDetailsfolderUnderTemperature                    PrinterStatusDetails = "FolderUnderTemperature"
	PrinterStatusDetailsfolderUnrecoverableFailure                PrinterStatusDetails = "FolderUnrecoverableFailure"
	PrinterStatusDetailsfolderUnrecoverableStorageError           PrinterStatusDetails = "FolderUnrecoverableStorageError"
	PrinterStatusDetailsfolderWarmingUp                           PrinterStatusDetails = "FolderWarmingUp"
	PrinterStatusDetailsfuserOverTemp                             PrinterStatusDetails = "FuserOverTemp"
	PrinterStatusDetailsfuserUnderTemp                            PrinterStatusDetails = "FuserUnderTemp"
	PrinterStatusDetailshibernate                                 PrinterStatusDetails = "Hibernate"
	PrinterStatusDetailsholdNewJobs                               PrinterStatusDetails = "HoldNewJobs"
	PrinterStatusDetailsidentifyPrinterRequested                  PrinterStatusDetails = "IdentifyPrinterRequested"
	PrinterStatusDetailsimprinterAdded                            PrinterStatusDetails = "ImprinterAdded"
	PrinterStatusDetailsimprinterAlmostEmpty                      PrinterStatusDetails = "ImprinterAlmostEmpty"
	PrinterStatusDetailsimprinterAlmostFull                       PrinterStatusDetails = "ImprinterAlmostFull"
	PrinterStatusDetailsimprinterAtLimit                          PrinterStatusDetails = "ImprinterAtLimit"
	PrinterStatusDetailsimprinterClosed                           PrinterStatusDetails = "ImprinterClosed"
	PrinterStatusDetailsimprinterConfigurationChange              PrinterStatusDetails = "ImprinterConfigurationChange"
	PrinterStatusDetailsimprinterCoverClosed                      PrinterStatusDetails = "ImprinterCoverClosed"
	PrinterStatusDetailsimprinterCoverOpen                        PrinterStatusDetails = "ImprinterCoverOpen"
	PrinterStatusDetailsimprinterEmpty                            PrinterStatusDetails = "ImprinterEmpty"
	PrinterStatusDetailsimprinterFull                             PrinterStatusDetails = "ImprinterFull"
	PrinterStatusDetailsimprinterInterlockClosed                  PrinterStatusDetails = "ImprinterInterlockClosed"
	PrinterStatusDetailsimprinterInterlockOpen                    PrinterStatusDetails = "ImprinterInterlockOpen"
	PrinterStatusDetailsimprinterJam                              PrinterStatusDetails = "ImprinterJam"
	PrinterStatusDetailsimprinterLifeAlmostOver                   PrinterStatusDetails = "ImprinterLifeAlmostOver"
	PrinterStatusDetailsimprinterLifeOver                         PrinterStatusDetails = "ImprinterLifeOver"
	PrinterStatusDetailsimprinterMemoryExhausted                  PrinterStatusDetails = "ImprinterMemoryExhausted"
	PrinterStatusDetailsimprinterMissing                          PrinterStatusDetails = "ImprinterMissing"
	PrinterStatusDetailsimprinterMotorFailure                     PrinterStatusDetails = "ImprinterMotorFailure"
	PrinterStatusDetailsimprinterNearLimit                        PrinterStatusDetails = "ImprinterNearLimit"
	PrinterStatusDetailsimprinterOffline                          PrinterStatusDetails = "ImprinterOffline"
	PrinterStatusDetailsimprinterOpened                           PrinterStatusDetails = "ImprinterOpened"
	PrinterStatusDetailsimprinterOverTemperature                  PrinterStatusDetails = "ImprinterOverTemperature"
	PrinterStatusDetailsimprinterPowerSaver                       PrinterStatusDetails = "ImprinterPowerSaver"
	PrinterStatusDetailsimprinterRecoverableFailure               PrinterStatusDetails = "ImprinterRecoverableFailure"
	PrinterStatusDetailsimprinterRecoverableStorage               PrinterStatusDetails = "ImprinterRecoverableStorage"
	PrinterStatusDetailsimprinterRemoved                          PrinterStatusDetails = "ImprinterRemoved"
	PrinterStatusDetailsimprinterResourceAdded                    PrinterStatusDetails = "ImprinterResourceAdded"
	PrinterStatusDetailsimprinterResourceRemoved                  PrinterStatusDetails = "ImprinterResourceRemoved"
	PrinterStatusDetailsimprinterThermistorFailure                PrinterStatusDetails = "ImprinterThermistorFailure"
	PrinterStatusDetailsimprinterTimingFailure                    PrinterStatusDetails = "ImprinterTimingFailure"
	PrinterStatusDetailsimprinterTurnedOff                        PrinterStatusDetails = "ImprinterTurnedOff"
	PrinterStatusDetailsimprinterTurnedOn                         PrinterStatusDetails = "ImprinterTurnedOn"
	PrinterStatusDetailsimprinterUnderTemperature                 PrinterStatusDetails = "ImprinterUnderTemperature"
	PrinterStatusDetailsimprinterUnrecoverableFailure             PrinterStatusDetails = "ImprinterUnrecoverableFailure"
	PrinterStatusDetailsimprinterUnrecoverableStorageError        PrinterStatusDetails = "ImprinterUnrecoverableStorageError"
	PrinterStatusDetailsimprinterWarmingUp                        PrinterStatusDetails = "ImprinterWarmingUp"
	PrinterStatusDetailsinputCannotFeedSizeSelected               PrinterStatusDetails = "InputCannotFeedSizeSelected"
	PrinterStatusDetailsinputManualInputRequest                   PrinterStatusDetails = "InputManualInputRequest"
	PrinterStatusDetailsinputMediaColorChange                     PrinterStatusDetails = "InputMediaColorChange"
	PrinterStatusDetailsinputMediaFormPartsChange                 PrinterStatusDetails = "InputMediaFormPartsChange"
	PrinterStatusDetailsinputMediaSizeChange                      PrinterStatusDetails = "InputMediaSizeChange"
	PrinterStatusDetailsinputMediaTrayFailure                     PrinterStatusDetails = "InputMediaTrayFailure"
	PrinterStatusDetailsinputMediaTrayFeedError                   PrinterStatusDetails = "InputMediaTrayFeedError"
	PrinterStatusDetailsinputMediaTrayJam                         PrinterStatusDetails = "InputMediaTrayJam"
	PrinterStatusDetailsinputMediaTypeChange                      PrinterStatusDetails = "InputMediaTypeChange"
	PrinterStatusDetailsinputMediaWeightChange                    PrinterStatusDetails = "InputMediaWeightChange"
	PrinterStatusDetailsinputPickRollerFailure                    PrinterStatusDetails = "InputPickRollerFailure"
	PrinterStatusDetailsinputPickRollerLifeOver                   PrinterStatusDetails = "InputPickRollerLifeOver"
	PrinterStatusDetailsinputPickRollerLifeWarn                   PrinterStatusDetails = "InputPickRollerLifeWarn"
	PrinterStatusDetailsinputPickRollerMissing                    PrinterStatusDetails = "InputPickRollerMissing"
	PrinterStatusDetailsinputTrayElevationFailure                 PrinterStatusDetails = "InputTrayElevationFailure"
	PrinterStatusDetailsinputTrayMissing                          PrinterStatusDetails = "InputTrayMissing"
	PrinterStatusDetailsinputTrayPositionFailure                  PrinterStatusDetails = "InputTrayPositionFailure"
	PrinterStatusDetailsinserterAdded                             PrinterStatusDetails = "InserterAdded"
	PrinterStatusDetailsinserterAlmostEmpty                       PrinterStatusDetails = "InserterAlmostEmpty"
	PrinterStatusDetailsinserterAlmostFull                        PrinterStatusDetails = "InserterAlmostFull"
	PrinterStatusDetailsinserterAtLimit                           PrinterStatusDetails = "InserterAtLimit"
	PrinterStatusDetailsinserterClosed                            PrinterStatusDetails = "InserterClosed"
	PrinterStatusDetailsinserterConfigurationChange               PrinterStatusDetails = "InserterConfigurationChange"
	PrinterStatusDetailsinserterCoverClosed                       PrinterStatusDetails = "InserterCoverClosed"
	PrinterStatusDetailsinserterCoverOpen                         PrinterStatusDetails = "InserterCoverOpen"
	PrinterStatusDetailsinserterEmpty                             PrinterStatusDetails = "InserterEmpty"
	PrinterStatusDetailsinserterFull                              PrinterStatusDetails = "InserterFull"
	PrinterStatusDetailsinserterInterlockClosed                   PrinterStatusDetails = "InserterInterlockClosed"
	PrinterStatusDetailsinserterInterlockOpen                     PrinterStatusDetails = "InserterInterlockOpen"
	PrinterStatusDetailsinserterJam                               PrinterStatusDetails = "InserterJam"
	PrinterStatusDetailsinserterLifeAlmostOver                    PrinterStatusDetails = "InserterLifeAlmostOver"
	PrinterStatusDetailsinserterLifeOver                          PrinterStatusDetails = "InserterLifeOver"
	PrinterStatusDetailsinserterMemoryExhausted                   PrinterStatusDetails = "InserterMemoryExhausted"
	PrinterStatusDetailsinserterMissing                           PrinterStatusDetails = "InserterMissing"
	PrinterStatusDetailsinserterMotorFailure                      PrinterStatusDetails = "InserterMotorFailure"
	PrinterStatusDetailsinserterNearLimit                         PrinterStatusDetails = "InserterNearLimit"
	PrinterStatusDetailsinserterOffline                           PrinterStatusDetails = "InserterOffline"
	PrinterStatusDetailsinserterOpened                            PrinterStatusDetails = "InserterOpened"
	PrinterStatusDetailsinserterOverTemperature                   PrinterStatusDetails = "InserterOverTemperature"
	PrinterStatusDetailsinserterPowerSaver                        PrinterStatusDetails = "InserterPowerSaver"
	PrinterStatusDetailsinserterRecoverableFailure                PrinterStatusDetails = "InserterRecoverableFailure"
	PrinterStatusDetailsinserterRecoverableStorage                PrinterStatusDetails = "InserterRecoverableStorage"
	PrinterStatusDetailsinserterRemoved                           PrinterStatusDetails = "InserterRemoved"
	PrinterStatusDetailsinserterResourceAdded                     PrinterStatusDetails = "InserterResourceAdded"
	PrinterStatusDetailsinserterResourceRemoved                   PrinterStatusDetails = "InserterResourceRemoved"
	PrinterStatusDetailsinserterThermistorFailure                 PrinterStatusDetails = "InserterThermistorFailure"
	PrinterStatusDetailsinserterTimingFailure                     PrinterStatusDetails = "InserterTimingFailure"
	PrinterStatusDetailsinserterTurnedOff                         PrinterStatusDetails = "InserterTurnedOff"
	PrinterStatusDetailsinserterTurnedOn                          PrinterStatusDetails = "InserterTurnedOn"
	PrinterStatusDetailsinserterUnderTemperature                  PrinterStatusDetails = "InserterUnderTemperature"
	PrinterStatusDetailsinserterUnrecoverableFailure              PrinterStatusDetails = "InserterUnrecoverableFailure"
	PrinterStatusDetailsinserterUnrecoverableStorageError         PrinterStatusDetails = "InserterUnrecoverableStorageError"
	PrinterStatusDetailsinserterWarmingUp                         PrinterStatusDetails = "InserterWarmingUp"
	PrinterStatusDetailsinterlockClosed                           PrinterStatusDetails = "InterlockClosed"
	PrinterStatusDetailsinterlockOpen                             PrinterStatusDetails = "InterlockOpen"
	PrinterStatusDetailsinterpreterCartridgeAdded                 PrinterStatusDetails = "InterpreterCartridgeAdded"
	PrinterStatusDetailsinterpreterCartridgeDeleted               PrinterStatusDetails = "InterpreterCartridgeDeleted"
	PrinterStatusDetailsinterpreterComplexPageEncountered         PrinterStatusDetails = "InterpreterComplexPageEncountered"
	PrinterStatusDetailsinterpreterMemoryDecrease                 PrinterStatusDetails = "InterpreterMemoryDecrease"
	PrinterStatusDetailsinterpreterMemoryIncrease                 PrinterStatusDetails = "InterpreterMemoryIncrease"
	PrinterStatusDetailsinterpreterResourceAdded                  PrinterStatusDetails = "InterpreterResourceAdded"
	PrinterStatusDetailsinterpreterResourceDeleted                PrinterStatusDetails = "InterpreterResourceDeleted"
	PrinterStatusDetailsinterpreterResourceUnavailable            PrinterStatusDetails = "InterpreterResourceUnavailable"
	PrinterStatusDetailslampAtEol                                 PrinterStatusDetails = "LampAtEol"
	PrinterStatusDetailslampFailure                               PrinterStatusDetails = "LampFailure"
	PrinterStatusDetailslampNearEol                               PrinterStatusDetails = "LampNearEol"
	PrinterStatusDetailslaserAtEol                                PrinterStatusDetails = "LaserAtEol"
	PrinterStatusDetailslaserFailure                              PrinterStatusDetails = "LaserFailure"
	PrinterStatusDetailslaserNearEol                              PrinterStatusDetails = "LaserNearEol"
	PrinterStatusDetailsmakeEnvelopeAdded                         PrinterStatusDetails = "MakeEnvelopeAdded"
	PrinterStatusDetailsmakeEnvelopeAlmostEmpty                   PrinterStatusDetails = "MakeEnvelopeAlmostEmpty"
	PrinterStatusDetailsmakeEnvelopeAlmostFull                    PrinterStatusDetails = "MakeEnvelopeAlmostFull"
	PrinterStatusDetailsmakeEnvelopeAtLimit                       PrinterStatusDetails = "MakeEnvelopeAtLimit"
	PrinterStatusDetailsmakeEnvelopeClosed                        PrinterStatusDetails = "MakeEnvelopeClosed"
	PrinterStatusDetailsmakeEnvelopeConfigurationChange           PrinterStatusDetails = "MakeEnvelopeConfigurationChange"
	PrinterStatusDetailsmakeEnvelopeCoverClosed                   PrinterStatusDetails = "MakeEnvelopeCoverClosed"
	PrinterStatusDetailsmakeEnvelopeCoverOpen                     PrinterStatusDetails = "MakeEnvelopeCoverOpen"
	PrinterStatusDetailsmakeEnvelopeEmpty                         PrinterStatusDetails = "MakeEnvelopeEmpty"
	PrinterStatusDetailsmakeEnvelopeFull                          PrinterStatusDetails = "MakeEnvelopeFull"
	PrinterStatusDetailsmakeEnvelopeInterlockClosed               PrinterStatusDetails = "MakeEnvelopeInterlockClosed"
	PrinterStatusDetailsmakeEnvelopeInterlockOpen                 PrinterStatusDetails = "MakeEnvelopeInterlockOpen"
	PrinterStatusDetailsmakeEnvelopeJam                           PrinterStatusDetails = "MakeEnvelopeJam"
	PrinterStatusDetailsmakeEnvelopeLifeAlmostOver                PrinterStatusDetails = "MakeEnvelopeLifeAlmostOver"
	PrinterStatusDetailsmakeEnvelopeLifeOver                      PrinterStatusDetails = "MakeEnvelopeLifeOver"
	PrinterStatusDetailsmakeEnvelopeMemoryExhausted               PrinterStatusDetails = "MakeEnvelopeMemoryExhausted"
	PrinterStatusDetailsmakeEnvelopeMissing                       PrinterStatusDetails = "MakeEnvelopeMissing"
	PrinterStatusDetailsmakeEnvelopeMotorFailure                  PrinterStatusDetails = "MakeEnvelopeMotorFailure"
	PrinterStatusDetailsmakeEnvelopeNearLimit                     PrinterStatusDetails = "MakeEnvelopeNearLimit"
	PrinterStatusDetailsmakeEnvelopeOffline                       PrinterStatusDetails = "MakeEnvelopeOffline"
	PrinterStatusDetailsmakeEnvelopeOpened                        PrinterStatusDetails = "MakeEnvelopeOpened"
	PrinterStatusDetailsmakeEnvelopeOverTemperature               PrinterStatusDetails = "MakeEnvelopeOverTemperature"
	PrinterStatusDetailsmakeEnvelopePowerSaver                    PrinterStatusDetails = "MakeEnvelopePowerSaver"
	PrinterStatusDetailsmakeEnvelopeRecoverableFailure            PrinterStatusDetails = "MakeEnvelopeRecoverableFailure"
	PrinterStatusDetailsmakeEnvelopeRecoverableStorage            PrinterStatusDetails = "MakeEnvelopeRecoverableStorage"
	PrinterStatusDetailsmakeEnvelopeRemoved                       PrinterStatusDetails = "MakeEnvelopeRemoved"
	PrinterStatusDetailsmakeEnvelopeResourceAdded                 PrinterStatusDetails = "MakeEnvelopeResourceAdded"
	PrinterStatusDetailsmakeEnvelopeResourceRemoved               PrinterStatusDetails = "MakeEnvelopeResourceRemoved"
	PrinterStatusDetailsmakeEnvelopeThermistorFailure             PrinterStatusDetails = "MakeEnvelopeThermistorFailure"
	PrinterStatusDetailsmakeEnvelopeTimingFailure                 PrinterStatusDetails = "MakeEnvelopeTimingFailure"
	PrinterStatusDetailsmakeEnvelopeTurnedOff                     PrinterStatusDetails = "MakeEnvelopeTurnedOff"
	PrinterStatusDetailsmakeEnvelopeTurnedOn                      PrinterStatusDetails = "MakeEnvelopeTurnedOn"
	PrinterStatusDetailsmakeEnvelopeUnderTemperature              PrinterStatusDetails = "MakeEnvelopeUnderTemperature"
	PrinterStatusDetailsmakeEnvelopeUnrecoverableFailure          PrinterStatusDetails = "MakeEnvelopeUnrecoverableFailure"
	PrinterStatusDetailsmakeEnvelopeUnrecoverableStorageError     PrinterStatusDetails = "MakeEnvelopeUnrecoverableStorageError"
	PrinterStatusDetailsmakeEnvelopeWarmingUp                     PrinterStatusDetails = "MakeEnvelopeWarmingUp"
	PrinterStatusDetailsmarkerAdjustingPrintQuality               PrinterStatusDetails = "MarkerAdjustingPrintQuality"
	PrinterStatusDetailsmarkerCleanerMissing                      PrinterStatusDetails = "MarkerCleanerMissing"
	PrinterStatusDetailsmarkerDeveloperAlmostEmpty                PrinterStatusDetails = "MarkerDeveloperAlmostEmpty"
	PrinterStatusDetailsmarkerDeveloperEmpty                      PrinterStatusDetails = "MarkerDeveloperEmpty"
	PrinterStatusDetailsmarkerDeveloperMissing                    PrinterStatusDetails = "MarkerDeveloperMissing"
	PrinterStatusDetailsmarkerFuserMissing                        PrinterStatusDetails = "MarkerFuserMissing"
	PrinterStatusDetailsmarkerFuserThermistorFailure              PrinterStatusDetails = "MarkerFuserThermistorFailure"
	PrinterStatusDetailsmarkerFuserTimingFailure                  PrinterStatusDetails = "MarkerFuserTimingFailure"
	PrinterStatusDetailsmarkerInkAlmostEmpty                      PrinterStatusDetails = "MarkerInkAlmostEmpty"
	PrinterStatusDetailsmarkerInkEmpty                            PrinterStatusDetails = "MarkerInkEmpty"
	PrinterStatusDetailsmarkerInkMissing                          PrinterStatusDetails = "MarkerInkMissing"
	PrinterStatusDetailsmarkerOpcMissing                          PrinterStatusDetails = "MarkerOpcMissing"
	PrinterStatusDetailsmarkerPrintRibbonAlmostEmpty              PrinterStatusDetails = "MarkerPrintRibbonAlmostEmpty"
	PrinterStatusDetailsmarkerPrintRibbonEmpty                    PrinterStatusDetails = "MarkerPrintRibbonEmpty"
	PrinterStatusDetailsmarkerPrintRibbonMissing                  PrinterStatusDetails = "MarkerPrintRibbonMissing"
	PrinterStatusDetailsmarkerSupplyAlmostEmpty                   PrinterStatusDetails = "MarkerSupplyAlmostEmpty"
	PrinterStatusDetailsmarkerSupplyEmpty                         PrinterStatusDetails = "MarkerSupplyEmpty"
	PrinterStatusDetailsmarkerSupplyLow                           PrinterStatusDetails = "MarkerSupplyLow"
	PrinterStatusDetailsmarkerSupplyMissing                       PrinterStatusDetails = "MarkerSupplyMissing"
	PrinterStatusDetailsmarkerTonerCartridgeMissing               PrinterStatusDetails = "MarkerTonerCartridgeMissing"
	PrinterStatusDetailsmarkerTonerMissing                        PrinterStatusDetails = "MarkerTonerMissing"
	PrinterStatusDetailsmarkerWasteAlmostFull                     PrinterStatusDetails = "MarkerWasteAlmostFull"
	PrinterStatusDetailsmarkerWasteFull                           PrinterStatusDetails = "MarkerWasteFull"
	PrinterStatusDetailsmarkerWasteInkReceptacleAlmostFull        PrinterStatusDetails = "MarkerWasteInkReceptacleAlmostFull"
	PrinterStatusDetailsmarkerWasteInkReceptacleFull              PrinterStatusDetails = "MarkerWasteInkReceptacleFull"
	PrinterStatusDetailsmarkerWasteInkReceptacleMissing           PrinterStatusDetails = "MarkerWasteInkReceptacleMissing"
	PrinterStatusDetailsmarkerWasteMissing                        PrinterStatusDetails = "MarkerWasteMissing"
	PrinterStatusDetailsmarkerWasteTonerReceptacleAlmostFull      PrinterStatusDetails = "MarkerWasteTonerReceptacleAlmostFull"
	PrinterStatusDetailsmarkerWasteTonerReceptacleFull            PrinterStatusDetails = "MarkerWasteTonerReceptacleFull"
	PrinterStatusDetailsmarkerWasteTonerReceptacleMissing         PrinterStatusDetails = "MarkerWasteTonerReceptacleMissing"
	PrinterStatusDetailsmaterialEmpty                             PrinterStatusDetails = "MaterialEmpty"
	PrinterStatusDetailsmaterialLow                               PrinterStatusDetails = "MaterialLow"
	PrinterStatusDetailsmaterialNeeded                            PrinterStatusDetails = "MaterialNeeded"
	PrinterStatusDetailsmediaDrying                               PrinterStatusDetails = "MediaDrying"
	PrinterStatusDetailsmediaEmpty                                PrinterStatusDetails = "MediaEmpty"
	PrinterStatusDetailsmediaJam                                  PrinterStatusDetails = "MediaJam"
	PrinterStatusDetailsmediaLow                                  PrinterStatusDetails = "MediaLow"
	PrinterStatusDetailsmediaNeeded                               PrinterStatusDetails = "MediaNeeded"
	PrinterStatusDetailsmediaPathCannotDuplexMediaSelected        PrinterStatusDetails = "MediaPathCannotDuplexMediaSelected"
	PrinterStatusDetailsmediaPathFailure                          PrinterStatusDetails = "MediaPathFailure"
	PrinterStatusDetailsmediaPathInputEmpty                       PrinterStatusDetails = "MediaPathInputEmpty"
	PrinterStatusDetailsmediaPathInputFeedError                   PrinterStatusDetails = "MediaPathInputFeedError"
	PrinterStatusDetailsmediaPathInputJam                         PrinterStatusDetails = "MediaPathInputJam"
	PrinterStatusDetailsmediaPathInputRequest                     PrinterStatusDetails = "MediaPathInputRequest"
	PrinterStatusDetailsmediaPathJam                              PrinterStatusDetails = "MediaPathJam"
	PrinterStatusDetailsmediaPathMediaTrayAlmostFull              PrinterStatusDetails = "MediaPathMediaTrayAlmostFull"
	PrinterStatusDetailsmediaPathMediaTrayFull                    PrinterStatusDetails = "MediaPathMediaTrayFull"
	PrinterStatusDetailsmediaPathMediaTrayMissing                 PrinterStatusDetails = "MediaPathMediaTrayMissing"
	PrinterStatusDetailsmediaPathOutputFeedError                  PrinterStatusDetails = "MediaPathOutputFeedError"
	PrinterStatusDetailsmediaPathOutputFull                       PrinterStatusDetails = "MediaPathOutputFull"
	PrinterStatusDetailsmediaPathOutputJam                        PrinterStatusDetails = "MediaPathOutputJam"
	PrinterStatusDetailsmediaPathPickRollerFailure                PrinterStatusDetails = "MediaPathPickRollerFailure"
	PrinterStatusDetailsmediaPathPickRollerLifeOver               PrinterStatusDetails = "MediaPathPickRollerLifeOver"
	PrinterStatusDetailsmediaPathPickRollerLifeWarn               PrinterStatusDetails = "MediaPathPickRollerLifeWarn"
	PrinterStatusDetailsmediaPathPickRollerMissing                PrinterStatusDetails = "MediaPathPickRollerMissing"
	PrinterStatusDetailsmotorFailure                              PrinterStatusDetails = "MotorFailure"
	PrinterStatusDetailsmovingToPaused                            PrinterStatusDetails = "MovingToPaused"
	PrinterStatusDetailsnone                                      PrinterStatusDetails = "None"
	PrinterStatusDetailsopticalPhotoConductorLifeOver             PrinterStatusDetails = "OpticalPhotoConductorLifeOver"
	PrinterStatusDetailsopticalPhotoConductorNearEndOfLife        PrinterStatusDetails = "OpticalPhotoConductorNearEndOfLife"
	PrinterStatusDetailsother                                     PrinterStatusDetails = "Other"
	PrinterStatusDetailsoutputAreaAlmostFull                      PrinterStatusDetails = "OutputAreaAlmostFull"
	PrinterStatusDetailsoutputAreaFull                            PrinterStatusDetails = "OutputAreaFull"
	PrinterStatusDetailsoutputMailboxSelectFailure                PrinterStatusDetails = "OutputMailboxSelectFailure"
	PrinterStatusDetailsoutputMediaTrayFailure                    PrinterStatusDetails = "OutputMediaTrayFailure"
	PrinterStatusDetailsoutputMediaTrayFeedError                  PrinterStatusDetails = "OutputMediaTrayFeedError"
	PrinterStatusDetailsoutputMediaTrayJam                        PrinterStatusDetails = "OutputMediaTrayJam"
	PrinterStatusDetailsoutputTrayMissing                         PrinterStatusDetails = "OutputTrayMissing"
	PrinterStatusDetailspaused                                    PrinterStatusDetails = "Paused"
	PrinterStatusDetailsperforaterAdded                           PrinterStatusDetails = "PerforaterAdded"
	PrinterStatusDetailsperforaterAlmostEmpty                     PrinterStatusDetails = "PerforaterAlmostEmpty"
	PrinterStatusDetailsperforaterAlmostFull                      PrinterStatusDetails = "PerforaterAlmostFull"
	PrinterStatusDetailsperforaterAtLimit                         PrinterStatusDetails = "PerforaterAtLimit"
	PrinterStatusDetailsperforaterClosed                          PrinterStatusDetails = "PerforaterClosed"
	PrinterStatusDetailsperforaterConfigurationChange             PrinterStatusDetails = "PerforaterConfigurationChange"
	PrinterStatusDetailsperforaterCoverClosed                     PrinterStatusDetails = "PerforaterCoverClosed"
	PrinterStatusDetailsperforaterCoverOpen                       PrinterStatusDetails = "PerforaterCoverOpen"
	PrinterStatusDetailsperforaterEmpty                           PrinterStatusDetails = "PerforaterEmpty"
	PrinterStatusDetailsperforaterFull                            PrinterStatusDetails = "PerforaterFull"
	PrinterStatusDetailsperforaterInterlockClosed                 PrinterStatusDetails = "PerforaterInterlockClosed"
	PrinterStatusDetailsperforaterInterlockOpen                   PrinterStatusDetails = "PerforaterInterlockOpen"
	PrinterStatusDetailsperforaterJam                             PrinterStatusDetails = "PerforaterJam"
	PrinterStatusDetailsperforaterLifeAlmostOver                  PrinterStatusDetails = "PerforaterLifeAlmostOver"
	PrinterStatusDetailsperforaterLifeOver                        PrinterStatusDetails = "PerforaterLifeOver"
	PrinterStatusDetailsperforaterMemoryExhausted                 PrinterStatusDetails = "PerforaterMemoryExhausted"
	PrinterStatusDetailsperforaterMissing                         PrinterStatusDetails = "PerforaterMissing"
	PrinterStatusDetailsperforaterMotorFailure                    PrinterStatusDetails = "PerforaterMotorFailure"
	PrinterStatusDetailsperforaterNearLimit                       PrinterStatusDetails = "PerforaterNearLimit"
	PrinterStatusDetailsperforaterOffline                         PrinterStatusDetails = "PerforaterOffline"
	PrinterStatusDetailsperforaterOpened                          PrinterStatusDetails = "PerforaterOpened"
	PrinterStatusDetailsperforaterOverTemperature                 PrinterStatusDetails = "PerforaterOverTemperature"
	PrinterStatusDetailsperforaterPowerSaver                      PrinterStatusDetails = "PerforaterPowerSaver"
	PrinterStatusDetailsperforaterRecoverableFailure              PrinterStatusDetails = "PerforaterRecoverableFailure"
	PrinterStatusDetailsperforaterRecoverableStorage              PrinterStatusDetails = "PerforaterRecoverableStorage"
	PrinterStatusDetailsperforaterRemoved                         PrinterStatusDetails = "PerforaterRemoved"
	PrinterStatusDetailsperforaterResourceAdded                   PrinterStatusDetails = "PerforaterResourceAdded"
	PrinterStatusDetailsperforaterResourceRemoved                 PrinterStatusDetails = "PerforaterResourceRemoved"
	PrinterStatusDetailsperforaterThermistorFailure               PrinterStatusDetails = "PerforaterThermistorFailure"
	PrinterStatusDetailsperforaterTimingFailure                   PrinterStatusDetails = "PerforaterTimingFailure"
	PrinterStatusDetailsperforaterTurnedOff                       PrinterStatusDetails = "PerforaterTurnedOff"
	PrinterStatusDetailsperforaterTurnedOn                        PrinterStatusDetails = "PerforaterTurnedOn"
	PrinterStatusDetailsperforaterUnderTemperature                PrinterStatusDetails = "PerforaterUnderTemperature"
	PrinterStatusDetailsperforaterUnrecoverableFailure            PrinterStatusDetails = "PerforaterUnrecoverableFailure"
	PrinterStatusDetailsperforaterUnrecoverableStorageError       PrinterStatusDetails = "PerforaterUnrecoverableStorageError"
	PrinterStatusDetailsperforaterWarmingUp                       PrinterStatusDetails = "PerforaterWarmingUp"
	PrinterStatusDetailsplatformCooling                           PrinterStatusDetails = "PlatformCooling"
	PrinterStatusDetailsplatformFailure                           PrinterStatusDetails = "PlatformFailure"
	PrinterStatusDetailsplatformHeating                           PrinterStatusDetails = "PlatformHeating"
	PrinterStatusDetailsplatformTemperatureHigh                   PrinterStatusDetails = "PlatformTemperatureHigh"
	PrinterStatusDetailsplatformTemperatureLow                    PrinterStatusDetails = "PlatformTemperatureLow"
	PrinterStatusDetailspowerDown                                 PrinterStatusDetails = "PowerDown"
	PrinterStatusDetailspowerUp                                   PrinterStatusDetails = "PowerUp"
	PrinterStatusDetailsprinterManualReset                        PrinterStatusDetails = "PrinterManualReset"
	PrinterStatusDetailsprinterNmsReset                           PrinterStatusDetails = "PrinterNmsReset"
	PrinterStatusDetailsprinterReadyToPrint                       PrinterStatusDetails = "PrinterReadyToPrint"
	PrinterStatusDetailspuncherAdded                              PrinterStatusDetails = "PuncherAdded"
	PrinterStatusDetailspuncherAlmostEmpty                        PrinterStatusDetails = "PuncherAlmostEmpty"
	PrinterStatusDetailspuncherAlmostFull                         PrinterStatusDetails = "PuncherAlmostFull"
	PrinterStatusDetailspuncherAtLimit                            PrinterStatusDetails = "PuncherAtLimit"
	PrinterStatusDetailspuncherClosed                             PrinterStatusDetails = "PuncherClosed"
	PrinterStatusDetailspuncherConfigurationChange                PrinterStatusDetails = "PuncherConfigurationChange"
	PrinterStatusDetailspuncherCoverClosed                        PrinterStatusDetails = "PuncherCoverClosed"
	PrinterStatusDetailspuncherCoverOpen                          PrinterStatusDetails = "PuncherCoverOpen"
	PrinterStatusDetailspuncherEmpty                              PrinterStatusDetails = "PuncherEmpty"
	PrinterStatusDetailspuncherFull                               PrinterStatusDetails = "PuncherFull"
	PrinterStatusDetailspuncherInterlockClosed                    PrinterStatusDetails = "PuncherInterlockClosed"
	PrinterStatusDetailspuncherInterlockOpen                      PrinterStatusDetails = "PuncherInterlockOpen"
	PrinterStatusDetailspuncherJam                                PrinterStatusDetails = "PuncherJam"
	PrinterStatusDetailspuncherLifeAlmostOver                     PrinterStatusDetails = "PuncherLifeAlmostOver"
	PrinterStatusDetailspuncherLifeOver                           PrinterStatusDetails = "PuncherLifeOver"
	PrinterStatusDetailspuncherMemoryExhausted                    PrinterStatusDetails = "PuncherMemoryExhausted"
	PrinterStatusDetailspuncherMissing                            PrinterStatusDetails = "PuncherMissing"
	PrinterStatusDetailspuncherMotorFailure                       PrinterStatusDetails = "PuncherMotorFailure"
	PrinterStatusDetailspuncherNearLimit                          PrinterStatusDetails = "PuncherNearLimit"
	PrinterStatusDetailspuncherOffline                            PrinterStatusDetails = "PuncherOffline"
	PrinterStatusDetailspuncherOpened                             PrinterStatusDetails = "PuncherOpened"
	PrinterStatusDetailspuncherOverTemperature                    PrinterStatusDetails = "PuncherOverTemperature"
	PrinterStatusDetailspuncherPowerSaver                         PrinterStatusDetails = "PuncherPowerSaver"
	PrinterStatusDetailspuncherRecoverableFailure                 PrinterStatusDetails = "PuncherRecoverableFailure"
	PrinterStatusDetailspuncherRecoverableStorage                 PrinterStatusDetails = "PuncherRecoverableStorage"
	PrinterStatusDetailspuncherRemoved                            PrinterStatusDetails = "PuncherRemoved"
	PrinterStatusDetailspuncherResourceAdded                      PrinterStatusDetails = "PuncherResourceAdded"
	PrinterStatusDetailspuncherResourceRemoved                    PrinterStatusDetails = "PuncherResourceRemoved"
	PrinterStatusDetailspuncherThermistorFailure                  PrinterStatusDetails = "PuncherThermistorFailure"
	PrinterStatusDetailspuncherTimingFailure                      PrinterStatusDetails = "PuncherTimingFailure"
	PrinterStatusDetailspuncherTurnedOff                          PrinterStatusDetails = "PuncherTurnedOff"
	PrinterStatusDetailspuncherTurnedOn                           PrinterStatusDetails = "PuncherTurnedOn"
	PrinterStatusDetailspuncherUnderTemperature                   PrinterStatusDetails = "PuncherUnderTemperature"
	PrinterStatusDetailspuncherUnrecoverableFailure               PrinterStatusDetails = "PuncherUnrecoverableFailure"
	PrinterStatusDetailspuncherUnrecoverableStorageError          PrinterStatusDetails = "PuncherUnrecoverableStorageError"
	PrinterStatusDetailspuncherWarmingUp                          PrinterStatusDetails = "PuncherWarmingUp"
	PrinterStatusDetailsresuming                                  PrinterStatusDetails = "Resuming"
	PrinterStatusDetailsscanMediaPathFailure                      PrinterStatusDetails = "ScanMediaPathFailure"
	PrinterStatusDetailsscanMediaPathInputEmpty                   PrinterStatusDetails = "ScanMediaPathInputEmpty"
	PrinterStatusDetailsscanMediaPathInputFeedError               PrinterStatusDetails = "ScanMediaPathInputFeedError"
	PrinterStatusDetailsscanMediaPathInputJam                     PrinterStatusDetails = "ScanMediaPathInputJam"
	PrinterStatusDetailsscanMediaPathInputRequest                 PrinterStatusDetails = "ScanMediaPathInputRequest"
	PrinterStatusDetailsscanMediaPathJam                          PrinterStatusDetails = "ScanMediaPathJam"
	PrinterStatusDetailsscanMediaPathOutputFeedError              PrinterStatusDetails = "ScanMediaPathOutputFeedError"
	PrinterStatusDetailsscanMediaPathOutputFull                   PrinterStatusDetails = "ScanMediaPathOutputFull"
	PrinterStatusDetailsscanMediaPathOutputJam                    PrinterStatusDetails = "ScanMediaPathOutputJam"
	PrinterStatusDetailsscanMediaPathPickRollerFailure            PrinterStatusDetails = "ScanMediaPathPickRollerFailure"
	PrinterStatusDetailsscanMediaPathPickRollerLifeOver           PrinterStatusDetails = "ScanMediaPathPickRollerLifeOver"
	PrinterStatusDetailsscanMediaPathPickRollerLifeWarn           PrinterStatusDetails = "ScanMediaPathPickRollerLifeWarn"
	PrinterStatusDetailsscanMediaPathPickRollerMissing            PrinterStatusDetails = "ScanMediaPathPickRollerMissing"
	PrinterStatusDetailsscanMediaPathTrayAlmostFull               PrinterStatusDetails = "ScanMediaPathTrayAlmostFull"
	PrinterStatusDetailsscanMediaPathTrayFull                     PrinterStatusDetails = "ScanMediaPathTrayFull"
	PrinterStatusDetailsscanMediaPathTrayMissing                  PrinterStatusDetails = "ScanMediaPathTrayMissing"
	PrinterStatusDetailsscannerLightFailure                       PrinterStatusDetails = "ScannerLightFailure"
	PrinterStatusDetailsscannerLightLifeAlmostOver                PrinterStatusDetails = "ScannerLightLifeAlmostOver"
	PrinterStatusDetailsscannerLightLifeOver                      PrinterStatusDetails = "ScannerLightLifeOver"
	PrinterStatusDetailsscannerLightMissing                       PrinterStatusDetails = "ScannerLightMissing"
	PrinterStatusDetailsscannerSensorFailure                      PrinterStatusDetails = "ScannerSensorFailure"
	PrinterStatusDetailsscannerSensorLifeAlmostOver               PrinterStatusDetails = "ScannerSensorLifeAlmostOver"
	PrinterStatusDetailsscannerSensorLifeOver                     PrinterStatusDetails = "ScannerSensorLifeOver"
	PrinterStatusDetailsscannerSensorMissing                      PrinterStatusDetails = "ScannerSensorMissing"
	PrinterStatusDetailsseparationCutterAdded                     PrinterStatusDetails = "SeparationCutterAdded"
	PrinterStatusDetailsseparationCutterAlmostEmpty               PrinterStatusDetails = "SeparationCutterAlmostEmpty"
	PrinterStatusDetailsseparationCutterAlmostFull                PrinterStatusDetails = "SeparationCutterAlmostFull"
	PrinterStatusDetailsseparationCutterAtLimit                   PrinterStatusDetails = "SeparationCutterAtLimit"
	PrinterStatusDetailsseparationCutterClosed                    PrinterStatusDetails = "SeparationCutterClosed"
	PrinterStatusDetailsseparationCutterConfigurationChange       PrinterStatusDetails = "SeparationCutterConfigurationChange"
	PrinterStatusDetailsseparationCutterCoverClosed               PrinterStatusDetails = "SeparationCutterCoverClosed"
	PrinterStatusDetailsseparationCutterCoverOpen                 PrinterStatusDetails = "SeparationCutterCoverOpen"
	PrinterStatusDetailsseparationCutterEmpty                     PrinterStatusDetails = "SeparationCutterEmpty"
	PrinterStatusDetailsseparationCutterFull                      PrinterStatusDetails = "SeparationCutterFull"
	PrinterStatusDetailsseparationCutterInterlockClosed           PrinterStatusDetails = "SeparationCutterInterlockClosed"
	PrinterStatusDetailsseparationCutterInterlockOpen             PrinterStatusDetails = "SeparationCutterInterlockOpen"
	PrinterStatusDetailsseparationCutterJam                       PrinterStatusDetails = "SeparationCutterJam"
	PrinterStatusDetailsseparationCutterLifeAlmostOver            PrinterStatusDetails = "SeparationCutterLifeAlmostOver"
	PrinterStatusDetailsseparationCutterLifeOver                  PrinterStatusDetails = "SeparationCutterLifeOver"
	PrinterStatusDetailsseparationCutterMemoryExhausted           PrinterStatusDetails = "SeparationCutterMemoryExhausted"
	PrinterStatusDetailsseparationCutterMissing                   PrinterStatusDetails = "SeparationCutterMissing"
	PrinterStatusDetailsseparationCutterMotorFailure              PrinterStatusDetails = "SeparationCutterMotorFailure"
	PrinterStatusDetailsseparationCutterNearLimit                 PrinterStatusDetails = "SeparationCutterNearLimit"
	PrinterStatusDetailsseparationCutterOffline                   PrinterStatusDetails = "SeparationCutterOffline"
	PrinterStatusDetailsseparationCutterOpened                    PrinterStatusDetails = "SeparationCutterOpened"
	PrinterStatusDetailsseparationCutterOverTemperature           PrinterStatusDetails = "SeparationCutterOverTemperature"
	PrinterStatusDetailsseparationCutterPowerSaver                PrinterStatusDetails = "SeparationCutterPowerSaver"
	PrinterStatusDetailsseparationCutterRecoverableFailure        PrinterStatusDetails = "SeparationCutterRecoverableFailure"
	PrinterStatusDetailsseparationCutterRecoverableStorage        PrinterStatusDetails = "SeparationCutterRecoverableStorage"
	PrinterStatusDetailsseparationCutterRemoved                   PrinterStatusDetails = "SeparationCutterRemoved"
	PrinterStatusDetailsseparationCutterResourceAdded             PrinterStatusDetails = "SeparationCutterResourceAdded"
	PrinterStatusDetailsseparationCutterResourceRemoved           PrinterStatusDetails = "SeparationCutterResourceRemoved"
	PrinterStatusDetailsseparationCutterThermistorFailure         PrinterStatusDetails = "SeparationCutterThermistorFailure"
	PrinterStatusDetailsseparationCutterTimingFailure             PrinterStatusDetails = "SeparationCutterTimingFailure"
	PrinterStatusDetailsseparationCutterTurnedOff                 PrinterStatusDetails = "SeparationCutterTurnedOff"
	PrinterStatusDetailsseparationCutterTurnedOn                  PrinterStatusDetails = "SeparationCutterTurnedOn"
	PrinterStatusDetailsseparationCutterUnderTemperature          PrinterStatusDetails = "SeparationCutterUnderTemperature"
	PrinterStatusDetailsseparationCutterUnrecoverableFailure      PrinterStatusDetails = "SeparationCutterUnrecoverableFailure"
	PrinterStatusDetailsseparationCutterUnrecoverableStorageError PrinterStatusDetails = "SeparationCutterUnrecoverableStorageError"
	PrinterStatusDetailsseparationCutterWarmingUp                 PrinterStatusDetails = "SeparationCutterWarmingUp"
	PrinterStatusDetailssheetRotatorAdded                         PrinterStatusDetails = "SheetRotatorAdded"
	PrinterStatusDetailssheetRotatorAlmostEmpty                   PrinterStatusDetails = "SheetRotatorAlmostEmpty"
	PrinterStatusDetailssheetRotatorAlmostFull                    PrinterStatusDetails = "SheetRotatorAlmostFull"
	PrinterStatusDetailssheetRotatorAtLimit                       PrinterStatusDetails = "SheetRotatorAtLimit"
	PrinterStatusDetailssheetRotatorClosed                        PrinterStatusDetails = "SheetRotatorClosed"
	PrinterStatusDetailssheetRotatorConfigurationChange           PrinterStatusDetails = "SheetRotatorConfigurationChange"
	PrinterStatusDetailssheetRotatorCoverClosed                   PrinterStatusDetails = "SheetRotatorCoverClosed"
	PrinterStatusDetailssheetRotatorCoverOpen                     PrinterStatusDetails = "SheetRotatorCoverOpen"
	PrinterStatusDetailssheetRotatorEmpty                         PrinterStatusDetails = "SheetRotatorEmpty"
	PrinterStatusDetailssheetRotatorFull                          PrinterStatusDetails = "SheetRotatorFull"
	PrinterStatusDetailssheetRotatorInterlockClosed               PrinterStatusDetails = "SheetRotatorInterlockClosed"
	PrinterStatusDetailssheetRotatorInterlockOpen                 PrinterStatusDetails = "SheetRotatorInterlockOpen"
	PrinterStatusDetailssheetRotatorJam                           PrinterStatusDetails = "SheetRotatorJam"
	PrinterStatusDetailssheetRotatorLifeAlmostOver                PrinterStatusDetails = "SheetRotatorLifeAlmostOver"
	PrinterStatusDetailssheetRotatorLifeOver                      PrinterStatusDetails = "SheetRotatorLifeOver"
	PrinterStatusDetailssheetRotatorMemoryExhausted               PrinterStatusDetails = "SheetRotatorMemoryExhausted"
	PrinterStatusDetailssheetRotatorMissing                       PrinterStatusDetails = "SheetRotatorMissing"
	PrinterStatusDetailssheetRotatorMotorFailure                  PrinterStatusDetails = "SheetRotatorMotorFailure"
	PrinterStatusDetailssheetRotatorNearLimit                     PrinterStatusDetails = "SheetRotatorNearLimit"
	PrinterStatusDetailssheetRotatorOffline                       PrinterStatusDetails = "SheetRotatorOffline"
	PrinterStatusDetailssheetRotatorOpened                        PrinterStatusDetails = "SheetRotatorOpened"
	PrinterStatusDetailssheetRotatorOverTemperature               PrinterStatusDetails = "SheetRotatorOverTemperature"
	PrinterStatusDetailssheetRotatorPowerSaver                    PrinterStatusDetails = "SheetRotatorPowerSaver"
	PrinterStatusDetailssheetRotatorRecoverableFailure            PrinterStatusDetails = "SheetRotatorRecoverableFailure"
	PrinterStatusDetailssheetRotatorRecoverableStorage            PrinterStatusDetails = "SheetRotatorRecoverableStorage"
	PrinterStatusDetailssheetRotatorRemoved                       PrinterStatusDetails = "SheetRotatorRemoved"
	PrinterStatusDetailssheetRotatorResourceAdded                 PrinterStatusDetails = "SheetRotatorResourceAdded"
	PrinterStatusDetailssheetRotatorResourceRemoved               PrinterStatusDetails = "SheetRotatorResourceRemoved"
	PrinterStatusDetailssheetRotatorThermistorFailure             PrinterStatusDetails = "SheetRotatorThermistorFailure"
	PrinterStatusDetailssheetRotatorTimingFailure                 PrinterStatusDetails = "SheetRotatorTimingFailure"
	PrinterStatusDetailssheetRotatorTurnedOff                     PrinterStatusDetails = "SheetRotatorTurnedOff"
	PrinterStatusDetailssheetRotatorTurnedOn                      PrinterStatusDetails = "SheetRotatorTurnedOn"
	PrinterStatusDetailssheetRotatorUnderTemperature              PrinterStatusDetails = "SheetRotatorUnderTemperature"
	PrinterStatusDetailssheetRotatorUnrecoverableFailure          PrinterStatusDetails = "SheetRotatorUnrecoverableFailure"
	PrinterStatusDetailssheetRotatorUnrecoverableStorageError     PrinterStatusDetails = "SheetRotatorUnrecoverableStorageError"
	PrinterStatusDetailssheetRotatorWarmingUp                     PrinterStatusDetails = "SheetRotatorWarmingUp"
	PrinterStatusDetailsshutdown                                  PrinterStatusDetails = "Shutdown"
	PrinterStatusDetailsslitterAdded                              PrinterStatusDetails = "SlitterAdded"
	PrinterStatusDetailsslitterAlmostEmpty                        PrinterStatusDetails = "SlitterAlmostEmpty"
	PrinterStatusDetailsslitterAlmostFull                         PrinterStatusDetails = "SlitterAlmostFull"
	PrinterStatusDetailsslitterAtLimit                            PrinterStatusDetails = "SlitterAtLimit"
	PrinterStatusDetailsslitterClosed                             PrinterStatusDetails = "SlitterClosed"
	PrinterStatusDetailsslitterConfigurationChange                PrinterStatusDetails = "SlitterConfigurationChange"
	PrinterStatusDetailsslitterCoverClosed                        PrinterStatusDetails = "SlitterCoverClosed"
	PrinterStatusDetailsslitterCoverOpen                          PrinterStatusDetails = "SlitterCoverOpen"
	PrinterStatusDetailsslitterEmpty                              PrinterStatusDetails = "SlitterEmpty"
	PrinterStatusDetailsslitterFull                               PrinterStatusDetails = "SlitterFull"
	PrinterStatusDetailsslitterInterlockClosed                    PrinterStatusDetails = "SlitterInterlockClosed"
	PrinterStatusDetailsslitterInterlockOpen                      PrinterStatusDetails = "SlitterInterlockOpen"
	PrinterStatusDetailsslitterJam                                PrinterStatusDetails = "SlitterJam"
	PrinterStatusDetailsslitterLifeAlmostOver                     PrinterStatusDetails = "SlitterLifeAlmostOver"
	PrinterStatusDetailsslitterLifeOver                           PrinterStatusDetails = "SlitterLifeOver"
	PrinterStatusDetailsslitterMemoryExhausted                    PrinterStatusDetails = "SlitterMemoryExhausted"
	PrinterStatusDetailsslitterMissing                            PrinterStatusDetails = "SlitterMissing"
	PrinterStatusDetailsslitterMotorFailure                       PrinterStatusDetails = "SlitterMotorFailure"
	PrinterStatusDetailsslitterNearLimit                          PrinterStatusDetails = "SlitterNearLimit"
	PrinterStatusDetailsslitterOffline                            PrinterStatusDetails = "SlitterOffline"
	PrinterStatusDetailsslitterOpened                             PrinterStatusDetails = "SlitterOpened"
	PrinterStatusDetailsslitterOverTemperature                    PrinterStatusDetails = "SlitterOverTemperature"
	PrinterStatusDetailsslitterPowerSaver                         PrinterStatusDetails = "SlitterPowerSaver"
	PrinterStatusDetailsslitterRecoverableFailure                 PrinterStatusDetails = "SlitterRecoverableFailure"
	PrinterStatusDetailsslitterRecoverableStorage                 PrinterStatusDetails = "SlitterRecoverableStorage"
	PrinterStatusDetailsslitterRemoved                            PrinterStatusDetails = "SlitterRemoved"
	PrinterStatusDetailsslitterResourceAdded                      PrinterStatusDetails = "SlitterResourceAdded"
	PrinterStatusDetailsslitterResourceRemoved                    PrinterStatusDetails = "SlitterResourceRemoved"
	PrinterStatusDetailsslitterThermistorFailure                  PrinterStatusDetails = "SlitterThermistorFailure"
	PrinterStatusDetailsslitterTimingFailure                      PrinterStatusDetails = "SlitterTimingFailure"
	PrinterStatusDetailsslitterTurnedOff                          PrinterStatusDetails = "SlitterTurnedOff"
	PrinterStatusDetailsslitterTurnedOn                           PrinterStatusDetails = "SlitterTurnedOn"
	PrinterStatusDetailsslitterUnderTemperature                   PrinterStatusDetails = "SlitterUnderTemperature"
	PrinterStatusDetailsslitterUnrecoverableFailure               PrinterStatusDetails = "SlitterUnrecoverableFailure"
	PrinterStatusDetailsslitterUnrecoverableStorageError          PrinterStatusDetails = "SlitterUnrecoverableStorageError"
	PrinterStatusDetailsslitterWarmingUp                          PrinterStatusDetails = "SlitterWarmingUp"
	PrinterStatusDetailsspoolAreaFull                             PrinterStatusDetails = "SpoolAreaFull"
	PrinterStatusDetailsstackerAdded                              PrinterStatusDetails = "StackerAdded"
	PrinterStatusDetailsstackerAlmostEmpty                        PrinterStatusDetails = "StackerAlmostEmpty"
	PrinterStatusDetailsstackerAlmostFull                         PrinterStatusDetails = "StackerAlmostFull"
	PrinterStatusDetailsstackerAtLimit                            PrinterStatusDetails = "StackerAtLimit"
	PrinterStatusDetailsstackerClosed                             PrinterStatusDetails = "StackerClosed"
	PrinterStatusDetailsstackerConfigurationChange                PrinterStatusDetails = "StackerConfigurationChange"
	PrinterStatusDetailsstackerCoverClosed                        PrinterStatusDetails = "StackerCoverClosed"
	PrinterStatusDetailsstackerCoverOpen                          PrinterStatusDetails = "StackerCoverOpen"
	PrinterStatusDetailsstackerEmpty                              PrinterStatusDetails = "StackerEmpty"
	PrinterStatusDetailsstackerFull                               PrinterStatusDetails = "StackerFull"
	PrinterStatusDetailsstackerInterlockClosed                    PrinterStatusDetails = "StackerInterlockClosed"
	PrinterStatusDetailsstackerInterlockOpen                      PrinterStatusDetails = "StackerInterlockOpen"
	PrinterStatusDetailsstackerJam                                PrinterStatusDetails = "StackerJam"
	PrinterStatusDetailsstackerLifeAlmostOver                     PrinterStatusDetails = "StackerLifeAlmostOver"
	PrinterStatusDetailsstackerLifeOver                           PrinterStatusDetails = "StackerLifeOver"
	PrinterStatusDetailsstackerMemoryExhausted                    PrinterStatusDetails = "StackerMemoryExhausted"
	PrinterStatusDetailsstackerMissing                            PrinterStatusDetails = "StackerMissing"
	PrinterStatusDetailsstackerMotorFailure                       PrinterStatusDetails = "StackerMotorFailure"
	PrinterStatusDetailsstackerNearLimit                          PrinterStatusDetails = "StackerNearLimit"
	PrinterStatusDetailsstackerOffline                            PrinterStatusDetails = "StackerOffline"
	PrinterStatusDetailsstackerOpened                             PrinterStatusDetails = "StackerOpened"
	PrinterStatusDetailsstackerOverTemperature                    PrinterStatusDetails = "StackerOverTemperature"
	PrinterStatusDetailsstackerPowerSaver                         PrinterStatusDetails = "StackerPowerSaver"
	PrinterStatusDetailsstackerRecoverableFailure                 PrinterStatusDetails = "StackerRecoverableFailure"
	PrinterStatusDetailsstackerRecoverableStorage                 PrinterStatusDetails = "StackerRecoverableStorage"
	PrinterStatusDetailsstackerRemoved                            PrinterStatusDetails = "StackerRemoved"
	PrinterStatusDetailsstackerResourceAdded                      PrinterStatusDetails = "StackerResourceAdded"
	PrinterStatusDetailsstackerResourceRemoved                    PrinterStatusDetails = "StackerResourceRemoved"
	PrinterStatusDetailsstackerThermistorFailure                  PrinterStatusDetails = "StackerThermistorFailure"
	PrinterStatusDetailsstackerTimingFailure                      PrinterStatusDetails = "StackerTimingFailure"
	PrinterStatusDetailsstackerTurnedOff                          PrinterStatusDetails = "StackerTurnedOff"
	PrinterStatusDetailsstackerTurnedOn                           PrinterStatusDetails = "StackerTurnedOn"
	PrinterStatusDetailsstackerUnderTemperature                   PrinterStatusDetails = "StackerUnderTemperature"
	PrinterStatusDetailsstackerUnrecoverableFailure               PrinterStatusDetails = "StackerUnrecoverableFailure"
	PrinterStatusDetailsstackerUnrecoverableStorageError          PrinterStatusDetails = "StackerUnrecoverableStorageError"
	PrinterStatusDetailsstackerWarmingUp                          PrinterStatusDetails = "StackerWarmingUp"
	PrinterStatusDetailsstandby                                   PrinterStatusDetails = "Standby"
	PrinterStatusDetailsstaplerAdded                              PrinterStatusDetails = "StaplerAdded"
	PrinterStatusDetailsstaplerAlmostEmpty                        PrinterStatusDetails = "StaplerAlmostEmpty"
	PrinterStatusDetailsstaplerAlmostFull                         PrinterStatusDetails = "StaplerAlmostFull"
	PrinterStatusDetailsstaplerAtLimit                            PrinterStatusDetails = "StaplerAtLimit"
	PrinterStatusDetailsstaplerClosed                             PrinterStatusDetails = "StaplerClosed"
	PrinterStatusDetailsstaplerConfigurationChange                PrinterStatusDetails = "StaplerConfigurationChange"
	PrinterStatusDetailsstaplerCoverClosed                        PrinterStatusDetails = "StaplerCoverClosed"
	PrinterStatusDetailsstaplerCoverOpen                          PrinterStatusDetails = "StaplerCoverOpen"
	PrinterStatusDetailsstaplerEmpty                              PrinterStatusDetails = "StaplerEmpty"
	PrinterStatusDetailsstaplerFull                               PrinterStatusDetails = "StaplerFull"
	PrinterStatusDetailsstaplerInterlockClosed                    PrinterStatusDetails = "StaplerInterlockClosed"
	PrinterStatusDetailsstaplerInterlockOpen                      PrinterStatusDetails = "StaplerInterlockOpen"
	PrinterStatusDetailsstaplerJam                                PrinterStatusDetails = "StaplerJam"
	PrinterStatusDetailsstaplerLifeAlmostOver                     PrinterStatusDetails = "StaplerLifeAlmostOver"
	PrinterStatusDetailsstaplerLifeOver                           PrinterStatusDetails = "StaplerLifeOver"
	PrinterStatusDetailsstaplerMemoryExhausted                    PrinterStatusDetails = "StaplerMemoryExhausted"
	PrinterStatusDetailsstaplerMissing                            PrinterStatusDetails = "StaplerMissing"
	PrinterStatusDetailsstaplerMotorFailure                       PrinterStatusDetails = "StaplerMotorFailure"
	PrinterStatusDetailsstaplerNearLimit                          PrinterStatusDetails = "StaplerNearLimit"
	PrinterStatusDetailsstaplerOffline                            PrinterStatusDetails = "StaplerOffline"
	PrinterStatusDetailsstaplerOpened                             PrinterStatusDetails = "StaplerOpened"
	PrinterStatusDetailsstaplerOverTemperature                    PrinterStatusDetails = "StaplerOverTemperature"
	PrinterStatusDetailsstaplerPowerSaver                         PrinterStatusDetails = "StaplerPowerSaver"
	PrinterStatusDetailsstaplerRecoverableFailure                 PrinterStatusDetails = "StaplerRecoverableFailure"
	PrinterStatusDetailsstaplerRecoverableStorage                 PrinterStatusDetails = "StaplerRecoverableStorage"
	PrinterStatusDetailsstaplerRemoved                            PrinterStatusDetails = "StaplerRemoved"
	PrinterStatusDetailsstaplerResourceAdded                      PrinterStatusDetails = "StaplerResourceAdded"
	PrinterStatusDetailsstaplerResourceRemoved                    PrinterStatusDetails = "StaplerResourceRemoved"
	PrinterStatusDetailsstaplerThermistorFailure                  PrinterStatusDetails = "StaplerThermistorFailure"
	PrinterStatusDetailsstaplerTimingFailure                      PrinterStatusDetails = "StaplerTimingFailure"
	PrinterStatusDetailsstaplerTurnedOff                          PrinterStatusDetails = "StaplerTurnedOff"
	PrinterStatusDetailsstaplerTurnedOn                           PrinterStatusDetails = "StaplerTurnedOn"
	PrinterStatusDetailsstaplerUnderTemperature                   PrinterStatusDetails = "StaplerUnderTemperature"
	PrinterStatusDetailsstaplerUnrecoverableFailure               PrinterStatusDetails = "StaplerUnrecoverableFailure"
	PrinterStatusDetailsstaplerUnrecoverableStorageError          PrinterStatusDetails = "StaplerUnrecoverableStorageError"
	PrinterStatusDetailsstaplerWarmingUp                          PrinterStatusDetails = "StaplerWarmingUp"
	PrinterStatusDetailsstitcherAdded                             PrinterStatusDetails = "StitcherAdded"
	PrinterStatusDetailsstitcherAlmostEmpty                       PrinterStatusDetails = "StitcherAlmostEmpty"
	PrinterStatusDetailsstitcherAlmostFull                        PrinterStatusDetails = "StitcherAlmostFull"
	PrinterStatusDetailsstitcherAtLimit                           PrinterStatusDetails = "StitcherAtLimit"
	PrinterStatusDetailsstitcherClosed                            PrinterStatusDetails = "StitcherClosed"
	PrinterStatusDetailsstitcherConfigurationChange               PrinterStatusDetails = "StitcherConfigurationChange"
	PrinterStatusDetailsstitcherCoverClosed                       PrinterStatusDetails = "StitcherCoverClosed"
	PrinterStatusDetailsstitcherCoverOpen                         PrinterStatusDetails = "StitcherCoverOpen"
	PrinterStatusDetailsstitcherEmpty                             PrinterStatusDetails = "StitcherEmpty"
	PrinterStatusDetailsstitcherFull                              PrinterStatusDetails = "StitcherFull"
	PrinterStatusDetailsstitcherInterlockClosed                   PrinterStatusDetails = "StitcherInterlockClosed"
	PrinterStatusDetailsstitcherInterlockOpen                     PrinterStatusDetails = "StitcherInterlockOpen"
	PrinterStatusDetailsstitcherJam                               PrinterStatusDetails = "StitcherJam"
	PrinterStatusDetailsstitcherLifeAlmostOver                    PrinterStatusDetails = "StitcherLifeAlmostOver"
	PrinterStatusDetailsstitcherLifeOver                          PrinterStatusDetails = "StitcherLifeOver"
	PrinterStatusDetailsstitcherMemoryExhausted                   PrinterStatusDetails = "StitcherMemoryExhausted"
	PrinterStatusDetailsstitcherMissing                           PrinterStatusDetails = "StitcherMissing"
	PrinterStatusDetailsstitcherMotorFailure                      PrinterStatusDetails = "StitcherMotorFailure"
	PrinterStatusDetailsstitcherNearLimit                         PrinterStatusDetails = "StitcherNearLimit"
	PrinterStatusDetailsstitcherOffline                           PrinterStatusDetails = "StitcherOffline"
	PrinterStatusDetailsstitcherOpened                            PrinterStatusDetails = "StitcherOpened"
	PrinterStatusDetailsstitcherOverTemperature                   PrinterStatusDetails = "StitcherOverTemperature"
	PrinterStatusDetailsstitcherPowerSaver                        PrinterStatusDetails = "StitcherPowerSaver"
	PrinterStatusDetailsstitcherRecoverableFailure                PrinterStatusDetails = "StitcherRecoverableFailure"
	PrinterStatusDetailsstitcherRecoverableStorage                PrinterStatusDetails = "StitcherRecoverableStorage"
	PrinterStatusDetailsstitcherRemoved                           PrinterStatusDetails = "StitcherRemoved"
	PrinterStatusDetailsstitcherResourceAdded                     PrinterStatusDetails = "StitcherResourceAdded"
	PrinterStatusDetailsstitcherResourceRemoved                   PrinterStatusDetails = "StitcherResourceRemoved"
	PrinterStatusDetailsstitcherThermistorFailure                 PrinterStatusDetails = "StitcherThermistorFailure"
	PrinterStatusDetailsstitcherTimingFailure                     PrinterStatusDetails = "StitcherTimingFailure"
	PrinterStatusDetailsstitcherTurnedOff                         PrinterStatusDetails = "StitcherTurnedOff"
	PrinterStatusDetailsstitcherTurnedOn                          PrinterStatusDetails = "StitcherTurnedOn"
	PrinterStatusDetailsstitcherUnderTemperature                  PrinterStatusDetails = "StitcherUnderTemperature"
	PrinterStatusDetailsstitcherUnrecoverableFailure              PrinterStatusDetails = "StitcherUnrecoverableFailure"
	PrinterStatusDetailsstitcherUnrecoverableStorageError         PrinterStatusDetails = "StitcherUnrecoverableStorageError"
	PrinterStatusDetailsstitcherWarmingUp                         PrinterStatusDetails = "StitcherWarmingUp"
	PrinterStatusDetailsstoppedPartially                          PrinterStatusDetails = "StoppedPartially"
	PrinterStatusDetailsstopping                                  PrinterStatusDetails = "Stopping"
	PrinterStatusDetailssubunitAdded                              PrinterStatusDetails = "SubunitAdded"
	PrinterStatusDetailssubunitAlmostEmpty                        PrinterStatusDetails = "SubunitAlmostEmpty"
	PrinterStatusDetailssubunitAlmostFull                         PrinterStatusDetails = "SubunitAlmostFull"
	PrinterStatusDetailssubunitAtLimit                            PrinterStatusDetails = "SubunitAtLimit"
	PrinterStatusDetailssubunitClosed                             PrinterStatusDetails = "SubunitClosed"
	PrinterStatusDetailssubunitCoolingDown                        PrinterStatusDetails = "SubunitCoolingDown"
	PrinterStatusDetailssubunitEmpty                              PrinterStatusDetails = "SubunitEmpty"
	PrinterStatusDetailssubunitFull                               PrinterStatusDetails = "SubunitFull"
	PrinterStatusDetailssubunitLifeAlmostOver                     PrinterStatusDetails = "SubunitLifeAlmostOver"
	PrinterStatusDetailssubunitLifeOver                           PrinterStatusDetails = "SubunitLifeOver"
	PrinterStatusDetailssubunitMemoryExhausted                    PrinterStatusDetails = "SubunitMemoryExhausted"
	PrinterStatusDetailssubunitMissing                            PrinterStatusDetails = "SubunitMissing"
	PrinterStatusDetailssubunitMotorFailure                       PrinterStatusDetails = "SubunitMotorFailure"
	PrinterStatusDetailssubunitNearLimit                          PrinterStatusDetails = "SubunitNearLimit"
	PrinterStatusDetailssubunitOffline                            PrinterStatusDetails = "SubunitOffline"
	PrinterStatusDetailssubunitOpened                             PrinterStatusDetails = "SubunitOpened"
	PrinterStatusDetailssubunitOverTemperature                    PrinterStatusDetails = "SubunitOverTemperature"
	PrinterStatusDetailssubunitPowerSaver                         PrinterStatusDetails = "SubunitPowerSaver"
	PrinterStatusDetailssubunitRecoverableFailure                 PrinterStatusDetails = "SubunitRecoverableFailure"
	PrinterStatusDetailssubunitRecoverableStorage                 PrinterStatusDetails = "SubunitRecoverableStorage"
	PrinterStatusDetailssubunitRemoved                            PrinterStatusDetails = "SubunitRemoved"
	PrinterStatusDetailssubunitResourceAdded                      PrinterStatusDetails = "SubunitResourceAdded"
	PrinterStatusDetailssubunitResourceRemoved                    PrinterStatusDetails = "SubunitResourceRemoved"
	PrinterStatusDetailssubunitThermistorFailure                  PrinterStatusDetails = "SubunitThermistorFailure"
	PrinterStatusDetailssubunitTimingFailure                      PrinterStatusDetails = "SubunitTimingFailure"
	PrinterStatusDetailssubunitTurnedOff                          PrinterStatusDetails = "SubunitTurnedOff"
	PrinterStatusDetailssubunitTurnedOn                           PrinterStatusDetails = "SubunitTurnedOn"
	PrinterStatusDetailssubunitUnderTemperature                   PrinterStatusDetails = "SubunitUnderTemperature"
	PrinterStatusDetailssubunitUnrecoverableFailure               PrinterStatusDetails = "SubunitUnrecoverableFailure"
	PrinterStatusDetailssubunitUnrecoverableStorage               PrinterStatusDetails = "SubunitUnrecoverableStorage"
	PrinterStatusDetailssubunitWarmingUp                          PrinterStatusDetails = "SubunitWarmingUp"
	PrinterStatusDetailssuspend                                   PrinterStatusDetails = "Suspend"
	PrinterStatusDetailstesting                                   PrinterStatusDetails = "Testing"
	PrinterStatusDetailstimedOut                                  PrinterStatusDetails = "TimedOut"
	PrinterStatusDetailstonerEmpty                                PrinterStatusDetails = "TonerEmpty"
	PrinterStatusDetailstonerLow                                  PrinterStatusDetails = "TonerLow"
	PrinterStatusDetailstrimmerAdded                              PrinterStatusDetails = "TrimmerAdded"
	PrinterStatusDetailstrimmerAlmostEmpty                        PrinterStatusDetails = "TrimmerAlmostEmpty"
	PrinterStatusDetailstrimmerAlmostFull                         PrinterStatusDetails = "TrimmerAlmostFull"
	PrinterStatusDetailstrimmerAtLimit                            PrinterStatusDetails = "TrimmerAtLimit"
	PrinterStatusDetailstrimmerClosed                             PrinterStatusDetails = "TrimmerClosed"
	PrinterStatusDetailstrimmerConfigurationChange                PrinterStatusDetails = "TrimmerConfigurationChange"
	PrinterStatusDetailstrimmerCoverClosed                        PrinterStatusDetails = "TrimmerCoverClosed"
	PrinterStatusDetailstrimmerCoverOpen                          PrinterStatusDetails = "TrimmerCoverOpen"
	PrinterStatusDetailstrimmerEmpty                              PrinterStatusDetails = "TrimmerEmpty"
	PrinterStatusDetailstrimmerFull                               PrinterStatusDetails = "TrimmerFull"
	PrinterStatusDetailstrimmerInterlockClosed                    PrinterStatusDetails = "TrimmerInterlockClosed"
	PrinterStatusDetailstrimmerInterlockOpen                      PrinterStatusDetails = "TrimmerInterlockOpen"
	PrinterStatusDetailstrimmerJam                                PrinterStatusDetails = "TrimmerJam"
	PrinterStatusDetailstrimmerLifeAlmostOver                     PrinterStatusDetails = "TrimmerLifeAlmostOver"
	PrinterStatusDetailstrimmerLifeOver                           PrinterStatusDetails = "TrimmerLifeOver"
	PrinterStatusDetailstrimmerMemoryExhausted                    PrinterStatusDetails = "TrimmerMemoryExhausted"
	PrinterStatusDetailstrimmerMissing                            PrinterStatusDetails = "TrimmerMissing"
	PrinterStatusDetailstrimmerMotorFailure                       PrinterStatusDetails = "TrimmerMotorFailure"
	PrinterStatusDetailstrimmerNearLimit                          PrinterStatusDetails = "TrimmerNearLimit"
	PrinterStatusDetailstrimmerOffline                            PrinterStatusDetails = "TrimmerOffline"
	PrinterStatusDetailstrimmerOpened                             PrinterStatusDetails = "TrimmerOpened"
	PrinterStatusDetailstrimmerOverTemperature                    PrinterStatusDetails = "TrimmerOverTemperature"
	PrinterStatusDetailstrimmerPowerSaver                         PrinterStatusDetails = "TrimmerPowerSaver"
	PrinterStatusDetailstrimmerRecoverableFailure                 PrinterStatusDetails = "TrimmerRecoverableFailure"
	PrinterStatusDetailstrimmerRecoverableStorage                 PrinterStatusDetails = "TrimmerRecoverableStorage"
	PrinterStatusDetailstrimmerRemoved                            PrinterStatusDetails = "TrimmerRemoved"
	PrinterStatusDetailstrimmerResourceAdded                      PrinterStatusDetails = "TrimmerResourceAdded"
	PrinterStatusDetailstrimmerResourceRemoved                    PrinterStatusDetails = "TrimmerResourceRemoved"
	PrinterStatusDetailstrimmerThermistorFailure                  PrinterStatusDetails = "TrimmerThermistorFailure"
	PrinterStatusDetailstrimmerTimingFailure                      PrinterStatusDetails = "TrimmerTimingFailure"
	PrinterStatusDetailstrimmerTurnedOff                          PrinterStatusDetails = "TrimmerTurnedOff"
	PrinterStatusDetailstrimmerTurnedOn                           PrinterStatusDetails = "TrimmerTurnedOn"
	PrinterStatusDetailstrimmerUnderTemperature                   PrinterStatusDetails = "TrimmerUnderTemperature"
	PrinterStatusDetailstrimmerUnrecoverableFailure               PrinterStatusDetails = "TrimmerUnrecoverableFailure"
	PrinterStatusDetailstrimmerUnrecoverableStorageError          PrinterStatusDetails = "TrimmerUnrecoverableStorageError"
	PrinterStatusDetailstrimmerWarmingUp                          PrinterStatusDetails = "TrimmerWarmingUp"
	PrinterStatusDetailsunknown                                   PrinterStatusDetails = "Unknown"
	PrinterStatusDetailswrapperAdded                              PrinterStatusDetails = "WrapperAdded"
	PrinterStatusDetailswrapperAlmostEmpty                        PrinterStatusDetails = "WrapperAlmostEmpty"
	PrinterStatusDetailswrapperAlmostFull                         PrinterStatusDetails = "WrapperAlmostFull"
	PrinterStatusDetailswrapperAtLimit                            PrinterStatusDetails = "WrapperAtLimit"
	PrinterStatusDetailswrapperClosed                             PrinterStatusDetails = "WrapperClosed"
	PrinterStatusDetailswrapperConfigurationChange                PrinterStatusDetails = "WrapperConfigurationChange"
	PrinterStatusDetailswrapperCoverClosed                        PrinterStatusDetails = "WrapperCoverClosed"
	PrinterStatusDetailswrapperCoverOpen                          PrinterStatusDetails = "WrapperCoverOpen"
	PrinterStatusDetailswrapperEmpty                              PrinterStatusDetails = "WrapperEmpty"
	PrinterStatusDetailswrapperFull                               PrinterStatusDetails = "WrapperFull"
	PrinterStatusDetailswrapperInterlockClosed                    PrinterStatusDetails = "WrapperInterlockClosed"
	PrinterStatusDetailswrapperInterlockOpen                      PrinterStatusDetails = "WrapperInterlockOpen"
	PrinterStatusDetailswrapperJam                                PrinterStatusDetails = "WrapperJam"
	PrinterStatusDetailswrapperLifeAlmostOver                     PrinterStatusDetails = "WrapperLifeAlmostOver"
	PrinterStatusDetailswrapperLifeOver                           PrinterStatusDetails = "WrapperLifeOver"
	PrinterStatusDetailswrapperMemoryExhausted                    PrinterStatusDetails = "WrapperMemoryExhausted"
	PrinterStatusDetailswrapperMissing                            PrinterStatusDetails = "WrapperMissing"
	PrinterStatusDetailswrapperMotorFailure                       PrinterStatusDetails = "WrapperMotorFailure"
	PrinterStatusDetailswrapperNearLimit                          PrinterStatusDetails = "WrapperNearLimit"
	PrinterStatusDetailswrapperOffline                            PrinterStatusDetails = "WrapperOffline"
	PrinterStatusDetailswrapperOpened                             PrinterStatusDetails = "WrapperOpened"
	PrinterStatusDetailswrapperOverTemperature                    PrinterStatusDetails = "WrapperOverTemperature"
	PrinterStatusDetailswrapperPowerSaver                         PrinterStatusDetails = "WrapperPowerSaver"
	PrinterStatusDetailswrapperRecoverableFailure                 PrinterStatusDetails = "WrapperRecoverableFailure"
	PrinterStatusDetailswrapperRecoverableStorage                 PrinterStatusDetails = "WrapperRecoverableStorage"
	PrinterStatusDetailswrapperRemoved                            PrinterStatusDetails = "WrapperRemoved"
	PrinterStatusDetailswrapperResourceAdded                      PrinterStatusDetails = "WrapperResourceAdded"
	PrinterStatusDetailswrapperResourceRemoved                    PrinterStatusDetails = "WrapperResourceRemoved"
	PrinterStatusDetailswrapperThermistorFailure                  PrinterStatusDetails = "WrapperThermistorFailure"
	PrinterStatusDetailswrapperTimingFailure                      PrinterStatusDetails = "WrapperTimingFailure"
	PrinterStatusDetailswrapperTurnedOff                          PrinterStatusDetails = "WrapperTurnedOff"
	PrinterStatusDetailswrapperTurnedOn                           PrinterStatusDetails = "WrapperTurnedOn"
	PrinterStatusDetailswrapperUnderTemperature                   PrinterStatusDetails = "WrapperUnderTemperature"
	PrinterStatusDetailswrapperUnrecoverableFailure               PrinterStatusDetails = "WrapperUnrecoverableFailure"
	PrinterStatusDetailswrapperUnrecoverableStorageError          PrinterStatusDetails = "WrapperUnrecoverableStorageError"
	PrinterStatusDetailswrapperWarmingUp                          PrinterStatusDetails = "WrapperWarmingUp"
)

func PossibleValuesForPrinterStatusDetails() []string {
	return []string{
		string(PrinterStatusDetailsalertRemovalOfBinaryChangeEntry),
		string(PrinterStatusDetailsbanderAdded),
		string(PrinterStatusDetailsbanderAlmostEmpty),
		string(PrinterStatusDetailsbanderAlmostFull),
		string(PrinterStatusDetailsbanderAtLimit),
		string(PrinterStatusDetailsbanderClosed),
		string(PrinterStatusDetailsbanderConfigurationChange),
		string(PrinterStatusDetailsbanderCoverClosed),
		string(PrinterStatusDetailsbanderCoverOpen),
		string(PrinterStatusDetailsbanderEmpty),
		string(PrinterStatusDetailsbanderFull),
		string(PrinterStatusDetailsbanderInterlockClosed),
		string(PrinterStatusDetailsbanderInterlockOpen),
		string(PrinterStatusDetailsbanderJam),
		string(PrinterStatusDetailsbanderLifeAlmostOver),
		string(PrinterStatusDetailsbanderLifeOver),
		string(PrinterStatusDetailsbanderMemoryExhausted),
		string(PrinterStatusDetailsbanderMissing),
		string(PrinterStatusDetailsbanderMotorFailure),
		string(PrinterStatusDetailsbanderNearLimit),
		string(PrinterStatusDetailsbanderOffline),
		string(PrinterStatusDetailsbanderOpened),
		string(PrinterStatusDetailsbanderOverTemperature),
		string(PrinterStatusDetailsbanderPowerSaver),
		string(PrinterStatusDetailsbanderRecoverableFailure),
		string(PrinterStatusDetailsbanderRecoverableStorage),
		string(PrinterStatusDetailsbanderRemoved),
		string(PrinterStatusDetailsbanderResourceAdded),
		string(PrinterStatusDetailsbanderResourceRemoved),
		string(PrinterStatusDetailsbanderThermistorFailure),
		string(PrinterStatusDetailsbanderTimingFailure),
		string(PrinterStatusDetailsbanderTurnedOff),
		string(PrinterStatusDetailsbanderTurnedOn),
		string(PrinterStatusDetailsbanderUnderTemperature),
		string(PrinterStatusDetailsbanderUnrecoverableFailure),
		string(PrinterStatusDetailsbanderUnrecoverableStorageError),
		string(PrinterStatusDetailsbanderWarmingUp),
		string(PrinterStatusDetailsbinderAdded),
		string(PrinterStatusDetailsbinderAlmostEmpty),
		string(PrinterStatusDetailsbinderAlmostFull),
		string(PrinterStatusDetailsbinderAtLimit),
		string(PrinterStatusDetailsbinderClosed),
		string(PrinterStatusDetailsbinderConfigurationChange),
		string(PrinterStatusDetailsbinderCoverClosed),
		string(PrinterStatusDetailsbinderCoverOpen),
		string(PrinterStatusDetailsbinderEmpty),
		string(PrinterStatusDetailsbinderFull),
		string(PrinterStatusDetailsbinderInterlockClosed),
		string(PrinterStatusDetailsbinderInterlockOpen),
		string(PrinterStatusDetailsbinderJam),
		string(PrinterStatusDetailsbinderLifeAlmostOver),
		string(PrinterStatusDetailsbinderLifeOver),
		string(PrinterStatusDetailsbinderMemoryExhausted),
		string(PrinterStatusDetailsbinderMissing),
		string(PrinterStatusDetailsbinderMotorFailure),
		string(PrinterStatusDetailsbinderNearLimit),
		string(PrinterStatusDetailsbinderOffline),
		string(PrinterStatusDetailsbinderOpened),
		string(PrinterStatusDetailsbinderOverTemperature),
		string(PrinterStatusDetailsbinderPowerSaver),
		string(PrinterStatusDetailsbinderRecoverableFailure),
		string(PrinterStatusDetailsbinderRecoverableStorage),
		string(PrinterStatusDetailsbinderRemoved),
		string(PrinterStatusDetailsbinderResourceAdded),
		string(PrinterStatusDetailsbinderResourceRemoved),
		string(PrinterStatusDetailsbinderThermistorFailure),
		string(PrinterStatusDetailsbinderTimingFailure),
		string(PrinterStatusDetailsbinderTurnedOff),
		string(PrinterStatusDetailsbinderTurnedOn),
		string(PrinterStatusDetailsbinderUnderTemperature),
		string(PrinterStatusDetailsbinderUnrecoverableFailure),
		string(PrinterStatusDetailsbinderUnrecoverableStorageError),
		string(PrinterStatusDetailsbinderWarmingUp),
		string(PrinterStatusDetailscameraFailure),
		string(PrinterStatusDetailschamberCooling),
		string(PrinterStatusDetailschamberFailure),
		string(PrinterStatusDetailschamberHeating),
		string(PrinterStatusDetailschamberTemperatureHigh),
		string(PrinterStatusDetailschamberTemperatureLow),
		string(PrinterStatusDetailscleanerLifeAlmostOver),
		string(PrinterStatusDetailscleanerLifeOver),
		string(PrinterStatusDetailsconfigurationChange),
		string(PrinterStatusDetailsconnectingToDevice),
		string(PrinterStatusDetailscoverOpen),
		string(PrinterStatusDetailsdeactivated),
		string(PrinterStatusDetailsdeleted),
		string(PrinterStatusDetailsdeveloperEmpty),
		string(PrinterStatusDetailsdeveloperLow),
		string(PrinterStatusDetailsdieCutterAdded),
		string(PrinterStatusDetailsdieCutterAlmostEmpty),
		string(PrinterStatusDetailsdieCutterAlmostFull),
		string(PrinterStatusDetailsdieCutterAtLimit),
		string(PrinterStatusDetailsdieCutterClosed),
		string(PrinterStatusDetailsdieCutterConfigurationChange),
		string(PrinterStatusDetailsdieCutterCoverClosed),
		string(PrinterStatusDetailsdieCutterCoverOpen),
		string(PrinterStatusDetailsdieCutterEmpty),
		string(PrinterStatusDetailsdieCutterFull),
		string(PrinterStatusDetailsdieCutterInterlockClosed),
		string(PrinterStatusDetailsdieCutterInterlockOpen),
		string(PrinterStatusDetailsdieCutterJam),
		string(PrinterStatusDetailsdieCutterLifeAlmostOver),
		string(PrinterStatusDetailsdieCutterLifeOver),
		string(PrinterStatusDetailsdieCutterMemoryExhausted),
		string(PrinterStatusDetailsdieCutterMissing),
		string(PrinterStatusDetailsdieCutterMotorFailure),
		string(PrinterStatusDetailsdieCutterNearLimit),
		string(PrinterStatusDetailsdieCutterOffline),
		string(PrinterStatusDetailsdieCutterOpened),
		string(PrinterStatusDetailsdieCutterOverTemperature),
		string(PrinterStatusDetailsdieCutterPowerSaver),
		string(PrinterStatusDetailsdieCutterRecoverableFailure),
		string(PrinterStatusDetailsdieCutterRecoverableStorage),
		string(PrinterStatusDetailsdieCutterRemoved),
		string(PrinterStatusDetailsdieCutterResourceAdded),
		string(PrinterStatusDetailsdieCutterResourceRemoved),
		string(PrinterStatusDetailsdieCutterThermistorFailure),
		string(PrinterStatusDetailsdieCutterTimingFailure),
		string(PrinterStatusDetailsdieCutterTurnedOff),
		string(PrinterStatusDetailsdieCutterTurnedOn),
		string(PrinterStatusDetailsdieCutterUnderTemperature),
		string(PrinterStatusDetailsdieCutterUnrecoverableFailure),
		string(PrinterStatusDetailsdieCutterUnrecoverableStorageError),
		string(PrinterStatusDetailsdieCutterWarmingUp),
		string(PrinterStatusDetailsdoorOpen),
		string(PrinterStatusDetailsextruderCooling),
		string(PrinterStatusDetailsextruderFailure),
		string(PrinterStatusDetailsextruderHeating),
		string(PrinterStatusDetailsextruderJam),
		string(PrinterStatusDetailsextruderTemperatureHigh),
		string(PrinterStatusDetailsextruderTemperatureLow),
		string(PrinterStatusDetailsfanFailure),
		string(PrinterStatusDetailsfaxModemLifeAlmostOver),
		string(PrinterStatusDetailsfaxModemLifeOver),
		string(PrinterStatusDetailsfaxModemMissing),
		string(PrinterStatusDetailsfaxModemTurnedOff),
		string(PrinterStatusDetailsfaxModemTurnedOn),
		string(PrinterStatusDetailsfolderAdded),
		string(PrinterStatusDetailsfolderAlmostEmpty),
		string(PrinterStatusDetailsfolderAlmostFull),
		string(PrinterStatusDetailsfolderAtLimit),
		string(PrinterStatusDetailsfolderClosed),
		string(PrinterStatusDetailsfolderConfigurationChange),
		string(PrinterStatusDetailsfolderCoverClosed),
		string(PrinterStatusDetailsfolderCoverOpen),
		string(PrinterStatusDetailsfolderEmpty),
		string(PrinterStatusDetailsfolderFull),
		string(PrinterStatusDetailsfolderInterlockClosed),
		string(PrinterStatusDetailsfolderInterlockOpen),
		string(PrinterStatusDetailsfolderJam),
		string(PrinterStatusDetailsfolderLifeAlmostOver),
		string(PrinterStatusDetailsfolderLifeOver),
		string(PrinterStatusDetailsfolderMemoryExhausted),
		string(PrinterStatusDetailsfolderMissing),
		string(PrinterStatusDetailsfolderMotorFailure),
		string(PrinterStatusDetailsfolderNearLimit),
		string(PrinterStatusDetailsfolderOffline),
		string(PrinterStatusDetailsfolderOpened),
		string(PrinterStatusDetailsfolderOverTemperature),
		string(PrinterStatusDetailsfolderPowerSaver),
		string(PrinterStatusDetailsfolderRecoverableFailure),
		string(PrinterStatusDetailsfolderRecoverableStorage),
		string(PrinterStatusDetailsfolderRemoved),
		string(PrinterStatusDetailsfolderResourceAdded),
		string(PrinterStatusDetailsfolderResourceRemoved),
		string(PrinterStatusDetailsfolderThermistorFailure),
		string(PrinterStatusDetailsfolderTimingFailure),
		string(PrinterStatusDetailsfolderTurnedOff),
		string(PrinterStatusDetailsfolderTurnedOn),
		string(PrinterStatusDetailsfolderUnderTemperature),
		string(PrinterStatusDetailsfolderUnrecoverableFailure),
		string(PrinterStatusDetailsfolderUnrecoverableStorageError),
		string(PrinterStatusDetailsfolderWarmingUp),
		string(PrinterStatusDetailsfuserOverTemp),
		string(PrinterStatusDetailsfuserUnderTemp),
		string(PrinterStatusDetailshibernate),
		string(PrinterStatusDetailsholdNewJobs),
		string(PrinterStatusDetailsidentifyPrinterRequested),
		string(PrinterStatusDetailsimprinterAdded),
		string(PrinterStatusDetailsimprinterAlmostEmpty),
		string(PrinterStatusDetailsimprinterAlmostFull),
		string(PrinterStatusDetailsimprinterAtLimit),
		string(PrinterStatusDetailsimprinterClosed),
		string(PrinterStatusDetailsimprinterConfigurationChange),
		string(PrinterStatusDetailsimprinterCoverClosed),
		string(PrinterStatusDetailsimprinterCoverOpen),
		string(PrinterStatusDetailsimprinterEmpty),
		string(PrinterStatusDetailsimprinterFull),
		string(PrinterStatusDetailsimprinterInterlockClosed),
		string(PrinterStatusDetailsimprinterInterlockOpen),
		string(PrinterStatusDetailsimprinterJam),
		string(PrinterStatusDetailsimprinterLifeAlmostOver),
		string(PrinterStatusDetailsimprinterLifeOver),
		string(PrinterStatusDetailsimprinterMemoryExhausted),
		string(PrinterStatusDetailsimprinterMissing),
		string(PrinterStatusDetailsimprinterMotorFailure),
		string(PrinterStatusDetailsimprinterNearLimit),
		string(PrinterStatusDetailsimprinterOffline),
		string(PrinterStatusDetailsimprinterOpened),
		string(PrinterStatusDetailsimprinterOverTemperature),
		string(PrinterStatusDetailsimprinterPowerSaver),
		string(PrinterStatusDetailsimprinterRecoverableFailure),
		string(PrinterStatusDetailsimprinterRecoverableStorage),
		string(PrinterStatusDetailsimprinterRemoved),
		string(PrinterStatusDetailsimprinterResourceAdded),
		string(PrinterStatusDetailsimprinterResourceRemoved),
		string(PrinterStatusDetailsimprinterThermistorFailure),
		string(PrinterStatusDetailsimprinterTimingFailure),
		string(PrinterStatusDetailsimprinterTurnedOff),
		string(PrinterStatusDetailsimprinterTurnedOn),
		string(PrinterStatusDetailsimprinterUnderTemperature),
		string(PrinterStatusDetailsimprinterUnrecoverableFailure),
		string(PrinterStatusDetailsimprinterUnrecoverableStorageError),
		string(PrinterStatusDetailsimprinterWarmingUp),
		string(PrinterStatusDetailsinputCannotFeedSizeSelected),
		string(PrinterStatusDetailsinputManualInputRequest),
		string(PrinterStatusDetailsinputMediaColorChange),
		string(PrinterStatusDetailsinputMediaFormPartsChange),
		string(PrinterStatusDetailsinputMediaSizeChange),
		string(PrinterStatusDetailsinputMediaTrayFailure),
		string(PrinterStatusDetailsinputMediaTrayFeedError),
		string(PrinterStatusDetailsinputMediaTrayJam),
		string(PrinterStatusDetailsinputMediaTypeChange),
		string(PrinterStatusDetailsinputMediaWeightChange),
		string(PrinterStatusDetailsinputPickRollerFailure),
		string(PrinterStatusDetailsinputPickRollerLifeOver),
		string(PrinterStatusDetailsinputPickRollerLifeWarn),
		string(PrinterStatusDetailsinputPickRollerMissing),
		string(PrinterStatusDetailsinputTrayElevationFailure),
		string(PrinterStatusDetailsinputTrayMissing),
		string(PrinterStatusDetailsinputTrayPositionFailure),
		string(PrinterStatusDetailsinserterAdded),
		string(PrinterStatusDetailsinserterAlmostEmpty),
		string(PrinterStatusDetailsinserterAlmostFull),
		string(PrinterStatusDetailsinserterAtLimit),
		string(PrinterStatusDetailsinserterClosed),
		string(PrinterStatusDetailsinserterConfigurationChange),
		string(PrinterStatusDetailsinserterCoverClosed),
		string(PrinterStatusDetailsinserterCoverOpen),
		string(PrinterStatusDetailsinserterEmpty),
		string(PrinterStatusDetailsinserterFull),
		string(PrinterStatusDetailsinserterInterlockClosed),
		string(PrinterStatusDetailsinserterInterlockOpen),
		string(PrinterStatusDetailsinserterJam),
		string(PrinterStatusDetailsinserterLifeAlmostOver),
		string(PrinterStatusDetailsinserterLifeOver),
		string(PrinterStatusDetailsinserterMemoryExhausted),
		string(PrinterStatusDetailsinserterMissing),
		string(PrinterStatusDetailsinserterMotorFailure),
		string(PrinterStatusDetailsinserterNearLimit),
		string(PrinterStatusDetailsinserterOffline),
		string(PrinterStatusDetailsinserterOpened),
		string(PrinterStatusDetailsinserterOverTemperature),
		string(PrinterStatusDetailsinserterPowerSaver),
		string(PrinterStatusDetailsinserterRecoverableFailure),
		string(PrinterStatusDetailsinserterRecoverableStorage),
		string(PrinterStatusDetailsinserterRemoved),
		string(PrinterStatusDetailsinserterResourceAdded),
		string(PrinterStatusDetailsinserterResourceRemoved),
		string(PrinterStatusDetailsinserterThermistorFailure),
		string(PrinterStatusDetailsinserterTimingFailure),
		string(PrinterStatusDetailsinserterTurnedOff),
		string(PrinterStatusDetailsinserterTurnedOn),
		string(PrinterStatusDetailsinserterUnderTemperature),
		string(PrinterStatusDetailsinserterUnrecoverableFailure),
		string(PrinterStatusDetailsinserterUnrecoverableStorageError),
		string(PrinterStatusDetailsinserterWarmingUp),
		string(PrinterStatusDetailsinterlockClosed),
		string(PrinterStatusDetailsinterlockOpen),
		string(PrinterStatusDetailsinterpreterCartridgeAdded),
		string(PrinterStatusDetailsinterpreterCartridgeDeleted),
		string(PrinterStatusDetailsinterpreterComplexPageEncountered),
		string(PrinterStatusDetailsinterpreterMemoryDecrease),
		string(PrinterStatusDetailsinterpreterMemoryIncrease),
		string(PrinterStatusDetailsinterpreterResourceAdded),
		string(PrinterStatusDetailsinterpreterResourceDeleted),
		string(PrinterStatusDetailsinterpreterResourceUnavailable),
		string(PrinterStatusDetailslampAtEol),
		string(PrinterStatusDetailslampFailure),
		string(PrinterStatusDetailslampNearEol),
		string(PrinterStatusDetailslaserAtEol),
		string(PrinterStatusDetailslaserFailure),
		string(PrinterStatusDetailslaserNearEol),
		string(PrinterStatusDetailsmakeEnvelopeAdded),
		string(PrinterStatusDetailsmakeEnvelopeAlmostEmpty),
		string(PrinterStatusDetailsmakeEnvelopeAlmostFull),
		string(PrinterStatusDetailsmakeEnvelopeAtLimit),
		string(PrinterStatusDetailsmakeEnvelopeClosed),
		string(PrinterStatusDetailsmakeEnvelopeConfigurationChange),
		string(PrinterStatusDetailsmakeEnvelopeCoverClosed),
		string(PrinterStatusDetailsmakeEnvelopeCoverOpen),
		string(PrinterStatusDetailsmakeEnvelopeEmpty),
		string(PrinterStatusDetailsmakeEnvelopeFull),
		string(PrinterStatusDetailsmakeEnvelopeInterlockClosed),
		string(PrinterStatusDetailsmakeEnvelopeInterlockOpen),
		string(PrinterStatusDetailsmakeEnvelopeJam),
		string(PrinterStatusDetailsmakeEnvelopeLifeAlmostOver),
		string(PrinterStatusDetailsmakeEnvelopeLifeOver),
		string(PrinterStatusDetailsmakeEnvelopeMemoryExhausted),
		string(PrinterStatusDetailsmakeEnvelopeMissing),
		string(PrinterStatusDetailsmakeEnvelopeMotorFailure),
		string(PrinterStatusDetailsmakeEnvelopeNearLimit),
		string(PrinterStatusDetailsmakeEnvelopeOffline),
		string(PrinterStatusDetailsmakeEnvelopeOpened),
		string(PrinterStatusDetailsmakeEnvelopeOverTemperature),
		string(PrinterStatusDetailsmakeEnvelopePowerSaver),
		string(PrinterStatusDetailsmakeEnvelopeRecoverableFailure),
		string(PrinterStatusDetailsmakeEnvelopeRecoverableStorage),
		string(PrinterStatusDetailsmakeEnvelopeRemoved),
		string(PrinterStatusDetailsmakeEnvelopeResourceAdded),
		string(PrinterStatusDetailsmakeEnvelopeResourceRemoved),
		string(PrinterStatusDetailsmakeEnvelopeThermistorFailure),
		string(PrinterStatusDetailsmakeEnvelopeTimingFailure),
		string(PrinterStatusDetailsmakeEnvelopeTurnedOff),
		string(PrinterStatusDetailsmakeEnvelopeTurnedOn),
		string(PrinterStatusDetailsmakeEnvelopeUnderTemperature),
		string(PrinterStatusDetailsmakeEnvelopeUnrecoverableFailure),
		string(PrinterStatusDetailsmakeEnvelopeUnrecoverableStorageError),
		string(PrinterStatusDetailsmakeEnvelopeWarmingUp),
		string(PrinterStatusDetailsmarkerAdjustingPrintQuality),
		string(PrinterStatusDetailsmarkerCleanerMissing),
		string(PrinterStatusDetailsmarkerDeveloperAlmostEmpty),
		string(PrinterStatusDetailsmarkerDeveloperEmpty),
		string(PrinterStatusDetailsmarkerDeveloperMissing),
		string(PrinterStatusDetailsmarkerFuserMissing),
		string(PrinterStatusDetailsmarkerFuserThermistorFailure),
		string(PrinterStatusDetailsmarkerFuserTimingFailure),
		string(PrinterStatusDetailsmarkerInkAlmostEmpty),
		string(PrinterStatusDetailsmarkerInkEmpty),
		string(PrinterStatusDetailsmarkerInkMissing),
		string(PrinterStatusDetailsmarkerOpcMissing),
		string(PrinterStatusDetailsmarkerPrintRibbonAlmostEmpty),
		string(PrinterStatusDetailsmarkerPrintRibbonEmpty),
		string(PrinterStatusDetailsmarkerPrintRibbonMissing),
		string(PrinterStatusDetailsmarkerSupplyAlmostEmpty),
		string(PrinterStatusDetailsmarkerSupplyEmpty),
		string(PrinterStatusDetailsmarkerSupplyLow),
		string(PrinterStatusDetailsmarkerSupplyMissing),
		string(PrinterStatusDetailsmarkerTonerCartridgeMissing),
		string(PrinterStatusDetailsmarkerTonerMissing),
		string(PrinterStatusDetailsmarkerWasteAlmostFull),
		string(PrinterStatusDetailsmarkerWasteFull),
		string(PrinterStatusDetailsmarkerWasteInkReceptacleAlmostFull),
		string(PrinterStatusDetailsmarkerWasteInkReceptacleFull),
		string(PrinterStatusDetailsmarkerWasteInkReceptacleMissing),
		string(PrinterStatusDetailsmarkerWasteMissing),
		string(PrinterStatusDetailsmarkerWasteTonerReceptacleAlmostFull),
		string(PrinterStatusDetailsmarkerWasteTonerReceptacleFull),
		string(PrinterStatusDetailsmarkerWasteTonerReceptacleMissing),
		string(PrinterStatusDetailsmaterialEmpty),
		string(PrinterStatusDetailsmaterialLow),
		string(PrinterStatusDetailsmaterialNeeded),
		string(PrinterStatusDetailsmediaDrying),
		string(PrinterStatusDetailsmediaEmpty),
		string(PrinterStatusDetailsmediaJam),
		string(PrinterStatusDetailsmediaLow),
		string(PrinterStatusDetailsmediaNeeded),
		string(PrinterStatusDetailsmediaPathCannotDuplexMediaSelected),
		string(PrinterStatusDetailsmediaPathFailure),
		string(PrinterStatusDetailsmediaPathInputEmpty),
		string(PrinterStatusDetailsmediaPathInputFeedError),
		string(PrinterStatusDetailsmediaPathInputJam),
		string(PrinterStatusDetailsmediaPathInputRequest),
		string(PrinterStatusDetailsmediaPathJam),
		string(PrinterStatusDetailsmediaPathMediaTrayAlmostFull),
		string(PrinterStatusDetailsmediaPathMediaTrayFull),
		string(PrinterStatusDetailsmediaPathMediaTrayMissing),
		string(PrinterStatusDetailsmediaPathOutputFeedError),
		string(PrinterStatusDetailsmediaPathOutputFull),
		string(PrinterStatusDetailsmediaPathOutputJam),
		string(PrinterStatusDetailsmediaPathPickRollerFailure),
		string(PrinterStatusDetailsmediaPathPickRollerLifeOver),
		string(PrinterStatusDetailsmediaPathPickRollerLifeWarn),
		string(PrinterStatusDetailsmediaPathPickRollerMissing),
		string(PrinterStatusDetailsmotorFailure),
		string(PrinterStatusDetailsmovingToPaused),
		string(PrinterStatusDetailsnone),
		string(PrinterStatusDetailsopticalPhotoConductorLifeOver),
		string(PrinterStatusDetailsopticalPhotoConductorNearEndOfLife),
		string(PrinterStatusDetailsother),
		string(PrinterStatusDetailsoutputAreaAlmostFull),
		string(PrinterStatusDetailsoutputAreaFull),
		string(PrinterStatusDetailsoutputMailboxSelectFailure),
		string(PrinterStatusDetailsoutputMediaTrayFailure),
		string(PrinterStatusDetailsoutputMediaTrayFeedError),
		string(PrinterStatusDetailsoutputMediaTrayJam),
		string(PrinterStatusDetailsoutputTrayMissing),
		string(PrinterStatusDetailspaused),
		string(PrinterStatusDetailsperforaterAdded),
		string(PrinterStatusDetailsperforaterAlmostEmpty),
		string(PrinterStatusDetailsperforaterAlmostFull),
		string(PrinterStatusDetailsperforaterAtLimit),
		string(PrinterStatusDetailsperforaterClosed),
		string(PrinterStatusDetailsperforaterConfigurationChange),
		string(PrinterStatusDetailsperforaterCoverClosed),
		string(PrinterStatusDetailsperforaterCoverOpen),
		string(PrinterStatusDetailsperforaterEmpty),
		string(PrinterStatusDetailsperforaterFull),
		string(PrinterStatusDetailsperforaterInterlockClosed),
		string(PrinterStatusDetailsperforaterInterlockOpen),
		string(PrinterStatusDetailsperforaterJam),
		string(PrinterStatusDetailsperforaterLifeAlmostOver),
		string(PrinterStatusDetailsperforaterLifeOver),
		string(PrinterStatusDetailsperforaterMemoryExhausted),
		string(PrinterStatusDetailsperforaterMissing),
		string(PrinterStatusDetailsperforaterMotorFailure),
		string(PrinterStatusDetailsperforaterNearLimit),
		string(PrinterStatusDetailsperforaterOffline),
		string(PrinterStatusDetailsperforaterOpened),
		string(PrinterStatusDetailsperforaterOverTemperature),
		string(PrinterStatusDetailsperforaterPowerSaver),
		string(PrinterStatusDetailsperforaterRecoverableFailure),
		string(PrinterStatusDetailsperforaterRecoverableStorage),
		string(PrinterStatusDetailsperforaterRemoved),
		string(PrinterStatusDetailsperforaterResourceAdded),
		string(PrinterStatusDetailsperforaterResourceRemoved),
		string(PrinterStatusDetailsperforaterThermistorFailure),
		string(PrinterStatusDetailsperforaterTimingFailure),
		string(PrinterStatusDetailsperforaterTurnedOff),
		string(PrinterStatusDetailsperforaterTurnedOn),
		string(PrinterStatusDetailsperforaterUnderTemperature),
		string(PrinterStatusDetailsperforaterUnrecoverableFailure),
		string(PrinterStatusDetailsperforaterUnrecoverableStorageError),
		string(PrinterStatusDetailsperforaterWarmingUp),
		string(PrinterStatusDetailsplatformCooling),
		string(PrinterStatusDetailsplatformFailure),
		string(PrinterStatusDetailsplatformHeating),
		string(PrinterStatusDetailsplatformTemperatureHigh),
		string(PrinterStatusDetailsplatformTemperatureLow),
		string(PrinterStatusDetailspowerDown),
		string(PrinterStatusDetailspowerUp),
		string(PrinterStatusDetailsprinterManualReset),
		string(PrinterStatusDetailsprinterNmsReset),
		string(PrinterStatusDetailsprinterReadyToPrint),
		string(PrinterStatusDetailspuncherAdded),
		string(PrinterStatusDetailspuncherAlmostEmpty),
		string(PrinterStatusDetailspuncherAlmostFull),
		string(PrinterStatusDetailspuncherAtLimit),
		string(PrinterStatusDetailspuncherClosed),
		string(PrinterStatusDetailspuncherConfigurationChange),
		string(PrinterStatusDetailspuncherCoverClosed),
		string(PrinterStatusDetailspuncherCoverOpen),
		string(PrinterStatusDetailspuncherEmpty),
		string(PrinterStatusDetailspuncherFull),
		string(PrinterStatusDetailspuncherInterlockClosed),
		string(PrinterStatusDetailspuncherInterlockOpen),
		string(PrinterStatusDetailspuncherJam),
		string(PrinterStatusDetailspuncherLifeAlmostOver),
		string(PrinterStatusDetailspuncherLifeOver),
		string(PrinterStatusDetailspuncherMemoryExhausted),
		string(PrinterStatusDetailspuncherMissing),
		string(PrinterStatusDetailspuncherMotorFailure),
		string(PrinterStatusDetailspuncherNearLimit),
		string(PrinterStatusDetailspuncherOffline),
		string(PrinterStatusDetailspuncherOpened),
		string(PrinterStatusDetailspuncherOverTemperature),
		string(PrinterStatusDetailspuncherPowerSaver),
		string(PrinterStatusDetailspuncherRecoverableFailure),
		string(PrinterStatusDetailspuncherRecoverableStorage),
		string(PrinterStatusDetailspuncherRemoved),
		string(PrinterStatusDetailspuncherResourceAdded),
		string(PrinterStatusDetailspuncherResourceRemoved),
		string(PrinterStatusDetailspuncherThermistorFailure),
		string(PrinterStatusDetailspuncherTimingFailure),
		string(PrinterStatusDetailspuncherTurnedOff),
		string(PrinterStatusDetailspuncherTurnedOn),
		string(PrinterStatusDetailspuncherUnderTemperature),
		string(PrinterStatusDetailspuncherUnrecoverableFailure),
		string(PrinterStatusDetailspuncherUnrecoverableStorageError),
		string(PrinterStatusDetailspuncherWarmingUp),
		string(PrinterStatusDetailsresuming),
		string(PrinterStatusDetailsscanMediaPathFailure),
		string(PrinterStatusDetailsscanMediaPathInputEmpty),
		string(PrinterStatusDetailsscanMediaPathInputFeedError),
		string(PrinterStatusDetailsscanMediaPathInputJam),
		string(PrinterStatusDetailsscanMediaPathInputRequest),
		string(PrinterStatusDetailsscanMediaPathJam),
		string(PrinterStatusDetailsscanMediaPathOutputFeedError),
		string(PrinterStatusDetailsscanMediaPathOutputFull),
		string(PrinterStatusDetailsscanMediaPathOutputJam),
		string(PrinterStatusDetailsscanMediaPathPickRollerFailure),
		string(PrinterStatusDetailsscanMediaPathPickRollerLifeOver),
		string(PrinterStatusDetailsscanMediaPathPickRollerLifeWarn),
		string(PrinterStatusDetailsscanMediaPathPickRollerMissing),
		string(PrinterStatusDetailsscanMediaPathTrayAlmostFull),
		string(PrinterStatusDetailsscanMediaPathTrayFull),
		string(PrinterStatusDetailsscanMediaPathTrayMissing),
		string(PrinterStatusDetailsscannerLightFailure),
		string(PrinterStatusDetailsscannerLightLifeAlmostOver),
		string(PrinterStatusDetailsscannerLightLifeOver),
		string(PrinterStatusDetailsscannerLightMissing),
		string(PrinterStatusDetailsscannerSensorFailure),
		string(PrinterStatusDetailsscannerSensorLifeAlmostOver),
		string(PrinterStatusDetailsscannerSensorLifeOver),
		string(PrinterStatusDetailsscannerSensorMissing),
		string(PrinterStatusDetailsseparationCutterAdded),
		string(PrinterStatusDetailsseparationCutterAlmostEmpty),
		string(PrinterStatusDetailsseparationCutterAlmostFull),
		string(PrinterStatusDetailsseparationCutterAtLimit),
		string(PrinterStatusDetailsseparationCutterClosed),
		string(PrinterStatusDetailsseparationCutterConfigurationChange),
		string(PrinterStatusDetailsseparationCutterCoverClosed),
		string(PrinterStatusDetailsseparationCutterCoverOpen),
		string(PrinterStatusDetailsseparationCutterEmpty),
		string(PrinterStatusDetailsseparationCutterFull),
		string(PrinterStatusDetailsseparationCutterInterlockClosed),
		string(PrinterStatusDetailsseparationCutterInterlockOpen),
		string(PrinterStatusDetailsseparationCutterJam),
		string(PrinterStatusDetailsseparationCutterLifeAlmostOver),
		string(PrinterStatusDetailsseparationCutterLifeOver),
		string(PrinterStatusDetailsseparationCutterMemoryExhausted),
		string(PrinterStatusDetailsseparationCutterMissing),
		string(PrinterStatusDetailsseparationCutterMotorFailure),
		string(PrinterStatusDetailsseparationCutterNearLimit),
		string(PrinterStatusDetailsseparationCutterOffline),
		string(PrinterStatusDetailsseparationCutterOpened),
		string(PrinterStatusDetailsseparationCutterOverTemperature),
		string(PrinterStatusDetailsseparationCutterPowerSaver),
		string(PrinterStatusDetailsseparationCutterRecoverableFailure),
		string(PrinterStatusDetailsseparationCutterRecoverableStorage),
		string(PrinterStatusDetailsseparationCutterRemoved),
		string(PrinterStatusDetailsseparationCutterResourceAdded),
		string(PrinterStatusDetailsseparationCutterResourceRemoved),
		string(PrinterStatusDetailsseparationCutterThermistorFailure),
		string(PrinterStatusDetailsseparationCutterTimingFailure),
		string(PrinterStatusDetailsseparationCutterTurnedOff),
		string(PrinterStatusDetailsseparationCutterTurnedOn),
		string(PrinterStatusDetailsseparationCutterUnderTemperature),
		string(PrinterStatusDetailsseparationCutterUnrecoverableFailure),
		string(PrinterStatusDetailsseparationCutterUnrecoverableStorageError),
		string(PrinterStatusDetailsseparationCutterWarmingUp),
		string(PrinterStatusDetailssheetRotatorAdded),
		string(PrinterStatusDetailssheetRotatorAlmostEmpty),
		string(PrinterStatusDetailssheetRotatorAlmostFull),
		string(PrinterStatusDetailssheetRotatorAtLimit),
		string(PrinterStatusDetailssheetRotatorClosed),
		string(PrinterStatusDetailssheetRotatorConfigurationChange),
		string(PrinterStatusDetailssheetRotatorCoverClosed),
		string(PrinterStatusDetailssheetRotatorCoverOpen),
		string(PrinterStatusDetailssheetRotatorEmpty),
		string(PrinterStatusDetailssheetRotatorFull),
		string(PrinterStatusDetailssheetRotatorInterlockClosed),
		string(PrinterStatusDetailssheetRotatorInterlockOpen),
		string(PrinterStatusDetailssheetRotatorJam),
		string(PrinterStatusDetailssheetRotatorLifeAlmostOver),
		string(PrinterStatusDetailssheetRotatorLifeOver),
		string(PrinterStatusDetailssheetRotatorMemoryExhausted),
		string(PrinterStatusDetailssheetRotatorMissing),
		string(PrinterStatusDetailssheetRotatorMotorFailure),
		string(PrinterStatusDetailssheetRotatorNearLimit),
		string(PrinterStatusDetailssheetRotatorOffline),
		string(PrinterStatusDetailssheetRotatorOpened),
		string(PrinterStatusDetailssheetRotatorOverTemperature),
		string(PrinterStatusDetailssheetRotatorPowerSaver),
		string(PrinterStatusDetailssheetRotatorRecoverableFailure),
		string(PrinterStatusDetailssheetRotatorRecoverableStorage),
		string(PrinterStatusDetailssheetRotatorRemoved),
		string(PrinterStatusDetailssheetRotatorResourceAdded),
		string(PrinterStatusDetailssheetRotatorResourceRemoved),
		string(PrinterStatusDetailssheetRotatorThermistorFailure),
		string(PrinterStatusDetailssheetRotatorTimingFailure),
		string(PrinterStatusDetailssheetRotatorTurnedOff),
		string(PrinterStatusDetailssheetRotatorTurnedOn),
		string(PrinterStatusDetailssheetRotatorUnderTemperature),
		string(PrinterStatusDetailssheetRotatorUnrecoverableFailure),
		string(PrinterStatusDetailssheetRotatorUnrecoverableStorageError),
		string(PrinterStatusDetailssheetRotatorWarmingUp),
		string(PrinterStatusDetailsshutdown),
		string(PrinterStatusDetailsslitterAdded),
		string(PrinterStatusDetailsslitterAlmostEmpty),
		string(PrinterStatusDetailsslitterAlmostFull),
		string(PrinterStatusDetailsslitterAtLimit),
		string(PrinterStatusDetailsslitterClosed),
		string(PrinterStatusDetailsslitterConfigurationChange),
		string(PrinterStatusDetailsslitterCoverClosed),
		string(PrinterStatusDetailsslitterCoverOpen),
		string(PrinterStatusDetailsslitterEmpty),
		string(PrinterStatusDetailsslitterFull),
		string(PrinterStatusDetailsslitterInterlockClosed),
		string(PrinterStatusDetailsslitterInterlockOpen),
		string(PrinterStatusDetailsslitterJam),
		string(PrinterStatusDetailsslitterLifeAlmostOver),
		string(PrinterStatusDetailsslitterLifeOver),
		string(PrinterStatusDetailsslitterMemoryExhausted),
		string(PrinterStatusDetailsslitterMissing),
		string(PrinterStatusDetailsslitterMotorFailure),
		string(PrinterStatusDetailsslitterNearLimit),
		string(PrinterStatusDetailsslitterOffline),
		string(PrinterStatusDetailsslitterOpened),
		string(PrinterStatusDetailsslitterOverTemperature),
		string(PrinterStatusDetailsslitterPowerSaver),
		string(PrinterStatusDetailsslitterRecoverableFailure),
		string(PrinterStatusDetailsslitterRecoverableStorage),
		string(PrinterStatusDetailsslitterRemoved),
		string(PrinterStatusDetailsslitterResourceAdded),
		string(PrinterStatusDetailsslitterResourceRemoved),
		string(PrinterStatusDetailsslitterThermistorFailure),
		string(PrinterStatusDetailsslitterTimingFailure),
		string(PrinterStatusDetailsslitterTurnedOff),
		string(PrinterStatusDetailsslitterTurnedOn),
		string(PrinterStatusDetailsslitterUnderTemperature),
		string(PrinterStatusDetailsslitterUnrecoverableFailure),
		string(PrinterStatusDetailsslitterUnrecoverableStorageError),
		string(PrinterStatusDetailsslitterWarmingUp),
		string(PrinterStatusDetailsspoolAreaFull),
		string(PrinterStatusDetailsstackerAdded),
		string(PrinterStatusDetailsstackerAlmostEmpty),
		string(PrinterStatusDetailsstackerAlmostFull),
		string(PrinterStatusDetailsstackerAtLimit),
		string(PrinterStatusDetailsstackerClosed),
		string(PrinterStatusDetailsstackerConfigurationChange),
		string(PrinterStatusDetailsstackerCoverClosed),
		string(PrinterStatusDetailsstackerCoverOpen),
		string(PrinterStatusDetailsstackerEmpty),
		string(PrinterStatusDetailsstackerFull),
		string(PrinterStatusDetailsstackerInterlockClosed),
		string(PrinterStatusDetailsstackerInterlockOpen),
		string(PrinterStatusDetailsstackerJam),
		string(PrinterStatusDetailsstackerLifeAlmostOver),
		string(PrinterStatusDetailsstackerLifeOver),
		string(PrinterStatusDetailsstackerMemoryExhausted),
		string(PrinterStatusDetailsstackerMissing),
		string(PrinterStatusDetailsstackerMotorFailure),
		string(PrinterStatusDetailsstackerNearLimit),
		string(PrinterStatusDetailsstackerOffline),
		string(PrinterStatusDetailsstackerOpened),
		string(PrinterStatusDetailsstackerOverTemperature),
		string(PrinterStatusDetailsstackerPowerSaver),
		string(PrinterStatusDetailsstackerRecoverableFailure),
		string(PrinterStatusDetailsstackerRecoverableStorage),
		string(PrinterStatusDetailsstackerRemoved),
		string(PrinterStatusDetailsstackerResourceAdded),
		string(PrinterStatusDetailsstackerResourceRemoved),
		string(PrinterStatusDetailsstackerThermistorFailure),
		string(PrinterStatusDetailsstackerTimingFailure),
		string(PrinterStatusDetailsstackerTurnedOff),
		string(PrinterStatusDetailsstackerTurnedOn),
		string(PrinterStatusDetailsstackerUnderTemperature),
		string(PrinterStatusDetailsstackerUnrecoverableFailure),
		string(PrinterStatusDetailsstackerUnrecoverableStorageError),
		string(PrinterStatusDetailsstackerWarmingUp),
		string(PrinterStatusDetailsstandby),
		string(PrinterStatusDetailsstaplerAdded),
		string(PrinterStatusDetailsstaplerAlmostEmpty),
		string(PrinterStatusDetailsstaplerAlmostFull),
		string(PrinterStatusDetailsstaplerAtLimit),
		string(PrinterStatusDetailsstaplerClosed),
		string(PrinterStatusDetailsstaplerConfigurationChange),
		string(PrinterStatusDetailsstaplerCoverClosed),
		string(PrinterStatusDetailsstaplerCoverOpen),
		string(PrinterStatusDetailsstaplerEmpty),
		string(PrinterStatusDetailsstaplerFull),
		string(PrinterStatusDetailsstaplerInterlockClosed),
		string(PrinterStatusDetailsstaplerInterlockOpen),
		string(PrinterStatusDetailsstaplerJam),
		string(PrinterStatusDetailsstaplerLifeAlmostOver),
		string(PrinterStatusDetailsstaplerLifeOver),
		string(PrinterStatusDetailsstaplerMemoryExhausted),
		string(PrinterStatusDetailsstaplerMissing),
		string(PrinterStatusDetailsstaplerMotorFailure),
		string(PrinterStatusDetailsstaplerNearLimit),
		string(PrinterStatusDetailsstaplerOffline),
		string(PrinterStatusDetailsstaplerOpened),
		string(PrinterStatusDetailsstaplerOverTemperature),
		string(PrinterStatusDetailsstaplerPowerSaver),
		string(PrinterStatusDetailsstaplerRecoverableFailure),
		string(PrinterStatusDetailsstaplerRecoverableStorage),
		string(PrinterStatusDetailsstaplerRemoved),
		string(PrinterStatusDetailsstaplerResourceAdded),
		string(PrinterStatusDetailsstaplerResourceRemoved),
		string(PrinterStatusDetailsstaplerThermistorFailure),
		string(PrinterStatusDetailsstaplerTimingFailure),
		string(PrinterStatusDetailsstaplerTurnedOff),
		string(PrinterStatusDetailsstaplerTurnedOn),
		string(PrinterStatusDetailsstaplerUnderTemperature),
		string(PrinterStatusDetailsstaplerUnrecoverableFailure),
		string(PrinterStatusDetailsstaplerUnrecoverableStorageError),
		string(PrinterStatusDetailsstaplerWarmingUp),
		string(PrinterStatusDetailsstitcherAdded),
		string(PrinterStatusDetailsstitcherAlmostEmpty),
		string(PrinterStatusDetailsstitcherAlmostFull),
		string(PrinterStatusDetailsstitcherAtLimit),
		string(PrinterStatusDetailsstitcherClosed),
		string(PrinterStatusDetailsstitcherConfigurationChange),
		string(PrinterStatusDetailsstitcherCoverClosed),
		string(PrinterStatusDetailsstitcherCoverOpen),
		string(PrinterStatusDetailsstitcherEmpty),
		string(PrinterStatusDetailsstitcherFull),
		string(PrinterStatusDetailsstitcherInterlockClosed),
		string(PrinterStatusDetailsstitcherInterlockOpen),
		string(PrinterStatusDetailsstitcherJam),
		string(PrinterStatusDetailsstitcherLifeAlmostOver),
		string(PrinterStatusDetailsstitcherLifeOver),
		string(PrinterStatusDetailsstitcherMemoryExhausted),
		string(PrinterStatusDetailsstitcherMissing),
		string(PrinterStatusDetailsstitcherMotorFailure),
		string(PrinterStatusDetailsstitcherNearLimit),
		string(PrinterStatusDetailsstitcherOffline),
		string(PrinterStatusDetailsstitcherOpened),
		string(PrinterStatusDetailsstitcherOverTemperature),
		string(PrinterStatusDetailsstitcherPowerSaver),
		string(PrinterStatusDetailsstitcherRecoverableFailure),
		string(PrinterStatusDetailsstitcherRecoverableStorage),
		string(PrinterStatusDetailsstitcherRemoved),
		string(PrinterStatusDetailsstitcherResourceAdded),
		string(PrinterStatusDetailsstitcherResourceRemoved),
		string(PrinterStatusDetailsstitcherThermistorFailure),
		string(PrinterStatusDetailsstitcherTimingFailure),
		string(PrinterStatusDetailsstitcherTurnedOff),
		string(PrinterStatusDetailsstitcherTurnedOn),
		string(PrinterStatusDetailsstitcherUnderTemperature),
		string(PrinterStatusDetailsstitcherUnrecoverableFailure),
		string(PrinterStatusDetailsstitcherUnrecoverableStorageError),
		string(PrinterStatusDetailsstitcherWarmingUp),
		string(PrinterStatusDetailsstoppedPartially),
		string(PrinterStatusDetailsstopping),
		string(PrinterStatusDetailssubunitAdded),
		string(PrinterStatusDetailssubunitAlmostEmpty),
		string(PrinterStatusDetailssubunitAlmostFull),
		string(PrinterStatusDetailssubunitAtLimit),
		string(PrinterStatusDetailssubunitClosed),
		string(PrinterStatusDetailssubunitCoolingDown),
		string(PrinterStatusDetailssubunitEmpty),
		string(PrinterStatusDetailssubunitFull),
		string(PrinterStatusDetailssubunitLifeAlmostOver),
		string(PrinterStatusDetailssubunitLifeOver),
		string(PrinterStatusDetailssubunitMemoryExhausted),
		string(PrinterStatusDetailssubunitMissing),
		string(PrinterStatusDetailssubunitMotorFailure),
		string(PrinterStatusDetailssubunitNearLimit),
		string(PrinterStatusDetailssubunitOffline),
		string(PrinterStatusDetailssubunitOpened),
		string(PrinterStatusDetailssubunitOverTemperature),
		string(PrinterStatusDetailssubunitPowerSaver),
		string(PrinterStatusDetailssubunitRecoverableFailure),
		string(PrinterStatusDetailssubunitRecoverableStorage),
		string(PrinterStatusDetailssubunitRemoved),
		string(PrinterStatusDetailssubunitResourceAdded),
		string(PrinterStatusDetailssubunitResourceRemoved),
		string(PrinterStatusDetailssubunitThermistorFailure),
		string(PrinterStatusDetailssubunitTimingFailure),
		string(PrinterStatusDetailssubunitTurnedOff),
		string(PrinterStatusDetailssubunitTurnedOn),
		string(PrinterStatusDetailssubunitUnderTemperature),
		string(PrinterStatusDetailssubunitUnrecoverableFailure),
		string(PrinterStatusDetailssubunitUnrecoverableStorage),
		string(PrinterStatusDetailssubunitWarmingUp),
		string(PrinterStatusDetailssuspend),
		string(PrinterStatusDetailstesting),
		string(PrinterStatusDetailstimedOut),
		string(PrinterStatusDetailstonerEmpty),
		string(PrinterStatusDetailstonerLow),
		string(PrinterStatusDetailstrimmerAdded),
		string(PrinterStatusDetailstrimmerAlmostEmpty),
		string(PrinterStatusDetailstrimmerAlmostFull),
		string(PrinterStatusDetailstrimmerAtLimit),
		string(PrinterStatusDetailstrimmerClosed),
		string(PrinterStatusDetailstrimmerConfigurationChange),
		string(PrinterStatusDetailstrimmerCoverClosed),
		string(PrinterStatusDetailstrimmerCoverOpen),
		string(PrinterStatusDetailstrimmerEmpty),
		string(PrinterStatusDetailstrimmerFull),
		string(PrinterStatusDetailstrimmerInterlockClosed),
		string(PrinterStatusDetailstrimmerInterlockOpen),
		string(PrinterStatusDetailstrimmerJam),
		string(PrinterStatusDetailstrimmerLifeAlmostOver),
		string(PrinterStatusDetailstrimmerLifeOver),
		string(PrinterStatusDetailstrimmerMemoryExhausted),
		string(PrinterStatusDetailstrimmerMissing),
		string(PrinterStatusDetailstrimmerMotorFailure),
		string(PrinterStatusDetailstrimmerNearLimit),
		string(PrinterStatusDetailstrimmerOffline),
		string(PrinterStatusDetailstrimmerOpened),
		string(PrinterStatusDetailstrimmerOverTemperature),
		string(PrinterStatusDetailstrimmerPowerSaver),
		string(PrinterStatusDetailstrimmerRecoverableFailure),
		string(PrinterStatusDetailstrimmerRecoverableStorage),
		string(PrinterStatusDetailstrimmerRemoved),
		string(PrinterStatusDetailstrimmerResourceAdded),
		string(PrinterStatusDetailstrimmerResourceRemoved),
		string(PrinterStatusDetailstrimmerThermistorFailure),
		string(PrinterStatusDetailstrimmerTimingFailure),
		string(PrinterStatusDetailstrimmerTurnedOff),
		string(PrinterStatusDetailstrimmerTurnedOn),
		string(PrinterStatusDetailstrimmerUnderTemperature),
		string(PrinterStatusDetailstrimmerUnrecoverableFailure),
		string(PrinterStatusDetailstrimmerUnrecoverableStorageError),
		string(PrinterStatusDetailstrimmerWarmingUp),
		string(PrinterStatusDetailsunknown),
		string(PrinterStatusDetailswrapperAdded),
		string(PrinterStatusDetailswrapperAlmostEmpty),
		string(PrinterStatusDetailswrapperAlmostFull),
		string(PrinterStatusDetailswrapperAtLimit),
		string(PrinterStatusDetailswrapperClosed),
		string(PrinterStatusDetailswrapperConfigurationChange),
		string(PrinterStatusDetailswrapperCoverClosed),
		string(PrinterStatusDetailswrapperCoverOpen),
		string(PrinterStatusDetailswrapperEmpty),
		string(PrinterStatusDetailswrapperFull),
		string(PrinterStatusDetailswrapperInterlockClosed),
		string(PrinterStatusDetailswrapperInterlockOpen),
		string(PrinterStatusDetailswrapperJam),
		string(PrinterStatusDetailswrapperLifeAlmostOver),
		string(PrinterStatusDetailswrapperLifeOver),
		string(PrinterStatusDetailswrapperMemoryExhausted),
		string(PrinterStatusDetailswrapperMissing),
		string(PrinterStatusDetailswrapperMotorFailure),
		string(PrinterStatusDetailswrapperNearLimit),
		string(PrinterStatusDetailswrapperOffline),
		string(PrinterStatusDetailswrapperOpened),
		string(PrinterStatusDetailswrapperOverTemperature),
		string(PrinterStatusDetailswrapperPowerSaver),
		string(PrinterStatusDetailswrapperRecoverableFailure),
		string(PrinterStatusDetailswrapperRecoverableStorage),
		string(PrinterStatusDetailswrapperRemoved),
		string(PrinterStatusDetailswrapperResourceAdded),
		string(PrinterStatusDetailswrapperResourceRemoved),
		string(PrinterStatusDetailswrapperThermistorFailure),
		string(PrinterStatusDetailswrapperTimingFailure),
		string(PrinterStatusDetailswrapperTurnedOff),
		string(PrinterStatusDetailswrapperTurnedOn),
		string(PrinterStatusDetailswrapperUnderTemperature),
		string(PrinterStatusDetailswrapperUnrecoverableFailure),
		string(PrinterStatusDetailswrapperUnrecoverableStorageError),
		string(PrinterStatusDetailswrapperWarmingUp),
	}
}

func parsePrinterStatusDetails(input string) (*PrinterStatusDetails, error) {
	vals := map[string]PrinterStatusDetails{
		"alertremovalofbinarychangeentry":           PrinterStatusDetailsalertRemovalOfBinaryChangeEntry,
		"banderadded":                               PrinterStatusDetailsbanderAdded,
		"banderalmostempty":                         PrinterStatusDetailsbanderAlmostEmpty,
		"banderalmostfull":                          PrinterStatusDetailsbanderAlmostFull,
		"banderatlimit":                             PrinterStatusDetailsbanderAtLimit,
		"banderclosed":                              PrinterStatusDetailsbanderClosed,
		"banderconfigurationchange":                 PrinterStatusDetailsbanderConfigurationChange,
		"bandercoverclosed":                         PrinterStatusDetailsbanderCoverClosed,
		"bandercoveropen":                           PrinterStatusDetailsbanderCoverOpen,
		"banderempty":                               PrinterStatusDetailsbanderEmpty,
		"banderfull":                                PrinterStatusDetailsbanderFull,
		"banderinterlockclosed":                     PrinterStatusDetailsbanderInterlockClosed,
		"banderinterlockopen":                       PrinterStatusDetailsbanderInterlockOpen,
		"banderjam":                                 PrinterStatusDetailsbanderJam,
		"banderlifealmostover":                      PrinterStatusDetailsbanderLifeAlmostOver,
		"banderlifeover":                            PrinterStatusDetailsbanderLifeOver,
		"bandermemoryexhausted":                     PrinterStatusDetailsbanderMemoryExhausted,
		"bandermissing":                             PrinterStatusDetailsbanderMissing,
		"bandermotorfailure":                        PrinterStatusDetailsbanderMotorFailure,
		"bandernearlimit":                           PrinterStatusDetailsbanderNearLimit,
		"banderoffline":                             PrinterStatusDetailsbanderOffline,
		"banderopened":                              PrinterStatusDetailsbanderOpened,
		"banderovertemperature":                     PrinterStatusDetailsbanderOverTemperature,
		"banderpowersaver":                          PrinterStatusDetailsbanderPowerSaver,
		"banderrecoverablefailure":                  PrinterStatusDetailsbanderRecoverableFailure,
		"banderrecoverablestorage":                  PrinterStatusDetailsbanderRecoverableStorage,
		"banderremoved":                             PrinterStatusDetailsbanderRemoved,
		"banderresourceadded":                       PrinterStatusDetailsbanderResourceAdded,
		"banderresourceremoved":                     PrinterStatusDetailsbanderResourceRemoved,
		"banderthermistorfailure":                   PrinterStatusDetailsbanderThermistorFailure,
		"bandertimingfailure":                       PrinterStatusDetailsbanderTimingFailure,
		"banderturnedoff":                           PrinterStatusDetailsbanderTurnedOff,
		"banderturnedon":                            PrinterStatusDetailsbanderTurnedOn,
		"banderundertemperature":                    PrinterStatusDetailsbanderUnderTemperature,
		"banderunrecoverablefailure":                PrinterStatusDetailsbanderUnrecoverableFailure,
		"banderunrecoverablestorageerror":           PrinterStatusDetailsbanderUnrecoverableStorageError,
		"banderwarmingup":                           PrinterStatusDetailsbanderWarmingUp,
		"binderadded":                               PrinterStatusDetailsbinderAdded,
		"binderalmostempty":                         PrinterStatusDetailsbinderAlmostEmpty,
		"binderalmostfull":                          PrinterStatusDetailsbinderAlmostFull,
		"binderatlimit":                             PrinterStatusDetailsbinderAtLimit,
		"binderclosed":                              PrinterStatusDetailsbinderClosed,
		"binderconfigurationchange":                 PrinterStatusDetailsbinderConfigurationChange,
		"bindercoverclosed":                         PrinterStatusDetailsbinderCoverClosed,
		"bindercoveropen":                           PrinterStatusDetailsbinderCoverOpen,
		"binderempty":                               PrinterStatusDetailsbinderEmpty,
		"binderfull":                                PrinterStatusDetailsbinderFull,
		"binderinterlockclosed":                     PrinterStatusDetailsbinderInterlockClosed,
		"binderinterlockopen":                       PrinterStatusDetailsbinderInterlockOpen,
		"binderjam":                                 PrinterStatusDetailsbinderJam,
		"binderlifealmostover":                      PrinterStatusDetailsbinderLifeAlmostOver,
		"binderlifeover":                            PrinterStatusDetailsbinderLifeOver,
		"bindermemoryexhausted":                     PrinterStatusDetailsbinderMemoryExhausted,
		"bindermissing":                             PrinterStatusDetailsbinderMissing,
		"bindermotorfailure":                        PrinterStatusDetailsbinderMotorFailure,
		"bindernearlimit":                           PrinterStatusDetailsbinderNearLimit,
		"binderoffline":                             PrinterStatusDetailsbinderOffline,
		"binderopened":                              PrinterStatusDetailsbinderOpened,
		"binderovertemperature":                     PrinterStatusDetailsbinderOverTemperature,
		"binderpowersaver":                          PrinterStatusDetailsbinderPowerSaver,
		"binderrecoverablefailure":                  PrinterStatusDetailsbinderRecoverableFailure,
		"binderrecoverablestorage":                  PrinterStatusDetailsbinderRecoverableStorage,
		"binderremoved":                             PrinterStatusDetailsbinderRemoved,
		"binderresourceadded":                       PrinterStatusDetailsbinderResourceAdded,
		"binderresourceremoved":                     PrinterStatusDetailsbinderResourceRemoved,
		"binderthermistorfailure":                   PrinterStatusDetailsbinderThermistorFailure,
		"bindertimingfailure":                       PrinterStatusDetailsbinderTimingFailure,
		"binderturnedoff":                           PrinterStatusDetailsbinderTurnedOff,
		"binderturnedon":                            PrinterStatusDetailsbinderTurnedOn,
		"binderundertemperature":                    PrinterStatusDetailsbinderUnderTemperature,
		"binderunrecoverablefailure":                PrinterStatusDetailsbinderUnrecoverableFailure,
		"binderunrecoverablestorageerror":           PrinterStatusDetailsbinderUnrecoverableStorageError,
		"binderwarmingup":                           PrinterStatusDetailsbinderWarmingUp,
		"camerafailure":                             PrinterStatusDetailscameraFailure,
		"chambercooling":                            PrinterStatusDetailschamberCooling,
		"chamberfailure":                            PrinterStatusDetailschamberFailure,
		"chamberheating":                            PrinterStatusDetailschamberHeating,
		"chambertemperaturehigh":                    PrinterStatusDetailschamberTemperatureHigh,
		"chambertemperaturelow":                     PrinterStatusDetailschamberTemperatureLow,
		"cleanerlifealmostover":                     PrinterStatusDetailscleanerLifeAlmostOver,
		"cleanerlifeover":                           PrinterStatusDetailscleanerLifeOver,
		"configurationchange":                       PrinterStatusDetailsconfigurationChange,
		"connectingtodevice":                        PrinterStatusDetailsconnectingToDevice,
		"coveropen":                                 PrinterStatusDetailscoverOpen,
		"deactivated":                               PrinterStatusDetailsdeactivated,
		"deleted":                                   PrinterStatusDetailsdeleted,
		"developerempty":                            PrinterStatusDetailsdeveloperEmpty,
		"developerlow":                              PrinterStatusDetailsdeveloperLow,
		"diecutteradded":                            PrinterStatusDetailsdieCutterAdded,
		"diecutteralmostempty":                      PrinterStatusDetailsdieCutterAlmostEmpty,
		"diecutteralmostfull":                       PrinterStatusDetailsdieCutterAlmostFull,
		"diecutteratlimit":                          PrinterStatusDetailsdieCutterAtLimit,
		"diecutterclosed":                           PrinterStatusDetailsdieCutterClosed,
		"diecutterconfigurationchange":              PrinterStatusDetailsdieCutterConfigurationChange,
		"diecuttercoverclosed":                      PrinterStatusDetailsdieCutterCoverClosed,
		"diecuttercoveropen":                        PrinterStatusDetailsdieCutterCoverOpen,
		"diecutterempty":                            PrinterStatusDetailsdieCutterEmpty,
		"diecutterfull":                             PrinterStatusDetailsdieCutterFull,
		"diecutterinterlockclosed":                  PrinterStatusDetailsdieCutterInterlockClosed,
		"diecutterinterlockopen":                    PrinterStatusDetailsdieCutterInterlockOpen,
		"diecutterjam":                              PrinterStatusDetailsdieCutterJam,
		"diecutterlifealmostover":                   PrinterStatusDetailsdieCutterLifeAlmostOver,
		"diecutterlifeover":                         PrinterStatusDetailsdieCutterLifeOver,
		"diecuttermemoryexhausted":                  PrinterStatusDetailsdieCutterMemoryExhausted,
		"diecuttermissing":                          PrinterStatusDetailsdieCutterMissing,
		"diecuttermotorfailure":                     PrinterStatusDetailsdieCutterMotorFailure,
		"diecutternearlimit":                        PrinterStatusDetailsdieCutterNearLimit,
		"diecutteroffline":                          PrinterStatusDetailsdieCutterOffline,
		"diecutteropened":                           PrinterStatusDetailsdieCutterOpened,
		"diecutterovertemperature":                  PrinterStatusDetailsdieCutterOverTemperature,
		"diecutterpowersaver":                       PrinterStatusDetailsdieCutterPowerSaver,
		"diecutterrecoverablefailure":               PrinterStatusDetailsdieCutterRecoverableFailure,
		"diecutterrecoverablestorage":               PrinterStatusDetailsdieCutterRecoverableStorage,
		"diecutterremoved":                          PrinterStatusDetailsdieCutterRemoved,
		"diecutterresourceadded":                    PrinterStatusDetailsdieCutterResourceAdded,
		"diecutterresourceremoved":                  PrinterStatusDetailsdieCutterResourceRemoved,
		"diecutterthermistorfailure":                PrinterStatusDetailsdieCutterThermistorFailure,
		"diecuttertimingfailure":                    PrinterStatusDetailsdieCutterTimingFailure,
		"diecutterturnedoff":                        PrinterStatusDetailsdieCutterTurnedOff,
		"diecutterturnedon":                         PrinterStatusDetailsdieCutterTurnedOn,
		"diecutterundertemperature":                 PrinterStatusDetailsdieCutterUnderTemperature,
		"diecutterunrecoverablefailure":             PrinterStatusDetailsdieCutterUnrecoverableFailure,
		"diecutterunrecoverablestorageerror":        PrinterStatusDetailsdieCutterUnrecoverableStorageError,
		"diecutterwarmingup":                        PrinterStatusDetailsdieCutterWarmingUp,
		"dooropen":                                  PrinterStatusDetailsdoorOpen,
		"extrudercooling":                           PrinterStatusDetailsextruderCooling,
		"extruderfailure":                           PrinterStatusDetailsextruderFailure,
		"extruderheating":                           PrinterStatusDetailsextruderHeating,
		"extruderjam":                               PrinterStatusDetailsextruderJam,
		"extrudertemperaturehigh":                   PrinterStatusDetailsextruderTemperatureHigh,
		"extrudertemperaturelow":                    PrinterStatusDetailsextruderTemperatureLow,
		"fanfailure":                                PrinterStatusDetailsfanFailure,
		"faxmodemlifealmostover":                    PrinterStatusDetailsfaxModemLifeAlmostOver,
		"faxmodemlifeover":                          PrinterStatusDetailsfaxModemLifeOver,
		"faxmodemmissing":                           PrinterStatusDetailsfaxModemMissing,
		"faxmodemturnedoff":                         PrinterStatusDetailsfaxModemTurnedOff,
		"faxmodemturnedon":                          PrinterStatusDetailsfaxModemTurnedOn,
		"folderadded":                               PrinterStatusDetailsfolderAdded,
		"folderalmostempty":                         PrinterStatusDetailsfolderAlmostEmpty,
		"folderalmostfull":                          PrinterStatusDetailsfolderAlmostFull,
		"folderatlimit":                             PrinterStatusDetailsfolderAtLimit,
		"folderclosed":                              PrinterStatusDetailsfolderClosed,
		"folderconfigurationchange":                 PrinterStatusDetailsfolderConfigurationChange,
		"foldercoverclosed":                         PrinterStatusDetailsfolderCoverClosed,
		"foldercoveropen":                           PrinterStatusDetailsfolderCoverOpen,
		"folderempty":                               PrinterStatusDetailsfolderEmpty,
		"folderfull":                                PrinterStatusDetailsfolderFull,
		"folderinterlockclosed":                     PrinterStatusDetailsfolderInterlockClosed,
		"folderinterlockopen":                       PrinterStatusDetailsfolderInterlockOpen,
		"folderjam":                                 PrinterStatusDetailsfolderJam,
		"folderlifealmostover":                      PrinterStatusDetailsfolderLifeAlmostOver,
		"folderlifeover":                            PrinterStatusDetailsfolderLifeOver,
		"foldermemoryexhausted":                     PrinterStatusDetailsfolderMemoryExhausted,
		"foldermissing":                             PrinterStatusDetailsfolderMissing,
		"foldermotorfailure":                        PrinterStatusDetailsfolderMotorFailure,
		"foldernearlimit":                           PrinterStatusDetailsfolderNearLimit,
		"folderoffline":                             PrinterStatusDetailsfolderOffline,
		"folderopened":                              PrinterStatusDetailsfolderOpened,
		"folderovertemperature":                     PrinterStatusDetailsfolderOverTemperature,
		"folderpowersaver":                          PrinterStatusDetailsfolderPowerSaver,
		"folderrecoverablefailure":                  PrinterStatusDetailsfolderRecoverableFailure,
		"folderrecoverablestorage":                  PrinterStatusDetailsfolderRecoverableStorage,
		"folderremoved":                             PrinterStatusDetailsfolderRemoved,
		"folderresourceadded":                       PrinterStatusDetailsfolderResourceAdded,
		"folderresourceremoved":                     PrinterStatusDetailsfolderResourceRemoved,
		"folderthermistorfailure":                   PrinterStatusDetailsfolderThermistorFailure,
		"foldertimingfailure":                       PrinterStatusDetailsfolderTimingFailure,
		"folderturnedoff":                           PrinterStatusDetailsfolderTurnedOff,
		"folderturnedon":                            PrinterStatusDetailsfolderTurnedOn,
		"folderundertemperature":                    PrinterStatusDetailsfolderUnderTemperature,
		"folderunrecoverablefailure":                PrinterStatusDetailsfolderUnrecoverableFailure,
		"folderunrecoverablestorageerror":           PrinterStatusDetailsfolderUnrecoverableStorageError,
		"folderwarmingup":                           PrinterStatusDetailsfolderWarmingUp,
		"fuserovertemp":                             PrinterStatusDetailsfuserOverTemp,
		"fuserundertemp":                            PrinterStatusDetailsfuserUnderTemp,
		"hibernate":                                 PrinterStatusDetailshibernate,
		"holdnewjobs":                               PrinterStatusDetailsholdNewJobs,
		"identifyprinterrequested":                  PrinterStatusDetailsidentifyPrinterRequested,
		"imprinteradded":                            PrinterStatusDetailsimprinterAdded,
		"imprinteralmostempty":                      PrinterStatusDetailsimprinterAlmostEmpty,
		"imprinteralmostfull":                       PrinterStatusDetailsimprinterAlmostFull,
		"imprinteratlimit":                          PrinterStatusDetailsimprinterAtLimit,
		"imprinterclosed":                           PrinterStatusDetailsimprinterClosed,
		"imprinterconfigurationchange":              PrinterStatusDetailsimprinterConfigurationChange,
		"imprintercoverclosed":                      PrinterStatusDetailsimprinterCoverClosed,
		"imprintercoveropen":                        PrinterStatusDetailsimprinterCoverOpen,
		"imprinterempty":                            PrinterStatusDetailsimprinterEmpty,
		"imprinterfull":                             PrinterStatusDetailsimprinterFull,
		"imprinterinterlockclosed":                  PrinterStatusDetailsimprinterInterlockClosed,
		"imprinterinterlockopen":                    PrinterStatusDetailsimprinterInterlockOpen,
		"imprinterjam":                              PrinterStatusDetailsimprinterJam,
		"imprinterlifealmostover":                   PrinterStatusDetailsimprinterLifeAlmostOver,
		"imprinterlifeover":                         PrinterStatusDetailsimprinterLifeOver,
		"imprintermemoryexhausted":                  PrinterStatusDetailsimprinterMemoryExhausted,
		"imprintermissing":                          PrinterStatusDetailsimprinterMissing,
		"imprintermotorfailure":                     PrinterStatusDetailsimprinterMotorFailure,
		"imprinternearlimit":                        PrinterStatusDetailsimprinterNearLimit,
		"imprinteroffline":                          PrinterStatusDetailsimprinterOffline,
		"imprinteropened":                           PrinterStatusDetailsimprinterOpened,
		"imprinterovertemperature":                  PrinterStatusDetailsimprinterOverTemperature,
		"imprinterpowersaver":                       PrinterStatusDetailsimprinterPowerSaver,
		"imprinterrecoverablefailure":               PrinterStatusDetailsimprinterRecoverableFailure,
		"imprinterrecoverablestorage":               PrinterStatusDetailsimprinterRecoverableStorage,
		"imprinterremoved":                          PrinterStatusDetailsimprinterRemoved,
		"imprinterresourceadded":                    PrinterStatusDetailsimprinterResourceAdded,
		"imprinterresourceremoved":                  PrinterStatusDetailsimprinterResourceRemoved,
		"imprinterthermistorfailure":                PrinterStatusDetailsimprinterThermistorFailure,
		"imprintertimingfailure":                    PrinterStatusDetailsimprinterTimingFailure,
		"imprinterturnedoff":                        PrinterStatusDetailsimprinterTurnedOff,
		"imprinterturnedon":                         PrinterStatusDetailsimprinterTurnedOn,
		"imprinterundertemperature":                 PrinterStatusDetailsimprinterUnderTemperature,
		"imprinterunrecoverablefailure":             PrinterStatusDetailsimprinterUnrecoverableFailure,
		"imprinterunrecoverablestorageerror":        PrinterStatusDetailsimprinterUnrecoverableStorageError,
		"imprinterwarmingup":                        PrinterStatusDetailsimprinterWarmingUp,
		"inputcannotfeedsizeselected":               PrinterStatusDetailsinputCannotFeedSizeSelected,
		"inputmanualinputrequest":                   PrinterStatusDetailsinputManualInputRequest,
		"inputmediacolorchange":                     PrinterStatusDetailsinputMediaColorChange,
		"inputmediaformpartschange":                 PrinterStatusDetailsinputMediaFormPartsChange,
		"inputmediasizechange":                      PrinterStatusDetailsinputMediaSizeChange,
		"inputmediatrayfailure":                     PrinterStatusDetailsinputMediaTrayFailure,
		"inputmediatrayfeederror":                   PrinterStatusDetailsinputMediaTrayFeedError,
		"inputmediatrayjam":                         PrinterStatusDetailsinputMediaTrayJam,
		"inputmediatypechange":                      PrinterStatusDetailsinputMediaTypeChange,
		"inputmediaweightchange":                    PrinterStatusDetailsinputMediaWeightChange,
		"inputpickrollerfailure":                    PrinterStatusDetailsinputPickRollerFailure,
		"inputpickrollerlifeover":                   PrinterStatusDetailsinputPickRollerLifeOver,
		"inputpickrollerlifewarn":                   PrinterStatusDetailsinputPickRollerLifeWarn,
		"inputpickrollermissing":                    PrinterStatusDetailsinputPickRollerMissing,
		"inputtrayelevationfailure":                 PrinterStatusDetailsinputTrayElevationFailure,
		"inputtraymissing":                          PrinterStatusDetailsinputTrayMissing,
		"inputtraypositionfailure":                  PrinterStatusDetailsinputTrayPositionFailure,
		"inserteradded":                             PrinterStatusDetailsinserterAdded,
		"inserteralmostempty":                       PrinterStatusDetailsinserterAlmostEmpty,
		"inserteralmostfull":                        PrinterStatusDetailsinserterAlmostFull,
		"inserteratlimit":                           PrinterStatusDetailsinserterAtLimit,
		"inserterclosed":                            PrinterStatusDetailsinserterClosed,
		"inserterconfigurationchange":               PrinterStatusDetailsinserterConfigurationChange,
		"insertercoverclosed":                       PrinterStatusDetailsinserterCoverClosed,
		"insertercoveropen":                         PrinterStatusDetailsinserterCoverOpen,
		"inserterempty":                             PrinterStatusDetailsinserterEmpty,
		"inserterfull":                              PrinterStatusDetailsinserterFull,
		"inserterinterlockclosed":                   PrinterStatusDetailsinserterInterlockClosed,
		"inserterinterlockopen":                     PrinterStatusDetailsinserterInterlockOpen,
		"inserterjam":                               PrinterStatusDetailsinserterJam,
		"inserterlifealmostover":                    PrinterStatusDetailsinserterLifeAlmostOver,
		"inserterlifeover":                          PrinterStatusDetailsinserterLifeOver,
		"insertermemoryexhausted":                   PrinterStatusDetailsinserterMemoryExhausted,
		"insertermissing":                           PrinterStatusDetailsinserterMissing,
		"insertermotorfailure":                      PrinterStatusDetailsinserterMotorFailure,
		"inserternearlimit":                         PrinterStatusDetailsinserterNearLimit,
		"inserteroffline":                           PrinterStatusDetailsinserterOffline,
		"inserteropened":                            PrinterStatusDetailsinserterOpened,
		"inserterovertemperature":                   PrinterStatusDetailsinserterOverTemperature,
		"inserterpowersaver":                        PrinterStatusDetailsinserterPowerSaver,
		"inserterrecoverablefailure":                PrinterStatusDetailsinserterRecoverableFailure,
		"inserterrecoverablestorage":                PrinterStatusDetailsinserterRecoverableStorage,
		"inserterremoved":                           PrinterStatusDetailsinserterRemoved,
		"inserterresourceadded":                     PrinterStatusDetailsinserterResourceAdded,
		"inserterresourceremoved":                   PrinterStatusDetailsinserterResourceRemoved,
		"inserterthermistorfailure":                 PrinterStatusDetailsinserterThermistorFailure,
		"insertertimingfailure":                     PrinterStatusDetailsinserterTimingFailure,
		"inserterturnedoff":                         PrinterStatusDetailsinserterTurnedOff,
		"inserterturnedon":                          PrinterStatusDetailsinserterTurnedOn,
		"inserterundertemperature":                  PrinterStatusDetailsinserterUnderTemperature,
		"inserterunrecoverablefailure":              PrinterStatusDetailsinserterUnrecoverableFailure,
		"inserterunrecoverablestorageerror":         PrinterStatusDetailsinserterUnrecoverableStorageError,
		"inserterwarmingup":                         PrinterStatusDetailsinserterWarmingUp,
		"interlockclosed":                           PrinterStatusDetailsinterlockClosed,
		"interlockopen":                             PrinterStatusDetailsinterlockOpen,
		"interpretercartridgeadded":                 PrinterStatusDetailsinterpreterCartridgeAdded,
		"interpretercartridgedeleted":               PrinterStatusDetailsinterpreterCartridgeDeleted,
		"interpretercomplexpageencountered":         PrinterStatusDetailsinterpreterComplexPageEncountered,
		"interpretermemorydecrease":                 PrinterStatusDetailsinterpreterMemoryDecrease,
		"interpretermemoryincrease":                 PrinterStatusDetailsinterpreterMemoryIncrease,
		"interpreterresourceadded":                  PrinterStatusDetailsinterpreterResourceAdded,
		"interpreterresourcedeleted":                PrinterStatusDetailsinterpreterResourceDeleted,
		"interpreterresourceunavailable":            PrinterStatusDetailsinterpreterResourceUnavailable,
		"lampateol":                                 PrinterStatusDetailslampAtEol,
		"lampfailure":                               PrinterStatusDetailslampFailure,
		"lampneareol":                               PrinterStatusDetailslampNearEol,
		"laserateol":                                PrinterStatusDetailslaserAtEol,
		"laserfailure":                              PrinterStatusDetailslaserFailure,
		"laserneareol":                              PrinterStatusDetailslaserNearEol,
		"makeenvelopeadded":                         PrinterStatusDetailsmakeEnvelopeAdded,
		"makeenvelopealmostempty":                   PrinterStatusDetailsmakeEnvelopeAlmostEmpty,
		"makeenvelopealmostfull":                    PrinterStatusDetailsmakeEnvelopeAlmostFull,
		"makeenvelopeatlimit":                       PrinterStatusDetailsmakeEnvelopeAtLimit,
		"makeenvelopeclosed":                        PrinterStatusDetailsmakeEnvelopeClosed,
		"makeenvelopeconfigurationchange":           PrinterStatusDetailsmakeEnvelopeConfigurationChange,
		"makeenvelopecoverclosed":                   PrinterStatusDetailsmakeEnvelopeCoverClosed,
		"makeenvelopecoveropen":                     PrinterStatusDetailsmakeEnvelopeCoverOpen,
		"makeenvelopeempty":                         PrinterStatusDetailsmakeEnvelopeEmpty,
		"makeenvelopefull":                          PrinterStatusDetailsmakeEnvelopeFull,
		"makeenvelopeinterlockclosed":               PrinterStatusDetailsmakeEnvelopeInterlockClosed,
		"makeenvelopeinterlockopen":                 PrinterStatusDetailsmakeEnvelopeInterlockOpen,
		"makeenvelopejam":                           PrinterStatusDetailsmakeEnvelopeJam,
		"makeenvelopelifealmostover":                PrinterStatusDetailsmakeEnvelopeLifeAlmostOver,
		"makeenvelopelifeover":                      PrinterStatusDetailsmakeEnvelopeLifeOver,
		"makeenvelopememoryexhausted":               PrinterStatusDetailsmakeEnvelopeMemoryExhausted,
		"makeenvelopemissing":                       PrinterStatusDetailsmakeEnvelopeMissing,
		"makeenvelopemotorfailure":                  PrinterStatusDetailsmakeEnvelopeMotorFailure,
		"makeenvelopenearlimit":                     PrinterStatusDetailsmakeEnvelopeNearLimit,
		"makeenvelopeoffline":                       PrinterStatusDetailsmakeEnvelopeOffline,
		"makeenvelopeopened":                        PrinterStatusDetailsmakeEnvelopeOpened,
		"makeenvelopeovertemperature":               PrinterStatusDetailsmakeEnvelopeOverTemperature,
		"makeenvelopepowersaver":                    PrinterStatusDetailsmakeEnvelopePowerSaver,
		"makeenveloperecoverablefailure":            PrinterStatusDetailsmakeEnvelopeRecoverableFailure,
		"makeenveloperecoverablestorage":            PrinterStatusDetailsmakeEnvelopeRecoverableStorage,
		"makeenveloperemoved":                       PrinterStatusDetailsmakeEnvelopeRemoved,
		"makeenveloperesourceadded":                 PrinterStatusDetailsmakeEnvelopeResourceAdded,
		"makeenveloperesourceremoved":               PrinterStatusDetailsmakeEnvelopeResourceRemoved,
		"makeenvelopethermistorfailure":             PrinterStatusDetailsmakeEnvelopeThermistorFailure,
		"makeenvelopetimingfailure":                 PrinterStatusDetailsmakeEnvelopeTimingFailure,
		"makeenvelopeturnedoff":                     PrinterStatusDetailsmakeEnvelopeTurnedOff,
		"makeenvelopeturnedon":                      PrinterStatusDetailsmakeEnvelopeTurnedOn,
		"makeenvelopeundertemperature":              PrinterStatusDetailsmakeEnvelopeUnderTemperature,
		"makeenvelopeunrecoverablefailure":          PrinterStatusDetailsmakeEnvelopeUnrecoverableFailure,
		"makeenvelopeunrecoverablestorageerror":     PrinterStatusDetailsmakeEnvelopeUnrecoverableStorageError,
		"makeenvelopewarmingup":                     PrinterStatusDetailsmakeEnvelopeWarmingUp,
		"markeradjustingprintquality":               PrinterStatusDetailsmarkerAdjustingPrintQuality,
		"markercleanermissing":                      PrinterStatusDetailsmarkerCleanerMissing,
		"markerdeveloperalmostempty":                PrinterStatusDetailsmarkerDeveloperAlmostEmpty,
		"markerdeveloperempty":                      PrinterStatusDetailsmarkerDeveloperEmpty,
		"markerdevelopermissing":                    PrinterStatusDetailsmarkerDeveloperMissing,
		"markerfusermissing":                        PrinterStatusDetailsmarkerFuserMissing,
		"markerfuserthermistorfailure":              PrinterStatusDetailsmarkerFuserThermistorFailure,
		"markerfusertimingfailure":                  PrinterStatusDetailsmarkerFuserTimingFailure,
		"markerinkalmostempty":                      PrinterStatusDetailsmarkerInkAlmostEmpty,
		"markerinkempty":                            PrinterStatusDetailsmarkerInkEmpty,
		"markerinkmissing":                          PrinterStatusDetailsmarkerInkMissing,
		"markeropcmissing":                          PrinterStatusDetailsmarkerOpcMissing,
		"markerprintribbonalmostempty":              PrinterStatusDetailsmarkerPrintRibbonAlmostEmpty,
		"markerprintribbonempty":                    PrinterStatusDetailsmarkerPrintRibbonEmpty,
		"markerprintribbonmissing":                  PrinterStatusDetailsmarkerPrintRibbonMissing,
		"markersupplyalmostempty":                   PrinterStatusDetailsmarkerSupplyAlmostEmpty,
		"markersupplyempty":                         PrinterStatusDetailsmarkerSupplyEmpty,
		"markersupplylow":                           PrinterStatusDetailsmarkerSupplyLow,
		"markersupplymissing":                       PrinterStatusDetailsmarkerSupplyMissing,
		"markertonercartridgemissing":               PrinterStatusDetailsmarkerTonerCartridgeMissing,
		"markertonermissing":                        PrinterStatusDetailsmarkerTonerMissing,
		"markerwastealmostfull":                     PrinterStatusDetailsmarkerWasteAlmostFull,
		"markerwastefull":                           PrinterStatusDetailsmarkerWasteFull,
		"markerwasteinkreceptaclealmostfull":        PrinterStatusDetailsmarkerWasteInkReceptacleAlmostFull,
		"markerwasteinkreceptaclefull":              PrinterStatusDetailsmarkerWasteInkReceptacleFull,
		"markerwasteinkreceptaclemissing":           PrinterStatusDetailsmarkerWasteInkReceptacleMissing,
		"markerwastemissing":                        PrinterStatusDetailsmarkerWasteMissing,
		"markerwastetonerreceptaclealmostfull":      PrinterStatusDetailsmarkerWasteTonerReceptacleAlmostFull,
		"markerwastetonerreceptaclefull":            PrinterStatusDetailsmarkerWasteTonerReceptacleFull,
		"markerwastetonerreceptaclemissing":         PrinterStatusDetailsmarkerWasteTonerReceptacleMissing,
		"materialempty":                             PrinterStatusDetailsmaterialEmpty,
		"materiallow":                               PrinterStatusDetailsmaterialLow,
		"materialneeded":                            PrinterStatusDetailsmaterialNeeded,
		"mediadrying":                               PrinterStatusDetailsmediaDrying,
		"mediaempty":                                PrinterStatusDetailsmediaEmpty,
		"mediajam":                                  PrinterStatusDetailsmediaJam,
		"medialow":                                  PrinterStatusDetailsmediaLow,
		"medianeeded":                               PrinterStatusDetailsmediaNeeded,
		"mediapathcannotduplexmediaselected":        PrinterStatusDetailsmediaPathCannotDuplexMediaSelected,
		"mediapathfailure":                          PrinterStatusDetailsmediaPathFailure,
		"mediapathinputempty":                       PrinterStatusDetailsmediaPathInputEmpty,
		"mediapathinputfeederror":                   PrinterStatusDetailsmediaPathInputFeedError,
		"mediapathinputjam":                         PrinterStatusDetailsmediaPathInputJam,
		"mediapathinputrequest":                     PrinterStatusDetailsmediaPathInputRequest,
		"mediapathjam":                              PrinterStatusDetailsmediaPathJam,
		"mediapathmediatrayalmostfull":              PrinterStatusDetailsmediaPathMediaTrayAlmostFull,
		"mediapathmediatrayfull":                    PrinterStatusDetailsmediaPathMediaTrayFull,
		"mediapathmediatraymissing":                 PrinterStatusDetailsmediaPathMediaTrayMissing,
		"mediapathoutputfeederror":                  PrinterStatusDetailsmediaPathOutputFeedError,
		"mediapathoutputfull":                       PrinterStatusDetailsmediaPathOutputFull,
		"mediapathoutputjam":                        PrinterStatusDetailsmediaPathOutputJam,
		"mediapathpickrollerfailure":                PrinterStatusDetailsmediaPathPickRollerFailure,
		"mediapathpickrollerlifeover":               PrinterStatusDetailsmediaPathPickRollerLifeOver,
		"mediapathpickrollerlifewarn":               PrinterStatusDetailsmediaPathPickRollerLifeWarn,
		"mediapathpickrollermissing":                PrinterStatusDetailsmediaPathPickRollerMissing,
		"motorfailure":                              PrinterStatusDetailsmotorFailure,
		"movingtopaused":                            PrinterStatusDetailsmovingToPaused,
		"none":                                      PrinterStatusDetailsnone,
		"opticalphotoconductorlifeover":             PrinterStatusDetailsopticalPhotoConductorLifeOver,
		"opticalphotoconductornearendoflife":        PrinterStatusDetailsopticalPhotoConductorNearEndOfLife,
		"other":                                     PrinterStatusDetailsother,
		"outputareaalmostfull":                      PrinterStatusDetailsoutputAreaAlmostFull,
		"outputareafull":                            PrinterStatusDetailsoutputAreaFull,
		"outputmailboxselectfailure":                PrinterStatusDetailsoutputMailboxSelectFailure,
		"outputmediatrayfailure":                    PrinterStatusDetailsoutputMediaTrayFailure,
		"outputmediatrayfeederror":                  PrinterStatusDetailsoutputMediaTrayFeedError,
		"outputmediatrayjam":                        PrinterStatusDetailsoutputMediaTrayJam,
		"outputtraymissing":                         PrinterStatusDetailsoutputTrayMissing,
		"paused":                                    PrinterStatusDetailspaused,
		"perforateradded":                           PrinterStatusDetailsperforaterAdded,
		"perforateralmostempty":                     PrinterStatusDetailsperforaterAlmostEmpty,
		"perforateralmostfull":                      PrinterStatusDetailsperforaterAlmostFull,
		"perforateratlimit":                         PrinterStatusDetailsperforaterAtLimit,
		"perforaterclosed":                          PrinterStatusDetailsperforaterClosed,
		"perforaterconfigurationchange":             PrinterStatusDetailsperforaterConfigurationChange,
		"perforatercoverclosed":                     PrinterStatusDetailsperforaterCoverClosed,
		"perforatercoveropen":                       PrinterStatusDetailsperforaterCoverOpen,
		"perforaterempty":                           PrinterStatusDetailsperforaterEmpty,
		"perforaterfull":                            PrinterStatusDetailsperforaterFull,
		"perforaterinterlockclosed":                 PrinterStatusDetailsperforaterInterlockClosed,
		"perforaterinterlockopen":                   PrinterStatusDetailsperforaterInterlockOpen,
		"perforaterjam":                             PrinterStatusDetailsperforaterJam,
		"perforaterlifealmostover":                  PrinterStatusDetailsperforaterLifeAlmostOver,
		"perforaterlifeover":                        PrinterStatusDetailsperforaterLifeOver,
		"perforatermemoryexhausted":                 PrinterStatusDetailsperforaterMemoryExhausted,
		"perforatermissing":                         PrinterStatusDetailsperforaterMissing,
		"perforatermotorfailure":                    PrinterStatusDetailsperforaterMotorFailure,
		"perforaternearlimit":                       PrinterStatusDetailsperforaterNearLimit,
		"perforateroffline":                         PrinterStatusDetailsperforaterOffline,
		"perforateropened":                          PrinterStatusDetailsperforaterOpened,
		"perforaterovertemperature":                 PrinterStatusDetailsperforaterOverTemperature,
		"perforaterpowersaver":                      PrinterStatusDetailsperforaterPowerSaver,
		"perforaterrecoverablefailure":              PrinterStatusDetailsperforaterRecoverableFailure,
		"perforaterrecoverablestorage":              PrinterStatusDetailsperforaterRecoverableStorage,
		"perforaterremoved":                         PrinterStatusDetailsperforaterRemoved,
		"perforaterresourceadded":                   PrinterStatusDetailsperforaterResourceAdded,
		"perforaterresourceremoved":                 PrinterStatusDetailsperforaterResourceRemoved,
		"perforaterthermistorfailure":               PrinterStatusDetailsperforaterThermistorFailure,
		"perforatertimingfailure":                   PrinterStatusDetailsperforaterTimingFailure,
		"perforaterturnedoff":                       PrinterStatusDetailsperforaterTurnedOff,
		"perforaterturnedon":                        PrinterStatusDetailsperforaterTurnedOn,
		"perforaterundertemperature":                PrinterStatusDetailsperforaterUnderTemperature,
		"perforaterunrecoverablefailure":            PrinterStatusDetailsperforaterUnrecoverableFailure,
		"perforaterunrecoverablestorageerror":       PrinterStatusDetailsperforaterUnrecoverableStorageError,
		"perforaterwarmingup":                       PrinterStatusDetailsperforaterWarmingUp,
		"platformcooling":                           PrinterStatusDetailsplatformCooling,
		"platformfailure":                           PrinterStatusDetailsplatformFailure,
		"platformheating":                           PrinterStatusDetailsplatformHeating,
		"platformtemperaturehigh":                   PrinterStatusDetailsplatformTemperatureHigh,
		"platformtemperaturelow":                    PrinterStatusDetailsplatformTemperatureLow,
		"powerdown":                                 PrinterStatusDetailspowerDown,
		"powerup":                                   PrinterStatusDetailspowerUp,
		"printermanualreset":                        PrinterStatusDetailsprinterManualReset,
		"printernmsreset":                           PrinterStatusDetailsprinterNmsReset,
		"printerreadytoprint":                       PrinterStatusDetailsprinterReadyToPrint,
		"puncheradded":                              PrinterStatusDetailspuncherAdded,
		"puncheralmostempty":                        PrinterStatusDetailspuncherAlmostEmpty,
		"puncheralmostfull":                         PrinterStatusDetailspuncherAlmostFull,
		"puncheratlimit":                            PrinterStatusDetailspuncherAtLimit,
		"puncherclosed":                             PrinterStatusDetailspuncherClosed,
		"puncherconfigurationchange":                PrinterStatusDetailspuncherConfigurationChange,
		"punchercoverclosed":                        PrinterStatusDetailspuncherCoverClosed,
		"punchercoveropen":                          PrinterStatusDetailspuncherCoverOpen,
		"puncherempty":                              PrinterStatusDetailspuncherEmpty,
		"puncherfull":                               PrinterStatusDetailspuncherFull,
		"puncherinterlockclosed":                    PrinterStatusDetailspuncherInterlockClosed,
		"puncherinterlockopen":                      PrinterStatusDetailspuncherInterlockOpen,
		"puncherjam":                                PrinterStatusDetailspuncherJam,
		"puncherlifealmostover":                     PrinterStatusDetailspuncherLifeAlmostOver,
		"puncherlifeover":                           PrinterStatusDetailspuncherLifeOver,
		"punchermemoryexhausted":                    PrinterStatusDetailspuncherMemoryExhausted,
		"punchermissing":                            PrinterStatusDetailspuncherMissing,
		"punchermotorfailure":                       PrinterStatusDetailspuncherMotorFailure,
		"punchernearlimit":                          PrinterStatusDetailspuncherNearLimit,
		"puncheroffline":                            PrinterStatusDetailspuncherOffline,
		"puncheropened":                             PrinterStatusDetailspuncherOpened,
		"puncherovertemperature":                    PrinterStatusDetailspuncherOverTemperature,
		"puncherpowersaver":                         PrinterStatusDetailspuncherPowerSaver,
		"puncherrecoverablefailure":                 PrinterStatusDetailspuncherRecoverableFailure,
		"puncherrecoverablestorage":                 PrinterStatusDetailspuncherRecoverableStorage,
		"puncherremoved":                            PrinterStatusDetailspuncherRemoved,
		"puncherresourceadded":                      PrinterStatusDetailspuncherResourceAdded,
		"puncherresourceremoved":                    PrinterStatusDetailspuncherResourceRemoved,
		"puncherthermistorfailure":                  PrinterStatusDetailspuncherThermistorFailure,
		"punchertimingfailure":                      PrinterStatusDetailspuncherTimingFailure,
		"puncherturnedoff":                          PrinterStatusDetailspuncherTurnedOff,
		"puncherturnedon":                           PrinterStatusDetailspuncherTurnedOn,
		"puncherundertemperature":                   PrinterStatusDetailspuncherUnderTemperature,
		"puncherunrecoverablefailure":               PrinterStatusDetailspuncherUnrecoverableFailure,
		"puncherunrecoverablestorageerror":          PrinterStatusDetailspuncherUnrecoverableStorageError,
		"puncherwarmingup":                          PrinterStatusDetailspuncherWarmingUp,
		"resuming":                                  PrinterStatusDetailsresuming,
		"scanmediapathfailure":                      PrinterStatusDetailsscanMediaPathFailure,
		"scanmediapathinputempty":                   PrinterStatusDetailsscanMediaPathInputEmpty,
		"scanmediapathinputfeederror":               PrinterStatusDetailsscanMediaPathInputFeedError,
		"scanmediapathinputjam":                     PrinterStatusDetailsscanMediaPathInputJam,
		"scanmediapathinputrequest":                 PrinterStatusDetailsscanMediaPathInputRequest,
		"scanmediapathjam":                          PrinterStatusDetailsscanMediaPathJam,
		"scanmediapathoutputfeederror":              PrinterStatusDetailsscanMediaPathOutputFeedError,
		"scanmediapathoutputfull":                   PrinterStatusDetailsscanMediaPathOutputFull,
		"scanmediapathoutputjam":                    PrinterStatusDetailsscanMediaPathOutputJam,
		"scanmediapathpickrollerfailure":            PrinterStatusDetailsscanMediaPathPickRollerFailure,
		"scanmediapathpickrollerlifeover":           PrinterStatusDetailsscanMediaPathPickRollerLifeOver,
		"scanmediapathpickrollerlifewarn":           PrinterStatusDetailsscanMediaPathPickRollerLifeWarn,
		"scanmediapathpickrollermissing":            PrinterStatusDetailsscanMediaPathPickRollerMissing,
		"scanmediapathtrayalmostfull":               PrinterStatusDetailsscanMediaPathTrayAlmostFull,
		"scanmediapathtrayfull":                     PrinterStatusDetailsscanMediaPathTrayFull,
		"scanmediapathtraymissing":                  PrinterStatusDetailsscanMediaPathTrayMissing,
		"scannerlightfailure":                       PrinterStatusDetailsscannerLightFailure,
		"scannerlightlifealmostover":                PrinterStatusDetailsscannerLightLifeAlmostOver,
		"scannerlightlifeover":                      PrinterStatusDetailsscannerLightLifeOver,
		"scannerlightmissing":                       PrinterStatusDetailsscannerLightMissing,
		"scannersensorfailure":                      PrinterStatusDetailsscannerSensorFailure,
		"scannersensorlifealmostover":               PrinterStatusDetailsscannerSensorLifeAlmostOver,
		"scannersensorlifeover":                     PrinterStatusDetailsscannerSensorLifeOver,
		"scannersensormissing":                      PrinterStatusDetailsscannerSensorMissing,
		"separationcutteradded":                     PrinterStatusDetailsseparationCutterAdded,
		"separationcutteralmostempty":               PrinterStatusDetailsseparationCutterAlmostEmpty,
		"separationcutteralmostfull":                PrinterStatusDetailsseparationCutterAlmostFull,
		"separationcutteratlimit":                   PrinterStatusDetailsseparationCutterAtLimit,
		"separationcutterclosed":                    PrinterStatusDetailsseparationCutterClosed,
		"separationcutterconfigurationchange":       PrinterStatusDetailsseparationCutterConfigurationChange,
		"separationcuttercoverclosed":               PrinterStatusDetailsseparationCutterCoverClosed,
		"separationcuttercoveropen":                 PrinterStatusDetailsseparationCutterCoverOpen,
		"separationcutterempty":                     PrinterStatusDetailsseparationCutterEmpty,
		"separationcutterfull":                      PrinterStatusDetailsseparationCutterFull,
		"separationcutterinterlockclosed":           PrinterStatusDetailsseparationCutterInterlockClosed,
		"separationcutterinterlockopen":             PrinterStatusDetailsseparationCutterInterlockOpen,
		"separationcutterjam":                       PrinterStatusDetailsseparationCutterJam,
		"separationcutterlifealmostover":            PrinterStatusDetailsseparationCutterLifeAlmostOver,
		"separationcutterlifeover":                  PrinterStatusDetailsseparationCutterLifeOver,
		"separationcuttermemoryexhausted":           PrinterStatusDetailsseparationCutterMemoryExhausted,
		"separationcuttermissing":                   PrinterStatusDetailsseparationCutterMissing,
		"separationcuttermotorfailure":              PrinterStatusDetailsseparationCutterMotorFailure,
		"separationcutternearlimit":                 PrinterStatusDetailsseparationCutterNearLimit,
		"separationcutteroffline":                   PrinterStatusDetailsseparationCutterOffline,
		"separationcutteropened":                    PrinterStatusDetailsseparationCutterOpened,
		"separationcutterovertemperature":           PrinterStatusDetailsseparationCutterOverTemperature,
		"separationcutterpowersaver":                PrinterStatusDetailsseparationCutterPowerSaver,
		"separationcutterrecoverablefailure":        PrinterStatusDetailsseparationCutterRecoverableFailure,
		"separationcutterrecoverablestorage":        PrinterStatusDetailsseparationCutterRecoverableStorage,
		"separationcutterremoved":                   PrinterStatusDetailsseparationCutterRemoved,
		"separationcutterresourceadded":             PrinterStatusDetailsseparationCutterResourceAdded,
		"separationcutterresourceremoved":           PrinterStatusDetailsseparationCutterResourceRemoved,
		"separationcutterthermistorfailure":         PrinterStatusDetailsseparationCutterThermistorFailure,
		"separationcuttertimingfailure":             PrinterStatusDetailsseparationCutterTimingFailure,
		"separationcutterturnedoff":                 PrinterStatusDetailsseparationCutterTurnedOff,
		"separationcutterturnedon":                  PrinterStatusDetailsseparationCutterTurnedOn,
		"separationcutterundertemperature":          PrinterStatusDetailsseparationCutterUnderTemperature,
		"separationcutterunrecoverablefailure":      PrinterStatusDetailsseparationCutterUnrecoverableFailure,
		"separationcutterunrecoverablestorageerror": PrinterStatusDetailsseparationCutterUnrecoverableStorageError,
		"separationcutterwarmingup":                 PrinterStatusDetailsseparationCutterWarmingUp,
		"sheetrotatoradded":                         PrinterStatusDetailssheetRotatorAdded,
		"sheetrotatoralmostempty":                   PrinterStatusDetailssheetRotatorAlmostEmpty,
		"sheetrotatoralmostfull":                    PrinterStatusDetailssheetRotatorAlmostFull,
		"sheetrotatoratlimit":                       PrinterStatusDetailssheetRotatorAtLimit,
		"sheetrotatorclosed":                        PrinterStatusDetailssheetRotatorClosed,
		"sheetrotatorconfigurationchange":           PrinterStatusDetailssheetRotatorConfigurationChange,
		"sheetrotatorcoverclosed":                   PrinterStatusDetailssheetRotatorCoverClosed,
		"sheetrotatorcoveropen":                     PrinterStatusDetailssheetRotatorCoverOpen,
		"sheetrotatorempty":                         PrinterStatusDetailssheetRotatorEmpty,
		"sheetrotatorfull":                          PrinterStatusDetailssheetRotatorFull,
		"sheetrotatorinterlockclosed":               PrinterStatusDetailssheetRotatorInterlockClosed,
		"sheetrotatorinterlockopen":                 PrinterStatusDetailssheetRotatorInterlockOpen,
		"sheetrotatorjam":                           PrinterStatusDetailssheetRotatorJam,
		"sheetrotatorlifealmostover":                PrinterStatusDetailssheetRotatorLifeAlmostOver,
		"sheetrotatorlifeover":                      PrinterStatusDetailssheetRotatorLifeOver,
		"sheetrotatormemoryexhausted":               PrinterStatusDetailssheetRotatorMemoryExhausted,
		"sheetrotatormissing":                       PrinterStatusDetailssheetRotatorMissing,
		"sheetrotatormotorfailure":                  PrinterStatusDetailssheetRotatorMotorFailure,
		"sheetrotatornearlimit":                     PrinterStatusDetailssheetRotatorNearLimit,
		"sheetrotatoroffline":                       PrinterStatusDetailssheetRotatorOffline,
		"sheetrotatoropened":                        PrinterStatusDetailssheetRotatorOpened,
		"sheetrotatorovertemperature":               PrinterStatusDetailssheetRotatorOverTemperature,
		"sheetrotatorpowersaver":                    PrinterStatusDetailssheetRotatorPowerSaver,
		"sheetrotatorrecoverablefailure":            PrinterStatusDetailssheetRotatorRecoverableFailure,
		"sheetrotatorrecoverablestorage":            PrinterStatusDetailssheetRotatorRecoverableStorage,
		"sheetrotatorremoved":                       PrinterStatusDetailssheetRotatorRemoved,
		"sheetrotatorresourceadded":                 PrinterStatusDetailssheetRotatorResourceAdded,
		"sheetrotatorresourceremoved":               PrinterStatusDetailssheetRotatorResourceRemoved,
		"sheetrotatorthermistorfailure":             PrinterStatusDetailssheetRotatorThermistorFailure,
		"sheetrotatortimingfailure":                 PrinterStatusDetailssheetRotatorTimingFailure,
		"sheetrotatorturnedoff":                     PrinterStatusDetailssheetRotatorTurnedOff,
		"sheetrotatorturnedon":                      PrinterStatusDetailssheetRotatorTurnedOn,
		"sheetrotatorundertemperature":              PrinterStatusDetailssheetRotatorUnderTemperature,
		"sheetrotatorunrecoverablefailure":          PrinterStatusDetailssheetRotatorUnrecoverableFailure,
		"sheetrotatorunrecoverablestorageerror":     PrinterStatusDetailssheetRotatorUnrecoverableStorageError,
		"sheetrotatorwarmingup":                     PrinterStatusDetailssheetRotatorWarmingUp,
		"shutdown":                                  PrinterStatusDetailsshutdown,
		"slitteradded":                              PrinterStatusDetailsslitterAdded,
		"slitteralmostempty":                        PrinterStatusDetailsslitterAlmostEmpty,
		"slitteralmostfull":                         PrinterStatusDetailsslitterAlmostFull,
		"slitteratlimit":                            PrinterStatusDetailsslitterAtLimit,
		"slitterclosed":                             PrinterStatusDetailsslitterClosed,
		"slitterconfigurationchange":                PrinterStatusDetailsslitterConfigurationChange,
		"slittercoverclosed":                        PrinterStatusDetailsslitterCoverClosed,
		"slittercoveropen":                          PrinterStatusDetailsslitterCoverOpen,
		"slitterempty":                              PrinterStatusDetailsslitterEmpty,
		"slitterfull":                               PrinterStatusDetailsslitterFull,
		"slitterinterlockclosed":                    PrinterStatusDetailsslitterInterlockClosed,
		"slitterinterlockopen":                      PrinterStatusDetailsslitterInterlockOpen,
		"slitterjam":                                PrinterStatusDetailsslitterJam,
		"slitterlifealmostover":                     PrinterStatusDetailsslitterLifeAlmostOver,
		"slitterlifeover":                           PrinterStatusDetailsslitterLifeOver,
		"slittermemoryexhausted":                    PrinterStatusDetailsslitterMemoryExhausted,
		"slittermissing":                            PrinterStatusDetailsslitterMissing,
		"slittermotorfailure":                       PrinterStatusDetailsslitterMotorFailure,
		"slitternearlimit":                          PrinterStatusDetailsslitterNearLimit,
		"slitteroffline":                            PrinterStatusDetailsslitterOffline,
		"slitteropened":                             PrinterStatusDetailsslitterOpened,
		"slitterovertemperature":                    PrinterStatusDetailsslitterOverTemperature,
		"slitterpowersaver":                         PrinterStatusDetailsslitterPowerSaver,
		"slitterrecoverablefailure":                 PrinterStatusDetailsslitterRecoverableFailure,
		"slitterrecoverablestorage":                 PrinterStatusDetailsslitterRecoverableStorage,
		"slitterremoved":                            PrinterStatusDetailsslitterRemoved,
		"slitterresourceadded":                      PrinterStatusDetailsslitterResourceAdded,
		"slitterresourceremoved":                    PrinterStatusDetailsslitterResourceRemoved,
		"slitterthermistorfailure":                  PrinterStatusDetailsslitterThermistorFailure,
		"slittertimingfailure":                      PrinterStatusDetailsslitterTimingFailure,
		"slitterturnedoff":                          PrinterStatusDetailsslitterTurnedOff,
		"slitterturnedon":                           PrinterStatusDetailsslitterTurnedOn,
		"slitterundertemperature":                   PrinterStatusDetailsslitterUnderTemperature,
		"slitterunrecoverablefailure":               PrinterStatusDetailsslitterUnrecoverableFailure,
		"slitterunrecoverablestorageerror":          PrinterStatusDetailsslitterUnrecoverableStorageError,
		"slitterwarmingup":                          PrinterStatusDetailsslitterWarmingUp,
		"spoolareafull":                             PrinterStatusDetailsspoolAreaFull,
		"stackeradded":                              PrinterStatusDetailsstackerAdded,
		"stackeralmostempty":                        PrinterStatusDetailsstackerAlmostEmpty,
		"stackeralmostfull":                         PrinterStatusDetailsstackerAlmostFull,
		"stackeratlimit":                            PrinterStatusDetailsstackerAtLimit,
		"stackerclosed":                             PrinterStatusDetailsstackerClosed,
		"stackerconfigurationchange":                PrinterStatusDetailsstackerConfigurationChange,
		"stackercoverclosed":                        PrinterStatusDetailsstackerCoverClosed,
		"stackercoveropen":                          PrinterStatusDetailsstackerCoverOpen,
		"stackerempty":                              PrinterStatusDetailsstackerEmpty,
		"stackerfull":                               PrinterStatusDetailsstackerFull,
		"stackerinterlockclosed":                    PrinterStatusDetailsstackerInterlockClosed,
		"stackerinterlockopen":                      PrinterStatusDetailsstackerInterlockOpen,
		"stackerjam":                                PrinterStatusDetailsstackerJam,
		"stackerlifealmostover":                     PrinterStatusDetailsstackerLifeAlmostOver,
		"stackerlifeover":                           PrinterStatusDetailsstackerLifeOver,
		"stackermemoryexhausted":                    PrinterStatusDetailsstackerMemoryExhausted,
		"stackermissing":                            PrinterStatusDetailsstackerMissing,
		"stackermotorfailure":                       PrinterStatusDetailsstackerMotorFailure,
		"stackernearlimit":                          PrinterStatusDetailsstackerNearLimit,
		"stackeroffline":                            PrinterStatusDetailsstackerOffline,
		"stackeropened":                             PrinterStatusDetailsstackerOpened,
		"stackerovertemperature":                    PrinterStatusDetailsstackerOverTemperature,
		"stackerpowersaver":                         PrinterStatusDetailsstackerPowerSaver,
		"stackerrecoverablefailure":                 PrinterStatusDetailsstackerRecoverableFailure,
		"stackerrecoverablestorage":                 PrinterStatusDetailsstackerRecoverableStorage,
		"stackerremoved":                            PrinterStatusDetailsstackerRemoved,
		"stackerresourceadded":                      PrinterStatusDetailsstackerResourceAdded,
		"stackerresourceremoved":                    PrinterStatusDetailsstackerResourceRemoved,
		"stackerthermistorfailure":                  PrinterStatusDetailsstackerThermistorFailure,
		"stackertimingfailure":                      PrinterStatusDetailsstackerTimingFailure,
		"stackerturnedoff":                          PrinterStatusDetailsstackerTurnedOff,
		"stackerturnedon":                           PrinterStatusDetailsstackerTurnedOn,
		"stackerundertemperature":                   PrinterStatusDetailsstackerUnderTemperature,
		"stackerunrecoverablefailure":               PrinterStatusDetailsstackerUnrecoverableFailure,
		"stackerunrecoverablestorageerror":          PrinterStatusDetailsstackerUnrecoverableStorageError,
		"stackerwarmingup":                          PrinterStatusDetailsstackerWarmingUp,
		"standby":                                   PrinterStatusDetailsstandby,
		"stapleradded":                              PrinterStatusDetailsstaplerAdded,
		"stapleralmostempty":                        PrinterStatusDetailsstaplerAlmostEmpty,
		"stapleralmostfull":                         PrinterStatusDetailsstaplerAlmostFull,
		"stapleratlimit":                            PrinterStatusDetailsstaplerAtLimit,
		"staplerclosed":                             PrinterStatusDetailsstaplerClosed,
		"staplerconfigurationchange":                PrinterStatusDetailsstaplerConfigurationChange,
		"staplercoverclosed":                        PrinterStatusDetailsstaplerCoverClosed,
		"staplercoveropen":                          PrinterStatusDetailsstaplerCoverOpen,
		"staplerempty":                              PrinterStatusDetailsstaplerEmpty,
		"staplerfull":                               PrinterStatusDetailsstaplerFull,
		"staplerinterlockclosed":                    PrinterStatusDetailsstaplerInterlockClosed,
		"staplerinterlockopen":                      PrinterStatusDetailsstaplerInterlockOpen,
		"staplerjam":                                PrinterStatusDetailsstaplerJam,
		"staplerlifealmostover":                     PrinterStatusDetailsstaplerLifeAlmostOver,
		"staplerlifeover":                           PrinterStatusDetailsstaplerLifeOver,
		"staplermemoryexhausted":                    PrinterStatusDetailsstaplerMemoryExhausted,
		"staplermissing":                            PrinterStatusDetailsstaplerMissing,
		"staplermotorfailure":                       PrinterStatusDetailsstaplerMotorFailure,
		"staplernearlimit":                          PrinterStatusDetailsstaplerNearLimit,
		"stapleroffline":                            PrinterStatusDetailsstaplerOffline,
		"stapleropened":                             PrinterStatusDetailsstaplerOpened,
		"staplerovertemperature":                    PrinterStatusDetailsstaplerOverTemperature,
		"staplerpowersaver":                         PrinterStatusDetailsstaplerPowerSaver,
		"staplerrecoverablefailure":                 PrinterStatusDetailsstaplerRecoverableFailure,
		"staplerrecoverablestorage":                 PrinterStatusDetailsstaplerRecoverableStorage,
		"staplerremoved":                            PrinterStatusDetailsstaplerRemoved,
		"staplerresourceadded":                      PrinterStatusDetailsstaplerResourceAdded,
		"staplerresourceremoved":                    PrinterStatusDetailsstaplerResourceRemoved,
		"staplerthermistorfailure":                  PrinterStatusDetailsstaplerThermistorFailure,
		"staplertimingfailure":                      PrinterStatusDetailsstaplerTimingFailure,
		"staplerturnedoff":                          PrinterStatusDetailsstaplerTurnedOff,
		"staplerturnedon":                           PrinterStatusDetailsstaplerTurnedOn,
		"staplerundertemperature":                   PrinterStatusDetailsstaplerUnderTemperature,
		"staplerunrecoverablefailure":               PrinterStatusDetailsstaplerUnrecoverableFailure,
		"staplerunrecoverablestorageerror":          PrinterStatusDetailsstaplerUnrecoverableStorageError,
		"staplerwarmingup":                          PrinterStatusDetailsstaplerWarmingUp,
		"stitcheradded":                             PrinterStatusDetailsstitcherAdded,
		"stitcheralmostempty":                       PrinterStatusDetailsstitcherAlmostEmpty,
		"stitcheralmostfull":                        PrinterStatusDetailsstitcherAlmostFull,
		"stitcheratlimit":                           PrinterStatusDetailsstitcherAtLimit,
		"stitcherclosed":                            PrinterStatusDetailsstitcherClosed,
		"stitcherconfigurationchange":               PrinterStatusDetailsstitcherConfigurationChange,
		"stitchercoverclosed":                       PrinterStatusDetailsstitcherCoverClosed,
		"stitchercoveropen":                         PrinterStatusDetailsstitcherCoverOpen,
		"stitcherempty":                             PrinterStatusDetailsstitcherEmpty,
		"stitcherfull":                              PrinterStatusDetailsstitcherFull,
		"stitcherinterlockclosed":                   PrinterStatusDetailsstitcherInterlockClosed,
		"stitcherinterlockopen":                     PrinterStatusDetailsstitcherInterlockOpen,
		"stitcherjam":                               PrinterStatusDetailsstitcherJam,
		"stitcherlifealmostover":                    PrinterStatusDetailsstitcherLifeAlmostOver,
		"stitcherlifeover":                          PrinterStatusDetailsstitcherLifeOver,
		"stitchermemoryexhausted":                   PrinterStatusDetailsstitcherMemoryExhausted,
		"stitchermissing":                           PrinterStatusDetailsstitcherMissing,
		"stitchermotorfailure":                      PrinterStatusDetailsstitcherMotorFailure,
		"stitchernearlimit":                         PrinterStatusDetailsstitcherNearLimit,
		"stitcheroffline":                           PrinterStatusDetailsstitcherOffline,
		"stitcheropened":                            PrinterStatusDetailsstitcherOpened,
		"stitcherovertemperature":                   PrinterStatusDetailsstitcherOverTemperature,
		"stitcherpowersaver":                        PrinterStatusDetailsstitcherPowerSaver,
		"stitcherrecoverablefailure":                PrinterStatusDetailsstitcherRecoverableFailure,
		"stitcherrecoverablestorage":                PrinterStatusDetailsstitcherRecoverableStorage,
		"stitcherremoved":                           PrinterStatusDetailsstitcherRemoved,
		"stitcherresourceadded":                     PrinterStatusDetailsstitcherResourceAdded,
		"stitcherresourceremoved":                   PrinterStatusDetailsstitcherResourceRemoved,
		"stitcherthermistorfailure":                 PrinterStatusDetailsstitcherThermistorFailure,
		"stitchertimingfailure":                     PrinterStatusDetailsstitcherTimingFailure,
		"stitcherturnedoff":                         PrinterStatusDetailsstitcherTurnedOff,
		"stitcherturnedon":                          PrinterStatusDetailsstitcherTurnedOn,
		"stitcherundertemperature":                  PrinterStatusDetailsstitcherUnderTemperature,
		"stitcherunrecoverablefailure":              PrinterStatusDetailsstitcherUnrecoverableFailure,
		"stitcherunrecoverablestorageerror":         PrinterStatusDetailsstitcherUnrecoverableStorageError,
		"stitcherwarmingup":                         PrinterStatusDetailsstitcherWarmingUp,
		"stoppedpartially":                          PrinterStatusDetailsstoppedPartially,
		"stopping":                                  PrinterStatusDetailsstopping,
		"subunitadded":                              PrinterStatusDetailssubunitAdded,
		"subunitalmostempty":                        PrinterStatusDetailssubunitAlmostEmpty,
		"subunitalmostfull":                         PrinterStatusDetailssubunitAlmostFull,
		"subunitatlimit":                            PrinterStatusDetailssubunitAtLimit,
		"subunitclosed":                             PrinterStatusDetailssubunitClosed,
		"subunitcoolingdown":                        PrinterStatusDetailssubunitCoolingDown,
		"subunitempty":                              PrinterStatusDetailssubunitEmpty,
		"subunitfull":                               PrinterStatusDetailssubunitFull,
		"subunitlifealmostover":                     PrinterStatusDetailssubunitLifeAlmostOver,
		"subunitlifeover":                           PrinterStatusDetailssubunitLifeOver,
		"subunitmemoryexhausted":                    PrinterStatusDetailssubunitMemoryExhausted,
		"subunitmissing":                            PrinterStatusDetailssubunitMissing,
		"subunitmotorfailure":                       PrinterStatusDetailssubunitMotorFailure,
		"subunitnearlimit":                          PrinterStatusDetailssubunitNearLimit,
		"subunitoffline":                            PrinterStatusDetailssubunitOffline,
		"subunitopened":                             PrinterStatusDetailssubunitOpened,
		"subunitovertemperature":                    PrinterStatusDetailssubunitOverTemperature,
		"subunitpowersaver":                         PrinterStatusDetailssubunitPowerSaver,
		"subunitrecoverablefailure":                 PrinterStatusDetailssubunitRecoverableFailure,
		"subunitrecoverablestorage":                 PrinterStatusDetailssubunitRecoverableStorage,
		"subunitremoved":                            PrinterStatusDetailssubunitRemoved,
		"subunitresourceadded":                      PrinterStatusDetailssubunitResourceAdded,
		"subunitresourceremoved":                    PrinterStatusDetailssubunitResourceRemoved,
		"subunitthermistorfailure":                  PrinterStatusDetailssubunitThermistorFailure,
		"subunittimingfailure":                      PrinterStatusDetailssubunitTimingFailure,
		"subunitturnedoff":                          PrinterStatusDetailssubunitTurnedOff,
		"subunitturnedon":                           PrinterStatusDetailssubunitTurnedOn,
		"subunitundertemperature":                   PrinterStatusDetailssubunitUnderTemperature,
		"subunitunrecoverablefailure":               PrinterStatusDetailssubunitUnrecoverableFailure,
		"subunitunrecoverablestorage":               PrinterStatusDetailssubunitUnrecoverableStorage,
		"subunitwarmingup":                          PrinterStatusDetailssubunitWarmingUp,
		"suspend":                                   PrinterStatusDetailssuspend,
		"testing":                                   PrinterStatusDetailstesting,
		"timedout":                                  PrinterStatusDetailstimedOut,
		"tonerempty":                                PrinterStatusDetailstonerEmpty,
		"tonerlow":                                  PrinterStatusDetailstonerLow,
		"trimmeradded":                              PrinterStatusDetailstrimmerAdded,
		"trimmeralmostempty":                        PrinterStatusDetailstrimmerAlmostEmpty,
		"trimmeralmostfull":                         PrinterStatusDetailstrimmerAlmostFull,
		"trimmeratlimit":                            PrinterStatusDetailstrimmerAtLimit,
		"trimmerclosed":                             PrinterStatusDetailstrimmerClosed,
		"trimmerconfigurationchange":                PrinterStatusDetailstrimmerConfigurationChange,
		"trimmercoverclosed":                        PrinterStatusDetailstrimmerCoverClosed,
		"trimmercoveropen":                          PrinterStatusDetailstrimmerCoverOpen,
		"trimmerempty":                              PrinterStatusDetailstrimmerEmpty,
		"trimmerfull":                               PrinterStatusDetailstrimmerFull,
		"trimmerinterlockclosed":                    PrinterStatusDetailstrimmerInterlockClosed,
		"trimmerinterlockopen":                      PrinterStatusDetailstrimmerInterlockOpen,
		"trimmerjam":                                PrinterStatusDetailstrimmerJam,
		"trimmerlifealmostover":                     PrinterStatusDetailstrimmerLifeAlmostOver,
		"trimmerlifeover":                           PrinterStatusDetailstrimmerLifeOver,
		"trimmermemoryexhausted":                    PrinterStatusDetailstrimmerMemoryExhausted,
		"trimmermissing":                            PrinterStatusDetailstrimmerMissing,
		"trimmermotorfailure":                       PrinterStatusDetailstrimmerMotorFailure,
		"trimmernearlimit":                          PrinterStatusDetailstrimmerNearLimit,
		"trimmeroffline":                            PrinterStatusDetailstrimmerOffline,
		"trimmeropened":                             PrinterStatusDetailstrimmerOpened,
		"trimmerovertemperature":                    PrinterStatusDetailstrimmerOverTemperature,
		"trimmerpowersaver":                         PrinterStatusDetailstrimmerPowerSaver,
		"trimmerrecoverablefailure":                 PrinterStatusDetailstrimmerRecoverableFailure,
		"trimmerrecoverablestorage":                 PrinterStatusDetailstrimmerRecoverableStorage,
		"trimmerremoved":                            PrinterStatusDetailstrimmerRemoved,
		"trimmerresourceadded":                      PrinterStatusDetailstrimmerResourceAdded,
		"trimmerresourceremoved":                    PrinterStatusDetailstrimmerResourceRemoved,
		"trimmerthermistorfailure":                  PrinterStatusDetailstrimmerThermistorFailure,
		"trimmertimingfailure":                      PrinterStatusDetailstrimmerTimingFailure,
		"trimmerturnedoff":                          PrinterStatusDetailstrimmerTurnedOff,
		"trimmerturnedon":                           PrinterStatusDetailstrimmerTurnedOn,
		"trimmerundertemperature":                   PrinterStatusDetailstrimmerUnderTemperature,
		"trimmerunrecoverablefailure":               PrinterStatusDetailstrimmerUnrecoverableFailure,
		"trimmerunrecoverablestorageerror":          PrinterStatusDetailstrimmerUnrecoverableStorageError,
		"trimmerwarmingup":                          PrinterStatusDetailstrimmerWarmingUp,
		"unknown":                                   PrinterStatusDetailsunknown,
		"wrapperadded":                              PrinterStatusDetailswrapperAdded,
		"wrapperalmostempty":                        PrinterStatusDetailswrapperAlmostEmpty,
		"wrapperalmostfull":                         PrinterStatusDetailswrapperAlmostFull,
		"wrapperatlimit":                            PrinterStatusDetailswrapperAtLimit,
		"wrapperclosed":                             PrinterStatusDetailswrapperClosed,
		"wrapperconfigurationchange":                PrinterStatusDetailswrapperConfigurationChange,
		"wrappercoverclosed":                        PrinterStatusDetailswrapperCoverClosed,
		"wrappercoveropen":                          PrinterStatusDetailswrapperCoverOpen,
		"wrapperempty":                              PrinterStatusDetailswrapperEmpty,
		"wrapperfull":                               PrinterStatusDetailswrapperFull,
		"wrapperinterlockclosed":                    PrinterStatusDetailswrapperInterlockClosed,
		"wrapperinterlockopen":                      PrinterStatusDetailswrapperInterlockOpen,
		"wrapperjam":                                PrinterStatusDetailswrapperJam,
		"wrapperlifealmostover":                     PrinterStatusDetailswrapperLifeAlmostOver,
		"wrapperlifeover":                           PrinterStatusDetailswrapperLifeOver,
		"wrappermemoryexhausted":                    PrinterStatusDetailswrapperMemoryExhausted,
		"wrappermissing":                            PrinterStatusDetailswrapperMissing,
		"wrappermotorfailure":                       PrinterStatusDetailswrapperMotorFailure,
		"wrappernearlimit":                          PrinterStatusDetailswrapperNearLimit,
		"wrapperoffline":                            PrinterStatusDetailswrapperOffline,
		"wrapperopened":                             PrinterStatusDetailswrapperOpened,
		"wrapperovertemperature":                    PrinterStatusDetailswrapperOverTemperature,
		"wrapperpowersaver":                         PrinterStatusDetailswrapperPowerSaver,
		"wrapperrecoverablefailure":                 PrinterStatusDetailswrapperRecoverableFailure,
		"wrapperrecoverablestorage":                 PrinterStatusDetailswrapperRecoverableStorage,
		"wrapperremoved":                            PrinterStatusDetailswrapperRemoved,
		"wrapperresourceadded":                      PrinterStatusDetailswrapperResourceAdded,
		"wrapperresourceremoved":                    PrinterStatusDetailswrapperResourceRemoved,
		"wrapperthermistorfailure":                  PrinterStatusDetailswrapperThermistorFailure,
		"wrappertimingfailure":                      PrinterStatusDetailswrapperTimingFailure,
		"wrapperturnedoff":                          PrinterStatusDetailswrapperTurnedOff,
		"wrapperturnedon":                           PrinterStatusDetailswrapperTurnedOn,
		"wrapperundertemperature":                   PrinterStatusDetailswrapperUnderTemperature,
		"wrapperunrecoverablefailure":               PrinterStatusDetailswrapperUnrecoverableFailure,
		"wrapperunrecoverablestorageerror":          PrinterStatusDetailswrapperUnrecoverableStorageError,
		"wrapperwarmingup":                          PrinterStatusDetailswrapperWarmingUp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterStatusDetails(input)
	return &out, nil
}
