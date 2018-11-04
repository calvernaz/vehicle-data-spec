package gofakeit

import "time"
import (
	"github.com/calvernaz/w3c-vehicle-data/types/transmission-gear"
	"github.com/calvernaz/w3c-vehicle-data/types/vehicle-type"
	"github.com/calvernaz/w3c-vehicle-data/types/zone"
	"github.com/calvernaz/w3c-vehicle-data/types/transmission-mode"
	"github.com/calvernaz/w3c-vehicle-data/types/fuel-type"
	"github.com/calvernaz/w3c-vehicle-data/types/parking-brake-status"
	"github.com/calvernaz/w3c-vehicle-data/types/alarm-status"
	"github.com/calvernaz/w3c-vehicle-data/types/lane-departure-status"
	"github.com/calvernaz/w3c-vehicle-data/types/convertible-root-status"
	"github.com/calvernaz/w3c-vehicle-data/types/wiper-control"
	"github.com/calvernaz/w3c-vehicle-data/types/identification-type"
	"github.com/calvernaz/w3c-vehicle-data/types/door-open-status"
	"github.com/calvernaz/w3c-vehicle-data/types/driver-mode"
	"github.com/calvernaz/w3c-vehicle-data/types/button-event"
	"github.com/calvernaz/w3c-vehicle-data/types/occupant-status"
	"github.com/calvernaz/w3c-vehicle-data/types/airflow-direction"
)

// Vehicle Data
// https://rawgit.com/w3c/automotive-bg/master/data_spec.html

//
// Configuration and Identification interfaces
//

// The Identification interface provides identification information about a vehicle.
type Identification struct {
	// Vehicle Identification Number (ISO 3833)
	VIN string
	// World Manufacturer Identifier defined by SAE ISO 3780:2009. 3 characters.
	WMI string
	// Vehicle type
	VehicleType vehicle_type.VehicleType
	// Brand name
	Brand string
	// Vehicle model
	Model string
	// Vehicle model year
	Year uint16
}

// The SizeConfiguration interface provides size and shape information about a vehicle as a whole.
type SizeConfiguration struct {
	// Widest dimension of the vehicle (not including the side mirrors) (Unit: millimeters Note: Number may be an
	// approximation, and should not be expected to be exact.)
	Width uint16
	// Distance from the ground to the highest point of the vehicle (not including antennas)
	// (Unit: millimeters Note: Number may be an approximation, and should not be expected to be exact.)
	Height uint16
	// Distance from front bumper to rear bumper (Unit: millimeters Note: Number may be an approximation,
	// and should not be expected to be exact.)
	Length uint16
	// List of car doors, organized in "rows" with number doors in each row.(Per Row - Min: 0, Max: 3)
	DoorsCount []uint16
	// Total number of doors on the vehicle (all doors opening to the interior, including hatchbacks) (Min: 0, Max: 10)
	TotalDoors uint16
}

// The FuelConfiguration interface provides information about the fuel configuration of a vehicle.
// A dictionary has been used to allow an associated array of values for vehicles which use multiple fuels.
type FuelConfiguration struct {
	//  Type of fuel used by vehicle. If the vehicle uses multiple fuels, fuelType returns an array of fuel types.
	FuelType []fuel_type.FuelType
	//  Location on the vehicle with access to the fuel door
	RefuelPosition zone.Zone
}

// The TransmissionConfiguration interface provides transmission-gear configuration information information about a vehicle.
type TransmissionConfiguration struct {
	// Transmission gear type
	TransmissionGearType transmission_gear.TransmissionGearType
}

// The WheelConfiguration interface provides wheel configuration information about a vehicle.
type WheelConfiguration struct {
	// Radius of the front wheel (Unit: millimeters)
	WheelRadius uint16
	// Zone for requested attribute
	Zone zone.Zone
}

// The SteeringWheelConfiguration interface provides steering wheel configuration information information about a
// vehicle.
type SteeringWheelConfiguration struct {
	// True if steering wheel is on left side of vehicle
	SteeringWheelLeft bool
	// Steering wheel position as percentage of extension from the dash (Unit: percentage, 0%:closest to dash,
	// 100%:farthest from dash)
	SteeringWheelTelescopingPosition uint16
	// Steering wheel position as percentage of tilt (Unit: percentage, 0%:tilted lowest downward-facing position,
	// 100%:highest upward-facing position)
	SteeringWheelPositionTilt uint16
}

