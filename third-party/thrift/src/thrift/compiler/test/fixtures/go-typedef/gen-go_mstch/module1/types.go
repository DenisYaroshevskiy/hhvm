// @generated by Thrift for [[[ program path ]]]
// This file is probably not the place you want to edit!

package module1 // [[[ program thrift source path ]]]

import (
  "fmt"

  module0 "module0"
  "github.com/facebook/fbthrift/thrift/lib/go/thrift"
)

var _ = module0.GoUnusedProtection__

// (needed to ensure safety because of naive import list construction)
var _ = fmt.Printf
var _ = thrift.ZERO


type Plate = string

type State string

type Year = int32

type Drivers = []string

type Accessory = module0.Accessory

type CarPartName = module0.PartName

type Car = Automobile

type Automobile struct {
    Plate Plate `thrift:"plate,1" json:"plate" db:"plate"`
    PreviousPlate *Plate `thrift:"previous_plate,2,optional" json:"previous_plate,omitempty" db:"previous_plate"`
    FirstPlate *Plate `thrift:"first_plate,3,optional" json:"first_plate,omitempty" db:"first_plate"`
    Year Year `thrift:"year,4" json:"year" db:"year"`
    Drivers Drivers `thrift:"drivers,5" json:"drivers" db:"drivers"`
    Accessories []*Accessory `thrift:"Accessories,6" json:"Accessories" db:"Accessories"`
    PartNames map[int32]*CarPartName `thrift:"PartNames,7" json:"PartNames" db:"PartNames"`
}
// Compile time interface enforcer
var _ thrift.Struct = &Automobile{}

func NewAutomobile() *Automobile {
    return (&Automobile{}).
        SetFirstPlate("0000")
}

// Deprecated: Use NewAutomobile().PreviousPlate instead.
var Automobile_PreviousPlate_DEFAULT = NewAutomobile().PreviousPlate

// Deprecated: Use NewAutomobile().FirstPlate instead.
var Automobile_FirstPlate_DEFAULT = NewAutomobile().FirstPlate

func (x *Automobile) GetPlate() Plate {
    return x.Plate
}

func (x *Automobile) GetPreviousPlate() *Plate {
    return x.PreviousPlate
}

func (x *Automobile) GetFirstPlate() *Plate {
    return x.FirstPlate
}

func (x *Automobile) GetYear() Year {
    return x.Year
}

func (x *Automobile) GetDrivers() Drivers {
    return x.Drivers
}

func (x *Automobile) GetAccessories() []*Accessory {
    return x.Accessories
}

func (x *Automobile) GetPartNames() map[int32]*CarPartName {
    return x.PartNames
}

func (x *Automobile) SetPlate(value Plate) *Automobile {
    x.Plate = value
    return x
}

func (x *Automobile) SetPreviousPlate(value Plate) *Automobile {
    x.PreviousPlate = &value
    return x
}

func (x *Automobile) SetFirstPlate(value Plate) *Automobile {
    x.FirstPlate = &value
    return x
}

func (x *Automobile) SetYear(value Year) *Automobile {
    x.Year = value
    return x
}

func (x *Automobile) SetDrivers(value Drivers) *Automobile {
    x.Drivers = value
    return x
}

func (x *Automobile) SetAccessories(value []*Accessory) *Automobile {
    x.Accessories = value
    return x
}

func (x *Automobile) SetPartNames(value map[int32]*CarPartName) *Automobile {
    x.PartNames = value
    return x
}


func (x *Automobile) IsSetPreviousPlate() bool {
    return x.PreviousPlate != nil
}

func (x *Automobile) IsSetFirstPlate() bool {
    return x.FirstPlate != nil
}


func (x *Automobile) IsSetDrivers() bool {
    return x.Drivers != nil
}

func (x *Automobile) IsSetAccessories() bool {
    return x.Accessories != nil
}

func (x *Automobile) IsSetPartNames() bool {
    return x.PartNames != nil
}

