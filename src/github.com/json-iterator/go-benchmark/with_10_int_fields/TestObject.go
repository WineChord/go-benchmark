// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package with_10_int_fields

// import (
// 	"bytes"
// 	"fmt"
// 	"git.apache.org/thrift.git/lib/go/thrift"
// )

// // (needed to ensure safety because of naive import list construction.)
// var _ = thrift.ZERO
// var _ = fmt.Printf
// var _ = bytes.Equal

// // Attributes:
// //  - Field1
// //  - Field2
// //  - Field3
// //  - Field4
// //  - Field5
// //  - Field6
// //  - Field7
// //  - Field8
// //  - Field9
// //  - Field10
// type ThriftTestObject struct {
//   Field1 int32 `thrift:"field1,1" db:"field1" json:"field1"`
//   Field2 int32 `thrift:"field2,2" db:"field2" json:"field2"`
//   Field3 int32 `thrift:"field3,3" db:"field3" json:"field3"`
//   Field4 int32 `thrift:"field4,4" db:"field4" json:"field4"`
//   Field5 int32 `thrift:"field5,5" db:"field5" json:"field5"`
//   Field6 int32 `thrift:"field6,6" db:"field6" json:"field6"`
//   Field7 int32 `thrift:"field7,7" db:"field7" json:"field7"`
//   Field8 int32 `thrift:"field8,8" db:"field8" json:"field8"`
//   Field9 int32 `thrift:"field9,9" db:"field9" json:"field9"`
//   Field10 int32 `thrift:"field10,10" db:"field10" json:"field10"`
// }

// func NewThriftTestObject() *ThriftTestObject {
//   return &ThriftTestObject{}
// }


// func (p *ThriftTestObject) GetField1() int32 {
//   return p.Field1
// }

// func (p *ThriftTestObject) GetField2() int32 {
//   return p.Field2
// }

// func (p *ThriftTestObject) GetField3() int32 {
//   return p.Field3
// }

// func (p *ThriftTestObject) GetField4() int32 {
//   return p.Field4
// }

// func (p *ThriftTestObject) GetField5() int32 {
//   return p.Field5
// }

// func (p *ThriftTestObject) GetField6() int32 {
//   return p.Field6
// }

// func (p *ThriftTestObject) GetField7() int32 {
//   return p.Field7
// }

// func (p *ThriftTestObject) GetField8() int32 {
//   return p.Field8
// }

// func (p *ThriftTestObject) GetField9() int32 {
//   return p.Field9
// }

// func (p *ThriftTestObject) GetField10() int32 {
//   return p.Field10
// }
// func (p *ThriftTestObject) Read(iprot thrift.TProtocol) error {
//   if _, err := iprot.ReadStructBegin(); err != nil {
//     return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
//   }


//   for {
//     _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
//     if err != nil {
//       return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
//     }
//     if fieldTypeId == thrift.STOP { break; }
//     switch fieldId {
//     case 1:
//       if err := p.ReadField1(iprot); err != nil {
//         return err
//       }
//     case 2:
//       if err := p.ReadField2(iprot); err != nil {
//         return err
//       }
//     case 3:
//       if err := p.ReadField3(iprot); err != nil {
//         return err
//       }
//     case 4:
//       if err := p.ReadField4(iprot); err != nil {
//         return err
//       }
//     case 5:
//       if err := p.ReadField5(iprot); err != nil {
//         return err
//       }
//     case 6:
//       if err := p.ReadField6(iprot); err != nil {
//         return err
//       }
//     case 7:
//       if err := p.ReadField7(iprot); err != nil {
//         return err
//       }
//     case 8:
//       if err := p.ReadField8(iprot); err != nil {
//         return err
//       }
//     case 9:
//       if err := p.ReadField9(iprot); err != nil {
//         return err
//       }
//     case 10:
//       if err := p.ReadField10(iprot); err != nil {
//         return err
//       }
//     default:
//       if err := iprot.Skip(fieldTypeId); err != nil {
//         return err
//       }
//     }
//     if err := iprot.ReadFieldEnd(); err != nil {
//       return err
//     }
//   }
//   if err := iprot.ReadStructEnd(); err != nil {
//     return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
//   }
//   return nil
// }

// func (p *ThriftTestObject)  ReadField1(iprot thrift.TProtocol) error {
//   if v, err := iprot.ReadI32(); err != nil {
//   return thrift.PrependError("error reading field 1: ", err)
// } else {
//   p.Field1 = v
// }
//   return nil
// }

// func (p *ThriftTestObject)  ReadField2(iprot thrift.TProtocol) error {
//   if v, err := iprot.ReadI32(); err != nil {
//   return thrift.PrependError("error reading field 2: ", err)
// } else {
//   p.Field2 = v
// }
//   return nil
// }