//
// Running Status interfaces
//

// The VehicleSpeed interface represents vehicle speed information
type VehicleSpeed struct {
	// Vehicle speed (Unit: meters per hour)
	Speed uint16
}

// The WheelSpeed interface represents wheel speed information.
type WheelSpeed struct {
	// Wheel speed (Unit: meters per hour)
	Speed uint16
	// Zone for requested attribute
	Zone zone.Zone
}

// The EngineSpeed interface represents engine speed information.
type EngineSpeed struct {
	// Engine speed (Unit: rotations per minute)
	Speed uint64
}

// The VehiclePowerModeType interface represents position of the ignition switch.
type VehiclePowerModeType struct {
	// Position of the ignition switch
	Value vehicle_type.VehicleType
}

// The PowertrainTorque interface represents powertrain torque.
type PowertrainTorque struct {
	// Powertrain torque (Unit: newton meters)
	Value uint16
}

// The PowertrainTorque interface represents powertrain torque.
type AcceleratorPedalPosition struct {
	// Powertrain torque (Unit: newton meters)
	Value uint16
}

// The ThrottlePosition represents position of the throttle.
type ThrottlePosition struct {
	// Throttle position as a percentage (Unit: percentage, 0%: closed, 100%: fully open)
	Value uint16
}

// The Trip interface represents trip meter.
type Trip struct {
	// Distance travelled based on trip meter (Unit: meters)
	Distance uint64
	// Average speed based on trip meter (Unit: kilometers per hour)
	AverageSpeed uint16
	// Fuel consumed based on trip meter (Unit: milliliters per 100 kilometers)
	FuelConsumption uint16
	// Trip meters
	Meters []Trip
}

// The Transmission interface represents the current transmission-gear gear and mode.
type Transmission struct {
	// Transmission gear position. Range 0 - 10
	Gear byte
	// Transmission Mode (see TransmissionMode)
	Mode transmission_mode.TransmissionMode
}

// The CruiseControlStatus interface represents cruise control settings
type CruiseControlStatus struct {
	// Whether or not the Cruise Control system is on (true) or off (false)
	Status bool
	// Target Cruise Control speed in kilometers per hour (Unit: kilometers per hour)
	Speed uint16
}

// The LightStatus interface represents exterior light statuses.
type LightStatus struct {
	// Headlight status: on (true), off (false)
	Head bool
	// Right turn signal status: on (true), off (false)
	RightTurn bool
	// Left turn signal status: on (true), off (false)
	LeftTurn bool
	// Brake light status: on (true), off (false)
	Brake bool
	// Fog light status: on (true), off (false)
	Fog bool
	// Hazard light status: on (true), off (false)
	Hazard bool
	// Parking light status: on (true), off (false)
	Parking bool
	// HighBeam light status: on (true), off (false)
	HighBeam bool
	// Whether automatic head lights status: activated (true) or not (false)
	AutomaticHeadLights bool
	// Whether dynamic high beam status: activated (true) or not (false)
	DynamicHighBeam bool
	// Zone for requested attribute
	Zone zone.Zone
}

// The InteriorLightStatus interface represents interior light status.
type InteriorLightStatus struct {
	// Interior light status for the given zone: on (true), off (false)
	Status bool
	//  Zone for requested attribute
	Zone zone.Zone
}

// The Horn interface represents horn status.
type Horn struct {
	// Horn status: on (true) or off (false)
	Status bool
}

// The Chime interface represents chime status.
type Chime struct {
	// Chime status when a door is open: on (true) or off (false)
	Status bool
}

// The Fuel interface represents vehicle fuel status.
type Fuel struct {
	// Fuel level as a percentage of fullness
	Level uint16
	// Estimated fuel range (Unit: meters)
	Range uint64
	// Instant fuel consumption in per distance travelled (Unit: milliliters per 100 kilometers)
	InstantConsumption uint64
	// Average fuel consumption in per distance travelled (Unit: milliliters per 100 kilometers).
	// Setting this to any value should reset the counter to '0'
	AverageConsumption uint64
	// Fuel consumed since engine start; (Unit: milliliters per 100 kilometers) resets to 0 each restart
	FuelConsumedSinceRestart uint64
	// Time elapsed since vehicle restart (Unit: seconds)
	TimeSinceRestart uint64
}