func (x *Automobile) writeField1(p thrift.Protocol) error {  // Plate
    if err := p.WriteFieldBegin("plate", thrift.STRING, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetPlate()
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Automobile) writeField2(p thrift.Protocol) error {  // PreviousPlate
    if !x.IsSetPreviousPlate() {
        return nil
    }

    if err := p.WriteFieldBegin("previous_plate", thrift.STRING, 2); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := *x.GetPreviousPlate()
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Automobile) writeField3(p thrift.Protocol) error {  // FirstPlate
    if !x.IsSetFirstPlate() {
        return nil
    }

    if err := p.WriteFieldBegin("first_plate", thrift.STRING, 3); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := *x.GetFirstPlate()
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Automobile) writeField4(p thrift.Protocol) error {  // Year
    if err := p.WriteFieldBegin("year", thrift.I32, 4); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetYear()
    if err := p.WriteI32(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Automobile) writeField5(p thrift.Protocol) error {  // Drivers
    if !x.IsSetDrivers() {
        return nil
    }

    if err := p.WriteFieldBegin("drivers", thrift.LIST, 5); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetDrivers()
    if err := p.WriteListBegin(thrift.STRING, len(item)); err != nil {
    return thrift.PrependError("error writing list begin: ", err)
}
for _, v := range item {
    {
        item := v
        if err := p.WriteString(item); err != nil {
    return err
}
    }
}
if err := p.WriteListEnd(); err != nil {
    return thrift.PrependError("error writing list end: ", err)
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Automobile) writeField6(p thrift.Protocol) error {  // Accessories
    if !x.IsSetAccessories() {
        return nil
    }

    if err := p.WriteFieldBegin("Accessories", thrift.LIST, 6); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetAccessories()
    if err := p.WriteListBegin(thrift.STRUCT, len(item)); err != nil {
    return thrift.PrependError("error writing list begin: ", err)
}
for _, v := range item {
    {
        item := v
        if err := item.Write(p); err != nil {
    return err
}
    }
}
if err := p.WriteListEnd(); err != nil {
    return thrift.PrependError("error writing list end: ", err)
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Automobile) writeField7(p thrift.Protocol) error {  // PartNames
    if !x.IsSetPartNames() {
        return nil
    }

    if err := p.WriteFieldBegin("PartNames", thrift.MAP, 7); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetPartNames()
    if err := p.WriteMapBegin(thrift.I32, thrift.STRUCT, len(item)); err != nil {
    return thrift.PrependError("error writing map begin: ", err)
}
for k, v := range item {
    {
        item := k
        if err := p.WriteI32(item); err != nil {
    return err
}
    }

    {
        item := v
        if err := item.Write(p); err != nil {
    return err
}
    }
}
if err := p.WriteMapEnd(); err != nil {
    return thrift.PrependError("error writing map end: ", err)
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Automobile) readField1(p thrift.Protocol) error {  // Plate
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetPlate(result)
    return nil
}

func (x *Automobile) readField2(p thrift.Protocol) error {  // PreviousPlate
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetPreviousPlate(result)
    return nil
}

func (x *Automobile) readField3(p thrift.Protocol) error {  // FirstPlate
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetFirstPlate(result)
    return nil
}

func (x *Automobile) readField4(p thrift.Protocol) error {  // Year
    result, err := p.ReadI32()
if err != nil {
    return err
}

    x.SetYear(result)
    return nil
}

func (x *Automobile) readField5(p thrift.Protocol) error {  // Drivers
    _ /* elemType */, size, err := p.ReadListBegin()
if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
}

listResult := make([]string, 0, size)
for i := 0; i < size; i++ {
    var elem string
    {
        result, err := p.ReadString()
if err != nil {
    return err
}
        elem = result
    }
    listResult = append(listResult, elem)
}

if err := p.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
}
result := listResult

    x.SetDrivers(result)
    return nil
}

func (x *Automobile) readField6(p thrift.Protocol) error {  // Accessories
    _ /* elemType */, size, err := p.ReadListBegin()
if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
}

listResult := make([]*Accessory, 0, size)
for i := 0; i < size; i++ {
    var elem Accessory
    {
        result := *module0.NewAccessory()
err := result.Read(p)
if err != nil {
    return err
}
        elem = result
    }
    listResult = append(listResult, &elem)
}

if err := p.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
}
result := listResult

    x.SetAccessories(result)
    return nil
}

func (x *Automobile) readField7(p thrift.Protocol) error {  // PartNames
    _ /* keyType */, _ /* valueType */, size, err := p.ReadMapBegin()
if err != nil {
    return thrift.PrependError("error reading map begin: ", err)
}

mapResult := make(map[int32]*CarPartName, size)
for i := 0; i < size; i++ {
    var key int32
    {
        result, err := p.ReadI32()
if err != nil {
    return err
}
        key = result
    }

    var value *CarPartName
    {
        result := *module0.NewPartName()
err := result.Read(p)
if err != nil {
    return err
}
        value = &result
    }

    mapResult[key] = value
}

if err := p.ReadMapEnd(); err != nil {
    return thrift.PrependError("error reading map end: ", err)
}
result := mapResult

    x.SetPartNames(result)
    return nil
}

func (x *Automobile) String() string {
    return fmt.Sprintf("%+v", x)
}


// Deprecated: Use Automobile.Set* methods instead or set the fields directly.
type AutomobileBuilder struct {
    obj *Automobile
}

func NewAutomobileBuilder() *AutomobileBuilder {
    return &AutomobileBuilder{
        obj: NewAutomobile(),
    }
}

func (x *AutomobileBuilder) Plate(value Plate) *AutomobileBuilder {
    x.obj.Plate = value
    return x
}

func (x *AutomobileBuilder) PreviousPlate(value *Plate) *AutomobileBuilder {
    x.obj.PreviousPlate = value
    return x
}

func (x *AutomobileBuilder) FirstPlate(value *Plate) *AutomobileBuilder {
    x.obj.FirstPlate = value
    return x
}

func (x *AutomobileBuilder) Year(value Year) *AutomobileBuilder {
    x.obj.Year = value
    return x
}

func (x *AutomobileBuilder) Drivers(value Drivers) *AutomobileBuilder {
    x.obj.Drivers = value
    return x
}

func (x *AutomobileBuilder) Accessories(value []*Accessory) *AutomobileBuilder {
    x.obj.Accessories = value
    return x
}

func (x *AutomobileBuilder) PartNames(value map[int32]*CarPartName) *AutomobileBuilder {
    x.obj.PartNames = value
    return x
}

func (x *AutomobileBuilder) Emit() *Automobile {
    var objCopy Automobile = *x.obj
    return &objCopy
}
func (x *Automobile) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("Automobile"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := x.writeField2(p); err != nil {
        return err
    }

    if err := x.writeField3(p); err != nil {
        return err
    }

    if err := x.writeField4(p); err != nil {
        return err
    }

    if err := x.writeField5(p); err != nil {
        return err
    }

    if err := x.writeField6(p); err != nil {
        return err
    }

    if err := x.writeField7(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *Automobile) Read(p thrift.Protocol) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, typ, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if typ == thrift.STOP {
            break;
        }

        switch id {
        case 1:  // plate
            if err := x.readField1(p); err != nil {
                return err
            }
        case 2:  // previous_plate
            if err := x.readField2(p); err != nil {
                return err
            }
        case 3:  // first_plate
            if err := x.readField3(p); err != nil {
                return err
            }
        case 4:  // year
            if err := x.readField4(p); err != nil {
                return err
            }
        case 5:  // drivers
            if err := x.readField5(p); err != nil {
                return err
            }
        case 6:  // Accessories
            if err := x.readField6(p); err != nil {
                return err
            }
        case 7:  // PartNames
            if err := x.readField7(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(typ); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

type MapKey struct {
    Num int64 `thrift:"num,1" json:"num" db:"num"`
    Strval string `thrift:"strval,2" json:"strval" db:"strval"`
}
// Compile time interface enforcer
var _ thrift.Struct = &MapKey{}

func NewMapKey() *MapKey {
    return (&MapKey{})
}

func (x *MapKey) GetNum() int64 {
    return x.Num
}

func (x *MapKey) GetStrval() string {
    return x.Strval
}

func (x *MapKey) SetNum(value int64) *MapKey {
    x.Num = value
    return x
}

func (x *MapKey) SetStrval(value string) *MapKey {
    x.Strval = value
    return x
}



func (x *MapKey) writeField1(p thrift.Protocol) error {  // Num
    if err := p.WriteFieldBegin("num", thrift.I64, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetNum()
    if err := p.WriteI64(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *MapKey) writeField2(p thrift.Protocol) error {  // Strval
    if err := p.WriteFieldBegin("strval", thrift.STRING, 2); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetStrval()
    if err := p.WriteString(item); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *MapKey) readField1(p thrift.Protocol) error {  // Num
    result, err := p.ReadI64()
if err != nil {
    return err
}

    x.SetNum(result)
    return nil
}

func (x *MapKey) readField2(p thrift.Protocol) error {  // Strval
    result, err := p.ReadString()
if err != nil {
    return err
}

    x.SetStrval(result)
    return nil
}

func (x *MapKey) String() string {
    return fmt.Sprintf("%+v", x)
}


// Deprecated: Use MapKey.Set* methods instead or set the fields directly.
type MapKeyBuilder struct {
    obj *MapKey
}

func NewMapKeyBuilder() *MapKeyBuilder {
    return &MapKeyBuilder{
        obj: NewMapKey(),
    }
}

func (x *MapKeyBuilder) Num(value int64) *MapKeyBuilder {
    x.obj.Num = value
    return x
}

func (x *MapKeyBuilder) Strval(value string) *MapKeyBuilder {
    x.obj.Strval = value
    return x
}

func (x *MapKeyBuilder) Emit() *MapKey {
    var objCopy MapKey = *x.obj
    return &objCopy
}
func (x *MapKey) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("MapKey"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := x.writeField2(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *MapKey) Read(p thrift.Protocol) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, typ, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if typ == thrift.STOP {
            break;
        }

        switch id {
        case 1:  // num
            if err := x.readField1(p); err != nil {
                return err
            }
        case 2:  // strval
            if err := x.readField2(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(typ); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

type MapContainer struct {
    Mapval map[*MapKey]string `thrift:"mapval,1" json:"mapval" db:"mapval"`
}
// Compile time interface enforcer
var _ thrift.Struct = &MapContainer{}

func NewMapContainer() *MapContainer {
    return (&MapContainer{})
}

func (x *MapContainer) GetMapval() map[*MapKey]string {
    return x.Mapval
}

func (x *MapContainer) SetMapval(value map[*MapKey]string) *MapContainer {
    x.Mapval = value
    return x
}

func (x *MapContainer) IsSetMapval() bool {
    return x.Mapval != nil
}

func (x *MapContainer) writeField1(p thrift.Protocol) error {  // Mapval
    if !x.IsSetMapval() {
        return nil
    }

    if err := p.WriteFieldBegin("mapval", thrift.MAP, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetMapval()
    if err := p.WriteMapBegin(thrift.STRUCT, thrift.STRING, len(item)); err != nil {
    return thrift.PrependError("error writing map begin: ", err)
}
for k, v := range item {
    {
        item := k
        if err := item.Write(p); err != nil {
    return err
}
    }

    {
        item := v
        if err := p.WriteString(item); err != nil {
    return err
}
    }
}
if err := p.WriteMapEnd(); err != nil {
    return thrift.PrependError("error writing map end: ", err)
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *MapContainer) readField1(p thrift.Protocol) error {  // Mapval
    _ /* keyType */, _ /* valueType */, size, err := p.ReadMapBegin()
if err != nil {
    return thrift.PrependError("error reading map begin: ", err)
}

mapResult := make(map[*MapKey]string, size)
for i := 0; i < size; i++ {
    var key *MapKey
    {
        result := *NewMapKey()
err := result.Read(p)
if err != nil {
    return err
}
        key = &result
    }

    var value string
    {
        result, err := p.ReadString()
if err != nil {
    return err
}
        value = result
    }

    mapResult[key] = value
}

if err := p.ReadMapEnd(); err != nil {
    return thrift.PrependError("error reading map end: ", err)
}
result := mapResult

    x.SetMapval(result)
    return nil
}

func (x *MapContainer) String() string {
    return fmt.Sprintf("%+v", x)
}


// Deprecated: Use MapContainer.Set* methods instead or set the fields directly.
type MapContainerBuilder struct {
    obj *MapContainer
}

func NewMapContainerBuilder() *MapContainerBuilder {
    return &MapContainerBuilder{
        obj: NewMapContainer(),
    }
}

func (x *MapContainerBuilder) Mapval(value map[*MapKey]string) *MapContainerBuilder {
    x.obj.Mapval = value
    return x
}

func (x *MapContainerBuilder) Emit() *MapContainer {
    var objCopy MapContainer = *x.obj
    return &objCopy
}
func (x *MapContainer) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("MapContainer"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *MapContainer) Read(p thrift.Protocol) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, typ, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if typ == thrift.STOP {
            break;
        }

        switch id {
        case 1:  // mapval
            if err := x.readField1(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(typ); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

type Pair struct {
    Automobile *Automobile `thrift:"automobile,1" json:"automobile" db:"automobile"`
    Car *Car `thrift:"car,2" json:"car" db:"car"`
}
// Compile time interface enforcer
var _ thrift.Struct = &Pair{}

func NewPair() *Pair {
    return (&Pair{})
}

// Deprecated: Use NewPair().Automobile instead.
var Pair_Automobile_DEFAULT = NewPair().Automobile

// Deprecated: Use NewPair().Car instead.
var Pair_Car_DEFAULT = NewPair().Car

func (x *Pair) GetAutomobile() *Automobile {
    return x.Automobile
}

func (x *Pair) GetCar() *Car {
    return x.Car
}

func (x *Pair) SetAutomobile(value Automobile) *Pair {
    x.Automobile = &value
    return x
}

func (x *Pair) SetCar(value Car) *Pair {
    x.Car = &value
    return x
}

func (x *Pair) IsSetAutomobile() bool {
    return x.Automobile != nil
}

func (x *Pair) IsSetCar() bool {
    return x.Car != nil
}

func (x *Pair) writeField1(p thrift.Protocol) error {  // Automobile
    if !x.IsSetAutomobile() {
        return nil
    }

    if err := p.WriteFieldBegin("automobile", thrift.STRUCT, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetAutomobile()
    if err := item.Write(p); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Pair) writeField2(p thrift.Protocol) error {  // Car
    if !x.IsSetCar() {
        return nil
    }

    if err := p.WriteFieldBegin("car", thrift.STRUCT, 2); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetCar()
    if err := item.Write(p); err != nil {
    return err
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Pair) readField1(p thrift.Protocol) error {  // Automobile
    result := *NewAutomobile()
err := result.Read(p)
if err != nil {
    return err
}

    x.SetAutomobile(result)
    return nil
}

func (x *Pair) readField2(p thrift.Protocol) error {  // Car
    result := *NewAutomobile()
err := result.Read(p)
if err != nil {
    return err
}

    x.SetCar(result)
    return nil
}

func (x *Pair) String() string {
    return fmt.Sprintf("%+v", x)
}


// Deprecated: Use Pair.Set* methods instead or set the fields directly.
type PairBuilder struct {
    obj *Pair
}

func NewPairBuilder() *PairBuilder {
    return &PairBuilder{
        obj: NewPair(),
    }
}

func (x *PairBuilder) Automobile(value *Automobile) *PairBuilder {
    x.obj.Automobile = value
    return x
}

func (x *PairBuilder) Car(value *Car) *PairBuilder {
    x.obj.Car = value
    return x
}

func (x *PairBuilder) Emit() *Pair {
    var objCopy Pair = *x.obj
    return &objCopy
}
func (x *Pair) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("Pair"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := x.writeField2(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *Pair) Read(p thrift.Protocol) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, typ, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if typ == thrift.STOP {
            break;
        }

        switch id {
        case 1:  // automobile
            if err := x.readField1(p); err != nil {
                return err
            }
        case 2:  // car
            if err := x.readField2(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(typ); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}

type Collection struct {
    Automobiles []*Automobile `thrift:"automobiles,1" json:"automobiles" db:"automobiles"`
    Cars []*Car `thrift:"cars,2" json:"cars" db:"cars"`
}
// Compile time interface enforcer
var _ thrift.Struct = &Collection{}

func NewCollection() *Collection {
    return (&Collection{})
}

func (x *Collection) GetAutomobiles() []*Automobile {
    return x.Automobiles
}

func (x *Collection) GetCars() []*Car {
    return x.Cars
}

func (x *Collection) SetAutomobiles(value []*Automobile) *Collection {
    x.Automobiles = value
    return x
}

func (x *Collection) SetCars(value []*Car) *Collection {
    x.Cars = value
    return x
}

func (x *Collection) IsSetAutomobiles() bool {
    return x.Automobiles != nil
}

func (x *Collection) IsSetCars() bool {
    return x.Cars != nil
}

func (x *Collection) writeField1(p thrift.Protocol) error {  // Automobiles
    if !x.IsSetAutomobiles() {
        return nil
    }

    if err := p.WriteFieldBegin("automobiles", thrift.LIST, 1); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetAutomobiles()
    if err := p.WriteListBegin(thrift.STRUCT, len(item)); err != nil {
    return thrift.PrependError("error writing list begin: ", err)
}
for _, v := range item {
    {
        item := v
        if err := item.Write(p); err != nil {
    return err
}
    }
}
if err := p.WriteListEnd(); err != nil {
    return thrift.PrependError("error writing list end: ", err)
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Collection) writeField2(p thrift.Protocol) error {  // Cars
    if !x.IsSetCars() {
        return nil
    }

    if err := p.WriteFieldBegin("cars", thrift.LIST, 2); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field begin error: ", x), err)
    }

    item := x.GetCars()
    if err := p.WriteListBegin(thrift.STRUCT, len(item)); err != nil {
    return thrift.PrependError("error writing list begin: ", err)
}
for _, v := range item {
    {
        item := v
        if err := item.Write(p); err != nil {
    return err
}
    }
}
if err := p.WriteListEnd(); err != nil {
    return thrift.PrependError("error writing list end: ", err)
}

    if err := p.WriteFieldEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field end error: ", x), err)
    }
    return nil
}

func (x *Collection) readField1(p thrift.Protocol) error {  // Automobiles
    _ /* elemType */, size, err := p.ReadListBegin()
if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
}

listResult := make([]*Automobile, 0, size)
for i := 0; i < size; i++ {
    var elem Automobile
    {
        result := *NewAutomobile()
err := result.Read(p)
if err != nil {
    return err
}
        elem = result
    }
    listResult = append(listResult, &elem)
}

if err := p.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
}
result := listResult

    x.SetAutomobiles(result)
    return nil
}

func (x *Collection) readField2(p thrift.Protocol) error {  // Cars
    _ /* elemType */, size, err := p.ReadListBegin()
if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
}

listResult := make([]*Car, 0, size)
for i := 0; i < size; i++ {
    var elem Car
    {
        result := *NewAutomobile()
err := result.Read(p)
if err != nil {
    return err
}
        elem = result
    }
    listResult = append(listResult, &elem)
}

if err := p.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
}
result := listResult

    x.SetCars(result)
    return nil
}

func (x *Collection) String() string {
    return fmt.Sprintf("%+v", x)
}


// Deprecated: Use Collection.Set* methods instead or set the fields directly.
type CollectionBuilder struct {
    obj *Collection
}

func NewCollectionBuilder() *CollectionBuilder {
    return &CollectionBuilder{
        obj: NewCollection(),
    }
}

func (x *CollectionBuilder) Automobiles(value []*Automobile) *CollectionBuilder {
    x.obj.Automobiles = value
    return x
}

func (x *CollectionBuilder) Cars(value []*Car) *CollectionBuilder {
    x.obj.Cars = value
    return x
}

func (x *CollectionBuilder) Emit() *Collection {
    var objCopy Collection = *x.obj
    return &objCopy
}
func (x *Collection) Write(p thrift.Protocol) error {
    if err := p.WriteStructBegin("Collection"); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", x), err)
    }

    if err := x.writeField1(p); err != nil {
        return err
    }

    if err := x.writeField2(p); err != nil {
        return err
    }

    if err := p.WriteFieldStop(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", x), err)
    }

    if err := p.WriteStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", x), err)
    }
    return nil
}

func (x *Collection) Read(p thrift.Protocol) error {
    if _, err := p.ReadStructBegin(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read error: ", x), err)
    }

    for {
        _, typ, id, err := p.ReadFieldBegin()
        if err != nil {
            return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", x, id), err)
        }

        if typ == thrift.STOP {
            break;
        }

        switch id {
        case 1:  // automobiles
            if err := x.readField1(p); err != nil {
                return err
            }
        case 2:  // cars
            if err := x.readField2(p); err != nil {
                return err
            }
        default:
            if err := p.Skip(typ); err != nil {
                return err
            }
        }

        if err := p.ReadFieldEnd(); err != nil {
            return err
        }
    }

    if err := p.ReadStructEnd(); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", x), err)
    }

    return nil
}