// func (p *ThriftTestObject)  ReadField3(iprot thrift.TProtocol) error {
//   if v, err := iprot.ReadI32(); err != nil {
//   return thrift.PrependError("error reading field 3: ", err)
// } else {
//   p.Field3 = v
// }
//   return nil
// }

// func (p *ThriftTestObject)  ReadField4(iprot thrift.TProtocol) error {
//   if v, err := iprot.ReadI32(); err != nil {
//   return thrift.PrependError("error reading field 4: ", err)
// } else {
//   p.Field4 = v
// }
//   return nil
// }

// func (p *ThriftTestObject)  ReadField5(iprot thrift.TProtocol) error {
//   if v, err := iprot.ReadI32(); err != nil {
//   return thrift.PrependError("error reading field 5: ", err)
// } else {
//   p.Field5 = v
// }
//   return nil
// }

// func (p *ThriftTestObject)  ReadField6(iprot thrift.TProtocol) error {
//   if v, err := iprot.ReadI32(); err != nil {
//   return thrift.PrependError("error reading field 6: ", err)
// } else {
//   p.Field6 = v
// }
//   return nil
// }

// func (p *ThriftTestObject)  ReadField7(iprot thrift.TProtocol) error {
//   if v, err := iprot.ReadI32(); err != nil {
//   return thrift.PrependError("error reading field 7: ", err)
// } else {
//   p.Field7 = v
// }
//   return nil
// }

// func (p *ThriftTestObject)  ReadField8(iprot thrift.TProtocol) error {
//   if v, err := iprot.ReadI32(); err != nil {
//   return thrift.PrependError("error reading field 8: ", err)
// } else {
//   p.Field8 = v
// }
//   return nil
// }

// func (p *ThriftTestObject)  ReadField9(iprot thrift.TProtocol) error {
//   if v, err := iprot.ReadI32(); err != nil {
//   return thrift.PrependError("error reading field 9: ", err)
// } else {
//   p.Field9 = v
// }
//   return nil
// }

// func (p *ThriftTestObject)  ReadField10(iprot thrift.TProtocol) error {
//   if v, err := iprot.ReadI32(); err != nil {
//   return thrift.PrependError("error reading field 10: ", err)
// } else {
//   p.Field10 = v
// }
//   return nil
// }

// func (p *ThriftTestObject) Write(oprot thrift.TProtocol) error {
//   if err := oprot.WriteStructBegin("ThriftTestObject"); err != nil {
//     return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
//   if p != nil {
//     if err := p.writeField1(oprot); err != nil { return err }
//     if err := p.writeField2(oprot); err != nil { return err }
//     if err := p.writeField3(oprot); err != nil { return err }
//     if err := p.writeField4(oprot); err != nil { return err }
//     if err := p.writeField5(oprot); err != nil { return err }
//     if err := p.writeField6(oprot); err != nil { return err }
//     if err := p.writeField7(oprot); err != nil { return err }
//     if err := p.writeField8(oprot); err != nil { return err }
//     if err := p.writeField9(oprot); err != nil { return err }
//     if err := p.writeField10(oprot); err != nil { return err }
//   }
//   if err := oprot.WriteFieldStop(); err != nil {
//     return thrift.PrependError("write field stop error: ", err) }
//   if err := oprot.WriteStructEnd(); err != nil {
//     return thrift.PrependError("write struct stop error: ", err) }
//   return nil
// }

// func (p *ThriftTestObject) writeField1(oprot thrift.TProtocol) (err error) {
//   if err := oprot.WriteFieldBegin("field1", thrift.I32, 1); err != nil {
//     return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:field1: ", p), err) }
//   if err := oprot.WriteI32(int32(p.Field1)); err != nil {
//   return thrift.PrependError(fmt.Sprintf("%T.field1 (1) field write error: ", p), err) }
//   if err := oprot.WriteFieldEnd(); err != nil {
//     return thrift.PrependError(fmt.Sprintf("%T write field end error 1:field1: ", p), err) }
//   return err
// }

// func (p *ThriftTestObject) writeField2(oprot thrift.TProtocol) (err error) {
//   if err := oprot.WriteFieldBegin("field2", thrift.I32, 2); err != nil {
//     return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:field2: ", p), err) }
//   if err := oprot.WriteI32(int32(p.Field2)); err != nil {
//   return thrift.PrependError(fmt.Sprintf("%T.field2 (2) field write error: ", p), err) }
//   if err := oprot.WriteFieldEnd(); err != nil {
//     return thrift.PrependError(fmt.Sprintf("%T write field end error 2:field2: ", p), err) }
//   return err
// }

// func (p *ThriftTestObject) writeField3(oprot thrift.TProtocol) (err error) {
//   if err := oprot.WriteFieldBegin("field3", thrift.I32, 3); err != nil {
//     return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:field3: ", p), err) }
//   if err := oprot.WriteI32(int32(p.Field3)); err != nil {
//   return thrift.PrependError(fmt.Sprintf("%T.field3 (3) field write error: ", p), err) }
//   if err := oprot.WriteFieldEnd(); err != nil {
//     return thrift.PrependError(fmt.Sprintf("%T write field end error 3:field3: ", p), err) }
//   return err
// }