// The EngineOil interface represents engine oil status
type EngineOil struct {
	// Engine oil level (Unit: percentage, 0%: empty, 100%: full
	Level uint16
	// Remaining engine oil life (Unit: percentage, 0%:no life remaining, 100%: full life remaining
	LifeRemaining uint16
	// Engine Oil Temperature (Unit: celcius)
	Temperature int64
	// Engine Oil Pressure (Unit: kilopascals)
	Pressure uint16
	// Engine oil change indicator status: change oil (true) or no change (false)
	Change bool
}

// The Acceleration interface represents vehicle acceleration
type Acceleration struct {
	// Acceleration on the "X" axis (Unit: centimeters per second squared)
	X int64
	// Acceleration on the "Y" axis (Unit: centimeters per second squared)
	Y int64
	// Acceleration on the "Z" axis (Unit: centimeters per second squared)
	Z int64
}

// The EngineCoolant represents values related to engine coolant.
type EngineCoolant struct {
	// Engine coolant level (Unit: percentage 0%: empty, 100%: full)
	Level byte
	// Engine coolant temperature (Unit: celcius)
	Temperature int16
}

// The SteeringWheel represents steering wheel data.
type SteeringWheel struct {
	// Angle of steering wheel off centerline (Unit: degrees -:degrees to the left, +:degrees to the right)
	Angle int16
}

// The WheelTick number of ticks per second.
type WheelTick struct {
	// Number of ticks per second (Unit: ticks per second)
	Value uint64
	// Zone for requested attribute
	Zone zone.Zone
}

// The IgnitionTime represents status of ignition.
type IgnitionTime struct {
	// Time at ignition on
	IgnitionOnTime time.Time
	// Time at ignition off
	IgnitionOffTime time.Time
}

// The YawRate represents vehicle yaw rate.
type YawRate struct {
	// Yaw rate of vehicle. (Unit: degrees per second)
	Value int16
}

// The BrakeOperation represents vehicle brake operation
type BrakeOperation struct {
	// Whether brake pedal is depressed or not. true: brake pedal is depressed, false: brake pedal is not depressed
	BrakePedalDepressed bool
}

// The ButtonEvent represents button press events from the steering wheel or other source
type ButtonEvent struct {
	// The type of event
	State button_event.ButtonEventType
}

// The DrivingMode interface provides information about whether or not the vehicle is driving.
// DrivingMode is an abstract data type that may combine several other data types such as vehicle speed,
// transmission-gear gear, etc. Typical usage would be to disable certain functions in the application if the vehicle is
// not safe to operate those functions to avoid driver distraction.
type DrivingMode struct {
	// True if vehicle is in driving mode
	Mode bool
}

// The NightMode interface provides information about whether or not it is night time.
// NightMode is an abstract data type that may combine several other data types such as exterior brightness,
// time of day, sunrise/sunset, etc to determine whether or not it is night time.
// Typical usage is to change the UI theme to a darker theme during the night
type NightMode struct {
	// True if it is night time
	Mode bool
}

//
// Maintenance interfaces
//

// The Odometer interface provides information about the distance that the vehicle has traveled
type Odometer struct {
	// The distance traveled by vehicle since start (Unit: meters).
	distanceSinceStart uint64
	// The total distance traveled by the vehicle (Unit: meters).
	distanceTotal uint64
}

// The TransmissionOil interface provides information about the state of a vehicles transmission-gear oil.
type TransmissionOil struct {
	// Transmission oil wear (Unit: percentage, 0: no wear, 100: completely worn).
	Wear byte
	// Current temperature of the transmission-gear oil(Unit: celsius)
	Temperature byte
}

// The TransmissionClutch interface provides information about the state of a vehicles transmission-gear clutch.
type TransmissionClutch struct {
	// Transmission clutch wear (Unit: percentage, 0%: no wear, 100%: completely worn).
	Wear byte
}

// The BrakeMaintenance interface provides information about the maintenance state of a vehicles brakes.
type BrakeMaintenance struct {
	// Brake fluid level (Unit: percentage, 0%: empty, 100%: full).
	FluidLevel byte
	// True if brake fluid level: low (true), not low (false)
	FluidLevelLow byte
	// Brake pad wear (Unit: percentage, 0%: no wear, 100%: completely worn).
	PadWear byte
	// True if brakes are worn: worn (true), not worn (false)
	BrakesWorn bool
	// Zone for requested attribute
	Zone zone.Zone
}

// The WasherFluid interface provides information about the state of a vehicles washer fluid
type WasherFluid struct {
	// Washer fluid level (Unit: percentage, 0%: empty, 100%: full).
	Level uint16
	// True if washer fluid level is low: low (true), not low: (false)
	LevelLow bool
}

// The MalfunctionIndicator interface provides information about the state of a vehicles Malfunction Indicator lamp.
type MalfunctionIndicator struct {
	// True if malfunction indicator lamp is on: lamp on (true), lamp not on (false)
	On bool
}

// The BatteryStatus interface provides information about the state of a vehicles battery
type BatteryStatus struct {
	// Battery charge level (Unit: percentage, 0%: empty, 100%: full).
	ChargeLevel byte
	// Battery voltage (Unit: volts).
	Voltage uint16
	// Battery current (Unit: amperes).
	Current uint16
	// Zone for requested attribute
	Zone zone.Zone
}

// The Tire interface provides information about the state of a vehicles tires
type Tire struct {
	// True if any tire pressure is low: pressure low (true), pressure not low (false)
	PressureLow bool
	// Tire pressure (Unit: kilopascal).
	Pressure uint16
	// Tire temperature (Unit: celsius).
	Temperature int16
	// Zone for requested attribute
	Zone zone.Zone
}

// The Diagnostic interface represents Diagnostic interface to malfunction indicator light information
type Diagnostic struct {
	// Engine runtime (Unit: seconds)
	AccumulatedEngineRuntime uint64
	// Distance travelled with the malfunction indicator light on (Unit: meters)
	DistanceWithMILOn uint64
	// Distance travelled since the codes were last cleared (Unit: meters)
	DistanceSinceCodeCleared uint64
	// Time elapsed with the malfunction indicator light on (Unit: seconds)
	TimeRunMILOn uint64
	// Time elapsed since the trouble codes were last cleared (Unit: seconds)
	TimeTroubleCodeClear uint64
}

//
// Personalization interfaces
//

// The LanguageConfiguration interface provides language information about a vehicle
type LanguageConfiguration struct {
	// Language identifier based on two-letter codes as specified in ISO 639-1
	Language string
}

// The UnitsOfMeasure interface provides information about the measurement system and units of measure of a vehicle.
type UnitsOfMeasure struct {
	// measurement system currently being used by vehicle. 'true' means the current measurement system is MKS-km(liter).
	// 'false' means it is US customary units-mile(gallon).
	IsMKSSystem bool
	// Fuel unit of measurement. The value is one of both "litter" and "gallon".
	UnitsFuelVolume string
	// Distance unit of measurement. The value is one of both "km" and "mile".
	UnitsDistance string
	// Speed unit of measurement. The value is one of both "km/h" and "mph".
	UnitsSpeed string
	// Fuel consumption unit of measurement. The value is one of following values: "l/100", "mpg", "km/l".
	UnitsFuelConsumption string
}

// The Mirror interface provides or sets information about mirrors in vehicle.
type Mirror struct {
	// Mirror tilt position in percentage distance travelled, from downward-facing to upward-facing position
	// (Unit: percentage, 0%:center position, -100%:fully downward, 100%:full upward)
	MirrorTilt byte
	// Mirror pan position in percentage distance travelled, from left to right position
	// (Unit: percentage, %0:center position, -100%:fully left, 100%:fully right)
	MirrorPan byte
	// Zone for requested attribute
	Zone zone.Zone
}