// func (p *ThriftTestObject) writeField4(oprot thrift.TProtocol) (err error) {
//   if err := oprot.WriteFieldBegin("field4", thrift.I32, 4); err != nil {
//     return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:field4: ", p), err) }
//   if err := oprot.WriteI32(int32(p.Field4)); err != nil {
//   return thrift.PrependError(fmt.Sprintf("%T.field4 (4) field write error: ", p), err) }
//   if err := oprot.WriteFieldEnd(); err != nil {
//     return thrift.PrependError(fmt.Sprintf("%T write field end error 4:field4: ", p), err) }
//   return err
// }

// func (p *ThriftTestObject) writeField5(oprot thrift.TProtocol) (err error) {
//   if err := oprot.WriteFieldBegin("field5", thrift.I32, 5); err != nil {
//     return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:field5: ", p), err) }
//   if err := oprot.WriteI32(int32(p.Field5)); err != nil {
//   return thrift.PrependError(fmt.Sprintf("%T.field5 (5) field write error: ", p), err) }
//   if err := oprot.WriteFieldEnd(); err != nil {
//     return thrift.PrependError(fmt.Sprintf("%T write field end error 5:field5: ", p), err) }
//   return err
// }

// func (p *ThriftTestObject) writeField6(oprot thrift.TProtocol) (err error) {
//   if err := oprot.WriteFieldBegin("field6", thrift.I32, 6); err != nil {
//     return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:field6: ", p), err) }
//   if err := oprot.WriteI32(int32(p.Field6)); err != nil {
//   return thrift.PrependError(fmt.Sprintf("%T.field6 (6) field write error: ", p), err) }
//   if err := oprot.WriteFieldEnd(); err != nil {
//     return thrift.PrependError(fmt.Sprintf("%T write field end error 6:field6: ", p), err) }
//   return err
// }

// func (p *ThriftTestObject) writeField7(oprot thrift.TProtocol) (err error) {
//   if err := oprot.WriteFieldBegin("field7", thrift.I32, 7); err != nil {
//     return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:field7: ", p), err) }
//   if err := oprot.WriteI32(int32(p.Field7)); err != nil {
//   return thrift.PrependError(fmt.Sprintf("%T.field7 (7) field write error: ", p), err) }
//   if err := oprot.WriteFieldEnd(); err != nil {
//     return thrift.PrependError(fmt.Sprintf("%T write field end error 7:field7: ", p), err) }
//   return err
// }

// func (p *ThriftTestObject) writeField8(oprot thrift.TProtocol) (err error) {
//   if err := oprot.WriteFieldBegin("field8", thrift.I32, 8); err != nil {
//     return thrift.PrependError(fmt.Sprintf("%T write field begin error 8:field8: ", p), err) }
//   if err := oprot.WriteI32(int32(p.Field8)); err != nil {
//   return thrift.PrependError(fmt.Sprintf("%T.field8 (8) field write error: ", p), err) }
//   if err := oprot.WriteFieldEnd(); err != nil {
//     return thrift.PrependError(fmt.Sprintf("%T write field end error 8:field8: ", p), err) }
//   return err
// }

// func (p *ThriftTestObject) writeField9(oprot thrift.TProtocol) (err error) {
//   if err := oprot.WriteFieldBegin("field9", thrift.I32, 9); err != nil {
//     return thrift.PrependError(fmt.Sprintf("%T write field begin error 9:field9: ", p), err) }
//   if err := oprot.WriteI32(int32(p.Field9)); err != nil {
//   return thrift.PrependError(fmt.Sprintf("%T.field9 (9) field write error: ", p), err) }
//   if err := oprot.WriteFieldEnd(); err != nil {
//     return thrift.PrependError(fmt.Sprintf("%T write field end error 9:field9: ", p), err) }
//   return err
// }

// func (p *ThriftTestObject) writeField10(oprot thrift.TProtocol) (err error) {
//   if err := oprot.WriteFieldBegin("field10", thrift.I32, 10); err != nil {
//     return thrift.PrependError(fmt.Sprintf("%T write field begin error 10:field10: ", p), err) }
//   if err := oprot.WriteI32(int32(p.Field10)); err != nil {
//   return thrift.PrependError(fmt.Sprintf("%T.field10 (10) field write error: ", p), err) }
//   if err := oprot.WriteFieldEnd(); err != nil {
//     return thrift.PrependError(fmt.Sprintf("%T write field end error 10:field10: ", p), err) }
//   return err
// }

// func (p *ThriftTestObject) String() string {
//   if p == nil {
//     return "<nil>"
//   }
//   return fmt.Sprintf("ThriftTestObject(%+v)", *p)
// }