// The SeatAdjustment interface provides or sets information about seats in vehicle.
type SeatAdjustment struct {
	// Seat back recline position as percent to completely reclined
	// (Unit: percentage, 0%: fully forward, 100%: fully reclined)
	ReclineSeatBack byte
	// seat slide position as percentage of distance travelled away from forwardmost position
	// (Unit: percentage, 0%: farthest forward, 100%: farthest back)
	SeatSlide byte
	// Seat cushion height position as a percentage of upward distance travelled
	// (Unit: percentage, 0%: lowest. 100%: highest)
	SeatCushionHeight byte
	// Headrest position as a percentage of upward distance travelled (Unit: percentage, 0%: lowest, 100%: highest)
	SeatHeadrest byte
	// Back cushion position as a percentage of lumbar curvature (Unit: percentage, 0%: flat, 100%: maximum curvature)
	SeatBackCushion byte
	// Sides of back cushion position as a percentage of curvature (Unit: percentage, 0%: flat, 100%: maximum curvature)
	SeatSideCushion byte
	// Zone for requested attribute
	Zone zone.Zone
}

// The DriveMode interface provides or sets information about a vehicles drive mode.
type DriverMode struct {
	// Vehicle drive mode
	DriveMode driver_mode.DriveModeType
}

// The DashboardIllumination interface provides or sets information about dashboard illumination in vehicle
type DashboardIllumination struct {
	//  illumination of dashboard as a percentage (Unit: percentage, 0%: none, 100%: maximum illumination)
	DashboardIllumination byte
}

// The VehicleSound interface provides or sets information about vehicle sound
type VehicleSound struct {
	// Active noise control status: not-activated (false), activated (true)
	ActiveNoiseControlMode bool
	// Engine sound enhancement mode where a null string means not-activated, and any other value represents a
	// manufacture specific setting. See availableSounds.
	EngineSoundEnhancementMode string
	// Array of available sounds. See engineSoundEnhancementMode
	AvailableSounds []string
}

//
// DrivingSafety Interfaces
//

// The AntilockBrakingSystem interface provides status of ABS(Antilock Braking System) status and setting.
type AntilockBrakingSystem struct {
	// Whether or not the ABS Setting is enabled: enabled (true) or disabled (false)
	Enabled bool
	// Whether or not the ABS is engaged: engaged (true) or idle (false)
	Engaged bool
}

// The TractionControlSystem interface provides status of TCS(Traction Control System) status and setting.
type TractionControlSystem struct {
	// Whether or not the TCS Setting is enabled: enabled (true) or disabled (false)
	Enabled bool
	// Whether or not the TCS is engaged: engaged (true) or idle (false)
	Engaged bool
}

// The ElectronicStabilityControl interface provides status of ESC(Electronic Stability Control) status and setting.
type ElectronicStabilitySystem struct {
	// Whether or not the ESC Setting is enabled: enabled (true) or disabled (false)
	Enabled bool
	// Whether or not the ESC is engaged: engaged (true) or idle (false)
	Engaged bool
}

// The TopSpeedLimit interface provides the current setting of top speed limit of the vehicle.
type TopSpeedLimit struct {
	// Vehicle top speed limit (Unit: kilometers per hour)
	Speed uint16
}

// The AirbagStatus interface provides the current status of airbags in each zones of the vehicle.
type AirbagStatus struct {
	// Whether or not the airbag is activaged: activated (true) or deactivated (false)
	Activated bool
	// Whether the airbag is deployed: deployed (true) or not (false)
	Deployed bool
	// Zone for requested attribute
	Zone zone.Zone
}

// The Door interface provides the current status of doors in each zones of the vehicle
type Door struct {
	// The status of door's open status
	Status door_open_status.DoorOpenStatus
	// Whether or not the door is locked: locked (true) or unlocked (false)
	Lock bool
	// Zone for requested attribute
	Zone zone.Zone
}

// The ChildSafetyLock interface provides the current setting of Child Safety Lock.
type ChildSafetyLock struct {
	// Whether or not the Child Safety Lock is locked: locked (true) or unlocked (false)
	Lock bool
	// Zone for requested attribute
	Zone zone.Zone
}

// The Seat interface provides the current occupant information and seatbelt status of a seat in different
// zones of the vehicle.
type Seat struct {
	// Status of seat occupant
	Occupant occupant_status.OccupantStatus
	// Whether or not the seat belt is fastened: fastened (true) or unfastened (false)
	SeatBelt bool
	// Occupant identifier
	OccupantName string
	// Identification type
	IdentificationType identification_type.IdentificationType
	// Zone for requested attribute
	Zone zone.Zone
}

//
// Climate interfaces
//

// The Temperature interface provides information about the current temperature of outside or inside vehicle.
type Temperature struct {
	// The current temperature of the air inside of the vehicle (Unit: celsius)
	InteriorTemperature float64
	// The current temperature of the air around the vehicle (Unit: celsius)
	ExteriorTemperature float64
}

// The RainSensor interface provides information about ambient light levels.
type RailSensor struct {
	// The amount of rain detected by the rain sensor. level of rain intensity (0: No Rain, 10:Heaviest Rain)
	RainIntensity byte
	// Zone for requested attribute
	Zone zone.Zone
}

// The WiperStatus interface represents the status of wiper operation.
type WiperStatus struct {
	// Current speed interval of wiping windshield
	WiperSpeed wiper_control.WiperControl
	// Current setting of the front wiper controller. It can be used to send user's request for changing setting.
	WiperSetting wiper_control.WiperControl
}

// The Defrost interface represents the status of wiper operation.
type Defrost struct {
	// Current status of the defrost switch for window. It can be used to send user's request for changing setting
	DefrostWindow bool
	// Current status of the defrost switch for mirrors. It can be used to send user's request for changing setting.
	DefrostMirrors bool
	// Zone for requested attribute
	Zone zone.Zone
}

// The Sunroof interface represents the current status of Sunroof.
type Sunroof struct {
	// Current status of Sunroof as a percentage of openness (0%: closed, 100%: fully opened)
	Openness byte
	// Current status of Sunroof as a percentage of tilted (0%: closed, 100%: maximum tilted)
	Tilt byte
	// Zone for requested attribute
	Zone zone.Zone
}

// The ConvertibleRoof interface represents the current status of Convertible Roof.
type ConvertibleRoof struct {
	// Current status of Convertible Roof.
	Status convertible_root_status.ConvertibleRoofStatus
	// Current setting of Convertible Roof. This is used to open (true) and close (false).
	Setting bool
}

// The SideWindow interface represents the current status of openness of side windows.
type SlideWindow struct {
	// Whether or not the window is locked: locked (true) or unlocked (false)
	Lock bool
	// Current status of the side window as a percentage of openness. (0%: Closed, 100%: Fully Opened)
	Openness byte
	// Zone for requested attribute
	Zone zone.Zone
}

// The ClimateControl interface represents the current setting of the climate control equipments such as heater
// and air conditioner.
type ClimateControl struct {
	// Current status of the direction of the air flow through the ventilation system
	AirflowDirection airflow_direction.AirflowDirection
	// Current status of the fan speed of the air flowing (0: off, 1: weakest, 10: strongest )
	FanSpeedLevel byte
	// Current setting of the desired temperature (Unit: celsius)
	TargetTemperature byte
	// Current status of the air conditioning system: on (true) or off (false)
	AirConditioning bool
	// Current status of the heating system: on (true) or off (false)
	Heater bool
	// Current status of the seat warmer ( 0: off, 1: least warm, 10: warmest )
	SeatHeater byte
	// Current status of the seat ventilation ( 0: off, 1: least warm, 10: warmest )
	SeatCooler byte
	// Current setting of air recirculation: on (true) or pulling in outside air (false)
	AirRecirculation bool
	// Current status of steering wheel heater ( 0: off, 1: least warm, 10: warmest ).
	SteeringWheelHeater byte
	// Zone for requested attribute
	Zone zone.Zone
}

// The AtmosphericPressure interface provides information about the current atmospheric pressure outside of the vehicle.
type AtmosphericPressure struct {
	// The current atmospheric pressure outside of the vehicle (Unit: hectopascal)
	Pressure uint16
}

//
// Vision and Parking interfaces
//

// The LaneDepartureDetection interface represents the current status of the lane departure warning function.
type LaneDepartureDetection struct {
	// Current status of Lane departure warning function
	Status lane_departure_status.LaneDepartureStatus
}

// The Alarm interface represents the current status of the vehicle alarm system.
type Alarm struct {
	// Current status of vehicle alarm system.
	Status alarm_status.AlarmStatus
}

// The ParkingBrake interface represents the current status of the parking brake.
type ParkingBrake struct {
	// Current status of parking brake.
	Status parking_braking_status.ParkingBrakeStatus
}
